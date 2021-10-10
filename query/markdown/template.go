package markdown

import (
	"encoding/json"
	"path"
	"text/template"
	"time"

	"github.com/flexzuu/website/query/client"
)

// Define a template.
const post = `
{{. | frontmatter }}

{{.Content}}
`
const autoImage = "auto_image"
const resize = "resize=width:400,fit:max"
const compress = "compress"

func assetUrl(handle string) string {
	return path.Join("https://cdn.filestackcontent.com", autoImage, resize, compress, handle)
}

type postFrontmatter struct {
	Title  string `json:"title,omitempty"`
	Date   string `json:"date,omitempty"`
	Author string `json:"author,omitempty"`

	Cover            string `json:"cover,omitempty"`
	CoverAttribution string `json:"cover_attribution,omitempty"`

	Description string `json:"description,omitempty"`
}

func postFrontmatterRender(p client.PostsPostsPost) (string, error) {
	bs, err := json.MarshalIndent(postFrontmatter{
		Title:  p.Title,
		Date:   p.UpdatedAt.Format(time.RFC3339),
		Author: p.UpdatedBy.Name,

		Cover:            assetUrl(p.CoverImage.Handle),
		CoverAttribution: p.CoverImage.Attribution,

		Description: p.Description,
	}, "", "\t")

	if err != nil {
		return "", err
	}
	return string(bs), nil
}

/*
+++
title = "Hello Friend"
date = "2021-10-09 22:44:26.124072 +0000 +0000"
author = "Jonas Faber"
cover = "https:/cdn.filestackcontent.com/auto_image/resize=width:400,fit:max/compress/V2jDA0ZyRg6TZMn6yIEf"
coverAttribution = "Photo by [Drew Beamer](https://unsplash.com/@drew_beamer?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) on [Unsplash](https://unsplash.com/s/photos/hello?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)
  "
description = "A simple description"
+++
*/

// Create a new template and parse the letter into it.
var PostTemplate = template.Must(template.New("post").Funcs(template.FuncMap{
	"frontmatter": postFrontmatterRender,
}).Parse(post))

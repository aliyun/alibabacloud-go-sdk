// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageURL interface {
	dara.Model
	String() string
	GoString() string
	SetThumbnail(v string) *ImageURL
	GetThumbnail() *string
	SetURL(v string) *ImageURL
	GetURL() *string
}

type ImageURL struct {
	Thumbnail *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	URL       *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s ImageURL) String() string {
	return dara.Prettify(s)
}

func (s ImageURL) GoString() string {
	return s.String()
}

func (s *ImageURL) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *ImageURL) GetURL() *string {
	return s.URL
}

func (s *ImageURL) SetThumbnail(v string) *ImageURL {
	s.Thumbnail = &v
	return s
}

func (s *ImageURL) SetURL(v string) *ImageURL {
	s.URL = &v
	return s
}

func (s *ImageURL) Validate() error {
	return dara.Validate(s)
}

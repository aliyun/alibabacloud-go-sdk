// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnifiedImageItem interface {
	dara.Model
	String() string
	GoString() string
	SetHeight(v int32) *UnifiedImageItem
	GetHeight() *int32
	SetHostPageUrl(v string) *UnifiedImageItem
	GetHostPageUrl() *string
	SetImageUrl(v string) *UnifiedImageItem
	GetImageUrl() *string
	SetPublishedTime(v string) *UnifiedImageItem
	GetPublishedTime() *string
	SetTitle(v string) *UnifiedImageItem
	GetTitle() *string
	SetWidth(v int32) *UnifiedImageItem
	GetWidth() *int32
}

type UnifiedImageItem struct {
	// Height
	//
	// example:
	//
	// 1330
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// URL of the web page to which the image belongs
	//
	// example:
	//
	// http://mbd.baidu.com/newspage/data/dtlandingsuper?nid=dt_4541580238898912926
	HostPageUrl *string `json:"hostPageUrl,omitempty" xml:"hostPageUrl,omitempty"`
	// Image URL
	//
	// example:
	//
	// http://pic.rmb.bdstatic.com/bjh/bb87f566c0c/241218/f7936f25837b20211e5ef88d7980c143.jpeg
	ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	// Web page published time, in ISO time format
	//
	// example:
	//
	// 2022-07-05T00:54:42+08:00
	PublishedTime *string `json:"publishedTime,omitempty" xml:"publishedTime,omitempty"`
	// Image title
	//
	// example:
	//
	// 阿里巴巴现在以主业上为主,以电商为主,以能够盈利的业务为主
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Image width
	//
	// example:
	//
	// 1000
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s UnifiedImageItem) String() string {
	return dara.Prettify(s)
}

func (s UnifiedImageItem) GoString() string {
	return s.String()
}

func (s *UnifiedImageItem) GetHeight() *int32 {
	return s.Height
}

func (s *UnifiedImageItem) GetHostPageUrl() *string {
	return s.HostPageUrl
}

func (s *UnifiedImageItem) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *UnifiedImageItem) GetPublishedTime() *string {
	return s.PublishedTime
}

func (s *UnifiedImageItem) GetTitle() *string {
	return s.Title
}

func (s *UnifiedImageItem) GetWidth() *int32 {
	return s.Width
}

func (s *UnifiedImageItem) SetHeight(v int32) *UnifiedImageItem {
	s.Height = &v
	return s
}

func (s *UnifiedImageItem) SetHostPageUrl(v string) *UnifiedImageItem {
	s.HostPageUrl = &v
	return s
}

func (s *UnifiedImageItem) SetImageUrl(v string) *UnifiedImageItem {
	s.ImageUrl = &v
	return s
}

func (s *UnifiedImageItem) SetPublishedTime(v string) *UnifiedImageItem {
	s.PublishedTime = &v
	return s
}

func (s *UnifiedImageItem) SetTitle(v string) *UnifiedImageItem {
	s.Title = &v
	return s
}

func (s *UnifiedImageItem) SetWidth(v int32) *UnifiedImageItem {
	s.Width = &v
	return s
}

func (s *UnifiedImageItem) Validate() error {
	return dara.Validate(s)
}

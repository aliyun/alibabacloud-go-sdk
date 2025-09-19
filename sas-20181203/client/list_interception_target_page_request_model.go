// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterceptionTargetPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListInterceptionTargetPageRequest
	GetAppName() *string
	SetCurrentPage(v int32) *ListInterceptionTargetPageRequest
	GetCurrentPage() *int32
	SetImageList(v []*string) *ListInterceptionTargetPageRequest
	GetImageList() []*string
	SetNamespace(v string) *ListInterceptionTargetPageRequest
	GetNamespace() *string
	SetPageSize(v int32) *ListInterceptionTargetPageRequest
	GetPageSize() *int32
	SetTagList(v []*string) *ListInterceptionTargetPageRequest
	GetTagList() []*string
	SetTargetName(v string) *ListInterceptionTargetPageRequest
	GetTargetName() *string
	SetTargetType(v string) *ListInterceptionTargetPageRequest
	GetTargetType() *string
}

type ListInterceptionTargetPageRequest struct {
	// The name of the application to which the network object belongs.
	//
	// example:
	//
	// frontend
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The images of the network object.
	ImageList []*string `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	// The namespace to which the network object belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The labels specified for the network object.
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The name of the network object.
	//
	// example:
	//
	// source-test-obj-0****
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the network object. Valid values:
	//
	// 	- IMAGE
	//
	// example:
	//
	// IMAGE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListInterceptionTargetPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInterceptionTargetPageRequest) GoString() string {
	return s.String()
}

func (s *ListInterceptionTargetPageRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListInterceptionTargetPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListInterceptionTargetPageRequest) GetImageList() []*string {
	return s.ImageList
}

func (s *ListInterceptionTargetPageRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListInterceptionTargetPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInterceptionTargetPageRequest) GetTagList() []*string {
	return s.TagList
}

func (s *ListInterceptionTargetPageRequest) GetTargetName() *string {
	return s.TargetName
}

func (s *ListInterceptionTargetPageRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListInterceptionTargetPageRequest) SetAppName(v string) *ListInterceptionTargetPageRequest {
	s.AppName = &v
	return s
}

func (s *ListInterceptionTargetPageRequest) SetCurrentPage(v int32) *ListInterceptionTargetPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListInterceptionTargetPageRequest) SetImageList(v []*string) *ListInterceptionTargetPageRequest {
	s.ImageList = v
	return s
}

func (s *ListInterceptionTargetPageRequest) SetNamespace(v string) *ListInterceptionTargetPageRequest {
	s.Namespace = &v
	return s
}

func (s *ListInterceptionTargetPageRequest) SetPageSize(v int32) *ListInterceptionTargetPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListInterceptionTargetPageRequest) SetTagList(v []*string) *ListInterceptionTargetPageRequest {
	s.TagList = v
	return s
}

func (s *ListInterceptionTargetPageRequest) SetTargetName(v string) *ListInterceptionTargetPageRequest {
	s.TargetName = &v
	return s
}

func (s *ListInterceptionTargetPageRequest) SetTargetType(v string) *ListInterceptionTargetPageRequest {
	s.TargetType = &v
	return s
}

func (s *ListInterceptionTargetPageRequest) Validate() error {
	return dara.Validate(s)
}

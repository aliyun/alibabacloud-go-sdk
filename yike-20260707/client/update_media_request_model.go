// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppendTags(v bool) *UpdateMediaRequest
	GetAppendTags() *bool
	SetCategoryId(v int64) *UpdateMediaRequest
	GetCategoryId() *int64
	SetCoverURL(v string) *UpdateMediaRequest
	GetCoverURL() *string
	SetDescription(v string) *UpdateMediaRequest
	GetDescription() *string
	SetDynamicMetaData(v string) *UpdateMediaRequest
	GetDynamicMetaData() *string
	SetInputURL(v string) *UpdateMediaRequest
	GetInputURL() *string
	SetMediaId(v string) *UpdateMediaRequest
	GetMediaId() *string
	SetMediaTags(v string) *UpdateMediaRequest
	GetMediaTags() *string
	SetTitle(v string) *UpdateMediaRequest
	GetTitle() *string
	SetUserData(v string) *UpdateMediaRequest
	GetUserData() *string
}

type UpdateMediaRequest struct {
	AppendTags *bool  `json:"AppendTags,omitempty" xml:"AppendTags,omitempty"`
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// https://bullet-time-20240910.oss-cn-shanghai.aliyuncs.com/ice-generated/a97255309a7c71f093d3e7f6d75a6302/snapshots/normal/2a4030b1950443048f3e3f81489d57eb-00001.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {}
	DynamicMetaData *string `json:"DynamicMetaData,omitempty" xml:"DynamicMetaData,omitempty"`
	// example:
	//
	// https://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// 剪映动画
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// {}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s UpdateMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaRequest) GetAppendTags() *bool {
	return s.AppendTags
}

func (s *UpdateMediaRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *UpdateMediaRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *UpdateMediaRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMediaRequest) GetDynamicMetaData() *string {
	return s.DynamicMetaData
}

func (s *UpdateMediaRequest) GetInputURL() *string {
	return s.InputURL
}

func (s *UpdateMediaRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *UpdateMediaRequest) GetMediaTags() *string {
	return s.MediaTags
}

func (s *UpdateMediaRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateMediaRequest) GetUserData() *string {
	return s.UserData
}

func (s *UpdateMediaRequest) SetAppendTags(v bool) *UpdateMediaRequest {
	s.AppendTags = &v
	return s
}

func (s *UpdateMediaRequest) SetCategoryId(v int64) *UpdateMediaRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateMediaRequest) SetCoverURL(v string) *UpdateMediaRequest {
	s.CoverURL = &v
	return s
}

func (s *UpdateMediaRequest) SetDescription(v string) *UpdateMediaRequest {
	s.Description = &v
	return s
}

func (s *UpdateMediaRequest) SetDynamicMetaData(v string) *UpdateMediaRequest {
	s.DynamicMetaData = &v
	return s
}

func (s *UpdateMediaRequest) SetInputURL(v string) *UpdateMediaRequest {
	s.InputURL = &v
	return s
}

func (s *UpdateMediaRequest) SetMediaId(v string) *UpdateMediaRequest {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaRequest) SetMediaTags(v string) *UpdateMediaRequest {
	s.MediaTags = &v
	return s
}

func (s *UpdateMediaRequest) SetTitle(v string) *UpdateMediaRequest {
	s.Title = &v
	return s
}

func (s *UpdateMediaRequest) SetUserData(v string) *UpdateMediaRequest {
	s.UserData = &v
	return s
}

func (s *UpdateMediaRequest) Validate() error {
	return dara.Validate(s)
}

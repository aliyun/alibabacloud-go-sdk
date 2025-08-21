// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *RemoveTagRequest
	GetAppKey() *int64
	SetTagName(v string) *RemoveTagRequest
	GetTagName() *string
}

type RemoveTagRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_tag
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s RemoveTagRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagRequest) GoString() string {
	return s.String()
}

func (s *RemoveTagRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *RemoveTagRequest) GetTagName() *string {
	return s.TagName
}

func (s *RemoveTagRequest) SetAppKey(v int64) *RemoveTagRequest {
	s.AppKey = &v
	return s
}

func (s *RemoveTagRequest) SetTagName(v string) *RemoveTagRequest {
	s.TagName = &v
	return s
}

func (s *RemoveTagRequest) Validate() error {
	return dara.Validate(s)
}

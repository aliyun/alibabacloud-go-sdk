// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *ListTagsRequest
	GetAppKey() *int64
}

type ListTagsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s ListTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagsRequest) GoString() string {
	return s.String()
}

func (s *ListTagsRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *ListTagsRequest) SetAppKey(v int64) *ListTagsRequest {
	s.AppKey = &v
	return s
}

func (s *ListTagsRequest) Validate() error {
	return dara.Validate(s)
}

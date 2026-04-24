// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListApiKeysRequest
	GetInstanceId() *string
	SetPage(v int32) *ListApiKeysRequest
	GetPage() *int32
	SetPageSize(v int32) *ListApiKeysRequest
	GetPageSize() *int32
}

type ListApiKeysRequest struct {
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysRequest) GoString() string {
	return s.String()
}

func (s *ListApiKeysRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApiKeysRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListApiKeysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApiKeysRequest) SetInstanceId(v string) *ListApiKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApiKeysRequest) SetPage(v int32) *ListApiKeysRequest {
	s.Page = &v
	return s
}

func (s *ListApiKeysRequest) SetPageSize(v int32) *ListApiKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListApiKeysRequest) Validate() error {
	return dara.Validate(s)
}

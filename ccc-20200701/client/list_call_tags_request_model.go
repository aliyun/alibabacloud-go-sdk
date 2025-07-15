// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListCallTagsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListCallTagsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCallTagsRequest
	GetPageSize() *int32
}

type ListCallTagsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCallTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCallTagsRequest) GoString() string {
	return s.String()
}

func (s *ListCallTagsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCallTagsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCallTagsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCallTagsRequest) SetInstanceId(v string) *ListCallTagsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCallTagsRequest) SetPageNumber(v int32) *ListCallTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCallTagsRequest) SetPageSize(v int32) *ListCallTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCallTagsRequest) Validate() error {
	return dara.Validate(s)
}

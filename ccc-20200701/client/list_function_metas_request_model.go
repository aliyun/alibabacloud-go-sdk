// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionMetasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHasHttpTrigger(v bool) *ListFunctionMetasRequest
	GetHasHttpTrigger() *bool
	SetInstanceId(v string) *ListFunctionMetasRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListFunctionMetasRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFunctionMetasRequest
	GetPageSize() *int32
}

type ListFunctionMetasRequest struct {
	// example:
	//
	// true
	HasHttpTrigger *bool `json:"HasHttpTrigger,omitempty" xml:"HasHttpTrigger,omitempty"`
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

func (s ListFunctionMetasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionMetasRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionMetasRequest) GetHasHttpTrigger() *bool {
	return s.HasHttpTrigger
}

func (s *ListFunctionMetasRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFunctionMetasRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFunctionMetasRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFunctionMetasRequest) SetHasHttpTrigger(v bool) *ListFunctionMetasRequest {
	s.HasHttpTrigger = &v
	return s
}

func (s *ListFunctionMetasRequest) SetInstanceId(v string) *ListFunctionMetasRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFunctionMetasRequest) SetPageNumber(v int32) *ListFunctionMetasRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionMetasRequest) SetPageSize(v int32) *ListFunctionMetasRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionMetasRequest) Validate() error {
	return dara.Validate(s)
}

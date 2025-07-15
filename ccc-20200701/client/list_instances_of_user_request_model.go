// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesOfUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListInstancesOfUserRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesOfUserRequest
	GetPageSize() *int32
}

type ListInstancesOfUserRequest struct {
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

func (s ListInstancesOfUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesOfUserRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesOfUserRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesOfUserRequest) SetPageNumber(v int32) *ListInstancesOfUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesOfUserRequest) SetPageSize(v int32) *ListInstancesOfUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesOfUserRequest) Validate() error {
	return dara.Validate(s)
}

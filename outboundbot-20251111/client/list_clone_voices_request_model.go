// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloneVoicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListCloneVoicesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListCloneVoicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloneVoicesRequest
	GetPageSize() *int32
	SetStatus(v string) *ListCloneVoicesRequest
	GetStatus() *string
}

type ListCloneVoicesRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status. Use Published.
	//
	// example:
	//
	// Published
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCloneVoicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloneVoicesRequest) GoString() string {
	return s.String()
}

func (s *ListCloneVoicesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloneVoicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloneVoicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloneVoicesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCloneVoicesRequest) SetInstanceId(v string) *ListCloneVoicesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCloneVoicesRequest) SetPageNumber(v int32) *ListCloneVoicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloneVoicesRequest) SetPageSize(v int32) *ListCloneVoicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloneVoicesRequest) SetStatus(v string) *ListCloneVoicesRequest {
	s.Status = &v
	return s
}

func (s *ListCloneVoicesRequest) Validate() error {
	return dara.Validate(s)
}

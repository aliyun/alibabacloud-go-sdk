// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutboundCallNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListOutboundCallNumbersRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListOutboundCallNumbersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOutboundCallNumbersRequest
	GetPageSize() *int32
}

type ListOutboundCallNumbersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
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

func (s ListOutboundCallNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundCallNumbersRequest) GoString() string {
	return s.String()
}

func (s *ListOutboundCallNumbersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOutboundCallNumbersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOutboundCallNumbersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOutboundCallNumbersRequest) SetInstanceId(v string) *ListOutboundCallNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOutboundCallNumbersRequest) SetPageNumber(v int32) *ListOutboundCallNumbersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOutboundCallNumbersRequest) SetPageSize(v int32) *ListOutboundCallNumbersRequest {
	s.PageSize = &v
	return s
}

func (s *ListOutboundCallNumbersRequest) Validate() error {
	return dara.Validate(s)
}

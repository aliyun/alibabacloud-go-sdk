// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutboundCallRestrictionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListOutboundCallRestrictionsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListOutboundCallRestrictionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOutboundCallRestrictionsRequest
	GetPageSize() *int32
	SetPolicy(v int32) *ListOutboundCallRestrictionsRequest
	GetPolicy() *int32
}

type ListOutboundCallRestrictionsRequest struct {
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
	// The policy. Valid values:
	//
	// 0: blacklist.
	//
	// 1: whitelist.
	//
	// example:
	//
	// 0
	Policy *int32 `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s ListOutboundCallRestrictionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundCallRestrictionsRequest) GoString() string {
	return s.String()
}

func (s *ListOutboundCallRestrictionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOutboundCallRestrictionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOutboundCallRestrictionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOutboundCallRestrictionsRequest) GetPolicy() *int32 {
	return s.Policy
}

func (s *ListOutboundCallRestrictionsRequest) SetInstanceId(v string) *ListOutboundCallRestrictionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOutboundCallRestrictionsRequest) SetPageNumber(v int32) *ListOutboundCallRestrictionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOutboundCallRestrictionsRequest) SetPageSize(v int32) *ListOutboundCallRestrictionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOutboundCallRestrictionsRequest) SetPolicy(v int32) *ListOutboundCallRestrictionsRequest {
	s.Policy = &v
	return s
}

func (s *ListOutboundCallRestrictionsRequest) Validate() error {
	return dara.Validate(s)
}

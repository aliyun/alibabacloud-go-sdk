// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListNumbersRequest
	GetInstanceId() *string
	SetNumber(v string) *ListNumbersRequest
	GetNumber() *string
	SetPageNumber(v int32) *ListNumbersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNumbersRequest
	GetPageSize() *int32
	SetSipTrunkId(v string) *ListNumbersRequest
	GetSipTrunkId() *string
	SetTenantId(v string) *ListNumbersRequest
	GetTenantId() *string
}

type ListNumbersRequest struct {
	// The instance ID.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The bound phone number.
	//
	// example:
	//
	// 057100001234
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The SIP trunk ID.
	//
	// example:
	//
	// 159d4b8361804ebda21242216a8981c2
	SipTrunkId *string `json:"SipTrunkId,omitempty" xml:"SipTrunkId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1234567890
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNumbersRequest) GoString() string {
	return s.String()
}

func (s *ListNumbersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNumbersRequest) GetNumber() *string {
	return s.Number
}

func (s *ListNumbersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNumbersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNumbersRequest) GetSipTrunkId() *string {
	return s.SipTrunkId
}

func (s *ListNumbersRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListNumbersRequest) SetInstanceId(v string) *ListNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNumbersRequest) SetNumber(v string) *ListNumbersRequest {
	s.Number = &v
	return s
}

func (s *ListNumbersRequest) SetPageNumber(v int32) *ListNumbersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNumbersRequest) SetPageSize(v int32) *ListNumbersRequest {
	s.PageSize = &v
	return s
}

func (s *ListNumbersRequest) SetSipTrunkId(v string) *ListNumbersRequest {
	s.SipTrunkId = &v
	return s
}

func (s *ListNumbersRequest) SetTenantId(v string) *ListNumbersRequest {
	s.TenantId = &v
	return s
}

func (s *ListNumbersRequest) Validate() error {
	return dara.Validate(s)
}

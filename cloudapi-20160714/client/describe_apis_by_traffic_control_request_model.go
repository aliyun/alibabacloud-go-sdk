// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisByTrafficControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeApisByTrafficControlRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisByTrafficControlRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeApisByTrafficControlRequest
	GetSecurityToken() *string
	SetTrafficControlId(v string) *DescribeApisByTrafficControlRequest
	GetTrafficControlId() *string
}

type DescribeApisByTrafficControlRequest struct {
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The security token included in the WebSocket request header. The system uses this token to authenticate the request.
	//
	// example:
	//
	// 9abe3317-3e22-4957-ab9f-dd893d0ac956
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the throttling policy that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
}

func (s DescribeApisByTrafficControlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByTrafficControlRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisByTrafficControlRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisByTrafficControlRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisByTrafficControlRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApisByTrafficControlRequest) GetTrafficControlId() *string {
	return s.TrafficControlId
}

func (s *DescribeApisByTrafficControlRequest) SetPageNumber(v int32) *DescribeApisByTrafficControlRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByTrafficControlRequest) SetPageSize(v int32) *DescribeApisByTrafficControlRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByTrafficControlRequest) SetSecurityToken(v string) *DescribeApisByTrafficControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApisByTrafficControlRequest) SetTrafficControlId(v string) *DescribeApisByTrafficControlRequest {
	s.TrafficControlId = &v
	return s
}

func (s *DescribeApisByTrafficControlRequest) Validate() error {
	return dara.Validate(s)
}

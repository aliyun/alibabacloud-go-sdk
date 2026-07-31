// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserResourcePackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeUserResourcePackageRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeUserResourcePackageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeUserResourcePackageRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeUserResourcePackageRequest
	GetSecurityToken() *string
	SetSortField(v string) *DescribeUserResourcePackageRequest
	GetSortField() *string
	SetSortRule(v string) *DescribeUserResourcePackageRequest
	GetSortRule() *string
	SetStatus(v string) *DescribeUserResourcePackageRequest
	GetStatus() *string
}

type DescribeUserResourcePackageRequest struct {
	// The resource plan instance ID.
	//
	// example:
	//
	// ****_ResourcePack-cn-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The sorting field. Valid values:
	//
	// - startTime: the effective period of the instance.
	//
	// - endTime: the expiration time of the instance.
	//
	// example:
	//
	// startTime
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The sorting collation. Default value: desc. Valid values:
	//
	// - asc
	//
	// - desc
	//
	// example:
	//
	// desc
	SortRule *string `json:"SortRule,omitempty" xml:"SortRule,omitempty"`
	// The status of the resource plan. Default value: valid. Valid values:
	//
	// - valid: Valid.
	//
	// - invalid: Invalid.
	//
	// - exhaust: Exhausted.
	//
	// example:
	//
	// valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeUserResourcePackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcePackageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserResourcePackageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeUserResourcePackageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUserResourcePackageRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeUserResourcePackageRequest) GetSortField() *string {
	return s.SortField
}

func (s *DescribeUserResourcePackageRequest) GetSortRule() *string {
	return s.SortRule
}

func (s *DescribeUserResourcePackageRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeUserResourcePackageRequest) SetInstanceId(v string) *DescribeUserResourcePackageRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserResourcePackageRequest) SetPageNumber(v int32) *DescribeUserResourcePackageRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserResourcePackageRequest) SetPageSize(v int32) *DescribeUserResourcePackageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserResourcePackageRequest) SetSecurityToken(v string) *DescribeUserResourcePackageRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUserResourcePackageRequest) SetSortField(v string) *DescribeUserResourcePackageRequest {
	s.SortField = &v
	return s
}

func (s *DescribeUserResourcePackageRequest) SetSortRule(v string) *DescribeUserResourcePackageRequest {
	s.SortRule = &v
	return s
}

func (s *DescribeUserResourcePackageRequest) SetStatus(v string) *DescribeUserResourcePackageRequest {
	s.Status = &v
	return s
}

func (s *DescribeUserResourcePackageRequest) Validate() error {
	return dara.Validate(s)
}

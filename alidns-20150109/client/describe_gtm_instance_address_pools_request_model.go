// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstanceAddressPoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeGtmInstanceAddressPoolsRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeGtmInstanceAddressPoolsRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeGtmInstanceAddressPoolsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGtmInstanceAddressPoolsRequest
	GetPageSize() *int32
}

type DescribeGtmInstanceAddressPoolsRequest struct {
	// The ID of the GTM instance that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// gtmtest
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language used by the user.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of the page to return. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Maximum value: **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGtmInstanceAddressPoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGtmInstanceAddressPoolsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmInstanceAddressPoolsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGtmInstanceAddressPoolsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGtmInstanceAddressPoolsRequest) SetInstanceId(v string) *DescribeGtmInstanceAddressPoolsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsRequest) SetLang(v string) *DescribeGtmInstanceAddressPoolsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsRequest) SetPageNumber(v int32) *DescribeGtmInstanceAddressPoolsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsRequest) SetPageSize(v int32) *DescribeGtmInstanceAddressPoolsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsRequest) Validate() error {
	return dara.Validate(s)
}

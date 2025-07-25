// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceAddressPoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDnsGtmInstanceAddressPoolsRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeDnsGtmInstanceAddressPoolsRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeDnsGtmInstanceAddressPoolsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDnsGtmInstanceAddressPoolsRequest
	GetPageSize() *int32
}

type DescribeDnsGtmInstanceAddressPoolsRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language to return some response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDnsGtmInstanceAddressPoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsGtmInstanceAddressPoolsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmInstanceAddressPoolsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDnsGtmInstanceAddressPoolsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDnsGtmInstanceAddressPoolsRequest) SetInstanceId(v string) *DescribeDnsGtmInstanceAddressPoolsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsRequest) SetLang(v string) *DescribeDnsGtmInstanceAddressPoolsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsRequest) SetPageNumber(v int32) *DescribeDnsGtmInstanceAddressPoolsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsRequest) SetPageSize(v int32) *DescribeDnsGtmInstanceAddressPoolsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsRequest) Validate() error {
	return dara.Validate(s)
}

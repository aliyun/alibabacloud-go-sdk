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
	// The instance ID. For more information, see [DescribeDnsGtmInstances](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describednsgtminstances?spm=a2c63.p38356.help-menu-search-29697.d_0).
	//
	// This parameter is required.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response. Valid values: en, zh, and ja. The default value is en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. The value starts from 1. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The maximum value is 100. The default value is 20.
	//
	// example:
	//
	// 20
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

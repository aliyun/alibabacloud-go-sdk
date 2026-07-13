// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceSystemCnameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDnsGtmInstanceSystemCnameRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeDnsGtmInstanceSystemCnameRequest
	GetLang() *string
}

type DescribeDnsGtmInstanceSystemCnameRequest struct {
	// The instance ID. <props="china">Call [DescribeDnsGtmInstances](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describednsgtminstances?spm=a2c4g.11186623.help-menu-29697.d_0_5_1_3_8_8.2aea3618RlSR9K) to obtain the ID.
	//
	// <props="intl">Call [DescribeDnsGtmInstances](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describednsgtminstances?spm=a2c63.p38356.help-menu-search-29697.d_0) to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response. The default value is en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeDnsGtmInstanceSystemCnameRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceSystemCnameRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceSystemCnameRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsGtmInstanceSystemCnameRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmInstanceSystemCnameRequest) SetInstanceId(v string) *DescribeDnsGtmInstanceSystemCnameRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmInstanceSystemCnameRequest) SetLang(v string) *DescribeDnsGtmInstanceSystemCnameRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmInstanceSystemCnameRequest) Validate() error {
	return dara.Validate(s)
}

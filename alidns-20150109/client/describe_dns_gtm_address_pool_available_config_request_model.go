// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAddressPoolAvailableConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDnsGtmAddressPoolAvailableConfigRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeDnsGtmAddressPoolAvailableConfigRequest
	GetLang() *string
}

type DescribeDnsGtmAddressPoolAvailableConfigRequest struct {
	// The ID of the instance.<props="china"> You can call [DescribeDnsGtmInstances](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describednsgtminstances?spm=a2c4g.11186623.help-menu-29697.d_0_5_1_3_8_8.2aea3618RlSR9K) to obtain the instance ID.
	//
	// <props="intl">You can call [DescribeDnsGtmInstances](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describednsgtminstances?spm=a2c63.p38356.help-menu-search-29697.d_0) to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of some returned parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeDnsGtmAddressPoolAvailableConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAddressPoolAvailableConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigRequest) SetInstanceId(v string) *DescribeDnsGtmAddressPoolAvailableConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigRequest) SetLang(v string) *DescribeDnsGtmAddressPoolAvailableConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigRequest) Validate() error {
	return dara.Validate(s)
}

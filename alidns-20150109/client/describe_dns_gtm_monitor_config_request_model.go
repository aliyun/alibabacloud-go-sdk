// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmMonitorConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDnsGtmMonitorConfigRequest
	GetLang() *string
	SetMonitorConfigId(v string) *DescribeDnsGtmMonitorConfigRequest
	GetMonitorConfigId() *string
}

type DescribeDnsGtmMonitorConfigRequest struct {
	// The language of the response. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the health check configuration.<props="china"> For more information, see [DescribeDnsGtmInstanceAddressPool](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describednsgtminstanceaddresspool?spm=a2c4g.11186623.help-menu-29697.d_0_5_1_3_9_6.7db77000nMCPI1).<props="intl"> For more information, see [DescribeDnsGtmInstanceAddressPool](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describednsgtminstanceaddresspool?spm=a2c63.p38356.help-menu-search-29697.d_0).
	//
	// This parameter is required.
	//
	// example:
	//
	// Moni*******
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
}

func (s DescribeDnsGtmMonitorConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmMonitorConfigRequest) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *DescribeDnsGtmMonitorConfigRequest) SetLang(v string) *DescribeDnsGtmMonitorConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigRequest) SetMonitorConfigId(v string) *DescribeDnsGtmMonitorConfigRequest {
	s.MonitorConfigId = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsProductInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDnsProductInstanceRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeDnsProductInstanceRequest
	GetLang() *string
	SetUserClientIp(v string) *DescribeDnsProductInstanceRequest
	GetUserClientIp() *string
}

type DescribeDnsProductInstanceRequest struct {
	// The instance ID. <props="china">You can call [DescribeDomainInfo](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describedomaininfo?spm=a2c4g.11186623.help-menu-search-29697.d_0) to obtain the instance ID.<props="intl">You can call [DescribeDomainInfo](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomaininfo?spm=a2c63.p38356.help-menu-search-29697.d_0) to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-8fxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// Default: en
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The client IP address.
	//
	// example:
	//
	// 192.0.2.0
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeDnsProductInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsProductInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsProductInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsProductInstanceRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsProductInstanceRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DescribeDnsProductInstanceRequest) SetInstanceId(v string) *DescribeDnsProductInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsProductInstanceRequest) SetLang(v string) *DescribeDnsProductInstanceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsProductInstanceRequest) SetUserClientIp(v string) *DescribeDnsProductInstanceRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsProductInstanceRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDomainOfDnsProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *ChangeDomainOfDnsProductRequest
	GetForce() *bool
	SetInstanceId(v string) *ChangeDomainOfDnsProductRequest
	GetInstanceId() *string
	SetLang(v string) *ChangeDomainOfDnsProductRequest
	GetLang() *string
	SetNewDomain(v string) *ChangeDomainOfDnsProductRequest
	GetNewDomain() *string
	SetUserClientIp(v string) *ChangeDomainOfDnsProductRequest
	GetUserClientIp() *string
}

type ChangeDomainOfDnsProductRequest struct {
	// Specifies whether to forcibly bind a domain name to the instance. Valid values:
	//
	// 	- **false****: no**
	//
	// 	- **true**: **yes**
	//
	// Default value: **false**.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The ID of the Alibaba Cloud Domain Name System (DNS) instance.
	//
	// You can call the [ListCloudGtmInstances ](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-listcloudgtminstances?spm=a2c63.p38356.help-menu-search-29697.d_0)operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-7sb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// Default value: **zh**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The domain name that you want to bind to the instance. If you leave this parameter empty, the domain name that is bound to the instance is unbound from the instance.
	//
	// example:
	//
	// newdomain.com
	NewDomain *string `json:"NewDomain,omitempty" xml:"NewDomain,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 1.1.1.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s ChangeDomainOfDnsProductRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeDomainOfDnsProductRequest) GoString() string {
	return s.String()
}

func (s *ChangeDomainOfDnsProductRequest) GetForce() *bool {
	return s.Force
}

func (s *ChangeDomainOfDnsProductRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ChangeDomainOfDnsProductRequest) GetLang() *string {
	return s.Lang
}

func (s *ChangeDomainOfDnsProductRequest) GetNewDomain() *string {
	return s.NewDomain
}

func (s *ChangeDomainOfDnsProductRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *ChangeDomainOfDnsProductRequest) SetForce(v bool) *ChangeDomainOfDnsProductRequest {
	s.Force = &v
	return s
}

func (s *ChangeDomainOfDnsProductRequest) SetInstanceId(v string) *ChangeDomainOfDnsProductRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeDomainOfDnsProductRequest) SetLang(v string) *ChangeDomainOfDnsProductRequest {
	s.Lang = &v
	return s
}

func (s *ChangeDomainOfDnsProductRequest) SetNewDomain(v string) *ChangeDomainOfDnsProductRequest {
	s.NewDomain = &v
	return s
}

func (s *ChangeDomainOfDnsProductRequest) SetUserClientIp(v string) *ChangeDomainOfDnsProductRequest {
	s.UserClientIp = &v
	return s
}

func (s *ChangeDomainOfDnsProductRequest) Validate() error {
	return dara.Validate(s)
}

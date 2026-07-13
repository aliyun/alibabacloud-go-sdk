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
	// Specifies whether to forcefully attach the domain name.
	//
	// Valid values:
	//
	// - **false**: No
	//
	// - **true**: Yes
	//
	// The default value is **false**.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The ID of the Cloud DNS product.
	//
	// You can obtain the ID by calling [ListCloudGtmInstances](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-listcloudgtminstances?spm=a2c63.p38356.help-menu-search-29697.d_0).
	//
	// This parameter is required.
	//
	// example:
	//
	// i-7XX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the request and response.
	//
	// Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// The default value is **zh**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The domain name that you want to attach. If you leave this parameter empty, the currently attached domain name is detached.
	//
	// example:
	//
	// example.com
	NewDomain *string `json:"NewDomain,omitempty" xml:"NewDomain,omitempty"`
	// The client IP address.
	//
	// example:
	//
	// 1.1.XX.XX
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

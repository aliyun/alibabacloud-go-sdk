// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDnsHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryDnsHostRequest
	GetInstanceId() *string
	SetLang(v string) *QueryDnsHostRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryDnsHostRequest
	GetUserClientIp() *string
}

type QueryDnsHostRequest struct {
	// The ID of the domain name instance. Call the QueryDomainList API to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ST2017120814571100001303
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language for returned error messages. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The user\\"s IP address.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDnsHostRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDnsHostRequest) GoString() string {
	return s.String()
}

func (s *QueryDnsHostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryDnsHostRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryDnsHostRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryDnsHostRequest) SetInstanceId(v string) *QueryDnsHostRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryDnsHostRequest) SetLang(v string) *QueryDnsHostRequest {
	s.Lang = &v
	return s
}

func (s *QueryDnsHostRequest) SetUserClientIp(v string) *QueryDnsHostRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryDnsHostRequest) Validate() error {
	return dara.Validate(s)
}

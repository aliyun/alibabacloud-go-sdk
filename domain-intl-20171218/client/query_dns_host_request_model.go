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
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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

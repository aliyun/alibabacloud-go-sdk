// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainByInstanceIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryDomainByInstanceIdRequest
	GetInstanceId() *string
	SetLang(v string) *QueryDomainByInstanceIdRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryDomainByInstanceIdRequest
	GetUserClientIp() *string
}

type QueryDomainByInstanceIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// S20131205001****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainByInstanceIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByInstanceIdRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainByInstanceIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryDomainByInstanceIdRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryDomainByInstanceIdRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryDomainByInstanceIdRequest) SetInstanceId(v string) *QueryDomainByInstanceIdRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryDomainByInstanceIdRequest) SetLang(v string) *QueryDomainByInstanceIdRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainByInstanceIdRequest) SetUserClientIp(v string) *QueryDomainByInstanceIdRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryDomainByInstanceIdRequest) Validate() error {
	return dara.Validate(s)
}

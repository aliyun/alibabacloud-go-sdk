// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDSRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QueryDSRecordRequest
	GetDomainName() *string
	SetLang(v string) *QueryDSRecordRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryDSRecordRequest
	GetUserClientIp() *string
}

type QueryDSRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDSRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDSRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryDSRecordRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDSRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryDSRecordRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryDSRecordRequest) SetDomainName(v string) *QueryDSRecordRequest {
	s.DomainName = &v
	return s
}

func (s *QueryDSRecordRequest) SetLang(v string) *QueryDSRecordRequest {
	s.Lang = &v
	return s
}

func (s *QueryDSRecordRequest) SetUserClientIp(v string) *QueryDSRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryDSRecordRequest) Validate() error {
	return dara.Validate(s)
}

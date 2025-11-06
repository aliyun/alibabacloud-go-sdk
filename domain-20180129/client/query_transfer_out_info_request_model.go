// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTransferOutInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QueryTransferOutInfoRequest
	GetDomainName() *string
	SetLang(v string) *QueryTransferOutInfoRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryTransferOutInfoRequest
	GetUserClientIp() *string
}

type QueryTransferOutInfoRequest struct {
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

func (s QueryTransferOutInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTransferOutInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryTransferOutInfoRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryTransferOutInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryTransferOutInfoRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryTransferOutInfoRequest) SetDomainName(v string) *QueryTransferOutInfoRequest {
	s.DomainName = &v
	return s
}

func (s *QueryTransferOutInfoRequest) SetLang(v string) *QueryTransferOutInfoRequest {
	s.Lang = &v
	return s
}

func (s *QueryTransferOutInfoRequest) SetUserClientIp(v string) *QueryTransferOutInfoRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryTransferOutInfoRequest) Validate() error {
	return dara.Validate(s)
}

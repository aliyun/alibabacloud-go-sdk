// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCancelingTransferInRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SaveSingleTaskForCancelingTransferInRequest
	GetDomainName() *string
	SetLang(v string) *SaveSingleTaskForCancelingTransferInRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForCancelingTransferInRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForCancelingTransferInRequest struct {
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

func (s SaveSingleTaskForCancelingTransferInRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCancelingTransferInRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCancelingTransferInRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForCancelingTransferInRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForCancelingTransferInRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForCancelingTransferInRequest) SetDomainName(v string) *SaveSingleTaskForCancelingTransferInRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferInRequest) SetLang(v string) *SaveSingleTaskForCancelingTransferInRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferInRequest) SetUserClientIp(v string) *SaveSingleTaskForCancelingTransferInRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferInRequest) Validate() error {
	return dara.Validate(s)
}

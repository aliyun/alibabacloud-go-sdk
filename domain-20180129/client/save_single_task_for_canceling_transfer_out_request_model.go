// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCancelingTransferOutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SaveSingleTaskForCancelingTransferOutRequest
	GetDomainName() *string
	SetLang(v string) *SaveSingleTaskForCancelingTransferOutRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForCancelingTransferOutRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForCancelingTransferOutRequest struct {
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

func (s SaveSingleTaskForCancelingTransferOutRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCancelingTransferOutRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCancelingTransferOutRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForCancelingTransferOutRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForCancelingTransferOutRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForCancelingTransferOutRequest) SetDomainName(v string) *SaveSingleTaskForCancelingTransferOutRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferOutRequest) SetLang(v string) *SaveSingleTaskForCancelingTransferOutRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferOutRequest) SetUserClientIp(v string) *SaveSingleTaskForCancelingTransferOutRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferOutRequest) Validate() error {
	return dara.Validate(s)
}

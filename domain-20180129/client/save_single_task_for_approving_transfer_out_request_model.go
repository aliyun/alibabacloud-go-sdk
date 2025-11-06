// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForApprovingTransferOutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SaveSingleTaskForApprovingTransferOutRequest
	GetDomainName() *string
	SetLang(v string) *SaveSingleTaskForApprovingTransferOutRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForApprovingTransferOutRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForApprovingTransferOutRequest struct {
	// This parameter is required.
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

func (s SaveSingleTaskForApprovingTransferOutRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForApprovingTransferOutRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForApprovingTransferOutRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForApprovingTransferOutRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForApprovingTransferOutRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForApprovingTransferOutRequest) SetDomainName(v string) *SaveSingleTaskForApprovingTransferOutRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForApprovingTransferOutRequest) SetLang(v string) *SaveSingleTaskForApprovingTransferOutRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForApprovingTransferOutRequest) SetUserClientIp(v string) *SaveSingleTaskForApprovingTransferOutRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForApprovingTransferOutRequest) Validate() error {
	return dara.Validate(s)
}

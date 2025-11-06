// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForApplyQuickTransferOutOpenlyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SaveSingleTaskForApplyQuickTransferOutOpenlyRequest
	GetDomainName() *string
	SetLang(v string) *SaveSingleTaskForApplyQuickTransferOutOpenlyRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForApplyQuickTransferOutOpenlyRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForApplyQuickTransferOutOpenlyRequest struct {
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

func (s SaveSingleTaskForApplyQuickTransferOutOpenlyRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForApplyQuickTransferOutOpenlyRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyRequest) SetDomainName(v string) *SaveSingleTaskForApplyQuickTransferOutOpenlyRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyRequest) SetLang(v string) *SaveSingleTaskForApplyQuickTransferOutOpenlyRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyRequest) SetUserClientIp(v string) *SaveSingleTaskForApplyQuickTransferOutOpenlyRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyRequest) Validate() error {
	return dara.Validate(s)
}

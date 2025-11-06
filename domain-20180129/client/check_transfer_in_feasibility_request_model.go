// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTransferInFeasibilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *CheckTransferInFeasibilityRequest
	GetDomainName() *string
	SetLang(v string) *CheckTransferInFeasibilityRequest
	GetLang() *string
	SetTransferAuthorizationCode(v string) *CheckTransferInFeasibilityRequest
	GetTransferAuthorizationCode() *string
	SetUserClientIp(v string) *CheckTransferInFeasibilityRequest
	GetUserClientIp() *string
}

type CheckTransferInFeasibilityRequest struct {
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
	// test
	TransferAuthorizationCode *string `json:"TransferAuthorizationCode,omitempty" xml:"TransferAuthorizationCode,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s CheckTransferInFeasibilityRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckTransferInFeasibilityRequest) GoString() string {
	return s.String()
}

func (s *CheckTransferInFeasibilityRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CheckTransferInFeasibilityRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckTransferInFeasibilityRequest) GetTransferAuthorizationCode() *string {
	return s.TransferAuthorizationCode
}

func (s *CheckTransferInFeasibilityRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *CheckTransferInFeasibilityRequest) SetDomainName(v string) *CheckTransferInFeasibilityRequest {
	s.DomainName = &v
	return s
}

func (s *CheckTransferInFeasibilityRequest) SetLang(v string) *CheckTransferInFeasibilityRequest {
	s.Lang = &v
	return s
}

func (s *CheckTransferInFeasibilityRequest) SetTransferAuthorizationCode(v string) *CheckTransferInFeasibilityRequest {
	s.TransferAuthorizationCode = &v
	return s
}

func (s *CheckTransferInFeasibilityRequest) SetUserClientIp(v string) *CheckTransferInFeasibilityRequest {
	s.UserClientIp = &v
	return s
}

func (s *CheckTransferInFeasibilityRequest) Validate() error {
	return dara.Validate(s)
}

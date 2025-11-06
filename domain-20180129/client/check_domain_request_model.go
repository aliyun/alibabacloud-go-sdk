// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *CheckDomainRequest
	GetDomainName() *string
	SetFeeCommand(v string) *CheckDomainRequest
	GetFeeCommand() *string
	SetFeeCurrency(v string) *CheckDomainRequest
	GetFeeCurrency() *string
	SetFeePeriod(v int32) *CheckDomainRequest
	GetFeePeriod() *int32
	SetLang(v string) *CheckDomainRequest
	GetLang() *string
}

type CheckDomainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test**.xin
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// create
	FeeCommand *string `json:"FeeCommand,omitempty" xml:"FeeCommand,omitempty"`
	// example:
	//
	// USD
	FeeCurrency *string `json:"FeeCurrency,omitempty" xml:"FeeCurrency,omitempty"`
	// example:
	//
	// 1
	FeePeriod *int32 `json:"FeePeriod,omitempty" xml:"FeePeriod,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s CheckDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDomainRequest) GoString() string {
	return s.String()
}

func (s *CheckDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CheckDomainRequest) GetFeeCommand() *string {
	return s.FeeCommand
}

func (s *CheckDomainRequest) GetFeeCurrency() *string {
	return s.FeeCurrency
}

func (s *CheckDomainRequest) GetFeePeriod() *int32 {
	return s.FeePeriod
}

func (s *CheckDomainRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckDomainRequest) SetDomainName(v string) *CheckDomainRequest {
	s.DomainName = &v
	return s
}

func (s *CheckDomainRequest) SetFeeCommand(v string) *CheckDomainRequest {
	s.FeeCommand = &v
	return s
}

func (s *CheckDomainRequest) SetFeeCurrency(v string) *CheckDomainRequest {
	s.FeeCurrency = &v
	return s
}

func (s *CheckDomainRequest) SetFeePeriod(v int32) *CheckDomainRequest {
	s.FeePeriod = &v
	return s
}

func (s *CheckDomainRequest) SetLang(v string) *CheckDomainRequest {
	s.Lang = &v
	return s
}

func (s *CheckDomainRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *AddDomainBackupRequest
	GetDomainName() *string
	SetLang(v string) *AddDomainBackupRequest
	GetLang() *string
	SetPeriodType(v string) *AddDomainBackupRequest
	GetPeriodType() *string
}

type AddDomainBackupRequest struct {
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.aliyun.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The backup cycle. Valid values:
	//
	// 	- DAY: backs up data on a daily basis.
	//
	// 	- HOUR: backs up data on an hourly basis.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
}

func (s AddDomainBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDomainBackupRequest) GoString() string {
	return s.String()
}

func (s *AddDomainBackupRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddDomainBackupRequest) GetLang() *string {
	return s.Lang
}

func (s *AddDomainBackupRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *AddDomainBackupRequest) SetDomainName(v string) *AddDomainBackupRequest {
	s.DomainName = &v
	return s
}

func (s *AddDomainBackupRequest) SetLang(v string) *AddDomainBackupRequest {
	s.Lang = &v
	return s
}

func (s *AddDomainBackupRequest) SetPeriodType(v string) *AddDomainBackupRequest {
	s.PeriodType = &v
	return s
}

func (s *AddDomainBackupRequest) Validate() error {
	return dara.Validate(s)
}

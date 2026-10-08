// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckProcessingServerLockApplyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *CheckProcessingServerLockApplyRequest
	GetDomainName() *string
	SetFeePeriod(v int32) *CheckProcessingServerLockApplyRequest
	GetFeePeriod() *int32
	SetLang(v string) *CheckProcessingServerLockApplyRequest
	GetLang() *string
	SetUserClientIp(v string) *CheckProcessingServerLockApplyRequest
	GetUserClientIp() *string
}

type CheckProcessingServerLockApplyRequest struct {
	// The domain name to be checked.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Registration period in years. Unit: **year(s)**. Valid range: **1 to 10*	- years.
	//
	// example:
	//
	// 1
	FeePeriod *int32 `json:"FeePeriod,omitempty" xml:"FeePeriod,omitempty"`
	// Language of error messages returned by the API. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// User IP address. You can set it to **127.0.0.1**.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s CheckProcessingServerLockApplyRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckProcessingServerLockApplyRequest) GoString() string {
	return s.String()
}

func (s *CheckProcessingServerLockApplyRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CheckProcessingServerLockApplyRequest) GetFeePeriod() *int32 {
	return s.FeePeriod
}

func (s *CheckProcessingServerLockApplyRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckProcessingServerLockApplyRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *CheckProcessingServerLockApplyRequest) SetDomainName(v string) *CheckProcessingServerLockApplyRequest {
	s.DomainName = &v
	return s
}

func (s *CheckProcessingServerLockApplyRequest) SetFeePeriod(v int32) *CheckProcessingServerLockApplyRequest {
	s.FeePeriod = &v
	return s
}

func (s *CheckProcessingServerLockApplyRequest) SetLang(v string) *CheckProcessingServerLockApplyRequest {
	s.Lang = &v
	return s
}

func (s *CheckProcessingServerLockApplyRequest) SetUserClientIp(v string) *CheckProcessingServerLockApplyRequest {
	s.UserClientIp = &v
	return s
}

func (s *CheckProcessingServerLockApplyRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMaxYearOfServerLockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckAction(v string) *CheckMaxYearOfServerLockRequest
	GetCheckAction() *string
	SetDomainName(v string) *CheckMaxYearOfServerLockRequest
	GetDomainName() *string
	SetLang(v string) *CheckMaxYearOfServerLockRequest
	GetLang() *string
	SetUserClientIp(v string) *CheckMaxYearOfServerLockRequest
	GetUserClientIp() *string
}

type CheckMaxYearOfServerLockRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// activate
	CheckAction *string `json:"CheckAction,omitempty" xml:"CheckAction,omitempty"`
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

func (s CheckMaxYearOfServerLockRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckMaxYearOfServerLockRequest) GoString() string {
	return s.String()
}

func (s *CheckMaxYearOfServerLockRequest) GetCheckAction() *string {
	return s.CheckAction
}

func (s *CheckMaxYearOfServerLockRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CheckMaxYearOfServerLockRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckMaxYearOfServerLockRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *CheckMaxYearOfServerLockRequest) SetCheckAction(v string) *CheckMaxYearOfServerLockRequest {
	s.CheckAction = &v
	return s
}

func (s *CheckMaxYearOfServerLockRequest) SetDomainName(v string) *CheckMaxYearOfServerLockRequest {
	s.DomainName = &v
	return s
}

func (s *CheckMaxYearOfServerLockRequest) SetLang(v string) *CheckMaxYearOfServerLockRequest {
	s.Lang = &v
	return s
}

func (s *CheckMaxYearOfServerLockRequest) SetUserClientIp(v string) *CheckMaxYearOfServerLockRequest {
	s.UserClientIp = &v
	return s
}

func (s *CheckMaxYearOfServerLockRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityScoreRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalType(v string) *GetSecurityScoreRuleRequest
	GetCalType() *string
	SetLang(v string) *GetSecurityScoreRuleRequest
	GetLang() *string
}

type GetSecurityScoreRuleRequest struct {
	// The old or new version of the security score rule. If you set this parameter to **home_security_score**, the new version of the security score rule is returned. Otherwise, the old version of the security score rule is returned by default.
	//
	// example:
	//
	// home_security_score
	CalType *string `json:"CalType,omitempty" xml:"CalType,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetSecurityScoreRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityScoreRuleRequest) GoString() string {
	return s.String()
}

func (s *GetSecurityScoreRuleRequest) GetCalType() *string {
	return s.CalType
}

func (s *GetSecurityScoreRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *GetSecurityScoreRuleRequest) SetCalType(v string) *GetSecurityScoreRuleRequest {
	s.CalType = &v
	return s
}

func (s *GetSecurityScoreRuleRequest) SetLang(v string) *GetSecurityScoreRuleRequest {
	s.Lang = &v
	return s
}

func (s *GetSecurityScoreRuleRequest) Validate() error {
	return dara.Validate(s)
}

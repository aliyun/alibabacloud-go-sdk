// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBuildRiskDefineRuleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetBuildRiskDefineRuleConfigRequest
	GetLang() *string
}

type GetBuildRiskDefineRuleConfigRequest struct {
	// The language of the content within the request and response. Valid values:
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

func (s GetBuildRiskDefineRuleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBuildRiskDefineRuleConfigRequest) GoString() string {
	return s.String()
}

func (s *GetBuildRiskDefineRuleConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *GetBuildRiskDefineRuleConfigRequest) SetLang(v string) *GetBuildRiskDefineRuleConfigRequest {
	s.Lang = &v
	return s
}

func (s *GetBuildRiskDefineRuleConfigRequest) Validate() error {
	return dara.Validate(s)
}

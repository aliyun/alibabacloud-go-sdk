// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSensitiveDefineRuleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetSensitiveDefineRuleConfigRequest
	GetLang() *string
	SetSource(v string) *GetSensitiveDefineRuleConfigRequest
	GetSource() *string
}

type GetSensitiveDefineRuleConfigRequest struct {
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
	// The source of the check rules. Valid values:
	//
	// 	- **image**: image.
	//
	// 	- **agentless**: agentless detection.
	//
	// example:
	//
	// image
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetSensitiveDefineRuleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSensitiveDefineRuleConfigRequest) GoString() string {
	return s.String()
}

func (s *GetSensitiveDefineRuleConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *GetSensitiveDefineRuleConfigRequest) GetSource() *string {
	return s.Source
}

func (s *GetSensitiveDefineRuleConfigRequest) SetLang(v string) *GetSensitiveDefineRuleConfigRequest {
	s.Lang = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigRequest) SetSource(v string) *GetSensitiveDefineRuleConfigRequest {
	s.Source = &v
	return s
}

func (s *GetSensitiveDefineRuleConfigRequest) Validate() error {
	return dara.Validate(s)
}

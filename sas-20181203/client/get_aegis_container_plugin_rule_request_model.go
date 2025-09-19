// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAegisContainerPluginRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetAegisContainerPluginRuleRequest
	GetId() *string
	SetLang(v string) *GetAegisContainerPluginRuleRequest
	GetLang() *string
	SetRuleType(v int32) *GetAegisContainerPluginRuleRequest
	GetRuleType() *int32
}

type GetAegisContainerPluginRuleRequest struct {
	// The ID of the defense rule against container escapes.
	//
	// >  You can call the [ListAegisContainerPluginRule](~~ListAegisContainerPluginRule~~) operation to obtain the ID.
	//
	// example:
	//
	// 1141****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **0**: custom rule
	//
	// 	- **1**: system rule
	//
	// example:
	//
	// 0
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s GetAegisContainerPluginRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAegisContainerPluginRuleRequest) GoString() string {
	return s.String()
}

func (s *GetAegisContainerPluginRuleRequest) GetId() *string {
	return s.Id
}

func (s *GetAegisContainerPluginRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *GetAegisContainerPluginRuleRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *GetAegisContainerPluginRuleRequest) SetId(v string) *GetAegisContainerPluginRuleRequest {
	s.Id = &v
	return s
}

func (s *GetAegisContainerPluginRuleRequest) SetLang(v string) *GetAegisContainerPluginRuleRequest {
	s.Lang = &v
	return s
}

func (s *GetAegisContainerPluginRuleRequest) SetRuleType(v int32) *GetAegisContainerPluginRuleRequest {
	s.RuleType = &v
	return s
}

func (s *GetAegisContainerPluginRuleRequest) Validate() error {
	return dara.Validate(s)
}

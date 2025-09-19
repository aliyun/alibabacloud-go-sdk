// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAegisContainerPluginRuleCriteriaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetAegisContainerPluginRuleCriteriaRequest
	GetLang() *string
	SetValue(v string) *GetAegisContainerPluginRuleCriteriaRequest
	GetValue() *string
}

type GetAegisContainerPluginRuleCriteriaRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The value of the search condition.
	//
	// example:
	//
	// ss
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAegisContainerPluginRuleCriteriaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAegisContainerPluginRuleCriteriaRequest) GoString() string {
	return s.String()
}

func (s *GetAegisContainerPluginRuleCriteriaRequest) GetLang() *string {
	return s.Lang
}

func (s *GetAegisContainerPluginRuleCriteriaRequest) GetValue() *string {
	return s.Value
}

func (s *GetAegisContainerPluginRuleCriteriaRequest) SetLang(v string) *GetAegisContainerPluginRuleCriteriaRequest {
	s.Lang = &v
	return s
}

func (s *GetAegisContainerPluginRuleCriteriaRequest) SetValue(v string) *GetAegisContainerPluginRuleCriteriaRequest {
	s.Value = &v
	return s
}

func (s *GetAegisContainerPluginRuleCriteriaRequest) Validate() error {
	return dara.Validate(s)
}

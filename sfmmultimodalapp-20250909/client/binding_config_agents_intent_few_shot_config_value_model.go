// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindingConfigAgentsIntentFewShotConfigValue interface {
	dara.Model
	String() string
	GoString() string
	SetQuery(v string) *BindingConfigAgentsIntentFewShotConfigValue
	GetQuery() *string
	SetParameters(v map[string]interface{}) *BindingConfigAgentsIntentFewShotConfigValue
	GetParameters() map[string]interface{}
}

type BindingConfigAgentsIntentFewShotConfigValue struct {
	Query      *string                `json:"Query,omitempty" xml:"Query,omitempty"`
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s BindingConfigAgentsIntentFewShotConfigValue) String() string {
	return dara.Prettify(s)
}

func (s BindingConfigAgentsIntentFewShotConfigValue) GoString() string {
	return s.String()
}

func (s *BindingConfigAgentsIntentFewShotConfigValue) GetQuery() *string {
	return s.Query
}

func (s *BindingConfigAgentsIntentFewShotConfigValue) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *BindingConfigAgentsIntentFewShotConfigValue) SetQuery(v string) *BindingConfigAgentsIntentFewShotConfigValue {
	s.Query = &v
	return s
}

func (s *BindingConfigAgentsIntentFewShotConfigValue) SetParameters(v map[string]interface{}) *BindingConfigAgentsIntentFewShotConfigValue {
	s.Parameters = v
	return s
}

func (s *BindingConfigAgentsIntentFewShotConfigValue) Validate() error {
	return dara.Validate(s)
}

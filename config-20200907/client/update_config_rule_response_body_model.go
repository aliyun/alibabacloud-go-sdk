// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleId(v string) *UpdateConfigRuleResponseBody
	GetConfigRuleId() *string
	SetRequestId(v string) *UpdateConfigRuleResponseBody
	GetRequestId() *string
}

type UpdateConfigRuleResponseBody struct {
	// The ID of the rule.
	//
	// example:
	//
	// cr-a260626622af0005****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConfigRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigRuleResponseBody) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *UpdateConfigRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConfigRuleResponseBody) SetConfigRuleId(v string) *UpdateConfigRuleResponseBody {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateConfigRuleResponseBody) SetRequestId(v string) *UpdateConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

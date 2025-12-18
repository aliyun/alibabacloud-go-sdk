// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleId(v string) *CreateConfigRuleResponseBody
	GetConfigRuleId() *string
	SetRequestId(v string) *CreateConfigRuleResponseBody
	GetRequestId() *string
}

type CreateConfigRuleResponseBody struct {
	// The rule ID.
	//
	// example:
	//
	// cr-5772ba41209e007b****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConfigRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigRuleResponseBody) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *CreateConfigRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConfigRuleResponseBody) SetConfigRuleId(v string) *CreateConfigRuleResponseBody {
	s.ConfigRuleId = &v
	return s
}

func (s *CreateConfigRuleResponseBody) SetRequestId(v string) *CreateConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConfigRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

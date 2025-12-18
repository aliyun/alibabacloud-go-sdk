// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartConfigRuleEvaluationByResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleId(v string) *StartConfigRuleEvaluationByResourceResponseBody
	GetConfigRuleId() *string
	SetRequestId(v string) *StartConfigRuleEvaluationByResourceResponseBody
	GetRequestId() *string
}

type StartConfigRuleEvaluationByResourceResponseBody struct {
	// example:
	//
	// cr-2da35180a8d1008e****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// example:
	//
	// A7A0FFF8-0B44-40C6-8BBF-3A185EFDF3F7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartConfigRuleEvaluationByResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartConfigRuleEvaluationByResourceResponseBody) GoString() string {
	return s.String()
}

func (s *StartConfigRuleEvaluationByResourceResponseBody) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *StartConfigRuleEvaluationByResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartConfigRuleEvaluationByResourceResponseBody) SetConfigRuleId(v string) *StartConfigRuleEvaluationByResourceResponseBody {
	s.ConfigRuleId = &v
	return s
}

func (s *StartConfigRuleEvaluationByResourceResponseBody) SetRequestId(v string) *StartConfigRuleEvaluationByResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartConfigRuleEvaluationByResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

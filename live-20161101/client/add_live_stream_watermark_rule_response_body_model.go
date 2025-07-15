// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveStreamWatermarkRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveStreamWatermarkRuleResponseBody
	GetRequestId() *string
	SetRuleId(v string) *AddLiveStreamWatermarkRuleResponseBody
	GetRuleId() *string
}

type AddLiveStreamWatermarkRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5c6a2a0df228-4a64-af62-20e91b96****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the watermark rule.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s AddLiveStreamWatermarkRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveStreamWatermarkRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveStreamWatermarkRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveStreamWatermarkRuleResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *AddLiveStreamWatermarkRuleResponseBody) SetRequestId(v string) *AddLiveStreamWatermarkRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveStreamWatermarkRuleResponseBody) SetRuleId(v string) *AddLiveStreamWatermarkRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *AddLiveStreamWatermarkRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAIProduceRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveAIProduceRulesResponseBody
	GetRequestId() *string
	SetRulesId(v string) *AddLiveAIProduceRulesResponseBody
	GetRulesId() *string
}

type AddLiveAIProduceRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0df228-4a64-af62-20e91b96****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the subtitle rule.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	RulesId *string `json:"RulesId,omitempty" xml:"RulesId,omitempty"`
}

func (s AddLiveAIProduceRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAIProduceRulesResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveAIProduceRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveAIProduceRulesResponseBody) GetRulesId() *string {
	return s.RulesId
}

func (s *AddLiveAIProduceRulesResponseBody) SetRequestId(v string) *AddLiveAIProduceRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveAIProduceRulesResponseBody) SetRulesId(v string) *AddLiveAIProduceRulesResponseBody {
	s.RulesId = &v
	return s
}

func (s *AddLiveAIProduceRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

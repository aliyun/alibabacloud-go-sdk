// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAIProduceRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveAIProduceRulesResponseBody
	GetRequestId() *string
}

type UpdateLiveAIProduceRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0df228-4a64-af62-20e91b96****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveAIProduceRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAIProduceRulesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveAIProduceRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveAIProduceRulesResponseBody) SetRequestId(v string) *UpdateLiveAIProduceRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveAIProduceRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

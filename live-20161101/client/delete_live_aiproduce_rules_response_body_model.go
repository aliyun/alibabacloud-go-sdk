// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAIProduceRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveAIProduceRulesResponseBody
	GetRequestId() *string
}

type DeleteLiveAIProduceRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0df228-4a64- af62-20e91b96****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveAIProduceRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAIProduceRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveAIProduceRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveAIProduceRulesResponseBody) SetRequestId(v string) *DeleteLiveAIProduceRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveAIProduceRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

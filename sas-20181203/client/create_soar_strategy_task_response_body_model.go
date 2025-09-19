// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSoarStrategyTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSoarStrategyTaskResponseBody
	GetRequestId() *string
	SetStrategyTaskId(v int64) *CreateSoarStrategyTaskResponseBody
	GetStrategyTaskId() *int64
}

type CreateSoarStrategyTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 43313389-DED8-5BB7-8CB9-F22CDEB744DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the policy task.
	//
	// example:
	//
	// 10688
	StrategyTaskId *int64 `json:"StrategyTaskId,omitempty" xml:"StrategyTaskId,omitempty"`
}

func (s CreateSoarStrategyTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSoarStrategyTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSoarStrategyTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSoarStrategyTaskResponseBody) GetStrategyTaskId() *int64 {
	return s.StrategyTaskId
}

func (s *CreateSoarStrategyTaskResponseBody) SetRequestId(v string) *CreateSoarStrategyTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSoarStrategyTaskResponseBody) SetStrategyTaskId(v int64) *CreateSoarStrategyTaskResponseBody {
	s.StrategyTaskId = &v
	return s
}

func (s *CreateSoarStrategyTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

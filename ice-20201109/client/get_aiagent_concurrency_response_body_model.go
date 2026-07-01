// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIAgentConcurrencyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActiveCount(v int32) *GetAIAgentConcurrencyResponseBody
	GetActiveCount() *int32
	SetRequestId(v string) *GetAIAgentConcurrencyResponseBody
	GetRequestId() *string
}

type GetAIAgentConcurrencyResponseBody struct {
	// The number of active concurrent calls.
	//
	// example:
	//
	// 39
	ActiveCount *int32 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAIAgentConcurrencyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAIAgentConcurrencyResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIAgentConcurrencyResponseBody) GetActiveCount() *int32 {
	return s.ActiveCount
}

func (s *GetAIAgentConcurrencyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAIAgentConcurrencyResponseBody) SetActiveCount(v int32) *GetAIAgentConcurrencyResponseBody {
	s.ActiveCount = &v
	return s
}

func (s *GetAIAgentConcurrencyResponseBody) SetRequestId(v string) *GetAIAgentConcurrencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIAgentConcurrencyResponseBody) Validate() error {
	return dara.Validate(s)
}

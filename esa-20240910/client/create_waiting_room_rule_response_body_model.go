// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWaitingRoomRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWaitingRoomRuleResponseBody
	GetRequestId() *string
	SetWaitingRoomRuleId(v int64) *CreateWaitingRoomRuleResponseBody
	GetWaitingRoomRuleId() *int64
}

type CreateWaitingRoomRuleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Waiting room bypass rule ID.
	//
	// example:
	//
	// 420072638347264
	WaitingRoomRuleId *int64 `json:"WaitingRoomRuleId,omitempty" xml:"WaitingRoomRuleId,omitempty"`
}

func (s CreateWaitingRoomRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWaitingRoomRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWaitingRoomRuleResponseBody) GetWaitingRoomRuleId() *int64 {
	return s.WaitingRoomRuleId
}

func (s *CreateWaitingRoomRuleResponseBody) SetRequestId(v string) *CreateWaitingRoomRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWaitingRoomRuleResponseBody) SetWaitingRoomRuleId(v int64) *CreateWaitingRoomRuleResponseBody {
	s.WaitingRoomRuleId = &v
	return s
}

func (s *CreateWaitingRoomRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

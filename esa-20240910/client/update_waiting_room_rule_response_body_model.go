// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaitingRoomRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWaitingRoomRuleResponseBody
	GetRequestId() *string
}

type UpdateWaitingRoomRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWaitingRoomRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaitingRoomRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWaitingRoomRuleResponseBody) SetRequestId(v string) *UpdateWaitingRoomRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWaitingRoomRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

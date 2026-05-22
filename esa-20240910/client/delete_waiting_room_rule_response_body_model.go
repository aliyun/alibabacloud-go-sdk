// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWaitingRoomRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWaitingRoomRuleResponseBody
	GetRequestId() *string
}

type DeleteWaitingRoomRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWaitingRoomRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWaitingRoomRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWaitingRoomRuleResponseBody) SetRequestId(v string) *DeleteWaitingRoomRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWaitingRoomRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

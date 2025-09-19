// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlarmEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAlarmEventResponseBody
	GetRequestId() *string
}

type UpdateAlarmEventResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AC1E0E53-DEF3-5D96-B170-19130EA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAlarmEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmEventResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlarmEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlarmEventResponseBody) SetRequestId(v string) *UpdateAlarmEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlarmEventResponseBody) Validate() error {
	return dara.Validate(s)
}

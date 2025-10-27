// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteScheduleResponseBody
	GetRequestId() *string
}

type DeleteScheduleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScheduleResponseBody) SetRequestId(v string) *DeleteScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}

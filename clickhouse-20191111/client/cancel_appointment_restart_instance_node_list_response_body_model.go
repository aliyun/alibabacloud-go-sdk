// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAppointmentRestartInstanceNodeListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelAppointmentRestartInstanceNodeListResponseBody
	GetRequestId() *string
}

type CancelAppointmentRestartInstanceNodeListResponseBody struct {
	// example:
	//
	// 1F488A93-83FD-540F-9B67-0333AF64E6A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelAppointmentRestartInstanceNodeListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelAppointmentRestartInstanceNodeListResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAppointmentRestartInstanceNodeListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelAppointmentRestartInstanceNodeListResponseBody) SetRequestId(v string) *CancelAppointmentRestartInstanceNodeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAppointmentRestartInstanceNodeListResponseBody) Validate() error {
	return dara.Validate(s)
}

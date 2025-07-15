// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtcMPUEventSubResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRtcMPUEventSubResponseBody
	GetRequestId() *string
}

type UpdateRtcMPUEventSubResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRtcMPUEventSubResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcMPUEventSubResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRtcMPUEventSubResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRtcMPUEventSubResponseBody) SetRequestId(v string) *UpdateRtcMPUEventSubResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRtcMPUEventSubResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRtcMPUEventSubResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRtcMPUEventSubResponseBody
	GetRequestId() *string
}

type DeleteRtcMPUEventSubResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRtcMPUEventSubResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRtcMPUEventSubResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRtcMPUEventSubResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRtcMPUEventSubResponseBody) SetRequestId(v string) *DeleteRtcMPUEventSubResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRtcMPUEventSubResponseBody) Validate() error {
	return dara.Validate(s)
}

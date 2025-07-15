// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRtcMPUEventSubResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRtcMPUEventSubResponseBody
	GetRequestId() *string
	SetSubId(v string) *CreateRtcMPUEventSubResponseBody
	GetSubId() *string
}

type CreateRtcMPUEventSubResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the subscription.
	//
	// example:
	//
	// Sub-******9799B2C4500******
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
}

func (s CreateRtcMPUEventSubResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRtcMPUEventSubResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRtcMPUEventSubResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRtcMPUEventSubResponseBody) GetSubId() *string {
	return s.SubId
}

func (s *CreateRtcMPUEventSubResponseBody) SetRequestId(v string) *CreateRtcMPUEventSubResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRtcMPUEventSubResponseBody) SetSubId(v string) *CreateRtcMPUEventSubResponseBody {
	s.SubId = &v
	return s
}

func (s *CreateRtcMPUEventSubResponseBody) Validate() error {
	return dara.Validate(s)
}

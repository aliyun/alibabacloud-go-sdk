// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRtcSendSipDtmfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RtcSendSipDtmfResponseBody
	GetRequestId() *string
}

type RtcSendSipDtmfResponseBody struct {
	// The request ID. This is the unique identifier returned for each API call.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RtcSendSipDtmfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RtcSendSipDtmfResponseBody) GoString() string {
	return s.String()
}

func (s *RtcSendSipDtmfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RtcSendSipDtmfResponseBody) SetRequestId(v string) *RtcSendSipDtmfResponseBody {
	s.RequestId = &v
	return s
}

func (s *RtcSendSipDtmfResponseBody) Validate() error {
	return dara.Validate(s)
}

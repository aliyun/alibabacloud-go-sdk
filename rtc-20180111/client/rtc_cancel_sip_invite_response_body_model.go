// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRtcCancelSipInviteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RtcCancelSipInviteResponseBody
	GetRequestId() *string
}

type RtcCancelSipInviteResponseBody struct {
	// example:
	//
	// EB3C3C75-74D5-5F01-9F3C-9974261BAED4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RtcCancelSipInviteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RtcCancelSipInviteResponseBody) GoString() string {
	return s.String()
}

func (s *RtcCancelSipInviteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RtcCancelSipInviteResponseBody) SetRequestId(v string) *RtcCancelSipInviteResponseBody {
	s.RequestId = &v
	return s
}

func (s *RtcCancelSipInviteResponseBody) Validate() error {
	return dara.Validate(s)
}

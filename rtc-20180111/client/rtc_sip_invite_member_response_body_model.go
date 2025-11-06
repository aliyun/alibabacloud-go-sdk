// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRtcSipInviteMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RtcSipInviteMemberResponseBody
	GetRequestId() *string
}

type RtcSipInviteMemberResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RtcSipInviteMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RtcSipInviteMemberResponseBody) GoString() string {
	return s.String()
}

func (s *RtcSipInviteMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RtcSipInviteMemberResponseBody) SetRequestId(v string) *RtcSipInviteMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *RtcSipInviteMemberResponseBody) Validate() error {
	return dara.Validate(s)
}

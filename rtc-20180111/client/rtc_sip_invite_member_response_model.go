// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRtcSipInviteMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RtcSipInviteMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RtcSipInviteMemberResponse
	GetStatusCode() *int32
	SetBody(v *RtcSipInviteMemberResponseBody) *RtcSipInviteMemberResponse
	GetBody() *RtcSipInviteMemberResponseBody
}

type RtcSipInviteMemberResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RtcSipInviteMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RtcSipInviteMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s RtcSipInviteMemberResponse) GoString() string {
	return s.String()
}

func (s *RtcSipInviteMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RtcSipInviteMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RtcSipInviteMemberResponse) GetBody() *RtcSipInviteMemberResponseBody {
	return s.Body
}

func (s *RtcSipInviteMemberResponse) SetHeaders(v map[string]*string) *RtcSipInviteMemberResponse {
	s.Headers = v
	return s
}

func (s *RtcSipInviteMemberResponse) SetStatusCode(v int32) *RtcSipInviteMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *RtcSipInviteMemberResponse) SetBody(v *RtcSipInviteMemberResponseBody) *RtcSipInviteMemberResponse {
	s.Body = v
	return s
}

func (s *RtcSipInviteMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

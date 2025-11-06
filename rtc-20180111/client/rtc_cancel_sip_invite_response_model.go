// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRtcCancelSipInviteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RtcCancelSipInviteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RtcCancelSipInviteResponse
	GetStatusCode() *int32
	SetBody(v *RtcCancelSipInviteResponseBody) *RtcCancelSipInviteResponse
	GetBody() *RtcCancelSipInviteResponseBody
}

type RtcCancelSipInviteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RtcCancelSipInviteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RtcCancelSipInviteResponse) String() string {
	return dara.Prettify(s)
}

func (s RtcCancelSipInviteResponse) GoString() string {
	return s.String()
}

func (s *RtcCancelSipInviteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RtcCancelSipInviteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RtcCancelSipInviteResponse) GetBody() *RtcCancelSipInviteResponseBody {
	return s.Body
}

func (s *RtcCancelSipInviteResponse) SetHeaders(v map[string]*string) *RtcCancelSipInviteResponse {
	s.Headers = v
	return s
}

func (s *RtcCancelSipInviteResponse) SetStatusCode(v int32) *RtcCancelSipInviteResponse {
	s.StatusCode = &v
	return s
}

func (s *RtcCancelSipInviteResponse) SetBody(v *RtcCancelSipInviteResponseBody) *RtcCancelSipInviteResponse {
	s.Body = v
	return s
}

func (s *RtcCancelSipInviteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

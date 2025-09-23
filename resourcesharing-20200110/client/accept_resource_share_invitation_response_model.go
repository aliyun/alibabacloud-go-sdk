// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptResourceShareInvitationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AcceptResourceShareInvitationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AcceptResourceShareInvitationResponse
	GetStatusCode() *int32
	SetBody(v *AcceptResourceShareInvitationResponseBody) *AcceptResourceShareInvitationResponse
	GetBody() *AcceptResourceShareInvitationResponseBody
}

type AcceptResourceShareInvitationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptResourceShareInvitationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptResourceShareInvitationResponse) String() string {
	return dara.Prettify(s)
}

func (s AcceptResourceShareInvitationResponse) GoString() string {
	return s.String()
}

func (s *AcceptResourceShareInvitationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AcceptResourceShareInvitationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AcceptResourceShareInvitationResponse) GetBody() *AcceptResourceShareInvitationResponseBody {
	return s.Body
}

func (s *AcceptResourceShareInvitationResponse) SetHeaders(v map[string]*string) *AcceptResourceShareInvitationResponse {
	s.Headers = v
	return s
}

func (s *AcceptResourceShareInvitationResponse) SetStatusCode(v int32) *AcceptResourceShareInvitationResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptResourceShareInvitationResponse) SetBody(v *AcceptResourceShareInvitationResponseBody) *AcceptResourceShareInvitationResponse {
	s.Body = v
	return s
}

func (s *AcceptResourceShareInvitationResponse) Validate() error {
	return dara.Validate(s)
}

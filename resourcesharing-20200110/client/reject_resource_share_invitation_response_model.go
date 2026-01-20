// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectResourceShareInvitationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RejectResourceShareInvitationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RejectResourceShareInvitationResponse
	GetStatusCode() *int32
	SetBody(v *RejectResourceShareInvitationResponseBody) *RejectResourceShareInvitationResponse
	GetBody() *RejectResourceShareInvitationResponseBody
}

type RejectResourceShareInvitationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RejectResourceShareInvitationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RejectResourceShareInvitationResponse) String() string {
	return dara.Prettify(s)
}

func (s RejectResourceShareInvitationResponse) GoString() string {
	return s.String()
}

func (s *RejectResourceShareInvitationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RejectResourceShareInvitationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RejectResourceShareInvitationResponse) GetBody() *RejectResourceShareInvitationResponseBody {
	return s.Body
}

func (s *RejectResourceShareInvitationResponse) SetHeaders(v map[string]*string) *RejectResourceShareInvitationResponse {
	s.Headers = v
	return s
}

func (s *RejectResourceShareInvitationResponse) SetStatusCode(v int32) *RejectResourceShareInvitationResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectResourceShareInvitationResponse) SetBody(v *RejectResourceShareInvitationResponseBody) *RejectResourceShareInvitationResponse {
	s.Body = v
	return s
}

func (s *RejectResourceShareInvitationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

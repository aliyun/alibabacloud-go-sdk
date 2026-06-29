// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeTokenPlanInviteLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeTokenPlanInviteLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeTokenPlanInviteLinkResponse
	GetStatusCode() *int32
	SetBody(v *RevokeTokenPlanInviteLinkResponseBody) *RevokeTokenPlanInviteLinkResponse
	GetBody() *RevokeTokenPlanInviteLinkResponseBody
}

type RevokeTokenPlanInviteLinkResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeTokenPlanInviteLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeTokenPlanInviteLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeTokenPlanInviteLinkResponse) GoString() string {
	return s.String()
}

func (s *RevokeTokenPlanInviteLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeTokenPlanInviteLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeTokenPlanInviteLinkResponse) GetBody() *RevokeTokenPlanInviteLinkResponseBody {
	return s.Body
}

func (s *RevokeTokenPlanInviteLinkResponse) SetHeaders(v map[string]*string) *RevokeTokenPlanInviteLinkResponse {
	s.Headers = v
	return s
}

func (s *RevokeTokenPlanInviteLinkResponse) SetStatusCode(v int32) *RevokeTokenPlanInviteLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeTokenPlanInviteLinkResponse) SetBody(v *RevokeTokenPlanInviteLinkResponseBody) *RevokeTokenPlanInviteLinkResponse {
	s.Body = v
	return s
}

func (s *RevokeTokenPlanInviteLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

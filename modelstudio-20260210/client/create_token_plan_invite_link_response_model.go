// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTokenPlanInviteLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTokenPlanInviteLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTokenPlanInviteLinkResponse
	GetStatusCode() *int32
	SetBody(v *CreateTokenPlanInviteLinkResponseBody) *CreateTokenPlanInviteLinkResponse
	GetBody() *CreateTokenPlanInviteLinkResponseBody
}

type CreateTokenPlanInviteLinkResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTokenPlanInviteLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTokenPlanInviteLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenPlanInviteLinkResponse) GoString() string {
	return s.String()
}

func (s *CreateTokenPlanInviteLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTokenPlanInviteLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTokenPlanInviteLinkResponse) GetBody() *CreateTokenPlanInviteLinkResponseBody {
	return s.Body
}

func (s *CreateTokenPlanInviteLinkResponse) SetHeaders(v map[string]*string) *CreateTokenPlanInviteLinkResponse {
	s.Headers = v
	return s
}

func (s *CreateTokenPlanInviteLinkResponse) SetStatusCode(v int32) *CreateTokenPlanInviteLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTokenPlanInviteLinkResponse) SetBody(v *CreateTokenPlanInviteLinkResponseBody) *CreateTokenPlanInviteLinkResponse {
	s.Body = v
	return s
}

func (s *CreateTokenPlanInviteLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

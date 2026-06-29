// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenPlanInviteLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTokenPlanInviteLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTokenPlanInviteLinkResponse
	GetStatusCode() *int32
	SetBody(v *GetTokenPlanInviteLinkResponseBody) *GetTokenPlanInviteLinkResponse
	GetBody() *GetTokenPlanInviteLinkResponseBody
}

type GetTokenPlanInviteLinkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTokenPlanInviteLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTokenPlanInviteLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTokenPlanInviteLinkResponse) GoString() string {
	return s.String()
}

func (s *GetTokenPlanInviteLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTokenPlanInviteLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTokenPlanInviteLinkResponse) GetBody() *GetTokenPlanInviteLinkResponseBody {
	return s.Body
}

func (s *GetTokenPlanInviteLinkResponse) SetHeaders(v map[string]*string) *GetTokenPlanInviteLinkResponse {
	s.Headers = v
	return s
}

func (s *GetTokenPlanInviteLinkResponse) SetStatusCode(v int32) *GetTokenPlanInviteLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenPlanInviteLinkResponse) SetBody(v *GetTokenPlanInviteLinkResponseBody) *GetTokenPlanInviteLinkResponse {
	s.Body = v
	return s
}

func (s *GetTokenPlanInviteLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

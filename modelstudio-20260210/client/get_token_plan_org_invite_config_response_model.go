// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenPlanOrgInviteConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTokenPlanOrgInviteConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTokenPlanOrgInviteConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetTokenPlanOrgInviteConfigResponseBody) *GetTokenPlanOrgInviteConfigResponse
	GetBody() *GetTokenPlanOrgInviteConfigResponseBody
}

type GetTokenPlanOrgInviteConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTokenPlanOrgInviteConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTokenPlanOrgInviteConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTokenPlanOrgInviteConfigResponse) GoString() string {
	return s.String()
}

func (s *GetTokenPlanOrgInviteConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTokenPlanOrgInviteConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTokenPlanOrgInviteConfigResponse) GetBody() *GetTokenPlanOrgInviteConfigResponseBody {
	return s.Body
}

func (s *GetTokenPlanOrgInviteConfigResponse) SetHeaders(v map[string]*string) *GetTokenPlanOrgInviteConfigResponse {
	s.Headers = v
	return s
}

func (s *GetTokenPlanOrgInviteConfigResponse) SetStatusCode(v int32) *GetTokenPlanOrgInviteConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenPlanOrgInviteConfigResponse) SetBody(v *GetTokenPlanOrgInviteConfigResponseBody) *GetTokenPlanOrgInviteConfigResponse {
	s.Body = v
	return s
}

func (s *GetTokenPlanOrgInviteConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

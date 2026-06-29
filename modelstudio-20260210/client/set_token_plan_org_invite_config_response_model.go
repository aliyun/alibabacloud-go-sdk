// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTokenPlanOrgInviteConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetTokenPlanOrgInviteConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetTokenPlanOrgInviteConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetTokenPlanOrgInviteConfigResponseBody) *SetTokenPlanOrgInviteConfigResponse
	GetBody() *SetTokenPlanOrgInviteConfigResponseBody
}

type SetTokenPlanOrgInviteConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetTokenPlanOrgInviteConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetTokenPlanOrgInviteConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetTokenPlanOrgInviteConfigResponse) GoString() string {
	return s.String()
}

func (s *SetTokenPlanOrgInviteConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetTokenPlanOrgInviteConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetTokenPlanOrgInviteConfigResponse) GetBody() *SetTokenPlanOrgInviteConfigResponseBody {
	return s.Body
}

func (s *SetTokenPlanOrgInviteConfigResponse) SetHeaders(v map[string]*string) *SetTokenPlanOrgInviteConfigResponse {
	s.Headers = v
	return s
}

func (s *SetTokenPlanOrgInviteConfigResponse) SetStatusCode(v int32) *SetTokenPlanOrgInviteConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetTokenPlanOrgInviteConfigResponse) SetBody(v *SetTokenPlanOrgInviteConfigResponseBody) *SetTokenPlanOrgInviteConfigResponse {
	s.Body = v
	return s
}

func (s *SetTokenPlanOrgInviteConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

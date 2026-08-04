// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterResetMemberAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterResetMemberAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterResetMemberAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterResetMemberAuthorizationResponseBody) *ModelRouterResetMemberAuthorizationResponse
	GetBody() *ModelRouterResetMemberAuthorizationResponseBody
}

type ModelRouterResetMemberAuthorizationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterResetMemberAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterResetMemberAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterResetMemberAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterResetMemberAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterResetMemberAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterResetMemberAuthorizationResponse) GetBody() *ModelRouterResetMemberAuthorizationResponseBody {
	return s.Body
}

func (s *ModelRouterResetMemberAuthorizationResponse) SetHeaders(v map[string]*string) *ModelRouterResetMemberAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterResetMemberAuthorizationResponse) SetStatusCode(v int32) *ModelRouterResetMemberAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterResetMemberAuthorizationResponse) SetBody(v *ModelRouterResetMemberAuthorizationResponseBody) *ModelRouterResetMemberAuthorizationResponse {
	s.Body = v
	return s
}

func (s *ModelRouterResetMemberAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

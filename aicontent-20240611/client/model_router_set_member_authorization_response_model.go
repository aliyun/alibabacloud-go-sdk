// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterSetMemberAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterSetMemberAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterSetMemberAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterSetMemberAuthorizationResponseBody) *ModelRouterSetMemberAuthorizationResponse
	GetBody() *ModelRouterSetMemberAuthorizationResponseBody
}

type ModelRouterSetMemberAuthorizationResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterSetMemberAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterSetMemberAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterSetMemberAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterSetMemberAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterSetMemberAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterSetMemberAuthorizationResponse) GetBody() *ModelRouterSetMemberAuthorizationResponseBody {
	return s.Body
}

func (s *ModelRouterSetMemberAuthorizationResponse) SetHeaders(v map[string]*string) *ModelRouterSetMemberAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterSetMemberAuthorizationResponse) SetStatusCode(v int32) *ModelRouterSetMemberAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterSetMemberAuthorizationResponse) SetBody(v *ModelRouterSetMemberAuthorizationResponseBody) *ModelRouterSetMemberAuthorizationResponse {
	s.Body = v
	return s
}

func (s *ModelRouterSetMemberAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

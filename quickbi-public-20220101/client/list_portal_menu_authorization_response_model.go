// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPortalMenuAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPortalMenuAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPortalMenuAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *ListPortalMenuAuthorizationResponseBody) *ListPortalMenuAuthorizationResponse
	GetBody() *ListPortalMenuAuthorizationResponseBody
}

type ListPortalMenuAuthorizationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPortalMenuAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPortalMenuAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPortalMenuAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *ListPortalMenuAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPortalMenuAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPortalMenuAuthorizationResponse) GetBody() *ListPortalMenuAuthorizationResponseBody {
	return s.Body
}

func (s *ListPortalMenuAuthorizationResponse) SetHeaders(v map[string]*string) *ListPortalMenuAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *ListPortalMenuAuthorizationResponse) SetStatusCode(v int32) *ListPortalMenuAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPortalMenuAuthorizationResponse) SetBody(v *ListPortalMenuAuthorizationResponseBody) *ListPortalMenuAuthorizationResponse {
	s.Body = v
	return s
}

func (s *ListPortalMenuAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

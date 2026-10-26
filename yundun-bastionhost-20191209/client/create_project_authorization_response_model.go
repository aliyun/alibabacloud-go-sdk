// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProjectAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProjectAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *CreateProjectAuthorizationResponseBody) *CreateProjectAuthorizationResponse
	GetBody() *CreateProjectAuthorizationResponseBody
}

type CreateProjectAuthorizationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProjectAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProjectAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProjectAuthorizationResponse) GetBody() *CreateProjectAuthorizationResponseBody {
	return s.Body
}

func (s *CreateProjectAuthorizationResponse) SetHeaders(v map[string]*string) *CreateProjectAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectAuthorizationResponse) SetStatusCode(v int32) *CreateProjectAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectAuthorizationResponse) SetBody(v *CreateProjectAuthorizationResponseBody) *CreateProjectAuthorizationResponse {
	s.Body = v
	return s
}

func (s *CreateProjectAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

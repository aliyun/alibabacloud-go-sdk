// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAuthorizationResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAuthorizationResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAuthorizationResourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateAuthorizationResourceResponseBody) *CreateAuthorizationResourceResponse
	GetBody() *CreateAuthorizationResourceResponseBody
}

type CreateAuthorizationResourceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAuthorizationResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAuthorizationResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAuthorizationResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAuthorizationResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAuthorizationResourceResponse) GetBody() *CreateAuthorizationResourceResponseBody {
	return s.Body
}

func (s *CreateAuthorizationResourceResponse) SetHeaders(v map[string]*string) *CreateAuthorizationResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateAuthorizationResourceResponse) SetStatusCode(v int32) *CreateAuthorizationResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAuthorizationResourceResponse) SetBody(v *CreateAuthorizationResourceResponseBody) *CreateAuthorizationResourceResponse {
	s.Body = v
	return s
}

func (s *CreateAuthorizationResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterCreateUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterCreateUserResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterCreateUserResponseBody) *ModelRouterCreateUserResponse
	GetBody() *ModelRouterCreateUserResponseBody
}

type ModelRouterCreateUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterCreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterCreateUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateUserResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterCreateUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterCreateUserResponse) GetBody() *ModelRouterCreateUserResponseBody {
	return s.Body
}

func (s *ModelRouterCreateUserResponse) SetHeaders(v map[string]*string) *ModelRouterCreateUserResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterCreateUserResponse) SetStatusCode(v int32) *ModelRouterCreateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterCreateUserResponse) SetBody(v *ModelRouterCreateUserResponseBody) *ModelRouterCreateUserResponse {
	s.Body = v
	return s
}

func (s *ModelRouterCreateUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

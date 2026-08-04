// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterUpdateUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterUpdateUserResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterUpdateUserResponseBody) *ModelRouterUpdateUserResponse
	GetBody() *ModelRouterUpdateUserResponseBody
}

type ModelRouterUpdateUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterUpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterUpdateUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateUserResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterUpdateUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterUpdateUserResponse) GetBody() *ModelRouterUpdateUserResponseBody {
	return s.Body
}

func (s *ModelRouterUpdateUserResponse) SetHeaders(v map[string]*string) *ModelRouterUpdateUserResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterUpdateUserResponse) SetStatusCode(v int32) *ModelRouterUpdateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterUpdateUserResponse) SetBody(v *ModelRouterUpdateUserResponseBody) *ModelRouterUpdateUserResponse {
	s.Body = v
	return s
}

func (s *ModelRouterUpdateUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

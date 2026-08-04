// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterDeleteUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterDeleteUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterDeleteUserResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterDeleteUserResponseBody) *ModelRouterDeleteUserResponse
	GetBody() *ModelRouterDeleteUserResponseBody
}

type ModelRouterDeleteUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterDeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterDeleteUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterDeleteUserResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterDeleteUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterDeleteUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterDeleteUserResponse) GetBody() *ModelRouterDeleteUserResponseBody {
	return s.Body
}

func (s *ModelRouterDeleteUserResponse) SetHeaders(v map[string]*string) *ModelRouterDeleteUserResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterDeleteUserResponse) SetStatusCode(v int32) *ModelRouterDeleteUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterDeleteUserResponse) SetBody(v *ModelRouterDeleteUserResponseBody) *ModelRouterDeleteUserResponse {
	s.Body = v
	return s
}

func (s *ModelRouterDeleteUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

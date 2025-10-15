// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationAuthorizationTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationAuthorizationTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationAuthorizationTypeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationAuthorizationTypeResponseBody) *UpdateApplicationAuthorizationTypeResponse
	GetBody() *UpdateApplicationAuthorizationTypeResponseBody
}

type UpdateApplicationAuthorizationTypeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationAuthorizationTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationAuthorizationTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationAuthorizationTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationAuthorizationTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationAuthorizationTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationAuthorizationTypeResponse) GetBody() *UpdateApplicationAuthorizationTypeResponseBody {
	return s.Body
}

func (s *UpdateApplicationAuthorizationTypeResponse) SetHeaders(v map[string]*string) *UpdateApplicationAuthorizationTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationAuthorizationTypeResponse) SetStatusCode(v int32) *UpdateApplicationAuthorizationTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationAuthorizationTypeResponse) SetBody(v *UpdateApplicationAuthorizationTypeResponseBody) *UpdateApplicationAuthorizationTypeResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationAuthorizationTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNumberBindingsOfInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNumberBindingsOfInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNumberBindingsOfInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNumberBindingsOfInstanceResponseBody) *UpdateNumberBindingsOfInstanceResponse
	GetBody() *UpdateNumberBindingsOfInstanceResponseBody
}

type UpdateNumberBindingsOfInstanceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNumberBindingsOfInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNumberBindingsOfInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNumberBindingsOfInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateNumberBindingsOfInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNumberBindingsOfInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNumberBindingsOfInstanceResponse) GetBody() *UpdateNumberBindingsOfInstanceResponseBody {
	return s.Body
}

func (s *UpdateNumberBindingsOfInstanceResponse) SetHeaders(v map[string]*string) *UpdateNumberBindingsOfInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateNumberBindingsOfInstanceResponse) SetStatusCode(v int32) *UpdateNumberBindingsOfInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNumberBindingsOfInstanceResponse) SetBody(v *UpdateNumberBindingsOfInstanceResponseBody) *UpdateNumberBindingsOfInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateNumberBindingsOfInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

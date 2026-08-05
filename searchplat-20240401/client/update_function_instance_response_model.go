// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFunctionInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFunctionInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFunctionInstanceResponseBody) *UpdateFunctionInstanceResponse
	GetBody() *UpdateFunctionInstanceResponseBody
}

type UpdateFunctionInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFunctionInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFunctionInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFunctionInstanceResponse) GetBody() *UpdateFunctionInstanceResponseBody {
	return s.Body
}

func (s *UpdateFunctionInstanceResponse) SetHeaders(v map[string]*string) *UpdateFunctionInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionInstanceResponse) SetStatusCode(v int32) *UpdateFunctionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFunctionInstanceResponse) SetBody(v *UpdateFunctionInstanceResponseBody) *UpdateFunctionInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateFunctionInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

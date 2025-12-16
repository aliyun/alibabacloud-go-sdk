// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionDefaultInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFunctionDefaultInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFunctionDefaultInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFunctionDefaultInstanceResponseBody) *UpdateFunctionDefaultInstanceResponse
	GetBody() *UpdateFunctionDefaultInstanceResponseBody
}

type UpdateFunctionDefaultInstanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFunctionDefaultInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFunctionDefaultInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionDefaultInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionDefaultInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFunctionDefaultInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFunctionDefaultInstanceResponse) GetBody() *UpdateFunctionDefaultInstanceResponseBody {
	return s.Body
}

func (s *UpdateFunctionDefaultInstanceResponse) SetHeaders(v map[string]*string) *UpdateFunctionDefaultInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponse) SetStatusCode(v int32) *UpdateFunctionDefaultInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponse) SetBody(v *UpdateFunctionDefaultInstanceResponseBody) *UpdateFunctionDefaultInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

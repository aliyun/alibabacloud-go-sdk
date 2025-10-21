// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVariableResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVariableResponseBody) *UpdateVariableResponse
	GetBody() *UpdateVariableResponseBody
}

type UpdateVariableResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVariableResponse) GoString() string {
	return s.String()
}

func (s *UpdateVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVariableResponse) GetBody() *UpdateVariableResponseBody {
	return s.Body
}

func (s *UpdateVariableResponse) SetHeaders(v map[string]*string) *UpdateVariableResponse {
	s.Headers = v
	return s
}

func (s *UpdateVariableResponse) SetStatusCode(v int32) *UpdateVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVariableResponse) SetBody(v *UpdateVariableResponseBody) *UpdateVariableResponse {
	s.Body = v
	return s
}

func (s *UpdateVariableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

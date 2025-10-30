// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQueryVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateQueryVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateQueryVariableResponse
	GetStatusCode() *int32
	SetBody(v *UpdateQueryVariableResponseBody) *UpdateQueryVariableResponse
	GetBody() *UpdateQueryVariableResponseBody
}

type UpdateQueryVariableResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQueryVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQueryVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateQueryVariableResponse) GoString() string {
	return s.String()
}

func (s *UpdateQueryVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateQueryVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateQueryVariableResponse) GetBody() *UpdateQueryVariableResponseBody {
	return s.Body
}

func (s *UpdateQueryVariableResponse) SetHeaders(v map[string]*string) *UpdateQueryVariableResponse {
	s.Headers = v
	return s
}

func (s *UpdateQueryVariableResponse) SetStatusCode(v int32) *UpdateQueryVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQueryVariableResponse) SetBody(v *UpdateQueryVariableResponseBody) *UpdateQueryVariableResponse {
	s.Body = v
	return s
}

func (s *UpdateQueryVariableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

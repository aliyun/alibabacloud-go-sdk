// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVariableResponse
	GetStatusCode() *int32
	SetBody(v *CreateVariableResponseBody) *CreateVariableResponse
	GetBody() *CreateVariableResponseBody
}

type CreateVariableResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVariableResponse) GoString() string {
	return s.String()
}

func (s *CreateVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVariableResponse) GetBody() *CreateVariableResponseBody {
	return s.Body
}

func (s *CreateVariableResponse) SetHeaders(v map[string]*string) *CreateVariableResponse {
	s.Headers = v
	return s
}

func (s *CreateVariableResponse) SetStatusCode(v int32) *CreateVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVariableResponse) SetBody(v *CreateVariableResponseBody) *CreateVariableResponse {
	s.Body = v
	return s
}

func (s *CreateVariableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

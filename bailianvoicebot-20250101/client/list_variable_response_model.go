// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVariableResponse
	GetStatusCode() *int32
	SetBody(v *ListVariableResponseBody) *ListVariableResponse
	GetBody() *ListVariableResponseBody
}

type ListVariableResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVariableResponse) GoString() string {
	return s.String()
}

func (s *ListVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVariableResponse) GetBody() *ListVariableResponseBody {
	return s.Body
}

func (s *ListVariableResponse) SetHeaders(v map[string]*string) *ListVariableResponse {
	s.Headers = v
	return s
}

func (s *ListVariableResponse) SetStatusCode(v int32) *ListVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVariableResponse) SetBody(v *ListVariableResponseBody) *ListVariableResponse {
	s.Body = v
	return s
}

func (s *ListVariableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

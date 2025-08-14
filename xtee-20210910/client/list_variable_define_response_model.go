// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVariableDefineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVariableDefineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVariableDefineResponse
	GetStatusCode() *int32
	SetBody(v *ListVariableDefineResponseBody) *ListVariableDefineResponse
	GetBody() *ListVariableDefineResponseBody
}

type ListVariableDefineResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVariableDefineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVariableDefineResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVariableDefineResponse) GoString() string {
	return s.String()
}

func (s *ListVariableDefineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVariableDefineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVariableDefineResponse) GetBody() *ListVariableDefineResponseBody {
	return s.Body
}

func (s *ListVariableDefineResponse) SetHeaders(v map[string]*string) *ListVariableDefineResponse {
	s.Headers = v
	return s
}

func (s *ListVariableDefineResponse) SetStatusCode(v int32) *ListVariableDefineResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVariableDefineResponse) SetBody(v *ListVariableDefineResponseBody) *ListVariableDefineResponse {
	s.Body = v
	return s
}

func (s *ListVariableDefineResponse) Validate() error {
	return dara.Validate(s)
}

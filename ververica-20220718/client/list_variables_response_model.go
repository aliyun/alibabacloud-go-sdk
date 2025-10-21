// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVariablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVariablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVariablesResponse
	GetStatusCode() *int32
	SetBody(v *ListVariablesResponseBody) *ListVariablesResponse
	GetBody() *ListVariablesResponseBody
}

type ListVariablesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVariablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVariablesResponse) GoString() string {
	return s.String()
}

func (s *ListVariablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVariablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVariablesResponse) GetBody() *ListVariablesResponseBody {
	return s.Body
}

func (s *ListVariablesResponse) SetHeaders(v map[string]*string) *ListVariablesResponse {
	s.Headers = v
	return s
}

func (s *ListVariablesResponse) SetStatusCode(v int32) *ListVariablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVariablesResponse) SetBody(v *ListVariablesResponseBody) *ListVariablesResponse {
	s.Body = v
	return s
}

func (s *ListVariablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

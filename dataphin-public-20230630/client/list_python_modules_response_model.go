// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPythonModulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPythonModulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPythonModulesResponse
	GetStatusCode() *int32
	SetBody(v *ListPythonModulesResponseBody) *ListPythonModulesResponse
	GetBody() *ListPythonModulesResponseBody
}

type ListPythonModulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPythonModulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPythonModulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPythonModulesResponse) GoString() string {
	return s.String()
}

func (s *ListPythonModulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPythonModulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPythonModulesResponse) GetBody() *ListPythonModulesResponseBody {
	return s.Body
}

func (s *ListPythonModulesResponse) SetHeaders(v map[string]*string) *ListPythonModulesResponse {
	s.Headers = v
	return s
}

func (s *ListPythonModulesResponse) SetStatusCode(v int32) *ListPythonModulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPythonModulesResponse) SetBody(v *ListPythonModulesResponseBody) *ListPythonModulesResponse {
	s.Body = v
	return s
}

func (s *ListPythonModulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

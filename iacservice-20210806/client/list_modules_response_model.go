// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModulesResponse
	GetStatusCode() *int32
	SetBody(v *ListModulesResponseBody) *ListModulesResponse
	GetBody() *ListModulesResponseBody
}

type ListModulesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModulesResponse) GoString() string {
	return s.String()
}

func (s *ListModulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModulesResponse) GetBody() *ListModulesResponseBody {
	return s.Body
}

func (s *ListModulesResponse) SetHeaders(v map[string]*string) *ListModulesResponse {
	s.Headers = v
	return s
}

func (s *ListModulesResponse) SetStatusCode(v int32) *ListModulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModulesResponse) SetBody(v *ListModulesResponseBody) *ListModulesResponse {
	s.Body = v
	return s
}

func (s *ListModulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSandboxesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSandboxesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSandboxesResponse
	GetStatusCode() *int32
	SetBody(v *ListSandboxesResult) *ListSandboxesResponse
	GetBody() *ListSandboxesResult
}

type ListSandboxesResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSandboxesResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSandboxesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSandboxesResponse) GoString() string {
	return s.String()
}

func (s *ListSandboxesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSandboxesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSandboxesResponse) GetBody() *ListSandboxesResult {
	return s.Body
}

func (s *ListSandboxesResponse) SetHeaders(v map[string]*string) *ListSandboxesResponse {
	s.Headers = v
	return s
}

func (s *ListSandboxesResponse) SetStatusCode(v int32) *ListSandboxesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSandboxesResponse) SetBody(v *ListSandboxesResult) *ListSandboxesResponse {
	s.Body = v
	return s
}

func (s *ListSandboxesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

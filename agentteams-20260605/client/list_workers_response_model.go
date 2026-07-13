// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkersResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkersResponseBody) *ListWorkersResponse
	GetBody() *ListWorkersResponseBody
}

type ListWorkersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkersResponse) GoString() string {
	return s.String()
}

func (s *ListWorkersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkersResponse) GetBody() *ListWorkersResponseBody {
	return s.Body
}

func (s *ListWorkersResponse) SetHeaders(v map[string]*string) *ListWorkersResponse {
	s.Headers = v
	return s
}

func (s *ListWorkersResponse) SetStatusCode(v int32) *ListWorkersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkersResponse) SetBody(v *ListWorkersResponseBody) *ListWorkersResponse {
	s.Body = v
	return s
}

func (s *ListWorkersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

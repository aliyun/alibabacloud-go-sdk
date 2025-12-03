// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkitemTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkitemTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkitemTimeResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkitemTimeResponseBody) *ListWorkitemTimeResponse
	GetBody() *ListWorkitemTimeResponseBody
}

type ListWorkitemTimeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkitemTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkitemTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkitemTimeResponse) GoString() string {
	return s.String()
}

func (s *ListWorkitemTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkitemTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkitemTimeResponse) GetBody() *ListWorkitemTimeResponseBody {
	return s.Body
}

func (s *ListWorkitemTimeResponse) SetHeaders(v map[string]*string) *ListWorkitemTimeResponse {
	s.Headers = v
	return s
}

func (s *ListWorkitemTimeResponse) SetStatusCode(v int32) *ListWorkitemTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkitemTimeResponse) SetBody(v *ListWorkitemTimeResponseBody) *ListWorkitemTimeResponse {
	s.Body = v
	return s
}

func (s *ListWorkitemTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

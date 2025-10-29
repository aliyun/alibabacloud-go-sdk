// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPartitionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPartitionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPartitionsResponse
	GetStatusCode() *int32
	SetBody(v *ListPartitionsResponseBody) *ListPartitionsResponse
	GetBody() *ListPartitionsResponseBody
}

type ListPartitionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPartitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPartitionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPartitionsResponse) GoString() string {
	return s.String()
}

func (s *ListPartitionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPartitionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPartitionsResponse) GetBody() *ListPartitionsResponseBody {
	return s.Body
}

func (s *ListPartitionsResponse) SetHeaders(v map[string]*string) *ListPartitionsResponse {
	s.Headers = v
	return s
}

func (s *ListPartitionsResponse) SetStatusCode(v int32) *ListPartitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPartitionsResponse) SetBody(v *ListPartitionsResponseBody) *ListPartitionsResponse {
	s.Body = v
	return s
}

func (s *ListPartitionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

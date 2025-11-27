// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAddonsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAddonsResponse
	GetStatusCode() *int32
	SetBody(v *ListAddonsResponseBody) *ListAddonsResponse
	GetBody() *ListAddonsResponseBody
}

type ListAddonsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAddonsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAddonsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponse) GoString() string {
	return s.String()
}

func (s *ListAddonsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAddonsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAddonsResponse) GetBody() *ListAddonsResponseBody {
	return s.Body
}

func (s *ListAddonsResponse) SetHeaders(v map[string]*string) *ListAddonsResponse {
	s.Headers = v
	return s
}

func (s *ListAddonsResponse) SetStatusCode(v int32) *ListAddonsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAddonsResponse) SetBody(v *ListAddonsResponseBody) *ListAddonsResponse {
	s.Body = v
	return s
}

func (s *ListAddonsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

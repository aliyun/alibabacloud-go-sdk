// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNumbersResponse
	GetStatusCode() *int32
	SetBody(v *ListNumbersResponseBody) *ListNumbersResponse
	GetBody() *ListNumbersResponseBody
}

type ListNumbersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNumbersResponse) GoString() string {
	return s.String()
}

func (s *ListNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNumbersResponse) GetBody() *ListNumbersResponseBody {
	return s.Body
}

func (s *ListNumbersResponse) SetHeaders(v map[string]*string) *ListNumbersResponse {
	s.Headers = v
	return s
}

func (s *ListNumbersResponse) SetStatusCode(v int32) *ListNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNumbersResponse) SetBody(v *ListNumbersResponseBody) *ListNumbersResponse {
	s.Body = v
	return s
}

func (s *ListNumbersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

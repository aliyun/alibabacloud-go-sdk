// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypeVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceTypeVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceTypeVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceTypeVersionsResponseBody) *ListResourceTypeVersionsResponse
	GetBody() *ListResourceTypeVersionsResponseBody
}

type ListResourceTypeVersionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceTypeVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceTypeVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypeVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceTypeVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceTypeVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceTypeVersionsResponse) GetBody() *ListResourceTypeVersionsResponseBody {
	return s.Body
}

func (s *ListResourceTypeVersionsResponse) SetHeaders(v map[string]*string) *ListResourceTypeVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceTypeVersionsResponse) SetStatusCode(v int32) *ListResourceTypeVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceTypeVersionsResponse) SetBody(v *ListResourceTypeVersionsResponseBody) *ListResourceTypeVersionsResponse {
	s.Body = v
	return s
}

func (s *ListResourceTypeVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

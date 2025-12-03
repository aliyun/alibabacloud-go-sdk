// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisByBackendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisByBackendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisByBackendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisByBackendResponseBody) *DescribeApisByBackendResponse
	GetBody() *DescribeApisByBackendResponseBody
}

type DescribeApisByBackendResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisByBackendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisByBackendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByBackendResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisByBackendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisByBackendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisByBackendResponse) GetBody() *DescribeApisByBackendResponseBody {
	return s.Body
}

func (s *DescribeApisByBackendResponse) SetHeaders(v map[string]*string) *DescribeApisByBackendResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisByBackendResponse) SetStatusCode(v int32) *DescribeApisByBackendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisByBackendResponse) SetBody(v *DescribeApisByBackendResponseBody) *DescribeApisByBackendResponse {
	s.Body = v
	return s
}

func (s *DescribeApisByBackendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

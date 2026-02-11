// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventDisposeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventDisposeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventDisposeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventDisposeResponseBody) *DescribeEventDisposeResponse
	GetBody() *DescribeEventDisposeResponseBody
}

type DescribeEventDisposeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventDisposeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventDisposeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDisposeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventDisposeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventDisposeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventDisposeResponse) GetBody() *DescribeEventDisposeResponseBody {
	return s.Body
}

func (s *DescribeEventDisposeResponse) SetHeaders(v map[string]*string) *DescribeEventDisposeResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventDisposeResponse) SetStatusCode(v int32) *DescribeEventDisposeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventDisposeResponse) SetBody(v *DescribeEventDisposeResponseBody) *DescribeEventDisposeResponse {
	s.Body = v
	return s
}

func (s *DescribeEventDisposeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

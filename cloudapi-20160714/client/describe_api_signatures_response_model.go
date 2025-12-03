// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiSignaturesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiSignaturesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiSignaturesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiSignaturesResponseBody) *DescribeApiSignaturesResponse
	GetBody() *DescribeApiSignaturesResponseBody
}

type DescribeApiSignaturesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiSignaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiSignaturesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiSignaturesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiSignaturesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiSignaturesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiSignaturesResponse) GetBody() *DescribeApiSignaturesResponseBody {
	return s.Body
}

func (s *DescribeApiSignaturesResponse) SetHeaders(v map[string]*string) *DescribeApiSignaturesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiSignaturesResponse) SetStatusCode(v int32) *DescribeApiSignaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiSignaturesResponse) SetBody(v *DescribeApiSignaturesResponseBody) *DescribeApiSignaturesResponse {
	s.Body = v
	return s
}

func (s *DescribeApiSignaturesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

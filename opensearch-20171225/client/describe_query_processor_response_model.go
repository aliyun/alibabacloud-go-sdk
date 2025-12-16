// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryProcessorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQueryProcessorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQueryProcessorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQueryProcessorResponseBody) *DescribeQueryProcessorResponse
	GetBody() *DescribeQueryProcessorResponseBody
}

type DescribeQueryProcessorResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQueryProcessorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryProcessorResponse) GoString() string {
	return s.String()
}

func (s *DescribeQueryProcessorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQueryProcessorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQueryProcessorResponse) GetBody() *DescribeQueryProcessorResponseBody {
	return s.Body
}

func (s *DescribeQueryProcessorResponse) SetHeaders(v map[string]*string) *DescribeQueryProcessorResponse {
	s.Headers = v
	return s
}

func (s *DescribeQueryProcessorResponse) SetStatusCode(v int32) *DescribeQueryProcessorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQueryProcessorResponse) SetBody(v *DescribeQueryProcessorResponseBody) *DescribeQueryProcessorResponse {
	s.Body = v
	return s
}

func (s *DescribeQueryProcessorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

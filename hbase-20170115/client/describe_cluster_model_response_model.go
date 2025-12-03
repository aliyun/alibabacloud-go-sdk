// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterModelResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterModelResponseBody) *DescribeClusterModelResponse
	GetBody() *DescribeClusterModelResponseBody
}

type DescribeClusterModelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterModelResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterModelResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterModelResponse) GetBody() *DescribeClusterModelResponseBody {
	return s.Body
}

func (s *DescribeClusterModelResponse) SetHeaders(v map[string]*string) *DescribeClusterModelResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterModelResponse) SetStatusCode(v int32) *DescribeClusterModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterModelResponse) SetBody(v *DescribeClusterModelResponseBody) *DescribeClusterModelResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

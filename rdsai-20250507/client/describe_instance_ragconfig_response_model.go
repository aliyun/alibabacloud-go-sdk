// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRAGConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceRAGConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceRAGConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceRAGConfigResponseBody) *DescribeInstanceRAGConfigResponse
	GetBody() *DescribeInstanceRAGConfigResponseBody
}

type DescribeInstanceRAGConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceRAGConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceRAGConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRAGConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRAGConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceRAGConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceRAGConfigResponse) GetBody() *DescribeInstanceRAGConfigResponseBody {
	return s.Body
}

func (s *DescribeInstanceRAGConfigResponse) SetHeaders(v map[string]*string) *DescribeInstanceRAGConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceRAGConfigResponse) SetStatusCode(v int32) *DescribeInstanceRAGConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceRAGConfigResponse) SetBody(v *DescribeInstanceRAGConfigResponseBody) *DescribeInstanceRAGConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceRAGConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

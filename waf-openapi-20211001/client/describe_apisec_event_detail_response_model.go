// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecEventDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecEventDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecEventDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecEventDetailResponseBody) *DescribeApisecEventDetailResponse
	GetBody() *DescribeApisecEventDetailResponseBody
}

type DescribeApisecEventDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecEventDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecEventDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecEventDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecEventDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecEventDetailResponse) GetBody() *DescribeApisecEventDetailResponseBody {
	return s.Body
}

func (s *DescribeApisecEventDetailResponse) SetHeaders(v map[string]*string) *DescribeApisecEventDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecEventDetailResponse) SetStatusCode(v int32) *DescribeApisecEventDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecEventDetailResponse) SetBody(v *DescribeApisecEventDetailResponseBody) *DescribeApisecEventDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecEventDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

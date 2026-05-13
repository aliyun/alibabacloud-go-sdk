// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTaskDetailResponseBody) *DescribeTaskDetailResponse
	GetBody() *DescribeTaskDetailResponseBody
}

type DescribeTaskDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTaskDetailResponse) GetBody() *DescribeTaskDetailResponseBody {
	return s.Body
}

func (s *DescribeTaskDetailResponse) SetHeaders(v map[string]*string) *DescribeTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskDetailResponse) SetStatusCode(v int32) *DescribeTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskDetailResponse) SetBody(v *DescribeTaskDetailResponseBody) *DescribeTaskDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

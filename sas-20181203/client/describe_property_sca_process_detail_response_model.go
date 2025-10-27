// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyScaProcessDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertyScaProcessDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertyScaProcessDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertyScaProcessDetailResponseBody) *DescribePropertyScaProcessDetailResponse
	GetBody() *DescribePropertyScaProcessDetailResponseBody
}

type DescribePropertyScaProcessDetailResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertyScaProcessDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertyScaProcessDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaProcessDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaProcessDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertyScaProcessDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertyScaProcessDetailResponse) GetBody() *DescribePropertyScaProcessDetailResponseBody {
	return s.Body
}

func (s *DescribePropertyScaProcessDetailResponse) SetHeaders(v map[string]*string) *DescribePropertyScaProcessDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyScaProcessDetailResponse) SetStatusCode(v int32) *DescribePropertyScaProcessDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponse) SetBody(v *DescribePropertyScaProcessDetailResponseBody) *DescribePropertyScaProcessDetailResponse {
	s.Body = v
	return s
}

func (s *DescribePropertyScaProcessDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

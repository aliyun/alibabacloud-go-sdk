// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoScalingHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoScalingHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoScalingHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoScalingHistoryResponseBody) *DescribeAutoScalingHistoryResponse
	GetBody() *DescribeAutoScalingHistoryResponseBody
}

type DescribeAutoScalingHistoryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoScalingHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoScalingHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoScalingHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoScalingHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoScalingHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoScalingHistoryResponse) GetBody() *DescribeAutoScalingHistoryResponseBody {
	return s.Body
}

func (s *DescribeAutoScalingHistoryResponse) SetHeaders(v map[string]*string) *DescribeAutoScalingHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoScalingHistoryResponse) SetStatusCode(v int32) *DescribeAutoScalingHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponse) SetBody(v *DescribeAutoScalingHistoryResponseBody) *DescribeAutoScalingHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoScalingHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

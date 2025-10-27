// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterBasicInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterBasicInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterBasicInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterBasicInfoResponseBody) *DescribeClusterBasicInfoResponse
	GetBody() *DescribeClusterBasicInfoResponseBody
}

type DescribeClusterBasicInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterBasicInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterBasicInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterBasicInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterBasicInfoResponse) GetBody() *DescribeClusterBasicInfoResponseBody {
	return s.Body
}

func (s *DescribeClusterBasicInfoResponse) SetHeaders(v map[string]*string) *DescribeClusterBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterBasicInfoResponse) SetStatusCode(v int32) *DescribeClusterBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterBasicInfoResponse) SetBody(v *DescribeClusterBasicInfoResponseBody) *DescribeClusterBasicInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterBasicInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

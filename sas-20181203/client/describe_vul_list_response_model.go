// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVulListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVulListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVulListResponseBody) *DescribeVulListResponse
	GetBody() *DescribeVulListResponseBody
}

type DescribeVulListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVulListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVulListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVulListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVulListResponse) GetBody() *DescribeVulListResponseBody {
	return s.Body
}

func (s *DescribeVulListResponse) SetHeaders(v map[string]*string) *DescribeVulListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulListResponse) SetStatusCode(v int32) *DescribeVulListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulListResponse) SetBody(v *DescribeVulListResponseBody) *DescribeVulListResponse {
	s.Body = v
	return s
}

func (s *DescribeVulListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVfwIPSConfigListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVfwIPSConfigListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVfwIPSConfigListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVfwIPSConfigListResponseBody) *DescribeVfwIPSConfigListResponse
	GetBody() *DescribeVfwIPSConfigListResponseBody
}

type DescribeVfwIPSConfigListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVfwIPSConfigListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVfwIPSConfigListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVfwIPSConfigListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVfwIPSConfigListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVfwIPSConfigListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVfwIPSConfigListResponse) GetBody() *DescribeVfwIPSConfigListResponseBody {
	return s.Body
}

func (s *DescribeVfwIPSConfigListResponse) SetHeaders(v map[string]*string) *DescribeVfwIPSConfigListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVfwIPSConfigListResponse) SetStatusCode(v int32) *DescribeVfwIPSConfigListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVfwIPSConfigListResponse) SetBody(v *DescribeVfwIPSConfigListResponseBody) *DescribeVfwIPSConfigListResponse {
	s.Body = v
	return s
}

func (s *DescribeVfwIPSConfigListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

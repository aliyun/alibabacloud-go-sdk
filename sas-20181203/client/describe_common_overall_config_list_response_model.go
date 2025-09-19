// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonOverallConfigListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCommonOverallConfigListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCommonOverallConfigListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCommonOverallConfigListResponseBody) *DescribeCommonOverallConfigListResponse
	GetBody() *DescribeCommonOverallConfigListResponseBody
}

type DescribeCommonOverallConfigListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCommonOverallConfigListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCommonOverallConfigListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonOverallConfigListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCommonOverallConfigListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCommonOverallConfigListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCommonOverallConfigListResponse) GetBody() *DescribeCommonOverallConfigListResponseBody {
	return s.Body
}

func (s *DescribeCommonOverallConfigListResponse) SetHeaders(v map[string]*string) *DescribeCommonOverallConfigListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCommonOverallConfigListResponse) SetStatusCode(v int32) *DescribeCommonOverallConfigListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCommonOverallConfigListResponse) SetBody(v *DescribeCommonOverallConfigListResponseBody) *DescribeCommonOverallConfigListResponse {
	s.Body = v
	return s
}

func (s *DescribeCommonOverallConfigListResponse) Validate() error {
	return dara.Validate(s)
}

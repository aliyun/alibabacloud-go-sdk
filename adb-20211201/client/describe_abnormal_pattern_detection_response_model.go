// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAbnormalPatternDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAbnormalPatternDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAbnormalPatternDetectionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAbnormalPatternDetectionResponseBody) *DescribeAbnormalPatternDetectionResponse
	GetBody() *DescribeAbnormalPatternDetectionResponseBody
}

type DescribeAbnormalPatternDetectionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAbnormalPatternDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAbnormalPatternDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAbnormalPatternDetectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeAbnormalPatternDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAbnormalPatternDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAbnormalPatternDetectionResponse) GetBody() *DescribeAbnormalPatternDetectionResponseBody {
	return s.Body
}

func (s *DescribeAbnormalPatternDetectionResponse) SetHeaders(v map[string]*string) *DescribeAbnormalPatternDetectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponse) SetStatusCode(v int32) *DescribeAbnormalPatternDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponse) SetBody(v *DescribeAbnormalPatternDetectionResponseBody) *DescribeAbnormalPatternDetectionResponse {
	s.Body = v
	return s
}

func (s *DescribeAbnormalPatternDetectionResponse) Validate() error {
	return dara.Validate(s)
}

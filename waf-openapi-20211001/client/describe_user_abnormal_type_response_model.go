// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAbnormalTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserAbnormalTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserAbnormalTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserAbnormalTypeResponseBody) *DescribeUserAbnormalTypeResponse
	GetBody() *DescribeUserAbnormalTypeResponseBody
}

type DescribeUserAbnormalTypeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserAbnormalTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserAbnormalTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAbnormalTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserAbnormalTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserAbnormalTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserAbnormalTypeResponse) GetBody() *DescribeUserAbnormalTypeResponseBody {
	return s.Body
}

func (s *DescribeUserAbnormalTypeResponse) SetHeaders(v map[string]*string) *DescribeUserAbnormalTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserAbnormalTypeResponse) SetStatusCode(v int32) *DescribeUserAbnormalTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserAbnormalTypeResponse) SetBody(v *DescribeUserAbnormalTypeResponseBody) *DescribeUserAbnormalTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeUserAbnormalTypeResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventLevelCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventLevelCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventLevelCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventLevelCountResponseBody) *DescribeEventLevelCountResponse
	GetBody() *DescribeEventLevelCountResponseBody
}

type DescribeEventLevelCountResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventLevelCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventLevelCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventLevelCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventLevelCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventLevelCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventLevelCountResponse) GetBody() *DescribeEventLevelCountResponseBody {
	return s.Body
}

func (s *DescribeEventLevelCountResponse) SetHeaders(v map[string]*string) *DescribeEventLevelCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventLevelCountResponse) SetStatusCode(v int32) *DescribeEventLevelCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventLevelCountResponse) SetBody(v *DescribeEventLevelCountResponseBody) *DescribeEventLevelCountResponse {
	s.Body = v
	return s
}

func (s *DescribeEventLevelCountResponse) Validate() error {
	return dara.Validate(s)
}

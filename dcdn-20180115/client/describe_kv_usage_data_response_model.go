// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKvUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKvUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKvUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKvUsageDataResponseBody) *DescribeKvUsageDataResponse
	GetBody() *DescribeKvUsageDataResponseBody
}

type DescribeKvUsageDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKvUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKvUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKvUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeKvUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKvUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKvUsageDataResponse) GetBody() *DescribeKvUsageDataResponseBody {
	return s.Body
}

func (s *DescribeKvUsageDataResponse) SetHeaders(v map[string]*string) *DescribeKvUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeKvUsageDataResponse) SetStatusCode(v int32) *DescribeKvUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKvUsageDataResponse) SetBody(v *DescribeKvUsageDataResponseBody) *DescribeKvUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeKvUsageDataResponse) Validate() error {
	return dara.Validate(s)
}

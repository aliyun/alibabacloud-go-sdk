// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZoneInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeZoneInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeZoneInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeZoneInfoResponseBody) *DescribeZoneInfoResponse
	GetBody() *DescribeZoneInfoResponseBody
}

type DescribeZoneInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeZoneInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeZoneInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeZoneInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeZoneInfoResponse) GetBody() *DescribeZoneInfoResponseBody {
	return s.Body
}

func (s *DescribeZoneInfoResponse) SetHeaders(v map[string]*string) *DescribeZoneInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeZoneInfoResponse) SetStatusCode(v int32) *DescribeZoneInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZoneInfoResponse) SetBody(v *DescribeZoneInfoResponseBody) *DescribeZoneInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeZoneInfoResponse) Validate() error {
	return dara.Validate(s)
}

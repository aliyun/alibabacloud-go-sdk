// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecursionZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecursionZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecursionZoneResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecursionZoneResponseBody) *DescribeRecursionZoneResponse
	GetBody() *DescribeRecursionZoneResponseBody
}

type DescribeRecursionZoneResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecursionZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecursionZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecursionZoneResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecursionZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecursionZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecursionZoneResponse) GetBody() *DescribeRecursionZoneResponseBody {
	return s.Body
}

func (s *DescribeRecursionZoneResponse) SetHeaders(v map[string]*string) *DescribeRecursionZoneResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecursionZoneResponse) SetStatusCode(v int32) *DescribeRecursionZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecursionZoneResponse) SetBody(v *DescribeRecursionZoneResponseBody) *DescribeRecursionZoneResponse {
	s.Body = v
	return s
}

func (s *DescribeRecursionZoneResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWorkZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWorkZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWorkZonesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWorkZonesResponseBody) *DescribeWorkZonesResponse
	GetBody() *DescribeWorkZonesResponseBody
}

type DescribeWorkZonesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWorkZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWorkZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeWorkZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWorkZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWorkZonesResponse) GetBody() *DescribeWorkZonesResponseBody {
	return s.Body
}

func (s *DescribeWorkZonesResponse) SetHeaders(v map[string]*string) *DescribeWorkZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeWorkZonesResponse) SetStatusCode(v int32) *DescribeWorkZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWorkZonesResponse) SetBody(v *DescribeWorkZonesResponseBody) *DescribeWorkZonesResponse {
	s.Body = v
	return s
}

func (s *DescribeWorkZonesResponse) Validate() error {
	return dara.Validate(s)
}

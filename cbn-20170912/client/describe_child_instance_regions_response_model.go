// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChildInstanceRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChildInstanceRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChildInstanceRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChildInstanceRegionsResponseBody) *DescribeChildInstanceRegionsResponse
	GetBody() *DescribeChildInstanceRegionsResponseBody
}

type DescribeChildInstanceRegionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChildInstanceRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChildInstanceRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChildInstanceRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeChildInstanceRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChildInstanceRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChildInstanceRegionsResponse) GetBody() *DescribeChildInstanceRegionsResponseBody {
	return s.Body
}

func (s *DescribeChildInstanceRegionsResponse) SetHeaders(v map[string]*string) *DescribeChildInstanceRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeChildInstanceRegionsResponse) SetStatusCode(v int32) *DescribeChildInstanceRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChildInstanceRegionsResponse) SetBody(v *DescribeChildInstanceRegionsResponseBody) *DescribeChildInstanceRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeChildInstanceRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

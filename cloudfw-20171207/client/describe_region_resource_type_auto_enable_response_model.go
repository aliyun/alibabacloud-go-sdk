// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionResourceTypeAutoEnableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRegionResourceTypeAutoEnableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRegionResourceTypeAutoEnableResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRegionResourceTypeAutoEnableResponseBody) *DescribeRegionResourceTypeAutoEnableResponse
	GetBody() *DescribeRegionResourceTypeAutoEnableResponseBody
}

type DescribeRegionResourceTypeAutoEnableResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionResourceTypeAutoEnableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionResourceTypeAutoEnableResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceTypeAutoEnableResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceTypeAutoEnableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRegionResourceTypeAutoEnableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRegionResourceTypeAutoEnableResponse) GetBody() *DescribeRegionResourceTypeAutoEnableResponseBody {
	return s.Body
}

func (s *DescribeRegionResourceTypeAutoEnableResponse) SetHeaders(v map[string]*string) *DescribeRegionResourceTypeAutoEnableResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionResourceTypeAutoEnableResponse) SetStatusCode(v int32) *DescribeRegionResourceTypeAutoEnableResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionResourceTypeAutoEnableResponse) SetBody(v *DescribeRegionResourceTypeAutoEnableResponseBody) *DescribeRegionResourceTypeAutoEnableResponse {
	s.Body = v
	return s
}

func (s *DescribeRegionResourceTypeAutoEnableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsForRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventsForRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventsForRegionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventsForRegionResponseBody) *DescribeEventsForRegionResponse
	GetBody() *DescribeEventsForRegionResponseBody
}

type DescribeEventsForRegionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventsForRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventsForRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsForRegionResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventsForRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventsForRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventsForRegionResponse) GetBody() *DescribeEventsForRegionResponseBody {
	return s.Body
}

func (s *DescribeEventsForRegionResponse) SetHeaders(v map[string]*string) *DescribeEventsForRegionResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventsForRegionResponse) SetStatusCode(v int32) *DescribeEventsForRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventsForRegionResponse) SetBody(v *DescribeEventsForRegionResponseBody) *DescribeEventsForRegionResponse {
	s.Body = v
	return s
}

func (s *DescribeEventsForRegionResponse) Validate() error {
	return dara.Validate(s)
}

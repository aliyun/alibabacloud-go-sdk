// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiTrafficControlsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiTrafficControlsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiTrafficControlsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiTrafficControlsResponseBody) *DescribeApiTrafficControlsResponse
	GetBody() *DescribeApiTrafficControlsResponseBody
}

type DescribeApiTrafficControlsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiTrafficControlsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiTrafficControlsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiTrafficControlsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficControlsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiTrafficControlsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiTrafficControlsResponse) GetBody() *DescribeApiTrafficControlsResponseBody {
	return s.Body
}

func (s *DescribeApiTrafficControlsResponse) SetHeaders(v map[string]*string) *DescribeApiTrafficControlsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiTrafficControlsResponse) SetStatusCode(v int32) *DescribeApiTrafficControlsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiTrafficControlsResponse) SetBody(v *DescribeApiTrafficControlsResponseBody) *DescribeApiTrafficControlsResponse {
	s.Body = v
	return s
}

func (s *DescribeApiTrafficControlsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

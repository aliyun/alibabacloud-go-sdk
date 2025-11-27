// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsPullStreamInfoConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsPullStreamInfoConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsPullStreamInfoConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsPullStreamInfoConfigResponseBody) *DescribeVsPullStreamInfoConfigResponse
	GetBody() *DescribeVsPullStreamInfoConfigResponseBody
}

type DescribeVsPullStreamInfoConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsPullStreamInfoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsPullStreamInfoConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsPullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsPullStreamInfoConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsPullStreamInfoConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsPullStreamInfoConfigResponse) GetBody() *DescribeVsPullStreamInfoConfigResponseBody {
	return s.Body
}

func (s *DescribeVsPullStreamInfoConfigResponse) SetHeaders(v map[string]*string) *DescribeVsPullStreamInfoConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponse) SetStatusCode(v int32) *DescribeVsPullStreamInfoConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponse) SetBody(v *DescribeVsPullStreamInfoConfigResponseBody) *DescribeVsPullStreamInfoConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeVsPullStreamInfoConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

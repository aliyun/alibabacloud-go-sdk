// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrafficLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTrafficLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTrafficLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTrafficLogResponseBody) *DescribeTrafficLogResponse
	GetBody() *DescribeTrafficLogResponseBody
}

type DescribeTrafficLogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrafficLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrafficLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrafficLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTrafficLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTrafficLogResponse) GetBody() *DescribeTrafficLogResponseBody {
	return s.Body
}

func (s *DescribeTrafficLogResponse) SetHeaders(v map[string]*string) *DescribeTrafficLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrafficLogResponse) SetStatusCode(v int32) *DescribeTrafficLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrafficLogResponse) SetBody(v *DescribeTrafficLogResponseBody) *DescribeTrafficLogResponse {
	s.Body = v
	return s
}

func (s *DescribeTrafficLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

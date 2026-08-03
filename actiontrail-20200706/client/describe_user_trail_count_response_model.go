// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserTrailCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserTrailCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserTrailCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserTrailCountResponseBody) *DescribeUserTrailCountResponse
	GetBody() *DescribeUserTrailCountResponseBody
}

type DescribeUserTrailCountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserTrailCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserTrailCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTrailCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserTrailCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserTrailCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserTrailCountResponse) GetBody() *DescribeUserTrailCountResponseBody {
	return s.Body
}

func (s *DescribeUserTrailCountResponse) SetHeaders(v map[string]*string) *DescribeUserTrailCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserTrailCountResponse) SetStatusCode(v int32) *DescribeUserTrailCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserTrailCountResponse) SetBody(v *DescribeUserTrailCountResponseBody) *DescribeUserTrailCountResponse {
	s.Body = v
	return s
}

func (s *DescribeUserTrailCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParentPlatformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeParentPlatformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeParentPlatformResponse
	GetStatusCode() *int32
	SetBody(v *DescribeParentPlatformResponseBody) *DescribeParentPlatformResponse
	GetBody() *DescribeParentPlatformResponseBody
}

type DescribeParentPlatformResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParentPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeParentPlatformResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeParentPlatformResponse) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeParentPlatformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeParentPlatformResponse) GetBody() *DescribeParentPlatformResponseBody {
	return s.Body
}

func (s *DescribeParentPlatformResponse) SetHeaders(v map[string]*string) *DescribeParentPlatformResponse {
	s.Headers = v
	return s
}

func (s *DescribeParentPlatformResponse) SetStatusCode(v int32) *DescribeParentPlatformResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParentPlatformResponse) SetBody(v *DescribeParentPlatformResponseBody) *DescribeParentPlatformResponse {
	s.Body = v
	return s
}

func (s *DescribeParentPlatformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

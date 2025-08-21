// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParentPlatformsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeParentPlatformsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeParentPlatformsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeParentPlatformsResponseBody) *DescribeParentPlatformsResponse
	GetBody() *DescribeParentPlatformsResponseBody
}

type DescribeParentPlatformsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParentPlatformsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeParentPlatformsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeParentPlatformsResponse) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeParentPlatformsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeParentPlatformsResponse) GetBody() *DescribeParentPlatformsResponseBody {
	return s.Body
}

func (s *DescribeParentPlatformsResponse) SetHeaders(v map[string]*string) *DescribeParentPlatformsResponse {
	s.Headers = v
	return s
}

func (s *DescribeParentPlatformsResponse) SetStatusCode(v int32) *DescribeParentPlatformsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParentPlatformsResponse) SetBody(v *DescribeParentPlatformsResponseBody) *DescribeParentPlatformsResponse {
	s.Body = v
	return s
}

func (s *DescribeParentPlatformsResponse) Validate() error {
	return dara.Validate(s)
}

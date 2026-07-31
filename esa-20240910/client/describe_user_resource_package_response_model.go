// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserResourcePackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserResourcePackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserResourcePackageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserResourcePackageResponseBody) *DescribeUserResourcePackageResponse
	GetBody() *DescribeUserResourcePackageResponseBody
}

type DescribeUserResourcePackageResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserResourcePackageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcePackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserResourcePackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserResourcePackageResponse) GetBody() *DescribeUserResourcePackageResponseBody {
	return s.Body
}

func (s *DescribeUserResourcePackageResponse) SetHeaders(v map[string]*string) *DescribeUserResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserResourcePackageResponse) SetStatusCode(v int32) *DescribeUserResourcePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserResourcePackageResponse) SetBody(v *DescribeUserResourcePackageResponseBody) *DescribeUserResourcePackageResponse {
	s.Body = v
	return s
}

func (s *DescribeUserResourcePackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

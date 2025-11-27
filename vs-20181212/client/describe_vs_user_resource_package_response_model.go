// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsUserResourcePackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsUserResourcePackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsUserResourcePackageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsUserResourcePackageResponseBody) *DescribeVsUserResourcePackageResponse
	GetBody() *DescribeVsUserResourcePackageResponseBody
}

type DescribeVsUserResourcePackageResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsUserResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsUserResourcePackageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsUserResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsUserResourcePackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsUserResourcePackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsUserResourcePackageResponse) GetBody() *DescribeVsUserResourcePackageResponseBody {
	return s.Body
}

func (s *DescribeVsUserResourcePackageResponse) SetHeaders(v map[string]*string) *DescribeVsUserResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsUserResourcePackageResponse) SetStatusCode(v int32) *DescribeVsUserResourcePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsUserResourcePackageResponse) SetBody(v *DescribeVsUserResourcePackageResponseBody) *DescribeVsUserResourcePackageResponse {
	s.Body = v
	return s
}

func (s *DescribeVsUserResourcePackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

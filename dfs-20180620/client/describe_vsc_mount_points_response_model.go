// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVscMountPointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVscMountPointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVscMountPointsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVscMountPointsResponseBody) *DescribeVscMountPointsResponse
	GetBody() *DescribeVscMountPointsResponseBody
}

type DescribeVscMountPointsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVscMountPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVscMountPointsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscMountPointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVscMountPointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVscMountPointsResponse) GetBody() *DescribeVscMountPointsResponseBody {
	return s.Body
}

func (s *DescribeVscMountPointsResponse) SetHeaders(v map[string]*string) *DescribeVscMountPointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVscMountPointsResponse) SetStatusCode(v int32) *DescribeVscMountPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVscMountPointsResponse) SetBody(v *DescribeVscMountPointsResponseBody) *DescribeVscMountPointsResponse {
	s.Body = v
	return s
}

func (s *DescribeVscMountPointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

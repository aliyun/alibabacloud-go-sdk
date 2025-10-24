// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAndroidInstanceGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAndroidInstanceGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAndroidInstanceGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAndroidInstanceGroupsResponseBody) *DescribeAndroidInstanceGroupsResponse
	GetBody() *DescribeAndroidInstanceGroupsResponseBody
}

type DescribeAndroidInstanceGroupsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAndroidInstanceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAndroidInstanceGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstanceGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstanceGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAndroidInstanceGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAndroidInstanceGroupsResponse) GetBody() *DescribeAndroidInstanceGroupsResponseBody {
	return s.Body
}

func (s *DescribeAndroidInstanceGroupsResponse) SetHeaders(v map[string]*string) *DescribeAndroidInstanceGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponse) SetStatusCode(v int32) *DescribeAndroidInstanceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponse) SetBody(v *DescribeAndroidInstanceGroupsResponseBody) *DescribeAndroidInstanceGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

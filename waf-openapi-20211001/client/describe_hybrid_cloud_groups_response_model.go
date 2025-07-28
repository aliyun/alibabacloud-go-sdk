// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudGroupsResponseBody) *DescribeHybridCloudGroupsResponse
	GetBody() *DescribeHybridCloudGroupsResponseBody
}

type DescribeHybridCloudGroupsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudGroupsResponse) GetBody() *DescribeHybridCloudGroupsResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudGroupsResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudGroupsResponse) SetStatusCode(v int32) *DescribeHybridCloudGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponse) SetBody(v *DescribeHybridCloudGroupsResponseBody) *DescribeHybridCloudGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudGroupsResponse) Validate() error {
	return dara.Validate(s)
}

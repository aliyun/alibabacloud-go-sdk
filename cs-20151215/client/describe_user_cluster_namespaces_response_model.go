// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserClusterNamespacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserClusterNamespacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserClusterNamespacesResponse
	GetStatusCode() *int32
	SetBody(v []*string) *DescribeUserClusterNamespacesResponse
	GetBody() []*string
}

type DescribeUserClusterNamespacesResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*string          `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s DescribeUserClusterNamespacesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserClusterNamespacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserClusterNamespacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserClusterNamespacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserClusterNamespacesResponse) GetBody() []*string {
	return s.Body
}

func (s *DescribeUserClusterNamespacesResponse) SetHeaders(v map[string]*string) *DescribeUserClusterNamespacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserClusterNamespacesResponse) SetStatusCode(v int32) *DescribeUserClusterNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserClusterNamespacesResponse) SetBody(v []*string) *DescribeUserClusterNamespacesResponse {
	s.Body = v
	return s
}

func (s *DescribeUserClusterNamespacesResponse) Validate() error {
	return dara.Validate(s)
}

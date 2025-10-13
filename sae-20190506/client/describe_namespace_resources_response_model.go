// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespaceResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNamespaceResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNamespaceResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNamespaceResourcesResponseBody) *DescribeNamespaceResourcesResponse
	GetBody() *DescribeNamespaceResourcesResponseBody
}

type DescribeNamespaceResourcesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNamespaceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNamespaceResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNamespaceResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNamespaceResourcesResponse) GetBody() *DescribeNamespaceResourcesResponseBody {
	return s.Body
}

func (s *DescribeNamespaceResourcesResponse) SetHeaders(v map[string]*string) *DescribeNamespaceResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNamespaceResourcesResponse) SetStatusCode(v int32) *DescribeNamespaceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNamespaceResourcesResponse) SetBody(v *DescribeNamespaceResourcesResponseBody) *DescribeNamespaceResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeNamespaceResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

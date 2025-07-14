// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespaceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNamespaceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNamespaceListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNamespaceListResponseBody) *DescribeNamespaceListResponse
	GetBody() *DescribeNamespaceListResponseBody
}

type DescribeNamespaceListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNamespaceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNamespaceListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNamespaceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNamespaceListResponse) GetBody() *DescribeNamespaceListResponseBody {
	return s.Body
}

func (s *DescribeNamespaceListResponse) SetHeaders(v map[string]*string) *DescribeNamespaceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeNamespaceListResponse) SetStatusCode(v int32) *DescribeNamespaceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNamespaceListResponse) SetBody(v *DescribeNamespaceListResponseBody) *DescribeNamespaceListResponse {
	s.Body = v
	return s
}

func (s *DescribeNamespaceListResponse) Validate() error {
	return dara.Validate(s)
}

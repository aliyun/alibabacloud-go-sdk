// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNamespaceResponseBody) *DescribeNamespaceResponse
	GetBody() *DescribeNamespaceResponseBody
}

type DescribeNamespaceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNamespaceResponse) GetBody() *DescribeNamespaceResponseBody {
	return s.Body
}

func (s *DescribeNamespaceResponse) SetHeaders(v map[string]*string) *DescribeNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DescribeNamespaceResponse) SetStatusCode(v int32) *DescribeNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNamespaceResponse) SetBody(v *DescribeNamespaceResponseBody) *DescribeNamespaceResponse {
	s.Body = v
	return s
}

func (s *DescribeNamespaceResponse) Validate() error {
	return dara.Validate(s)
}

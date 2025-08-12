// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridMonitorNamespaceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridMonitorNamespaceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridMonitorNamespaceListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridMonitorNamespaceListResponseBody) *DescribeHybridMonitorNamespaceListResponse
	GetBody() *DescribeHybridMonitorNamespaceListResponseBody
}

type DescribeHybridMonitorNamespaceListResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridMonitorNamespaceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridMonitorNamespaceListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorNamespaceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorNamespaceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridMonitorNamespaceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridMonitorNamespaceListResponse) GetBody() *DescribeHybridMonitorNamespaceListResponseBody {
	return s.Body
}

func (s *DescribeHybridMonitorNamespaceListResponse) SetHeaders(v map[string]*string) *DescribeHybridMonitorNamespaceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponse) SetStatusCode(v int32) *DescribeHybridMonitorNamespaceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponse) SetBody(v *DescribeHybridMonitorNamespaceListResponseBody) *DescribeHybridMonitorNamespaceListResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponse) Validate() error {
	return dara.Validate(s)
}

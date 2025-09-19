// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridProxyLinkedClientListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridProxyLinkedClientListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridProxyLinkedClientListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridProxyLinkedClientListResponseBody) *DescribeHybridProxyLinkedClientListResponse
	GetBody() *DescribeHybridProxyLinkedClientListResponseBody
}

type DescribeHybridProxyLinkedClientListResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridProxyLinkedClientListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridProxyLinkedClientListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyLinkedClientListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyLinkedClientListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridProxyLinkedClientListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridProxyLinkedClientListResponse) GetBody() *DescribeHybridProxyLinkedClientListResponseBody {
	return s.Body
}

func (s *DescribeHybridProxyLinkedClientListResponse) SetHeaders(v map[string]*string) *DescribeHybridProxyLinkedClientListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponse) SetStatusCode(v int32) *DescribeHybridProxyLinkedClientListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponse) SetBody(v *DescribeHybridProxyLinkedClientListResponseBody) *DescribeHybridProxyLinkedClientListResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridProxyLinkedClientListResponse) Validate() error {
	return dara.Validate(s)
}

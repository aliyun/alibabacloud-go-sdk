// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackendListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackendListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackendListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackendListResponseBody) *DescribeBackendListResponse
	GetBody() *DescribeBackendListResponseBody
}

type DescribeBackendListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackendListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackendListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackendListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackendListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackendListResponse) GetBody() *DescribeBackendListResponseBody {
	return s.Body
}

func (s *DescribeBackendListResponse) SetHeaders(v map[string]*string) *DescribeBackendListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackendListResponse) SetStatusCode(v int32) *DescribeBackendListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackendListResponse) SetBody(v *DescribeBackendListResponseBody) *DescribeBackendListResponse {
	s.Body = v
	return s
}

func (s *DescribeBackendListResponse) Validate() error {
	return dara.Validate(s)
}

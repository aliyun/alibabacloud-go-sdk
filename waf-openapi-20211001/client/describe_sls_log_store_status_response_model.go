// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsLogStoreStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlsLogStoreStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlsLogStoreStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlsLogStoreStatusResponseBody) *DescribeSlsLogStoreStatusResponse
	GetBody() *DescribeSlsLogStoreStatusResponseBody
}

type DescribeSlsLogStoreStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsLogStoreStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlsLogStoreStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsLogStoreStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogStoreStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlsLogStoreStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlsLogStoreStatusResponse) GetBody() *DescribeSlsLogStoreStatusResponseBody {
	return s.Body
}

func (s *DescribeSlsLogStoreStatusResponse) SetHeaders(v map[string]*string) *DescribeSlsLogStoreStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsLogStoreStatusResponse) SetStatusCode(v int32) *DescribeSlsLogStoreStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsLogStoreStatusResponse) SetBody(v *DescribeSlsLogStoreStatusResponseBody) *DescribeSlsLogStoreStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeSlsLogStoreStatusResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsLogStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlsLogStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlsLogStoreResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlsLogStoreResponseBody) *DescribeSlsLogStoreResponse
	GetBody() *DescribeSlsLogStoreResponseBody
}

type DescribeSlsLogStoreResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsLogStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlsLogStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsLogStoreResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlsLogStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlsLogStoreResponse) GetBody() *DescribeSlsLogStoreResponseBody {
	return s.Body
}

func (s *DescribeSlsLogStoreResponse) SetHeaders(v map[string]*string) *DescribeSlsLogStoreResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsLogStoreResponse) SetStatusCode(v int32) *DescribeSlsLogStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsLogStoreResponse) SetBody(v *DescribeSlsLogStoreResponseBody) *DescribeSlsLogStoreResponse {
	s.Body = v
	return s
}

func (s *DescribeSlsLogStoreResponse) Validate() error {
	return dara.Validate(s)
}

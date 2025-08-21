// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsLogstoreInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlsLogstoreInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlsLogstoreInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlsLogstoreInfoResponseBody) *DescribeSlsLogstoreInfoResponse
	GetBody() *DescribeSlsLogstoreInfoResponseBody
}

type DescribeSlsLogstoreInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsLogstoreInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlsLogstoreInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsLogstoreInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogstoreInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlsLogstoreInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlsLogstoreInfoResponse) GetBody() *DescribeSlsLogstoreInfoResponseBody {
	return s.Body
}

func (s *DescribeSlsLogstoreInfoResponse) SetHeaders(v map[string]*string) *DescribeSlsLogstoreInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsLogstoreInfoResponse) SetStatusCode(v int32) *DescribeSlsLogstoreInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponse) SetBody(v *DescribeSlsLogstoreInfoResponseBody) *DescribeSlsLogstoreInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeSlsLogstoreInfoResponse) Validate() error {
	return dara.Validate(s)
}

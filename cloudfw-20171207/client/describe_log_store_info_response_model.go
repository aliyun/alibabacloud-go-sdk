// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogStoreInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogStoreInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogStoreInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogStoreInfoResponseBody) *DescribeLogStoreInfoResponse
	GetBody() *DescribeLogStoreInfoResponseBody
}

type DescribeLogStoreInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogStoreInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogStoreInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogStoreInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogStoreInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogStoreInfoResponse) GetBody() *DescribeLogStoreInfoResponseBody {
	return s.Body
}

func (s *DescribeLogStoreInfoResponse) SetHeaders(v map[string]*string) *DescribeLogStoreInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogStoreInfoResponse) SetStatusCode(v int32) *DescribeLogStoreInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogStoreInfoResponse) SetBody(v *DescribeLogStoreInfoResponseBody) *DescribeLogStoreInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeLogStoreInfoResponse) Validate() error {
	return dara.Validate(s)
}

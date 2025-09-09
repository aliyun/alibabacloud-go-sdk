// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataAssetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataAssetsResponseBody) *DescribeDataAssetsResponse
	GetBody() *DescribeDataAssetsResponseBody
}

type DescribeDataAssetsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAssetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataAssetsResponse) GetBody() *DescribeDataAssetsResponseBody {
	return s.Body
}

func (s *DescribeDataAssetsResponse) SetHeaders(v map[string]*string) *DescribeDataAssetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataAssetsResponse) SetStatusCode(v int32) *DescribeDataAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataAssetsResponse) SetBody(v *DescribeDataAssetsResponseBody) *DescribeDataAssetsResponse {
	s.Body = v
	return s
}

func (s *DescribeDataAssetsResponse) Validate() error {
	return dara.Validate(s)
}

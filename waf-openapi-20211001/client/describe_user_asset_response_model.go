// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserAssetResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserAssetResponseBody) *DescribeUserAssetResponse
	GetBody() *DescribeUserAssetResponseBody
}

type DescribeUserAssetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAssetResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserAssetResponse) GetBody() *DescribeUserAssetResponseBody {
	return s.Body
}

func (s *DescribeUserAssetResponse) SetHeaders(v map[string]*string) *DescribeUserAssetResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserAssetResponse) SetStatusCode(v int32) *DescribeUserAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserAssetResponse) SetBody(v *DescribeUserAssetResponseBody) *DescribeUserAssetResponse {
	s.Body = v
	return s
}

func (s *DescribeUserAssetResponse) Validate() error {
	return dara.Validate(s)
}

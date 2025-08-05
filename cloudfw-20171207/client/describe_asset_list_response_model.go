// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAssetListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAssetListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAssetListResponseBody) *DescribeAssetListResponse
	GetBody() *DescribeAssetListResponseBody
}

type DescribeAssetListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAssetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAssetListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAssetListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAssetListResponse) GetBody() *DescribeAssetListResponseBody {
	return s.Body
}

func (s *DescribeAssetListResponse) SetHeaders(v map[string]*string) *DescribeAssetListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetListResponse) SetStatusCode(v int32) *DescribeAssetListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetListResponse) SetBody(v *DescribeAssetListResponseBody) *DescribeAssetListResponse {
	s.Body = v
	return s
}

func (s *DescribeAssetListResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetDetailByUuidsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAssetDetailByUuidsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAssetDetailByUuidsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAssetDetailByUuidsResponseBody) *DescribeAssetDetailByUuidsResponse
	GetBody() *DescribeAssetDetailByUuidsResponseBody
}

type DescribeAssetDetailByUuidsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAssetDetailByUuidsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAssetDetailByUuidsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetDetailByUuidsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAssetDetailByUuidsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAssetDetailByUuidsResponse) GetBody() *DescribeAssetDetailByUuidsResponseBody {
	return s.Body
}

func (s *DescribeAssetDetailByUuidsResponse) SetHeaders(v map[string]*string) *DescribeAssetDetailByUuidsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetDetailByUuidsResponse) SetStatusCode(v int32) *DescribeAssetDetailByUuidsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponse) SetBody(v *DescribeAssetDetailByUuidsResponseBody) *DescribeAssetDetailByUuidsResponse {
	s.Body = v
	return s
}

func (s *DescribeAssetDetailByUuidsResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAssetStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAssetStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAssetStatisticResponseBody) *DescribeAssetStatisticResponse
	GetBody() *DescribeAssetStatisticResponseBody
}

type DescribeAssetStatisticResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAssetStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAssetStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAssetStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAssetStatisticResponse) GetBody() *DescribeAssetStatisticResponseBody {
	return s.Body
}

func (s *DescribeAssetStatisticResponse) SetHeaders(v map[string]*string) *DescribeAssetStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetStatisticResponse) SetStatusCode(v int32) *DescribeAssetStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetStatisticResponse) SetBody(v *DescribeAssetStatisticResponseBody) *DescribeAssetStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeAssetStatisticResponse) Validate() error {
	return dara.Validate(s)
}

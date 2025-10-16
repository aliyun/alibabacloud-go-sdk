// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetRiskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAssetRiskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAssetRiskListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAssetRiskListResponseBody) *DescribeAssetRiskListResponse
	GetBody() *DescribeAssetRiskListResponseBody
}

type DescribeAssetRiskListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAssetRiskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAssetRiskListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetRiskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetRiskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAssetRiskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAssetRiskListResponse) GetBody() *DescribeAssetRiskListResponseBody {
	return s.Body
}

func (s *DescribeAssetRiskListResponse) SetHeaders(v map[string]*string) *DescribeAssetRiskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetRiskListResponse) SetStatusCode(v int32) *DescribeAssetRiskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetRiskListResponse) SetBody(v *DescribeAssetRiskListResponseBody) *DescribeAssetRiskListResponse {
	s.Body = v
	return s
}

func (s *DescribeAssetRiskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

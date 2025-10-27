// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAssetSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAssetSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAssetSummaryResponseBody) *DescribeAssetSummaryResponse
	GetBody() *DescribeAssetSummaryResponseBody
}

type DescribeAssetSummaryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAssetSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAssetSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAssetSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAssetSummaryResponse) GetBody() *DescribeAssetSummaryResponseBody {
	return s.Body
}

func (s *DescribeAssetSummaryResponse) SetHeaders(v map[string]*string) *DescribeAssetSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetSummaryResponse) SetStatusCode(v int32) *DescribeAssetSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetSummaryResponse) SetBody(v *DescribeAssetSummaryResponseBody) *DescribeAssetSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeAssetSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

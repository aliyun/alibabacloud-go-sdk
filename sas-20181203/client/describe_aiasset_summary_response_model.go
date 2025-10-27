// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIAssetSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAIAssetSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAIAssetSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAIAssetSummaryResponseBody) *DescribeAIAssetSummaryResponse
	GetBody() *DescribeAIAssetSummaryResponseBody
}

type DescribeAIAssetSummaryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAIAssetSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAIAssetSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIAssetSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIAssetSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAIAssetSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAIAssetSummaryResponse) GetBody() *DescribeAIAssetSummaryResponseBody {
	return s.Body
}

func (s *DescribeAIAssetSummaryResponse) SetHeaders(v map[string]*string) *DescribeAIAssetSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIAssetSummaryResponse) SetStatusCode(v int32) *DescribeAIAssetSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAIAssetSummaryResponse) SetBody(v *DescribeAIAssetSummaryResponseBody) *DescribeAIAssetSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeAIAssetSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCfwRiskLevelSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCfwRiskLevelSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCfwRiskLevelSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCfwRiskLevelSummaryResponseBody) *DescribeCfwRiskLevelSummaryResponse
	GetBody() *DescribeCfwRiskLevelSummaryResponseBody
}

type DescribeCfwRiskLevelSummaryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCfwRiskLevelSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCfwRiskLevelSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCfwRiskLevelSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeCfwRiskLevelSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCfwRiskLevelSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCfwRiskLevelSummaryResponse) GetBody() *DescribeCfwRiskLevelSummaryResponseBody {
	return s.Body
}

func (s *DescribeCfwRiskLevelSummaryResponse) SetHeaders(v map[string]*string) *DescribeCfwRiskLevelSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeCfwRiskLevelSummaryResponse) SetStatusCode(v int32) *DescribeCfwRiskLevelSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCfwRiskLevelSummaryResponse) SetBody(v *DescribeCfwRiskLevelSummaryResponseBody) *DescribeCfwRiskLevelSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeCfwRiskLevelSummaryResponse) Validate() error {
	return dara.Validate(s)
}

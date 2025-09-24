// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansCoverageDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSavingsPlansCoverageDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSavingsPlansCoverageDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSavingsPlansCoverageDetailResponseBody) *DescribeSavingsPlansCoverageDetailResponse
	GetBody() *DescribeSavingsPlansCoverageDetailResponseBody
}

type DescribeSavingsPlansCoverageDetailResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSavingsPlansCoverageDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSavingsPlansCoverageDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSavingsPlansCoverageDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSavingsPlansCoverageDetailResponse) GetBody() *DescribeSavingsPlansCoverageDetailResponseBody {
	return s.Body
}

func (s *DescribeSavingsPlansCoverageDetailResponse) SetHeaders(v map[string]*string) *DescribeSavingsPlansCoverageDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponse) SetStatusCode(v int32) *DescribeSavingsPlansCoverageDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponse) SetBody(v *DescribeSavingsPlansCoverageDetailResponseBody) *DescribeSavingsPlansCoverageDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponse) Validate() error {
	return dara.Validate(s)
}

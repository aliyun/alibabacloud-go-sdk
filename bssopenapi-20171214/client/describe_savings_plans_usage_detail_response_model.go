// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansUsageDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSavingsPlansUsageDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSavingsPlansUsageDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSavingsPlansUsageDetailResponseBody) *DescribeSavingsPlansUsageDetailResponse
	GetBody() *DescribeSavingsPlansUsageDetailResponseBody
}

type DescribeSavingsPlansUsageDetailResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSavingsPlansUsageDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSavingsPlansUsageDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSavingsPlansUsageDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSavingsPlansUsageDetailResponse) GetBody() *DescribeSavingsPlansUsageDetailResponseBody {
	return s.Body
}

func (s *DescribeSavingsPlansUsageDetailResponse) SetHeaders(v map[string]*string) *DescribeSavingsPlansUsageDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponse) SetStatusCode(v int32) *DescribeSavingsPlansUsageDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponse) SetBody(v *DescribeSavingsPlansUsageDetailResponseBody) *DescribeSavingsPlansUsageDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

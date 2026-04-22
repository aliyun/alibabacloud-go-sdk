// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePvtzStatisticsSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePvtzStatisticsSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePvtzStatisticsSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribePvtzStatisticsSummaryResponseBody) *DescribePvtzStatisticsSummaryResponse
	GetBody() *DescribePvtzStatisticsSummaryResponseBody
}

type DescribePvtzStatisticsSummaryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePvtzStatisticsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePvtzStatisticsSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePvtzStatisticsSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePvtzStatisticsSummaryResponse) GetBody() *DescribePvtzStatisticsSummaryResponseBody {
	return s.Body
}

func (s *DescribePvtzStatisticsSummaryResponse) SetHeaders(v map[string]*string) *DescribePvtzStatisticsSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponse) SetStatusCode(v int32) *DescribePvtzStatisticsSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponse) SetBody(v *DescribePvtzStatisticsSummaryResponseBody) *DescribePvtzStatisticsSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

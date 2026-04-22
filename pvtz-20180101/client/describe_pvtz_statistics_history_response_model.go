// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePvtzStatisticsHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePvtzStatisticsHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePvtzStatisticsHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribePvtzStatisticsHistoryResponseBody) *DescribePvtzStatisticsHistoryResponse
	GetBody() *DescribePvtzStatisticsHistoryResponseBody
}

type DescribePvtzStatisticsHistoryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePvtzStatisticsHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePvtzStatisticsHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePvtzStatisticsHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePvtzStatisticsHistoryResponse) GetBody() *DescribePvtzStatisticsHistoryResponseBody {
	return s.Body
}

func (s *DescribePvtzStatisticsHistoryResponse) SetHeaders(v map[string]*string) *DescribePvtzStatisticsHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribePvtzStatisticsHistoryResponse) SetStatusCode(v int32) *DescribePvtzStatisticsHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePvtzStatisticsHistoryResponse) SetBody(v *DescribePvtzStatisticsHistoryResponseBody) *DescribePvtzStatisticsHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribePvtzStatisticsHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

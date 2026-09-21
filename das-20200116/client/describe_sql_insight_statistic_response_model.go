// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlInsightStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSqlInsightStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSqlInsightStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSqlInsightStatisticResponseBody) *DescribeSqlInsightStatisticResponse
	GetBody() *DescribeSqlInsightStatisticResponseBody
}

type DescribeSqlInsightStatisticResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSqlInsightStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSqlInsightStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlInsightStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlInsightStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSqlInsightStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSqlInsightStatisticResponse) GetBody() *DescribeSqlInsightStatisticResponseBody {
	return s.Body
}

func (s *DescribeSqlInsightStatisticResponse) SetHeaders(v map[string]*string) *DescribeSqlInsightStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlInsightStatisticResponse) SetStatusCode(v int32) *DescribeSqlInsightStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlInsightStatisticResponse) SetBody(v *DescribeSqlInsightStatisticResponseBody) *DescribeSqlInsightStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeSqlInsightStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

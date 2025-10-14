// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsRequestStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePdnsRequestStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePdnsRequestStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePdnsRequestStatisticsResponseBody) *DescribePdnsRequestStatisticsResponse
	GetBody() *DescribePdnsRequestStatisticsResponseBody
}

type DescribePdnsRequestStatisticsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePdnsRequestStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePdnsRequestStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsRequestStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribePdnsRequestStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePdnsRequestStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePdnsRequestStatisticsResponse) GetBody() *DescribePdnsRequestStatisticsResponseBody {
	return s.Body
}

func (s *DescribePdnsRequestStatisticsResponse) SetHeaders(v map[string]*string) *DescribePdnsRequestStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribePdnsRequestStatisticsResponse) SetStatusCode(v int32) *DescribePdnsRequestStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponse) SetBody(v *DescribePdnsRequestStatisticsResponseBody) *DescribePdnsRequestStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribePdnsRequestStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

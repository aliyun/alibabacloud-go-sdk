// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagOnlineClientStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagOnlineClientStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagOnlineClientStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagOnlineClientStatisticsResponseBody) *DescribeSagOnlineClientStatisticsResponse
	GetBody() *DescribeSagOnlineClientStatisticsResponseBody
}

type DescribeSagOnlineClientStatisticsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagOnlineClientStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagOnlineClientStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagOnlineClientStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagOnlineClientStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagOnlineClientStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagOnlineClientStatisticsResponse) GetBody() *DescribeSagOnlineClientStatisticsResponseBody {
	return s.Body
}

func (s *DescribeSagOnlineClientStatisticsResponse) SetHeaders(v map[string]*string) *DescribeSagOnlineClientStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagOnlineClientStatisticsResponse) SetStatusCode(v int32) *DescribeSagOnlineClientStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagOnlineClientStatisticsResponse) SetBody(v *DescribeSagOnlineClientStatisticsResponseBody) *DescribeSagOnlineClientStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeSagOnlineClientStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

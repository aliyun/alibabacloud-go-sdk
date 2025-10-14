// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteTimeSeriesDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSiteTimeSeriesDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSiteTimeSeriesDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSiteTimeSeriesDataResponseBody) *DescribeSiteTimeSeriesDataResponse
	GetBody() *DescribeSiteTimeSeriesDataResponseBody
}

type DescribeSiteTimeSeriesDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSiteTimeSeriesDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSiteTimeSeriesDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteTimeSeriesDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteTimeSeriesDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSiteTimeSeriesDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSiteTimeSeriesDataResponse) GetBody() *DescribeSiteTimeSeriesDataResponseBody {
	return s.Body
}

func (s *DescribeSiteTimeSeriesDataResponse) SetHeaders(v map[string]*string) *DescribeSiteTimeSeriesDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponse) SetStatusCode(v int32) *DescribeSiteTimeSeriesDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponse) SetBody(v *DescribeSiteTimeSeriesDataResponseBody) *DescribeSiteTimeSeriesDataResponse {
	s.Body = v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

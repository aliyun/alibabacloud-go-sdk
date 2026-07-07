// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteWafTimeSeriesDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSiteWafTimeSeriesDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSiteWafTimeSeriesDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSiteWafTimeSeriesDataResponseBody) *DescribeSiteWafTimeSeriesDataResponse
	GetBody() *DescribeSiteWafTimeSeriesDataResponseBody
}

type DescribeSiteWafTimeSeriesDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSiteWafTimeSeriesDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSiteWafTimeSeriesDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteWafTimeSeriesDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteWafTimeSeriesDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSiteWafTimeSeriesDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSiteWafTimeSeriesDataResponse) GetBody() *DescribeSiteWafTimeSeriesDataResponseBody {
	return s.Body
}

func (s *DescribeSiteWafTimeSeriesDataResponse) SetHeaders(v map[string]*string) *DescribeSiteWafTimeSeriesDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponse) SetStatusCode(v int32) *DescribeSiteWafTimeSeriesDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponse) SetBody(v *DescribeSiteWafTimeSeriesDataResponseBody) *DescribeSiteWafTimeSeriesDataResponse {
	s.Body = v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

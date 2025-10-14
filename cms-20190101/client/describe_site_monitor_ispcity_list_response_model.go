// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorISPCityListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSiteMonitorISPCityListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSiteMonitorISPCityListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSiteMonitorISPCityListResponseBody) *DescribeSiteMonitorISPCityListResponse
	GetBody() *DescribeSiteMonitorISPCityListResponseBody
}

type DescribeSiteMonitorISPCityListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSiteMonitorISPCityListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSiteMonitorISPCityListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorISPCityListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorISPCityListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSiteMonitorISPCityListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSiteMonitorISPCityListResponse) GetBody() *DescribeSiteMonitorISPCityListResponseBody {
	return s.Body
}

func (s *DescribeSiteMonitorISPCityListResponse) SetHeaders(v map[string]*string) *DescribeSiteMonitorISPCityListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponse) SetStatusCode(v int32) *DescribeSiteMonitorISPCityListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponse) SetBody(v *DescribeSiteMonitorISPCityListResponseBody) *DescribeSiteMonitorISPCityListResponse {
	s.Body = v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

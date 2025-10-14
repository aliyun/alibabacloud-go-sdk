// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSiteMonitorDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSiteMonitorDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSiteMonitorDataResponseBody) *DescribeSiteMonitorDataResponse
	GetBody() *DescribeSiteMonitorDataResponseBody
}

type DescribeSiteMonitorDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSiteMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSiteMonitorDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSiteMonitorDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSiteMonitorDataResponse) GetBody() *DescribeSiteMonitorDataResponseBody {
	return s.Body
}

func (s *DescribeSiteMonitorDataResponse) SetHeaders(v map[string]*string) *DescribeSiteMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteMonitorDataResponse) SetStatusCode(v int32) *DescribeSiteMonitorDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSiteMonitorDataResponse) SetBody(v *DescribeSiteMonitorDataResponseBody) *DescribeSiteMonitorDataResponse {
	s.Body = v
	return s
}

func (s *DescribeSiteMonitorDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSiteMonitorQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSiteMonitorQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSiteMonitorQuotaResponseBody) *DescribeSiteMonitorQuotaResponse
	GetBody() *DescribeSiteMonitorQuotaResponseBody
}

type DescribeSiteMonitorQuotaResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSiteMonitorQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSiteMonitorQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSiteMonitorQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSiteMonitorQuotaResponse) GetBody() *DescribeSiteMonitorQuotaResponseBody {
	return s.Body
}

func (s *DescribeSiteMonitorQuotaResponse) SetHeaders(v map[string]*string) *DescribeSiteMonitorQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteMonitorQuotaResponse) SetStatusCode(v int32) *DescribeSiteMonitorQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponse) SetBody(v *DescribeSiteMonitorQuotaResponseBody) *DescribeSiteMonitorQuotaResponse {
	s.Body = v
	return s
}

func (s *DescribeSiteMonitorQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainMonitoringUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainMonitoringUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainMonitoringUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainMonitoringUsageDataResponseBody) *DescribeLiveDomainMonitoringUsageDataResponse
	GetBody() *DescribeLiveDomainMonitoringUsageDataResponseBody
}

type DescribeLiveDomainMonitoringUsageDataResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainMonitoringUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainMonitoringUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainMonitoringUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMonitoringUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainMonitoringUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainMonitoringUsageDataResponse) GetBody() *DescribeLiveDomainMonitoringUsageDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainMonitoringUsageDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainMonitoringUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponse) SetStatusCode(v int32) *DescribeLiveDomainMonitoringUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponse) SetBody(v *DescribeLiveDomainMonitoringUsageDataResponseBody) *DescribeLiveDomainMonitoringUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataResponse) Validate() error {
	return dara.Validate(s)
}

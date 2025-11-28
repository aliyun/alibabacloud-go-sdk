// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDashboardDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDashboardDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDashboardDeviceInfoResponseBody) *DescribeDashboardDeviceInfoResponse
	GetBody() *DescribeDashboardDeviceInfoResponseBody
}

type DescribeDashboardDeviceInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDashboardDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDashboardDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDashboardDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDashboardDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDashboardDeviceInfoResponse) GetBody() *DescribeDashboardDeviceInfoResponseBody {
	return s.Body
}

func (s *DescribeDashboardDeviceInfoResponse) SetHeaders(v map[string]*string) *DescribeDashboardDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDashboardDeviceInfoResponse) SetStatusCode(v int32) *DescribeDashboardDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDashboardDeviceInfoResponse) SetBody(v *DescribeDashboardDeviceInfoResponseBody) *DescribeDashboardDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDashboardDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

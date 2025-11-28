// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardMemberDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDashboardMemberDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDashboardMemberDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDashboardMemberDeviceInfoResponseBody) *DescribeDashboardMemberDeviceInfoResponse
	GetBody() *DescribeDashboardMemberDeviceInfoResponseBody
}

type DescribeDashboardMemberDeviceInfoResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDashboardMemberDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDashboardMemberDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardMemberDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDashboardMemberDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDashboardMemberDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDashboardMemberDeviceInfoResponse) GetBody() *DescribeDashboardMemberDeviceInfoResponseBody {
	return s.Body
}

func (s *DescribeDashboardMemberDeviceInfoResponse) SetHeaders(v map[string]*string) *DescribeDashboardMemberDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoResponse) SetStatusCode(v int32) *DescribeDashboardMemberDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoResponse) SetBody(v *DescribeDashboardMemberDeviceInfoResponseBody) *DescribeDashboardMemberDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDashboardMemberDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

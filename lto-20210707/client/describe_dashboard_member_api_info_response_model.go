// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardMemberApiInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDashboardMemberApiInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDashboardMemberApiInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDashboardMemberApiInfoResponseBody) *DescribeDashboardMemberApiInfoResponse
	GetBody() *DescribeDashboardMemberApiInfoResponseBody
}

type DescribeDashboardMemberApiInfoResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDashboardMemberApiInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDashboardMemberApiInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardMemberApiInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDashboardMemberApiInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDashboardMemberApiInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDashboardMemberApiInfoResponse) GetBody() *DescribeDashboardMemberApiInfoResponseBody {
	return s.Body
}

func (s *DescribeDashboardMemberApiInfoResponse) SetHeaders(v map[string]*string) *DescribeDashboardMemberApiInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDashboardMemberApiInfoResponse) SetStatusCode(v int32) *DescribeDashboardMemberApiInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDashboardMemberApiInfoResponse) SetBody(v *DescribeDashboardMemberApiInfoResponseBody) *DescribeDashboardMemberApiInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDashboardMemberApiInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

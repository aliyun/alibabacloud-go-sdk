// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardApiInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDashboardApiInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDashboardApiInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDashboardApiInfoResponseBody) *DescribeDashboardApiInfoResponse
	GetBody() *DescribeDashboardApiInfoResponseBody
}

type DescribeDashboardApiInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDashboardApiInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDashboardApiInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardApiInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDashboardApiInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDashboardApiInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDashboardApiInfoResponse) GetBody() *DescribeDashboardApiInfoResponseBody {
	return s.Body
}

func (s *DescribeDashboardApiInfoResponse) SetHeaders(v map[string]*string) *DescribeDashboardApiInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDashboardApiInfoResponse) SetStatusCode(v int32) *DescribeDashboardApiInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDashboardApiInfoResponse) SetBody(v *DescribeDashboardApiInfoResponseBody) *DescribeDashboardApiInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDashboardApiInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

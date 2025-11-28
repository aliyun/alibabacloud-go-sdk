// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardBaseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDashboardBaseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDashboardBaseInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDashboardBaseInfoResponseBody) *DescribeDashboardBaseInfoResponse
	GetBody() *DescribeDashboardBaseInfoResponseBody
}

type DescribeDashboardBaseInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDashboardBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDashboardBaseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDashboardBaseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDashboardBaseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDashboardBaseInfoResponse) GetBody() *DescribeDashboardBaseInfoResponseBody {
	return s.Body
}

func (s *DescribeDashboardBaseInfoResponse) SetHeaders(v map[string]*string) *DescribeDashboardBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDashboardBaseInfoResponse) SetStatusCode(v int32) *DescribeDashboardBaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDashboardBaseInfoResponse) SetBody(v *DescribeDashboardBaseInfoResponseBody) *DescribeDashboardBaseInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDashboardBaseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

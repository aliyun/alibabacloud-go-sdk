// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlarmBannerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlarmBannerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlarmBannerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlarmBannerResponseBody) *DescribeAlarmBannerResponse
	GetBody() *DescribeAlarmBannerResponseBody
}

type DescribeAlarmBannerResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlarmBannerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlarmBannerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmBannerResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlarmBannerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlarmBannerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlarmBannerResponse) GetBody() *DescribeAlarmBannerResponseBody {
	return s.Body
}

func (s *DescribeAlarmBannerResponse) SetHeaders(v map[string]*string) *DescribeAlarmBannerResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlarmBannerResponse) SetStatusCode(v int32) *DescribeAlarmBannerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlarmBannerResponse) SetBody(v *DescribeAlarmBannerResponseBody) *DescribeAlarmBannerResponse {
	s.Body = v
	return s
}

func (s *DescribeAlarmBannerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

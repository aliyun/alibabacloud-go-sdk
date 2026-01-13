// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadMeteringDailyDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadMeteringDailyDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadMeteringDailyDetailResponse
	GetStatusCode() *int32
	SetBody(v *DownloadMeteringDailyDetailResponseBody) *DownloadMeteringDailyDetailResponse
	GetBody() *DownloadMeteringDailyDetailResponseBody
}

type DownloadMeteringDailyDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadMeteringDailyDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadMeteringDailyDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadMeteringDailyDetailResponse) GoString() string {
	return s.String()
}

func (s *DownloadMeteringDailyDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadMeteringDailyDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadMeteringDailyDetailResponse) GetBody() *DownloadMeteringDailyDetailResponseBody {
	return s.Body
}

func (s *DownloadMeteringDailyDetailResponse) SetHeaders(v map[string]*string) *DownloadMeteringDailyDetailResponse {
	s.Headers = v
	return s
}

func (s *DownloadMeteringDailyDetailResponse) SetStatusCode(v int32) *DownloadMeteringDailyDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadMeteringDailyDetailResponse) SetBody(v *DownloadMeteringDailyDetailResponseBody) *DownloadMeteringDailyDetailResponse {
	s.Body = v
	return s
}

func (s *DownloadMeteringDailyDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

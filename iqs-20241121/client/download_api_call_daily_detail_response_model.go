// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadApiCallDailyDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadApiCallDailyDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadApiCallDailyDetailResponse
	GetStatusCode() *int32
	SetBody(v *DownloadApiCallDailyDetailResponseBody) *DownloadApiCallDailyDetailResponse
	GetBody() *DownloadApiCallDailyDetailResponseBody
}

type DownloadApiCallDailyDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadApiCallDailyDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadApiCallDailyDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadApiCallDailyDetailResponse) GoString() string {
	return s.String()
}

func (s *DownloadApiCallDailyDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadApiCallDailyDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadApiCallDailyDetailResponse) GetBody() *DownloadApiCallDailyDetailResponseBody {
	return s.Body
}

func (s *DownloadApiCallDailyDetailResponse) SetHeaders(v map[string]*string) *DownloadApiCallDailyDetailResponse {
	s.Headers = v
	return s
}

func (s *DownloadApiCallDailyDetailResponse) SetStatusCode(v int32) *DownloadApiCallDailyDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadApiCallDailyDetailResponse) SetBody(v *DownloadApiCallDailyDetailResponseBody) *DownloadApiCallDailyDetailResponse {
	s.Body = v
	return s
}

func (s *DownloadApiCallDailyDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

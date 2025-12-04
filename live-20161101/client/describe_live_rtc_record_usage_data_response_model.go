// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRtcRecordUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveRtcRecordUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveRtcRecordUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveRtcRecordUsageDataResponseBody) *DescribeLiveRtcRecordUsageDataResponse
	GetBody() *DescribeLiveRtcRecordUsageDataResponseBody
}

type DescribeLiveRtcRecordUsageDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveRtcRecordUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveRtcRecordUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRtcRecordUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveRtcRecordUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveRtcRecordUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveRtcRecordUsageDataResponse) GetBody() *DescribeLiveRtcRecordUsageDataResponseBody {
	return s.Body
}

func (s *DescribeLiveRtcRecordUsageDataResponse) SetHeaders(v map[string]*string) *DescribeLiveRtcRecordUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponse) SetStatusCode(v int32) *DescribeLiveRtcRecordUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponse) SetBody(v *DescribeLiveRtcRecordUsageDataResponseBody) *DescribeLiveRtcRecordUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveRtcRecordUsageDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

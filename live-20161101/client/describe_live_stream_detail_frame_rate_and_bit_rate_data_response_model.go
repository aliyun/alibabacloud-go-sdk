// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamDetailFrameRateAndBitRateDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponse
	GetBody() *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody
}

type DescribeLiveStreamDetailFrameRateAndBitRateDataResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamDetailFrameRateAndBitRateDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamDetailFrameRateAndBitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponse) GetBody() *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponse) SetStatusCode(v int32) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponse) SetBody(v *DescribeLiveStreamDetailFrameRateAndBitRateDataResponseBody) *DescribeLiveStreamDetailFrameRateAndBitRateDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamDetailFrameRateAndBitRateDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

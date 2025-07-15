// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamBitRateDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamBitRateDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamBitRateDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamBitRateDataResponseBody) *DescribeLiveStreamBitRateDataResponse
	GetBody() *DescribeLiveStreamBitRateDataResponseBody
}

type DescribeLiveStreamBitRateDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamBitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamBitRateDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamBitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamBitRateDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamBitRateDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamBitRateDataResponse) GetBody() *DescribeLiveStreamBitRateDataResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamBitRateDataResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamBitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponse) SetStatusCode(v int32) *DescribeLiveStreamBitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponse) SetBody(v *DescribeLiveStreamBitRateDataResponseBody) *DescribeLiveStreamBitRateDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamBitRateDataResponse) Validate() error {
	return dara.Validate(s)
}

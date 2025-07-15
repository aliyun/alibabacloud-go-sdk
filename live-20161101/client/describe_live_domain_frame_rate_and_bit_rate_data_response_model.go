// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainFrameRateAndBitRateDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainFrameRateAndBitRateDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainFrameRateAndBitRateDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainFrameRateAndBitRateDataResponseBody) *DescribeLiveDomainFrameRateAndBitRateDataResponse
	GetBody() *DescribeLiveDomainFrameRateAndBitRateDataResponseBody
}

type DescribeLiveDomainFrameRateAndBitRateDataResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainFrameRateAndBitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainFrameRateAndBitRateDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainFrameRateAndBitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponse) GetBody() *DescribeLiveDomainFrameRateAndBitRateDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainFrameRateAndBitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponse) SetStatusCode(v int32) *DescribeLiveDomainFrameRateAndBitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponse) SetBody(v *DescribeLiveDomainFrameRateAndBitRateDataResponseBody) *DescribeLiveDomainFrameRateAndBitRateDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainFrameRateAndBitRateDataResponse) Validate() error {
	return dara.Validate(s)
}

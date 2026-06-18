// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheReservePriceGapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCacheReservePriceGapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCacheReservePriceGapResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCacheReservePriceGapResponseBody) *DescribeCacheReservePriceGapResponse
	GetBody() *DescribeCacheReservePriceGapResponseBody
}

type DescribeCacheReservePriceGapResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCacheReservePriceGapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCacheReservePriceGapResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheReservePriceGapResponse) GoString() string {
	return s.String()
}

func (s *DescribeCacheReservePriceGapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCacheReservePriceGapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCacheReservePriceGapResponse) GetBody() *DescribeCacheReservePriceGapResponseBody {
	return s.Body
}

func (s *DescribeCacheReservePriceGapResponse) SetHeaders(v map[string]*string) *DescribeCacheReservePriceGapResponse {
	s.Headers = v
	return s
}

func (s *DescribeCacheReservePriceGapResponse) SetStatusCode(v int32) *DescribeCacheReservePriceGapResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCacheReservePriceGapResponse) SetBody(v *DescribeCacheReservePriceGapResponseBody) *DescribeCacheReservePriceGapResponse {
	s.Body = v
	return s
}

func (s *DescribeCacheReservePriceGapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

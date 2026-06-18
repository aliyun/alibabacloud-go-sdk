// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheReservePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCacheReservePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCacheReservePriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCacheReservePriceResponseBody) *DescribeCacheReservePriceResponse
	GetBody() *DescribeCacheReservePriceResponseBody
}

type DescribeCacheReservePriceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCacheReservePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCacheReservePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheReservePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeCacheReservePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCacheReservePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCacheReservePriceResponse) GetBody() *DescribeCacheReservePriceResponseBody {
	return s.Body
}

func (s *DescribeCacheReservePriceResponse) SetHeaders(v map[string]*string) *DescribeCacheReservePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeCacheReservePriceResponse) SetStatusCode(v int32) *DescribeCacheReservePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCacheReservePriceResponse) SetBody(v *DescribeCacheReservePriceResponseBody) *DescribeCacheReservePriceResponse {
	s.Body = v
	return s
}

func (s *DescribeCacheReservePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

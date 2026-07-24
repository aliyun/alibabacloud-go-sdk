// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDoSPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDoSPriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDoSPriceResponseBody) *DescribeDDoSPriceResponse
	GetBody() *DescribeDDoSPriceResponseBody
}

type DescribeDDoSPriceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDoSPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDoSPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDoSPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDoSPriceResponse) GetBody() *DescribeDDoSPriceResponseBody {
	return s.Body
}

func (s *DescribeDDoSPriceResponse) SetHeaders(v map[string]*string) *DescribeDDoSPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDoSPriceResponse) SetStatusCode(v int32) *DescribeDDoSPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDoSPriceResponse) SetBody(v *DescribeDDoSPriceResponseBody) *DescribeDDoSPriceResponse {
	s.Body = v
	return s
}

func (s *DescribeDDoSPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

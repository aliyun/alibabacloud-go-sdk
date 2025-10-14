// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCollectFlightLowestPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CollectFlightLowestPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CollectFlightLowestPriceResponse
	GetStatusCode() *int32
	SetBody(v *CollectFlightLowestPriceResponseBody) *CollectFlightLowestPriceResponse
	GetBody() *CollectFlightLowestPriceResponseBody
}

type CollectFlightLowestPriceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CollectFlightLowestPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CollectFlightLowestPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s CollectFlightLowestPriceResponse) GoString() string {
	return s.String()
}

func (s *CollectFlightLowestPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CollectFlightLowestPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CollectFlightLowestPriceResponse) GetBody() *CollectFlightLowestPriceResponseBody {
	return s.Body
}

func (s *CollectFlightLowestPriceResponse) SetHeaders(v map[string]*string) *CollectFlightLowestPriceResponse {
	s.Headers = v
	return s
}

func (s *CollectFlightLowestPriceResponse) SetStatusCode(v int32) *CollectFlightLowestPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *CollectFlightLowestPriceResponse) SetBody(v *CollectFlightLowestPriceResponseBody) *CollectFlightLowestPriceResponse {
	s.Body = v
	return s
}

func (s *CollectFlightLowestPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

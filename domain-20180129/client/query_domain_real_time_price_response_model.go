// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainRealTimePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDomainRealTimePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDomainRealTimePriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryDomainRealTimePriceResponseBody) *QueryDomainRealTimePriceResponse
	GetBody() *QueryDomainRealTimePriceResponseBody
}

type QueryDomainRealTimePriceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDomainRealTimePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDomainRealTimePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainRealTimePriceResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainRealTimePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDomainRealTimePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDomainRealTimePriceResponse) GetBody() *QueryDomainRealTimePriceResponseBody {
	return s.Body
}

func (s *QueryDomainRealTimePriceResponse) SetHeaders(v map[string]*string) *QueryDomainRealTimePriceResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainRealTimePriceResponse) SetStatusCode(v int32) *QueryDomainRealTimePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainRealTimePriceResponse) SetBody(v *QueryDomainRealTimePriceResponseBody) *QueryDomainRealTimePriceResponse {
	s.Body = v
	return s
}

func (s *QueryDomainRealTimePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

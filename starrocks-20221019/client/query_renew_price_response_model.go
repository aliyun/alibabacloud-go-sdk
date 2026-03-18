// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRenewPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRenewPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRenewPriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryRenewPriceResponseBody) *QueryRenewPriceResponse
	GetBody() *QueryRenewPriceResponseBody
}

type QueryRenewPriceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRenewPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRenewPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryRenewPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRenewPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRenewPriceResponse) GetBody() *QueryRenewPriceResponseBody {
	return s.Body
}

func (s *QueryRenewPriceResponse) SetHeaders(v map[string]*string) *QueryRenewPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryRenewPriceResponse) SetStatusCode(v int32) *QueryRenewPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRenewPriceResponse) SetBody(v *QueryRenewPriceResponseBody) *QueryRenewPriceResponse {
	s.Body = v
	return s
}

func (s *QueryRenewPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

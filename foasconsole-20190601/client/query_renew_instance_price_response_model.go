// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRenewInstancePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRenewInstancePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRenewInstancePriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryRenewInstancePriceResponseBody) *QueryRenewInstancePriceResponse
	GetBody() *QueryRenewInstancePriceResponseBody
}

type QueryRenewInstancePriceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRenewInstancePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRenewInstancePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewInstancePriceResponse) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRenewInstancePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRenewInstancePriceResponse) GetBody() *QueryRenewInstancePriceResponseBody {
	return s.Body
}

func (s *QueryRenewInstancePriceResponse) SetHeaders(v map[string]*string) *QueryRenewInstancePriceResponse {
	s.Headers = v
	return s
}

func (s *QueryRenewInstancePriceResponse) SetStatusCode(v int32) *QueryRenewInstancePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRenewInstancePriceResponse) SetBody(v *QueryRenewInstancePriceResponseBody) *QueryRenewInstancePriceResponse {
	s.Body = v
	return s
}

func (s *QueryRenewInstancePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

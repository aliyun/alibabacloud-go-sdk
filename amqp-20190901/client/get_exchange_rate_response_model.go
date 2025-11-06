// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExchangeRateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExchangeRateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExchangeRateResponse
	GetStatusCode() *int32
	SetBody(v *GetExchangeRateResponseBody) *GetExchangeRateResponse
	GetBody() *GetExchangeRateResponseBody
}

type GetExchangeRateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExchangeRateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExchangeRateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExchangeRateResponse) GoString() string {
	return s.String()
}

func (s *GetExchangeRateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExchangeRateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExchangeRateResponse) GetBody() *GetExchangeRateResponseBody {
	return s.Body
}

func (s *GetExchangeRateResponse) SetHeaders(v map[string]*string) *GetExchangeRateResponse {
	s.Headers = v
	return s
}

func (s *GetExchangeRateResponse) SetStatusCode(v int32) *GetExchangeRateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExchangeRateResponse) SetBody(v *GetExchangeRateResponseBody) *GetExchangeRateResponse {
	s.Body = v
	return s
}

func (s *GetExchangeRateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

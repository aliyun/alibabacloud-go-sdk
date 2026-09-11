// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApiPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApiPriceResponse
	GetStatusCode() *int32
	SetBody(v *GetApiPriceResponseBody) *GetApiPriceResponse
	GetBody() *GetApiPriceResponseBody
}

type GetApiPriceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApiPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApiPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApiPriceResponse) GoString() string {
	return s.String()
}

func (s *GetApiPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApiPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApiPriceResponse) GetBody() *GetApiPriceResponseBody {
	return s.Body
}

func (s *GetApiPriceResponse) SetHeaders(v map[string]*string) *GetApiPriceResponse {
	s.Headers = v
	return s
}

func (s *GetApiPriceResponse) SetStatusCode(v int32) *GetApiPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApiPriceResponse) SetBody(v *GetApiPriceResponseBody) *GetApiPriceResponse {
	s.Body = v
	return s
}

func (s *GetApiPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

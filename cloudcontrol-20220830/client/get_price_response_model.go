// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPriceResponse
	GetStatusCode() *int32
	SetBody(v *GetPriceResponseBody) *GetPriceResponse
	GetBody() *GetPriceResponseBody
}

type GetPriceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPriceResponse) GoString() string {
	return s.String()
}

func (s *GetPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPriceResponse) GetBody() *GetPriceResponseBody {
	return s.Body
}

func (s *GetPriceResponse) SetHeaders(v map[string]*string) *GetPriceResponse {
	s.Headers = v
	return s
}

func (s *GetPriceResponse) SetStatusCode(v int32) *GetPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPriceResponse) SetBody(v *GetPriceResponseBody) *GetPriceResponse {
	s.Body = v
	return s
}

func (s *GetPriceResponse) Validate() error {
	return dara.Validate(s)
}

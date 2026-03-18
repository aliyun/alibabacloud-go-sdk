// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEnableMultiAzPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEnableMultiAzPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEnableMultiAzPriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryEnableMultiAzPriceResponseBody) *QueryEnableMultiAzPriceResponse
	GetBody() *QueryEnableMultiAzPriceResponseBody
}

type QueryEnableMultiAzPriceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEnableMultiAzPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEnableMultiAzPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEnableMultiAzPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryEnableMultiAzPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEnableMultiAzPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEnableMultiAzPriceResponse) GetBody() *QueryEnableMultiAzPriceResponseBody {
	return s.Body
}

func (s *QueryEnableMultiAzPriceResponse) SetHeaders(v map[string]*string) *QueryEnableMultiAzPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryEnableMultiAzPriceResponse) SetStatusCode(v int32) *QueryEnableMultiAzPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEnableMultiAzPriceResponse) SetBody(v *QueryEnableMultiAzPriceResponseBody) *QueryEnableMultiAzPriceResponse {
	s.Body = v
	return s
}

func (s *QueryEnableMultiAzPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

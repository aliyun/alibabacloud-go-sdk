// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConvertInstancePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryConvertInstancePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryConvertInstancePriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryConvertInstancePriceResponseBody) *QueryConvertInstancePriceResponse
	GetBody() *QueryConvertInstancePriceResponseBody
}

type QueryConvertInstancePriceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConvertInstancePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConvertInstancePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertInstancePriceResponse) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryConvertInstancePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryConvertInstancePriceResponse) GetBody() *QueryConvertInstancePriceResponseBody {
	return s.Body
}

func (s *QueryConvertInstancePriceResponse) SetHeaders(v map[string]*string) *QueryConvertInstancePriceResponse {
	s.Headers = v
	return s
}

func (s *QueryConvertInstancePriceResponse) SetStatusCode(v int32) *QueryConvertInstancePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConvertInstancePriceResponse) SetBody(v *QueryConvertInstancePriceResponseBody) *QueryConvertInstancePriceResponse {
	s.Body = v
	return s
}

func (s *QueryConvertInstancePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

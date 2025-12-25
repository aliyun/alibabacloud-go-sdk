// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConvertPrepayInstancePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryConvertPrepayInstancePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryConvertPrepayInstancePriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryConvertPrepayInstancePriceResponseBody) *QueryConvertPrepayInstancePriceResponse
	GetBody() *QueryConvertPrepayInstancePriceResponseBody
}

type QueryConvertPrepayInstancePriceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConvertPrepayInstancePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConvertPrepayInstancePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceResponse) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryConvertPrepayInstancePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryConvertPrepayInstancePriceResponse) GetBody() *QueryConvertPrepayInstancePriceResponseBody {
	return s.Body
}

func (s *QueryConvertPrepayInstancePriceResponse) SetHeaders(v map[string]*string) *QueryConvertPrepayInstancePriceResponse {
	s.Headers = v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponse) SetStatusCode(v int32) *QueryConvertPrepayInstancePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponse) SetBody(v *QueryConvertPrepayInstancePriceResponseBody) *QueryConvertPrepayInstancePriceResponse {
	s.Body = v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

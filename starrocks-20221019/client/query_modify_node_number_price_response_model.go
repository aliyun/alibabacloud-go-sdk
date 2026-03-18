// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyNodeNumberPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryModifyNodeNumberPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryModifyNodeNumberPriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryModifyNodeNumberPriceResponseBody) *QueryModifyNodeNumberPriceResponse
	GetBody() *QueryModifyNodeNumberPriceResponseBody
}

type QueryModifyNodeNumberPriceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryModifyNodeNumberPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryModifyNodeNumberPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyNodeNumberPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryModifyNodeNumberPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryModifyNodeNumberPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryModifyNodeNumberPriceResponse) GetBody() *QueryModifyNodeNumberPriceResponseBody {
	return s.Body
}

func (s *QueryModifyNodeNumberPriceResponse) SetHeaders(v map[string]*string) *QueryModifyNodeNumberPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryModifyNodeNumberPriceResponse) SetStatusCode(v int32) *QueryModifyNodeNumberPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryModifyNodeNumberPriceResponse) SetBody(v *QueryModifyNodeNumberPriceResponseBody) *QueryModifyNodeNumberPriceResponse {
	s.Body = v
	return s
}

func (s *QueryModifyNodeNumberPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

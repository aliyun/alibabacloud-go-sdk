// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCreateInstancePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCreateInstancePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCreateInstancePriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryCreateInstancePriceResponseBody) *QueryCreateInstancePriceResponse
	GetBody() *QueryCreateInstancePriceResponseBody
}

type QueryCreateInstancePriceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCreateInstancePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCreateInstancePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceResponse) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCreateInstancePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCreateInstancePriceResponse) GetBody() *QueryCreateInstancePriceResponseBody {
	return s.Body
}

func (s *QueryCreateInstancePriceResponse) SetHeaders(v map[string]*string) *QueryCreateInstancePriceResponse {
	s.Headers = v
	return s
}

func (s *QueryCreateInstancePriceResponse) SetStatusCode(v int32) *QueryCreateInstancePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCreateInstancePriceResponse) SetBody(v *QueryCreateInstancePriceResponseBody) *QueryCreateInstancePriceResponse {
	s.Body = v
	return s
}

func (s *QueryCreateInstancePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

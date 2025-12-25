// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyInstancePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryModifyInstancePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryModifyInstancePriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryModifyInstancePriceResponseBody) *QueryModifyInstancePriceResponse
	GetBody() *QueryModifyInstancePriceResponseBody
}

type QueryModifyInstancePriceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryModifyInstancePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryModifyInstancePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyInstancePriceResponse) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryModifyInstancePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryModifyInstancePriceResponse) GetBody() *QueryModifyInstancePriceResponseBody {
	return s.Body
}

func (s *QueryModifyInstancePriceResponse) SetHeaders(v map[string]*string) *QueryModifyInstancePriceResponse {
	s.Headers = v
	return s
}

func (s *QueryModifyInstancePriceResponse) SetStatusCode(v int32) *QueryModifyInstancePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryModifyInstancePriceResponse) SetBody(v *QueryModifyInstancePriceResponseBody) *QueryModifyInstancePriceResponse {
	s.Body = v
	return s
}

func (s *QueryModifyInstancePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

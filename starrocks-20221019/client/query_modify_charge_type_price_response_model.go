// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyChargeTypePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryModifyChargeTypePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryModifyChargeTypePriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryModifyChargeTypePriceResponseBody) *QueryModifyChargeTypePriceResponse
	GetBody() *QueryModifyChargeTypePriceResponseBody
}

type QueryModifyChargeTypePriceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryModifyChargeTypePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryModifyChargeTypePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyChargeTypePriceResponse) GoString() string {
	return s.String()
}

func (s *QueryModifyChargeTypePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryModifyChargeTypePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryModifyChargeTypePriceResponse) GetBody() *QueryModifyChargeTypePriceResponseBody {
	return s.Body
}

func (s *QueryModifyChargeTypePriceResponse) SetHeaders(v map[string]*string) *QueryModifyChargeTypePriceResponse {
	s.Headers = v
	return s
}

func (s *QueryModifyChargeTypePriceResponse) SetStatusCode(v int32) *QueryModifyChargeTypePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryModifyChargeTypePriceResponse) SetBody(v *QueryModifyChargeTypePriceResponseBody) *QueryModifyChargeTypePriceResponse {
	s.Body = v
	return s
}

func (s *QueryModifyChargeTypePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

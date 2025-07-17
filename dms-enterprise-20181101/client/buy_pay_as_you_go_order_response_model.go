// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuyPayAsYouGoOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BuyPayAsYouGoOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BuyPayAsYouGoOrderResponse
	GetStatusCode() *int32
	SetBody(v *BuyPayAsYouGoOrderResponseBody) *BuyPayAsYouGoOrderResponse
	GetBody() *BuyPayAsYouGoOrderResponseBody
}

type BuyPayAsYouGoOrderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BuyPayAsYouGoOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BuyPayAsYouGoOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s BuyPayAsYouGoOrderResponse) GoString() string {
	return s.String()
}

func (s *BuyPayAsYouGoOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BuyPayAsYouGoOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BuyPayAsYouGoOrderResponse) GetBody() *BuyPayAsYouGoOrderResponseBody {
	return s.Body
}

func (s *BuyPayAsYouGoOrderResponse) SetHeaders(v map[string]*string) *BuyPayAsYouGoOrderResponse {
	s.Headers = v
	return s
}

func (s *BuyPayAsYouGoOrderResponse) SetStatusCode(v int32) *BuyPayAsYouGoOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *BuyPayAsYouGoOrderResponse) SetBody(v *BuyPayAsYouGoOrderResponseBody) *BuyPayAsYouGoOrderResponse {
	s.Body = v
	return s
}

func (s *BuyPayAsYouGoOrderResponse) Validate() error {
	return dara.Validate(s)
}

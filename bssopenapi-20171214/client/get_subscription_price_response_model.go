// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSubscriptionPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSubscriptionPriceResponse
	GetStatusCode() *int32
	SetBody(v *GetSubscriptionPriceResponseBody) *GetSubscriptionPriceResponse
	GetBody() *GetSubscriptionPriceResponseBody
}

type GetSubscriptionPriceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubscriptionPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubscriptionPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionPriceResponse) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSubscriptionPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSubscriptionPriceResponse) GetBody() *GetSubscriptionPriceResponseBody {
	return s.Body
}

func (s *GetSubscriptionPriceResponse) SetHeaders(v map[string]*string) *GetSubscriptionPriceResponse {
	s.Headers = v
	return s
}

func (s *GetSubscriptionPriceResponse) SetStatusCode(v int32) *GetSubscriptionPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubscriptionPriceResponse) SetBody(v *GetSubscriptionPriceResponseBody) *GetSubscriptionPriceResponse {
	s.Body = v
	return s
}

func (s *GetSubscriptionPriceResponse) Validate() error {
	return dara.Validate(s)
}

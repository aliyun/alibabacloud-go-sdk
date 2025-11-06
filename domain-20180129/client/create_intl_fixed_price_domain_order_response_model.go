// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntlFixedPriceDomainOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIntlFixedPriceDomainOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIntlFixedPriceDomainOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateIntlFixedPriceDomainOrderResponseBody) *CreateIntlFixedPriceDomainOrderResponse
	GetBody() *CreateIntlFixedPriceDomainOrderResponseBody
}

type CreateIntlFixedPriceDomainOrderResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIntlFixedPriceDomainOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIntlFixedPriceDomainOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIntlFixedPriceDomainOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateIntlFixedPriceDomainOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIntlFixedPriceDomainOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIntlFixedPriceDomainOrderResponse) GetBody() *CreateIntlFixedPriceDomainOrderResponseBody {
	return s.Body
}

func (s *CreateIntlFixedPriceDomainOrderResponse) SetHeaders(v map[string]*string) *CreateIntlFixedPriceDomainOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateIntlFixedPriceDomainOrderResponse) SetStatusCode(v int32) *CreateIntlFixedPriceDomainOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIntlFixedPriceDomainOrderResponse) SetBody(v *CreateIntlFixedPriceDomainOrderResponseBody) *CreateIntlFixedPriceDomainOrderResponse {
	s.Body = v
	return s
}

func (s *CreateIntlFixedPriceDomainOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

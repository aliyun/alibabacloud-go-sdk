// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPurchaserShopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPurchaserShopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPurchaserShopResponse
	GetStatusCode() *int32
	SetBody(v *Shop) *GetPurchaserShopResponse
	GetBody() *Shop
}

type GetPurchaserShopResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Shop              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPurchaserShopResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPurchaserShopResponse) GoString() string {
	return s.String()
}

func (s *GetPurchaserShopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPurchaserShopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPurchaserShopResponse) GetBody() *Shop {
	return s.Body
}

func (s *GetPurchaserShopResponse) SetHeaders(v map[string]*string) *GetPurchaserShopResponse {
	s.Headers = v
	return s
}

func (s *GetPurchaserShopResponse) SetStatusCode(v int32) *GetPurchaserShopResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPurchaserShopResponse) SetBody(v *Shop) *GetPurchaserShopResponse {
	s.Body = v
	return s
}

func (s *GetPurchaserShopResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

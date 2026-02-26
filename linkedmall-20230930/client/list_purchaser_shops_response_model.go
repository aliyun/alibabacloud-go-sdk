// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPurchaserShopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPurchaserShopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPurchaserShopsResponse
	GetStatusCode() *int32
	SetBody(v *ShopPageResult) *ListPurchaserShopsResponse
	GetBody() *ShopPageResult
}

type ListPurchaserShopsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShopPageResult    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPurchaserShopsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPurchaserShopsResponse) GoString() string {
	return s.String()
}

func (s *ListPurchaserShopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPurchaserShopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPurchaserShopsResponse) GetBody() *ShopPageResult {
	return s.Body
}

func (s *ListPurchaserShopsResponse) SetHeaders(v map[string]*string) *ListPurchaserShopsResponse {
	s.Headers = v
	return s
}

func (s *ListPurchaserShopsResponse) SetStatusCode(v int32) *ListPurchaserShopsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPurchaserShopsResponse) SetBody(v *ShopPageResult) *ListPurchaserShopsResponse {
	s.Body = v
	return s
}

func (s *ListPurchaserShopsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

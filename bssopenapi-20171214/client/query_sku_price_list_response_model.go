// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySkuPriceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySkuPriceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySkuPriceListResponse
	GetStatusCode() *int32
	SetBody(v *QuerySkuPriceListResponseBody) *QuerySkuPriceListResponse
	GetBody() *QuerySkuPriceListResponseBody
}

type QuerySkuPriceListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySkuPriceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySkuPriceListResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySkuPriceListResponse) GoString() string {
	return s.String()
}

func (s *QuerySkuPriceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySkuPriceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySkuPriceListResponse) GetBody() *QuerySkuPriceListResponseBody {
	return s.Body
}

func (s *QuerySkuPriceListResponse) SetHeaders(v map[string]*string) *QuerySkuPriceListResponse {
	s.Headers = v
	return s
}

func (s *QuerySkuPriceListResponse) SetStatusCode(v int32) *QuerySkuPriceListResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySkuPriceListResponse) SetBody(v *QuerySkuPriceListResponseBody) *QuerySkuPriceListResponse {
	s.Body = v
	return s
}

func (s *QuerySkuPriceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

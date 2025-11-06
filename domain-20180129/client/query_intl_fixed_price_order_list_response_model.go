// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIntlFixedPriceOrderListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryIntlFixedPriceOrderListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryIntlFixedPriceOrderListResponse
	GetStatusCode() *int32
	SetBody(v *QueryIntlFixedPriceOrderListResponseBody) *QueryIntlFixedPriceOrderListResponse
	GetBody() *QueryIntlFixedPriceOrderListResponseBody
}

type QueryIntlFixedPriceOrderListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryIntlFixedPriceOrderListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryIntlFixedPriceOrderListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryIntlFixedPriceOrderListResponse) GoString() string {
	return s.String()
}

func (s *QueryIntlFixedPriceOrderListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryIntlFixedPriceOrderListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryIntlFixedPriceOrderListResponse) GetBody() *QueryIntlFixedPriceOrderListResponseBody {
	return s.Body
}

func (s *QueryIntlFixedPriceOrderListResponse) SetHeaders(v map[string]*string) *QueryIntlFixedPriceOrderListResponse {
	s.Headers = v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponse) SetStatusCode(v int32) *QueryIntlFixedPriceOrderListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponse) SetBody(v *QueryIntlFixedPriceOrderListResponseBody) *QueryIntlFixedPriceOrderListResponse {
	s.Body = v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

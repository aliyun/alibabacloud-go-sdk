// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseCacheReserveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PurchaseCacheReserveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PurchaseCacheReserveResponse
	GetStatusCode() *int32
	SetBody(v *PurchaseCacheReserveResponseBody) *PurchaseCacheReserveResponse
	GetBody() *PurchaseCacheReserveResponseBody
}

type PurchaseCacheReserveResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseCacheReserveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PurchaseCacheReserveResponse) String() string {
	return dara.Prettify(s)
}

func (s PurchaseCacheReserveResponse) GoString() string {
	return s.String()
}

func (s *PurchaseCacheReserveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PurchaseCacheReserveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PurchaseCacheReserveResponse) GetBody() *PurchaseCacheReserveResponseBody {
	return s.Body
}

func (s *PurchaseCacheReserveResponse) SetHeaders(v map[string]*string) *PurchaseCacheReserveResponse {
	s.Headers = v
	return s
}

func (s *PurchaseCacheReserveResponse) SetStatusCode(v int32) *PurchaseCacheReserveResponse {
	s.StatusCode = &v
	return s
}

func (s *PurchaseCacheReserveResponse) SetBody(v *PurchaseCacheReserveResponseBody) *PurchaseCacheReserveResponse {
	s.Body = v
	return s
}

func (s *PurchaseCacheReserveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

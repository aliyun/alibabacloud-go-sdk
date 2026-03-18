// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPurchasedServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPurchasedServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPurchasedServiceResponse
	GetStatusCode() *int32
	SetBody(v *QueryPurchasedServiceResponseBody) *QueryPurchasedServiceResponse
	GetBody() *QueryPurchasedServiceResponseBody
}

type QueryPurchasedServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPurchasedServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPurchasedServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPurchasedServiceResponse) GoString() string {
	return s.String()
}

func (s *QueryPurchasedServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPurchasedServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPurchasedServiceResponse) GetBody() *QueryPurchasedServiceResponseBody {
	return s.Body
}

func (s *QueryPurchasedServiceResponse) SetHeaders(v map[string]*string) *QueryPurchasedServiceResponse {
	s.Headers = v
	return s
}

func (s *QueryPurchasedServiceResponse) SetStatusCode(v int32) *QueryPurchasedServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPurchasedServiceResponse) SetBody(v *QueryPurchasedServiceResponseBody) *QueryPurchasedServiceResponse {
	s.Body = v
	return s
}

func (s *QueryPurchasedServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCashCouponsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCashCouponsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCashCouponsResponse
	GetStatusCode() *int32
	SetBody(v *QueryCashCouponsResponseBody) *QueryCashCouponsResponse
	GetBody() *QueryCashCouponsResponseBody
}

type QueryCashCouponsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCashCouponsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCashCouponsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCashCouponsResponse) GoString() string {
	return s.String()
}

func (s *QueryCashCouponsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCashCouponsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCashCouponsResponse) GetBody() *QueryCashCouponsResponseBody {
	return s.Body
}

func (s *QueryCashCouponsResponse) SetHeaders(v map[string]*string) *QueryCashCouponsResponse {
	s.Headers = v
	return s
}

func (s *QueryCashCouponsResponse) SetStatusCode(v int32) *QueryCashCouponsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCashCouponsResponse) SetBody(v *QueryCashCouponsResponseBody) *QueryCashCouponsResponse {
	s.Body = v
	return s
}

func (s *QueryCashCouponsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

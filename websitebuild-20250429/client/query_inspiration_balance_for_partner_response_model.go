// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInspirationBalanceForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInspirationBalanceForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInspirationBalanceForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *QueryInspirationBalanceForPartnerResponseBody) *QueryInspirationBalanceForPartnerResponse
	GetBody() *QueryInspirationBalanceForPartnerResponseBody
}

type QueryInspirationBalanceForPartnerResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInspirationBalanceForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInspirationBalanceForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationBalanceForPartnerResponse) GoString() string {
	return s.String()
}

func (s *QueryInspirationBalanceForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInspirationBalanceForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInspirationBalanceForPartnerResponse) GetBody() *QueryInspirationBalanceForPartnerResponseBody {
	return s.Body
}

func (s *QueryInspirationBalanceForPartnerResponse) SetHeaders(v map[string]*string) *QueryInspirationBalanceForPartnerResponse {
	s.Headers = v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponse) SetStatusCode(v int32) *QueryInspirationBalanceForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponse) SetBody(v *QueryInspirationBalanceForPartnerResponseBody) *QueryInspirationBalanceForPartnerResponse {
	s.Body = v
	return s
}

func (s *QueryInspirationBalanceForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

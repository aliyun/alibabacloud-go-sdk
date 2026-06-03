// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGateVerifyBillingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryGateVerifyBillingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryGateVerifyBillingResponse
	GetStatusCode() *int32
	SetBody(v *QueryGateVerifyBillingResponseBody) *QueryGateVerifyBillingResponse
	GetBody() *QueryGateVerifyBillingResponseBody
}

type QueryGateVerifyBillingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGateVerifyBillingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGateVerifyBillingResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyBillingResponse) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryGateVerifyBillingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryGateVerifyBillingResponse) GetBody() *QueryGateVerifyBillingResponseBody {
	return s.Body
}

func (s *QueryGateVerifyBillingResponse) SetHeaders(v map[string]*string) *QueryGateVerifyBillingResponse {
	s.Headers = v
	return s
}

func (s *QueryGateVerifyBillingResponse) SetStatusCode(v int32) *QueryGateVerifyBillingResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGateVerifyBillingResponse) SetBody(v *QueryGateVerifyBillingResponseBody) *QueryGateVerifyBillingResponse {
	s.Body = v
	return s
}

func (s *QueryGateVerifyBillingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGateVerifyBillingPublicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryGateVerifyBillingPublicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryGateVerifyBillingPublicResponse
	GetStatusCode() *int32
	SetBody(v *QueryGateVerifyBillingPublicResponseBody) *QueryGateVerifyBillingPublicResponse
	GetBody() *QueryGateVerifyBillingPublicResponseBody
}

type QueryGateVerifyBillingPublicResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGateVerifyBillingPublicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGateVerifyBillingPublicResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyBillingPublicResponse) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingPublicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryGateVerifyBillingPublicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryGateVerifyBillingPublicResponse) GetBody() *QueryGateVerifyBillingPublicResponseBody {
	return s.Body
}

func (s *QueryGateVerifyBillingPublicResponse) SetHeaders(v map[string]*string) *QueryGateVerifyBillingPublicResponse {
	s.Headers = v
	return s
}

func (s *QueryGateVerifyBillingPublicResponse) SetStatusCode(v int32) *QueryGateVerifyBillingPublicResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponse) SetBody(v *QueryGateVerifyBillingPublicResponseBody) *QueryGateVerifyBillingPublicResponse {
	s.Body = v
	return s
}

func (s *QueryGateVerifyBillingPublicResponse) Validate() error {
	return dara.Validate(s)
}

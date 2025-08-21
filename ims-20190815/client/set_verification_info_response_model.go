// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVerificationInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetVerificationInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetVerificationInfoResponse
	GetStatusCode() *int32
	SetBody(v *SetVerificationInfoResponseBody) *SetVerificationInfoResponse
	GetBody() *SetVerificationInfoResponseBody
}

type SetVerificationInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetVerificationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetVerificationInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s SetVerificationInfoResponse) GoString() string {
	return s.String()
}

func (s *SetVerificationInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetVerificationInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetVerificationInfoResponse) GetBody() *SetVerificationInfoResponseBody {
	return s.Body
}

func (s *SetVerificationInfoResponse) SetHeaders(v map[string]*string) *SetVerificationInfoResponse {
	s.Headers = v
	return s
}

func (s *SetVerificationInfoResponse) SetStatusCode(v int32) *SetVerificationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SetVerificationInfoResponse) SetBody(v *SetVerificationInfoResponseBody) *SetVerificationInfoResponse {
	s.Body = v
	return s
}

func (s *SetVerificationInfoResponse) Validate() error {
	return dara.Validate(s)
}

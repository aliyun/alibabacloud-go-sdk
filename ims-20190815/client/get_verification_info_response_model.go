// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVerificationInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVerificationInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVerificationInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetVerificationInfoResponseBody) *GetVerificationInfoResponse
	GetBody() *GetVerificationInfoResponseBody
}

type GetVerificationInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVerificationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVerificationInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVerificationInfoResponse) GoString() string {
	return s.String()
}

func (s *GetVerificationInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVerificationInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVerificationInfoResponse) GetBody() *GetVerificationInfoResponseBody {
	return s.Body
}

func (s *GetVerificationInfoResponse) SetHeaders(v map[string]*string) *GetVerificationInfoResponse {
	s.Headers = v
	return s
}

func (s *GetVerificationInfoResponse) SetStatusCode(v int32) *GetVerificationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVerificationInfoResponse) SetBody(v *GetVerificationInfoResponseBody) *GetVerificationInfoResponse {
	s.Body = v
	return s
}

func (s *GetVerificationInfoResponse) Validate() error {
	return dara.Validate(s)
}

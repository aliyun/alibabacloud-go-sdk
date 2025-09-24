// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetResellerUserQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetResellerUserQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetResellerUserQuotaResponse
	GetStatusCode() *int32
	SetBody(v *SetResellerUserQuotaResponseBody) *SetResellerUserQuotaResponse
	GetBody() *SetResellerUserQuotaResponseBody
}

type SetResellerUserQuotaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetResellerUserQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetResellerUserQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s SetResellerUserQuotaResponse) GoString() string {
	return s.String()
}

func (s *SetResellerUserQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetResellerUserQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetResellerUserQuotaResponse) GetBody() *SetResellerUserQuotaResponseBody {
	return s.Body
}

func (s *SetResellerUserQuotaResponse) SetHeaders(v map[string]*string) *SetResellerUserQuotaResponse {
	s.Headers = v
	return s
}

func (s *SetResellerUserQuotaResponse) SetStatusCode(v int32) *SetResellerUserQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *SetResellerUserQuotaResponse) SetBody(v *SetResellerUserQuotaResponseBody) *SetResellerUserQuotaResponse {
	s.Body = v
	return s
}

func (s *SetResellerUserQuotaResponse) Validate() error {
	return dara.Validate(s)
}

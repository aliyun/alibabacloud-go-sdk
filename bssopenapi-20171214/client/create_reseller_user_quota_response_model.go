// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResellerUserQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResellerUserQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResellerUserQuotaResponse
	GetStatusCode() *int32
	SetBody(v *CreateResellerUserQuotaResponseBody) *CreateResellerUserQuotaResponse
	GetBody() *CreateResellerUserQuotaResponseBody
}

type CreateResellerUserQuotaResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResellerUserQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResellerUserQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResellerUserQuotaResponse) GoString() string {
	return s.String()
}

func (s *CreateResellerUserQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResellerUserQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResellerUserQuotaResponse) GetBody() *CreateResellerUserQuotaResponseBody {
	return s.Body
}

func (s *CreateResellerUserQuotaResponse) SetHeaders(v map[string]*string) *CreateResellerUserQuotaResponse {
	s.Headers = v
	return s
}

func (s *CreateResellerUserQuotaResponse) SetStatusCode(v int32) *CreateResellerUserQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResellerUserQuotaResponse) SetBody(v *CreateResellerUserQuotaResponseBody) *CreateResellerUserQuotaResponse {
	s.Body = v
	return s
}

func (s *CreateResellerUserQuotaResponse) Validate() error {
	return dara.Validate(s)
}

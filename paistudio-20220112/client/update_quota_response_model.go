// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateQuotaResponse
	GetStatusCode() *int32
	SetBody(v *UpdateQuotaResponseBody) *UpdateQuotaResponse
	GetBody() *UpdateQuotaResponseBody
}

type UpdateQuotaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaResponse) GoString() string {
	return s.String()
}

func (s *UpdateQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateQuotaResponse) GetBody() *UpdateQuotaResponseBody {
	return s.Body
}

func (s *UpdateQuotaResponse) SetHeaders(v map[string]*string) *UpdateQuotaResponse {
	s.Headers = v
	return s
}

func (s *UpdateQuotaResponse) SetStatusCode(v int32) *UpdateQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQuotaResponse) SetBody(v *UpdateQuotaResponseBody) *UpdateQuotaResponse {
	s.Body = v
	return s
}

func (s *UpdateQuotaResponse) Validate() error {
	return dara.Validate(s)
}

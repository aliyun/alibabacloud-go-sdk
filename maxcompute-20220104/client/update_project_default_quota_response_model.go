// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectDefaultQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProjectDefaultQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProjectDefaultQuotaResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProjectDefaultQuotaResponseBody) *UpdateProjectDefaultQuotaResponse
	GetBody() *UpdateProjectDefaultQuotaResponseBody
}

type UpdateProjectDefaultQuotaResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectDefaultQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectDefaultQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectDefaultQuotaResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectDefaultQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProjectDefaultQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProjectDefaultQuotaResponse) GetBody() *UpdateProjectDefaultQuotaResponseBody {
	return s.Body
}

func (s *UpdateProjectDefaultQuotaResponse) SetHeaders(v map[string]*string) *UpdateProjectDefaultQuotaResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectDefaultQuotaResponse) SetStatusCode(v int32) *UpdateProjectDefaultQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectDefaultQuotaResponse) SetBody(v *UpdateProjectDefaultQuotaResponseBody) *UpdateProjectDefaultQuotaResponse {
	s.Body = v
	return s
}

func (s *UpdateProjectDefaultQuotaResponse) Validate() error {
	return dara.Validate(s)
}

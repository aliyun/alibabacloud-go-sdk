// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeSubQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateComputeSubQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateComputeSubQuotaResponse
	GetStatusCode() *int32
	SetBody(v *UpdateComputeSubQuotaResponseBody) *UpdateComputeSubQuotaResponse
	GetBody() *UpdateComputeSubQuotaResponseBody
}

type UpdateComputeSubQuotaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComputeSubQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComputeSubQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeSubQuotaResponse) GoString() string {
	return s.String()
}

func (s *UpdateComputeSubQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateComputeSubQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateComputeSubQuotaResponse) GetBody() *UpdateComputeSubQuotaResponseBody {
	return s.Body
}

func (s *UpdateComputeSubQuotaResponse) SetHeaders(v map[string]*string) *UpdateComputeSubQuotaResponse {
	s.Headers = v
	return s
}

func (s *UpdateComputeSubQuotaResponse) SetStatusCode(v int32) *UpdateComputeSubQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComputeSubQuotaResponse) SetBody(v *UpdateComputeSubQuotaResponseBody) *UpdateComputeSubQuotaResponse {
	s.Body = v
	return s
}

func (s *UpdateComputeSubQuotaResponse) Validate() error {
	return dara.Validate(s)
}

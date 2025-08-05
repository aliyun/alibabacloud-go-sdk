// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeQuotaPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateComputeQuotaPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateComputeQuotaPlanResponse
	GetStatusCode() *int32
	SetBody(v *UpdateComputeQuotaPlanResponseBody) *UpdateComputeQuotaPlanResponse
	GetBody() *UpdateComputeQuotaPlanResponseBody
}

type UpdateComputeQuotaPlanResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComputeQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComputeQuotaPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateComputeQuotaPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateComputeQuotaPlanResponse) GetBody() *UpdateComputeQuotaPlanResponseBody {
	return s.Body
}

func (s *UpdateComputeQuotaPlanResponse) SetHeaders(v map[string]*string) *UpdateComputeQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateComputeQuotaPlanResponse) SetStatusCode(v int32) *UpdateComputeQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponse) SetBody(v *UpdateComputeQuotaPlanResponseBody) *UpdateComputeQuotaPlanResponse {
	s.Body = v
	return s
}

func (s *UpdateComputeQuotaPlanResponse) Validate() error {
	return dara.Validate(s)
}

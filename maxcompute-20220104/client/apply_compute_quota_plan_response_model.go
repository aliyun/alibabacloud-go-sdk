// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyComputeQuotaPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyComputeQuotaPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyComputeQuotaPlanResponse
	GetStatusCode() *int32
	SetBody(v *ApplyComputeQuotaPlanResponseBody) *ApplyComputeQuotaPlanResponse
	GetBody() *ApplyComputeQuotaPlanResponseBody
}

type ApplyComputeQuotaPlanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyComputeQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyComputeQuotaPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyComputeQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *ApplyComputeQuotaPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyComputeQuotaPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyComputeQuotaPlanResponse) GetBody() *ApplyComputeQuotaPlanResponseBody {
	return s.Body
}

func (s *ApplyComputeQuotaPlanResponse) SetHeaders(v map[string]*string) *ApplyComputeQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *ApplyComputeQuotaPlanResponse) SetStatusCode(v int32) *ApplyComputeQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyComputeQuotaPlanResponse) SetBody(v *ApplyComputeQuotaPlanResponseBody) *ApplyComputeQuotaPlanResponse {
	s.Body = v
	return s
}

func (s *ApplyComputeQuotaPlanResponse) Validate() error {
	return dara.Validate(s)
}

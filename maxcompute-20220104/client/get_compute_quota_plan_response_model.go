// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeQuotaPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetComputeQuotaPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetComputeQuotaPlanResponse
	GetStatusCode() *int32
	SetBody(v *GetComputeQuotaPlanResponseBody) *GetComputeQuotaPlanResponse
	GetBody() *GetComputeQuotaPlanResponseBody
}

type GetComputeQuotaPlanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComputeQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComputeQuotaPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s GetComputeQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetComputeQuotaPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetComputeQuotaPlanResponse) GetBody() *GetComputeQuotaPlanResponseBody {
	return s.Body
}

func (s *GetComputeQuotaPlanResponse) SetHeaders(v map[string]*string) *GetComputeQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *GetComputeQuotaPlanResponse) SetStatusCode(v int32) *GetComputeQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComputeQuotaPlanResponse) SetBody(v *GetComputeQuotaPlanResponseBody) *GetComputeQuotaPlanResponse {
	s.Body = v
	return s
}

func (s *GetComputeQuotaPlanResponse) Validate() error {
	return dara.Validate(s)
}

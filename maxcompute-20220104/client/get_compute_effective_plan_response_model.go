// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeEffectivePlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetComputeEffectivePlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetComputeEffectivePlanResponse
	GetStatusCode() *int32
	SetBody(v *GetComputeEffectivePlanResponseBody) *GetComputeEffectivePlanResponse
	GetBody() *GetComputeEffectivePlanResponseBody
}

type GetComputeEffectivePlanResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComputeEffectivePlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComputeEffectivePlanResponse) String() string {
	return dara.Prettify(s)
}

func (s GetComputeEffectivePlanResponse) GoString() string {
	return s.String()
}

func (s *GetComputeEffectivePlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetComputeEffectivePlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetComputeEffectivePlanResponse) GetBody() *GetComputeEffectivePlanResponseBody {
	return s.Body
}

func (s *GetComputeEffectivePlanResponse) SetHeaders(v map[string]*string) *GetComputeEffectivePlanResponse {
	s.Headers = v
	return s
}

func (s *GetComputeEffectivePlanResponse) SetStatusCode(v int32) *GetComputeEffectivePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComputeEffectivePlanResponse) SetBody(v *GetComputeEffectivePlanResponseBody) *GetComputeEffectivePlanResponse {
	s.Body = v
	return s
}

func (s *GetComputeEffectivePlanResponse) Validate() error {
	return dara.Validate(s)
}

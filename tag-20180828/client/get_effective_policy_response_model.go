// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEffectivePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEffectivePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEffectivePolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetEffectivePolicyResponseBody) *GetEffectivePolicyResponse
	GetBody() *GetEffectivePolicyResponseBody
}

type GetEffectivePolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEffectivePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEffectivePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEffectivePolicyResponse) GoString() string {
	return s.String()
}

func (s *GetEffectivePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEffectivePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEffectivePolicyResponse) GetBody() *GetEffectivePolicyResponseBody {
	return s.Body
}

func (s *GetEffectivePolicyResponse) SetHeaders(v map[string]*string) *GetEffectivePolicyResponse {
	s.Headers = v
	return s
}

func (s *GetEffectivePolicyResponse) SetStatusCode(v int32) *GetEffectivePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEffectivePolicyResponse) SetBody(v *GetEffectivePolicyResponseBody) *GetEffectivePolicyResponse {
	s.Body = v
	return s
}

func (s *GetEffectivePolicyResponse) Validate() error {
	return dara.Validate(s)
}

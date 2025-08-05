// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyBindingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePolicyBindingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePolicyBindingsResponse
	GetStatusCode() *int32
	SetBody(v *CreatePolicyBindingsResponseBody) *CreatePolicyBindingsResponse
	GetBody() *CreatePolicyBindingsResponseBody
}

type CreatePolicyBindingsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolicyBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolicyBindingsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyBindingsResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyBindingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePolicyBindingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePolicyBindingsResponse) GetBody() *CreatePolicyBindingsResponseBody {
	return s.Body
}

func (s *CreatePolicyBindingsResponse) SetHeaders(v map[string]*string) *CreatePolicyBindingsResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyBindingsResponse) SetStatusCode(v int32) *CreatePolicyBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolicyBindingsResponse) SetBody(v *CreatePolicyBindingsResponseBody) *CreatePolicyBindingsResponse {
	s.Body = v
	return s
}

func (s *CreatePolicyBindingsResponse) Validate() error {
	return dara.Validate(s)
}

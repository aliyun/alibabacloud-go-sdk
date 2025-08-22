// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnWafPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDcdnWafPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDcdnWafPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateDcdnWafPolicyResponseBody) *CreateDcdnWafPolicyResponse
	GetBody() *CreateDcdnWafPolicyResponseBody
}

type CreateDcdnWafPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDcdnWafPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDcdnWafPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnWafPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateDcdnWafPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDcdnWafPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDcdnWafPolicyResponse) GetBody() *CreateDcdnWafPolicyResponseBody {
	return s.Body
}

func (s *CreateDcdnWafPolicyResponse) SetHeaders(v map[string]*string) *CreateDcdnWafPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateDcdnWafPolicyResponse) SetStatusCode(v int32) *CreateDcdnWafPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDcdnWafPolicyResponse) SetBody(v *CreateDcdnWafPolicyResponseBody) *CreateDcdnWafPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateDcdnWafPolicyResponse) Validate() error {
	return dara.Validate(s)
}

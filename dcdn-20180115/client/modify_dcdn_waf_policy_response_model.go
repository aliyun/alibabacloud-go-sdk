// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDcdnWafPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDcdnWafPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDcdnWafPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDcdnWafPolicyResponseBody) *ModifyDcdnWafPolicyResponse
	GetBody() *ModifyDcdnWafPolicyResponseBody
}

type ModifyDcdnWafPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDcdnWafPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDcdnWafPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDcdnWafPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyDcdnWafPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDcdnWafPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDcdnWafPolicyResponse) GetBody() *ModifyDcdnWafPolicyResponseBody {
	return s.Body
}

func (s *ModifyDcdnWafPolicyResponse) SetHeaders(v map[string]*string) *ModifyDcdnWafPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyDcdnWafPolicyResponse) SetStatusCode(v int32) *ModifyDcdnWafPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDcdnWafPolicyResponse) SetBody(v *ModifyDcdnWafPolicyResponseBody) *ModifyDcdnWafPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyDcdnWafPolicyResponse) Validate() error {
	return dara.Validate(s)
}

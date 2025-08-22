// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDcdnWafPolicyDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDcdnWafPolicyDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDcdnWafPolicyDomainsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDcdnWafPolicyDomainsResponseBody) *ModifyDcdnWafPolicyDomainsResponse
	GetBody() *ModifyDcdnWafPolicyDomainsResponseBody
}

type ModifyDcdnWafPolicyDomainsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDcdnWafPolicyDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDcdnWafPolicyDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDcdnWafPolicyDomainsResponse) GoString() string {
	return s.String()
}

func (s *ModifyDcdnWafPolicyDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDcdnWafPolicyDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDcdnWafPolicyDomainsResponse) GetBody() *ModifyDcdnWafPolicyDomainsResponseBody {
	return s.Body
}

func (s *ModifyDcdnWafPolicyDomainsResponse) SetHeaders(v map[string]*string) *ModifyDcdnWafPolicyDomainsResponse {
	s.Headers = v
	return s
}

func (s *ModifyDcdnWafPolicyDomainsResponse) SetStatusCode(v int32) *ModifyDcdnWafPolicyDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDcdnWafPolicyDomainsResponse) SetBody(v *ModifyDcdnWafPolicyDomainsResponseBody) *ModifyDcdnWafPolicyDomainsResponse {
	s.Body = v
	return s
}

func (s *ModifyDcdnWafPolicyDomainsResponse) Validate() error {
	return dara.Validate(s)
}

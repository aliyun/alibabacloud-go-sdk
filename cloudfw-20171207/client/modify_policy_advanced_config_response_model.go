// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyAdvancedConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPolicyAdvancedConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPolicyAdvancedConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPolicyAdvancedConfigResponseBody) *ModifyPolicyAdvancedConfigResponse
	GetBody() *ModifyPolicyAdvancedConfigResponseBody
}

type ModifyPolicyAdvancedConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolicyAdvancedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolicyAdvancedConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyAdvancedConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyAdvancedConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPolicyAdvancedConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPolicyAdvancedConfigResponse) GetBody() *ModifyPolicyAdvancedConfigResponseBody {
	return s.Body
}

func (s *ModifyPolicyAdvancedConfigResponse) SetHeaders(v map[string]*string) *ModifyPolicyAdvancedConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolicyAdvancedConfigResponse) SetStatusCode(v int32) *ModifyPolicyAdvancedConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolicyAdvancedConfigResponse) SetBody(v *ModifyPolicyAdvancedConfigResponseBody) *ModifyPolicyAdvancedConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyPolicyAdvancedConfigResponse) Validate() error {
	return dara.Validate(s)
}

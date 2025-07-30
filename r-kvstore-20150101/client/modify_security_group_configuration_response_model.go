// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySecurityGroupConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySecurityGroupConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *ModifySecurityGroupConfigurationResponseBody) *ModifySecurityGroupConfigurationResponse
	GetBody() *ModifySecurityGroupConfigurationResponseBody
}

type ModifySecurityGroupConfigurationResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecurityGroupConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecurityGroupConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySecurityGroupConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySecurityGroupConfigurationResponse) GetBody() *ModifySecurityGroupConfigurationResponseBody {
	return s.Body
}

func (s *ModifySecurityGroupConfigurationResponse) SetHeaders(v map[string]*string) *ModifySecurityGroupConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityGroupConfigurationResponse) SetStatusCode(v int32) *ModifySecurityGroupConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityGroupConfigurationResponse) SetBody(v *ModifySecurityGroupConfigurationResponseBody) *ModifySecurityGroupConfigurationResponse {
	s.Body = v
	return s
}

func (s *ModifySecurityGroupConfigurationResponse) Validate() error {
	return dara.Validate(s)
}

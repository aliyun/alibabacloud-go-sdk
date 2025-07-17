// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAlertConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAlertConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAlertConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAlertConfigurationResponseBody) *ModifyAlertConfigurationResponse
	GetBody() *ModifyAlertConfigurationResponseBody
}

type ModifyAlertConfigurationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAlertConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAlertConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAlertConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyAlertConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAlertConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAlertConfigurationResponse) GetBody() *ModifyAlertConfigurationResponseBody {
	return s.Body
}

func (s *ModifyAlertConfigurationResponse) SetHeaders(v map[string]*string) *ModifyAlertConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifyAlertConfigurationResponse) SetStatusCode(v int32) *ModifyAlertConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAlertConfigurationResponse) SetBody(v *ModifyAlertConfigurationResponseBody) *ModifyAlertConfigurationResponse {
	s.Body = v
	return s
}

func (s *ModifyAlertConfigurationResponse) Validate() error {
	return dara.Validate(s)
}

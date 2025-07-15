// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApiConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApiConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApiConfigurationResponseBody) *ModifyApiConfigurationResponse
	GetBody() *ModifyApiConfigurationResponseBody
}

type ModifyApiConfigurationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApiConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApiConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyApiConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApiConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApiConfigurationResponse) GetBody() *ModifyApiConfigurationResponseBody {
	return s.Body
}

func (s *ModifyApiConfigurationResponse) SetHeaders(v map[string]*string) *ModifyApiConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifyApiConfigurationResponse) SetStatusCode(v int32) *ModifyApiConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApiConfigurationResponse) SetBody(v *ModifyApiConfigurationResponseBody) *ModifyApiConfigurationResponse {
	s.Body = v
	return s
}

func (s *ModifyApiConfigurationResponse) Validate() error {
	return dara.Validate(s)
}

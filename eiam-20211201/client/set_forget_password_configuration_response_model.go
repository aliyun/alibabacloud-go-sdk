// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetForgetPasswordConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetForgetPasswordConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetForgetPasswordConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *SetForgetPasswordConfigurationResponseBody) *SetForgetPasswordConfigurationResponse
	GetBody() *SetForgetPasswordConfigurationResponseBody
}

type SetForgetPasswordConfigurationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetForgetPasswordConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetForgetPasswordConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s SetForgetPasswordConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetForgetPasswordConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetForgetPasswordConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetForgetPasswordConfigurationResponse) GetBody() *SetForgetPasswordConfigurationResponseBody {
	return s.Body
}

func (s *SetForgetPasswordConfigurationResponse) SetHeaders(v map[string]*string) *SetForgetPasswordConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetForgetPasswordConfigurationResponse) SetStatusCode(v int32) *SetForgetPasswordConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetForgetPasswordConfigurationResponse) SetBody(v *SetForgetPasswordConfigurationResponseBody) *SetForgetPasswordConfigurationResponse {
	s.Body = v
	return s
}

func (s *SetForgetPasswordConfigurationResponse) Validate() error {
	return dara.Validate(s)
}

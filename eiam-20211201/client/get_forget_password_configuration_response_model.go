// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetForgetPasswordConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetForgetPasswordConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetForgetPasswordConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetForgetPasswordConfigurationResponseBody) *GetForgetPasswordConfigurationResponse
	GetBody() *GetForgetPasswordConfigurationResponseBody
}

type GetForgetPasswordConfigurationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetForgetPasswordConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetForgetPasswordConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetForgetPasswordConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetForgetPasswordConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetForgetPasswordConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetForgetPasswordConfigurationResponse) GetBody() *GetForgetPasswordConfigurationResponseBody {
	return s.Body
}

func (s *GetForgetPasswordConfigurationResponse) SetHeaders(v map[string]*string) *GetForgetPasswordConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetForgetPasswordConfigurationResponse) SetStatusCode(v int32) *GetForgetPasswordConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetForgetPasswordConfigurationResponse) SetBody(v *GetForgetPasswordConfigurationResponseBody) *GetForgetPasswordConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetForgetPasswordConfigurationResponse) Validate() error {
	return dara.Validate(s)
}

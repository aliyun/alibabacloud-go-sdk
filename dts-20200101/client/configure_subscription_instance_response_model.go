// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSubscriptionInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigureSubscriptionInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigureSubscriptionInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ConfigureSubscriptionInstanceResponseBody) *ConfigureSubscriptionInstanceResponse
	GetBody() *ConfigureSubscriptionInstanceResponseBody
}

type ConfigureSubscriptionInstanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigureSubscriptionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigureSubscriptionInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSubscriptionInstanceResponse) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigureSubscriptionInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigureSubscriptionInstanceResponse) GetBody() *ConfigureSubscriptionInstanceResponseBody {
	return s.Body
}

func (s *ConfigureSubscriptionInstanceResponse) SetHeaders(v map[string]*string) *ConfigureSubscriptionInstanceResponse {
	s.Headers = v
	return s
}

func (s *ConfigureSubscriptionInstanceResponse) SetStatusCode(v int32) *ConfigureSubscriptionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureSubscriptionInstanceResponse) SetBody(v *ConfigureSubscriptionInstanceResponseBody) *ConfigureSubscriptionInstanceResponse {
	s.Body = v
	return s
}

func (s *ConfigureSubscriptionInstanceResponse) Validate() error {
	return dara.Validate(s)
}

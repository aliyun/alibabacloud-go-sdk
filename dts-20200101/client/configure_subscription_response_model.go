// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigureSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigureSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *ConfigureSubscriptionResponseBody) *ConfigureSubscriptionResponse
	GetBody() *ConfigureSubscriptionResponseBody
}

type ConfigureSubscriptionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigureSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigureSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigureSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigureSubscriptionResponse) GetBody() *ConfigureSubscriptionResponseBody {
	return s.Body
}

func (s *ConfigureSubscriptionResponse) SetHeaders(v map[string]*string) *ConfigureSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *ConfigureSubscriptionResponse) SetStatusCode(v int32) *ConfigureSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureSubscriptionResponse) SetBody(v *ConfigureSubscriptionResponseBody) *ConfigureSubscriptionResponse {
	s.Body = v
	return s
}

func (s *ConfigureSubscriptionResponse) Validate() error {
	return dara.Validate(s)
}

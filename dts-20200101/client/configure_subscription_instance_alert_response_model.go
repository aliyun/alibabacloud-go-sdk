// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSubscriptionInstanceAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigureSubscriptionInstanceAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigureSubscriptionInstanceAlertResponse
	GetStatusCode() *int32
	SetBody(v *ConfigureSubscriptionInstanceAlertResponseBody) *ConfigureSubscriptionInstanceAlertResponse
	GetBody() *ConfigureSubscriptionInstanceAlertResponseBody
}

type ConfigureSubscriptionInstanceAlertResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigureSubscriptionInstanceAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigureSubscriptionInstanceAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSubscriptionInstanceAlertResponse) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigureSubscriptionInstanceAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigureSubscriptionInstanceAlertResponse) GetBody() *ConfigureSubscriptionInstanceAlertResponseBody {
	return s.Body
}

func (s *ConfigureSubscriptionInstanceAlertResponse) SetHeaders(v map[string]*string) *ConfigureSubscriptionInstanceAlertResponse {
	s.Headers = v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertResponse) SetStatusCode(v int32) *ConfigureSubscriptionInstanceAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertResponse) SetBody(v *ConfigureSubscriptionInstanceAlertResponseBody) *ConfigureSubscriptionInstanceAlertResponse {
	s.Body = v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

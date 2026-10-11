// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEgressProxyConfigurationAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEgressProxyConfigurationAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEgressProxyConfigurationAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEgressProxyConfigurationAttributeResponseBody) *UpdateEgressProxyConfigurationAttributeResponse
	GetBody() *UpdateEgressProxyConfigurationAttributeResponseBody
}

type UpdateEgressProxyConfigurationAttributeResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEgressProxyConfigurationAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEgressProxyConfigurationAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyConfigurationAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyConfigurationAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEgressProxyConfigurationAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEgressProxyConfigurationAttributeResponse) GetBody() *UpdateEgressProxyConfigurationAttributeResponseBody {
	return s.Body
}

func (s *UpdateEgressProxyConfigurationAttributeResponse) SetHeaders(v map[string]*string) *UpdateEgressProxyConfigurationAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeResponse) SetStatusCode(v int32) *UpdateEgressProxyConfigurationAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeResponse) SetBody(v *UpdateEgressProxyConfigurationAttributeResponseBody) *UpdateEgressProxyConfigurationAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

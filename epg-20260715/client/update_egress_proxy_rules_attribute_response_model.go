// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEgressProxyRulesAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEgressProxyRulesAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEgressProxyRulesAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEgressProxyRulesAttributeResponseBody) *UpdateEgressProxyRulesAttributeResponse
	GetBody() *UpdateEgressProxyRulesAttributeResponseBody
}

type UpdateEgressProxyRulesAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEgressProxyRulesAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEgressProxyRulesAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyRulesAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyRulesAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEgressProxyRulesAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEgressProxyRulesAttributeResponse) GetBody() *UpdateEgressProxyRulesAttributeResponseBody {
	return s.Body
}

func (s *UpdateEgressProxyRulesAttributeResponse) SetHeaders(v map[string]*string) *UpdateEgressProxyRulesAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateEgressProxyRulesAttributeResponse) SetStatusCode(v int32) *UpdateEgressProxyRulesAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEgressProxyRulesAttributeResponse) SetBody(v *UpdateEgressProxyRulesAttributeResponseBody) *UpdateEgressProxyRulesAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateEgressProxyRulesAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

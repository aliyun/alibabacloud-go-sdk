// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRuleAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRuleAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRuleAttributeResponseBody) *UpdateRuleAttributeResponse
	GetBody() *UpdateRuleAttributeResponseBody
}

type UpdateRuleAttributeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRuleAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRuleAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRuleAttributeResponse) GetBody() *UpdateRuleAttributeResponseBody {
	return s.Body
}

func (s *UpdateRuleAttributeResponse) SetHeaders(v map[string]*string) *UpdateRuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRuleAttributeResponse) SetStatusCode(v int32) *UpdateRuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRuleAttributeResponse) SetBody(v *UpdateRuleAttributeResponseBody) *UpdateRuleAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateRuleAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

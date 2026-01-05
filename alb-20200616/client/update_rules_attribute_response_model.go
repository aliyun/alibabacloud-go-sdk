// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRulesAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRulesAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRulesAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRulesAttributeResponseBody) *UpdateRulesAttributeResponse
	GetBody() *UpdateRulesAttributeResponseBody
}

type UpdateRulesAttributeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRulesAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRulesAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRulesAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRulesAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRulesAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRulesAttributeResponse) GetBody() *UpdateRulesAttributeResponseBody {
	return s.Body
}

func (s *UpdateRulesAttributeResponse) SetHeaders(v map[string]*string) *UpdateRulesAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRulesAttributeResponse) SetStatusCode(v int32) *UpdateRulesAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRulesAttributeResponse) SetBody(v *UpdateRulesAttributeResponseBody) *UpdateRulesAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateRulesAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

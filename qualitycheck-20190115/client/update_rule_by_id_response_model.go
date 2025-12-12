// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRuleByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRuleByIdResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRuleByIdResponseBody) *UpdateRuleByIdResponse
	GetBody() *UpdateRuleByIdResponseBody
}

type UpdateRuleByIdResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRuleByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRuleByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleByIdResponse) GoString() string {
	return s.String()
}

func (s *UpdateRuleByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRuleByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRuleByIdResponse) GetBody() *UpdateRuleByIdResponseBody {
	return s.Body
}

func (s *UpdateRuleByIdResponse) SetHeaders(v map[string]*string) *UpdateRuleByIdResponse {
	s.Headers = v
	return s
}

func (s *UpdateRuleByIdResponse) SetStatusCode(v int32) *UpdateRuleByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRuleByIdResponse) SetBody(v *UpdateRuleByIdResponseBody) *UpdateRuleByIdResponse {
	s.Body = v
	return s
}

func (s *UpdateRuleByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

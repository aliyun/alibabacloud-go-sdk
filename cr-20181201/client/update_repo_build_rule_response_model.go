// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepoBuildRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRepoBuildRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRepoBuildRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRepoBuildRuleResponseBody) *UpdateRepoBuildRuleResponse
	GetBody() *UpdateRepoBuildRuleResponseBody
}

type UpdateRepoBuildRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRepoBuildRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRepoBuildRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepoBuildRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRepoBuildRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRepoBuildRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRepoBuildRuleResponse) GetBody() *UpdateRepoBuildRuleResponseBody {
	return s.Body
}

func (s *UpdateRepoBuildRuleResponse) SetHeaders(v map[string]*string) *UpdateRepoBuildRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRepoBuildRuleResponse) SetStatusCode(v int32) *UpdateRepoBuildRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRepoBuildRuleResponse) SetBody(v *UpdateRepoBuildRuleResponseBody) *UpdateRepoBuildRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateRepoBuildRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

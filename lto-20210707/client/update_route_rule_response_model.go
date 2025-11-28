// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRouteRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRouteRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRouteRuleResponseBody) *UpdateRouteRuleResponse
	GetBody() *UpdateRouteRuleResponseBody
}

type UpdateRouteRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRouteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRouteRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRouteRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRouteRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRouteRuleResponse) GetBody() *UpdateRouteRuleResponseBody {
	return s.Body
}

func (s *UpdateRouteRuleResponse) SetHeaders(v map[string]*string) *UpdateRouteRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRouteRuleResponse) SetStatusCode(v int32) *UpdateRouteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRouteRuleResponse) SetBody(v *UpdateRouteRuleResponseBody) *UpdateRouteRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateRouteRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

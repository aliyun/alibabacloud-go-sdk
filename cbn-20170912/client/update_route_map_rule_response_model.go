// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteMapRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRouteMapRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRouteMapRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRouteMapRuleResponseBody) *UpdateRouteMapRuleResponse
	GetBody() *UpdateRouteMapRuleResponseBody
}

type UpdateRouteMapRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRouteMapRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRouteMapRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteMapRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRouteMapRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRouteMapRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRouteMapRuleResponse) GetBody() *UpdateRouteMapRuleResponseBody {
	return s.Body
}

func (s *UpdateRouteMapRuleResponse) SetHeaders(v map[string]*string) *UpdateRouteMapRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRouteMapRuleResponse) SetStatusCode(v int32) *UpdateRouteMapRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRouteMapRuleResponse) SetBody(v *UpdateRouteMapRuleResponseBody) *UpdateRouteMapRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateRouteMapRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

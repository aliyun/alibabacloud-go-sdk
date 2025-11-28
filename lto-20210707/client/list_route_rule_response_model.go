// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRouteRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRouteRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRouteRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListRouteRuleResponseBody) *ListRouteRuleResponse
	GetBody() *ListRouteRuleResponseBody
}

type ListRouteRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRouteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRouteRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRouteRuleResponse) GoString() string {
	return s.String()
}

func (s *ListRouteRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRouteRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRouteRuleResponse) GetBody() *ListRouteRuleResponseBody {
	return s.Body
}

func (s *ListRouteRuleResponse) SetHeaders(v map[string]*string) *ListRouteRuleResponse {
	s.Headers = v
	return s
}

func (s *ListRouteRuleResponse) SetStatusCode(v int32) *ListRouteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRouteRuleResponse) SetBody(v *ListRouteRuleResponseBody) *ListRouteRuleResponse {
	s.Body = v
	return s
}

func (s *ListRouteRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

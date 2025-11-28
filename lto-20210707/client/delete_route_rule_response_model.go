// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRouteRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRouteRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRouteRuleResponseBody) *DeleteRouteRuleResponse
	GetBody() *DeleteRouteRuleResponseBody
}

type DeleteRouteRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRouteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRouteRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRouteRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRouteRuleResponse) GetBody() *DeleteRouteRuleResponseBody {
	return s.Body
}

func (s *DeleteRouteRuleResponse) SetHeaders(v map[string]*string) *DeleteRouteRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteRuleResponse) SetStatusCode(v int32) *DeleteRouteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouteRuleResponse) SetBody(v *DeleteRouteRuleResponseBody) *DeleteRouteRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteRouteRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

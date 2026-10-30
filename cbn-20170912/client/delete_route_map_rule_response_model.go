// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteMapRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRouteMapRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRouteMapRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRouteMapRuleResponseBody) *DeleteRouteMapRuleResponse
	GetBody() *DeleteRouteMapRuleResponseBody
}

type DeleteRouteMapRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRouteMapRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRouteMapRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteMapRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteMapRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRouteMapRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRouteMapRuleResponse) GetBody() *DeleteRouteMapRuleResponseBody {
	return s.Body
}

func (s *DeleteRouteMapRuleResponse) SetHeaders(v map[string]*string) *DeleteRouteMapRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteMapRuleResponse) SetStatusCode(v int32) *DeleteRouteMapRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouteMapRuleResponse) SetBody(v *DeleteRouteMapRuleResponseBody) *DeleteRouteMapRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteRouteMapRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

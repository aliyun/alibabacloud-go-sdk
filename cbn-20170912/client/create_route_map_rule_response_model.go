// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteMapRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRouteMapRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRouteMapRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateRouteMapRuleResponseBody) *CreateRouteMapRuleResponse
	GetBody() *CreateRouteMapRuleResponseBody
}

type CreateRouteMapRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRouteMapRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRouteMapRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteMapRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRouteMapRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRouteMapRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRouteMapRuleResponse) GetBody() *CreateRouteMapRuleResponseBody {
	return s.Body
}

func (s *CreateRouteMapRuleResponse) SetHeaders(v map[string]*string) *CreateRouteMapRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRouteMapRuleResponse) SetStatusCode(v int32) *CreateRouteMapRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRouteMapRuleResponse) SetBody(v *CreateRouteMapRuleResponseBody) *CreateRouteMapRuleResponse {
	s.Body = v
	return s
}

func (s *CreateRouteMapRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

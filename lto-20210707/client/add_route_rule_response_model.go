// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRouteRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRouteRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRouteRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddRouteRuleResponseBody) *AddRouteRuleResponse
	GetBody() *AddRouteRuleResponseBody
}

type AddRouteRuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRouteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRouteRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRouteRuleResponse) GoString() string {
	return s.String()
}

func (s *AddRouteRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRouteRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRouteRuleResponse) GetBody() *AddRouteRuleResponseBody {
	return s.Body
}

func (s *AddRouteRuleResponse) SetHeaders(v map[string]*string) *AddRouteRuleResponse {
	s.Headers = v
	return s
}

func (s *AddRouteRuleResponse) SetStatusCode(v int32) *AddRouteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRouteRuleResponse) SetBody(v *AddRouteRuleResponseBody) *AddRouteRuleResponse {
	s.Body = v
	return s
}

func (s *AddRouteRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

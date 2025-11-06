// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebFlowRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWebFlowRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWebFlowRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWebFlowRulesResponseBody) *DeleteWebFlowRulesResponse
	GetBody() *DeleteWebFlowRulesResponseBody
}

type DeleteWebFlowRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWebFlowRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWebFlowRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebFlowRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebFlowRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWebFlowRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWebFlowRulesResponse) GetBody() *DeleteWebFlowRulesResponseBody {
	return s.Body
}

func (s *DeleteWebFlowRulesResponse) SetHeaders(v map[string]*string) *DeleteWebFlowRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebFlowRulesResponse) SetStatusCode(v int32) *DeleteWebFlowRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWebFlowRulesResponse) SetBody(v *DeleteWebFlowRulesResponseBody) *DeleteWebFlowRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteWebFlowRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

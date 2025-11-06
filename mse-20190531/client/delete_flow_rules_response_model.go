// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFlowRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFlowRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFlowRulesResponseBody) *DeleteFlowRulesResponse
	GetBody() *DeleteFlowRulesResponseBody
}

type DeleteFlowRulesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFlowRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFlowRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFlowRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFlowRulesResponse) GetBody() *DeleteFlowRulesResponseBody {
	return s.Body
}

func (s *DeleteFlowRulesResponse) SetHeaders(v map[string]*string) *DeleteFlowRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowRulesResponse) SetStatusCode(v int32) *DeleteFlowRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlowRulesResponse) SetBody(v *DeleteFlowRulesResponseBody) *DeleteFlowRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteFlowRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

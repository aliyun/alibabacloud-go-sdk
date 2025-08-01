// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebFlowRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWebFlowRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWebFlowRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListWebFlowRulesResponseBody) *ListWebFlowRulesResponse
	GetBody() *ListWebFlowRulesResponseBody
}

type ListWebFlowRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWebFlowRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWebFlowRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWebFlowRulesResponse) GoString() string {
	return s.String()
}

func (s *ListWebFlowRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWebFlowRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWebFlowRulesResponse) GetBody() *ListWebFlowRulesResponseBody {
	return s.Body
}

func (s *ListWebFlowRulesResponse) SetHeaders(v map[string]*string) *ListWebFlowRulesResponse {
	s.Headers = v
	return s
}

func (s *ListWebFlowRulesResponse) SetStatusCode(v int32) *ListWebFlowRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWebFlowRulesResponse) SetBody(v *ListWebFlowRulesResponseBody) *ListWebFlowRulesResponse {
	s.Body = v
	return s
}

func (s *ListWebFlowRulesResponse) Validate() error {
	return dara.Validate(s)
}

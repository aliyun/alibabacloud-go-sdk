// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFlowRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFlowRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListFlowRulesResponseBody) *ListFlowRulesResponse
	GetBody() *ListFlowRulesResponseBody
}

type ListFlowRulesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlowRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlowRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFlowRulesResponse) GoString() string {
	return s.String()
}

func (s *ListFlowRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFlowRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFlowRulesResponse) GetBody() *ListFlowRulesResponseBody {
	return s.Body
}

func (s *ListFlowRulesResponse) SetHeaders(v map[string]*string) *ListFlowRulesResponse {
	s.Headers = v
	return s
}

func (s *ListFlowRulesResponse) SetStatusCode(v int32) *ListFlowRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlowRulesResponse) SetBody(v *ListFlowRulesResponseBody) *ListFlowRulesResponse {
	s.Body = v
	return s
}

func (s *ListFlowRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

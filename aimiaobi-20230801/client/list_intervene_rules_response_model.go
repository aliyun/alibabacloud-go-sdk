// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterveneRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInterveneRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInterveneRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListInterveneRulesResponseBody) *ListInterveneRulesResponse
	GetBody() *ListInterveneRulesResponseBody
}

type ListInterveneRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterveneRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterveneRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInterveneRulesResponse) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInterveneRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInterveneRulesResponse) GetBody() *ListInterveneRulesResponseBody {
	return s.Body
}

func (s *ListInterveneRulesResponse) SetHeaders(v map[string]*string) *ListInterveneRulesResponse {
	s.Headers = v
	return s
}

func (s *ListInterveneRulesResponse) SetStatusCode(v int32) *ListInterveneRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterveneRulesResponse) SetBody(v *ListInterveneRulesResponseBody) *ListInterveneRulesResponse {
	s.Body = v
	return s
}

func (s *ListInterveneRulesResponse) Validate() error {
	return dara.Validate(s)
}

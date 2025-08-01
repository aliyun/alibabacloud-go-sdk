// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIsolationRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIsolationRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIsolationRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListIsolationRulesResponseBody) *ListIsolationRulesResponse
	GetBody() *ListIsolationRulesResponseBody
}

type ListIsolationRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIsolationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIsolationRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIsolationRulesResponse) GoString() string {
	return s.String()
}

func (s *ListIsolationRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIsolationRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIsolationRulesResponse) GetBody() *ListIsolationRulesResponseBody {
	return s.Body
}

func (s *ListIsolationRulesResponse) SetHeaders(v map[string]*string) *ListIsolationRulesResponse {
	s.Headers = v
	return s
}

func (s *ListIsolationRulesResponse) SetStatusCode(v int32) *ListIsolationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIsolationRulesResponse) SetBody(v *ListIsolationRulesResponseBody) *ListIsolationRulesResponse {
	s.Body = v
	return s
}

func (s *ListIsolationRulesResponse) Validate() error {
	return dara.Validate(s)
}

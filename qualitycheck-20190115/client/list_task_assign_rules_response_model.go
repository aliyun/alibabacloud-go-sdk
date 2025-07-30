// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskAssignRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskAssignRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskAssignRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskAssignRulesResponseBody) *ListTaskAssignRulesResponse
	GetBody() *ListTaskAssignRulesResponseBody
}

type ListTaskAssignRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskAssignRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskAssignRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskAssignRulesResponse) GoString() string {
	return s.String()
}

func (s *ListTaskAssignRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskAssignRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskAssignRulesResponse) GetBody() *ListTaskAssignRulesResponseBody {
	return s.Body
}

func (s *ListTaskAssignRulesResponse) SetHeaders(v map[string]*string) *ListTaskAssignRulesResponse {
	s.Headers = v
	return s
}

func (s *ListTaskAssignRulesResponse) SetStatusCode(v int32) *ListTaskAssignRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskAssignRulesResponse) SetBody(v *ListTaskAssignRulesResponseBody) *ListTaskAssignRulesResponse {
	s.Body = v
	return s
}

func (s *ListTaskAssignRulesResponse) Validate() error {
	return dara.Validate(s)
}

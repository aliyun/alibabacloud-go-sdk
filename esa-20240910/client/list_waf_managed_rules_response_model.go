// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafManagedRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWafManagedRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWafManagedRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListWafManagedRulesResponseBody) *ListWafManagedRulesResponse
	GetBody() *ListWafManagedRulesResponseBody
}

type ListWafManagedRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWafManagedRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWafManagedRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWafManagedRulesResponse) GoString() string {
	return s.String()
}

func (s *ListWafManagedRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWafManagedRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWafManagedRulesResponse) GetBody() *ListWafManagedRulesResponseBody {
	return s.Body
}

func (s *ListWafManagedRulesResponse) SetHeaders(v map[string]*string) *ListWafManagedRulesResponse {
	s.Headers = v
	return s
}

func (s *ListWafManagedRulesResponse) SetStatusCode(v int32) *ListWafManagedRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWafManagedRulesResponse) SetBody(v *ListWafManagedRulesResponseBody) *ListWafManagedRulesResponse {
	s.Body = v
	return s
}

func (s *ListWafManagedRulesResponse) Validate() error {
	return dara.Validate(s)
}

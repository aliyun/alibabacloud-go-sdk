// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemClientRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSystemClientRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSystemClientRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListSystemClientRulesResponseBody) *ListSystemClientRulesResponse
	GetBody() *ListSystemClientRulesResponseBody
}

type ListSystemClientRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSystemClientRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSystemClientRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSystemClientRulesResponse) GoString() string {
	return s.String()
}

func (s *ListSystemClientRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSystemClientRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSystemClientRulesResponse) GetBody() *ListSystemClientRulesResponseBody {
	return s.Body
}

func (s *ListSystemClientRulesResponse) SetHeaders(v map[string]*string) *ListSystemClientRulesResponse {
	s.Headers = v
	return s
}

func (s *ListSystemClientRulesResponse) SetStatusCode(v int32) *ListSystemClientRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSystemClientRulesResponse) SetBody(v *ListSystemClientRulesResponseBody) *ListSystemClientRulesResponse {
	s.Body = v
	return s
}

func (s *ListSystemClientRulesResponse) Validate() error {
	return dara.Validate(s)
}

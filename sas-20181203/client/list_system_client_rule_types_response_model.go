// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemClientRuleTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSystemClientRuleTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSystemClientRuleTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListSystemClientRuleTypesResponseBody) *ListSystemClientRuleTypesResponse
	GetBody() *ListSystemClientRuleTypesResponseBody
}

type ListSystemClientRuleTypesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSystemClientRuleTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSystemClientRuleTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSystemClientRuleTypesResponse) GoString() string {
	return s.String()
}

func (s *ListSystemClientRuleTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSystemClientRuleTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSystemClientRuleTypesResponse) GetBody() *ListSystemClientRuleTypesResponseBody {
	return s.Body
}

func (s *ListSystemClientRuleTypesResponse) SetHeaders(v map[string]*string) *ListSystemClientRuleTypesResponse {
	s.Headers = v
	return s
}

func (s *ListSystemClientRuleTypesResponse) SetStatusCode(v int32) *ListSystemClientRuleTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSystemClientRuleTypesResponse) SetBody(v *ListSystemClientRuleTypesResponseBody) *ListSystemClientRuleTypesResponse {
	s.Body = v
	return s
}

func (s *ListSystemClientRuleTypesResponse) Validate() error {
	return dara.Validate(s)
}

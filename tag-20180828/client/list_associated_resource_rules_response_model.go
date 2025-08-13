// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssociatedResourceRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAssociatedResourceRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAssociatedResourceRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListAssociatedResourceRulesResponseBody) *ListAssociatedResourceRulesResponse
	GetBody() *ListAssociatedResourceRulesResponseBody
}

type ListAssociatedResourceRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssociatedResourceRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssociatedResourceRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAssociatedResourceRulesResponse) GoString() string {
	return s.String()
}

func (s *ListAssociatedResourceRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAssociatedResourceRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAssociatedResourceRulesResponse) GetBody() *ListAssociatedResourceRulesResponseBody {
	return s.Body
}

func (s *ListAssociatedResourceRulesResponse) SetHeaders(v map[string]*string) *ListAssociatedResourceRulesResponse {
	s.Headers = v
	return s
}

func (s *ListAssociatedResourceRulesResponse) SetStatusCode(v int32) *ListAssociatedResourceRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssociatedResourceRulesResponse) SetBody(v *ListAssociatedResourceRulesResponseBody) *ListAssociatedResourceRulesResponse {
	s.Body = v
	return s
}

func (s *ListAssociatedResourceRulesResponse) Validate() error {
	return dara.Validate(s)
}

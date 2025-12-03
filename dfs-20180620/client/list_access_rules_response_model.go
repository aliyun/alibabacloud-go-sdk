// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccessRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccessRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListAccessRulesResponseBody) *ListAccessRulesResponse
	GetBody() *ListAccessRulesResponseBody
}

type ListAccessRulesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccessRulesResponse) GoString() string {
	return s.String()
}

func (s *ListAccessRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccessRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccessRulesResponse) GetBody() *ListAccessRulesResponseBody {
	return s.Body
}

func (s *ListAccessRulesResponse) SetHeaders(v map[string]*string) *ListAccessRulesResponse {
	s.Headers = v
	return s
}

func (s *ListAccessRulesResponse) SetStatusCode(v int32) *ListAccessRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessRulesResponse) SetBody(v *ListAccessRulesResponseBody) *ListAccessRulesResponse {
	s.Body = v
	return s
}

func (s *ListAccessRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

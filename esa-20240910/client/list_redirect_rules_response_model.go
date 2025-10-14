// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRedirectRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRedirectRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRedirectRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListRedirectRulesResponseBody) *ListRedirectRulesResponse
	GetBody() *ListRedirectRulesResponseBody
}

type ListRedirectRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRedirectRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRedirectRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRedirectRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRedirectRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRedirectRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRedirectRulesResponse) GetBody() *ListRedirectRulesResponseBody {
	return s.Body
}

func (s *ListRedirectRulesResponse) SetHeaders(v map[string]*string) *ListRedirectRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRedirectRulesResponse) SetStatusCode(v int32) *ListRedirectRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRedirectRulesResponse) SetBody(v *ListRedirectRulesResponseBody) *ListRedirectRulesResponse {
	s.Body = v
	return s
}

func (s *ListRedirectRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

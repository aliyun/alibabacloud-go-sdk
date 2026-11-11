// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListErrorPagesRedirectRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListErrorPagesRedirectRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListErrorPagesRedirectRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListErrorPagesRedirectRulesResponseBody) *ListErrorPagesRedirectRulesResponse
	GetBody() *ListErrorPagesRedirectRulesResponseBody
}

type ListErrorPagesRedirectRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListErrorPagesRedirectRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListErrorPagesRedirectRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListErrorPagesRedirectRulesResponse) GoString() string {
	return s.String()
}

func (s *ListErrorPagesRedirectRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListErrorPagesRedirectRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListErrorPagesRedirectRulesResponse) GetBody() *ListErrorPagesRedirectRulesResponseBody {
	return s.Body
}

func (s *ListErrorPagesRedirectRulesResponse) SetHeaders(v map[string]*string) *ListErrorPagesRedirectRulesResponse {
	s.Headers = v
	return s
}

func (s *ListErrorPagesRedirectRulesResponse) SetStatusCode(v int32) *ListErrorPagesRedirectRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListErrorPagesRedirectRulesResponse) SetBody(v *ListErrorPagesRedirectRulesResponseBody) *ListErrorPagesRedirectRulesResponse {
	s.Body = v
	return s
}

func (s *ListErrorPagesRedirectRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationRulesForApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizationRulesForApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizationRulesForApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizationRulesForApplicationResponseBody) *ListAuthorizationRulesForApplicationResponse
	GetBody() *ListAuthorizationRulesForApplicationResponseBody
}

type ListAuthorizationRulesForApplicationResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizationRulesForApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizationRulesForApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesForApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesForApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizationRulesForApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizationRulesForApplicationResponse) GetBody() *ListAuthorizationRulesForApplicationResponseBody {
	return s.Body
}

func (s *ListAuthorizationRulesForApplicationResponse) SetHeaders(v map[string]*string) *ListAuthorizationRulesForApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizationRulesForApplicationResponse) SetStatusCode(v int32) *ListAuthorizationRulesForApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizationRulesForApplicationResponse) SetBody(v *ListAuthorizationRulesForApplicationResponseBody) *ListAuthorizationRulesForApplicationResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizationRulesForApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

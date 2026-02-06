// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomResponseCodeRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomResponseCodeRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomResponseCodeRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomResponseCodeRulesResponseBody) *ListCustomResponseCodeRulesResponse
	GetBody() *ListCustomResponseCodeRulesResponseBody
}

type ListCustomResponseCodeRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomResponseCodeRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomResponseCodeRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomResponseCodeRulesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomResponseCodeRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomResponseCodeRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomResponseCodeRulesResponse) GetBody() *ListCustomResponseCodeRulesResponseBody {
	return s.Body
}

func (s *ListCustomResponseCodeRulesResponse) SetHeaders(v map[string]*string) *ListCustomResponseCodeRulesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomResponseCodeRulesResponse) SetStatusCode(v int32) *ListCustomResponseCodeRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomResponseCodeRulesResponse) SetBody(v *ListCustomResponseCodeRulesResponseBody) *ListCustomResponseCodeRulesResponse {
	s.Body = v
	return s
}

func (s *ListCustomResponseCodeRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

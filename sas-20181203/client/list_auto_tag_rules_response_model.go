// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoTagRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAutoTagRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAutoTagRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListAutoTagRulesResponseBody) *ListAutoTagRulesResponse
	GetBody() *ListAutoTagRulesResponseBody
}

type ListAutoTagRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutoTagRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutoTagRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAutoTagRulesResponse) GoString() string {
	return s.String()
}

func (s *ListAutoTagRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAutoTagRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAutoTagRulesResponse) GetBody() *ListAutoTagRulesResponseBody {
	return s.Body
}

func (s *ListAutoTagRulesResponse) SetHeaders(v map[string]*string) *ListAutoTagRulesResponse {
	s.Headers = v
	return s
}

func (s *ListAutoTagRulesResponse) SetStatusCode(v int32) *ListAutoTagRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoTagRulesResponse) SetBody(v *ListAutoTagRulesResponseBody) *ListAutoTagRulesResponse {
	s.Body = v
	return s
}

func (s *ListAutoTagRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

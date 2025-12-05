// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListRulesResponseBody) *ListRulesResponse
	GetBody() *ListRulesResponseBody
}

type ListRulesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRulesResponse) GetBody() *ListRulesResponseBody {
	return s.Body
}

func (s *ListRulesResponse) SetHeaders(v map[string]*string) *ListRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRulesResponse) SetStatusCode(v int32) *ListRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRulesResponse) SetBody(v *ListRulesResponseBody) *ListRulesResponse {
	s.Body = v
	return s
}

func (s *ListRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

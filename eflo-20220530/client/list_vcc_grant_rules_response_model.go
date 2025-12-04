// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVccGrantRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVccGrantRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVccGrantRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListVccGrantRulesResponseBody) *ListVccGrantRulesResponse
	GetBody() *ListVccGrantRulesResponseBody
}

type ListVccGrantRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVccGrantRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVccGrantRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVccGrantRulesResponse) GoString() string {
	return s.String()
}

func (s *ListVccGrantRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVccGrantRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVccGrantRulesResponse) GetBody() *ListVccGrantRulesResponseBody {
	return s.Body
}

func (s *ListVccGrantRulesResponse) SetHeaders(v map[string]*string) *ListVccGrantRulesResponse {
	s.Headers = v
	return s
}

func (s *ListVccGrantRulesResponse) SetStatusCode(v int32) *ListVccGrantRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVccGrantRulesResponse) SetBody(v *ListVccGrantRulesResponseBody) *ListVccGrantRulesResponse {
	s.Body = v
	return s
}

func (s *ListVccGrantRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

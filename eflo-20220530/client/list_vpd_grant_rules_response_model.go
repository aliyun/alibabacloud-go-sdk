// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpdGrantRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpdGrantRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpdGrantRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListVpdGrantRulesResponseBody) *ListVpdGrantRulesResponse
	GetBody() *ListVpdGrantRulesResponseBody
}

type ListVpdGrantRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpdGrantRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpdGrantRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpdGrantRulesResponse) GoString() string {
	return s.String()
}

func (s *ListVpdGrantRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpdGrantRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpdGrantRulesResponse) GetBody() *ListVpdGrantRulesResponseBody {
	return s.Body
}

func (s *ListVpdGrantRulesResponse) SetHeaders(v map[string]*string) *ListVpdGrantRulesResponse {
	s.Headers = v
	return s
}

func (s *ListVpdGrantRulesResponse) SetStatusCode(v int32) *ListVpdGrantRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpdGrantRulesResponse) SetBody(v *ListVpdGrantRulesResponseBody) *ListVpdGrantRulesResponse {
	s.Body = v
	return s
}

func (s *ListVpdGrantRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

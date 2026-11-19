// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMemoryRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMemoryRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListMemoryRulesResponseBody) *ListMemoryRulesResponse
	GetBody() *ListMemoryRulesResponseBody
}

type ListMemoryRulesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMemoryRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMemoryRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryRulesResponse) GoString() string {
	return s.String()
}

func (s *ListMemoryRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMemoryRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMemoryRulesResponse) GetBody() *ListMemoryRulesResponseBody {
	return s.Body
}

func (s *ListMemoryRulesResponse) SetHeaders(v map[string]*string) *ListMemoryRulesResponse {
	s.Headers = v
	return s
}

func (s *ListMemoryRulesResponse) SetStatusCode(v int32) *ListMemoryRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemoryRulesResponse) SetBody(v *ListMemoryRulesResponseBody) *ListMemoryRulesResponse {
	s.Body = v
	return s
}

func (s *ListMemoryRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

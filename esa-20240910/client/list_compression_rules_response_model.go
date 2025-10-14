// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompressionRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCompressionRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCompressionRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListCompressionRulesResponseBody) *ListCompressionRulesResponse
	GetBody() *ListCompressionRulesResponseBody
}

type ListCompressionRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCompressionRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCompressionRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCompressionRulesResponse) GoString() string {
	return s.String()
}

func (s *ListCompressionRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCompressionRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCompressionRulesResponse) GetBody() *ListCompressionRulesResponseBody {
	return s.Body
}

func (s *ListCompressionRulesResponse) SetHeaders(v map[string]*string) *ListCompressionRulesResponse {
	s.Headers = v
	return s
}

func (s *ListCompressionRulesResponse) SetStatusCode(v int32) *ListCompressionRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCompressionRulesResponse) SetBody(v *ListCompressionRulesResponseBody) *ListCompressionRulesResponse {
	s.Body = v
	return s
}

func (s *ListCompressionRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

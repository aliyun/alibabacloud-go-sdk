// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpRequestHeaderModificationRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHttpRequestHeaderModificationRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHttpRequestHeaderModificationRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListHttpRequestHeaderModificationRulesResponseBody) *ListHttpRequestHeaderModificationRulesResponse
	GetBody() *ListHttpRequestHeaderModificationRulesResponseBody
}

type ListHttpRequestHeaderModificationRulesResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHttpRequestHeaderModificationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHttpRequestHeaderModificationRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHttpRequestHeaderModificationRulesResponse) GoString() string {
	return s.String()
}

func (s *ListHttpRequestHeaderModificationRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHttpRequestHeaderModificationRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHttpRequestHeaderModificationRulesResponse) GetBody() *ListHttpRequestHeaderModificationRulesResponseBody {
	return s.Body
}

func (s *ListHttpRequestHeaderModificationRulesResponse) SetHeaders(v map[string]*string) *ListHttpRequestHeaderModificationRulesResponse {
	s.Headers = v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponse) SetStatusCode(v int32) *ListHttpRequestHeaderModificationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponse) SetBody(v *ListHttpRequestHeaderModificationRulesResponseBody) *ListHttpRequestHeaderModificationRulesResponse {
	s.Body = v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

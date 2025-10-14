// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpIncomingRequestHeaderModificationRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHttpIncomingRequestHeaderModificationRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHttpIncomingRequestHeaderModificationRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListHttpIncomingRequestHeaderModificationRulesResponseBody) *ListHttpIncomingRequestHeaderModificationRulesResponse
	GetBody() *ListHttpIncomingRequestHeaderModificationRulesResponseBody
}

type ListHttpIncomingRequestHeaderModificationRulesResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHttpIncomingRequestHeaderModificationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHttpIncomingRequestHeaderModificationRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHttpIncomingRequestHeaderModificationRulesResponse) GoString() string {
	return s.String()
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponse) GetBody() *ListHttpIncomingRequestHeaderModificationRulesResponseBody {
	return s.Body
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponse) SetHeaders(v map[string]*string) *ListHttpIncomingRequestHeaderModificationRulesResponse {
	s.Headers = v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponse) SetStatusCode(v int32) *ListHttpIncomingRequestHeaderModificationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponse) SetBody(v *ListHttpIncomingRequestHeaderModificationRulesResponseBody) *ListHttpIncomingRequestHeaderModificationRulesResponse {
	s.Body = v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

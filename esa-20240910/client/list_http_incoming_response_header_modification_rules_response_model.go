// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpIncomingResponseHeaderModificationRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHttpIncomingResponseHeaderModificationRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHttpIncomingResponseHeaderModificationRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListHttpIncomingResponseHeaderModificationRulesResponseBody) *ListHttpIncomingResponseHeaderModificationRulesResponse
	GetBody() *ListHttpIncomingResponseHeaderModificationRulesResponseBody
}

type ListHttpIncomingResponseHeaderModificationRulesResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHttpIncomingResponseHeaderModificationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHttpIncomingResponseHeaderModificationRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHttpIncomingResponseHeaderModificationRulesResponse) GoString() string {
	return s.String()
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponse) GetBody() *ListHttpIncomingResponseHeaderModificationRulesResponseBody {
	return s.Body
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponse) SetHeaders(v map[string]*string) *ListHttpIncomingResponseHeaderModificationRulesResponse {
	s.Headers = v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponse) SetStatusCode(v int32) *ListHttpIncomingResponseHeaderModificationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponse) SetBody(v *ListHttpIncomingResponseHeaderModificationRulesResponseBody) *ListHttpIncomingResponseHeaderModificationRulesResponse {
	s.Body = v
	return s
}

func (s *ListHttpIncomingResponseHeaderModificationRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecuritySuggestionNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecuritySuggestionNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecuritySuggestionNumberResponse
	GetStatusCode() *int32
	SetBody(v *GetSecuritySuggestionNumberResponseBody) *GetSecuritySuggestionNumberResponse
	GetBody() *GetSecuritySuggestionNumberResponseBody
}

type GetSecuritySuggestionNumberResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecuritySuggestionNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecuritySuggestionNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySuggestionNumberResponse) GoString() string {
	return s.String()
}

func (s *GetSecuritySuggestionNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecuritySuggestionNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecuritySuggestionNumberResponse) GetBody() *GetSecuritySuggestionNumberResponseBody {
	return s.Body
}

func (s *GetSecuritySuggestionNumberResponse) SetHeaders(v map[string]*string) *GetSecuritySuggestionNumberResponse {
	s.Headers = v
	return s
}

func (s *GetSecuritySuggestionNumberResponse) SetStatusCode(v int32) *GetSecuritySuggestionNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecuritySuggestionNumberResponse) SetBody(v *GetSecuritySuggestionNumberResponseBody) *GetSecuritySuggestionNumberResponse {
	s.Body = v
	return s
}

func (s *GetSecuritySuggestionNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

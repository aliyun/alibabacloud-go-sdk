// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecuritySuggestionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecuritySuggestionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecuritySuggestionListResponse
	GetStatusCode() *int32
	SetBody(v *GetSecuritySuggestionListResponseBody) *GetSecuritySuggestionListResponse
	GetBody() *GetSecuritySuggestionListResponseBody
}

type GetSecuritySuggestionListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecuritySuggestionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecuritySuggestionListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySuggestionListResponse) GoString() string {
	return s.String()
}

func (s *GetSecuritySuggestionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecuritySuggestionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecuritySuggestionListResponse) GetBody() *GetSecuritySuggestionListResponseBody {
	return s.Body
}

func (s *GetSecuritySuggestionListResponse) SetHeaders(v map[string]*string) *GetSecuritySuggestionListResponse {
	s.Headers = v
	return s
}

func (s *GetSecuritySuggestionListResponse) SetStatusCode(v int32) *GetSecuritySuggestionListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecuritySuggestionListResponse) SetBody(v *GetSecuritySuggestionListResponseBody) *GetSecuritySuggestionListResponse {
	s.Body = v
	return s
}

func (s *GetSecuritySuggestionListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

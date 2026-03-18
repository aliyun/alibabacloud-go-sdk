// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantArticlesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTextbookAssistantArticlesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTextbookAssistantArticlesResponse
	GetStatusCode() *int32
	SetBody(v *ListTextbookAssistantArticlesResponseBody) *ListTextbookAssistantArticlesResponse
	GetBody() *ListTextbookAssistantArticlesResponseBody
}

type ListTextbookAssistantArticlesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTextbookAssistantArticlesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextbookAssistantArticlesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantArticlesResponse) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticlesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTextbookAssistantArticlesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTextbookAssistantArticlesResponse) GetBody() *ListTextbookAssistantArticlesResponseBody {
	return s.Body
}

func (s *ListTextbookAssistantArticlesResponse) SetHeaders(v map[string]*string) *ListTextbookAssistantArticlesResponse {
	s.Headers = v
	return s
}

func (s *ListTextbookAssistantArticlesResponse) SetStatusCode(v int32) *ListTextbookAssistantArticlesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextbookAssistantArticlesResponse) SetBody(v *ListTextbookAssistantArticlesResponseBody) *ListTextbookAssistantArticlesResponse {
	s.Body = v
	return s
}

func (s *ListTextbookAssistantArticlesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

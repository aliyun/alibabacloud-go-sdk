// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantArticleDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTextbookAssistantArticleDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTextbookAssistantArticleDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListTextbookAssistantArticleDetailsResponseBody) *ListTextbookAssistantArticleDetailsResponse
	GetBody() *ListTextbookAssistantArticleDetailsResponseBody
}

type ListTextbookAssistantArticleDetailsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTextbookAssistantArticleDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTextbookAssistantArticleDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTextbookAssistantArticleDetailsResponse) GetBody() *ListTextbookAssistantArticleDetailsResponseBody {
	return s.Body
}

func (s *ListTextbookAssistantArticleDetailsResponse) SetHeaders(v map[string]*string) *ListTextbookAssistantArticleDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponse) SetStatusCode(v int32) *ListTextbookAssistantArticleDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponse) SetBody(v *ListTextbookAssistantArticleDetailsResponseBody) *ListTextbookAssistantArticleDetailsResponse {
	s.Body = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

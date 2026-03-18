// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantSceneDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTextbookAssistantSceneDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTextbookAssistantSceneDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListTextbookAssistantSceneDetailsResponseBody) *ListTextbookAssistantSceneDetailsResponse
	GetBody() *ListTextbookAssistantSceneDetailsResponseBody
}

type ListTextbookAssistantSceneDetailsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTextbookAssistantSceneDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextbookAssistantSceneDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTextbookAssistantSceneDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTextbookAssistantSceneDetailsResponse) GetBody() *ListTextbookAssistantSceneDetailsResponseBody {
	return s.Body
}

func (s *ListTextbookAssistantSceneDetailsResponse) SetHeaders(v map[string]*string) *ListTextbookAssistantSceneDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponse) SetStatusCode(v int32) *ListTextbookAssistantSceneDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponse) SetBody(v *ListTextbookAssistantSceneDetailsResponseBody) *ListTextbookAssistantSceneDetailsResponse {
	s.Body = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

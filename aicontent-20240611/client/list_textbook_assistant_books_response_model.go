// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantBooksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTextbookAssistantBooksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTextbookAssistantBooksResponse
	GetStatusCode() *int32
	SetBody(v *ListTextbookAssistantBooksResponseBody) *ListTextbookAssistantBooksResponse
	GetBody() *ListTextbookAssistantBooksResponseBody
}

type ListTextbookAssistantBooksResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTextbookAssistantBooksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextbookAssistantBooksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBooksResponse) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBooksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTextbookAssistantBooksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTextbookAssistantBooksResponse) GetBody() *ListTextbookAssistantBooksResponseBody {
	return s.Body
}

func (s *ListTextbookAssistantBooksResponse) SetHeaders(v map[string]*string) *ListTextbookAssistantBooksResponse {
	s.Headers = v
	return s
}

func (s *ListTextbookAssistantBooksResponse) SetStatusCode(v int32) *ListTextbookAssistantBooksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextbookAssistantBooksResponse) SetBody(v *ListTextbookAssistantBooksResponseBody) *ListTextbookAssistantBooksResponse {
	s.Body = v
	return s
}

func (s *ListTextbookAssistantBooksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

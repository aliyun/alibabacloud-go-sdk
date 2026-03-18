// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantBookDirectoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTextbookAssistantBookDirectoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTextbookAssistantBookDirectoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListTextbookAssistantBookDirectoriesResponseBody) *ListTextbookAssistantBookDirectoriesResponse
	GetBody() *ListTextbookAssistantBookDirectoriesResponseBody
}

type ListTextbookAssistantBookDirectoriesResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTextbookAssistantBookDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextbookAssistantBookDirectoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTextbookAssistantBookDirectoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTextbookAssistantBookDirectoriesResponse) GetBody() *ListTextbookAssistantBookDirectoriesResponseBody {
	return s.Body
}

func (s *ListTextbookAssistantBookDirectoriesResponse) SetHeaders(v map[string]*string) *ListTextbookAssistantBookDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponse) SetStatusCode(v int32) *ListTextbookAssistantBookDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponse) SetBody(v *ListTextbookAssistantBookDirectoriesResponseBody) *ListTextbookAssistantBookDirectoriesResponse {
	s.Body = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

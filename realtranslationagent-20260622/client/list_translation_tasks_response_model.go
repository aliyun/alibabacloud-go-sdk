// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTranslationTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTranslationTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTranslationTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListTranslationTasksResponseBody) *ListTranslationTasksResponse
	GetBody() *ListTranslationTasksResponseBody
}

type ListTranslationTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTranslationTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTranslationTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTranslationTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTranslationTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTranslationTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTranslationTasksResponse) GetBody() *ListTranslationTasksResponseBody {
	return s.Body
}

func (s *ListTranslationTasksResponse) SetHeaders(v map[string]*string) *ListTranslationTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTranslationTasksResponse) SetStatusCode(v int32) *ListTranslationTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTranslationTasksResponse) SetBody(v *ListTranslationTasksResponseBody) *ListTranslationTasksResponse {
	s.Body = v
	return s
}

func (s *ListTranslationTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGlobalQuestionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGlobalQuestionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGlobalQuestionsResponse
	GetStatusCode() *int32
	SetBody(v *ListGlobalQuestionsResponseBody) *ListGlobalQuestionsResponse
	GetBody() *ListGlobalQuestionsResponseBody
}

type ListGlobalQuestionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGlobalQuestionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGlobalQuestionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGlobalQuestionsResponse) GoString() string {
	return s.String()
}

func (s *ListGlobalQuestionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGlobalQuestionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGlobalQuestionsResponse) GetBody() *ListGlobalQuestionsResponseBody {
	return s.Body
}

func (s *ListGlobalQuestionsResponse) SetHeaders(v map[string]*string) *ListGlobalQuestionsResponse {
	s.Headers = v
	return s
}

func (s *ListGlobalQuestionsResponse) SetStatusCode(v int32) *ListGlobalQuestionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGlobalQuestionsResponse) SetBody(v *ListGlobalQuestionsResponseBody) *ListGlobalQuestionsResponse {
	s.Body = v
	return s
}

func (s *ListGlobalQuestionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

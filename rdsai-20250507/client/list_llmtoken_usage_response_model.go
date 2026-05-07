// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLLMTokenUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLLMTokenUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLLMTokenUsageResponse
	GetStatusCode() *int32
	SetBody(v *ListLLMTokenUsageResponseBody) *ListLLMTokenUsageResponse
	GetBody() *ListLLMTokenUsageResponseBody
}

type ListLLMTokenUsageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLLMTokenUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLLMTokenUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLLMTokenUsageResponse) GoString() string {
	return s.String()
}

func (s *ListLLMTokenUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLLMTokenUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLLMTokenUsageResponse) GetBody() *ListLLMTokenUsageResponseBody {
	return s.Body
}

func (s *ListLLMTokenUsageResponse) SetHeaders(v map[string]*string) *ListLLMTokenUsageResponse {
	s.Headers = v
	return s
}

func (s *ListLLMTokenUsageResponse) SetStatusCode(v int32) *ListLLMTokenUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLLMTokenUsageResponse) SetBody(v *ListLLMTokenUsageResponseBody) *ListLLMTokenUsageResponse {
	s.Body = v
	return s
}

func (s *ListLLMTokenUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

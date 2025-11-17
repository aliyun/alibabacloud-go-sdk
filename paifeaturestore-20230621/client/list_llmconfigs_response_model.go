// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLLMConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLLMConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLLMConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListLLMConfigsResponseBody) *ListLLMConfigsResponse
	GetBody() *ListLLMConfigsResponseBody
}

type ListLLMConfigsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLLMConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLLMConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLLMConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListLLMConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLLMConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLLMConfigsResponse) GetBody() *ListLLMConfigsResponseBody {
	return s.Body
}

func (s *ListLLMConfigsResponse) SetHeaders(v map[string]*string) *ListLLMConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListLLMConfigsResponse) SetStatusCode(v int32) *ListLLMConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLLMConfigsResponse) SetBody(v *ListLLMConfigsResponseBody) *ListLLMConfigsResponse {
	s.Body = v
	return s
}

func (s *ListLLMConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromptTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPromptTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPromptTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListPromptTemplatesResponseBody) *ListPromptTemplatesResponse
	GetBody() *ListPromptTemplatesResponseBody
}

type ListPromptTemplatesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPromptTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPromptTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPromptTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListPromptTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPromptTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPromptTemplatesResponse) GetBody() *ListPromptTemplatesResponseBody {
	return s.Body
}

func (s *ListPromptTemplatesResponse) SetHeaders(v map[string]*string) *ListPromptTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListPromptTemplatesResponse) SetStatusCode(v int32) *ListPromptTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPromptTemplatesResponse) SetBody(v *ListPromptTemplatesResponseBody) *ListPromptTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListPromptTemplatesResponse) Validate() error {
	return dara.Validate(s)
}

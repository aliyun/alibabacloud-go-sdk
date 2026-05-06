// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLlmTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLlmTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLlmTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListLlmTemplatesResponseBody) *ListLlmTemplatesResponse
	GetBody() *ListLlmTemplatesResponseBody
}

type ListLlmTemplatesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLlmTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLlmTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLlmTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListLlmTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLlmTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLlmTemplatesResponse) GetBody() *ListLlmTemplatesResponseBody {
	return s.Body
}

func (s *ListLlmTemplatesResponse) SetHeaders(v map[string]*string) *ListLlmTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListLlmTemplatesResponse) SetStatusCode(v int32) *ListLlmTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLlmTemplatesResponse) SetBody(v *ListLlmTemplatesResponseBody) *ListLlmTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListLlmTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

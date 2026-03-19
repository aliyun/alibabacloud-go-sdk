// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPptTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPptTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPptTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListPptTemplatesResponseBody) *ListPptTemplatesResponse
	GetBody() *ListPptTemplatesResponseBody
}

type ListPptTemplatesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPptTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPptTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPptTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListPptTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPptTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPptTemplatesResponse) GetBody() *ListPptTemplatesResponseBody {
	return s.Body
}

func (s *ListPptTemplatesResponse) SetHeaders(v map[string]*string) *ListPptTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListPptTemplatesResponse) SetStatusCode(v int32) *ListPptTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPptTemplatesResponse) SetBody(v *ListPptTemplatesResponseBody) *ListPptTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListPptTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

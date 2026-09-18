// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDistillationTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDistillationTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDistillationTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListDistillationTemplatesResponseBody) *ListDistillationTemplatesResponse
	GetBody() *ListDistillationTemplatesResponseBody
}

type ListDistillationTemplatesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDistillationTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDistillationTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDistillationTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListDistillationTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDistillationTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDistillationTemplatesResponse) GetBody() *ListDistillationTemplatesResponseBody {
	return s.Body
}

func (s *ListDistillationTemplatesResponse) SetHeaders(v map[string]*string) *ListDistillationTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListDistillationTemplatesResponse) SetStatusCode(v int32) *ListDistillationTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDistillationTemplatesResponse) SetBody(v *ListDistillationTemplatesResponseBody) *ListDistillationTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListDistillationTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

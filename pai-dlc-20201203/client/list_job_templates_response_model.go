// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListJobTemplatesResponseBody) *ListJobTemplatesResponse
	GetBody() *ListJobTemplatesResponseBody
}

type ListJobTemplatesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobTemplatesResponse) GetBody() *ListJobTemplatesResponseBody {
	return s.Body
}

func (s *ListJobTemplatesResponse) SetHeaders(v map[string]*string) *ListJobTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListJobTemplatesResponse) SetStatusCode(v int32) *ListJobTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobTemplatesResponse) SetBody(v *ListJobTemplatesResponseBody) *ListJobTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListJobTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

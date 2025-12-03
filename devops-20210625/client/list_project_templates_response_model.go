// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectTemplatesResponseBody) *ListProjectTemplatesResponse
	GetBody() *ListProjectTemplatesResponseBody
}

type ListProjectTemplatesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListProjectTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectTemplatesResponse) GetBody() *ListProjectTemplatesResponseBody {
	return s.Body
}

func (s *ListProjectTemplatesResponse) SetHeaders(v map[string]*string) *ListProjectTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListProjectTemplatesResponse) SetStatusCode(v int32) *ListProjectTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectTemplatesResponse) SetBody(v *ListProjectTemplatesResponseBody) *ListProjectTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListProjectTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

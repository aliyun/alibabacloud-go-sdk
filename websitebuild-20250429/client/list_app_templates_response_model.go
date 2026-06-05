// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListAppTemplatesResponseBody) *ListAppTemplatesResponse
	GetBody() *ListAppTemplatesResponseBody
}

type ListAppTemplatesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppTemplatesResponse) GetBody() *ListAppTemplatesResponseBody {
	return s.Body
}

func (s *ListAppTemplatesResponse) SetHeaders(v map[string]*string) *ListAppTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListAppTemplatesResponse) SetStatusCode(v int32) *ListAppTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppTemplatesResponse) SetBody(v *ListAppTemplatesResponseBody) *ListAppTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListAppTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

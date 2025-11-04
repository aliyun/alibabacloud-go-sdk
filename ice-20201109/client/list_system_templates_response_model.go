// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSystemTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSystemTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListSystemTemplatesResponseBody) *ListSystemTemplatesResponse
	GetBody() *ListSystemTemplatesResponseBody
}

type ListSystemTemplatesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSystemTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSystemTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSystemTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListSystemTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSystemTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSystemTemplatesResponse) GetBody() *ListSystemTemplatesResponseBody {
	return s.Body
}

func (s *ListSystemTemplatesResponse) SetHeaders(v map[string]*string) *ListSystemTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListSystemTemplatesResponse) SetStatusCode(v int32) *ListSystemTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSystemTemplatesResponse) SetBody(v *ListSystemTemplatesResponseBody) *ListSystemTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListSystemTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

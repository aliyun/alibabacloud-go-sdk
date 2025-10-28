// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConfigTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConfigTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListConfigTemplatesResponseBody) *ListConfigTemplatesResponse
	GetBody() *ListConfigTemplatesResponseBody
}

type ListConfigTemplatesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConfigTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConfigTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConfigTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListConfigTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConfigTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConfigTemplatesResponse) GetBody() *ListConfigTemplatesResponseBody {
	return s.Body
}

func (s *ListConfigTemplatesResponse) SetHeaders(v map[string]*string) *ListConfigTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListConfigTemplatesResponse) SetStatusCode(v int32) *ListConfigTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigTemplatesResponse) SetBody(v *ListConfigTemplatesResponseBody) *ListConfigTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListConfigTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

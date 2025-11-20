// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListTemplatesResult) *ListTemplatesResponse
	GetBody() *ListTemplatesResult
}

type ListTemplatesResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTemplatesResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTemplatesResponse) GetBody() *ListTemplatesResult {
	return s.Body
}

func (s *ListTemplatesResponse) SetHeaders(v map[string]*string) *ListTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListTemplatesResponse) SetStatusCode(v int32) *ListTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplatesResponse) SetBody(v *ListTemplatesResult) *ListTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

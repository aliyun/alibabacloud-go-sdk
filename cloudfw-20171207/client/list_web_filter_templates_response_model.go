// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebFilterTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWebFilterTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWebFilterTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListWebFilterTemplatesResponseBody) *ListWebFilterTemplatesResponse
	GetBody() *ListWebFilterTemplatesResponseBody
}

type ListWebFilterTemplatesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWebFilterTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWebFilterTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWebFilterTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListWebFilterTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWebFilterTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWebFilterTemplatesResponse) GetBody() *ListWebFilterTemplatesResponseBody {
	return s.Body
}

func (s *ListWebFilterTemplatesResponse) SetHeaders(v map[string]*string) *ListWebFilterTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListWebFilterTemplatesResponse) SetStatusCode(v int32) *ListWebFilterTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWebFilterTemplatesResponse) SetBody(v *ListWebFilterTemplatesResponseBody) *ListWebFilterTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListWebFilterTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

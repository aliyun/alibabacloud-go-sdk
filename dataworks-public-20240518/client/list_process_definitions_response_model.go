// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProcessDefinitionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProcessDefinitionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProcessDefinitionsResponse
	GetStatusCode() *int32
	SetBody(v *ListProcessDefinitionsResponseBody) *ListProcessDefinitionsResponse
	GetBody() *ListProcessDefinitionsResponseBody
}

type ListProcessDefinitionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProcessDefinitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProcessDefinitionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProcessDefinitionsResponse) GoString() string {
	return s.String()
}

func (s *ListProcessDefinitionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProcessDefinitionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProcessDefinitionsResponse) GetBody() *ListProcessDefinitionsResponseBody {
	return s.Body
}

func (s *ListProcessDefinitionsResponse) SetHeaders(v map[string]*string) *ListProcessDefinitionsResponse {
	s.Headers = v
	return s
}

func (s *ListProcessDefinitionsResponse) SetStatusCode(v int32) *ListProcessDefinitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProcessDefinitionsResponse) SetBody(v *ListProcessDefinitionsResponseBody) *ListProcessDefinitionsResponse {
	s.Body = v
	return s
}

func (s *ListProcessDefinitionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

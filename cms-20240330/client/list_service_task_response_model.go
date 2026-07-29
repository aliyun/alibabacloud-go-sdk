// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceTaskResponseBody) *ListServiceTaskResponse
	GetBody() *ListServiceTaskResponseBody
}

type ListServiceTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTaskResponse) GoString() string {
	return s.String()
}

func (s *ListServiceTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceTaskResponse) GetBody() *ListServiceTaskResponseBody {
	return s.Body
}

func (s *ListServiceTaskResponse) SetHeaders(v map[string]*string) *ListServiceTaskResponse {
	s.Headers = v
	return s
}

func (s *ListServiceTaskResponse) SetStatusCode(v int32) *ListServiceTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceTaskResponse) SetBody(v *ListServiceTaskResponseBody) *ListServiceTaskResponse {
	s.Body = v
	return s
}

func (s *ListServiceTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

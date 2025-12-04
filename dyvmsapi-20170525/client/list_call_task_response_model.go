// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCallTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCallTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListCallTaskResponseBody) *ListCallTaskResponse
	GetBody() *ListCallTaskResponseBody
}

type ListCallTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCallTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCallTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCallTaskResponse) GoString() string {
	return s.String()
}

func (s *ListCallTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCallTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCallTaskResponse) GetBody() *ListCallTaskResponseBody {
	return s.Body
}

func (s *ListCallTaskResponse) SetHeaders(v map[string]*string) *ListCallTaskResponse {
	s.Headers = v
	return s
}

func (s *ListCallTaskResponse) SetStatusCode(v int32) *ListCallTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCallTaskResponse) SetBody(v *ListCallTaskResponseBody) *ListCallTaskResponse {
	s.Body = v
	return s
}

func (s *ListCallTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

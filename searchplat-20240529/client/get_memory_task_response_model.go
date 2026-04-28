// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMemoryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMemoryTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetMemoryTaskResponseBody) *GetMemoryTaskResponse
	GetBody() *GetMemoryTaskResponseBody
}

type GetMemoryTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemoryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemoryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryTaskResponse) GoString() string {
	return s.String()
}

func (s *GetMemoryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMemoryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMemoryTaskResponse) GetBody() *GetMemoryTaskResponseBody {
	return s.Body
}

func (s *GetMemoryTaskResponse) SetHeaders(v map[string]*string) *GetMemoryTaskResponse {
	s.Headers = v
	return s
}

func (s *GetMemoryTaskResponse) SetStatusCode(v int32) *GetMemoryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemoryTaskResponse) SetBody(v *GetMemoryTaskResponseBody) *GetMemoryTaskResponse {
	s.Body = v
	return s
}

func (s *GetMemoryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

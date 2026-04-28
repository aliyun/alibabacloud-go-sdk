// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryHealthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMemoryHealthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMemoryHealthResponse
	GetStatusCode() *int32
	SetBody(v *GetMemoryHealthResponseBody) *GetMemoryHealthResponse
	GetBody() *GetMemoryHealthResponseBody
}

type GetMemoryHealthResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemoryHealthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemoryHealthResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryHealthResponse) GoString() string {
	return s.String()
}

func (s *GetMemoryHealthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMemoryHealthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMemoryHealthResponse) GetBody() *GetMemoryHealthResponseBody {
	return s.Body
}

func (s *GetMemoryHealthResponse) SetHeaders(v map[string]*string) *GetMemoryHealthResponse {
	s.Headers = v
	return s
}

func (s *GetMemoryHealthResponse) SetStatusCode(v int32) *GetMemoryHealthResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemoryHealthResponse) SetBody(v *GetMemoryHealthResponseBody) *GetMemoryHealthResponse {
	s.Body = v
	return s
}

func (s *GetMemoryHealthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

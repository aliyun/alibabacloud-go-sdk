// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemorySessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMemorySessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMemorySessionResponse
	GetStatusCode() *int32
	SetBody(v *GetMemorySessionResponseBody) *GetMemorySessionResponse
	GetBody() *GetMemorySessionResponseBody
}

type GetMemorySessionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemorySessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemorySessionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMemorySessionResponse) GoString() string {
	return s.String()
}

func (s *GetMemorySessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMemorySessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMemorySessionResponse) GetBody() *GetMemorySessionResponseBody {
	return s.Body
}

func (s *GetMemorySessionResponse) SetHeaders(v map[string]*string) *GetMemorySessionResponse {
	s.Headers = v
	return s
}

func (s *GetMemorySessionResponse) SetStatusCode(v int32) *GetMemorySessionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemorySessionResponse) SetBody(v *GetMemorySessionResponseBody) *GetMemorySessionResponse {
	s.Body = v
	return s
}

func (s *GetMemorySessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

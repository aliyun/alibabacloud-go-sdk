// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMemoryResponse
	GetStatusCode() *int32
	SetBody(v *GetMemoryResponseBody) *GetMemoryResponse
	GetBody() *GetMemoryResponseBody
}

type GetMemoryResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryResponse) GoString() string {
	return s.String()
}

func (s *GetMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMemoryResponse) GetBody() *GetMemoryResponseBody {
	return s.Body
}

func (s *GetMemoryResponse) SetHeaders(v map[string]*string) *GetMemoryResponse {
	s.Headers = v
	return s
}

func (s *GetMemoryResponse) SetStatusCode(v int32) *GetMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemoryResponse) SetBody(v *GetMemoryResponseBody) *GetMemoryResponse {
	s.Body = v
	return s
}

func (s *GetMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

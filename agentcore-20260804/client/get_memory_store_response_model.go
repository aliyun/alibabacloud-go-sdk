// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMemoryStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMemoryStoreResponse
	GetStatusCode() *int32
	SetBody(v *GetMemoryStoreResponseBody) *GetMemoryStoreResponse
	GetBody() *GetMemoryStoreResponseBody
}

type GetMemoryStoreResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemoryStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemoryStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryStoreResponse) GoString() string {
	return s.String()
}

func (s *GetMemoryStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMemoryStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMemoryStoreResponse) GetBody() *GetMemoryStoreResponseBody {
	return s.Body
}

func (s *GetMemoryStoreResponse) SetHeaders(v map[string]*string) *GetMemoryStoreResponse {
	s.Headers = v
	return s
}

func (s *GetMemoryStoreResponse) SetStatusCode(v int32) *GetMemoryStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemoryStoreResponse) SetBody(v *GetMemoryStoreResponseBody) *GetMemoryStoreResponse {
	s.Body = v
	return s
}

func (s *GetMemoryStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

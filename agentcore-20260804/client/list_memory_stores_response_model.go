// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryStoresResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMemoryStoresResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMemoryStoresResponse
	GetStatusCode() *int32
	SetBody(v *ListMemoryStoresResponseBody) *ListMemoryStoresResponse
	GetBody() *ListMemoryStoresResponseBody
}

type ListMemoryStoresResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMemoryStoresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMemoryStoresResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryStoresResponse) GoString() string {
	return s.String()
}

func (s *ListMemoryStoresResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMemoryStoresResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMemoryStoresResponse) GetBody() *ListMemoryStoresResponseBody {
	return s.Body
}

func (s *ListMemoryStoresResponse) SetHeaders(v map[string]*string) *ListMemoryStoresResponse {
	s.Headers = v
	return s
}

func (s *ListMemoryStoresResponse) SetStatusCode(v int32) *ListMemoryStoresResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemoryStoresResponse) SetBody(v *ListMemoryStoresResponseBody) *ListMemoryStoresResponse {
	s.Body = v
	return s
}

func (s *ListMemoryStoresResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

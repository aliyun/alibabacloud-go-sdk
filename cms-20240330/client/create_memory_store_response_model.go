// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMemoryStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMemoryStoreResponse
	GetStatusCode() *int32
	SetBody(v *CreateMemoryStoreResponseBody) *CreateMemoryStoreResponse
	GetBody() *CreateMemoryStoreResponseBody
}

type CreateMemoryStoreResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMemoryStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemoryStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryStoreResponse) GoString() string {
	return s.String()
}

func (s *CreateMemoryStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMemoryStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMemoryStoreResponse) GetBody() *CreateMemoryStoreResponseBody {
	return s.Body
}

func (s *CreateMemoryStoreResponse) SetHeaders(v map[string]*string) *CreateMemoryStoreResponse {
	s.Headers = v
	return s
}

func (s *CreateMemoryStoreResponse) SetStatusCode(v int32) *CreateMemoryStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMemoryStoreResponse) SetBody(v *CreateMemoryStoreResponseBody) *CreateMemoryStoreResponse {
	s.Body = v
	return s
}

func (s *CreateMemoryStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

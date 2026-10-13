// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMemoryStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMemoryStoreResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMemoryStoreResponseBody) *UpdateMemoryStoreResponse
	GetBody() *UpdateMemoryStoreResponseBody
}

type UpdateMemoryStoreResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMemoryStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemoryStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryStoreResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemoryStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMemoryStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMemoryStoreResponse) GetBody() *UpdateMemoryStoreResponseBody {
	return s.Body
}

func (s *UpdateMemoryStoreResponse) SetHeaders(v map[string]*string) *UpdateMemoryStoreResponse {
	s.Headers = v
	return s
}

func (s *UpdateMemoryStoreResponse) SetStatusCode(v int32) *UpdateMemoryStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMemoryStoreResponse) SetBody(v *UpdateMemoryStoreResponseBody) *UpdateMemoryStoreResponse {
	s.Body = v
	return s
}

func (s *UpdateMemoryStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

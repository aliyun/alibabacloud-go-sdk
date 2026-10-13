// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMemoryStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMemoryStoreResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMemoryStoreResponseBody) *DeleteMemoryStoreResponse
	GetBody() *DeleteMemoryStoreResponseBody
}

type DeleteMemoryStoreResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMemoryStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMemoryStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryStoreResponse) GoString() string {
	return s.String()
}

func (s *DeleteMemoryStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMemoryStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMemoryStoreResponse) GetBody() *DeleteMemoryStoreResponseBody {
	return s.Body
}

func (s *DeleteMemoryStoreResponse) SetHeaders(v map[string]*string) *DeleteMemoryStoreResponse {
	s.Headers = v
	return s
}

func (s *DeleteMemoryStoreResponse) SetStatusCode(v int32) *DeleteMemoryStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMemoryStoreResponse) SetBody(v *DeleteMemoryStoreResponseBody) *DeleteMemoryStoreResponse {
	s.Body = v
	return s
}

func (s *DeleteMemoryStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

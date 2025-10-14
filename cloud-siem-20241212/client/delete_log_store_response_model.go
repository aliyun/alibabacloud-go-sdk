// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLogStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLogStoreResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLogStoreResponseBody) *DeleteLogStoreResponse
	GetBody() *DeleteLogStoreResponseBody
}

type DeleteLogStoreResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLogStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLogStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogStoreResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLogStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLogStoreResponse) GetBody() *DeleteLogStoreResponseBody {
	return s.Body
}

func (s *DeleteLogStoreResponse) SetHeaders(v map[string]*string) *DeleteLogStoreResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogStoreResponse) SetStatusCode(v int32) *DeleteLogStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLogStoreResponse) SetBody(v *DeleteLogStoreResponseBody) *DeleteLogStoreResponse {
	s.Body = v
	return s
}

func (s *DeleteLogStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContextStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContextStoreResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContextStoreResponseBody) *DeleteContextStoreResponse
	GetBody() *DeleteContextStoreResponseBody
}

type DeleteContextStoreResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContextStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContextStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextStoreResponse) GoString() string {
	return s.String()
}

func (s *DeleteContextStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContextStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContextStoreResponse) GetBody() *DeleteContextStoreResponseBody {
	return s.Body
}

func (s *DeleteContextStoreResponse) SetHeaders(v map[string]*string) *DeleteContextStoreResponse {
	s.Headers = v
	return s
}

func (s *DeleteContextStoreResponse) SetStatusCode(v int32) *DeleteContextStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContextStoreResponse) SetBody(v *DeleteContextStoreResponseBody) *DeleteContextStoreResponse {
	s.Body = v
	return s
}

func (s *DeleteContextStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

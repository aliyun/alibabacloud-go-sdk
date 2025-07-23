// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEntityStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEntityStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEntityStoreResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEntityStoreResponseBody) *DeleteEntityStoreResponse
	GetBody() *DeleteEntityStoreResponseBody
}

type DeleteEntityStoreResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEntityStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEntityStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEntityStoreResponse) GoString() string {
	return s.String()
}

func (s *DeleteEntityStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEntityStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEntityStoreResponse) GetBody() *DeleteEntityStoreResponseBody {
	return s.Body
}

func (s *DeleteEntityStoreResponse) SetHeaders(v map[string]*string) *DeleteEntityStoreResponse {
	s.Headers = v
	return s
}

func (s *DeleteEntityStoreResponse) SetStatusCode(v int32) *DeleteEntityStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEntityStoreResponse) SetBody(v *DeleteEntityStoreResponseBody) *DeleteEntityStoreResponse {
	s.Body = v
	return s
}

func (s *DeleteEntityStoreResponse) Validate() error {
	return dara.Validate(s)
}

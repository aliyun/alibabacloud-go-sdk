// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKvNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKvNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKvNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKvNamespaceResponseBody) *DeleteKvNamespaceResponse
	GetBody() *DeleteKvNamespaceResponseBody
}

type DeleteKvNamespaceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKvNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKvNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKvNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteKvNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKvNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKvNamespaceResponse) GetBody() *DeleteKvNamespaceResponseBody {
	return s.Body
}

func (s *DeleteKvNamespaceResponse) SetHeaders(v map[string]*string) *DeleteKvNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteKvNamespaceResponse) SetStatusCode(v int32) *DeleteKvNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKvNamespaceResponse) SetBody(v *DeleteKvNamespaceResponseBody) *DeleteKvNamespaceResponse {
	s.Body = v
	return s
}

func (s *DeleteKvNamespaceResponse) Validate() error {
	return dara.Validate(s)
}

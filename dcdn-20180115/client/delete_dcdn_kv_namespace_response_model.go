// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnKvNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDcdnKvNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDcdnKvNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDcdnKvNamespaceResponseBody) *DeleteDcdnKvNamespaceResponse
	GetBody() *DeleteDcdnKvNamespaceResponseBody
}

type DeleteDcdnKvNamespaceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDcdnKvNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDcdnKvNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnKvNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnKvNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDcdnKvNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDcdnKvNamespaceResponse) GetBody() *DeleteDcdnKvNamespaceResponseBody {
	return s.Body
}

func (s *DeleteDcdnKvNamespaceResponse) SetHeaders(v map[string]*string) *DeleteDcdnKvNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnKvNamespaceResponse) SetStatusCode(v int32) *DeleteDcdnKvNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDcdnKvNamespaceResponse) SetBody(v *DeleteDcdnKvNamespaceResponseBody) *DeleteDcdnKvNamespaceResponse {
	s.Body = v
	return s
}

func (s *DeleteDcdnKvNamespaceResponse) Validate() error {
	return dara.Validate(s)
}

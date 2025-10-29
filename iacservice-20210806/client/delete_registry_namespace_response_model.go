// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegistryNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRegistryNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRegistryNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRegistryNamespaceResponseBody) *DeleteRegistryNamespaceResponse
	GetBody() *DeleteRegistryNamespaceResponseBody
}

type DeleteRegistryNamespaceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRegistryNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRegistryNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegistryNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteRegistryNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRegistryNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRegistryNamespaceResponse) GetBody() *DeleteRegistryNamespaceResponseBody {
	return s.Body
}

func (s *DeleteRegistryNamespaceResponse) SetHeaders(v map[string]*string) *DeleteRegistryNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteRegistryNamespaceResponse) SetStatusCode(v int32) *DeleteRegistryNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRegistryNamespaceResponse) SetBody(v *DeleteRegistryNamespaceResponseBody) *DeleteRegistryNamespaceResponse {
	s.Body = v
	return s
}

func (s *DeleteRegistryNamespaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

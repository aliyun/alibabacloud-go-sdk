// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRegistryNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRegistryNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRegistryNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateRegistryNamespaceResponseBody) *CreateRegistryNamespaceResponse
	GetBody() *CreateRegistryNamespaceResponseBody
}

type CreateRegistryNamespaceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRegistryNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRegistryNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRegistryNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateRegistryNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRegistryNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRegistryNamespaceResponse) GetBody() *CreateRegistryNamespaceResponseBody {
	return s.Body
}

func (s *CreateRegistryNamespaceResponse) SetHeaders(v map[string]*string) *CreateRegistryNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateRegistryNamespaceResponse) SetStatusCode(v int32) *CreateRegistryNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRegistryNamespaceResponse) SetBody(v *CreateRegistryNamespaceResponseBody) *CreateRegistryNamespaceResponse {
	s.Body = v
	return s
}

func (s *CreateRegistryNamespaceResponse) Validate() error {
	return dara.Validate(s)
}

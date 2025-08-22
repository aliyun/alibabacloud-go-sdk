// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegistryNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRegistryNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRegistryNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *GetRegistryNamespaceResponseBody) *GetRegistryNamespaceResponse
	GetBody() *GetRegistryNamespaceResponseBody
}

type GetRegistryNamespaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRegistryNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRegistryNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRegistryNamespaceResponse) GoString() string {
	return s.String()
}

func (s *GetRegistryNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRegistryNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRegistryNamespaceResponse) GetBody() *GetRegistryNamespaceResponseBody {
	return s.Body
}

func (s *GetRegistryNamespaceResponse) SetHeaders(v map[string]*string) *GetRegistryNamespaceResponse {
	s.Headers = v
	return s
}

func (s *GetRegistryNamespaceResponse) SetStatusCode(v int32) *GetRegistryNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegistryNamespaceResponse) SetBody(v *GetRegistryNamespaceResponseBody) *GetRegistryNamespaceResponse {
	s.Body = v
	return s
}

func (s *GetRegistryNamespaceResponse) Validate() error {
	return dara.Validate(s)
}

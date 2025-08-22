// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRegistryNamespaceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRegistryNamespaceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRegistryNamespaceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRegistryNamespaceAttributeResponseBody) *UpdateRegistryNamespaceAttributeResponse
	GetBody() *UpdateRegistryNamespaceAttributeResponseBody
}

type UpdateRegistryNamespaceAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRegistryNamespaceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRegistryNamespaceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRegistryNamespaceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRegistryNamespaceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRegistryNamespaceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRegistryNamespaceAttributeResponse) GetBody() *UpdateRegistryNamespaceAttributeResponseBody {
	return s.Body
}

func (s *UpdateRegistryNamespaceAttributeResponse) SetHeaders(v map[string]*string) *UpdateRegistryNamespaceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRegistryNamespaceAttributeResponse) SetStatusCode(v int32) *UpdateRegistryNamespaceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRegistryNamespaceAttributeResponse) SetBody(v *UpdateRegistryNamespaceAttributeResponseBody) *UpdateRegistryNamespaceAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateRegistryNamespaceAttributeResponse) Validate() error {
	return dara.Validate(s)
}

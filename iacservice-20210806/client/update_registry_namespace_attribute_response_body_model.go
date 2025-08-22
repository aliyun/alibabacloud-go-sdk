// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRegistryNamespaceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceName(v string) *UpdateRegistryNamespaceAttributeResponseBody
	GetNamespaceName() *string
	SetRequestId(v string) *UpdateRegistryNamespaceAttributeResponseBody
	GetRequestId() *string
}

type UpdateRegistryNamespaceAttributeResponseBody struct {
	// example:
	//
	// test
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	// example:
	//
	// CA5C5B39-D1DC-5309-8E97-B9A91DA21094
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateRegistryNamespaceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRegistryNamespaceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRegistryNamespaceAttributeResponseBody) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *UpdateRegistryNamespaceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRegistryNamespaceAttributeResponseBody) SetNamespaceName(v string) *UpdateRegistryNamespaceAttributeResponseBody {
	s.NamespaceName = &v
	return s
}

func (s *UpdateRegistryNamespaceAttributeResponseBody) SetRequestId(v string) *UpdateRegistryNamespaceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRegistryNamespaceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

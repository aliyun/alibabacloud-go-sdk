// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRegistryNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceName(v string) *CreateRegistryNamespaceResponseBody
	GetNamespaceName() *string
	SetRequestId(v string) *CreateRegistryNamespaceResponseBody
	GetRequestId() *string
}

type CreateRegistryNamespaceResponseBody struct {
	// example:
	//
	// iac
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	// example:
	//
	// B4672AE3-C313-5B7A-BB24-45345570D398
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateRegistryNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRegistryNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRegistryNamespaceResponseBody) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *CreateRegistryNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRegistryNamespaceResponseBody) SetNamespaceName(v string) *CreateRegistryNamespaceResponseBody {
	s.NamespaceName = &v
	return s
}

func (s *CreateRegistryNamespaceResponseBody) SetRequestId(v string) *CreateRegistryNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRegistryNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegistryNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v *GetRegistryNamespaceResponseBodyNamespace) *GetRegistryNamespaceResponseBody
	GetNamespace() *GetRegistryNamespaceResponseBodyNamespace
	SetRequestId(v string) *GetRegistryNamespaceResponseBody
	GetRequestId() *string
}

type GetRegistryNamespaceResponseBody struct {
	Namespace *GetRegistryNamespaceResponseBodyNamespace `json:"namespace,omitempty" xml:"namespace,omitempty" type:"Struct"`
	// example:
	//
	// 26684763-5BAB-58C8-BA4F-9D622AB7AD14
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRegistryNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRegistryNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegistryNamespaceResponseBody) GetNamespace() *GetRegistryNamespaceResponseBodyNamespace {
	return s.Namespace
}

func (s *GetRegistryNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRegistryNamespaceResponseBody) SetNamespace(v *GetRegistryNamespaceResponseBodyNamespace) *GetRegistryNamespaceResponseBody {
	s.Namespace = v
	return s
}

func (s *GetRegistryNamespaceResponseBody) SetRequestId(v string) *GetRegistryNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegistryNamespaceResponseBody) Validate() error {
	if s.Namespace != nil {
		if err := s.Namespace.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRegistryNamespaceResponseBodyNamespace struct {
	// example:
	//
	// private
	Acl *string `json:"acl,omitempty" xml:"acl,omitempty"`
	// example:
	//
	// 2025-03-20T02:18:29Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// demo
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// admin
	Maintainer *string `json:"maintainer,omitempty" xml:"maintainer,omitempty"`
	// example:
	//
	// 21
	Modules *int32 `json:"modules,omitempty" xml:"modules,omitempty"`
	// example:
	//
	// test_namespace
	NamespaceName  *string  `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	SharedAccounts []*int64 `json:"sharedAccounts,omitempty" xml:"sharedAccounts,omitempty" type:"Repeated"`
	// example:
	//
	// share
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetRegistryNamespaceResponseBodyNamespace) String() string {
	return dara.Prettify(s)
}

func (s GetRegistryNamespaceResponseBodyNamespace) GoString() string {
	return s.String()
}

func (s *GetRegistryNamespaceResponseBodyNamespace) GetAcl() *string {
	return s.Acl
}

func (s *GetRegistryNamespaceResponseBodyNamespace) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRegistryNamespaceResponseBodyNamespace) GetDescription() *string {
	return s.Description
}

func (s *GetRegistryNamespaceResponseBodyNamespace) GetMaintainer() *string {
	return s.Maintainer
}

func (s *GetRegistryNamespaceResponseBodyNamespace) GetModules() *int32 {
	return s.Modules
}

func (s *GetRegistryNamespaceResponseBodyNamespace) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *GetRegistryNamespaceResponseBodyNamespace) GetSharedAccounts() []*int64 {
	return s.SharedAccounts
}

func (s *GetRegistryNamespaceResponseBodyNamespace) GetType() *string {
	return s.Type
}

func (s *GetRegistryNamespaceResponseBodyNamespace) SetAcl(v string) *GetRegistryNamespaceResponseBodyNamespace {
	s.Acl = &v
	return s
}

func (s *GetRegistryNamespaceResponseBodyNamespace) SetCreateTime(v string) *GetRegistryNamespaceResponseBodyNamespace {
	s.CreateTime = &v
	return s
}

func (s *GetRegistryNamespaceResponseBodyNamespace) SetDescription(v string) *GetRegistryNamespaceResponseBodyNamespace {
	s.Description = &v
	return s
}

func (s *GetRegistryNamespaceResponseBodyNamespace) SetMaintainer(v string) *GetRegistryNamespaceResponseBodyNamespace {
	s.Maintainer = &v
	return s
}

func (s *GetRegistryNamespaceResponseBodyNamespace) SetModules(v int32) *GetRegistryNamespaceResponseBodyNamespace {
	s.Modules = &v
	return s
}

func (s *GetRegistryNamespaceResponseBodyNamespace) SetNamespaceName(v string) *GetRegistryNamespaceResponseBodyNamespace {
	s.NamespaceName = &v
	return s
}

func (s *GetRegistryNamespaceResponseBodyNamespace) SetSharedAccounts(v []*int64) *GetRegistryNamespaceResponseBodyNamespace {
	s.SharedAccounts = v
	return s
}

func (s *GetRegistryNamespaceResponseBodyNamespace) SetType(v string) *GetRegistryNamespaceResponseBodyNamespace {
	s.Type = &v
	return s
}

func (s *GetRegistryNamespaceResponseBodyNamespace) Validate() error {
	return dara.Validate(s)
}

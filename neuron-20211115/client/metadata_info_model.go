// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetadataInfo interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *MetadataInfo
	GetEnterpriseId() *int64
	SetEnv(v string) *MetadataInfo
	GetEnv() *string
	SetId(v int64) *MetadataInfo
	GetId() *int64
	SetInstanceId(v int64) *MetadataInfo
	GetInstanceId() *int64
	SetNamespaceId(v int64) *MetadataInfo
	GetNamespaceId() *int64
	SetNamespaceName(v string) *MetadataInfo
	GetNamespaceName() *string
	SetOrgId(v int32) *MetadataInfo
	GetOrgId() *int32
}

type MetadataInfo struct {
	// example:
	//
	// 1
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1
	InstanceId *int64 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 1
	NamespaceId *int64 `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// product
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	// example:
	//
	// 1
	OrgId *int32 `json:"orgId,omitempty" xml:"orgId,omitempty"`
}

func (s MetadataInfo) String() string {
	return dara.Prettify(s)
}

func (s MetadataInfo) GoString() string {
	return s.String()
}

func (s *MetadataInfo) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *MetadataInfo) GetEnv() *string {
	return s.Env
}

func (s *MetadataInfo) GetId() *int64 {
	return s.Id
}

func (s *MetadataInfo) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *MetadataInfo) GetNamespaceId() *int64 {
	return s.NamespaceId
}

func (s *MetadataInfo) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *MetadataInfo) GetOrgId() *int32 {
	return s.OrgId
}

func (s *MetadataInfo) SetEnterpriseId(v int64) *MetadataInfo {
	s.EnterpriseId = &v
	return s
}

func (s *MetadataInfo) SetEnv(v string) *MetadataInfo {
	s.Env = &v
	return s
}

func (s *MetadataInfo) SetId(v int64) *MetadataInfo {
	s.Id = &v
	return s
}

func (s *MetadataInfo) SetInstanceId(v int64) *MetadataInfo {
	s.InstanceId = &v
	return s
}

func (s *MetadataInfo) SetNamespaceId(v int64) *MetadataInfo {
	s.NamespaceId = &v
	return s
}

func (s *MetadataInfo) SetNamespaceName(v string) *MetadataInfo {
	s.NamespaceName = &v
	return s
}

func (s *MetadataInfo) SetOrgId(v int32) *MetadataInfo {
	s.OrgId = &v
	return s
}

func (s *MetadataInfo) Validate() error {
	return dara.Validate(s)
}

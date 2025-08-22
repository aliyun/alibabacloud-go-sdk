// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegistryNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListRegistryNamespacesResponseBody
	GetCount() *int64
	SetMaxResults(v int32) *ListRegistryNamespacesResponseBody
	GetMaxResults() *int32
	SetNamespaces(v []*ListRegistryNamespacesResponseBodyNamespaces) *ListRegistryNamespacesResponseBody
	GetNamespaces() []*ListRegistryNamespacesResponseBodyNamespaces
	SetNextToken(v string) *ListRegistryNamespacesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRegistryNamespacesResponseBody
	GetRequestId() *string
}

type ListRegistryNamespacesResponseBody struct {
	// example:
	//
	// 53
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// 24
	MaxResults *int32                                          `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	Namespaces []*ListRegistryNamespacesResponseBodyNamespaces `json:"namespaces,omitempty" xml:"namespaces,omitempty" type:"Repeated"`
	// example:
	//
	// IPTL1Zpr1andEF4lQ3XAYFTgtpI04QQpc5dyKpESXBc=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 5FFB0033-A016-5A9D-9283-C123AAA7F71D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRegistryNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRegistryNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegistryNamespacesResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListRegistryNamespacesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRegistryNamespacesResponseBody) GetNamespaces() []*ListRegistryNamespacesResponseBodyNamespaces {
	return s.Namespaces
}

func (s *ListRegistryNamespacesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRegistryNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRegistryNamespacesResponseBody) SetCount(v int64) *ListRegistryNamespacesResponseBody {
	s.Count = &v
	return s
}

func (s *ListRegistryNamespacesResponseBody) SetMaxResults(v int32) *ListRegistryNamespacesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRegistryNamespacesResponseBody) SetNamespaces(v []*ListRegistryNamespacesResponseBodyNamespaces) *ListRegistryNamespacesResponseBody {
	s.Namespaces = v
	return s
}

func (s *ListRegistryNamespacesResponseBody) SetNextToken(v string) *ListRegistryNamespacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRegistryNamespacesResponseBody) SetRequestId(v string) *ListRegistryNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegistryNamespacesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRegistryNamespacesResponseBodyNamespaces struct {
	// example:
	//
	// private
	Acl *string `json:"acl,omitempty" xml:"acl,omitempty"`
	// example:
	//
	// 2025-01-15T02:16:58Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// dd
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
	// shared
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListRegistryNamespacesResponseBodyNamespaces) String() string {
	return dara.Prettify(s)
}

func (s ListRegistryNamespacesResponseBodyNamespaces) GoString() string {
	return s.String()
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) GetAcl() *string {
	return s.Acl
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) GetDescription() *string {
	return s.Description
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) GetMaintainer() *string {
	return s.Maintainer
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) GetModules() *int32 {
	return s.Modules
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) GetSharedAccounts() []*int64 {
	return s.SharedAccounts
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) GetType() *string {
	return s.Type
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) SetAcl(v string) *ListRegistryNamespacesResponseBodyNamespaces {
	s.Acl = &v
	return s
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) SetCreateTime(v string) *ListRegistryNamespacesResponseBodyNamespaces {
	s.CreateTime = &v
	return s
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) SetDescription(v string) *ListRegistryNamespacesResponseBodyNamespaces {
	s.Description = &v
	return s
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) SetMaintainer(v string) *ListRegistryNamespacesResponseBodyNamespaces {
	s.Maintainer = &v
	return s
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) SetModules(v int32) *ListRegistryNamespacesResponseBodyNamespaces {
	s.Modules = &v
	return s
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) SetNamespaceName(v string) *ListRegistryNamespacesResponseBodyNamespaces {
	s.NamespaceName = &v
	return s
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) SetSharedAccounts(v []*int64) *ListRegistryNamespacesResponseBodyNamespaces {
	s.SharedAccounts = v
	return s
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) SetType(v string) *ListRegistryNamespacesResponseBodyNamespaces {
	s.Type = &v
	return s
}

func (s *ListRegistryNamespacesResponseBodyNamespaces) Validate() error {
	return dara.Validate(s)
}

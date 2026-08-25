// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegistryModulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListRegistryModulesResponseBody
	GetCount() *int64
	SetMaxResults(v int32) *ListRegistryModulesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListRegistryModulesResponseBody
	GetNextToken() *string
	SetRegistryModules(v []*ListRegistryModulesResponseBodyRegistryModules) *ListRegistryModulesResponseBody
	GetRegistryModules() []*ListRegistryModulesResponseBodyRegistryModules
	SetRequestId(v string) *ListRegistryModulesResponseBody
	GetRequestId() *string
}

type ListRegistryModulesResponseBody struct {
	// The total number of records.
	//
	// example:
	//
	// 173
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page. A value of null indicates that no more pages are available.
	//
	// example:
	//
	// iRdnbADPQp4dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The list of Registry modules.
	RegistryModules []*ListRegistryModulesResponseBodyRegistryModules `json:"registryModules,omitempty" xml:"registryModules,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D25216A9-C0F7-5A3A-A7E4-2B3D4F3A355D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRegistryModulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRegistryModulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegistryModulesResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListRegistryModulesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRegistryModulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRegistryModulesResponseBody) GetRegistryModules() []*ListRegistryModulesResponseBodyRegistryModules {
	return s.RegistryModules
}

func (s *ListRegistryModulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRegistryModulesResponseBody) SetCount(v int64) *ListRegistryModulesResponseBody {
	s.Count = &v
	return s
}

func (s *ListRegistryModulesResponseBody) SetMaxResults(v int32) *ListRegistryModulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRegistryModulesResponseBody) SetNextToken(v string) *ListRegistryModulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRegistryModulesResponseBody) SetRegistryModules(v []*ListRegistryModulesResponseBodyRegistryModules) *ListRegistryModulesResponseBody {
	s.RegistryModules = v
	return s
}

func (s *ListRegistryModulesResponseBody) SetRequestId(v string) *ListRegistryModulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegistryModulesResponseBody) Validate() error {
	if s.RegistryModules != nil {
		for _, item := range s.RegistryModules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRegistryModulesResponseBodyRegistryModules struct {
	// The permission. private: private.
	//
	// example:
	//
	// private
	Acl *string `json:"acl,omitempty" xml:"acl,omitempty"`
	// The creation time in the format of YYYY-MM-DD HH:mm:ss. The returned value does not include a time zone identifier.
	//
	// example:
	//
	// 2025-06-10 16:16:04
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the Registry module.
	//
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The number of downloads.
	//
	// example:
	//
	// 32
	Downloads *int32 `json:"downloads,omitempty" xml:"downloads,omitempty"`
	// The name of the Registry module.
	//
	// example:
	//
	// ModuleName
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// NamespaceName
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	// The provider type. alicloud: Alibaba Cloud.
	//
	// example:
	//
	// alicloud
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// The list of accounts with which the Registry module is shared.
	SharedAccounts []*int64 `json:"sharedAccounts,omitempty" xml:"sharedAccounts,omitempty" type:"Repeated"`
	// The module source.
	//
	// example:
	//
	// namespaceName/ModuleName
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The module source URL.
	//
	// example:
	//
	// URL
	SourceUrl *string `json:"sourceUrl,omitempty" xml:"sourceUrl,omitempty"`
	// The templatetype.
	//
	// example:
	//
	// system
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The latest version number.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListRegistryModulesResponseBodyRegistryModules) String() string {
	return dara.Prettify(s)
}

func (s ListRegistryModulesResponseBodyRegistryModules) GoString() string {
	return s.String()
}

func (s *ListRegistryModulesResponseBodyRegistryModules) GetAcl() *string {
	return s.Acl
}

func (s *ListRegistryModulesResponseBodyRegistryModules) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRegistryModulesResponseBodyRegistryModules) GetDescription() *string {
	return s.Description
}

func (s *ListRegistryModulesResponseBodyRegistryModules) GetDownloads() *int32 {
	return s.Downloads
}

func (s *ListRegistryModulesResponseBodyRegistryModules) GetModuleName() *string {
	return s.ModuleName
}

func (s *ListRegistryModulesResponseBodyRegistryModules) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListRegistryModulesResponseBodyRegistryModules) GetProvider() *string {
	return s.Provider
}

func (s *ListRegistryModulesResponseBodyRegistryModules) GetSharedAccounts() []*int64 {
	return s.SharedAccounts
}

func (s *ListRegistryModulesResponseBodyRegistryModules) GetSource() *string {
	return s.Source
}

func (s *ListRegistryModulesResponseBodyRegistryModules) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *ListRegistryModulesResponseBodyRegistryModules) GetType() *string {
	return s.Type
}

func (s *ListRegistryModulesResponseBodyRegistryModules) GetVersion() *string {
	return s.Version
}

func (s *ListRegistryModulesResponseBodyRegistryModules) SetAcl(v string) *ListRegistryModulesResponseBodyRegistryModules {
	s.Acl = &v
	return s
}

func (s *ListRegistryModulesResponseBodyRegistryModules) SetCreateTime(v string) *ListRegistryModulesResponseBodyRegistryModules {
	s.CreateTime = &v
	return s
}

func (s *ListRegistryModulesResponseBodyRegistryModules) SetDescription(v string) *ListRegistryModulesResponseBodyRegistryModules {
	s.Description = &v
	return s
}

func (s *ListRegistryModulesResponseBodyRegistryModules) SetDownloads(v int32) *ListRegistryModulesResponseBodyRegistryModules {
	s.Downloads = &v
	return s
}

func (s *ListRegistryModulesResponseBodyRegistryModules) SetModuleName(v string) *ListRegistryModulesResponseBodyRegistryModules {
	s.ModuleName = &v
	return s
}

func (s *ListRegistryModulesResponseBodyRegistryModules) SetNamespaceName(v string) *ListRegistryModulesResponseBodyRegistryModules {
	s.NamespaceName = &v
	return s
}

func (s *ListRegistryModulesResponseBodyRegistryModules) SetProvider(v string) *ListRegistryModulesResponseBodyRegistryModules {
	s.Provider = &v
	return s
}

func (s *ListRegistryModulesResponseBodyRegistryModules) SetSharedAccounts(v []*int64) *ListRegistryModulesResponseBodyRegistryModules {
	s.SharedAccounts = v
	return s
}

func (s *ListRegistryModulesResponseBodyRegistryModules) SetSource(v string) *ListRegistryModulesResponseBodyRegistryModules {
	s.Source = &v
	return s
}

func (s *ListRegistryModulesResponseBodyRegistryModules) SetSourceUrl(v string) *ListRegistryModulesResponseBodyRegistryModules {
	s.SourceUrl = &v
	return s
}

func (s *ListRegistryModulesResponseBodyRegistryModules) SetType(v string) *ListRegistryModulesResponseBodyRegistryModules {
	s.Type = &v
	return s
}

func (s *ListRegistryModulesResponseBodyRegistryModules) SetVersion(v string) *ListRegistryModulesResponseBodyRegistryModules {
	s.Version = &v
	return s
}

func (s *ListRegistryModulesResponseBodyRegistryModules) Validate() error {
	return dara.Validate(s)
}

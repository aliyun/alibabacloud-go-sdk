// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegistryModuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegistryModule(v *GetRegistryModuleResponseBodyRegistryModule) *GetRegistryModuleResponseBody
	GetRegistryModule() *GetRegistryModuleResponseBodyRegistryModule
	SetRequestId(v string) *GetRegistryModuleResponseBody
	GetRequestId() *string
}

type GetRegistryModuleResponseBody struct {
	// The Registry template.
	RegistryModule *GetRegistryModuleResponseBodyRegistryModule `json:"registryModule,omitempty" xml:"registryModule,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5B5AD471-5036-581B-AC9B-7D5EECED877A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRegistryModuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRegistryModuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegistryModuleResponseBody) GetRegistryModule() *GetRegistryModuleResponseBodyRegistryModule {
	return s.RegistryModule
}

func (s *GetRegistryModuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRegistryModuleResponseBody) SetRegistryModule(v *GetRegistryModuleResponseBodyRegistryModule) *GetRegistryModuleResponseBody {
	s.RegistryModule = v
	return s
}

func (s *GetRegistryModuleResponseBody) SetRequestId(v string) *GetRegistryModuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegistryModuleResponseBody) Validate() error {
	if s.RegistryModule != nil {
		if err := s.RegistryModule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRegistryModuleResponseBodyRegistryModule struct {
	// The permissions. Valid values:
	//
	// - private: Private.
	//
	// example:
	//
	// private
	Acl *string `json:"acl,omitempty" xml:"acl,omitempty"`
	// The creation time in the YYYY-MM-DD HH:mm:ss format. The returned value does not include a time zone identifier.
	//
	// example:
	//
	// 2025-05-28 13:39:05
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The template description.
	//
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The number of downloads.
	//
	// example:
	//
	// 23
	Downloads *int32 `json:"downloads,omitempty" xml:"downloads,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// ecs-cluster
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// NamespaceName
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	// The provider type. Valid values:
	//
	// - alicloud: Alibaba Cloud.
	//
	// example:
	//
	// alicloud
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// The list of accounts that share the Registry template.
	SharedAccounts []*int64 `json:"sharedAccounts,omitempty" xml:"sharedAccounts,omitempty" type:"Repeated"`
	// The template source, which is concatenated in the <NamespaceName>/<ModuleName> format.
	//
	// example:
	//
	// namespaceName/ModuleName
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The URL of the template source.
	//
	// example:
	//
	// URL
	SourceUrl *string `json:"sourceUrl,omitempty" xml:"sourceUrl,omitempty"`
	// The workspace type. Valid values:
	//
	// - system: Public module.
	//
	// - self: Custom module.
	//
	// - shared: Shared module.
	//
	// - community: Community module.
	//
	// example:
	//
	// system
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The latest version.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetRegistryModuleResponseBodyRegistryModule) String() string {
	return dara.Prettify(s)
}

func (s GetRegistryModuleResponseBodyRegistryModule) GoString() string {
	return s.String()
}

func (s *GetRegistryModuleResponseBodyRegistryModule) GetAcl() *string {
	return s.Acl
}

func (s *GetRegistryModuleResponseBodyRegistryModule) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRegistryModuleResponseBodyRegistryModule) GetDescription() *string {
	return s.Description
}

func (s *GetRegistryModuleResponseBodyRegistryModule) GetDownloads() *int32 {
	return s.Downloads
}

func (s *GetRegistryModuleResponseBodyRegistryModule) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetRegistryModuleResponseBodyRegistryModule) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *GetRegistryModuleResponseBodyRegistryModule) GetProvider() *string {
	return s.Provider
}

func (s *GetRegistryModuleResponseBodyRegistryModule) GetSharedAccounts() []*int64 {
	return s.SharedAccounts
}

func (s *GetRegistryModuleResponseBodyRegistryModule) GetSource() *string {
	return s.Source
}

func (s *GetRegistryModuleResponseBodyRegistryModule) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *GetRegistryModuleResponseBodyRegistryModule) GetType() *string {
	return s.Type
}

func (s *GetRegistryModuleResponseBodyRegistryModule) GetVersion() *string {
	return s.Version
}

func (s *GetRegistryModuleResponseBodyRegistryModule) SetAcl(v string) *GetRegistryModuleResponseBodyRegistryModule {
	s.Acl = &v
	return s
}

func (s *GetRegistryModuleResponseBodyRegistryModule) SetCreateTime(v string) *GetRegistryModuleResponseBodyRegistryModule {
	s.CreateTime = &v
	return s
}

func (s *GetRegistryModuleResponseBodyRegistryModule) SetDescription(v string) *GetRegistryModuleResponseBodyRegistryModule {
	s.Description = &v
	return s
}

func (s *GetRegistryModuleResponseBodyRegistryModule) SetDownloads(v int32) *GetRegistryModuleResponseBodyRegistryModule {
	s.Downloads = &v
	return s
}

func (s *GetRegistryModuleResponseBodyRegistryModule) SetModuleName(v string) *GetRegistryModuleResponseBodyRegistryModule {
	s.ModuleName = &v
	return s
}

func (s *GetRegistryModuleResponseBodyRegistryModule) SetNamespaceName(v string) *GetRegistryModuleResponseBodyRegistryModule {
	s.NamespaceName = &v
	return s
}

func (s *GetRegistryModuleResponseBodyRegistryModule) SetProvider(v string) *GetRegistryModuleResponseBodyRegistryModule {
	s.Provider = &v
	return s
}

func (s *GetRegistryModuleResponseBodyRegistryModule) SetSharedAccounts(v []*int64) *GetRegistryModuleResponseBodyRegistryModule {
	s.SharedAccounts = v
	return s
}

func (s *GetRegistryModuleResponseBodyRegistryModule) SetSource(v string) *GetRegistryModuleResponseBodyRegistryModule {
	s.Source = &v
	return s
}

func (s *GetRegistryModuleResponseBodyRegistryModule) SetSourceUrl(v string) *GetRegistryModuleResponseBodyRegistryModule {
	s.SourceUrl = &v
	return s
}

func (s *GetRegistryModuleResponseBodyRegistryModule) SetType(v string) *GetRegistryModuleResponseBodyRegistryModule {
	s.Type = &v
	return s
}

func (s *GetRegistryModuleResponseBodyRegistryModule) SetVersion(v string) *GetRegistryModuleResponseBodyRegistryModule {
	s.Version = &v
	return s
}

func (s *GetRegistryModuleResponseBodyRegistryModule) Validate() error {
	return dara.Validate(s)
}

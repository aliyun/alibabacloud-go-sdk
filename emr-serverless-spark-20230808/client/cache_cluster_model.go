// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCacheCluster interface {
	dara.Model
	String() string
	GoString() string
	SetBindedWorkspaces(v []*string) *CacheCluster
	GetBindedWorkspaces() []*string
	SetCacheClusterId(v string) *CacheCluster
	GetCacheClusterId() *string
	SetCacheClusterName(v string) *CacheCluster
	GetCacheClusterName() *string
	SetCachesets(v []*CacheClusterCachesets) *CacheCluster
	GetCachesets() []*CacheClusterCachesets
	SetClusterId(v string) *CacheCluster
	GetClusterId() *string
	SetConfiguration(v string) *CacheCluster
	GetConfiguration() *string
	SetConfigurations(v []*CacheClusterConfigurations) *CacheCluster
	GetConfigurations() []*CacheClusterConfigurations
	SetCreateTime(v string) *CacheCluster
	GetCreateTime() *string
	SetCreator(v string) *CacheCluster
	GetCreator() *string
	SetExtra(v string) *CacheCluster
	GetExtra() *string
	SetGmtCreated(v int64) *CacheCluster
	GetGmtCreated() *int64
	SetGmtModified(v int64) *CacheCluster
	GetGmtModified() *int64
	SetModifier(v string) *CacheCluster
	GetModifier() *string
	SetName(v string) *CacheCluster
	GetName() *string
	SetPaymentType(v string) *CacheCluster
	GetPaymentType() *string
	SetRegionId(v string) *CacheCluster
	GetRegionId() *string
	SetResourceGroupId(v string) *CacheCluster
	GetResourceGroupId() *string
	SetResourceSpec(v *CacheClusterResourceSpec) *CacheCluster
	GetResourceSpec() *CacheClusterResourceSpec
	SetState(v string) *CacheCluster
	GetState() *string
	SetTables(v []*CacheClusterTables) *CacheCluster
	GetTables() []*CacheClusterTables
	SetTags(v []*CacheClusterTags) *CacheCluster
	GetTags() []*CacheClusterTags
	SetUsedResourceSpec(v *CacheClusterUsedResourceSpec) *CacheCluster
	GetUsedResourceSpec() *CacheClusterUsedResourceSpec
	SetVersion(v string) *CacheCluster
	GetVersion() *string
}

type CacheCluster struct {
	BindedWorkspaces []*string                     `json:"bindedWorkspaces,omitempty" xml:"bindedWorkspaces,omitempty" type:"Repeated"`
	CacheClusterId   *string                       `json:"cacheClusterId,omitempty" xml:"cacheClusterId,omitempty"`
	CacheClusterName *string                       `json:"cacheClusterName,omitempty" xml:"cacheClusterName,omitempty"`
	Cachesets        []*CacheClusterCachesets      `json:"cachesets,omitempty" xml:"cachesets,omitempty" type:"Repeated"`
	ClusterId        *string                       `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	Configuration    *string                       `json:"configuration,omitempty" xml:"configuration,omitempty"`
	Configurations   []*CacheClusterConfigurations `json:"configurations,omitempty" xml:"configurations,omitempty" type:"Repeated"`
	CreateTime       *string                       `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator          *string                       `json:"creator,omitempty" xml:"creator,omitempty"`
	Extra            *string                       `json:"extra,omitempty" xml:"extra,omitempty"`
	GmtCreated       *int64                        `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	GmtModified      *int64                        `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Modifier         *string                       `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Name             *string                       `json:"name,omitempty" xml:"name,omitempty"`
	PaymentType      *string                       `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	RegionId         *string                       `json:"regionId,omitempty" xml:"regionId,omitempty"`
	ResourceGroupId  *string                       `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	ResourceSpec     *CacheClusterResourceSpec     `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty" type:"Struct"`
	State            *string                       `json:"state,omitempty" xml:"state,omitempty"`
	Tables           []*CacheClusterTables         `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
	Tags             []*CacheClusterTags           `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	UsedResourceSpec *CacheClusterUsedResourceSpec `json:"usedResourceSpec,omitempty" xml:"usedResourceSpec,omitempty" type:"Struct"`
	Version          *string                       `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CacheCluster) String() string {
	return dara.Prettify(s)
}

func (s CacheCluster) GoString() string {
	return s.String()
}

func (s *CacheCluster) GetBindedWorkspaces() []*string {
	return s.BindedWorkspaces
}

func (s *CacheCluster) GetCacheClusterId() *string {
	return s.CacheClusterId
}

func (s *CacheCluster) GetCacheClusterName() *string {
	return s.CacheClusterName
}

func (s *CacheCluster) GetCachesets() []*CacheClusterCachesets {
	return s.Cachesets
}

func (s *CacheCluster) GetClusterId() *string {
	return s.ClusterId
}

func (s *CacheCluster) GetConfiguration() *string {
	return s.Configuration
}

func (s *CacheCluster) GetConfigurations() []*CacheClusterConfigurations {
	return s.Configurations
}

func (s *CacheCluster) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CacheCluster) GetCreator() *string {
	return s.Creator
}

func (s *CacheCluster) GetExtra() *string {
	return s.Extra
}

func (s *CacheCluster) GetGmtCreated() *int64 {
	return s.GmtCreated
}

func (s *CacheCluster) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *CacheCluster) GetModifier() *string {
	return s.Modifier
}

func (s *CacheCluster) GetName() *string {
	return s.Name
}

func (s *CacheCluster) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CacheCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *CacheCluster) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CacheCluster) GetResourceSpec() *CacheClusterResourceSpec {
	return s.ResourceSpec
}

func (s *CacheCluster) GetState() *string {
	return s.State
}

func (s *CacheCluster) GetTables() []*CacheClusterTables {
	return s.Tables
}

func (s *CacheCluster) GetTags() []*CacheClusterTags {
	return s.Tags
}

func (s *CacheCluster) GetUsedResourceSpec() *CacheClusterUsedResourceSpec {
	return s.UsedResourceSpec
}

func (s *CacheCluster) GetVersion() *string {
	return s.Version
}

func (s *CacheCluster) SetBindedWorkspaces(v []*string) *CacheCluster {
	s.BindedWorkspaces = v
	return s
}

func (s *CacheCluster) SetCacheClusterId(v string) *CacheCluster {
	s.CacheClusterId = &v
	return s
}

func (s *CacheCluster) SetCacheClusterName(v string) *CacheCluster {
	s.CacheClusterName = &v
	return s
}

func (s *CacheCluster) SetCachesets(v []*CacheClusterCachesets) *CacheCluster {
	s.Cachesets = v
	return s
}

func (s *CacheCluster) SetClusterId(v string) *CacheCluster {
	s.ClusterId = &v
	return s
}

func (s *CacheCluster) SetConfiguration(v string) *CacheCluster {
	s.Configuration = &v
	return s
}

func (s *CacheCluster) SetConfigurations(v []*CacheClusterConfigurations) *CacheCluster {
	s.Configurations = v
	return s
}

func (s *CacheCluster) SetCreateTime(v string) *CacheCluster {
	s.CreateTime = &v
	return s
}

func (s *CacheCluster) SetCreator(v string) *CacheCluster {
	s.Creator = &v
	return s
}

func (s *CacheCluster) SetExtra(v string) *CacheCluster {
	s.Extra = &v
	return s
}

func (s *CacheCluster) SetGmtCreated(v int64) *CacheCluster {
	s.GmtCreated = &v
	return s
}

func (s *CacheCluster) SetGmtModified(v int64) *CacheCluster {
	s.GmtModified = &v
	return s
}

func (s *CacheCluster) SetModifier(v string) *CacheCluster {
	s.Modifier = &v
	return s
}

func (s *CacheCluster) SetName(v string) *CacheCluster {
	s.Name = &v
	return s
}

func (s *CacheCluster) SetPaymentType(v string) *CacheCluster {
	s.PaymentType = &v
	return s
}

func (s *CacheCluster) SetRegionId(v string) *CacheCluster {
	s.RegionId = &v
	return s
}

func (s *CacheCluster) SetResourceGroupId(v string) *CacheCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *CacheCluster) SetResourceSpec(v *CacheClusterResourceSpec) *CacheCluster {
	s.ResourceSpec = v
	return s
}

func (s *CacheCluster) SetState(v string) *CacheCluster {
	s.State = &v
	return s
}

func (s *CacheCluster) SetTables(v []*CacheClusterTables) *CacheCluster {
	s.Tables = v
	return s
}

func (s *CacheCluster) SetTags(v []*CacheClusterTags) *CacheCluster {
	s.Tags = v
	return s
}

func (s *CacheCluster) SetUsedResourceSpec(v *CacheClusterUsedResourceSpec) *CacheCluster {
	s.UsedResourceSpec = v
	return s
}

func (s *CacheCluster) SetVersion(v string) *CacheCluster {
	s.Version = &v
	return s
}

func (s *CacheCluster) Validate() error {
	if s.Cachesets != nil {
		for _, item := range s.Cachesets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Configurations != nil {
		for _, item := range s.Configurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	if s.Tables != nil {
		for _, item := range s.Tables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UsedResourceSpec != nil {
		if err := s.UsedResourceSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CacheClusterCachesets struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s CacheClusterCachesets) String() string {
	return dara.Prettify(s)
}

func (s CacheClusterCachesets) GoString() string {
	return s.String()
}

func (s *CacheClusterCachesets) GetName() *string {
	return s.Name
}

func (s *CacheClusterCachesets) GetPath() *string {
	return s.Path
}

func (s *CacheClusterCachesets) SetName(v string) *CacheClusterCachesets {
	s.Name = &v
	return s
}

func (s *CacheClusterCachesets) SetPath(v string) *CacheClusterCachesets {
	s.Path = &v
	return s
}

func (s *CacheClusterCachesets) Validate() error {
	return dara.Validate(s)
}

type CacheClusterConfigurations struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// cacheset.xml
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CacheClusterConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CacheClusterConfigurations) GoString() string {
	return s.String()
}

func (s *CacheClusterConfigurations) GetContent() *string {
	return s.Content
}

func (s *CacheClusterConfigurations) GetName() *string {
	return s.Name
}

func (s *CacheClusterConfigurations) SetContent(v string) *CacheClusterConfigurations {
	s.Content = &v
	return s
}

func (s *CacheClusterConfigurations) SetName(v string) *CacheClusterConfigurations {
	s.Name = &v
	return s
}

func (s *CacheClusterConfigurations) Validate() error {
	return dara.Validate(s)
}

type CacheClusterResourceSpec struct {
	BandWidth *int64 `json:"bandWidth,omitempty" xml:"bandWidth,omitempty"`
	Ha        *bool  `json:"ha,omitempty" xml:"ha,omitempty"`
	Storage   *int64 `json:"storage,omitempty" xml:"storage,omitempty"`
}

func (s CacheClusterResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s CacheClusterResourceSpec) GoString() string {
	return s.String()
}

func (s *CacheClusterResourceSpec) GetBandWidth() *int64 {
	return s.BandWidth
}

func (s *CacheClusterResourceSpec) GetHa() *bool {
	return s.Ha
}

func (s *CacheClusterResourceSpec) GetStorage() *int64 {
	return s.Storage
}

func (s *CacheClusterResourceSpec) SetBandWidth(v int64) *CacheClusterResourceSpec {
	s.BandWidth = &v
	return s
}

func (s *CacheClusterResourceSpec) SetHa(v bool) *CacheClusterResourceSpec {
	s.Ha = &v
	return s
}

func (s *CacheClusterResourceSpec) SetStorage(v int64) *CacheClusterResourceSpec {
	s.Storage = &v
	return s
}

func (s *CacheClusterResourceSpec) Validate() error {
	return dara.Validate(s)
}

type CacheClusterTables struct {
	CatalogId       *string `json:"catalogId,omitempty" xml:"catalogId,omitempty"`
	CatalogProvider *string `json:"catalogProvider,omitempty" xml:"catalogProvider,omitempty"`
	Database        *string `json:"database,omitempty" xml:"database,omitempty"`
	Table           *string `json:"table,omitempty" xml:"table,omitempty"`
	WorkspaceId     *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CacheClusterTables) String() string {
	return dara.Prettify(s)
}

func (s CacheClusterTables) GoString() string {
	return s.String()
}

func (s *CacheClusterTables) GetCatalogId() *string {
	return s.CatalogId
}

func (s *CacheClusterTables) GetCatalogProvider() *string {
	return s.CatalogProvider
}

func (s *CacheClusterTables) GetDatabase() *string {
	return s.Database
}

func (s *CacheClusterTables) GetTable() *string {
	return s.Table
}

func (s *CacheClusterTables) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CacheClusterTables) SetCatalogId(v string) *CacheClusterTables {
	s.CatalogId = &v
	return s
}

func (s *CacheClusterTables) SetCatalogProvider(v string) *CacheClusterTables {
	s.CatalogProvider = &v
	return s
}

func (s *CacheClusterTables) SetDatabase(v string) *CacheClusterTables {
	s.Database = &v
	return s
}

func (s *CacheClusterTables) SetTable(v string) *CacheClusterTables {
	s.Table = &v
	return s
}

func (s *CacheClusterTables) SetWorkspaceId(v string) *CacheClusterTables {
	s.WorkspaceId = &v
	return s
}

func (s *CacheClusterTables) Validate() error {
	return dara.Validate(s)
}

type CacheClusterTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CacheClusterTags) String() string {
	return dara.Prettify(s)
}

func (s CacheClusterTags) GoString() string {
	return s.String()
}

func (s *CacheClusterTags) GetKey() *string {
	return s.Key
}

func (s *CacheClusterTags) GetValue() *string {
	return s.Value
}

func (s *CacheClusterTags) SetKey(v string) *CacheClusterTags {
	s.Key = &v
	return s
}

func (s *CacheClusterTags) SetValue(v string) *CacheClusterTags {
	s.Value = &v
	return s
}

func (s *CacheClusterTags) Validate() error {
	return dara.Validate(s)
}

type CacheClusterUsedResourceSpec struct {
	BandWidth *int64 `json:"bandWidth,omitempty" xml:"bandWidth,omitempty"`
	Storage   *int64 `json:"storage,omitempty" xml:"storage,omitempty"`
}

func (s CacheClusterUsedResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s CacheClusterUsedResourceSpec) GoString() string {
	return s.String()
}

func (s *CacheClusterUsedResourceSpec) GetBandWidth() *int64 {
	return s.BandWidth
}

func (s *CacheClusterUsedResourceSpec) GetStorage() *int64 {
	return s.Storage
}

func (s *CacheClusterUsedResourceSpec) SetBandWidth(v int64) *CacheClusterUsedResourceSpec {
	s.BandWidth = &v
	return s
}

func (s *CacheClusterUsedResourceSpec) SetStorage(v int64) *CacheClusterUsedResourceSpec {
	s.Storage = &v
	return s
}

func (s *CacheClusterUsedResourceSpec) Validate() error {
	return dara.Validate(s)
}

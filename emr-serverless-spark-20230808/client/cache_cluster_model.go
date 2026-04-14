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
	SetCachesets(v []*CacheClusterCachesets) *CacheCluster
	GetCachesets() []*CacheClusterCachesets
	SetClusterId(v string) *CacheCluster
	GetClusterId() *string
	SetConfiguration(v string) *CacheCluster
	GetConfiguration() *string
	SetConfigurations(v []*CacheClusterConfigurations) *CacheCluster
	GetConfigurations() []*CacheClusterConfigurations
	SetCreator(v string) *CacheCluster
	GetCreator() *string
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
	SetResourceSpec(v *CacheClusterResourceSpec) *CacheCluster
	GetResourceSpec() *CacheClusterResourceSpec
	SetState(v string) *CacheCluster
	GetState() *string
	SetUsedResourceSpec(v *CacheClusterUsedResourceSpec) *CacheCluster
	GetUsedResourceSpec() *CacheClusterUsedResourceSpec
}

type CacheCluster struct {
	BindedWorkspaces []*string                     `json:"bindedWorkspaces,omitempty" xml:"bindedWorkspaces,omitempty" type:"Repeated"`
	Cachesets        []*CacheClusterCachesets      `json:"cachesets,omitempty" xml:"cachesets,omitempty" type:"Repeated"`
	ClusterId        *string                       `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	Configuration    *string                       `json:"configuration,omitempty" xml:"configuration,omitempty"`
	Configurations   []*CacheClusterConfigurations `json:"configurations,omitempty" xml:"configurations,omitempty" type:"Repeated"`
	Creator          *string                       `json:"creator,omitempty" xml:"creator,omitempty"`
	GmtCreated       *int64                        `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	GmtModified      *int64                        `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Modifier         *string                       `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Name             *string                       `json:"name,omitempty" xml:"name,omitempty"`
	PaymentType      *string                       `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	RegionId         *string                       `json:"regionId,omitempty" xml:"regionId,omitempty"`
	ResourceSpec     *CacheClusterResourceSpec     `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty" type:"Struct"`
	State            *string                       `json:"state,omitempty" xml:"state,omitempty"`
	UsedResourceSpec *CacheClusterUsedResourceSpec `json:"usedResourceSpec,omitempty" xml:"usedResourceSpec,omitempty" type:"Struct"`
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

func (s *CacheCluster) GetCreator() *string {
	return s.Creator
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

func (s *CacheCluster) GetResourceSpec() *CacheClusterResourceSpec {
	return s.ResourceSpec
}

func (s *CacheCluster) GetState() *string {
	return s.State
}

func (s *CacheCluster) GetUsedResourceSpec() *CacheClusterUsedResourceSpec {
	return s.UsedResourceSpec
}

func (s *CacheCluster) SetBindedWorkspaces(v []*string) *CacheCluster {
	s.BindedWorkspaces = v
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

func (s *CacheCluster) SetCreator(v string) *CacheCluster {
	s.Creator = &v
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

func (s *CacheCluster) SetResourceSpec(v *CacheClusterResourceSpec) *CacheCluster {
	s.ResourceSpec = v
	return s
}

func (s *CacheCluster) SetState(v string) *CacheCluster {
	s.State = &v
	return s
}

func (s *CacheCluster) SetUsedResourceSpec(v *CacheClusterUsedResourceSpec) *CacheCluster {
	s.UsedResourceSpec = v
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

func (s *CacheClusterResourceSpec) GetStorage() *int64 {
	return s.Storage
}

func (s *CacheClusterResourceSpec) SetBandWidth(v int64) *CacheClusterResourceSpec {
	s.BandWidth = &v
	return s
}

func (s *CacheClusterResourceSpec) SetStorage(v int64) *CacheClusterResourceSpec {
	s.Storage = &v
	return s
}

func (s *CacheClusterResourceSpec) Validate() error {
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

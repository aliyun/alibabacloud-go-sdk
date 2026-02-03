// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCacheCluster interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CacheCluster
	GetClusterId() *string
	SetClusterName(v string) *CacheCluster
	GetClusterName() *string
	SetCreatedAt(v int64) *CacheCluster
	GetCreatedAt() *int64
	SetCreatedBy(v string) *CacheCluster
	GetCreatedBy() *string
	SetDeployInstanceVersion(v string) *CacheCluster
	GetDeployInstanceVersion() *string
	SetDeployOptionsVersion(v int64) *CacheCluster
	GetDeployOptionsVersion() *int64
	SetInstanceVersion(v string) *CacheCluster
	GetInstanceVersion() *string
	SetOptions(v map[string]*string) *CacheCluster
	GetOptions() map[string]*string
	SetOptionsVersion(v int64) *CacheCluster
	GetOptionsVersion() *int64
	SetStatus(v string) *CacheCluster
	GetStatus() *string
	SetUpdatedAt(v int64) *CacheCluster
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *CacheCluster
	GetUpdatedBy() *string
	SetVSwitches(v []*CacheClusterVSwitches) *CacheCluster
	GetVSwitches() []*CacheClusterVSwitches
	SetVpcId(v string) *CacheCluster
	GetVpcId() *string
}

type CacheCluster struct {
	ClusterId             *string                  `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	ClusterName           *string                  `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	CreatedAt             *int64                   `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy             *string                  `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	DeployInstanceVersion *string                  `json:"deployInstanceVersion,omitempty" xml:"deployInstanceVersion,omitempty"`
	DeployOptionsVersion  *int64                   `json:"deployOptionsVersion,omitempty" xml:"deployOptionsVersion,omitempty"`
	InstanceVersion       *string                  `json:"instanceVersion,omitempty" xml:"instanceVersion,omitempty"`
	Options               map[string]*string       `json:"options,omitempty" xml:"options,omitempty"`
	OptionsVersion        *int64                   `json:"optionsVersion,omitempty" xml:"optionsVersion,omitempty"`
	Status                *string                  `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt             *int64                   `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy             *string                  `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
	VSwitches             []*CacheClusterVSwitches `json:"vSwitches,omitempty" xml:"vSwitches,omitempty" type:"Repeated"`
	VpcId                 *string                  `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s CacheCluster) String() string {
	return dara.Prettify(s)
}

func (s CacheCluster) GoString() string {
	return s.String()
}

func (s *CacheCluster) GetClusterId() *string {
	return s.ClusterId
}

func (s *CacheCluster) GetClusterName() *string {
	return s.ClusterName
}

func (s *CacheCluster) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *CacheCluster) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *CacheCluster) GetDeployInstanceVersion() *string {
	return s.DeployInstanceVersion
}

func (s *CacheCluster) GetDeployOptionsVersion() *int64 {
	return s.DeployOptionsVersion
}

func (s *CacheCluster) GetInstanceVersion() *string {
	return s.InstanceVersion
}

func (s *CacheCluster) GetOptions() map[string]*string {
	return s.Options
}

func (s *CacheCluster) GetOptionsVersion() *int64 {
	return s.OptionsVersion
}

func (s *CacheCluster) GetStatus() *string {
	return s.Status
}

func (s *CacheCluster) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *CacheCluster) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *CacheCluster) GetVSwitches() []*CacheClusterVSwitches {
	return s.VSwitches
}

func (s *CacheCluster) GetVpcId() *string {
	return s.VpcId
}

func (s *CacheCluster) SetClusterId(v string) *CacheCluster {
	s.ClusterId = &v
	return s
}

func (s *CacheCluster) SetClusterName(v string) *CacheCluster {
	s.ClusterName = &v
	return s
}

func (s *CacheCluster) SetCreatedAt(v int64) *CacheCluster {
	s.CreatedAt = &v
	return s
}

func (s *CacheCluster) SetCreatedBy(v string) *CacheCluster {
	s.CreatedBy = &v
	return s
}

func (s *CacheCluster) SetDeployInstanceVersion(v string) *CacheCluster {
	s.DeployInstanceVersion = &v
	return s
}

func (s *CacheCluster) SetDeployOptionsVersion(v int64) *CacheCluster {
	s.DeployOptionsVersion = &v
	return s
}

func (s *CacheCluster) SetInstanceVersion(v string) *CacheCluster {
	s.InstanceVersion = &v
	return s
}

func (s *CacheCluster) SetOptions(v map[string]*string) *CacheCluster {
	s.Options = v
	return s
}

func (s *CacheCluster) SetOptionsVersion(v int64) *CacheCluster {
	s.OptionsVersion = &v
	return s
}

func (s *CacheCluster) SetStatus(v string) *CacheCluster {
	s.Status = &v
	return s
}

func (s *CacheCluster) SetUpdatedAt(v int64) *CacheCluster {
	s.UpdatedAt = &v
	return s
}

func (s *CacheCluster) SetUpdatedBy(v string) *CacheCluster {
	s.UpdatedBy = &v
	return s
}

func (s *CacheCluster) SetVSwitches(v []*CacheClusterVSwitches) *CacheCluster {
	s.VSwitches = v
	return s
}

func (s *CacheCluster) SetVpcId(v string) *CacheCluster {
	s.VpcId = &v
	return s
}

func (s *CacheCluster) Validate() error {
	if s.VSwitches != nil {
		for _, item := range s.VSwitches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CacheClusterVSwitches struct {
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	Zone      *string `json:"zone,omitempty" xml:"zone,omitempty"`
}

func (s CacheClusterVSwitches) String() string {
	return dara.Prettify(s)
}

func (s CacheClusterVSwitches) GoString() string {
	return s.String()
}

func (s *CacheClusterVSwitches) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CacheClusterVSwitches) GetZone() *string {
	return s.Zone
}

func (s *CacheClusterVSwitches) SetVSwitchId(v string) *CacheClusterVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *CacheClusterVSwitches) SetZone(v string) *CacheClusterVSwitches {
	s.Zone = &v
	return s
}

func (s *CacheClusterVSwitches) Validate() error {
	return dara.Validate(s)
}

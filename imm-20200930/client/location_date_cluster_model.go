// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLocationDateCluster interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v []*Address) *LocationDateCluster
	GetAddresses() []*Address
	SetCreateTime(v string) *LocationDateCluster
	GetCreateTime() *string
	SetCustomId(v string) *LocationDateCluster
	GetCustomId() *string
	SetCustomLabels(v map[string]interface{}) *LocationDateCluster
	GetCustomLabels() map[string]interface{}
	SetLocationDateClusterEndTime(v string) *LocationDateCluster
	GetLocationDateClusterEndTime() *string
	SetLocationDateClusterLevel(v string) *LocationDateCluster
	GetLocationDateClusterLevel() *string
	SetLocationDateClusterStartTime(v string) *LocationDateCluster
	GetLocationDateClusterStartTime() *string
	SetObjectId(v string) *LocationDateCluster
	GetObjectId() *string
	SetTitle(v string) *LocationDateCluster
	GetTitle() *string
	SetUpdateTime(v string) *LocationDateCluster
	GetUpdateTime() *string
}

type LocationDateCluster struct {
	Addresses                    []*Address             `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	CreateTime                   *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomId                     *string                `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels                 map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	LocationDateClusterEndTime   *string                `json:"LocationDateClusterEndTime,omitempty" xml:"LocationDateClusterEndTime,omitempty"`
	LocationDateClusterLevel     *string                `json:"LocationDateClusterLevel,omitempty" xml:"LocationDateClusterLevel,omitempty"`
	LocationDateClusterStartTime *string                `json:"LocationDateClusterStartTime,omitempty" xml:"LocationDateClusterStartTime,omitempty"`
	ObjectId                     *string                `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	Title                        *string                `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTime                   *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s LocationDateCluster) String() string {
	return dara.Prettify(s)
}

func (s LocationDateCluster) GoString() string {
	return s.String()
}

func (s *LocationDateCluster) GetAddresses() []*Address {
	return s.Addresses
}

func (s *LocationDateCluster) GetCreateTime() *string {
	return s.CreateTime
}

func (s *LocationDateCluster) GetCustomId() *string {
	return s.CustomId
}

func (s *LocationDateCluster) GetCustomLabels() map[string]interface{} {
	return s.CustomLabels
}

func (s *LocationDateCluster) GetLocationDateClusterEndTime() *string {
	return s.LocationDateClusterEndTime
}

func (s *LocationDateCluster) GetLocationDateClusterLevel() *string {
	return s.LocationDateClusterLevel
}

func (s *LocationDateCluster) GetLocationDateClusterStartTime() *string {
	return s.LocationDateClusterStartTime
}

func (s *LocationDateCluster) GetObjectId() *string {
	return s.ObjectId
}

func (s *LocationDateCluster) GetTitle() *string {
	return s.Title
}

func (s *LocationDateCluster) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *LocationDateCluster) SetAddresses(v []*Address) *LocationDateCluster {
	s.Addresses = v
	return s
}

func (s *LocationDateCluster) SetCreateTime(v string) *LocationDateCluster {
	s.CreateTime = &v
	return s
}

func (s *LocationDateCluster) SetCustomId(v string) *LocationDateCluster {
	s.CustomId = &v
	return s
}

func (s *LocationDateCluster) SetCustomLabels(v map[string]interface{}) *LocationDateCluster {
	s.CustomLabels = v
	return s
}

func (s *LocationDateCluster) SetLocationDateClusterEndTime(v string) *LocationDateCluster {
	s.LocationDateClusterEndTime = &v
	return s
}

func (s *LocationDateCluster) SetLocationDateClusterLevel(v string) *LocationDateCluster {
	s.LocationDateClusterLevel = &v
	return s
}

func (s *LocationDateCluster) SetLocationDateClusterStartTime(v string) *LocationDateCluster {
	s.LocationDateClusterStartTime = &v
	return s
}

func (s *LocationDateCluster) SetObjectId(v string) *LocationDateCluster {
	s.ObjectId = &v
	return s
}

func (s *LocationDateCluster) SetTitle(v string) *LocationDateCluster {
	s.Title = &v
	return s
}

func (s *LocationDateCluster) SetUpdateTime(v string) *LocationDateCluster {
	s.UpdateTime = &v
	return s
}

func (s *LocationDateCluster) Validate() error {
	if s.Addresses != nil {
		for _, item := range s.Addresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

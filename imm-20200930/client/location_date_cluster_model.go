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
	// The addresses.
	Addresses []*Address `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// The time when the spatiotemporal cluster was created.
	//
	// example:
	//
	// 2022-11-16T13:14:34.882523669+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The custom ID.
	//
	// example:
	//
	// user-01
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// The custom labels.
	//
	// example:
	//
	// {
	//
	//       "User": "Jane"
	//
	// }
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The end time of the spatiotemporal cluster.
	//
	// example:
	//
	// 2022-05-02T23:59:59.999999999+08:00
	LocationDateClusterEndTime *string `json:"LocationDateClusterEndTime,omitempty" xml:"LocationDateClusterEndTime,omitempty"`
	// The administrative level of the spatiotemporal cluster.
	//
	// Enumerated values:
	//
	// 	- country
	//
	// 	- province
	//
	// 	- city
	//
	// 	- district
	//
	// 	- township
	//
	// example:
	//
	// province
	LocationDateClusterLevel *string `json:"LocationDateClusterLevel,omitempty" xml:"LocationDateClusterLevel,omitempty"`
	// The start time of the spatiotemporal cluster.
	//
	// example:
	//
	// 2022-05-01T00:00:00+08:00
	LocationDateClusterStartTime *string `json:"LocationDateClusterStartTime,omitempty" xml:"LocationDateClusterStartTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// location-date-cluster-14f48cb3-079d-4595-80c4-5735284b****
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The custom title.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The time when the spatiotemporal cluster was updated.
	//
	// example:
	//
	// 2022-11-16T13:15:05.65746784+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

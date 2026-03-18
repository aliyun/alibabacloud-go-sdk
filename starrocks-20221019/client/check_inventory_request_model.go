// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInventoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterInfo(v string) *CheckInventoryRequest
	GetClusterInfo() *string
	SetZoneId(v string) *CheckInventoryRequest
	GetZoneId() *string
}

type CheckInventoryRequest struct {
	// example:
	//
	// {
	//
	//   "clusterType": "sr",
	//
	//   "regionId": "cn-hangzhou",
	//
	//   "packageType": "official",
	//
	//   "runMode": "shared_data",
	//
	//   "beResourceSpec": {
	//
	//     "cu": 8,
	//
	//     "storageSize": 200,
	//
	//     "nodeNumber": 3,
	//
	//     "diskNumber": 1,
	//
	//     "storagePerformanceLevel": "pl1",
	//
	//     "diskType": "essd",
	//
	//     "specType": "standard"
	//
	//   },
	//
	//   "feResourceSpec": {
	//
	//     "cu": 8,
	//
	//     "storageSize": 100,
	//
	//     "nodeNumber": 3,
	//
	//     "specType": "standard"
	//
	//   }
	//
	// }
	ClusterInfo *string `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty"`
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CheckInventoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckInventoryRequest) GoString() string {
	return s.String()
}

func (s *CheckInventoryRequest) GetClusterInfo() *string {
	return s.ClusterInfo
}

func (s *CheckInventoryRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CheckInventoryRequest) SetClusterInfo(v string) *CheckInventoryRequest {
	s.ClusterInfo = &v
	return s
}

func (s *CheckInventoryRequest) SetZoneId(v string) *CheckInventoryRequest {
	s.ZoneId = &v
	return s
}

func (s *CheckInventoryRequest) Validate() error {
	return dara.Validate(s)
}

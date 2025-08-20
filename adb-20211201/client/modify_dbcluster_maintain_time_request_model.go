// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterMaintainTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterMaintainTimeRequest
	GetDBClusterId() *string
	SetMaintainTime(v string) *ModifyDBClusterMaintainTimeRequest
	GetMaintainTime() *string
}

type ModifyDBClusterMaintainTimeRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The maintenance window of the cluster. It must be in the hh:mmZ-hh:mmZ format.
	//
	// > The interval must be 1 hour and start and end at the beginning of an hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22:00Z-23:00Z
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
}

func (s ModifyDBClusterMaintainTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterMaintainTimeRequest) GetMaintainTime() *string {
	return s.MaintainTime
}

func (s *ModifyDBClusterMaintainTimeRequest) SetDBClusterId(v string) *ModifyDBClusterMaintainTimeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetMaintainTime(v string) *ModifyDBClusterMaintainTimeRequest {
	s.MaintainTime = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOpenSearchClassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceDiskSize(v int32) *ModifyOpenSearchClassRequest
	GetDBInstanceDiskSize() *int32
	SetDBInstanceName(v string) *ModifyOpenSearchClassRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ModifyOpenSearchClassRequest
	GetRegionId() *string
	SetSearchClassCode(v string) *ModifyOpenSearchClassRequest
	GetSearchClassCode() *string
}

type ModifyOpenSearchClassRequest struct {
	// The target disk size per node, in GB. If not specified, the current disk size is retained.
	//
	// example:
	//
	// 500
	DBInstanceDiskSize *int32 `json:"DBInstanceDiskSize,omitempty" xml:"DBInstanceDiskSize,omitempty"`
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-spsil01pww4hfz
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The node specifications code of the PolarDB-X Search data node. This parameter is required. Active node specifications depend on the region and sales configuration, and must differ from the current node specifications.
	//
	// example:
	//
	// opensearch.sn2ne.large.1
	SearchClassCode *string `json:"SearchClassCode,omitempty" xml:"SearchClassCode,omitempty"`
}

func (s ModifyOpenSearchClassRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenSearchClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyOpenSearchClassRequest) GetDBInstanceDiskSize() *int32 {
	return s.DBInstanceDiskSize
}

func (s *ModifyOpenSearchClassRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyOpenSearchClassRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyOpenSearchClassRequest) GetSearchClassCode() *string {
	return s.SearchClassCode
}

func (s *ModifyOpenSearchClassRequest) SetDBInstanceDiskSize(v int32) *ModifyOpenSearchClassRequest {
	s.DBInstanceDiskSize = &v
	return s
}

func (s *ModifyOpenSearchClassRequest) SetDBInstanceName(v string) *ModifyOpenSearchClassRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyOpenSearchClassRequest) SetRegionId(v string) *ModifyOpenSearchClassRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOpenSearchClassRequest) SetSearchClassCode(v string) *ModifyOpenSearchClassRequest {
	s.SearchClassCode = &v
	return s
}

func (s *ModifyOpenSearchClassRequest) Validate() error {
	return dara.Validate(s)
}

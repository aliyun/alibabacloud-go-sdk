// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolarFsInstanceDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyPolarFsInstanceDescriptionRequest
	GetDBClusterId() *string
	SetPolarFsInstanceDescription(v string) *ModifyPolarFsInstanceDescriptionRequest
	GetPolarFsInstanceDescription() *string
	SetPolarFsInstanceId(v string) *ModifyPolarFsInstanceDescriptionRequest
	GetPolarFsInstanceId() *string
	SetRegionId(v string) *ModifyPolarFsInstanceDescriptionRequest
	GetRegionId() *string
}

type ModifyPolarFsInstanceDescriptionRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query all clusters in the specified region, including cluster IDs.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The description of the PolarFS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pcs-****************
	PolarFsInstanceDescription *string `json:"PolarFsInstanceDescription,omitempty" xml:"PolarFsInstanceDescription,omitempty"`
	// The PolarFS instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pcs-****************
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query all available regions for the current account, including region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyPolarFsInstanceDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolarFsInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolarFsInstanceDescriptionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyPolarFsInstanceDescriptionRequest) GetPolarFsInstanceDescription() *string {
	return s.PolarFsInstanceDescription
}

func (s *ModifyPolarFsInstanceDescriptionRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *ModifyPolarFsInstanceDescriptionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPolarFsInstanceDescriptionRequest) SetDBClusterId(v string) *ModifyPolarFsInstanceDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyPolarFsInstanceDescriptionRequest) SetPolarFsInstanceDescription(v string) *ModifyPolarFsInstanceDescriptionRequest {
	s.PolarFsInstanceDescription = &v
	return s
}

func (s *ModifyPolarFsInstanceDescriptionRequest) SetPolarFsInstanceId(v string) *ModifyPolarFsInstanceDescriptionRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *ModifyPolarFsInstanceDescriptionRequest) SetRegionId(v string) *ModifyPolarFsInstanceDescriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPolarFsInstanceDescriptionRequest) Validate() error {
	return dara.Validate(s)
}

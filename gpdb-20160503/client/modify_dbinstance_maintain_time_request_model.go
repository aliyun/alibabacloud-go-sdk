// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceMaintainTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceMaintainTimeRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *ModifyDBInstanceMaintainTimeRequest
	GetEndTime() *string
	SetResourceGroupId(v string) *ModifyDBInstanceMaintainTimeRequest
	GetResourceGroupId() *string
	SetStartTime(v string) *ModifyDBInstanceMaintainTimeRequest
	GetStartTime() *string
}

type ModifyDBInstanceMaintainTimeRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end time of the maintenance window. The end time must be later than the start time. Specify the time in the HH:mmZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 03:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is no longer used.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start time of the maintenance window. Specify the time in the HH:mmZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 02:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyDBInstanceMaintainTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetDBInstanceId(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetEndTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetResourceGroupId(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetStartTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) Validate() error {
	return dara.Validate(s)
}

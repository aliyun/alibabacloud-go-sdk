// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGadInstanceDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *DescribeGadInstanceDetailResponseBody
	GetCreateTime() *int64
	SetDbEngineType(v string) *DescribeGadInstanceDetailResponseBody
	GetDbEngineType() *string
	SetDbListMode(v int32) *DescribeGadInstanceDetailResponseBody
	GetDbListMode() *int32
	SetDynamicCode(v string) *DescribeGadInstanceDetailResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeGadInstanceDetailResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeGadInstanceDetailResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeGadInstanceDetailResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *DescribeGadInstanceDetailResponseBody
	GetHttpStatusCode() *string
	SetInstanceId(v string) *DescribeGadInstanceDetailResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeGadInstanceDetailResponseBody
	GetInstanceName() *string
	SetInstanceType(v string) *DescribeGadInstanceDetailResponseBody
	GetInstanceType() *string
	SetMasterDbInstanceId(v string) *DescribeGadInstanceDetailResponseBody
	GetMasterDbInstanceId() *string
	SetMasterDbInstanceRegion(v string) *DescribeGadInstanceDetailResponseBody
	GetMasterDbInstanceRegion() *string
	SetMasterDblnstanceZoneld(v string) *DescribeGadInstanceDetailResponseBody
	GetMasterDblnstanceZoneld() *string
	SetMasterEngineArchType(v int32) *DescribeGadInstanceDetailResponseBody
	GetMasterEngineArchType() *int32
	SetRegionId(v string) *DescribeGadInstanceDetailResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeGadInstanceDetailResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeGadInstanceDetailResponseBody
	GetResourceGroupId() *string
	SetSlaveDbInstances(v *DescribeGadInstanceDetailResponseBodySlaveDbInstances) *DescribeGadInstanceDetailResponseBody
	GetSlaveDbInstances() *DescribeGadInstanceDetailResponseBodySlaveDbInstances
	SetStatus(v string) *DescribeGadInstanceDetailResponseBody
	GetStatus() *string
	SetSuccess(v string) *DescribeGadInstanceDetailResponseBody
	GetSuccess() *string
}

type DescribeGadInstanceDetailResponseBody struct {
	// The creation timestamp.
	//
	// example:
	//
	// 1770393523000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The database engine.
	//
	// example:
	//
	// MySQL
	DbEngineType *string `json:"DbEngineType,omitempty" xml:"DbEngineType,omitempty"`
	// The database list mode in disaster recovery scenarios.
	//
	// example:
	//
	// 1
	DbListMode *int32 `json:"DbListMode,omitempty" xml:"DbListMode,omitempty"`
	// The dynamic error code. This parameter will be deprecated soon.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message related to the request.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned when an error occurs during the call.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned when an error occurs during the call.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the global active instance group.
	//
	// example:
	//
	// rm-bp162d4tp0500****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the global active instance group.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The GAD instance type.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance ID of the primary instance.
	//
	// example:
	//
	// rm-bp162d4tp0500**
	MasterDbInstanceId *string `json:"MasterDbInstanceId,omitempty" xml:"MasterDbInstanceId,omitempty"`
	// The region of the primary instance.
	//
	// example:
	//
	// cn-hangzhou
	MasterDbInstanceRegion *string `json:"MasterDbInstanceRegion,omitempty" xml:"MasterDbInstanceRegion,omitempty"`
	// The zone of the primary instance.
	//
	// example:
	//
	// cn-hangzhou
	MasterDblnstanceZoneld *string `json:"MasterDblnstanceZoneld,omitempty" xml:"MasterDblnstanceZoneld,omitempty"`
	// The architecture type of the primary MongoDB instance. This parameter is specific to MongoDB. Valid values:
	//
	// - 1: replica set
	//
	// - 2: sharded cluster
	//
	// example:
	//
	// 1
	MasterEngineArchType *int32 `json:"MasterEngineArchType,omitempty" xml:"MasterEngineArchType,omitempty"`
	// The region ID. You can specify this parameter to indicate the region where the instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C166D79D-436B-45F0-B5A5-25E1959F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId  *string                                                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SlaveDbInstances *DescribeGadInstanceDetailResponseBodySlaveDbInstances `json:"SlaveDbInstances,omitempty" xml:"SlaveDbInstances,omitempty" type:"Struct"`
	// The status of the global active instance group.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGadInstanceDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDetailResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeGadInstanceDetailResponseBody) GetDbEngineType() *string {
	return s.DbEngineType
}

func (s *DescribeGadInstanceDetailResponseBody) GetDbListMode() *int32 {
	return s.DbListMode
}

func (s *DescribeGadInstanceDetailResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeGadInstanceDetailResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeGadInstanceDetailResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeGadInstanceDetailResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeGadInstanceDetailResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeGadInstanceDetailResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGadInstanceDetailResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeGadInstanceDetailResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeGadInstanceDetailResponseBody) GetMasterDbInstanceId() *string {
	return s.MasterDbInstanceId
}

func (s *DescribeGadInstanceDetailResponseBody) GetMasterDbInstanceRegion() *string {
	return s.MasterDbInstanceRegion
}

func (s *DescribeGadInstanceDetailResponseBody) GetMasterDblnstanceZoneld() *string {
	return s.MasterDblnstanceZoneld
}

func (s *DescribeGadInstanceDetailResponseBody) GetMasterEngineArchType() *int32 {
	return s.MasterEngineArchType
}

func (s *DescribeGadInstanceDetailResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGadInstanceDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGadInstanceDetailResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGadInstanceDetailResponseBody) GetSlaveDbInstances() *DescribeGadInstanceDetailResponseBodySlaveDbInstances {
	return s.SlaveDbInstances
}

func (s *DescribeGadInstanceDetailResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDetailResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeGadInstanceDetailResponseBody) SetCreateTime(v int64) *DescribeGadInstanceDetailResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetDbEngineType(v string) *DescribeGadInstanceDetailResponseBody {
	s.DbEngineType = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetDbListMode(v int32) *DescribeGadInstanceDetailResponseBody {
	s.DbListMode = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetDynamicCode(v string) *DescribeGadInstanceDetailResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetDynamicMessage(v string) *DescribeGadInstanceDetailResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetErrCode(v string) *DescribeGadInstanceDetailResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetErrMessage(v string) *DescribeGadInstanceDetailResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetHttpStatusCode(v string) *DescribeGadInstanceDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetInstanceId(v string) *DescribeGadInstanceDetailResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetInstanceName(v string) *DescribeGadInstanceDetailResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetInstanceType(v string) *DescribeGadInstanceDetailResponseBody {
	s.InstanceType = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetMasterDbInstanceId(v string) *DescribeGadInstanceDetailResponseBody {
	s.MasterDbInstanceId = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetMasterDbInstanceRegion(v string) *DescribeGadInstanceDetailResponseBody {
	s.MasterDbInstanceRegion = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetMasterDblnstanceZoneld(v string) *DescribeGadInstanceDetailResponseBody {
	s.MasterDblnstanceZoneld = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetMasterEngineArchType(v int32) *DescribeGadInstanceDetailResponseBody {
	s.MasterEngineArchType = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetRegionId(v string) *DescribeGadInstanceDetailResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetRequestId(v string) *DescribeGadInstanceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetResourceGroupId(v string) *DescribeGadInstanceDetailResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetSlaveDbInstances(v *DescribeGadInstanceDetailResponseBodySlaveDbInstances) *DescribeGadInstanceDetailResponseBody {
	s.SlaveDbInstances = v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetStatus(v string) *DescribeGadInstanceDetailResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) SetSuccess(v string) *DescribeGadInstanceDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBody) Validate() error {
	if s.SlaveDbInstances != nil {
		if err := s.SlaveDbInstances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGadInstanceDetailResponseBodySlaveDbInstances struct {
	SlaveDbInstances []*DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances `json:"SlaveDbInstances,omitempty" xml:"SlaveDbInstances,omitempty" type:"Repeated"`
}

func (s DescribeGadInstanceDetailResponseBodySlaveDbInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDetailResponseBodySlaveDbInstances) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDetailResponseBodySlaveDbInstances) GetSlaveDbInstances() []*DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances {
	return s.SlaveDbInstances
}

func (s *DescribeGadInstanceDetailResponseBodySlaveDbInstances) SetSlaveDbInstances(v []*DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances) *DescribeGadInstanceDetailResponseBodySlaveDbInstances {
	s.SlaveDbInstances = v
	return s
}

func (s *DescribeGadInstanceDetailResponseBodySlaveDbInstances) Validate() error {
	if s.SlaveDbInstances != nil {
		for _, item := range s.SlaveDbInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances struct {
	CreateTime       *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DbInstanceId     *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	DbInstanceRegion *string `json:"DbInstanceRegion,omitempty" xml:"DbInstanceRegion,omitempty"`
	DtsInstanceId    *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	SyncArchitecture *string `json:"SyncArchitecture,omitempty" xml:"SyncArchitecture,omitempty"`
}

func (s DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances) GetDbInstanceRegion() *string {
	return s.DbInstanceRegion
}

func (s *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances) GetSyncArchitecture() *string {
	return s.SyncArchitecture
}

func (s *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances) SetCreateTime(v int64) *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances) SetDbInstanceId(v string) *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances) SetDbInstanceRegion(v string) *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances {
	s.DbInstanceRegion = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances) SetDtsInstanceId(v string) *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances {
	s.DtsInstanceId = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances) SetSyncArchitecture(v string) *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances {
	s.SyncArchitecture = &v
	return s
}

func (s *DescribeGadInstanceDetailResponseBodySlaveDbInstancesSlaveDbInstances) Validate() error {
	return dara.Validate(s)
}

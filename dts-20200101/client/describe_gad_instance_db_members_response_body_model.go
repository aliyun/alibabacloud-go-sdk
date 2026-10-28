// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGadInstanceDbMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *DescribeGadInstanceDbMembersResponseBody
	GetCreateTime() *int64
	SetDbEngineType(v string) *DescribeGadInstanceDbMembersResponseBody
	GetDbEngineType() *string
	SetDynamicCode(v string) *DescribeGadInstanceDbMembersResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeGadInstanceDbMembersResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeGadInstanceDbMembersResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeGadInstanceDbMembersResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *DescribeGadInstanceDbMembersResponseBody
	GetHttpStatusCode() *string
	SetInstanceId(v string) *DescribeGadInstanceDbMembersResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeGadInstanceDbMembersResponseBody
	GetInstanceName() *string
	SetMasterDbInstanceId(v string) *DescribeGadInstanceDbMembersResponseBody
	GetMasterDbInstanceId() *string
	SetRegionId(v string) *DescribeGadInstanceDbMembersResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeGadInstanceDbMembersResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeGadInstanceDbMembersResponseBody
	GetResourceGroupId() *string
	SetSlaveDbInstances(v *DescribeGadInstanceDbMembersResponseBodySlaveDbInstances) *DescribeGadInstanceDbMembersResponseBody
	GetSlaveDbInstances() *DescribeGadInstanceDbMembersResponseBodySlaveDbInstances
	SetStatus(v string) *DescribeGadInstanceDbMembersResponseBody
	GetStatus() *string
	SetSuccess(v string) *DescribeGadInstanceDbMembersResponseBody
	GetSuccess() *string
}

type DescribeGadInstanceDbMembersResponseBody struct {
	// The timestamp when the instance was created.
	//
	// example:
	//
	// 1770393523000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The database engine.
	//
	// example:
	//
	// SQLServer
	DbEngineType *string `json:"DbEngineType,omitempty" xml:"DbEngineType,omitempty"`
	// The dynamic error code associated with the request.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the **%s*	- placeholder in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **Type**, it indicates that the **Type*	- request parameter is invalid.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the call fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call fails.
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
	// The instance ID.
	//
	// example:
	//
	// gad-pfx10jt8b8m***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance group.
	//
	// example:
	//
	// gad-sqlserver-**-set4
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance ID of the primary role.
	//
	// example:
	//
	// rm-j6c685****1kb05
	MasterDbInstanceId *string `json:"MasterDbInstanceId,omitempty" xml:"MasterDbInstanceId,omitempty"`
	// The region where the DTS instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FC3BAAF2-74E3-4471-8EB5-96202D6A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId  *string                                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SlaveDbInstances *DescribeGadInstanceDbMembersResponseBodySlaveDbInstances `json:"SlaveDbInstances,omitempty" xml:"SlaveDbInstances,omitempty" type:"Struct"`
	// The status of the active geo-redundancy instance group.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGadInstanceDbMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDbMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetDbEngineType() *string {
	return s.DbEngineType
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetMasterDbInstanceId() *string {
	return s.MasterDbInstanceId
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetSlaveDbInstances() *DescribeGadInstanceDbMembersResponseBodySlaveDbInstances {
	return s.SlaveDbInstances
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDbMembersResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetCreateTime(v int64) *DescribeGadInstanceDbMembersResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetDbEngineType(v string) *DescribeGadInstanceDbMembersResponseBody {
	s.DbEngineType = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetDynamicCode(v string) *DescribeGadInstanceDbMembersResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetDynamicMessage(v string) *DescribeGadInstanceDbMembersResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetErrCode(v string) *DescribeGadInstanceDbMembersResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetErrMessage(v string) *DescribeGadInstanceDbMembersResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetHttpStatusCode(v string) *DescribeGadInstanceDbMembersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetInstanceId(v string) *DescribeGadInstanceDbMembersResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetInstanceName(v string) *DescribeGadInstanceDbMembersResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetMasterDbInstanceId(v string) *DescribeGadInstanceDbMembersResponseBody {
	s.MasterDbInstanceId = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetRegionId(v string) *DescribeGadInstanceDbMembersResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetRequestId(v string) *DescribeGadInstanceDbMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetResourceGroupId(v string) *DescribeGadInstanceDbMembersResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetSlaveDbInstances(v *DescribeGadInstanceDbMembersResponseBodySlaveDbInstances) *DescribeGadInstanceDbMembersResponseBody {
	s.SlaveDbInstances = v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetStatus(v string) *DescribeGadInstanceDbMembersResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) SetSuccess(v string) *DescribeGadInstanceDbMembersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBody) Validate() error {
	if s.SlaveDbInstances != nil {
		if err := s.SlaveDbInstances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGadInstanceDbMembersResponseBodySlaveDbInstances struct {
	SlaveDbInstances []*DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances `json:"SlaveDbInstances,omitempty" xml:"SlaveDbInstances,omitempty" type:"Repeated"`
}

func (s DescribeGadInstanceDbMembersResponseBodySlaveDbInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDbMembersResponseBodySlaveDbInstances) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDbMembersResponseBodySlaveDbInstances) GetSlaveDbInstances() []*DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances {
	return s.SlaveDbInstances
}

func (s *DescribeGadInstanceDbMembersResponseBodySlaveDbInstances) SetSlaveDbInstances(v []*DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances) *DescribeGadInstanceDbMembersResponseBodySlaveDbInstances {
	s.SlaveDbInstances = v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBodySlaveDbInstances) Validate() error {
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

type DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances struct {
	CreateTime       *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DbInstanceId     *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	DtsInstanceId    *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	SyncArchitecture *string `json:"SyncArchitecture,omitempty" xml:"SyncArchitecture,omitempty"`
}

func (s DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances) GetSyncArchitecture() *string {
	return s.SyncArchitecture
}

func (s *DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances) SetCreateTime(v int64) *DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances) SetDbInstanceId(v string) *DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances) SetDtsInstanceId(v string) *DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances {
	s.DtsInstanceId = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances) SetSyncArchitecture(v string) *DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances {
	s.SyncArchitecture = &v
	return s
}

func (s *DescribeGadInstanceDbMembersResponseBodySlaveDbInstancesSlaveDbInstances) Validate() error {
	return dara.Validate(s)
}

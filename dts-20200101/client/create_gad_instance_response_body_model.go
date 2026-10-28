// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGadInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *CreateGadInstanceResponseBody
	GetCreateTime() *int64
	SetDbEngineType(v string) *CreateGadInstanceResponseBody
	GetDbEngineType() *string
	SetDynamicCode(v string) *CreateGadInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateGadInstanceResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *CreateGadInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateGadInstanceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *CreateGadInstanceResponseBody
	GetHttpStatusCode() *string
	SetInstanceId(v string) *CreateGadInstanceResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *CreateGadInstanceResponseBody
	GetInstanceName() *string
	SetInstanceType(v string) *CreateGadInstanceResponseBody
	GetInstanceType() *string
	SetMasterDbInstanceId(v string) *CreateGadInstanceResponseBody
	GetMasterDbInstanceId() *string
	SetRegionId(v string) *CreateGadInstanceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *CreateGadInstanceResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateGadInstanceResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *CreateGadInstanceResponseBody
	GetStatus() *string
	SetSuccess(v string) *CreateGadInstanceResponseBody
	GetSuccess() *string
}

type CreateGadInstanceResponseBody struct {
	// The timestamp when the instance was created.
	//
	// example:
	//
	// 1733821662000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The database engine.
	//
	// example:
	//
	// SQLServer
	DbEngineType *string `json:"DbEngineType,omitempty" xml:"DbEngineType,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the %s placeholder in the error message of the ErrMessage response parameter.
	//
	// > If ErrMessage returns The Value of Input Parameter %s is not valid and DynamicMessage returns JobId, the JobId request parameter is invalid.
	//
	// example:
	//
	// Type
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned when the call fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned when the call fails.
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
	// rm-bpxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the group. The name must be globally unique.
	//
	// example:
	//
	// gad-sqlserver-hk-set4
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type. Valid values: DR (disaster recovery) and GAD (active geo-redundancy).
	//
	// example:
	//
	// DR
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the primary ApsaraDB RDS instance.
	//
	// example:
	//
	// rm-j6c497vzndgd8m***
	MasterDbInstanceId *string `json:"MasterDbInstanceId,omitempty" xml:"MasterDbInstanceId,omitempty"`
	// The region of the primary instance.
	//
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C306C198-7807-409D-930A-D6CE6C32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek3d****4mxya
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the group. The status is Running immediately after the group is created.
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
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGadInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGadInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGadInstanceResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateGadInstanceResponseBody) GetDbEngineType() *string {
	return s.DbEngineType
}

func (s *CreateGadInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateGadInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateGadInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateGadInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateGadInstanceResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateGadInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateGadInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateGadInstanceResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateGadInstanceResponseBody) GetMasterDbInstanceId() *string {
	return s.MasterDbInstanceId
}

func (s *CreateGadInstanceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGadInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGadInstanceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateGadInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateGadInstanceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateGadInstanceResponseBody) SetCreateTime(v int64) *CreateGadInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateGadInstanceResponseBody) SetDbEngineType(v string) *CreateGadInstanceResponseBody {
	s.DbEngineType = &v
	return s
}

func (s *CreateGadInstanceResponseBody) SetDynamicCode(v string) *CreateGadInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateGadInstanceResponseBody) SetDynamicMessage(v string) *CreateGadInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateGadInstanceResponseBody) SetErrCode(v string) *CreateGadInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateGadInstanceResponseBody) SetErrMessage(v string) *CreateGadInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateGadInstanceResponseBody) SetHttpStatusCode(v string) *CreateGadInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateGadInstanceResponseBody) SetInstanceId(v string) *CreateGadInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateGadInstanceResponseBody) SetInstanceName(v string) *CreateGadInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *CreateGadInstanceResponseBody) SetInstanceType(v string) *CreateGadInstanceResponseBody {
	s.InstanceType = &v
	return s
}

func (s *CreateGadInstanceResponseBody) SetMasterDbInstanceId(v string) *CreateGadInstanceResponseBody {
	s.MasterDbInstanceId = &v
	return s
}

func (s *CreateGadInstanceResponseBody) SetRegionId(v string) *CreateGadInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *CreateGadInstanceResponseBody) SetRequestId(v string) *CreateGadInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGadInstanceResponseBody) SetResourceGroupId(v string) *CreateGadInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateGadInstanceResponseBody) SetStatus(v string) *CreateGadInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateGadInstanceResponseBody) SetSuccess(v string) *CreateGadInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGadInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

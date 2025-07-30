// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGadInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *DescribeGadInstancesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeGadInstancesResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeGadInstancesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeGadInstancesResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *DescribeGadInstancesResponseBody
	GetHttpStatusCode() *string
	SetInstances(v *DescribeGadInstancesResponseBodyInstances) *DescribeGadInstancesResponseBody
	GetInstances() *DescribeGadInstancesResponseBodyInstances
	SetPageNumber(v int32) *DescribeGadInstancesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeGadInstancesResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeGadInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeGadInstancesResponseBody
	GetSuccess() *string
	SetTotalRecordCount(v int32) *DescribeGadInstancesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeGadInstancesResponseBody struct {
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// present environment is not support,so skip.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Instances      *DescribeGadInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// C166D79D-436B-45F0-B5A5-25E1959F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 22
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeGadInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGadInstancesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeGadInstancesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeGadInstancesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeGadInstancesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeGadInstancesResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeGadInstancesResponseBody) GetInstances() *DescribeGadInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeGadInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGadInstancesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeGadInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGadInstancesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeGadInstancesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeGadInstancesResponseBody) SetDynamicCode(v string) *DescribeGadInstancesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeGadInstancesResponseBody) SetDynamicMessage(v string) *DescribeGadInstancesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeGadInstancesResponseBody) SetErrCode(v string) *DescribeGadInstancesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeGadInstancesResponseBody) SetErrMessage(v string) *DescribeGadInstancesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeGadInstancesResponseBody) SetHttpStatusCode(v string) *DescribeGadInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeGadInstancesResponseBody) SetInstances(v *DescribeGadInstancesResponseBodyInstances) *DescribeGadInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeGadInstancesResponseBody) SetPageNumber(v int32) *DescribeGadInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGadInstancesResponseBody) SetPageRecordCount(v int32) *DescribeGadInstancesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeGadInstancesResponseBody) SetRequestId(v string) *DescribeGadInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGadInstancesResponseBody) SetSuccess(v string) *DescribeGadInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGadInstancesResponseBody) SetTotalRecordCount(v int32) *DescribeGadInstancesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeGadInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstancesResponseBodyInstances struct {
	Instances []*DescribeGadInstancesResponseBodyInstancesInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
}

func (s DescribeGadInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeGadInstancesResponseBodyInstances) GetInstances() []*DescribeGadInstancesResponseBodyInstancesInstances {
	return s.Instances
}

func (s *DescribeGadInstancesResponseBodyInstances) SetInstances(v []*DescribeGadInstancesResponseBodyInstancesInstances) *DescribeGadInstancesResponseBodyInstances {
	s.Instances = v
	return s
}

func (s *DescribeGadInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstancesResponseBodyInstancesInstances struct {
	// example:
	//
	// 2024-05-29 23:55:58
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// MySQL
	DbEngineType *string `json:"DbEngineType,omitempty" xml:"DbEngineType,omitempty"`
	// example:
	//
	// 2
	DbInstanceCount *int32 `json:"DbInstanceCount,omitempty" xml:"DbInstanceCount,omitempty"`
	// example:
	//
	// rg-a76s8afa****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	InstanceRegion *string `json:"InstanceRegion,omitempty" xml:"InstanceRegion,omitempty"`
	// example:
	//
	// DR
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// rm-sadfasfa****
	MasterDbInstanceId *string `json:"MasterDbInstanceId,omitempty" xml:"MasterDbInstanceId,omitempty"`
	// example:
	//
	// test
	MasterDbInstanceName *string `json:"MasterDbInstanceName,omitempty" xml:"MasterDbInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	MasterDbInstanceRegion *string `json:"MasterDbInstanceRegion,omitempty" xml:"MasterDbInstanceRegion,omitempty"`
	// example:
	//
	// cn-hangzhou
	MasterDbInstanceZoneId *string `json:"MasterDbInstanceZoneId,omitempty" xml:"MasterDbInstanceZoneId,omitempty"`
	MasterEngineArchType   *int32  `json:"MasterEngineArchType,omitempty" xml:"MasterEngineArchType,omitempty"`
	// example:
	//
	// rg-aekzq276dmnaxqa
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstancesResponseBodyInstancesInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstancesResponseBodyInstancesInstances) GoString() string {
	return s.String()
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) GetDbEngineType() *string {
	return s.DbEngineType
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) GetDbInstanceCount() *int32 {
	return s.DbInstanceCount
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) GetInstanceRegion() *string {
	return s.InstanceRegion
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) GetMasterDbInstanceId() *string {
	return s.MasterDbInstanceId
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) GetMasterDbInstanceName() *string {
	return s.MasterDbInstanceName
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) GetMasterDbInstanceRegion() *string {
	return s.MasterDbInstanceRegion
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) GetMasterDbInstanceZoneId() *string {
	return s.MasterDbInstanceZoneId
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) GetMasterEngineArchType() *int32 {
	return s.MasterEngineArchType
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) SetCreateTime(v int64) *DescribeGadInstancesResponseBodyInstancesInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) SetDbEngineType(v string) *DescribeGadInstancesResponseBodyInstancesInstances {
	s.DbEngineType = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) SetDbInstanceCount(v int32) *DescribeGadInstancesResponseBodyInstancesInstances {
	s.DbInstanceCount = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) SetInstanceId(v string) *DescribeGadInstancesResponseBodyInstancesInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) SetInstanceName(v string) *DescribeGadInstancesResponseBodyInstancesInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) SetInstanceRegion(v string) *DescribeGadInstancesResponseBodyInstancesInstances {
	s.InstanceRegion = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) SetInstanceType(v string) *DescribeGadInstancesResponseBodyInstancesInstances {
	s.InstanceType = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) SetMasterDbInstanceId(v string) *DescribeGadInstancesResponseBodyInstancesInstances {
	s.MasterDbInstanceId = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) SetMasterDbInstanceName(v string) *DescribeGadInstancesResponseBodyInstancesInstances {
	s.MasterDbInstanceName = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) SetMasterDbInstanceRegion(v string) *DescribeGadInstancesResponseBodyInstancesInstances {
	s.MasterDbInstanceRegion = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) SetMasterDbInstanceZoneId(v string) *DescribeGadInstancesResponseBodyInstancesInstances {
	s.MasterDbInstanceZoneId = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) SetMasterEngineArchType(v int32) *DescribeGadInstancesResponseBodyInstancesInstances {
	s.MasterEngineArchType = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) SetResourceGroupId(v string) *DescribeGadInstancesResponseBodyInstancesInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) SetStatus(v string) *DescribeGadInstancesResponseBodyInstancesInstances {
	s.Status = &v
	return s
}

func (s *DescribeGadInstancesResponseBodyInstancesInstances) Validate() error {
	return dara.Validate(s)
}

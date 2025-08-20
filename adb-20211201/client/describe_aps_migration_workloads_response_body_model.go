// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsMigrationWorkloadsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMigrationWorkloads(v []*DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) *DescribeApsMigrationWorkloadsResponseBody
	GetMigrationWorkloads() []*DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads
	SetPageNumber(v int32) *DescribeApsMigrationWorkloadsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApsMigrationWorkloadsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApsMigrationWorkloadsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApsMigrationWorkloadsResponseBody
	GetTotalCount() *int32
}

type DescribeApsMigrationWorkloadsResponseBody struct {
	// The queried migration workloads.
	MigrationWorkloads []*DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads `json:"MigrationWorkloads,omitempty" xml:"MigrationWorkloads,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******-3EEC-57F0-9F06-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApsMigrationWorkloadsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsMigrationWorkloadsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApsMigrationWorkloadsResponseBody) GetMigrationWorkloads() []*DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads {
	return s.MigrationWorkloads
}

func (s *DescribeApsMigrationWorkloadsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApsMigrationWorkloadsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApsMigrationWorkloadsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApsMigrationWorkloadsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApsMigrationWorkloadsResponseBody) SetMigrationWorkloads(v []*DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) *DescribeApsMigrationWorkloadsResponseBody {
	s.MigrationWorkloads = v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBody) SetPageNumber(v int32) *DescribeApsMigrationWorkloadsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBody) SetPageSize(v int32) *DescribeApsMigrationWorkloadsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBody) SetRequestId(v string) *DescribeApsMigrationWorkloadsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBody) SetTotalCount(v int32) *DescribeApsMigrationWorkloadsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads struct {
	// The number of AnalyticDB compute units (ACUs).
	//
	// example:
	//
	// -
	AcuCount *int32 `json:"AcuCount,omitempty" xml:"AcuCount,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2021-06-21T02:15:16Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error message.
	//
	// example:
	//
	// -
	FailedMsg *string `json:"FailedMsg,omitempty" xml:"FailedMsg,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The maximum response time.
	//
	// example:
	//
	// 1000
	MaxRT *string `json:"MaxRT,omitempty" xml:"MaxRT,omitempty"`
	// The time when the migration job was modified.
	//
	// example:
	//
	// 2021-06-21T02:15:16Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the workload.
	//
	// example:
	//
	// TEST-001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The OSS URL.
	//
	// example:
	//
	// oss://******
	OssLocation *string `json:"OssLocation,omitempty" xml:"OssLocation,omitempty"`
	// The status.
	//
	// example:
	//
	// COMPLETED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The destination type.
	//
	// example:
	//
	// OSS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The sub-type of the workload.
	//
	// example:
	//
	// test
	WorkloadSubType *string `json:"WorkloadSubType,omitempty" xml:"WorkloadSubType,omitempty"`
}

func (s DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) GoString() string {
	return s.String()
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) GetAcuCount() *int32 {
	return s.AcuCount
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) GetFailedMsg() *string {
	return s.FailedMsg
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) GetId() *string {
	return s.Id
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) GetMaxRT() *string {
	return s.MaxRT
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) GetName() *string {
	return s.Name
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) GetOssLocation() *string {
	return s.OssLocation
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) GetState() *string {
	return s.State
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) GetWorkloadSubType() *string {
	return s.WorkloadSubType
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) SetAcuCount(v int32) *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads {
	s.AcuCount = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) SetCreateTime(v string) *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads {
	s.CreateTime = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) SetFailedMsg(v string) *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads {
	s.FailedMsg = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) SetId(v string) *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads {
	s.Id = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) SetMaxRT(v string) *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads {
	s.MaxRT = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) SetModifyTime(v string) *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads {
	s.ModifyTime = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) SetName(v string) *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads {
	s.Name = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) SetOssLocation(v string) *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads {
	s.OssLocation = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) SetState(v string) *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads {
	s.State = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) SetTargetType(v string) *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads {
	s.TargetType = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) SetWorkloadSubType(v string) *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads {
	s.WorkloadSubType = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponseBodyMigrationWorkloads) Validate() error {
	return dara.Validate(s)
}

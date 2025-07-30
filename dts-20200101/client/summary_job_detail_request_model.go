// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSummaryJobDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *SummaryJobDetailRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *SummaryJobDetailRequest
	GetDtsJobId() *string
	SetJobCode(v string) *SummaryJobDetailRequest
	GetJobCode() *string
	SetRegionId(v string) *SummaryJobDetailRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *SummaryJobDetailRequest
	GetResourceGroupId() *string
	SetStructType(v string) *SummaryJobDetailRequest
	GetStructType() *string
	SetSynchronizationDirection(v string) *SummaryJobDetailRequest
	GetSynchronizationDirection() *string
	SetZeroEtlJob(v bool) *SummaryJobDetailRequest
	GetZeroEtlJob() *bool
}

type SummaryJobDetailRequest struct {
	// The ID of the data migration or data synchronization instance.
	//
	// >  You must specify at least one of the DtsJobId and DtsInstanceId parameters.
	//
	// example:
	//
	// dtsl3m1213ye7l****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the data migration or data synchronization task.
	//
	// >  You must specify at least one of the DtsJobId and DtsInstanceId parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// l3m1213ye7l****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The phase of the data migration task. Valid values:
	//
	// 	- **02**: The task is in the schema migration phase.
	//
	// 	- **03**: The task is in the incremental migration phase.
	//
	// This parameter is required.
	//
	// example:
	//
	// 02
	JobCode *string `json:"JobCode,omitempty" xml:"JobCode,omitempty"`
	// The region ID of the DTS instance. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aek25bwhtt22cjq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of schema definition. Valid values:
	//
	// 	- **before**: schema migration or initial schema synchronization
	//
	// 	- **after**: DDL operations performed during incremental data migration or synchronization
	//
	// example:
	//
	// before
	StructType *string `json:"StructType,omitempty" xml:"StructType,omitempty"`
	// The synchronization direction of the data synchronization task. Valid values:
	//
	// 	- **Forward**: Data is synchronized from the source database to the destination database.
	//
	// 	- **Reverse**: Data is synchronized from the destination database to the source database.
	//
	// >
	//
	// 	- Default value: **Forward**.
	//
	// 	- You can set this parameter to **Reverse*	- to delete the reverse synchronization task only if the topology is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// Whether it is a seamless integration (Zero-ETL) task, the value can be:
	//
	// - **false**: No. - **true**: Yes.
	//
	// example:
	//
	// true
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s SummaryJobDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s SummaryJobDetailRequest) GoString() string {
	return s.String()
}

func (s *SummaryJobDetailRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *SummaryJobDetailRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *SummaryJobDetailRequest) GetJobCode() *string {
	return s.JobCode
}

func (s *SummaryJobDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SummaryJobDetailRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SummaryJobDetailRequest) GetStructType() *string {
	return s.StructType
}

func (s *SummaryJobDetailRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *SummaryJobDetailRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *SummaryJobDetailRequest) SetDtsInstanceId(v string) *SummaryJobDetailRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *SummaryJobDetailRequest) SetDtsJobId(v string) *SummaryJobDetailRequest {
	s.DtsJobId = &v
	return s
}

func (s *SummaryJobDetailRequest) SetJobCode(v string) *SummaryJobDetailRequest {
	s.JobCode = &v
	return s
}

func (s *SummaryJobDetailRequest) SetRegionId(v string) *SummaryJobDetailRequest {
	s.RegionId = &v
	return s
}

func (s *SummaryJobDetailRequest) SetResourceGroupId(v string) *SummaryJobDetailRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SummaryJobDetailRequest) SetStructType(v string) *SummaryJobDetailRequest {
	s.StructType = &v
	return s
}

func (s *SummaryJobDetailRequest) SetSynchronizationDirection(v string) *SummaryJobDetailRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *SummaryJobDetailRequest) SetZeroEtlJob(v bool) *SummaryJobDetailRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *SummaryJobDetailRequest) Validate() error {
	return dara.Validate(s)
}

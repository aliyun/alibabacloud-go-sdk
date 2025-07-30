// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendDtsJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobIds(v string) *SuspendDtsJobsRequest
	GetDtsJobIds() *string
	SetRegionId(v string) *SuspendDtsJobsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *SuspendDtsJobsRequest
	GetResourceGroupId() *string
	SetZeroEtlJob(v bool) *SuspendDtsJobsRequest
	GetZeroEtlJob() *bool
}

type SuspendDtsJobsRequest struct {
	// The ID of the data migration or data synchronization task.
	//
	// >
	//
	// 	- For multiple tasks, separate them with commas (,).
	//
	// 	- You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hfi12iv4z7e****
	DtsJobIds *string `json:"DtsJobIds,omitempty" xml:"DtsJobIds,omitempty"`
	// The ID of the region in which the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource GroupId
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Whether it is a seamless integration (Zero-ETL) task, the value can be:
	//
	// - **false**: No. - **true**: Yes.
	//
	// example:
	//
	// true
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s SuspendDtsJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendDtsJobsRequest) GoString() string {
	return s.String()
}

func (s *SuspendDtsJobsRequest) GetDtsJobIds() *string {
	return s.DtsJobIds
}

func (s *SuspendDtsJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SuspendDtsJobsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SuspendDtsJobsRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *SuspendDtsJobsRequest) SetDtsJobIds(v string) *SuspendDtsJobsRequest {
	s.DtsJobIds = &v
	return s
}

func (s *SuspendDtsJobsRequest) SetRegionId(v string) *SuspendDtsJobsRequest {
	s.RegionId = &v
	return s
}

func (s *SuspendDtsJobsRequest) SetResourceGroupId(v string) *SuspendDtsJobsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SuspendDtsJobsRequest) SetZeroEtlJob(v bool) *SuspendDtsJobsRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *SuspendDtsJobsRequest) Validate() error {
	return dara.Validate(s)
}

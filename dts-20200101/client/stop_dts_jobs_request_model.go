// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDtsJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobIds(v string) *StopDtsJobsRequest
	GetDtsJobIds() *string
	SetRegionId(v string) *StopDtsJobsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *StopDtsJobsRequest
	GetResourceGroupId() *string
	SetZeroEtlJob(v bool) *StopDtsJobsRequest
	GetZeroEtlJob() *bool
}

type StopDtsJobsRequest struct {
	// The IDs of the migration and synchronization tasks.
	//
	// > - Separate multiple task IDs with commas (**,**).
	//
	// - Call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to obtain DTS task IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// l5o11f9029c****
	DtsJobIds *string `json:"DtsJobIds,omitempty" xml:"DtsJobIds,omitempty"`
	// The region where the DTS instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzsf6yoxhfpva
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether the node is a seamless integration (Zero-ETL) node. Valid values:
	//
	// - **false**: No.
	//
	// - **true**: Yes.
	//
	// example:
	//
	// true
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s StopDtsJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDtsJobsRequest) GoString() string {
	return s.String()
}

func (s *StopDtsJobsRequest) GetDtsJobIds() *string {
	return s.DtsJobIds
}

func (s *StopDtsJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopDtsJobsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *StopDtsJobsRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *StopDtsJobsRequest) SetDtsJobIds(v string) *StopDtsJobsRequest {
	s.DtsJobIds = &v
	return s
}

func (s *StopDtsJobsRequest) SetRegionId(v string) *StopDtsJobsRequest {
	s.RegionId = &v
	return s
}

func (s *StopDtsJobsRequest) SetResourceGroupId(v string) *StopDtsJobsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *StopDtsJobsRequest) SetZeroEtlJob(v bool) *StopDtsJobsRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *StopDtsJobsRequest) Validate() error {
	return dara.Validate(s)
}

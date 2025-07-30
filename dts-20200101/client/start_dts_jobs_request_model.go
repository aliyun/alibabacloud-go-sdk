// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDtsJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobIds(v string) *StartDtsJobsRequest
	GetDtsJobIds() *string
	SetRegionId(v string) *StartDtsJobsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *StartDtsJobsRequest
	GetResourceGroupId() *string
	SetZeroEtlJob(v bool) *StartDtsJobsRequest
	GetZeroEtlJob() *bool
}

type StartDtsJobsRequest struct {
	// The ID of the data migration or data synchronization task.
	//
	// >
	//
	// 	- For multiple tasks, separate them with commas (**,**).
	//
	// 	- You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// l5o11f9029c****
	DtsJobIds *string `json:"DtsJobIds,omitempty" xml:"DtsJobIds,omitempty"`
	// The region ID of the DTS instance. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
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

func (s StartDtsJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDtsJobsRequest) GoString() string {
	return s.String()
}

func (s *StartDtsJobsRequest) GetDtsJobIds() *string {
	return s.DtsJobIds
}

func (s *StartDtsJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartDtsJobsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *StartDtsJobsRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *StartDtsJobsRequest) SetDtsJobIds(v string) *StartDtsJobsRequest {
	s.DtsJobIds = &v
	return s
}

func (s *StartDtsJobsRequest) SetRegionId(v string) *StartDtsJobsRequest {
	s.RegionId = &v
	return s
}

func (s *StartDtsJobsRequest) SetResourceGroupId(v string) *StartDtsJobsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *StartDtsJobsRequest) SetZeroEtlJob(v bool) *StartDtsJobsRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *StartDtsJobsRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDtsJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobIds(v string) *DeleteDtsJobsRequest
	GetDtsJobIds() *string
	SetRegionId(v string) *DeleteDtsJobsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteDtsJobsRequest
	GetResourceGroupId() *string
	SetZeroEtlJob(v bool) *DeleteDtsJobsRequest
	GetZeroEtlJob() *bool
}

type DeleteDtsJobsRequest struct {
	// The IDs of the data migration, data synchronization, or change tracking tasks.
	//
	// > - You can specify up to 10 task IDs, separated by commas (,).
	//
	// - You can call [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) to query the DTS task IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// l5o11f9029c****
	DtsJobIds *string `json:"DtsJobIds,omitempty" xml:"DtsJobIds,omitempty"`
	// The region in which the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// A special parameter for specific business scenarios. You do not need to configure this parameter.
	//
	// example:
	//
	// rg3m1213ye7l****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// A special parameter for specific business scenarios. You do not need to configure this parameter.
	//
	// example:
	//
	// true
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s DeleteDtsJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDtsJobsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDtsJobsRequest) GetDtsJobIds() *string {
	return s.DtsJobIds
}

func (s *DeleteDtsJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDtsJobsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteDtsJobsRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *DeleteDtsJobsRequest) SetDtsJobIds(v string) *DeleteDtsJobsRequest {
	s.DtsJobIds = &v
	return s
}

func (s *DeleteDtsJobsRequest) SetRegionId(v string) *DeleteDtsJobsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDtsJobsRequest) SetResourceGroupId(v string) *DeleteDtsJobsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteDtsJobsRequest) SetZeroEtlJob(v bool) *DeleteDtsJobsRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *DeleteDtsJobsRequest) Validate() error {
	return dara.Validate(s)
}

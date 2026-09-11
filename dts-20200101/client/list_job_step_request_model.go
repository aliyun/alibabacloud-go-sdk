// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobStepRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *ListJobStepRequest
	GetDtsJobId() *string
	SetRegionId(v string) *ListJobStepRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListJobStepRequest
	GetResourceGroupId() *string
	SetZeroEtlJob(v bool) *ListJobStepRequest
	GetZeroEtlJob() *bool
}

type ListJobStepRequest struct {
	// Migration or sync task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// k2gm967v16f****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The region where the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aek3dcgyq7p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the job is a seamless integration (Zero-ETL) job. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s ListJobStepRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobStepRequest) GoString() string {
	return s.String()
}

func (s *ListJobStepRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ListJobStepRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListJobStepRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListJobStepRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *ListJobStepRequest) SetDtsJobId(v string) *ListJobStepRequest {
	s.DtsJobId = &v
	return s
}

func (s *ListJobStepRequest) SetRegionId(v string) *ListJobStepRequest {
	s.RegionId = &v
	return s
}

func (s *ListJobStepRequest) SetResourceGroupId(v string) *ListJobStepRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListJobStepRequest) SetZeroEtlJob(v bool) *ListJobStepRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *ListJobStepRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipFullJobTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *SkipFullJobTableRequest
	GetDtsJobId() *string
	SetJobProgressId(v string) *SkipFullJobTableRequest
	GetJobProgressId() *string
	SetRegionId(v string) *SkipFullJobTableRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *SkipFullJobTableRequest
	GetResourceGroupId() *string
	SetZeroEtlJob(v bool) *SkipFullJobTableRequest
	GetZeroEtlJob() *bool
}

type SkipFullJobTableRequest struct {
	// The ID of the data migration, data synchronization, or change tracking task. **[Deprecated]*	- This parameter is deprecated and no longer takes effect.
	//
	// example:
	//
	// l3m1213ye7l****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The primary key ID of JobProgress. Each object in a full data migration task corresponds to a JobProgress record. Specify this parameter to determine the tables to skip.
	//
	// example:
	//
	// 123
	JobProgressId *string `json:"JobProgressId,omitempty" xml:"JobProgressId,omitempty"`
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
	// rg-aek2ilvoxlrdcby
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether the node is a seamless integration (Zero-ETL) node. Valid values:
	//
	// - **true**: yes
	//
	// - **false**: no
	//
	// example:
	//
	// true
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s SkipFullJobTableRequest) String() string {
	return dara.Prettify(s)
}

func (s SkipFullJobTableRequest) GoString() string {
	return s.String()
}

func (s *SkipFullJobTableRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *SkipFullJobTableRequest) GetJobProgressId() *string {
	return s.JobProgressId
}

func (s *SkipFullJobTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SkipFullJobTableRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SkipFullJobTableRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *SkipFullJobTableRequest) SetDtsJobId(v string) *SkipFullJobTableRequest {
	s.DtsJobId = &v
	return s
}

func (s *SkipFullJobTableRequest) SetJobProgressId(v string) *SkipFullJobTableRequest {
	s.JobProgressId = &v
	return s
}

func (s *SkipFullJobTableRequest) SetRegionId(v string) *SkipFullJobTableRequest {
	s.RegionId = &v
	return s
}

func (s *SkipFullJobTableRequest) SetResourceGroupId(v string) *SkipFullJobTableRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SkipFullJobTableRequest) SetZeroEtlJob(v bool) *SkipFullJobTableRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *SkipFullJobTableRequest) Validate() error {
	return dara.Validate(s)
}

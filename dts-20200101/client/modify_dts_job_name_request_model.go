// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *ModifyDtsJobNameRequest
	GetDtsJobId() *string
	SetDtsJobName(v string) *ModifyDtsJobNameRequest
	GetDtsJobName() *string
	SetRegionId(v string) *ModifyDtsJobNameRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyDtsJobNameRequest
	GetResourceGroupId() *string
	SetZeroEtlJob(v bool) *ModifyDtsJobNameRequest
	GetZeroEtlJob() *bool
}

type ModifyDtsJobNameRequest struct {
	// The ID of the DTS task. The DTS task can be a data migration, data synchronization, or change tracking task.
	//
	// This parameter is required.
	//
	// example:
	//
	// l3m1213ye7l****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The new name of the DTS task.
	//
	// >  We recommend that you specify a descriptive name for easy identification. You do not need to use a unique name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtstest
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// The ID of the region in which the DTS instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource GroupId
	//
	// example:
	//
	// rg-aekzfkjjb5gyy6i
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

func (s ModifyDtsJobNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobNameRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ModifyDtsJobNameRequest) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *ModifyDtsJobNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDtsJobNameRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDtsJobNameRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *ModifyDtsJobNameRequest) SetDtsJobId(v string) *ModifyDtsJobNameRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifyDtsJobNameRequest) SetDtsJobName(v string) *ModifyDtsJobNameRequest {
	s.DtsJobName = &v
	return s
}

func (s *ModifyDtsJobNameRequest) SetRegionId(v string) *ModifyDtsJobNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDtsJobNameRequest) SetResourceGroupId(v string) *ModifyDtsJobNameRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDtsJobNameRequest) SetZeroEtlJob(v bool) *ModifyDtsJobNameRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *ModifyDtsJobNameRequest) Validate() error {
	return dara.Validate(s)
}

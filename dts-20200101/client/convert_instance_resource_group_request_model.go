// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertInstanceResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *ConvertInstanceResourceGroupRequest
	GetDtsJobId() *string
	SetNewResourceGroupId(v string) *ConvertInstanceResourceGroupRequest
	GetNewResourceGroupId() *string
	SetRegionId(v string) *ConvertInstanceResourceGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ConvertInstanceResourceGroupRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *ConvertInstanceResourceGroupRequest
	GetResourceId() *string
	SetZeroEtlJob(v bool) *ConvertInstanceResourceGroupRequest
	GetZeroEtlJob() *bool
}

type ConvertInstanceResourceGroupRequest struct {
	// This historical parameter does not take effect and is not required.
	//
	// example:
	//
	// m4312mab158****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The ID of new resource group. You can obtain the ID on the Resource Group page in the Resource Management console. For more information, see [View basic information about a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-aek2r4fkrqw****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The ID of the region in which the Data Transmission Service (DTS) instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is only for special services and not required.
	//
	// example:
	//
	// rg-3m1213ye7l****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the DTS instance. You can view the ID in the **ID/Name*	- column on the task page in the console.
	//
	// >  This parameter is required.
	//
	// example:
	//
	// dtszhc12zp727o****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is only for special services and not required.
	//
	// example:
	//
	// true
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s ConvertInstanceResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertInstanceResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ConvertInstanceResourceGroupRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ConvertInstanceResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *ConvertInstanceResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConvertInstanceResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConvertInstanceResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ConvertInstanceResourceGroupRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *ConvertInstanceResourceGroupRequest) SetDtsJobId(v string) *ConvertInstanceResourceGroupRequest {
	s.DtsJobId = &v
	return s
}

func (s *ConvertInstanceResourceGroupRequest) SetNewResourceGroupId(v string) *ConvertInstanceResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ConvertInstanceResourceGroupRequest) SetRegionId(v string) *ConvertInstanceResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ConvertInstanceResourceGroupRequest) SetResourceGroupId(v string) *ConvertInstanceResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConvertInstanceResourceGroupRequest) SetResourceId(v string) *ConvertInstanceResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ConvertInstanceResourceGroupRequest) SetZeroEtlJob(v bool) *ConvertInstanceResourceGroupRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *ConvertInstanceResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}

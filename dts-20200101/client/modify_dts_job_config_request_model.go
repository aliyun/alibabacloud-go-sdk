// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *ModifyDtsJobConfigRequest
	GetDtsJobId() *string
	SetOwnerId(v string) *ModifyDtsJobConfigRequest
	GetOwnerId() *string
	SetParameters(v string) *ModifyDtsJobConfigRequest
	GetParameters() *string
	SetRegionId(v string) *ModifyDtsJobConfigRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyDtsJobConfigRequest
	GetResourceGroupId() *string
}

type ModifyDtsJobConfigRequest struct {
	// The ID of the DTS task. You can call [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) to query the task ID.
	//
	// example:
	//
	// lxsn87r328d****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	OwnerId  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The parameters to be modified, in JSON character string format. For more information, see [Parameters metric description](https://help.aliyun.com/document_detail/2536412.html).
	//
	// example:
	//
	// [{"module":"07","name":"sink.connection.idle.second","value":60},{"module":"07","name":"sink.batch.size.maximum","value":64}]
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region in which the instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyDtsJobConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobConfigRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ModifyDtsJobConfigRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyDtsJobConfigRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyDtsJobConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDtsJobConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDtsJobConfigRequest) SetDtsJobId(v string) *ModifyDtsJobConfigRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifyDtsJobConfigRequest) SetOwnerId(v string) *ModifyDtsJobConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDtsJobConfigRequest) SetParameters(v string) *ModifyDtsJobConfigRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyDtsJobConfigRequest) SetRegionId(v string) *ModifyDtsJobConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDtsJobConfigRequest) SetResourceGroupId(v string) *ModifyDtsJobConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDtsJobConfigRequest) Validate() error {
	return dara.Validate(s)
}

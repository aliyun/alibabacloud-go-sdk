// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBaaSFabricBizChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaaSFabricChannelId(v string) *AddBaaSFabricBizChainRequest
	GetBaaSFabricChannelId() *string
	SetBaaSFabricConsortiumId(v string) *AddBaaSFabricBizChainRequest
	GetBaaSFabricConsortiumId() *string
	SetBaaSFabricOrganizationId(v string) *AddBaaSFabricBizChainRequest
	GetBaaSFabricOrganizationId() *string
	SetContractTemplateIdList(v string) *AddBaaSFabricBizChainRequest
	GetContractTemplateIdList() *string
	SetName(v string) *AddBaaSFabricBizChainRequest
	GetName() *string
	SetRegionId(v string) *AddBaaSFabricBizChainRequest
	GetRegionId() *string
	SetRemark(v string) *AddBaaSFabricBizChainRequest
	GetRemark() *string
}

type AddBaaSFabricBizChainRequest struct {
	// This parameter is required.
	BaaSFabricChannelId *string `json:"BaaSFabricChannelId,omitempty" xml:"BaaSFabricChannelId,omitempty"`
	// This parameter is required.
	BaaSFabricConsortiumId *string `json:"BaaSFabricConsortiumId,omitempty" xml:"BaaSFabricConsortiumId,omitempty"`
	// This parameter is required.
	BaaSFabricOrganizationId *string `json:"BaaSFabricOrganizationId,omitempty" xml:"BaaSFabricOrganizationId,omitempty"`
	ContractTemplateIdList   *string `json:"ContractTemplateIdList,omitempty" xml:"ContractTemplateIdList,omitempty"`
	// This parameter is required.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s AddBaaSFabricBizChainRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBaaSFabricBizChainRequest) GoString() string {
	return s.String()
}

func (s *AddBaaSFabricBizChainRequest) GetBaaSFabricChannelId() *string {
	return s.BaaSFabricChannelId
}

func (s *AddBaaSFabricBizChainRequest) GetBaaSFabricConsortiumId() *string {
	return s.BaaSFabricConsortiumId
}

func (s *AddBaaSFabricBizChainRequest) GetBaaSFabricOrganizationId() *string {
	return s.BaaSFabricOrganizationId
}

func (s *AddBaaSFabricBizChainRequest) GetContractTemplateIdList() *string {
	return s.ContractTemplateIdList
}

func (s *AddBaaSFabricBizChainRequest) GetName() *string {
	return s.Name
}

func (s *AddBaaSFabricBizChainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddBaaSFabricBizChainRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddBaaSFabricBizChainRequest) SetBaaSFabricChannelId(v string) *AddBaaSFabricBizChainRequest {
	s.BaaSFabricChannelId = &v
	return s
}

func (s *AddBaaSFabricBizChainRequest) SetBaaSFabricConsortiumId(v string) *AddBaaSFabricBizChainRequest {
	s.BaaSFabricConsortiumId = &v
	return s
}

func (s *AddBaaSFabricBizChainRequest) SetBaaSFabricOrganizationId(v string) *AddBaaSFabricBizChainRequest {
	s.BaaSFabricOrganizationId = &v
	return s
}

func (s *AddBaaSFabricBizChainRequest) SetContractTemplateIdList(v string) *AddBaaSFabricBizChainRequest {
	s.ContractTemplateIdList = &v
	return s
}

func (s *AddBaaSFabricBizChainRequest) SetName(v string) *AddBaaSFabricBizChainRequest {
	s.Name = &v
	return s
}

func (s *AddBaaSFabricBizChainRequest) SetRegionId(v string) *AddBaaSFabricBizChainRequest {
	s.RegionId = &v
	return s
}

func (s *AddBaaSFabricBizChainRequest) SetRemark(v string) *AddBaaSFabricBizChainRequest {
	s.Remark = &v
	return s
}

func (s *AddBaaSFabricBizChainRequest) Validate() error {
	return dara.Validate(s)
}

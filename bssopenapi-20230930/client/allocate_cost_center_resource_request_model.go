// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateCostCenterResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromCostCenterId(v int64) *AllocateCostCenterResourceRequest
	GetFromCostCenterId() *int64
	SetFromOwnerAccountId(v int64) *AllocateCostCenterResourceRequest
	GetFromOwnerAccountId() *int64
	SetNbid(v string) *AllocateCostCenterResourceRequest
	GetNbid() *string
	SetResourceInstanceList(v []*AllocateCostCenterResourceRequestResourceInstanceList) *AllocateCostCenterResourceRequest
	GetResourceInstanceList() []*AllocateCostCenterResourceRequestResourceInstanceList
	SetToCostCenterId(v int64) *AllocateCostCenterResourceRequest
	GetToCostCenterId() *int64
}

type AllocateCostCenterResourceRequest struct {
	// example:
	//
	// 637180
	FromCostCenterId *int64 `json:"FromCostCenterId,omitempty" xml:"FromCostCenterId,omitempty"`
	// example:
	//
	// 1529600453335198
	FromOwnerAccountId *int64 `json:"FromOwnerAccountId,omitempty" xml:"FromOwnerAccountId,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// This parameter is required.
	ResourceInstanceList []*AllocateCostCenterResourceRequestResourceInstanceList `json:"ResourceInstanceList,omitempty" xml:"ResourceInstanceList,omitempty" type:"Repeated"`
	// example:
	//
	// 638288
	ToCostCenterId *int64 `json:"ToCostCenterId,omitempty" xml:"ToCostCenterId,omitempty"`
}

func (s AllocateCostCenterResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateCostCenterResourceRequest) GoString() string {
	return s.String()
}

func (s *AllocateCostCenterResourceRequest) GetFromCostCenterId() *int64 {
	return s.FromCostCenterId
}

func (s *AllocateCostCenterResourceRequest) GetFromOwnerAccountId() *int64 {
	return s.FromOwnerAccountId
}

func (s *AllocateCostCenterResourceRequest) GetNbid() *string {
	return s.Nbid
}

func (s *AllocateCostCenterResourceRequest) GetResourceInstanceList() []*AllocateCostCenterResourceRequestResourceInstanceList {
	return s.ResourceInstanceList
}

func (s *AllocateCostCenterResourceRequest) GetToCostCenterId() *int64 {
	return s.ToCostCenterId
}

func (s *AllocateCostCenterResourceRequest) SetFromCostCenterId(v int64) *AllocateCostCenterResourceRequest {
	s.FromCostCenterId = &v
	return s
}

func (s *AllocateCostCenterResourceRequest) SetFromOwnerAccountId(v int64) *AllocateCostCenterResourceRequest {
	s.FromOwnerAccountId = &v
	return s
}

func (s *AllocateCostCenterResourceRequest) SetNbid(v string) *AllocateCostCenterResourceRequest {
	s.Nbid = &v
	return s
}

func (s *AllocateCostCenterResourceRequest) SetResourceInstanceList(v []*AllocateCostCenterResourceRequestResourceInstanceList) *AllocateCostCenterResourceRequest {
	s.ResourceInstanceList = v
	return s
}

func (s *AllocateCostCenterResourceRequest) SetToCostCenterId(v int64) *AllocateCostCenterResourceRequest {
	s.ToCostCenterId = &v
	return s
}

func (s *AllocateCostCenterResourceRequest) Validate() error {
	if s.ResourceInstanceList != nil {
		for _, item := range s.ResourceInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AllocateCostCenterResourceRequestResourceInstanceList struct {
	// example:
	//
	// qwer1-cn-beijing
	ApportionCode *string `json:"ApportionCode,omitempty" xml:"ApportionCode,omitempty"`
	// example:
	//
	// split-item-test1
	ApportionName *string `json:"ApportionName,omitempty" xml:"ApportionName,omitempty"`
	// example:
	//
	// oss
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// RESOURCE_UDR
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ecs
	PipCode *string `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	// example:
	//
	// related-resource
	RelatedResources *string `json:"RelatedResources,omitempty" xml:"RelatedResources,omitempty"`
	// example:
	//
	// xihe_mpp
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// example:
	//
	// cn-hangzhou;standard
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// ecs-test-1
	ResourceNick *string `json:"ResourceNick,omitempty" xml:"ResourceNick,omitempty"`
	// example:
	//
	// AUTO_ALLOCATE
	ResourceSource *string `json:"ResourceSource,omitempty" xml:"ResourceSource,omitempty"`
	// example:
	//
	// 0
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// example:
	//
	// tag-test1
	ResourceTag *string `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty"`
	// example:
	//
	// SCU
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 273394581313325532
	ResourceUserId *int64 `json:"ResourceUserId,omitempty" xml:"ResourceUserId,omitempty"`
	// example:
	//
	// test
	ResourceUserName *string `json:"ResourceUserName,omitempty" xml:"ResourceUserName,omitempty"`
}

func (s AllocateCostCenterResourceRequestResourceInstanceList) String() string {
	return dara.Prettify(s)
}

func (s AllocateCostCenterResourceRequestResourceInstanceList) GoString() string {
	return s.String()
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetApportionCode() *string {
	return s.ApportionCode
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetApportionName() *string {
	return s.ApportionName
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetCommodityName() *string {
	return s.CommodityName
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetPipCode() *string {
	return s.PipCode
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetRelatedResources() *string {
	return s.RelatedResources
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetResourceId() *string {
	return s.ResourceId
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetResourceNick() *string {
	return s.ResourceNick
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetResourceSource() *string {
	return s.ResourceSource
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetResourceTag() *string {
	return s.ResourceTag
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetResourceType() *string {
	return s.ResourceType
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetResourceUserId() *int64 {
	return s.ResourceUserId
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) GetResourceUserName() *string {
	return s.ResourceUserName
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetApportionCode(v string) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.ApportionCode = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetApportionName(v string) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.ApportionName = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetCommodityCode(v string) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.CommodityCode = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetCommodityName(v string) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.CommodityName = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetInstanceId(v string) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.InstanceId = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetPipCode(v string) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.PipCode = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetRelatedResources(v string) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.RelatedResources = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetResourceGroup(v string) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.ResourceGroup = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetResourceId(v string) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.ResourceId = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetResourceNick(v string) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.ResourceNick = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetResourceSource(v string) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.ResourceSource = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetResourceStatus(v string) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.ResourceStatus = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetResourceTag(v string) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.ResourceTag = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetResourceType(v string) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.ResourceType = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetResourceUserId(v int64) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.ResourceUserId = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) SetResourceUserName(v string) *AllocateCostCenterResourceRequestResourceInstanceList {
	s.ResourceUserName = &v
	return s
}

func (s *AllocateCostCenterResourceRequestResourceInstanceList) Validate() error {
	return dara.Validate(s)
}

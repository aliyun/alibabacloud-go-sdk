// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWorkspaceResponseBody
	GetRequestId() *string
	SetResult(v *CreateWorkspaceResponseBodyResult) *CreateWorkspaceResponseBody
	GetResult() *CreateWorkspaceResponseBodyResult
}

type CreateWorkspaceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2BA0504F-B179-586D-8210-A7C7C09A9907
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Response result
	Result *CreateWorkspaceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkspaceResponseBody) GetResult() *CreateWorkspaceResponseBodyResult {
	return s.Result
}

func (s *CreateWorkspaceResponseBody) SetRequestId(v string) *CreateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetResult(v *CreateWorkspaceResponseBodyResult) *CreateWorkspaceResponseBody {
	s.Result = v
	return s
}

func (s *CreateWorkspaceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkspaceResponseBodyResult struct {
	// Billing type
	//
	// - POSTPAY: Pay-as-you-go
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// Commodity code
	//
	// example:
	//
	// opensearch_platform_public_cn
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// Workspace description
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Custom domain name prefix
	//
	// example:
	//
	// defalult-xxxx
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// Engine type
	//
	// example:
	//
	// rag
	EngineType *string `json:"engineType,omitempty" xml:"engineType,omitempty"`
	// Workspace ID
	//
	// example:
	//
	// 120142804
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Workspace instance ID
	//
	// example:
	//
	// ops-xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Workspace name
	//
	// example:
	//
	// default
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Quota
	Quota *CreateWorkspaceResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// Resource group ID
	//
	// example:
	//
	// rg-xxxxxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Tags
	Tags []*CreateWorkspaceResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// Type
	//
	// example:
	//
	// standard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateWorkspaceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBodyResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateWorkspaceResponseBodyResult) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *CreateWorkspaceResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkspaceResponseBodyResult) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateWorkspaceResponseBodyResult) GetEngineType() *string {
	return s.EngineType
}

func (s *CreateWorkspaceResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *CreateWorkspaceResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateWorkspaceResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateWorkspaceResponseBodyResult) GetQuota() *CreateWorkspaceResponseBodyResultQuota {
	return s.Quota
}

func (s *CreateWorkspaceResponseBodyResult) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateWorkspaceResponseBodyResult) GetTags() []*CreateWorkspaceResponseBodyResultTags {
	return s.Tags
}

func (s *CreateWorkspaceResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *CreateWorkspaceResponseBodyResult) SetChargeType(v string) *CreateWorkspaceResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *CreateWorkspaceResponseBodyResult) SetCommodityCode(v string) *CreateWorkspaceResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *CreateWorkspaceResponseBodyResult) SetDescription(v string) *CreateWorkspaceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceResponseBodyResult) SetDomainName(v string) *CreateWorkspaceResponseBodyResult {
	s.DomainName = &v
	return s
}

func (s *CreateWorkspaceResponseBodyResult) SetEngineType(v string) *CreateWorkspaceResponseBodyResult {
	s.EngineType = &v
	return s
}

func (s *CreateWorkspaceResponseBodyResult) SetId(v string) *CreateWorkspaceResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateWorkspaceResponseBodyResult) SetInstanceId(v string) *CreateWorkspaceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *CreateWorkspaceResponseBodyResult) SetName(v string) *CreateWorkspaceResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceResponseBodyResult) SetQuota(v *CreateWorkspaceResponseBodyResultQuota) *CreateWorkspaceResponseBodyResult {
	s.Quota = v
	return s
}

func (s *CreateWorkspaceResponseBodyResult) SetResourceGroupId(v string) *CreateWorkspaceResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateWorkspaceResponseBodyResult) SetTags(v []*CreateWorkspaceResponseBodyResultTags) *CreateWorkspaceResponseBodyResult {
	s.Tags = v
	return s
}

func (s *CreateWorkspaceResponseBodyResult) SetType(v string) *CreateWorkspaceResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateWorkspaceResponseBodyResult) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateWorkspaceResponseBodyResultQuota struct {
	// Compute resource
	//
	// example:
	//
	// 0
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// Storage capacity
	//
	// example:
	//
	// 0
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// Specification
	//
	// example:
	//
	// rag.share.common
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s CreateWorkspaceResponseBodyResultQuota) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBodyResultQuota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *CreateWorkspaceResponseBodyResultQuota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *CreateWorkspaceResponseBodyResultQuota) GetSpec() *string {
	return s.Spec
}

func (s *CreateWorkspaceResponseBodyResultQuota) SetComputeResource(v int32) *CreateWorkspaceResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *CreateWorkspaceResponseBodyResultQuota) SetDocSize(v int32) *CreateWorkspaceResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *CreateWorkspaceResponseBodyResultQuota) SetSpec(v string) *CreateWorkspaceResponseBodyResultQuota {
	s.Spec = &v
	return s
}

func (s *CreateWorkspaceResponseBodyResultQuota) Validate() error {
	return dara.Validate(s)
}

type CreateWorkspaceResponseBodyResultTags struct {
	// Tag key
	//
	// example:
	//
	// a
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// Tag value
	//
	// example:
	//
	// c
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s CreateWorkspaceResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBodyResultTags) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateWorkspaceResponseBodyResultTags) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateWorkspaceResponseBodyResultTags) SetTagKey(v string) *CreateWorkspaceResponseBodyResultTags {
	s.TagKey = &v
	return s
}

func (s *CreateWorkspaceResponseBodyResultTags) SetTagValue(v string) *CreateWorkspaceResponseBodyResultTags {
	s.TagValue = &v
	return s
}

func (s *CreateWorkspaceResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}

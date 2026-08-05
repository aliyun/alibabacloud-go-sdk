// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWorkspaceResponseBody
	GetRequestId() *string
	SetResult(v *UpdateWorkspaceResponseBodyResult) *UpdateWorkspaceResponseBody
	GetResult() *UpdateWorkspaceResponseBodyResult
}

type UpdateWorkspaceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B7C901ED-2BC1-5CFB-BE23-242DE5E3BA5C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result *UpdateWorkspaceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkspaceResponseBody) GetResult() *UpdateWorkspaceResponseBodyResult {
	return s.Result
}

func (s *UpdateWorkspaceResponseBody) SetRequestId(v string) *UpdateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkspaceResponseBody) SetResult(v *UpdateWorkspaceResponseBodyResult) *UpdateWorkspaceResponseBody {
	s.Result = v
	return s
}

func (s *UpdateWorkspaceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkspaceResponseBodyResult struct {
	// The billing type. Valid values:
	//
	// - POSTPAY: pay-as-you-go.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The description.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The custom domain name prefix.
	//
	// example:
	//
	// default-xxx
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// The engine type.
	//
	// example:
	//
	// rag
	EngineType *string `json:"engineType,omitempty" xml:"engineType,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 1222212
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ops-cn-em93wcq0s001
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// default
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether the production is complete. Valid values:
	//
	// - 0: In production.
	//
	// - 1: Production complete.
	//
	// example:
	//
	// 1
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The quota.
	Quota *UpdateWorkspaceResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aeky6pyhbh6j3dy
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The tags.
	Tags []*UpdateWorkspaceResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The type.
	//
	// example:
	//
	// standard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateWorkspaceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBodyResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *UpdateWorkspaceResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkspaceResponseBodyResult) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateWorkspaceResponseBodyResult) GetEngineType() *string {
	return s.EngineType
}

func (s *UpdateWorkspaceResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *UpdateWorkspaceResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateWorkspaceResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateWorkspaceResponseBodyResult) GetProduced() *int32 {
	return s.Produced
}

func (s *UpdateWorkspaceResponseBodyResult) GetQuota() *UpdateWorkspaceResponseBodyResultQuota {
	return s.Quota
}

func (s *UpdateWorkspaceResponseBodyResult) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateWorkspaceResponseBodyResult) GetTags() []*UpdateWorkspaceResponseBodyResultTags {
	return s.Tags
}

func (s *UpdateWorkspaceResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *UpdateWorkspaceResponseBodyResult) SetChargeType(v string) *UpdateWorkspaceResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyResult) SetDescription(v string) *UpdateWorkspaceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyResult) SetDomainName(v string) *UpdateWorkspaceResponseBodyResult {
	s.DomainName = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyResult) SetEngineType(v string) *UpdateWorkspaceResponseBodyResult {
	s.EngineType = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyResult) SetId(v string) *UpdateWorkspaceResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyResult) SetInstanceId(v string) *UpdateWorkspaceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyResult) SetName(v string) *UpdateWorkspaceResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyResult) SetProduced(v int32) *UpdateWorkspaceResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyResult) SetQuota(v *UpdateWorkspaceResponseBodyResultQuota) *UpdateWorkspaceResponseBodyResult {
	s.Quota = v
	return s
}

func (s *UpdateWorkspaceResponseBodyResult) SetResourceGroupId(v string) *UpdateWorkspaceResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyResult) SetTags(v []*UpdateWorkspaceResponseBodyResultTags) *UpdateWorkspaceResponseBodyResult {
	s.Tags = v
	return s
}

func (s *UpdateWorkspaceResponseBodyResult) SetType(v string) *UpdateWorkspaceResponseBodyResult {
	s.Type = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyResult) Validate() error {
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

type UpdateWorkspaceResponseBodyResultQuota struct {
	// The compute resource.
	//
	// example:
	//
	// 0
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity.
	//
	// example:
	//
	// 0
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications.
	//
	// example:
	//
	// rag.share.common
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s UpdateWorkspaceResponseBodyResultQuota) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBodyResultQuota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *UpdateWorkspaceResponseBodyResultQuota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *UpdateWorkspaceResponseBodyResultQuota) GetSpec() *string {
	return s.Spec
}

func (s *UpdateWorkspaceResponseBodyResultQuota) SetComputeResource(v int32) *UpdateWorkspaceResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyResultQuota) SetDocSize(v int32) *UpdateWorkspaceResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyResultQuota) SetSpec(v string) *UpdateWorkspaceResponseBodyResultQuota {
	s.Spec = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyResultQuota) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkspaceResponseBodyResultTags struct {
	// The tag key.
	//
	// example:
	//
	// a
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s UpdateWorkspaceResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBodyResultTags) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateWorkspaceResponseBodyResultTags) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateWorkspaceResponseBodyResultTags) SetTagKey(v string) *UpdateWorkspaceResponseBodyResultTags {
	s.TagKey = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyResultTags) SetTagValue(v string) *UpdateWorkspaceResponseBodyResultTags {
	s.TagValue = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}

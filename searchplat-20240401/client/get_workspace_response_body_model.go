// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWorkspaceResponseBody
	GetRequestId() *string
	SetResult(v *GetWorkspaceResponseBodyResult) *GetWorkspaceResponseBody
	GetResult() *GetWorkspaceResponseBodyResult
}

type GetWorkspaceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 58113A95-1858-5674-87E5-192AEE6FD9DD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result *GetWorkspaceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspaceResponseBody) GetResult() *GetWorkspaceResponseBodyResult {
	return s.Result
}

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetResult(v *GetWorkspaceResponseBodyResult) *GetWorkspaceResponseBody {
	s.Result = v
	return s
}

func (s *GetWorkspaceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspaceResponseBodyResult struct {
	// apiToken
	//
	// example:
	//
	// apiToken
	ApiToken *string `json:"apiToken,omitempty" xml:"apiToken,omitempty"`
	// The billing type. Valid values:
	//
	// - POSTPAY: pay-as-you-go.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The commodity code.
	//
	// example:
	//
	// commodityCode
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
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
	// The quota.
	Quota *GetWorkspaceResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzaowqymbb4ki
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The status.
	//
	// example:
	//
	// ""
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tags.
	Tags []*GetWorkspaceResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The type.
	//
	// example:
	//
	// standard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetWorkspaceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyResult) GetApiToken() *string {
	return s.ApiToken
}

func (s *GetWorkspaceResponseBodyResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetWorkspaceResponseBodyResult) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetWorkspaceResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *GetWorkspaceResponseBodyResult) GetDomainName() *string {
	return s.DomainName
}

func (s *GetWorkspaceResponseBodyResult) GetEngineType() *string {
	return s.EngineType
}

func (s *GetWorkspaceResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *GetWorkspaceResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkspaceResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetWorkspaceResponseBodyResult) GetQuota() *GetWorkspaceResponseBodyResultQuota {
	return s.Quota
}

func (s *GetWorkspaceResponseBodyResult) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetWorkspaceResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetWorkspaceResponseBodyResult) GetTags() []*GetWorkspaceResponseBodyResultTags {
	return s.Tags
}

func (s *GetWorkspaceResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetWorkspaceResponseBodyResult) SetApiToken(v string) *GetWorkspaceResponseBodyResult {
	s.ApiToken = &v
	return s
}

func (s *GetWorkspaceResponseBodyResult) SetChargeType(v string) *GetWorkspaceResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *GetWorkspaceResponseBodyResult) SetCommodityCode(v string) *GetWorkspaceResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *GetWorkspaceResponseBodyResult) SetDescription(v string) *GetWorkspaceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetWorkspaceResponseBodyResult) SetDomainName(v string) *GetWorkspaceResponseBodyResult {
	s.DomainName = &v
	return s
}

func (s *GetWorkspaceResponseBodyResult) SetEngineType(v string) *GetWorkspaceResponseBodyResult {
	s.EngineType = &v
	return s
}

func (s *GetWorkspaceResponseBodyResult) SetId(v string) *GetWorkspaceResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetWorkspaceResponseBodyResult) SetInstanceId(v string) *GetWorkspaceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *GetWorkspaceResponseBodyResult) SetName(v string) *GetWorkspaceResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetWorkspaceResponseBodyResult) SetQuota(v *GetWorkspaceResponseBodyResultQuota) *GetWorkspaceResponseBodyResult {
	s.Quota = v
	return s
}

func (s *GetWorkspaceResponseBodyResult) SetResourceGroupId(v string) *GetWorkspaceResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *GetWorkspaceResponseBodyResult) SetStatus(v string) *GetWorkspaceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetWorkspaceResponseBodyResult) SetTags(v []*GetWorkspaceResponseBodyResultTags) *GetWorkspaceResponseBodyResult {
	s.Tags = v
	return s
}

func (s *GetWorkspaceResponseBodyResult) SetType(v string) *GetWorkspaceResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetWorkspaceResponseBodyResult) Validate() error {
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

type GetWorkspaceResponseBodyResultQuota struct {
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

func (s GetWorkspaceResponseBodyResultQuota) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyResultQuota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *GetWorkspaceResponseBodyResultQuota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *GetWorkspaceResponseBodyResultQuota) GetSpec() *string {
	return s.Spec
}

func (s *GetWorkspaceResponseBodyResultQuota) SetComputeResource(v int32) *GetWorkspaceResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *GetWorkspaceResponseBodyResultQuota) SetDocSize(v int32) *GetWorkspaceResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *GetWorkspaceResponseBodyResultQuota) SetSpec(v string) *GetWorkspaceResponseBodyResultQuota {
	s.Spec = &v
	return s
}

func (s *GetWorkspaceResponseBodyResultQuota) Validate() error {
	return dara.Validate(s)
}

type GetWorkspaceResponseBodyResultTags struct {
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
	// x
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s GetWorkspaceResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyResultTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetWorkspaceResponseBodyResultTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetWorkspaceResponseBodyResultTags) SetTagKey(v string) *GetWorkspaceResponseBodyResultTags {
	s.TagKey = &v
	return s
}

func (s *GetWorkspaceResponseBodyResultTags) SetTagValue(v string) *GetWorkspaceResponseBodyResultTags {
	s.TagValue = &v
	return s
}

func (s *GetWorkspaceResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}

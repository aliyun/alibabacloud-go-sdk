// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWorkspacesResponseBody
	GetRequestId() *string
	SetResult(v []*ListWorkspacesResponseBodyResult) *ListWorkspacesResponseBody
	GetResult() []*ListWorkspacesResponseBodyResult
	SetTotalCount(v int32) *ListWorkspacesResponseBody
	GetTotalCount() *int32
}

type ListWorkspacesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 33E4F0CA-F766-5803-B11C-70DC57A5A6E4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	Result []*ListWorkspacesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListWorkspacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspacesResponseBody) GetResult() []*ListWorkspacesResponseBodyResult {
	return s.Result
}

func (s *ListWorkspacesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetResult(v []*ListWorkspacesResponseBodyResult) *ListWorkspacesResponseBody {
	s.Result = v
	return s
}

func (s *ListWorkspacesResponseBody) SetTotalCount(v int32) *ListWorkspacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspacesResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspacesResponseBodyResult struct {
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
	// The description.
	//
	// example:
	//
	// 22222
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
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The quota information.
	Quota *ListWorkspacesResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzvlxzgo5b4si
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The tags.
	Tags []*ListWorkspacesResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The type.
	//
	// example:
	//
	// standard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListWorkspacesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyResult) GetApiToken() *string {
	return s.ApiToken
}

func (s *ListWorkspacesResponseBodyResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListWorkspacesResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListWorkspacesResponseBodyResult) GetDomainName() *string {
	return s.DomainName
}

func (s *ListWorkspacesResponseBodyResult) GetEngineType() *string {
	return s.EngineType
}

func (s *ListWorkspacesResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListWorkspacesResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWorkspacesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListWorkspacesResponseBodyResult) GetQuota() *ListWorkspacesResponseBodyResultQuota {
	return s.Quota
}

func (s *ListWorkspacesResponseBodyResult) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListWorkspacesResponseBodyResult) GetTags() []*ListWorkspacesResponseBodyResultTags {
	return s.Tags
}

func (s *ListWorkspacesResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListWorkspacesResponseBodyResult) SetApiToken(v string) *ListWorkspacesResponseBodyResult {
	s.ApiToken = &v
	return s
}

func (s *ListWorkspacesResponseBodyResult) SetChargeType(v string) *ListWorkspacesResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ListWorkspacesResponseBodyResult) SetDescription(v string) *ListWorkspacesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListWorkspacesResponseBodyResult) SetDomainName(v string) *ListWorkspacesResponseBodyResult {
	s.DomainName = &v
	return s
}

func (s *ListWorkspacesResponseBodyResult) SetEngineType(v string) *ListWorkspacesResponseBodyResult {
	s.EngineType = &v
	return s
}

func (s *ListWorkspacesResponseBodyResult) SetId(v string) *ListWorkspacesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListWorkspacesResponseBodyResult) SetInstanceId(v string) *ListWorkspacesResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyResult) SetName(v string) *ListWorkspacesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListWorkspacesResponseBodyResult) SetQuota(v *ListWorkspacesResponseBodyResultQuota) *ListWorkspacesResponseBodyResult {
	s.Quota = v
	return s
}

func (s *ListWorkspacesResponseBodyResult) SetResourceGroupId(v string) *ListWorkspacesResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *ListWorkspacesResponseBodyResult) SetTags(v []*ListWorkspacesResponseBodyResultTags) *ListWorkspacesResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListWorkspacesResponseBodyResult) SetType(v string) *ListWorkspacesResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListWorkspacesResponseBodyResult) Validate() error {
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

type ListWorkspacesResponseBodyResultQuota struct {
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
	// rag.share.compute
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ListWorkspacesResponseBodyResultQuota) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyResultQuota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *ListWorkspacesResponseBodyResultQuota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *ListWorkspacesResponseBodyResultQuota) GetSpec() *string {
	return s.Spec
}

func (s *ListWorkspacesResponseBodyResultQuota) SetComputeResource(v int32) *ListWorkspacesResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *ListWorkspacesResponseBodyResultQuota) SetDocSize(v int32) *ListWorkspacesResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *ListWorkspacesResponseBodyResultQuota) SetSpec(v string) *ListWorkspacesResponseBodyResultQuota {
	s.Spec = &v
	return s
}

func (s *ListWorkspacesResponseBodyResultQuota) Validate() error {
	return dara.Validate(s)
}

type ListWorkspacesResponseBodyResultTags struct {
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
	// c
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListWorkspacesResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyResultTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListWorkspacesResponseBodyResultTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListWorkspacesResponseBodyResultTags) SetTagKey(v string) *ListWorkspacesResponseBodyResultTags {
	s.TagKey = &v
	return s
}

func (s *ListWorkspacesResponseBodyResultTags) SetTagValue(v string) *ListWorkspacesResponseBodyResultTags {
	s.TagValue = &v
	return s
}

func (s *ListWorkspacesResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}

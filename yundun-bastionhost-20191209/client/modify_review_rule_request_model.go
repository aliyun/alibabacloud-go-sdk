// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReviewRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyReviewRuleRequest
	GetDescription() *string
	SetInstanceId(v string) *ModifyReviewRuleRequest
	GetInstanceId() *string
	SetName(v string) *ModifyReviewRuleRequest
	GetName() *string
	SetRegionId(v string) *ModifyReviewRuleRequest
	GetRegionId() *string
	SetReviewRuleId(v int64) *ModifyReviewRuleRequest
	GetReviewRuleId() *int64
	SetScopeType(v int32) *ModifyReviewRuleRequest
	GetScopeType() *int32
}

type ModifyReviewRuleRequest struct {
	// The new description of the audit policy. The description can be up to 500 characters in length.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-wwo36qbv601
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new name of the audit policy. The name must be 1 to 128 characters in length. It cannot start with a special character and can contain only the following special characters: periods (.), underscores (_), hyphens (-), backslashes (\\), and spaces.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the audit policy.
	//
	// > You can invoke the ListReviewRules operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReviewRuleId *int64 `json:"ReviewRuleId,omitempty" xml:"ReviewRuleId,omitempty"`
	// The type of the asset scope rule for auditable assets. Valid values:
	//
	// - **1**: All projects.
	//
	// - **2**: Specified projects.
	//
	// example:
	//
	// 2
	ScopeType *int32 `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s ModifyReviewRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyReviewRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyReviewRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyReviewRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyReviewRuleRequest) GetName() *string {
	return s.Name
}

func (s *ModifyReviewRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyReviewRuleRequest) GetReviewRuleId() *int64 {
	return s.ReviewRuleId
}

func (s *ModifyReviewRuleRequest) GetScopeType() *int32 {
	return s.ScopeType
}

func (s *ModifyReviewRuleRequest) SetDescription(v string) *ModifyReviewRuleRequest {
	s.Description = &v
	return s
}

func (s *ModifyReviewRuleRequest) SetInstanceId(v string) *ModifyReviewRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyReviewRuleRequest) SetName(v string) *ModifyReviewRuleRequest {
	s.Name = &v
	return s
}

func (s *ModifyReviewRuleRequest) SetRegionId(v string) *ModifyReviewRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyReviewRuleRequest) SetReviewRuleId(v int64) *ModifyReviewRuleRequest {
	s.ReviewRuleId = &v
	return s
}

func (s *ModifyReviewRuleRequest) SetScopeType(v int32) *ModifyReviewRuleRequest {
	s.ScopeType = &v
	return s
}

func (s *ModifyReviewRuleRequest) Validate() error {
	return dara.Validate(s)
}

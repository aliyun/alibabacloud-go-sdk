// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReviewRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateReviewRuleRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateReviewRuleRequest
	GetInstanceId() *string
	SetName(v string) *CreateReviewRuleRequest
	GetName() *string
	SetRegionId(v string) *CreateReviewRuleRequest
	GetRegionId() *string
	SetScopeType(v int32) *CreateReviewRuleRequest
	GetScopeType() *int32
}

type CreateReviewRuleRequest struct {
	// The description of the audit policy. The description can be up to 500 characters in length.
	//
	// example:
	//
	// session
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the bastion host instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the audit policy. The name must be 1 to 128 characters in length. It cannot start with a special character and can contain only the following special characters: periods (.), underscores (_), hyphens (-), backslashes (\\), and spaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the auditable asset scope rule. Valid values:
	//
	// - **1**: All projects.
	//
	// - **2**: Specified projects.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ScopeType *int32 `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s CreateReviewRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReviewRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateReviewRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateReviewRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateReviewRuleRequest) GetName() *string {
	return s.Name
}

func (s *CreateReviewRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateReviewRuleRequest) GetScopeType() *int32 {
	return s.ScopeType
}

func (s *CreateReviewRuleRequest) SetDescription(v string) *CreateReviewRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateReviewRuleRequest) SetInstanceId(v string) *CreateReviewRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateReviewRuleRequest) SetName(v string) *CreateReviewRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateReviewRuleRequest) SetRegionId(v string) *CreateReviewRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateReviewRuleRequest) SetScopeType(v int32) *CreateReviewRuleRequest {
	s.ScopeType = &v
	return s
}

func (s *CreateReviewRuleRequest) Validate() error {
	return dara.Validate(s)
}

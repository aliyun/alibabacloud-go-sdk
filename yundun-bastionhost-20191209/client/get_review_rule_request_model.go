// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReviewRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetReviewRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetReviewRuleRequest
	GetRegionId() *string
	SetReviewRuleId(v int64) *GetReviewRuleRequest
	GetReviewRuleId() *int64
}

type GetReviewRuleRequest struct {
	// The instance ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The audit policy ID.
	//
	// > You can invoke the ListReviewRules operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReviewRuleId *int64 `json:"ReviewRuleId,omitempty" xml:"ReviewRuleId,omitempty"`
}

func (s GetReviewRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReviewRuleRequest) GoString() string {
	return s.String()
}

func (s *GetReviewRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetReviewRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetReviewRuleRequest) GetReviewRuleId() *int64 {
	return s.ReviewRuleId
}

func (s *GetReviewRuleRequest) SetInstanceId(v string) *GetReviewRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetReviewRuleRequest) SetRegionId(v string) *GetReviewRuleRequest {
	s.RegionId = &v
	return s
}

func (s *GetReviewRuleRequest) SetReviewRuleId(v int64) *GetReviewRuleRequest {
	s.ReviewRuleId = &v
	return s
}

func (s *GetReviewRuleRequest) Validate() error {
	return dara.Validate(s)
}

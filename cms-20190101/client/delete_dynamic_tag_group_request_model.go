// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDynamicTagGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicTagRuleId(v string) *DeleteDynamicTagGroupRequest
	GetDynamicTagRuleId() *string
	SetRegionId(v string) *DeleteDynamicTagGroupRequest
	GetRegionId() *string
}

type DeleteDynamicTagGroupRequest struct {
	// The ID of the tag rule.
	//
	// For information about how to obtain the ID of a tag rule, see [DescribeDynamicTagRuleList](https://help.aliyun.com/document_detail/150126.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 6b882d9a-5117-42e2-9d0c-4749a0c6****
	DynamicTagRuleId *string `json:"DynamicTagRuleId,omitempty" xml:"DynamicTagRuleId,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDynamicTagGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDynamicTagGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDynamicTagGroupRequest) GetDynamicTagRuleId() *string {
	return s.DynamicTagRuleId
}

func (s *DeleteDynamicTagGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDynamicTagGroupRequest) SetDynamicTagRuleId(v string) *DeleteDynamicTagGroupRequest {
	s.DynamicTagRuleId = &v
	return s
}

func (s *DeleteDynamicTagGroupRequest) SetRegionId(v string) *DeleteDynamicTagGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDynamicTagGroupRequest) Validate() error {
	return dara.Validate(s)
}

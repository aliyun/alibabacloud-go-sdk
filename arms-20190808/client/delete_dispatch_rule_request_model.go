// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDispatchRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteDispatchRuleRequest
	GetId() *string
	SetRegionId(v string) *DeleteDispatchRuleRequest
	GetRegionId() *string
}

type DeleteDispatchRuleRequest struct {
	// The ID of the dispatch policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDispatchRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDispatchRuleRequest) GetId() *string {
	return s.Id
}

func (s *DeleteDispatchRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDispatchRuleRequest) SetId(v string) *DeleteDispatchRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteDispatchRuleRequest) SetRegionId(v string) *DeleteDispatchRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDispatchRuleRequest) Validate() error {
	return dara.Validate(s)
}

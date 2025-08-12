// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRuleBlackListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteMetricRuleBlackListRequest
	GetId() *string
	SetRegionId(v string) *DeleteMetricRuleBlackListRequest
	GetRegionId() *string
}

type DeleteMetricRuleBlackListRequest struct {
	// The IDs of the blacklist policies. Separate multiple IDs with commas (,). You can specify up to 50 IDs.
	//
	// For more information about how to obtain the ID of a blacklist policy, see [DescribeMetricRuleBlackList](https://help.aliyun.com/document_detail/457257.html).
	//
	// >  You can also set this parameter to a JSON array. Example: `["a9ad2ac2-3ed9-11ed-b878-0242ac12****","5cb8a9a4-198f-4651-a353-f8b28788****"]`.
	//
	// This parameter is required.
	//
	// example:
	//
	// a9ad2ac2-3ed9-11ed-b878-0242ac12****
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteMetricRuleBlackListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRuleBlackListRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleBlackListRequest) GetId() *string {
	return s.Id
}

func (s *DeleteMetricRuleBlackListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMetricRuleBlackListRequest) SetId(v string) *DeleteMetricRuleBlackListRequest {
	s.Id = &v
	return s
}

func (s *DeleteMetricRuleBlackListRequest) SetRegionId(v string) *DeleteMetricRuleBlackListRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMetricRuleBlackListRequest) Validate() error {
	return dara.Validate(s)
}

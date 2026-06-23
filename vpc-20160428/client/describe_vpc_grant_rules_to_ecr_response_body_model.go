// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcGrantRulesToEcrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGrantRuleModels(v []*DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) *DescribeVpcGrantRulesToEcrResponseBody
	GetGrantRuleModels() []*DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels
	SetNextToken(v string) *DescribeVpcGrantRulesToEcrResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeVpcGrantRulesToEcrResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeVpcGrantRulesToEcrResponseBody
	GetTotalCount() *string
}

type DescribeVpcGrantRulesToEcrResponseBody struct {
	// A list of authorization rules.
	GrantRuleModels []*DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels `json:"GrantRuleModels,omitempty" xml:"GrantRuleModels,omitempty" type:"Repeated"`
	// The pagination token for the next page of results. A non-empty value indicates that more results are available. If this parameter is not returned or is empty, all results have been retrieved.
	//
	// - If **NextToken*	- is empty, there are no more queries.
	//
	// - If **NextToken*	- has a value, the value is the token to start the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 66342E8E-5E87-5FF9-80C7-C3E3571A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries. Note: When you perform a paged query by using the `MaxResults` and `NextToken` parameters, the value of this parameter is for reference only.
	//
	// example:
	//
	// 10
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVpcGrantRulesToEcrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcGrantRulesToEcrResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcGrantRulesToEcrResponseBody) GetGrantRuleModels() []*DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels {
	return s.GrantRuleModels
}

func (s *DescribeVpcGrantRulesToEcrResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVpcGrantRulesToEcrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcGrantRulesToEcrResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeVpcGrantRulesToEcrResponseBody) SetGrantRuleModels(v []*DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) *DescribeVpcGrantRulesToEcrResponseBody {
	s.GrantRuleModels = v
	return s
}

func (s *DescribeVpcGrantRulesToEcrResponseBody) SetNextToken(v string) *DescribeVpcGrantRulesToEcrResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrResponseBody) SetRequestId(v string) *DescribeVpcGrantRulesToEcrResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrResponseBody) SetTotalCount(v string) *DescribeVpcGrantRulesToEcrResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrResponseBody) Validate() error {
	if s.GrantRuleModels != nil {
		for _, item := range s.GrantRuleModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels struct {
	// The time when the authorization was created.
	//
	// example:
	//
	// 2024-09-09T02:14:51Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the Express Connect Router.
	//
	// example:
	//
	// ecr-tz7w3chlaptxr2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the Alibaba Cloud account (main account) that owns the Express Connect Router.
	//
	// example:
	//
	// 192732132151****
	EcrOwnerId *int64 `json:"EcrOwnerId,omitempty" xml:"EcrOwnerId,omitempty"`
	// The ID of the network instance.
	//
	// example:
	//
	// vpc-wz9ek66wd7tl5xqpy****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Alibaba Cloud account (main account) that owns the network instance.
	//
	// example:
	//
	// 192745367151****
	InstanceUid *int64 `json:"InstanceUid,omitempty" xml:"InstanceUid,omitempty"`
	// The ID of the region where the network instance is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The type of the network instance. Valid values:
	//
	// - **VBR**: virtual border router
	//
	// - **VPC**: Virtual Private Cloud
	//
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) GoString() string {
	return s.String()
}

func (s *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) GetEcrOwnerId() *int64 {
	return s.EcrOwnerId
}

func (s *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) GetInstanceUid() *int64 {
	return s.InstanceUid
}

func (s *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) GetType() *string {
	return s.Type
}

func (s *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) SetCreationTime(v string) *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels {
	s.CreationTime = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) SetEcrId(v string) *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels {
	s.EcrId = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) SetEcrOwnerId(v int64) *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels {
	s.EcrOwnerId = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) SetInstanceId(v string) *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels {
	s.InstanceId = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) SetInstanceUid(v int64) *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels {
	s.InstanceUid = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) SetRegionNo(v string) *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) SetType(v string) *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels {
	s.Type = &v
	return s
}

func (s *DescribeVpcGrantRulesToEcrResponseBodyGrantRuleModels) Validate() error {
	return dara.Validate(s)
}

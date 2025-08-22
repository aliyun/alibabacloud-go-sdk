// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafScenesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseScenes(v []*DescribeDcdnWafScenesResponseBodyDefenseScenes) *DescribeDcdnWafScenesResponseBody
	GetDefenseScenes() []*DescribeDcdnWafScenesResponseBodyDefenseScenes
	SetRequestId(v string) *DescribeDcdnWafScenesResponseBody
	GetRequestId() *string
}

type DescribeDcdnWafScenesResponseBody struct {
	// The types of the protection policies.
	DefenseScenes []*DescribeDcdnWafScenesResponseBodyDefenseScenes `json:"DefenseScenes,omitempty" xml:"DefenseScenes,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 3D7BB13C-DD84-5654-A835-B8E1385DE274
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnWafScenesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafScenesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafScenesResponseBody) GetDefenseScenes() []*DescribeDcdnWafScenesResponseBodyDefenseScenes {
	return s.DefenseScenes
}

func (s *DescribeDcdnWafScenesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafScenesResponseBody) SetDefenseScenes(v []*DescribeDcdnWafScenesResponseBodyDefenseScenes) *DescribeDcdnWafScenesResponseBody {
	s.DefenseScenes = v
	return s
}

func (s *DescribeDcdnWafScenesResponseBody) SetRequestId(v string) *DescribeDcdnWafScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafScenesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafScenesResponseBodyDefenseScenes struct {
	// The type of the protection policy, which is the same as the DefenseScenes parameter in request parameters.
	//
	// example:
	//
	// waf_group
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The total number of policies of this type that were configured.
	//
	// example:
	//
	// 10
	PolicyCount *int32 `json:"PolicyCount,omitempty" xml:"PolicyCount,omitempty"`
	// The total number of protection rules that were configured in this type of the policy.
	//
	// example:
	//
	// 12
	RuleCount *int32 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
}

func (s DescribeDcdnWafScenesResponseBodyDefenseScenes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafScenesResponseBodyDefenseScenes) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafScenesResponseBodyDefenseScenes) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDcdnWafScenesResponseBodyDefenseScenes) GetPolicyCount() *int32 {
	return s.PolicyCount
}

func (s *DescribeDcdnWafScenesResponseBodyDefenseScenes) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *DescribeDcdnWafScenesResponseBodyDefenseScenes) SetDefenseScene(v string) *DescribeDcdnWafScenesResponseBodyDefenseScenes {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDcdnWafScenesResponseBodyDefenseScenes) SetPolicyCount(v int32) *DescribeDcdnWafScenesResponseBodyDefenseScenes {
	s.PolicyCount = &v
	return s
}

func (s *DescribeDcdnWafScenesResponseBodyDefenseScenes) SetRuleCount(v int32) *DescribeDcdnWafScenesResponseBodyDefenseScenes {
	s.RuleCount = &v
	return s
}

func (s *DescribeDcdnWafScenesResponseBodyDefenseScenes) Validate() error {
	return dara.Validate(s)
}

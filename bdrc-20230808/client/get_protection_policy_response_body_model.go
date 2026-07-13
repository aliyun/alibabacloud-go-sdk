// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProtectionPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetProtectionPolicyResponseBodyData) *GetProtectionPolicyResponseBody
	GetData() *GetProtectionPolicyResponseBodyData
	SetRequestId(v string) *GetProtectionPolicyResponseBody
	GetRequestId() *string
}

type GetProtectionPolicyResponseBody struct {
	// The data returned.
	Data *GetProtectionPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique ID of the request.
	//
	// example:
	//
	// AE43C4CB-8074-5EBD-9806-8CA6D12800B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProtectionPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProtectionPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetProtectionPolicyResponseBody) GetData() *GetProtectionPolicyResponseBodyData {
	return s.Data
}

func (s *GetProtectionPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProtectionPolicyResponseBody) SetData(v *GetProtectionPolicyResponseBodyData) *GetProtectionPolicyResponseBody {
	s.Data = v
	return s
}

func (s *GetProtectionPolicyResponseBody) SetRequestId(v string) *GetProtectionPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProtectionPolicyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProtectionPolicyResponseBodyData struct {
	// The list of associated resource category IDs.
	BoundResourceCategoryIds []*string `json:"BoundResourceCategoryIds,omitempty" xml:"BoundResourceCategoryIds,omitempty" type:"Repeated"`
	// The time when the policy was last applied.
	//
	// example:
	//
	// 1742167218
	LatestApplyTime *int64 `json:"LatestApplyTime,omitempty" xml:"LatestApplyTime,omitempty"`
	// The ID of the latest application task.
	//
	// example:
	//
	// t-123***7890
	LatestTaskId *string `json:"LatestTaskId,omitempty" xml:"LatestTaskId,omitempty"`
	// The protection policy ID.
	//
	// example:
	//
	// p-123***7890
	ProtectionPolicyId *string `json:"ProtectionPolicyId,omitempty" xml:"ProtectionPolicyId,omitempty"`
	// The protection policy name.
	//
	// example:
	//
	// MyProtectionPolicy
	ProtectionPolicyName *string `json:"ProtectionPolicyName,omitempty" xml:"ProtectionPolicyName,omitempty"`
	// The region ID of the protection policy.
	//
	// example:
	//
	// cn-hangzhou
	ProtectionPolicyRegionId *string `json:"ProtectionPolicyRegionId,omitempty" xml:"ProtectionPolicyRegionId,omitempty"`
	// The list of configured sub-protection policies.
	SubProtectionPolicies []*GetProtectionPolicyResponseBodyDataSubProtectionPolicies `json:"SubProtectionPolicies,omitempty" xml:"SubProtectionPolicies,omitempty" type:"Repeated"`
}

func (s GetProtectionPolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetProtectionPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProtectionPolicyResponseBodyData) GetBoundResourceCategoryIds() []*string {
	return s.BoundResourceCategoryIds
}

func (s *GetProtectionPolicyResponseBodyData) GetLatestApplyTime() *int64 {
	return s.LatestApplyTime
}

func (s *GetProtectionPolicyResponseBodyData) GetLatestTaskId() *string {
	return s.LatestTaskId
}

func (s *GetProtectionPolicyResponseBodyData) GetProtectionPolicyId() *string {
	return s.ProtectionPolicyId
}

func (s *GetProtectionPolicyResponseBodyData) GetProtectionPolicyName() *string {
	return s.ProtectionPolicyName
}

func (s *GetProtectionPolicyResponseBodyData) GetProtectionPolicyRegionId() *string {
	return s.ProtectionPolicyRegionId
}

func (s *GetProtectionPolicyResponseBodyData) GetSubProtectionPolicies() []*GetProtectionPolicyResponseBodyDataSubProtectionPolicies {
	return s.SubProtectionPolicies
}

func (s *GetProtectionPolicyResponseBodyData) SetBoundResourceCategoryIds(v []*string) *GetProtectionPolicyResponseBodyData {
	s.BoundResourceCategoryIds = v
	return s
}

func (s *GetProtectionPolicyResponseBodyData) SetLatestApplyTime(v int64) *GetProtectionPolicyResponseBodyData {
	s.LatestApplyTime = &v
	return s
}

func (s *GetProtectionPolicyResponseBodyData) SetLatestTaskId(v string) *GetProtectionPolicyResponseBodyData {
	s.LatestTaskId = &v
	return s
}

func (s *GetProtectionPolicyResponseBodyData) SetProtectionPolicyId(v string) *GetProtectionPolicyResponseBodyData {
	s.ProtectionPolicyId = &v
	return s
}

func (s *GetProtectionPolicyResponseBodyData) SetProtectionPolicyName(v string) *GetProtectionPolicyResponseBodyData {
	s.ProtectionPolicyName = &v
	return s
}

func (s *GetProtectionPolicyResponseBodyData) SetProtectionPolicyRegionId(v string) *GetProtectionPolicyResponseBodyData {
	s.ProtectionPolicyRegionId = &v
	return s
}

func (s *GetProtectionPolicyResponseBodyData) SetSubProtectionPolicies(v []*GetProtectionPolicyResponseBodyDataSubProtectionPolicies) *GetProtectionPolicyResponseBodyData {
	s.SubProtectionPolicies = v
	return s
}

func (s *GetProtectionPolicyResponseBodyData) Validate() error {
	if s.SubProtectionPolicies != nil {
		for _, item := range s.SubProtectionPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetProtectionPolicyResponseBodyDataSubProtectionPolicies struct {
	// The sub-protection policy configuration.
	//
	// example:
	//
	// {\\"autoSnapshotPolicyId\\":\\"sp-123***7890\\"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The sub-protection policy type.
	//
	// example:
	//
	// ECS_AUTO_SNAPSHOT_POLICY
	SubProtectionPolicyType *string `json:"SubProtectionPolicyType,omitempty" xml:"SubProtectionPolicyType,omitempty"`
}

func (s GetProtectionPolicyResponseBodyDataSubProtectionPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetProtectionPolicyResponseBodyDataSubProtectionPolicies) GoString() string {
	return s.String()
}

func (s *GetProtectionPolicyResponseBodyDataSubProtectionPolicies) GetConfig() *string {
	return s.Config
}

func (s *GetProtectionPolicyResponseBodyDataSubProtectionPolicies) GetSubProtectionPolicyType() *string {
	return s.SubProtectionPolicyType
}

func (s *GetProtectionPolicyResponseBodyDataSubProtectionPolicies) SetConfig(v string) *GetProtectionPolicyResponseBodyDataSubProtectionPolicies {
	s.Config = &v
	return s
}

func (s *GetProtectionPolicyResponseBodyDataSubProtectionPolicies) SetSubProtectionPolicyType(v string) *GetProtectionPolicyResponseBodyDataSubProtectionPolicies {
	s.SubProtectionPolicyType = &v
	return s
}

func (s *GetProtectionPolicyResponseBodyDataSubProtectionPolicies) Validate() error {
	return dara.Validate(s)
}

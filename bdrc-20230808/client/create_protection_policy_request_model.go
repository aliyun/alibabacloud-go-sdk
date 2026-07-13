// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProtectionPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBoundResourceCategoryIds(v []*string) *CreateProtectionPolicyRequest
	GetBoundResourceCategoryIds() []*string
	SetClientToken(v string) *CreateProtectionPolicyRequest
	GetClientToken() *string
	SetProtectionPolicyName(v string) *CreateProtectionPolicyRequest
	GetProtectionPolicyName() *string
	SetProtectionPolicyRegionId(v string) *CreateProtectionPolicyRequest
	GetProtectionPolicyRegionId() *string
	SetSubProtectionPolicies(v []*CreateProtectionPolicyRequestSubProtectionPolicies) *CreateProtectionPolicyRequest
	GetSubProtectionPolicies() []*CreateProtectionPolicyRequestSubProtectionPolicies
}

type CreateProtectionPolicyRequest struct {
	// The IDs of associated resource categories.
	BoundResourceCategoryIds []*string `json:"BoundResourceCategoryIds,omitempty" xml:"BoundResourceCategoryIds,omitempty" type:"Repeated"`
	// The client token used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// a1b2c3d4-****-****-****-a1b2c3d4f5e6
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the protection policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// My***Policy
	ProtectionPolicyName *string `json:"ProtectionPolicyName,omitempty" xml:"ProtectionPolicyName,omitempty"`
	// The region ID of the protection policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ProtectionPolicyRegionId *string `json:"ProtectionPolicyRegionId,omitempty" xml:"ProtectionPolicyRegionId,omitempty"`
	// The sub-protection policies.
	//
	// This parameter is required.
	SubProtectionPolicies []*CreateProtectionPolicyRequestSubProtectionPolicies `json:"SubProtectionPolicies,omitempty" xml:"SubProtectionPolicies,omitempty" type:"Repeated"`
}

func (s CreateProtectionPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectionPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateProtectionPolicyRequest) GetBoundResourceCategoryIds() []*string {
	return s.BoundResourceCategoryIds
}

func (s *CreateProtectionPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateProtectionPolicyRequest) GetProtectionPolicyName() *string {
	return s.ProtectionPolicyName
}

func (s *CreateProtectionPolicyRequest) GetProtectionPolicyRegionId() *string {
	return s.ProtectionPolicyRegionId
}

func (s *CreateProtectionPolicyRequest) GetSubProtectionPolicies() []*CreateProtectionPolicyRequestSubProtectionPolicies {
	return s.SubProtectionPolicies
}

func (s *CreateProtectionPolicyRequest) SetBoundResourceCategoryIds(v []*string) *CreateProtectionPolicyRequest {
	s.BoundResourceCategoryIds = v
	return s
}

func (s *CreateProtectionPolicyRequest) SetClientToken(v string) *CreateProtectionPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProtectionPolicyRequest) SetProtectionPolicyName(v string) *CreateProtectionPolicyRequest {
	s.ProtectionPolicyName = &v
	return s
}

func (s *CreateProtectionPolicyRequest) SetProtectionPolicyRegionId(v string) *CreateProtectionPolicyRequest {
	s.ProtectionPolicyRegionId = &v
	return s
}

func (s *CreateProtectionPolicyRequest) SetSubProtectionPolicies(v []*CreateProtectionPolicyRequestSubProtectionPolicies) *CreateProtectionPolicyRequest {
	s.SubProtectionPolicies = v
	return s
}

func (s *CreateProtectionPolicyRequest) Validate() error {
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

type CreateProtectionPolicyRequestSubProtectionPolicies struct {
	// The configuration of the sub-protection policy.
	//
	// example:
	//
	// {\\"PlaybookUuid\\": \\"2093d1ea-0651-48a6-bea2-fa7157285dc1\\", \\"ParamType\\": \\"custom\\", \\"InputParams\\": \\"\\"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The type of the sub-protection policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_AUTO_SNAPSHOT_POLICY
	SubProtectionPolicyType *string `json:"SubProtectionPolicyType,omitempty" xml:"SubProtectionPolicyType,omitempty"`
}

func (s CreateProtectionPolicyRequestSubProtectionPolicies) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectionPolicyRequestSubProtectionPolicies) GoString() string {
	return s.String()
}

func (s *CreateProtectionPolicyRequestSubProtectionPolicies) GetConfig() *string {
	return s.Config
}

func (s *CreateProtectionPolicyRequestSubProtectionPolicies) GetSubProtectionPolicyType() *string {
	return s.SubProtectionPolicyType
}

func (s *CreateProtectionPolicyRequestSubProtectionPolicies) SetConfig(v string) *CreateProtectionPolicyRequestSubProtectionPolicies {
	s.Config = &v
	return s
}

func (s *CreateProtectionPolicyRequestSubProtectionPolicies) SetSubProtectionPolicyType(v string) *CreateProtectionPolicyRequestSubProtectionPolicies {
	s.SubProtectionPolicyType = &v
	return s
}

func (s *CreateProtectionPolicyRequestSubProtectionPolicies) Validate() error {
	return dara.Validate(s)
}

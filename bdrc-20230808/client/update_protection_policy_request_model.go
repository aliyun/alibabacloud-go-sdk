// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProtectionPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBoundResourceCategoryIds(v []*string) *UpdateProtectionPolicyRequest
	GetBoundResourceCategoryIds() []*string
	SetClientToken(v string) *UpdateProtectionPolicyRequest
	GetClientToken() *string
	SetProtectionPolicyName(v string) *UpdateProtectionPolicyRequest
	GetProtectionPolicyName() *string
	SetSubProtectionPolicies(v []*UpdateProtectionPolicyRequestSubProtectionPolicies) *UpdateProtectionPolicyRequest
	GetSubProtectionPolicies() []*UpdateProtectionPolicyRequestSubProtectionPolicies
}

type UpdateProtectionPolicyRequest struct {
	// The list of associated resource category IDs.
	BoundResourceCategoryIds []*string `json:"BoundResourceCategoryIds,omitempty" xml:"BoundResourceCategoryIds,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters. If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// a1b2c3d4-****-****-****-a1b2c3d4f5e6
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The protection policy name.
	//
	// example:
	//
	// My***Policy
	ProtectionPolicyName *string `json:"ProtectionPolicyName,omitempty" xml:"ProtectionPolicyName,omitempty"`
	// The list of enabled sub-protection policies.
	SubProtectionPolicies []*UpdateProtectionPolicyRequestSubProtectionPolicies `json:"SubProtectionPolicies,omitempty" xml:"SubProtectionPolicies,omitempty" type:"Repeated"`
}

func (s UpdateProtectionPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectionPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateProtectionPolicyRequest) GetBoundResourceCategoryIds() []*string {
	return s.BoundResourceCategoryIds
}

func (s *UpdateProtectionPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateProtectionPolicyRequest) GetProtectionPolicyName() *string {
	return s.ProtectionPolicyName
}

func (s *UpdateProtectionPolicyRequest) GetSubProtectionPolicies() []*UpdateProtectionPolicyRequestSubProtectionPolicies {
	return s.SubProtectionPolicies
}

func (s *UpdateProtectionPolicyRequest) SetBoundResourceCategoryIds(v []*string) *UpdateProtectionPolicyRequest {
	s.BoundResourceCategoryIds = v
	return s
}

func (s *UpdateProtectionPolicyRequest) SetClientToken(v string) *UpdateProtectionPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateProtectionPolicyRequest) SetProtectionPolicyName(v string) *UpdateProtectionPolicyRequest {
	s.ProtectionPolicyName = &v
	return s
}

func (s *UpdateProtectionPolicyRequest) SetSubProtectionPolicies(v []*UpdateProtectionPolicyRequestSubProtectionPolicies) *UpdateProtectionPolicyRequest {
	s.SubProtectionPolicies = v
	return s
}

func (s *UpdateProtectionPolicyRequest) Validate() error {
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

type UpdateProtectionPolicyRequestSubProtectionPolicies struct {
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

func (s UpdateProtectionPolicyRequestSubProtectionPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectionPolicyRequestSubProtectionPolicies) GoString() string {
	return s.String()
}

func (s *UpdateProtectionPolicyRequestSubProtectionPolicies) GetConfig() *string {
	return s.Config
}

func (s *UpdateProtectionPolicyRequestSubProtectionPolicies) GetSubProtectionPolicyType() *string {
	return s.SubProtectionPolicyType
}

func (s *UpdateProtectionPolicyRequestSubProtectionPolicies) SetConfig(v string) *UpdateProtectionPolicyRequestSubProtectionPolicies {
	s.Config = &v
	return s
}

func (s *UpdateProtectionPolicyRequestSubProtectionPolicies) SetSubProtectionPolicyType(v string) *UpdateProtectionPolicyRequestSubProtectionPolicies {
	s.SubProtectionPolicyType = &v
	return s
}

func (s *UpdateProtectionPolicyRequestSubProtectionPolicies) Validate() error {
	return dara.Validate(s)
}

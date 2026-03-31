// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecodeDiagnosticMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDecodedDiagnosticMessage(v *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) *DecodeDiagnosticMessageResponseBody
	GetDecodedDiagnosticMessage() *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage
	SetRequestId(v string) *DecodeDiagnosticMessageResponseBody
	GetRequestId() *string
}

type DecodeDiagnosticMessageResponseBody struct {
	// The decoded diagnostic information.
	DecodedDiagnosticMessage *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage `json:"DecodedDiagnosticMessage,omitempty" xml:"DecodedDiagnosticMessage,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D2331703-AADF-5564-BA9B-26CD51A33BA0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DecodeDiagnosticMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DecodeDiagnosticMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DecodeDiagnosticMessageResponseBody) GetDecodedDiagnosticMessage() *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage {
	return s.DecodedDiagnosticMessage
}

func (s *DecodeDiagnosticMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DecodeDiagnosticMessageResponseBody) SetDecodedDiagnosticMessage(v *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) *DecodeDiagnosticMessageResponseBody {
	s.DecodedDiagnosticMessage = v
	return s
}

func (s *DecodeDiagnosticMessageResponseBody) SetRequestId(v string) *DecodeDiagnosticMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBody) Validate() error {
	if s.DecodedDiagnosticMessage != nil {
		if err := s.DecodedDiagnosticMessage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage struct {
	// The operation that is used for authentication in the request.
	//
	// example:
	//
	// ram:DecodeDiagnosticMessage
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The conditions that are used for authentication in the request.
	AuthConditions []*DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions `json:"AuthConditions,omitempty" xml:"AuthConditions,omitempty" type:"Repeated"`
	// The operator that is used for authentication in the request.
	AuthPrincipal *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal `json:"AuthPrincipal,omitempty" xml:"AuthPrincipal,omitempty" type:"Struct"`
	// The resource that is used for authentication in the request.
	//
	// example:
	//
	// *
	AuthResource *string `json:"AuthResource,omitempty" xml:"AuthResource,omitempty"`
	// Indicates whether the access denied error is caused by an explicit deny.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ExplicitDeny *bool `json:"ExplicitDeny,omitempty" xml:"ExplicitDeny,omitempty"`
	// The policies that are matched.
	MatchedPolicies []*DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies `json:"MatchedPolicies,omitempty" xml:"MatchedPolicies,omitempty" type:"Repeated"`
	// The type of the policy that causes the access denied error.
	//
	// Valid values:
	//
	// 	- AssumeRolePolicy: role-specific trust policy
	//
	// 	- ControlPolicy: control policy
	//
	// 	- AccountLevelIdentityBasedPolicy: identity-based policy at the account level
	//
	// 	- ResourceGroupLevelIdentityBasedPolicy: identity-based policy at the resource group level
	//
	// 	- SessionPolicy: session policy
	//
	// example:
	//
	// AccountLevelIdentityBasedPolicy
	NoPermissionPolicyType *string `json:"NoPermissionPolicyType,omitempty" xml:"NoPermissionPolicyType,omitempty"`
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) String() string {
	return dara.Prettify(s)
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) GoString() string {
	return s.String()
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) GetAuthConditions() []*DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions {
	return s.AuthConditions
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) GetAuthPrincipal() *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal {
	return s.AuthPrincipal
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) GetAuthResource() *string {
	return s.AuthResource
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) GetExplicitDeny() *bool {
	return s.ExplicitDeny
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) GetMatchedPolicies() []*DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies {
	return s.MatchedPolicies
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) GetNoPermissionPolicyType() *string {
	return s.NoPermissionPolicyType
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) SetAuthAction(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage {
	s.AuthAction = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) SetAuthConditions(v []*DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage {
	s.AuthConditions = v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) SetAuthPrincipal(v *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage {
	s.AuthPrincipal = v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) SetAuthResource(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage {
	s.AuthResource = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) SetExplicitDeny(v bool) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage {
	s.ExplicitDeny = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) SetMatchedPolicies(v []*DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage {
	s.MatchedPolicies = v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) SetNoPermissionPolicyType(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage {
	s.NoPermissionPolicyType = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessage) Validate() error {
	if s.AuthConditions != nil {
		for _, item := range s.AuthConditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AuthPrincipal != nil {
		if err := s.AuthPrincipal.Validate(); err != nil {
			return err
		}
	}
	if s.MatchedPolicies != nil {
		for _, item := range s.MatchedPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions struct {
	// The key of the condition.
	//
	// example:
	//
	// acs:SourceIp
	ConditionKey *string `json:"ConditionKey,omitempty" xml:"ConditionKey,omitempty"`
	// The values that correspond to the key.
	ConditionValues []*string `json:"ConditionValues,omitempty" xml:"ConditionValues,omitempty" type:"Repeated"`
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions) String() string {
	return dara.Prettify(s)
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions) GoString() string {
	return s.String()
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions) GetConditionKey() *string {
	return s.ConditionKey
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions) GetConditionValues() []*string {
	return s.ConditionValues
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions) SetConditionKey(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions {
	s.ConditionKey = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions) SetConditionValues(v []*string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions {
	s.ConditionValues = v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthConditions) Validate() error {
	return dara.Validate(s)
}

type DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal struct {
	// The identity.
	//
	// 	- If the operator is a RAM user, the ID of the user is displayed.
	//
	// 	- If the operator is a RAM role, the name and session name of the role are displayed. Example: RoleName:RoleSessionName.
	//
	// 	- If the operator is an SSO federated identity, the type and name of the identity provider (IdP) are displayed. Example: saml-provider/AzureAD.
	//
	// example:
	//
	// 28877424437521****
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The ID of the Alibaba Cloud account to which the identity belongs.
	//
	// example:
	//
	// 196813200012****
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The identity type that is used for authentication in the request.
	//
	// Valid values:
	//
	// 	- SubUser: RAM user
	//
	// 	- AssumedRoleUser: RAM role
	//
	// 	- Federated: SSO federated identity
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) String() string {
	return dara.Prettify(s)
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) GoString() string {
	return s.String()
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) SetAuthPrincipalDisplayName(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) SetAuthPrincipalOwnerId(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) SetAuthPrincipalType(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal {
	s.AuthPrincipalType = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageAuthPrincipal) Validate() error {
	return dara.Validate(s)
}

type DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies struct {
	// The type of the entity to which the policy is attached.
	//
	// Valid values:
	//
	// 	- RamUser: RAM user
	//
	// 	- RamRole: RAM role
	//
	// 	- ResourceDirectoryTarget: entity in a resource directory
	//
	// 	- RamGroup: RAM user group
	//
	// example:
	//
	// RamUser
	AttachedEntityType *string `json:"AttachedEntityType,omitempty" xml:"AttachedEntityType,omitempty"`
	// The authorization scope of the policy.
	//
	// Valid values:
	//
	// 	- Account: Alibaba Cloud account
	//
	// 	- Folder: folder in the resource directory
	//
	// 	- ResourceGroup: resource group
	//
	// example:
	//
	// Account
	AttachedScope *string `json:"AttachedScope,omitempty" xml:"AttachedScope,omitempty"`
	// The effect of the policy.
	//
	// Valid values:
	//
	// 	- Deny
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Allow
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Deny
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// The identifier of the policy.
	//
	// 	- Control policy: the ID of the control policy
	//
	// 	- RAM policy: the name of the policy
	//
	// example:
	//
	// MyPolicyName
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	// The type of the policy.
	//
	// Valid values:
	//
	// 	- Custom: custom policy
	//
	// 	- System: system policy
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The version number of the policy.
	//
	// > Only custom policies have version numbers.
	//
	// example:
	//
	// v1
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) String() string {
	return dara.Prettify(s)
}

func (s DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) GoString() string {
	return s.String()
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) GetAttachedEntityType() *string {
	return s.AttachedEntityType
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) GetAttachedScope() *string {
	return s.AttachedScope
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) GetEffect() *string {
	return s.Effect
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) SetAttachedEntityType(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies {
	s.AttachedEntityType = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) SetAttachedScope(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies {
	s.AttachedScope = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) SetEffect(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies {
	s.Effect = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) SetPolicyIdentifier(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies {
	s.PolicyIdentifier = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) SetPolicyType(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies {
	s.PolicyType = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) SetPolicyVersion(v string) *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies {
	s.PolicyVersion = &v
	return s
}

func (s *DecodeDiagnosticMessageResponseBodyDecodedDiagnosticMessageMatchedPolicies) Validate() error {
	return dara.Validate(s)
}

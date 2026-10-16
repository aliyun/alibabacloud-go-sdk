// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicySetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicySet(v *CreatePolicySetResponseBodyPolicySet) *CreatePolicySetResponseBody
	GetPolicySet() *CreatePolicySetResponseBodyPolicySet
	SetRequestId(v string) *CreatePolicySetResponseBody
	GetRequestId() *string
}

type CreatePolicySetResponseBody struct {
	// The policy set details.
	PolicySet *CreatePolicySetResponseBodyPolicySet `json:"PolicySet,omitempty" xml:"PolicySet,omitempty" type:"Struct"`
	// The unique identifier of the request.
	//
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolicySetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicySetResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicySetResponseBody) GetPolicySet() *CreatePolicySetResponseBodyPolicySet {
	return s.PolicySet
}

func (s *CreatePolicySetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolicySetResponseBody) SetPolicySet(v *CreatePolicySetResponseBodyPolicySet) *CreatePolicySetResponseBody {
	s.PolicySet = v
	return s
}

func (s *CreatePolicySetResponseBody) SetRequestId(v string) *CreatePolicySetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicySetResponseBody) Validate() error {
	if s.PolicySet != nil {
		if err := s.PolicySet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePolicySetResponseBodyPolicySet struct {
	// The creation time in UTC ISO 8601 format.
	//
	// example:
	//
	// 2026-05-08T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ARN of the policy set.
	//
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:policyset/default-policy-set
	PolicySetArn *string `json:"PolicySetArn,omitempty" xml:"PolicySetArn,omitempty"`
	// The name of the policy set. The name must be unique within the current Alibaba Cloud account.
	//
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
	// The source platform.
	//
	// example:
	//
	// Manuel
	SourcePlatform *string `json:"SourcePlatform,omitempty" xml:"SourcePlatform,omitempty"`
}

func (s CreatePolicySetResponseBodyPolicySet) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicySetResponseBodyPolicySet) GoString() string {
	return s.String()
}

func (s *CreatePolicySetResponseBodyPolicySet) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreatePolicySetResponseBodyPolicySet) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicySetResponseBodyPolicySet) GetPolicySetArn() *string {
	return s.PolicySetArn
}

func (s *CreatePolicySetResponseBodyPolicySet) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *CreatePolicySetResponseBodyPolicySet) GetSourcePlatform() *string {
	return s.SourcePlatform
}

func (s *CreatePolicySetResponseBodyPolicySet) SetCreateTime(v string) *CreatePolicySetResponseBodyPolicySet {
	s.CreateTime = &v
	return s
}

func (s *CreatePolicySetResponseBodyPolicySet) SetDescription(v string) *CreatePolicySetResponseBodyPolicySet {
	s.Description = &v
	return s
}

func (s *CreatePolicySetResponseBodyPolicySet) SetPolicySetArn(v string) *CreatePolicySetResponseBodyPolicySet {
	s.PolicySetArn = &v
	return s
}

func (s *CreatePolicySetResponseBodyPolicySet) SetPolicySetName(v string) *CreatePolicySetResponseBodyPolicySet {
	s.PolicySetName = &v
	return s
}

func (s *CreatePolicySetResponseBodyPolicySet) SetSourcePlatform(v string) *CreatePolicySetResponseBodyPolicySet {
	s.SourcePlatform = &v
	return s
}

func (s *CreatePolicySetResponseBodyPolicySet) Validate() error {
	return dara.Validate(s)
}

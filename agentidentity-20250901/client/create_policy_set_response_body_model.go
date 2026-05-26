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
	PolicySet *CreatePolicySetResponseBodyPolicySet `json:"PolicySet,omitempty" xml:"PolicySet,omitempty" type:"Struct"`
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
	// example:
	//
	// 2026-05-08T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:policyset/default-policy-set
	PolicySetArn *string `json:"PolicySetArn,omitempty" xml:"PolicySetArn,omitempty"`
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
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

func (s *CreatePolicySetResponseBodyPolicySet) Validate() error {
	return dara.Validate(s)
}

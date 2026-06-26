// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicySetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicySet(v *GetPolicySetResponseBodyPolicySet) *GetPolicySetResponseBody
	GetPolicySet() *GetPolicySetResponseBodyPolicySet
	SetRequestId(v string) *GetPolicySetResponseBody
	GetRequestId() *string
}

type GetPolicySetResponseBody struct {
	PolicySet *GetPolicySetResponseBodyPolicySet `json:"PolicySet,omitempty" xml:"PolicySet,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPolicySetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolicySetResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicySetResponseBody) GetPolicySet() *GetPolicySetResponseBodyPolicySet {
	return s.PolicySet
}

func (s *GetPolicySetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolicySetResponseBody) SetPolicySet(v *GetPolicySetResponseBodyPolicySet) *GetPolicySetResponseBody {
	s.PolicySet = v
	return s
}

func (s *GetPolicySetResponseBody) SetRequestId(v string) *GetPolicySetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolicySetResponseBody) Validate() error {
	if s.PolicySet != nil {
		if err := s.PolicySet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPolicySetResponseBodyPolicySet struct {
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicySetArn  *string `json:"PolicySetArn,omitempty" xml:"PolicySetArn,omitempty"`
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetPolicySetResponseBodyPolicySet) String() string {
	return dara.Prettify(s)
}

func (s GetPolicySetResponseBodyPolicySet) GoString() string {
	return s.String()
}

func (s *GetPolicySetResponseBodyPolicySet) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPolicySetResponseBodyPolicySet) GetDescription() *string {
	return s.Description
}

func (s *GetPolicySetResponseBodyPolicySet) GetPolicySetArn() *string {
	return s.PolicySetArn
}

func (s *GetPolicySetResponseBodyPolicySet) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *GetPolicySetResponseBodyPolicySet) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetPolicySetResponseBodyPolicySet) SetCreateTime(v string) *GetPolicySetResponseBodyPolicySet {
	s.CreateTime = &v
	return s
}

func (s *GetPolicySetResponseBodyPolicySet) SetDescription(v string) *GetPolicySetResponseBodyPolicySet {
	s.Description = &v
	return s
}

func (s *GetPolicySetResponseBodyPolicySet) SetPolicySetArn(v string) *GetPolicySetResponseBodyPolicySet {
	s.PolicySetArn = &v
	return s
}

func (s *GetPolicySetResponseBodyPolicySet) SetPolicySetName(v string) *GetPolicySetResponseBodyPolicySet {
	s.PolicySetName = &v
	return s
}

func (s *GetPolicySetResponseBodyPolicySet) SetUpdateTime(v string) *GetPolicySetResponseBodyPolicySet {
	s.UpdateTime = &v
	return s
}

func (s *GetPolicySetResponseBodyPolicySet) Validate() error {
	return dara.Validate(s)
}

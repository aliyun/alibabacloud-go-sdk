// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppPolicy(v *CreateAppPolicyResponseBodyAppPolicy) *CreateAppPolicyResponseBody
	GetAppPolicy() *CreateAppPolicyResponseBodyAppPolicy
	SetRequestId(v string) *CreateAppPolicyResponseBody
	GetRequestId() *string
}

type CreateAppPolicyResponseBody struct {
	AppPolicy *CreateAppPolicyResponseBodyAppPolicy `json:"AppPolicy,omitempty" xml:"AppPolicy,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppPolicyResponseBody) GetAppPolicy() *CreateAppPolicyResponseBodyAppPolicy {
	return s.AppPolicy
}

func (s *CreateAppPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppPolicyResponseBody) SetAppPolicy(v *CreateAppPolicyResponseBodyAppPolicy) *CreateAppPolicyResponseBody {
	s.AppPolicy = v
	return s
}

func (s *CreateAppPolicyResponseBody) SetRequestId(v string) *CreateAppPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateAppPolicyResponseBodyAppPolicy struct {
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	PolicyName       *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType       *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PolicyValue      *string `json:"PolicyValue,omitempty" xml:"PolicyValue,omitempty"`
	Products         *string `json:"Products,omitempty" xml:"Products,omitempty"`
}

func (s CreateAppPolicyResponseBodyAppPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateAppPolicyResponseBodyAppPolicy) GoString() string {
	return s.String()
}

func (s *CreateAppPolicyResponseBodyAppPolicy) GetCreationTime() *string {
	return s.CreationTime
}

func (s *CreateAppPolicyResponseBodyAppPolicy) GetDescription() *string {
	return s.Description
}

func (s *CreateAppPolicyResponseBodyAppPolicy) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *CreateAppPolicyResponseBodyAppPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreateAppPolicyResponseBodyAppPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateAppPolicyResponseBodyAppPolicy) GetPolicyValue() *string {
	return s.PolicyValue
}

func (s *CreateAppPolicyResponseBodyAppPolicy) GetProducts() *string {
	return s.Products
}

func (s *CreateAppPolicyResponseBodyAppPolicy) SetCreationTime(v string) *CreateAppPolicyResponseBodyAppPolicy {
	s.CreationTime = &v
	return s
}

func (s *CreateAppPolicyResponseBodyAppPolicy) SetDescription(v string) *CreateAppPolicyResponseBodyAppPolicy {
	s.Description = &v
	return s
}

func (s *CreateAppPolicyResponseBodyAppPolicy) SetModificationTime(v string) *CreateAppPolicyResponseBodyAppPolicy {
	s.ModificationTime = &v
	return s
}

func (s *CreateAppPolicyResponseBodyAppPolicy) SetPolicyName(v string) *CreateAppPolicyResponseBodyAppPolicy {
	s.PolicyName = &v
	return s
}

func (s *CreateAppPolicyResponseBodyAppPolicy) SetPolicyType(v string) *CreateAppPolicyResponseBodyAppPolicy {
	s.PolicyType = &v
	return s
}

func (s *CreateAppPolicyResponseBodyAppPolicy) SetPolicyValue(v string) *CreateAppPolicyResponseBodyAppPolicy {
	s.PolicyValue = &v
	return s
}

func (s *CreateAppPolicyResponseBodyAppPolicy) SetProducts(v string) *CreateAppPolicyResponseBodyAppPolicy {
	s.Products = &v
	return s
}

func (s *CreateAppPolicyResponseBodyAppPolicy) Validate() error {
	return dara.Validate(s)
}

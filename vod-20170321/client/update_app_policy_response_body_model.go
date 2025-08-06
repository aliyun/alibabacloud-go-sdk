// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppPolicy(v *UpdateAppPolicyResponseBodyAppPolicy) *UpdateAppPolicyResponseBody
	GetAppPolicy() *UpdateAppPolicyResponseBodyAppPolicy
	SetRequestId(v string) *UpdateAppPolicyResponseBody
	GetRequestId() *string
}

type UpdateAppPolicyResponseBody struct {
	AppPolicy *UpdateAppPolicyResponseBodyAppPolicy `json:"AppPolicy,omitempty" xml:"AppPolicy,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppPolicyResponseBody) GetAppPolicy() *UpdateAppPolicyResponseBodyAppPolicy {
	return s.AppPolicy
}

func (s *UpdateAppPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppPolicyResponseBody) SetAppPolicy(v *UpdateAppPolicyResponseBodyAppPolicy) *UpdateAppPolicyResponseBody {
	s.AppPolicy = v
	return s
}

func (s *UpdateAppPolicyResponseBody) SetRequestId(v string) *UpdateAppPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateAppPolicyResponseBodyAppPolicy struct {
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	PolicyName       *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType       *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PolicyValue      *string `json:"PolicyValue,omitempty" xml:"PolicyValue,omitempty"`
	Products         *string `json:"Products,omitempty" xml:"Products,omitempty"`
}

func (s UpdateAppPolicyResponseBodyAppPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppPolicyResponseBodyAppPolicy) GoString() string {
	return s.String()
}

func (s *UpdateAppPolicyResponseBodyAppPolicy) GetCreationTime() *string {
	return s.CreationTime
}

func (s *UpdateAppPolicyResponseBodyAppPolicy) GetDescription() *string {
	return s.Description
}

func (s *UpdateAppPolicyResponseBodyAppPolicy) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *UpdateAppPolicyResponseBodyAppPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *UpdateAppPolicyResponseBodyAppPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateAppPolicyResponseBodyAppPolicy) GetPolicyValue() *string {
	return s.PolicyValue
}

func (s *UpdateAppPolicyResponseBodyAppPolicy) GetProducts() *string {
	return s.Products
}

func (s *UpdateAppPolicyResponseBodyAppPolicy) SetCreationTime(v string) *UpdateAppPolicyResponseBodyAppPolicy {
	s.CreationTime = &v
	return s
}

func (s *UpdateAppPolicyResponseBodyAppPolicy) SetDescription(v string) *UpdateAppPolicyResponseBodyAppPolicy {
	s.Description = &v
	return s
}

func (s *UpdateAppPolicyResponseBodyAppPolicy) SetModificationTime(v string) *UpdateAppPolicyResponseBodyAppPolicy {
	s.ModificationTime = &v
	return s
}

func (s *UpdateAppPolicyResponseBodyAppPolicy) SetPolicyName(v string) *UpdateAppPolicyResponseBodyAppPolicy {
	s.PolicyName = &v
	return s
}

func (s *UpdateAppPolicyResponseBodyAppPolicy) SetPolicyType(v string) *UpdateAppPolicyResponseBodyAppPolicy {
	s.PolicyType = &v
	return s
}

func (s *UpdateAppPolicyResponseBodyAppPolicy) SetPolicyValue(v string) *UpdateAppPolicyResponseBodyAppPolicy {
	s.PolicyValue = &v
	return s
}

func (s *UpdateAppPolicyResponseBodyAppPolicy) SetProducts(v string) *UpdateAppPolicyResponseBodyAppPolicy {
	s.Products = &v
	return s
}

func (s *UpdateAppPolicyResponseBodyAppPolicy) Validate() error {
	return dara.Validate(s)
}

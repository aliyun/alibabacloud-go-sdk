// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppPolicyList(v []*GetAppPoliciesResponseBodyAppPolicyList) *GetAppPoliciesResponseBody
	GetAppPolicyList() []*GetAppPoliciesResponseBodyAppPolicyList
	SetNonExistPolicyNames(v []*string) *GetAppPoliciesResponseBody
	GetNonExistPolicyNames() []*string
	SetRequestId(v string) *GetAppPoliciesResponseBody
	GetRequestId() *string
}

type GetAppPoliciesResponseBody struct {
	AppPolicyList       []*GetAppPoliciesResponseBodyAppPolicyList `json:"AppPolicyList,omitempty" xml:"AppPolicyList,omitempty" type:"Repeated"`
	NonExistPolicyNames []*string                                  `json:"NonExistPolicyNames,omitempty" xml:"NonExistPolicyNames,omitempty" type:"Repeated"`
	RequestId           *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppPoliciesResponseBody) GetAppPolicyList() []*GetAppPoliciesResponseBodyAppPolicyList {
	return s.AppPolicyList
}

func (s *GetAppPoliciesResponseBody) GetNonExistPolicyNames() []*string {
	return s.NonExistPolicyNames
}

func (s *GetAppPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppPoliciesResponseBody) SetAppPolicyList(v []*GetAppPoliciesResponseBodyAppPolicyList) *GetAppPoliciesResponseBody {
	s.AppPolicyList = v
	return s
}

func (s *GetAppPoliciesResponseBody) SetNonExistPolicyNames(v []*string) *GetAppPoliciesResponseBody {
	s.NonExistPolicyNames = v
	return s
}

func (s *GetAppPoliciesResponseBody) SetRequestId(v string) *GetAppPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAppPoliciesResponseBodyAppPolicyList struct {
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	PolicyName       *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType       *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PolicyValue      *string `json:"PolicyValue,omitempty" xml:"PolicyValue,omitempty"`
	Products         *string `json:"Products,omitempty" xml:"Products,omitempty"`
}

func (s GetAppPoliciesResponseBodyAppPolicyList) String() string {
	return dara.Prettify(s)
}

func (s GetAppPoliciesResponseBodyAppPolicyList) GoString() string {
	return s.String()
}

func (s *GetAppPoliciesResponseBodyAppPolicyList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetAppPoliciesResponseBodyAppPolicyList) GetDescription() *string {
	return s.Description
}

func (s *GetAppPoliciesResponseBodyAppPolicyList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *GetAppPoliciesResponseBodyAppPolicyList) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetAppPoliciesResponseBodyAppPolicyList) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetAppPoliciesResponseBodyAppPolicyList) GetPolicyValue() *string {
	return s.PolicyValue
}

func (s *GetAppPoliciesResponseBodyAppPolicyList) GetProducts() *string {
	return s.Products
}

func (s *GetAppPoliciesResponseBodyAppPolicyList) SetCreationTime(v string) *GetAppPoliciesResponseBodyAppPolicyList {
	s.CreationTime = &v
	return s
}

func (s *GetAppPoliciesResponseBodyAppPolicyList) SetDescription(v string) *GetAppPoliciesResponseBodyAppPolicyList {
	s.Description = &v
	return s
}

func (s *GetAppPoliciesResponseBodyAppPolicyList) SetModificationTime(v string) *GetAppPoliciesResponseBodyAppPolicyList {
	s.ModificationTime = &v
	return s
}

func (s *GetAppPoliciesResponseBodyAppPolicyList) SetPolicyName(v string) *GetAppPoliciesResponseBodyAppPolicyList {
	s.PolicyName = &v
	return s
}

func (s *GetAppPoliciesResponseBodyAppPolicyList) SetPolicyType(v string) *GetAppPoliciesResponseBodyAppPolicyList {
	s.PolicyType = &v
	return s
}

func (s *GetAppPoliciesResponseBodyAppPolicyList) SetPolicyValue(v string) *GetAppPoliciesResponseBodyAppPolicyList {
	s.PolicyValue = &v
	return s
}

func (s *GetAppPoliciesResponseBodyAppPolicyList) SetProducts(v string) *GetAppPoliciesResponseBodyAppPolicyList {
	s.Products = &v
	return s
}

func (s *GetAppPoliciesResponseBodyAppPolicyList) Validate() error {
	return dara.Validate(s)
}

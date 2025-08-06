// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppPolicyList(v []*ListAppPolicyResponseBodyAppPolicyList) *ListAppPolicyResponseBody
	GetAppPolicyList() []*ListAppPolicyResponseBodyAppPolicyList
	SetRequestId(v string) *ListAppPolicyResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListAppPolicyResponseBody
	GetTotal() *int64
}

type ListAppPolicyResponseBody struct {
	AppPolicyList []*ListAppPolicyResponseBodyAppPolicyList `json:"AppPolicyList,omitempty" xml:"AppPolicyList,omitempty" type:"Repeated"`
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total         *int64                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAppPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppPolicyResponseBody) GetAppPolicyList() []*ListAppPolicyResponseBodyAppPolicyList {
	return s.AppPolicyList
}

func (s *ListAppPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppPolicyResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAppPolicyResponseBody) SetAppPolicyList(v []*ListAppPolicyResponseBodyAppPolicyList) *ListAppPolicyResponseBody {
	s.AppPolicyList = v
	return s
}

func (s *ListAppPolicyResponseBody) SetRequestId(v string) *ListAppPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppPolicyResponseBody) SetTotal(v int64) *ListAppPolicyResponseBody {
	s.Total = &v
	return s
}

func (s *ListAppPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAppPolicyResponseBodyAppPolicyList struct {
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	PolicyName       *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType       *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PolicyValue      *string `json:"PolicyValue,omitempty" xml:"PolicyValue,omitempty"`
	Products         *string `json:"Products,omitempty" xml:"Products,omitempty"`
}

func (s ListAppPolicyResponseBodyAppPolicyList) String() string {
	return dara.Prettify(s)
}

func (s ListAppPolicyResponseBodyAppPolicyList) GoString() string {
	return s.String()
}

func (s *ListAppPolicyResponseBodyAppPolicyList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAppPolicyResponseBodyAppPolicyList) GetDescription() *string {
	return s.Description
}

func (s *ListAppPolicyResponseBodyAppPolicyList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *ListAppPolicyResponseBodyAppPolicyList) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListAppPolicyResponseBodyAppPolicyList) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListAppPolicyResponseBodyAppPolicyList) GetPolicyValue() *string {
	return s.PolicyValue
}

func (s *ListAppPolicyResponseBodyAppPolicyList) GetProducts() *string {
	return s.Products
}

func (s *ListAppPolicyResponseBodyAppPolicyList) SetCreationTime(v string) *ListAppPolicyResponseBodyAppPolicyList {
	s.CreationTime = &v
	return s
}

func (s *ListAppPolicyResponseBodyAppPolicyList) SetDescription(v string) *ListAppPolicyResponseBodyAppPolicyList {
	s.Description = &v
	return s
}

func (s *ListAppPolicyResponseBodyAppPolicyList) SetModificationTime(v string) *ListAppPolicyResponseBodyAppPolicyList {
	s.ModificationTime = &v
	return s
}

func (s *ListAppPolicyResponseBodyAppPolicyList) SetPolicyName(v string) *ListAppPolicyResponseBodyAppPolicyList {
	s.PolicyName = &v
	return s
}

func (s *ListAppPolicyResponseBodyAppPolicyList) SetPolicyType(v string) *ListAppPolicyResponseBodyAppPolicyList {
	s.PolicyType = &v
	return s
}

func (s *ListAppPolicyResponseBodyAppPolicyList) SetPolicyValue(v string) *ListAppPolicyResponseBodyAppPolicyList {
	s.PolicyValue = &v
	return s
}

func (s *ListAppPolicyResponseBodyAppPolicyList) SetProducts(v string) *ListAppPolicyResponseBodyAppPolicyList {
	s.Products = &v
	return s
}

func (s *ListAppPolicyResponseBodyAppPolicyList) Validate() error {
	return dara.Validate(s)
}

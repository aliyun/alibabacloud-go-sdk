// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomPrivacyPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomPrivacyPolicies(v []*ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) *ListCustomPrivacyPoliciesResponseBody
	GetCustomPrivacyPolicies() []*ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies
	SetMaxResults(v int64) *ListCustomPrivacyPoliciesResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListCustomPrivacyPoliciesResponseBody
	GetNextToken() *string
	SetPreviousToken(v string) *ListCustomPrivacyPoliciesResponseBody
	GetPreviousToken() *string
	SetRequestId(v string) *ListCustomPrivacyPoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCustomPrivacyPoliciesResponseBody
	GetTotalCount() *int64
}

type ListCustomPrivacyPoliciesResponseBody struct {
	// The list of custom terms.
	CustomPrivacyPolicies []*ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies `json:"CustomPrivacyPolicies,omitempty" xml:"CustomPrivacyPolicies,omitempty" type:"Repeated"`
	// The number of entries per page in a paged query.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned by this call.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The pagination token returned by this call.
	//
	// example:
	//
	// PTxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries in the list.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomPrivacyPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPrivacyPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomPrivacyPoliciesResponseBody) GetCustomPrivacyPolicies() []*ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies {
	return s.CustomPrivacyPolicies
}

func (s *ListCustomPrivacyPoliciesResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListCustomPrivacyPoliciesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomPrivacyPoliciesResponseBody) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListCustomPrivacyPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomPrivacyPoliciesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCustomPrivacyPoliciesResponseBody) SetCustomPrivacyPolicies(v []*ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) *ListCustomPrivacyPoliciesResponseBody {
	s.CustomPrivacyPolicies = v
	return s
}

func (s *ListCustomPrivacyPoliciesResponseBody) SetMaxResults(v int64) *ListCustomPrivacyPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCustomPrivacyPoliciesResponseBody) SetNextToken(v string) *ListCustomPrivacyPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCustomPrivacyPoliciesResponseBody) SetPreviousToken(v string) *ListCustomPrivacyPoliciesResponseBody {
	s.PreviousToken = &v
	return s
}

func (s *ListCustomPrivacyPoliciesResponseBody) SetRequestId(v string) *ListCustomPrivacyPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomPrivacyPoliciesResponseBody) SetTotalCount(v int64) *ListCustomPrivacyPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomPrivacyPoliciesResponseBody) Validate() error {
	if s.CustomPrivacyPolicies != nil {
		for _, item := range s.CustomPrivacyPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies struct {
	// The custom term ID.
	//
	// example:
	//
	// pp_xxxxx
	CustomPrivacyPolicyId *string `json:"CustomPrivacyPolicyId,omitempty" xml:"CustomPrivacyPolicyId,omitempty"`
	// The custom term name.
	//
	// example:
	//
	// Custom Privacy Policy Name
	CustomPrivacyPolicyName *string `json:"CustomPrivacyPolicyName,omitempty" xml:"CustomPrivacyPolicyName,omitempty"`
	// The default language term entry. The value is obtained from the LanguageCode field returned by the ListLanguages operation.
	//
	// example:
	//
	// zh-Hans-CN
	DefaultLanguageCode *string `json:"DefaultLanguageCode,omitempty" xml:"DefaultLanguageCode,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The custom term status.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The consent type of the custom term.
	//
	// example:
	//
	// implied_consent
	UserConsentType *string `json:"UserConsentType,omitempty" xml:"UserConsentType,omitempty"`
}

func (s ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) GoString() string {
	return s.String()
}

func (s *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) GetCustomPrivacyPolicyId() *string {
	return s.CustomPrivacyPolicyId
}

func (s *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) GetCustomPrivacyPolicyName() *string {
	return s.CustomPrivacyPolicyName
}

func (s *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) GetDefaultLanguageCode() *string {
	return s.DefaultLanguageCode
}

func (s *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) GetStatus() *string {
	return s.Status
}

func (s *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) GetUserConsentType() *string {
	return s.UserConsentType
}

func (s *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) SetCustomPrivacyPolicyId(v string) *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies {
	s.CustomPrivacyPolicyId = &v
	return s
}

func (s *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) SetCustomPrivacyPolicyName(v string) *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies {
	s.CustomPrivacyPolicyName = &v
	return s
}

func (s *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) SetDefaultLanguageCode(v string) *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies {
	s.DefaultLanguageCode = &v
	return s
}

func (s *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) SetInstanceId(v string) *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies {
	s.InstanceId = &v
	return s
}

func (s *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) SetStatus(v string) *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies {
	s.Status = &v
	return s
}

func (s *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) SetUserConsentType(v string) *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies {
	s.UserConsentType = &v
	return s
}

func (s *ListCustomPrivacyPoliciesResponseBodyCustomPrivacyPolicies) Validate() error {
	return dara.Validate(s)
}

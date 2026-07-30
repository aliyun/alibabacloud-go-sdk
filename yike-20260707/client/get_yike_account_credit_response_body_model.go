// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeAccountCreditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreditInfo(v *GetYikeAccountCreditResponseBodyCreditInfo) *GetYikeAccountCreditResponseBody
	GetCreditInfo() *GetYikeAccountCreditResponseBodyCreditInfo
	SetMembershipInfo(v *GetYikeAccountCreditResponseBodyMembershipInfo) *GetYikeAccountCreditResponseBody
	GetMembershipInfo() *GetYikeAccountCreditResponseBodyMembershipInfo
	SetRequestId(v string) *GetYikeAccountCreditResponseBody
	GetRequestId() *string
}

type GetYikeAccountCreditResponseBody struct {
	// The credit information.
	CreditInfo *GetYikeAccountCreditResponseBodyCreditInfo `json:"CreditInfo,omitempty" xml:"CreditInfo,omitempty" type:"Struct"`
	// The membership information.
	MembershipInfo *GetYikeAccountCreditResponseBodyMembershipInfo `json:"MembershipInfo,omitempty" xml:"MembershipInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetYikeAccountCreditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAccountCreditResponseBody) GoString() string {
	return s.String()
}

func (s *GetYikeAccountCreditResponseBody) GetCreditInfo() *GetYikeAccountCreditResponseBodyCreditInfo {
	return s.CreditInfo
}

func (s *GetYikeAccountCreditResponseBody) GetMembershipInfo() *GetYikeAccountCreditResponseBodyMembershipInfo {
	return s.MembershipInfo
}

func (s *GetYikeAccountCreditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetYikeAccountCreditResponseBody) SetCreditInfo(v *GetYikeAccountCreditResponseBodyCreditInfo) *GetYikeAccountCreditResponseBody {
	s.CreditInfo = v
	return s
}

func (s *GetYikeAccountCreditResponseBody) SetMembershipInfo(v *GetYikeAccountCreditResponseBodyMembershipInfo) *GetYikeAccountCreditResponseBody {
	s.MembershipInfo = v
	return s
}

func (s *GetYikeAccountCreditResponseBody) SetRequestId(v string) *GetYikeAccountCreditResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYikeAccountCreditResponseBody) Validate() error {
	if s.CreditInfo != nil {
		if err := s.CreditInfo.Validate(); err != nil {
			return err
		}
	}
	if s.MembershipInfo != nil {
		if err := s.MembershipInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetYikeAccountCreditResponseBodyCreditInfo struct {
	// The total granted credits.
	//
	// example:
	//
	// 200
	GrantedCreditQuota *float64 `json:"GrantedCreditQuota,omitempty" xml:"GrantedCreditQuota,omitempty"`
	// The remaining granted credits.
	//
	// example:
	//
	// 0
	GrantedCreditQuotaUsage *float64 `json:"GrantedCreditQuotaUsage,omitempty" xml:"GrantedCreditQuotaUsage,omitempty"`
	// The total credits of the booster pack.
	//
	// example:
	//
	// 20000
	PackCreditQuota *float64 `json:"PackCreditQuota,omitempty" xml:"PackCreditQuota,omitempty"`
	// The remaining credits of the booster pack.
	//
	// example:
	//
	// 5000
	PackCreditQuotaUsage *float64 `json:"PackCreditQuotaUsage,omitempty" xml:"PackCreditQuotaUsage,omitempty"`
	// The total credits of the membership plan.
	//
	// example:
	//
	// 10000
	ResourceCreditQuota *float64 `json:"ResourceCreditQuota,omitempty" xml:"ResourceCreditQuota,omitempty"`
	// The remaining credits of the membership plan.
	//
	// example:
	//
	// 2000
	ResourceCreditQuotaUsage *float64 `json:"ResourceCreditQuotaUsage,omitempty" xml:"ResourceCreditQuotaUsage,omitempty"`
}

func (s GetYikeAccountCreditResponseBodyCreditInfo) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAccountCreditResponseBodyCreditInfo) GoString() string {
	return s.String()
}

func (s *GetYikeAccountCreditResponseBodyCreditInfo) GetGrantedCreditQuota() *float64 {
	return s.GrantedCreditQuota
}

func (s *GetYikeAccountCreditResponseBodyCreditInfo) GetGrantedCreditQuotaUsage() *float64 {
	return s.GrantedCreditQuotaUsage
}

func (s *GetYikeAccountCreditResponseBodyCreditInfo) GetPackCreditQuota() *float64 {
	return s.PackCreditQuota
}

func (s *GetYikeAccountCreditResponseBodyCreditInfo) GetPackCreditQuotaUsage() *float64 {
	return s.PackCreditQuotaUsage
}

func (s *GetYikeAccountCreditResponseBodyCreditInfo) GetResourceCreditQuota() *float64 {
	return s.ResourceCreditQuota
}

func (s *GetYikeAccountCreditResponseBodyCreditInfo) GetResourceCreditQuotaUsage() *float64 {
	return s.ResourceCreditQuotaUsage
}

func (s *GetYikeAccountCreditResponseBodyCreditInfo) SetGrantedCreditQuota(v float64) *GetYikeAccountCreditResponseBodyCreditInfo {
	s.GrantedCreditQuota = &v
	return s
}

func (s *GetYikeAccountCreditResponseBodyCreditInfo) SetGrantedCreditQuotaUsage(v float64) *GetYikeAccountCreditResponseBodyCreditInfo {
	s.GrantedCreditQuotaUsage = &v
	return s
}

func (s *GetYikeAccountCreditResponseBodyCreditInfo) SetPackCreditQuota(v float64) *GetYikeAccountCreditResponseBodyCreditInfo {
	s.PackCreditQuota = &v
	return s
}

func (s *GetYikeAccountCreditResponseBodyCreditInfo) SetPackCreditQuotaUsage(v float64) *GetYikeAccountCreditResponseBodyCreditInfo {
	s.PackCreditQuotaUsage = &v
	return s
}

func (s *GetYikeAccountCreditResponseBodyCreditInfo) SetResourceCreditQuota(v float64) *GetYikeAccountCreditResponseBodyCreditInfo {
	s.ResourceCreditQuota = &v
	return s
}

func (s *GetYikeAccountCreditResponseBodyCreditInfo) SetResourceCreditQuotaUsage(v float64) *GetYikeAccountCreditResponseBodyCreditInfo {
	s.ResourceCreditQuotaUsage = &v
	return s
}

func (s *GetYikeAccountCreditResponseBodyCreditInfo) Validate() error {
	return dara.Validate(s)
}

type GetYikeAccountCreditResponseBodyMembershipInfo struct {
	// The end time.
	//
	// example:
	//
	// 1784179281
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The membership level. Valid values:
	//
	// - basic: Basic Edition.
	//
	// - standard: Standard Edition.
	//
	// - professional: Ultimate Edition.
	//
	// - ultra: Ultra Edition.
	//
	// example:
	//
	// basic
	Membership *string `json:"Membership,omitempty" xml:"Membership,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1784784081
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetYikeAccountCreditResponseBodyMembershipInfo) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAccountCreditResponseBodyMembershipInfo) GoString() string {
	return s.String()
}

func (s *GetYikeAccountCreditResponseBodyMembershipInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *GetYikeAccountCreditResponseBodyMembershipInfo) GetMembership() *string {
	return s.Membership
}

func (s *GetYikeAccountCreditResponseBodyMembershipInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *GetYikeAccountCreditResponseBodyMembershipInfo) SetEndTime(v string) *GetYikeAccountCreditResponseBodyMembershipInfo {
	s.EndTime = &v
	return s
}

func (s *GetYikeAccountCreditResponseBodyMembershipInfo) SetMembership(v string) *GetYikeAccountCreditResponseBodyMembershipInfo {
	s.Membership = &v
	return s
}

func (s *GetYikeAccountCreditResponseBodyMembershipInfo) SetStartTime(v string) *GetYikeAccountCreditResponseBodyMembershipInfo {
	s.StartTime = &v
	return s
}

func (s *GetYikeAccountCreditResponseBodyMembershipInfo) Validate() error {
	return dara.Validate(s)
}

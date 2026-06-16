// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUsersForAuthorizationRuleResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListUsersForAuthorizationRuleResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUsersForAuthorizationRuleResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListUsersForAuthorizationRuleResponseBody
	GetTotalCount() *int64
	SetUsers(v []*ListUsersForAuthorizationRuleResponseBodyUsers) *ListUsersForAuthorizationRuleResponseBody
	GetUsers() []*ListUsersForAuthorizationRuleResponseBodyUsers
}

type ListUsersForAuthorizationRuleResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token returned from the current call. Use this token to start the next paged query.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of accounts.
	Users []*ListUsersForAuthorizationRuleResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersForAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersForAuthorizationRuleResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersForAuthorizationRuleResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUsersForAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersForAuthorizationRuleResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUsersForAuthorizationRuleResponseBody) GetUsers() []*ListUsersForAuthorizationRuleResponseBodyUsers {
	return s.Users
}

func (s *ListUsersForAuthorizationRuleResponseBody) SetMaxResults(v int32) *ListUsersForAuthorizationRuleResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUsersForAuthorizationRuleResponseBody) SetNextToken(v string) *ListUsersForAuthorizationRuleResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUsersForAuthorizationRuleResponseBody) SetRequestId(v string) *ListUsersForAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersForAuthorizationRuleResponseBody) SetTotalCount(v int64) *ListUsersForAuthorizationRuleResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersForAuthorizationRuleResponseBody) SetUsers(v []*ListUsersForAuthorizationRuleResponseBodyUsers) *ListUsersForAuthorizationRuleResponseBody {
	s.Users = v
	return s
}

func (s *ListUsersForAuthorizationRuleResponseBody) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersForAuthorizationRuleResponseBodyUsers struct {
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The time range of the validity period. This parameter takes effect only when **ValidityType*	- is set to **time_bound**.
	ValidityPeriod *ListUsersForAuthorizationRuleResponseBodyUsersValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// The type of the validity period for the relationship. Valid values:
	//
	// - permanent: permanent
	//
	// - time_bound: custom time range
	//
	// example:
	//
	// permanent
	ValidityType *string `json:"ValidityType,omitempty" xml:"ValidityType,omitempty"`
}

func (s ListUsersForAuthorizationRuleResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForAuthorizationRuleResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersForAuthorizationRuleResponseBodyUsers) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUsersForAuthorizationRuleResponseBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersForAuthorizationRuleResponseBodyUsers) GetValidityPeriod() *ListUsersForAuthorizationRuleResponseBodyUsersValidityPeriod {
	return s.ValidityPeriod
}

func (s *ListUsersForAuthorizationRuleResponseBodyUsers) GetValidityType() *string {
	return s.ValidityType
}

func (s *ListUsersForAuthorizationRuleResponseBodyUsers) SetInstanceId(v string) *ListUsersForAuthorizationRuleResponseBodyUsers {
	s.InstanceId = &v
	return s
}

func (s *ListUsersForAuthorizationRuleResponseBodyUsers) SetUserId(v string) *ListUsersForAuthorizationRuleResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListUsersForAuthorizationRuleResponseBodyUsers) SetValidityPeriod(v *ListUsersForAuthorizationRuleResponseBodyUsersValidityPeriod) *ListUsersForAuthorizationRuleResponseBodyUsers {
	s.ValidityPeriod = v
	return s
}

func (s *ListUsersForAuthorizationRuleResponseBodyUsers) SetValidityType(v string) *ListUsersForAuthorizationRuleResponseBodyUsers {
	s.ValidityType = &v
	return s
}

func (s *ListUsersForAuthorizationRuleResponseBodyUsers) Validate() error {
	if s.ValidityPeriod != nil {
		if err := s.ValidityPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUsersForAuthorizationRuleResponseBodyUsersValidityPeriod struct {
	// The end time of the validity period. This is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1704042061000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the validity period. This is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1704042061000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListUsersForAuthorizationRuleResponseBodyUsersValidityPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForAuthorizationRuleResponseBodyUsersValidityPeriod) GoString() string {
	return s.String()
}

func (s *ListUsersForAuthorizationRuleResponseBodyUsersValidityPeriod) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListUsersForAuthorizationRuleResponseBodyUsersValidityPeriod) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListUsersForAuthorizationRuleResponseBodyUsersValidityPeriod) SetEndTime(v int64) *ListUsersForAuthorizationRuleResponseBodyUsersValidityPeriod {
	s.EndTime = &v
	return s
}

func (s *ListUsersForAuthorizationRuleResponseBodyUsersValidityPeriod) SetStartTime(v int64) *ListUsersForAuthorizationRuleResponseBodyUsersValidityPeriod {
	s.StartTime = &v
	return s
}

func (s *ListUsersForAuthorizationRuleResponseBodyUsersValidityPeriod) Validate() error {
	return dara.Validate(s)
}

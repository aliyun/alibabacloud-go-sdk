// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v []*ListGroupsForAuthorizationRuleResponseBodyGroups) *ListGroupsForAuthorizationRuleResponseBody
	GetGroups() []*ListGroupsForAuthorizationRuleResponseBodyGroups
	SetMaxResults(v int32) *ListGroupsForAuthorizationRuleResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListGroupsForAuthorizationRuleResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListGroupsForAuthorizationRuleResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListGroupsForAuthorizationRuleResponseBody
	GetTotalCount() *int64
}

type ListGroupsForAuthorizationRuleResponseBody struct {
	// The list of groups.
	Groups []*ListGroupsForAuthorizationRuleResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is returned for the next page.
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
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupsForAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsForAuthorizationRuleResponseBody) GetGroups() []*ListGroupsForAuthorizationRuleResponseBodyGroups {
	return s.Groups
}

func (s *ListGroupsForAuthorizationRuleResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGroupsForAuthorizationRuleResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupsForAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupsForAuthorizationRuleResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListGroupsForAuthorizationRuleResponseBody) SetGroups(v []*ListGroupsForAuthorizationRuleResponseBodyGroups) *ListGroupsForAuthorizationRuleResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsForAuthorizationRuleResponseBody) SetMaxResults(v int32) *ListGroupsForAuthorizationRuleResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsForAuthorizationRuleResponseBody) SetNextToken(v string) *ListGroupsForAuthorizationRuleResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGroupsForAuthorizationRuleResponseBody) SetRequestId(v string) *ListGroupsForAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsForAuthorizationRuleResponseBody) SetTotalCount(v int64) *ListGroupsForAuthorizationRuleResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGroupsForAuthorizationRuleResponseBody) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupsForAuthorizationRuleResponseBodyGroups struct {
	// The group ID.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time range of the validity period. This parameter takes effect only when **ValidityType*	- is set to **time_bound**.
	ValidityPeriod *ListGroupsForAuthorizationRuleResponseBodyGroupsValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// The type of the validity period of the relationship. Valid values:
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

func (s ListGroupsForAuthorizationRuleResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForAuthorizationRuleResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsForAuthorizationRuleResponseBodyGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupsForAuthorizationRuleResponseBodyGroups) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGroupsForAuthorizationRuleResponseBodyGroups) GetValidityPeriod() *ListGroupsForAuthorizationRuleResponseBodyGroupsValidityPeriod {
	return s.ValidityPeriod
}

func (s *ListGroupsForAuthorizationRuleResponseBodyGroups) GetValidityType() *string {
	return s.ValidityType
}

func (s *ListGroupsForAuthorizationRuleResponseBodyGroups) SetGroupId(v string) *ListGroupsForAuthorizationRuleResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListGroupsForAuthorizationRuleResponseBodyGroups) SetInstanceId(v string) *ListGroupsForAuthorizationRuleResponseBodyGroups {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsForAuthorizationRuleResponseBodyGroups) SetValidityPeriod(v *ListGroupsForAuthorizationRuleResponseBodyGroupsValidityPeriod) *ListGroupsForAuthorizationRuleResponseBodyGroups {
	s.ValidityPeriod = v
	return s
}

func (s *ListGroupsForAuthorizationRuleResponseBodyGroups) SetValidityType(v string) *ListGroupsForAuthorizationRuleResponseBodyGroups {
	s.ValidityType = &v
	return s
}

func (s *ListGroupsForAuthorizationRuleResponseBodyGroups) Validate() error {
	if s.ValidityPeriod != nil {
		if err := s.ValidityPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGroupsForAuthorizationRuleResponseBodyGroupsValidityPeriod struct {
	// The end time of the validity period. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1704042061000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the validity period. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1704042061000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListGroupsForAuthorizationRuleResponseBodyGroupsValidityPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForAuthorizationRuleResponseBodyGroupsValidityPeriod) GoString() string {
	return s.String()
}

func (s *ListGroupsForAuthorizationRuleResponseBodyGroupsValidityPeriod) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListGroupsForAuthorizationRuleResponseBodyGroupsValidityPeriod) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListGroupsForAuthorizationRuleResponseBodyGroupsValidityPeriod) SetEndTime(v int64) *ListGroupsForAuthorizationRuleResponseBodyGroupsValidityPeriod {
	s.EndTime = &v
	return s
}

func (s *ListGroupsForAuthorizationRuleResponseBodyGroupsValidityPeriod) SetStartTime(v int64) *ListGroupsForAuthorizationRuleResponseBodyGroupsValidityPeriod {
	s.StartTime = &v
	return s
}

func (s *ListGroupsForAuthorizationRuleResponseBodyGroupsValidityPeriod) Validate() error {
	return dara.Validate(s)
}

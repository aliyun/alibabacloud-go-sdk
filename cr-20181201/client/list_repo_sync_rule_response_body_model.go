// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoSyncRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRepoSyncRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListRepoSyncRuleResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListRepoSyncRuleResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepoSyncRuleResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRepoSyncRuleResponseBody
	GetRequestId() *string
	SetSyncRules(v []*ListRepoSyncRuleResponseBodySyncRules) *ListRepoSyncRuleResponseBody
	GetSyncRules() []*ListRepoSyncRuleResponseBodySyncRules
	SetTotalCount(v int32) *ListRepoSyncRuleResponseBody
	GetTotalCount() *int32
}

type ListRepoSyncRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 838D1602-6D8F-47FB-B60A-656645D2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried synchronization rules.
	SyncRules []*ListRepoSyncRuleResponseBodySyncRules `json:"SyncRules,omitempty" xml:"SyncRules,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRepoSyncRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepoSyncRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoSyncRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRepoSyncRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListRepoSyncRuleResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRepoSyncRuleResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRepoSyncRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepoSyncRuleResponseBody) GetSyncRules() []*ListRepoSyncRuleResponseBodySyncRules {
	return s.SyncRules
}

func (s *ListRepoSyncRuleResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRepoSyncRuleResponseBody) SetCode(v string) *ListRepoSyncRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoSyncRuleResponseBody) SetIsSuccess(v bool) *ListRepoSyncRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoSyncRuleResponseBody) SetPageNo(v int32) *ListRepoSyncRuleResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepoSyncRuleResponseBody) SetPageSize(v int32) *ListRepoSyncRuleResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepoSyncRuleResponseBody) SetRequestId(v string) *ListRepoSyncRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoSyncRuleResponseBody) SetSyncRules(v []*ListRepoSyncRuleResponseBodySyncRules) *ListRepoSyncRuleResponseBody {
	s.SyncRules = v
	return s
}

func (s *ListRepoSyncRuleResponseBody) SetTotalCount(v int32) *ListRepoSyncRuleResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRepoSyncRuleResponseBody) Validate() error {
	if s.SyncRules != nil {
		for _, item := range s.SyncRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRepoSyncRuleResponseBodySyncRules struct {
	// The time when the synchronization rule was created.
	//
	// example:
	//
	// 1572604642000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the synchronization is performed across Alibaba Cloud accounts. Valid values:
	//
	// 	- `true`: Images are synchronized across Alibaba Cloud accounts.
	//
	// 	- `false`: Images are synchronized within the same Alibaba Cloud account.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	CrossUser *bool `json:"CrossUser,omitempty" xml:"CrossUser,omitempty"`
	// The ID of the source instance.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	LocalInstanceId *string `json:"LocalInstanceId,omitempty" xml:"LocalInstanceId,omitempty"`
	// The name of the namespace in the source instance.
	//
	// example:
	//
	// test
	LocalNamespaceName *string `json:"LocalNamespaceName,omitempty" xml:"LocalNamespaceName,omitempty"`
	// The region ID of the source instance.
	//
	// example:
	//
	// cn-shanghai
	LocalRegionId *string `json:"LocalRegionId,omitempty" xml:"LocalRegionId,omitempty"`
	// The name of the repository in the source instance.
	//
	// example:
	//
	// test-repo-local
	LocalRepoName *string `json:"LocalRepoName,omitempty" xml:"LocalRepoName,omitempty"`
	// The time when the synchronization rule was last modified.
	//
	// example:
	//
	// 1572604642000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The regular expression that is used to filter repositories.
	//
	// >  This parameter is valid only when SyncScope is set to `NAMESPACE`.
	//
	// example:
	//
	// .*
	RepoNameFilter *string `json:"RepoNameFilter,omitempty" xml:"RepoNameFilter,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- `FROM`: Images are synchronized from the source instance to the destination instance.
	//
	// 	- `TO`: Images are synchronized from the destination instance to the source instance.
	//
	// example:
	//
	// FROM
	SyncDirection *string `json:"SyncDirection,omitempty" xml:"SyncDirection,omitempty"`
	// The ID of the synchronization rule.
	//
	// example:
	//
	// crsr-7lph66uloi6h****
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
	// The name of the synchronization rule.
	//
	// example:
	//
	// sync-rule-1
	SyncRuleName *string `json:"SyncRuleName,omitempty" xml:"SyncRuleName,omitempty"`
	// The synchronization scope. Valid values:
	//
	// 	- `NAMESPACE`: synchronizes the image tags in a namespace that meet the synchronization rule.
	//
	// 	- `REPO`: synchronizes the image tags in an image repository that meet the synchronization rule.
	//
	// example:
	//
	// NAMESPACE
	SyncScope *string `json:"SyncScope,omitempty" xml:"SyncScope,omitempty"`
	// The policy that is applied to trigger the synchronization rule. Valid values:
	//
	// 	- `INITIATIVE`: The synchronization rule is positively triggered.
	//
	// 	- `PASSIVE`: The synchronization rule is passively triggered.
	//
	// example:
	//
	// PASSIVE
	SyncTrigger *string `json:"SyncTrigger,omitempty" xml:"SyncTrigger,omitempty"`
	// The regular expression that is used to filter image tags.
	//
	// example:
	//
	// .*
	TagFilter *string `json:"TagFilter,omitempty" xml:"TagFilter,omitempty"`
	// The ID of the destination instance.
	//
	// example:
	//
	// cri-k77rd2eo9ztt****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The name of the namespace in the destination instance.
	//
	// example:
	//
	// test
	TargetNamespaceName *string `json:"TargetNamespaceName,omitempty" xml:"TargetNamespaceName,omitempty"`
	// The region ID of the destination instance.
	//
	// example:
	//
	// cn-shenzhen
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
	// The name of the repository in the destination instance.
	//
	// example:
	//
	// test-repo-target
	TargetRepoName *string `json:"TargetRepoName,omitempty" xml:"TargetRepoName,omitempty"`
}

func (s ListRepoSyncRuleResponseBodySyncRules) String() string {
	return dara.Prettify(s)
}

func (s ListRepoSyncRuleResponseBodySyncRules) GoString() string {
	return s.String()
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetCrossUser() *bool {
	return s.CrossUser
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetLocalInstanceId() *string {
	return s.LocalInstanceId
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetLocalNamespaceName() *string {
	return s.LocalNamespaceName
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetLocalRegionId() *string {
	return s.LocalRegionId
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetLocalRepoName() *string {
	return s.LocalRepoName
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetRepoNameFilter() *string {
	return s.RepoNameFilter
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetSyncDirection() *string {
	return s.SyncDirection
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetSyncRuleId() *string {
	return s.SyncRuleId
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetSyncRuleName() *string {
	return s.SyncRuleName
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetSyncScope() *string {
	return s.SyncScope
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetSyncTrigger() *string {
	return s.SyncTrigger
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetTagFilter() *string {
	return s.TagFilter
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetTargetNamespaceName() *string {
	return s.TargetNamespaceName
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetTargetRegionId() *string {
	return s.TargetRegionId
}

func (s *ListRepoSyncRuleResponseBodySyncRules) GetTargetRepoName() *string {
	return s.TargetRepoName
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetCreateTime(v int64) *ListRepoSyncRuleResponseBodySyncRules {
	s.CreateTime = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetCrossUser(v bool) *ListRepoSyncRuleResponseBodySyncRules {
	s.CrossUser = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetLocalInstanceId(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.LocalInstanceId = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetLocalNamespaceName(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.LocalNamespaceName = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetLocalRegionId(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.LocalRegionId = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetLocalRepoName(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.LocalRepoName = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetModifiedTime(v int64) *ListRepoSyncRuleResponseBodySyncRules {
	s.ModifiedTime = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetRepoNameFilter(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.RepoNameFilter = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetSyncDirection(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.SyncDirection = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetSyncRuleId(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.SyncRuleId = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetSyncRuleName(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.SyncRuleName = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetSyncScope(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.SyncScope = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetSyncTrigger(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.SyncTrigger = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetTagFilter(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.TagFilter = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetTargetInstanceId(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.TargetInstanceId = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetTargetNamespaceName(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.TargetNamespaceName = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetTargetRegionId(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.TargetRegionId = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetTargetRepoName(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.TargetRepoName = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) Validate() error {
	return dara.Validate(s)
}

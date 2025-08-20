// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoSyncRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateRepoSyncRuleRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *CreateRepoSyncRuleRequest
	GetNamespaceName() *string
	SetRepoName(v string) *CreateRepoSyncRuleRequest
	GetRepoName() *string
	SetRepoNameFilter(v string) *CreateRepoSyncRuleRequest
	GetRepoNameFilter() *string
	SetSyncRuleName(v string) *CreateRepoSyncRuleRequest
	GetSyncRuleName() *string
	SetSyncScope(v string) *CreateRepoSyncRuleRequest
	GetSyncScope() *string
	SetSyncTrigger(v string) *CreateRepoSyncRuleRequest
	GetSyncTrigger() *string
	SetTagFilter(v string) *CreateRepoSyncRuleRequest
	GetTagFilter() *string
	SetTargetInstanceId(v string) *CreateRepoSyncRuleRequest
	GetTargetInstanceId() *string
	SetTargetNamespaceName(v string) *CreateRepoSyncRuleRequest
	GetTargetNamespaceName() *string
	SetTargetRegionId(v string) *CreateRepoSyncRuleRequest
	GetTargetRegionId() *string
	SetTargetRepoName(v string) *CreateRepoSyncRuleRequest
	GetTargetRepoName() *string
	SetTargetUserId(v string) *CreateRepoSyncRuleRequest
	GetTargetUserId() *string
}

type CreateRepoSyncRuleRequest struct {
	// The source instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-hpdfkc6utbaq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace name of the source instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The name of the image repository in the source instance.
	//
	// example:
	//
	// repo1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The regular expression that is used to filter repositories.
	//
	// >  This parameter is valid only when SyncScope is set to `NAMESPACE`.
	//
	// example:
	//
	// .*
	RepoNameFilter *string `json:"RepoNameFilter,omitempty" xml:"RepoNameFilter,omitempty"`
	// The name of the image synchronization rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// rule
	SyncRuleName *string `json:"SyncRuleName,omitempty" xml:"SyncRuleName,omitempty"`
	// The synchronization scope. Valid values:
	//
	// 	- `REPO`: synchronizes the image tags in an image repository that meet the synchronization rule.
	//
	// 	- `NAMESPACE`: synchronizes the image tags in a namespace that meet the synchronization rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// REPO
	SyncScope *string `json:"SyncScope,omitempty" xml:"SyncScope,omitempty"`
	// The mode of triggering the synchronization rule. Valid values:
	//
	// 	- `INITIATIVE`: manually triggers the synchronization rule.
	//
	// 	- `PASSIVE`: automatically triggers the synchronization rule.
	//
	// example:
	//
	// PASSIVE
	SyncTrigger *string `json:"SyncTrigger,omitempty" xml:"SyncTrigger,omitempty"`
	// The regular expression that is used to filter image tags.
	//
	// This parameter is required.
	//
	// example:
	//
	// .*
	TagFilter *string `json:"TagFilter,omitempty" xml:"TagFilter,omitempty"`
	// The destination instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-ibxs3piklys3****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The namespace name of the destination instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	TargetNamespaceName *string `json:"TargetNamespaceName,omitempty" xml:"TargetNamespaceName,omitempty"`
	// The region ID of the destination instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
	// The name of the image repository in the destination instance.
	//
	// example:
	//
	// repo1
	TargetRepoName *string `json:"TargetRepoName,omitempty" xml:"TargetRepoName,omitempty"`
	// The user ID (UID) of the account to which the destination instance belongs.
	//
	// >  If you synchronize images across accounts, you must use the UID.
	//
	// example:
	//
	// 12645940***
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s CreateRepoSyncRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoSyncRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRepoSyncRuleRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *CreateRepoSyncRuleRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *CreateRepoSyncRuleRequest) GetRepoNameFilter() *string {
	return s.RepoNameFilter
}

func (s *CreateRepoSyncRuleRequest) GetSyncRuleName() *string {
	return s.SyncRuleName
}

func (s *CreateRepoSyncRuleRequest) GetSyncScope() *string {
	return s.SyncScope
}

func (s *CreateRepoSyncRuleRequest) GetSyncTrigger() *string {
	return s.SyncTrigger
}

func (s *CreateRepoSyncRuleRequest) GetTagFilter() *string {
	return s.TagFilter
}

func (s *CreateRepoSyncRuleRequest) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *CreateRepoSyncRuleRequest) GetTargetNamespaceName() *string {
	return s.TargetNamespaceName
}

func (s *CreateRepoSyncRuleRequest) GetTargetRegionId() *string {
	return s.TargetRegionId
}

func (s *CreateRepoSyncRuleRequest) GetTargetRepoName() *string {
	return s.TargetRepoName
}

func (s *CreateRepoSyncRuleRequest) GetTargetUserId() *string {
	return s.TargetUserId
}

func (s *CreateRepoSyncRuleRequest) SetInstanceId(v string) *CreateRepoSyncRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetNamespaceName(v string) *CreateRepoSyncRuleRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetRepoName(v string) *CreateRepoSyncRuleRequest {
	s.RepoName = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetRepoNameFilter(v string) *CreateRepoSyncRuleRequest {
	s.RepoNameFilter = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetSyncRuleName(v string) *CreateRepoSyncRuleRequest {
	s.SyncRuleName = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetSyncScope(v string) *CreateRepoSyncRuleRequest {
	s.SyncScope = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetSyncTrigger(v string) *CreateRepoSyncRuleRequest {
	s.SyncTrigger = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetTagFilter(v string) *CreateRepoSyncRuleRequest {
	s.TagFilter = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetTargetInstanceId(v string) *CreateRepoSyncRuleRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetTargetNamespaceName(v string) *CreateRepoSyncRuleRequest {
	s.TargetNamespaceName = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetTargetRegionId(v string) *CreateRepoSyncRuleRequest {
	s.TargetRegionId = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetTargetRepoName(v string) *CreateRepoSyncRuleRequest {
	s.TargetRepoName = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetTargetUserId(v string) *CreateRepoSyncRuleRequest {
	s.TargetUserId = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) Validate() error {
	return dara.Validate(s)
}

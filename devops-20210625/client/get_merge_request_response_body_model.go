// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMergeRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetMergeRequestResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMergeRequestResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetMergeRequestResponseBody
	GetRequestId() *string
	SetResult(v *GetMergeRequestResponseBodyResult) *GetMergeRequestResponseBody
	GetResult() *GetMergeRequestResponseBodyResult
	SetSuccess(v bool) *GetMergeRequestResponseBody
	GetSuccess() *bool
}

type GetMergeRequestResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// F7B85D1B-D1C2-140F-A039-341859F130B9
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetMergeRequestResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMergeRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMergeRequestResponseBody) GoString() string {
	return s.String()
}

func (s *GetMergeRequestResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMergeRequestResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMergeRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMergeRequestResponseBody) GetResult() *GetMergeRequestResponseBodyResult {
	return s.Result
}

func (s *GetMergeRequestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMergeRequestResponseBody) SetErrorCode(v string) *GetMergeRequestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMergeRequestResponseBody) SetErrorMessage(v string) *GetMergeRequestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMergeRequestResponseBody) SetRequestId(v string) *GetMergeRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMergeRequestResponseBody) SetResult(v *GetMergeRequestResponseBodyResult) *GetMergeRequestResponseBody {
	s.Result = v
	return s
}

func (s *GetMergeRequestResponseBody) SetSuccess(v bool) *GetMergeRequestResponseBody {
	s.Success = &v
	return s
}

func (s *GetMergeRequestResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMergeRequestResponseBodyResult struct {
	// example:
	//
	// 1
	Ahead *int32 `json:"ahead,omitempty" xml:"ahead,omitempty"`
	// example:
	//
	// true
	AllRequirementsPass *bool                                    `json:"allRequirementsPass,omitempty" xml:"allRequirementsPass,omitempty"`
	Author              *GetMergeRequestResponseBodyResultAuthor `json:"author,omitempty" xml:"author,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Behind *int32 `json:"behind,omitempty" xml:"behind,omitempty"`
	// example:
	//
	// WEB
	CreateFrom *string `json:"createFrom,omitempty" xml:"createFrom,omitempty"`
	// example:
	//
	// 2023-05-30T02:53:36Z
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// xxx
	DetailUrl *string `json:"detailUrl,omitempty" xml:"detailUrl,omitempty"`
	// example:
	//
	// 1
	LocalId        *int64  `json:"localId,omitempty" xml:"localId,omitempty"`
	MergedRevision *string `json:"mergedRevision,omitempty" xml:"mergedRevision,omitempty"`
	// example:
	//
	// bca90244c4b749e0b109df52ac0eb570
	MrBizId *string `json:"mrBizId,omitempty" xml:"mrBizId,omitempty"`
	// example:
	//
	// CODE_REVIEW
	MrType *string `json:"mrType,omitempty" xml:"mrType,omitempty"`
	// example:
	//
	// 2369234
	ProjectId *int64                                        `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Reviewers []*GetMergeRequestResponseBodyResultReviewers `json:"reviewers,omitempty" xml:"reviewers,omitempty" type:"Repeated"`
	// example:
	//
	// test-merge-request
	SourceBranch *string `json:"sourceBranch,omitempty" xml:"sourceBranch,omitempty"`
	// example:
	//
	// 2369234
	SourceProjectId *int64 `json:"sourceProjectId,omitempty" xml:"sourceProjectId,omitempty"`
	// example:
	//
	// UNDER_REVIEW
	Status      *string                                         `json:"status,omitempty" xml:"status,omitempty"`
	Subscribers []*GetMergeRequestResponseBodyResultSubscribers `json:"subscribers,omitempty" xml:"subscribers,omitempty" type:"Repeated"`
	// example:
	//
	// true
	SupportMergeFastForwardOnly *bool `json:"supportMergeFastForwardOnly,omitempty" xml:"supportMergeFastForwardOnly,omitempty"`
	// example:
	//
	// master
	TargetBranch *string `json:"targetBranch,omitempty" xml:"targetBranch,omitempty"`
	// example:
	//
	// 2369234
	TargetProjectId                *int64  `json:"targetProjectId,omitempty" xml:"targetProjectId,omitempty"`
	TargetProjectNameWithNamespace *string `json:"targetProjectNameWithNamespace,omitempty" xml:"targetProjectNameWithNamespace,omitempty"`
	// example:
	//
	// orgId/test-group/test-target-repo
	TargetProjectPathWithNamespace *string                                    `json:"targetProjectPathWithNamespace,omitempty" xml:"targetProjectPathWithNamespace,omitempty"`
	Title                          *string                                    `json:"title,omitempty" xml:"title,omitempty"`
	TodoList                       *GetMergeRequestResponseBodyResultTodoList `json:"todoList,omitempty" xml:"todoList,omitempty" type:"Struct"`
	// example:
	//
	// 2023-05-30T02:53:36Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// xxx
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s GetMergeRequestResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetMergeRequestResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMergeRequestResponseBodyResult) GetAhead() *int32 {
	return s.Ahead
}

func (s *GetMergeRequestResponseBodyResult) GetAllRequirementsPass() *bool {
	return s.AllRequirementsPass
}

func (s *GetMergeRequestResponseBodyResult) GetAuthor() *GetMergeRequestResponseBodyResultAuthor {
	return s.Author
}

func (s *GetMergeRequestResponseBodyResult) GetBehind() *int32 {
	return s.Behind
}

func (s *GetMergeRequestResponseBodyResult) GetCreateFrom() *string {
	return s.CreateFrom
}

func (s *GetMergeRequestResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMergeRequestResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *GetMergeRequestResponseBodyResult) GetDetailUrl() *string {
	return s.DetailUrl
}

func (s *GetMergeRequestResponseBodyResult) GetLocalId() *int64 {
	return s.LocalId
}

func (s *GetMergeRequestResponseBodyResult) GetMergedRevision() *string {
	return s.MergedRevision
}

func (s *GetMergeRequestResponseBodyResult) GetMrBizId() *string {
	return s.MrBizId
}

func (s *GetMergeRequestResponseBodyResult) GetMrType() *string {
	return s.MrType
}

func (s *GetMergeRequestResponseBodyResult) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetMergeRequestResponseBodyResult) GetReviewers() []*GetMergeRequestResponseBodyResultReviewers {
	return s.Reviewers
}

func (s *GetMergeRequestResponseBodyResult) GetSourceBranch() *string {
	return s.SourceBranch
}

func (s *GetMergeRequestResponseBodyResult) GetSourceProjectId() *int64 {
	return s.SourceProjectId
}

func (s *GetMergeRequestResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetMergeRequestResponseBodyResult) GetSubscribers() []*GetMergeRequestResponseBodyResultSubscribers {
	return s.Subscribers
}

func (s *GetMergeRequestResponseBodyResult) GetSupportMergeFastForwardOnly() *bool {
	return s.SupportMergeFastForwardOnly
}

func (s *GetMergeRequestResponseBodyResult) GetTargetBranch() *string {
	return s.TargetBranch
}

func (s *GetMergeRequestResponseBodyResult) GetTargetProjectId() *int64 {
	return s.TargetProjectId
}

func (s *GetMergeRequestResponseBodyResult) GetTargetProjectNameWithNamespace() *string {
	return s.TargetProjectNameWithNamespace
}

func (s *GetMergeRequestResponseBodyResult) GetTargetProjectPathWithNamespace() *string {
	return s.TargetProjectPathWithNamespace
}

func (s *GetMergeRequestResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *GetMergeRequestResponseBodyResult) GetTodoList() *GetMergeRequestResponseBodyResultTodoList {
	return s.TodoList
}

func (s *GetMergeRequestResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetMergeRequestResponseBodyResult) GetWebUrl() *string {
	return s.WebUrl
}

func (s *GetMergeRequestResponseBodyResult) SetAhead(v int32) *GetMergeRequestResponseBodyResult {
	s.Ahead = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetAllRequirementsPass(v bool) *GetMergeRequestResponseBodyResult {
	s.AllRequirementsPass = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetAuthor(v *GetMergeRequestResponseBodyResultAuthor) *GetMergeRequestResponseBodyResult {
	s.Author = v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetBehind(v int32) *GetMergeRequestResponseBodyResult {
	s.Behind = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetCreateFrom(v string) *GetMergeRequestResponseBodyResult {
	s.CreateFrom = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetCreateTime(v string) *GetMergeRequestResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetDescription(v string) *GetMergeRequestResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetDetailUrl(v string) *GetMergeRequestResponseBodyResult {
	s.DetailUrl = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetLocalId(v int64) *GetMergeRequestResponseBodyResult {
	s.LocalId = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetMergedRevision(v string) *GetMergeRequestResponseBodyResult {
	s.MergedRevision = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetMrBizId(v string) *GetMergeRequestResponseBodyResult {
	s.MrBizId = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetMrType(v string) *GetMergeRequestResponseBodyResult {
	s.MrType = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetProjectId(v int64) *GetMergeRequestResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetReviewers(v []*GetMergeRequestResponseBodyResultReviewers) *GetMergeRequestResponseBodyResult {
	s.Reviewers = v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetSourceBranch(v string) *GetMergeRequestResponseBodyResult {
	s.SourceBranch = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetSourceProjectId(v int64) *GetMergeRequestResponseBodyResult {
	s.SourceProjectId = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetStatus(v string) *GetMergeRequestResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetSubscribers(v []*GetMergeRequestResponseBodyResultSubscribers) *GetMergeRequestResponseBodyResult {
	s.Subscribers = v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetSupportMergeFastForwardOnly(v bool) *GetMergeRequestResponseBodyResult {
	s.SupportMergeFastForwardOnly = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetTargetBranch(v string) *GetMergeRequestResponseBodyResult {
	s.TargetBranch = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetTargetProjectId(v int64) *GetMergeRequestResponseBodyResult {
	s.TargetProjectId = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetTargetProjectNameWithNamespace(v string) *GetMergeRequestResponseBodyResult {
	s.TargetProjectNameWithNamespace = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetTargetProjectPathWithNamespace(v string) *GetMergeRequestResponseBodyResult {
	s.TargetProjectPathWithNamespace = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetTitle(v string) *GetMergeRequestResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetTodoList(v *GetMergeRequestResponseBodyResultTodoList) *GetMergeRequestResponseBodyResult {
	s.TodoList = v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetUpdateTime(v string) *GetMergeRequestResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) SetWebUrl(v string) *GetMergeRequestResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *GetMergeRequestResponseBodyResult) Validate() error {
	if s.Author != nil {
		if err := s.Author.Validate(); err != nil {
			return err
		}
	}
	if s.Reviewers != nil {
		for _, item := range s.Reviewers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Subscribers != nil {
		for _, item := range s.Subscribers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TodoList != nil {
		if err := s.TodoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMergeRequestResponseBodyResultAuthor struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 19927
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// test-codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// root-test-codeup
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetMergeRequestResponseBodyResultAuthor) String() string {
	return dara.Prettify(s)
}

func (s GetMergeRequestResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *GetMergeRequestResponseBodyResultAuthor) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetMergeRequestResponseBodyResultAuthor) GetEmail() *string {
	return s.Email
}

func (s *GetMergeRequestResponseBodyResultAuthor) GetId() *int64 {
	return s.Id
}

func (s *GetMergeRequestResponseBodyResultAuthor) GetName() *string {
	return s.Name
}

func (s *GetMergeRequestResponseBodyResultAuthor) GetState() *string {
	return s.State
}

func (s *GetMergeRequestResponseBodyResultAuthor) GetUsername() *string {
	return s.Username
}

func (s *GetMergeRequestResponseBodyResultAuthor) SetAvatarUrl(v string) *GetMergeRequestResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultAuthor) SetEmail(v string) *GetMergeRequestResponseBodyResultAuthor {
	s.Email = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultAuthor) SetId(v int64) *GetMergeRequestResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultAuthor) SetName(v string) *GetMergeRequestResponseBodyResultAuthor {
	s.Name = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultAuthor) SetState(v string) *GetMergeRequestResponseBodyResultAuthor {
	s.State = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultAuthor) SetUsername(v string) *GetMergeRequestResponseBodyResultAuthor {
	s.Username = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultAuthor) Validate() error {
	return dara.Validate(s)
}

type GetMergeRequestResponseBodyResultReviewers struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// username@example.com
	Email        *string `json:"email,omitempty" xml:"email,omitempty"`
	HasCommented *bool   `json:"hasCommented,omitempty" xml:"hasCommented,omitempty"`
	// example:
	//
	// false
	HasReviewed *bool `json:"hasReviewed,omitempty" xml:"hasReviewed,omitempty"`
	// example:
	//
	// 90452
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// test-codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// NOT_PASS
	ReviewOpinionStatus *string `json:"reviewOpinionStatus,omitempty" xml:"reviewOpinionStatus,omitempty"`
	ReviewTime          *string `json:"reviewTime,omitempty" xml:"reviewTime,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// root-test-codeup
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetMergeRequestResponseBodyResultReviewers) String() string {
	return dara.Prettify(s)
}

func (s GetMergeRequestResponseBodyResultReviewers) GoString() string {
	return s.String()
}

func (s *GetMergeRequestResponseBodyResultReviewers) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetMergeRequestResponseBodyResultReviewers) GetEmail() *string {
	return s.Email
}

func (s *GetMergeRequestResponseBodyResultReviewers) GetHasCommented() *bool {
	return s.HasCommented
}

func (s *GetMergeRequestResponseBodyResultReviewers) GetHasReviewed() *bool {
	return s.HasReviewed
}

func (s *GetMergeRequestResponseBodyResultReviewers) GetId() *int64 {
	return s.Id
}

func (s *GetMergeRequestResponseBodyResultReviewers) GetName() *string {
	return s.Name
}

func (s *GetMergeRequestResponseBodyResultReviewers) GetReviewOpinionStatus() *string {
	return s.ReviewOpinionStatus
}

func (s *GetMergeRequestResponseBodyResultReviewers) GetReviewTime() *string {
	return s.ReviewTime
}

func (s *GetMergeRequestResponseBodyResultReviewers) GetState() *string {
	return s.State
}

func (s *GetMergeRequestResponseBodyResultReviewers) GetUsername() *string {
	return s.Username
}

func (s *GetMergeRequestResponseBodyResultReviewers) SetAvatarUrl(v string) *GetMergeRequestResponseBodyResultReviewers {
	s.AvatarUrl = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultReviewers) SetEmail(v string) *GetMergeRequestResponseBodyResultReviewers {
	s.Email = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultReviewers) SetHasCommented(v bool) *GetMergeRequestResponseBodyResultReviewers {
	s.HasCommented = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultReviewers) SetHasReviewed(v bool) *GetMergeRequestResponseBodyResultReviewers {
	s.HasReviewed = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultReviewers) SetId(v int64) *GetMergeRequestResponseBodyResultReviewers {
	s.Id = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultReviewers) SetName(v string) *GetMergeRequestResponseBodyResultReviewers {
	s.Name = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultReviewers) SetReviewOpinionStatus(v string) *GetMergeRequestResponseBodyResultReviewers {
	s.ReviewOpinionStatus = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultReviewers) SetReviewTime(v string) *GetMergeRequestResponseBodyResultReviewers {
	s.ReviewTime = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultReviewers) SetState(v string) *GetMergeRequestResponseBodyResultReviewers {
	s.State = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultReviewers) SetUsername(v string) *GetMergeRequestResponseBodyResultReviewers {
	s.Username = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultReviewers) Validate() error {
	return dara.Validate(s)
}

type GetMergeRequestResponseBodyResultSubscribers struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 90452
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// test-subscriber
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// test-subscriber
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetMergeRequestResponseBodyResultSubscribers) String() string {
	return dara.Prettify(s)
}

func (s GetMergeRequestResponseBodyResultSubscribers) GoString() string {
	return s.String()
}

func (s *GetMergeRequestResponseBodyResultSubscribers) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetMergeRequestResponseBodyResultSubscribers) GetEmail() *string {
	return s.Email
}

func (s *GetMergeRequestResponseBodyResultSubscribers) GetId() *int64 {
	return s.Id
}

func (s *GetMergeRequestResponseBodyResultSubscribers) GetName() *string {
	return s.Name
}

func (s *GetMergeRequestResponseBodyResultSubscribers) GetState() *string {
	return s.State
}

func (s *GetMergeRequestResponseBodyResultSubscribers) GetUsername() *string {
	return s.Username
}

func (s *GetMergeRequestResponseBodyResultSubscribers) SetAvatarUrl(v string) *GetMergeRequestResponseBodyResultSubscribers {
	s.AvatarUrl = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultSubscribers) SetEmail(v string) *GetMergeRequestResponseBodyResultSubscribers {
	s.Email = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultSubscribers) SetId(v int64) *GetMergeRequestResponseBodyResultSubscribers {
	s.Id = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultSubscribers) SetName(v string) *GetMergeRequestResponseBodyResultSubscribers {
	s.Name = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultSubscribers) SetState(v string) *GetMergeRequestResponseBodyResultSubscribers {
	s.State = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultSubscribers) SetUsername(v string) *GetMergeRequestResponseBodyResultSubscribers {
	s.Username = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultSubscribers) Validate() error {
	return dara.Validate(s)
}

type GetMergeRequestResponseBodyResultTodoList struct {
	RequirementCheckItems []*GetMergeRequestResponseBodyResultTodoListRequirementCheckItems `json:"requirementCheckItems,omitempty" xml:"requirementCheckItems,omitempty" type:"Repeated"`
}

func (s GetMergeRequestResponseBodyResultTodoList) String() string {
	return dara.Prettify(s)
}

func (s GetMergeRequestResponseBodyResultTodoList) GoString() string {
	return s.String()
}

func (s *GetMergeRequestResponseBodyResultTodoList) GetRequirementCheckItems() []*GetMergeRequestResponseBodyResultTodoListRequirementCheckItems {
	return s.RequirementCheckItems
}

func (s *GetMergeRequestResponseBodyResultTodoList) SetRequirementCheckItems(v []*GetMergeRequestResponseBodyResultTodoListRequirementCheckItems) *GetMergeRequestResponseBodyResultTodoList {
	s.RequirementCheckItems = v
	return s
}

func (s *GetMergeRequestResponseBodyResultTodoList) Validate() error {
	if s.RequirementCheckItems != nil {
		for _, item := range s.RequirementCheckItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMergeRequestResponseBodyResultTodoListRequirementCheckItems struct {
	// example:
	//
	// COMMENTS_CHECK
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// example:
	//
	// true
	Pass *bool `json:"pass,omitempty" xml:"pass,omitempty"`
}

func (s GetMergeRequestResponseBodyResultTodoListRequirementCheckItems) String() string {
	return dara.Prettify(s)
}

func (s GetMergeRequestResponseBodyResultTodoListRequirementCheckItems) GoString() string {
	return s.String()
}

func (s *GetMergeRequestResponseBodyResultTodoListRequirementCheckItems) GetItemType() *string {
	return s.ItemType
}

func (s *GetMergeRequestResponseBodyResultTodoListRequirementCheckItems) GetPass() *bool {
	return s.Pass
}

func (s *GetMergeRequestResponseBodyResultTodoListRequirementCheckItems) SetItemType(v string) *GetMergeRequestResponseBodyResultTodoListRequirementCheckItems {
	s.ItemType = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultTodoListRequirementCheckItems) SetPass(v bool) *GetMergeRequestResponseBodyResultTodoListRequirementCheckItems {
	s.Pass = &v
	return s
}

func (s *GetMergeRequestResponseBodyResultTodoListRequirementCheckItems) Validate() error {
	return dara.Validate(s)
}

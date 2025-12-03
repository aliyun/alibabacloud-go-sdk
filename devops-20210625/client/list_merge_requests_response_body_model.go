// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMergeRequestsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListMergeRequestsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListMergeRequestsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListMergeRequestsResponseBody
	GetRequestId() *string
	SetResult(v []*ListMergeRequestsResponseBodyResult) *ListMergeRequestsResponseBody
	GetResult() []*ListMergeRequestsResponseBodyResult
	SetSuccess(v bool) *ListMergeRequestsResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListMergeRequestsResponseBody
	GetTotal() *int64
}

type ListMergeRequestsResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListMergeRequestsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 30
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMergeRequestsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMergeRequestsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListMergeRequestsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMergeRequestsResponseBody) GetResult() []*ListMergeRequestsResponseBodyResult {
	return s.Result
}

func (s *ListMergeRequestsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMergeRequestsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListMergeRequestsResponseBody) SetErrorCode(v string) *ListMergeRequestsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMergeRequestsResponseBody) SetErrorMessage(v string) *ListMergeRequestsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMergeRequestsResponseBody) SetRequestId(v string) *ListMergeRequestsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMergeRequestsResponseBody) SetResult(v []*ListMergeRequestsResponseBodyResult) *ListMergeRequestsResponseBody {
	s.Result = v
	return s
}

func (s *ListMergeRequestsResponseBody) SetSuccess(v bool) *ListMergeRequestsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMergeRequestsResponseBody) SetTotal(v int64) *ListMergeRequestsResponseBody {
	s.Total = &v
	return s
}

func (s *ListMergeRequestsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMergeRequestsResponseBodyResult struct {
	Author *ListMergeRequestsResponseBodyResultAuthor `json:"author,omitempty" xml:"author,omitempty" type:"Struct"`
	// example:
	//
	// 2022-03-18 14:24:54
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// WEB
	CreationMethod *string `json:"creationMethod,omitempty" xml:"creationMethod,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// xxx
	DetailUrl *string `json:"detailUrl,omitempty" xml:"detailUrl,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1
	Iid    *int64                                       `json:"iid,omitempty" xml:"iid,omitempty"`
	Labels []*ListMergeRequestsResponseBodyResultLabels `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	LocalId        *int64  `json:"localId,omitempty" xml:"localId,omitempty"`
	MergedRevision *string `json:"mergedRevision,omitempty" xml:"mergedRevision,omitempty"`
	// example:
	//
	// bca90244c4b749e0b109df52ac0eb570
	MrBizId           *string `json:"mrBizId,omitempty" xml:"mrBizId,omitempty"`
	NameWithNamespace *string `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	// example:
	//
	// true
	NewMergeRequestIdentifier *bool `json:"newMergeRequestIdentifier,omitempty" xml:"newMergeRequestIdentifier,omitempty"`
	// example:
	//
	// UNDER_REVIEW
	NewVersionState *string `json:"newVersionState,omitempty" xml:"newVersionState,omitempty"`
	// example:
	//
	// 2369234
	ProjectId *int64                                          `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Reviewers []*ListMergeRequestsResponseBodyResultReviewers `json:"reviewers,omitempty" xml:"reviewers,omitempty" type:"Repeated"`
	// example:
	//
	// test-merge-source-branch
	SourceBranch *string `json:"sourceBranch,omitempty" xml:"sourceBranch,omitempty"`
	// example:
	//
	// 2876119
	SourceProjectId *int64 `json:"sourceProjectId,omitempty" xml:"sourceProjectId,omitempty"`
	// example:
	//
	// BRANCH
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// example:
	//
	// git@xxx:xxx/test/test.git
	SshUrl *string `json:"sshUrl,omitempty" xml:"sshUrl,omitempty"`
	// example:
	//
	// opened
	State       *string                                           `json:"state,omitempty" xml:"state,omitempty"`
	Subscribers []*ListMergeRequestsResponseBodyResultSubscribers `json:"subscribers,omitempty" xml:"subscribers,omitempty" type:"Repeated"`
	// example:
	//
	// true
	SupportMergeFFOnly *bool `json:"supportMergeFFOnly,omitempty" xml:"supportMergeFFOnly,omitempty"`
	// example:
	//
	// test-merge-target-branch
	TargetBranch *string `json:"targetBranch,omitempty" xml:"targetBranch,omitempty"`
	// example:
	//
	// 2876119
	TargetProjectId *int64 `json:"targetProjectId,omitempty" xml:"targetProjectId,omitempty"`
	// example:
	//
	// BRANCH
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// ""
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
	// example:
	//
	// false
	WorkInProgress *bool `json:"workInProgress,omitempty" xml:"workInProgress,omitempty"`
}

func (s ListMergeRequestsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponseBodyResult) GetAuthor() *ListMergeRequestsResponseBodyResultAuthor {
	return s.Author
}

func (s *ListMergeRequestsResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListMergeRequestsResponseBodyResult) GetCreationMethod() *string {
	return s.CreationMethod
}

func (s *ListMergeRequestsResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListMergeRequestsResponseBodyResult) GetDetailUrl() *string {
	return s.DetailUrl
}

func (s *ListMergeRequestsResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListMergeRequestsResponseBodyResult) GetIid() *int64 {
	return s.Iid
}

func (s *ListMergeRequestsResponseBodyResult) GetLabels() []*ListMergeRequestsResponseBodyResultLabels {
	return s.Labels
}

func (s *ListMergeRequestsResponseBodyResult) GetLocalId() *int64 {
	return s.LocalId
}

func (s *ListMergeRequestsResponseBodyResult) GetMergedRevision() *string {
	return s.MergedRevision
}

func (s *ListMergeRequestsResponseBodyResult) GetMrBizId() *string {
	return s.MrBizId
}

func (s *ListMergeRequestsResponseBodyResult) GetNameWithNamespace() *string {
	return s.NameWithNamespace
}

func (s *ListMergeRequestsResponseBodyResult) GetNewMergeRequestIdentifier() *bool {
	return s.NewMergeRequestIdentifier
}

func (s *ListMergeRequestsResponseBodyResult) GetNewVersionState() *string {
	return s.NewVersionState
}

func (s *ListMergeRequestsResponseBodyResult) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListMergeRequestsResponseBodyResult) GetReviewers() []*ListMergeRequestsResponseBodyResultReviewers {
	return s.Reviewers
}

func (s *ListMergeRequestsResponseBodyResult) GetSourceBranch() *string {
	return s.SourceBranch
}

func (s *ListMergeRequestsResponseBodyResult) GetSourceProjectId() *int64 {
	return s.SourceProjectId
}

func (s *ListMergeRequestsResponseBodyResult) GetSourceType() *string {
	return s.SourceType
}

func (s *ListMergeRequestsResponseBodyResult) GetSshUrl() *string {
	return s.SshUrl
}

func (s *ListMergeRequestsResponseBodyResult) GetState() *string {
	return s.State
}

func (s *ListMergeRequestsResponseBodyResult) GetSubscribers() []*ListMergeRequestsResponseBodyResultSubscribers {
	return s.Subscribers
}

func (s *ListMergeRequestsResponseBodyResult) GetSupportMergeFFOnly() *bool {
	return s.SupportMergeFFOnly
}

func (s *ListMergeRequestsResponseBodyResult) GetTargetBranch() *string {
	return s.TargetBranch
}

func (s *ListMergeRequestsResponseBodyResult) GetTargetProjectId() *int64 {
	return s.TargetProjectId
}

func (s *ListMergeRequestsResponseBodyResult) GetTargetType() *string {
	return s.TargetType
}

func (s *ListMergeRequestsResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *ListMergeRequestsResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListMergeRequestsResponseBodyResult) GetWebUrl() *string {
	return s.WebUrl
}

func (s *ListMergeRequestsResponseBodyResult) GetWorkInProgress() *bool {
	return s.WorkInProgress
}

func (s *ListMergeRequestsResponseBodyResult) SetAuthor(v *ListMergeRequestsResponseBodyResultAuthor) *ListMergeRequestsResponseBodyResult {
	s.Author = v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetCreatedAt(v string) *ListMergeRequestsResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetCreationMethod(v string) *ListMergeRequestsResponseBodyResult {
	s.CreationMethod = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetDescription(v string) *ListMergeRequestsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetDetailUrl(v string) *ListMergeRequestsResponseBodyResult {
	s.DetailUrl = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetId(v int64) *ListMergeRequestsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetIid(v int64) *ListMergeRequestsResponseBodyResult {
	s.Iid = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetLabels(v []*ListMergeRequestsResponseBodyResultLabels) *ListMergeRequestsResponseBodyResult {
	s.Labels = v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetLocalId(v int64) *ListMergeRequestsResponseBodyResult {
	s.LocalId = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetMergedRevision(v string) *ListMergeRequestsResponseBodyResult {
	s.MergedRevision = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetMrBizId(v string) *ListMergeRequestsResponseBodyResult {
	s.MrBizId = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetNameWithNamespace(v string) *ListMergeRequestsResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetNewMergeRequestIdentifier(v bool) *ListMergeRequestsResponseBodyResult {
	s.NewMergeRequestIdentifier = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetNewVersionState(v string) *ListMergeRequestsResponseBodyResult {
	s.NewVersionState = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetProjectId(v int64) *ListMergeRequestsResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetReviewers(v []*ListMergeRequestsResponseBodyResultReviewers) *ListMergeRequestsResponseBodyResult {
	s.Reviewers = v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetSourceBranch(v string) *ListMergeRequestsResponseBodyResult {
	s.SourceBranch = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetSourceProjectId(v int64) *ListMergeRequestsResponseBodyResult {
	s.SourceProjectId = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetSourceType(v string) *ListMergeRequestsResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetSshUrl(v string) *ListMergeRequestsResponseBodyResult {
	s.SshUrl = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetState(v string) *ListMergeRequestsResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetSubscribers(v []*ListMergeRequestsResponseBodyResultSubscribers) *ListMergeRequestsResponseBodyResult {
	s.Subscribers = v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetSupportMergeFFOnly(v bool) *ListMergeRequestsResponseBodyResult {
	s.SupportMergeFFOnly = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetTargetBranch(v string) *ListMergeRequestsResponseBodyResult {
	s.TargetBranch = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetTargetProjectId(v int64) *ListMergeRequestsResponseBodyResult {
	s.TargetProjectId = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetTargetType(v string) *ListMergeRequestsResponseBodyResult {
	s.TargetType = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetTitle(v string) *ListMergeRequestsResponseBodyResult {
	s.Title = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetUpdatedAt(v string) *ListMergeRequestsResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetWebUrl(v string) *ListMergeRequestsResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetWorkInProgress(v bool) *ListMergeRequestsResponseBodyResult {
	s.WorkInProgress = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) Validate() error {
	if s.Author != nil {
		if err := s.Author.Validate(); err != nil {
			return err
		}
	}
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	return nil
}

type ListMergeRequestsResponseBodyResultAuthor struct {
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

func (s ListMergeRequestsResponseBodyResultAuthor) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestsResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponseBodyResultAuthor) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListMergeRequestsResponseBodyResultAuthor) GetEmail() *string {
	return s.Email
}

func (s *ListMergeRequestsResponseBodyResultAuthor) GetId() *int64 {
	return s.Id
}

func (s *ListMergeRequestsResponseBodyResultAuthor) GetName() *string {
	return s.Name
}

func (s *ListMergeRequestsResponseBodyResultAuthor) GetState() *string {
	return s.State
}

func (s *ListMergeRequestsResponseBodyResultAuthor) GetUsername() *string {
	return s.Username
}

func (s *ListMergeRequestsResponseBodyResultAuthor) SetAvatarUrl(v string) *ListMergeRequestsResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultAuthor) SetEmail(v string) *ListMergeRequestsResponseBodyResultAuthor {
	s.Email = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultAuthor) SetId(v int64) *ListMergeRequestsResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultAuthor) SetName(v string) *ListMergeRequestsResponseBodyResultAuthor {
	s.Name = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultAuthor) SetState(v string) *ListMergeRequestsResponseBodyResultAuthor {
	s.State = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultAuthor) SetUsername(v string) *ListMergeRequestsResponseBodyResultAuthor {
	s.Username = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultAuthor) Validate() error {
	return dara.Validate(s)
}

type ListMergeRequestsResponseBodyResultLabels struct {
	Color       *string `json:"color,omitempty" xml:"color,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListMergeRequestsResponseBodyResultLabels) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestsResponseBodyResultLabels) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponseBodyResultLabels) GetColor() *string {
	return s.Color
}

func (s *ListMergeRequestsResponseBodyResultLabels) GetDescription() *string {
	return s.Description
}

func (s *ListMergeRequestsResponseBodyResultLabels) GetId() *string {
	return s.Id
}

func (s *ListMergeRequestsResponseBodyResultLabels) GetName() *string {
	return s.Name
}

func (s *ListMergeRequestsResponseBodyResultLabels) SetColor(v string) *ListMergeRequestsResponseBodyResultLabels {
	s.Color = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultLabels) SetDescription(v string) *ListMergeRequestsResponseBodyResultLabels {
	s.Description = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultLabels) SetId(v string) *ListMergeRequestsResponseBodyResultLabels {
	s.Id = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultLabels) SetName(v string) *ListMergeRequestsResponseBodyResultLabels {
	s.Name = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultLabels) Validate() error {
	return dara.Validate(s)
}

type ListMergeRequestsResponseBodyResultReviewers struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// username@example.com
	Email        *string `json:"email,omitempty" xml:"email,omitempty"`
	HasCommented *bool   `json:"hasCommented,omitempty" xml:"hasCommented,omitempty"`
	HasReviewed  *bool   `json:"hasReviewed,omitempty" xml:"hasReviewed,omitempty"`
	// example:
	//
	// 43127
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// test-review-user
	Name                *string `json:"name,omitempty" xml:"name,omitempty"`
	ReviewOpinionStatus *string `json:"reviewOpinionStatus,omitempty" xml:"reviewOpinionStatus,omitempty"`
	ReviewTime          *string `json:"reviewTime,omitempty" xml:"reviewTime,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// approved
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// root-test-review-user
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListMergeRequestsResponseBodyResultReviewers) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestsResponseBodyResultReviewers) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponseBodyResultReviewers) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListMergeRequestsResponseBodyResultReviewers) GetEmail() *string {
	return s.Email
}

func (s *ListMergeRequestsResponseBodyResultReviewers) GetHasCommented() *bool {
	return s.HasCommented
}

func (s *ListMergeRequestsResponseBodyResultReviewers) GetHasReviewed() *bool {
	return s.HasReviewed
}

func (s *ListMergeRequestsResponseBodyResultReviewers) GetId() *int64 {
	return s.Id
}

func (s *ListMergeRequestsResponseBodyResultReviewers) GetName() *string {
	return s.Name
}

func (s *ListMergeRequestsResponseBodyResultReviewers) GetReviewOpinionStatus() *string {
	return s.ReviewOpinionStatus
}

func (s *ListMergeRequestsResponseBodyResultReviewers) GetReviewTime() *string {
	return s.ReviewTime
}

func (s *ListMergeRequestsResponseBodyResultReviewers) GetState() *string {
	return s.State
}

func (s *ListMergeRequestsResponseBodyResultReviewers) GetStatus() *string {
	return s.Status
}

func (s *ListMergeRequestsResponseBodyResultReviewers) GetUsername() *string {
	return s.Username
}

func (s *ListMergeRequestsResponseBodyResultReviewers) SetAvatarUrl(v string) *ListMergeRequestsResponseBodyResultReviewers {
	s.AvatarUrl = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultReviewers) SetEmail(v string) *ListMergeRequestsResponseBodyResultReviewers {
	s.Email = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultReviewers) SetHasCommented(v bool) *ListMergeRequestsResponseBodyResultReviewers {
	s.HasCommented = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultReviewers) SetHasReviewed(v bool) *ListMergeRequestsResponseBodyResultReviewers {
	s.HasReviewed = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultReviewers) SetId(v int64) *ListMergeRequestsResponseBodyResultReviewers {
	s.Id = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultReviewers) SetName(v string) *ListMergeRequestsResponseBodyResultReviewers {
	s.Name = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultReviewers) SetReviewOpinionStatus(v string) *ListMergeRequestsResponseBodyResultReviewers {
	s.ReviewOpinionStatus = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultReviewers) SetReviewTime(v string) *ListMergeRequestsResponseBodyResultReviewers {
	s.ReviewTime = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultReviewers) SetState(v string) *ListMergeRequestsResponseBodyResultReviewers {
	s.State = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultReviewers) SetStatus(v string) *ListMergeRequestsResponseBodyResultReviewers {
	s.Status = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultReviewers) SetUsername(v string) *ListMergeRequestsResponseBodyResultReviewers {
	s.Username = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultReviewers) Validate() error {
	return dara.Validate(s)
}

type ListMergeRequestsResponseBodyResultSubscribers struct {
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
	// 1876119
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
	// root-test-subscriber
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListMergeRequestsResponseBodyResultSubscribers) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestsResponseBodyResultSubscribers) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponseBodyResultSubscribers) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListMergeRequestsResponseBodyResultSubscribers) GetEmail() *string {
	return s.Email
}

func (s *ListMergeRequestsResponseBodyResultSubscribers) GetId() *int64 {
	return s.Id
}

func (s *ListMergeRequestsResponseBodyResultSubscribers) GetName() *string {
	return s.Name
}

func (s *ListMergeRequestsResponseBodyResultSubscribers) GetState() *string {
	return s.State
}

func (s *ListMergeRequestsResponseBodyResultSubscribers) GetUsername() *string {
	return s.Username
}

func (s *ListMergeRequestsResponseBodyResultSubscribers) SetAvatarUrl(v string) *ListMergeRequestsResponseBodyResultSubscribers {
	s.AvatarUrl = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultSubscribers) SetEmail(v string) *ListMergeRequestsResponseBodyResultSubscribers {
	s.Email = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultSubscribers) SetId(v int64) *ListMergeRequestsResponseBodyResultSubscribers {
	s.Id = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultSubscribers) SetName(v string) *ListMergeRequestsResponseBodyResultSubscribers {
	s.Name = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultSubscribers) SetState(v string) *ListMergeRequestsResponseBodyResultSubscribers {
	s.State = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultSubscribers) SetUsername(v string) *ListMergeRequestsResponseBodyResultSubscribers {
	s.Username = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultSubscribers) Validate() error {
	return dara.Validate(s)
}

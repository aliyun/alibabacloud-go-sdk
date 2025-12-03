// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMergeRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateMergeRequestResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateMergeRequestResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateMergeRequestResponseBody
	GetRequestId() *string
	SetResult(v *CreateMergeRequestResponseBodyResult) *CreateMergeRequestResponseBody
	GetResult() *CreateMergeRequestResponseBodyResult
	SetSuccess(v bool) *CreateMergeRequestResponseBody
	GetSuccess() *bool
}

type CreateMergeRequestResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// F7B85D1B-D1C2-140F-A039-341859F130B9
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateMergeRequestResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateMergeRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMergeRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMergeRequestResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateMergeRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMergeRequestResponseBody) GetResult() *CreateMergeRequestResponseBodyResult {
	return s.Result
}

func (s *CreateMergeRequestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMergeRequestResponseBody) SetErrorCode(v string) *CreateMergeRequestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateMergeRequestResponseBody) SetErrorMessage(v string) *CreateMergeRequestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateMergeRequestResponseBody) SetRequestId(v string) *CreateMergeRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMergeRequestResponseBody) SetResult(v *CreateMergeRequestResponseBodyResult) *CreateMergeRequestResponseBody {
	s.Result = v
	return s
}

func (s *CreateMergeRequestResponseBody) SetSuccess(v bool) *CreateMergeRequestResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMergeRequestResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMergeRequestResponseBodyResult struct {
	// example:
	//
	// 2
	Ahead *int32 `json:"ahead,omitempty" xml:"ahead,omitempty"`
	// example:
	//
	// true
	AllRequirementsPass *bool                                       `json:"allRequirementsPass,omitempty" xml:"allRequirementsPass,omitempty"`
	Author              *CreateMergeRequestResponseBodyResultAuthor `json:"author,omitempty" xml:"author,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Behind *int32 `json:"behind,omitempty" xml:"behind,omitempty"`
	// example:
	//
	// WEB
	CreateFrom *string `json:"createFrom,omitempty" xml:"createFrom,omitempty"`
	// example:
	//
	// 2023-06-02T03:41:22Z
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// xxx
	DetailUrl *string `json:"detailUrl,omitempty" xml:"detailUrl,omitempty"`
	// example:
	//
	// 1
	LocalId *int64 `json:"localId,omitempty" xml:"localId,omitempty"`
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
	ProjectId *int64                                           `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Reviewers []*CreateMergeRequestResponseBodyResultReviewers `json:"reviewers,omitempty" xml:"reviewers,omitempty" type:"Repeated"`
	// example:
	//
	// sourceBranch
	SourceBranch *string `json:"sourceBranch,omitempty" xml:"sourceBranch,omitempty"`
	// example:
	//
	// 2369234
	SourceProjectId *int64 `json:"sourceProjectId,omitempty" xml:"sourceProjectId,omitempty"`
	// example:
	//
	// UNDER_REVIEW
	Status      *string                                            `json:"status,omitempty" xml:"status,omitempty"`
	Subscribers []*CreateMergeRequestResponseBodyResultSubscribers `json:"subscribers,omitempty" xml:"subscribers,omitempty" type:"Repeated"`
	// example:
	//
	// true
	SupportMergeFastForwardOnly *bool `json:"supportMergeFastForwardOnly,omitempty" xml:"supportMergeFastForwardOnly,omitempty"`
	// example:
	//
	// targetBranch
	TargetBranch *string `json:"targetBranch,omitempty" xml:"targetBranch,omitempty"`
	// example:
	//
	// 2369234
	TargetProjectId *int64  `json:"targetProjectId,omitempty" xml:"targetProjectId,omitempty"`
	Title           *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2023-06-02T03:41:22Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// https://codeup.aliyun.com/xxx/test/test
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s CreateMergeRequestResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResult) GetAhead() *int32 {
	return s.Ahead
}

func (s *CreateMergeRequestResponseBodyResult) GetAllRequirementsPass() *bool {
	return s.AllRequirementsPass
}

func (s *CreateMergeRequestResponseBodyResult) GetAuthor() *CreateMergeRequestResponseBodyResultAuthor {
	return s.Author
}

func (s *CreateMergeRequestResponseBodyResult) GetBehind() *int32 {
	return s.Behind
}

func (s *CreateMergeRequestResponseBodyResult) GetCreateFrom() *string {
	return s.CreateFrom
}

func (s *CreateMergeRequestResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateMergeRequestResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *CreateMergeRequestResponseBodyResult) GetDetailUrl() *string {
	return s.DetailUrl
}

func (s *CreateMergeRequestResponseBodyResult) GetLocalId() *int64 {
	return s.LocalId
}

func (s *CreateMergeRequestResponseBodyResult) GetMrBizId() *string {
	return s.MrBizId
}

func (s *CreateMergeRequestResponseBodyResult) GetMrType() *string {
	return s.MrType
}

func (s *CreateMergeRequestResponseBodyResult) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateMergeRequestResponseBodyResult) GetReviewers() []*CreateMergeRequestResponseBodyResultReviewers {
	return s.Reviewers
}

func (s *CreateMergeRequestResponseBodyResult) GetSourceBranch() *string {
	return s.SourceBranch
}

func (s *CreateMergeRequestResponseBodyResult) GetSourceProjectId() *int64 {
	return s.SourceProjectId
}

func (s *CreateMergeRequestResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *CreateMergeRequestResponseBodyResult) GetSubscribers() []*CreateMergeRequestResponseBodyResultSubscribers {
	return s.Subscribers
}

func (s *CreateMergeRequestResponseBodyResult) GetSupportMergeFastForwardOnly() *bool {
	return s.SupportMergeFastForwardOnly
}

func (s *CreateMergeRequestResponseBodyResult) GetTargetBranch() *string {
	return s.TargetBranch
}

func (s *CreateMergeRequestResponseBodyResult) GetTargetProjectId() *int64 {
	return s.TargetProjectId
}

func (s *CreateMergeRequestResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *CreateMergeRequestResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateMergeRequestResponseBodyResult) GetWebUrl() *string {
	return s.WebUrl
}

func (s *CreateMergeRequestResponseBodyResult) SetAhead(v int32) *CreateMergeRequestResponseBodyResult {
	s.Ahead = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetAllRequirementsPass(v bool) *CreateMergeRequestResponseBodyResult {
	s.AllRequirementsPass = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetAuthor(v *CreateMergeRequestResponseBodyResultAuthor) *CreateMergeRequestResponseBodyResult {
	s.Author = v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetBehind(v int32) *CreateMergeRequestResponseBodyResult {
	s.Behind = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetCreateFrom(v string) *CreateMergeRequestResponseBodyResult {
	s.CreateFrom = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetCreateTime(v string) *CreateMergeRequestResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetDescription(v string) *CreateMergeRequestResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetDetailUrl(v string) *CreateMergeRequestResponseBodyResult {
	s.DetailUrl = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetLocalId(v int64) *CreateMergeRequestResponseBodyResult {
	s.LocalId = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetMrBizId(v string) *CreateMergeRequestResponseBodyResult {
	s.MrBizId = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetMrType(v string) *CreateMergeRequestResponseBodyResult {
	s.MrType = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetProjectId(v int64) *CreateMergeRequestResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetReviewers(v []*CreateMergeRequestResponseBodyResultReviewers) *CreateMergeRequestResponseBodyResult {
	s.Reviewers = v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetSourceBranch(v string) *CreateMergeRequestResponseBodyResult {
	s.SourceBranch = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetSourceProjectId(v int64) *CreateMergeRequestResponseBodyResult {
	s.SourceProjectId = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetStatus(v string) *CreateMergeRequestResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetSubscribers(v []*CreateMergeRequestResponseBodyResultSubscribers) *CreateMergeRequestResponseBodyResult {
	s.Subscribers = v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetSupportMergeFastForwardOnly(v bool) *CreateMergeRequestResponseBodyResult {
	s.SupportMergeFastForwardOnly = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetTargetBranch(v string) *CreateMergeRequestResponseBodyResult {
	s.TargetBranch = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetTargetProjectId(v int64) *CreateMergeRequestResponseBodyResult {
	s.TargetProjectId = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetTitle(v string) *CreateMergeRequestResponseBodyResult {
	s.Title = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetUpdateTime(v string) *CreateMergeRequestResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetWebUrl(v string) *CreateMergeRequestResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) Validate() error {
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
	return nil
}

type CreateMergeRequestResponseBodyResultAuthor struct {
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
	// 19230
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
	// test-codeup-nickname
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s CreateMergeRequestResponseBodyResultAuthor) String() string {
	return dara.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultAuthor) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *CreateMergeRequestResponseBodyResultAuthor) GetEmail() *string {
	return s.Email
}

func (s *CreateMergeRequestResponseBodyResultAuthor) GetId() *int64 {
	return s.Id
}

func (s *CreateMergeRequestResponseBodyResultAuthor) GetName() *string {
	return s.Name
}

func (s *CreateMergeRequestResponseBodyResultAuthor) GetState() *string {
	return s.State
}

func (s *CreateMergeRequestResponseBodyResultAuthor) GetUsername() *string {
	return s.Username
}

func (s *CreateMergeRequestResponseBodyResultAuthor) SetAvatarUrl(v string) *CreateMergeRequestResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAuthor) SetEmail(v string) *CreateMergeRequestResponseBodyResultAuthor {
	s.Email = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAuthor) SetId(v int64) *CreateMergeRequestResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAuthor) SetName(v string) *CreateMergeRequestResponseBodyResultAuthor {
	s.Name = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAuthor) SetState(v string) *CreateMergeRequestResponseBodyResultAuthor {
	s.State = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAuthor) SetUsername(v string) *CreateMergeRequestResponseBodyResultAuthor {
	s.Username = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAuthor) Validate() error {
	return dara.Validate(s)
}

type CreateMergeRequestResponseBodyResultReviewers struct {
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
	// false
	HasReviewed *bool `json:"hasReviewed,omitempty" xml:"hasReviewed,omitempty"`
	// example:
	//
	// 7905
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// PASS
	ReviewOpinionStatus *string `json:"reviewOpinionStatus,omitempty" xml:"reviewOpinionStatus,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// root-codeup
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s CreateMergeRequestResponseBodyResultReviewers) String() string {
	return dara.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultReviewers) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultReviewers) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *CreateMergeRequestResponseBodyResultReviewers) GetEmail() *string {
	return s.Email
}

func (s *CreateMergeRequestResponseBodyResultReviewers) GetHasReviewed() *bool {
	return s.HasReviewed
}

func (s *CreateMergeRequestResponseBodyResultReviewers) GetId() *int64 {
	return s.Id
}

func (s *CreateMergeRequestResponseBodyResultReviewers) GetName() *string {
	return s.Name
}

func (s *CreateMergeRequestResponseBodyResultReviewers) GetReviewOpinionStatus() *string {
	return s.ReviewOpinionStatus
}

func (s *CreateMergeRequestResponseBodyResultReviewers) GetState() *string {
	return s.State
}

func (s *CreateMergeRequestResponseBodyResultReviewers) GetUsername() *string {
	return s.Username
}

func (s *CreateMergeRequestResponseBodyResultReviewers) SetAvatarUrl(v string) *CreateMergeRequestResponseBodyResultReviewers {
	s.AvatarUrl = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultReviewers) SetEmail(v string) *CreateMergeRequestResponseBodyResultReviewers {
	s.Email = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultReviewers) SetHasReviewed(v bool) *CreateMergeRequestResponseBodyResultReviewers {
	s.HasReviewed = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultReviewers) SetId(v int64) *CreateMergeRequestResponseBodyResultReviewers {
	s.Id = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultReviewers) SetName(v string) *CreateMergeRequestResponseBodyResultReviewers {
	s.Name = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultReviewers) SetReviewOpinionStatus(v string) *CreateMergeRequestResponseBodyResultReviewers {
	s.ReviewOpinionStatus = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultReviewers) SetState(v string) *CreateMergeRequestResponseBodyResultReviewers {
	s.State = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultReviewers) SetUsername(v string) *CreateMergeRequestResponseBodyResultReviewers {
	s.Username = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultReviewers) Validate() error {
	return dara.Validate(s)
}

type CreateMergeRequestResponseBodyResultSubscribers struct {
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
	// 10092
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

func (s CreateMergeRequestResponseBodyResultSubscribers) String() string {
	return dara.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultSubscribers) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultSubscribers) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *CreateMergeRequestResponseBodyResultSubscribers) GetEmail() *string {
	return s.Email
}

func (s *CreateMergeRequestResponseBodyResultSubscribers) GetId() *int64 {
	return s.Id
}

func (s *CreateMergeRequestResponseBodyResultSubscribers) GetName() *string {
	return s.Name
}

func (s *CreateMergeRequestResponseBodyResultSubscribers) GetState() *string {
	return s.State
}

func (s *CreateMergeRequestResponseBodyResultSubscribers) GetUsername() *string {
	return s.Username
}

func (s *CreateMergeRequestResponseBodyResultSubscribers) SetAvatarUrl(v string) *CreateMergeRequestResponseBodyResultSubscribers {
	s.AvatarUrl = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultSubscribers) SetEmail(v string) *CreateMergeRequestResponseBodyResultSubscribers {
	s.Email = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultSubscribers) SetId(v int64) *CreateMergeRequestResponseBodyResultSubscribers {
	s.Id = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultSubscribers) SetName(v string) *CreateMergeRequestResponseBodyResultSubscribers {
	s.Name = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultSubscribers) SetState(v string) *CreateMergeRequestResponseBodyResultSubscribers {
	s.State = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultSubscribers) SetUsername(v string) *CreateMergeRequestResponseBodyResultSubscribers {
	s.Username = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultSubscribers) Validate() error {
	return dara.Validate(s)
}

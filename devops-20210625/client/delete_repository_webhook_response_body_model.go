// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepositoryWebhookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteRepositoryWebhookResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteRepositoryWebhookResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteRepositoryWebhookResponseBody
	GetRequestId() *string
	SetResult(v *DeleteRepositoryWebhookResponseBodyResult) *DeleteRepositoryWebhookResponseBody
	GetResult() *DeleteRepositoryWebhookResponseBodyResult
	SetSuccess(v bool) *DeleteRepositoryWebhookResponseBody
	GetSuccess() *bool
}

type DeleteRepositoryWebhookResponseBody struct {
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
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DeleteRepositoryWebhookResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteRepositoryWebhookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryWebhookResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteRepositoryWebhookResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteRepositoryWebhookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRepositoryWebhookResponseBody) GetResult() *DeleteRepositoryWebhookResponseBodyResult {
	return s.Result
}

func (s *DeleteRepositoryWebhookResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRepositoryWebhookResponseBody) SetErrorCode(v string) *DeleteRepositoryWebhookResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBody) SetErrorMessage(v string) *DeleteRepositoryWebhookResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBody) SetRequestId(v string) *DeleteRepositoryWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBody) SetResult(v *DeleteRepositoryWebhookResponseBodyResult) *DeleteRepositoryWebhookResponseBody {
	s.Result = v
	return s
}

func (s *DeleteRepositoryWebhookResponseBody) SetSuccess(v bool) *DeleteRepositoryWebhookResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteRepositoryWebhookResponseBodyResult struct {
	// example:
	//
	// false
	BuildEvents *bool `json:"buildEvents,omitempty" xml:"buildEvents,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	CreatedAt   *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// false
	EnableSslVerification *bool `json:"enableSslVerification,omitempty" xml:"enableSslVerification,omitempty"`
	// example:
	//
	// 524836
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// false
	IssuesEvents *bool `json:"issuesEvents,omitempty" xml:"issuesEvents,omitempty"`
	// example:
	//
	// ""
	LastTestResult *string `json:"lastTestResult,omitempty" xml:"lastTestResult,omitempty"`
	// example:
	//
	// true
	MergeRequestsEvents *bool `json:"mergeRequestsEvents,omitempty" xml:"mergeRequestsEvents,omitempty"`
	// example:
	//
	// true
	NoteEvents *bool `json:"noteEvents,omitempty" xml:"noteEvents,omitempty"`
	// example:
	//
	// 2369234
	ProjectId *int64 `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// true
	PushEvents *bool `json:"pushEvents,omitempty" xml:"pushEvents,omitempty"`
	// example:
	//
	// xxxx
	SecretToken *string `json:"secretToken,omitempty" xml:"secretToken,omitempty"`
	// example:
	//
	// true
	TagPushEvents *bool `json:"tagPushEvents,omitempty" xml:"tagPushEvents,omitempty"`
	// example:
	//
	// ProjectHook
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// https://xxxxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s DeleteRepositoryWebhookResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryWebhookResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryWebhookResponseBodyResult) GetBuildEvents() *bool {
	return s.BuildEvents
}

func (s *DeleteRepositoryWebhookResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DeleteRepositoryWebhookResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *DeleteRepositoryWebhookResponseBodyResult) GetEnableSslVerification() *bool {
	return s.EnableSslVerification
}

func (s *DeleteRepositoryWebhookResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *DeleteRepositoryWebhookResponseBodyResult) GetIssuesEvents() *bool {
	return s.IssuesEvents
}

func (s *DeleteRepositoryWebhookResponseBodyResult) GetLastTestResult() *string {
	return s.LastTestResult
}

func (s *DeleteRepositoryWebhookResponseBodyResult) GetMergeRequestsEvents() *bool {
	return s.MergeRequestsEvents
}

func (s *DeleteRepositoryWebhookResponseBodyResult) GetNoteEvents() *bool {
	return s.NoteEvents
}

func (s *DeleteRepositoryWebhookResponseBodyResult) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteRepositoryWebhookResponseBodyResult) GetPushEvents() *bool {
	return s.PushEvents
}

func (s *DeleteRepositoryWebhookResponseBodyResult) GetSecretToken() *string {
	return s.SecretToken
}

func (s *DeleteRepositoryWebhookResponseBodyResult) GetTagPushEvents() *bool {
	return s.TagPushEvents
}

func (s *DeleteRepositoryWebhookResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *DeleteRepositoryWebhookResponseBodyResult) GetUrl() *string {
	return s.Url
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetBuildEvents(v bool) *DeleteRepositoryWebhookResponseBodyResult {
	s.BuildEvents = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetCreatedAt(v string) *DeleteRepositoryWebhookResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetDescription(v string) *DeleteRepositoryWebhookResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetEnableSslVerification(v bool) *DeleteRepositoryWebhookResponseBodyResult {
	s.EnableSslVerification = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetId(v int64) *DeleteRepositoryWebhookResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetIssuesEvents(v bool) *DeleteRepositoryWebhookResponseBodyResult {
	s.IssuesEvents = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetLastTestResult(v string) *DeleteRepositoryWebhookResponseBodyResult {
	s.LastTestResult = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetMergeRequestsEvents(v bool) *DeleteRepositoryWebhookResponseBodyResult {
	s.MergeRequestsEvents = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetNoteEvents(v bool) *DeleteRepositoryWebhookResponseBodyResult {
	s.NoteEvents = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetProjectId(v int64) *DeleteRepositoryWebhookResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetPushEvents(v bool) *DeleteRepositoryWebhookResponseBodyResult {
	s.PushEvents = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetSecretToken(v string) *DeleteRepositoryWebhookResponseBodyResult {
	s.SecretToken = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetTagPushEvents(v bool) *DeleteRepositoryWebhookResponseBodyResult {
	s.TagPushEvents = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetType(v string) *DeleteRepositoryWebhookResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetUrl(v string) *DeleteRepositoryWebhookResponseBodyResult {
	s.Url = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

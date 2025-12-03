// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryWebhookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListRepositoryWebhookResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListRepositoryWebhookResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListRepositoryWebhookResponseBody
	GetRequestId() *string
	SetResult(v []*ListRepositoryWebhookResponseBodyResult) *ListRepositoryWebhookResponseBody
	GetResult() []*ListRepositoryWebhookResponseBodyResult
	SetSuccess(v bool) *ListRepositoryWebhookResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListRepositoryWebhookResponseBody
	GetTotal() *int64
}

type ListRepositoryWebhookResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// HC93CE1A-8D7A-13A9-8306-7465DE2E5C0F
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListRepositoryWebhookResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListRepositoryWebhookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryWebhookResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListRepositoryWebhookResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListRepositoryWebhookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepositoryWebhookResponseBody) GetResult() []*ListRepositoryWebhookResponseBodyResult {
	return s.Result
}

func (s *ListRepositoryWebhookResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRepositoryWebhookResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListRepositoryWebhookResponseBody) SetErrorCode(v string) *ListRepositoryWebhookResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetErrorMessage(v string) *ListRepositoryWebhookResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetRequestId(v string) *ListRepositoryWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetResult(v []*ListRepositoryWebhookResponseBodyResult) *ListRepositoryWebhookResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetSuccess(v bool) *ListRepositoryWebhookResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetTotal(v int64) *ListRepositoryWebhookResponseBody {
	s.Total = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) Validate() error {
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

type ListRepositoryWebhookResponseBodyResult struct {
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
	// 16776
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LastTestResult *string `json:"lastTestResult,omitempty" xml:"lastTestResult,omitempty"`
	// example:
	//
	// true
	MergeRequestsEvents *bool `json:"mergeRequestsEvents,omitempty" xml:"mergeRequestsEvents,omitempty"`
	// example:
	//
	// false
	NoteEvents *bool  `json:"noteEvents,omitempty" xml:"noteEvents,omitempty"`
	ProjectId  *int64 `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// false
	PushEvents *bool `json:"pushEvents,omitempty" xml:"pushEvents,omitempty"`
	// example:
	//
	// xxxx
	SecretToken *string `json:"secretToken,omitempty" xml:"secretToken,omitempty"`
	// example:
	//
	// false
	TagPushEvents *bool `json:"tagPushEvents,omitempty" xml:"tagPushEvents,omitempty"`
	// example:
	//
	// https://xxxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListRepositoryWebhookResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryWebhookResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryWebhookResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListRepositoryWebhookResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListRepositoryWebhookResponseBodyResult) GetEnableSslVerification() *bool {
	return s.EnableSslVerification
}

func (s *ListRepositoryWebhookResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListRepositoryWebhookResponseBodyResult) GetLastTestResult() *string {
	return s.LastTestResult
}

func (s *ListRepositoryWebhookResponseBodyResult) GetMergeRequestsEvents() *bool {
	return s.MergeRequestsEvents
}

func (s *ListRepositoryWebhookResponseBodyResult) GetNoteEvents() *bool {
	return s.NoteEvents
}

func (s *ListRepositoryWebhookResponseBodyResult) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListRepositoryWebhookResponseBodyResult) GetPushEvents() *bool {
	return s.PushEvents
}

func (s *ListRepositoryWebhookResponseBodyResult) GetSecretToken() *string {
	return s.SecretToken
}

func (s *ListRepositoryWebhookResponseBodyResult) GetTagPushEvents() *bool {
	return s.TagPushEvents
}

func (s *ListRepositoryWebhookResponseBodyResult) GetUrl() *string {
	return s.Url
}

func (s *ListRepositoryWebhookResponseBodyResult) SetCreatedAt(v string) *ListRepositoryWebhookResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetDescription(v string) *ListRepositoryWebhookResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetEnableSslVerification(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.EnableSslVerification = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetId(v int64) *ListRepositoryWebhookResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetLastTestResult(v string) *ListRepositoryWebhookResponseBodyResult {
	s.LastTestResult = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetMergeRequestsEvents(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.MergeRequestsEvents = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetNoteEvents(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.NoteEvents = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetProjectId(v int64) *ListRepositoryWebhookResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetPushEvents(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.PushEvents = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetSecretToken(v string) *ListRepositoryWebhookResponseBodyResult {
	s.SecretToken = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetTagPushEvents(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.TagPushEvents = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetUrl(v string) *ListRepositoryWebhookResponseBodyResult {
	s.Url = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

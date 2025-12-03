// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWebhookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *AddWebhookResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddWebhookResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddWebhookResponseBody
	GetRequestId() *string
	SetResult(v *AddWebhookResponseBodyResult) *AddWebhookResponseBody
	GetResult() *AddWebhookResponseBodyResult
	SetSuccess(v bool) *AddWebhookResponseBody
	GetSuccess() *bool
}

type AddWebhookResponseBody struct {
	// example:
	//
	// SYSTEM_NOT_FOUND_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// “”
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 6177543A-8D54-5736-A93B-E0195A1512CB
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *AddWebhookResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddWebhookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *AddWebhookResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddWebhookResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddWebhookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddWebhookResponseBody) GetResult() *AddWebhookResponseBodyResult {
	return s.Result
}

func (s *AddWebhookResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddWebhookResponseBody) SetErrorCode(v string) *AddWebhookResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddWebhookResponseBody) SetErrorMessage(v string) *AddWebhookResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddWebhookResponseBody) SetRequestId(v string) *AddWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWebhookResponseBody) SetResult(v *AddWebhookResponseBodyResult) *AddWebhookResponseBody {
	s.Result = v
	return s
}

func (s *AddWebhookResponseBody) SetSuccess(v bool) *AddWebhookResponseBody {
	s.Success = &v
	return s
}

func (s *AddWebhookResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddWebhookResponseBodyResult struct {
	// example:
	//
	// 2022-03-12 12:00:00
	CreatedAt   *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// false
	EnableSslVerification *bool `json:"enableSslVerification,omitempty" xml:"enableSslVerification,omitempty"`
	// example:
	//
	// 30815
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// ok
	LastTestResult *string `json:"lastTestResult,omitempty" xml:"lastTestResult,omitempty"`
	// example:
	//
	// true
	MergeRequestsEvents *bool `json:"mergeRequestsEvents,omitempty" xml:"mergeRequestsEvents,omitempty"`
	// example:
	//
	// false
	NoteEvents *bool `json:"noteEvents,omitempty" xml:"noteEvents,omitempty"`
	// example:
	//
	// true
	PushEvents *bool `json:"pushEvents,omitempty" xml:"pushEvents,omitempty"`
	// example:
	//
	// 2835387
	RepositoryId *int64 `json:"repositoryId,omitempty" xml:"repositoryId,omitempty"`
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
	// https://xxxxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s AddWebhookResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s AddWebhookResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddWebhookResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *AddWebhookResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *AddWebhookResponseBodyResult) GetEnableSslVerification() *bool {
	return s.EnableSslVerification
}

func (s *AddWebhookResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *AddWebhookResponseBodyResult) GetLastTestResult() *string {
	return s.LastTestResult
}

func (s *AddWebhookResponseBodyResult) GetMergeRequestsEvents() *bool {
	return s.MergeRequestsEvents
}

func (s *AddWebhookResponseBodyResult) GetNoteEvents() *bool {
	return s.NoteEvents
}

func (s *AddWebhookResponseBodyResult) GetPushEvents() *bool {
	return s.PushEvents
}

func (s *AddWebhookResponseBodyResult) GetRepositoryId() *int64 {
	return s.RepositoryId
}

func (s *AddWebhookResponseBodyResult) GetSecretToken() *string {
	return s.SecretToken
}

func (s *AddWebhookResponseBodyResult) GetTagPushEvents() *bool {
	return s.TagPushEvents
}

func (s *AddWebhookResponseBodyResult) GetUrl() *string {
	return s.Url
}

func (s *AddWebhookResponseBodyResult) SetCreatedAt(v string) *AddWebhookResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetDescription(v string) *AddWebhookResponseBodyResult {
	s.Description = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetEnableSslVerification(v bool) *AddWebhookResponseBodyResult {
	s.EnableSslVerification = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetId(v int64) *AddWebhookResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetLastTestResult(v string) *AddWebhookResponseBodyResult {
	s.LastTestResult = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetMergeRequestsEvents(v bool) *AddWebhookResponseBodyResult {
	s.MergeRequestsEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetNoteEvents(v bool) *AddWebhookResponseBodyResult {
	s.NoteEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetPushEvents(v bool) *AddWebhookResponseBodyResult {
	s.PushEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetRepositoryId(v int64) *AddWebhookResponseBodyResult {
	s.RepositoryId = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetSecretToken(v string) *AddWebhookResponseBodyResult {
	s.SecretToken = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetTagPushEvents(v bool) *AddWebhookResponseBodyResult {
	s.TagPushEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetUrl(v string) *AddWebhookResponseBodyResult {
	s.Url = &v
	return s
}

func (s *AddWebhookResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

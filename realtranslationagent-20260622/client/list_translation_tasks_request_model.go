// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTranslationTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKey(v string) *ListTranslationTasksRequest
	GetAPIKey() *string
	SetEndTime(v string) *ListTranslationTasksRequest
	GetEndTime() *string
	SetMaxResults(v int32) *ListTranslationTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTranslationTasksRequest
	GetNextToken() *string
	SetOriginalFileName(v string) *ListTranslationTasksRequest
	GetOriginalFileName() *string
	SetSourceLanguage(v string) *ListTranslationTasksRequest
	GetSourceLanguage() *string
	SetStartTime(v string) *ListTranslationTasksRequest
	GetStartTime() *string
	SetStatus(v string) *ListTranslationTasksRequest
	GetStatus() *string
	SetTargetLanguage(v string) *ListTranslationTasksRequest
	GetTargetLanguage() *string
	SetTaskId(v string) *ListTranslationTasksRequest
	GetTaskId() *string
}

type ListTranslationTasksRequest struct {
	APIKey           *string `json:"APIKey,omitempty" xml:"APIKey,omitempty"`
	EndTime          *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxResults       *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OriginalFileName *string `json:"OriginalFileName,omitempty" xml:"OriginalFileName,omitempty"`
	SourceLanguage   *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetLanguage   *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	TaskId           *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListTranslationTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTranslationTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTranslationTasksRequest) GetAPIKey() *string {
	return s.APIKey
}

func (s *ListTranslationTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTranslationTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTranslationTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTranslationTasksRequest) GetOriginalFileName() *string {
	return s.OriginalFileName
}

func (s *ListTranslationTasksRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *ListTranslationTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTranslationTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTranslationTasksRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *ListTranslationTasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTranslationTasksRequest) SetAPIKey(v string) *ListTranslationTasksRequest {
	s.APIKey = &v
	return s
}

func (s *ListTranslationTasksRequest) SetEndTime(v string) *ListTranslationTasksRequest {
	s.EndTime = &v
	return s
}

func (s *ListTranslationTasksRequest) SetMaxResults(v int32) *ListTranslationTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTranslationTasksRequest) SetNextToken(v string) *ListTranslationTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListTranslationTasksRequest) SetOriginalFileName(v string) *ListTranslationTasksRequest {
	s.OriginalFileName = &v
	return s
}

func (s *ListTranslationTasksRequest) SetSourceLanguage(v string) *ListTranslationTasksRequest {
	s.SourceLanguage = &v
	return s
}

func (s *ListTranslationTasksRequest) SetStartTime(v string) *ListTranslationTasksRequest {
	s.StartTime = &v
	return s
}

func (s *ListTranslationTasksRequest) SetStatus(v string) *ListTranslationTasksRequest {
	s.Status = &v
	return s
}

func (s *ListTranslationTasksRequest) SetTargetLanguage(v string) *ListTranslationTasksRequest {
	s.TargetLanguage = &v
	return s
}

func (s *ListTranslationTasksRequest) SetTaskId(v string) *ListTranslationTasksRequest {
	s.TaskId = &v
	return s
}

func (s *ListTranslationTasksRequest) Validate() error {
	return dara.Validate(s)
}

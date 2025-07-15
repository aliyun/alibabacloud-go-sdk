// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchTaskDialogueDatasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeContent(v bool) *ListSearchTaskDialogueDatasRequest
	GetIncludeContent() *bool
	SetMultimodalSearchType(v string) *ListSearchTaskDialogueDatasRequest
	GetMultimodalSearchType() *string
	SetOriginalSessionId(v string) *ListSearchTaskDialogueDatasRequest
	GetOriginalSessionId() *string
	SetPageNumber(v int32) *ListSearchTaskDialogueDatasRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSearchTaskDialogueDatasRequest
	GetPageSize() *int32
	SetQuery(v string) *ListSearchTaskDialogueDatasRequest
	GetQuery() *string
	SetSearchModel(v string) *ListSearchTaskDialogueDatasRequest
	GetSearchModel() *string
	SetSearchModelDataValue(v string) *ListSearchTaskDialogueDatasRequest
	GetSearchModelDataValue() *string
	SetSessionId(v string) *ListSearchTaskDialogueDatasRequest
	GetSessionId() *string
	SetTaskId(v string) *ListSearchTaskDialogueDatasRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *ListSearchTaskDialogueDatasRequest
	GetWorkspaceId() *string
}

type ListSearchTaskDialogueDatasRequest struct {
	// example:
	//
	// true
	IncludeContent *bool `json:"IncludeContent,omitempty" xml:"IncludeContent,omitempty"`
	// example:
	//
	// text
	MultimodalSearchType *string `json:"MultimodalSearchType,omitempty" xml:"MultimodalSearchType,omitempty"`
	// example:
	//
	// xx
	OriginalSessionId *string `json:"OriginalSessionId,omitempty" xml:"OriginalSessionId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xx
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// ClusterGenerate
	SearchModel *string `json:"SearchModel,omitempty" xml:"SearchModel,omitempty"`
	// example:
	//
	// xxx
	SearchModelDataValue *string `json:"SearchModelDataValue,omitempty" xml:"SearchModelDataValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xx
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListSearchTaskDialogueDatasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTaskDialogueDatasRequest) GoString() string {
	return s.String()
}

func (s *ListSearchTaskDialogueDatasRequest) GetIncludeContent() *bool {
	return s.IncludeContent
}

func (s *ListSearchTaskDialogueDatasRequest) GetMultimodalSearchType() *string {
	return s.MultimodalSearchType
}

func (s *ListSearchTaskDialogueDatasRequest) GetOriginalSessionId() *string {
	return s.OriginalSessionId
}

func (s *ListSearchTaskDialogueDatasRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSearchTaskDialogueDatasRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSearchTaskDialogueDatasRequest) GetQuery() *string {
	return s.Query
}

func (s *ListSearchTaskDialogueDatasRequest) GetSearchModel() *string {
	return s.SearchModel
}

func (s *ListSearchTaskDialogueDatasRequest) GetSearchModelDataValue() *string {
	return s.SearchModelDataValue
}

func (s *ListSearchTaskDialogueDatasRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListSearchTaskDialogueDatasRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListSearchTaskDialogueDatasRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListSearchTaskDialogueDatasRequest) SetIncludeContent(v bool) *ListSearchTaskDialogueDatasRequest {
	s.IncludeContent = &v
	return s
}

func (s *ListSearchTaskDialogueDatasRequest) SetMultimodalSearchType(v string) *ListSearchTaskDialogueDatasRequest {
	s.MultimodalSearchType = &v
	return s
}

func (s *ListSearchTaskDialogueDatasRequest) SetOriginalSessionId(v string) *ListSearchTaskDialogueDatasRequest {
	s.OriginalSessionId = &v
	return s
}

func (s *ListSearchTaskDialogueDatasRequest) SetPageNumber(v int32) *ListSearchTaskDialogueDatasRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSearchTaskDialogueDatasRequest) SetPageSize(v int32) *ListSearchTaskDialogueDatasRequest {
	s.PageSize = &v
	return s
}

func (s *ListSearchTaskDialogueDatasRequest) SetQuery(v string) *ListSearchTaskDialogueDatasRequest {
	s.Query = &v
	return s
}

func (s *ListSearchTaskDialogueDatasRequest) SetSearchModel(v string) *ListSearchTaskDialogueDatasRequest {
	s.SearchModel = &v
	return s
}

func (s *ListSearchTaskDialogueDatasRequest) SetSearchModelDataValue(v string) *ListSearchTaskDialogueDatasRequest {
	s.SearchModelDataValue = &v
	return s
}

func (s *ListSearchTaskDialogueDatasRequest) SetSessionId(v string) *ListSearchTaskDialogueDatasRequest {
	s.SessionId = &v
	return s
}

func (s *ListSearchTaskDialogueDatasRequest) SetTaskId(v string) *ListSearchTaskDialogueDatasRequest {
	s.TaskId = &v
	return s
}

func (s *ListSearchTaskDialogueDatasRequest) SetWorkspaceId(v string) *ListSearchTaskDialogueDatasRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListSearchTaskDialogueDatasRequest) Validate() error {
	return dara.Validate(s)
}

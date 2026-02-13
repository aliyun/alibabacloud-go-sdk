// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDataAgentSessionResponseBodyData) *ListDataAgentSessionResponseBody
	GetData() []*ListDataAgentSessionResponseBodyData
	SetErrorCode(v string) *ListDataAgentSessionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataAgentSessionResponseBody
	GetErrorMessage() *string
	SetPageNumber(v int32) *ListDataAgentSessionResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataAgentSessionResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataAgentSessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataAgentSessionResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListDataAgentSessionResponseBody
	GetTotal() *int32
	SetTotalPages(v int32) *ListDataAgentSessionResponseBody
	GetTotalPages() *int32
}

type ListDataAgentSessionResponseBody struct {
	Data []*ListDataAgentSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E0D21075-CD3E-4D98-8264-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListDataAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataAgentSessionResponseBody) GetData() []*ListDataAgentSessionResponseBodyData {
	return s.Data
}

func (s *ListDataAgentSessionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataAgentSessionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataAgentSessionResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataAgentSessionResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataAgentSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataAgentSessionResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListDataAgentSessionResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListDataAgentSessionResponseBody) SetData(v []*ListDataAgentSessionResponseBodyData) *ListDataAgentSessionResponseBody {
	s.Data = v
	return s
}

func (s *ListDataAgentSessionResponseBody) SetErrorCode(v string) *ListDataAgentSessionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataAgentSessionResponseBody) SetErrorMessage(v string) *ListDataAgentSessionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataAgentSessionResponseBody) SetPageNumber(v int32) *ListDataAgentSessionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentSessionResponseBody) SetPageSize(v int32) *ListDataAgentSessionResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentSessionResponseBody) SetRequestId(v string) *ListDataAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataAgentSessionResponseBody) SetSuccess(v bool) *ListDataAgentSessionResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataAgentSessionResponseBody) SetTotal(v int32) *ListDataAgentSessionResponseBody {
	s.Total = &v
	return s
}

func (s *ListDataAgentSessionResponseBody) SetTotalPages(v int32) *ListDataAgentSessionResponseBody {
	s.TotalPages = &v
	return s
}

func (s *ListDataAgentSessionResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataAgentSessionResponseBodyData struct {
	// example:
	//
	// cu0cs*******mf
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// RUNNING
	AgentStatus *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// example:
	//
	// 1731645908000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	FavoriteInWorkspace *bool `json:"FavoriteInWorkspace,omitempty" xml:"FavoriteInWorkspace,omitempty"`
	// example:
	//
	// f-8*******01m
	File *string `json:"File,omitempty" xml:"File,omitempty"`
	// example:
	//
	// true
	Saved         *bool                                              `json:"Saved,omitempty" xml:"Saved,omitempty"`
	SessionConfig *ListDataAgentSessionResponseBodyDataSessionConfig `json:"SessionConfig,omitempty" xml:"SessionConfig,omitempty" type:"Struct"`
	// example:
	//
	// h8r********4fch
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// RUNNING
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2096******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListDataAgentSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataAgentSessionResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *ListDataAgentSessionResponseBodyData) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *ListDataAgentSessionResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataAgentSessionResponseBodyData) GetFavoriteInWorkspace() *bool {
	return s.FavoriteInWorkspace
}

func (s *ListDataAgentSessionResponseBodyData) GetFile() *string {
	return s.File
}

func (s *ListDataAgentSessionResponseBodyData) GetSaved() *bool {
	return s.Saved
}

func (s *ListDataAgentSessionResponseBodyData) GetSessionConfig() *ListDataAgentSessionResponseBodyDataSessionConfig {
	return s.SessionConfig
}

func (s *ListDataAgentSessionResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *ListDataAgentSessionResponseBodyData) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *ListDataAgentSessionResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListDataAgentSessionResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ListDataAgentSessionResponseBodyData) SetAgentId(v string) *ListDataAgentSessionResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyData) SetAgentStatus(v string) *ListDataAgentSessionResponseBodyData {
	s.AgentStatus = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyData) SetCreateTime(v int64) *ListDataAgentSessionResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyData) SetFavoriteInWorkspace(v bool) *ListDataAgentSessionResponseBodyData {
	s.FavoriteInWorkspace = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyData) SetFile(v string) *ListDataAgentSessionResponseBodyData {
	s.File = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyData) SetSaved(v bool) *ListDataAgentSessionResponseBodyData {
	s.Saved = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyData) SetSessionConfig(v *ListDataAgentSessionResponseBodyDataSessionConfig) *ListDataAgentSessionResponseBodyData {
	s.SessionConfig = v
	return s
}

func (s *ListDataAgentSessionResponseBodyData) SetSessionId(v string) *ListDataAgentSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyData) SetSessionStatus(v string) *ListDataAgentSessionResponseBodyData {
	s.SessionStatus = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyData) SetTitle(v string) *ListDataAgentSessionResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyData) SetUserId(v string) *ListDataAgentSessionResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyData) Validate() error {
	if s.SessionConfig != nil {
		if err := s.SessionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataAgentSessionResponseBodyDataSessionConfig struct {
	// example:
	//
	// ca-e*******ckd
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// example:
	//
	// prod
	CustomAgentStage *string `json:"CustomAgentStage,omitempty" xml:"CustomAgentStage,omitempty"`
	// example:
	//
	// true
	EnableSearch *bool `json:"EnableSearch,omitempty" xml:"EnableSearch,omitempty"`
	// example:
	//
	// CHINESE
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// ANALYSIS
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// user-oss-bucket
	UserOssBucket *string `json:"UserOssBucket,omitempty" xml:"UserOssBucket,omitempty"`
}

func (s ListDataAgentSessionResponseBodyDataSessionConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentSessionResponseBodyDataSessionConfig) GoString() string {
	return s.String()
}

func (s *ListDataAgentSessionResponseBodyDataSessionConfig) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *ListDataAgentSessionResponseBodyDataSessionConfig) GetCustomAgentStage() *string {
	return s.CustomAgentStage
}

func (s *ListDataAgentSessionResponseBodyDataSessionConfig) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *ListDataAgentSessionResponseBodyDataSessionConfig) GetLanguage() *string {
	return s.Language
}

func (s *ListDataAgentSessionResponseBodyDataSessionConfig) GetMode() *string {
	return s.Mode
}

func (s *ListDataAgentSessionResponseBodyDataSessionConfig) GetUserOssBucket() *string {
	return s.UserOssBucket
}

func (s *ListDataAgentSessionResponseBodyDataSessionConfig) SetCustomAgentId(v string) *ListDataAgentSessionResponseBodyDataSessionConfig {
	s.CustomAgentId = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyDataSessionConfig) SetCustomAgentStage(v string) *ListDataAgentSessionResponseBodyDataSessionConfig {
	s.CustomAgentStage = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyDataSessionConfig) SetEnableSearch(v bool) *ListDataAgentSessionResponseBodyDataSessionConfig {
	s.EnableSearch = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyDataSessionConfig) SetLanguage(v string) *ListDataAgentSessionResponseBodyDataSessionConfig {
	s.Language = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyDataSessionConfig) SetMode(v string) *ListDataAgentSessionResponseBodyDataSessionConfig {
	s.Mode = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyDataSessionConfig) SetUserOssBucket(v string) *ListDataAgentSessionResponseBodyDataSessionConfig {
	s.UserOssBucket = &v
	return s
}

func (s *ListDataAgentSessionResponseBodyDataSessionConfig) Validate() error {
	return dara.Validate(s)
}

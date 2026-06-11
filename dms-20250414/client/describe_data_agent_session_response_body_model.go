// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDataAgentSessionResponseBodyData) *DescribeDataAgentSessionResponseBody
	GetData() *DescribeDataAgentSessionResponseBodyData
	SetErrorCode(v string) *DescribeDataAgentSessionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DescribeDataAgentSessionResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DescribeDataAgentSessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDataAgentSessionResponseBody
	GetSuccess() *bool
}

type DescribeDataAgentSessionResponseBody struct {
	// The response data.
	Data *DescribeDataAgentSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 18****-*****-*******7A3122F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDataAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentSessionResponseBody) GetData() *DescribeDataAgentSessionResponseBodyData {
	return s.Data
}

func (s *DescribeDataAgentSessionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeDataAgentSessionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDataAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataAgentSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDataAgentSessionResponseBody) SetData(v *DescribeDataAgentSessionResponseBodyData) *DescribeDataAgentSessionResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDataAgentSessionResponseBody) SetErrorCode(v string) *DescribeDataAgentSessionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBody) SetErrorMessage(v string) *DescribeDataAgentSessionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBody) SetRequestId(v string) *DescribeDataAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBody) SetSuccess(v bool) *DescribeDataAgentSessionResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataAgentSessionResponseBodyData struct {
	// The ID of the agent.
	//
	// example:
	//
	// cu0cs*******mf
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The status of the agent.
	//
	// example:
	//
	// RUNNING
	AgentStatus *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// The session replay history.
	ChatHistoryLocations []*DescribeDataAgentSessionResponseBodyDataChatHistoryLocations `json:"ChatHistoryLocations,omitempty" xml:"ChatHistoryLocations,omitempty" type:"Repeated"`
	// The timestamp indicating when the session was created.
	//
	// example:
	//
	// 1731645908000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the current user has favorited the session in the workspace.
	//
	// example:
	//
	// true
	FavoriteInWorkspace *string `json:"FavoriteInWorkspace,omitempty" xml:"FavoriteInWorkspace,omitempty"`
	// The ID of the file.
	//
	// example:
	//
	// f-8*******01m
	File *string `json:"File,omitempty" xml:"File,omitempty"`
	// Indicates whether the current user has favorited the session.
	//
	// example:
	//
	// true
	Saved *bool `json:"Saved,omitempty" xml:"Saved,omitempty"`
	// The configuration of the session.
	SessionConfig *DescribeDataAgentSessionResponseBodyDataSessionConfig `json:"SessionConfig,omitempty" xml:"SessionConfig,omitempty" type:"Struct"`
	// The ID of the agent session.
	//
	// example:
	//
	// 976*********p
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The status of the session.
	//
	// example:
	//
	// RUNNING
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	// The title of the session.
	//
	// example:
	//
	// 分析一下这份文件，给出报告。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The ID of the session owner.
	//
	// example:
	//
	// 2096******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeDataAgentSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentSessionResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *DescribeDataAgentSessionResponseBodyData) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *DescribeDataAgentSessionResponseBodyData) GetChatHistoryLocations() []*DescribeDataAgentSessionResponseBodyDataChatHistoryLocations {
	return s.ChatHistoryLocations
}

func (s *DescribeDataAgentSessionResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeDataAgentSessionResponseBodyData) GetFavoriteInWorkspace() *string {
	return s.FavoriteInWorkspace
}

func (s *DescribeDataAgentSessionResponseBodyData) GetFile() *string {
	return s.File
}

func (s *DescribeDataAgentSessionResponseBodyData) GetSaved() *bool {
	return s.Saved
}

func (s *DescribeDataAgentSessionResponseBodyData) GetSessionConfig() *DescribeDataAgentSessionResponseBodyDataSessionConfig {
	return s.SessionConfig
}

func (s *DescribeDataAgentSessionResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeDataAgentSessionResponseBodyData) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *DescribeDataAgentSessionResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *DescribeDataAgentSessionResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *DescribeDataAgentSessionResponseBodyData) SetAgentId(v string) *DescribeDataAgentSessionResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyData) SetAgentStatus(v string) *DescribeDataAgentSessionResponseBodyData {
	s.AgentStatus = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyData) SetChatHistoryLocations(v []*DescribeDataAgentSessionResponseBodyDataChatHistoryLocations) *DescribeDataAgentSessionResponseBodyData {
	s.ChatHistoryLocations = v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyData) SetCreateTime(v int64) *DescribeDataAgentSessionResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyData) SetFavoriteInWorkspace(v string) *DescribeDataAgentSessionResponseBodyData {
	s.FavoriteInWorkspace = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyData) SetFile(v string) *DescribeDataAgentSessionResponseBodyData {
	s.File = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyData) SetSaved(v bool) *DescribeDataAgentSessionResponseBodyData {
	s.Saved = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyData) SetSessionConfig(v *DescribeDataAgentSessionResponseBodyDataSessionConfig) *DescribeDataAgentSessionResponseBodyData {
	s.SessionConfig = v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyData) SetSessionId(v string) *DescribeDataAgentSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyData) SetSessionStatus(v string) *DescribeDataAgentSessionResponseBodyData {
	s.SessionStatus = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyData) SetTitle(v string) *DescribeDataAgentSessionResponseBodyData {
	s.Title = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyData) SetUserId(v string) *DescribeDataAgentSessionResponseBodyData {
	s.UserId = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyData) Validate() error {
	if s.ChatHistoryLocations != nil {
		for _, item := range s.ChatHistoryLocations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SessionConfig != nil {
		if err := s.SessionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataAgentSessionResponseBodyDataChatHistoryLocations struct {
	// The key of the session replay history item.
	//
	// example:
	//
	// testKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The OSS download URL for the session replay history item.
	//
	// example:
	//
	// ****
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeDataAgentSessionResponseBodyDataChatHistoryLocations) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentSessionResponseBodyDataChatHistoryLocations) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentSessionResponseBodyDataChatHistoryLocations) GetKey() *string {
	return s.Key
}

func (s *DescribeDataAgentSessionResponseBodyDataChatHistoryLocations) GetUrl() *string {
	return s.Url
}

func (s *DescribeDataAgentSessionResponseBodyDataChatHistoryLocations) SetKey(v string) *DescribeDataAgentSessionResponseBodyDataChatHistoryLocations {
	s.Key = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataChatHistoryLocations) SetUrl(v string) *DescribeDataAgentSessionResponseBodyDataChatHistoryLocations {
	s.Url = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataChatHistoryLocations) Validate() error {
	return dara.Validate(s)
}

type DescribeDataAgentSessionResponseBodyDataSessionConfig struct {
	// The ID of the custom agent.
	//
	// example:
	//
	// ca-e*******ckd
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// The stage of the custom agent. Valid values:
	//
	// - **debug**: The test stage.
	//
	// - **prod**: The production stage.
	//
	// example:
	//
	// debug
	CustomAgentStage *string `json:"CustomAgentStage,omitempty" xml:"CustomAgentStage,omitempty"`
	// Indicates whether web search is enabled.
	//
	// example:
	//
	// True
	EnableSearch *bool     `json:"EnableSearch,omitempty" xml:"EnableSearch,omitempty"`
	EncryptKey   *string   `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
	EncryptType  *string   `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	KbUuidList   []*string `json:"KbUuidList,omitempty" xml:"KbUuidList,omitempty" type:"Repeated"`
	// The language. Valid values:
	//
	// - **CHINESE**: Chinese
	//
	// - **ENGLISH**: English
	//
	// example:
	//
	// CHINESE
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// A list of MCP server IDs.
	McpServerIds []*string `json:"McpServerIds,omitempty" xml:"McpServerIds,omitempty" type:"Repeated"`
	// The mode. Valid values:
	//
	// - **ASK_DATA**: quick inquiry mode
	//
	// - **ANALYSIS**: analysis mode
	//
	// - **INSIGHT**: insight mode
	//
	// example:
	//
	// ANALYSIS
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ReportPageWidth *int64  `json:"ReportPageWidth,omitempty" xml:"ReportPageWidth,omitempty"`
	ReportWaterMark *string `json:"ReportWaterMark,omitempty" xml:"ReportWaterMark,omitempty"`
	// The name of the user\\"s OSS bucket.
	//
	// - The service can upload analysis files and reports to this bucket.
	//
	// example:
	//
	// user-oss-bucket
	UserOssBucket *string `json:"UserOssBucket,omitempty" xml:"UserOssBucket,omitempty"`
}

func (s DescribeDataAgentSessionResponseBodyDataSessionConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentSessionResponseBodyDataSessionConfig) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) GetCustomAgentStage() *string {
	return s.CustomAgentStage
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) GetEncryptKey() *string {
	return s.EncryptKey
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) GetEncryptType() *string {
	return s.EncryptType
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) GetKbUuidList() []*string {
	return s.KbUuidList
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) GetLanguage() *string {
	return s.Language
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) GetMcpServerIds() []*string {
	return s.McpServerIds
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) GetMode() *string {
	return s.Mode
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) GetReportPageWidth() *int64 {
	return s.ReportPageWidth
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) GetReportWaterMark() *string {
	return s.ReportWaterMark
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) GetUserOssBucket() *string {
	return s.UserOssBucket
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) SetCustomAgentId(v string) *DescribeDataAgentSessionResponseBodyDataSessionConfig {
	s.CustomAgentId = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) SetCustomAgentStage(v string) *DescribeDataAgentSessionResponseBodyDataSessionConfig {
	s.CustomAgentStage = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) SetEnableSearch(v bool) *DescribeDataAgentSessionResponseBodyDataSessionConfig {
	s.EnableSearch = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) SetEncryptKey(v string) *DescribeDataAgentSessionResponseBodyDataSessionConfig {
	s.EncryptKey = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) SetEncryptType(v string) *DescribeDataAgentSessionResponseBodyDataSessionConfig {
	s.EncryptType = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) SetKbUuidList(v []*string) *DescribeDataAgentSessionResponseBodyDataSessionConfig {
	s.KbUuidList = v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) SetLanguage(v string) *DescribeDataAgentSessionResponseBodyDataSessionConfig {
	s.Language = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) SetMcpServerIds(v []*string) *DescribeDataAgentSessionResponseBodyDataSessionConfig {
	s.McpServerIds = v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) SetMode(v string) *DescribeDataAgentSessionResponseBodyDataSessionConfig {
	s.Mode = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) SetReportPageWidth(v int64) *DescribeDataAgentSessionResponseBodyDataSessionConfig {
	s.ReportPageWidth = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) SetReportWaterMark(v string) *DescribeDataAgentSessionResponseBodyDataSessionConfig {
	s.ReportWaterMark = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) SetUserOssBucket(v string) *DescribeDataAgentSessionResponseBodyDataSessionConfig {
	s.UserOssBucket = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataSessionConfig) Validate() error {
	return dara.Validate(s)
}

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
	// The response struct.
	Data *DescribeDataAgentSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 18****-*****-*******7A3122F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return value. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
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
	// The current agent ID.
	//
	// example:
	//
	// cu0cs*******mf
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The current agent status.
	//
	// example:
	//
	// RUNNING
	AgentStatus *string                                              `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	Artifacts   []*DescribeDataAgentSessionResponseBodyDataArtifacts `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	// The chat replay history.
	ChatHistoryLocations []*DescribeDataAgentSessionResponseBodyDataChatHistoryLocations `json:"ChatHistoryLocations,omitempty" xml:"ChatHistoryLocations,omitempty" type:"Repeated"`
	// The session creation time.
	//
	// example:
	//
	// 1731645908000
	CreateTime  *int64                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataSources []*DescribeDataAgentSessionResponseBodyDataDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// Indicates whether the session is saved as a favorite in the workspace by the current logged-in user.
	//
	// example:
	//
	// true
	FavoriteInWorkspace *string `json:"FavoriteInWorkspace,omitempty" xml:"FavoriteInWorkspace,omitempty"`
	// The file ID.
	//
	// example:
	//
	// f-8*******01m
	File          *string                                                  `json:"File,omitempty" xml:"File,omitempty"`
	RecallResults []*DescribeDataAgentSessionResponseBodyDataRecallResults `json:"RecallResults,omitempty" xml:"RecallResults,omitempty" type:"Repeated"`
	// Indicates whether the session is saved as a favorite by the current logged-in user.
	//
	// example:
	//
	// true
	Saved *bool `json:"Saved,omitempty" xml:"Saved,omitempty"`
	// The session configuration item.
	SessionConfig *DescribeDataAgentSessionResponseBodyDataSessionConfig `json:"SessionConfig,omitempty" xml:"SessionConfig,omitempty" type:"Struct"`
	// The agent session ID.
	//
	// example:
	//
	// 976*********p
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The session status.
	//
	// example:
	//
	// RUNNING
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	// The title.
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

func (s *DescribeDataAgentSessionResponseBodyData) GetArtifacts() []*DescribeDataAgentSessionResponseBodyDataArtifacts {
	return s.Artifacts
}

func (s *DescribeDataAgentSessionResponseBodyData) GetChatHistoryLocations() []*DescribeDataAgentSessionResponseBodyDataChatHistoryLocations {
	return s.ChatHistoryLocations
}

func (s *DescribeDataAgentSessionResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeDataAgentSessionResponseBodyData) GetDataSources() []*DescribeDataAgentSessionResponseBodyDataDataSources {
	return s.DataSources
}

func (s *DescribeDataAgentSessionResponseBodyData) GetFavoriteInWorkspace() *string {
	return s.FavoriteInWorkspace
}

func (s *DescribeDataAgentSessionResponseBodyData) GetFile() *string {
	return s.File
}

func (s *DescribeDataAgentSessionResponseBodyData) GetRecallResults() []*DescribeDataAgentSessionResponseBodyDataRecallResults {
	return s.RecallResults
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

func (s *DescribeDataAgentSessionResponseBodyData) SetArtifacts(v []*DescribeDataAgentSessionResponseBodyDataArtifacts) *DescribeDataAgentSessionResponseBodyData {
	s.Artifacts = v
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

func (s *DescribeDataAgentSessionResponseBodyData) SetDataSources(v []*DescribeDataAgentSessionResponseBodyDataDataSources) *DescribeDataAgentSessionResponseBodyData {
	s.DataSources = v
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

func (s *DescribeDataAgentSessionResponseBodyData) SetRecallResults(v []*DescribeDataAgentSessionResponseBodyDataRecallResults) *DescribeDataAgentSessionResponseBodyData {
	s.RecallResults = v
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
	if s.Artifacts != nil {
		for _, item := range s.Artifacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ChatHistoryLocations != nil {
		for _, item := range s.ChatHistoryLocations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DataSources != nil {
		for _, item := range s.DataSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RecallResults != nil {
		for _, item := range s.RecallResults {
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

type DescribeDataAgentSessionResponseBodyDataArtifacts struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FinishTime  *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ReceiveTime *string `json:"ReceiveTime,omitempty" xml:"ReceiveTime,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDataAgentSessionResponseBodyDataArtifacts) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentSessionResponseBodyDataArtifacts) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) GetDescription() *string {
	return s.Description
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) GetId() *string {
	return s.Id
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) GetName() *string {
	return s.Name
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) GetReceiveTime() *string {
	return s.ReceiveTime
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) GetStatus() *string {
	return s.Status
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) GetType() *string {
	return s.Type
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) SetDescription(v string) *DescribeDataAgentSessionResponseBodyDataArtifacts {
	s.Description = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) SetFinishTime(v string) *DescribeDataAgentSessionResponseBodyDataArtifacts {
	s.FinishTime = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) SetId(v string) *DescribeDataAgentSessionResponseBodyDataArtifacts {
	s.Id = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) SetName(v string) *DescribeDataAgentSessionResponseBodyDataArtifacts {
	s.Name = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) SetReceiveTime(v string) *DescribeDataAgentSessionResponseBodyDataArtifacts {
	s.ReceiveTime = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) SetStartTime(v string) *DescribeDataAgentSessionResponseBodyDataArtifacts {
	s.StartTime = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) SetStatus(v string) *DescribeDataAgentSessionResponseBodyDataArtifacts {
	s.Status = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) SetType(v string) *DescribeDataAgentSessionResponseBodyDataArtifacts {
	s.Type = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataArtifacts) Validate() error {
	return dara.Validate(s)
}

type DescribeDataAgentSessionResponseBodyDataChatHistoryLocations struct {
	// The key of the chat replay history.
	//
	// example:
	//
	// testKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The OSS download URL of the chat replay history.
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

type DescribeDataAgentSessionResponseBodyDataDataSources struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Detail   *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
}

func (s DescribeDataAgentSessionResponseBodyDataDataSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentSessionResponseBodyDataDataSources) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentSessionResponseBodyDataDataSources) GetCategory() *string {
	return s.Category
}

func (s *DescribeDataAgentSessionResponseBodyDataDataSources) GetDetail() *string {
	return s.Detail
}

func (s *DescribeDataAgentSessionResponseBodyDataDataSources) SetCategory(v string) *DescribeDataAgentSessionResponseBodyDataDataSources {
	s.Category = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataDataSources) SetDetail(v string) *DescribeDataAgentSessionResponseBodyDataDataSources {
	s.Detail = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataDataSources) Validate() error {
	return dara.Validate(s)
}

type DescribeDataAgentSessionResponseBodyDataRecallResults struct {
	Content *string  `json:"Content,omitempty" xml:"Content,omitempty"`
	Score   *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	Type    *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDataAgentSessionResponseBodyDataRecallResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentSessionResponseBodyDataRecallResults) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentSessionResponseBodyDataRecallResults) GetContent() *string {
	return s.Content
}

func (s *DescribeDataAgentSessionResponseBodyDataRecallResults) GetScore() *float64 {
	return s.Score
}

func (s *DescribeDataAgentSessionResponseBodyDataRecallResults) GetType() *string {
	return s.Type
}

func (s *DescribeDataAgentSessionResponseBodyDataRecallResults) SetContent(v string) *DescribeDataAgentSessionResponseBodyDataRecallResults {
	s.Content = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataRecallResults) SetScore(v float64) *DescribeDataAgentSessionResponseBodyDataRecallResults {
	s.Score = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataRecallResults) SetType(v string) *DescribeDataAgentSessionResponseBodyDataRecallResults {
	s.Type = &v
	return s
}

func (s *DescribeDataAgentSessionResponseBodyDataRecallResults) Validate() error {
	return dara.Validate(s)
}

type DescribeDataAgentSessionResponseBodyDataSessionConfig struct {
	// The custom agent ID.
	//
	// example:
	//
	// ca-e*******ckd
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// The stage of the custom agent. Valid values:
	//
	// - **debug**: Debug stage.
	//
	// - **prod**: Production stage.
	//
	// example:
	//
	// debug
	CustomAgentStage *string `json:"CustomAgentStage,omitempty" xml:"CustomAgentStage,omitempty"`
	// Specifies whether to enable web search.
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
	// - **CHINESE**: Chinese.
	//
	// - **ENGLISH**: English.
	//
	// example:
	//
	// CHINESE
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The list of MCP server IDs in the session configuration.
	McpServerIds []*string `json:"McpServerIds,omitempty" xml:"McpServerIds,omitempty" type:"Repeated"`
	// The mode. Valid values:
	//
	// - **ASK_DATA**: Ask data mode.
	//
	// - **ANALYSIS**: Analysis mode.
	//
	// - **INSIGHT**: Insight mode.
	//
	// example:
	//
	// ANALYSIS
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ReportPageWidth *int64  `json:"ReportPageWidth,omitempty" xml:"ReportPageWidth,omitempty"`
	ReportWaterMark *string `json:"ReportWaterMark,omitempty" xml:"ReportWaterMark,omitempty"`
	// The name of the user OSS bucket.
	//
	// - Analysis process files and report artifacts can be uploaded to the user-specified OSS bucket.
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

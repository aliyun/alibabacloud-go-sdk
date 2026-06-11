// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDataAgentSessionResponseBodyData) *CreateDataAgentSessionResponseBody
	GetData() *CreateDataAgentSessionResponseBodyData
	SetErrorCode(v string) *CreateDataAgentSessionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataAgentSessionResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDataAgentSessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataAgentSessionResponseBody
	GetSuccess() *bool
}

type CreateDataAgentSessionResponseBody struct {
	// The response structure.
	Data *CreateDataAgentSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CB***********3F1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
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

func (s CreateDataAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataAgentSessionResponseBody) GetData() *CreateDataAgentSessionResponseBodyData {
	return s.Data
}

func (s *CreateDataAgentSessionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataAgentSessionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataAgentSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataAgentSessionResponseBody) SetData(v *CreateDataAgentSessionResponseBodyData) *CreateDataAgentSessionResponseBody {
	s.Data = v
	return s
}

func (s *CreateDataAgentSessionResponseBody) SetErrorCode(v string) *CreateDataAgentSessionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataAgentSessionResponseBody) SetErrorMessage(v string) *CreateDataAgentSessionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataAgentSessionResponseBody) SetRequestId(v string) *CreateDataAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataAgentSessionResponseBody) SetSuccess(v bool) *CreateDataAgentSessionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataAgentSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataAgentSessionResponseBodyData struct {
	// The agent ID.
	//
	// example:
	//
	// cu0cs*******mf
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The agent status.
	//
	// example:
	//
	// RUNNING
	AgentStatus *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// The time when the session was created, in Unix milliseconds.
	//
	// example:
	//
	// 1765262307992
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the associated file.
	//
	// example:
	//
	// f-8*******01m
	File *string `json:"File,omitempty" xml:"File,omitempty"`
	// Indicates whether the current user has favorited the session.
	//
	// example:
	//
	// false
	Saved *bool `json:"Saved,omitempty" xml:"Saved,omitempty"`
	// The session configuration.
	SessionConfig *CreateDataAgentSessionResponseBodyDataSessionConfig `json:"SessionConfig,omitempty" xml:"SessionConfig,omitempty" type:"Struct"`
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
	// The session title.
	//
	// example:
	//
	// 帮我分析一下这份数据，给出报告。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateDataAgentSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDataAgentSessionResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateDataAgentSessionResponseBodyData) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *CreateDataAgentSessionResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateDataAgentSessionResponseBodyData) GetFile() *string {
	return s.File
}

func (s *CreateDataAgentSessionResponseBodyData) GetSaved() *bool {
	return s.Saved
}

func (s *CreateDataAgentSessionResponseBodyData) GetSessionConfig() *CreateDataAgentSessionResponseBodyDataSessionConfig {
	return s.SessionConfig
}

func (s *CreateDataAgentSessionResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateDataAgentSessionResponseBodyData) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *CreateDataAgentSessionResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *CreateDataAgentSessionResponseBodyData) SetAgentId(v string) *CreateDataAgentSessionResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyData) SetAgentStatus(v string) *CreateDataAgentSessionResponseBodyData {
	s.AgentStatus = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyData) SetCreateTime(v int64) *CreateDataAgentSessionResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyData) SetFile(v string) *CreateDataAgentSessionResponseBodyData {
	s.File = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyData) SetSaved(v bool) *CreateDataAgentSessionResponseBodyData {
	s.Saved = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyData) SetSessionConfig(v *CreateDataAgentSessionResponseBodyDataSessionConfig) *CreateDataAgentSessionResponseBodyData {
	s.SessionConfig = v
	return s
}

func (s *CreateDataAgentSessionResponseBodyData) SetSessionId(v string) *CreateDataAgentSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyData) SetSessionStatus(v string) *CreateDataAgentSessionResponseBodyData {
	s.SessionStatus = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyData) SetTitle(v string) *CreateDataAgentSessionResponseBodyData {
	s.Title = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyData) Validate() error {
	if s.SessionConfig != nil {
		if err := s.SessionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataAgentSessionResponseBodyDataSessionConfig struct {
	// The custom agent ID.
	//
	// example:
	//
	// ca-e*******ckd
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// The stage of the custom agent. Valid values:
	//
	// - **debug**: Debug stage
	//
	// - **prod**: Production stage
	//
	// example:
	//
	// debug
	CustomAgentStage *string `json:"CustomAgentStage,omitempty" xml:"CustomAgentStage,omitempty"`
	// Indicates whether web search is enabled.
	//
	// example:
	//
	// true
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
	// - **ASK_DATA**: Quick Inquiry Mode
	//
	// - **ANALYSIS**: Analysis Mode
	//
	// - **INSIGHT**: Insight Mode
	//
	// example:
	//
	// ANALYSIS
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ReportPageWidth *int64  `json:"ReportPageWidth,omitempty" xml:"ReportPageWidth,omitempty"`
	ReportWaterMark *string `json:"ReportWaterMark,omitempty" xml:"ReportWaterMark,omitempty"`
	// The name of the user\\"s OSS bucket.
	//
	// - Analysis files and report artifacts can be uploaded to this OSS bucket.
	//
	// example:
	//
	// user-oss-bucket
	UserOssBucket *string `json:"UserOssBucket,omitempty" xml:"UserOssBucket,omitempty"`
}

func (s CreateDataAgentSessionResponseBodyDataSessionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentSessionResponseBodyDataSessionConfig) GoString() string {
	return s.String()
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) GetCustomAgentStage() *string {
	return s.CustomAgentStage
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) GetEncryptKey() *string {
	return s.EncryptKey
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) GetEncryptType() *string {
	return s.EncryptType
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) GetKbUuidList() []*string {
	return s.KbUuidList
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) GetLanguage() *string {
	return s.Language
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) GetMcpServerIds() []*string {
	return s.McpServerIds
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) GetMode() *string {
	return s.Mode
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) GetReportPageWidth() *int64 {
	return s.ReportPageWidth
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) GetReportWaterMark() *string {
	return s.ReportWaterMark
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) GetUserOssBucket() *string {
	return s.UserOssBucket
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) SetCustomAgentId(v string) *CreateDataAgentSessionResponseBodyDataSessionConfig {
	s.CustomAgentId = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) SetCustomAgentStage(v string) *CreateDataAgentSessionResponseBodyDataSessionConfig {
	s.CustomAgentStage = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) SetEnableSearch(v bool) *CreateDataAgentSessionResponseBodyDataSessionConfig {
	s.EnableSearch = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) SetEncryptKey(v string) *CreateDataAgentSessionResponseBodyDataSessionConfig {
	s.EncryptKey = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) SetEncryptType(v string) *CreateDataAgentSessionResponseBodyDataSessionConfig {
	s.EncryptType = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) SetKbUuidList(v []*string) *CreateDataAgentSessionResponseBodyDataSessionConfig {
	s.KbUuidList = v
	return s
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) SetLanguage(v string) *CreateDataAgentSessionResponseBodyDataSessionConfig {
	s.Language = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) SetMcpServerIds(v []*string) *CreateDataAgentSessionResponseBodyDataSessionConfig {
	s.McpServerIds = v
	return s
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) SetMode(v string) *CreateDataAgentSessionResponseBodyDataSessionConfig {
	s.Mode = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) SetReportPageWidth(v int64) *CreateDataAgentSessionResponseBodyDataSessionConfig {
	s.ReportPageWidth = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) SetReportWaterMark(v string) *CreateDataAgentSessionResponseBodyDataSessionConfig {
	s.ReportWaterMark = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) SetUserOssBucket(v string) *CreateDataAgentSessionResponseBodyDataSessionConfig {
	s.UserOssBucket = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) Validate() error {
	return dara.Validate(s)
}

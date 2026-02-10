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
	Data *DescribeDataAgentSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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
	// example:
	//
	// cu0cs*******mf
	AgentId              *string                                                         `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentStatus          *string                                                         `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	ChatHistoryLocations []*DescribeDataAgentSessionResponseBodyDataChatHistoryLocations `json:"ChatHistoryLocations,omitempty" xml:"ChatHistoryLocations,omitempty" type:"Repeated"`
	// example:
	//
	// 1731645908000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	FavoriteInWorkspace *string `json:"FavoriteInWorkspace,omitempty" xml:"FavoriteInWorkspace,omitempty"`
	// example:
	//
	// f-8*******01m
	File *string `json:"File,omitempty" xml:"File,omitempty"`
	// example:
	//
	// true
	Saved         *bool                                                  `json:"Saved,omitempty" xml:"Saved,omitempty"`
	SessionConfig *DescribeDataAgentSessionResponseBodyDataSessionConfig `json:"SessionConfig,omitempty" xml:"SessionConfig,omitempty" type:"Struct"`
	// example:
	//
	// 976*********p
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	// example:
	//
	// ca-e*******ckd
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// example:
	//
	// debug
	CustomAgentStage *string `json:"CustomAgentStage,omitempty" xml:"CustomAgentStage,omitempty"`
	// example:
	//
	// True
	EnableSearch *bool   `json:"EnableSearch,omitempty" xml:"EnableSearch,omitempty"`
	EncryptKey   *string `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
	EncryptType  *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// example:
	//
	// CHINESE
	Language     *string   `json:"Language,omitempty" xml:"Language,omitempty"`
	McpServerIds []*string `json:"McpServerIds,omitempty" xml:"McpServerIds,omitempty" type:"Repeated"`
	// example:
	//
	// ANALYSIS
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ReportPageWidth *int64  `json:"ReportPageWidth,omitempty" xml:"ReportPageWidth,omitempty"`
	ReportWaterMark *string `json:"ReportWaterMark,omitempty" xml:"ReportWaterMark,omitempty"`
	UserOssBucket   *string `json:"UserOssBucket,omitempty" xml:"UserOssBucket,omitempty"`
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

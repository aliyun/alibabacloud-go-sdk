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
	Data *CreateDataAgentSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 1CB***********3F1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// Agent Id
	//
	// example:
	//
	// cu0cs*******mf
	AgentId     *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentStatus *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// example:
	//
	// 1765262307992
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// f-8*******01m
	File *string `json:"File,omitempty" xml:"File,omitempty"`
	// example:
	//
	// false
	Saved         *bool                                                `json:"Saved,omitempty" xml:"Saved,omitempty"`
	SessionConfig *CreateDataAgentSessionResponseBodyDataSessionConfig `json:"SessionConfig,omitempty" xml:"SessionConfig,omitempty" type:"Struct"`
	// example:
	//
	// 976*********p
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
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

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) GetLanguage() *string {
	return s.Language
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) GetMode() *string {
	return s.Mode
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

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) SetLanguage(v string) *CreateDataAgentSessionResponseBodyDataSessionConfig {
	s.Language = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) SetMode(v string) *CreateDataAgentSessionResponseBodyDataSessionConfig {
	s.Mode = &v
	return s
}

func (s *CreateDataAgentSessionResponseBodyDataSessionConfig) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAIAgentTemplateConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAvatarChat3D(v *AIAgentTemplateConfigAvatarChat3D) *AIAgentTemplateConfig
	GetAvatarChat3D() *AIAgentTemplateConfigAvatarChat3D
	SetVisionChat(v *AIAgentTemplateConfigVisionChat) *AIAgentTemplateConfig
	GetVisionChat() *AIAgentTemplateConfigVisionChat
	SetVoiceChat(v *AIAgentTemplateConfigVoiceChat) *AIAgentTemplateConfig
	GetVoiceChat() *AIAgentTemplateConfigVoiceChat
}

type AIAgentTemplateConfig struct {
	AvatarChat3D *AIAgentTemplateConfigAvatarChat3D `json:"AvatarChat3D,omitempty" xml:"AvatarChat3D,omitempty" type:"Struct"`
	VisionChat   *AIAgentTemplateConfigVisionChat   `json:"VisionChat,omitempty" xml:"VisionChat,omitempty" type:"Struct"`
	VoiceChat    *AIAgentTemplateConfigVoiceChat    `json:"VoiceChat,omitempty" xml:"VoiceChat,omitempty" type:"Struct"`
}

func (s AIAgentTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentTemplateConfig) GoString() string {
	return s.String()
}

func (s *AIAgentTemplateConfig) GetAvatarChat3D() *AIAgentTemplateConfigAvatarChat3D {
	return s.AvatarChat3D
}

func (s *AIAgentTemplateConfig) GetVisionChat() *AIAgentTemplateConfigVisionChat {
	return s.VisionChat
}

func (s *AIAgentTemplateConfig) GetVoiceChat() *AIAgentTemplateConfigVoiceChat {
	return s.VoiceChat
}

func (s *AIAgentTemplateConfig) SetAvatarChat3D(v *AIAgentTemplateConfigAvatarChat3D) *AIAgentTemplateConfig {
	s.AvatarChat3D = v
	return s
}

func (s *AIAgentTemplateConfig) SetVisionChat(v *AIAgentTemplateConfigVisionChat) *AIAgentTemplateConfig {
	s.VisionChat = v
	return s
}

func (s *AIAgentTemplateConfig) SetVoiceChat(v *AIAgentTemplateConfigVoiceChat) *AIAgentTemplateConfig {
	s.VoiceChat = v
	return s
}

func (s *AIAgentTemplateConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentTemplateConfigAvatarChat3D struct {
	AsrHotWords              []*string                                      `json:"AsrHotWords,omitempty" xml:"AsrHotWords,omitempty" type:"Repeated"`
	AsrLanguageId            *string                                        `json:"AsrLanguageId,omitempty" xml:"AsrLanguageId,omitempty"`
	AsrMaxSilence            *int32                                         `json:"AsrMaxSilence,omitempty" xml:"AsrMaxSilence,omitempty"`
	AvatarId                 *string                                        `json:"AvatarId,omitempty" xml:"AvatarId,omitempty"`
	BailianAppParams         *string                                        `json:"BailianAppParams,omitempty" xml:"BailianAppParams,omitempty"`
	CharBreak                *bool                                          `json:"CharBreak,omitempty" xml:"CharBreak,omitempty"`
	EnableIntelligentSegment *bool                                          `json:"EnableIntelligentSegment,omitempty" xml:"EnableIntelligentSegment,omitempty"`
	EnablePushToTalk         *bool                                          `json:"EnablePushToTalk,omitempty" xml:"EnablePushToTalk,omitempty"`
	EnableVoiceInterrupt     *bool                                          `json:"EnableVoiceInterrupt,omitempty" xml:"EnableVoiceInterrupt,omitempty"`
	GracefulShutdown         *bool                                          `json:"GracefulShutdown,omitempty" xml:"GracefulShutdown,omitempty"`
	Greeting                 *string                                        `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	InterruptWords           []*string                                      `json:"InterruptWords,omitempty" xml:"InterruptWords,omitempty" type:"Repeated"`
	LlmHistory               []*AIAgentTemplateConfigAvatarChat3DLlmHistory `json:"LlmHistory,omitempty" xml:"LlmHistory,omitempty" type:"Repeated"`
	LlmHistoryLimit          *int32                                         `json:"LlmHistoryLimit,omitempty" xml:"LlmHistoryLimit,omitempty"`
	LlmSystemPrompt          *string                                        `json:"LlmSystemPrompt,omitempty" xml:"LlmSystemPrompt,omitempty"`
	MaxIdleTime              *int32                                         `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	UseVoiceprint            *bool                                          `json:"UseVoiceprint,omitempty" xml:"UseVoiceprint,omitempty"`
	UserOfflineTimeout       *int32                                         `json:"UserOfflineTimeout,omitempty" xml:"UserOfflineTimeout,omitempty"`
	UserOnlineTimeout        *int32                                         `json:"UserOnlineTimeout,omitempty" xml:"UserOnlineTimeout,omitempty"`
	VadLevel                 *int32                                         `json:"VadLevel,omitempty" xml:"VadLevel,omitempty"`
	VoiceId                  *string                                        `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	VoiceIdList              []*string                                      `json:"VoiceIdList,omitempty" xml:"VoiceIdList,omitempty" type:"Repeated"`
	VoiceprintId             *string                                        `json:"VoiceprintId,omitempty" xml:"VoiceprintId,omitempty"`
	Volume                   *int64                                         `json:"Volume,omitempty" xml:"Volume,omitempty"`
	WakeUpQuery              *string                                        `json:"WakeUpQuery,omitempty" xml:"WakeUpQuery,omitempty"`
	WorkflowOverrideParams   *string                                        `json:"WorkflowOverrideParams,omitempty" xml:"WorkflowOverrideParams,omitempty"`
}

func (s AIAgentTemplateConfigAvatarChat3D) String() string {
	return dara.Prettify(s)
}

func (s AIAgentTemplateConfigAvatarChat3D) GoString() string {
	return s.String()
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetAsrHotWords() []*string {
	return s.AsrHotWords
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetAsrLanguageId() *string {
	return s.AsrLanguageId
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetAsrMaxSilence() *int32 {
	return s.AsrMaxSilence
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetAvatarId() *string {
	return s.AvatarId
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetBailianAppParams() *string {
	return s.BailianAppParams
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetCharBreak() *bool {
	return s.CharBreak
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetEnableIntelligentSegment() *bool {
	return s.EnableIntelligentSegment
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetEnablePushToTalk() *bool {
	return s.EnablePushToTalk
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetEnableVoiceInterrupt() *bool {
	return s.EnableVoiceInterrupt
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetGracefulShutdown() *bool {
	return s.GracefulShutdown
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetGreeting() *string {
	return s.Greeting
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetInterruptWords() []*string {
	return s.InterruptWords
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetLlmHistory() []*AIAgentTemplateConfigAvatarChat3DLlmHistory {
	return s.LlmHistory
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetLlmHistoryLimit() *int32 {
	return s.LlmHistoryLimit
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetLlmSystemPrompt() *string {
	return s.LlmSystemPrompt
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetMaxIdleTime() *int32 {
	return s.MaxIdleTime
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetUseVoiceprint() *bool {
	return s.UseVoiceprint
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetUserOfflineTimeout() *int32 {
	return s.UserOfflineTimeout
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetUserOnlineTimeout() *int32 {
	return s.UserOnlineTimeout
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetVadLevel() *int32 {
	return s.VadLevel
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetVoiceId() *string {
	return s.VoiceId
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetVoiceIdList() []*string {
	return s.VoiceIdList
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetVoiceprintId() *string {
	return s.VoiceprintId
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetVolume() *int64 {
	return s.Volume
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetWakeUpQuery() *string {
	return s.WakeUpQuery
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetWorkflowOverrideParams() *string {
	return s.WorkflowOverrideParams
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetAsrHotWords(v []*string) *AIAgentTemplateConfigAvatarChat3D {
	s.AsrHotWords = v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetAsrLanguageId(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.AsrLanguageId = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetAsrMaxSilence(v int32) *AIAgentTemplateConfigAvatarChat3D {
	s.AsrMaxSilence = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetAvatarId(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.AvatarId = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetBailianAppParams(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.BailianAppParams = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetCharBreak(v bool) *AIAgentTemplateConfigAvatarChat3D {
	s.CharBreak = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetEnableIntelligentSegment(v bool) *AIAgentTemplateConfigAvatarChat3D {
	s.EnableIntelligentSegment = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetEnablePushToTalk(v bool) *AIAgentTemplateConfigAvatarChat3D {
	s.EnablePushToTalk = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetEnableVoiceInterrupt(v bool) *AIAgentTemplateConfigAvatarChat3D {
	s.EnableVoiceInterrupt = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetGracefulShutdown(v bool) *AIAgentTemplateConfigAvatarChat3D {
	s.GracefulShutdown = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetGreeting(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.Greeting = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetInterruptWords(v []*string) *AIAgentTemplateConfigAvatarChat3D {
	s.InterruptWords = v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetLlmHistory(v []*AIAgentTemplateConfigAvatarChat3DLlmHistory) *AIAgentTemplateConfigAvatarChat3D {
	s.LlmHistory = v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetLlmHistoryLimit(v int32) *AIAgentTemplateConfigAvatarChat3D {
	s.LlmHistoryLimit = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetLlmSystemPrompt(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.LlmSystemPrompt = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetMaxIdleTime(v int32) *AIAgentTemplateConfigAvatarChat3D {
	s.MaxIdleTime = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetUseVoiceprint(v bool) *AIAgentTemplateConfigAvatarChat3D {
	s.UseVoiceprint = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetUserOfflineTimeout(v int32) *AIAgentTemplateConfigAvatarChat3D {
	s.UserOfflineTimeout = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetUserOnlineTimeout(v int32) *AIAgentTemplateConfigAvatarChat3D {
	s.UserOnlineTimeout = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetVadLevel(v int32) *AIAgentTemplateConfigAvatarChat3D {
	s.VadLevel = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetVoiceId(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.VoiceId = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetVoiceIdList(v []*string) *AIAgentTemplateConfigAvatarChat3D {
	s.VoiceIdList = v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetVoiceprintId(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.VoiceprintId = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetVolume(v int64) *AIAgentTemplateConfigAvatarChat3D {
	s.Volume = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetWakeUpQuery(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.WakeUpQuery = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetWorkflowOverrideParams(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.WorkflowOverrideParams = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) Validate() error {
	return dara.Validate(s)
}

type AIAgentTemplateConfigAvatarChat3DLlmHistory struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Role    *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s AIAgentTemplateConfigAvatarChat3DLlmHistory) String() string {
	return dara.Prettify(s)
}

func (s AIAgentTemplateConfigAvatarChat3DLlmHistory) GoString() string {
	return s.String()
}

func (s *AIAgentTemplateConfigAvatarChat3DLlmHistory) GetContent() *string {
	return s.Content
}

func (s *AIAgentTemplateConfigAvatarChat3DLlmHistory) GetRole() *string {
	return s.Role
}

func (s *AIAgentTemplateConfigAvatarChat3DLlmHistory) SetContent(v string) *AIAgentTemplateConfigAvatarChat3DLlmHistory {
	s.Content = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3DLlmHistory) SetRole(v string) *AIAgentTemplateConfigAvatarChat3DLlmHistory {
	s.Role = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3DLlmHistory) Validate() error {
	return dara.Validate(s)
}

type AIAgentTemplateConfigVisionChat struct {
	AsrHotWords              []*string                                    `json:"AsrHotWords,omitempty" xml:"AsrHotWords,omitempty" type:"Repeated"`
	AsrLanguageId            *string                                      `json:"AsrLanguageId,omitempty" xml:"AsrLanguageId,omitempty"`
	AsrMaxSilence            *int32                                       `json:"AsrMaxSilence,omitempty" xml:"AsrMaxSilence,omitempty"`
	BailianAppParams         *string                                      `json:"BailianAppParams,omitempty" xml:"BailianAppParams,omitempty"`
	CharBreak                *bool                                        `json:"CharBreak,omitempty" xml:"CharBreak,omitempty"`
	EnableIntelligentSegment *bool                                        `json:"EnableIntelligentSegment,omitempty" xml:"EnableIntelligentSegment,omitempty"`
	EnablePushToTalk         *bool                                        `json:"EnablePushToTalk,omitempty" xml:"EnablePushToTalk,omitempty"`
	EnableVoiceInterrupt     *bool                                        `json:"EnableVoiceInterrupt,omitempty" xml:"EnableVoiceInterrupt,omitempty"`
	GracefulShutdown         *bool                                        `json:"GracefulShutdown,omitempty" xml:"GracefulShutdown,omitempty"`
	Greeting                 *string                                      `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	InterruptWords           []*string                                    `json:"InterruptWords,omitempty" xml:"InterruptWords,omitempty" type:"Repeated"`
	LlmHistory               []*AIAgentTemplateConfigVisionChatLlmHistory `json:"LlmHistory,omitempty" xml:"LlmHistory,omitempty" type:"Repeated"`
	LlmHistoryLimit          *int32                                       `json:"LlmHistoryLimit,omitempty" xml:"LlmHistoryLimit,omitempty"`
	LlmSystemPrompt          *string                                      `json:"LlmSystemPrompt,omitempty" xml:"LlmSystemPrompt,omitempty"`
	MaxIdleTime              *int32                                       `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	UseVoiceprint            *bool                                        `json:"UseVoiceprint,omitempty" xml:"UseVoiceprint,omitempty"`
	UserOfflineTimeout       *int32                                       `json:"UserOfflineTimeout,omitempty" xml:"UserOfflineTimeout,omitempty"`
	UserOnlineTimeout        *int32                                       `json:"UserOnlineTimeout,omitempty" xml:"UserOnlineTimeout,omitempty"`
	VadLevel                 *int32                                       `json:"VadLevel,omitempty" xml:"VadLevel,omitempty"`
	VoiceId                  *string                                      `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	VoiceIdList              []*string                                    `json:"VoiceIdList,omitempty" xml:"VoiceIdList,omitempty" type:"Repeated"`
	VoiceprintId             *string                                      `json:"VoiceprintId,omitempty" xml:"VoiceprintId,omitempty"`
	Volume                   *int64                                       `json:"Volume,omitempty" xml:"Volume,omitempty"`
	WakeUpQuery              *string                                      `json:"WakeUpQuery,omitempty" xml:"WakeUpQuery,omitempty"`
	WorkflowOverrideParams   *string                                      `json:"WorkflowOverrideParams,omitempty" xml:"WorkflowOverrideParams,omitempty"`
}

func (s AIAgentTemplateConfigVisionChat) String() string {
	return dara.Prettify(s)
}

func (s AIAgentTemplateConfigVisionChat) GoString() string {
	return s.String()
}

func (s *AIAgentTemplateConfigVisionChat) GetAsrHotWords() []*string {
	return s.AsrHotWords
}

func (s *AIAgentTemplateConfigVisionChat) GetAsrLanguageId() *string {
	return s.AsrLanguageId
}

func (s *AIAgentTemplateConfigVisionChat) GetAsrMaxSilence() *int32 {
	return s.AsrMaxSilence
}

func (s *AIAgentTemplateConfigVisionChat) GetBailianAppParams() *string {
	return s.BailianAppParams
}

func (s *AIAgentTemplateConfigVisionChat) GetCharBreak() *bool {
	return s.CharBreak
}

func (s *AIAgentTemplateConfigVisionChat) GetEnableIntelligentSegment() *bool {
	return s.EnableIntelligentSegment
}

func (s *AIAgentTemplateConfigVisionChat) GetEnablePushToTalk() *bool {
	return s.EnablePushToTalk
}

func (s *AIAgentTemplateConfigVisionChat) GetEnableVoiceInterrupt() *bool {
	return s.EnableVoiceInterrupt
}

func (s *AIAgentTemplateConfigVisionChat) GetGracefulShutdown() *bool {
	return s.GracefulShutdown
}

func (s *AIAgentTemplateConfigVisionChat) GetGreeting() *string {
	return s.Greeting
}

func (s *AIAgentTemplateConfigVisionChat) GetInterruptWords() []*string {
	return s.InterruptWords
}

func (s *AIAgentTemplateConfigVisionChat) GetLlmHistory() []*AIAgentTemplateConfigVisionChatLlmHistory {
	return s.LlmHistory
}

func (s *AIAgentTemplateConfigVisionChat) GetLlmHistoryLimit() *int32 {
	return s.LlmHistoryLimit
}

func (s *AIAgentTemplateConfigVisionChat) GetLlmSystemPrompt() *string {
	return s.LlmSystemPrompt
}

func (s *AIAgentTemplateConfigVisionChat) GetMaxIdleTime() *int32 {
	return s.MaxIdleTime
}

func (s *AIAgentTemplateConfigVisionChat) GetUseVoiceprint() *bool {
	return s.UseVoiceprint
}

func (s *AIAgentTemplateConfigVisionChat) GetUserOfflineTimeout() *int32 {
	return s.UserOfflineTimeout
}

func (s *AIAgentTemplateConfigVisionChat) GetUserOnlineTimeout() *int32 {
	return s.UserOnlineTimeout
}

func (s *AIAgentTemplateConfigVisionChat) GetVadLevel() *int32 {
	return s.VadLevel
}

func (s *AIAgentTemplateConfigVisionChat) GetVoiceId() *string {
	return s.VoiceId
}

func (s *AIAgentTemplateConfigVisionChat) GetVoiceIdList() []*string {
	return s.VoiceIdList
}

func (s *AIAgentTemplateConfigVisionChat) GetVoiceprintId() *string {
	return s.VoiceprintId
}

func (s *AIAgentTemplateConfigVisionChat) GetVolume() *int64 {
	return s.Volume
}

func (s *AIAgentTemplateConfigVisionChat) GetWakeUpQuery() *string {
	return s.WakeUpQuery
}

func (s *AIAgentTemplateConfigVisionChat) GetWorkflowOverrideParams() *string {
	return s.WorkflowOverrideParams
}

func (s *AIAgentTemplateConfigVisionChat) SetAsrHotWords(v []*string) *AIAgentTemplateConfigVisionChat {
	s.AsrHotWords = v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetAsrLanguageId(v string) *AIAgentTemplateConfigVisionChat {
	s.AsrLanguageId = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetAsrMaxSilence(v int32) *AIAgentTemplateConfigVisionChat {
	s.AsrMaxSilence = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetBailianAppParams(v string) *AIAgentTemplateConfigVisionChat {
	s.BailianAppParams = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetCharBreak(v bool) *AIAgentTemplateConfigVisionChat {
	s.CharBreak = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetEnableIntelligentSegment(v bool) *AIAgentTemplateConfigVisionChat {
	s.EnableIntelligentSegment = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetEnablePushToTalk(v bool) *AIAgentTemplateConfigVisionChat {
	s.EnablePushToTalk = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetEnableVoiceInterrupt(v bool) *AIAgentTemplateConfigVisionChat {
	s.EnableVoiceInterrupt = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetGracefulShutdown(v bool) *AIAgentTemplateConfigVisionChat {
	s.GracefulShutdown = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetGreeting(v string) *AIAgentTemplateConfigVisionChat {
	s.Greeting = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetInterruptWords(v []*string) *AIAgentTemplateConfigVisionChat {
	s.InterruptWords = v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetLlmHistory(v []*AIAgentTemplateConfigVisionChatLlmHistory) *AIAgentTemplateConfigVisionChat {
	s.LlmHistory = v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetLlmHistoryLimit(v int32) *AIAgentTemplateConfigVisionChat {
	s.LlmHistoryLimit = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetLlmSystemPrompt(v string) *AIAgentTemplateConfigVisionChat {
	s.LlmSystemPrompt = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetMaxIdleTime(v int32) *AIAgentTemplateConfigVisionChat {
	s.MaxIdleTime = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetUseVoiceprint(v bool) *AIAgentTemplateConfigVisionChat {
	s.UseVoiceprint = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetUserOfflineTimeout(v int32) *AIAgentTemplateConfigVisionChat {
	s.UserOfflineTimeout = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetUserOnlineTimeout(v int32) *AIAgentTemplateConfigVisionChat {
	s.UserOnlineTimeout = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetVadLevel(v int32) *AIAgentTemplateConfigVisionChat {
	s.VadLevel = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetVoiceId(v string) *AIAgentTemplateConfigVisionChat {
	s.VoiceId = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetVoiceIdList(v []*string) *AIAgentTemplateConfigVisionChat {
	s.VoiceIdList = v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetVoiceprintId(v string) *AIAgentTemplateConfigVisionChat {
	s.VoiceprintId = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetVolume(v int64) *AIAgentTemplateConfigVisionChat {
	s.Volume = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetWakeUpQuery(v string) *AIAgentTemplateConfigVisionChat {
	s.WakeUpQuery = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetWorkflowOverrideParams(v string) *AIAgentTemplateConfigVisionChat {
	s.WorkflowOverrideParams = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) Validate() error {
	return dara.Validate(s)
}

type AIAgentTemplateConfigVisionChatLlmHistory struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Role    *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s AIAgentTemplateConfigVisionChatLlmHistory) String() string {
	return dara.Prettify(s)
}

func (s AIAgentTemplateConfigVisionChatLlmHistory) GoString() string {
	return s.String()
}

func (s *AIAgentTemplateConfigVisionChatLlmHistory) GetContent() *string {
	return s.Content
}

func (s *AIAgentTemplateConfigVisionChatLlmHistory) GetRole() *string {
	return s.Role
}

func (s *AIAgentTemplateConfigVisionChatLlmHistory) SetContent(v string) *AIAgentTemplateConfigVisionChatLlmHistory {
	s.Content = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChatLlmHistory) SetRole(v string) *AIAgentTemplateConfigVisionChatLlmHistory {
	s.Role = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChatLlmHistory) Validate() error {
	return dara.Validate(s)
}

type AIAgentTemplateConfigVoiceChat struct {
	AsrHotWords              []*string                                   `json:"AsrHotWords,omitempty" xml:"AsrHotWords,omitempty" type:"Repeated"`
	AsrLanguageId            *string                                     `json:"AsrLanguageId,omitempty" xml:"AsrLanguageId,omitempty"`
	AsrMaxSilence            *int32                                      `json:"AsrMaxSilence,omitempty" xml:"AsrMaxSilence,omitempty"`
	AvatarUrl                *string                                     `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	AvatarUrlType            *string                                     `json:"AvatarUrlType,omitempty" xml:"AvatarUrlType,omitempty"`
	BailianAppParams         *string                                     `json:"BailianAppParams,omitempty" xml:"BailianAppParams,omitempty"`
	CharBreak                *bool                                       `json:"CharBreak,omitempty" xml:"CharBreak,omitempty"`
	EnableIntelligentSegment *bool                                       `json:"EnableIntelligentSegment,omitempty" xml:"EnableIntelligentSegment,omitempty"`
	EnablePushToTalk         *bool                                       `json:"EnablePushToTalk,omitempty" xml:"EnablePushToTalk,omitempty"`
	EnableVoiceInterrupt     *bool                                       `json:"EnableVoiceInterrupt,omitempty" xml:"EnableVoiceInterrupt,omitempty"`
	GracefulShutdown         *bool                                       `json:"GracefulShutdown,omitempty" xml:"GracefulShutdown,omitempty"`
	Greeting                 *string                                     `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	InterruptWords           []*string                                   `json:"InterruptWords,omitempty" xml:"InterruptWords,omitempty" type:"Repeated"`
	LlmHistory               []*AIAgentTemplateConfigVoiceChatLlmHistory `json:"LlmHistory,omitempty" xml:"LlmHistory,omitempty" type:"Repeated"`
	LlmHistoryLimit          *int32                                      `json:"LlmHistoryLimit,omitempty" xml:"LlmHistoryLimit,omitempty"`
	LlmSystemPrompt          *string                                     `json:"LlmSystemPrompt,omitempty" xml:"LlmSystemPrompt,omitempty"`
	MaxIdleTime              *int32                                      `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	UseVoiceprint            *bool                                       `json:"UseVoiceprint,omitempty" xml:"UseVoiceprint,omitempty"`
	UserOfflineTimeout       *int32                                      `json:"UserOfflineTimeout,omitempty" xml:"UserOfflineTimeout,omitempty"`
	UserOnlineTimeout        *int32                                      `json:"UserOnlineTimeout,omitempty" xml:"UserOnlineTimeout,omitempty"`
	VadLevel                 *int32                                      `json:"VadLevel,omitempty" xml:"VadLevel,omitempty"`
	VoiceId                  *string                                     `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	VoiceIdList              []*string                                   `json:"VoiceIdList,omitempty" xml:"VoiceIdList,omitempty" type:"Repeated"`
	VoiceprintId             *string                                     `json:"VoiceprintId,omitempty" xml:"VoiceprintId,omitempty"`
	Volume                   *int64                                      `json:"Volume,omitempty" xml:"Volume,omitempty"`
	WakeUpQuery              *string                                     `json:"WakeUpQuery,omitempty" xml:"WakeUpQuery,omitempty"`
	WorkflowOverrideParams   *string                                     `json:"WorkflowOverrideParams,omitempty" xml:"WorkflowOverrideParams,omitempty"`
}

func (s AIAgentTemplateConfigVoiceChat) String() string {
	return dara.Prettify(s)
}

func (s AIAgentTemplateConfigVoiceChat) GoString() string {
	return s.String()
}

func (s *AIAgentTemplateConfigVoiceChat) GetAsrHotWords() []*string {
	return s.AsrHotWords
}

func (s *AIAgentTemplateConfigVoiceChat) GetAsrLanguageId() *string {
	return s.AsrLanguageId
}

func (s *AIAgentTemplateConfigVoiceChat) GetAsrMaxSilence() *int32 {
	return s.AsrMaxSilence
}

func (s *AIAgentTemplateConfigVoiceChat) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *AIAgentTemplateConfigVoiceChat) GetAvatarUrlType() *string {
	return s.AvatarUrlType
}

func (s *AIAgentTemplateConfigVoiceChat) GetBailianAppParams() *string {
	return s.BailianAppParams
}

func (s *AIAgentTemplateConfigVoiceChat) GetCharBreak() *bool {
	return s.CharBreak
}

func (s *AIAgentTemplateConfigVoiceChat) GetEnableIntelligentSegment() *bool {
	return s.EnableIntelligentSegment
}

func (s *AIAgentTemplateConfigVoiceChat) GetEnablePushToTalk() *bool {
	return s.EnablePushToTalk
}

func (s *AIAgentTemplateConfigVoiceChat) GetEnableVoiceInterrupt() *bool {
	return s.EnableVoiceInterrupt
}

func (s *AIAgentTemplateConfigVoiceChat) GetGracefulShutdown() *bool {
	return s.GracefulShutdown
}

func (s *AIAgentTemplateConfigVoiceChat) GetGreeting() *string {
	return s.Greeting
}

func (s *AIAgentTemplateConfigVoiceChat) GetInterruptWords() []*string {
	return s.InterruptWords
}

func (s *AIAgentTemplateConfigVoiceChat) GetLlmHistory() []*AIAgentTemplateConfigVoiceChatLlmHistory {
	return s.LlmHistory
}

func (s *AIAgentTemplateConfigVoiceChat) GetLlmHistoryLimit() *int32 {
	return s.LlmHistoryLimit
}

func (s *AIAgentTemplateConfigVoiceChat) GetLlmSystemPrompt() *string {
	return s.LlmSystemPrompt
}

func (s *AIAgentTemplateConfigVoiceChat) GetMaxIdleTime() *int32 {
	return s.MaxIdleTime
}

func (s *AIAgentTemplateConfigVoiceChat) GetUseVoiceprint() *bool {
	return s.UseVoiceprint
}

func (s *AIAgentTemplateConfigVoiceChat) GetUserOfflineTimeout() *int32 {
	return s.UserOfflineTimeout
}

func (s *AIAgentTemplateConfigVoiceChat) GetUserOnlineTimeout() *int32 {
	return s.UserOnlineTimeout
}

func (s *AIAgentTemplateConfigVoiceChat) GetVadLevel() *int32 {
	return s.VadLevel
}

func (s *AIAgentTemplateConfigVoiceChat) GetVoiceId() *string {
	return s.VoiceId
}

func (s *AIAgentTemplateConfigVoiceChat) GetVoiceIdList() []*string {
	return s.VoiceIdList
}

func (s *AIAgentTemplateConfigVoiceChat) GetVoiceprintId() *string {
	return s.VoiceprintId
}

func (s *AIAgentTemplateConfigVoiceChat) GetVolume() *int64 {
	return s.Volume
}

func (s *AIAgentTemplateConfigVoiceChat) GetWakeUpQuery() *string {
	return s.WakeUpQuery
}

func (s *AIAgentTemplateConfigVoiceChat) GetWorkflowOverrideParams() *string {
	return s.WorkflowOverrideParams
}

func (s *AIAgentTemplateConfigVoiceChat) SetAsrHotWords(v []*string) *AIAgentTemplateConfigVoiceChat {
	s.AsrHotWords = v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetAsrLanguageId(v string) *AIAgentTemplateConfigVoiceChat {
	s.AsrLanguageId = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetAsrMaxSilence(v int32) *AIAgentTemplateConfigVoiceChat {
	s.AsrMaxSilence = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetAvatarUrl(v string) *AIAgentTemplateConfigVoiceChat {
	s.AvatarUrl = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetAvatarUrlType(v string) *AIAgentTemplateConfigVoiceChat {
	s.AvatarUrlType = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetBailianAppParams(v string) *AIAgentTemplateConfigVoiceChat {
	s.BailianAppParams = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetCharBreak(v bool) *AIAgentTemplateConfigVoiceChat {
	s.CharBreak = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetEnableIntelligentSegment(v bool) *AIAgentTemplateConfigVoiceChat {
	s.EnableIntelligentSegment = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetEnablePushToTalk(v bool) *AIAgentTemplateConfigVoiceChat {
	s.EnablePushToTalk = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetEnableVoiceInterrupt(v bool) *AIAgentTemplateConfigVoiceChat {
	s.EnableVoiceInterrupt = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetGracefulShutdown(v bool) *AIAgentTemplateConfigVoiceChat {
	s.GracefulShutdown = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetGreeting(v string) *AIAgentTemplateConfigVoiceChat {
	s.Greeting = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetInterruptWords(v []*string) *AIAgentTemplateConfigVoiceChat {
	s.InterruptWords = v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetLlmHistory(v []*AIAgentTemplateConfigVoiceChatLlmHistory) *AIAgentTemplateConfigVoiceChat {
	s.LlmHistory = v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetLlmHistoryLimit(v int32) *AIAgentTemplateConfigVoiceChat {
	s.LlmHistoryLimit = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetLlmSystemPrompt(v string) *AIAgentTemplateConfigVoiceChat {
	s.LlmSystemPrompt = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetMaxIdleTime(v int32) *AIAgentTemplateConfigVoiceChat {
	s.MaxIdleTime = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetUseVoiceprint(v bool) *AIAgentTemplateConfigVoiceChat {
	s.UseVoiceprint = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetUserOfflineTimeout(v int32) *AIAgentTemplateConfigVoiceChat {
	s.UserOfflineTimeout = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetUserOnlineTimeout(v int32) *AIAgentTemplateConfigVoiceChat {
	s.UserOnlineTimeout = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetVadLevel(v int32) *AIAgentTemplateConfigVoiceChat {
	s.VadLevel = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetVoiceId(v string) *AIAgentTemplateConfigVoiceChat {
	s.VoiceId = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetVoiceIdList(v []*string) *AIAgentTemplateConfigVoiceChat {
	s.VoiceIdList = v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetVoiceprintId(v string) *AIAgentTemplateConfigVoiceChat {
	s.VoiceprintId = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetVolume(v int64) *AIAgentTemplateConfigVoiceChat {
	s.Volume = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetWakeUpQuery(v string) *AIAgentTemplateConfigVoiceChat {
	s.WakeUpQuery = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetWorkflowOverrideParams(v string) *AIAgentTemplateConfigVoiceChat {
	s.WorkflowOverrideParams = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) Validate() error {
	return dara.Validate(s)
}

type AIAgentTemplateConfigVoiceChatLlmHistory struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Role    *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s AIAgentTemplateConfigVoiceChatLlmHistory) String() string {
	return dara.Prettify(s)
}

func (s AIAgentTemplateConfigVoiceChatLlmHistory) GoString() string {
	return s.String()
}

func (s *AIAgentTemplateConfigVoiceChatLlmHistory) GetContent() *string {
	return s.Content
}

func (s *AIAgentTemplateConfigVoiceChatLlmHistory) GetRole() *string {
	return s.Role
}

func (s *AIAgentTemplateConfigVoiceChatLlmHistory) SetContent(v string) *AIAgentTemplateConfigVoiceChatLlmHistory {
	s.Content = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChatLlmHistory) SetRole(v string) *AIAgentTemplateConfigVoiceChatLlmHistory {
	s.Role = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChatLlmHistory) Validate() error {
	return dara.Validate(s)
}

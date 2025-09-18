// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMmAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeMmAppResponseBody
	GetAppId() *string
	SetAppName(v string) *DescribeMmAppResponseBody
	GetAppName() *string
	SetConversationConfig(v *DescribeMmAppResponseBodyConversationConfig) *DescribeMmAppResponseBody
	GetConversationConfig() *DescribeMmAppResponseBodyConversationConfig
	SetCreateUserId(v string) *DescribeMmAppResponseBody
	GetCreateUserId() *string
	SetCreateUserName(v string) *DescribeMmAppResponseBody
	GetCreateUserName() *string
	SetGmtCreate(v string) *DescribeMmAppResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *DescribeMmAppResponseBody
	GetGmtModified() *string
	SetModelConfig(v *DescribeMmAppResponseBodyModelConfig) *DescribeMmAppResponseBody
	GetModelConfig() *DescribeMmAppResponseBodyModelConfig
	SetModifyUserId(v string) *DescribeMmAppResponseBody
	GetModifyUserId() *string
	SetModifyUserName(v string) *DescribeMmAppResponseBody
	GetModifyUserName() *string
	SetPrompt(v string) *DescribeMmAppResponseBody
	GetPrompt() *string
	SetPublishVersion(v int64) *DescribeMmAppResponseBody
	GetPublishVersion() *int64
	SetRequestId(v string) *DescribeMmAppResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeMmAppResponseBody
	GetStatus() *string
}

type DescribeMmAppResponseBody struct {
	// example:
	//
	// mm_xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 多模态应用xxxx
	AppName            *string                                      `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ConversationConfig *DescribeMmAppResponseBodyConversationConfig `json:"ConversationConfig,omitempty" xml:"ConversationConfig,omitempty" type:"Struct"`
	// example:
	//
	// 243433
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// xxxx
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// xxx
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// xxx
	GmtModified *string                               `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ModelConfig *DescribeMmAppResponseBodyModelConfig `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty" type:"Struct"`
	// example:
	//
	// 56673435
	ModifyUserId *string `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	// example:
	//
	// xxxx
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// 提示词xxxx
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 1
	PublishVersion *int64 `json:"PublishVersion,omitempty" xml:"PublishVersion,omitempty"`
	// example:
	//
	// xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMmAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMmAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMmAppResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *DescribeMmAppResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DescribeMmAppResponseBody) GetConversationConfig() *DescribeMmAppResponseBodyConversationConfig {
	return s.ConversationConfig
}

func (s *DescribeMmAppResponseBody) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *DescribeMmAppResponseBody) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *DescribeMmAppResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeMmAppResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeMmAppResponseBody) GetModelConfig() *DescribeMmAppResponseBodyModelConfig {
	return s.ModelConfig
}

func (s *DescribeMmAppResponseBody) GetModifyUserId() *string {
	return s.ModifyUserId
}

func (s *DescribeMmAppResponseBody) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *DescribeMmAppResponseBody) GetPrompt() *string {
	return s.Prompt
}

func (s *DescribeMmAppResponseBody) GetPublishVersion() *int64 {
	return s.PublishVersion
}

func (s *DescribeMmAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMmAppResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeMmAppResponseBody) SetAppId(v string) *DescribeMmAppResponseBody {
	s.AppId = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetAppName(v string) *DescribeMmAppResponseBody {
	s.AppName = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetConversationConfig(v *DescribeMmAppResponseBodyConversationConfig) *DescribeMmAppResponseBody {
	s.ConversationConfig = v
	return s
}

func (s *DescribeMmAppResponseBody) SetCreateUserId(v string) *DescribeMmAppResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetCreateUserName(v string) *DescribeMmAppResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetGmtCreate(v string) *DescribeMmAppResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetGmtModified(v string) *DescribeMmAppResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetModelConfig(v *DescribeMmAppResponseBodyModelConfig) *DescribeMmAppResponseBody {
	s.ModelConfig = v
	return s
}

func (s *DescribeMmAppResponseBody) SetModifyUserId(v string) *DescribeMmAppResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetModifyUserName(v string) *DescribeMmAppResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetPrompt(v string) *DescribeMmAppResponseBody {
	s.Prompt = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetPublishVersion(v int64) *DescribeMmAppResponseBody {
	s.PublishVersion = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetRequestId(v string) *DescribeMmAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMmAppResponseBody) SetStatus(v string) *DescribeMmAppResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeMmAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMmAppResponseBodyConversationConfig struct {
	// example:
	//
	// xxxx
	AsrModel *string `json:"AsrModel,omitempty" xml:"AsrModel,omitempty"`
	OpenAsr  *bool   `json:"OpenAsr,omitempty" xml:"OpenAsr,omitempty"`
	OpenTts  *bool   `json:"OpenTts,omitempty" xml:"OpenTts,omitempty"`
	// example:
	//
	// xxxx
	TtsModel *string `json:"TtsModel,omitempty" xml:"TtsModel,omitempty"`
}

func (s DescribeMmAppResponseBodyConversationConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeMmAppResponseBodyConversationConfig) GoString() string {
	return s.String()
}

func (s *DescribeMmAppResponseBodyConversationConfig) GetAsrModel() *string {
	return s.AsrModel
}

func (s *DescribeMmAppResponseBodyConversationConfig) GetOpenAsr() *bool {
	return s.OpenAsr
}

func (s *DescribeMmAppResponseBodyConversationConfig) GetOpenTts() *bool {
	return s.OpenTts
}

func (s *DescribeMmAppResponseBodyConversationConfig) GetTtsModel() *string {
	return s.TtsModel
}

func (s *DescribeMmAppResponseBodyConversationConfig) SetAsrModel(v string) *DescribeMmAppResponseBodyConversationConfig {
	s.AsrModel = &v
	return s
}

func (s *DescribeMmAppResponseBodyConversationConfig) SetOpenAsr(v bool) *DescribeMmAppResponseBodyConversationConfig {
	s.OpenAsr = &v
	return s
}

func (s *DescribeMmAppResponseBodyConversationConfig) SetOpenTts(v bool) *DescribeMmAppResponseBodyConversationConfig {
	s.OpenTts = &v
	return s
}

func (s *DescribeMmAppResponseBodyConversationConfig) SetTtsModel(v string) *DescribeMmAppResponseBodyConversationConfig {
	s.TtsModel = &v
	return s
}

func (s *DescribeMmAppResponseBodyConversationConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeMmAppResponseBodyModelConfig struct {
	// example:
	//
	// 5
	HistoryLimit *int32 `json:"HistoryLimit,omitempty" xml:"HistoryLimit,omitempty"`
	// example:
	//
	// MMH
	ModelType     *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	OpenWebSearch *bool   `json:"OpenWebSearch,omitempty" xml:"OpenWebSearch,omitempty"`
	// example:
	//
	// xxxx
	TextModal *string `json:"TextModal,omitempty" xml:"TextModal,omitempty"`
}

func (s DescribeMmAppResponseBodyModelConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeMmAppResponseBodyModelConfig) GoString() string {
	return s.String()
}

func (s *DescribeMmAppResponseBodyModelConfig) GetHistoryLimit() *int32 {
	return s.HistoryLimit
}

func (s *DescribeMmAppResponseBodyModelConfig) GetModelType() *string {
	return s.ModelType
}

func (s *DescribeMmAppResponseBodyModelConfig) GetOpenWebSearch() *bool {
	return s.OpenWebSearch
}

func (s *DescribeMmAppResponseBodyModelConfig) GetTextModal() *string {
	return s.TextModal
}

func (s *DescribeMmAppResponseBodyModelConfig) SetHistoryLimit(v int32) *DescribeMmAppResponseBodyModelConfig {
	s.HistoryLimit = &v
	return s
}

func (s *DescribeMmAppResponseBodyModelConfig) SetModelType(v string) *DescribeMmAppResponseBodyModelConfig {
	s.ModelType = &v
	return s
}

func (s *DescribeMmAppResponseBodyModelConfig) SetOpenWebSearch(v bool) *DescribeMmAppResponseBodyModelConfig {
	s.OpenWebSearch = &v
	return s
}

func (s *DescribeMmAppResponseBodyModelConfig) SetTextModal(v string) *DescribeMmAppResponseBodyModelConfig {
	s.TextModal = &v
	return s
}

func (s *DescribeMmAppResponseBodyModelConfig) Validate() error {
	return dara.Validate(s)
}

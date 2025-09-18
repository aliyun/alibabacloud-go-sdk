// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInfoList(v []*ListMmAppResponseBodyAppInfoList) *ListMmAppResponseBody
	GetAppInfoList() []*ListMmAppResponseBodyAppInfoList
	SetPageNumber(v int32) *ListMmAppResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMmAppResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListMmAppResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListMmAppResponseBody
	GetTotalCount() *int32
}

type ListMmAppResponseBody struct {
	AppInfoList []*ListMmAppResponseBodyAppInfoList `json:"AppInfoList,omitempty" xml:"AppInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMmAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMmAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmAppResponseBody) GetAppInfoList() []*ListMmAppResponseBodyAppInfoList {
	return s.AppInfoList
}

func (s *ListMmAppResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMmAppResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMmAppResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMmAppResponseBody) SetAppInfoList(v []*ListMmAppResponseBodyAppInfoList) *ListMmAppResponseBody {
	s.AppInfoList = v
	return s
}

func (s *ListMmAppResponseBody) SetPageNumber(v int32) *ListMmAppResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListMmAppResponseBody) SetPageSize(v int32) *ListMmAppResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMmAppResponseBody) SetRequestId(v string) *ListMmAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMmAppResponseBody) SetTotalCount(v int32) *ListMmAppResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMmAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMmAppResponseBodyAppInfoList struct {
	// example:
	//
	// mm_xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 多模态
	AppName            *string                                             `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ConversationConfig *ListMmAppResponseBodyAppInfoListConversationConfig `json:"ConversationConfig,omitempty" xml:"ConversationConfig,omitempty" type:"Struct"`
	// example:
	//
	// 454564
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// xxx
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// xxx
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// xxx
	GmtModified *string                                      `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ModelConfig *ListMmAppResponseBodyAppInfoListModelConfig `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty" type:"Struct"`
	// example:
	//
	// 56445
	ModifyUserId *string `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	// example:
	//
	// xxx
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// 提示词
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 1
	PublishVersion *int64 `json:"PublishVersion,omitempty" xml:"PublishVersion,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMmAppResponseBodyAppInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListMmAppResponseBodyAppInfoList) GoString() string {
	return s.String()
}

func (s *ListMmAppResponseBodyAppInfoList) GetAppId() *string {
	return s.AppId
}

func (s *ListMmAppResponseBodyAppInfoList) GetAppName() *string {
	return s.AppName
}

func (s *ListMmAppResponseBodyAppInfoList) GetConversationConfig() *ListMmAppResponseBodyAppInfoListConversationConfig {
	return s.ConversationConfig
}

func (s *ListMmAppResponseBodyAppInfoList) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *ListMmAppResponseBodyAppInfoList) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *ListMmAppResponseBodyAppInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMmAppResponseBodyAppInfoList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMmAppResponseBodyAppInfoList) GetModelConfig() *ListMmAppResponseBodyAppInfoListModelConfig {
	return s.ModelConfig
}

func (s *ListMmAppResponseBodyAppInfoList) GetModifyUserId() *string {
	return s.ModifyUserId
}

func (s *ListMmAppResponseBodyAppInfoList) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *ListMmAppResponseBodyAppInfoList) GetPrompt() *string {
	return s.Prompt
}

func (s *ListMmAppResponseBodyAppInfoList) GetPublishVersion() *int64 {
	return s.PublishVersion
}

func (s *ListMmAppResponseBodyAppInfoList) GetStatus() *int32 {
	return s.Status
}

func (s *ListMmAppResponseBodyAppInfoList) SetAppId(v string) *ListMmAppResponseBodyAppInfoList {
	s.AppId = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoList) SetAppName(v string) *ListMmAppResponseBodyAppInfoList {
	s.AppName = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoList) SetConversationConfig(v *ListMmAppResponseBodyAppInfoListConversationConfig) *ListMmAppResponseBodyAppInfoList {
	s.ConversationConfig = v
	return s
}

func (s *ListMmAppResponseBodyAppInfoList) SetCreateUserId(v string) *ListMmAppResponseBodyAppInfoList {
	s.CreateUserId = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoList) SetCreateUserName(v string) *ListMmAppResponseBodyAppInfoList {
	s.CreateUserName = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoList) SetGmtCreate(v string) *ListMmAppResponseBodyAppInfoList {
	s.GmtCreate = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoList) SetGmtModified(v string) *ListMmAppResponseBodyAppInfoList {
	s.GmtModified = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoList) SetModelConfig(v *ListMmAppResponseBodyAppInfoListModelConfig) *ListMmAppResponseBodyAppInfoList {
	s.ModelConfig = v
	return s
}

func (s *ListMmAppResponseBodyAppInfoList) SetModifyUserId(v string) *ListMmAppResponseBodyAppInfoList {
	s.ModifyUserId = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoList) SetModifyUserName(v string) *ListMmAppResponseBodyAppInfoList {
	s.ModifyUserName = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoList) SetPrompt(v string) *ListMmAppResponseBodyAppInfoList {
	s.Prompt = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoList) SetPublishVersion(v int64) *ListMmAppResponseBodyAppInfoList {
	s.PublishVersion = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoList) SetStatus(v int32) *ListMmAppResponseBodyAppInfoList {
	s.Status = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoList) Validate() error {
	return dara.Validate(s)
}

type ListMmAppResponseBodyAppInfoListConversationConfig struct {
	// example:
	//
	// xxx
	AsrModel *string `json:"AsrModel,omitempty" xml:"AsrModel,omitempty"`
	OpenAsr  *bool   `json:"OpenAsr,omitempty" xml:"OpenAsr,omitempty"`
	OpenTts  *bool   `json:"OpenTts,omitempty" xml:"OpenTts,omitempty"`
	// example:
	//
	// xxx
	TtsModel *string `json:"TtsModel,omitempty" xml:"TtsModel,omitempty"`
}

func (s ListMmAppResponseBodyAppInfoListConversationConfig) String() string {
	return dara.Prettify(s)
}

func (s ListMmAppResponseBodyAppInfoListConversationConfig) GoString() string {
	return s.String()
}

func (s *ListMmAppResponseBodyAppInfoListConversationConfig) GetAsrModel() *string {
	return s.AsrModel
}

func (s *ListMmAppResponseBodyAppInfoListConversationConfig) GetOpenAsr() *bool {
	return s.OpenAsr
}

func (s *ListMmAppResponseBodyAppInfoListConversationConfig) GetOpenTts() *bool {
	return s.OpenTts
}

func (s *ListMmAppResponseBodyAppInfoListConversationConfig) GetTtsModel() *string {
	return s.TtsModel
}

func (s *ListMmAppResponseBodyAppInfoListConversationConfig) SetAsrModel(v string) *ListMmAppResponseBodyAppInfoListConversationConfig {
	s.AsrModel = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoListConversationConfig) SetOpenAsr(v bool) *ListMmAppResponseBodyAppInfoListConversationConfig {
	s.OpenAsr = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoListConversationConfig) SetOpenTts(v bool) *ListMmAppResponseBodyAppInfoListConversationConfig {
	s.OpenTts = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoListConversationConfig) SetTtsModel(v string) *ListMmAppResponseBodyAppInfoListConversationConfig {
	s.TtsModel = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoListConversationConfig) Validate() error {
	return dara.Validate(s)
}

type ListMmAppResponseBodyAppInfoListModelConfig struct {
	// example:
	//
	// 5
	HistoryLimit *string `json:"HistoryLimit,omitempty" xml:"HistoryLimit,omitempty"`
	// example:
	//
	// MMH
	ModelType     *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	OpenWebSearch *bool   `json:"OpenWebSearch,omitempty" xml:"OpenWebSearch,omitempty"`
	// example:
	//
	// xxx
	TextModal *string `json:"TextModal,omitempty" xml:"TextModal,omitempty"`
}

func (s ListMmAppResponseBodyAppInfoListModelConfig) String() string {
	return dara.Prettify(s)
}

func (s ListMmAppResponseBodyAppInfoListModelConfig) GoString() string {
	return s.String()
}

func (s *ListMmAppResponseBodyAppInfoListModelConfig) GetHistoryLimit() *string {
	return s.HistoryLimit
}

func (s *ListMmAppResponseBodyAppInfoListModelConfig) GetModelType() *string {
	return s.ModelType
}

func (s *ListMmAppResponseBodyAppInfoListModelConfig) GetOpenWebSearch() *bool {
	return s.OpenWebSearch
}

func (s *ListMmAppResponseBodyAppInfoListModelConfig) GetTextModal() *string {
	return s.TextModal
}

func (s *ListMmAppResponseBodyAppInfoListModelConfig) SetHistoryLimit(v string) *ListMmAppResponseBodyAppInfoListModelConfig {
	s.HistoryLimit = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoListModelConfig) SetModelType(v string) *ListMmAppResponseBodyAppInfoListModelConfig {
	s.ModelType = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoListModelConfig) SetOpenWebSearch(v bool) *ListMmAppResponseBodyAppInfoListModelConfig {
	s.OpenWebSearch = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoListModelConfig) SetTextModal(v string) *ListMmAppResponseBodyAppInfoListModelConfig {
	s.TextModal = &v
	return s
}

func (s *ListMmAppResponseBodyAppInfoListModelConfig) Validate() error {
	return dara.Validate(s)
}

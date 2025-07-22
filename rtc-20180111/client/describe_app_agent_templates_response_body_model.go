// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppAgentTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAppAgentTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v []*DescribeAppAgentTemplatesResponseBodyTemplates) *DescribeAppAgentTemplatesResponseBody
	GetTemplates() []*DescribeAppAgentTemplatesResponseBodyTemplates
	SetTotalNum(v int64) *DescribeAppAgentTemplatesResponseBody
	GetTotalNum() *int64
	SetTotalPage(v int64) *DescribeAppAgentTemplatesResponseBody
	GetTotalPage() *int64
}

type DescribeAppAgentTemplatesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 30D41049-D02D-1C21-86AE-B3E5FD805C27
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates []*DescribeAppAgentTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeAppAgentTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppAgentTemplatesResponseBody) GetTemplates() []*DescribeAppAgentTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *DescribeAppAgentTemplatesResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *DescribeAppAgentTemplatesResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeAppAgentTemplatesResponseBody) SetRequestId(v string) *DescribeAppAgentTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBody) SetTemplates(v []*DescribeAppAgentTemplatesResponseBodyTemplates) *DescribeAppAgentTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBody) SetTotalNum(v int64) *DescribeAppAgentTemplatesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBody) SetTotalPage(v int64) *DescribeAppAgentTemplatesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAgentTemplatesResponseBodyTemplates struct {
	AsrConfig *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty" type:"Struct"`
	// example:
	//
	// 1
	ChatMode *int32 `json:"ChatMode,omitempty" xml:"ChatMode,omitempty"`
	// example:
	//
	// 2020-09-04T06:22:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 你好，机器人。
	Greeting *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	// example:
	//
	// wv7N****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	InterruptMode *int32                                                   `json:"InterruptMode,omitempty" xml:"InterruptMode,omitempty"`
	LlmConfig     *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig `json:"LlmConfig,omitempty" xml:"LlmConfig,omitempty" type:"Struct"`
	// example:
	//
	// 测试
	Name      *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	TtsConfig *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAppAgentTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetAsrConfig() *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig {
	return s.AsrConfig
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetChatMode() *int32 {
	return s.ChatMode
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetGreeting() *string {
	return s.Greeting
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetId() *string {
	return s.Id
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetInterruptMode() *int32 {
	return s.InterruptMode
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetLlmConfig() *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	return s.LlmConfig
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetName() *string {
	return s.Name
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetTtsConfig() *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	return s.TtsConfig
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetType() *int32 {
	return s.Type
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetAsrConfig(v *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.AsrConfig = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetChatMode(v int32) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.ChatMode = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetCreateTime(v string) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.CreateTime = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetGreeting(v string) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.Greeting = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetId(v string) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.Id = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetInterruptMode(v int32) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.InterruptMode = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetLlmConfig(v *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.LlmConfig = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetName(v string) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetTtsConfig(v *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.TtsConfig = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetType(v int32) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.Type = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig struct {
	// example:
	//
	// 300
	MaxSentenceSilence *int32  `json:"MaxSentenceSilence,omitempty" xml:"MaxSentenceSilence,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ecfadace11114cf08a7f07aceee798ad
	VocabularyId *string                                                               `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
	WordWeights  []*DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights `json:"WordWeights,omitempty" xml:"WordWeights,omitempty" type:"Repeated"`
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) GetMaxSentenceSilence() *int32 {
	return s.MaxSentenceSilence
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) GetName() *string {
	return s.Name
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) GetWordWeights() []*DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights {
	return s.WordWeights
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) SetMaxSentenceSilence(v int32) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig {
	s.MaxSentenceSilence = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) SetName(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig {
	s.Name = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) SetVocabularyId(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig {
	s.VocabularyId = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) SetWordWeights(v []*DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig {
	s.WordWeights = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1
	Weight *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Word   *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) GetLang() *string {
	return s.Lang
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) GetWord() *string {
	return s.Word
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) SetLang(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights {
	s.Lang = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) SetWeight(v int32) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights {
	s.Weight = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) SetWord(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights {
	s.Word = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig struct {
	// example:
	//
	// qW8GpBOdHK/pv9gdUSVLvQ==
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// 8
	HistoryDepth *int32 `json:"HistoryDepth,omitempty" xml:"HistoryDepth,omitempty"`
	// example:
	//
	// 1024
	MaxToken *int32 `json:"MaxToken,omitempty" xml:"MaxToken,omitempty"`
	// example:
	//
	// deepseek-r1
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// llm
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 0.8
	Temperature *float32 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// example:
	//
	// 0.8
	TopP *float32 `json:"TopP,omitempty" xml:"TopP,omitempty"`
	// example:
	//
	// https://test.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// aliyun
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetHistoryDepth() *int32 {
	return s.HistoryDepth
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetMaxToken() *int32 {
	return s.MaxToken
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetModel() *string {
	return s.Model
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetName() *string {
	return s.Name
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetPrompt() *string {
	return s.Prompt
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetTemperature() *float32 {
	return s.Temperature
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetTopP() *float32 {
	return s.TopP
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetUrl() *string {
	return s.Url
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetVendor() *string {
	return s.Vendor
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetApiKey(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.ApiKey = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetHistoryDepth(v int32) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.HistoryDepth = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetMaxToken(v int32) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.MaxToken = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetModel(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.Model = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetName(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.Name = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetPrompt(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.Prompt = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetTemperature(v float32) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.Temperature = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetTopP(v float32) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.TopP = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetUrl(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.Url = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetVendor(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.Vendor = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig struct {
	// example:
	//
	// N5448VFGI2mXJU8a/A03VQ==
	ApiKey         *string  `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	FilterBrackets []*int32 `json:"FilterBrackets,omitempty" xml:"FilterBrackets,omitempty" type:"Repeated"`
	// example:
	//
	// cosyvoice-v1
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// Tts
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.8
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	// example:
	//
	// 50
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// example:
	//
	// aliyun
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// example:
	//
	// longwan
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetFilterBrackets() []*int32 {
	return s.FilterBrackets
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetModel() *string {
	return s.Model
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetName() *string {
	return s.Name
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetPitch() *float32 {
	return s.Pitch
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetRate() *float32 {
	return s.Rate
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetVendor() *string {
	return s.Vendor
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetVoice() *string {
	return s.Voice
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetApiKey(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.ApiKey = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetFilterBrackets(v []*int32) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.FilterBrackets = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetModel(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.Model = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetName(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.Name = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetPitch(v float32) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.Pitch = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetRate(v float32) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.Rate = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetVendor(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.Vendor = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetVoice(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.Voice = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetVolume(v int32) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.Volume = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) Validate() error {
	return dara.Validate(s)
}

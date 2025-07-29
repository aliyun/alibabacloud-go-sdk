// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunEnterpriseVocAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAkProxy(v string) *RunEnterpriseVocAnalysisRequest
	GetAkProxy() *string
	SetApiKey(v string) *RunEnterpriseVocAnalysisRequest
	GetApiKey() *string
	SetContent(v string) *RunEnterpriseVocAnalysisRequest
	GetContent() *string
	SetExtraInfo(v string) *RunEnterpriseVocAnalysisRequest
	GetExtraInfo() *string
	SetFilterTags(v []*RunEnterpriseVocAnalysisRequestFilterTags) *RunEnterpriseVocAnalysisRequest
	GetFilterTags() []*RunEnterpriseVocAnalysisRequestFilterTags
	SetModelId(v string) *RunEnterpriseVocAnalysisRequest
	GetModelId() *string
	SetOutputFormat(v string) *RunEnterpriseVocAnalysisRequest
	GetOutputFormat() *string
	SetSourceTrace(v bool) *RunEnterpriseVocAnalysisRequest
	GetSourceTrace() *bool
	SetTags(v []*RunEnterpriseVocAnalysisRequestTags) *RunEnterpriseVocAnalysisRequest
	GetTags() []*RunEnterpriseVocAnalysisRequestTags
	SetTaskDescription(v string) *RunEnterpriseVocAnalysisRequest
	GetTaskDescription() *string
}

type RunEnterpriseVocAnalysisRequest struct {
	AkProxy *string `json:"akProxy,omitempty" xml:"akProxy,omitempty"`
	ApiKey  *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// 需要进行VOC分析的文本内容（content、contents、url、fileKey 四选一。优先级从小到大）
	//
	// example:
	//
	// 这是一段需要分析的文本内容
	Content   *string `json:"content,omitempty" xml:"content,omitempty"`
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// 过滤标签，用于筛选符合条件的内容。
	FilterTags []*RunEnterpriseVocAnalysisRequestFilterTags `json:"filterTags,omitempty" xml:"filterTags,omitempty" type:"Repeated"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// 指定返回结果的格式，支持json或text
	//
	// example:
	//
	// 按照如下格式输出：{"text1": "xxxx", "text2": "xxxx"}
	OutputFormat *string `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	SourceTrace  *bool   `json:"sourceTrace,omitempty" xml:"sourceTrace,omitempty"`
	// 业务标签体系，用于对文本内容进行分类和分析。
	Tags []*RunEnterpriseVocAnalysisRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 你是一名经验丰富的数据分析师，擅长从文本评论中提取结构化信息。你需要从用户评论列表中识别和提取出与以下四个维度相关的关键词和短语：
	//
	//
	//
	//             索引：输入评论JSON数组中的索引（从零开始）表示针对该条索引抽取的维度。
	//
	//             购买动机：描述用户购买产品的原因、需求或驱动力的关键词或短语。
	//
	//             未满足需求点：用户在使用产品过程中提到的未满足需求或问题的关键词或短语。
	//
	//             使用场景：用户提到的具体使用场景、使用方式或环境的关键词或短语。
	//
	//             正负面观点：明确表示用户对产品或服务的正面或负面看法的关键词或短语。
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
}

func (s RunEnterpriseVocAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s RunEnterpriseVocAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunEnterpriseVocAnalysisRequest) GetAkProxy() *string {
	return s.AkProxy
}

func (s *RunEnterpriseVocAnalysisRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *RunEnterpriseVocAnalysisRequest) GetContent() *string {
	return s.Content
}

func (s *RunEnterpriseVocAnalysisRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *RunEnterpriseVocAnalysisRequest) GetFilterTags() []*RunEnterpriseVocAnalysisRequestFilterTags {
	return s.FilterTags
}

func (s *RunEnterpriseVocAnalysisRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunEnterpriseVocAnalysisRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *RunEnterpriseVocAnalysisRequest) GetSourceTrace() *bool {
	return s.SourceTrace
}

func (s *RunEnterpriseVocAnalysisRequest) GetTags() []*RunEnterpriseVocAnalysisRequestTags {
	return s.Tags
}

func (s *RunEnterpriseVocAnalysisRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *RunEnterpriseVocAnalysisRequest) SetAkProxy(v string) *RunEnterpriseVocAnalysisRequest {
	s.AkProxy = &v
	return s
}

func (s *RunEnterpriseVocAnalysisRequest) SetApiKey(v string) *RunEnterpriseVocAnalysisRequest {
	s.ApiKey = &v
	return s
}

func (s *RunEnterpriseVocAnalysisRequest) SetContent(v string) *RunEnterpriseVocAnalysisRequest {
	s.Content = &v
	return s
}

func (s *RunEnterpriseVocAnalysisRequest) SetExtraInfo(v string) *RunEnterpriseVocAnalysisRequest {
	s.ExtraInfo = &v
	return s
}

func (s *RunEnterpriseVocAnalysisRequest) SetFilterTags(v []*RunEnterpriseVocAnalysisRequestFilterTags) *RunEnterpriseVocAnalysisRequest {
	s.FilterTags = v
	return s
}

func (s *RunEnterpriseVocAnalysisRequest) SetModelId(v string) *RunEnterpriseVocAnalysisRequest {
	s.ModelId = &v
	return s
}

func (s *RunEnterpriseVocAnalysisRequest) SetOutputFormat(v string) *RunEnterpriseVocAnalysisRequest {
	s.OutputFormat = &v
	return s
}

func (s *RunEnterpriseVocAnalysisRequest) SetSourceTrace(v bool) *RunEnterpriseVocAnalysisRequest {
	s.SourceTrace = &v
	return s
}

func (s *RunEnterpriseVocAnalysisRequest) SetTags(v []*RunEnterpriseVocAnalysisRequestTags) *RunEnterpriseVocAnalysisRequest {
	s.Tags = v
	return s
}

func (s *RunEnterpriseVocAnalysisRequest) SetTaskDescription(v string) *RunEnterpriseVocAnalysisRequest {
	s.TaskDescription = &v
	return s
}

func (s *RunEnterpriseVocAnalysisRequest) Validate() error {
	return dara.Validate(s)
}

type RunEnterpriseVocAnalysisRequestFilterTags struct {
	// 标签定义提示词
	//
	// example:
	//
	// 标签定义提示词
	TagDefinePrompt *string `json:"tagDefinePrompt,omitempty" xml:"tagDefinePrompt,omitempty"`
	// 标签名称
	//
	// example:
	//
	// 标签名称
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s RunEnterpriseVocAnalysisRequestFilterTags) String() string {
	return dara.Prettify(s)
}

func (s RunEnterpriseVocAnalysisRequestFilterTags) GoString() string {
	return s.String()
}

func (s *RunEnterpriseVocAnalysisRequestFilterTags) GetTagDefinePrompt() *string {
	return s.TagDefinePrompt
}

func (s *RunEnterpriseVocAnalysisRequestFilterTags) GetTagName() *string {
	return s.TagName
}

func (s *RunEnterpriseVocAnalysisRequestFilterTags) SetTagDefinePrompt(v string) *RunEnterpriseVocAnalysisRequestFilterTags {
	s.TagDefinePrompt = &v
	return s
}

func (s *RunEnterpriseVocAnalysisRequestFilterTags) SetTagName(v string) *RunEnterpriseVocAnalysisRequestFilterTags {
	s.TagName = &v
	return s
}

func (s *RunEnterpriseVocAnalysisRequestFilterTags) Validate() error {
	return dara.Validate(s)
}

type RunEnterpriseVocAnalysisRequestTags struct {
	// 标签定义提示词
	//
	// example:
	//
	// 标签定义提示词
	TagDefinePrompt *string `json:"tagDefinePrompt,omitempty" xml:"tagDefinePrompt,omitempty"`
	// 标签名称
	//
	// example:
	//
	// 标签名称
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s RunEnterpriseVocAnalysisRequestTags) String() string {
	return dara.Prettify(s)
}

func (s RunEnterpriseVocAnalysisRequestTags) GoString() string {
	return s.String()
}

func (s *RunEnterpriseVocAnalysisRequestTags) GetTagDefinePrompt() *string {
	return s.TagDefinePrompt
}

func (s *RunEnterpriseVocAnalysisRequestTags) GetTagName() *string {
	return s.TagName
}

func (s *RunEnterpriseVocAnalysisRequestTags) SetTagDefinePrompt(v string) *RunEnterpriseVocAnalysisRequestTags {
	s.TagDefinePrompt = &v
	return s
}

func (s *RunEnterpriseVocAnalysisRequestTags) SetTagName(v string) *RunEnterpriseVocAnalysisRequestTags {
	s.TagName = &v
	return s
}

func (s *RunEnterpriseVocAnalysisRequestTags) Validate() error {
	return dara.Validate(s)
}

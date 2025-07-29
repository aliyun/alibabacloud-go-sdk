// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunEnterpriseVocAnalysisShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAkProxy(v string) *RunEnterpriseVocAnalysisShrinkRequest
	GetAkProxy() *string
	SetApiKey(v string) *RunEnterpriseVocAnalysisShrinkRequest
	GetApiKey() *string
	SetContent(v string) *RunEnterpriseVocAnalysisShrinkRequest
	GetContent() *string
	SetExtraInfo(v string) *RunEnterpriseVocAnalysisShrinkRequest
	GetExtraInfo() *string
	SetFilterTagsShrink(v string) *RunEnterpriseVocAnalysisShrinkRequest
	GetFilterTagsShrink() *string
	SetModelId(v string) *RunEnterpriseVocAnalysisShrinkRequest
	GetModelId() *string
	SetOutputFormat(v string) *RunEnterpriseVocAnalysisShrinkRequest
	GetOutputFormat() *string
	SetSourceTrace(v bool) *RunEnterpriseVocAnalysisShrinkRequest
	GetSourceTrace() *bool
	SetTagsShrink(v string) *RunEnterpriseVocAnalysisShrinkRequest
	GetTagsShrink() *string
	SetTaskDescription(v string) *RunEnterpriseVocAnalysisShrinkRequest
	GetTaskDescription() *string
}

type RunEnterpriseVocAnalysisShrinkRequest struct {
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
	FilterTagsShrink *string `json:"filterTags,omitempty" xml:"filterTags,omitempty"`
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
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
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

func (s RunEnterpriseVocAnalysisShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunEnterpriseVocAnalysisShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) GetAkProxy() *string {
	return s.AkProxy
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) GetFilterTagsShrink() *string {
	return s.FilterTagsShrink
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) GetSourceTrace() *bool {
	return s.SourceTrace
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) SetAkProxy(v string) *RunEnterpriseVocAnalysisShrinkRequest {
	s.AkProxy = &v
	return s
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) SetApiKey(v string) *RunEnterpriseVocAnalysisShrinkRequest {
	s.ApiKey = &v
	return s
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) SetContent(v string) *RunEnterpriseVocAnalysisShrinkRequest {
	s.Content = &v
	return s
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) SetExtraInfo(v string) *RunEnterpriseVocAnalysisShrinkRequest {
	s.ExtraInfo = &v
	return s
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) SetFilterTagsShrink(v string) *RunEnterpriseVocAnalysisShrinkRequest {
	s.FilterTagsShrink = &v
	return s
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) SetModelId(v string) *RunEnterpriseVocAnalysisShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) SetOutputFormat(v string) *RunEnterpriseVocAnalysisShrinkRequest {
	s.OutputFormat = &v
	return s
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) SetSourceTrace(v bool) *RunEnterpriseVocAnalysisShrinkRequest {
	s.SourceTrace = &v
	return s
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) SetTagsShrink(v string) *RunEnterpriseVocAnalysisShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) SetTaskDescription(v string) *RunEnterpriseVocAnalysisShrinkRequest {
	s.TaskDescription = &v
	return s
}

func (s *RunEnterpriseVocAnalysisShrinkRequest) Validate() error {
	return dara.Validate(s)
}

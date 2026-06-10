// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogoTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogoVersion(v string) *CreateLogoTaskRequest
	GetLogoVersion() *string
	SetNegativePrompt(v string) *CreateLogoTaskRequest
	GetNegativePrompt() *string
	SetParameters(v string) *CreateLogoTaskRequest
	GetParameters() *string
	SetPrompt(v string) *CreateLogoTaskRequest
	GetPrompt() *string
}

type CreateLogoTaskRequest struct {
	LogoVersion *string `json:"LogoVersion,omitempty" xml:"LogoVersion,omitempty"`
	// Negative prompt
	//
	// example:
	//
	// 不生成政治相关图片内容
	NegativePrompt *string `json:"NegativePrompt,omitempty" xml:"NegativePrompt,omitempty"`
	// Job parameters
	//
	// example:
	//
	// {\\"ehpcutilParam\\":\\"sched job_submit --commandline \\\\\\"/ehpcdata/data/usersTest/huangqiaoyi-1725933699384/huangqiaoyi-1725933699384.slurm\\\\\\" --runasuser TestGfjnSimworks\\"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// Positive prompt
	//
	// example:
	//
	// 请严格遵照原文提取稿件的摘要，不要随意捏造和修改文中出现的新闻数据，不要改变原文的意思和新闻事实，不要在摘要中添加自己的观点。生成的摘要不要出现语法和标点的错误。在生成摘要时请去掉里面的记者、通讯员、稿件来源等信息。
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
}

func (s CreateLogoTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLogoTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateLogoTaskRequest) GetLogoVersion() *string {
	return s.LogoVersion
}

func (s *CreateLogoTaskRequest) GetNegativePrompt() *string {
	return s.NegativePrompt
}

func (s *CreateLogoTaskRequest) GetParameters() *string {
	return s.Parameters
}

func (s *CreateLogoTaskRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *CreateLogoTaskRequest) SetLogoVersion(v string) *CreateLogoTaskRequest {
	s.LogoVersion = &v
	return s
}

func (s *CreateLogoTaskRequest) SetNegativePrompt(v string) *CreateLogoTaskRequest {
	s.NegativePrompt = &v
	return s
}

func (s *CreateLogoTaskRequest) SetParameters(v string) *CreateLogoTaskRequest {
	s.Parameters = &v
	return s
}

func (s *CreateLogoTaskRequest) SetPrompt(v string) *CreateLogoTaskRequest {
	s.Prompt = &v
	return s
}

func (s *CreateLogoTaskRequest) Validate() error {
	return dara.Validate(s)
}

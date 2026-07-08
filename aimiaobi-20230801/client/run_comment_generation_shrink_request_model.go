// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCommentGenerationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowEmoji(v bool) *RunCommentGenerationShrinkRequest
	GetAllowEmoji() *bool
	SetExtraInfo(v string) *RunCommentGenerationShrinkRequest
	GetExtraInfo() *string
	SetLength(v string) *RunCommentGenerationShrinkRequest
	GetLength() *string
	SetLengthRangeShrink(v string) *RunCommentGenerationShrinkRequest
	GetLengthRangeShrink() *string
	SetModelId(v string) *RunCommentGenerationShrinkRequest
	GetModelId() *string
	SetNumComments(v string) *RunCommentGenerationShrinkRequest
	GetNumComments() *string
	SetSentimentShrink(v string) *RunCommentGenerationShrinkRequest
	GetSentimentShrink() *string
	SetSessionId(v string) *RunCommentGenerationShrinkRequest
	GetSessionId() *string
	SetSourceMaterial(v string) *RunCommentGenerationShrinkRequest
	GetSourceMaterial() *string
	SetStyle(v string) *RunCommentGenerationShrinkRequest
	GetStyle() *string
	SetTypeShrink(v string) *RunCommentGenerationShrinkRequest
	GetTypeShrink() *string
	SetWorkspaceId(v string) *RunCommentGenerationShrinkRequest
	GetWorkspaceId() *string
}

type RunCommentGenerationShrinkRequest struct {
	// Set to true to allow emoji in comments. Default is false.
	//
	// example:
	//
	// true
	AllowEmoji *bool `json:"AllowEmoji,omitempty" xml:"AllowEmoji,omitempty"`
	// Additional instructions.
	//
	// example:
	//
	// 不要输出额外其他信息
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// Length of each comment in characters.
	//
	// example:
	//
	// 20
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// Length distribution.
	//
	// Valid keys:
	//
	// - short (up to 20 characters)
	//
	// - medium (20–50 characters)
	//
	// - long (50–100 characters)
	//
	// This parameter is required.
	//
	// example:
	//
	// {"short":"50","long":“50”}
	LengthRangeShrink *string `json:"LengthRange,omitempty" xml:"LengthRange,omitempty"`
	// ID of the model to use.
	//
	// example:
	//
	// quanmiao-max、quanmiao-plus
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// Number of comments to generate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	NumComments *string `json:"NumComments,omitempty" xml:"NumComments,omitempty"`
	// Sentiment distribution.
	//
	// Valid keys:
	//
	// - positive
	//
	// - neutral
	//
	// - negative
	//
	// This parameter is required.
	//
	// example:
	//
	// {"positive":"50","negative":"50"}
	SentimentShrink *string `json:"Sentiment,omitempty" xml:"Sentiment,omitempty"`
	// Session ID.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// Article to comment on.
	//
	// This parameter is required.
	//
	// example:
	//
	// ps5是sony新一代的游戏机，他创新性的...
	SourceMaterial *string `json:"SourceMaterial,omitempty" xml:"SourceMaterial,omitempty"`
	// Tone of the comments.
	//
	// example:
	//
	// 积极正面
	Style *string `json:"Style,omitempty" xml:"Style,omitempty"`
	// Comment type.
	//
	// Valid keys:
	//
	// - emotion (expresses emotion)
	//
	// - opinion (states an opinion)
	//
	// - interaction (encourages interaction)
	//
	// - experience (shares experience)
	//
	// - humor (uses humor)
	//
	// This parameter is required.
	//
	// example:
	//
	// {"emotion":"50","opinion":"50"}
	TypeShrink *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Path parameter. The unique identifier of your Alibaba Cloud Model Studio workspace. To get this ID, see [Get the workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-3kcs1w3lltrtbfkr
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunCommentGenerationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCommentGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunCommentGenerationShrinkRequest) GetAllowEmoji() *bool {
	return s.AllowEmoji
}

func (s *RunCommentGenerationShrinkRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *RunCommentGenerationShrinkRequest) GetLength() *string {
	return s.Length
}

func (s *RunCommentGenerationShrinkRequest) GetLengthRangeShrink() *string {
	return s.LengthRangeShrink
}

func (s *RunCommentGenerationShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunCommentGenerationShrinkRequest) GetNumComments() *string {
	return s.NumComments
}

func (s *RunCommentGenerationShrinkRequest) GetSentimentShrink() *string {
	return s.SentimentShrink
}

func (s *RunCommentGenerationShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunCommentGenerationShrinkRequest) GetSourceMaterial() *string {
	return s.SourceMaterial
}

func (s *RunCommentGenerationShrinkRequest) GetStyle() *string {
	return s.Style
}

func (s *RunCommentGenerationShrinkRequest) GetTypeShrink() *string {
	return s.TypeShrink
}

func (s *RunCommentGenerationShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunCommentGenerationShrinkRequest) SetAllowEmoji(v bool) *RunCommentGenerationShrinkRequest {
	s.AllowEmoji = &v
	return s
}

func (s *RunCommentGenerationShrinkRequest) SetExtraInfo(v string) *RunCommentGenerationShrinkRequest {
	s.ExtraInfo = &v
	return s
}

func (s *RunCommentGenerationShrinkRequest) SetLength(v string) *RunCommentGenerationShrinkRequest {
	s.Length = &v
	return s
}

func (s *RunCommentGenerationShrinkRequest) SetLengthRangeShrink(v string) *RunCommentGenerationShrinkRequest {
	s.LengthRangeShrink = &v
	return s
}

func (s *RunCommentGenerationShrinkRequest) SetModelId(v string) *RunCommentGenerationShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunCommentGenerationShrinkRequest) SetNumComments(v string) *RunCommentGenerationShrinkRequest {
	s.NumComments = &v
	return s
}

func (s *RunCommentGenerationShrinkRequest) SetSentimentShrink(v string) *RunCommentGenerationShrinkRequest {
	s.SentimentShrink = &v
	return s
}

func (s *RunCommentGenerationShrinkRequest) SetSessionId(v string) *RunCommentGenerationShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *RunCommentGenerationShrinkRequest) SetSourceMaterial(v string) *RunCommentGenerationShrinkRequest {
	s.SourceMaterial = &v
	return s
}

func (s *RunCommentGenerationShrinkRequest) SetStyle(v string) *RunCommentGenerationShrinkRequest {
	s.Style = &v
	return s
}

func (s *RunCommentGenerationShrinkRequest) SetTypeShrink(v string) *RunCommentGenerationShrinkRequest {
	s.TypeShrink = &v
	return s
}

func (s *RunCommentGenerationShrinkRequest) SetWorkspaceId(v string) *RunCommentGenerationShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunCommentGenerationShrinkRequest) Validate() error {
	return dara.Validate(s)
}

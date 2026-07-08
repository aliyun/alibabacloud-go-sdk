// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCommentGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowEmoji(v bool) *RunCommentGenerationRequest
	GetAllowEmoji() *bool
	SetExtraInfo(v string) *RunCommentGenerationRequest
	GetExtraInfo() *string
	SetLength(v string) *RunCommentGenerationRequest
	GetLength() *string
	SetLengthRange(v map[string]interface{}) *RunCommentGenerationRequest
	GetLengthRange() map[string]interface{}
	SetModelId(v string) *RunCommentGenerationRequest
	GetModelId() *string
	SetNumComments(v string) *RunCommentGenerationRequest
	GetNumComments() *string
	SetSentiment(v map[string]interface{}) *RunCommentGenerationRequest
	GetSentiment() map[string]interface{}
	SetSessionId(v string) *RunCommentGenerationRequest
	GetSessionId() *string
	SetSourceMaterial(v string) *RunCommentGenerationRequest
	GetSourceMaterial() *string
	SetStyle(v string) *RunCommentGenerationRequest
	GetStyle() *string
	SetType(v map[string]interface{}) *RunCommentGenerationRequest
	GetType() map[string]interface{}
	SetWorkspaceId(v string) *RunCommentGenerationRequest
	GetWorkspaceId() *string
}

type RunCommentGenerationRequest struct {
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
	LengthRange map[string]interface{} `json:"LengthRange,omitempty" xml:"LengthRange,omitempty"`
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
	Sentiment map[string]interface{} `json:"Sentiment,omitempty" xml:"Sentiment,omitempty"`
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
	Type map[string]interface{} `json:"Type,omitempty" xml:"Type,omitempty"`
	// Path parameter. The unique identifier of your Alibaba Cloud Model Studio workspace. To get this ID, see [Get the workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-3kcs1w3lltrtbfkr
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunCommentGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCommentGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunCommentGenerationRequest) GetAllowEmoji() *bool {
	return s.AllowEmoji
}

func (s *RunCommentGenerationRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *RunCommentGenerationRequest) GetLength() *string {
	return s.Length
}

func (s *RunCommentGenerationRequest) GetLengthRange() map[string]interface{} {
	return s.LengthRange
}

func (s *RunCommentGenerationRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunCommentGenerationRequest) GetNumComments() *string {
	return s.NumComments
}

func (s *RunCommentGenerationRequest) GetSentiment() map[string]interface{} {
	return s.Sentiment
}

func (s *RunCommentGenerationRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunCommentGenerationRequest) GetSourceMaterial() *string {
	return s.SourceMaterial
}

func (s *RunCommentGenerationRequest) GetStyle() *string {
	return s.Style
}

func (s *RunCommentGenerationRequest) GetType() map[string]interface{} {
	return s.Type
}

func (s *RunCommentGenerationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunCommentGenerationRequest) SetAllowEmoji(v bool) *RunCommentGenerationRequest {
	s.AllowEmoji = &v
	return s
}

func (s *RunCommentGenerationRequest) SetExtraInfo(v string) *RunCommentGenerationRequest {
	s.ExtraInfo = &v
	return s
}

func (s *RunCommentGenerationRequest) SetLength(v string) *RunCommentGenerationRequest {
	s.Length = &v
	return s
}

func (s *RunCommentGenerationRequest) SetLengthRange(v map[string]interface{}) *RunCommentGenerationRequest {
	s.LengthRange = v
	return s
}

func (s *RunCommentGenerationRequest) SetModelId(v string) *RunCommentGenerationRequest {
	s.ModelId = &v
	return s
}

func (s *RunCommentGenerationRequest) SetNumComments(v string) *RunCommentGenerationRequest {
	s.NumComments = &v
	return s
}

func (s *RunCommentGenerationRequest) SetSentiment(v map[string]interface{}) *RunCommentGenerationRequest {
	s.Sentiment = v
	return s
}

func (s *RunCommentGenerationRequest) SetSessionId(v string) *RunCommentGenerationRequest {
	s.SessionId = &v
	return s
}

func (s *RunCommentGenerationRequest) SetSourceMaterial(v string) *RunCommentGenerationRequest {
	s.SourceMaterial = &v
	return s
}

func (s *RunCommentGenerationRequest) SetStyle(v string) *RunCommentGenerationRequest {
	s.Style = &v
	return s
}

func (s *RunCommentGenerationRequest) SetType(v map[string]interface{}) *RunCommentGenerationRequest {
	s.Type = v
	return s
}

func (s *RunCommentGenerationRequest) SetWorkspaceId(v string) *RunCommentGenerationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunCommentGenerationRequest) Validate() error {
	return dara.Validate(s)
}

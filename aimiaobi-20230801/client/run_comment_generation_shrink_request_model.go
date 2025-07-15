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
	// example:
	//
	// true
	AllowEmoji *bool   `json:"AllowEmoji,omitempty" xml:"AllowEmoji,omitempty"`
	ExtraInfo  *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// example:
	//
	// 20
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"short":"50","long":“50”}
	LengthRangeShrink *string `json:"LengthRange,omitempty" xml:"LengthRange,omitempty"`
	ModelId           *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	NumComments *string `json:"NumComments,omitempty" xml:"NumComments,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"positive":"50","negative":"50"}
	SentimentShrink *string `json:"Sentiment,omitempty" xml:"Sentiment,omitempty"`
	SessionId       *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	SourceMaterial *string `json:"SourceMaterial,omitempty" xml:"SourceMaterial,omitempty"`
	Style          *string `json:"Style,omitempty" xml:"Style,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"emotion":"50","opinion":"50"}
	TypeShrink *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

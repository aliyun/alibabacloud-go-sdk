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
	LengthRange map[string]interface{} `json:"LengthRange,omitempty" xml:"LengthRange,omitempty"`
	ModelId     *string                `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
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
	Sentiment map[string]interface{} `json:"Sentiment,omitempty" xml:"Sentiment,omitempty"`
	SessionId *string                `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	SourceMaterial *string `json:"SourceMaterial,omitempty" xml:"SourceMaterial,omitempty"`
	Style          *string `json:"Style,omitempty" xml:"Style,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"emotion":"50","opinion":"50"}
	Type map[string]interface{} `json:"Type,omitempty" xml:"Type,omitempty"`
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

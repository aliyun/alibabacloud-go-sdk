// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocBrainmapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCleanCache(v bool) *RunDocBrainmapRequest
	GetCleanCache() *bool
	SetDocId(v string) *RunDocBrainmapRequest
	GetDocId() *string
	SetModelName(v string) *RunDocBrainmapRequest
	GetModelName() *string
	SetNodeNumber(v int32) *RunDocBrainmapRequest
	GetNodeNumber() *int32
	SetPrompt(v string) *RunDocBrainmapRequest
	GetPrompt() *string
	SetResponseFormat(v int32) *RunDocBrainmapRequest
	GetResponseFormat() *int32
	SetSessionId(v string) *RunDocBrainmapRequest
	GetSessionId() *string
	SetWordNumber(v int32) *RunDocBrainmapRequest
	GetWordNumber() *int32
	SetWorkspaceId(v string) *RunDocBrainmapRequest
	GetWorkspaceId() *string
	SetReferenceContent(v string) *RunDocBrainmapRequest
	GetReferenceContent() *string
}

type RunDocBrainmapRequest struct {
	// Indicates whether to clear the previous cache.
	//
	// example:
	//
	// true
	CleanCache *bool `json:"CleanCache,omitempty" xml:"CleanCache,omitempty"`
	// The document ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// The name of the model to use.
	//
	// example:
	//
	// quanmiao-max、quanmiao-plus
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The number of nodes to generate at the second level of the mind map.
	//
	// example:
	//
	// 3
	NodeNumber *int32 `json:"NodeNumber,omitempty" xml:"NodeNumber,omitempty"`
	// A custom prompt to guide the mind map generation.
	//
	// example:
	//
	// 请按英文输出
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 0
	ResponseFormat *int32 `json:"ResponseFormat,omitempty" xml:"ResponseFormat,omitempty"`
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The maximum number of words in each node.
	//
	// example:
	//
	// 20
	WordNumber *int32 `json:"WordNumber,omitempty" xml:"WordNumber,omitempty"`
	// The ID of the Model Studio workspace. For more information, see [How to use a workspace](https://help.aliyun.com/document_detail/2782167.html).
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The source content for generating the mind map. This parameter takes precedence over `DocId`.
	//
	// example:
	//
	// 要生成脑图的内容
	ReferenceContent *string `json:"referenceContent,omitempty" xml:"referenceContent,omitempty"`
}

func (s RunDocBrainmapRequest) String() string {
	return dara.Prettify(s)
}

func (s RunDocBrainmapRequest) GoString() string {
	return s.String()
}

func (s *RunDocBrainmapRequest) GetCleanCache() *bool {
	return s.CleanCache
}

func (s *RunDocBrainmapRequest) GetDocId() *string {
	return s.DocId
}

func (s *RunDocBrainmapRequest) GetModelName() *string {
	return s.ModelName
}

func (s *RunDocBrainmapRequest) GetNodeNumber() *int32 {
	return s.NodeNumber
}

func (s *RunDocBrainmapRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunDocBrainmapRequest) GetResponseFormat() *int32 {
	return s.ResponseFormat
}

func (s *RunDocBrainmapRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDocBrainmapRequest) GetWordNumber() *int32 {
	return s.WordNumber
}

func (s *RunDocBrainmapRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunDocBrainmapRequest) GetReferenceContent() *string {
	return s.ReferenceContent
}

func (s *RunDocBrainmapRequest) SetCleanCache(v bool) *RunDocBrainmapRequest {
	s.CleanCache = &v
	return s
}

func (s *RunDocBrainmapRequest) SetDocId(v string) *RunDocBrainmapRequest {
	s.DocId = &v
	return s
}

func (s *RunDocBrainmapRequest) SetModelName(v string) *RunDocBrainmapRequest {
	s.ModelName = &v
	return s
}

func (s *RunDocBrainmapRequest) SetNodeNumber(v int32) *RunDocBrainmapRequest {
	s.NodeNumber = &v
	return s
}

func (s *RunDocBrainmapRequest) SetPrompt(v string) *RunDocBrainmapRequest {
	s.Prompt = &v
	return s
}

func (s *RunDocBrainmapRequest) SetResponseFormat(v int32) *RunDocBrainmapRequest {
	s.ResponseFormat = &v
	return s
}

func (s *RunDocBrainmapRequest) SetSessionId(v string) *RunDocBrainmapRequest {
	s.SessionId = &v
	return s
}

func (s *RunDocBrainmapRequest) SetWordNumber(v int32) *RunDocBrainmapRequest {
	s.WordNumber = &v
	return s
}

func (s *RunDocBrainmapRequest) SetWorkspaceId(v string) *RunDocBrainmapRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunDocBrainmapRequest) SetReferenceContent(v string) *RunDocBrainmapRequest {
	s.ReferenceContent = &v
	return s
}

func (s *RunDocBrainmapRequest) Validate() error {
	return dara.Validate(s)
}

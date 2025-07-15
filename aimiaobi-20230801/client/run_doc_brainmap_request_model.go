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
	CleanCache *bool `json:"CleanCache,omitempty" xml:"CleanCache,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	DocId      *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	ModelName  *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	NodeNumber *int32  `json:"NodeNumber,omitempty" xml:"NodeNumber,omitempty"`
	Prompt     *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	WordNumber *int32  `json:"WordNumber,omitempty" xml:"WordNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId      *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

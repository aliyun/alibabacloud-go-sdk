// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunBookBrainmapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCleanCache(v bool) *RunBookBrainmapRequest
	GetCleanCache() *bool
	SetDocId(v string) *RunBookBrainmapRequest
	GetDocId() *string
	SetNodeNumber(v int32) *RunBookBrainmapRequest
	GetNodeNumber() *int32
	SetPrompt(v string) *RunBookBrainmapRequest
	GetPrompt() *string
	SetSessionId(v string) *RunBookBrainmapRequest
	GetSessionId() *string
	SetWordNumber(v int32) *RunBookBrainmapRequest
	GetWordNumber() *int32
	SetWorkspaceId(v string) *RunBookBrainmapRequest
	GetWorkspaceId() *string
}

type RunBookBrainmapRequest struct {
	// example:
	//
	// true
	CleanCache *bool `json:"CleanCache,omitempty" xml:"CleanCache,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 3
	NodeNumber *int32  `json:"NodeNumber,omitempty" xml:"NodeNumber,omitempty"`
	Prompt     *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 20
	WordNumber *int32 `json:"WordNumber,omitempty" xml:"WordNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-hx72jf15gqyobvd9
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunBookBrainmapRequest) String() string {
	return dara.Prettify(s)
}

func (s RunBookBrainmapRequest) GoString() string {
	return s.String()
}

func (s *RunBookBrainmapRequest) GetCleanCache() *bool {
	return s.CleanCache
}

func (s *RunBookBrainmapRequest) GetDocId() *string {
	return s.DocId
}

func (s *RunBookBrainmapRequest) GetNodeNumber() *int32 {
	return s.NodeNumber
}

func (s *RunBookBrainmapRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunBookBrainmapRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunBookBrainmapRequest) GetWordNumber() *int32 {
	return s.WordNumber
}

func (s *RunBookBrainmapRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunBookBrainmapRequest) SetCleanCache(v bool) *RunBookBrainmapRequest {
	s.CleanCache = &v
	return s
}

func (s *RunBookBrainmapRequest) SetDocId(v string) *RunBookBrainmapRequest {
	s.DocId = &v
	return s
}

func (s *RunBookBrainmapRequest) SetNodeNumber(v int32) *RunBookBrainmapRequest {
	s.NodeNumber = &v
	return s
}

func (s *RunBookBrainmapRequest) SetPrompt(v string) *RunBookBrainmapRequest {
	s.Prompt = &v
	return s
}

func (s *RunBookBrainmapRequest) SetSessionId(v string) *RunBookBrainmapRequest {
	s.SessionId = &v
	return s
}

func (s *RunBookBrainmapRequest) SetWordNumber(v int32) *RunBookBrainmapRequest {
	s.WordNumber = &v
	return s
}

func (s *RunBookBrainmapRequest) SetWorkspaceId(v string) *RunBookBrainmapRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunBookBrainmapRequest) Validate() error {
	return dara.Validate(s)
}

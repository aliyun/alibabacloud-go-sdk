// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTopicSelectionMergeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrompt(v string) *RunTopicSelectionMergeRequest
	GetPrompt() *string
	SetTopics(v []*TopicSelection) *RunTopicSelectionMergeRequest
	GetTopics() []*TopicSelection
	SetWorkspaceId(v string) *RunTopicSelectionMergeRequest
	GetWorkspaceId() *string
}

type RunTopicSelectionMergeRequest struct {
	// Custom merge prompt
	//
	// example:
	//
	// 请从xxxx的角度，分析xxxx事件
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// List of topic perspectives to merge
	//
	// This parameter is required.
	Topics []*TopicSelection `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
	// [Workspace ID](https://help.aliyun.com/document_detail/2782167.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunTopicSelectionMergeRequest) String() string {
	return dara.Prettify(s)
}

func (s RunTopicSelectionMergeRequest) GoString() string {
	return s.String()
}

func (s *RunTopicSelectionMergeRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunTopicSelectionMergeRequest) GetTopics() []*TopicSelection {
	return s.Topics
}

func (s *RunTopicSelectionMergeRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunTopicSelectionMergeRequest) SetPrompt(v string) *RunTopicSelectionMergeRequest {
	s.Prompt = &v
	return s
}

func (s *RunTopicSelectionMergeRequest) SetTopics(v []*TopicSelection) *RunTopicSelectionMergeRequest {
	s.Topics = v
	return s
}

func (s *RunTopicSelectionMergeRequest) SetWorkspaceId(v string) *RunTopicSelectionMergeRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunTopicSelectionMergeRequest) Validate() error {
	if s.Topics != nil {
		for _, item := range s.Topics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContextualRetrievalShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *ContextualRetrievalShrinkRequest
	GetDatasetName() *string
	SetMessagesShrink(v string) *ContextualRetrievalShrinkRequest
	GetMessagesShrink() *string
	SetProjectName(v string) *ContextualRetrievalShrinkRequest
	GetProjectName() *string
	SetRecallOnly(v bool) *ContextualRetrievalShrinkRequest
	GetRecallOnly() *bool
	SetSmartClusterIdsShrink(v string) *ContextualRetrievalShrinkRequest
	GetSmartClusterIdsShrink() *string
}

type ContextualRetrievalShrinkRequest struct {
	// The dataset used for retrieval.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The conversation history and tool calling history. The latest message is at the end (index n-1), and the oldest message is at the beginning (index 0). The messages must be in user-assistant pairs, with a total count of 2*n+1, and the length of the latest question cannot exceed 1,000 characters. The conversation history is limited to 100 messages.
	//
	// This parameter is required.
	MessagesShrink *string `json:"Messages,omitempty" xml:"Messages,omitempty"`
	// The name of the project. For more information about how to obtain the project name, see [Create a project](https://www.alibabacloud.com/help/en/imm/getting-started/create-a-project-1).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Specifies whether to enable only the recall process (embedding search). If this parameter is set to true, the returned data is not reranked, which allows you to customize the reranking process. Default value: false.
	//
	// example:
	//
	// false
	RecallOnly *bool `json:"RecallOnly,omitempty" xml:"RecallOnly,omitempty"`
	// The list of smart cluster IDs, which are used to retrieve files within specific smart clusters.
	SmartClusterIdsShrink *string `json:"SmartClusterIds,omitempty" xml:"SmartClusterIds,omitempty"`
}

func (s ContextualRetrievalShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ContextualRetrievalShrinkRequest) GoString() string {
	return s.String()
}

func (s *ContextualRetrievalShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ContextualRetrievalShrinkRequest) GetMessagesShrink() *string {
	return s.MessagesShrink
}

func (s *ContextualRetrievalShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ContextualRetrievalShrinkRequest) GetRecallOnly() *bool {
	return s.RecallOnly
}

func (s *ContextualRetrievalShrinkRequest) GetSmartClusterIdsShrink() *string {
	return s.SmartClusterIdsShrink
}

func (s *ContextualRetrievalShrinkRequest) SetDatasetName(v string) *ContextualRetrievalShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *ContextualRetrievalShrinkRequest) SetMessagesShrink(v string) *ContextualRetrievalShrinkRequest {
	s.MessagesShrink = &v
	return s
}

func (s *ContextualRetrievalShrinkRequest) SetProjectName(v string) *ContextualRetrievalShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *ContextualRetrievalShrinkRequest) SetRecallOnly(v bool) *ContextualRetrievalShrinkRequest {
	s.RecallOnly = &v
	return s
}

func (s *ContextualRetrievalShrinkRequest) SetSmartClusterIdsShrink(v string) *ContextualRetrievalShrinkRequest {
	s.SmartClusterIdsShrink = &v
	return s
}

func (s *ContextualRetrievalShrinkRequest) Validate() error {
	return dara.Validate(s)
}

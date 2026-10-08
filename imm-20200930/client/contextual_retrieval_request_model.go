// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContextualRetrievalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *ContextualRetrievalRequest
	GetDatasetName() *string
	SetMessages(v []*ContextualMessage) *ContextualRetrievalRequest
	GetMessages() []*ContextualMessage
	SetProjectName(v string) *ContextualRetrievalRequest
	GetProjectName() *string
	SetRecallOnly(v bool) *ContextualRetrievalRequest
	GetRecallOnly() *bool
	SetSmartClusterIds(v []*string) *ContextualRetrievalRequest
	GetSmartClusterIds() []*string
}

type ContextualRetrievalRequest struct {
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
	Messages []*ContextualMessage `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
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
	SmartClusterIds []*string `json:"SmartClusterIds,omitempty" xml:"SmartClusterIds,omitempty" type:"Repeated"`
}

func (s ContextualRetrievalRequest) String() string {
	return dara.Prettify(s)
}

func (s ContextualRetrievalRequest) GoString() string {
	return s.String()
}

func (s *ContextualRetrievalRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ContextualRetrievalRequest) GetMessages() []*ContextualMessage {
	return s.Messages
}

func (s *ContextualRetrievalRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ContextualRetrievalRequest) GetRecallOnly() *bool {
	return s.RecallOnly
}

func (s *ContextualRetrievalRequest) GetSmartClusterIds() []*string {
	return s.SmartClusterIds
}

func (s *ContextualRetrievalRequest) SetDatasetName(v string) *ContextualRetrievalRequest {
	s.DatasetName = &v
	return s
}

func (s *ContextualRetrievalRequest) SetMessages(v []*ContextualMessage) *ContextualRetrievalRequest {
	s.Messages = v
	return s
}

func (s *ContextualRetrievalRequest) SetProjectName(v string) *ContextualRetrievalRequest {
	s.ProjectName = &v
	return s
}

func (s *ContextualRetrievalRequest) SetRecallOnly(v bool) *ContextualRetrievalRequest {
	s.RecallOnly = &v
	return s
}

func (s *ContextualRetrievalRequest) SetSmartClusterIds(v []*string) *ContextualRetrievalRequest {
	s.SmartClusterIds = v
	return s
}

func (s *ContextualRetrievalRequest) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

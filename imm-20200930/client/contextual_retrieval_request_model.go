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
	// The dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The conversation or tool invocation history. The latest message is at the end of the list (with an index number of n-1), whereas the earliest message is at the beginning of the list (with an index number of 0). Historical messages must be provided in user-assistant pairs. The maximum number of messages that you can specify is 2\\*n+1. The current question cannot exceed 1,000 characters in length. The maximum number of historical messages allowed is 100.
	//
	// This parameter is required.
	Messages []*ContextualMessage `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/zh/imm/getting-started/create-a-project-1?spm=a2c4g.11186623.help-menu-search-62354.d_0).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Indicates whether to enable recall-only (embedding-based search). If you set this parameter to true, returned results have not been re-ranked and can be ranked in custom order. Default value: false.
	//
	// example:
	//
	// false
	RecallOnly *bool `json:"RecallOnly,omitempty" xml:"RecallOnly,omitempty"`
	// The IDs of clusters from which results are retrieved.
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

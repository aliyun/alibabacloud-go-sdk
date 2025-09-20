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
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// This parameter is required.
	Messages []*ContextualMessage `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// false
	RecallOnly      *bool     `json:"RecallOnly,omitempty" xml:"RecallOnly,omitempty"`
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
	return dara.Validate(s)
}

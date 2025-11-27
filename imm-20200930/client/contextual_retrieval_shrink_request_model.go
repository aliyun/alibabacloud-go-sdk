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
	MessagesShrink *string `json:"Messages,omitempty" xml:"Messages,omitempty"`
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

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
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// This parameter is required.
	MessagesShrink *string `json:"Messages,omitempty" xml:"Messages,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// false
	RecallOnly            *bool   `json:"RecallOnly,omitempty" xml:"RecallOnly,omitempty"`
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

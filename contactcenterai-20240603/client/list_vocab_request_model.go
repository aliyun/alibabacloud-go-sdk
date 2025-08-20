// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVocabRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *ListVocabRequest
	GetWorkspaceId() *string
}

type ListVocabRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// llm-jhfr****8v
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListVocabRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVocabRequest) GoString() string {
	return s.String()
}

func (s *ListVocabRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListVocabRequest) SetWorkspaceId(v string) *ListVocabRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListVocabRequest) Validate() error {
	return dara.Validate(s)
}

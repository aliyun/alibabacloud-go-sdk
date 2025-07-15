// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocIds(v []*string) *DeleteDocsRequest
	GetDocIds() []*string
	SetWorkspaceId(v string) *DeleteDocsRequest
	GetWorkspaceId() *string
}

type DeleteDocsRequest struct {
	// This parameter is required.
	DocIds []*string `json:"DocIds,omitempty" xml:"DocIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDocsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocsRequest) GetDocIds() []*string {
	return s.DocIds
}

func (s *DeleteDocsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteDocsRequest) SetDocIds(v []*string) *DeleteDocsRequest {
	s.DocIds = v
	return s
}

func (s *DeleteDocsRequest) SetWorkspaceId(v string) *DeleteDocsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDocsRequest) Validate() error {
	return dara.Validate(s)
}

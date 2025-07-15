// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *GetDocInfoRequest
	GetCategoryId() *string
	SetDocId(v string) *GetDocInfoRequest
	GetDocId() *string
	SetWorkspaceId(v string) *GetDocInfoRequest
	GetWorkspaceId() *string
}

type GetDocInfoRequest struct {
	// example:
	//
	// default
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDocInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDocInfoRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *GetDocInfoRequest) GetDocId() *string {
	return s.DocId
}

func (s *GetDocInfoRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDocInfoRequest) SetCategoryId(v string) *GetDocInfoRequest {
	s.CategoryId = &v
	return s
}

func (s *GetDocInfoRequest) SetDocId(v string) *GetDocInfoRequest {
	s.DocId = &v
	return s
}

func (s *GetDocInfoRequest) SetWorkspaceId(v string) *GetDocInfoRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDocInfoRequest) Validate() error {
	return dara.Validate(s)
}

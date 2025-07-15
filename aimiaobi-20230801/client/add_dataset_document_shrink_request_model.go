// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDatasetDocumentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *AddDatasetDocumentShrinkRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *AddDatasetDocumentShrinkRequest
	GetDatasetName() *string
	SetDocumentShrink(v string) *AddDatasetDocumentShrinkRequest
	GetDocumentShrink() *string
	SetWorkspaceId(v string) *AddDatasetDocumentShrinkRequest
	GetWorkspaceId() *string
}

type AddDatasetDocumentShrinkRequest struct {
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// 数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// This parameter is required.
	DocumentShrink *string `json:"Document,omitempty" xml:"Document,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddDatasetDocumentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDatasetDocumentShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddDatasetDocumentShrinkRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *AddDatasetDocumentShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *AddDatasetDocumentShrinkRequest) GetDocumentShrink() *string {
	return s.DocumentShrink
}

func (s *AddDatasetDocumentShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddDatasetDocumentShrinkRequest) SetDatasetId(v int64) *AddDatasetDocumentShrinkRequest {
	s.DatasetId = &v
	return s
}

func (s *AddDatasetDocumentShrinkRequest) SetDatasetName(v string) *AddDatasetDocumentShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *AddDatasetDocumentShrinkRequest) SetDocumentShrink(v string) *AddDatasetDocumentShrinkRequest {
	s.DocumentShrink = &v
	return s
}

func (s *AddDatasetDocumentShrinkRequest) SetWorkspaceId(v string) *AddDatasetDocumentShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddDatasetDocumentShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetDocumentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *UpdateDatasetDocumentShrinkRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *UpdateDatasetDocumentShrinkRequest
	GetDatasetName() *string
	SetDocumentShrink(v string) *UpdateDatasetDocumentShrinkRequest
	GetDocumentShrink() *string
	SetWorkspaceId(v string) *UpdateDatasetDocumentShrinkRequest
	GetWorkspaceId() *string
}

type UpdateDatasetDocumentShrinkRequest struct {
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
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDatasetDocumentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetDocumentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetDocumentShrinkRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *UpdateDatasetDocumentShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateDatasetDocumentShrinkRequest) GetDocumentShrink() *string {
	return s.DocumentShrink
}

func (s *UpdateDatasetDocumentShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateDatasetDocumentShrinkRequest) SetDatasetId(v int64) *UpdateDatasetDocumentShrinkRequest {
	s.DatasetId = &v
	return s
}

func (s *UpdateDatasetDocumentShrinkRequest) SetDatasetName(v string) *UpdateDatasetDocumentShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateDatasetDocumentShrinkRequest) SetDocumentShrink(v string) *UpdateDatasetDocumentShrinkRequest {
	s.DocumentShrink = &v
	return s
}

func (s *UpdateDatasetDocumentShrinkRequest) SetWorkspaceId(v string) *UpdateDatasetDocumentShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDatasetDocumentShrinkRequest) Validate() error {
	return dara.Validate(s)
}

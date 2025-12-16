// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetDocumentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *GetDatasetDocumentShrinkRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *GetDatasetDocumentShrinkRequest
	GetDatasetName() *string
	SetDocId(v string) *GetDatasetDocumentShrinkRequest
	GetDocId() *string
	SetDocUuid(v string) *GetDatasetDocumentShrinkRequest
	GetDocUuid() *string
	SetIncludeFieldsShrink(v string) *GetDatasetDocumentShrinkRequest
	GetIncludeFieldsShrink() *string
	SetWorkspaceId(v string) *GetDatasetDocumentShrinkRequest
	GetWorkspaceId() *string
}

type GetDatasetDocumentShrinkRequest struct {
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// 数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// xxx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// xxx
	DocUuid             *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	IncludeFieldsShrink *string `json:"IncludeFields,omitempty" xml:"IncludeFields,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDatasetDocumentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetDocumentShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetDocumentShrinkRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *GetDatasetDocumentShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetDatasetDocumentShrinkRequest) GetDocId() *string {
	return s.DocId
}

func (s *GetDatasetDocumentShrinkRequest) GetDocUuid() *string {
	return s.DocUuid
}

func (s *GetDatasetDocumentShrinkRequest) GetIncludeFieldsShrink() *string {
	return s.IncludeFieldsShrink
}

func (s *GetDatasetDocumentShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDatasetDocumentShrinkRequest) SetDatasetId(v int64) *GetDatasetDocumentShrinkRequest {
	s.DatasetId = &v
	return s
}

func (s *GetDatasetDocumentShrinkRequest) SetDatasetName(v string) *GetDatasetDocumentShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *GetDatasetDocumentShrinkRequest) SetDocId(v string) *GetDatasetDocumentShrinkRequest {
	s.DocId = &v
	return s
}

func (s *GetDatasetDocumentShrinkRequest) SetDocUuid(v string) *GetDatasetDocumentShrinkRequest {
	s.DocUuid = &v
	return s
}

func (s *GetDatasetDocumentShrinkRequest) SetIncludeFieldsShrink(v string) *GetDatasetDocumentShrinkRequest {
	s.IncludeFieldsShrink = &v
	return s
}

func (s *GetDatasetDocumentShrinkRequest) SetWorkspaceId(v string) *GetDatasetDocumentShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDatasetDocumentShrinkRequest) Validate() error {
	return dara.Validate(s)
}

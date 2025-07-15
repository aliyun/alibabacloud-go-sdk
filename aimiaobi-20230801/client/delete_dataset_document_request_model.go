// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *DeleteDatasetDocumentRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *DeleteDatasetDocumentRequest
	GetDatasetName() *string
	SetDocId(v string) *DeleteDatasetDocumentRequest
	GetDocId() *string
	SetDocUuid(v string) *DeleteDatasetDocumentRequest
	GetDocUuid() *string
	SetWorkspaceId(v string) *DeleteDatasetDocumentRequest
	GetWorkspaceId() *string
}

type DeleteDatasetDocumentRequest struct {
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
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDatasetDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetDocumentRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetDocumentRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *DeleteDatasetDocumentRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *DeleteDatasetDocumentRequest) GetDocId() *string {
	return s.DocId
}

func (s *DeleteDatasetDocumentRequest) GetDocUuid() *string {
	return s.DocUuid
}

func (s *DeleteDatasetDocumentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteDatasetDocumentRequest) SetDatasetId(v int64) *DeleteDatasetDocumentRequest {
	s.DatasetId = &v
	return s
}

func (s *DeleteDatasetDocumentRequest) SetDatasetName(v string) *DeleteDatasetDocumentRequest {
	s.DatasetName = &v
	return s
}

func (s *DeleteDatasetDocumentRequest) SetDocId(v string) *DeleteDatasetDocumentRequest {
	s.DocId = &v
	return s
}

func (s *DeleteDatasetDocumentRequest) SetDocUuid(v string) *DeleteDatasetDocumentRequest {
	s.DocUuid = &v
	return s
}

func (s *DeleteDatasetDocumentRequest) SetWorkspaceId(v string) *DeleteDatasetDocumentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDatasetDocumentRequest) Validate() error {
	return dara.Validate(s)
}

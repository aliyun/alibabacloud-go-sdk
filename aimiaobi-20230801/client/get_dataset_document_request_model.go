// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *GetDatasetDocumentRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *GetDatasetDocumentRequest
	GetDatasetName() *string
	SetDocId(v string) *GetDatasetDocumentRequest
	GetDocId() *string
	SetDocUuid(v string) *GetDatasetDocumentRequest
	GetDocUuid() *string
	SetWorkspaceId(v string) *GetDatasetDocumentRequest
	GetWorkspaceId() *string
}

type GetDatasetDocumentRequest struct {
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
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDatasetDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetDocumentRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetDocumentRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *GetDatasetDocumentRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetDatasetDocumentRequest) GetDocId() *string {
	return s.DocId
}

func (s *GetDatasetDocumentRequest) GetDocUuid() *string {
	return s.DocUuid
}

func (s *GetDatasetDocumentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDatasetDocumentRequest) SetDatasetId(v int64) *GetDatasetDocumentRequest {
	s.DatasetId = &v
	return s
}

func (s *GetDatasetDocumentRequest) SetDatasetName(v string) *GetDatasetDocumentRequest {
	s.DatasetName = &v
	return s
}

func (s *GetDatasetDocumentRequest) SetDocId(v string) *GetDatasetDocumentRequest {
	s.DocId = &v
	return s
}

func (s *GetDatasetDocumentRequest) SetDocUuid(v string) *GetDatasetDocumentRequest {
	s.DocUuid = &v
	return s
}

func (s *GetDatasetDocumentRequest) SetWorkspaceId(v string) *GetDatasetDocumentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDatasetDocumentRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *UpdateDatasetDocumentRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *UpdateDatasetDocumentRequest
	GetDatasetName() *string
	SetDocument(v *UpdateDatasetDocumentRequestDocument) *UpdateDatasetDocumentRequest
	GetDocument() *UpdateDatasetDocumentRequestDocument
	SetWorkspaceId(v string) *UpdateDatasetDocumentRequest
	GetWorkspaceId() *string
}

type UpdateDatasetDocumentRequest struct {
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// 数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// This parameter is required.
	Document *UpdateDatasetDocumentRequestDocument `json:"Document,omitempty" xml:"Document,omitempty" type:"Struct"`
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDatasetDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetDocumentRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetDocumentRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *UpdateDatasetDocumentRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateDatasetDocumentRequest) GetDocument() *UpdateDatasetDocumentRequestDocument {
	return s.Document
}

func (s *UpdateDatasetDocumentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateDatasetDocumentRequest) SetDatasetId(v int64) *UpdateDatasetDocumentRequest {
	s.DatasetId = &v
	return s
}

func (s *UpdateDatasetDocumentRequest) SetDatasetName(v string) *UpdateDatasetDocumentRequest {
	s.DatasetName = &v
	return s
}

func (s *UpdateDatasetDocumentRequest) SetDocument(v *UpdateDatasetDocumentRequestDocument) *UpdateDatasetDocumentRequest {
	s.Document = v
	return s
}

func (s *UpdateDatasetDocumentRequest) SetWorkspaceId(v string) *UpdateDatasetDocumentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDatasetDocumentRequest) Validate() error {
	if s.Document != nil {
		if err := s.Document.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetDocumentRequestDocument struct {
	// example:
	//
	// 用户指定的文档唯一ID
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 内部文档唯一ID
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateDatasetDocumentRequestDocument) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetDocumentRequestDocument) GoString() string {
	return s.String()
}

func (s *UpdateDatasetDocumentRequestDocument) GetDocId() *string {
	return s.DocId
}

func (s *UpdateDatasetDocumentRequestDocument) GetDocUuid() *string {
	return s.DocUuid
}

func (s *UpdateDatasetDocumentRequestDocument) GetTitle() *string {
	return s.Title
}

func (s *UpdateDatasetDocumentRequestDocument) SetDocId(v string) *UpdateDatasetDocumentRequestDocument {
	s.DocId = &v
	return s
}

func (s *UpdateDatasetDocumentRequestDocument) SetDocUuid(v string) *UpdateDatasetDocumentRequestDocument {
	s.DocUuid = &v
	return s
}

func (s *UpdateDatasetDocumentRequestDocument) SetTitle(v string) *UpdateDatasetDocumentRequestDocument {
	s.Title = &v
	return s
}

func (s *UpdateDatasetDocumentRequestDocument) Validate() error {
	return dara.Validate(s)
}

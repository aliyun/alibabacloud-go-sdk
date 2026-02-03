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
	CategoryUuid *string `json:"CategoryUuid,omitempty" xml:"CategoryUuid,omitempty"`
	// example:
	//
	// 用户指定的文档唯一ID
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 内部文档唯一ID
	DocUuid *string   `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	Extend1 *string   `json:"Extend1,omitempty" xml:"Extend1,omitempty"`
	Extend2 *string   `json:"Extend2,omitempty" xml:"Extend2,omitempty"`
	Extend3 *string   `json:"Extend3,omitempty" xml:"Extend3,omitempty"`
	Tags    []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *UpdateDatasetDocumentRequestDocument) GetCategoryUuid() *string {
	return s.CategoryUuid
}

func (s *UpdateDatasetDocumentRequestDocument) GetDocId() *string {
	return s.DocId
}

func (s *UpdateDatasetDocumentRequestDocument) GetDocUuid() *string {
	return s.DocUuid
}

func (s *UpdateDatasetDocumentRequestDocument) GetExtend1() *string {
	return s.Extend1
}

func (s *UpdateDatasetDocumentRequestDocument) GetExtend2() *string {
	return s.Extend2
}

func (s *UpdateDatasetDocumentRequestDocument) GetExtend3() *string {
	return s.Extend3
}

func (s *UpdateDatasetDocumentRequestDocument) GetTags() []*string {
	return s.Tags
}

func (s *UpdateDatasetDocumentRequestDocument) GetTitle() *string {
	return s.Title
}

func (s *UpdateDatasetDocumentRequestDocument) SetCategoryUuid(v string) *UpdateDatasetDocumentRequestDocument {
	s.CategoryUuid = &v
	return s
}

func (s *UpdateDatasetDocumentRequestDocument) SetDocId(v string) *UpdateDatasetDocumentRequestDocument {
	s.DocId = &v
	return s
}

func (s *UpdateDatasetDocumentRequestDocument) SetDocUuid(v string) *UpdateDatasetDocumentRequestDocument {
	s.DocUuid = &v
	return s
}

func (s *UpdateDatasetDocumentRequestDocument) SetExtend1(v string) *UpdateDatasetDocumentRequestDocument {
	s.Extend1 = &v
	return s
}

func (s *UpdateDatasetDocumentRequestDocument) SetExtend2(v string) *UpdateDatasetDocumentRequestDocument {
	s.Extend2 = &v
	return s
}

func (s *UpdateDatasetDocumentRequestDocument) SetExtend3(v string) *UpdateDatasetDocumentRequestDocument {
	s.Extend3 = &v
	return s
}

func (s *UpdateDatasetDocumentRequestDocument) SetTags(v []*string) *UpdateDatasetDocumentRequestDocument {
	s.Tags = v
	return s
}

func (s *UpdateDatasetDocumentRequestDocument) SetTitle(v string) *UpdateDatasetDocumentRequestDocument {
	s.Title = &v
	return s
}

func (s *UpdateDatasetDocumentRequestDocument) Validate() error {
	return dara.Validate(s)
}

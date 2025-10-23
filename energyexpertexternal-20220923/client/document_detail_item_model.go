// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentDetailItem interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DocumentDetailItem
	GetCreateTime() *string
	SetDocHash(v string) *DocumentDetailItem
	GetDocHash() *string
	SetDocName(v string) *DocumentDetailItem
	GetDocName() *string
	SetDocUrl(v string) *DocumentDetailItem
	GetDocUrl() *string
	SetFolderId(v string) *DocumentDetailItem
	GetFolderId() *string
	SetFolderName(v string) *DocumentDetailItem
	GetFolderName() *string
	SetId(v int64) *DocumentDetailItem
	GetId() *int64
	SetJobId(v string) *DocumentDetailItem
	GetJobId() *string
	SetJobStatus(v string) *DocumentDetailItem
	GetJobStatus() *string
	SetOriginDocName(v string) *DocumentDetailItem
	GetOriginDocName() *string
	SetOriginDocUrl(v string) *DocumentDetailItem
	GetOriginDocUrl() *string
	SetUpdateTime(v string) *DocumentDetailItem
	GetUpdateTime() *string
}

type DocumentDetailItem struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	CreateTime    *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DocHash       *string `json:"docHash,omitempty" xml:"docHash,omitempty"`
	DocName       *string `json:"docName,omitempty" xml:"docName,omitempty"`
	DocUrl        *string `json:"docUrl,omitempty" xml:"docUrl,omitempty"`
	FolderId      *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	FolderName    *string `json:"folderName,omitempty" xml:"folderName,omitempty"`
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	JobId         *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	JobStatus     *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	OriginDocName *string `json:"originDocName,omitempty" xml:"originDocName,omitempty"`
	OriginDocUrl  *string `json:"originDocUrl,omitempty" xml:"originDocUrl,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s DocumentDetailItem) String() string {
	return dara.Prettify(s)
}

func (s DocumentDetailItem) GoString() string {
	return s.String()
}

func (s *DocumentDetailItem) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DocumentDetailItem) GetDocHash() *string {
	return s.DocHash
}

func (s *DocumentDetailItem) GetDocName() *string {
	return s.DocName
}

func (s *DocumentDetailItem) GetDocUrl() *string {
	return s.DocUrl
}

func (s *DocumentDetailItem) GetFolderId() *string {
	return s.FolderId
}

func (s *DocumentDetailItem) GetFolderName() *string {
	return s.FolderName
}

func (s *DocumentDetailItem) GetId() *int64 {
	return s.Id
}

func (s *DocumentDetailItem) GetJobId() *string {
	return s.JobId
}

func (s *DocumentDetailItem) GetJobStatus() *string {
	return s.JobStatus
}

func (s *DocumentDetailItem) GetOriginDocName() *string {
	return s.OriginDocName
}

func (s *DocumentDetailItem) GetOriginDocUrl() *string {
	return s.OriginDocUrl
}

func (s *DocumentDetailItem) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DocumentDetailItem) SetCreateTime(v string) *DocumentDetailItem {
	s.CreateTime = &v
	return s
}

func (s *DocumentDetailItem) SetDocHash(v string) *DocumentDetailItem {
	s.DocHash = &v
	return s
}

func (s *DocumentDetailItem) SetDocName(v string) *DocumentDetailItem {
	s.DocName = &v
	return s
}

func (s *DocumentDetailItem) SetDocUrl(v string) *DocumentDetailItem {
	s.DocUrl = &v
	return s
}

func (s *DocumentDetailItem) SetFolderId(v string) *DocumentDetailItem {
	s.FolderId = &v
	return s
}

func (s *DocumentDetailItem) SetFolderName(v string) *DocumentDetailItem {
	s.FolderName = &v
	return s
}

func (s *DocumentDetailItem) SetId(v int64) *DocumentDetailItem {
	s.Id = &v
	return s
}

func (s *DocumentDetailItem) SetJobId(v string) *DocumentDetailItem {
	s.JobId = &v
	return s
}

func (s *DocumentDetailItem) SetJobStatus(v string) *DocumentDetailItem {
	s.JobStatus = &v
	return s
}

func (s *DocumentDetailItem) SetOriginDocName(v string) *DocumentDetailItem {
	s.OriginDocName = &v
	return s
}

func (s *DocumentDetailItem) SetOriginDocUrl(v string) *DocumentDetailItem {
	s.OriginDocUrl = &v
	return s
}

func (s *DocumentDetailItem) SetUpdateTime(v string) *DocumentDetailItem {
	s.UpdateTime = &v
	return s
}

func (s *DocumentDetailItem) Validate() error {
	return dara.Validate(s)
}

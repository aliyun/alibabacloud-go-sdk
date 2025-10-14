// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExportStatus(v string) *GetExportTaskResponseBody
	GetExportStatus() *string
	SetExportType(v string) *GetExportTaskResponseBody
	GetExportType() *string
	SetFileName(v string) *GetExportTaskResponseBody
	GetFileName() *string
	SetGmtCreate(v string) *GetExportTaskResponseBody
	GetGmtCreate() *string
	SetId(v int64) *GetExportTaskResponseBody
	GetId() *int64
	SetLink(v string) *GetExportTaskResponseBody
	GetLink() *string
	SetProgress(v int32) *GetExportTaskResponseBody
	GetProgress() *int32
	SetRequestId(v string) *GetExportTaskResponseBody
	GetRequestId() *string
}

type GetExportTaskResponseBody struct {
	// example:
	//
	// success
	ExportStatus *string `json:"ExportStatus,omitempty" xml:"ExportStatus,omitempty"`
	// example:
	//
	// incident_list
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// example:
	//
	// event_1193842352994186_17344888****.xlsx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 1605076118000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 400185
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// https://cloud-siem-user-dataset.oss-cn-shanghai.aliyuncs.com/export_file/17661858******5/event_17661858******5_175748****.xlsx
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// example:
	//
	// 66
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetExportTaskResponseBody) GetExportStatus() *string {
	return s.ExportStatus
}

func (s *GetExportTaskResponseBody) GetExportType() *string {
	return s.ExportType
}

func (s *GetExportTaskResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *GetExportTaskResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetExportTaskResponseBody) GetId() *int64 {
	return s.Id
}

func (s *GetExportTaskResponseBody) GetLink() *string {
	return s.Link
}

func (s *GetExportTaskResponseBody) GetProgress() *int32 {
	return s.Progress
}

func (s *GetExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExportTaskResponseBody) SetExportStatus(v string) *GetExportTaskResponseBody {
	s.ExportStatus = &v
	return s
}

func (s *GetExportTaskResponseBody) SetExportType(v string) *GetExportTaskResponseBody {
	s.ExportType = &v
	return s
}

func (s *GetExportTaskResponseBody) SetFileName(v string) *GetExportTaskResponseBody {
	s.FileName = &v
	return s
}

func (s *GetExportTaskResponseBody) SetGmtCreate(v string) *GetExportTaskResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetExportTaskResponseBody) SetId(v int64) *GetExportTaskResponseBody {
	s.Id = &v
	return s
}

func (s *GetExportTaskResponseBody) SetLink(v string) *GetExportTaskResponseBody {
	s.Link = &v
	return s
}

func (s *GetExportTaskResponseBody) SetProgress(v int32) *GetExportTaskResponseBody {
	s.Progress = &v
	return s
}

func (s *GetExportTaskResponseBody) SetRequestId(v string) *GetExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

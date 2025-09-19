// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventExportInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExportStatus(v string) *DescribeSuspEventExportInfoResponseBody
	GetExportStatus() *string
	SetFileName(v string) *DescribeSuspEventExportInfoResponseBody
	GetFileName() *string
	SetGmtCreate(v int64) *DescribeSuspEventExportInfoResponseBody
	GetGmtCreate() *int64
	SetGmtModified(v int64) *DescribeSuspEventExportInfoResponseBody
	GetGmtModified() *int64
	SetId(v int32) *DescribeSuspEventExportInfoResponseBody
	GetId() *int32
	SetLink(v string) *DescribeSuspEventExportInfoResponseBody
	GetLink() *string
	SetProgress(v int32) *DescribeSuspEventExportInfoResponseBody
	GetProgress() *int32
	SetProperties(v string) *DescribeSuspEventExportInfoResponseBody
	GetProperties() *string
	SetRequestId(v string) *DescribeSuspEventExportInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeSuspEventExportInfoResponseBody
	GetTotalCount() *int32
	SetType(v string) *DescribeSuspEventExportInfoResponseBody
	GetType() *string
}

type DescribeSuspEventExportInfoResponseBody struct {
	// The handling status for the exception. Valid values:
	//
	// 	- **exporting**: in progress
	//
	// 	- **success**: successful
	//
	// 	- **failed**: failed
	//
	// 	- **pending**: pending
	//
	// example:
	//
	// success
	ExportStatus *string `json:"ExportStatus,omitempty" xml:"ExportStatus,omitempty"`
	// The name of the exported file.
	//
	// example:
	//
	// suspicious_event_20221221_1671590521234.zip
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The time when the export task was created.
	//
	// example:
	//
	// 2022-12-20T15:18Z
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the export task was modified.
	//
	// example:
	//
	// 2022-12-20T15:18Z
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the export task.
	//
	// example:
	//
	// 11
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The URL at which you can download the exported Excel file.
	//
	// example:
	//
	// http://suspicious-xxxxxxx.oss-cn-shanghai.aliyuncs.com/xxxxxxxxxxx/suspicious_event_20221221_1671590525269.zip?Expires=1671594125&OSSAccessKeyId=yourAccessKeyID&Signature=xxxxxxxxxxxxxxxxxx
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// The progress percentage of the export task.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The exported parameters of exceptions.
	//
	// example:
	//
	// id,eventSubType,eventDetail,level,status,ip,instanceName,desc,lastTime,operateTime,note
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578AXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of exceptions exported.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The type of the export task. The value is fixed as suspiciousEvent.
	//
	// example:
	//
	// suspiciousEvent
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSuspEventExportInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventExportInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventExportInfoResponseBody) GetExportStatus() *string {
	return s.ExportStatus
}

func (s *DescribeSuspEventExportInfoResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *DescribeSuspEventExportInfoResponseBody) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeSuspEventExportInfoResponseBody) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeSuspEventExportInfoResponseBody) GetId() *int32 {
	return s.Id
}

func (s *DescribeSuspEventExportInfoResponseBody) GetLink() *string {
	return s.Link
}

func (s *DescribeSuspEventExportInfoResponseBody) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeSuspEventExportInfoResponseBody) GetProperties() *string {
	return s.Properties
}

func (s *DescribeSuspEventExportInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSuspEventExportInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSuspEventExportInfoResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeSuspEventExportInfoResponseBody) SetExportStatus(v string) *DescribeSuspEventExportInfoResponseBody {
	s.ExportStatus = &v
	return s
}

func (s *DescribeSuspEventExportInfoResponseBody) SetFileName(v string) *DescribeSuspEventExportInfoResponseBody {
	s.FileName = &v
	return s
}

func (s *DescribeSuspEventExportInfoResponseBody) SetGmtCreate(v int64) *DescribeSuspEventExportInfoResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSuspEventExportInfoResponseBody) SetGmtModified(v int64) *DescribeSuspEventExportInfoResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeSuspEventExportInfoResponseBody) SetId(v int32) *DescribeSuspEventExportInfoResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventExportInfoResponseBody) SetLink(v string) *DescribeSuspEventExportInfoResponseBody {
	s.Link = &v
	return s
}

func (s *DescribeSuspEventExportInfoResponseBody) SetProgress(v int32) *DescribeSuspEventExportInfoResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribeSuspEventExportInfoResponseBody) SetProperties(v string) *DescribeSuspEventExportInfoResponseBody {
	s.Properties = &v
	return s
}

func (s *DescribeSuspEventExportInfoResponseBody) SetRequestId(v string) *DescribeSuspEventExportInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventExportInfoResponseBody) SetTotalCount(v int32) *DescribeSuspEventExportInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSuspEventExportInfoResponseBody) SetType(v string) *DescribeSuspEventExportInfoResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeSuspEventExportInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

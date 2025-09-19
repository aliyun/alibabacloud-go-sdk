// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWarningExportInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentCount(v int32) *DescribeWarningExportInfoResponseBody
	GetCurrentCount() *int32
	SetExportStatus(v string) *DescribeWarningExportInfoResponseBody
	GetExportStatus() *string
	SetFileName(v string) *DescribeWarningExportInfoResponseBody
	GetFileName() *string
	SetId(v int64) *DescribeWarningExportInfoResponseBody
	GetId() *int64
	SetLink(v string) *DescribeWarningExportInfoResponseBody
	GetLink() *string
	SetMessage(v string) *DescribeWarningExportInfoResponseBody
	GetMessage() *string
	SetProgress(v int32) *DescribeWarningExportInfoResponseBody
	GetProgress() *int32
	SetRequestId(v string) *DescribeWarningExportInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWarningExportInfoResponseBody
	GetTotalCount() *int32
}

type DescribeWarningExportInfoResponseBody struct {
	// The number of baseline entries that are exported.
	//
	// example:
	//
	// 1
	CurrentCount *int32 `json:"CurrentCount,omitempty" xml:"CurrentCount,omitempty"`
	// The status of the export task.
	//
	// Valid values:
	//
	// 	- **init**: The task is being initialized.
	//
	// 	- **exporting**: The task is in progress.
	//
	// 	- **success**: The task is successful.
	//
	// example:
	//
	// success
	ExportStatus *string `json:"ExportStatus,omitempty" xml:"ExportStatus,omitempty"`
	// The name of the exported Excel file.
	//
	// example:
	//
	// app_20210917
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The ID of the export task for the baseline check result.
	//
	// example:
	//
	// 131231
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The URL at which you can download the exported Excel file.
	//
	// example:
	//
	// https://eds.aliyun.com/notification/entitle/64b5c3e2-e52b-4d29-9617-e7e6d74XXXX
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// The message that shows the task result. The value is fixed as **successful**, which indicates that the export task is complete.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The progress percentage of the export task.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 79CFF74D-E967-5407-8A78-EE03B925FDAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of baseline entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWarningExportInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarningExportInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWarningExportInfoResponseBody) GetCurrentCount() *int32 {
	return s.CurrentCount
}

func (s *DescribeWarningExportInfoResponseBody) GetExportStatus() *string {
	return s.ExportStatus
}

func (s *DescribeWarningExportInfoResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *DescribeWarningExportInfoResponseBody) GetId() *int64 {
	return s.Id
}

func (s *DescribeWarningExportInfoResponseBody) GetLink() *string {
	return s.Link
}

func (s *DescribeWarningExportInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeWarningExportInfoResponseBody) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeWarningExportInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWarningExportInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWarningExportInfoResponseBody) SetCurrentCount(v int32) *DescribeWarningExportInfoResponseBody {
	s.CurrentCount = &v
	return s
}

func (s *DescribeWarningExportInfoResponseBody) SetExportStatus(v string) *DescribeWarningExportInfoResponseBody {
	s.ExportStatus = &v
	return s
}

func (s *DescribeWarningExportInfoResponseBody) SetFileName(v string) *DescribeWarningExportInfoResponseBody {
	s.FileName = &v
	return s
}

func (s *DescribeWarningExportInfoResponseBody) SetId(v int64) *DescribeWarningExportInfoResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeWarningExportInfoResponseBody) SetLink(v string) *DescribeWarningExportInfoResponseBody {
	s.Link = &v
	return s
}

func (s *DescribeWarningExportInfoResponseBody) SetMessage(v string) *DescribeWarningExportInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeWarningExportInfoResponseBody) SetProgress(v int32) *DescribeWarningExportInfoResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribeWarningExportInfoResponseBody) SetRequestId(v string) *DescribeWarningExportInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWarningExportInfoResponseBody) SetTotalCount(v int32) *DescribeWarningExportInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWarningExportInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

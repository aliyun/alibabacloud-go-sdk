// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHcExportInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentCount(v int32) *DescribeHcExportInfoResponseBody
	GetCurrentCount() *int32
	SetFileName(v string) *DescribeHcExportInfoResponseBody
	GetFileName() *string
	SetGmtCreate(v int64) *DescribeHcExportInfoResponseBody
	GetGmtCreate() *int64
	SetId(v int64) *DescribeHcExportInfoResponseBody
	GetId() *int64
	SetLink(v string) *DescribeHcExportInfoResponseBody
	GetLink() *string
	SetProgress(v int32) *DescribeHcExportInfoResponseBody
	GetProgress() *int32
	SetRequestId(v string) *DescribeHcExportInfoResponseBody
	GetRequestId() *string
	SetResultStatus(v string) *DescribeHcExportInfoResponseBody
	GetResultStatus() *string
	SetTotalCount(v int32) *DescribeHcExportInfoResponseBody
	GetTotalCount() *int32
}

type DescribeHcExportInfoResponseBody struct {
	// The number of exported entries.
	//
	// example:
	//
	// 148
	CurrentCount *int32 `json:"CurrentCount,omitempty" xml:"CurrentCount,omitempty"`
	// The name of the exported file.
	//
	// example:
	//
	// health_check_export_2022****
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The time when the export task was created.
	//
	// example:
	//
	// 2022-11-03T15:15Z
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the export task.
	//
	// example:
	//
	// 1082278
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The download URL for the exported file.
	//
	// example:
	//
	// https://hc-export.oss-cn-shanghai.aliyuncs.com/export_hc/health_check_export_20221222_1671699255808.zip?Expires=1672304056&OSSAccessKeyId=****&Signature=****
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// The progress percentage of the export task.
	//
	// example:
	//
	// 89
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3C2C94CF-ED08-50C0-BC72-C5029251****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the export task. Valid values:
	//
	// 	- **exporting**: The task is in progress.
	//
	// 	- **success**: The task is complete.
	//
	// example:
	//
	// exporting
	ResultStatus *string `json:"ResultStatus,omitempty" xml:"ResultStatus,omitempty"`
	// The total number of exported entries.
	//
	// example:
	//
	// 624
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHcExportInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHcExportInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHcExportInfoResponseBody) GetCurrentCount() *int32 {
	return s.CurrentCount
}

func (s *DescribeHcExportInfoResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *DescribeHcExportInfoResponseBody) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeHcExportInfoResponseBody) GetId() *int64 {
	return s.Id
}

func (s *DescribeHcExportInfoResponseBody) GetLink() *string {
	return s.Link
}

func (s *DescribeHcExportInfoResponseBody) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeHcExportInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHcExportInfoResponseBody) GetResultStatus() *string {
	return s.ResultStatus
}

func (s *DescribeHcExportInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHcExportInfoResponseBody) SetCurrentCount(v int32) *DescribeHcExportInfoResponseBody {
	s.CurrentCount = &v
	return s
}

func (s *DescribeHcExportInfoResponseBody) SetFileName(v string) *DescribeHcExportInfoResponseBody {
	s.FileName = &v
	return s
}

func (s *DescribeHcExportInfoResponseBody) SetGmtCreate(v int64) *DescribeHcExportInfoResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeHcExportInfoResponseBody) SetId(v int64) *DescribeHcExportInfoResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeHcExportInfoResponseBody) SetLink(v string) *DescribeHcExportInfoResponseBody {
	s.Link = &v
	return s
}

func (s *DescribeHcExportInfoResponseBody) SetProgress(v int32) *DescribeHcExportInfoResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribeHcExportInfoResponseBody) SetRequestId(v string) *DescribeHcExportInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHcExportInfoResponseBody) SetResultStatus(v string) *DescribeHcExportInfoResponseBody {
	s.ResultStatus = &v
	return s
}

func (s *DescribeHcExportInfoResponseBody) SetTotalCount(v int32) *DescribeHcExportInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHcExportInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

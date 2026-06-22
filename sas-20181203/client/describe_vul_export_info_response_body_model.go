// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulExportInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentCount(v int32) *DescribeVulExportInfoResponseBody
	GetCurrentCount() *int32
	SetExportStatus(v string) *DescribeVulExportInfoResponseBody
	GetExportStatus() *string
	SetFileName(v string) *DescribeVulExportInfoResponseBody
	GetFileName() *string
	SetId(v int64) *DescribeVulExportInfoResponseBody
	GetId() *int64
	SetLink(v string) *DescribeVulExportInfoResponseBody
	GetLink() *string
	SetMessage(v string) *DescribeVulExportInfoResponseBody
	GetMessage() *string
	SetProgress(v int32) *DescribeVulExportInfoResponseBody
	GetProgress() *int32
	SetRequestId(v string) *DescribeVulExportInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVulExportInfoResponseBody
	GetTotalCount() *int32
}

type DescribeVulExportInfoResponseBody struct {
	// The number of data entries that have been exported.
	//
	// example:
	//
	// 1
	CurrentCount *int32 `json:"CurrentCount,omitempty" xml:"CurrentCount,omitempty"`
	// The status of the export task. Valid values:
	//
	// - **init**: initializing
	//
	// - **exporting**: exporting in progress
	//
	// - **success**: export successful.
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
	// The ID of the vulnerability export task.
	//
	// example:
	//
	// 14356
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The download link of the exported Excel file.
	//
	// example:
	//
	// http://www.aliyun.com
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// The message about the export result. Fixed value: **success**. This value indicates that the export is successful.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The export progress percentage (%).
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The request ID. Alibaba Cloud generates a unique identifier for each API request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 4C1AE3F3-18FA-4108-BXXX-AFA1A032756C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of data entries in the exported Excel file.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVulExportInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulExportInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulExportInfoResponseBody) GetCurrentCount() *int32 {
	return s.CurrentCount
}

func (s *DescribeVulExportInfoResponseBody) GetExportStatus() *string {
	return s.ExportStatus
}

func (s *DescribeVulExportInfoResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *DescribeVulExportInfoResponseBody) GetId() *int64 {
	return s.Id
}

func (s *DescribeVulExportInfoResponseBody) GetLink() *string {
	return s.Link
}

func (s *DescribeVulExportInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeVulExportInfoResponseBody) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeVulExportInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVulExportInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVulExportInfoResponseBody) SetCurrentCount(v int32) *DescribeVulExportInfoResponseBody {
	s.CurrentCount = &v
	return s
}

func (s *DescribeVulExportInfoResponseBody) SetExportStatus(v string) *DescribeVulExportInfoResponseBody {
	s.ExportStatus = &v
	return s
}

func (s *DescribeVulExportInfoResponseBody) SetFileName(v string) *DescribeVulExportInfoResponseBody {
	s.FileName = &v
	return s
}

func (s *DescribeVulExportInfoResponseBody) SetId(v int64) *DescribeVulExportInfoResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeVulExportInfoResponseBody) SetLink(v string) *DescribeVulExportInfoResponseBody {
	s.Link = &v
	return s
}

func (s *DescribeVulExportInfoResponseBody) SetMessage(v string) *DescribeVulExportInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeVulExportInfoResponseBody) SetProgress(v int32) *DescribeVulExportInfoResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribeVulExportInfoResponseBody) SetRequestId(v string) *DescribeVulExportInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulExportInfoResponseBody) SetTotalCount(v int32) *DescribeVulExportInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulExportInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

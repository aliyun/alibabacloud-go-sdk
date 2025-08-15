// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExportInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentCount(v int32) *DescribeExportInfoResponseBody
	GetCurrentCount() *int32
	SetExportStatus(v string) *DescribeExportInfoResponseBody
	GetExportStatus() *string
	SetFileName(v string) *DescribeExportInfoResponseBody
	GetFileName() *string
	SetId(v int64) *DescribeExportInfoResponseBody
	GetId() *int64
	SetLink(v string) *DescribeExportInfoResponseBody
	GetLink() *string
	SetMessage(v string) *DescribeExportInfoResponseBody
	GetMessage() *string
	SetProgress(v int32) *DescribeExportInfoResponseBody
	GetProgress() *int32
	SetRequestId(v string) *DescribeExportInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeExportInfoResponseBody
	GetTotalCount() *int32
}

type DescribeExportInfoResponseBody struct {
	// example:
	//
	// 13
	CurrentCount *int32 `json:"CurrentCount,omitempty" xml:"CurrentCount,omitempty"`
	// example:
	//
	// success
	ExportStatus *string `json:"ExportStatus,omitempty" xml:"ExportStatus,omitempty"`
	// example:
	//
	// health_check_export_20******
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 991012
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// http://www.aliyun.com******
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 61
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// AAC546A5-5EDC-5939-86A3-56DFAF******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 28
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExportInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExportInfoResponseBody) GetCurrentCount() *int32 {
	return s.CurrentCount
}

func (s *DescribeExportInfoResponseBody) GetExportStatus() *string {
	return s.ExportStatus
}

func (s *DescribeExportInfoResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *DescribeExportInfoResponseBody) GetId() *int64 {
	return s.Id
}

func (s *DescribeExportInfoResponseBody) GetLink() *string {
	return s.Link
}

func (s *DescribeExportInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeExportInfoResponseBody) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeExportInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExportInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeExportInfoResponseBody) SetCurrentCount(v int32) *DescribeExportInfoResponseBody {
	s.CurrentCount = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetExportStatus(v string) *DescribeExportInfoResponseBody {
	s.ExportStatus = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetFileName(v string) *DescribeExportInfoResponseBody {
	s.FileName = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetId(v int64) *DescribeExportInfoResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetLink(v string) *DescribeExportInfoResponseBody {
	s.Link = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetMessage(v string) *DescribeExportInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetProgress(v int32) *DescribeExportInfoResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetRequestId(v string) *DescribeExportInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetTotalCount(v int32) *DescribeExportInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeExportInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

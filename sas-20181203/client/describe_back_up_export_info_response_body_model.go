// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackUpExportInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeBackUpExportInfoResponseBodyData) *DescribeBackUpExportInfoResponseBody
	GetData() []*DescribeBackUpExportInfoResponseBodyData
	SetPageInfo(v *DescribeBackUpExportInfoResponseBodyPageInfo) *DescribeBackUpExportInfoResponseBody
	GetPageInfo() *DescribeBackUpExportInfoResponseBodyPageInfo
	SetRequestId(v string) *DescribeBackUpExportInfoResponseBody
	GetRequestId() *string
}

type DescribeBackUpExportInfoResponseBody struct {
	// The data returned.
	Data []*DescribeBackUpExportInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeBackUpExportInfoResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578ABF384
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackUpExportInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackUpExportInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackUpExportInfoResponseBody) GetData() []*DescribeBackUpExportInfoResponseBodyData {
	return s.Data
}

func (s *DescribeBackUpExportInfoResponseBody) GetPageInfo() *DescribeBackUpExportInfoResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeBackUpExportInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackUpExportInfoResponseBody) SetData(v []*DescribeBackUpExportInfoResponseBodyData) *DescribeBackUpExportInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBackUpExportInfoResponseBody) SetPageInfo(v *DescribeBackUpExportInfoResponseBodyPageInfo) *DescribeBackUpExportInfoResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeBackUpExportInfoResponseBody) SetRequestId(v string) *DescribeBackUpExportInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackUpExportInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackUpExportInfoResponseBodyData struct {
	// The number of exported entries.
	//
	// example:
	//
	// 29
	CurrentCount *int32 `json:"CurrentCount,omitempty" xml:"CurrentCount,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// suspicious_event_20221203
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The time when the export task was created.
	//
	// example:
	//
	// 1671607025000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the export task.
	//
	// example:
	//
	// 273698***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The URL at which you can download the archived information.
	//
	// example:
	//
	// http://xxx.oss-cn-xxx.aliyuncs.com/export/assetInstance_20221221_1671606250570.zip
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// The error message that is returned when the export task fails.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The progress percentage of the export task.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of the export task. Valid values:
	//
	// 	- **init**: The task is being initialized.
	//
	// 	- **exporting**: The task is in progress.
	//
	// 	- **success**: The task is complete.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of entries in the file.
	//
	// example:
	//
	// 29
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackUpExportInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackUpExportInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackUpExportInfoResponseBodyData) GetCurrentCount() *int32 {
	return s.CurrentCount
}

func (s *DescribeBackUpExportInfoResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *DescribeBackUpExportInfoResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeBackUpExportInfoResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DescribeBackUpExportInfoResponseBodyData) GetLink() *string {
	return s.Link
}

func (s *DescribeBackUpExportInfoResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DescribeBackUpExportInfoResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeBackUpExportInfoResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeBackUpExportInfoResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBackUpExportInfoResponseBodyData) SetCurrentCount(v int32) *DescribeBackUpExportInfoResponseBodyData {
	s.CurrentCount = &v
	return s
}

func (s *DescribeBackUpExportInfoResponseBodyData) SetFileName(v string) *DescribeBackUpExportInfoResponseBodyData {
	s.FileName = &v
	return s
}

func (s *DescribeBackUpExportInfoResponseBodyData) SetGmtCreate(v int64) *DescribeBackUpExportInfoResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeBackUpExportInfoResponseBodyData) SetId(v int64) *DescribeBackUpExportInfoResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeBackUpExportInfoResponseBodyData) SetLink(v string) *DescribeBackUpExportInfoResponseBodyData {
	s.Link = &v
	return s
}

func (s *DescribeBackUpExportInfoResponseBodyData) SetMessage(v string) *DescribeBackUpExportInfoResponseBodyData {
	s.Message = &v
	return s
}

func (s *DescribeBackUpExportInfoResponseBodyData) SetProgress(v int32) *DescribeBackUpExportInfoResponseBodyData {
	s.Progress = &v
	return s
}

func (s *DescribeBackUpExportInfoResponseBodyData) SetStatus(v string) *DescribeBackUpExportInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeBackUpExportInfoResponseBodyData) SetTotalCount(v int32) *DescribeBackUpExportInfoResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackUpExportInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeBackUpExportInfoResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 9
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 29
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackUpExportInfoResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackUpExportInfoResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeBackUpExportInfoResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeBackUpExportInfoResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeBackUpExportInfoResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackUpExportInfoResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBackUpExportInfoResponseBodyPageInfo) SetCount(v int32) *DescribeBackUpExportInfoResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeBackUpExportInfoResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeBackUpExportInfoResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBackUpExportInfoResponseBodyPageInfo) SetPageSize(v int32) *DescribeBackUpExportInfoResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeBackUpExportInfoResponseBodyPageInfo) SetTotalCount(v int32) *DescribeBackUpExportInfoResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackUpExportInfoResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

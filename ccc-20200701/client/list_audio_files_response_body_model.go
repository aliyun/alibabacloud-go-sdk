// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAudioFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAudioFilesResponseBody
	GetCode() *string
	SetData(v *ListAudioFilesResponseBodyData) *ListAudioFilesResponseBody
	GetData() *ListAudioFilesResponseBodyData
	SetHttpStatusCode(v int32) *ListAudioFilesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAudioFilesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAudioFilesResponseBody
	GetRequestId() *string
}

type ListAudioFilesResponseBody struct {
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListAudioFilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 53223330-EBF1-586B-A2CB-93C3B711FDA0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAudioFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAudioFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAudioFilesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAudioFilesResponseBody) GetData() *ListAudioFilesResponseBodyData {
	return s.Data
}

func (s *ListAudioFilesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAudioFilesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAudioFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAudioFilesResponseBody) SetCode(v string) *ListAudioFilesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAudioFilesResponseBody) SetData(v *ListAudioFilesResponseBodyData) *ListAudioFilesResponseBody {
	s.Data = v
	return s
}

func (s *ListAudioFilesResponseBody) SetHttpStatusCode(v int32) *ListAudioFilesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAudioFilesResponseBody) SetMessage(v string) *ListAudioFilesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAudioFilesResponseBody) SetRequestId(v string) *ListAudioFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAudioFilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAudioFilesResponseBodyData struct {
	List []*ListAudioFilesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAudioFilesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAudioFilesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAudioFilesResponseBodyData) GetList() []*ListAudioFilesResponseBodyDataList {
	return s.List
}

func (s *ListAudioFilesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAudioFilesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAudioFilesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAudioFilesResponseBodyData) SetList(v []*ListAudioFilesResponseBodyDataList) *ListAudioFilesResponseBodyData {
	s.List = v
	return s
}

func (s *ListAudioFilesResponseBodyData) SetPageNumber(v int32) *ListAudioFilesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAudioFilesResponseBodyData) SetPageSize(v int32) *ListAudioFilesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAudioFilesResponseBodyData) SetTotalCount(v int32) *ListAudioFilesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListAudioFilesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListAudioFilesResponseBodyDataList struct {
	// example:
	//
	// test-file.wav
	AudioFileName *string `json:"AudioFileName,omitempty" xml:"AudioFileName,omitempty"`
	// example:
	//
	// d5cd7a94-3b6a-47d2-b7fd-0b1cd839bf77
	AudioResourceId *string `json:"AudioResourceId,omitempty" xml:"AudioResourceId,omitempty"`
	AuditResult     *string `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	// example:
	//
	// 2021-03-05 17:35:45.0
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ccc-test/test-file.wav
	OssFileKey *string `json:"OssFileKey,omitempty" xml:"OssFileKey,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2021-03-08 15:34:49.0
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	Usage       *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListAudioFilesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListAudioFilesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListAudioFilesResponseBodyDataList) GetAudioFileName() *string {
	return s.AudioFileName
}

func (s *ListAudioFilesResponseBodyDataList) GetAudioResourceId() *string {
	return s.AudioResourceId
}

func (s *ListAudioFilesResponseBodyDataList) GetAuditResult() *string {
	return s.AuditResult
}

func (s *ListAudioFilesResponseBodyDataList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListAudioFilesResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAudioFilesResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListAudioFilesResponseBodyDataList) GetOssFileKey() *string {
	return s.OssFileKey
}

func (s *ListAudioFilesResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *ListAudioFilesResponseBodyDataList) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *ListAudioFilesResponseBodyDataList) GetUsage() *string {
	return s.Usage
}

func (s *ListAudioFilesResponseBodyDataList) SetAudioFileName(v string) *ListAudioFilesResponseBodyDataList {
	s.AudioFileName = &v
	return s
}

func (s *ListAudioFilesResponseBodyDataList) SetAudioResourceId(v string) *ListAudioFilesResponseBodyDataList {
	s.AudioResourceId = &v
	return s
}

func (s *ListAudioFilesResponseBodyDataList) SetAuditResult(v string) *ListAudioFilesResponseBodyDataList {
	s.AuditResult = &v
	return s
}

func (s *ListAudioFilesResponseBodyDataList) SetCreatedTime(v string) *ListAudioFilesResponseBodyDataList {
	s.CreatedTime = &v
	return s
}

func (s *ListAudioFilesResponseBodyDataList) SetInstanceId(v string) *ListAudioFilesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListAudioFilesResponseBodyDataList) SetName(v string) *ListAudioFilesResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListAudioFilesResponseBodyDataList) SetOssFileKey(v string) *ListAudioFilesResponseBodyDataList {
	s.OssFileKey = &v
	return s
}

func (s *ListAudioFilesResponseBodyDataList) SetStatus(v string) *ListAudioFilesResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListAudioFilesResponseBodyDataList) SetUpdatedTime(v string) *ListAudioFilesResponseBodyDataList {
	s.UpdatedTime = &v
	return s
}

func (s *ListAudioFilesResponseBodyDataList) SetUsage(v string) *ListAudioFilesResponseBodyDataList {
	s.Usage = &v
	return s
}

func (s *ListAudioFilesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

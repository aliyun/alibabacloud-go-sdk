// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectComputeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SelectComputeTaskResponseBody
	GetCode() *string
	SetData(v *SelectComputeTaskResponseBodyData) *SelectComputeTaskResponseBody
	GetData() *SelectComputeTaskResponseBodyData
	SetMsg(v string) *SelectComputeTaskResponseBody
	GetMsg() *string
	SetRequestId(v string) *SelectComputeTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SelectComputeTaskResponseBody
	GetSuccess() *bool
}

type SelectComputeTaskResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SelectComputeTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                            `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SelectComputeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SelectComputeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SelectComputeTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SelectComputeTaskResponseBody) GetData() *SelectComputeTaskResponseBodyData {
	return s.Data
}

func (s *SelectComputeTaskResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *SelectComputeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SelectComputeTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SelectComputeTaskResponseBody) SetCode(v string) *SelectComputeTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SelectComputeTaskResponseBody) SetData(v *SelectComputeTaskResponseBodyData) *SelectComputeTaskResponseBody {
	s.Data = v
	return s
}

func (s *SelectComputeTaskResponseBody) SetMsg(v string) *SelectComputeTaskResponseBody {
	s.Msg = &v
	return s
}

func (s *SelectComputeTaskResponseBody) SetRequestId(v string) *SelectComputeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SelectComputeTaskResponseBody) SetSuccess(v bool) *SelectComputeTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SelectComputeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type SelectComputeTaskResponseBodyData struct {
	AppId               *int64                                                `json:"appId,omitempty" xml:"appId,omitempty"`
	BcId                *int64                                                `json:"bcId,omitempty" xml:"bcId,omitempty"`
	ComputeOssFileTitle *string                                               `json:"computeOssFileTitle,omitempty" xml:"computeOssFileTitle,omitempty"`
	DatasetIds          *string                                               `json:"datasetIds,omitempty" xml:"datasetIds,omitempty"`
	ExportOssFileList   []*SelectComputeTaskResponseBodyDataExportOssFileList `json:"exportOssFileList,omitempty" xml:"exportOssFileList,omitempty" type:"Repeated"`
	ExtInfo             *string                                               `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	FileNum             *int64                                                `json:"fileNum,omitempty" xml:"fileNum,omitempty"`
	Name                *string                                               `json:"name,omitempty" xml:"name,omitempty"`
	Remarks             *string                                               `json:"remarks,omitempty" xml:"remarks,omitempty"`
	Status              *string                                               `json:"status,omitempty" xml:"status,omitempty"`
	TaskResultList      []*SelectComputeTaskResponseBodyDataTaskResultList    `json:"taskResultList,omitempty" xml:"taskResultList,omitempty" type:"Repeated"`
}

func (s SelectComputeTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SelectComputeTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SelectComputeTaskResponseBodyData) GetAppId() *int64 {
	return s.AppId
}

func (s *SelectComputeTaskResponseBodyData) GetBcId() *int64 {
	return s.BcId
}

func (s *SelectComputeTaskResponseBodyData) GetComputeOssFileTitle() *string {
	return s.ComputeOssFileTitle
}

func (s *SelectComputeTaskResponseBodyData) GetDatasetIds() *string {
	return s.DatasetIds
}

func (s *SelectComputeTaskResponseBodyData) GetExportOssFileList() []*SelectComputeTaskResponseBodyDataExportOssFileList {
	return s.ExportOssFileList
}

func (s *SelectComputeTaskResponseBodyData) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *SelectComputeTaskResponseBodyData) GetFileNum() *int64 {
	return s.FileNum
}

func (s *SelectComputeTaskResponseBodyData) GetName() *string {
	return s.Name
}

func (s *SelectComputeTaskResponseBodyData) GetRemarks() *string {
	return s.Remarks
}

func (s *SelectComputeTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SelectComputeTaskResponseBodyData) GetTaskResultList() []*SelectComputeTaskResponseBodyDataTaskResultList {
	return s.TaskResultList
}

func (s *SelectComputeTaskResponseBodyData) SetAppId(v int64) *SelectComputeTaskResponseBodyData {
	s.AppId = &v
	return s
}

func (s *SelectComputeTaskResponseBodyData) SetBcId(v int64) *SelectComputeTaskResponseBodyData {
	s.BcId = &v
	return s
}

func (s *SelectComputeTaskResponseBodyData) SetComputeOssFileTitle(v string) *SelectComputeTaskResponseBodyData {
	s.ComputeOssFileTitle = &v
	return s
}

func (s *SelectComputeTaskResponseBodyData) SetDatasetIds(v string) *SelectComputeTaskResponseBodyData {
	s.DatasetIds = &v
	return s
}

func (s *SelectComputeTaskResponseBodyData) SetExportOssFileList(v []*SelectComputeTaskResponseBodyDataExportOssFileList) *SelectComputeTaskResponseBodyData {
	s.ExportOssFileList = v
	return s
}

func (s *SelectComputeTaskResponseBodyData) SetExtInfo(v string) *SelectComputeTaskResponseBodyData {
	s.ExtInfo = &v
	return s
}

func (s *SelectComputeTaskResponseBodyData) SetFileNum(v int64) *SelectComputeTaskResponseBodyData {
	s.FileNum = &v
	return s
}

func (s *SelectComputeTaskResponseBodyData) SetName(v string) *SelectComputeTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *SelectComputeTaskResponseBodyData) SetRemarks(v string) *SelectComputeTaskResponseBodyData {
	s.Remarks = &v
	return s
}

func (s *SelectComputeTaskResponseBodyData) SetStatus(v string) *SelectComputeTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *SelectComputeTaskResponseBodyData) SetTaskResultList(v []*SelectComputeTaskResponseBodyDataTaskResultList) *SelectComputeTaskResponseBodyData {
	s.TaskResultList = v
	return s
}

func (s *SelectComputeTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type SelectComputeTaskResponseBodyDataExportOssFileList struct {
	DownLoadUrl *string `json:"downLoadUrl,omitempty" xml:"downLoadUrl,omitempty"`
	Password    *string `json:"password,omitempty" xml:"password,omitempty"`
}

func (s SelectComputeTaskResponseBodyDataExportOssFileList) String() string {
	return dara.Prettify(s)
}

func (s SelectComputeTaskResponseBodyDataExportOssFileList) GoString() string {
	return s.String()
}

func (s *SelectComputeTaskResponseBodyDataExportOssFileList) GetDownLoadUrl() *string {
	return s.DownLoadUrl
}

func (s *SelectComputeTaskResponseBodyDataExportOssFileList) GetPassword() *string {
	return s.Password
}

func (s *SelectComputeTaskResponseBodyDataExportOssFileList) SetDownLoadUrl(v string) *SelectComputeTaskResponseBodyDataExportOssFileList {
	s.DownLoadUrl = &v
	return s
}

func (s *SelectComputeTaskResponseBodyDataExportOssFileList) SetPassword(v string) *SelectComputeTaskResponseBodyDataExportOssFileList {
	s.Password = &v
	return s
}

func (s *SelectComputeTaskResponseBodyDataExportOssFileList) Validate() error {
	return dara.Validate(s)
}

type SelectComputeTaskResponseBodyDataTaskResultList struct {
	BcId    *int64 `json:"bcId,omitempty" xml:"bcId,omitempty"`
	Code    *int32 `json:"code,omitempty" xml:"code,omitempty"`
	LineNum *int64 `json:"lineNum,omitempty" xml:"lineNum,omitempty"`
}

func (s SelectComputeTaskResponseBodyDataTaskResultList) String() string {
	return dara.Prettify(s)
}

func (s SelectComputeTaskResponseBodyDataTaskResultList) GoString() string {
	return s.String()
}

func (s *SelectComputeTaskResponseBodyDataTaskResultList) GetBcId() *int64 {
	return s.BcId
}

func (s *SelectComputeTaskResponseBodyDataTaskResultList) GetCode() *int32 {
	return s.Code
}

func (s *SelectComputeTaskResponseBodyDataTaskResultList) GetLineNum() *int64 {
	return s.LineNum
}

func (s *SelectComputeTaskResponseBodyDataTaskResultList) SetBcId(v int64) *SelectComputeTaskResponseBodyDataTaskResultList {
	s.BcId = &v
	return s
}

func (s *SelectComputeTaskResponseBodyDataTaskResultList) SetCode(v int32) *SelectComputeTaskResponseBodyDataTaskResultList {
	s.Code = &v
	return s
}

func (s *SelectComputeTaskResponseBodyDataTaskResultList) SetLineNum(v int64) *SelectComputeTaskResponseBodyDataTaskResultList {
	s.LineNum = &v
	return s
}

func (s *SelectComputeTaskResponseBodyDataTaskResultList) Validate() error {
	return dara.Validate(s)
}

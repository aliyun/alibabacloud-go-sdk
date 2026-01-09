// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectComputeTask2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SelectComputeTask2ResponseBody
	GetCode() *string
	SetData(v *SelectComputeTask2ResponseBodyData) *SelectComputeTask2ResponseBody
	GetData() *SelectComputeTask2ResponseBodyData
	SetMsg(v string) *SelectComputeTask2ResponseBody
	GetMsg() *string
	SetRequestId(v string) *SelectComputeTask2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SelectComputeTask2ResponseBody
	GetSuccess() *bool
}

type SelectComputeTask2ResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SelectComputeTask2ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                             `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SelectComputeTask2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SelectComputeTask2ResponseBody) GoString() string {
	return s.String()
}

func (s *SelectComputeTask2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *SelectComputeTask2ResponseBody) GetData() *SelectComputeTask2ResponseBodyData {
	return s.Data
}

func (s *SelectComputeTask2ResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *SelectComputeTask2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SelectComputeTask2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SelectComputeTask2ResponseBody) SetCode(v string) *SelectComputeTask2ResponseBody {
	s.Code = &v
	return s
}

func (s *SelectComputeTask2ResponseBody) SetData(v *SelectComputeTask2ResponseBodyData) *SelectComputeTask2ResponseBody {
	s.Data = v
	return s
}

func (s *SelectComputeTask2ResponseBody) SetMsg(v string) *SelectComputeTask2ResponseBody {
	s.Msg = &v
	return s
}

func (s *SelectComputeTask2ResponseBody) SetRequestId(v string) *SelectComputeTask2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *SelectComputeTask2ResponseBody) SetSuccess(v bool) *SelectComputeTask2ResponseBody {
	s.Success = &v
	return s
}

func (s *SelectComputeTask2ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SelectComputeTask2ResponseBodyData struct {
	AppId               *int64                                                 `json:"appId,omitempty" xml:"appId,omitempty"`
	BcId                *int64                                                 `json:"bcId,omitempty" xml:"bcId,omitempty"`
	ComputeOssFileTitle *string                                                `json:"computeOssFileTitle,omitempty" xml:"computeOssFileTitle,omitempty"`
	DatasetIds          *string                                                `json:"datasetIds,omitempty" xml:"datasetIds,omitempty"`
	ExportOssFileList   []*SelectComputeTask2ResponseBodyDataExportOssFileList `json:"exportOssFileList,omitempty" xml:"exportOssFileList,omitempty" type:"Repeated"`
	ExtInfo             *string                                                `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	FileNum             *int64                                                 `json:"fileNum,omitempty" xml:"fileNum,omitempty"`
	Hint                *string                                                `json:"hint,omitempty" xml:"hint,omitempty"`
	Name                *string                                                `json:"name,omitempty" xml:"name,omitempty"`
	Remarks             *string                                                `json:"remarks,omitempty" xml:"remarks,omitempty"`
	Status              *string                                                `json:"status,omitempty" xml:"status,omitempty"`
	TaskResultList      []*SelectComputeTask2ResponseBodyDataTaskResultList    `json:"taskResultList,omitempty" xml:"taskResultList,omitempty" type:"Repeated"`
}

func (s SelectComputeTask2ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SelectComputeTask2ResponseBodyData) GoString() string {
	return s.String()
}

func (s *SelectComputeTask2ResponseBodyData) GetAppId() *int64 {
	return s.AppId
}

func (s *SelectComputeTask2ResponseBodyData) GetBcId() *int64 {
	return s.BcId
}

func (s *SelectComputeTask2ResponseBodyData) GetComputeOssFileTitle() *string {
	return s.ComputeOssFileTitle
}

func (s *SelectComputeTask2ResponseBodyData) GetDatasetIds() *string {
	return s.DatasetIds
}

func (s *SelectComputeTask2ResponseBodyData) GetExportOssFileList() []*SelectComputeTask2ResponseBodyDataExportOssFileList {
	return s.ExportOssFileList
}

func (s *SelectComputeTask2ResponseBodyData) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *SelectComputeTask2ResponseBodyData) GetFileNum() *int64 {
	return s.FileNum
}

func (s *SelectComputeTask2ResponseBodyData) GetHint() *string {
	return s.Hint
}

func (s *SelectComputeTask2ResponseBodyData) GetName() *string {
	return s.Name
}

func (s *SelectComputeTask2ResponseBodyData) GetRemarks() *string {
	return s.Remarks
}

func (s *SelectComputeTask2ResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SelectComputeTask2ResponseBodyData) GetTaskResultList() []*SelectComputeTask2ResponseBodyDataTaskResultList {
	return s.TaskResultList
}

func (s *SelectComputeTask2ResponseBodyData) SetAppId(v int64) *SelectComputeTask2ResponseBodyData {
	s.AppId = &v
	return s
}

func (s *SelectComputeTask2ResponseBodyData) SetBcId(v int64) *SelectComputeTask2ResponseBodyData {
	s.BcId = &v
	return s
}

func (s *SelectComputeTask2ResponseBodyData) SetComputeOssFileTitle(v string) *SelectComputeTask2ResponseBodyData {
	s.ComputeOssFileTitle = &v
	return s
}

func (s *SelectComputeTask2ResponseBodyData) SetDatasetIds(v string) *SelectComputeTask2ResponseBodyData {
	s.DatasetIds = &v
	return s
}

func (s *SelectComputeTask2ResponseBodyData) SetExportOssFileList(v []*SelectComputeTask2ResponseBodyDataExportOssFileList) *SelectComputeTask2ResponseBodyData {
	s.ExportOssFileList = v
	return s
}

func (s *SelectComputeTask2ResponseBodyData) SetExtInfo(v string) *SelectComputeTask2ResponseBodyData {
	s.ExtInfo = &v
	return s
}

func (s *SelectComputeTask2ResponseBodyData) SetFileNum(v int64) *SelectComputeTask2ResponseBodyData {
	s.FileNum = &v
	return s
}

func (s *SelectComputeTask2ResponseBodyData) SetHint(v string) *SelectComputeTask2ResponseBodyData {
	s.Hint = &v
	return s
}

func (s *SelectComputeTask2ResponseBodyData) SetName(v string) *SelectComputeTask2ResponseBodyData {
	s.Name = &v
	return s
}

func (s *SelectComputeTask2ResponseBodyData) SetRemarks(v string) *SelectComputeTask2ResponseBodyData {
	s.Remarks = &v
	return s
}

func (s *SelectComputeTask2ResponseBodyData) SetStatus(v string) *SelectComputeTask2ResponseBodyData {
	s.Status = &v
	return s
}

func (s *SelectComputeTask2ResponseBodyData) SetTaskResultList(v []*SelectComputeTask2ResponseBodyDataTaskResultList) *SelectComputeTask2ResponseBodyData {
	s.TaskResultList = v
	return s
}

func (s *SelectComputeTask2ResponseBodyData) Validate() error {
	if s.ExportOssFileList != nil {
		for _, item := range s.ExportOssFileList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TaskResultList != nil {
		for _, item := range s.TaskResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SelectComputeTask2ResponseBodyDataExportOssFileList struct {
	DownLoadUrl *string `json:"downLoadUrl,omitempty" xml:"downLoadUrl,omitempty"`
	Password    *string `json:"password,omitempty" xml:"password,omitempty"`
}

func (s SelectComputeTask2ResponseBodyDataExportOssFileList) String() string {
	return dara.Prettify(s)
}

func (s SelectComputeTask2ResponseBodyDataExportOssFileList) GoString() string {
	return s.String()
}

func (s *SelectComputeTask2ResponseBodyDataExportOssFileList) GetDownLoadUrl() *string {
	return s.DownLoadUrl
}

func (s *SelectComputeTask2ResponseBodyDataExportOssFileList) GetPassword() *string {
	return s.Password
}

func (s *SelectComputeTask2ResponseBodyDataExportOssFileList) SetDownLoadUrl(v string) *SelectComputeTask2ResponseBodyDataExportOssFileList {
	s.DownLoadUrl = &v
	return s
}

func (s *SelectComputeTask2ResponseBodyDataExportOssFileList) SetPassword(v string) *SelectComputeTask2ResponseBodyDataExportOssFileList {
	s.Password = &v
	return s
}

func (s *SelectComputeTask2ResponseBodyDataExportOssFileList) Validate() error {
	return dara.Validate(s)
}

type SelectComputeTask2ResponseBodyDataTaskResultList struct {
	BcId    *int64 `json:"bcId,omitempty" xml:"bcId,omitempty"`
	Code    *int32 `json:"code,omitempty" xml:"code,omitempty"`
	LineNum *int64 `json:"lineNum,omitempty" xml:"lineNum,omitempty"`
}

func (s SelectComputeTask2ResponseBodyDataTaskResultList) String() string {
	return dara.Prettify(s)
}

func (s SelectComputeTask2ResponseBodyDataTaskResultList) GoString() string {
	return s.String()
}

func (s *SelectComputeTask2ResponseBodyDataTaskResultList) GetBcId() *int64 {
	return s.BcId
}

func (s *SelectComputeTask2ResponseBodyDataTaskResultList) GetCode() *int32 {
	return s.Code
}

func (s *SelectComputeTask2ResponseBodyDataTaskResultList) GetLineNum() *int64 {
	return s.LineNum
}

func (s *SelectComputeTask2ResponseBodyDataTaskResultList) SetBcId(v int64) *SelectComputeTask2ResponseBodyDataTaskResultList {
	s.BcId = &v
	return s
}

func (s *SelectComputeTask2ResponseBodyDataTaskResultList) SetCode(v int32) *SelectComputeTask2ResponseBodyDataTaskResultList {
	s.Code = &v
	return s
}

func (s *SelectComputeTask2ResponseBodyDataTaskResultList) SetLineNum(v int64) *SelectComputeTask2ResponseBodyDataTaskResultList {
	s.LineNum = &v
	return s
}

func (s *SelectComputeTask2ResponseBodyDataTaskResultList) Validate() error {
	return dara.Validate(s)
}

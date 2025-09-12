// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeTask2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListComputeTask2ResponseBody
	GetCode() *string
	SetData(v *ListComputeTask2ResponseBodyData) *ListComputeTask2ResponseBody
	GetData() *ListComputeTask2ResponseBodyData
	SetMsg(v string) *ListComputeTask2ResponseBody
	GetMsg() *string
	SetRequestId(v string) *ListComputeTask2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListComputeTask2ResponseBody
	GetSuccess() *bool
}

type ListComputeTask2ResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListComputeTask2ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                           `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListComputeTask2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComputeTask2ResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeTask2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListComputeTask2ResponseBody) GetData() *ListComputeTask2ResponseBodyData {
	return s.Data
}

func (s *ListComputeTask2ResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ListComputeTask2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComputeTask2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListComputeTask2ResponseBody) SetCode(v string) *ListComputeTask2ResponseBody {
	s.Code = &v
	return s
}

func (s *ListComputeTask2ResponseBody) SetData(v *ListComputeTask2ResponseBodyData) *ListComputeTask2ResponseBody {
	s.Data = v
	return s
}

func (s *ListComputeTask2ResponseBody) SetMsg(v string) *ListComputeTask2ResponseBody {
	s.Msg = &v
	return s
}

func (s *ListComputeTask2ResponseBody) SetRequestId(v string) *ListComputeTask2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComputeTask2ResponseBody) SetSuccess(v bool) *ListComputeTask2ResponseBody {
	s.Success = &v
	return s
}

func (s *ListComputeTask2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListComputeTask2ResponseBodyData struct {
	Data     []*ListComputeTask2ResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	TotalNum *int32                                  `json:"totalNum,omitempty" xml:"totalNum,omitempty"`
}

func (s ListComputeTask2ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListComputeTask2ResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListComputeTask2ResponseBodyData) GetData() []*ListComputeTask2ResponseBodyDataData {
	return s.Data
}

func (s *ListComputeTask2ResponseBodyData) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListComputeTask2ResponseBodyData) SetData(v []*ListComputeTask2ResponseBodyDataData) *ListComputeTask2ResponseBodyData {
	s.Data = v
	return s
}

func (s *ListComputeTask2ResponseBodyData) SetTotalNum(v int32) *ListComputeTask2ResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *ListComputeTask2ResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListComputeTask2ResponseBodyDataData struct {
	AppId               *int64                                                `json:"appId,omitempty" xml:"appId,omitempty"`
	BcId                *int64                                                `json:"bcId,omitempty" xml:"bcId,omitempty"`
	ComputeOssFileTitle *string                                               `json:"computeOssFileTitle,omitempty" xml:"computeOssFileTitle,omitempty"`
	DatasetIds          *string                                               `json:"datasetIds,omitempty" xml:"datasetIds,omitempty"`
	ExtInfo             *string                                               `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	FileNum             *int64                                                `json:"fileNum,omitempty" xml:"fileNum,omitempty"`
	Name                *string                                               `json:"name,omitempty" xml:"name,omitempty"`
	Remarks             *string                                               `json:"remarks,omitempty" xml:"remarks,omitempty"`
	Status              *string                                               `json:"status,omitempty" xml:"status,omitempty"`
	TaskResultList      []*ListComputeTask2ResponseBodyDataDataTaskResultList `json:"taskResultList,omitempty" xml:"taskResultList,omitempty" type:"Repeated"`
}

func (s ListComputeTask2ResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s ListComputeTask2ResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListComputeTask2ResponseBodyDataData) GetAppId() *int64 {
	return s.AppId
}

func (s *ListComputeTask2ResponseBodyDataData) GetBcId() *int64 {
	return s.BcId
}

func (s *ListComputeTask2ResponseBodyDataData) GetComputeOssFileTitle() *string {
	return s.ComputeOssFileTitle
}

func (s *ListComputeTask2ResponseBodyDataData) GetDatasetIds() *string {
	return s.DatasetIds
}

func (s *ListComputeTask2ResponseBodyDataData) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *ListComputeTask2ResponseBodyDataData) GetFileNum() *int64 {
	return s.FileNum
}

func (s *ListComputeTask2ResponseBodyDataData) GetName() *string {
	return s.Name
}

func (s *ListComputeTask2ResponseBodyDataData) GetRemarks() *string {
	return s.Remarks
}

func (s *ListComputeTask2ResponseBodyDataData) GetStatus() *string {
	return s.Status
}

func (s *ListComputeTask2ResponseBodyDataData) GetTaskResultList() []*ListComputeTask2ResponseBodyDataDataTaskResultList {
	return s.TaskResultList
}

func (s *ListComputeTask2ResponseBodyDataData) SetAppId(v int64) *ListComputeTask2ResponseBodyDataData {
	s.AppId = &v
	return s
}

func (s *ListComputeTask2ResponseBodyDataData) SetBcId(v int64) *ListComputeTask2ResponseBodyDataData {
	s.BcId = &v
	return s
}

func (s *ListComputeTask2ResponseBodyDataData) SetComputeOssFileTitle(v string) *ListComputeTask2ResponseBodyDataData {
	s.ComputeOssFileTitle = &v
	return s
}

func (s *ListComputeTask2ResponseBodyDataData) SetDatasetIds(v string) *ListComputeTask2ResponseBodyDataData {
	s.DatasetIds = &v
	return s
}

func (s *ListComputeTask2ResponseBodyDataData) SetExtInfo(v string) *ListComputeTask2ResponseBodyDataData {
	s.ExtInfo = &v
	return s
}

func (s *ListComputeTask2ResponseBodyDataData) SetFileNum(v int64) *ListComputeTask2ResponseBodyDataData {
	s.FileNum = &v
	return s
}

func (s *ListComputeTask2ResponseBodyDataData) SetName(v string) *ListComputeTask2ResponseBodyDataData {
	s.Name = &v
	return s
}

func (s *ListComputeTask2ResponseBodyDataData) SetRemarks(v string) *ListComputeTask2ResponseBodyDataData {
	s.Remarks = &v
	return s
}

func (s *ListComputeTask2ResponseBodyDataData) SetStatus(v string) *ListComputeTask2ResponseBodyDataData {
	s.Status = &v
	return s
}

func (s *ListComputeTask2ResponseBodyDataData) SetTaskResultList(v []*ListComputeTask2ResponseBodyDataDataTaskResultList) *ListComputeTask2ResponseBodyDataData {
	s.TaskResultList = v
	return s
}

func (s *ListComputeTask2ResponseBodyDataData) Validate() error {
	return dara.Validate(s)
}

type ListComputeTask2ResponseBodyDataDataTaskResultList struct {
	BcId    *int64 `json:"bcId,omitempty" xml:"bcId,omitempty"`
	Code    *int32 `json:"code,omitempty" xml:"code,omitempty"`
	LineNum *int64 `json:"lineNum,omitempty" xml:"lineNum,omitempty"`
}

func (s ListComputeTask2ResponseBodyDataDataTaskResultList) String() string {
	return dara.Prettify(s)
}

func (s ListComputeTask2ResponseBodyDataDataTaskResultList) GoString() string {
	return s.String()
}

func (s *ListComputeTask2ResponseBodyDataDataTaskResultList) GetBcId() *int64 {
	return s.BcId
}

func (s *ListComputeTask2ResponseBodyDataDataTaskResultList) GetCode() *int32 {
	return s.Code
}

func (s *ListComputeTask2ResponseBodyDataDataTaskResultList) GetLineNum() *int64 {
	return s.LineNum
}

func (s *ListComputeTask2ResponseBodyDataDataTaskResultList) SetBcId(v int64) *ListComputeTask2ResponseBodyDataDataTaskResultList {
	s.BcId = &v
	return s
}

func (s *ListComputeTask2ResponseBodyDataDataTaskResultList) SetCode(v int32) *ListComputeTask2ResponseBodyDataDataTaskResultList {
	s.Code = &v
	return s
}

func (s *ListComputeTask2ResponseBodyDataDataTaskResultList) SetLineNum(v int64) *ListComputeTask2ResponseBodyDataDataTaskResultList {
	s.LineNum = &v
	return s
}

func (s *ListComputeTask2ResponseBodyDataDataTaskResultList) Validate() error {
	return dara.Validate(s)
}

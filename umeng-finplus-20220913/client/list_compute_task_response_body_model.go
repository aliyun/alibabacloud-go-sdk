// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListComputeTaskResponseBody
	GetCode() *string
	SetData(v []*ListComputeTaskResponseBodyData) *ListComputeTaskResponseBody
	GetData() []*ListComputeTaskResponseBodyData
	SetMsg(v string) *ListComputeTaskResponseBody
	GetMsg() *string
	SetRequestId(v string) *ListComputeTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListComputeTaskResponseBody
	GetSuccess() *bool
}

type ListComputeTaskResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListComputeTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Msg       *string                            `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListComputeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComputeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListComputeTaskResponseBody) GetData() []*ListComputeTaskResponseBodyData {
	return s.Data
}

func (s *ListComputeTaskResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ListComputeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComputeTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListComputeTaskResponseBody) SetCode(v string) *ListComputeTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListComputeTaskResponseBody) SetData(v []*ListComputeTaskResponseBodyData) *ListComputeTaskResponseBody {
	s.Data = v
	return s
}

func (s *ListComputeTaskResponseBody) SetMsg(v string) *ListComputeTaskResponseBody {
	s.Msg = &v
	return s
}

func (s *ListComputeTaskResponseBody) SetRequestId(v string) *ListComputeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComputeTaskResponseBody) SetSuccess(v bool) *ListComputeTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ListComputeTaskResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListComputeTaskResponseBodyData struct {
	AppId               *int64                                           `json:"appId,omitempty" xml:"appId,omitempty"`
	BcId                *int64                                           `json:"bcId,omitempty" xml:"bcId,omitempty"`
	ComputeOssFileTitle *string                                          `json:"computeOssFileTitle,omitempty" xml:"computeOssFileTitle,omitempty"`
	DatasetIds          *string                                          `json:"datasetIds,omitempty" xml:"datasetIds,omitempty"`
	ExtInfo             *string                                          `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	FileNum             *int64                                           `json:"fileNum,omitempty" xml:"fileNum,omitempty"`
	Name                *string                                          `json:"name,omitempty" xml:"name,omitempty"`
	Remarks             *string                                          `json:"remarks,omitempty" xml:"remarks,omitempty"`
	Status              *string                                          `json:"status,omitempty" xml:"status,omitempty"`
	TaskResultList      []*ListComputeTaskResponseBodyDataTaskResultList `json:"taskResultList,omitempty" xml:"taskResultList,omitempty" type:"Repeated"`
}

func (s ListComputeTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListComputeTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListComputeTaskResponseBodyData) GetAppId() *int64 {
	return s.AppId
}

func (s *ListComputeTaskResponseBodyData) GetBcId() *int64 {
	return s.BcId
}

func (s *ListComputeTaskResponseBodyData) GetComputeOssFileTitle() *string {
	return s.ComputeOssFileTitle
}

func (s *ListComputeTaskResponseBodyData) GetDatasetIds() *string {
	return s.DatasetIds
}

func (s *ListComputeTaskResponseBodyData) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *ListComputeTaskResponseBodyData) GetFileNum() *int64 {
	return s.FileNum
}

func (s *ListComputeTaskResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListComputeTaskResponseBodyData) GetRemarks() *string {
	return s.Remarks
}

func (s *ListComputeTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListComputeTaskResponseBodyData) GetTaskResultList() []*ListComputeTaskResponseBodyDataTaskResultList {
	return s.TaskResultList
}

func (s *ListComputeTaskResponseBodyData) SetAppId(v int64) *ListComputeTaskResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListComputeTaskResponseBodyData) SetBcId(v int64) *ListComputeTaskResponseBodyData {
	s.BcId = &v
	return s
}

func (s *ListComputeTaskResponseBodyData) SetComputeOssFileTitle(v string) *ListComputeTaskResponseBodyData {
	s.ComputeOssFileTitle = &v
	return s
}

func (s *ListComputeTaskResponseBodyData) SetDatasetIds(v string) *ListComputeTaskResponseBodyData {
	s.DatasetIds = &v
	return s
}

func (s *ListComputeTaskResponseBodyData) SetExtInfo(v string) *ListComputeTaskResponseBodyData {
	s.ExtInfo = &v
	return s
}

func (s *ListComputeTaskResponseBodyData) SetFileNum(v int64) *ListComputeTaskResponseBodyData {
	s.FileNum = &v
	return s
}

func (s *ListComputeTaskResponseBodyData) SetName(v string) *ListComputeTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListComputeTaskResponseBodyData) SetRemarks(v string) *ListComputeTaskResponseBodyData {
	s.Remarks = &v
	return s
}

func (s *ListComputeTaskResponseBodyData) SetStatus(v string) *ListComputeTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListComputeTaskResponseBodyData) SetTaskResultList(v []*ListComputeTaskResponseBodyDataTaskResultList) *ListComputeTaskResponseBodyData {
	s.TaskResultList = v
	return s
}

func (s *ListComputeTaskResponseBodyData) Validate() error {
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

type ListComputeTaskResponseBodyDataTaskResultList struct {
	BcId    *int64 `json:"bcId,omitempty" xml:"bcId,omitempty"`
	Code    *int32 `json:"code,omitempty" xml:"code,omitempty"`
	LineNum *int64 `json:"lineNum,omitempty" xml:"lineNum,omitempty"`
}

func (s ListComputeTaskResponseBodyDataTaskResultList) String() string {
	return dara.Prettify(s)
}

func (s ListComputeTaskResponseBodyDataTaskResultList) GoString() string {
	return s.String()
}

func (s *ListComputeTaskResponseBodyDataTaskResultList) GetBcId() *int64 {
	return s.BcId
}

func (s *ListComputeTaskResponseBodyDataTaskResultList) GetCode() *int32 {
	return s.Code
}

func (s *ListComputeTaskResponseBodyDataTaskResultList) GetLineNum() *int64 {
	return s.LineNum
}

func (s *ListComputeTaskResponseBodyDataTaskResultList) SetBcId(v int64) *ListComputeTaskResponseBodyDataTaskResultList {
	s.BcId = &v
	return s
}

func (s *ListComputeTaskResponseBodyDataTaskResultList) SetCode(v int32) *ListComputeTaskResponseBodyDataTaskResultList {
	s.Code = &v
	return s
}

func (s *ListComputeTaskResponseBodyDataTaskResultList) SetLineNum(v int64) *ListComputeTaskResponseBodyDataTaskResultList {
	s.LineNum = &v
	return s
}

func (s *ListComputeTaskResponseBodyDataTaskResultList) Validate() error {
	return dara.Validate(s)
}

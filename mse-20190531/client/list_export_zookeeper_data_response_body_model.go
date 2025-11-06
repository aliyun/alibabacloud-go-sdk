// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExportZookeeperDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListExportZookeeperDataResponseBodyData) *ListExportZookeeperDataResponseBody
	GetData() []*ListExportZookeeperDataResponseBodyData
	SetDynamicMessage(v string) *ListExportZookeeperDataResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *ListExportZookeeperDataResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v string) *ListExportZookeeperDataResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *ListExportZookeeperDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListExportZookeeperDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListExportZookeeperDataResponseBody
	GetSuccess() *bool
}

type ListExportZookeeperDataResponseBody struct {
	// The details of the data.
	Data []*ListExportZookeeperDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
	//
	// > If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8BD1E58D-0755-42AC-A599-E6B55112****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListExportZookeeperDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExportZookeeperDataResponseBody) GoString() string {
	return s.String()
}

func (s *ListExportZookeeperDataResponseBody) GetData() []*ListExportZookeeperDataResponseBodyData {
	return s.Data
}

func (s *ListExportZookeeperDataResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListExportZookeeperDataResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListExportZookeeperDataResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ListExportZookeeperDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListExportZookeeperDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExportZookeeperDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListExportZookeeperDataResponseBody) SetData(v []*ListExportZookeeperDataResponseBodyData) *ListExportZookeeperDataResponseBody {
	s.Data = v
	return s
}

func (s *ListExportZookeeperDataResponseBody) SetDynamicMessage(v string) *ListExportZookeeperDataResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListExportZookeeperDataResponseBody) SetErrorCode(v string) *ListExportZookeeperDataResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListExportZookeeperDataResponseBody) SetHttpStatusCode(v string) *ListExportZookeeperDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListExportZookeeperDataResponseBody) SetMessage(v string) *ListExportZookeeperDataResponseBody {
	s.Message = &v
	return s
}

func (s *ListExportZookeeperDataResponseBody) SetRequestId(v string) *ListExportZookeeperDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExportZookeeperDataResponseBody) SetSuccess(v bool) *ListExportZookeeperDataResponseBody {
	s.Success = &v
	return s
}

func (s *ListExportZookeeperDataResponseBody) Validate() error {
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

type ListExportZookeeperDataResponseBodyData struct {
	// The details of the task.
	//
	// example:
	//
	// {}
	ContentMap *string `json:"ContentMap,omitempty" xml:"ContentMap,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 1631001140913
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of the object that is exported. Valid values:
	//
	// 	- transactionLog: transaction logs
	//
	// 	- snapshot: snapshots
	//
	// example:
	//
	// snapshot
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// The extension information that is in the JSON format. The extension information facilitates addition of parameters.
	//
	// example:
	//
	// {}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the associated task at the underlying layer. This parameter is used only to troubleshoot failures.
	//
	// example:
	//
	// 10
	KubeoneTaskIds *string `json:"KubeoneTaskIds,omitempty" xml:"KubeoneTaskIds,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- CREATE: The task is being created.
	//
	// 	- RUNNING: The task is being executed.
	//
	// 	- FINISH: The task is completed.
	//
	// 	- FAILED: The task failed.
	//
	// 	- EXPIRE: The task has expired.
	//
	// example:
	//
	// FINISH
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the task was updated.
	//
	// example:
	//
	// 1632979237663
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListExportZookeeperDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListExportZookeeperDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExportZookeeperDataResponseBodyData) GetContentMap() *string {
	return s.ContentMap
}

func (s *ListExportZookeeperDataResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListExportZookeeperDataResponseBodyData) GetExportType() *string {
	return s.ExportType
}

func (s *ListExportZookeeperDataResponseBodyData) GetExtend() *string {
	return s.Extend
}

func (s *ListExportZookeeperDataResponseBodyData) GetId() *int32 {
	return s.Id
}

func (s *ListExportZookeeperDataResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListExportZookeeperDataResponseBodyData) GetKubeoneTaskIds() *string {
	return s.KubeoneTaskIds
}

func (s *ListExportZookeeperDataResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListExportZookeeperDataResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListExportZookeeperDataResponseBodyData) SetContentMap(v string) *ListExportZookeeperDataResponseBodyData {
	s.ContentMap = &v
	return s
}

func (s *ListExportZookeeperDataResponseBodyData) SetCreateTime(v int64) *ListExportZookeeperDataResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListExportZookeeperDataResponseBodyData) SetExportType(v string) *ListExportZookeeperDataResponseBodyData {
	s.ExportType = &v
	return s
}

func (s *ListExportZookeeperDataResponseBodyData) SetExtend(v string) *ListExportZookeeperDataResponseBodyData {
	s.Extend = &v
	return s
}

func (s *ListExportZookeeperDataResponseBodyData) SetId(v int32) *ListExportZookeeperDataResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListExportZookeeperDataResponseBodyData) SetInstanceId(v string) *ListExportZookeeperDataResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListExportZookeeperDataResponseBodyData) SetKubeoneTaskIds(v string) *ListExportZookeeperDataResponseBodyData {
	s.KubeoneTaskIds = &v
	return s
}

func (s *ListExportZookeeperDataResponseBodyData) SetStatus(v string) *ListExportZookeeperDataResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListExportZookeeperDataResponseBodyData) SetUpdateTime(v int64) *ListExportZookeeperDataResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListExportZookeeperDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}

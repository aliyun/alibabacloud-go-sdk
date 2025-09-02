// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRefDISyncTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListRefDISyncTasksResponseBodyData) *ListRefDISyncTasksResponseBody
	GetData() *ListRefDISyncTasksResponseBodyData
	SetRequestId(v string) *ListRefDISyncTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRefDISyncTasksResponseBody
	GetSuccess() *bool
}

type ListRefDISyncTasksResponseBody struct {
	// The returned result.
	Data *ListRefDISyncTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 0bc1411515937635973****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRefDISyncTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRefDISyncTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListRefDISyncTasksResponseBody) GetData() *ListRefDISyncTasksResponseBodyData {
	return s.Data
}

func (s *ListRefDISyncTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRefDISyncTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRefDISyncTasksResponseBody) SetData(v *ListRefDISyncTasksResponseBodyData) *ListRefDISyncTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListRefDISyncTasksResponseBody) SetRequestId(v string) *ListRefDISyncTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRefDISyncTasksResponseBody) SetSuccess(v bool) *ListRefDISyncTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListRefDISyncTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRefDISyncTasksResponseBodyData struct {
	// The details of the synchronization tasks. In most cases, a data source is used by multiple synchronization tasks. Therefore, the value of this parameter is an array. The following parameters are the elements in the array. The sample values of these parameters show the details of a synchronization task.
	DISyncTasks []*ListRefDISyncTasksResponseBodyDataDISyncTasks `json:"DISyncTasks,omitempty" xml:"DISyncTasks,omitempty" type:"Repeated"`
}

func (s ListRefDISyncTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRefDISyncTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRefDISyncTasksResponseBodyData) GetDISyncTasks() []*ListRefDISyncTasksResponseBodyDataDISyncTasks {
	return s.DISyncTasks
}

func (s *ListRefDISyncTasksResponseBodyData) SetDISyncTasks(v []*ListRefDISyncTasksResponseBodyDataDISyncTasks) *ListRefDISyncTasksResponseBodyData {
	s.DISyncTasks = v
	return s
}

func (s *ListRefDISyncTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListRefDISyncTasksResponseBodyDataDISyncTasks struct {
	// The destination of the synchronization task. If the synchronization task has multiple destinations, the return value is a JSON array, such as \\\\"odps_writer\\\\", \\\\"mysql\\\\". If the RefType parameter is set to to, the synchronization tasks that use the specified data source as the destination are returned. In this case, the value of this parameter indicates the specified data source.
	//
	// example:
	//
	// [\\"qcc_polardb2\\"]
	DiDestinationDatasource *string `json:"DiDestinationDatasource,omitempty" xml:"DiDestinationDatasource,omitempty"`
	// The source of the synchronization task. If the synchronization task has multiple sources, the return value is a JSON array, such as \\\\"odps_writer\\\\", \\\\"mysql\\\\". If the RefType parameter is set to from, the synchronization tasks that use the specified data source as the source are returned. In this case, the value of this parameter indicates the specified data source.
	//
	// example:
	//
	// [\\"odps_writer\\"]
	DiSourceDatasource *string `json:"DiSourceDatasource,omitempty" xml:"DiSourceDatasource,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 100000
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the synchronization task.
	//
	// example:
	//
	// abcd1234
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The type of the synchronization task. Valid values:
	//
	// 	- DI_OFFLINE: batch synchronization task
	//
	// 	- DI_REALTIME: real-time synchronization task
	//
	// example:
	//
	// DI_OFFLINE
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListRefDISyncTasksResponseBodyDataDISyncTasks) String() string {
	return dara.Prettify(s)
}

func (s ListRefDISyncTasksResponseBodyDataDISyncTasks) GoString() string {
	return s.String()
}

func (s *ListRefDISyncTasksResponseBodyDataDISyncTasks) GetDiDestinationDatasource() *string {
	return s.DiDestinationDatasource
}

func (s *ListRefDISyncTasksResponseBodyDataDISyncTasks) GetDiSourceDatasource() *string {
	return s.DiSourceDatasource
}

func (s *ListRefDISyncTasksResponseBodyDataDISyncTasks) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListRefDISyncTasksResponseBodyDataDISyncTasks) GetNodeName() *string {
	return s.NodeName
}

func (s *ListRefDISyncTasksResponseBodyDataDISyncTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *ListRefDISyncTasksResponseBodyDataDISyncTasks) SetDiDestinationDatasource(v string) *ListRefDISyncTasksResponseBodyDataDISyncTasks {
	s.DiDestinationDatasource = &v
	return s
}

func (s *ListRefDISyncTasksResponseBodyDataDISyncTasks) SetDiSourceDatasource(v string) *ListRefDISyncTasksResponseBodyDataDISyncTasks {
	s.DiSourceDatasource = &v
	return s
}

func (s *ListRefDISyncTasksResponseBodyDataDISyncTasks) SetNodeId(v int64) *ListRefDISyncTasksResponseBodyDataDISyncTasks {
	s.NodeId = &v
	return s
}

func (s *ListRefDISyncTasksResponseBodyDataDISyncTasks) SetNodeName(v string) *ListRefDISyncTasksResponseBodyDataDISyncTasks {
	s.NodeName = &v
	return s
}

func (s *ListRefDISyncTasksResponseBodyDataDISyncTasks) SetTaskType(v string) *ListRefDISyncTasksResponseBodyDataDISyncTasks {
	s.TaskType = &v
	return s
}

func (s *ListRefDISyncTasksResponseBodyDataDISyncTasks) Validate() error {
	return dara.Validate(s)
}

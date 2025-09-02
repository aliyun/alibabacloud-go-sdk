// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDISyncInstanceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDISyncInstanceInfoResponseBodyData) *GetDISyncInstanceInfoResponseBody
	GetData() *GetDISyncInstanceInfoResponseBodyData
	SetRequestId(v string) *GetDISyncInstanceInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDISyncInstanceInfoResponseBody
	GetSuccess() *bool
}

type GetDISyncInstanceInfoResponseBody struct {
	// The status of the real-time synchronization task or data synchronization solution.
	Data *GetDISyncInstanceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
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

func (s GetDISyncInstanceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDISyncInstanceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDISyncInstanceInfoResponseBody) GetData() *GetDISyncInstanceInfoResponseBodyData {
	return s.Data
}

func (s *GetDISyncInstanceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDISyncInstanceInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDISyncInstanceInfoResponseBody) SetData(v *GetDISyncInstanceInfoResponseBodyData) *GetDISyncInstanceInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDISyncInstanceInfoResponseBody) SetRequestId(v string) *GetDISyncInstanceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDISyncInstanceInfoResponseBody) SetSuccess(v bool) *GetDISyncInstanceInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetDISyncInstanceInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDISyncInstanceInfoResponseBodyData struct {
	// The cause of the failure to obtain the status of the real-time synchronization task or data synchronization solution. If the status of the real-time synchronization task or data synchronization solution is obtained, the value null is returned.
	//
	// example:
	//
	// fileId[100] is invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 	- If the TaskType parameter is set to DI_REALTIME, the Name parameter indicates the name of the real-time synchronization task.
	//
	// 	- If the TaskType parameter is set to DI_SOLUTION, the value null is returned.
	//
	// example:
	//
	// streamx_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 	- If the TaskType parameter is set to DI_REALTIME, the value null is returned.
	//
	// 	- If the TaskType parameter is set to DI_SOLUTION, the SolutionInfo parameter indicates the details of the data synchronization solution.
	SolutionInfo *GetDISyncInstanceInfoResponseBodyDataSolutionInfo `json:"SolutionInfo,omitempty" xml:"SolutionInfo,omitempty" type:"Struct"`
	// 	- If the TaskType parameter is set to DI_REALTIME, the Status parameter indicates the status of the real-time synchronization task. Valid values: PAUSE, NORUN, RUN, KILLING, and WAIT.
	//
	// 	- If the TaskType parameter is set to DI_SOLUTION, the Status parameter indicates the status of the data synchronization solution. Valid values: success and fail.
	//
	// example:
	//
	// RUN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDISyncInstanceInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDISyncInstanceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDISyncInstanceInfoResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GetDISyncInstanceInfoResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetDISyncInstanceInfoResponseBodyData) GetSolutionInfo() *GetDISyncInstanceInfoResponseBodyDataSolutionInfo {
	return s.SolutionInfo
}

func (s *GetDISyncInstanceInfoResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDISyncInstanceInfoResponseBodyData) SetMessage(v string) *GetDISyncInstanceInfoResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetDISyncInstanceInfoResponseBodyData) SetName(v string) *GetDISyncInstanceInfoResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetDISyncInstanceInfoResponseBodyData) SetSolutionInfo(v *GetDISyncInstanceInfoResponseBodyDataSolutionInfo) *GetDISyncInstanceInfoResponseBodyData {
	s.SolutionInfo = v
	return s
}

func (s *GetDISyncInstanceInfoResponseBodyData) SetStatus(v string) *GetDISyncInstanceInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDISyncInstanceInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDISyncInstanceInfoResponseBodyDataSolutionInfo struct {
	// The creator of the data synchronization solution.
	//
	// example:
	//
	// dataworks_3h1
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The ID of the data synchronization solution.
	//
	// example:
	//
	// 100
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The status of the data synchronization solution.
	//
	// example:
	//
	// run
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The step details of the data synchronization solution.
	StepDetail []*GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail `json:"StepDetail,omitempty" xml:"StepDetail,omitempty" type:"Repeated"`
}

func (s GetDISyncInstanceInfoResponseBodyDataSolutionInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDISyncInstanceInfoResponseBodyDataSolutionInfo) GoString() string {
	return s.String()
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfo) GetId() *int64 {
	return s.Id
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfo) GetStatus() *string {
	return s.Status
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfo) GetStepDetail() []*GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail {
	return s.StepDetail
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfo) SetCreatorName(v string) *GetDISyncInstanceInfoResponseBodyDataSolutionInfo {
	s.CreatorName = &v
	return s
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfo) SetId(v int64) *GetDISyncInstanceInfoResponseBodyDataSolutionInfo {
	s.Id = &v
	return s
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfo) SetStatus(v string) *GetDISyncInstanceInfoResponseBodyDataSolutionInfo {
	s.Status = &v
	return s
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfo) SetStepDetail(v []*GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail) *GetDISyncInstanceInfoResponseBodyDataSolutionInfo {
	s.StepDetail = v
	return s
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfo) Validate() error {
	return dara.Validate(s)
}

type GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail struct {
	// The information of the data synchronization solution.
	//
	// example:
	//
	// {\\"fusionProps\\":{\\"dataSource\\":[{\\"fileName\\":\\"holo_20221020161613\\",\\"status\\":\\"SUCCESS\\"}]},\\"customProps\\":{\\"showSN\\":true,\\"columns\\":[{\\"dataIndex\\":\\"fileName\\",\\"width\\":0.8,\\"title\\":\\"Real-time synchronization name\\"},{\\"dataIndex\\":\\"status\\",\\"width\\":0.2,\\"title\\":\\"Status\\"}]},\\"componentName\\":\\"Table\\"}
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	// The status of the step in the data synchronization solution.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the step in the data synchronization solution.
	//
	// example:
	//
	// 1
	StepId *int64 `json:"StepId,omitempty" xml:"StepId,omitempty"`
	// The name of the step in the data synchronization solution.
	//
	// example:
	//
	// Create a base table
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
}

func (s GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail) GoString() string {
	return s.String()
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail) GetInfo() *string {
	return s.Info
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail) GetStatus() *string {
	return s.Status
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail) GetStepId() *int64 {
	return s.StepId
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail) GetStepName() *string {
	return s.StepName
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail) SetInfo(v string) *GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail {
	s.Info = &v
	return s
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail) SetStatus(v string) *GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail {
	s.Status = &v
	return s
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail) SetStepId(v int64) *GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail {
	s.StepId = &v
	return s
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail) SetStepName(v string) *GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail {
	s.StepName = &v
	return s
}

func (s *GetDISyncInstanceInfoResponseBodyDataSolutionInfoStepDetail) Validate() error {
	return dara.Validate(s)
}

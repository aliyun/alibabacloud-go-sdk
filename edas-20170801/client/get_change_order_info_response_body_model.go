// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChangeOrderInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetChangeOrderInfoResponseBody
	GetCode() *int32
	SetMessage(v string) *GetChangeOrderInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetChangeOrderInfoResponseBody
	GetRequestId() *string
	SetChangeOrderInfo(v *GetChangeOrderInfoResponseBodyChangeOrderInfo) *GetChangeOrderInfoResponseBody
	GetChangeOrderInfo() *GetChangeOrderInfoResponseBodyChangeOrderInfo
}

type GetChangeOrderInfoResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4JFR-FV9F***************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the change process.
	ChangeOrderInfo *GetChangeOrderInfoResponseBodyChangeOrderInfo `json:"changeOrderInfo,omitempty" xml:"changeOrderInfo,omitempty" type:"Struct"`
}

func (s GetChangeOrderInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetChangeOrderInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChangeOrderInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChangeOrderInfoResponseBody) GetChangeOrderInfo() *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	return s.ChangeOrderInfo
}

func (s *GetChangeOrderInfoResponseBody) SetCode(v int32) *GetChangeOrderInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetChangeOrderInfoResponseBody) SetMessage(v string) *GetChangeOrderInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetChangeOrderInfoResponseBody) SetRequestId(v string) *GetChangeOrderInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBody) SetChangeOrderInfo(v *GetChangeOrderInfoResponseBodyChangeOrderInfo) *GetChangeOrderInfoResponseBody {
	s.ChangeOrderInfo = v
	return s
}

func (s *GetChangeOrderInfoResponseBody) Validate() error {
	if s.ChangeOrderInfo != nil {
		if err := s.ChangeOrderInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChangeOrderInfoResponseBodyChangeOrderInfo struct {
	// The number of batches for the change.
	//
	// example:
	//
	// 1
	BatchCount *int32 `json:"BatchCount,omitempty" xml:"BatchCount,omitempty"`
	// Indicates whether the change for the next batch is automatically or manually triggered when phased release is performed. Valid values:
	//
	// 	- Automatic: The change for the next batch is automatically triggered.
	//
	// 	- Manual: The change for the next batch is manually triggered.
	//
	// example:
	//
	// Automatic
	BatchType *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	// The description of the change process.
	//
	// example:
	//
	// Application scale-out
	ChangeOrderDescription *string `json:"ChangeOrderDescription,omitempty" xml:"ChangeOrderDescription,omitempty"`
	// The ID of the change process.
	//
	// example:
	//
	// 1074f3e2-e974-4a0e-****-************
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The type of the change process.
	//
	// example:
	//
	// Application Scale Out
	CoType *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	// The time when the change process is created.
	//
	// example:
	//
	// 2019-11-13 14:23:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The person in charge of the change process.
	//
	// example:
	//
	// edas_com***_****@******-*****.***
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// The description of the change process.
	//
	// example:
	//
	// IP of Scale-Out Instance: 47.107.XX.XX
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The information about the batches of the change task.
	PipelineInfoList *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList `json:"PipelineInfoList,omitempty" xml:"PipelineInfoList,omitempty" type:"Struct"`
	// The state of the change process. Valid values:
	//
	// 	- 0: ready
	//
	// 	- 1: in progress
	//
	// 	- 2: successful
	//
	// 	- 3: failed
	//
	// 	- 6: terminated
	//
	// 	- 7: partially executed
	//
	// 	- 8: wait for manual confirmation to trigger the next batch during a manual phased release
	//
	// 	- 9: wait to trigger the next batch during an automatic phased release
	//
	// 	- 10: failed due to a system exception
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether rollbacks are allowed. Valid values:
	//
	// 	- true: Rollbacks are allowed.
	//
	// 	- false: Rollbacks are not allowed.
	//
	// example:
	//
	// false
	SupportRollback *bool                                                 `json:"SupportRollback,omitempty" xml:"SupportRollback,omitempty"`
	Targets         *GetChangeOrderInfoResponseBodyChangeOrderInfoTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Struct"`
	// The throttling rules.
	TrafficControl *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl `json:"TrafficControl,omitempty" xml:"TrafficControl,omitempty" type:"Struct"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfo) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfo) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) GetBatchCount() *int32 {
	return s.BatchCount
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) GetBatchType() *string {
	return s.BatchType
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) GetChangeOrderDescription() *string {
	return s.ChangeOrderDescription
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) GetCoType() *string {
	return s.CoType
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) GetDesc() *string {
	return s.Desc
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) GetPipelineInfoList() *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList {
	return s.PipelineInfoList
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) GetStatus() *int32 {
	return s.Status
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) GetSupportRollback() *bool {
	return s.SupportRollback
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) GetTargets() *GetChangeOrderInfoResponseBodyChangeOrderInfoTargets {
	return s.Targets
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) GetTrafficControl() *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl {
	return s.TrafficControl
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetBatchCount(v int32) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.BatchCount = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetBatchType(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.BatchType = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetChangeOrderDescription(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.ChangeOrderDescription = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetChangeOrderId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.ChangeOrderId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetCoType(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.CoType = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetCreateTime(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.CreateTime = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetCreateUserId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.CreateUserId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetDesc(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.Desc = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetPipelineInfoList(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.PipelineInfoList = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetStatus(v int32) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.Status = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetSupportRollback(v bool) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.SupportRollback = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetTargets(v *GetChangeOrderInfoResponseBodyChangeOrderInfoTargets) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.Targets = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetTrafficControl(v *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.TrafficControl = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) Validate() error {
	if s.PipelineInfoList != nil {
		if err := s.PipelineInfoList.Validate(); err != nil {
			return err
		}
	}
	if s.Targets != nil {
		if err := s.Targets.Validate(); err != nil {
			return err
		}
	}
	if s.TrafficControl != nil {
		if err := s.TrafficControl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList struct {
	PipelineInfo []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo `json:"PipelineInfo,omitempty" xml:"PipelineInfo,omitempty" type:"Repeated"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList) GetPipelineInfo() []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo {
	return s.PipelineInfo
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList) SetPipelineInfo(v []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList {
	s.PipelineInfo = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList) Validate() error {
	if s.PipelineInfo != nil {
		for _, item := range s.PipelineInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo struct {
	// The ID of each batch for the change during the phased release.
	//
	// example:
	//
	// 4c4ee320-5e47-4a48-****-************
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The name of the batch.
	//
	// example:
	//
	// Batch: 1
	PipelineName *string `json:"PipelineName,omitempty" xml:"PipelineName,omitempty"`
	// The state of the change task. Valid values:
	//
	// 	- 0: ready
	//
	// 	- 1: in progress
	//
	// 	- 2: successful
	//
	// 	- 3: failed
	//
	// 	- 6: terminated
	//
	// 	- 8: wait for manual confirmation to trigger the next batch during a manual phased release
	//
	// 	- 9: wait to trigger the next batch during an automatic phased release
	//
	// 	- 10: failed due to a system exception
	//
	// example:
	//
	// 2
	PipelineStatus *int32 `json:"PipelineStatus,omitempty" xml:"PipelineStatus,omitempty"`
	// The execution results in each stage.
	StageDetailList *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList `json:"StageDetailList,omitempty" xml:"StageDetailList,omitempty" type:"Struct"`
	// The stages of the change process.
	StageList *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList `json:"StageList,omitempty" xml:"StageList,omitempty" type:"Struct"`
	// The time when the change task was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1583911702158
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time when the change task was updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1583911743633
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) GetPipelineId() *string {
	return s.PipelineId
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) GetPipelineName() *string {
	return s.PipelineName
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) GetPipelineStatus() *int32 {
	return s.PipelineStatus
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) GetStageDetailList() *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList {
	return s.StageDetailList
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) GetStageList() *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList {
	return s.StageList
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) SetPipelineId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo {
	s.PipelineId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) SetPipelineName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo {
	s.PipelineName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) SetPipelineStatus(v int32) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo {
	s.PipelineStatus = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) SetStageDetailList(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo {
	s.StageDetailList = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) SetStageList(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo {
	s.StageList = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) SetStartTime(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo {
	s.StartTime = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) SetUpdateTime(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo {
	s.UpdateTime = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) Validate() error {
	if s.StageDetailList != nil {
		if err := s.StageDetailList.Validate(); err != nil {
			return err
		}
	}
	if s.StageList != nil {
		if err := s.StageList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList struct {
	StageDetailDTO []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO `json:"StageDetailDTO,omitempty" xml:"StageDetailDTO,omitempty" type:"Repeated"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList) GetStageDetailDTO() []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO {
	return s.StageDetailDTO
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList) SetStageDetailDTO(v []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList {
	s.StageDetailDTO = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList) Validate() error {
	if s.StageDetailDTO != nil {
		for _, item := range s.StageDetailDTO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO struct {
	// The ID of the stage.
	//
	// example:
	//
	// d7561440-10a6-452f-8a90-62f6e7ec****
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The name of the stage.
	//
	// example:
	//
	// Process Start
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The status of the stage. Valid values:
	//
	// 	- 0: ready
	//
	// 	- 1: in progress
	//
	// 	- 2: successful
	//
	// 	- 3: failed
	//
	// 	- 6: terminated
	//
	// 	- 8: wait for manual confirmation to trigger the next batch during a manual phased release
	//
	// 	- 9: wait to trigger the next batch during an automatic phased release
	//
	// 	- 10: failed due to a system exception
	//
	// example:
	//
	// 2
	StageStatus *int32 `json:"StageStatus,omitempty" xml:"StageStatus,omitempty"`
	// The information about the task.
	TaskList *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Struct"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) GetStageId() *string {
	return s.StageId
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) GetStageName() *string {
	return s.StageName
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) GetStageStatus() *int32 {
	return s.StageStatus
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) GetTaskList() *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList {
	return s.TaskList
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) SetStageId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO {
	s.StageId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) SetStageName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO {
	s.StageName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) SetStageStatus(v int32) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO {
	s.StageStatus = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) SetTaskList(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO {
	s.TaskList = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) Validate() error {
	if s.TaskList != nil {
		if err := s.TaskList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList struct {
	TaskInfoDTO []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO `json:"TaskInfoDTO,omitempty" xml:"TaskInfoDTO,omitempty" type:"Repeated"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList) GetTaskInfoDTO() []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	return s.TaskInfoDTO
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList) SetTaskInfoDTO(v []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList {
	s.TaskInfoDTO = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList) Validate() error {
	if s.TaskInfoDTO != nil {
		for _, item := range s.TaskInfoDTO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO struct {
	// The type of the retry policy. Value 0 indicates no retry, value 1 indicates automatic retry, and value 2 indicates manual retry.
	//
	// example:
	//
	// 0
	RetryType *int32 `json:"RetryType,omitempty" xml:"RetryType,omitempty"`
	// Indicates whether errors that occur in the change process are ignored. Valid values:``
	//
	// 	- true: Errors that occur in the change process are ignored. This parameter can be set to true only when URL health checks are performed.
	//
	// 	- false: Errors that occur in the change process are not ignored.
	//
	// example:
	//
	// false
	ShowManualIgnorance *bool `json:"ShowManualIgnorance,omitempty" xml:"ShowManualIgnorance,omitempty"`
	// Error codes
	//
	// example:
	//
	// 400
	TaskErrorCode *string `json:"TaskErrorCode,omitempty" xml:"TaskErrorCode,omitempty"`
	// Indicates whether the task is error-tolerant. If the task can tolerate errors, the errors that occur in the change process are ignored and the next task is executed.
	//
	// 	- 0: The task is not error-tolerant.
	//
	// 	- 1: The task is error-tolerant.
	//
	// example:
	//
	// 0
	TaskErrorIgnorance *int32 `json:"TaskErrorIgnorance,omitempty" xml:"TaskErrorIgnorance,omitempty"`
	// The error message for the task.
	//
	// example:
	//
	// 400
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// d6d3b934-90a1-4ae8-8cbd-2446003d****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Task information
	//
	// example:
	//
	// [CALLBACK] 2020-03-11 15:28:44.781  requestId: c952ab99-8c5b-4ff1-9412-ae3bf9b1****, message: success
	TaskMessage *string `json:"TaskMessage,omitempty" xml:"TaskMessage,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// Build Image
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The state of the task. Valid values:
	//
	// 	- 0: ready
	//
	// 	- 1: in progress
	//
	// 	- 2: successful
	//
	// 	- 3: failed
	//
	// 	- 6: terminated
	//
	// 	- 8: wait for manual confirmation to trigger the next batch during a manual phased release
	//
	// 	- 9: wait to trigger the next batch during an automatic phased release
	//
	// 	- 10: failed due to a system exception
	//
	// example:
	//
	// 2
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) GetRetryType() *int32 {
	return s.RetryType
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) GetShowManualIgnorance() *bool {
	return s.ShowManualIgnorance
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) GetTaskErrorCode() *string {
	return s.TaskErrorCode
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) GetTaskErrorIgnorance() *int32 {
	return s.TaskErrorIgnorance
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) GetTaskId() *string {
	return s.TaskId
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) GetTaskMessage() *string {
	return s.TaskMessage
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) GetTaskName() *string {
	return s.TaskName
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetRetryType(v int32) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.RetryType = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetShowManualIgnorance(v bool) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.ShowManualIgnorance = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetTaskErrorCode(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.TaskErrorCode = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetTaskErrorIgnorance(v int32) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.TaskErrorIgnorance = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetTaskErrorMessage(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.TaskErrorMessage = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetTaskId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.TaskId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetTaskMessage(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.TaskMessage = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetTaskName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.TaskName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetTaskStatus(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.TaskStatus = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) Validate() error {
	return dara.Validate(s)
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList struct {
	StageInfoDTO []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO `json:"StageInfoDTO,omitempty" xml:"StageInfoDTO,omitempty" type:"Repeated"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList) GetStageInfoDTO() []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO {
	return s.StageInfoDTO
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList) SetStageInfoDTO(v []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList {
	s.StageInfoDTO = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList) Validate() error {
	if s.StageInfoDTO != nil {
		for _, item := range s.StageInfoDTO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO struct {
	// The ID of the stage.
	//
	// example:
	//
	// 358a143f-09a0-45e0-****-************@**_*******_*****
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The name of the stage.
	//
	// example:
	//
	// Scale Out
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The results of the task executed in the stage.
	StageResultDTO *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO `json:"StageResultDTO,omitempty" xml:"StageResultDTO,omitempty" type:"Struct"`
	// The state of the stage. Valid values:
	//
	// 	- 0: ready
	//
	// 	- 1: in progress
	//
	// 	- 2: successful
	//
	// 	- 3: failed
	//
	// 	- 6: terminated
	//
	// 	- 8: wait for manual confirmation to trigger the next batch during a manual phased release
	//
	// 	- 9: wait to trigger the next batch during an automatic phased release
	//
	// 	- 10: failed due to a system exception
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) GetStageId() *string {
	return s.StageId
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) GetStageName() *string {
	return s.StageName
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) GetStageResultDTO() *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO {
	return s.StageResultDTO
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) GetStatus() *int32 {
	return s.Status
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) SetStageId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO {
	s.StageId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) SetStageName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO {
	s.StageName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) SetStageResultDTO(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO {
	s.StageResultDTO = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) SetStatus(v int32) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO {
	s.Status = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) Validate() error {
	if s.StageResultDTO != nil {
		if err := s.StageResultDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO struct {
	// The results of the task executed on each Elastic Compute Service (ECS) instance in each stage.
	InstanceDTOList *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList `json:"InstanceDTOList,omitempty" xml:"InstanceDTOList,omitempty" type:"Struct"`
	// The results of tasks executed in each service-oriented stage.
	ServiceStage *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage `json:"ServiceStage,omitempty" xml:"ServiceStage,omitempty" type:"Struct"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO) GetInstanceDTOList() *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList {
	return s.InstanceDTOList
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO) GetServiceStage() *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage {
	return s.ServiceStage
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO) SetInstanceDTOList(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO {
	s.InstanceDTOList = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO) SetServiceStage(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO {
	s.ServiceStage = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO) Validate() error {
	if s.InstanceDTOList != nil {
		if err := s.InstanceDTOList.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceStage != nil {
		if err := s.ServiceStage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList struct {
	InstanceDTO []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO `json:"InstanceDTO,omitempty" xml:"InstanceDTO,omitempty" type:"Repeated"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList) GetInstanceDTO() []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO {
	return s.InstanceDTO
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList) SetInstanceDTO(v []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList {
	s.InstanceDTO = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList) Validate() error {
	if s.InstanceDTO != nil {
		for _, item := range s.InstanceDTO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO struct {
	// The IP address of the ECS instance.
	//
	// example:
	//
	// 47.XX.XX.12 (Public)<br>***.**.*.**	- (*******)
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// The name of the ECS instance.
	//
	// example:
	//
	// EDAS-scaled
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The results of the task executed on the ECS instance in each stage.
	InstanceStageDTOList *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList `json:"InstanceStageDTOList,omitempty" xml:"InstanceStageDTOList,omitempty" type:"Struct"`
	// The name of the node.
	//
	// example:
	//
	// canary-test
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// The state of the pod.
	//
	// example:
	//
	// In progress
	PodStatus *string `json:"PodStatus,omitempty" xml:"PodStatus,omitempty"`
	// The running state. Valid values:
	//
	// 	- 0: ready
	//
	// 	- 1: in progress
	//
	// 	- 2: successful
	//
	// 	- 3: failed
	//
	// 	- 6: terminated
	//
	// 	- 8: wait for manual confirmation to trigger the next batch during a manual phased release
	//
	// 	- 9: wait to trigger the next batch during an automatic phased release
	//
	// 	- 10: failed due to a system exception
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) GetInstanceIp() *string {
	return s.InstanceIp
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) GetInstanceStageDTOList() *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList {
	return s.InstanceStageDTOList
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) GetPodName() *string {
	return s.PodName
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) GetPodStatus() *string {
	return s.PodStatus
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) GetStatus() *int32 {
	return s.Status
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) SetInstanceIp(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO {
	s.InstanceIp = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) SetInstanceName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO {
	s.InstanceName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) SetInstanceStageDTOList(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO {
	s.InstanceStageDTOList = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) SetPodName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO {
	s.PodName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) SetPodStatus(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO {
	s.PodStatus = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) SetStatus(v int32) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO {
	s.Status = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) Validate() error {
	if s.InstanceStageDTOList != nil {
		if err := s.InstanceStageDTOList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList struct {
	InstanceStageDTO []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO `json:"InstanceStageDTO,omitempty" xml:"InstanceStageDTO,omitempty" type:"Repeated"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList) GetInstanceStageDTO() []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO {
	return s.InstanceStageDTO
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList) SetInstanceStageDTO(v []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList {
	s.InstanceStageDTO = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList) Validate() error {
	if s.InstanceStageDTO != nil {
		for _, item := range s.InstanceStageDTO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO struct {
	// The time when the execution stopped.
	//
	// example:
	//
	// 2020-03-11T07:28:52Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the stage.
	//
	// example:
	//
	// 5dd4c0f2-d81a-406f-****-************
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The information about the stage.
	//
	// example:
	//
	// Pulling image \\"registry-vpc.cn-hangzhou.aliyuncs.com****-user/1172745****_shared_repo:428084d6-265f-****-911a-7eb0d2c3****_15839117****\\
	StageMessage *string `json:"StageMessage,omitempty" xml:"StageMessage,omitempty"`
	// The name of the stage.
	//
	// example:
	//
	// scale out
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The time when the execution was started.
	//
	// example:
	//
	// 2020-03-11T07:28:49Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the stage. Valid values:
	//
	// 	- 0: ready
	//
	// 	- 1: in progress
	//
	// 	- 2: successful
	//
	// 	- 3: failed
	//
	// 	- 6: terminated
	//
	// 	- 8: wait for manual confirmation to trigger the next batch during a manual phased release
	//
	// 	- 9: wait to trigger the next batch during an automatic phased release
	//
	// 	- 10: failed due to a system exception
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) GetStageId() *string {
	return s.StageId
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) GetStageMessage() *string {
	return s.StageMessage
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) GetStageName() *string {
	return s.StageName
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) GetStartTime() *string {
	return s.StartTime
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) GetStatus() *int32 {
	return s.Status
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) SetFinishTime(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO {
	s.FinishTime = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) SetStageId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO {
	s.StageId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) SetStageMessage(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO {
	s.StageMessage = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) SetStageName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO {
	s.StageName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) SetStartTime(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO {
	s.StartTime = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) SetStatus(v int32) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO {
	s.Status = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) Validate() error {
	return dara.Validate(s)
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage struct {
	// The execution result in the stage.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the stage.
	//
	// example:
	//
	// 358a143f-09a0-45e0-****-************
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// Phase Name
	//
	// example:
	//
	// Enable Tengine
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The running state. Valid values:
	//
	// 	- 0: ready
	//
	// 	- 1: in progress
	//
	// 	- 2: successful
	//
	// 	- 3: failed
	//
	// 	- 6: terminated
	//
	// 	- 8: wait for manual confirmation to trigger the next batch during a manual phased release
	//
	// 	- 9: wait to trigger the next batch during an automatic phased release
	//
	// 	- 10: failed due to a system exception
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) GetMessage() *string {
	return s.Message
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) GetStageId() *string {
	return s.StageId
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) GetStageName() *string {
	return s.StageName
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) GetStatus() *int32 {
	return s.Status
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) SetMessage(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage {
	s.Message = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) SetStageId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage {
	s.StageId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) SetStageName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage {
	s.StageName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) SetStatus(v int32) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage {
	s.Status = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) Validate() error {
	return dara.Validate(s)
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoTargets struct {
	Items []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoTargets) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoTargets) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoTargets) GetItems() []*string {
	return s.Items
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoTargets) SetItems(v []*string) *GetChangeOrderInfoResponseBodyChangeOrderInfoTargets {
	s.Items = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoTargets) Validate() error {
	return dara.Validate(s)
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl struct {
	// The route forwarding policy.
	//
	// example:
	//
	// [{"app":"9c8247da-91b6-42bb-8f99-92a0b9c6f****","type":"GROUP"}]
	Routes *string `json:"Routes,omitempty" xml:"Routes,omitempty"`
	// The traffic routing rules.
	//
	// example:
	//
	// [{"conditionType":"content","conditions":[{"key":"name","operator":"EQ","strategy":"PARAM","values":["jim"]},{"key":"name","operator":"EQ","strategy":"COOKIE","values":["jim"]}],"percent":100,"protocol":"SPRINGCLOUD","triggerPolicy":"AND"}]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The description of throttling rules.
	//
	// example:
	//
	// This canary release batch is complete, and the user has confirmed to proceed to the next batch.
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) GetRoutes() *string {
	return s.Routes
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) GetRules() *string {
	return s.Rules
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) GetTips() *string {
	return s.Tips
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) SetRoutes(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl {
	s.Routes = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) SetRules(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl {
	s.Rules = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) SetTips(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl {
	s.Tips = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) Validate() error {
	return dara.Validate(s)
}

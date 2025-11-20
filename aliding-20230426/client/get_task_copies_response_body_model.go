// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskCopiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetTaskCopiesResponseBodyData) *GetTaskCopiesResponseBody
	GetData() []*GetTaskCopiesResponseBodyData
	SetPageNumber(v int64) *GetTaskCopiesResponseBody
	GetPageNumber() *int64
	SetRequestId(v string) *GetTaskCopiesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetTaskCopiesResponseBody
	GetTotalCount() *int64
	SetVendorRequestId(v string) *GetTaskCopiesResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetTaskCopiesResponseBody
	GetVendorType() *string
}

type GetTaskCopiesResponseBody struct {
	// example:
	//
	// [{}]
	Data []*GetTaskCopiesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetTaskCopiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskCopiesResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskCopiesResponseBody) GetData() []*GetTaskCopiesResponseBodyData {
	return s.Data
}

func (s *GetTaskCopiesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetTaskCopiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskCopiesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetTaskCopiesResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetTaskCopiesResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetTaskCopiesResponseBody) SetData(v []*GetTaskCopiesResponseBodyData) *GetTaskCopiesResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskCopiesResponseBody) SetPageNumber(v int64) *GetTaskCopiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetTaskCopiesResponseBody) SetRequestId(v string) *GetTaskCopiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskCopiesResponseBody) SetTotalCount(v int64) *GetTaskCopiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetTaskCopiesResponseBody) SetVendorRequestId(v string) *GetTaskCopiesResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetTaskCopiesResponseBody) SetVendorType(v string) *GetTaskCopiesResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetTaskCopiesResponseBody) Validate() error {
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

type GetTaskCopiesResponseBodyData struct {
	// example:
	//
	// [ "actxxx" ]
	ActionExecutorId []*string `json:"ActionExecutorId,omitempty" xml:"ActionExecutorId,omitempty" type:"Repeated"`
	// example:
	//
	// [ "name" ]
	ActionExecutorName []*string `json:"ActionExecutorName,omitempty" xml:"ActionExecutorName,omitempty" type:"Repeated"`
	// example:
	//
	// APP_XCxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// g02fbs08xxx
	CarbonActivityId *string `json:"CarbonActivityId,omitempty" xml:"CarbonActivityId,omitempty"`
	// example:
	//
	// 2020-01-01
	CreateTimeGMT *string `json:"CreateTimeGMT,omitempty" xml:"CreateTimeGMT,omitempty"`
	// example:
	//
	// []
	CurrentActivityInstances []*GetTaskCopiesResponseBodyDataCurrentActivityInstances `json:"CurrentActivityInstances,omitempty" xml:"CurrentActivityInstances,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	DataMap map[string]interface{} `json:"DataMap,omitempty" xml:"DataMap,omitempty"`
	// example:
	//
	// edit
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// 2020-01-01
	FinishTimeGMT *string `json:"FinishTimeGMT,omitempty" xml:"FinishTimeGMT,omitempty"`
	// example:
	//
	// formxxxx
	FormInstanceId *string `json:"FormInstanceId,omitempty" xml:"FormInstanceId,omitempty"`
	// example:
	//
	// uuid
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// value
	InstanceValue *string `json:"InstanceValue,omitempty" xml:"InstanceValue,omitempty"`
	// example:
	//
	// 2020-01-01
	ModifiedTimeGMT *string `json:"ModifiedTimeGMT,omitempty" xml:"ModifiedTimeGMT,omitempty"`
	// example:
	//
	// guyagsd
	OriginatorAvatar *string `json:"OriginatorAvatar,omitempty" xml:"OriginatorAvatar,omitempty"`
	// example:
	//
	// guyagsd
	OriginatorDisplayName *string `json:"OriginatorDisplayName,omitempty" xml:"OriginatorDisplayName,omitempty"`
	// example:
	//
	// 123456
	OriginatorId *string `json:"OriginatorId,omitempty" xml:"OriginatorId,omitempty"`
	// example:
	//
	// 同意
	ProcessApprovedResult *string `json:"ProcessApprovedResult,omitempty" xml:"ProcessApprovedResult,omitempty"`
	// example:
	//
	// 同意
	ProcessApprovedResultText *string `json:"ProcessApprovedResultText,omitempty" xml:"ProcessApprovedResultText,omitempty"`
	// example:
	//
	// code
	ProcessCode *string `json:"ProcessCode,omitempty" xml:"ProcessCode,omitempty"`
	// example:
	//
	// processxxxx
	ProcessId *int64 `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// example:
	//
	// instancexxxx
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// example:
	//
	// 同意
	ProcessInstanceStatus *string `json:"ProcessInstanceStatus,omitempty" xml:"ProcessInstanceStatus,omitempty"`
	// example:
	//
	// 同意
	ProcessInstanceStatusText *string `json:"ProcessInstanceStatusText,omitempty" xml:"ProcessInstanceStatusText,omitempty"`
	// example:
	//
	// 名称
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// example:
	//
	// 12345
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 1.0
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetTaskCopiesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTaskCopiesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskCopiesResponseBodyData) GetActionExecutorId() []*string {
	return s.ActionExecutorId
}

func (s *GetTaskCopiesResponseBodyData) GetActionExecutorName() []*string {
	return s.ActionExecutorName
}

func (s *GetTaskCopiesResponseBodyData) GetAppType() *string {
	return s.AppType
}

func (s *GetTaskCopiesResponseBodyData) GetCarbonActivityId() *string {
	return s.CarbonActivityId
}

func (s *GetTaskCopiesResponseBodyData) GetCreateTimeGMT() *string {
	return s.CreateTimeGMT
}

func (s *GetTaskCopiesResponseBodyData) GetCurrentActivityInstances() []*GetTaskCopiesResponseBodyDataCurrentActivityInstances {
	return s.CurrentActivityInstances
}

func (s *GetTaskCopiesResponseBodyData) GetDataMap() map[string]interface{} {
	return s.DataMap
}

func (s *GetTaskCopiesResponseBodyData) GetDataType() *string {
	return s.DataType
}

func (s *GetTaskCopiesResponseBodyData) GetFinishTimeGMT() *string {
	return s.FinishTimeGMT
}

func (s *GetTaskCopiesResponseBodyData) GetFormInstanceId() *string {
	return s.FormInstanceId
}

func (s *GetTaskCopiesResponseBodyData) GetFormUuid() *string {
	return s.FormUuid
}

func (s *GetTaskCopiesResponseBodyData) GetInstanceValue() *string {
	return s.InstanceValue
}

func (s *GetTaskCopiesResponseBodyData) GetModifiedTimeGMT() *string {
	return s.ModifiedTimeGMT
}

func (s *GetTaskCopiesResponseBodyData) GetOriginatorAvatar() *string {
	return s.OriginatorAvatar
}

func (s *GetTaskCopiesResponseBodyData) GetOriginatorDisplayName() *string {
	return s.OriginatorDisplayName
}

func (s *GetTaskCopiesResponseBodyData) GetOriginatorId() *string {
	return s.OriginatorId
}

func (s *GetTaskCopiesResponseBodyData) GetProcessApprovedResult() *string {
	return s.ProcessApprovedResult
}

func (s *GetTaskCopiesResponseBodyData) GetProcessApprovedResultText() *string {
	return s.ProcessApprovedResultText
}

func (s *GetTaskCopiesResponseBodyData) GetProcessCode() *string {
	return s.ProcessCode
}

func (s *GetTaskCopiesResponseBodyData) GetProcessId() *int64 {
	return s.ProcessId
}

func (s *GetTaskCopiesResponseBodyData) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetTaskCopiesResponseBodyData) GetProcessInstanceStatus() *string {
	return s.ProcessInstanceStatus
}

func (s *GetTaskCopiesResponseBodyData) GetProcessInstanceStatusText() *string {
	return s.ProcessInstanceStatusText
}

func (s *GetTaskCopiesResponseBodyData) GetProcessName() *string {
	return s.ProcessName
}

func (s *GetTaskCopiesResponseBodyData) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *GetTaskCopiesResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetTaskCopiesResponseBodyData) GetVersion() *int64 {
	return s.Version
}

func (s *GetTaskCopiesResponseBodyData) SetActionExecutorId(v []*string) *GetTaskCopiesResponseBodyData {
	s.ActionExecutorId = v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetActionExecutorName(v []*string) *GetTaskCopiesResponseBodyData {
	s.ActionExecutorName = v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetAppType(v string) *GetTaskCopiesResponseBodyData {
	s.AppType = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetCarbonActivityId(v string) *GetTaskCopiesResponseBodyData {
	s.CarbonActivityId = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetCreateTimeGMT(v string) *GetTaskCopiesResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetCurrentActivityInstances(v []*GetTaskCopiesResponseBodyDataCurrentActivityInstances) *GetTaskCopiesResponseBodyData {
	s.CurrentActivityInstances = v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetDataMap(v map[string]interface{}) *GetTaskCopiesResponseBodyData {
	s.DataMap = v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetDataType(v string) *GetTaskCopiesResponseBodyData {
	s.DataType = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetFinishTimeGMT(v string) *GetTaskCopiesResponseBodyData {
	s.FinishTimeGMT = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetFormInstanceId(v string) *GetTaskCopiesResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetFormUuid(v string) *GetTaskCopiesResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetInstanceValue(v string) *GetTaskCopiesResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetModifiedTimeGMT(v string) *GetTaskCopiesResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetOriginatorAvatar(v string) *GetTaskCopiesResponseBodyData {
	s.OriginatorAvatar = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetOriginatorDisplayName(v string) *GetTaskCopiesResponseBodyData {
	s.OriginatorDisplayName = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetOriginatorId(v string) *GetTaskCopiesResponseBodyData {
	s.OriginatorId = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessApprovedResult(v string) *GetTaskCopiesResponseBodyData {
	s.ProcessApprovedResult = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessApprovedResultText(v string) *GetTaskCopiesResponseBodyData {
	s.ProcessApprovedResultText = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessCode(v string) *GetTaskCopiesResponseBodyData {
	s.ProcessCode = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessId(v int64) *GetTaskCopiesResponseBodyData {
	s.ProcessId = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessInstanceId(v string) *GetTaskCopiesResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessInstanceStatus(v string) *GetTaskCopiesResponseBodyData {
	s.ProcessInstanceStatus = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessInstanceStatusText(v string) *GetTaskCopiesResponseBodyData {
	s.ProcessInstanceStatusText = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessName(v string) *GetTaskCopiesResponseBodyData {
	s.ProcessName = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetSerialNumber(v string) *GetTaskCopiesResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetTitle(v string) *GetTaskCopiesResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetVersion(v int64) *GetTaskCopiesResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) Validate() error {
	if s.CurrentActivityInstances != nil {
		for _, item := range s.CurrentActivityInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTaskCopiesResponseBodyDataCurrentActivityInstances struct {
	// example:
	//
	// act-xxaanfaf
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// running
	ActivityInstanceStatus *string `json:"ActivityInstanceStatus,omitempty" xml:"ActivityInstanceStatus,omitempty"`
	// example:
	//
	// act-12345
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// example:
	//
	// redirect task
	ActivityNameInEnglish *string `json:"ActivityNameInEnglish,omitempty" xml:"ActivityNameInEnglish,omitempty"`
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskCopiesResponseBodyDataCurrentActivityInstances) String() string {
	return dara.Prettify(s)
}

func (s GetTaskCopiesResponseBodyDataCurrentActivityInstances) GoString() string {
	return s.String()
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) GetActivityId() *string {
	return s.ActivityId
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) GetActivityInstanceStatus() *string {
	return s.ActivityInstanceStatus
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) GetActivityName() *string {
	return s.ActivityName
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) GetActivityNameInEnglish() *string {
	return s.ActivityNameInEnglish
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) GetId() *int64 {
	return s.Id
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) SetActivityId(v string) *GetTaskCopiesResponseBodyDataCurrentActivityInstances {
	s.ActivityId = &v
	return s
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) SetActivityInstanceStatus(v string) *GetTaskCopiesResponseBodyDataCurrentActivityInstances {
	s.ActivityInstanceStatus = &v
	return s
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) SetActivityName(v string) *GetTaskCopiesResponseBodyDataCurrentActivityInstances {
	s.ActivityName = &v
	return s
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) SetActivityNameInEnglish(v string) *GetTaskCopiesResponseBodyDataCurrentActivityInstances {
	s.ActivityNameInEnglish = &v
	return s
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) SetId(v int64) *GetTaskCopiesResponseBodyDataCurrentActivityInstances {
	s.Id = &v
	return s
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) Validate() error {
	return dara.Validate(s)
}

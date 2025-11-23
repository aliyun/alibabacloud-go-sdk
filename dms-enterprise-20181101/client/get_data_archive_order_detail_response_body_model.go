// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataArchiveOrderDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataArchiveOrderDetail(v *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) *GetDataArchiveOrderDetailResponseBody
	GetDataArchiveOrderDetail() *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail
	SetErrorCode(v string) *GetDataArchiveOrderDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataArchiveOrderDetailResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataArchiveOrderDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataArchiveOrderDetailResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *GetDataArchiveOrderDetailResponseBody
	GetTraceId() *string
}

type GetDataArchiveOrderDetailResponseBody struct {
	// The details of data archiving tickets.
	DataArchiveOrderDetail *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail `json:"DataArchiveOrderDetail,omitempty" xml:"DataArchiveOrderDetail,omitempty" type:"Struct"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// User [19929582****] not exist
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request, which is used to query logs and troubleshoot issues.
	//
	// example:
	//
	// 4161CE36-28DF-5191-8A6F-A17076A0B124
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Tracks service requests.
	//
	// example:
	//
	// 0a06e1e316757357507896067d3780
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s GetDataArchiveOrderDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataArchiveOrderDetailResponseBody) GetDataArchiveOrderDetail() *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	return s.DataArchiveOrderDetail
}

func (s *GetDataArchiveOrderDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataArchiveOrderDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataArchiveOrderDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataArchiveOrderDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataArchiveOrderDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *GetDataArchiveOrderDetailResponseBody) SetDataArchiveOrderDetail(v *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) *GetDataArchiveOrderDetailResponseBody {
	s.DataArchiveOrderDetail = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBody) SetErrorCode(v string) *GetDataArchiveOrderDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBody) SetErrorMessage(v string) *GetDataArchiveOrderDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBody) SetRequestId(v string) *GetDataArchiveOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBody) SetSuccess(v bool) *GetDataArchiveOrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBody) SetTraceId(v string) *GetDataArchiveOrderDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBody) Validate() error {
	if s.DataArchiveOrderDetail != nil {
		if err := s.DataArchiveOrderDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail struct {
	// The description of the data archiving tickets.
	//
	// example:
	//
	// Archiving of test results
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The user who submitted the ticket.
	//
	// example:
	//
	// dmstest
	Committer *string `json:"Committer,omitempty" xml:"Committer,omitempty"`
	// The ID of the user who submitted the ticket. The ID is a user ID and not the ID of an Alibaba Cloud account.
	//
	// example:
	//
	// 26***
	CommitterId *int64 `json:"CommitterId,omitempty" xml:"CommitterId,omitempty"`
	// The time when the ticket was created.
	//
	// example:
	//
	// 2023-05-15 16:00:48
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the ticket was last modified.
	//
	// example:
	//
	// 2023-05-23 16:00:48
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of data archiving tickets.
	//
	// example:
	//
	// 868****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The additional information about the ticket.
	PluginExtraData *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData `json:"PluginExtraData,omitempty" xml:"PluginExtraData,omitempty" type:"Struct"`
	// The ticket creation parameter. The value is a JSON string. For more information, see [PluginType parameter](https://help.aliyun.com/document_detail/429109.html).
	PluginParam *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam `json:"PluginParam,omitempty" xml:"PluginParam,omitempty" type:"Struct"`
	// The plug-in type that corresponds to the type of the ticket. The plug-in type for data archiving is DATA_ARCHIVE. For more information, see [PluginType parameter](https://help.aliyun.com/document_detail/429109.html).
	//
	// example:
	//
	// DATA_ARCHIVE
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	// The user IDs related to the ticket.
	RelatedUserList []*int64 `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	// The nicknames of the users that are related to the ticket.
	RelatedUserNickList []*string `json:"RelatedUserNickList,omitempty" xml:"RelatedUserNickList,omitempty" type:"Repeated"`
	// The status code of the ticket. Valid values:
	//
	// 	- **new**: newly created.
	//
	// 	- **toaudit**: being reviewed.
	//
	// 	- **Approved**: approved.
	//
	// 	- **reject**: rejected.
	//
	// 	- **processing**: being executed.
	//
	// 	- **Success**: successful.
	//
	// 	- **closed**: disabled.
	//
	// example:
	//
	// processing
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The status description of the ticket.
	//
	// example:
	//
	// a ticket task is being executed.
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// The ID of the approval process. You can call the [GetOrderBaseInfo](https://help.aliyun.com/document_detail/144642.html) operation to obtain the ID of the approval process.
	//
	// example:
	//
	// 29****
	WorkflowInstanceId *int64 `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
	// The description of the approval process.
	//
	// example:
	//
	// approved
	WorkflowStatusDesc *string `json:"WorkflowStatusDesc,omitempty" xml:"WorkflowStatusDesc,omitempty"`
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GoString() string {
	return s.String()
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GetComment() *string {
	return s.Comment
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GetCommitter() *string {
	return s.Committer
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GetCommitterId() *int64 {
	return s.CommitterId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GetId() *int64 {
	return s.Id
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GetPluginExtraData() *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData {
	return s.PluginExtraData
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GetPluginParam() *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam {
	return s.PluginParam
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GetPluginType() *string {
	return s.PluginType
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GetRelatedUserList() []*int64 {
	return s.RelatedUserList
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GetRelatedUserNickList() []*string {
	return s.RelatedUserNickList
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GetStatusCode() *string {
	return s.StatusCode
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) GetWorkflowStatusDesc() *string {
	return s.WorkflowStatusDesc
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) SetComment(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	s.Comment = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) SetCommitter(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	s.Committer = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) SetCommitterId(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	s.CommitterId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) SetGmtCreate(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	s.GmtCreate = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) SetGmtModified(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	s.GmtModified = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) SetId(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	s.Id = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) SetPluginExtraData(v *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	s.PluginExtraData = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) SetPluginParam(v *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	s.PluginParam = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) SetPluginType(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	s.PluginType = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) SetRelatedUserList(v []*int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	s.RelatedUserList = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) SetRelatedUserNickList(v []*string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	s.RelatedUserNickList = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) SetStatusCode(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	s.StatusCode = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) SetStatusDesc(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	s.StatusDesc = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) SetWorkflowInstanceId(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	s.WorkflowInstanceId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) SetWorkflowStatusDesc(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail {
	s.WorkflowStatusDesc = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetail) Validate() error {
	if s.PluginExtraData != nil {
		if err := s.PluginExtraData.Validate(); err != nil {
			return err
		}
	}
	if s.PluginParam != nil {
		if err := s.PluginParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData struct {
	// The information about the workflow.
	DagInfo *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo `json:"DagInfo,omitempty" xml:"DagInfo,omitempty" type:"Struct"`
	// The database information related to data archiving tickets.
	DbBaseInfo *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo `json:"DbBaseInfo,omitempty" xml:"DbBaseInfo,omitempty" type:"Struct"`
	// The total number of archiving tasks.
	//
	// example:
	//
	// 2
	InstanceTotal *int64 `json:"InstanceTotal,omitempty" xml:"InstanceTotal,omitempty"`
	// The list of archiving tasks.
	Instances []*GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The time when the next task is triggered.
	NextFireTimeResult *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataNextFireTimeResult `json:"NextFireTimeResult,omitempty" xml:"NextFireTimeResult,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 10
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the temporary table that is generated by the archiving task (indicated by the archiving task ID).
	//
	// example:
	//
	// {
	//
	//       "803***": [
	//
	//             "tmp_dms_21321_20230704144336_temp_test_check"
	//
	//       ]
	//
	// }
	TempTableNameMap map[string]interface{} `json:"TempTableNameMap,omitempty" xml:"TempTableNameMap,omitempty"`
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) GoString() string {
	return s.String()
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) GetDagInfo() *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	return s.DagInfo
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) GetDbBaseInfo() *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	return s.DbBaseInfo
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) GetInstanceTotal() *int64 {
	return s.InstanceTotal
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) GetInstances() []*GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances {
	return s.Instances
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) GetNextFireTimeResult() *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataNextFireTimeResult {
	return s.NextFireTimeResult
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) GetTempTableNameMap() map[string]interface{} {
	return s.TempTableNameMap
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) SetDagInfo(v *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData {
	s.DagInfo = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) SetDbBaseInfo(v *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData {
	s.DbBaseInfo = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) SetInstanceTotal(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData {
	s.InstanceTotal = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) SetInstances(v []*GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData {
	s.Instances = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) SetNextFireTimeResult(v *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataNextFireTimeResult) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData {
	s.NextFireTimeResult = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) SetPageIndex(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData {
	s.PageIndex = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) SetPageSize(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData {
	s.PageSize = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) SetTempTableNameMap(v map[string]interface{}) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData {
	s.TempTableNameMap = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraData) Validate() error {
	if s.DagInfo != nil {
		if err := s.DagInfo.Validate(); err != nil {
			return err
		}
	}
	if s.DbBaseInfo != nil {
		if err := s.DbBaseInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NextFireTimeResult != nil {
		if err := s.NextFireTimeResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo struct {
	// The ID of the user who created the task flow.
	//
	// example:
	//
	// 59****
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The start time for scheduling. The task flow is not scheduled before this point in time.
	//
	// example:
	//
	// 1970-01-01
	CronBeginDate *string `json:"CronBeginDate,omitempty" xml:"CronBeginDate,omitempty"`
	// The end time for scheduling. The task flow is not scheduled after this point in time.
	//
	// example:
	//
	// 9999-01-01
	CronEndDate *string `json:"CronEndDate,omitempty" xml:"CronEndDate,omitempty"`
	// Indicates whether the archiving task is a scheduled task. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CronTrigger *bool `json:"CronTrigger,omitempty" xml:"CronTrigger,omitempty"`
	// Indicates whether the task is used to develop warehouses.
	//
	// >  This field is a retained field that is not in use.
	//
	// example:
	//
	// false
	DWDevelop *bool `json:"DWDevelop,omitempty" xml:"DWDevelop,omitempty"`
	// The name of the workflow.
	//
	// example:
	//
	// data-archive-9099197
	DagName *string `json:"DagName,omitempty" xml:"DagName,omitempty"`
	// The ID of the owner of the workflow.
	//
	// example:
	//
	// 13****
	DagOwnerId *string `json:"DagOwnerId,omitempty" xml:"DagOwnerId,omitempty"`
	// The ID of the deployment record.
	//
	// example:
	//
	// 93***
	DeployId *int64 `json:"DeployId,omitempty" xml:"DeployId,omitempty"`
	// The description of the workflow.
	//
	// example:
	//
	// order id:9099197
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the editable workflow version.
	//
	// example:
	//
	// 24***
	EditDagId *int64 `json:"EditDagId,omitempty" xml:"EditDagId,omitempty"`
	// The time when the workflow was created.
	//
	// example:
	//
	// 2023-05-15 16:00:48
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the workflow was last modified.
	//
	// example:
	//
	// 2023-06-15 16:00:48
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the task flow.
	//
	// example:
	//
	// 24***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the workflow is public. Valid values:
	//
	// 	- **0**: not public.
	//
	// 	- **1**: public.
	//
	// example:
	//
	// 0
	IsPublic *int64 `json:"IsPublic,omitempty" xml:"IsPublic,omitempty"`
	// Indicates whether the task is a historical task. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Legacy *bool `json:"Legacy,omitempty" xml:"Legacy,omitempty"`
	// Indicates whether the task was created by the system. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 5***
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Indicates whether the workflow is triggered to run once. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	TriggerOnce *bool `json:"TriggerOnce,omitempty" xml:"TriggerOnce,omitempty"`
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GoString() string {
	return s.String()
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetCronBeginDate() *string {
	return s.CronBeginDate
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetCronEndDate() *string {
	return s.CronEndDate
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetCronTrigger() *bool {
	return s.CronTrigger
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetDWDevelop() *bool {
	return s.DWDevelop
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetDagName() *string {
	return s.DagName
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetDagOwnerId() *string {
	return s.DagOwnerId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetDeployId() *int64 {
	return s.DeployId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetDescription() *string {
	return s.Description
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetEditDagId() *int64 {
	return s.EditDagId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetId() *int64 {
	return s.Id
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetIsPublic() *int64 {
	return s.IsPublic
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetLegacy() *bool {
	return s.Legacy
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetSystem() *bool {
	return s.System
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) GetTriggerOnce() *bool {
	return s.TriggerOnce
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetCreatorId(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.CreatorId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetCronBeginDate(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.CronBeginDate = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetCronEndDate(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.CronEndDate = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetCronTrigger(v bool) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.CronTrigger = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetDWDevelop(v bool) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.DWDevelop = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetDagName(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.DagName = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetDagOwnerId(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.DagOwnerId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetDeployId(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.DeployId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetDescription(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.Description = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetEditDagId(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.EditDagId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetGmtCreate(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetGmtModified(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.GmtModified = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetId(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.Id = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetIsPublic(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.IsPublic = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetLegacy(v bool) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.Legacy = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetSystem(v bool) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.System = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetTenantId(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.TenantId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) SetTriggerOnce(v bool) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo {
	s.TriggerOnce = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDagInfo) Validate() error {
	return dara.Validate(s)
}

type GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo struct {
	// The alias of the database instance.
	//
	// example:
	//
	// tf-testAccDMSEnterpriseLogicDatabase853****
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The timeout period of queries on the database.
	//
	// example:
	//
	// 600
	AlterTimeout *int64 `json:"AlterTimeout,omitempty" xml:"AlterTimeout,omitempty"`
	// Indicates whether access control is enabled for data assets. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AssetControl *bool `json:"AssetControl,omitempty" xml:"AssetControl,omitempty"`
	// The name of the instance in the instance list.
	//
	// example:
	//
	// test
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// Indicates whether the instance is added to the DMS whitelist.
	//
	// example:
	//
	// whitelist_done
	ClusterNode *string `json:"ClusterNode,omitempty" xml:"ClusterNode,omitempty"`
	// The ID of the database. You can call the [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to query the ID of the database.
	//
	// >  You can call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) operation to query the ID of a physical database or the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) operation to query the ID of a logical database.
	//
	// example:
	//
	// 348****
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The type of the database. For information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// example:
	//
	// MySQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The ID of the database administrator (DBA) of the instance.
	//
	// example:
	//
	// 16****
	DbaId *int64 `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	// The nickname of the DBA of the instance.
	//
	// example:
	//
	// DBA
	DbaName *string `json:"DbaName,omitempty" xml:"DbaName,omitempty"`
	// The complete endpoint of the database.
	//
	// example:
	//
	// test@rm-2ze756u8837****.mysql.rds.aliyuncs.com:3306 [test]
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The encoding format of the database.
	//
	// example:
	//
	// utf8
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// The type of the environment to which the database belongs. Valid values:
	//
	// 	- **product**: production environment
	//
	// 	- **dev**: development environment
	//
	// 	- **pre**: staging environment
	//
	// 	- **test**: test environment
	//
	// 	- **sit**: system integration testing (SIT) environment
	//
	// 	- **uat**: user acceptance testing (UAT) environment
	//
	// 	- **pet**: stress testing environment
	//
	// 	- **stag**: STAG environment
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// Indicates whether the instance needs special attention. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Follow *bool `json:"Follow,omitempty" xml:"Follow,omitempty"`
	// The endpoint that is used to connect to the database.
	//
	// example:
	//
	// rm-2ze756u8837****.mysql.rds.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The region in which the database instance resides.
	//
	// example:
	//
	// cn-beijing
	Idc *string `json:"Idc,omitempty" xml:"Idc,omitempty"`
	// The name of the region in which the database instance resides.
	//
	// example:
	//
	// cn-beijing
	IdcTitle *string `json:"IdcTitle,omitempty" xml:"IdcTitle,omitempty"`
	// The ID of the instance to which the database belongs.
	//
	// example:
	//
	// 175****
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The source of the database instance.Valid values:
	//
	// 	- **RDS**: an ApsaraDB RDS instance.
	//
	// 	- **ECS_OWN**: a self-managed database deployed on an Elastic Compute Service (ECS) instance.
	//
	// 	- **PUBLIC_OWN**: a self-managed database instance that is connected over the Internet.
	//
	// 	- **VPC_ID**: a self-managed database instance in a virtual private cloud (VPC) that is connected over Express Connect circuits.
	//
	// 	- **GATEWAY**: a database instance connected by using a database gateway.
	//
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The time when the database information was last obtained.
	//
	// example:
	//
	// 2023-05-14 18:34:45
	LastSyncTime *string `json:"LastSyncTime,omitempty" xml:"LastSyncTime,omitempty"`
	// The instance level.
	//
	// example:
	//
	// medium
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Indicates whether the database is logical. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The IDs of the owners of the databases, which are stored as an array. You can call the [GetUser](https://help.aliyun.com/document_detail/147098.html) or [ListUsers](https://help.aliyun.com/document_detail/141938.html) operation to query the IDs of the owners.
	//
	// >  The value of OwnerIds is the same as the value of UserId
	OwnerIds []*int64 `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
	// The usernames of the database owners.
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
	// The port that is used to connect to the database.
	//
	// example:
	//
	// 3306
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name that is used to search for the database.
	//
	// example:
	//
	// test@rm-2ze756u8837****.mysql.rds.aliyuncs.com:3306 [test]
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The details of the control mode of the instance.
	StandardGroup *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup `json:"StandardGroup,omitempty" xml:"StandardGroup,omitempty" type:"Struct"`
	// The status of the database. Valid values:
	//
	// 	- **NORMAL**: The database is running as expected.
	//
	// 	- **DISABLE**: The database is disabled.
	//
	// 	- **OFFLINE**: The database is unpublished.
	//
	// 	- **NOT_EXIST**: The database does not exist.
	//
	// example:
	//
	// NORMAL
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The number of tables.
	//
	// example:
	//
	// 201
	TableCount *int64 `json:"TableCount,omitempty" xml:"TableCount,omitempty"`
	// The name of TNS.
	//
	// example:
	//
	// TNS_4010
	TnsName *string `json:"TnsName,omitempty" xml:"TnsName,omitempty"`
	// The unit type.
	//
	// example:
	//
	// -1
	UnitType *string `json:"UnitType,omitempty" xml:"UnitType,omitempty"`
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GoString() string {
	return s.String()
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetAlias() *string {
	return s.Alias
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetAlterTimeout() *int64 {
	return s.AlterTimeout
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetAssetControl() *bool {
	return s.AssetControl
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetCatalogName() *string {
	return s.CatalogName
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetClusterNode() *string {
	return s.ClusterNode
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetDbId() *int64 {
	return s.DbId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetDbType() *string {
	return s.DbType
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetDbaId() *int64 {
	return s.DbaId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetDbaName() *string {
	return s.DbaName
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetDescription() *string {
	return s.Description
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetEncoding() *string {
	return s.Encoding
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetEnvType() *string {
	return s.EnvType
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetFollow() *bool {
	return s.Follow
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetHost() *string {
	return s.Host
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetIdc() *string {
	return s.Idc
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetIdcTitle() *string {
	return s.IdcTitle
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetLastSyncTime() *string {
	return s.LastSyncTime
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetLevel() *string {
	return s.Level
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetLogic() *bool {
	return s.Logic
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetOwnerIds() []*int64 {
	return s.OwnerIds
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetOwnerNames() []*string {
	return s.OwnerNames
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetPort() *int64 {
	return s.Port
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetSearchName() *string {
	return s.SearchName
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetStandardGroup() *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup {
	return s.StandardGroup
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetState() *string {
	return s.State
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetTableCount() *int64 {
	return s.TableCount
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetTnsName() *string {
	return s.TnsName
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) GetUnitType() *string {
	return s.UnitType
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetAlias(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.Alias = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetAlterTimeout(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.AlterTimeout = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetAssetControl(v bool) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.AssetControl = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetCatalogName(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.CatalogName = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetClusterNode(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.ClusterNode = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetDbId(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.DbId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetDbType(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.DbType = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetDbaId(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.DbaId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetDbaName(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.DbaName = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetDescription(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.Description = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetEncoding(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.Encoding = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetEnvType(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.EnvType = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetFollow(v bool) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.Follow = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetHost(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.Host = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetIdc(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.Idc = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetIdcTitle(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.IdcTitle = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetInstanceId(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.InstanceId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetInstanceSource(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.InstanceSource = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetLastSyncTime(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.LastSyncTime = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetLevel(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.Level = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetLogic(v bool) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.Logic = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetOwnerIds(v []*int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.OwnerIds = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetOwnerNames(v []*string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.OwnerNames = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetPort(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.Port = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetSchemaName(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.SchemaName = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetSearchName(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.SearchName = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetStandardGroup(v *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.StandardGroup = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetState(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.State = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetTableCount(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.TableCount = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetTnsName(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.TnsName = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) SetUnitType(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo {
	s.UnitType = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfo) Validate() error {
	if s.StandardGroup != nil {
		if err := s.StandardGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup struct {
	// The type of the instance engine. For information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// example:
	//
	// MySQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The description of the security rule set.
	//
	// example:
	//
	// adb_mysql default
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the instance is managed in Flexible Management or Stable Change mode. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	FreeOrStable *bool `json:"FreeOrStable,omitempty" xml:"FreeOrStable,omitempty"`
	// The time when the security rule was created.
	//
	// example:
	//
	// 2020-05-24 14:12:32
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the security rule was last modified.
	//
	// example:
	//
	// 2020-05-25 14:12:32
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The type of the control mode of the instance. Valid values:
	//
	// 	- **COMMON**: The instance is managed in Security Collaboration mode.
	//
	// 	- **NONE_CONTROL**: The instance is managed in Flexible Management mode.
	//
	// 	- **STABLE**: The instance is managed in Stable Change mode.
	//
	// example:
	//
	// COMMON
	GroupMode *string `json:"GroupMode,omitempty" xml:"GroupMode,omitempty"`
	// The name of the security rule that corresponds to the control mode.
	//
	// example:
	//
	// adb_mysql default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the security rule.
	//
	// example:
	//
	// 24***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The user ID of the last modified security rule.
	//
	// example:
	//
	// 12****
	LastMenderId *int64 `json:"LastMenderId,omitempty" xml:"LastMenderId,omitempty"`
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) GoString() string {
	return s.String()
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) GetDbType() *string {
	return s.DbType
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) GetDescription() *string {
	return s.Description
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) GetFreeOrStable() *bool {
	return s.FreeOrStable
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) GetGroupMode() *string {
	return s.GroupMode
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) GetId() *int64 {
	return s.Id
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) GetLastMenderId() *int64 {
	return s.LastMenderId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) SetDbType(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup {
	s.DbType = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) SetDescription(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup {
	s.Description = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) SetFreeOrStable(v bool) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup {
	s.FreeOrStable = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) SetGmtCreate(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup {
	s.GmtCreate = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) SetGmtModified(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup {
	s.GmtModified = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) SetGroupMode(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup {
	s.GroupMode = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) SetGroupName(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup {
	s.GroupName = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) SetId(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup {
	s.Id = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) SetLastMenderId(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup {
	s.LastMenderId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataDbBaseInfoStandardGroup) Validate() error {
	return dara.Validate(s)
}

type GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances struct {
	// The business time of the task flow. The time is displayed in the yyyy-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-05-14 16:00:57
	BusinessTime *string `json:"BusinessTime,omitempty" xml:"BusinessTime,omitempty"`
	// The task flow ID. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to obtain the value of this parameter.
	//
	// example:
	//
	// 37***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The time when the task flow ended. The time is displayed in the yyyy-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2022-06-04 15:14:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the task flow was created.
	//
	// example:
	//
	// 2023-05-14 16:00:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the task flow was last modified.
	//
	// example:
	//
	// 2023-05-14 16:00:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the historical task flow.
	//
	// example:
	//
	// 32***
	HistoryDagId *int64 `json:"HistoryDagId,omitempty" xml:"HistoryDagId,omitempty"`
	// The ID of the instance in the task flow that is executed.
	//
	// example:
	//
	// 24***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The context of the previous execution of the task flow.
	//
	// example:
	//
	// {
	//
	//       "nodes": [
	//
	//             48**
	//
	//       ],
	//
	//       "edges": {}
	//
	// }
	LastRunningContext *string `json:"LastRunningContext,omitempty" xml:"LastRunningContext,omitempty"`
	// The context of the current execution of the task flow.
	//
	// example:
	//
	// 2023-05-15 16:37:48[GMT+08:00] INFO - Resource Control is active!\\n2023-05-15 16:37:48[GMT+08:00] INFO - Starting job j_4834 at Mon May 15 16:37:48 CST 2023
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **0**: The task is waiting for execution.
	//
	// 	- **1**: The task is in progress.
	//
	// 	- **2**: The task is suspended.
	//
	// 	- **3**: The task failed.
	//
	// 	- **4**: The task is successful.
	//
	// 	- **5**: The task is complete.
	//
	// example:
	//
	// 4
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 5***
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The mode in which the task flow was triggered. Valid values:
	//
	// 	- **0**: The task flow was triggered based on a schedule.
	//
	// 	- **1**: The task flow was manually triggered.
	//
	// example:
	//
	// 1
	TriggerType *int64 `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) GoString() string {
	return s.String()
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) GetBusinessTime() *string {
	return s.BusinessTime
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) GetDagId() *int64 {
	return s.DagId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) GetEndTime() *string {
	return s.EndTime
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) GetHistoryDagId() *int64 {
	return s.HistoryDagId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) GetId() *int64 {
	return s.Id
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) GetLastRunningContext() *string {
	return s.LastRunningContext
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) GetMsg() *string {
	return s.Msg
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) GetStatus() *int64 {
	return s.Status
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) GetTriggerType() *int64 {
	return s.TriggerType
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) GetVersion() *string {
	return s.Version
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) SetBusinessTime(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances {
	s.BusinessTime = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) SetDagId(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances {
	s.DagId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) SetEndTime(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances {
	s.EndTime = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) SetGmtCreate(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances {
	s.GmtCreate = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) SetGmtModified(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances {
	s.GmtModified = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) SetHistoryDagId(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances {
	s.HistoryDagId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) SetId(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances {
	s.Id = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) SetLastRunningContext(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances {
	s.LastRunningContext = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) SetMsg(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances {
	s.Msg = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) SetStatus(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances {
	s.Status = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) SetTenantId(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances {
	s.TenantId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) SetTriggerType(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances {
	s.TriggerType = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) SetVersion(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances {
	s.Version = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataInstances) Validate() error {
	return dara.Validate(s)
}

type GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataNextFireTimeResult struct {
	// The type of scheduled triggering.
	//
	// example:
	//
	// NOT_SET
	CronFireType *string `json:"CronFireType,omitempty" xml:"CronFireType,omitempty"`
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataNextFireTimeResult) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataNextFireTimeResult) GoString() string {
	return s.String()
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataNextFireTimeResult) GetCronFireType() *string {
	return s.CronFireType
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataNextFireTimeResult) SetCronFireType(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataNextFireTimeResult {
	s.CronFireType = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginExtraDataNextFireTimeResult) Validate() error {
	return dara.Validate(s)
}

type GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam struct {
	// The type of the archiving destination.
	//
	// example:
	//
	// inner_oss
	ArchiveMethod *string `json:"ArchiveMethod,omitempty" xml:"ArchiveMethod,omitempty"`
	// The schema of the database and table to be archived.
	//
	// example:
	//
	// test
	DbSchema *string `json:"DbSchema,omitempty" xml:"DbSchema,omitempty"`
	// Indicates whether the database is logical.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The post behavior of archiving.
	OrderAfter []*string `json:"OrderAfter,omitempty" xml:"OrderAfter,omitempty" type:"Repeated"`
	// The running method, which indicates whether to run the task immediately or at a specific point in time.
	//
	// example:
	//
	// now
	RunMethod *string `json:"RunMethod,omitempty" xml:"RunMethod,omitempty"`
	// The ID of the source database.
	//
	// example:
	//
	// 12***
	SourceDatabaseId *int64 `json:"SourceDatabaseId,omitempty" xml:"SourceDatabaseId,omitempty"`
	// The list of the archived tables and the filter conditions.
	TableIncludes []*GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParamTableIncludes `json:"TableIncludes,omitempty" xml:"TableIncludes,omitempty" type:"Repeated"`
	// The mapping of schemas.
	TableMapping []*string `json:"TableMapping,omitempty" xml:"TableMapping,omitempty" type:"Repeated"`
	// The ID of the destination instance.
	//
	// example:
	//
	// 12***
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The time variable defined for scheduled archiving.
	Variables []*string `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) GoString() string {
	return s.String()
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) GetArchiveMethod() *string {
	return s.ArchiveMethod
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) GetDbSchema() *string {
	return s.DbSchema
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) GetLogic() *bool {
	return s.Logic
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) GetOrderAfter() []*string {
	return s.OrderAfter
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) GetRunMethod() *string {
	return s.RunMethod
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) GetSourceDatabaseId() *int64 {
	return s.SourceDatabaseId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) GetTableIncludes() []*GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParamTableIncludes {
	return s.TableIncludes
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) GetTableMapping() []*string {
	return s.TableMapping
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) GetVariables() []*string {
	return s.Variables
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) SetArchiveMethod(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam {
	s.ArchiveMethod = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) SetDbSchema(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam {
	s.DbSchema = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) SetLogic(v bool) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam {
	s.Logic = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) SetOrderAfter(v []*string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam {
	s.OrderAfter = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) SetRunMethod(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam {
	s.RunMethod = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) SetSourceDatabaseId(v int64) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam {
	s.SourceDatabaseId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) SetTableIncludes(v []*GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParamTableIncludes) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam {
	s.TableIncludes = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) SetTableMapping(v []*string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam {
	s.TableMapping = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) SetTargetInstanceId(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam {
	s.TargetInstanceId = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) SetVariables(v []*string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam {
	s.Variables = v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParam) Validate() error {
	if s.TableIncludes != nil {
		for _, item := range s.TableIncludes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParamTableIncludes struct {
	// The table name.
	//
	// example:
	//
	// tm_insured_cb
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The filter condition.
	//
	// example:
	//
	// id<1000 or gmt_create<\\"2023-05-14 16:00:57\\"
	TableWhere *string `json:"TableWhere,omitempty" xml:"TableWhere,omitempty"`
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParamTableIncludes) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParamTableIncludes) GoString() string {
	return s.String()
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParamTableIncludes) GetTableName() *string {
	return s.TableName
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParamTableIncludes) GetTableWhere() *string {
	return s.TableWhere
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParamTableIncludes) SetTableName(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParamTableIncludes {
	s.TableName = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParamTableIncludes) SetTableWhere(v string) *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParamTableIncludes {
	s.TableWhere = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponseBodyDataArchiveOrderDetailPluginParamTableIncludes) Validate() error {
	return dara.Validate(s)
}

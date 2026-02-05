// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseSupportPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListEnterpriseSupportPlanResponseBody
	GetCode() *string
	SetData(v []*ListEnterpriseSupportPlanResponseBodyData) *ListEnterpriseSupportPlanResponseBody
	GetData() []*ListEnterpriseSupportPlanResponseBodyData
	SetHttpStatusCode(v int32) *ListEnterpriseSupportPlanResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListEnterpriseSupportPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEnterpriseSupportPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEnterpriseSupportPlanResponseBody
	GetSuccess() *bool
}

type ListEnterpriseSupportPlanResponseBody struct {
	Code           *string                                      `json:"code,omitempty" xml:"code,omitempty"`
	Data           []*ListEnterpriseSupportPlanResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                       `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                                      `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListEnterpriseSupportPlanResponseBody) GetData() []*ListEnterpriseSupportPlanResponseBodyData {
	return s.Data
}

func (s *ListEnterpriseSupportPlanResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListEnterpriseSupportPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEnterpriseSupportPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnterpriseSupportPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEnterpriseSupportPlanResponseBody) SetCode(v string) *ListEnterpriseSupportPlanResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBody) SetData(v []*ListEnterpriseSupportPlanResponseBodyData) *ListEnterpriseSupportPlanResponseBody {
	s.Data = v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBody) SetHttpStatusCode(v int32) *ListEnterpriseSupportPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBody) SetMessage(v string) *ListEnterpriseSupportPlanResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBody) SetRequestId(v string) *ListEnterpriseSupportPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBody) SetSuccess(v bool) *ListEnterpriseSupportPlanResponseBody {
	s.Success = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBody) Validate() error {
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

type ListEnterpriseSupportPlanResponseBodyData struct {
	CanApplyFreeOrder        *bool                                                    `json:"canApplyFreeOrder,omitempty" xml:"canApplyFreeOrder,omitempty"`
	CustomerId               *string                                                  `json:"customerId,omitempty" xml:"customerId,omitempty"`
	Docs                     []*ListEnterpriseSupportPlanResponseBodyDataDocs         `json:"docs,omitempty" xml:"docs,omitempty" type:"Repeated"`
	EndTime                  *string                                                  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	FirstPayTime             *string                                                  `json:"firstPayTime,omitempty" xml:"firstPayTime,omitempty"`
	FreeOrderApplyCode       *string                                                  `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	FreeOrderApplyId         *int64                                                   `json:"freeOrderApplyId,omitempty" xml:"freeOrderApplyId,omitempty"`
	FreeOrderApplyTime       *string                                                  `json:"freeOrderApplyTime,omitempty" xml:"freeOrderApplyTime,omitempty"`
	FreeOrderApprovedTime    *string                                                  `json:"freeOrderApprovedTime,omitempty" xml:"freeOrderApprovedTime,omitempty"`
	FreeOrderExpectStartTime *string                                                  `json:"freeOrderExpectStartTime,omitempty" xml:"freeOrderExpectStartTime,omitempty"`
	GtspProjectId            *int64                                                   `json:"gtspProjectId,omitempty" xml:"gtspProjectId,omitempty"`
	InstanceId               *string                                                  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Nodes                    []*ListEnterpriseSupportPlanResponseBodyDataNodes        `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	OperateInfos             []*ListEnterpriseSupportPlanResponseBodyDataOperateInfos `json:"operateInfos,omitempty" xml:"operateInfos,omitempty" type:"Repeated"`
	OrderIds                 []*int64                                                 `json:"orderIds,omitempty" xml:"orderIds,omitempty" type:"Repeated"`
	ServiceName              *string                                                  `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	ServiceStatus            *string                                                  `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty"`
	ServiceStatusName        *string                                                  `json:"serviceStatusName,omitempty" xml:"serviceStatusName,omitempty"`
	ServiceType              *string                                                  `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	SortTime                 *string                                                  `json:"sortTime,omitempty" xml:"sortTime,omitempty"`
	StartTime                *string                                                  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Tags                     []*ListEnterpriseSupportPlanResponseBodyDataTags         `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	TaskNum                  *int64                                                   `json:"taskNum,omitempty" xml:"taskNum,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetCanApplyFreeOrder() *bool {
	return s.CanApplyFreeOrder
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetDocs() []*ListEnterpriseSupportPlanResponseBodyDataDocs {
	return s.Docs
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetFirstPayTime() *string {
	return s.FirstPayTime
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetFreeOrderApplyCode() *string {
	return s.FreeOrderApplyCode
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetFreeOrderApplyId() *int64 {
	return s.FreeOrderApplyId
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetFreeOrderApplyTime() *string {
	return s.FreeOrderApplyTime
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetFreeOrderApprovedTime() *string {
	return s.FreeOrderApprovedTime
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetFreeOrderExpectStartTime() *string {
	return s.FreeOrderExpectStartTime
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetGtspProjectId() *int64 {
	return s.GtspProjectId
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetNodes() []*ListEnterpriseSupportPlanResponseBodyDataNodes {
	return s.Nodes
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetOperateInfos() []*ListEnterpriseSupportPlanResponseBodyDataOperateInfos {
	return s.OperateInfos
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetOrderIds() []*int64 {
	return s.OrderIds
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetServiceStatusName() *string {
	return s.ServiceStatusName
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetSortTime() *string {
	return s.SortTime
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetTags() []*ListEnterpriseSupportPlanResponseBodyDataTags {
	return s.Tags
}

func (s *ListEnterpriseSupportPlanResponseBodyData) GetTaskNum() *int64 {
	return s.TaskNum
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetCanApplyFreeOrder(v bool) *ListEnterpriseSupportPlanResponseBodyData {
	s.CanApplyFreeOrder = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetCustomerId(v string) *ListEnterpriseSupportPlanResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetDocs(v []*ListEnterpriseSupportPlanResponseBodyDataDocs) *ListEnterpriseSupportPlanResponseBodyData {
	s.Docs = v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetEndTime(v string) *ListEnterpriseSupportPlanResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetFirstPayTime(v string) *ListEnterpriseSupportPlanResponseBodyData {
	s.FirstPayTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetFreeOrderApplyCode(v string) *ListEnterpriseSupportPlanResponseBodyData {
	s.FreeOrderApplyCode = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetFreeOrderApplyId(v int64) *ListEnterpriseSupportPlanResponseBodyData {
	s.FreeOrderApplyId = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetFreeOrderApplyTime(v string) *ListEnterpriseSupportPlanResponseBodyData {
	s.FreeOrderApplyTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetFreeOrderApprovedTime(v string) *ListEnterpriseSupportPlanResponseBodyData {
	s.FreeOrderApprovedTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetFreeOrderExpectStartTime(v string) *ListEnterpriseSupportPlanResponseBodyData {
	s.FreeOrderExpectStartTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetGtspProjectId(v int64) *ListEnterpriseSupportPlanResponseBodyData {
	s.GtspProjectId = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetInstanceId(v string) *ListEnterpriseSupportPlanResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetNodes(v []*ListEnterpriseSupportPlanResponseBodyDataNodes) *ListEnterpriseSupportPlanResponseBodyData {
	s.Nodes = v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetOperateInfos(v []*ListEnterpriseSupportPlanResponseBodyDataOperateInfos) *ListEnterpriseSupportPlanResponseBodyData {
	s.OperateInfos = v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetOrderIds(v []*int64) *ListEnterpriseSupportPlanResponseBodyData {
	s.OrderIds = v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetServiceName(v string) *ListEnterpriseSupportPlanResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetServiceStatus(v string) *ListEnterpriseSupportPlanResponseBodyData {
	s.ServiceStatus = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetServiceStatusName(v string) *ListEnterpriseSupportPlanResponseBodyData {
	s.ServiceStatusName = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetServiceType(v string) *ListEnterpriseSupportPlanResponseBodyData {
	s.ServiceType = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetSortTime(v string) *ListEnterpriseSupportPlanResponseBodyData {
	s.SortTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetStartTime(v string) *ListEnterpriseSupportPlanResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetTags(v []*ListEnterpriseSupportPlanResponseBodyDataTags) *ListEnterpriseSupportPlanResponseBodyData {
	s.Tags = v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetTaskNum(v int64) *ListEnterpriseSupportPlanResponseBodyData {
	s.TaskNum = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) Validate() error {
	if s.Docs != nil {
		for _, item := range s.Docs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OperateInfos != nil {
		for _, item := range s.OperateInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEnterpriseSupportPlanResponseBodyDataDocs struct {
	Attachments        []*ListEnterpriseSupportPlanResponseBodyDataDocsAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	DocId              *int64                                                      `json:"docId,omitempty" xml:"docId,omitempty"`
	Evaluated          *bool                                                       `json:"evaluated,omitempty" xml:"evaluated,omitempty"`
	EvaluationUrl      *string                                                     `json:"evaluationUrl,omitempty" xml:"evaluationUrl,omitempty"`
	FileName           *string                                                     `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FreeOrderApplyCode *string                                                     `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	Name               *string                                                     `json:"name,omitempty" xml:"name,omitempty"`
	OrderId            *string                                                     `json:"orderId,omitempty" xml:"orderId,omitempty"`
	UploadTime         *string                                                     `json:"uploadTime,omitempty" xml:"uploadTime,omitempty"`
	Url                *string                                                     `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataDocs) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataDocs) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) GetAttachments() []*ListEnterpriseSupportPlanResponseBodyDataDocsAttachments {
	return s.Attachments
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) GetDocId() *int64 {
	return s.DocId
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) GetEvaluated() *bool {
	return s.Evaluated
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) GetEvaluationUrl() *string {
	return s.EvaluationUrl
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) GetFileName() *string {
	return s.FileName
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) GetFreeOrderApplyCode() *string {
	return s.FreeOrderApplyCode
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) GetName() *string {
	return s.Name
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) GetOrderId() *string {
	return s.OrderId
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) GetUploadTime() *string {
	return s.UploadTime
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) GetUrl() *string {
	return s.Url
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) SetAttachments(v []*ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) *ListEnterpriseSupportPlanResponseBodyDataDocs {
	s.Attachments = v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) SetDocId(v int64) *ListEnterpriseSupportPlanResponseBodyDataDocs {
	s.DocId = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) SetEvaluated(v bool) *ListEnterpriseSupportPlanResponseBodyDataDocs {
	s.Evaluated = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) SetEvaluationUrl(v string) *ListEnterpriseSupportPlanResponseBodyDataDocs {
	s.EvaluationUrl = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) SetFileName(v string) *ListEnterpriseSupportPlanResponseBodyDataDocs {
	s.FileName = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) SetFreeOrderApplyCode(v string) *ListEnterpriseSupportPlanResponseBodyDataDocs {
	s.FreeOrderApplyCode = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) SetName(v string) *ListEnterpriseSupportPlanResponseBodyDataDocs {
	s.Name = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) SetOrderId(v string) *ListEnterpriseSupportPlanResponseBodyDataDocs {
	s.OrderId = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) SetUploadTime(v string) *ListEnterpriseSupportPlanResponseBodyDataDocs {
	s.UploadTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) SetUrl(v string) *ListEnterpriseSupportPlanResponseBodyDataDocs {
	s.Url = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEnterpriseSupportPlanResponseBodyDataDocsAttachments struct {
	DocId              *int64  `json:"docId,omitempty" xml:"docId,omitempty"`
	Evaluated          *bool   `json:"evaluated,omitempty" xml:"evaluated,omitempty"`
	EvaluationUrl      *string `json:"evaluationUrl,omitempty" xml:"evaluationUrl,omitempty"`
	FileName           *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FreeOrderApplyCode *string `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	OrderId            *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	UploadTime         *string `json:"uploadTime,omitempty" xml:"uploadTime,omitempty"`
	Url                *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) GetDocId() *int64 {
	return s.DocId
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) GetEvaluated() *bool {
	return s.Evaluated
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) GetEvaluationUrl() *string {
	return s.EvaluationUrl
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) GetFileName() *string {
	return s.FileName
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) GetFreeOrderApplyCode() *string {
	return s.FreeOrderApplyCode
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) GetName() *string {
	return s.Name
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) GetOrderId() *string {
	return s.OrderId
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) GetUploadTime() *string {
	return s.UploadTime
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) GetUrl() *string {
	return s.Url
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) SetDocId(v int64) *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments {
	s.DocId = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) SetEvaluated(v bool) *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments {
	s.Evaluated = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) SetEvaluationUrl(v string) *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments {
	s.EvaluationUrl = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) SetFileName(v string) *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments {
	s.FileName = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) SetFreeOrderApplyCode(v string) *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments {
	s.FreeOrderApplyCode = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) SetName(v string) *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments {
	s.Name = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) SetOrderId(v string) *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments {
	s.OrderId = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) SetUploadTime(v string) *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments {
	s.UploadTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) SetUrl(v string) *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments {
	s.Url = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocsAttachments) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseSupportPlanResponseBodyDataNodes struct {
	DocNode               *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode               `json:"docNode,omitempty" xml:"docNode,omitempty" type:"Struct"`
	FinishNode            *ListEnterpriseSupportPlanResponseBodyDataNodesFinishNode            `json:"finishNode,omitempty" xml:"finishNode,omitempty" type:"Struct"`
	FreeOrderAuditNode    *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode    `json:"freeOrderAuditNode,omitempty" xml:"freeOrderAuditNode,omitempty" type:"Struct"`
	FreeOrderNode         *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode         `json:"freeOrderNode,omitempty" xml:"freeOrderNode,omitempty" type:"Struct"`
	Name                  *string                                                              `json:"name,omitempty" xml:"name,omitempty"`
	OrderDate             *int64                                                               `json:"orderDate,omitempty" xml:"orderDate,omitempty"`
	OrderNode             *ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode             `json:"orderNode,omitempty" xml:"orderNode,omitempty" type:"Struct"`
	ServiceImplementation *ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation `json:"serviceImplementation,omitempty" xml:"serviceImplementation,omitempty" type:"Struct"`
	Status                *int32                                                               `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) GetDocNode() *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode {
	return s.DocNode
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) GetFinishNode() *ListEnterpriseSupportPlanResponseBodyDataNodesFinishNode {
	return s.FinishNode
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) GetFreeOrderAuditNode() *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode {
	return s.FreeOrderAuditNode
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) GetFreeOrderNode() *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode {
	return s.FreeOrderNode
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) GetName() *string {
	return s.Name
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) GetOrderDate() *int64 {
	return s.OrderDate
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) GetOrderNode() *ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode {
	return s.OrderNode
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) GetServiceImplementation() *ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation {
	return s.ServiceImplementation
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) GetStatus() *int32 {
	return s.Status
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) SetDocNode(v *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) *ListEnterpriseSupportPlanResponseBodyDataNodes {
	s.DocNode = v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) SetFinishNode(v *ListEnterpriseSupportPlanResponseBodyDataNodesFinishNode) *ListEnterpriseSupportPlanResponseBodyDataNodes {
	s.FinishNode = v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) SetFreeOrderAuditNode(v *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode) *ListEnterpriseSupportPlanResponseBodyDataNodes {
	s.FreeOrderAuditNode = v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) SetFreeOrderNode(v *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode) *ListEnterpriseSupportPlanResponseBodyDataNodes {
	s.FreeOrderNode = v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) SetName(v string) *ListEnterpriseSupportPlanResponseBodyDataNodes {
	s.Name = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) SetOrderDate(v int64) *ListEnterpriseSupportPlanResponseBodyDataNodes {
	s.OrderDate = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) SetOrderNode(v *ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode) *ListEnterpriseSupportPlanResponseBodyDataNodes {
	s.OrderNode = v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) SetServiceImplementation(v *ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation) *ListEnterpriseSupportPlanResponseBodyDataNodes {
	s.ServiceImplementation = v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) SetStatus(v int32) *ListEnterpriseSupportPlanResponseBodyDataNodes {
	s.Status = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodes) Validate() error {
	if s.DocNode != nil {
		if err := s.DocNode.Validate(); err != nil {
			return err
		}
	}
	if s.FinishNode != nil {
		if err := s.FinishNode.Validate(); err != nil {
			return err
		}
	}
	if s.FreeOrderAuditNode != nil {
		if err := s.FreeOrderAuditNode.Validate(); err != nil {
			return err
		}
	}
	if s.FreeOrderNode != nil {
		if err := s.FreeOrderNode.Validate(); err != nil {
			return err
		}
	}
	if s.OrderNode != nil {
		if err := s.OrderNode.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceImplementation != nil {
		if err := s.ServiceImplementation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEnterpriseSupportPlanResponseBodyDataNodesDocNode struct {
	Attachments        []interface{} `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	DocId              *int64        `json:"docId,omitempty" xml:"docId,omitempty"`
	DocName            *string       `json:"docName,omitempty" xml:"docName,omitempty"`
	DocSubmitTime      *string       `json:"docSubmitTime,omitempty" xml:"docSubmitTime,omitempty"`
	FileName           *string       `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FreeOrderApplyCode *string       `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	OrderId            *string       `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) GetAttachments() []interface{} {
	return s.Attachments
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) GetDocId() *int64 {
	return s.DocId
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) GetDocName() *string {
	return s.DocName
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) GetDocSubmitTime() *string {
	return s.DocSubmitTime
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) GetFileName() *string {
	return s.FileName
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) GetFreeOrderApplyCode() *string {
	return s.FreeOrderApplyCode
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) GetOrderId() *string {
	return s.OrderId
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) SetAttachments(v []interface{}) *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode {
	s.Attachments = v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) SetDocId(v int64) *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode {
	s.DocId = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) SetDocName(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode {
	s.DocName = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) SetDocSubmitTime(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode {
	s.DocSubmitTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) SetFileName(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode {
	s.FileName = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) SetFreeOrderApplyCode(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode {
	s.FreeOrderApplyCode = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) SetOrderId(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode {
	s.OrderId = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseSupportPlanResponseBodyDataNodesFinishNode struct {
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesFinishNode) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesFinishNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFinishNode) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFinishNode) SetFinishTime(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesFinishNode {
	s.FinishTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFinishNode) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode struct {
	AuditTime  *string `json:"auditTime,omitempty" xml:"auditTime,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	StatusName *string `json:"statusName,omitempty" xml:"statusName,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode) GetAuditTime() *string {
	return s.AuditTime
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode) GetStatus() *string {
	return s.Status
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode) GetStatusName() *string {
	return s.StatusName
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode) SetAuditTime(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode {
	s.AuditTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode) SetStatus(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode {
	s.Status = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode) SetStatusName(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode {
	s.StatusName = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode struct {
	ApplyTime *string `json:"applyTime,omitempty" xml:"applyTime,omitempty"`
	Uid       *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode) GetApplyTime() *string {
	return s.ApplyTime
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode) GetUid() *string {
	return s.Uid
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode) SetApplyTime(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode {
	s.ApplyTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode) SetUid(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode {
	s.Uid = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode struct {
	PayTime *string `json:"payTime,omitempty" xml:"payTime,omitempty"`
	Uid     *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode) GetPayTime() *string {
	return s.PayTime
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode) GetUid() *string {
	return s.Uid
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode) SetPayTime(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode {
	s.PayTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode) SetUid(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode {
	s.Uid = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation struct {
	EndTime   *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation) GetEndTime() *string {
	return s.EndTime
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation) GetStartTime() *string {
	return s.StartTime
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation) SetEndTime(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation {
	s.EndTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation) SetStartTime(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation {
	s.StartTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseSupportPlanResponseBodyDataOperateInfos struct {
	CanClick *bool   `json:"canClick,omitempty" xml:"canClick,omitempty"`
	Text     *string `json:"text,omitempty" xml:"text,omitempty"`
	Url      *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataOperateInfos) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataOperateInfos) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataOperateInfos) GetCanClick() *bool {
	return s.CanClick
}

func (s *ListEnterpriseSupportPlanResponseBodyDataOperateInfos) GetText() *string {
	return s.Text
}

func (s *ListEnterpriseSupportPlanResponseBodyDataOperateInfos) GetUrl() *string {
	return s.Url
}

func (s *ListEnterpriseSupportPlanResponseBodyDataOperateInfos) SetCanClick(v bool) *ListEnterpriseSupportPlanResponseBodyDataOperateInfos {
	s.CanClick = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataOperateInfos) SetText(v string) *ListEnterpriseSupportPlanResponseBodyDataOperateInfos {
	s.Text = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataOperateInfos) SetUrl(v string) *ListEnterpriseSupportPlanResponseBodyDataOperateInfos {
	s.Url = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataOperateInfos) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseSupportPlanResponseBodyDataTags struct {
	ExtendInfos []*string `json:"extendInfos,omitempty" xml:"extendInfos,omitempty" type:"Repeated"`
	Show        *bool     `json:"show,omitempty" xml:"show,omitempty"`
	TagCode     *string   `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	TagName     *string   `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataTags) GetExtendInfos() []*string {
	return s.ExtendInfos
}

func (s *ListEnterpriseSupportPlanResponseBodyDataTags) GetShow() *bool {
	return s.Show
}

func (s *ListEnterpriseSupportPlanResponseBodyDataTags) GetTagCode() *string {
	return s.TagCode
}

func (s *ListEnterpriseSupportPlanResponseBodyDataTags) GetTagName() *string {
	return s.TagName
}

func (s *ListEnterpriseSupportPlanResponseBodyDataTags) SetExtendInfos(v []*string) *ListEnterpriseSupportPlanResponseBodyDataTags {
	s.ExtendInfos = v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataTags) SetShow(v bool) *ListEnterpriseSupportPlanResponseBodyDataTags {
	s.Show = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataTags) SetTagCode(v string) *ListEnterpriseSupportPlanResponseBodyDataTags {
	s.TagCode = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataTags) SetTagName(v string) *ListEnterpriseSupportPlanResponseBodyDataTags {
	s.TagName = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseSupportPlanSimpleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListEnterpriseSupportPlanSimpleResponseBody
	GetCode() *string
	SetData(v []*ListEnterpriseSupportPlanSimpleResponseBodyData) *ListEnterpriseSupportPlanSimpleResponseBody
	GetData() []*ListEnterpriseSupportPlanSimpleResponseBodyData
	SetHttpStatusCode(v int32) *ListEnterpriseSupportPlanSimpleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListEnterpriseSupportPlanSimpleResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEnterpriseSupportPlanSimpleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEnterpriseSupportPlanSimpleResponseBody
	GetSuccess() *bool
}

type ListEnterpriseSupportPlanSimpleResponseBody struct {
	Code           *string                                            `json:"code,omitempty" xml:"code,omitempty"`
	Data           []*ListEnterpriseSupportPlanSimpleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                             `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                                            `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListEnterpriseSupportPlanSimpleResponseBody) GetData() []*ListEnterpriseSupportPlanSimpleResponseBodyData {
	return s.Data
}

func (s *ListEnterpriseSupportPlanSimpleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListEnterpriseSupportPlanSimpleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEnterpriseSupportPlanSimpleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnterpriseSupportPlanSimpleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEnterpriseSupportPlanSimpleResponseBody) SetCode(v string) *ListEnterpriseSupportPlanSimpleResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBody) SetData(v []*ListEnterpriseSupportPlanSimpleResponseBodyData) *ListEnterpriseSupportPlanSimpleResponseBody {
	s.Data = v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBody) SetHttpStatusCode(v int32) *ListEnterpriseSupportPlanSimpleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBody) SetMessage(v string) *ListEnterpriseSupportPlanSimpleResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBody) SetRequestId(v string) *ListEnterpriseSupportPlanSimpleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBody) SetSuccess(v bool) *ListEnterpriseSupportPlanSimpleResponseBody {
	s.Success = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBody) Validate() error {
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

type ListEnterpriseSupportPlanSimpleResponseBodyData struct {
	CanApplyFreeOrder        *bool                                                          `json:"canApplyFreeOrder,omitempty" xml:"canApplyFreeOrder,omitempty"`
	CustomerId               *string                                                        `json:"customerId,omitempty" xml:"customerId,omitempty"`
	Docs                     []*ListEnterpriseSupportPlanSimpleResponseBodyDataDocs         `json:"docs,omitempty" xml:"docs,omitempty" type:"Repeated"`
	EndTime                  *string                                                        `json:"endTime,omitempty" xml:"endTime,omitempty"`
	FirstPayTime             *string                                                        `json:"firstPayTime,omitempty" xml:"firstPayTime,omitempty"`
	FreeOrderApplyCode       *string                                                        `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	FreeOrderApplyId         *int64                                                         `json:"freeOrderApplyId,omitempty" xml:"freeOrderApplyId,omitempty"`
	FreeOrderApplyTime       *string                                                        `json:"freeOrderApplyTime,omitempty" xml:"freeOrderApplyTime,omitempty"`
	FreeOrderApprovedTime    *string                                                        `json:"freeOrderApprovedTime,omitempty" xml:"freeOrderApprovedTime,omitempty"`
	FreeOrderExpectStartTime *string                                                        `json:"freeOrderExpectStartTime,omitempty" xml:"freeOrderExpectStartTime,omitempty"`
	InstanceId               *string                                                        `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Nodes                    []*ListEnterpriseSupportPlanSimpleResponseBodyDataNodes        `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	OperateInfos             []*ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos `json:"operateInfos,omitempty" xml:"operateInfos,omitempty" type:"Repeated"`
	OrderIds                 []*int64                                                       `json:"orderIds,omitempty" xml:"orderIds,omitempty" type:"Repeated"`
	ServiceName              *string                                                        `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	ServiceStatus            *string                                                        `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty"`
	ServiceStatusName        *string                                                        `json:"serviceStatusName,omitempty" xml:"serviceStatusName,omitempty"`
	ServiceType              *string                                                        `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	SortTime                 *string                                                        `json:"sortTime,omitempty" xml:"sortTime,omitempty"`
	StartTime                *string                                                        `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TaskNum                  *int64                                                         `json:"taskNum,omitempty" xml:"taskNum,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetCanApplyFreeOrder() *bool {
	return s.CanApplyFreeOrder
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetDocs() []*ListEnterpriseSupportPlanSimpleResponseBodyDataDocs {
	return s.Docs
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetFirstPayTime() *string {
	return s.FirstPayTime
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetFreeOrderApplyCode() *string {
	return s.FreeOrderApplyCode
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetFreeOrderApplyId() *int64 {
	return s.FreeOrderApplyId
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetFreeOrderApplyTime() *string {
	return s.FreeOrderApplyTime
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetFreeOrderApprovedTime() *string {
	return s.FreeOrderApprovedTime
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetFreeOrderExpectStartTime() *string {
	return s.FreeOrderExpectStartTime
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetNodes() []*ListEnterpriseSupportPlanSimpleResponseBodyDataNodes {
	return s.Nodes
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetOperateInfos() []*ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos {
	return s.OperateInfos
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetOrderIds() []*int64 {
	return s.OrderIds
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetServiceStatusName() *string {
	return s.ServiceStatusName
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetSortTime() *string {
	return s.SortTime
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) GetTaskNum() *int64 {
	return s.TaskNum
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetCanApplyFreeOrder(v bool) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.CanApplyFreeOrder = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetCustomerId(v string) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetDocs(v []*ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.Docs = v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetEndTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetFirstPayTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.FirstPayTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetFreeOrderApplyCode(v string) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.FreeOrderApplyCode = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetFreeOrderApplyId(v int64) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.FreeOrderApplyId = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetFreeOrderApplyTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.FreeOrderApplyTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetFreeOrderApprovedTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.FreeOrderApprovedTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetFreeOrderExpectStartTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.FreeOrderExpectStartTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetInstanceId(v string) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetNodes(v []*ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.Nodes = v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetOperateInfos(v []*ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.OperateInfos = v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetOrderIds(v []*int64) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.OrderIds = v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetServiceName(v string) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetServiceStatus(v string) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.ServiceStatus = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetServiceStatusName(v string) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.ServiceStatusName = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetServiceType(v string) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.ServiceType = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetSortTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.SortTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetStartTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetTaskNum(v int64) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.TaskNum = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) Validate() error {
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
	return nil
}

type ListEnterpriseSupportPlanSimpleResponseBodyDataDocs struct {
	DocId              *int64  `json:"docId,omitempty" xml:"docId,omitempty"`
	FileName           *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FreeOrderApplyCode *string `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	OrderId            *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	UploadTime         *string `json:"uploadTime,omitempty" xml:"uploadTime,omitempty"`
	Url                *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) GetDocId() *int64 {
	return s.DocId
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) GetFileName() *string {
	return s.FileName
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) GetFreeOrderApplyCode() *string {
	return s.FreeOrderApplyCode
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) GetName() *string {
	return s.Name
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) GetOrderId() *string {
	return s.OrderId
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) GetUploadTime() *string {
	return s.UploadTime
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) GetUrl() *string {
	return s.Url
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) SetDocId(v int64) *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs {
	s.DocId = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) SetFileName(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs {
	s.FileName = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) SetFreeOrderApplyCode(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs {
	s.FreeOrderApplyCode = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) SetName(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs {
	s.Name = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) SetOrderId(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs {
	s.OrderId = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) SetUploadTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs {
	s.UploadTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) SetUrl(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs {
	s.Url = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseSupportPlanSimpleResponseBodyDataNodes struct {
	DocNode               *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode               `json:"docNode,omitempty" xml:"docNode,omitempty" type:"Struct"`
	FinishNode            *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFinishNode            `json:"finishNode,omitempty" xml:"finishNode,omitempty" type:"Struct"`
	FreeOrderAuditNode    *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode    `json:"freeOrderAuditNode,omitempty" xml:"freeOrderAuditNode,omitempty" type:"Struct"`
	FreeOrderNode         *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode         `json:"freeOrderNode,omitempty" xml:"freeOrderNode,omitempty" type:"Struct"`
	Name                  *string                                                                    `json:"name,omitempty" xml:"name,omitempty"`
	OrderDate             *int64                                                                     `json:"orderDate,omitempty" xml:"orderDate,omitempty"`
	OrderNode             *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode             `json:"orderNode,omitempty" xml:"orderNode,omitempty" type:"Struct"`
	ServiceImplementation *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation `json:"serviceImplementation,omitempty" xml:"serviceImplementation,omitempty" type:"Struct"`
	Status                *int32                                                                     `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) GetDocNode() *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode {
	return s.DocNode
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) GetFinishNode() *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFinishNode {
	return s.FinishNode
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) GetFreeOrderAuditNode() *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode {
	return s.FreeOrderAuditNode
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) GetFreeOrderNode() *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode {
	return s.FreeOrderNode
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) GetName() *string {
	return s.Name
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) GetOrderDate() *int64 {
	return s.OrderDate
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) GetOrderNode() *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode {
	return s.OrderNode
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) GetServiceImplementation() *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation {
	return s.ServiceImplementation
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) GetStatus() *int32 {
	return s.Status
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) SetDocNode(v *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes {
	s.DocNode = v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) SetFinishNode(v *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFinishNode) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes {
	s.FinishNode = v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) SetFreeOrderAuditNode(v *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes {
	s.FreeOrderAuditNode = v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) SetFreeOrderNode(v *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes {
	s.FreeOrderNode = v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) SetName(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes {
	s.Name = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) SetOrderDate(v int64) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes {
	s.OrderDate = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) SetOrderNode(v *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes {
	s.OrderNode = v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) SetServiceImplementation(v *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes {
	s.ServiceImplementation = v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) SetStatus(v int32) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes {
	s.Status = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) Validate() error {
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

type ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode struct {
	DocId              *int64  `json:"docId,omitempty" xml:"docId,omitempty"`
	DocName            *string `json:"docName,omitempty" xml:"docName,omitempty"`
	DocSubmitTime      *string `json:"docSubmitTime,omitempty" xml:"docSubmitTime,omitempty"`
	FileName           *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FreeOrderApplyCode *string `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	OrderId            *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) GetDocId() *int64 {
	return s.DocId
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) GetDocName() *string {
	return s.DocName
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) GetDocSubmitTime() *string {
	return s.DocSubmitTime
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) GetFileName() *string {
	return s.FileName
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) GetFreeOrderApplyCode() *string {
	return s.FreeOrderApplyCode
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) GetOrderId() *string {
	return s.OrderId
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) SetDocId(v int64) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode {
	s.DocId = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) SetDocName(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode {
	s.DocName = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) SetDocSubmitTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode {
	s.DocSubmitTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) SetFileName(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode {
	s.FileName = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) SetFreeOrderApplyCode(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode {
	s.FreeOrderApplyCode = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) SetOrderId(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode {
	s.OrderId = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFinishNode struct {
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFinishNode) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFinishNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFinishNode) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFinishNode) SetFinishTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFinishNode {
	s.FinishTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFinishNode) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode struct {
	AuditTime  *string `json:"auditTime,omitempty" xml:"auditTime,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	StatusName *string `json:"statusName,omitempty" xml:"statusName,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode) GetAuditTime() *string {
	return s.AuditTime
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode) GetStatus() *string {
	return s.Status
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode) GetStatusName() *string {
	return s.StatusName
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode) SetAuditTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode {
	s.AuditTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode) SetStatus(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode {
	s.Status = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode) SetStatusName(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode {
	s.StatusName = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode struct {
	ApplyTime *string `json:"applyTime,omitempty" xml:"applyTime,omitempty"`
	Uid       *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode) GetApplyTime() *string {
	return s.ApplyTime
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode) GetUid() *string {
	return s.Uid
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode) SetApplyTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode {
	s.ApplyTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode) SetUid(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode {
	s.Uid = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode struct {
	PayTime *string `json:"payTime,omitempty" xml:"payTime,omitempty"`
	Uid     *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode) GetPayTime() *string {
	return s.PayTime
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode) GetUid() *string {
	return s.Uid
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode) SetPayTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode {
	s.PayTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode) SetUid(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode {
	s.Uid = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation struct {
	EndTime   *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation) GetEndTime() *string {
	return s.EndTime
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation) GetStartTime() *string {
	return s.StartTime
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation) SetEndTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation {
	s.EndTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation) SetStartTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation {
	s.StartTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation) Validate() error {
	return dara.Validate(s)
}

type ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos struct {
	CanClick *bool   `json:"canClick,omitempty" xml:"canClick,omitempty"`
	Text     *string `json:"text,omitempty" xml:"text,omitempty"`
	Url      *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos) GetCanClick() *bool {
	return s.CanClick
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos) GetText() *string {
	return s.Text
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos) GetUrl() *string {
	return s.Url
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos) SetCanClick(v bool) *ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos {
	s.CanClick = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos) SetText(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos {
	s.Text = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos) SetUrl(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos {
	s.Url = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataOperateInfos) Validate() error {
	return dara.Validate(s)
}

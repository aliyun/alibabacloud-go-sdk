// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DataValue struct {
	DocId      *int64  `json:"docId,omitempty" xml:"docId,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	FileName   *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	Url        *string `json:"url,omitempty" xml:"url,omitempty"`
	UploadTime *string `json:"uploadTime,omitempty" xml:"uploadTime,omitempty"`
	OrderId    *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	ApplyId    *string `json:"applyId,omitempty" xml:"applyId,omitempty"`
}

func (s DataValue) String() string {
	return tea.Prettify(s)
}

func (s DataValue) GoString() string {
	return s.String()
}

func (s *DataValue) SetDocId(v int64) *DataValue {
	s.DocId = &v
	return s
}

func (s *DataValue) SetName(v string) *DataValue {
	s.Name = &v
	return s
}

func (s *DataValue) SetFileName(v string) *DataValue {
	s.FileName = &v
	return s
}

func (s *DataValue) SetUrl(v string) *DataValue {
	s.Url = &v
	return s
}

func (s *DataValue) SetUploadTime(v string) *DataValue {
	s.UploadTime = &v
	return s
}

func (s *DataValue) SetOrderId(v string) *DataValue {
	s.OrderId = &v
	return s
}

func (s *DataValue) SetApplyId(v string) *DataValue {
	s.ApplyId = &v
	return s
}

type GetDownloadUrlRequest struct {
	FileId             *int64  `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileKey            *string `json:"fileKey,omitempty" xml:"fileKey,omitempty"`
	FreeOrderApplyCode *string `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	OrderId            *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	Scene              *string `json:"scene,omitempty" xml:"scene,omitempty"`
}

func (s GetDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlRequest) SetFileId(v int64) *GetDownloadUrlRequest {
	s.FileId = &v
	return s
}

func (s *GetDownloadUrlRequest) SetFileKey(v string) *GetDownloadUrlRequest {
	s.FileKey = &v
	return s
}

func (s *GetDownloadUrlRequest) SetFreeOrderApplyCode(v string) *GetDownloadUrlRequest {
	s.FreeOrderApplyCode = &v
	return s
}

func (s *GetDownloadUrlRequest) SetOrderId(v string) *GetDownloadUrlRequest {
	s.OrderId = &v
	return s
}

func (s *GetDownloadUrlRequest) SetScene(v string) *GetDownloadUrlRequest {
	s.Scene = &v
	return s
}

type GetDownloadUrlResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *string `json:"data,omitempty" xml:"data,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlResponseBody) SetCode(v string) *GetDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetData(v string) *GetDownloadUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetHttpStatusCode(v int32) *GetDownloadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetMessage(v string) *GetDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetRequestId(v string) *GetDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetSuccess(v bool) *GetDownloadUrlResponseBody {
	s.Success = &v
	return s
}

type GetDownloadUrlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlResponse) SetHeaders(v map[string]*string) *GetDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetDownloadUrlResponse) SetStatusCode(v int32) *GetDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDownloadUrlResponse) SetBody(v *GetDownloadUrlResponseBody) *GetDownloadUrlResponse {
	s.Body = v
	return s
}

type GetEnterpriseSupportPlanDetailRequest struct {
	FreeOrderApplyCodes []*string `json:"freeOrderApplyCodes,omitempty" xml:"freeOrderApplyCodes,omitempty" type:"Repeated"`
	OrderIds            []*int64  `json:"orderIds,omitempty" xml:"orderIds,omitempty" type:"Repeated"`
}

func (s GetEnterpriseSupportPlanDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseSupportPlanDetailRequest) GoString() string {
	return s.String()
}

func (s *GetEnterpriseSupportPlanDetailRequest) SetFreeOrderApplyCodes(v []*string) *GetEnterpriseSupportPlanDetailRequest {
	s.FreeOrderApplyCodes = v
	return s
}

func (s *GetEnterpriseSupportPlanDetailRequest) SetOrderIds(v []*int64) *GetEnterpriseSupportPlanDetailRequest {
	s.OrderIds = v
	return s
}

type GetEnterpriseSupportPlanDetailResponseBody struct {
	Code           *string                                         `json:"code,omitempty" xml:"code,omitempty"`
	Data           *GetEnterpriseSupportPlanDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                          `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                                         `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetEnterpriseSupportPlanDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseSupportPlanDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnterpriseSupportPlanDetailResponseBody) SetCode(v string) *GetEnterpriseSupportPlanDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBody) SetData(v *GetEnterpriseSupportPlanDetailResponseBodyData) *GetEnterpriseSupportPlanDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBody) SetHttpStatusCode(v int32) *GetEnterpriseSupportPlanDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBody) SetMessage(v string) *GetEnterpriseSupportPlanDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBody) SetRequestId(v string) *GetEnterpriseSupportPlanDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBody) SetSuccess(v bool) *GetEnterpriseSupportPlanDetailResponseBody {
	s.Success = &v
	return s
}

type GetEnterpriseSupportPlanDetailResponseBodyData struct {
	CanApplyFreeOrder        *bool                                                         `json:"canApplyFreeOrder,omitempty" xml:"canApplyFreeOrder,omitempty"`
	CustomerId               *int64                                                        `json:"customerId,omitempty" xml:"customerId,omitempty"`
	DingGroups               []*GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups   `json:"dingGroups,omitempty" xml:"dingGroups,omitempty" type:"Repeated"`
	Docs                     []*GetEnterpriseSupportPlanDetailResponseBodyDataDocs         `json:"docs,omitempty" xml:"docs,omitempty" type:"Repeated"`
	EndTime                  *string                                                       `json:"endTime,omitempty" xml:"endTime,omitempty"`
	FirstPayTime             *string                                                       `json:"firstPayTime,omitempty" xml:"firstPayTime,omitempty"`
	FreeOrderApplyCode       *string                                                       `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	FreeOrderApplyId         *int64                                                        `json:"freeOrderApplyId,omitempty" xml:"freeOrderApplyId,omitempty"`
	FreeOrderApplyTime       *string                                                       `json:"freeOrderApplyTime,omitempty" xml:"freeOrderApplyTime,omitempty"`
	FreeOrderApprovedTime    *string                                                       `json:"freeOrderApprovedTime,omitempty" xml:"freeOrderApprovedTime,omitempty"`
	FreeOrderExpectStartTime *string                                                       `json:"freeOrderExpectStartTime,omitempty" xml:"freeOrderExpectStartTime,omitempty"`
	InstanceId               *string                                                       `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Nodes                    []*GetEnterpriseSupportPlanDetailResponseBodyDataNodes        `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	OrderIds                 []*int64                                                      `json:"orderIds,omitempty" xml:"orderIds,omitempty" type:"Repeated"`
	ServiceItems             []*GetEnterpriseSupportPlanDetailResponseBodyDataServiceItems `json:"serviceItems,omitempty" xml:"serviceItems,omitempty" type:"Repeated"`
	ServiceName              *string                                                       `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	ServiceStatus            *string                                                       `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty"`
	ServiceStatusName        *string                                                       `json:"serviceStatusName,omitempty" xml:"serviceStatusName,omitempty"`
	ServiceType              *string                                                       `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	SortTime                 *string                                                       `json:"sortTime,omitempty" xml:"sortTime,omitempty"`
	StartTime                *string                                                       `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TaskNum                  *int64                                                        `json:"taskNum,omitempty" xml:"taskNum,omitempty"`
}

func (s GetEnterpriseSupportPlanDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseSupportPlanDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetCanApplyFreeOrder(v bool) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.CanApplyFreeOrder = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetCustomerId(v int64) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetDingGroups(v []*GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.DingGroups = v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetDocs(v []*GetEnterpriseSupportPlanDetailResponseBodyDataDocs) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.Docs = v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetEndTime(v string) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetFirstPayTime(v string) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.FirstPayTime = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetFreeOrderApplyCode(v string) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.FreeOrderApplyCode = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetFreeOrderApplyId(v int64) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.FreeOrderApplyId = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetFreeOrderApplyTime(v string) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.FreeOrderApplyTime = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetFreeOrderApprovedTime(v string) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.FreeOrderApprovedTime = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetFreeOrderExpectStartTime(v string) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.FreeOrderExpectStartTime = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetInstanceId(v string) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetNodes(v []*GetEnterpriseSupportPlanDetailResponseBodyDataNodes) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.Nodes = v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetOrderIds(v []*int64) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.OrderIds = v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetServiceItems(v []*GetEnterpriseSupportPlanDetailResponseBodyDataServiceItems) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.ServiceItems = v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetServiceName(v string) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetServiceStatus(v string) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.ServiceStatus = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetServiceStatusName(v string) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.ServiceStatusName = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetServiceType(v string) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.ServiceType = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetSortTime(v string) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.SortTime = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetStartTime(v string) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyData) SetTaskNum(v int64) *GetEnterpriseSupportPlanDetailResponseBodyData {
	s.TaskNum = &v
	return s
}

type GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups struct {
	MainDingDepartmentId *string `json:"mainDingDepartmentId,omitempty" xml:"mainDingDepartmentId,omitempty"`
	MainDingGroupId      *string `json:"mainDingGroupId,omitempty" xml:"mainDingGroupId,omitempty"`
	MainDingGroupName    *string `json:"mainDingGroupName,omitempty" xml:"mainDingGroupName,omitempty"`
	SubDingDepartmentId  *string `json:"subDingDepartmentId,omitempty" xml:"subDingDepartmentId,omitempty"`
	SubDingGroupId       *string `json:"subDingGroupId,omitempty" xml:"subDingGroupId,omitempty"`
	SubDingGroupName     *string `json:"subDingGroupName,omitempty" xml:"subDingGroupName,omitempty"`
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups) GoString() string {
	return s.String()
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups) SetMainDingDepartmentId(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups {
	s.MainDingDepartmentId = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups) SetMainDingGroupId(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups {
	s.MainDingGroupId = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups) SetMainDingGroupName(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups {
	s.MainDingGroupName = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups) SetSubDingDepartmentId(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups {
	s.SubDingDepartmentId = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups) SetSubDingGroupId(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups {
	s.SubDingGroupId = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups) SetSubDingGroupName(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataDingGroups {
	s.SubDingGroupName = &v
	return s
}

type GetEnterpriseSupportPlanDetailResponseBodyDataDocs struct {
	DocId              *int64  `json:"docId,omitempty" xml:"docId,omitempty"`
	FileName           *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FreeOrderApplyCode *string `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	OrderId            *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	UploadTime         *string `json:"uploadTime,omitempty" xml:"uploadTime,omitempty"`
	Url                *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataDocs) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataDocs) GoString() string {
	return s.String()
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataDocs) SetDocId(v int64) *GetEnterpriseSupportPlanDetailResponseBodyDataDocs {
	s.DocId = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataDocs) SetFileName(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataDocs {
	s.FileName = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataDocs) SetFreeOrderApplyCode(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataDocs {
	s.FreeOrderApplyCode = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataDocs) SetName(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataDocs {
	s.Name = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataDocs) SetOrderId(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataDocs {
	s.OrderId = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataDocs) SetUploadTime(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataDocs {
	s.UploadTime = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataDocs) SetUrl(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataDocs {
	s.Url = &v
	return s
}

type GetEnterpriseSupportPlanDetailResponseBodyDataNodes struct {
	DocNode               *GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode               `json:"docNode,omitempty" xml:"docNode,omitempty" type:"Struct"`
	FinishNode            *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFinishNode            `json:"finishNode,omitempty" xml:"finishNode,omitempty" type:"Struct"`
	FreeOrderAuditNode    *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderAuditNode    `json:"freeOrderAuditNode,omitempty" xml:"freeOrderAuditNode,omitempty" type:"Struct"`
	FreeOrderNode         *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderNode         `json:"freeOrderNode,omitempty" xml:"freeOrderNode,omitempty" type:"Struct"`
	Name                  *string                                                                   `json:"name,omitempty" xml:"name,omitempty"`
	OrderDate             *int64                                                                    `json:"orderDate,omitempty" xml:"orderDate,omitempty"`
	OrderNode             *GetEnterpriseSupportPlanDetailResponseBodyDataNodesOrderNode             `json:"orderNode,omitempty" xml:"orderNode,omitempty" type:"Struct"`
	ServiceImplementation *GetEnterpriseSupportPlanDetailResponseBodyDataNodesServiceImplementation `json:"serviceImplementation,omitempty" xml:"serviceImplementation,omitempty" type:"Struct"`
	Status                *int32                                                                    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataNodes) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodes) SetDocNode(v *GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode) *GetEnterpriseSupportPlanDetailResponseBodyDataNodes {
	s.DocNode = v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodes) SetFinishNode(v *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFinishNode) *GetEnterpriseSupportPlanDetailResponseBodyDataNodes {
	s.FinishNode = v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodes) SetFreeOrderAuditNode(v *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderAuditNode) *GetEnterpriseSupportPlanDetailResponseBodyDataNodes {
	s.FreeOrderAuditNode = v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodes) SetFreeOrderNode(v *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderNode) *GetEnterpriseSupportPlanDetailResponseBodyDataNodes {
	s.FreeOrderNode = v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodes) SetName(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataNodes {
	s.Name = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodes) SetOrderDate(v int64) *GetEnterpriseSupportPlanDetailResponseBodyDataNodes {
	s.OrderDate = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodes) SetOrderNode(v *GetEnterpriseSupportPlanDetailResponseBodyDataNodesOrderNode) *GetEnterpriseSupportPlanDetailResponseBodyDataNodes {
	s.OrderNode = v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodes) SetServiceImplementation(v *GetEnterpriseSupportPlanDetailResponseBodyDataNodesServiceImplementation) *GetEnterpriseSupportPlanDetailResponseBodyDataNodes {
	s.ServiceImplementation = v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodes) SetStatus(v int32) *GetEnterpriseSupportPlanDetailResponseBodyDataNodes {
	s.Status = &v
	return s
}

type GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode struct {
	DocId              *int64  `json:"docId,omitempty" xml:"docId,omitempty"`
	DocName            *string `json:"docName,omitempty" xml:"docName,omitempty"`
	DocSubmitTime      *string `json:"docSubmitTime,omitempty" xml:"docSubmitTime,omitempty"`
	FileName           *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FreeOrderApplyCode *string `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	OrderId            *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode) GoString() string {
	return s.String()
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode) SetDocId(v int64) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode {
	s.DocId = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode) SetDocName(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode {
	s.DocName = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode) SetDocSubmitTime(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode {
	s.DocSubmitTime = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode) SetFileName(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode {
	s.FileName = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode) SetFreeOrderApplyCode(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode {
	s.FreeOrderApplyCode = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode) SetOrderId(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesDocNode {
	s.OrderId = &v
	return s
}

type GetEnterpriseSupportPlanDetailResponseBodyDataNodesFinishNode struct {
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataNodesFinishNode) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataNodesFinishNode) GoString() string {
	return s.String()
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFinishNode) SetFinishTime(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFinishNode {
	s.FinishTime = &v
	return s
}

type GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderAuditNode struct {
	AuditTime  *string `json:"auditTime,omitempty" xml:"auditTime,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	StatusName *string `json:"statusName,omitempty" xml:"statusName,omitempty"`
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderAuditNode) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderAuditNode) GoString() string {
	return s.String()
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderAuditNode) SetAuditTime(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderAuditNode {
	s.AuditTime = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderAuditNode) SetStatus(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderAuditNode {
	s.Status = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderAuditNode) SetStatusName(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderAuditNode {
	s.StatusName = &v
	return s
}

type GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderNode struct {
	ApplyTime *string `json:"applyTime,omitempty" xml:"applyTime,omitempty"`
	Uid       *int64  `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderNode) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderNode) GoString() string {
	return s.String()
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderNode) SetApplyTime(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderNode {
	s.ApplyTime = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderNode) SetUid(v int64) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesFreeOrderNode {
	s.Uid = &v
	return s
}

type GetEnterpriseSupportPlanDetailResponseBodyDataNodesOrderNode struct {
	PayTime *string `json:"payTime,omitempty" xml:"payTime,omitempty"`
	Uid     *int64  `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataNodesOrderNode) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataNodesOrderNode) GoString() string {
	return s.String()
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesOrderNode) SetPayTime(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesOrderNode {
	s.PayTime = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesOrderNode) SetUid(v int64) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesOrderNode {
	s.Uid = &v
	return s
}

type GetEnterpriseSupportPlanDetailResponseBodyDataNodesServiceImplementation struct {
	EndTime   *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataNodesServiceImplementation) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataNodesServiceImplementation) GoString() string {
	return s.String()
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesServiceImplementation) SetEndTime(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesServiceImplementation {
	s.EndTime = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataNodesServiceImplementation) SetStartTime(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataNodesServiceImplementation {
	s.StartTime = &v
	return s
}

type GetEnterpriseSupportPlanDetailResponseBodyDataServiceItems struct {
	Content     *string                                                                  `json:"content,omitempty" xml:"content,omitempty"`
	Desc        *string                                                                  `json:"desc,omitempty" xml:"desc,omitempty"`
	Name        *string                                                                  `json:"name,omitempty" xml:"name,omitempty"`
	OperateList []*GetEnterpriseSupportPlanDetailResponseBodyDataServiceItemsOperateList `json:"operateList,omitempty" xml:"operateList,omitempty" type:"Repeated"`
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataServiceItems) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataServiceItems) GoString() string {
	return s.String()
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataServiceItems) SetContent(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataServiceItems {
	s.Content = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataServiceItems) SetDesc(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataServiceItems {
	s.Desc = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataServiceItems) SetName(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataServiceItems {
	s.Name = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataServiceItems) SetOperateList(v []*GetEnterpriseSupportPlanDetailResponseBodyDataServiceItemsOperateList) *GetEnterpriseSupportPlanDetailResponseBodyDataServiceItems {
	s.OperateList = v
	return s
}

type GetEnterpriseSupportPlanDetailResponseBodyDataServiceItemsOperateList struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Name1 *string `json:"name1,omitempty" xml:"name1,omitempty"`
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataServiceItemsOperateList) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseSupportPlanDetailResponseBodyDataServiceItemsOperateList) GoString() string {
	return s.String()
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataServiceItemsOperateList) SetName(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataServiceItemsOperateList {
	s.Name = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponseBodyDataServiceItemsOperateList) SetName1(v string) *GetEnterpriseSupportPlanDetailResponseBodyDataServiceItemsOperateList {
	s.Name1 = &v
	return s
}

type GetEnterpriseSupportPlanDetailResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEnterpriseSupportPlanDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEnterpriseSupportPlanDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseSupportPlanDetailResponse) GoString() string {
	return s.String()
}

func (s *GetEnterpriseSupportPlanDetailResponse) SetHeaders(v map[string]*string) *GetEnterpriseSupportPlanDetailResponse {
	s.Headers = v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponse) SetStatusCode(v int32) *GetEnterpriseSupportPlanDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnterpriseSupportPlanDetailResponse) SetBody(v *GetEnterpriseSupportPlanDetailResponseBody) *GetEnterpriseSupportPlanDetailResponse {
	s.Body = v
	return s
}

type GetPreViewUrlRequest struct {
	ApplyCode *string `json:"applyCode,omitempty" xml:"applyCode,omitempty"`
	FileId    *int64  `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileKey   *string `json:"fileKey,omitempty" xml:"fileKey,omitempty"`
	OrderId   *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	Scene     *string `json:"scene,omitempty" xml:"scene,omitempty"`
}

func (s GetPreViewUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPreViewUrlRequest) GoString() string {
	return s.String()
}

func (s *GetPreViewUrlRequest) SetApplyCode(v string) *GetPreViewUrlRequest {
	s.ApplyCode = &v
	return s
}

func (s *GetPreViewUrlRequest) SetFileId(v int64) *GetPreViewUrlRequest {
	s.FileId = &v
	return s
}

func (s *GetPreViewUrlRequest) SetFileKey(v string) *GetPreViewUrlRequest {
	s.FileKey = &v
	return s
}

func (s *GetPreViewUrlRequest) SetOrderId(v string) *GetPreViewUrlRequest {
	s.OrderId = &v
	return s
}

func (s *GetPreViewUrlRequest) SetScene(v string) *GetPreViewUrlRequest {
	s.Scene = &v
	return s
}

type GetPreViewUrlResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *string `json:"data,omitempty" xml:"data,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPreViewUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPreViewUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetPreViewUrlResponseBody) SetCode(v string) *GetPreViewUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetPreViewUrlResponseBody) SetData(v string) *GetPreViewUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetPreViewUrlResponseBody) SetHttpStatusCode(v int32) *GetPreViewUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPreViewUrlResponseBody) SetMessage(v string) *GetPreViewUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetPreViewUrlResponseBody) SetRequestId(v string) *GetPreViewUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPreViewUrlResponseBody) SetSuccess(v bool) *GetPreViewUrlResponseBody {
	s.Success = &v
	return s
}

type GetPreViewUrlResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPreViewUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPreViewUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPreViewUrlResponse) GoString() string {
	return s.String()
}

func (s *GetPreViewUrlResponse) SetHeaders(v map[string]*string) *GetPreViewUrlResponse {
	s.Headers = v
	return s
}

func (s *GetPreViewUrlResponse) SetStatusCode(v int32) *GetPreViewUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPreViewUrlResponse) SetBody(v *GetPreViewUrlResponseBody) *GetPreViewUrlResponse {
	s.Body = v
	return s
}

type GetServiceDetailRequest struct {
	ApplyCode *string `json:"applyCode,omitempty" xml:"applyCode,omitempty"`
}

func (s GetServiceDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailRequest) GoString() string {
	return s.String()
}

func (s *GetServiceDetailRequest) SetApplyCode(v string) *GetServiceDetailRequest {
	s.ApplyCode = &v
	return s
}

type GetServiceDetailResponseBody struct {
	Code           *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data           *GetServiceDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpStatusCode *int32                            `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                           `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetServiceDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBody) SetCode(v string) *GetServiceDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceDetailResponseBody) SetData(v *GetServiceDetailResponseBodyData) *GetServiceDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceDetailResponseBody) SetHttpStatusCode(v int32) *GetServiceDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetServiceDetailResponseBody) SetMessage(v string) *GetServiceDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceDetailResponseBody) SetRequestId(v string) *GetServiceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceDetailResponseBody) SetSuccess(v bool) *GetServiceDetailResponseBody {
	s.Success = &v
	return s
}

type GetServiceDetailResponseBodyData struct {
	ApplierId                       *string                                              `json:"applierId,omitempty" xml:"applierId,omitempty"`
	ApplyCode                       *string                                              `json:"applyCode,omitempty" xml:"applyCode,omitempty"`
	ApplyTime                       *int64                                               `json:"applyTime,omitempty" xml:"applyTime,omitempty"`
	Appointments                    []*GetServiceDetailResponseBodyDataAppointments      `json:"appointments,omitempty" xml:"appointments,omitempty" type:"Repeated"`
	BuyUrl                          *string                                              `json:"buyUrl,omitempty" xml:"buyUrl,omitempty"`
	CreatorEmpId                    *string                                              `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	CustomerName                    *string                                              `json:"customerName,omitempty" xml:"customerName,omitempty"`
	CycleService                    *bool                                                `json:"cycleService,omitempty" xml:"cycleService,omitempty"`
	ExecutedCount                   *int64                                               `json:"executedCount,omitempty" xml:"executedCount,omitempty"`
	FinishCount                     *int64                                               `json:"finishCount,omitempty" xml:"finishCount,omitempty"`
	FormMap                         map[string]interface{}                               `json:"formMap,omitempty" xml:"formMap,omitempty"`
	GmtCreate                       *string                                              `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified                     *string                                              `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                              *int64                                               `json:"id,omitempty" xml:"id,omitempty"`
	MergeSolutionAndReporterOneStep *bool                                                `json:"mergeSolutionAndReporterOneStep,omitempty" xml:"mergeSolutionAndReporterOneStep,omitempty"`
	ModifierEmpId                   *string                                              `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	PackDetails                     []map[string]interface{}                             `json:"packDetails,omitempty" xml:"packDetails,omitempty" type:"Repeated"`
	PayOrders                       []*GetServiceDetailResponseBodyDataPayOrders         `json:"payOrders,omitempty" xml:"payOrders,omitempty" type:"Repeated"`
	PayUrl                          *string                                              `json:"payUrl,omitempty" xml:"payUrl,omitempty"`
	PerformanceOrders               []*GetServiceDetailResponseBodyDataPerformanceOrders `json:"performanceOrders,omitempty" xml:"performanceOrders,omitempty" type:"Repeated"`
	PerformancePacks                []*GetServiceDetailResponseBodyDataPerformancePacks  `json:"performancePacks,omitempty" xml:"performancePacks,omitempty" type:"Repeated"`
	ReneWalUrl                      *string                                              `json:"reneWalUrl,omitempty" xml:"reneWalUrl,omitempty"`
	ServiceCode                     *string                                              `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	ServiceName                     *string                                              `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	ServiceReports                  []*GetServiceDetailResponseBodyDataServiceReports    `json:"serviceReports,omitempty" xml:"serviceReports,omitempty" type:"Repeated"`
	ServiceTimeRange                []*int64                                             `json:"serviceTimeRange,omitempty" xml:"serviceTimeRange,omitempty" type:"Repeated"`
	Status                          *string                                              `json:"status,omitempty" xml:"status,omitempty"`
	StatusCode                      *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	StatusStr                       *string                                              `json:"statusStr,omitempty" xml:"statusStr,omitempty"`
	TermOfValidity                  *string                                              `json:"termOfValidity,omitempty" xml:"termOfValidity,omitempty"`
	TotalPack                       *int32                                               `json:"totalPack,omitempty" xml:"totalPack,omitempty"`
	Type                            *string                                              `json:"type,omitempty" xml:"type,omitempty"`
	UsePack                         *int64                                               `json:"usePack,omitempty" xml:"usePack,omitempty"`
}

func (s GetServiceDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyData) SetApplierId(v string) *GetServiceDetailResponseBodyData {
	s.ApplierId = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetApplyCode(v string) *GetServiceDetailResponseBodyData {
	s.ApplyCode = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetApplyTime(v int64) *GetServiceDetailResponseBodyData {
	s.ApplyTime = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetAppointments(v []*GetServiceDetailResponseBodyDataAppointments) *GetServiceDetailResponseBodyData {
	s.Appointments = v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetBuyUrl(v string) *GetServiceDetailResponseBodyData {
	s.BuyUrl = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetCreatorEmpId(v string) *GetServiceDetailResponseBodyData {
	s.CreatorEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetCustomerName(v string) *GetServiceDetailResponseBodyData {
	s.CustomerName = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetCycleService(v bool) *GetServiceDetailResponseBodyData {
	s.CycleService = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetExecutedCount(v int64) *GetServiceDetailResponseBodyData {
	s.ExecutedCount = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetFinishCount(v int64) *GetServiceDetailResponseBodyData {
	s.FinishCount = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetFormMap(v map[string]interface{}) *GetServiceDetailResponseBodyData {
	s.FormMap = v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetGmtCreate(v string) *GetServiceDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetGmtModified(v string) *GetServiceDetailResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetId(v int64) *GetServiceDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetMergeSolutionAndReporterOneStep(v bool) *GetServiceDetailResponseBodyData {
	s.MergeSolutionAndReporterOneStep = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetModifierEmpId(v string) *GetServiceDetailResponseBodyData {
	s.ModifierEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetPackDetails(v []map[string]interface{}) *GetServiceDetailResponseBodyData {
	s.PackDetails = v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetPayOrders(v []*GetServiceDetailResponseBodyDataPayOrders) *GetServiceDetailResponseBodyData {
	s.PayOrders = v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetPayUrl(v string) *GetServiceDetailResponseBodyData {
	s.PayUrl = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetPerformanceOrders(v []*GetServiceDetailResponseBodyDataPerformanceOrders) *GetServiceDetailResponseBodyData {
	s.PerformanceOrders = v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetPerformancePacks(v []*GetServiceDetailResponseBodyDataPerformancePacks) *GetServiceDetailResponseBodyData {
	s.PerformancePacks = v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetReneWalUrl(v string) *GetServiceDetailResponseBodyData {
	s.ReneWalUrl = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetServiceCode(v string) *GetServiceDetailResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetServiceName(v string) *GetServiceDetailResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetServiceReports(v []*GetServiceDetailResponseBodyDataServiceReports) *GetServiceDetailResponseBodyData {
	s.ServiceReports = v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetServiceTimeRange(v []*int64) *GetServiceDetailResponseBodyData {
	s.ServiceTimeRange = v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetStatus(v string) *GetServiceDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetStatusCode(v int32) *GetServiceDetailResponseBodyData {
	s.StatusCode = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetStatusStr(v string) *GetServiceDetailResponseBodyData {
	s.StatusStr = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetTermOfValidity(v string) *GetServiceDetailResponseBodyData {
	s.TermOfValidity = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetTotalPack(v int32) *GetServiceDetailResponseBodyData {
	s.TotalPack = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetType(v string) *GetServiceDetailResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetUsePack(v int64) *GetServiceDetailResponseBodyData {
	s.UsePack = &v
	return s
}

type GetServiceDetailResponseBodyDataAppointments struct {
	HuhangId     *int64  `json:"huhangId,omitempty" xml:"huhangId,omitempty"`
	PurchaseCode *int32  `json:"purchaseCode,omitempty" xml:"purchaseCode,omitempty"`
	PurchaseDesc *string `json:"purchaseDesc,omitempty" xml:"purchaseDesc,omitempty"`
	SupportDays  *int32  `json:"supportDays,omitempty" xml:"supportDays,omitempty"`
	TravelDays   *int32  `json:"travelDays,omitempty" xml:"travelDays,omitempty"`
}

func (s GetServiceDetailResponseBodyDataAppointments) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataAppointments) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataAppointments) SetHuhangId(v int64) *GetServiceDetailResponseBodyDataAppointments {
	s.HuhangId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataAppointments) SetPurchaseCode(v int32) *GetServiceDetailResponseBodyDataAppointments {
	s.PurchaseCode = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataAppointments) SetPurchaseDesc(v string) *GetServiceDetailResponseBodyDataAppointments {
	s.PurchaseDesc = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataAppointments) SetSupportDays(v int32) *GetServiceDetailResponseBodyDataAppointments {
	s.SupportDays = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataAppointments) SetTravelDays(v int32) *GetServiceDetailResponseBodyDataAppointments {
	s.TravelDays = &v
	return s
}

type GetServiceDetailResponseBodyDataPayOrders struct {
	Amount               *string                `json:"amount,omitempty" xml:"amount,omitempty"`
	CompassCommodityCode *string                `json:"compassCommodityCode,omitempty" xml:"compassCommodityCode,omitempty"`
	CompassCommodityName *string                `json:"compassCommodityName,omitempty" xml:"compassCommodityName,omitempty"`
	CreatorEmpId         *string                `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	GmtCreate            *string                `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified          *string                `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                   *int64                 `json:"id,omitempty" xml:"id,omitempty"`
	ModifierEmpId        *string                `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	Operate              map[string]interface{} `json:"operate,omitempty" xml:"operate,omitempty"`
	OrderDetail          interface{}            `json:"orderDetail,omitempty" xml:"orderDetail,omitempty"`
	OrderId              *int64                 `json:"orderId,omitempty" xml:"orderId,omitempty"`
	OriginalPrice        *float64               `json:"originalPrice,omitempty" xml:"originalPrice,omitempty"`
	PayAmount            *float64               `json:"payAmount,omitempty" xml:"payAmount,omitempty"`
	PayTime              *string                `json:"payTime,omitempty" xml:"payTime,omitempty"`
	ProductName          *string                `json:"productName,omitempty" xml:"productName,omitempty"`
	ReneWalUrl           *string                `json:"reneWalUrl,omitempty" xml:"reneWalUrl,omitempty"`
	ServiceContentMap    map[string]interface{} `json:"serviceContentMap,omitempty" xml:"serviceContentMap,omitempty"`
	Status               *int32                 `json:"status,omitempty" xml:"status,omitempty"`
	StatusStr            *string                `json:"statusStr,omitempty" xml:"statusStr,omitempty"`
	SupportDays          *int32                 `json:"supportDays,omitempty" xml:"supportDays,omitempty"`
	Uid                  *string                `json:"uid,omitempty" xml:"uid,omitempty"`
	Url                  *string                `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPayOrders) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPayOrders) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetAmount(v string) *GetServiceDetailResponseBodyDataPayOrders {
	s.Amount = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetCompassCommodityCode(v string) *GetServiceDetailResponseBodyDataPayOrders {
	s.CompassCommodityCode = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetCompassCommodityName(v string) *GetServiceDetailResponseBodyDataPayOrders {
	s.CompassCommodityName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetCreatorEmpId(v string) *GetServiceDetailResponseBodyDataPayOrders {
	s.CreatorEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPayOrders {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPayOrders {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetId(v int64) *GetServiceDetailResponseBodyDataPayOrders {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetModifierEmpId(v string) *GetServiceDetailResponseBodyDataPayOrders {
	s.ModifierEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetOperate(v map[string]interface{}) *GetServiceDetailResponseBodyDataPayOrders {
	s.Operate = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetOrderDetail(v interface{}) *GetServiceDetailResponseBodyDataPayOrders {
	s.OrderDetail = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetOrderId(v int64) *GetServiceDetailResponseBodyDataPayOrders {
	s.OrderId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetOriginalPrice(v float64) *GetServiceDetailResponseBodyDataPayOrders {
	s.OriginalPrice = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetPayAmount(v float64) *GetServiceDetailResponseBodyDataPayOrders {
	s.PayAmount = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetPayTime(v string) *GetServiceDetailResponseBodyDataPayOrders {
	s.PayTime = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetProductName(v string) *GetServiceDetailResponseBodyDataPayOrders {
	s.ProductName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetReneWalUrl(v string) *GetServiceDetailResponseBodyDataPayOrders {
	s.ReneWalUrl = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetServiceContentMap(v map[string]interface{}) *GetServiceDetailResponseBodyDataPayOrders {
	s.ServiceContentMap = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetStatus(v int32) *GetServiceDetailResponseBodyDataPayOrders {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetStatusStr(v string) *GetServiceDetailResponseBodyDataPayOrders {
	s.StatusStr = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetSupportDays(v int32) *GetServiceDetailResponseBodyDataPayOrders {
	s.SupportDays = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetUid(v string) *GetServiceDetailResponseBodyDataPayOrders {
	s.Uid = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPayOrders) SetUrl(v string) *GetServiceDetailResponseBodyDataPayOrders {
	s.Url = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrders struct {
	ApplyFileVOList                 []*GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList     `json:"applyFileVOList,omitempty" xml:"applyFileVOList,omitempty" type:"Repeated"`
	AppointmentCode                 *string                                                                 `json:"appointmentCode,omitempty" xml:"appointmentCode,omitempty"`
	AppointmentEndTime              *int64                                                                  `json:"appointmentEndTime,omitempty" xml:"appointmentEndTime,omitempty"`
	AppointmentId                   *string                                                                 `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	AppointmentPassTime             *int64                                                                  `json:"appointmentPassTime,omitempty" xml:"appointmentPassTime,omitempty"`
	AppointmentTime                 *int64                                                                  `json:"appointmentTime,omitempty" xml:"appointmentTime,omitempty"`
	CommodityDesc                   *string                                                                 `json:"commodityDesc,omitempty" xml:"commodityDesc,omitempty"`
	CreatorEmpId                    *string                                                                 `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	CycleService                    *bool                                                                   `json:"cycleService,omitempty" xml:"cycleService,omitempty"`
	ExtList                         []*GetServiceDetailResponseBodyDataPerformanceOrdersExtList             `json:"extList,omitempty" xml:"extList,omitempty" type:"Repeated"`
	GmtCreate                       *string                                                                 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified                     *string                                                                 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                              *int64                                                                  `json:"id,omitempty" xml:"id,omitempty"`
	MergeSolutionAndReporterOneStep *bool                                                                   `json:"mergeSolutionAndReporterOneStep,omitempty" xml:"mergeSolutionAndReporterOneStep,omitempty"`
	ModifierEmpId                   *string                                                                 `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	OrderDetail                     interface{}                                                             `json:"orderDetail,omitempty" xml:"orderDetail,omitempty"`
	OrderId                         *int64                                                                  `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PackCount                       *int32                                                                  `json:"packCount,omitempty" xml:"packCount,omitempty"`
	PackDetails                     []map[string]interface{}                                                `json:"packDetails,omitempty" xml:"packDetails,omitempty" type:"Repeated"`
	PerformanceNodeDTOS             []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformanceNodeDTOS `json:"performanceNodeDTOS,omitempty" xml:"performanceNodeDTOS,omitempty" type:"Repeated"`
	PerformancePacks                []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks    `json:"performancePacks,omitempty" xml:"performancePacks,omitempty" type:"Repeated"`
	PurchasePackCode                *int32                                                                  `json:"purchasePackCode,omitempty" xml:"purchasePackCode,omitempty"`
	ServiceApplyId                  *int64                                                                  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	ServiceMonthReports             []*GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports `json:"serviceMonthReports,omitempty" xml:"serviceMonthReports,omitempty" type:"Repeated"`
	ServiceReports                  []*GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports      `json:"serviceReports,omitempty" xml:"serviceReports,omitempty" type:"Repeated"`
	ServiceSchemes                  []*GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes      `json:"serviceSchemes,omitempty" xml:"serviceSchemes,omitempty" type:"Repeated"`
	Status                          *int32                                                                  `json:"status,omitempty" xml:"status,omitempty"`
	StatusStr                       *string                                                                 `json:"statusStr,omitempty" xml:"statusStr,omitempty"`
	SupportTime                     []*int64                                                                `json:"supportTime,omitempty" xml:"supportTime,omitempty" type:"Repeated"`
	TamEngineers                    []*GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers        `json:"tamEngineers,omitempty" xml:"tamEngineers,omitempty" type:"Repeated"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrders) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrders) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetApplyFileVOList(v []*GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.ApplyFileVOList = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetAppointmentCode(v string) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.AppointmentCode = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetAppointmentEndTime(v int64) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.AppointmentEndTime = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetAppointmentPassTime(v int64) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.AppointmentPassTime = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetAppointmentTime(v int64) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.AppointmentTime = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetCommodityDesc(v string) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.CommodityDesc = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetCreatorEmpId(v string) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.CreatorEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetCycleService(v bool) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.CycleService = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetExtList(v []*GetServiceDetailResponseBodyDataPerformanceOrdersExtList) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.ExtList = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetMergeSolutionAndReporterOneStep(v bool) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.MergeSolutionAndReporterOneStep = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetModifierEmpId(v string) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.ModifierEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetOrderDetail(v interface{}) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.OrderDetail = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetOrderId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.OrderId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetPackCount(v int32) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.PackCount = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetPackDetails(v []map[string]interface{}) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.PackDetails = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetPerformanceNodeDTOS(v []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformanceNodeDTOS) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.PerformanceNodeDTOS = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetPerformancePacks(v []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.PerformancePacks = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetPurchasePackCode(v int32) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.PurchasePackCode = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetServiceMonthReports(v []*GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.ServiceMonthReports = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetServiceReports(v []*GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.ServiceReports = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetServiceSchemes(v []*GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.ServiceSchemes = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetStatusStr(v string) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.StatusStr = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetSupportTime(v []*int64) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.SupportTime = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrders) SetTamEngineers(v []*GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers) *GetServiceDetailResponseBodyDataPerformanceOrders {
	s.TamEngineers = v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList) SetBatchGroup(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList {
	s.BatchGroup = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList) SetCustomerId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList {
	s.CustomerId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList) SetFileName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList {
	s.FileName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList) SetFileType(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList {
	s.FileType = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList) SetId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList) SetRemarke(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList {
	s.Remarke = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList) SetUrl(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersApplyFileVOList {
	s.Url = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrdersExtList struct {
	KeyCode *string     `json:"keyCode,omitempty" xml:"keyCode,omitempty"`
	Name    *string     `json:"name,omitempty" xml:"name,omitempty"`
	Value   interface{} `json:"value,omitempty" xml:"value,omitempty"`
	View    *string     `json:"view,omitempty" xml:"view,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersExtList) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersExtList) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersExtList) SetKeyCode(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersExtList {
	s.KeyCode = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersExtList) SetName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersExtList {
	s.Name = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersExtList) SetValue(v interface{}) *GetServiceDetailResponseBodyDataPerformanceOrdersExtList {
	s.Value = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersExtList) SetView(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersExtList {
	s.View = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrdersPerformanceNodeDTOS struct {
	Display    *bool       `json:"display,omitempty" xml:"display,omitempty"`
	ExtendInfo interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Index      *int32      `json:"index,omitempty" xml:"index,omitempty"`
	NodeName   *string     `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	Status     *int32      `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformanceNodeDTOS) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformanceNodeDTOS) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformanceNodeDTOS) SetDisplay(v bool) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformanceNodeDTOS {
	s.Display = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformanceNodeDTOS) SetExtendInfo(v interface{}) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformanceNodeDTOS {
	s.ExtendInfo = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformanceNodeDTOS) SetIndex(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformanceNodeDTOS {
	s.Index = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformanceNodeDTOS) SetNodeName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformanceNodeDTOS {
	s.NodeName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformanceNodeDTOS) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformanceNodeDTOS {
	s.Status = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks struct {
	ApplyFileVOList                 []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList     `json:"applyFileVOList,omitempty" xml:"applyFileVOList,omitempty" type:"Repeated"`
	AppointmentCode                 *string                                                                                 `json:"appointmentCode,omitempty" xml:"appointmentCode,omitempty"`
	AppointmentEndTime              *int64                                                                                  `json:"appointmentEndTime,omitempty" xml:"appointmentEndTime,omitempty"`
	AppointmentId                   *string                                                                                 `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	AppointmentPassTime             *int64                                                                                  `json:"appointmentPassTime,omitempty" xml:"appointmentPassTime,omitempty"`
	AppointmentTime                 *int64                                                                                  `json:"appointmentTime,omitempty" xml:"appointmentTime,omitempty"`
	CommodityDesc                   *string                                                                                 `json:"commodityDesc,omitempty" xml:"commodityDesc,omitempty"`
	CreatorEmpId                    *string                                                                                 `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	CycleService                    *bool                                                                                   `json:"cycleService,omitempty" xml:"cycleService,omitempty"`
	ExtList                         []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksExtList             `json:"extList,omitempty" xml:"extList,omitempty" type:"Repeated"`
	GmtCreate                       *string                                                                                 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified                     *string                                                                                 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                              *int64                                                                                  `json:"id,omitempty" xml:"id,omitempty"`
	MergeSolutionAndReporterOneStep *bool                                                                                   `json:"mergeSolutionAndReporterOneStep,omitempty" xml:"mergeSolutionAndReporterOneStep,omitempty"`
	ModifierEmpId                   *string                                                                                 `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	OrderDetail                     interface{}                                                                             `json:"orderDetail,omitempty" xml:"orderDetail,omitempty"`
	OrderId                         *int64                                                                                  `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PerformanceNodeDTOS             []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksPerformanceNodeDTOS `json:"performanceNodeDTOS,omitempty" xml:"performanceNodeDTOS,omitempty" type:"Repeated"`
	PurchasePackCode                *int32                                                                                  `json:"purchasePackCode,omitempty" xml:"purchasePackCode,omitempty"`
	ServiceApplyId                  *int64                                                                                  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	ServiceMonthReports             []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports `json:"serviceMonthReports,omitempty" xml:"serviceMonthReports,omitempty" type:"Repeated"`
	ServiceReports                  []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports      `json:"serviceReports,omitempty" xml:"serviceReports,omitempty" type:"Repeated"`
	ServiceSchemes                  []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes      `json:"serviceSchemes,omitempty" xml:"serviceSchemes,omitempty" type:"Repeated"`
	Status                          *int32                                                                                  `json:"status,omitempty" xml:"status,omitempty"`
	StatusStr                       *string                                                                                 `json:"statusStr,omitempty" xml:"statusStr,omitempty"`
	SupportTime                     []*int64                                                                                `json:"supportTime,omitempty" xml:"supportTime,omitempty" type:"Repeated"`
	TamEngineers                    []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers        `json:"tamEngineers,omitempty" xml:"tamEngineers,omitempty" type:"Repeated"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetApplyFileVOList(v []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.ApplyFileVOList = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetAppointmentCode(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.AppointmentCode = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetAppointmentEndTime(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.AppointmentEndTime = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetAppointmentPassTime(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.AppointmentPassTime = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetAppointmentTime(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.AppointmentTime = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetCommodityDesc(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.CommodityDesc = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetCreatorEmpId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.CreatorEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetCycleService(v bool) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.CycleService = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetExtList(v []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksExtList) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.ExtList = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetMergeSolutionAndReporterOneStep(v bool) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.MergeSolutionAndReporterOneStep = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetModifierEmpId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.ModifierEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetOrderDetail(v interface{}) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.OrderDetail = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetOrderId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.OrderId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetPerformanceNodeDTOS(v []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksPerformanceNodeDTOS) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.PerformanceNodeDTOS = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetPurchasePackCode(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.PurchasePackCode = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetServiceMonthReports(v []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.ServiceMonthReports = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetServiceReports(v []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.ServiceReports = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetServiceSchemes(v []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.ServiceSchemes = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetStatusStr(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.StatusStr = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetSupportTime(v []*int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.SupportTime = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks) SetTamEngineers(v []*GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacks {
	s.TamEngineers = v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList) SetBatchGroup(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList {
	s.BatchGroup = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList) SetCustomerId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList {
	s.CustomerId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList) SetFileName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList {
	s.FileName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList) SetFileType(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList {
	s.FileType = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList) SetId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList) SetRemarke(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList {
	s.Remarke = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList) SetUrl(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksApplyFileVOList {
	s.Url = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksExtList struct {
	KeyCode *string     `json:"keyCode,omitempty" xml:"keyCode,omitempty"`
	Name    *string     `json:"name,omitempty" xml:"name,omitempty"`
	Value   interface{} `json:"value,omitempty" xml:"value,omitempty"`
	View    *string     `json:"view,omitempty" xml:"view,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksExtList) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksExtList) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksExtList) SetKeyCode(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksExtList {
	s.KeyCode = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksExtList) SetName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksExtList {
	s.Name = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksExtList) SetValue(v interface{}) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksExtList {
	s.Value = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksExtList) SetView(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksExtList {
	s.View = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksPerformanceNodeDTOS struct {
	Display    *bool       `json:"display,omitempty" xml:"display,omitempty"`
	ExtendInfo interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Index      *int32      `json:"index,omitempty" xml:"index,omitempty"`
	NodeName   *string     `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	Status     *int32      `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksPerformanceNodeDTOS) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksPerformanceNodeDTOS) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksPerformanceNodeDTOS) SetDisplay(v bool) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	s.Display = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksPerformanceNodeDTOS) SetExtendInfo(v interface{}) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	s.ExtendInfo = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksPerformanceNodeDTOS) SetIndex(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	s.Index = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksPerformanceNodeDTOS) SetNodeName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	s.NodeName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksPerformanceNodeDTOS) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	s.Status = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports) SetBatchGroup(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports {
	s.BatchGroup = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports) SetCustomerId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports {
	s.CustomerId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports) SetFileName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports {
	s.FileName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports) SetFileType(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports {
	s.FileType = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports) SetId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports) SetRemarke(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports {
	s.Remarke = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports) SetUrl(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceMonthReports {
	s.Url = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports) SetBatchGroup(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports {
	s.BatchGroup = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports) SetCustomerId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports {
	s.CustomerId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports) SetFileName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports {
	s.FileName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports) SetFileType(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports {
	s.FileType = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports) SetId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports) SetRemarke(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports {
	s.Remarke = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports) SetUrl(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceReports {
	s.Url = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes) SetBatchGroup(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes {
	s.BatchGroup = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes) SetCustomerId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes {
	s.CustomerId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes) SetFileName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes {
	s.FileName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes) SetFileType(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes {
	s.FileType = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes) SetId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes) SetRemarke(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes {
	s.Remarke = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes) SetUrl(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksServiceSchemes {
	s.Url = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers struct {
	CreatorEmpId  *string `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	GmtCreate     *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified   *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	HrStatus      *string `json:"hrStatus,omitempty" xml:"hrStatus,omitempty"`
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LastName      *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
	ModifierEmpId *string `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	NickNameEn    *string `json:"nickNameEn,omitempty" xml:"nickNameEn,omitempty"`
	RealmId       *int64  `json:"realmId,omitempty" xml:"realmId,omitempty"`
	Workid        *string `json:"workid,omitempty" xml:"workid,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers) SetCreatorEmpId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers {
	s.CreatorEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers) SetHrStatus(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers {
	s.HrStatus = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers) SetId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers) SetLastName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers {
	s.LastName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers) SetModifierEmpId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers {
	s.ModifierEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers) SetName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers {
	s.Name = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers) SetNickNameEn(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers {
	s.NickNameEn = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers) SetRealmId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers {
	s.RealmId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers) SetWorkid(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersPerformancePacksTamEngineers {
	s.Workid = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports) SetBatchGroup(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports {
	s.BatchGroup = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports) SetCustomerId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports {
	s.CustomerId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports) SetFileName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports {
	s.FileName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports) SetFileType(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports {
	s.FileType = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports) SetId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports) SetRemarke(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports {
	s.Remarke = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports) SetUrl(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceMonthReports {
	s.Url = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports) SetBatchGroup(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports {
	s.BatchGroup = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports) SetCustomerId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports {
	s.CustomerId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports) SetFileName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports {
	s.FileName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports) SetFileType(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports {
	s.FileType = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports) SetId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports) SetRemarke(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports {
	s.Remarke = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports) SetUrl(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceReports {
	s.Url = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes) SetBatchGroup(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes {
	s.BatchGroup = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes) SetCustomerId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes {
	s.CustomerId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes) SetFileName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes {
	s.FileName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes) SetFileType(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes {
	s.FileType = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes) SetId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes) SetRemarke(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes {
	s.Remarke = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes) SetUrl(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersServiceSchemes {
	s.Url = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers struct {
	CreatorEmpId  *string `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	GmtCreate     *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified   *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	HrStatus      *string `json:"hrStatus,omitempty" xml:"hrStatus,omitempty"`
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LastName      *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
	ModifierEmpId *string `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	NickNameEn    *string `json:"nickNameEn,omitempty" xml:"nickNameEn,omitempty"`
	RealmId       *int64  `json:"realmId,omitempty" xml:"realmId,omitempty"`
	Workid        *string `json:"workid,omitempty" xml:"workid,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers) SetCreatorEmpId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers {
	s.CreatorEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers) SetHrStatus(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers {
	s.HrStatus = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers) SetId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers) SetLastName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers {
	s.LastName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers) SetModifierEmpId(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers {
	s.ModifierEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers) SetName(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers {
	s.Name = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers) SetNickNameEn(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers {
	s.NickNameEn = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers) SetRealmId(v int64) *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers {
	s.RealmId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers) SetWorkid(v string) *GetServiceDetailResponseBodyDataPerformanceOrdersTamEngineers {
	s.Workid = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformancePacks struct {
	ApplyFileVOList                 []*GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList     `json:"applyFileVOList,omitempty" xml:"applyFileVOList,omitempty" type:"Repeated"`
	AppointmentCode                 *string                                                                `json:"appointmentCode,omitempty" xml:"appointmentCode,omitempty"`
	AppointmentEndTime              *int64                                                                 `json:"appointmentEndTime,omitempty" xml:"appointmentEndTime,omitempty"`
	AppointmentId                   *string                                                                `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	AppointmentPassTime             *int64                                                                 `json:"appointmentPassTime,omitempty" xml:"appointmentPassTime,omitempty"`
	AppointmentTime                 *int64                                                                 `json:"appointmentTime,omitempty" xml:"appointmentTime,omitempty"`
	CommodityDesc                   *string                                                                `json:"commodityDesc,omitempty" xml:"commodityDesc,omitempty"`
	CreatorEmpId                    *string                                                                `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	CycleService                    *bool                                                                  `json:"cycleService,omitempty" xml:"cycleService,omitempty"`
	ExtList                         []*GetServiceDetailResponseBodyDataPerformancePacksExtList             `json:"extList,omitempty" xml:"extList,omitempty" type:"Repeated"`
	GmtCreate                       *string                                                                `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified                     *string                                                                `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                              *int64                                                                 `json:"id,omitempty" xml:"id,omitempty"`
	MergeSolutionAndReporterOneStep *bool                                                                  `json:"mergeSolutionAndReporterOneStep,omitempty" xml:"mergeSolutionAndReporterOneStep,omitempty"`
	ModifierEmpId                   *string                                                                `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	OrderDetail                     interface{}                                                            `json:"orderDetail,omitempty" xml:"orderDetail,omitempty"`
	OrderId                         *int64                                                                 `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PerformanceNodeDTOS             []*GetServiceDetailResponseBodyDataPerformancePacksPerformanceNodeDTOS `json:"performanceNodeDTOS,omitempty" xml:"performanceNodeDTOS,omitempty" type:"Repeated"`
	PurchasePackCode                *int32                                                                 `json:"purchasePackCode,omitempty" xml:"purchasePackCode,omitempty"`
	ServiceApplyId                  *int64                                                                 `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	ServiceMonthReports             []*GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports `json:"serviceMonthReports,omitempty" xml:"serviceMonthReports,omitempty" type:"Repeated"`
	ServiceReports                  []*GetServiceDetailResponseBodyDataPerformancePacksServiceReports      `json:"serviceReports,omitempty" xml:"serviceReports,omitempty" type:"Repeated"`
	ServiceSchemes                  []*GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes      `json:"serviceSchemes,omitempty" xml:"serviceSchemes,omitempty" type:"Repeated"`
	Status                          *int32                                                                 `json:"status,omitempty" xml:"status,omitempty"`
	StatusStr                       *string                                                                `json:"statusStr,omitempty" xml:"statusStr,omitempty"`
	SupportTime                     []*int64                                                               `json:"supportTime,omitempty" xml:"supportTime,omitempty" type:"Repeated"`
	TamEngineers                    []*GetServiceDetailResponseBodyDataPerformancePacksTamEngineers        `json:"tamEngineers,omitempty" xml:"tamEngineers,omitempty" type:"Repeated"`
}

func (s GetServiceDetailResponseBodyDataPerformancePacks) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformancePacks) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetApplyFileVOList(v []*GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.ApplyFileVOList = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetAppointmentCode(v string) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.AppointmentCode = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetAppointmentEndTime(v int64) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.AppointmentEndTime = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetAppointmentPassTime(v int64) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.AppointmentPassTime = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetAppointmentTime(v int64) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.AppointmentTime = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetCommodityDesc(v string) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.CommodityDesc = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetCreatorEmpId(v string) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.CreatorEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetCycleService(v bool) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.CycleService = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetExtList(v []*GetServiceDetailResponseBodyDataPerformancePacksExtList) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.ExtList = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetId(v int64) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetMergeSolutionAndReporterOneStep(v bool) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.MergeSolutionAndReporterOneStep = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetModifierEmpId(v string) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.ModifierEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetOrderDetail(v interface{}) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.OrderDetail = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetOrderId(v int64) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.OrderId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetPerformanceNodeDTOS(v []*GetServiceDetailResponseBodyDataPerformancePacksPerformanceNodeDTOS) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.PerformanceNodeDTOS = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetPurchasePackCode(v int32) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.PurchasePackCode = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetServiceMonthReports(v []*GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.ServiceMonthReports = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetServiceReports(v []*GetServiceDetailResponseBodyDataPerformancePacksServiceReports) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.ServiceReports = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetServiceSchemes(v []*GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.ServiceSchemes = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetStatusStr(v string) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.StatusStr = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetSupportTime(v []*int64) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.SupportTime = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacks) SetTamEngineers(v []*GetServiceDetailResponseBodyDataPerformancePacksTamEngineers) *GetServiceDetailResponseBodyDataPerformancePacks {
	s.TamEngineers = v
	return s
}

type GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList) SetBatchGroup(v string) *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList {
	s.BatchGroup = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList) SetCustomerId(v string) *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList {
	s.CustomerId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList) SetFileName(v string) *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList {
	s.FileName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList) SetFileType(v int32) *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList {
	s.FileType = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList) SetId(v int64) *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList) SetRemarke(v string) *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList {
	s.Remarke = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList) SetUrl(v string) *GetServiceDetailResponseBodyDataPerformancePacksApplyFileVOList {
	s.Url = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformancePacksExtList struct {
	KeyCode *string     `json:"keyCode,omitempty" xml:"keyCode,omitempty"`
	Name    *string     `json:"name,omitempty" xml:"name,omitempty"`
	Value   interface{} `json:"value,omitempty" xml:"value,omitempty"`
	View    *string     `json:"view,omitempty" xml:"view,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformancePacksExtList) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformancePacksExtList) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksExtList) SetKeyCode(v string) *GetServiceDetailResponseBodyDataPerformancePacksExtList {
	s.KeyCode = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksExtList) SetName(v string) *GetServiceDetailResponseBodyDataPerformancePacksExtList {
	s.Name = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksExtList) SetValue(v interface{}) *GetServiceDetailResponseBodyDataPerformancePacksExtList {
	s.Value = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksExtList) SetView(v string) *GetServiceDetailResponseBodyDataPerformancePacksExtList {
	s.View = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformancePacksPerformanceNodeDTOS struct {
	Display    *bool       `json:"display,omitempty" xml:"display,omitempty"`
	ExtendInfo interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Index      *int32      `json:"index,omitempty" xml:"index,omitempty"`
	NodeName   *string     `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	Status     *int32      `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformancePacksPerformanceNodeDTOS) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformancePacksPerformanceNodeDTOS) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksPerformanceNodeDTOS) SetDisplay(v bool) *GetServiceDetailResponseBodyDataPerformancePacksPerformanceNodeDTOS {
	s.Display = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksPerformanceNodeDTOS) SetExtendInfo(v interface{}) *GetServiceDetailResponseBodyDataPerformancePacksPerformanceNodeDTOS {
	s.ExtendInfo = v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksPerformanceNodeDTOS) SetIndex(v int32) *GetServiceDetailResponseBodyDataPerformancePacksPerformanceNodeDTOS {
	s.Index = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksPerformanceNodeDTOS) SetNodeName(v string) *GetServiceDetailResponseBodyDataPerformancePacksPerformanceNodeDTOS {
	s.NodeName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksPerformanceNodeDTOS) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformancePacksPerformanceNodeDTOS {
	s.Status = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports) SetBatchGroup(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports {
	s.BatchGroup = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports) SetCustomerId(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports {
	s.CustomerId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports) SetFileName(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports {
	s.FileName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports) SetFileType(v int32) *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports {
	s.FileType = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports) SetId(v int64) *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports) SetRemarke(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports {
	s.Remarke = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports) SetUrl(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceMonthReports {
	s.Url = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformancePacksServiceReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformancePacksServiceReports) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformancePacksServiceReports) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceReports) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceReports {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceReports) SetBatchGroup(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceReports {
	s.BatchGroup = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceReports) SetCustomerId(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceReports {
	s.CustomerId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceReports) SetFileName(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceReports {
	s.FileName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceReports) SetFileType(v int32) *GetServiceDetailResponseBodyDataPerformancePacksServiceReports {
	s.FileType = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceReports) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceReports {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceReports) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceReports {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceReports) SetId(v int64) *GetServiceDetailResponseBodyDataPerformancePacksServiceReports {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceReports) SetRemarke(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceReports {
	s.Remarke = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceReports) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataPerformancePacksServiceReports {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceReports) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformancePacksServiceReports {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceReports) SetUrl(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceReports {
	s.Url = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes) SetBatchGroup(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes {
	s.BatchGroup = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes) SetCustomerId(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes {
	s.CustomerId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes) SetFileName(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes {
	s.FileName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes) SetFileType(v int32) *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes {
	s.FileType = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes) SetId(v int64) *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes) SetRemarke(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes {
	s.Remarke = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes) SetStatus(v int32) *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes) SetUrl(v string) *GetServiceDetailResponseBodyDataPerformancePacksServiceSchemes {
	s.Url = &v
	return s
}

type GetServiceDetailResponseBodyDataPerformancePacksTamEngineers struct {
	CreatorEmpId  *string `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	GmtCreate     *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified   *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	HrStatus      *string `json:"hrStatus,omitempty" xml:"hrStatus,omitempty"`
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LastName      *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
	ModifierEmpId *string `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	NickNameEn    *string `json:"nickNameEn,omitempty" xml:"nickNameEn,omitempty"`
	RealmId       *int64  `json:"realmId,omitempty" xml:"realmId,omitempty"`
	Workid        *string `json:"workid,omitempty" xml:"workid,omitempty"`
}

func (s GetServiceDetailResponseBodyDataPerformancePacksTamEngineers) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataPerformancePacksTamEngineers) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers) SetCreatorEmpId(v string) *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers {
	s.CreatorEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers) SetGmtModified(v string) *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers) SetHrStatus(v string) *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers {
	s.HrStatus = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers) SetId(v int64) *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers) SetLastName(v string) *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers {
	s.LastName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers) SetModifierEmpId(v string) *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers {
	s.ModifierEmpId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers) SetName(v string) *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers {
	s.Name = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers) SetNickNameEn(v string) *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers {
	s.NickNameEn = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers) SetRealmId(v int64) *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers {
	s.RealmId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers) SetWorkid(v string) *GetServiceDetailResponseBodyDataPerformancePacksTamEngineers {
	s.Workid = &v
	return s
}

type GetServiceDetailResponseBodyDataServiceReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetServiceDetailResponseBodyDataServiceReports) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataServiceReports) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataServiceReports) SetAppointmentId(v string) *GetServiceDetailResponseBodyDataServiceReports {
	s.AppointmentId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataServiceReports) SetBatchGroup(v string) *GetServiceDetailResponseBodyDataServiceReports {
	s.BatchGroup = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataServiceReports) SetCustomerId(v string) *GetServiceDetailResponseBodyDataServiceReports {
	s.CustomerId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataServiceReports) SetFileName(v string) *GetServiceDetailResponseBodyDataServiceReports {
	s.FileName = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataServiceReports) SetFileType(v int32) *GetServiceDetailResponseBodyDataServiceReports {
	s.FileType = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataServiceReports) SetGmtCreate(v string) *GetServiceDetailResponseBodyDataServiceReports {
	s.GmtCreate = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataServiceReports) SetGmtModified(v string) *GetServiceDetailResponseBodyDataServiceReports {
	s.GmtModified = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataServiceReports) SetId(v int64) *GetServiceDetailResponseBodyDataServiceReports {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataServiceReports) SetRemarke(v string) *GetServiceDetailResponseBodyDataServiceReports {
	s.Remarke = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataServiceReports) SetServiceApplyId(v int64) *GetServiceDetailResponseBodyDataServiceReports {
	s.ServiceApplyId = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataServiceReports) SetStatus(v int32) *GetServiceDetailResponseBodyDataServiceReports {
	s.Status = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataServiceReports) SetUrl(v string) *GetServiceDetailResponseBodyDataServiceReports {
	s.Url = &v
	return s
}

type GetServiceDetailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceDetailResponse) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponse) SetHeaders(v map[string]*string) *GetServiceDetailResponse {
	s.Headers = v
	return s
}

func (s *GetServiceDetailResponse) SetStatusCode(v int32) *GetServiceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceDetailResponse) SetBody(v *GetServiceDetailResponseBody) *GetServiceDetailResponse {
	s.Body = v
	return s
}

type GetYunQiTaskByRecordIdRequest struct {
	RecordId *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
}

func (s GetYunQiTaskByRecordIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetYunQiTaskByRecordIdRequest) GoString() string {
	return s.String()
}

func (s *GetYunQiTaskByRecordIdRequest) SetRecordId(v string) *GetYunQiTaskByRecordIdRequest {
	s.RecordId = &v
	return s
}

type GetYunQiTaskByRecordIdResponseBody struct {
	Code           *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data           *GetYunQiTaskByRecordIdResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetYunQiTaskByRecordIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetYunQiTaskByRecordIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetYunQiTaskByRecordIdResponseBody) SetCode(v string) *GetYunQiTaskByRecordIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBody) SetData(v *GetYunQiTaskByRecordIdResponseBodyData) *GetYunQiTaskByRecordIdResponseBody {
	s.Data = v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBody) SetHttpStatusCode(v int32) *GetYunQiTaskByRecordIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBody) SetMessage(v string) *GetYunQiTaskByRecordIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBody) SetRequestId(v string) *GetYunQiTaskByRecordIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBody) SetSuccess(v bool) *GetYunQiTaskByRecordIdResponseBody {
	s.Success = &v
	return s
}

type GetYunQiTaskByRecordIdResponseBodyData struct {
	ChatId               *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	CreateTime           *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorName          *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	EndTime              *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	EvaluationStar       *int32  `json:"evaluationStar,omitempty" xml:"evaluationStar,omitempty"`
	Important            *string `json:"important,omitempty" xml:"important,omitempty"`
	MainDingDepartmentId *string `json:"mainDingDepartmentId,omitempty" xml:"mainDingDepartmentId,omitempty"`
	MainDingGroupId      *string `json:"mainDingGroupId,omitempty" xml:"mainDingGroupId,omitempty"`
	MainDingGroupName    *string `json:"mainDingGroupName,omitempty" xml:"mainDingGroupName,omitempty"`
	ProductName          *string `json:"productName,omitempty" xml:"productName,omitempty"`
	RecordId             *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
	Status               *string `json:"status,omitempty" xml:"status,omitempty"`
	SubDingDepartmentId  *string `json:"subDingDepartmentId,omitempty" xml:"subDingDepartmentId,omitempty"`
	SubDingGroupId       *string `json:"subDingGroupId,omitempty" xml:"subDingGroupId,omitempty"`
	SubDingGroupName     *string `json:"subDingGroupName,omitempty" xml:"subDingGroupName,omitempty"`
	Title                *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetYunQiTaskByRecordIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetYunQiTaskByRecordIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetChatId(v string) *GetYunQiTaskByRecordIdResponseBodyData {
	s.ChatId = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetCreateTime(v int64) *GetYunQiTaskByRecordIdResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetCreatorName(v string) *GetYunQiTaskByRecordIdResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetEndTime(v int64) *GetYunQiTaskByRecordIdResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetEvaluationStar(v int32) *GetYunQiTaskByRecordIdResponseBodyData {
	s.EvaluationStar = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetImportant(v string) *GetYunQiTaskByRecordIdResponseBodyData {
	s.Important = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetMainDingDepartmentId(v string) *GetYunQiTaskByRecordIdResponseBodyData {
	s.MainDingDepartmentId = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetMainDingGroupId(v string) *GetYunQiTaskByRecordIdResponseBodyData {
	s.MainDingGroupId = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetMainDingGroupName(v string) *GetYunQiTaskByRecordIdResponseBodyData {
	s.MainDingGroupName = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetProductName(v string) *GetYunQiTaskByRecordIdResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetRecordId(v string) *GetYunQiTaskByRecordIdResponseBodyData {
	s.RecordId = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetStatus(v string) *GetYunQiTaskByRecordIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetSubDingDepartmentId(v string) *GetYunQiTaskByRecordIdResponseBodyData {
	s.SubDingDepartmentId = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetSubDingGroupId(v string) *GetYunQiTaskByRecordIdResponseBodyData {
	s.SubDingGroupId = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetSubDingGroupName(v string) *GetYunQiTaskByRecordIdResponseBodyData {
	s.SubDingGroupName = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponseBodyData) SetTitle(v string) *GetYunQiTaskByRecordIdResponseBodyData {
	s.Title = &v
	return s
}

type GetYunQiTaskByRecordIdResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetYunQiTaskByRecordIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYunQiTaskByRecordIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetYunQiTaskByRecordIdResponse) GoString() string {
	return s.String()
}

func (s *GetYunQiTaskByRecordIdResponse) SetHeaders(v map[string]*string) *GetYunQiTaskByRecordIdResponse {
	s.Headers = v
	return s
}

func (s *GetYunQiTaskByRecordIdResponse) SetStatusCode(v int32) *GetYunQiTaskByRecordIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetYunQiTaskByRecordIdResponse) SetBody(v *GetYunQiTaskByRecordIdResponseBody) *GetYunQiTaskByRecordIdResponse {
	s.Body = v
	return s
}

type ListDocsGroupByYearRequest struct {
	ApplyCodes      []*string `json:"applyCodes,omitempty" xml:"applyCodes,omitempty" type:"Repeated"`
	FileNameKeyword *string   `json:"fileNameKeyword,omitempty" xml:"fileNameKeyword,omitempty"`
	OrderIds        []*int64  `json:"orderIds,omitempty" xml:"orderIds,omitempty" type:"Repeated"`
	ProductCode     *string   `json:"productCode,omitempty" xml:"productCode,omitempty"`
	Scene           *string   `json:"scene,omitempty" xml:"scene,omitempty"`
}

func (s ListDocsGroupByYearRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDocsGroupByYearRequest) GoString() string {
	return s.String()
}

func (s *ListDocsGroupByYearRequest) SetApplyCodes(v []*string) *ListDocsGroupByYearRequest {
	s.ApplyCodes = v
	return s
}

func (s *ListDocsGroupByYearRequest) SetFileNameKeyword(v string) *ListDocsGroupByYearRequest {
	s.FileNameKeyword = &v
	return s
}

func (s *ListDocsGroupByYearRequest) SetOrderIds(v []*int64) *ListDocsGroupByYearRequest {
	s.OrderIds = v
	return s
}

func (s *ListDocsGroupByYearRequest) SetProductCode(v string) *ListDocsGroupByYearRequest {
	s.ProductCode = &v
	return s
}

func (s *ListDocsGroupByYearRequest) SetScene(v string) *ListDocsGroupByYearRequest {
	s.Scene = &v
	return s
}

type ListDocsGroupByYearResponseBody struct {
	Code           *string                 `json:"code,omitempty" xml:"code,omitempty"`
	Data           map[string][]*DataValue `json:"data,omitempty" xml:"data,omitempty"`
	HttpStatusCode *int32                  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                 `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListDocsGroupByYearResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDocsGroupByYearResponseBody) GoString() string {
	return s.String()
}

func (s *ListDocsGroupByYearResponseBody) SetCode(v string) *ListDocsGroupByYearResponseBody {
	s.Code = &v
	return s
}

func (s *ListDocsGroupByYearResponseBody) SetData(v map[string][]*DataValue) *ListDocsGroupByYearResponseBody {
	s.Data = v
	return s
}

func (s *ListDocsGroupByYearResponseBody) SetHttpStatusCode(v int32) *ListDocsGroupByYearResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDocsGroupByYearResponseBody) SetMessage(v string) *ListDocsGroupByYearResponseBody {
	s.Message = &v
	return s
}

func (s *ListDocsGroupByYearResponseBody) SetRequestId(v string) *ListDocsGroupByYearResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDocsGroupByYearResponseBody) SetSuccess(v bool) *ListDocsGroupByYearResponseBody {
	s.Success = &v
	return s
}

type ListDocsGroupByYearResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDocsGroupByYearResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDocsGroupByYearResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDocsGroupByYearResponse) GoString() string {
	return s.String()
}

func (s *ListDocsGroupByYearResponse) SetHeaders(v map[string]*string) *ListDocsGroupByYearResponse {
	s.Headers = v
	return s
}

func (s *ListDocsGroupByYearResponse) SetStatusCode(v int32) *ListDocsGroupByYearResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDocsGroupByYearResponse) SetBody(v *ListDocsGroupByYearResponseBody) *ListDocsGroupByYearResponse {
	s.Body = v
	return s
}

type ListEnterpriseSupportPlanRequest struct {
	PageNum  *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListEnterpriseSupportPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanRequest) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanRequest) SetPageNum(v int32) *ListEnterpriseSupportPlanRequest {
	s.PageNum = &v
	return s
}

func (s *ListEnterpriseSupportPlanRequest) SetPageSize(v int32) *ListEnterpriseSupportPlanRequest {
	s.PageSize = &v
	return s
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
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBody) GoString() string {
	return s.String()
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

type ListEnterpriseSupportPlanResponseBodyData struct {
	CanApplyFreeOrder        *bool                                                    `json:"canApplyFreeOrder,omitempty" xml:"canApplyFreeOrder,omitempty"`
	CustomerId               *int64                                                   `json:"customerId,omitempty" xml:"customerId,omitempty"`
	Docs                     []*ListEnterpriseSupportPlanResponseBodyDataDocs         `json:"docs,omitempty" xml:"docs,omitempty" type:"Repeated"`
	EndTime                  *string                                                  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	FirstPayTime             *string                                                  `json:"firstPayTime,omitempty" xml:"firstPayTime,omitempty"`
	FreeOrderApplyCode       *string                                                  `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	FreeOrderApplyId         *int64                                                   `json:"freeOrderApplyId,omitempty" xml:"freeOrderApplyId,omitempty"`
	FreeOrderApplyTime       *string                                                  `json:"freeOrderApplyTime,omitempty" xml:"freeOrderApplyTime,omitempty"`
	FreeOrderApprovedTime    *string                                                  `json:"freeOrderApprovedTime,omitempty" xml:"freeOrderApprovedTime,omitempty"`
	FreeOrderExpectStartTime *string                                                  `json:"freeOrderExpectStartTime,omitempty" xml:"freeOrderExpectStartTime,omitempty"`
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
	TaskNum                  *int64                                                   `json:"taskNum,omitempty" xml:"taskNum,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetCanApplyFreeOrder(v bool) *ListEnterpriseSupportPlanResponseBodyData {
	s.CanApplyFreeOrder = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyData) SetCustomerId(v int64) *ListEnterpriseSupportPlanResponseBodyData {
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

func (s *ListEnterpriseSupportPlanResponseBodyData) SetTaskNum(v int64) *ListEnterpriseSupportPlanResponseBodyData {
	s.TaskNum = &v
	return s
}

type ListEnterpriseSupportPlanResponseBodyDataDocs struct {
	DocId              *int64  `json:"docId,omitempty" xml:"docId,omitempty"`
	FileName           *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FreeOrderApplyCode *string `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	OrderId            *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	UploadTime         *string `json:"uploadTime,omitempty" xml:"uploadTime,omitempty"`
	Url                *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataDocs) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataDocs) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataDocs) SetDocId(v int64) *ListEnterpriseSupportPlanResponseBodyDataDocs {
	s.DocId = &v
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
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodes) GoString() string {
	return s.String()
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

type ListEnterpriseSupportPlanResponseBodyDataNodesDocNode struct {
	DocId              *int64  `json:"docId,omitempty" xml:"docId,omitempty"`
	DocName            *string `json:"docName,omitempty" xml:"docName,omitempty"`
	DocSubmitTime      *string `json:"docSubmitTime,omitempty" xml:"docSubmitTime,omitempty"`
	FileName           *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FreeOrderApplyCode *string `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	OrderId            *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesDocNode) GoString() string {
	return s.String()
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

type ListEnterpriseSupportPlanResponseBodyDataNodesFinishNode struct {
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesFinishNode) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesFinishNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFinishNode) SetFinishTime(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesFinishNode {
	s.FinishTime = &v
	return s
}

type ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode struct {
	AuditTime  *string `json:"auditTime,omitempty" xml:"auditTime,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	StatusName *string `json:"statusName,omitempty" xml:"statusName,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderAuditNode) GoString() string {
	return s.String()
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

type ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode struct {
	ApplyTime *string `json:"applyTime,omitempty" xml:"applyTime,omitempty"`
	Uid       *int64  `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode) SetApplyTime(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode {
	s.ApplyTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode) SetUid(v int64) *ListEnterpriseSupportPlanResponseBodyDataNodesFreeOrderNode {
	s.Uid = &v
	return s
}

type ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode struct {
	PayTime *string `json:"payTime,omitempty" xml:"payTime,omitempty"`
	Uid     *int64  `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode) SetPayTime(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode {
	s.PayTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode) SetUid(v int64) *ListEnterpriseSupportPlanResponseBodyDataNodesOrderNode {
	s.Uid = &v
	return s
}

type ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation struct {
	EndTime   *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation) SetEndTime(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation {
	s.EndTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation) SetStartTime(v string) *ListEnterpriseSupportPlanResponseBodyDataNodesServiceImplementation {
	s.StartTime = &v
	return s
}

type ListEnterpriseSupportPlanResponseBodyDataOperateInfos struct {
	CanClick *bool   `json:"canClick,omitempty" xml:"canClick,omitempty"`
	Text     *string `json:"text,omitempty" xml:"text,omitempty"`
	Url      *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListEnterpriseSupportPlanResponseBodyDataOperateInfos) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponseBodyDataOperateInfos) GoString() string {
	return s.String()
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

type ListEnterpriseSupportPlanResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnterpriseSupportPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnterpriseSupportPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponse) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponse) SetHeaders(v map[string]*string) *ListEnterpriseSupportPlanResponse {
	s.Headers = v
	return s
}

func (s *ListEnterpriseSupportPlanResponse) SetStatusCode(v int32) *ListEnterpriseSupportPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponse) SetBody(v *ListEnterpriseSupportPlanResponseBody) *ListEnterpriseSupportPlanResponse {
	s.Body = v
	return s
}

type ListEnterpriseSupportPlanSimpleRequest struct {
	PageNum  *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleRequest) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleRequest) SetPageNum(v int32) *ListEnterpriseSupportPlanSimpleRequest {
	s.PageNum = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleRequest) SetPageSize(v int32) *ListEnterpriseSupportPlanSimpleRequest {
	s.PageSize = &v
	return s
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
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBody) GoString() string {
	return s.String()
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

type ListEnterpriseSupportPlanSimpleResponseBodyData struct {
	CanApplyFreeOrder        *bool                                                   `json:"canApplyFreeOrder,omitempty" xml:"canApplyFreeOrder,omitempty"`
	CustomerId               *int64                                                  `json:"customerId,omitempty" xml:"customerId,omitempty"`
	Docs                     []*ListEnterpriseSupportPlanSimpleResponseBodyDataDocs  `json:"docs,omitempty" xml:"docs,omitempty" type:"Repeated"`
	EndTime                  *string                                                 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	FirstPayTime             *string                                                 `json:"firstPayTime,omitempty" xml:"firstPayTime,omitempty"`
	FreeOrderApplyCode       *string                                                 `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	FreeOrderApplyId         *int64                                                  `json:"freeOrderApplyId,omitempty" xml:"freeOrderApplyId,omitempty"`
	FreeOrderApplyTime       *string                                                 `json:"freeOrderApplyTime,omitempty" xml:"freeOrderApplyTime,omitempty"`
	FreeOrderApprovedTime    *string                                                 `json:"freeOrderApprovedTime,omitempty" xml:"freeOrderApprovedTime,omitempty"`
	FreeOrderExpectStartTime *string                                                 `json:"freeOrderExpectStartTime,omitempty" xml:"freeOrderExpectStartTime,omitempty"`
	InstanceId               *string                                                 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Nodes                    []*ListEnterpriseSupportPlanSimpleResponseBodyDataNodes `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	OrderIds                 []*int64                                                `json:"orderIds,omitempty" xml:"orderIds,omitempty" type:"Repeated"`
	ServiceName              *string                                                 `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	ServiceStatus            *string                                                 `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty"`
	ServiceStatusName        *string                                                 `json:"serviceStatusName,omitempty" xml:"serviceStatusName,omitempty"`
	ServiceType              *string                                                 `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	SortTime                 *string                                                 `json:"sortTime,omitempty" xml:"sortTime,omitempty"`
	StartTime                *string                                                 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TaskNum                  *int64                                                  `json:"taskNum,omitempty" xml:"taskNum,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetCanApplyFreeOrder(v bool) *ListEnterpriseSupportPlanSimpleResponseBodyData {
	s.CanApplyFreeOrder = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyData) SetCustomerId(v int64) *ListEnterpriseSupportPlanSimpleResponseBodyData {
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
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataDocs) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodes) GoString() string {
	return s.String()
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

type ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode struct {
	DocId              *int64  `json:"docId,omitempty" xml:"docId,omitempty"`
	DocName            *string `json:"docName,omitempty" xml:"docName,omitempty"`
	DocSubmitTime      *string `json:"docSubmitTime,omitempty" xml:"docSubmitTime,omitempty"`
	FileName           *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FreeOrderApplyCode *string `json:"freeOrderApplyCode,omitempty" xml:"freeOrderApplyCode,omitempty"`
	OrderId            *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesDocNode) GoString() string {
	return s.String()
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

type ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFinishNode struct {
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFinishNode) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFinishNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFinishNode) SetFinishTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFinishNode {
	s.FinishTime = &v
	return s
}

type ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode struct {
	AuditTime  *string `json:"auditTime,omitempty" xml:"auditTime,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	StatusName *string `json:"statusName,omitempty" xml:"statusName,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderAuditNode) GoString() string {
	return s.String()
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

type ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode struct {
	ApplyTime *string `json:"applyTime,omitempty" xml:"applyTime,omitempty"`
	Uid       *int64  `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode) SetApplyTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode {
	s.ApplyTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode) SetUid(v int64) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesFreeOrderNode {
	s.Uid = &v
	return s
}

type ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode struct {
	PayTime *string `json:"payTime,omitempty" xml:"payTime,omitempty"`
	Uid     *int64  `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode) SetPayTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode {
	s.PayTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode) SetUid(v int64) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesOrderNode {
	s.Uid = &v
	return s
}

type ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation struct {
	EndTime   *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation) SetEndTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation {
	s.EndTime = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation) SetStartTime(v string) *ListEnterpriseSupportPlanSimpleResponseBodyDataNodesServiceImplementation {
	s.StartTime = &v
	return s
}

type ListEnterpriseSupportPlanSimpleResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnterpriseSupportPlanSimpleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponse) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponse) SetHeaders(v map[string]*string) *ListEnterpriseSupportPlanSimpleResponse {
	s.Headers = v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponse) SetStatusCode(v int32) *ListEnterpriseSupportPlanSimpleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponse) SetBody(v *ListEnterpriseSupportPlanSimpleResponseBody) *ListEnterpriseSupportPlanSimpleResponse {
	s.Body = v
	return s
}

type ListServiceApplyRequest struct {
	ApplyType   []*string `json:"applyType,omitempty" xml:"applyType,omitempty" type:"Repeated"`
	EndDate     *int64    `json:"endDate,omitempty" xml:"endDate,omitempty"`
	PageNum     *int32    `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize    *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProductCode *int32    `json:"productCode,omitempty" xml:"productCode,omitempty"`
	StartDate   *int64    `json:"startDate,omitempty" xml:"startDate,omitempty"`
	Status      *string   `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListServiceApplyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyRequest) GoString() string {
	return s.String()
}

func (s *ListServiceApplyRequest) SetApplyType(v []*string) *ListServiceApplyRequest {
	s.ApplyType = v
	return s
}

func (s *ListServiceApplyRequest) SetEndDate(v int64) *ListServiceApplyRequest {
	s.EndDate = &v
	return s
}

func (s *ListServiceApplyRequest) SetPageNum(v int32) *ListServiceApplyRequest {
	s.PageNum = &v
	return s
}

func (s *ListServiceApplyRequest) SetPageSize(v int32) *ListServiceApplyRequest {
	s.PageSize = &v
	return s
}

func (s *ListServiceApplyRequest) SetProductCode(v int32) *ListServiceApplyRequest {
	s.ProductCode = &v
	return s
}

func (s *ListServiceApplyRequest) SetStartDate(v int64) *ListServiceApplyRequest {
	s.StartDate = &v
	return s
}

func (s *ListServiceApplyRequest) SetStatus(v string) *ListServiceApplyRequest {
	s.Status = &v
	return s
}

type ListServiceApplyResponseBody struct {
	Code           *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data           *ListServiceApplyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpStatusCode *int32                            `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                           `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListServiceApplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBody) SetCode(v string) *ListServiceApplyResponseBody {
	s.Code = &v
	return s
}

func (s *ListServiceApplyResponseBody) SetData(v *ListServiceApplyResponseBodyData) *ListServiceApplyResponseBody {
	s.Data = v
	return s
}

func (s *ListServiceApplyResponseBody) SetHttpStatusCode(v int32) *ListServiceApplyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListServiceApplyResponseBody) SetMessage(v string) *ListServiceApplyResponseBody {
	s.Message = &v
	return s
}

func (s *ListServiceApplyResponseBody) SetRequestId(v string) *ListServiceApplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceApplyResponseBody) SetSuccess(v bool) *ListServiceApplyResponseBody {
	s.Success = &v
	return s
}

type ListServiceApplyResponseBodyData struct {
	Extend   interface{}                             `json:"extend,omitempty" xml:"extend,omitempty"`
	List     []*ListServiceApplyResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListServiceApplyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyData) SetExtend(v interface{}) *ListServiceApplyResponseBodyData {
	s.Extend = v
	return s
}

func (s *ListServiceApplyResponseBodyData) SetList(v []*ListServiceApplyResponseBodyDataList) *ListServiceApplyResponseBodyData {
	s.List = v
	return s
}

func (s *ListServiceApplyResponseBodyData) SetPageNum(v int32) *ListServiceApplyResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListServiceApplyResponseBodyData) SetPageSize(v int32) *ListServiceApplyResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListServiceApplyResponseBodyData) SetTotal(v int32) *ListServiceApplyResponseBodyData {
	s.Total = &v
	return s
}

type ListServiceApplyResponseBodyDataList struct {
	ApplierId                       *string                                                  `json:"applierId,omitempty" xml:"applierId,omitempty"`
	ApplyCode                       *string                                                  `json:"applyCode,omitempty" xml:"applyCode,omitempty"`
	ApplyComponentDetails           [][]*string                                              `json:"applyComponentDetails,omitempty" xml:"applyComponentDetails,omitempty" type:"Repeated"`
	ApplyTime                       *int64                                                   `json:"applyTime,omitempty" xml:"applyTime,omitempty"`
	Appointments                    []*ListServiceApplyResponseBodyDataListAppointments      `json:"appointments,omitempty" xml:"appointments,omitempty" type:"Repeated"`
	BuyUrl                          *string                                                  `json:"buyUrl,omitempty" xml:"buyUrl,omitempty"`
	CreatorEmpId                    *string                                                  `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	CustomerName                    *string                                                  `json:"customerName,omitempty" xml:"customerName,omitempty"`
	CycleService                    *bool                                                    `json:"cycleService,omitempty" xml:"cycleService,omitempty"`
	ExecutedCount                   *int64                                                   `json:"executedCount,omitempty" xml:"executedCount,omitempty"`
	FinishCount                     *int64                                                   `json:"finishCount,omitempty" xml:"finishCount,omitempty"`
	GmtCreate                       *string                                                  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified                     *string                                                  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                              *int64                                                   `json:"id,omitempty" xml:"id,omitempty"`
	MergeSolutionAndReporterOneStep *bool                                                    `json:"mergeSolutionAndReporterOneStep,omitempty" xml:"mergeSolutionAndReporterOneStep,omitempty"`
	ModifierEmpId                   *string                                                  `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	PackDetails                     []map[string]interface{}                                 `json:"packDetails,omitempty" xml:"packDetails,omitempty" type:"Repeated"`
	PayOrders                       []*ListServiceApplyResponseBodyDataListPayOrders         `json:"payOrders,omitempty" xml:"payOrders,omitempty" type:"Repeated"`
	PayUrl                          *string                                                  `json:"payUrl,omitempty" xml:"payUrl,omitempty"`
	PerformanceOrders               []*ListServiceApplyResponseBodyDataListPerformanceOrders `json:"performanceOrders,omitempty" xml:"performanceOrders,omitempty" type:"Repeated"`
	PerformancePacks                []*ListServiceApplyResponseBodyDataListPerformancePacks  `json:"performancePacks,omitempty" xml:"performancePacks,omitempty" type:"Repeated"`
	ReneWalUrl                      *string                                                  `json:"reneWalUrl,omitempty" xml:"reneWalUrl,omitempty"`
	ServiceCode                     *string                                                  `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	ServiceName                     *string                                                  `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	ServiceReports                  []*ListServiceApplyResponseBodyDataListServiceReports    `json:"serviceReports,omitempty" xml:"serviceReports,omitempty" type:"Repeated"`
	ServiceTimeRange                []*int64                                                 `json:"serviceTimeRange,omitempty" xml:"serviceTimeRange,omitempty" type:"Repeated"`
	Status                          *string                                                  `json:"status,omitempty" xml:"status,omitempty"`
	StatusCode                      *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	StatusStr                       *string                                                  `json:"statusStr,omitempty" xml:"statusStr,omitempty"`
	TermOfValidity                  *string                                                  `json:"termOfValidity,omitempty" xml:"termOfValidity,omitempty"`
	TotalPack                       *int32                                                   `json:"totalPack,omitempty" xml:"totalPack,omitempty"`
	Type                            *string                                                  `json:"type,omitempty" xml:"type,omitempty"`
	UsePack                         *int64                                                   `json:"usePack,omitempty" xml:"usePack,omitempty"`
}

func (s ListServiceApplyResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataList) SetApplierId(v string) *ListServiceApplyResponseBodyDataList {
	s.ApplierId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetApplyCode(v string) *ListServiceApplyResponseBodyDataList {
	s.ApplyCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetApplyComponentDetails(v [][]*string) *ListServiceApplyResponseBodyDataList {
	s.ApplyComponentDetails = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetApplyTime(v int64) *ListServiceApplyResponseBodyDataList {
	s.ApplyTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetAppointments(v []*ListServiceApplyResponseBodyDataListAppointments) *ListServiceApplyResponseBodyDataList {
	s.Appointments = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetBuyUrl(v string) *ListServiceApplyResponseBodyDataList {
	s.BuyUrl = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataList {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetCustomerName(v string) *ListServiceApplyResponseBodyDataList {
	s.CustomerName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetCycleService(v bool) *ListServiceApplyResponseBodyDataList {
	s.CycleService = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetExecutedCount(v int64) *ListServiceApplyResponseBodyDataList {
	s.ExecutedCount = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetFinishCount(v int64) *ListServiceApplyResponseBodyDataList {
	s.FinishCount = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetGmtModified(v string) *ListServiceApplyResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetId(v int64) *ListServiceApplyResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetMergeSolutionAndReporterOneStep(v bool) *ListServiceApplyResponseBodyDataList {
	s.MergeSolutionAndReporterOneStep = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataList {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetPackDetails(v []map[string]interface{}) *ListServiceApplyResponseBodyDataList {
	s.PackDetails = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetPayOrders(v []*ListServiceApplyResponseBodyDataListPayOrders) *ListServiceApplyResponseBodyDataList {
	s.PayOrders = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetPayUrl(v string) *ListServiceApplyResponseBodyDataList {
	s.PayUrl = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetPerformanceOrders(v []*ListServiceApplyResponseBodyDataListPerformanceOrders) *ListServiceApplyResponseBodyDataList {
	s.PerformanceOrders = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetPerformancePacks(v []*ListServiceApplyResponseBodyDataListPerformancePacks) *ListServiceApplyResponseBodyDataList {
	s.PerformancePacks = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetReneWalUrl(v string) *ListServiceApplyResponseBodyDataList {
	s.ReneWalUrl = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetServiceCode(v string) *ListServiceApplyResponseBodyDataList {
	s.ServiceCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetServiceName(v string) *ListServiceApplyResponseBodyDataList {
	s.ServiceName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetServiceReports(v []*ListServiceApplyResponseBodyDataListServiceReports) *ListServiceApplyResponseBodyDataList {
	s.ServiceReports = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetServiceTimeRange(v []*int64) *ListServiceApplyResponseBodyDataList {
	s.ServiceTimeRange = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetStatus(v string) *ListServiceApplyResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetStatusCode(v int32) *ListServiceApplyResponseBodyDataList {
	s.StatusCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetStatusStr(v string) *ListServiceApplyResponseBodyDataList {
	s.StatusStr = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetTermOfValidity(v string) *ListServiceApplyResponseBodyDataList {
	s.TermOfValidity = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetTotalPack(v int32) *ListServiceApplyResponseBodyDataList {
	s.TotalPack = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetType(v string) *ListServiceApplyResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetUsePack(v int64) *ListServiceApplyResponseBodyDataList {
	s.UsePack = &v
	return s
}

type ListServiceApplyResponseBodyDataListAppointments struct {
	HuhangId     *int64  `json:"huhangId,omitempty" xml:"huhangId,omitempty"`
	PurchaseCode *int32  `json:"purchaseCode,omitempty" xml:"purchaseCode,omitempty"`
	PurchaseDesc *string `json:"purchaseDesc,omitempty" xml:"purchaseDesc,omitempty"`
	SupportDays  *int32  `json:"supportDays,omitempty" xml:"supportDays,omitempty"`
	TravelDays   *int32  `json:"travelDays,omitempty" xml:"travelDays,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListAppointments) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListAppointments) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListAppointments) SetHuhangId(v int64) *ListServiceApplyResponseBodyDataListAppointments {
	s.HuhangId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListAppointments) SetPurchaseCode(v int32) *ListServiceApplyResponseBodyDataListAppointments {
	s.PurchaseCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListAppointments) SetPurchaseDesc(v string) *ListServiceApplyResponseBodyDataListAppointments {
	s.PurchaseDesc = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListAppointments) SetSupportDays(v int32) *ListServiceApplyResponseBodyDataListAppointments {
	s.SupportDays = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListAppointments) SetTravelDays(v int32) *ListServiceApplyResponseBodyDataListAppointments {
	s.TravelDays = &v
	return s
}

type ListServiceApplyResponseBodyDataListPayOrders struct {
	Amount               *string                `json:"amount,omitempty" xml:"amount,omitempty"`
	CompassCommodityCode *string                `json:"compassCommodityCode,omitempty" xml:"compassCommodityCode,omitempty"`
	CompassCommodityName *string                `json:"compassCommodityName,omitempty" xml:"compassCommodityName,omitempty"`
	CreatorEmpId         *string                `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	GmtCreate            *string                `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified          *string                `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                   *int64                 `json:"id,omitempty" xml:"id,omitempty"`
	ModifierEmpId        *string                `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	Operate              map[string]interface{} `json:"operate,omitempty" xml:"operate,omitempty"`
	OrderDetail          interface{}            `json:"orderDetail,omitempty" xml:"orderDetail,omitempty"`
	OrderId              *int64                 `json:"orderId,omitempty" xml:"orderId,omitempty"`
	OriginalPrice        *float64               `json:"originalPrice,omitempty" xml:"originalPrice,omitempty"`
	PayAmount            *float64               `json:"payAmount,omitempty" xml:"payAmount,omitempty"`
	PayTime              *string                `json:"payTime,omitempty" xml:"payTime,omitempty"`
	ProductName          *string                `json:"productName,omitempty" xml:"productName,omitempty"`
	ReneWalUrl           *string                `json:"reneWalUrl,omitempty" xml:"reneWalUrl,omitempty"`
	ServiceContentMap    map[string]interface{} `json:"serviceContentMap,omitempty" xml:"serviceContentMap,omitempty"`
	Status               *int32                 `json:"status,omitempty" xml:"status,omitempty"`
	StatusStr            *string                `json:"statusStr,omitempty" xml:"statusStr,omitempty"`
	SupportDays          *int32                 `json:"supportDays,omitempty" xml:"supportDays,omitempty"`
	Uid                  *string                `json:"uid,omitempty" xml:"uid,omitempty"`
	Url                  *string                `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPayOrders) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPayOrders) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetAmount(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.Amount = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetCompassCommodityCode(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.CompassCommodityCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetCompassCommodityName(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.CompassCommodityName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetId(v int64) *ListServiceApplyResponseBodyDataListPayOrders {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetOperate(v map[string]interface{}) *ListServiceApplyResponseBodyDataListPayOrders {
	s.Operate = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetOrderDetail(v interface{}) *ListServiceApplyResponseBodyDataListPayOrders {
	s.OrderDetail = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetOrderId(v int64) *ListServiceApplyResponseBodyDataListPayOrders {
	s.OrderId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetOriginalPrice(v float64) *ListServiceApplyResponseBodyDataListPayOrders {
	s.OriginalPrice = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetPayAmount(v float64) *ListServiceApplyResponseBodyDataListPayOrders {
	s.PayAmount = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetPayTime(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.PayTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetProductName(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.ProductName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetReneWalUrl(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.ReneWalUrl = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetServiceContentMap(v map[string]interface{}) *ListServiceApplyResponseBodyDataListPayOrders {
	s.ServiceContentMap = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPayOrders {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetStatusStr(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.StatusStr = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetSupportDays(v int32) *ListServiceApplyResponseBodyDataListPayOrders {
	s.SupportDays = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetUid(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.Uid = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetUrl(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.Url = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrders struct {
	ApplyFileVOList                 []*ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList     `json:"applyFileVOList,omitempty" xml:"applyFileVOList,omitempty" type:"Repeated"`
	AppointmentCode                 *string                                                                     `json:"appointmentCode,omitempty" xml:"appointmentCode,omitempty"`
	AppointmentEndTime              *int64                                                                      `json:"appointmentEndTime,omitempty" xml:"appointmentEndTime,omitempty"`
	AppointmentId                   *string                                                                     `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	AppointmentPassTime             *int64                                                                      `json:"appointmentPassTime,omitempty" xml:"appointmentPassTime,omitempty"`
	AppointmentTime                 *int64                                                                      `json:"appointmentTime,omitempty" xml:"appointmentTime,omitempty"`
	CommodityDesc                   *string                                                                     `json:"commodityDesc,omitempty" xml:"commodityDesc,omitempty"`
	CreatorEmpId                    *string                                                                     `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	CycleService                    *bool                                                                       `json:"cycleService,omitempty" xml:"cycleService,omitempty"`
	ExtList                         []*ListServiceApplyResponseBodyDataListPerformanceOrdersExtList             `json:"extList,omitempty" xml:"extList,omitempty" type:"Repeated"`
	GmtCreate                       *string                                                                     `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified                     *string                                                                     `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                              *int64                                                                      `json:"id,omitempty" xml:"id,omitempty"`
	MergeSolutionAndReporterOneStep *bool                                                                       `json:"mergeSolutionAndReporterOneStep,omitempty" xml:"mergeSolutionAndReporterOneStep,omitempty"`
	ModifierEmpId                   *string                                                                     `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	NtmCommodityCode                *string                                                                     `json:"ntmCommodityCode,omitempty" xml:"ntmCommodityCode,omitempty"`
	OrderDetail                     interface{}                                                                 `json:"orderDetail,omitempty" xml:"orderDetail,omitempty"`
	OrderId                         *int64                                                                      `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PackCount                       *int32                                                                      `json:"packCount,omitempty" xml:"packCount,omitempty"`
	PackDetails                     []map[string]interface{}                                                    `json:"packDetails,omitempty" xml:"packDetails,omitempty" type:"Repeated"`
	PerformanceNodeDTOS             []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS `json:"performanceNodeDTOS,omitempty" xml:"performanceNodeDTOS,omitempty" type:"Repeated"`
	PerformancePacks                []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks    `json:"performancePacks,omitempty" xml:"performancePacks,omitempty" type:"Repeated"`
	PurchasePackCode                *int32                                                                      `json:"purchasePackCode,omitempty" xml:"purchasePackCode,omitempty"`
	ServiceApplyId                  *int64                                                                      `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	ServiceMonthReports             []*ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports `json:"serviceMonthReports,omitempty" xml:"serviceMonthReports,omitempty" type:"Repeated"`
	ServiceReports                  []*ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports      `json:"serviceReports,omitempty" xml:"serviceReports,omitempty" type:"Repeated"`
	ServiceSchemes                  []*ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes      `json:"serviceSchemes,omitempty" xml:"serviceSchemes,omitempty" type:"Repeated"`
	Status                          *int32                                                                      `json:"status,omitempty" xml:"status,omitempty"`
	StatusStr                       *string                                                                     `json:"statusStr,omitempty" xml:"statusStr,omitempty"`
	SupportTime                     []*int64                                                                    `json:"supportTime,omitempty" xml:"supportTime,omitempty" type:"Repeated"`
	TamEngineers                    []*ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers        `json:"tamEngineers,omitempty" xml:"tamEngineers,omitempty" type:"Repeated"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrders) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrders) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetApplyFileVOList(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.ApplyFileVOList = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetAppointmentCode(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.AppointmentCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetAppointmentEndTime(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.AppointmentEndTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetAppointmentPassTime(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.AppointmentPassTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetAppointmentTime(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.AppointmentTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetCommodityDesc(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.CommodityDesc = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetCycleService(v bool) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.CycleService = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetExtList(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.ExtList = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetMergeSolutionAndReporterOneStep(v bool) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.MergeSolutionAndReporterOneStep = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetNtmCommodityCode(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.NtmCommodityCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetOrderDetail(v interface{}) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.OrderDetail = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetOrderId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.OrderId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetPackCount(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.PackCount = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetPackDetails(v []map[string]interface{}) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.PackDetails = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetPerformanceNodeDTOS(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.PerformanceNodeDTOS = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetPerformancePacks(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.PerformancePacks = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetPurchasePackCode(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.PurchasePackCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetServiceMonthReports(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.ServiceMonthReports = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetServiceReports(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.ServiceReports = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetServiceSchemes(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.ServiceSchemes = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetStatusStr(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.StatusStr = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetSupportTime(v []*int64) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.SupportTime = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetTamEngineers(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.TamEngineers = v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.Url = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersExtList struct {
	KeyCode *string     `json:"keyCode,omitempty" xml:"keyCode,omitempty"`
	Name    *string     `json:"name,omitempty" xml:"name,omitempty"`
	Value   interface{} `json:"value,omitempty" xml:"value,omitempty"`
	View    *string     `json:"view,omitempty" xml:"view,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) SetKeyCode(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList {
	s.KeyCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) SetName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList {
	s.Name = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) SetValue(v interface{}) *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList {
	s.Value = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) SetView(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList {
	s.View = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS struct {
	Display    *bool       `json:"display,omitempty" xml:"display,omitempty"`
	ExtendInfo interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Index      *int32      `json:"index,omitempty" xml:"index,omitempty"`
	NodeName   *string     `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	Status     *int32      `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) SetDisplay(v bool) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS {
	s.Display = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) SetExtendInfo(v interface{}) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS {
	s.ExtendInfo = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) SetIndex(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS {
	s.Index = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) SetNodeName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS {
	s.NodeName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS {
	s.Status = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks struct {
	ApplyFileVOList                 []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList     `json:"applyFileVOList,omitempty" xml:"applyFileVOList,omitempty" type:"Repeated"`
	AppointmentCode                 *string                                                                                     `json:"appointmentCode,omitempty" xml:"appointmentCode,omitempty"`
	AppointmentEndTime              *int64                                                                                      `json:"appointmentEndTime,omitempty" xml:"appointmentEndTime,omitempty"`
	AppointmentId                   *string                                                                                     `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	AppointmentPassTime             *int64                                                                                      `json:"appointmentPassTime,omitempty" xml:"appointmentPassTime,omitempty"`
	AppointmentTime                 *int64                                                                                      `json:"appointmentTime,omitempty" xml:"appointmentTime,omitempty"`
	CommodityDesc                   *string                                                                                     `json:"commodityDesc,omitempty" xml:"commodityDesc,omitempty"`
	CreatorEmpId                    *string                                                                                     `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	CycleService                    *bool                                                                                       `json:"cycleService,omitempty" xml:"cycleService,omitempty"`
	ExtList                         []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList             `json:"extList,omitempty" xml:"extList,omitempty" type:"Repeated"`
	GmtCreate                       *string                                                                                     `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified                     *string                                                                                     `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                              *int64                                                                                      `json:"id,omitempty" xml:"id,omitempty"`
	MergeSolutionAndReporterOneStep *bool                                                                                       `json:"mergeSolutionAndReporterOneStep,omitempty" xml:"mergeSolutionAndReporterOneStep,omitempty"`
	ModifierEmpId                   *string                                                                                     `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	NtmCommodityCode                *string                                                                                     `json:"ntmCommodityCode,omitempty" xml:"ntmCommodityCode,omitempty"`
	OrderDetail                     interface{}                                                                                 `json:"orderDetail,omitempty" xml:"orderDetail,omitempty"`
	OrderId                         *int64                                                                                      `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PerformanceNodeDTOS             []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS `json:"performanceNodeDTOS,omitempty" xml:"performanceNodeDTOS,omitempty" type:"Repeated"`
	PurchasePackCode                *int32                                                                                      `json:"purchasePackCode,omitempty" xml:"purchasePackCode,omitempty"`
	ServiceApplyId                  *int64                                                                                      `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	ServiceMonthReports             []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports `json:"serviceMonthReports,omitempty" xml:"serviceMonthReports,omitempty" type:"Repeated"`
	ServiceReports                  []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports      `json:"serviceReports,omitempty" xml:"serviceReports,omitempty" type:"Repeated"`
	ServiceSchemes                  []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes      `json:"serviceSchemes,omitempty" xml:"serviceSchemes,omitempty" type:"Repeated"`
	Status                          *int32                                                                                      `json:"status,omitempty" xml:"status,omitempty"`
	StatusStr                       *string                                                                                     `json:"statusStr,omitempty" xml:"statusStr,omitempty"`
	SupportTime                     []*int64                                                                                    `json:"supportTime,omitempty" xml:"supportTime,omitempty" type:"Repeated"`
	TamEngineers                    []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers        `json:"tamEngineers,omitempty" xml:"tamEngineers,omitempty" type:"Repeated"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetApplyFileVOList(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.ApplyFileVOList = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetAppointmentCode(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.AppointmentCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetAppointmentEndTime(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.AppointmentEndTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetAppointmentPassTime(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.AppointmentPassTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetAppointmentTime(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.AppointmentTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetCommodityDesc(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.CommodityDesc = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetCycleService(v bool) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.CycleService = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetExtList(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.ExtList = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetMergeSolutionAndReporterOneStep(v bool) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.MergeSolutionAndReporterOneStep = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetNtmCommodityCode(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.NtmCommodityCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetOrderDetail(v interface{}) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.OrderDetail = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetOrderId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.OrderId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetPerformanceNodeDTOS(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.PerformanceNodeDTOS = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetPurchasePackCode(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.PurchasePackCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetServiceMonthReports(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.ServiceMonthReports = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetServiceReports(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.ServiceReports = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetServiceSchemes(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.ServiceSchemes = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetStatusStr(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.StatusStr = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetSupportTime(v []*int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.SupportTime = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetTamEngineers(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.TamEngineers = v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.Url = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList struct {
	KeyCode *string     `json:"keyCode,omitempty" xml:"keyCode,omitempty"`
	Name    *string     `json:"name,omitempty" xml:"name,omitempty"`
	Value   interface{} `json:"value,omitempty" xml:"value,omitempty"`
	View    *string     `json:"view,omitempty" xml:"view,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) SetKeyCode(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList {
	s.KeyCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) SetName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList {
	s.Name = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) SetValue(v interface{}) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList {
	s.Value = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) SetView(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList {
	s.View = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS struct {
	Display    *bool       `json:"display,omitempty" xml:"display,omitempty"`
	ExtendInfo interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Index      *int32      `json:"index,omitempty" xml:"index,omitempty"`
	NodeName   *string     `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	Status     *int32      `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) SetDisplay(v bool) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	s.Display = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) SetExtendInfo(v interface{}) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	s.ExtendInfo = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) SetIndex(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	s.Index = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) SetNodeName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	s.NodeName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	s.Status = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.Url = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.Url = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.Url = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers struct {
	CreatorEmpId  *string `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	GmtCreate     *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified   *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	HrStatus      *string `json:"hrStatus,omitempty" xml:"hrStatus,omitempty"`
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LastName      *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
	ModifierEmpId *string `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	NickNameEn    *string `json:"nickNameEn,omitempty" xml:"nickNameEn,omitempty"`
	RealmId       *int64  `json:"realmId,omitempty" xml:"realmId,omitempty"`
	Workid        *string `json:"workid,omitempty" xml:"workid,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetHrStatus(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.HrStatus = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetLastName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.LastName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.Name = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetNickNameEn(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.NickNameEn = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetRealmId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.RealmId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetWorkid(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.Workid = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.Url = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.Url = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.Url = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers struct {
	CreatorEmpId  *string `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	GmtCreate     *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified   *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	HrStatus      *string `json:"hrStatus,omitempty" xml:"hrStatus,omitempty"`
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LastName      *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
	ModifierEmpId *string `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	NickNameEn    *string `json:"nickNameEn,omitempty" xml:"nickNameEn,omitempty"`
	RealmId       *int64  `json:"realmId,omitempty" xml:"realmId,omitempty"`
	Workid        *string `json:"workid,omitempty" xml:"workid,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetHrStatus(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.HrStatus = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetLastName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.LastName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.Name = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetNickNameEn(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.NickNameEn = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetRealmId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.RealmId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetWorkid(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.Workid = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformancePacks struct {
	ApplyFileVOList                 []*ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList     `json:"applyFileVOList,omitempty" xml:"applyFileVOList,omitempty" type:"Repeated"`
	AppointmentCode                 *string                                                                    `json:"appointmentCode,omitempty" xml:"appointmentCode,omitempty"`
	AppointmentEndTime              *int64                                                                     `json:"appointmentEndTime,omitempty" xml:"appointmentEndTime,omitempty"`
	AppointmentId                   *string                                                                    `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	AppointmentPassTime             *int64                                                                     `json:"appointmentPassTime,omitempty" xml:"appointmentPassTime,omitempty"`
	AppointmentTime                 *int64                                                                     `json:"appointmentTime,omitempty" xml:"appointmentTime,omitempty"`
	CommodityDesc                   *string                                                                    `json:"commodityDesc,omitempty" xml:"commodityDesc,omitempty"`
	CreatorEmpId                    *string                                                                    `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	CycleService                    *bool                                                                      `json:"cycleService,omitempty" xml:"cycleService,omitempty"`
	ExtList                         []*ListServiceApplyResponseBodyDataListPerformancePacksExtList             `json:"extList,omitempty" xml:"extList,omitempty" type:"Repeated"`
	GmtCreate                       *string                                                                    `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified                     *string                                                                    `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                              *int64                                                                     `json:"id,omitempty" xml:"id,omitempty"`
	MergeSolutionAndReporterOneStep *bool                                                                      `json:"mergeSolutionAndReporterOneStep,omitempty" xml:"mergeSolutionAndReporterOneStep,omitempty"`
	ModifierEmpId                   *string                                                                    `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	NtmCommodityCode                *string                                                                    `json:"ntmCommodityCode,omitempty" xml:"ntmCommodityCode,omitempty"`
	OrderDetail                     interface{}                                                                `json:"orderDetail,omitempty" xml:"orderDetail,omitempty"`
	OrderId                         *int64                                                                     `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PerformanceNodeDTOS             []*ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS `json:"performanceNodeDTOS,omitempty" xml:"performanceNodeDTOS,omitempty" type:"Repeated"`
	PurchasePackCode                *int32                                                                     `json:"purchasePackCode,omitempty" xml:"purchasePackCode,omitempty"`
	ServiceApplyId                  *int64                                                                     `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	ServiceMonthReports             []*ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports `json:"serviceMonthReports,omitempty" xml:"serviceMonthReports,omitempty" type:"Repeated"`
	ServiceReports                  []*ListServiceApplyResponseBodyDataListPerformancePacksServiceReports      `json:"serviceReports,omitempty" xml:"serviceReports,omitempty" type:"Repeated"`
	ServiceSchemes                  []*ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes      `json:"serviceSchemes,omitempty" xml:"serviceSchemes,omitempty" type:"Repeated"`
	Status                          *int32                                                                     `json:"status,omitempty" xml:"status,omitempty"`
	StatusStr                       *string                                                                    `json:"statusStr,omitempty" xml:"statusStr,omitempty"`
	SupportTime                     []*int64                                                                   `json:"supportTime,omitempty" xml:"supportTime,omitempty" type:"Repeated"`
	TamEngineers                    []*ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers        `json:"tamEngineers,omitempty" xml:"tamEngineers,omitempty" type:"Repeated"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacks) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacks) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetApplyFileVOList(v []*ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.ApplyFileVOList = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetAppointmentCode(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.AppointmentCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetAppointmentEndTime(v int64) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.AppointmentEndTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetAppointmentPassTime(v int64) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.AppointmentPassTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetAppointmentTime(v int64) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.AppointmentTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetCommodityDesc(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.CommodityDesc = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetCycleService(v bool) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.CycleService = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetExtList(v []*ListServiceApplyResponseBodyDataListPerformancePacksExtList) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.ExtList = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetMergeSolutionAndReporterOneStep(v bool) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.MergeSolutionAndReporterOneStep = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetNtmCommodityCode(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.NtmCommodityCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetOrderDetail(v interface{}) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.OrderDetail = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetOrderId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.OrderId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetPerformanceNodeDTOS(v []*ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.PerformanceNodeDTOS = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetPurchasePackCode(v int32) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.PurchasePackCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetServiceMonthReports(v []*ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.ServiceMonthReports = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetServiceReports(v []*ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.ServiceReports = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetServiceSchemes(v []*ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.ServiceSchemes = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetStatusStr(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.StatusStr = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetSupportTime(v []*int64) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.SupportTime = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetTamEngineers(v []*ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.TamEngineers = v
	return s
}

type ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.Url = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformancePacksExtList struct {
	KeyCode *string     `json:"keyCode,omitempty" xml:"keyCode,omitempty"`
	Name    *string     `json:"name,omitempty" xml:"name,omitempty"`
	Value   interface{} `json:"value,omitempty" xml:"value,omitempty"`
	View    *string     `json:"view,omitempty" xml:"view,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksExtList) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksExtList) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksExtList) SetKeyCode(v string) *ListServiceApplyResponseBodyDataListPerformancePacksExtList {
	s.KeyCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksExtList) SetName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksExtList {
	s.Name = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksExtList) SetValue(v interface{}) *ListServiceApplyResponseBodyDataListPerformancePacksExtList {
	s.Value = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksExtList) SetView(v string) *ListServiceApplyResponseBodyDataListPerformancePacksExtList {
	s.View = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS struct {
	Display    *bool       `json:"display,omitempty" xml:"display,omitempty"`
	ExtendInfo interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Index      *int32      `json:"index,omitempty" xml:"index,omitempty"`
	NodeName   *string     `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	Status     *int32      `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) SetDisplay(v bool) *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS {
	s.Display = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) SetExtendInfo(v interface{}) *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS {
	s.ExtendInfo = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) SetIndex(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS {
	s.Index = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) SetNodeName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS {
	s.NodeName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS {
	s.Status = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.Url = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformancePacksServiceReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.Url = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.Url = &v
	return s
}

type ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers struct {
	CreatorEmpId  *string `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	GmtCreate     *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified   *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	HrStatus      *string `json:"hrStatus,omitempty" xml:"hrStatus,omitempty"`
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LastName      *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
	ModifierEmpId *string `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	NickNameEn    *string `json:"nickNameEn,omitempty" xml:"nickNameEn,omitempty"`
	RealmId       *int64  `json:"realmId,omitempty" xml:"realmId,omitempty"`
	Workid        *string `json:"workid,omitempty" xml:"workid,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetHrStatus(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.HrStatus = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetLastName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.LastName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.Name = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetNickNameEn(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.NickNameEn = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetRealmId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.RealmId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetWorkid(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.Workid = &v
	return s
}

type ListServiceApplyResponseBodyDataListServiceReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListServiceReports) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListServiceReports) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetFileName(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetFileType(v int32) *ListServiceApplyResponseBodyDataListServiceReports {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetId(v int64) *ListServiceApplyResponseBodyDataListServiceReports {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetRemarke(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListServiceReports {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetStatus(v int32) *ListServiceApplyResponseBodyDataListServiceReports {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetUrl(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.Url = &v
	return s
}

type ListServiceApplyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceApplyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceApplyResponse) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponse) SetHeaders(v map[string]*string) *ListServiceApplyResponse {
	s.Headers = v
	return s
}

func (s *ListServiceApplyResponse) SetStatusCode(v int32) *ListServiceApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceApplyResponse) SetBody(v *ListServiceApplyResponseBody) *ListServiceApplyResponse {
	s.Body = v
	return s
}

type ListYunQiTaskByUidRequest struct {
	CreateTimeEnd       *int64    `json:"createTimeEnd,omitempty" xml:"createTimeEnd,omitempty"`
	CreateTimeStart     *int64    `json:"createTimeStart,omitempty" xml:"createTimeStart,omitempty"`
	FreeOrderApplyCodes []*string `json:"freeOrderApplyCodes,omitempty" xml:"freeOrderApplyCodes,omitempty" type:"Repeated"`
	FreeOrderApplyIds   []*int64  `json:"freeOrderApplyIds,omitempty" xml:"freeOrderApplyIds,omitempty" type:"Repeated"`
	OrderIds            []*int64  `json:"orderIds,omitempty" xml:"orderIds,omitempty" type:"Repeated"`
	PageNum             *int32    `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize            *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StatusList          []*string `json:"statusList,omitempty" xml:"statusList,omitempty" type:"Repeated"`
}

func (s ListYunQiTaskByUidRequest) String() string {
	return tea.Prettify(s)
}

func (s ListYunQiTaskByUidRequest) GoString() string {
	return s.String()
}

func (s *ListYunQiTaskByUidRequest) SetCreateTimeEnd(v int64) *ListYunQiTaskByUidRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListYunQiTaskByUidRequest) SetCreateTimeStart(v int64) *ListYunQiTaskByUidRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListYunQiTaskByUidRequest) SetFreeOrderApplyCodes(v []*string) *ListYunQiTaskByUidRequest {
	s.FreeOrderApplyCodes = v
	return s
}

func (s *ListYunQiTaskByUidRequest) SetFreeOrderApplyIds(v []*int64) *ListYunQiTaskByUidRequest {
	s.FreeOrderApplyIds = v
	return s
}

func (s *ListYunQiTaskByUidRequest) SetOrderIds(v []*int64) *ListYunQiTaskByUidRequest {
	s.OrderIds = v
	return s
}

func (s *ListYunQiTaskByUidRequest) SetPageNum(v int32) *ListYunQiTaskByUidRequest {
	s.PageNum = &v
	return s
}

func (s *ListYunQiTaskByUidRequest) SetPageSize(v int32) *ListYunQiTaskByUidRequest {
	s.PageSize = &v
	return s
}

func (s *ListYunQiTaskByUidRequest) SetStatusList(v []*string) *ListYunQiTaskByUidRequest {
	s.StatusList = v
	return s
}

type ListYunQiTaskByUidResponseBody struct {
	Code           *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data           *ListYunQiTaskByUidResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpStatusCode *int32                              `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                             `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListYunQiTaskByUidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListYunQiTaskByUidResponseBody) GoString() string {
	return s.String()
}

func (s *ListYunQiTaskByUidResponseBody) SetCode(v string) *ListYunQiTaskByUidResponseBody {
	s.Code = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBody) SetData(v *ListYunQiTaskByUidResponseBodyData) *ListYunQiTaskByUidResponseBody {
	s.Data = v
	return s
}

func (s *ListYunQiTaskByUidResponseBody) SetHttpStatusCode(v int32) *ListYunQiTaskByUidResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBody) SetMessage(v string) *ListYunQiTaskByUidResponseBody {
	s.Message = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBody) SetRequestId(v string) *ListYunQiTaskByUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBody) SetSuccess(v bool) *ListYunQiTaskByUidResponseBody {
	s.Success = &v
	return s
}

type ListYunQiTaskByUidResponseBodyData struct {
	Extend   interface{}                               `json:"extend,omitempty" xml:"extend,omitempty"`
	List     []*ListYunQiTaskByUidResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                    `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListYunQiTaskByUidResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListYunQiTaskByUidResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListYunQiTaskByUidResponseBodyData) SetExtend(v interface{}) *ListYunQiTaskByUidResponseBodyData {
	s.Extend = v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyData) SetList(v []*ListYunQiTaskByUidResponseBodyDataList) *ListYunQiTaskByUidResponseBodyData {
	s.List = v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyData) SetPageNum(v int32) *ListYunQiTaskByUidResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyData) SetPageSize(v int32) *ListYunQiTaskByUidResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyData) SetTotal(v int32) *ListYunQiTaskByUidResponseBodyData {
	s.Total = &v
	return s
}

type ListYunQiTaskByUidResponseBodyDataList struct {
	ChatId               *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	CreateTime           *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorName          *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	EndTime              *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	EvaluationStar       *int32  `json:"evaluationStar,omitempty" xml:"evaluationStar,omitempty"`
	Important            *string `json:"important,omitempty" xml:"important,omitempty"`
	MainDingDepartmentId *string `json:"mainDingDepartmentId,omitempty" xml:"mainDingDepartmentId,omitempty"`
	MainDingGroupId      *string `json:"mainDingGroupId,omitempty" xml:"mainDingGroupId,omitempty"`
	MainDingGroupName    *string `json:"mainDingGroupName,omitempty" xml:"mainDingGroupName,omitempty"`
	ProductName          *string `json:"productName,omitempty" xml:"productName,omitempty"`
	RecordId             *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
	Status               *string `json:"status,omitempty" xml:"status,omitempty"`
	SubDingDepartmentId  *string `json:"subDingDepartmentId,omitempty" xml:"subDingDepartmentId,omitempty"`
	SubDingGroupId       *string `json:"subDingGroupId,omitempty" xml:"subDingGroupId,omitempty"`
	SubDingGroupName     *string `json:"subDingGroupName,omitempty" xml:"subDingGroupName,omitempty"`
	Title                *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListYunQiTaskByUidResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListYunQiTaskByUidResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetChatId(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.ChatId = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetCreateTime(v int64) *ListYunQiTaskByUidResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetCreatorName(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.CreatorName = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetEndTime(v int64) *ListYunQiTaskByUidResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetEvaluationStar(v int32) *ListYunQiTaskByUidResponseBodyDataList {
	s.EvaluationStar = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetImportant(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.Important = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetMainDingDepartmentId(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.MainDingDepartmentId = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetMainDingGroupId(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.MainDingGroupId = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetMainDingGroupName(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.MainDingGroupName = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetProductName(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.ProductName = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetRecordId(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.RecordId = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetStatus(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetSubDingDepartmentId(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.SubDingDepartmentId = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetSubDingGroupId(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.SubDingGroupId = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetSubDingGroupName(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.SubDingGroupName = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetTitle(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.Title = &v
	return s
}

type ListYunQiTaskByUidResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListYunQiTaskByUidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListYunQiTaskByUidResponse) String() string {
	return tea.Prettify(s)
}

func (s ListYunQiTaskByUidResponse) GoString() string {
	return s.String()
}

func (s *ListYunQiTaskByUidResponse) SetHeaders(v map[string]*string) *ListYunQiTaskByUidResponse {
	s.Headers = v
	return s
}

func (s *ListYunQiTaskByUidResponse) SetStatusCode(v int32) *ListYunQiTaskByUidResponse {
	s.StatusCode = &v
	return s
}

func (s *ListYunQiTaskByUidResponse) SetBody(v *ListYunQiTaskByUidResponseBody) *ListYunQiTaskByUidResponse {
	s.Body = v
	return s
}

type MarkFileReadedRequest struct {
	ApplyCode *string `json:"applyCode,omitempty" xml:"applyCode,omitempty"`
	FileId    *int64  `json:"fileId,omitempty" xml:"fileId,omitempty"`
	OrderId   *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	Scene     *string `json:"scene,omitempty" xml:"scene,omitempty"`
}

func (s MarkFileReadedRequest) String() string {
	return tea.Prettify(s)
}

func (s MarkFileReadedRequest) GoString() string {
	return s.String()
}

func (s *MarkFileReadedRequest) SetApplyCode(v string) *MarkFileReadedRequest {
	s.ApplyCode = &v
	return s
}

func (s *MarkFileReadedRequest) SetFileId(v int64) *MarkFileReadedRequest {
	s.FileId = &v
	return s
}

func (s *MarkFileReadedRequest) SetOrderId(v string) *MarkFileReadedRequest {
	s.OrderId = &v
	return s
}

func (s *MarkFileReadedRequest) SetScene(v string) *MarkFileReadedRequest {
	s.Scene = &v
	return s
}

type MarkFileReadedResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *bool   `json:"data,omitempty" xml:"data,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s MarkFileReadedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MarkFileReadedResponseBody) GoString() string {
	return s.String()
}

func (s *MarkFileReadedResponseBody) SetCode(v string) *MarkFileReadedResponseBody {
	s.Code = &v
	return s
}

func (s *MarkFileReadedResponseBody) SetData(v bool) *MarkFileReadedResponseBody {
	s.Data = &v
	return s
}

func (s *MarkFileReadedResponseBody) SetHttpStatusCode(v int32) *MarkFileReadedResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *MarkFileReadedResponseBody) SetMessage(v string) *MarkFileReadedResponseBody {
	s.Message = &v
	return s
}

func (s *MarkFileReadedResponseBody) SetRequestId(v string) *MarkFileReadedResponseBody {
	s.RequestId = &v
	return s
}

func (s *MarkFileReadedResponseBody) SetSuccess(v bool) *MarkFileReadedResponseBody {
	s.Success = &v
	return s
}

type MarkFileReadedResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MarkFileReadedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MarkFileReadedResponse) String() string {
	return tea.Prettify(s)
}

func (s MarkFileReadedResponse) GoString() string {
	return s.String()
}

func (s *MarkFileReadedResponse) SetHeaders(v map[string]*string) *MarkFileReadedResponse {
	s.Headers = v
	return s
}

func (s *MarkFileReadedResponse) SetStatusCode(v int32) *MarkFileReadedResponse {
	s.StatusCode = &v
	return s
}

func (s *MarkFileReadedResponse) SetBody(v *MarkFileReadedResponseBody) *MarkFileReadedResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("customerservice"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDownloadUrlWithOptions(request *GetDownloadUrlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["fileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.FileKey)) {
		body["fileKey"] = request.FileKey
	}

	if !tea.BoolValue(util.IsUnset(request.FreeOrderApplyCode)) {
		body["freeOrderApplyCode"] = request.FreeOrderApplyCode
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["orderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDownloadUrl"),
		Version:     tea.String("2023-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/customerWorkbench/pop/api/file/getDownloadUrl"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDownloadUrl(request *GetDownloadUrlRequest) (_result *GetDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDownloadUrlResponse{}
	_body, _err := client.GetDownloadUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnterpriseSupportPlanDetailWithOptions(request *GetEnterpriseSupportPlanDetailRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEnterpriseSupportPlanDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FreeOrderApplyCodes)) {
		body["freeOrderApplyCodes"] = request.FreeOrderApplyCodes
	}

	if !tea.BoolValue(util.IsUnset(request.OrderIds)) {
		body["orderIds"] = request.OrderIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnterpriseSupportPlanDetail"),
		Version:     tea.String("2023-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/customerWorkbench/pop/api/enterpriseSupport/getEnterpriseSupportPlanDetail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnterpriseSupportPlanDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEnterpriseSupportPlanDetail(request *GetEnterpriseSupportPlanDetailRequest) (_result *GetEnterpriseSupportPlanDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnterpriseSupportPlanDetailResponse{}
	_body, _err := client.GetEnterpriseSupportPlanDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPreViewUrlWithOptions(request *GetPreViewUrlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPreViewUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyCode)) {
		body["applyCode"] = request.ApplyCode
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["fileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.FileKey)) {
		body["fileKey"] = request.FileKey
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["orderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPreViewUrl"),
		Version:     tea.String("2023-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/customerWorkbench/pop/api/file/getPreViewUrl"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPreViewUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPreViewUrl(request *GetPreViewUrlRequest) (_result *GetPreViewUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPreViewUrlResponse{}
	_body, _err := client.GetPreViewUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceDetailWithOptions(request *GetServiceDetailRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyCode)) {
		body["applyCode"] = request.ApplyCode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceDetail"),
		Version:     tea.String("2023-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/customerWorkbench/pop/api/expert/service/getServiceDetail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceDetail(request *GetServiceDetailRequest) (_result *GetServiceDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceDetailResponse{}
	_body, _err := client.GetServiceDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetYunQiTaskByRecordIdWithOptions(request *GetYunQiTaskByRecordIdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetYunQiTaskByRecordIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["recordId"] = request.RecordId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetYunQiTaskByRecordId"),
		Version:     tea.String("2023-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/customerWorkbench/pop/api/record/getYunQiTaskByRecordId"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetYunQiTaskByRecordIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetYunQiTaskByRecordId(request *GetYunQiTaskByRecordIdRequest) (_result *GetYunQiTaskByRecordIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetYunQiTaskByRecordIdResponse{}
	_body, _err := client.GetYunQiTaskByRecordIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDocsGroupByYearWithOptions(request *ListDocsGroupByYearRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDocsGroupByYearResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyCodes)) {
		body["applyCodes"] = request.ApplyCodes
	}

	if !tea.BoolValue(util.IsUnset(request.FileNameKeyword)) {
		body["fileNameKeyword"] = request.FileNameKeyword
	}

	if !tea.BoolValue(util.IsUnset(request.OrderIds)) {
		body["orderIds"] = request.OrderIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["productCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDocsGroupByYear"),
		Version:     tea.String("2023-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/customerWorkbench/pop/api/file/listDocsGroupByYear"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDocsGroupByYearResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDocsGroupByYear(request *ListDocsGroupByYearRequest) (_result *ListDocsGroupByYearResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDocsGroupByYearResponse{}
	_body, _err := client.ListDocsGroupByYearWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnterpriseSupportPlanWithOptions(request *ListEnterpriseSupportPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnterpriseSupportPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnterpriseSupportPlan"),
		Version:     tea.String("2023-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/customerWorkbench/pop/api/enterpriseSupport/listEnterpriseSupportPlan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEnterpriseSupportPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnterpriseSupportPlan(request *ListEnterpriseSupportPlanRequest) (_result *ListEnterpriseSupportPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnterpriseSupportPlanResponse{}
	_body, _err := client.ListEnterpriseSupportPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnterpriseSupportPlanSimpleWithOptions(request *ListEnterpriseSupportPlanSimpleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnterpriseSupportPlanSimpleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnterpriseSupportPlanSimple"),
		Version:     tea.String("2023-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/customerWorkbench/pop/api/enterpriseSupport/listEnterpriseSupportPlanSimple"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEnterpriseSupportPlanSimpleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnterpriseSupportPlanSimple(request *ListEnterpriseSupportPlanSimpleRequest) (_result *ListEnterpriseSupportPlanSimpleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnterpriseSupportPlanSimpleResponse{}
	_body, _err := client.ListEnterpriseSupportPlanSimpleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServiceApplyWithOptions(request *ListServiceApplyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServiceApplyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyType)) {
		body["applyType"] = request.ApplyType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["productCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["startDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceApply"),
		Version:     tea.String("2023-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/customerWorkbench/pop/api/expert/service/listServiceApply"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceApplyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceApply(request *ListServiceApplyRequest) (_result *ListServiceApplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServiceApplyResponse{}
	_body, _err := client.ListServiceApplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListYunQiTaskByUidWithOptions(request *ListYunQiTaskByUidRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListYunQiTaskByUidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateTimeEnd)) {
		body["createTimeEnd"] = request.CreateTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeStart)) {
		body["createTimeStart"] = request.CreateTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.FreeOrderApplyCodes)) {
		body["freeOrderApplyCodes"] = request.FreeOrderApplyCodes
	}

	if !tea.BoolValue(util.IsUnset(request.FreeOrderApplyIds)) {
		body["freeOrderApplyIds"] = request.FreeOrderApplyIds
	}

	if !tea.BoolValue(util.IsUnset(request.OrderIds)) {
		body["orderIds"] = request.OrderIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StatusList)) {
		body["statusList"] = request.StatusList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListYunQiTaskByUid"),
		Version:     tea.String("2023-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/customerWorkbench/pop/api/record/listYunQiTaskByUid"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListYunQiTaskByUidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListYunQiTaskByUid(request *ListYunQiTaskByUidRequest) (_result *ListYunQiTaskByUidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListYunQiTaskByUidResponse{}
	_body, _err := client.ListYunQiTaskByUidWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MarkFileReadedWithOptions(request *MarkFileReadedRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *MarkFileReadedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyCode)) {
		body["applyCode"] = request.ApplyCode
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["fileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["orderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MarkFileReaded"),
		Version:     tea.String("2023-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/customerWorkbench/pop/api/file/markFileReaded"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &MarkFileReadedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MarkFileReaded(request *MarkFileReadedRequest) (_result *MarkFileReadedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MarkFileReadedResponse{}
	_body, _err := client.MarkFileReadedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionApplyOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplyOrders(v *ListPermissionApplyOrdersResponseBodyApplyOrders) *ListPermissionApplyOrdersResponseBody
	GetApplyOrders() *ListPermissionApplyOrdersResponseBodyApplyOrders
	SetRequestId(v string) *ListPermissionApplyOrdersResponseBody
	GetRequestId() *string
}

type ListPermissionApplyOrdersResponseBody struct {
	// The query results returned by page.
	ApplyOrders *ListPermissionApplyOrdersResponseBodyApplyOrders `json:"ApplyOrders,omitempty" xml:"ApplyOrders,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPermissionApplyOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionApplyOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionApplyOrdersResponseBody) GetApplyOrders() *ListPermissionApplyOrdersResponseBodyApplyOrders {
	return s.ApplyOrders
}

func (s *ListPermissionApplyOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPermissionApplyOrdersResponseBody) SetApplyOrders(v *ListPermissionApplyOrdersResponseBodyApplyOrders) *ListPermissionApplyOrdersResponseBody {
	s.ApplyOrders = v
	return s
}

func (s *ListPermissionApplyOrdersResponseBody) SetRequestId(v string) *ListPermissionApplyOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPermissionApplyOrdersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPermissionApplyOrdersResponseBodyApplyOrders struct {
	// The list of the permission request orders.
	ApplyOrder []*ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder `json:"ApplyOrder,omitempty" xml:"ApplyOrder,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 150
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPermissionApplyOrdersResponseBodyApplyOrders) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionApplyOrdersResponseBodyApplyOrders) GoString() string {
	return s.String()
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrders) GetApplyOrder() []*ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder {
	return s.ApplyOrder
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrders) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrders) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrders) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrders) SetApplyOrder(v []*ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) *ListPermissionApplyOrdersResponseBodyApplyOrders {
	s.ApplyOrder = v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrders) SetPageNumber(v int32) *ListPermissionApplyOrdersResponseBodyApplyOrders {
	s.PageNumber = &v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrders) SetPageSize(v int32) *ListPermissionApplyOrdersResponseBodyApplyOrders {
	s.PageSize = &v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrders) SetTotalCount(v int32) *ListPermissionApplyOrdersResponseBodyApplyOrders {
	s.TotalCount = &v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrders) Validate() error {
	return dara.Validate(s)
}

type ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder struct {
	// The ID of the Alibaba Cloud account that was used to submit the permission request order.
	//
	// example:
	//
	// 267842600408993176
	ApplyBaseId *string `json:"ApplyBaseId,omitempty" xml:"ApplyBaseId,omitempty"`
	// The time when the permission request order was submitted. The parameter value is a UNIX timestamp.
	//
	// example:
	//
	// 1615284086000
	ApplyTimestamp *int64 `json:"ApplyTimestamp,omitempty" xml:"ApplyTimestamp,omitempty"`
	// The content of the permission request order.
	ApproveContent          *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent `json:"ApproveContent,omitempty" xml:"ApproveContent,omitempty" type:"Struct"`
	FinishApprovalComment   *string                                                                   `json:"FinishApprovalComment,omitempty" xml:"FinishApprovalComment,omitempty"`
	FinishApprovalTimestamp *int64                                                                    `json:"FinishApprovalTimestamp,omitempty" xml:"FinishApprovalTimestamp,omitempty"`
	// The ID of the permission request order.
	//
	// example:
	//
	// ad8da78d-8135-455e-9486-27cf213fc140
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The status of the permission request order. Valid values:
	//
	// 	- 1: to be processed
	//
	// 	- 2: approved and authorized
	//
	// 	- 3: approved but authorization failed
	//
	// 	- 4: rejected
	//
	// example:
	//
	// 2
	FlowStatus *int32 `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty"`
}

func (s ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) GoString() string {
	return s.String()
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) GetApplyBaseId() *string {
	return s.ApplyBaseId
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) GetApplyTimestamp() *int64 {
	return s.ApplyTimestamp
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) GetApproveContent() *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent {
	return s.ApproveContent
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) GetFinishApprovalComment() *string {
	return s.FinishApprovalComment
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) GetFinishApprovalTimestamp() *int64 {
	return s.FinishApprovalTimestamp
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) GetFlowId() *string {
	return s.FlowId
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) GetFlowStatus() *int32 {
	return s.FlowStatus
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) SetApplyBaseId(v string) *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder {
	s.ApplyBaseId = &v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) SetApplyTimestamp(v int64) *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder {
	s.ApplyTimestamp = &v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) SetApproveContent(v *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent) *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder {
	s.ApproveContent = v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) SetFinishApprovalComment(v string) *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder {
	s.FinishApprovalComment = &v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) SetFinishApprovalTimestamp(v int64) *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder {
	s.FinishApprovalTimestamp = &v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) SetFlowId(v string) *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder {
	s.FlowId = &v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) SetFlowStatus(v int32) *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder {
	s.FlowStatus = &v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder) Validate() error {
	return dara.Validate(s)
}

type ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent struct {
	// The reason for your request. The administrator determines whether to approve the request based on the reason.
	//
	// example:
	//
	// I need to use this table
	ApplyReason *string `json:"ApplyReason,omitempty" xml:"ApplyReason,omitempty"`
	// The type of the permission request order. The parameter value is 1 and cannot be changed. This value indicates ACL-based authorization.
	//
	// example:
	//
	// 1
	OrderType *int32 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The content of the object on which you requested permissions.
	ProjectMeta *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMeta `json:"ProjectMeta,omitempty" xml:"ProjectMeta,omitempty" type:"Struct"`
}

func (s ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent) GoString() string {
	return s.String()
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent) GetApplyReason() *string {
	return s.ApplyReason
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent) GetOrderType() *int32 {
	return s.OrderType
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent) GetProjectMeta() *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMeta {
	return s.ProjectMeta
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent) SetApplyReason(v string) *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent {
	s.ApplyReason = &v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent) SetOrderType(v int32) *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent {
	s.OrderType = &v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent) SetProjectMeta(v *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMeta) *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent {
	s.ProjectMeta = v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent) Validate() error {
	return dara.Validate(s)
}

type ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMeta struct {
	// The information about the object on which you requested permissions.
	ObjectMetaList []*ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMetaObjectMetaList `json:"ObjectMetaList,omitempty" xml:"ObjectMetaList,omitempty" type:"Repeated"`
	// The name of the DataWorks workspace that is associated with the MaxCompute project in which you requested permissions on a table.
	//
	// example:
	//
	// aWorkspaceName
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMeta) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMeta) GoString() string {
	return s.String()
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMeta) GetObjectMetaList() []*ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMetaObjectMetaList {
	return s.ObjectMetaList
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMeta) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMeta) SetObjectMetaList(v []*ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMetaObjectMetaList) *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMeta {
	s.ObjectMetaList = v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMeta) SetWorkspaceName(v string) *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMeta {
	s.WorkspaceName = &v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMeta) Validate() error {
	return dara.Validate(s)
}

type ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMetaObjectMetaList struct {
	Actions []*string `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The name of the table on which you requested permissions.
	//
	// example:
	//
	// aTableName
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
}

func (s ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMetaObjectMetaList) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMetaObjectMetaList) GoString() string {
	return s.String()
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMetaObjectMetaList) GetActions() []*string {
	return s.Actions
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMetaObjectMetaList) GetObjectName() *string {
	return s.ObjectName
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMetaObjectMetaList) SetActions(v []*string) *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMetaObjectMetaList {
	s.Actions = v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMetaObjectMetaList) SetObjectName(v string) *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMetaObjectMetaList {
	s.ObjectName = &v
	return s
}

func (s *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMetaObjectMetaList) Validate() error {
	return dara.Validate(s)
}

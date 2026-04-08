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
	// The paginated query results of permission requests.
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
	if s.ApplyOrders != nil {
		if err := s.ApplyOrders.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPermissionApplyOrdersResponseBodyApplyOrders struct {
	// The list of permission requests.
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
	// The total number of permission requests returned.
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
	if s.ApplyOrder != nil {
		for _, item := range s.ApplyOrder {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrder struct {
	// The Alibaba Cloud account ID of the user who submitted the permission request.
	//
	// example:
	//
	// 267842600408993176
	ApplyBaseId *string `json:"ApplyBaseId,omitempty" xml:"ApplyBaseId,omitempty"`
	// The time when the permission request was submitted, in Unix timestamp format.
	//
	// example:
	//
	// 1615284086000
	ApplyTimestamp *int64 `json:"ApplyTimestamp,omitempty" xml:"ApplyTimestamp,omitempty"`
	// The content of the permission request.
	ApproveContent *ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent `json:"ApproveContent,omitempty" xml:"ApproveContent,omitempty" type:"Struct"`
	// The final approval comment.
	//
	// example:
	//
	// agree
	FinishApprovalComment *string `json:"FinishApprovalComment,omitempty" xml:"FinishApprovalComment,omitempty"`
	// The final approval timestamp. Displayed as a Unix timestamp.
	//
	// example:
	//
	// 1757496687000
	FinishApprovalTimestamp *int64 `json:"FinishApprovalTimestamp,omitempty" xml:"FinishApprovalTimestamp,omitempty"`
	// The permission request ID.
	//
	// example:
	//
	// ad8da78d-8135-455e-9486-27cf213fc140
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The status of the permission request. Valid values:
	//
	// 	- 1: Pending approval
	//
	// 	- 2: Approved and authorization succeeded
	//
	// 	- 3: Approved but authorization failed
	//
	// 	- 4: Rejected
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
	if s.ApproveContent != nil {
		if err := s.ApproveContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContent struct {
	// The reason for the permission request, which is used by administrators for evaluation and approval.
	//
	// example:
	//
	// I need to use this table
	ApplyReason *string `json:"ApplyReason,omitempty" xml:"ApplyReason,omitempty"`
	// The type of permission request. Only the value 1 is supported, which indicates an ACL permission request for objects.
	//
	// example:
	//
	// 1
	OrderType *int32 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The content of the requested object.
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
	if s.ProjectMeta != nil {
		if err := s.ProjectMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMeta struct {
	// The information about the requested object.
	ObjectMetaList []*ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMetaObjectMetaList `json:"ObjectMetaList,omitempty" xml:"ObjectMetaList,omitempty" type:"Repeated"`
	// The name of the DataWorks workspace that contains the MaxCompute project for which permissions are requested.
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
	if s.ObjectMetaList != nil {
		for _, item := range s.ObjectMetaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPermissionApplyOrdersResponseBodyApplyOrdersApplyOrderApproveContentProjectMetaObjectMetaList struct {
	// The operation type.
	Actions []*string `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The name of the requested table.
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

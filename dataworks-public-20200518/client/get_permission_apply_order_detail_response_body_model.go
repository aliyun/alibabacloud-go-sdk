// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermissionApplyOrderDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplyOrderDetail(v *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) *GetPermissionApplyOrderDetailResponseBody
	GetApplyOrderDetail() *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail
	SetRequestId(v string) *GetPermissionApplyOrderDetailResponseBody
	GetRequestId() *string
}

type GetPermissionApplyOrderDetailResponseBody struct {
	// The details of the permission request order.
	ApplyOrderDetail *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail `json:"ApplyOrderDetail,omitempty" xml:"ApplyOrderDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPermissionApplyOrderDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionApplyOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetPermissionApplyOrderDetailResponseBody) GetApplyOrderDetail() *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail {
	return s.ApplyOrderDetail
}

func (s *GetPermissionApplyOrderDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPermissionApplyOrderDetailResponseBody) SetApplyOrderDetail(v *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) *GetPermissionApplyOrderDetailResponseBody {
	s.ApplyOrderDetail = v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBody) SetRequestId(v string) *GetPermissionApplyOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail struct {
	// The ID of the Alibaba Cloud account that was used to submit the permission request order.
	//
	// example:
	//
	// 267842600408993176
	ApplyBaseId *string `json:"ApplyBaseId,omitempty" xml:"ApplyBaseId,omitempty"`
	// The time when the permission request order was submitted. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1615284086000
	ApplyTimestamp *int64 `json:"ApplyTimestamp,omitempty" xml:"ApplyTimestamp,omitempty"`
	// The list of Alibaba Cloud accounts that are used to process the permission request order.
	ApproveAccountList []*GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveAccountList `json:"ApproveAccountList,omitempty" xml:"ApproveAccountList,omitempty" type:"Repeated"`
	// The content of the permission request.
	ApproveContent           *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent `json:"ApproveContent,omitempty" xml:"ApproveContent,omitempty" type:"Struct"`
	FinishAapprovalTimestamp *int64                                                                   `json:"FinishAapprovalTimestamp,omitempty" xml:"FinishAapprovalTimestamp,omitempty"`
	FinishApprovalComment    *string                                                                  `json:"FinishApprovalComment,omitempty" xml:"FinishApprovalComment,omitempty"`
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
	// The information about the account that is used to request permissions.
	GranteeObjectList []*GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList `json:"GranteeObjectList,omitempty" xml:"GranteeObjectList,omitempty" type:"Repeated"`
}

func (s GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) GoString() string {
	return s.String()
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) GetApplyBaseId() *string {
	return s.ApplyBaseId
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) GetApplyTimestamp() *int64 {
	return s.ApplyTimestamp
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) GetApproveAccountList() []*GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveAccountList {
	return s.ApproveAccountList
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) GetApproveContent() *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent {
	return s.ApproveContent
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) GetFinishAapprovalTimestamp() *int64 {
	return s.FinishAapprovalTimestamp
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) GetFinishApprovalComment() *string {
	return s.FinishApprovalComment
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) GetFlowId() *string {
	return s.FlowId
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) GetFlowStatus() *int32 {
	return s.FlowStatus
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) GetGranteeObjectList() []*GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList {
	return s.GranteeObjectList
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) SetApplyBaseId(v string) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail {
	s.ApplyBaseId = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) SetApplyTimestamp(v int64) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail {
	s.ApplyTimestamp = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) SetApproveAccountList(v []*GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveAccountList) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail {
	s.ApproveAccountList = v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) SetApproveContent(v *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail {
	s.ApproveContent = v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) SetFinishAapprovalTimestamp(v int64) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail {
	s.FinishAapprovalTimestamp = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) SetFinishApprovalComment(v string) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail {
	s.FinishApprovalComment = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) SetFlowId(v string) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail {
	s.FlowId = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) SetFlowStatus(v int32) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail {
	s.FlowStatus = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) SetGranteeObjectList(v []*GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail {
	s.GranteeObjectList = v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetail) Validate() error {
	return dara.Validate(s)
}

type GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveAccountList struct {
	// The ID of the Alibaba Cloud account that is used to process the permission request order.
	//
	// example:
	//
	// 182293110403****
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
}

func (s GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveAccountList) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveAccountList) GoString() string {
	return s.String()
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveAccountList) GetBaseId() *string {
	return s.BaseId
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveAccountList) SetBaseId(v string) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveAccountList {
	s.BaseId = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveAccountList) Validate() error {
	return dara.Validate(s)
}

type GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent struct {
	// The reason of the permission request. The administrator processes the request based on the reason.
	//
	// example:
	//
	// I need to use this table
	ApplyReason *string `json:"ApplyReason,omitempty" xml:"ApplyReason,omitempty"`
	// The expiration time of the permissions that you request. The value is a UNIX timestamp. If LabelSecurity is disabled for the MaxCompute project in which you want to request permissions on the fields of a table, or the security level of the fields is 0 or is lower than or equal to the security level of the Alibaba Cloud account for which you want to request permissions, you can request only permanent permissions.
	//
	// example:
	//
	// 1617115071885
	Deadline *int64 `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// The type of the permission request order. The parameter value is 1 and cannot be changed. This value indicates ACL-based authorization.
	//
	// example:
	//
	// 1
	OrderType *int32 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The information about the project and workspace that are associated with the object on which you request permissions.
	ProjectMeta *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta `json:"ProjectMeta,omitempty" xml:"ProjectMeta,omitempty" type:"Struct"`
}

func (s GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent) GoString() string {
	return s.String()
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent) GetApplyReason() *string {
	return s.ApplyReason
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent) GetDeadline() *int64 {
	return s.Deadline
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent) GetOrderType() *int32 {
	return s.OrderType
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent) GetProjectMeta() *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta {
	return s.ProjectMeta
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent) SetApplyReason(v string) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent {
	s.ApplyReason = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent) SetDeadline(v int64) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent {
	s.Deadline = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent) SetOrderType(v int32) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent {
	s.OrderType = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent) SetProjectMeta(v *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent {
	s.ProjectMeta = v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContent) Validate() error {
	return dara.Validate(s)
}

type GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta struct {
	// The MaxCompute project to which the object on which you request permissions belongs.
	//
	// example:
	//
	// aMaxComputeProject
	MaxComputeProjectName *string `json:"MaxComputeProjectName,omitempty" xml:"MaxComputeProjectName,omitempty"`
	// The details about the object on which you request permissions.
	ObjectMetaList []*GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList `json:"ObjectMetaList,omitempty" xml:"ObjectMetaList,omitempty" type:"Repeated"`
	// The ID of the DataWorks workspace that is associated with the object on which you request permissions.
	//
	// example:
	//
	// 12345
	WorkspaceId *int32 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta) GoString() string {
	return s.String()
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta) GetMaxComputeProjectName() *string {
	return s.MaxComputeProjectName
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta) GetObjectMetaList() []*GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList {
	return s.ObjectMetaList
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta) GetWorkspaceId() *int32 {
	return s.WorkspaceId
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta) SetMaxComputeProjectName(v string) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta {
	s.MaxComputeProjectName = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta) SetObjectMetaList(v []*GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta {
	s.ObjectMetaList = v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta) SetWorkspaceId(v int32) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta {
	s.WorkspaceId = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMeta) Validate() error {
	return dara.Validate(s)
}

type GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList struct {
	Actions []*string `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The information about the column fields in the object on which you request permissions.
	ColumnMetaList []*GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList `json:"ColumnMetaList,omitempty" xml:"ColumnMetaList,omitempty" type:"Repeated"`
	// The name of the table on which you request permissions.
	//
	// example:
	//
	// aTableName
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
}

func (s GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList) GoString() string {
	return s.String()
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList) GetActions() []*string {
	return s.Actions
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList) GetColumnMetaList() []*GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList {
	return s.ColumnMetaList
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList) GetObjectName() *string {
	return s.ObjectName
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList) SetActions(v []*string) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList {
	s.Actions = v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList) SetColumnMetaList(v []*GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList {
	s.ColumnMetaList = v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList) SetObjectName(v string) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList {
	s.ObjectName = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaList) Validate() error {
	return dara.Validate(s)
}

type GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList struct {
	ColumnActions []*string `json:"ColumnActions,omitempty" xml:"ColumnActions,omitempty" type:"Repeated"`
	// The description of the column on which you request permissions.
	//
	// example:
	//
	// Field description
	ColumnComment *string `json:"ColumnComment,omitempty" xml:"ColumnComment,omitempty"`
	// The name of the column on which you request permissions.
	//
	// example:
	//
	// aColumnName
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The security level of the column on which you request permissions. Valid values: 0 to 9.
	//
	// example:
	//
	// 9
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList) GoString() string {
	return s.String()
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList) GetColumnActions() []*string {
	return s.ColumnActions
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList) GetColumnComment() *string {
	return s.ColumnComment
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList) SetColumnActions(v []*string) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList {
	s.ColumnActions = v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList) SetColumnComment(v string) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList {
	s.ColumnComment = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList) SetColumnName(v string) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList {
	s.ColumnName = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList) SetSecurityLevel(v string) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList {
	s.SecurityLevel = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailApproveContentProjectMetaObjectMetaListColumnMetaList) Validate() error {
	return dara.Validate(s)
}

type GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList struct {
	// The ID of the account that is used to request permissions.
	//
	// example:
	//
	// 267842600408993176
	GranteeId *string `json:"GranteeId,omitempty" xml:"GranteeId,omitempty"`
	// The name of the account that is used to request permissions. The name is in the same format as that of the account used to access the MaxCompute project.
	//
	// 	- If the account is an Alibaba Cloud account, the value is in the ALIYUN$+Account name format.
	//
	// 	- If the account is a RAM user, the value is in the RAM$+Account name format.
	//
	// example:
	//
	// RAM$dataworks_3h1_1:StsRamUser(StsRamUser)
	GranteeName *string `json:"GranteeName,omitempty" xml:"GranteeName,omitempty"`
	// The type of the subject that requests permissions. The value is fixed as 1, which indicates users.
	//
	// example:
	//
	// 1
	GranteeType *int32 `json:"GranteeType,omitempty" xml:"GranteeType,omitempty"`
	// The subtype of the subject that requests permissions. Valid values:
	//
	// 	- 101: production account
	//
	// 	- 103: individual account
	//
	// 	- 105: account that requests permissions for others
	//
	// example:
	//
	// 103
	GranteeTypeSub *int32 `json:"GranteeTypeSub,omitempty" xml:"GranteeTypeSub,omitempty"`
}

func (s GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList) GoString() string {
	return s.String()
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList) GetGranteeId() *string {
	return s.GranteeId
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList) GetGranteeName() *string {
	return s.GranteeName
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList) GetGranteeType() *int32 {
	return s.GranteeType
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList) GetGranteeTypeSub() *int32 {
	return s.GranteeTypeSub
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList) SetGranteeId(v string) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList {
	s.GranteeId = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList) SetGranteeName(v string) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList {
	s.GranteeName = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList) SetGranteeType(v int32) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList {
	s.GranteeType = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList) SetGranteeTypeSub(v int32) *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList {
	s.GranteeTypeSub = &v
	return s
}

func (s *GetPermissionApplyOrderDetailResponseBodyApplyOrderDetailGranteeObjectList) Validate() error {
	return dara.Validate(s)
}

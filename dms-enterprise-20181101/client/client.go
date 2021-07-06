// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type SubmitStructSyncOrderApprovalRequest struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Tid     *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s SubmitStructSyncOrderApprovalRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitStructSyncOrderApprovalRequest) GoString() string {
	return s.String()
}

func (s *SubmitStructSyncOrderApprovalRequest) SetOrderId(v int64) *SubmitStructSyncOrderApprovalRequest {
	s.OrderId = &v
	return s
}

func (s *SubmitStructSyncOrderApprovalRequest) SetTid(v int64) *SubmitStructSyncOrderApprovalRequest {
	s.Tid = &v
	return s
}

type SubmitStructSyncOrderApprovalResponseBody struct {
	// Id of the request
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage       *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode          *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	WorkflowInstanceId *int64  `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
}

func (s SubmitStructSyncOrderApprovalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitStructSyncOrderApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitStructSyncOrderApprovalResponseBody) SetRequestId(v string) *SubmitStructSyncOrderApprovalResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitStructSyncOrderApprovalResponseBody) SetSuccess(v bool) *SubmitStructSyncOrderApprovalResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitStructSyncOrderApprovalResponseBody) SetErrorMessage(v string) *SubmitStructSyncOrderApprovalResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SubmitStructSyncOrderApprovalResponseBody) SetErrorCode(v string) *SubmitStructSyncOrderApprovalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitStructSyncOrderApprovalResponseBody) SetWorkflowInstanceId(v int64) *SubmitStructSyncOrderApprovalResponseBody {
	s.WorkflowInstanceId = &v
	return s
}

type SubmitStructSyncOrderApprovalResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitStructSyncOrderApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitStructSyncOrderApprovalResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitStructSyncOrderApprovalResponse) GoString() string {
	return s.String()
}

func (s *SubmitStructSyncOrderApprovalResponse) SetHeaders(v map[string]*string) *SubmitStructSyncOrderApprovalResponse {
	s.Headers = v
	return s
}

func (s *SubmitStructSyncOrderApprovalResponse) SetBody(v *SubmitStructSyncOrderApprovalResponseBody) *SubmitStructSyncOrderApprovalResponse {
	s.Body = v
	return s
}

type ListDatabaseUserPermssionsRequest struct {
	Tid        *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	PermType   *string `json:"PermType,omitempty" xml:"PermType,omitempty"`
	DbId       *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Logic      *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDatabaseUserPermssionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatabaseUserPermssionsRequest) GoString() string {
	return s.String()
}

func (s *ListDatabaseUserPermssionsRequest) SetTid(v int64) *ListDatabaseUserPermssionsRequest {
	s.Tid = &v
	return s
}

func (s *ListDatabaseUserPermssionsRequest) SetPermType(v string) *ListDatabaseUserPermssionsRequest {
	s.PermType = &v
	return s
}

func (s *ListDatabaseUserPermssionsRequest) SetDbId(v string) *ListDatabaseUserPermssionsRequest {
	s.DbId = &v
	return s
}

func (s *ListDatabaseUserPermssionsRequest) SetLogic(v bool) *ListDatabaseUserPermssionsRequest {
	s.Logic = &v
	return s
}

func (s *ListDatabaseUserPermssionsRequest) SetUserName(v string) *ListDatabaseUserPermssionsRequest {
	s.UserName = &v
	return s
}

func (s *ListDatabaseUserPermssionsRequest) SetPageNumber(v int32) *ListDatabaseUserPermssionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatabaseUserPermssionsRequest) SetPageSize(v int32) *ListDatabaseUserPermssionsRequest {
	s.PageSize = &v
	return s
}

type ListDatabaseUserPermssionsResponseBody struct {
	TotalCount      *int64                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId       *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode       *string                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage    *string                                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	UserPermissions *ListDatabaseUserPermssionsResponseBodyUserPermissions `json:"UserPermissions,omitempty" xml:"UserPermissions,omitempty" type:"Struct"`
	Success         *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDatabaseUserPermssionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatabaseUserPermssionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabaseUserPermssionsResponseBody) SetTotalCount(v int64) *ListDatabaseUserPermssionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBody) SetRequestId(v string) *ListDatabaseUserPermssionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBody) SetErrorCode(v string) *ListDatabaseUserPermssionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBody) SetErrorMessage(v string) *ListDatabaseUserPermssionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBody) SetUserPermissions(v *ListDatabaseUserPermssionsResponseBodyUserPermissions) *ListDatabaseUserPermssionsResponseBody {
	s.UserPermissions = v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBody) SetSuccess(v bool) *ListDatabaseUserPermssionsResponseBody {
	s.Success = &v
	return s
}

type ListDatabaseUserPermssionsResponseBodyUserPermissions struct {
	UserPermission []*ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission `json:"UserPermission,omitempty" xml:"UserPermission,omitempty" type:"Repeated"`
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissions) String() string {
	return tea.Prettify(s)
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissions) GoString() string {
	return s.String()
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissions) SetUserPermission(v []*ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) *ListDatabaseUserPermssionsResponseBodyUserPermissions {
	s.UserPermission = v
	return s
}

type ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission struct {
	DbId         *string                                                                         `json:"DbId,omitempty" xml:"DbId,omitempty"`
	TableName    *string                                                                         `json:"TableName,omitempty" xml:"TableName,omitempty"`
	UserId       *string                                                                         `json:"UserId,omitempty" xml:"UserId,omitempty"`
	SchemaName   *string                                                                         `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	Logic        *bool                                                                           `json:"Logic,omitempty" xml:"Logic,omitempty"`
	UserNickName *string                                                                         `json:"UserNickName,omitempty" xml:"UserNickName,omitempty"`
	InstanceId   *string                                                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PermDetails  *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails `json:"PermDetails,omitempty" xml:"PermDetails,omitempty" type:"Struct"`
	EnvType      *string                                                                         `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	ColumnName   *string                                                                         `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	DbType       *string                                                                         `json:"DbType,omitempty" xml:"DbType,omitempty"`
	DsType       *string                                                                         `json:"DsType,omitempty" xml:"DsType,omitempty"`
	TableId      *string                                                                         `json:"TableId,omitempty" xml:"TableId,omitempty"`
	SearchName   *string                                                                         `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	Alias        *string                                                                         `json:"Alias,omitempty" xml:"Alias,omitempty"`
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) String() string {
	return tea.Prettify(s)
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) GoString() string {
	return s.String()
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetDbId(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.DbId = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetTableName(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.TableName = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetUserId(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.UserId = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetSchemaName(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.SchemaName = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetLogic(v bool) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.Logic = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetUserNickName(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.UserNickName = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetInstanceId(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.InstanceId = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetPermDetails(v *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.PermDetails = v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetEnvType(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.EnvType = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetColumnName(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.ColumnName = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetDbType(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.DbType = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetDsType(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.DsType = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetTableId(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.TableId = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetSearchName(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.SearchName = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission) SetAlias(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermission {
	s.Alias = &v
	return s
}

type ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails struct {
	PermDetail []*ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail `json:"PermDetail,omitempty" xml:"PermDetail,omitempty" type:"Repeated"`
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails) String() string {
	return tea.Prettify(s)
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails) GoString() string {
	return s.String()
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails) SetPermDetail(v []*ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetails {
	s.PermDetail = v
	return s
}

type ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail struct {
	OriginFrom   *string `json:"OriginFrom,omitempty" xml:"OriginFrom,omitempty"`
	PermType     *string `json:"PermType,omitempty" xml:"PermType,omitempty"`
	ExpireDate   *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	CreateDate   *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	UserAccessId *string `json:"UserAccessId,omitempty" xml:"UserAccessId,omitempty"`
	ExtraData    *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) String() string {
	return tea.Prettify(s)
}

func (s ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GoString() string {
	return s.String()
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetOriginFrom(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.OriginFrom = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetPermType(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.PermType = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetExpireDate(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.ExpireDate = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetCreateDate(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.CreateDate = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetUserAccessId(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.UserAccessId = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetExtraData(v string) *ListDatabaseUserPermssionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.ExtraData = &v
	return s
}

type ListDatabaseUserPermssionsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDatabaseUserPermssionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDatabaseUserPermssionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDatabaseUserPermssionsResponse) GoString() string {
	return s.String()
}

func (s *ListDatabaseUserPermssionsResponse) SetHeaders(v map[string]*string) *ListDatabaseUserPermssionsResponse {
	s.Headers = v
	return s
}

func (s *ListDatabaseUserPermssionsResponse) SetBody(v *ListDatabaseUserPermssionsResponseBody) *ListDatabaseUserPermssionsResponse {
	s.Body = v
	return s
}

type ListSensitiveColumnsRequest struct {
	Tid           *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	SchemaName    *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName     *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	ColumnName    *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSensitiveColumnsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveColumnsRequest) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsRequest) SetTid(v int64) *ListSensitiveColumnsRequest {
	s.Tid = &v
	return s
}

func (s *ListSensitiveColumnsRequest) SetSchemaName(v string) *ListSensitiveColumnsRequest {
	s.SchemaName = &v
	return s
}

func (s *ListSensitiveColumnsRequest) SetTableName(v string) *ListSensitiveColumnsRequest {
	s.TableName = &v
	return s
}

func (s *ListSensitiveColumnsRequest) SetColumnName(v string) *ListSensitiveColumnsRequest {
	s.ColumnName = &v
	return s
}

func (s *ListSensitiveColumnsRequest) SetSecurityLevel(v string) *ListSensitiveColumnsRequest {
	s.SecurityLevel = &v
	return s
}

func (s *ListSensitiveColumnsRequest) SetPageNumber(v int32) *ListSensitiveColumnsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSensitiveColumnsRequest) SetPageSize(v int32) *ListSensitiveColumnsRequest {
	s.PageSize = &v
	return s
}

type ListSensitiveColumnsResponseBody struct {
	TotalCount          *int64                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId           *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode           *string                                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage        *string                                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	SensitiveColumnList *ListSensitiveColumnsResponseBodySensitiveColumnList `json:"SensitiveColumnList,omitempty" xml:"SensitiveColumnList,omitempty" type:"Struct"`
	Success             *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSensitiveColumnsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsResponseBody) SetTotalCount(v int64) *ListSensitiveColumnsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSensitiveColumnsResponseBody) SetRequestId(v string) *ListSensitiveColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSensitiveColumnsResponseBody) SetErrorCode(v string) *ListSensitiveColumnsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSensitiveColumnsResponseBody) SetErrorMessage(v string) *ListSensitiveColumnsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSensitiveColumnsResponseBody) SetSensitiveColumnList(v *ListSensitiveColumnsResponseBodySensitiveColumnList) *ListSensitiveColumnsResponseBody {
	s.SensitiveColumnList = v
	return s
}

func (s *ListSensitiveColumnsResponseBody) SetSuccess(v bool) *ListSensitiveColumnsResponseBody {
	s.Success = &v
	return s
}

type ListSensitiveColumnsResponseBodySensitiveColumnList struct {
	SensitiveColumn []*ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn `json:"SensitiveColumn,omitempty" xml:"SensitiveColumn,omitempty" type:"Repeated"`
}

func (s ListSensitiveColumnsResponseBodySensitiveColumnList) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveColumnsResponseBodySensitiveColumnList) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnList) SetSensitiveColumn(v []*ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) *ListSensitiveColumnsResponseBodySensitiveColumnList {
	s.SensitiveColumn = v
	return s
}

type ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn struct {
	ColumnName    *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	TableName     *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	ColumnCount   *int64  `json:"ColumnCount,omitempty" xml:"ColumnCount,omitempty"`
	SchemaName    *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	FunctionType  *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
}

func (s ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) SetColumnName(v string) *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn {
	s.ColumnName = &v
	return s
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) SetTableName(v string) *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn {
	s.TableName = &v
	return s
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) SetSecurityLevel(v string) *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn {
	s.SecurityLevel = &v
	return s
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) SetColumnCount(v int64) *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn {
	s.ColumnCount = &v
	return s
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) SetSchemaName(v string) *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn {
	s.SchemaName = &v
	return s
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) SetFunctionType(v string) *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn {
	s.FunctionType = &v
	return s
}

type ListSensitiveColumnsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSensitiveColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSensitiveColumnsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveColumnsResponse) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsResponse) SetHeaders(v map[string]*string) *ListSensitiveColumnsResponse {
	s.Headers = v
	return s
}

func (s *ListSensitiveColumnsResponse) SetBody(v *ListSensitiveColumnsResponseBody) *ListSensitiveColumnsResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	Tid        *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	Role       *string `json:"Role,omitempty" xml:"Role,omitempty"`
	UserState  *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SearchKey  *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetTid(v int64) *ListUsersRequest {
	s.Tid = &v
	return s
}

func (s *ListUsersRequest) SetRole(v string) *ListUsersRequest {
	s.Role = &v
	return s
}

func (s *ListUsersRequest) SetUserState(v string) *ListUsersRequest {
	s.UserState = &v
	return s
}

func (s *ListUsersRequest) SetSearchKey(v string) *ListUsersRequest {
	s.SearchKey = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

type ListUsersResponseBody struct {
	TotalCount   *int64                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId    *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	UserList     *ListUsersResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Struct"`
	Success      *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetTotalCount(v int64) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetErrorCode(v string) *ListUsersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUsersResponseBody) SetErrorMessage(v string) *ListUsersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListUsersResponseBody) SetUserList(v *ListUsersResponseBodyUserList) *ListUsersResponseBody {
	s.UserList = v
	return s
}

func (s *ListUsersResponseBody) SetSuccess(v bool) *ListUsersResponseBody {
	s.Success = &v
	return s
}

type ListUsersResponseBodyUserList struct {
	User []*ListUsersResponseBodyUserListUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUserList) SetUser(v []*ListUsersResponseBodyUserListUser) *ListUsersResponseBodyUserList {
	s.User = v
	return s
}

type ListUsersResponseBodyUserListUser struct {
	State           *string                                        `json:"State,omitempty" xml:"State,omitempty"`
	CurResultCount  *int64                                         `json:"CurResultCount,omitempty" xml:"CurResultCount,omitempty"`
	UserId          *string                                        `json:"UserId,omitempty" xml:"UserId,omitempty"`
	LastLoginTime   *string                                        `json:"LastLoginTime,omitempty" xml:"LastLoginTime,omitempty"`
	MaxResultCount  *int64                                         `json:"MaxResultCount,omitempty" xml:"MaxResultCount,omitempty"`
	ParentUid       *string                                        `json:"ParentUid,omitempty" xml:"ParentUid,omitempty"`
	RoleIdList      *ListUsersResponseBodyUserListUserRoleIdList   `json:"RoleIdList,omitempty" xml:"RoleIdList,omitempty" type:"Struct"`
	RoleNameList    *ListUsersResponseBodyUserListUserRoleNameList `json:"RoleNameList,omitempty" xml:"RoleNameList,omitempty" type:"Struct"`
	NickName        *string                                        `json:"NickName,omitempty" xml:"NickName,omitempty"`
	MaxExecuteCount *int64                                         `json:"MaxExecuteCount,omitempty" xml:"MaxExecuteCount,omitempty"`
	CurExecuteCount *int64                                         `json:"CurExecuteCount,omitempty" xml:"CurExecuteCount,omitempty"`
	Mobile          *string                                        `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Uid             *string                                        `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListUsersResponseBodyUserListUser) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUserListUser) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUserListUser) SetState(v string) *ListUsersResponseBodyUserListUser {
	s.State = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetCurResultCount(v int64) *ListUsersResponseBodyUserListUser {
	s.CurResultCount = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetUserId(v string) *ListUsersResponseBodyUserListUser {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetLastLoginTime(v string) *ListUsersResponseBodyUserListUser {
	s.LastLoginTime = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetMaxResultCount(v int64) *ListUsersResponseBodyUserListUser {
	s.MaxResultCount = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetParentUid(v string) *ListUsersResponseBodyUserListUser {
	s.ParentUid = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetRoleIdList(v *ListUsersResponseBodyUserListUserRoleIdList) *ListUsersResponseBodyUserListUser {
	s.RoleIdList = v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetRoleNameList(v *ListUsersResponseBodyUserListUserRoleNameList) *ListUsersResponseBodyUserListUser {
	s.RoleNameList = v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetNickName(v string) *ListUsersResponseBodyUserListUser {
	s.NickName = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetMaxExecuteCount(v int64) *ListUsersResponseBodyUserListUser {
	s.MaxExecuteCount = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetCurExecuteCount(v int64) *ListUsersResponseBodyUserListUser {
	s.CurExecuteCount = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetMobile(v string) *ListUsersResponseBodyUserListUser {
	s.Mobile = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetUid(v string) *ListUsersResponseBodyUserListUser {
	s.Uid = &v
	return s
}

type ListUsersResponseBodyUserListUserRoleIdList struct {
	RoleIds []*int32 `json:"RoleIds,omitempty" xml:"RoleIds,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUserListUserRoleIdList) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUserListUserRoleIdList) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUserListUserRoleIdList) SetRoleIds(v []*int32) *ListUsersResponseBodyUserListUserRoleIdList {
	s.RoleIds = v
	return s
}

type ListUsersResponseBodyUserListUserRoleNameList struct {
	RoleNames []*string `json:"RoleNames,omitempty" xml:"RoleNames,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUserListUserRoleNameList) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUserListUserRoleNameList) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUserListUserRoleNameList) SetRoleNames(v []*string) *ListUsersResponseBodyUserListUserRoleNameList {
	s.RoleNames = v
	return s
}

type ListUsersResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type SubmitOrderApprovalRequest struct {
	Tid     *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s SubmitOrderApprovalRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitOrderApprovalRequest) GoString() string {
	return s.String()
}

func (s *SubmitOrderApprovalRequest) SetTid(v int64) *SubmitOrderApprovalRequest {
	s.Tid = &v
	return s
}

func (s *SubmitOrderApprovalRequest) SetOrderId(v int64) *SubmitOrderApprovalRequest {
	s.OrderId = &v
	return s
}

type SubmitOrderApprovalResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitOrderApprovalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitOrderApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitOrderApprovalResponseBody) SetRequestId(v string) *SubmitOrderApprovalResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitOrderApprovalResponseBody) SetErrorCode(v string) *SubmitOrderApprovalResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitOrderApprovalResponseBody) SetErrorMessage(v string) *SubmitOrderApprovalResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SubmitOrderApprovalResponseBody) SetSuccess(v bool) *SubmitOrderApprovalResponseBody {
	s.Success = &v
	return s
}

type SubmitOrderApprovalResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitOrderApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitOrderApprovalResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitOrderApprovalResponse) GoString() string {
	return s.String()
}

func (s *SubmitOrderApprovalResponse) SetHeaders(v map[string]*string) *SubmitOrderApprovalResponse {
	s.Headers = v
	return s
}

func (s *SubmitOrderApprovalResponse) SetBody(v *SubmitOrderApprovalResponseBody) *SubmitOrderApprovalResponse {
	s.Body = v
	return s
}

type GrantUserPermissionRequest struct {
	Tid        *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DsType     *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
	DbId       *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Logic      *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	TableId    *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	TableName  *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	PermTypes  *string `json:"PermTypes,omitempty" xml:"PermTypes,omitempty"`
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
}

func (s GrantUserPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantUserPermissionRequest) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionRequest) SetTid(v int64) *GrantUserPermissionRequest {
	s.Tid = &v
	return s
}

func (s *GrantUserPermissionRequest) SetUserId(v string) *GrantUserPermissionRequest {
	s.UserId = &v
	return s
}

func (s *GrantUserPermissionRequest) SetDsType(v string) *GrantUserPermissionRequest {
	s.DsType = &v
	return s
}

func (s *GrantUserPermissionRequest) SetDbId(v string) *GrantUserPermissionRequest {
	s.DbId = &v
	return s
}

func (s *GrantUserPermissionRequest) SetLogic(v bool) *GrantUserPermissionRequest {
	s.Logic = &v
	return s
}

func (s *GrantUserPermissionRequest) SetTableId(v string) *GrantUserPermissionRequest {
	s.TableId = &v
	return s
}

func (s *GrantUserPermissionRequest) SetTableName(v string) *GrantUserPermissionRequest {
	s.TableName = &v
	return s
}

func (s *GrantUserPermissionRequest) SetPermTypes(v string) *GrantUserPermissionRequest {
	s.PermTypes = &v
	return s
}

func (s *GrantUserPermissionRequest) SetExpireDate(v string) *GrantUserPermissionRequest {
	s.ExpireDate = &v
	return s
}

type GrantUserPermissionResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GrantUserPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantUserPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionResponseBody) SetRequestId(v string) *GrantUserPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantUserPermissionResponseBody) SetErrorCode(v string) *GrantUserPermissionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GrantUserPermissionResponseBody) SetErrorMessage(v string) *GrantUserPermissionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GrantUserPermissionResponseBody) SetSuccess(v bool) *GrantUserPermissionResponseBody {
	s.Success = &v
	return s
}

type GrantUserPermissionResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GrantUserPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantUserPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantUserPermissionResponse) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionResponse) SetHeaders(v map[string]*string) *GrantUserPermissionResponse {
	s.Headers = v
	return s
}

func (s *GrantUserPermissionResponse) SetBody(v *GrantUserPermissionResponseBody) *GrantUserPermissionResponse {
	s.Body = v
	return s
}

type GetMetaTableDetailInfoRequest struct {
	Tid       *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s GetMetaTableDetailInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMetaTableDetailInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTableDetailInfoRequest) SetTid(v int64) *GetMetaTableDetailInfoRequest {
	s.Tid = &v
	return s
}

func (s *GetMetaTableDetailInfoRequest) SetTableGuid(v string) *GetMetaTableDetailInfoRequest {
	s.TableGuid = &v
	return s
}

type GetMetaTableDetailInfoResponseBody struct {
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DetailInfo   *GetMetaTableDetailInfoResponseBodyDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Struct"`
	ErrorCode    *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetaTableDetailInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMetaTableDetailInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaTableDetailInfoResponseBody) SetRequestId(v string) *GetMetaTableDetailInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBody) SetDetailInfo(v *GetMetaTableDetailInfoResponseBodyDetailInfo) *GetMetaTableDetailInfoResponseBody {
	s.DetailInfo = v
	return s
}

func (s *GetMetaTableDetailInfoResponseBody) SetErrorCode(v string) *GetMetaTableDetailInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBody) SetErrorMessage(v string) *GetMetaTableDetailInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBody) SetSuccess(v bool) *GetMetaTableDetailInfoResponseBody {
	s.Success = &v
	return s
}

type GetMetaTableDetailInfoResponseBodyDetailInfo struct {
	IndexList  []*GetMetaTableDetailInfoResponseBodyDetailInfoIndexList  `json:"IndexList,omitempty" xml:"IndexList,omitempty" type:"Repeated"`
	ColumnList []*GetMetaTableDetailInfoResponseBodyDetailInfoColumnList `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Repeated"`
}

func (s GetMetaTableDetailInfoResponseBodyDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s GetMetaTableDetailInfoResponseBodyDetailInfo) GoString() string {
	return s.String()
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfo) SetIndexList(v []*GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) *GetMetaTableDetailInfoResponseBodyDetailInfo {
	s.IndexList = v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfo) SetColumnList(v []*GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) *GetMetaTableDetailInfoResponseBodyDetailInfo {
	s.ColumnList = v
	return s
}

type GetMetaTableDetailInfoResponseBodyDetailInfoIndexList struct {
	IndexColumns []*string `json:"IndexColumns,omitempty" xml:"IndexColumns,omitempty" type:"Repeated"`
	IndexName    *string   `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	Unique       *bool     `json:"Unique,omitempty" xml:"Unique,omitempty"`
	IndexType    *string   `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	IndexId      *string   `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
}

func (s GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) String() string {
	return tea.Prettify(s)
}

func (s GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) GoString() string {
	return s.String()
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) SetIndexColumns(v []*string) *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList {
	s.IndexColumns = v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) SetIndexName(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList {
	s.IndexName = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) SetUnique(v bool) *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList {
	s.Unique = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) SetIndexType(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList {
	s.IndexType = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList) SetIndexId(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoIndexList {
	s.IndexId = &v
	return s
}

type GetMetaTableDetailInfoResponseBodyDetailInfoColumnList struct {
	ColumnName    *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DataScale     *int32  `json:"DataScale,omitempty" xml:"DataScale,omitempty"`
	DataPrecision *int32  `json:"DataPrecision,omitempty" xml:"DataPrecision,omitempty"`
	ColumnType    *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	AutoIncrement *bool   `json:"AutoIncrement,omitempty" xml:"AutoIncrement,omitempty"`
	Position      *string `json:"Position,omitempty" xml:"Position,omitempty"`
	Nullable      *bool   `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	ColumnId      *string `json:"ColumnId,omitempty" xml:"ColumnId,omitempty"`
	DataLength    *int64  `json:"DataLength,omitempty" xml:"DataLength,omitempty"`
}

func (s GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) String() string {
	return tea.Prettify(s)
}

func (s GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) GoString() string {
	return s.String()
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetColumnName(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.ColumnName = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetDescription(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.Description = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetDataScale(v int32) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.DataScale = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetDataPrecision(v int32) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.DataPrecision = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetColumnType(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.ColumnType = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetAutoIncrement(v bool) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.AutoIncrement = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetPosition(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.Position = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetNullable(v bool) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.Nullable = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetColumnId(v string) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.ColumnId = &v
	return s
}

func (s *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList) SetDataLength(v int64) *GetMetaTableDetailInfoResponseBodyDetailInfoColumnList {
	s.DataLength = &v
	return s
}

type GetMetaTableDetailInfoResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMetaTableDetailInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMetaTableDetailInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMetaTableDetailInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMetaTableDetailInfoResponse) SetHeaders(v map[string]*string) *GetMetaTableDetailInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMetaTableDetailInfoResponse) SetBody(v *GetMetaTableDetailInfoResponseBody) *GetMetaTableDetailInfoResponse {
	s.Body = v
	return s
}

type GetDataCorrectSQLFileRequest struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Tid     *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataCorrectSQLFileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectSQLFileRequest) GoString() string {
	return s.String()
}

func (s *GetDataCorrectSQLFileRequest) SetOrderId(v int64) *GetDataCorrectSQLFileRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataCorrectSQLFileRequest) SetTid(v int64) *GetDataCorrectSQLFileRequest {
	s.Tid = &v
	return s
}

type GetDataCorrectSQLFileResponseBody struct {
	// Id of the request
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s GetDataCorrectSQLFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectSQLFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCorrectSQLFileResponseBody) SetRequestId(v string) *GetDataCorrectSQLFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCorrectSQLFileResponseBody) SetSuccess(v bool) *GetDataCorrectSQLFileResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCorrectSQLFileResponseBody) SetErrorMessage(v string) *GetDataCorrectSQLFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataCorrectSQLFileResponseBody) SetErrorCode(v string) *GetDataCorrectSQLFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataCorrectSQLFileResponseBody) SetFileUrl(v string) *GetDataCorrectSQLFileResponseBody {
	s.FileUrl = &v
	return s
}

type GetDataCorrectSQLFileResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDataCorrectSQLFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDataCorrectSQLFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectSQLFileResponse) GoString() string {
	return s.String()
}

func (s *GetDataCorrectSQLFileResponse) SetHeaders(v map[string]*string) *GetDataCorrectSQLFileResponse {
	s.Headers = v
	return s
}

func (s *GetDataCorrectSQLFileResponse) SetBody(v *GetDataCorrectSQLFileResponseBody) *GetDataCorrectSQLFileResponse {
	s.Body = v
	return s
}

type CreateFreeLockCorrectOrderRequest struct {
	Comment         *string                                 `json:"Comment,omitempty" xml:"Comment,omitempty"`
	RelatedUserList []*int64                                `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	Param           *CreateFreeLockCorrectOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	AttachmentKey   *string                                 `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	Tid             *int64                                  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateFreeLockCorrectOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFreeLockCorrectOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateFreeLockCorrectOrderRequest) SetComment(v string) *CreateFreeLockCorrectOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequest) SetRelatedUserList(v []*int64) *CreateFreeLockCorrectOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateFreeLockCorrectOrderRequest) SetParam(v *CreateFreeLockCorrectOrderRequestParam) *CreateFreeLockCorrectOrderRequest {
	s.Param = v
	return s
}

func (s *CreateFreeLockCorrectOrderRequest) SetAttachmentKey(v string) *CreateFreeLockCorrectOrderRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequest) SetTid(v int64) *CreateFreeLockCorrectOrderRequest {
	s.Tid = &v
	return s
}

type CreateFreeLockCorrectOrderRequestParam struct {
	Classify               *string                                             `json:"Classify,omitempty" xml:"Classify,omitempty"`
	ExecSQL                *string                                             `json:"ExecSQL,omitempty" xml:"ExecSQL,omitempty"`
	SqlType                *string                                             `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	AttachmentName         *string                                             `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	RollbackSQL            *string                                             `json:"RollbackSQL,omitempty" xml:"RollbackSQL,omitempty"`
	RollbackAttachmentName *string                                             `json:"RollbackAttachmentName,omitempty" xml:"RollbackAttachmentName,omitempty"`
	RollbackSqlType        *string                                             `json:"RollbackSqlType,omitempty" xml:"RollbackSqlType,omitempty"`
	DbItemList             []*CreateFreeLockCorrectOrderRequestParamDbItemList `json:"DbItemList,omitempty" xml:"DbItemList,omitempty" type:"Repeated"`
}

func (s CreateFreeLockCorrectOrderRequestParam) String() string {
	return tea.Prettify(s)
}

func (s CreateFreeLockCorrectOrderRequestParam) GoString() string {
	return s.String()
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetClassify(v string) *CreateFreeLockCorrectOrderRequestParam {
	s.Classify = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetExecSQL(v string) *CreateFreeLockCorrectOrderRequestParam {
	s.ExecSQL = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetSqlType(v string) *CreateFreeLockCorrectOrderRequestParam {
	s.SqlType = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetAttachmentName(v string) *CreateFreeLockCorrectOrderRequestParam {
	s.AttachmentName = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetRollbackSQL(v string) *CreateFreeLockCorrectOrderRequestParam {
	s.RollbackSQL = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetRollbackAttachmentName(v string) *CreateFreeLockCorrectOrderRequestParam {
	s.RollbackAttachmentName = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetRollbackSqlType(v string) *CreateFreeLockCorrectOrderRequestParam {
	s.RollbackSqlType = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParam) SetDbItemList(v []*CreateFreeLockCorrectOrderRequestParamDbItemList) *CreateFreeLockCorrectOrderRequestParam {
	s.DbItemList = v
	return s
}

type CreateFreeLockCorrectOrderRequestParamDbItemList struct {
	DbId  *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Logic *bool  `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s CreateFreeLockCorrectOrderRequestParamDbItemList) String() string {
	return tea.Prettify(s)
}

func (s CreateFreeLockCorrectOrderRequestParamDbItemList) GoString() string {
	return s.String()
}

func (s *CreateFreeLockCorrectOrderRequestParamDbItemList) SetDbId(v int64) *CreateFreeLockCorrectOrderRequestParamDbItemList {
	s.DbId = &v
	return s
}

func (s *CreateFreeLockCorrectOrderRequestParamDbItemList) SetLogic(v bool) *CreateFreeLockCorrectOrderRequestParamDbItemList {
	s.Logic = &v
	return s
}

type CreateFreeLockCorrectOrderShrinkRequest struct {
	Comment               *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	RelatedUserListShrink *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	ParamShrink           *string `json:"Param,omitempty" xml:"Param,omitempty"`
	AttachmentKey         *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	Tid                   *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateFreeLockCorrectOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFreeLockCorrectOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) SetComment(v string) *CreateFreeLockCorrectOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateFreeLockCorrectOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) SetParamShrink(v string) *CreateFreeLockCorrectOrderShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) SetAttachmentKey(v string) *CreateFreeLockCorrectOrderShrinkRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) SetTid(v int64) *CreateFreeLockCorrectOrderShrinkRequest {
	s.Tid = &v
	return s
}

type CreateFreeLockCorrectOrderResponseBody struct {
	// Id of the request
	RequestId         *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success           *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage      *string  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode         *string  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	CreateOrderResult []*int64 `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Repeated"`
}

func (s CreateFreeLockCorrectOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFreeLockCorrectOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFreeLockCorrectOrderResponseBody) SetRequestId(v string) *CreateFreeLockCorrectOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFreeLockCorrectOrderResponseBody) SetSuccess(v bool) *CreateFreeLockCorrectOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFreeLockCorrectOrderResponseBody) SetErrorMessage(v string) *CreateFreeLockCorrectOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateFreeLockCorrectOrderResponseBody) SetErrorCode(v string) *CreateFreeLockCorrectOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFreeLockCorrectOrderResponseBody) SetCreateOrderResult(v []*int64) *CreateFreeLockCorrectOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

type CreateFreeLockCorrectOrderResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFreeLockCorrectOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFreeLockCorrectOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFreeLockCorrectOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateFreeLockCorrectOrderResponse) SetHeaders(v map[string]*string) *CreateFreeLockCorrectOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateFreeLockCorrectOrderResponse) SetBody(v *CreateFreeLockCorrectOrderResponseBody) *CreateFreeLockCorrectOrderResponse {
	s.Body = v
	return s
}

type CreateOrderRequest struct {
	Tid             *int64                 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	Comment         *string                `json:"Comment,omitempty" xml:"Comment,omitempty"`
	PluginParam     map[string]interface{} `json:"PluginParam,omitempty" xml:"PluginParam,omitempty"`
	RelatedUserList *string                `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	PluginType      *string                `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	AttachmentKey   *string                `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) SetTid(v int64) *CreateOrderRequest {
	s.Tid = &v
	return s
}

func (s *CreateOrderRequest) SetComment(v string) *CreateOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateOrderRequest) SetPluginParam(v map[string]interface{}) *CreateOrderRequest {
	s.PluginParam = v
	return s
}

func (s *CreateOrderRequest) SetRelatedUserList(v string) *CreateOrderRequest {
	s.RelatedUserList = &v
	return s
}

func (s *CreateOrderRequest) SetPluginType(v string) *CreateOrderRequest {
	s.PluginType = &v
	return s
}

func (s *CreateOrderRequest) SetAttachmentKey(v string) *CreateOrderRequest {
	s.AttachmentKey = &v
	return s
}

type CreateOrderShrinkRequest struct {
	Tid               *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	Comment           *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	PluginParamShrink *string `json:"PluginParam,omitempty" xml:"PluginParam,omitempty"`
	RelatedUserList   *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	PluginType        *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	AttachmentKey     *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
}

func (s CreateOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderShrinkRequest) SetTid(v int64) *CreateOrderShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetComment(v string) *CreateOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetPluginParamShrink(v string) *CreateOrderShrinkRequest {
	s.PluginParamShrink = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetRelatedUserList(v string) *CreateOrderShrinkRequest {
	s.RelatedUserList = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetPluginType(v string) *CreateOrderShrinkRequest {
	s.PluginType = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetAttachmentKey(v string) *CreateOrderShrinkRequest {
	s.AttachmentKey = &v
	return s
}

type CreateOrderResponseBody struct {
	CreateOrderResult *CreateOrderResponseBodyCreateOrderResult `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Struct"`
	RequestId         *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode         *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage      *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success           *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) SetCreateOrderResult(v *CreateOrderResponseBodyCreateOrderResult) *CreateOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

func (s *CreateOrderResponseBody) SetRequestId(v string) *CreateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderResponseBody) SetErrorCode(v string) *CreateOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateOrderResponseBody) SetErrorMessage(v string) *CreateOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateOrderResponseBody) SetSuccess(v bool) *CreateOrderResponseBody {
	s.Success = &v
	return s
}

type CreateOrderResponseBodyCreateOrderResult struct {
	OrderIds []*int64 `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
}

func (s CreateOrderResponseBodyCreateOrderResult) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBodyCreateOrderResult) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBodyCreateOrderResult) SetOrderIds(v []*int64) *CreateOrderResponseBodyCreateOrderResult {
	s.OrderIds = v
	return s
}

type CreateOrderResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderResponse) SetHeaders(v map[string]*string) *CreateOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderResponse) SetBody(v *CreateOrderResponseBody) *CreateOrderResponse {
	s.Body = v
	return s
}

type ListDatabasesRequest struct {
	Tid        *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDatabasesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatabasesRequest) GoString() string {
	return s.String()
}

func (s *ListDatabasesRequest) SetTid(v int64) *ListDatabasesRequest {
	s.Tid = &v
	return s
}

func (s *ListDatabasesRequest) SetInstanceId(v string) *ListDatabasesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDatabasesRequest) SetPageNumber(v int32) *ListDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatabasesRequest) SetPageSize(v int32) *ListDatabasesRequest {
	s.PageSize = &v
	return s
}

type ListDatabasesResponseBody struct {
	TotalCount   *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DatabaseList *ListDatabasesResponseBodyDatabaseList `json:"DatabaseList,omitempty" xml:"DatabaseList,omitempty" type:"Struct"`
	ErrorCode    *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDatabasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBody) SetTotalCount(v int64) *ListDatabasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatabasesResponseBody) SetRequestId(v string) *ListDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabasesResponseBody) SetDatabaseList(v *ListDatabasesResponseBodyDatabaseList) *ListDatabasesResponseBody {
	s.DatabaseList = v
	return s
}

func (s *ListDatabasesResponseBody) SetErrorCode(v string) *ListDatabasesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDatabasesResponseBody) SetErrorMessage(v string) *ListDatabasesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDatabasesResponseBody) SetSuccess(v bool) *ListDatabasesResponseBody {
	s.Success = &v
	return s
}

type ListDatabasesResponseBodyDatabaseList struct {
	Database []*ListDatabasesResponseBodyDatabaseListDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Repeated"`
}

func (s ListDatabasesResponseBodyDatabaseList) String() string {
	return tea.Prettify(s)
}

func (s ListDatabasesResponseBodyDatabaseList) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyDatabaseList) SetDatabase(v []*ListDatabasesResponseBodyDatabaseListDatabase) *ListDatabasesResponseBodyDatabaseList {
	s.Database = v
	return s
}

type ListDatabasesResponseBodyDatabaseListDatabase struct {
	DatabaseId    *string                                                     `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	Host          *string                                                     `json:"Host,omitempty" xml:"Host,omitempty"`
	CatalogName   *string                                                     `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	DbaName       *string                                                     `json:"DbaName,omitempty" xml:"DbaName,omitempty"`
	State         *string                                                     `json:"State,omitempty" xml:"State,omitempty"`
	DbaId         *string                                                     `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	SchemaName    *string                                                     `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	InstanceId    *string                                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Port          *int32                                                      `json:"Port,omitempty" xml:"Port,omitempty"`
	EnvType       *string                                                     `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Sid           *string                                                     `json:"Sid,omitempty" xml:"Sid,omitempty"`
	OwnerIdList   *ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList   `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	Encoding      *string                                                     `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	DbType        *string                                                     `json:"DbType,omitempty" xml:"DbType,omitempty"`
	OwnerNameList *ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	SearchName    *string                                                     `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s ListDatabasesResponseBodyDatabaseListDatabase) String() string {
	return tea.Prettify(s)
}

func (s ListDatabasesResponseBodyDatabaseListDatabase) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetDatabaseId(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.DatabaseId = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetHost(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.Host = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetCatalogName(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.CatalogName = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetDbaName(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.DbaName = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetState(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.State = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetDbaId(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.DbaId = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetSchemaName(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.SchemaName = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetInstanceId(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.InstanceId = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetPort(v int32) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.Port = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetEnvType(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.EnvType = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetSid(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.Sid = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetOwnerIdList(v *ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.OwnerIdList = v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetEncoding(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.Encoding = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetDbType(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.DbType = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetOwnerNameList(v *ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.OwnerNameList = v
	return s
}

func (s *ListDatabasesResponseBodyDatabaseListDatabase) SetSearchName(v string) *ListDatabasesResponseBodyDatabaseListDatabase {
	s.SearchName = &v
	return s
}

type ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList) String() string {
	return tea.Prettify(s)
}

func (s ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList) SetOwnerIds(v []*string) *ListDatabasesResponseBodyDatabaseListDatabaseOwnerIdList {
	s.OwnerIds = v
	return s
}

type ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList) String() string {
	return tea.Prettify(s)
}

func (s ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList) SetOwnerNames(v []*string) *ListDatabasesResponseBodyDatabaseListDatabaseOwnerNameList {
	s.OwnerNames = v
	return s
}

type ListDatabasesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDatabasesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDatabasesResponse) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponse) SetHeaders(v map[string]*string) *ListDatabasesResponse {
	s.Headers = v
	return s
}

func (s *ListDatabasesResponse) SetBody(v *ListDatabasesResponseBody) *ListDatabasesResponse {
	s.Body = v
	return s
}

type ListUserPermissionsRequest struct {
	Tid          *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	PermType     *string `json:"PermType,omitempty" xml:"PermType,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Logic        *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	EnvType      *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	DbType       *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListUserPermissionsRequest) SetTid(v int64) *ListUserPermissionsRequest {
	s.Tid = &v
	return s
}

func (s *ListUserPermissionsRequest) SetPermType(v string) *ListUserPermissionsRequest {
	s.PermType = &v
	return s
}

func (s *ListUserPermissionsRequest) SetUserId(v string) *ListUserPermissionsRequest {
	s.UserId = &v
	return s
}

func (s *ListUserPermissionsRequest) SetDatabaseName(v string) *ListUserPermissionsRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListUserPermissionsRequest) SetLogic(v bool) *ListUserPermissionsRequest {
	s.Logic = &v
	return s
}

func (s *ListUserPermissionsRequest) SetEnvType(v string) *ListUserPermissionsRequest {
	s.EnvType = &v
	return s
}

func (s *ListUserPermissionsRequest) SetDbType(v string) *ListUserPermissionsRequest {
	s.DbType = &v
	return s
}

func (s *ListUserPermissionsRequest) SetPageNumber(v int32) *ListUserPermissionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserPermissionsRequest) SetPageSize(v int32) *ListUserPermissionsRequest {
	s.PageSize = &v
	return s
}

type ListUserPermissionsResponseBody struct {
	TotalCount      *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode       *string                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage    *string                                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	UserPermissions *ListUserPermissionsResponseBodyUserPermissions `json:"UserPermissions,omitempty" xml:"UserPermissions,omitempty" type:"Struct"`
	Success         *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUserPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserPermissionsResponseBody) SetTotalCount(v int64) *ListUserPermissionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserPermissionsResponseBody) SetRequestId(v string) *ListUserPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserPermissionsResponseBody) SetErrorCode(v string) *ListUserPermissionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUserPermissionsResponseBody) SetErrorMessage(v string) *ListUserPermissionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListUserPermissionsResponseBody) SetUserPermissions(v *ListUserPermissionsResponseBodyUserPermissions) *ListUserPermissionsResponseBody {
	s.UserPermissions = v
	return s
}

func (s *ListUserPermissionsResponseBody) SetSuccess(v bool) *ListUserPermissionsResponseBody {
	s.Success = &v
	return s
}

type ListUserPermissionsResponseBodyUserPermissions struct {
	UserPermission []*ListUserPermissionsResponseBodyUserPermissionsUserPermission `json:"UserPermission,omitempty" xml:"UserPermission,omitempty" type:"Repeated"`
}

func (s ListUserPermissionsResponseBodyUserPermissions) String() string {
	return tea.Prettify(s)
}

func (s ListUserPermissionsResponseBodyUserPermissions) GoString() string {
	return s.String()
}

func (s *ListUserPermissionsResponseBodyUserPermissions) SetUserPermission(v []*ListUserPermissionsResponseBodyUserPermissionsUserPermission) *ListUserPermissionsResponseBodyUserPermissions {
	s.UserPermission = v
	return s
}

type ListUserPermissionsResponseBodyUserPermissionsUserPermission struct {
	DbId         *string                                                                  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	TableName    *string                                                                  `json:"TableName,omitempty" xml:"TableName,omitempty"`
	UserId       *string                                                                  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	SchemaName   *string                                                                  `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	Logic        *bool                                                                    `json:"Logic,omitempty" xml:"Logic,omitempty"`
	UserNickName *string                                                                  `json:"UserNickName,omitempty" xml:"UserNickName,omitempty"`
	InstanceId   *string                                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PermDetails  *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails `json:"PermDetails,omitempty" xml:"PermDetails,omitempty" type:"Struct"`
	EnvType      *string                                                                  `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	ColumnName   *string                                                                  `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	DbType       *string                                                                  `json:"DbType,omitempty" xml:"DbType,omitempty"`
	DsType       *string                                                                  `json:"DsType,omitempty" xml:"DsType,omitempty"`
	TableId      *string                                                                  `json:"TableId,omitempty" xml:"TableId,omitempty"`
	SearchName   *string                                                                  `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	Alias        *string                                                                  `json:"Alias,omitempty" xml:"Alias,omitempty"`
}

func (s ListUserPermissionsResponseBodyUserPermissionsUserPermission) String() string {
	return tea.Prettify(s)
}

func (s ListUserPermissionsResponseBodyUserPermissionsUserPermission) GoString() string {
	return s.String()
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetDbId(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.DbId = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetTableName(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.TableName = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetUserId(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.UserId = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetSchemaName(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.SchemaName = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetLogic(v bool) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.Logic = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetUserNickName(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.UserNickName = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetInstanceId(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.InstanceId = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetPermDetails(v *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.PermDetails = v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetEnvType(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.EnvType = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetColumnName(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.ColumnName = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetDbType(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.DbType = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetDsType(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.DsType = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetTableId(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.TableId = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetSearchName(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.SearchName = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermission) SetAlias(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.Alias = &v
	return s
}

type ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails struct {
	PermDetail []*ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail `json:"PermDetail,omitempty" xml:"PermDetail,omitempty" type:"Repeated"`
}

func (s ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) String() string {
	return tea.Prettify(s)
}

func (s ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) GoString() string {
	return s.String()
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) SetPermDetail(v []*ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails {
	s.PermDetail = v
	return s
}

type ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail struct {
	OriginFrom   *string `json:"OriginFrom,omitempty" xml:"OriginFrom,omitempty"`
	PermType     *string `json:"PermType,omitempty" xml:"PermType,omitempty"`
	ExpireDate   *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	CreateDate   *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	UserAccessId *string `json:"UserAccessId,omitempty" xml:"UserAccessId,omitempty"`
	ExtraData    *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
}

func (s ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) String() string {
	return tea.Prettify(s)
}

func (s ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GoString() string {
	return s.String()
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetOriginFrom(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.OriginFrom = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetPermType(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.PermType = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetExpireDate(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.ExpireDate = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetCreateDate(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.CreateDate = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetUserAccessId(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.UserAccessId = &v
	return s
}

func (s *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetExtraData(v string) *ListUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.ExtraData = &v
	return s
}

type ListUserPermissionsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListUserPermissionsResponse) SetHeaders(v map[string]*string) *ListUserPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListUserPermissionsResponse) SetBody(v *ListUserPermissionsResponseBody) *ListUserPermissionsResponse {
	s.Body = v
	return s
}

type ListWorkFlowTemplatesRequest struct {
	Tid        *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s ListWorkFlowTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkFlowTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkFlowTemplatesRequest) SetTid(v int64) *ListWorkFlowTemplatesRequest {
	s.Tid = &v
	return s
}

func (s *ListWorkFlowTemplatesRequest) SetSearchName(v string) *ListWorkFlowTemplatesRequest {
	s.SearchName = &v
	return s
}

type ListWorkFlowTemplatesResponseBody struct {
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode         *string                                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage      *string                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success           *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
	WorkFlowTemplates *ListWorkFlowTemplatesResponseBodyWorkFlowTemplates `json:"WorkFlowTemplates,omitempty" xml:"WorkFlowTemplates,omitempty" type:"Struct"`
}

func (s ListWorkFlowTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkFlowTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkFlowTemplatesResponseBody) SetRequestId(v string) *ListWorkFlowTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBody) SetErrorCode(v string) *ListWorkFlowTemplatesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBody) SetErrorMessage(v string) *ListWorkFlowTemplatesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBody) SetSuccess(v bool) *ListWorkFlowTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBody) SetWorkFlowTemplates(v *ListWorkFlowTemplatesResponseBodyWorkFlowTemplates) *ListWorkFlowTemplatesResponseBody {
	s.WorkFlowTemplates = v
	return s
}

type ListWorkFlowTemplatesResponseBodyWorkFlowTemplates struct {
	WorkFlowTemplate []*ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate `json:"WorkFlowTemplate,omitempty" xml:"WorkFlowTemplate,omitempty" type:"Repeated"`
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplates) GoString() string {
	return s.String()
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplates) SetWorkFlowTemplate(v []*ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplates {
	s.WorkFlowTemplate = v
	return s
}

type ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate struct {
	IsSystem      *int32                                                                           `json:"IsSystem,omitempty" xml:"IsSystem,omitempty"`
	WorkflowNodes *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes `json:"WorkflowNodes,omitempty" xml:"WorkflowNodes,omitempty" type:"Struct"`
	Comment       *string                                                                          `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Enabled       *string                                                                          `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	TemplateName  *string                                                                          `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateId    *int64                                                                           `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	CreateUserId  *int64                                                                           `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) String() string {
	return tea.Prettify(s)
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) GoString() string {
	return s.String()
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) SetIsSystem(v int32) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate {
	s.IsSystem = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) SetWorkflowNodes(v *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate {
	s.WorkflowNodes = v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) SetComment(v string) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate {
	s.Comment = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) SetEnabled(v string) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate {
	s.Enabled = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) SetTemplateName(v string) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate {
	s.TemplateName = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) SetTemplateId(v int64) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate {
	s.TemplateId = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate) SetCreateUserId(v int64) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplate {
	s.CreateUserId = &v
	return s
}

type ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes struct {
	WorkflowNode []*ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode `json:"WorkflowNode,omitempty" xml:"WorkflowNode,omitempty" type:"Repeated"`
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes) String() string {
	return tea.Prettify(s)
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes) GoString() string {
	return s.String()
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes) SetWorkflowNode(v []*ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodes {
	s.WorkflowNode = v
	return s
}

type ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode struct {
	Comment      *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	NodeName     *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	Position     *int32  `json:"Position,omitempty" xml:"Position,omitempty"`
	CreateUserId *int64  `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	NodeId       *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) String() string {
	return tea.Prettify(s)
}

func (s ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) GoString() string {
	return s.String()
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) SetComment(v string) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode {
	s.Comment = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) SetNodeType(v string) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode {
	s.NodeType = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) SetNodeName(v string) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode {
	s.NodeName = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) SetPosition(v int32) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode {
	s.Position = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) SetCreateUserId(v int64) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode {
	s.CreateUserId = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) SetTemplateId(v int64) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode {
	s.TemplateId = &v
	return s
}

func (s *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode) SetNodeId(v int64) *ListWorkFlowTemplatesResponseBodyWorkFlowTemplatesWorkFlowTemplateWorkflowNodesWorkflowNode {
	s.NodeId = &v
	return s
}

type ListWorkFlowTemplatesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListWorkFlowTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWorkFlowTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkFlowTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkFlowTemplatesResponse) SetHeaders(v map[string]*string) *ListWorkFlowTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkFlowTemplatesResponse) SetBody(v *ListWorkFlowTemplatesResponseBody) *ListWorkFlowTemplatesResponse {
	s.Body = v
	return s
}

type GetDataExportOrderDetailRequest struct {
	Tid     *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s GetDataExportOrderDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataExportOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDataExportOrderDetailRequest) SetTid(v int64) *GetDataExportOrderDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetDataExportOrderDetailRequest) SetOrderId(v int64) *GetDataExportOrderDetailRequest {
	s.OrderId = &v
	return s
}

type GetDataExportOrderDetailResponseBody struct {
	RequestId             *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DataExportOrderDetail *GetDataExportOrderDetailResponseBodyDataExportOrderDetail `json:"DataExportOrderDetail,omitempty" xml:"DataExportOrderDetail,omitempty" type:"Struct"`
	ErrorCode             *string                                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage          *string                                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success               *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataExportOrderDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataExportOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataExportOrderDetailResponseBody) SetRequestId(v string) *GetDataExportOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBody) SetDataExportOrderDetail(v *GetDataExportOrderDetailResponseBodyDataExportOrderDetail) *GetDataExportOrderDetailResponseBody {
	s.DataExportOrderDetail = v
	return s
}

func (s *GetDataExportOrderDetailResponseBody) SetErrorCode(v string) *GetDataExportOrderDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBody) SetErrorMessage(v string) *GetDataExportOrderDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBody) SetSuccess(v bool) *GetDataExportOrderDetailResponseBody {
	s.Success = &v
	return s
}

type GetDataExportOrderDetailResponseBodyDataExportOrderDetail struct {
	KeyInfo     *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo     `json:"KeyInfo,omitempty" xml:"KeyInfo,omitempty" type:"Struct"`
	OrderDetail *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail `json:"OrderDetail,omitempty" xml:"OrderDetail,omitempty" type:"Struct"`
}

func (s GetDataExportOrderDetailResponseBodyDataExportOrderDetail) String() string {
	return tea.Prettify(s)
}

func (s GetDataExportOrderDetailResponseBodyDataExportOrderDetail) GoString() string {
	return s.String()
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetail) SetKeyInfo(v *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo) *GetDataExportOrderDetailResponseBodyDataExportOrderDetail {
	s.KeyInfo = v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetail) SetOrderDetail(v *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) *GetDataExportOrderDetailResponseBodyDataExportOrderDetail {
	s.OrderDetail = v
	return s
}

type GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo struct {
	PreCheckId *int64  `json:"PreCheckId,omitempty" xml:"PreCheckId,omitempty"`
	JobStatus  *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
}

func (s GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo) GoString() string {
	return s.String()
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo) SetPreCheckId(v int64) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo {
	s.PreCheckId = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo) SetJobStatus(v string) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailKeyInfo {
	s.JobStatus = &v
	return s
}

type GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail struct {
	DbId                   *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Database               *string `json:"Database,omitempty" xml:"Database,omitempty"`
	Classify               *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	ExeSQL                 *string `json:"ExeSQL,omitempty" xml:"ExeSQL,omitempty"`
	Logic                  *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	ActualAffectRows       *int64  `json:"ActualAffectRows,omitempty" xml:"ActualAffectRows,omitempty"`
	IgnoreAffectRows       *bool   `json:"IgnoreAffectRows,omitempty" xml:"IgnoreAffectRows,omitempty"`
	IgnoreAffectRowsReason *string `json:"IgnoreAffectRowsReason,omitempty" xml:"IgnoreAffectRowsReason,omitempty"`
	EnvType                *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
}

func (s GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) String() string {
	return tea.Prettify(s)
}

func (s GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) GoString() string {
	return s.String()
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetDbId(v int32) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.DbId = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetDatabase(v string) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.Database = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetClassify(v string) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.Classify = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetExeSQL(v string) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.ExeSQL = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetLogic(v bool) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.Logic = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetActualAffectRows(v int64) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.ActualAffectRows = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetIgnoreAffectRows(v bool) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.IgnoreAffectRows = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetIgnoreAffectRowsReason(v string) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.IgnoreAffectRowsReason = &v
	return s
}

func (s *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail) SetEnvType(v string) *GetDataExportOrderDetailResponseBodyDataExportOrderDetailOrderDetail {
	s.EnvType = &v
	return s
}

type GetDataExportOrderDetailResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDataExportOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDataExportOrderDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataExportOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDataExportOrderDetailResponse) SetHeaders(v map[string]*string) *GetDataExportOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDataExportOrderDetailResponse) SetBody(v *GetDataExportOrderDetailResponseBody) *GetDataExportOrderDetailResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	Tid            *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	SearchKey      *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	DbType         *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	EnvType        *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	NetType        *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	InstanceState  *string `json:"InstanceState,omitempty" xml:"InstanceState,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetTid(v int64) *ListInstancesRequest {
	s.Tid = &v
	return s
}

func (s *ListInstancesRequest) SetSearchKey(v string) *ListInstancesRequest {
	s.SearchKey = &v
	return s
}

func (s *ListInstancesRequest) SetDbType(v string) *ListInstancesRequest {
	s.DbType = &v
	return s
}

func (s *ListInstancesRequest) SetEnvType(v string) *ListInstancesRequest {
	s.EnvType = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceSource(v string) *ListInstancesRequest {
	s.InstanceSource = &v
	return s
}

func (s *ListInstancesRequest) SetNetType(v string) *ListInstancesRequest {
	s.NetType = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceState(v string) *ListInstancesRequest {
	s.InstanceState = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

type ListInstancesResponseBody struct {
	TotalCount   *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	InstanceList *ListInstancesResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetTotalCount(v int64) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetErrorCode(v string) *ListInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetErrorMessage(v string) *ListInstancesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListInstancesResponseBody) SetInstanceList(v *ListInstancesResponseBodyInstanceList) *ListInstancesResponseBody {
	s.InstanceList = v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v bool) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

type ListInstancesResponseBodyInstanceList struct {
	Instance []*ListInstancesResponseBodyInstanceListInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceList) SetInstance(v []*ListInstancesResponseBodyInstanceListInstance) *ListInstancesResponseBodyInstanceList {
	s.Instance = v
	return s
}

type ListInstancesResponseBodyInstanceListInstance struct {
	VpcId            *string                                                     `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	DatabaseUser     *string                                                     `json:"DatabaseUser,omitempty" xml:"DatabaseUser,omitempty"`
	DbaId            *string                                                     `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	UseDsql          *int32                                                      `json:"UseDsql,omitempty" xml:"UseDsql,omitempty"`
	Port             *int32                                                      `json:"Port,omitempty" xml:"Port,omitempty"`
	EcsInstanceId    *string                                                     `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	EnvType          *string                                                     `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Sid              *string                                                     `json:"Sid,omitempty" xml:"Sid,omitempty"`
	OwnerIdList      *ListInstancesResponseBodyInstanceListInstanceOwnerIdList   `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	SafeRuleId       *string                                                     `json:"SafeRuleId,omitempty" xml:"SafeRuleId,omitempty"`
	OwnerNameList    *ListInstancesResponseBodyInstanceListInstanceOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	DbaNickName      *string                                                     `json:"DbaNickName,omitempty" xml:"DbaNickName,omitempty"`
	QueryTimeout     *int32                                                      `json:"QueryTimeout,omitempty" xml:"QueryTimeout,omitempty"`
	InstanceSource   *string                                                     `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	Host             *string                                                     `json:"Host,omitempty" xml:"Host,omitempty"`
	State            *string                                                     `json:"State,omitempty" xml:"State,omitempty"`
	DataLinkName     *string                                                     `json:"DataLinkName,omitempty" xml:"DataLinkName,omitempty"`
	ExportTimeout    *int32                                                      `json:"ExportTimeout,omitempty" xml:"ExportTimeout,omitempty"`
	InstanceId       *string                                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType     *string                                                     `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	DatabasePassword *string                                                     `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	InstanceAlias    *string                                                     `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	DdlOnline        *int32                                                      `json:"DdlOnline,omitempty" xml:"DdlOnline,omitempty"`
	EcsRegion        *string                                                     `json:"EcsRegion,omitempty" xml:"EcsRegion,omitempty"`
}

func (s ListInstancesResponseBodyInstanceListInstance) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceListInstance) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetVpcId(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.VpcId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetDatabaseUser(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.DatabaseUser = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetDbaId(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.DbaId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetUseDsql(v int32) *ListInstancesResponseBodyInstanceListInstance {
	s.UseDsql = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetPort(v int32) *ListInstancesResponseBodyInstanceListInstance {
	s.Port = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetEcsInstanceId(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.EcsInstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetEnvType(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.EnvType = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetSid(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.Sid = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetOwnerIdList(v *ListInstancesResponseBodyInstanceListInstanceOwnerIdList) *ListInstancesResponseBodyInstanceListInstance {
	s.OwnerIdList = v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetSafeRuleId(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.SafeRuleId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetOwnerNameList(v *ListInstancesResponseBodyInstanceListInstanceOwnerNameList) *ListInstancesResponseBodyInstanceListInstance {
	s.OwnerNameList = v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetDbaNickName(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.DbaNickName = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetQueryTimeout(v int32) *ListInstancesResponseBodyInstanceListInstance {
	s.QueryTimeout = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetInstanceSource(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.InstanceSource = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetHost(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.Host = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetState(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.State = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetDataLinkName(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.DataLinkName = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetExportTimeout(v int32) *ListInstancesResponseBodyInstanceListInstance {
	s.ExportTimeout = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetInstanceId(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetInstanceType(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetDatabasePassword(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.DatabasePassword = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetInstanceAlias(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.InstanceAlias = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetDdlOnline(v int32) *ListInstancesResponseBodyInstanceListInstance {
	s.DdlOnline = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListInstance) SetEcsRegion(v string) *ListInstancesResponseBodyInstanceListInstance {
	s.EcsRegion = &v
	return s
}

type ListInstancesResponseBodyInstanceListInstanceOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyInstanceListInstanceOwnerIdList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceListInstanceOwnerIdList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceListInstanceOwnerIdList) SetOwnerIds(v []*string) *ListInstancesResponseBodyInstanceListInstanceOwnerIdList {
	s.OwnerIds = v
	return s
}

type ListInstancesResponseBodyInstanceListInstanceOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyInstanceListInstanceOwnerNameList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceListInstanceOwnerNameList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceListInstanceOwnerNameList) SetOwnerNames(v []*string) *ListInstancesResponseBodyInstanceListInstanceOwnerNameList {
	s.OwnerNames = v
	return s
}

type ListInstancesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type GetUserUploadFileJobRequest struct {
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	Tid    *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetUserUploadFileJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserUploadFileJobRequest) GoString() string {
	return s.String()
}

func (s *GetUserUploadFileJobRequest) SetJobKey(v string) *GetUserUploadFileJobRequest {
	s.JobKey = &v
	return s
}

func (s *GetUserUploadFileJobRequest) SetTid(v int64) *GetUserUploadFileJobRequest {
	s.Tid = &v
	return s
}

type GetUserUploadFileJobResponseBody struct {
	// Id of the request
	RequestId           *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success             *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage        *string                                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode           *string                                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	UploadFileJobDetail *GetUserUploadFileJobResponseBodyUploadFileJobDetail `json:"UploadFileJobDetail,omitempty" xml:"UploadFileJobDetail,omitempty" type:"Struct"`
}

func (s GetUserUploadFileJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserUploadFileJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserUploadFileJobResponseBody) SetRequestId(v string) *GetUserUploadFileJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserUploadFileJobResponseBody) SetSuccess(v bool) *GetUserUploadFileJobResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserUploadFileJobResponseBody) SetErrorMessage(v string) *GetUserUploadFileJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserUploadFileJobResponseBody) SetErrorCode(v string) *GetUserUploadFileJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserUploadFileJobResponseBody) SetUploadFileJobDetail(v *GetUserUploadFileJobResponseBodyUploadFileJobDetail) *GetUserUploadFileJobResponseBody {
	s.UploadFileJobDetail = v
	return s
}

type GetUserUploadFileJobResponseBodyUploadFileJobDetail struct {
	JobKey         *string                                                            `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	FileName       *string                                                            `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize       *int64                                                             `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileSource     *string                                                            `json:"FileSource,omitempty" xml:"FileSource,omitempty"`
	UploadType     *string                                                            `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
	UploadURL      *string                                                            `json:"UploadURL,omitempty" xml:"UploadURL,omitempty"`
	UploadOSSParam *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam `json:"UploadOSSParam,omitempty" xml:"UploadOSSParam,omitempty" type:"Struct"`
	UploadedSize   *int64                                                             `json:"UploadedSize,omitempty" xml:"UploadedSize,omitempty"`
	JobStatus      *string                                                            `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	JobStatusDesc  *string                                                            `json:"JobStatusDesc,omitempty" xml:"JobStatusDesc,omitempty"`
	AttachmentKey  *string                                                            `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
}

func (s GetUserUploadFileJobResponseBodyUploadFileJobDetail) String() string {
	return tea.Prettify(s)
}

func (s GetUserUploadFileJobResponseBodyUploadFileJobDetail) GoString() string {
	return s.String()
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetJobKey(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.JobKey = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetFileName(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.FileName = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetFileSize(v int64) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.FileSize = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetFileSource(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.FileSource = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetUploadType(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.UploadType = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetUploadURL(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.UploadURL = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetUploadOSSParam(v *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.UploadOSSParam = v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetUploadedSize(v int64) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.UploadedSize = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetJobStatus(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.JobStatus = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetJobStatusDesc(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.JobStatusDesc = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetAttachmentKey(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.AttachmentKey = &v
	return s
}

type GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam struct {
	Endpoint   *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
}

func (s GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) String() string {
	return tea.Prettify(s)
}

func (s GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) GoString() string {
	return s.String()
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) SetEndpoint(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam {
	s.Endpoint = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) SetBucketName(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam {
	s.BucketName = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) SetObjectName(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam {
	s.ObjectName = &v
	return s
}

type GetUserUploadFileJobResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserUploadFileJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserUploadFileJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserUploadFileJobResponse) GoString() string {
	return s.String()
}

func (s *GetUserUploadFileJobResponse) SetHeaders(v map[string]*string) *GetUserUploadFileJobResponse {
	s.Headers = v
	return s
}

func (s *GetUserUploadFileJobResponse) SetBody(v *GetUserUploadFileJobResponseBody) *GetUserUploadFileJobResponse {
	s.Body = v
	return s
}

type ListDDLPublishRecordsRequest struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Tid     *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListDDLPublishRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDDLPublishRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListDDLPublishRecordsRequest) SetOrderId(v int64) *ListDDLPublishRecordsRequest {
	s.OrderId = &v
	return s
}

func (s *ListDDLPublishRecordsRequest) SetTid(v int64) *ListDDLPublishRecordsRequest {
	s.Tid = &v
	return s
}

type ListDDLPublishRecordsResponseBody struct {
	// Id of the request
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success              *bool                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage         *string                                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode            *string                                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	DDLPublishRecordList []*ListDDLPublishRecordsResponseBodyDDLPublishRecordList `json:"DDLPublishRecordList,omitempty" xml:"DDLPublishRecordList,omitempty" type:"Repeated"`
}

func (s ListDDLPublishRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDDLPublishRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDDLPublishRecordsResponseBody) SetRequestId(v string) *ListDDLPublishRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBody) SetSuccess(v bool) *ListDDLPublishRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBody) SetErrorMessage(v string) *ListDDLPublishRecordsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBody) SetErrorCode(v string) *ListDDLPublishRecordsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBody) SetDDLPublishRecordList(v []*ListDDLPublishRecordsResponseBodyDDLPublishRecordList) *ListDDLPublishRecordsResponseBody {
	s.DDLPublishRecordList = v
	return s
}

type ListDDLPublishRecordsResponseBodyDDLPublishRecordList struct {
	AuditStatus         *string                                                                     `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AuditExpireTime     *string                                                                     `json:"AuditExpireTime,omitempty" xml:"AuditExpireTime,omitempty"`
	CreatorId           *int64                                                                      `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Finality            *bool                                                                       `json:"Finality,omitempty" xml:"Finality,omitempty"`
	FinalityReason      *string                                                                     `json:"FinalityReason,omitempty" xml:"FinalityReason,omitempty"`
	PublishStatus       *string                                                                     `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	RiskLevel           *string                                                                     `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	StatusDesc          *string                                                                     `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	WorkflowInstanceId  *int64                                                                      `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
	PublishTaskInfoList []*ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList `json:"PublishTaskInfoList,omitempty" xml:"PublishTaskInfoList,omitempty" type:"Repeated"`
}

func (s ListDDLPublishRecordsResponseBodyDDLPublishRecordList) String() string {
	return tea.Prettify(s)
}

func (s ListDDLPublishRecordsResponseBodyDDLPublishRecordList) GoString() string {
	return s.String()
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetAuditStatus(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.AuditStatus = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetAuditExpireTime(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.AuditExpireTime = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetCreatorId(v int64) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.CreatorId = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetFinality(v bool) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.Finality = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetFinalityReason(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.FinalityReason = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetPublishStatus(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.PublishStatus = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetRiskLevel(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.RiskLevel = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetStatusDesc(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.StatusDesc = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetWorkflowInstanceId(v int64) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetPublishTaskInfoList(v []*ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.PublishTaskInfoList = v
	return s
}

type ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList struct {
	DbId            *int64                                                                                    `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Logic           *bool                                                                                     `json:"Logic,omitempty" xml:"Logic,omitempty"`
	PlanTime        *string                                                                                   `json:"PlanTime,omitempty" xml:"PlanTime,omitempty"`
	PublishStrategy *string                                                                                   `json:"PublishStrategy,omitempty" xml:"PublishStrategy,omitempty"`
	StatusDesc      *string                                                                                   `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	TaskJobStatus   *string                                                                                   `json:"TaskJobStatus,omitempty" xml:"TaskJobStatus,omitempty"`
	PublishJobList  []*ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList `json:"PublishJobList,omitempty" xml:"PublishJobList,omitempty" type:"Repeated"`
}

func (s ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) GoString() string {
	return s.String()
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) SetDbId(v int64) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList {
	s.DbId = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) SetLogic(v bool) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList {
	s.Logic = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) SetPlanTime(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList {
	s.PlanTime = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) SetPublishStrategy(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList {
	s.PublishStrategy = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) SetStatusDesc(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList {
	s.StatusDesc = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) SetTaskJobStatus(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList {
	s.TaskJobStatus = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) SetPublishJobList(v []*ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList {
	s.PublishJobList = v
	return s
}

type ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList struct {
	ExecuteCount  *int64  `json:"ExecuteCount,omitempty" xml:"ExecuteCount,omitempty"`
	Scripts       *string `json:"Scripts,omitempty" xml:"Scripts,omitempty"`
	TableName     *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	StatusDesc    *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	TaskJobStatus *string `json:"TaskJobStatus,omitempty" xml:"TaskJobStatus,omitempty"`
	DBTaskGroupId *int64  `json:"DBTaskGroupId,omitempty" xml:"DBTaskGroupId,omitempty"`
}

func (s ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) String() string {
	return tea.Prettify(s)
}

func (s ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) GoString() string {
	return s.String()
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) SetExecuteCount(v int64) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList {
	s.ExecuteCount = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) SetScripts(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList {
	s.Scripts = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) SetTableName(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList {
	s.TableName = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) SetStatusDesc(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList {
	s.StatusDesc = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) SetTaskJobStatus(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList {
	s.TaskJobStatus = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) SetDBTaskGroupId(v int64) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList {
	s.DBTaskGroupId = &v
	return s
}

type ListDDLPublishRecordsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDDLPublishRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDDLPublishRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDDLPublishRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListDDLPublishRecordsResponse) SetHeaders(v map[string]*string) *ListDDLPublishRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListDDLPublishRecordsResponse) SetBody(v *ListDDLPublishRecordsResponseBody) *ListDDLPublishRecordsResponse {
	s.Body = v
	return s
}

type GetStructSyncJobDetailRequest struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Tid     *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetStructSyncJobDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncJobDetailRequest) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobDetailRequest) SetOrderId(v int64) *GetStructSyncJobDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetStructSyncJobDetailRequest) SetTid(v int64) *GetStructSyncJobDetailRequest {
	s.Tid = &v
	return s
}

type GetStructSyncJobDetailResponseBody struct {
	// Id of the request
	RequestId           *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success             *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage        *string                                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode           *string                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	StructSyncJobDetail *GetStructSyncJobDetailResponseBodyStructSyncJobDetail `json:"StructSyncJobDetail,omitempty" xml:"StructSyncJobDetail,omitempty" type:"Struct"`
}

func (s GetStructSyncJobDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobDetailResponseBody) SetRequestId(v string) *GetStructSyncJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBody) SetSuccess(v bool) *GetStructSyncJobDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBody) SetErrorMessage(v string) *GetStructSyncJobDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBody) SetErrorCode(v string) *GetStructSyncJobDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBody) SetStructSyncJobDetail(v *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) *GetStructSyncJobDetailResponseBody {
	s.StructSyncJobDetail = v
	return s
}

type GetStructSyncJobDetailResponseBodyStructSyncJobDetail struct {
	JobStatus     *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	TableAnalyzed *int64  `json:"TableAnalyzed,omitempty" xml:"TableAnalyzed,omitempty"`
	TableCount    *int64  `json:"TableCount,omitempty" xml:"TableCount,omitempty"`
	SqlCount      *int64  `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	ExecuteCount  *int64  `json:"ExecuteCount,omitempty" xml:"ExecuteCount,omitempty"`
	SecurityRule  *string `json:"SecurityRule,omitempty" xml:"SecurityRule,omitempty"`
	DBTaskGroupId *int64  `json:"DBTaskGroupId,omitempty" xml:"DBTaskGroupId,omitempty"`
}

func (s GetStructSyncJobDetailResponseBodyStructSyncJobDetail) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncJobDetailResponseBodyStructSyncJobDetail) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetJobStatus(v string) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.JobStatus = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetMessage(v string) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.Message = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetTableAnalyzed(v int64) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.TableAnalyzed = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetTableCount(v int64) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.TableCount = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetSqlCount(v int64) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.SqlCount = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetExecuteCount(v int64) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.ExecuteCount = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetSecurityRule(v string) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.SecurityRule = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetDBTaskGroupId(v int64) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.DBTaskGroupId = &v
	return s
}

type GetStructSyncJobDetailResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStructSyncJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStructSyncJobDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncJobDetailResponse) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobDetailResponse) SetHeaders(v map[string]*string) *GetStructSyncJobDetailResponse {
	s.Headers = v
	return s
}

func (s *GetStructSyncJobDetailResponse) SetBody(v *GetStructSyncJobDetailResponseBody) *GetStructSyncJobDetailResponse {
	s.Body = v
	return s
}

type CreateUploadOSSFileJobRequest struct {
	FileSource   *string                                    `json:"FileSource,omitempty" xml:"FileSource,omitempty"`
	FileName     *string                                    `json:"FileName,omitempty" xml:"FileName,omitempty"`
	UploadTarget *CreateUploadOSSFileJobRequestUploadTarget `json:"UploadTarget,omitempty" xml:"UploadTarget,omitempty" type:"Struct"`
	Tid          *int64                                     `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateUploadOSSFileJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadOSSFileJobRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadOSSFileJobRequest) SetFileSource(v string) *CreateUploadOSSFileJobRequest {
	s.FileSource = &v
	return s
}

func (s *CreateUploadOSSFileJobRequest) SetFileName(v string) *CreateUploadOSSFileJobRequest {
	s.FileName = &v
	return s
}

func (s *CreateUploadOSSFileJobRequest) SetUploadTarget(v *CreateUploadOSSFileJobRequestUploadTarget) *CreateUploadOSSFileJobRequest {
	s.UploadTarget = v
	return s
}

func (s *CreateUploadOSSFileJobRequest) SetTid(v int64) *CreateUploadOSSFileJobRequest {
	s.Tid = &v
	return s
}

type CreateUploadOSSFileJobRequestUploadTarget struct {
	Endpoint   *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
}

func (s CreateUploadOSSFileJobRequestUploadTarget) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadOSSFileJobRequestUploadTarget) GoString() string {
	return s.String()
}

func (s *CreateUploadOSSFileJobRequestUploadTarget) SetEndpoint(v string) *CreateUploadOSSFileJobRequestUploadTarget {
	s.Endpoint = &v
	return s
}

func (s *CreateUploadOSSFileJobRequestUploadTarget) SetBucketName(v string) *CreateUploadOSSFileJobRequestUploadTarget {
	s.BucketName = &v
	return s
}

func (s *CreateUploadOSSFileJobRequestUploadTarget) SetObjectName(v string) *CreateUploadOSSFileJobRequestUploadTarget {
	s.ObjectName = &v
	return s
}

type CreateUploadOSSFileJobShrinkRequest struct {
	FileSource         *string `json:"FileSource,omitempty" xml:"FileSource,omitempty"`
	FileName           *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	UploadTargetShrink *string `json:"UploadTarget,omitempty" xml:"UploadTarget,omitempty"`
	Tid                *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateUploadOSSFileJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadOSSFileJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadOSSFileJobShrinkRequest) SetFileSource(v string) *CreateUploadOSSFileJobShrinkRequest {
	s.FileSource = &v
	return s
}

func (s *CreateUploadOSSFileJobShrinkRequest) SetFileName(v string) *CreateUploadOSSFileJobShrinkRequest {
	s.FileName = &v
	return s
}

func (s *CreateUploadOSSFileJobShrinkRequest) SetUploadTargetShrink(v string) *CreateUploadOSSFileJobShrinkRequest {
	s.UploadTargetShrink = &v
	return s
}

func (s *CreateUploadOSSFileJobShrinkRequest) SetTid(v int64) *CreateUploadOSSFileJobShrinkRequest {
	s.Tid = &v
	return s
}

type CreateUploadOSSFileJobResponseBody struct {
	// Id of the request
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	JobKey       *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
}

func (s CreateUploadOSSFileJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadOSSFileJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadOSSFileJobResponseBody) SetRequestId(v string) *CreateUploadOSSFileJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadOSSFileJobResponseBody) SetSuccess(v bool) *CreateUploadOSSFileJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUploadOSSFileJobResponseBody) SetErrorMessage(v string) *CreateUploadOSSFileJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateUploadOSSFileJobResponseBody) SetErrorCode(v string) *CreateUploadOSSFileJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateUploadOSSFileJobResponseBody) SetJobKey(v string) *CreateUploadOSSFileJobResponseBody {
	s.JobKey = &v
	return s
}

type CreateUploadOSSFileJobResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUploadOSSFileJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUploadOSSFileJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadOSSFileJobResponse) GoString() string {
	return s.String()
}

func (s *CreateUploadOSSFileJobResponse) SetHeaders(v map[string]*string) *CreateUploadOSSFileJobResponse {
	s.Headers = v
	return s
}

func (s *CreateUploadOSSFileJobResponse) SetBody(v *CreateUploadOSSFileJobResponseBody) *CreateUploadOSSFileJobResponse {
	s.Body = v
	return s
}

type SearchDatabaseRequest struct {
	Tid          *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	SearchKey    *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	EnvType      *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	SearchRange  *string `json:"SearchRange,omitempty" xml:"SearchRange,omitempty"`
	SearchTarget *string `json:"SearchTarget,omitempty" xml:"SearchTarget,omitempty"`
	DbType       *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
}

func (s SearchDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchDatabaseRequest) GoString() string {
	return s.String()
}

func (s *SearchDatabaseRequest) SetTid(v int64) *SearchDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *SearchDatabaseRequest) SetSearchKey(v string) *SearchDatabaseRequest {
	s.SearchKey = &v
	return s
}

func (s *SearchDatabaseRequest) SetPageNumber(v int32) *SearchDatabaseRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchDatabaseRequest) SetPageSize(v int32) *SearchDatabaseRequest {
	s.PageSize = &v
	return s
}

func (s *SearchDatabaseRequest) SetEnvType(v string) *SearchDatabaseRequest {
	s.EnvType = &v
	return s
}

func (s *SearchDatabaseRequest) SetSearchRange(v string) *SearchDatabaseRequest {
	s.SearchRange = &v
	return s
}

func (s *SearchDatabaseRequest) SetSearchTarget(v string) *SearchDatabaseRequest {
	s.SearchTarget = &v
	return s
}

func (s *SearchDatabaseRequest) SetDbType(v string) *SearchDatabaseRequest {
	s.DbType = &v
	return s
}

type SearchDatabaseResponseBody struct {
	TotalCount         *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId          *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SearchDatabaseList *SearchDatabaseResponseBodySearchDatabaseList `json:"SearchDatabaseList,omitempty" xml:"SearchDatabaseList,omitempty" type:"Struct"`
	ErrorCode          *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage       *string                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success            *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *SearchDatabaseResponseBody) SetTotalCount(v int64) *SearchDatabaseResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchDatabaseResponseBody) SetRequestId(v string) *SearchDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchDatabaseResponseBody) SetSearchDatabaseList(v *SearchDatabaseResponseBodySearchDatabaseList) *SearchDatabaseResponseBody {
	s.SearchDatabaseList = v
	return s
}

func (s *SearchDatabaseResponseBody) SetErrorCode(v string) *SearchDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchDatabaseResponseBody) SetErrorMessage(v string) *SearchDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SearchDatabaseResponseBody) SetSuccess(v bool) *SearchDatabaseResponseBody {
	s.Success = &v
	return s
}

type SearchDatabaseResponseBodySearchDatabaseList struct {
	SearchDatabase []*SearchDatabaseResponseBodySearchDatabaseListSearchDatabase `json:"SearchDatabase,omitempty" xml:"SearchDatabase,omitempty" type:"Repeated"`
}

func (s SearchDatabaseResponseBodySearchDatabaseList) String() string {
	return tea.Prettify(s)
}

func (s SearchDatabaseResponseBodySearchDatabaseList) GoString() string {
	return s.String()
}

func (s *SearchDatabaseResponseBodySearchDatabaseList) SetSearchDatabase(v []*SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) *SearchDatabaseResponseBodySearchDatabaseList {
	s.SearchDatabase = v
	return s
}

type SearchDatabaseResponseBodySearchDatabaseListSearchDatabase struct {
	DatabaseId    *string                                                                  `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	Host          *string                                                                  `json:"Host,omitempty" xml:"Host,omitempty"`
	DbaId         *string                                                                  `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	SchemaName    *string                                                                  `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	Logic         *bool                                                                    `json:"Logic,omitempty" xml:"Logic,omitempty"`
	DatalinkName  *string                                                                  `json:"DatalinkName,omitempty" xml:"DatalinkName,omitempty"`
	Port          *int32                                                                   `json:"Port,omitempty" xml:"Port,omitempty"`
	EnvType       *string                                                                  `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Sid           *string                                                                  `json:"Sid,omitempty" xml:"Sid,omitempty"`
	OwnerIdList   *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList   `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	Encoding      *string                                                                  `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	DbType        *string                                                                  `json:"DbType,omitempty" xml:"DbType,omitempty"`
	OwnerNameList *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	SearchName    *string                                                                  `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	Alias         *string                                                                  `json:"Alias,omitempty" xml:"Alias,omitempty"`
}

func (s SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) String() string {
	return tea.Prettify(s)
}

func (s SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GoString() string {
	return s.String()
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetDatabaseId(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.DatabaseId = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetHost(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.Host = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetDbaId(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.DbaId = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetSchemaName(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.SchemaName = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetLogic(v bool) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.Logic = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetDatalinkName(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.DatalinkName = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetPort(v int32) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.Port = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetEnvType(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.EnvType = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetSid(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.Sid = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetOwnerIdList(v *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.OwnerIdList = v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetEncoding(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.Encoding = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetDbType(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.DbType = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetOwnerNameList(v *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.OwnerNameList = v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetSearchName(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.SearchName = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetAlias(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.Alias = &v
	return s
}

type SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList) String() string {
	return tea.Prettify(s)
}

func (s SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList) GoString() string {
	return s.String()
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList) SetOwnerIds(v []*string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList {
	s.OwnerIds = v
	return s
}

type SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList) String() string {
	return tea.Prettify(s)
}

func (s SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList) GoString() string {
	return s.String()
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList) SetOwnerNames(v []*string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList {
	s.OwnerNames = v
	return s
}

type SearchDatabaseResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchDatabaseResponse) GoString() string {
	return s.String()
}

func (s *SearchDatabaseResponse) SetHeaders(v map[string]*string) *SearchDatabaseResponse {
	s.Headers = v
	return s
}

func (s *SearchDatabaseResponse) SetBody(v *SearchDatabaseResponseBody) *SearchDatabaseResponse {
	s.Body = v
	return s
}

type GetDBTopologyRequest struct {
	LogicDbId *int64 `json:"LogicDbId,omitempty" xml:"LogicDbId,omitempty"`
	Tid       *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDBTopologyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDBTopologyRequest) GoString() string {
	return s.String()
}

func (s *GetDBTopologyRequest) SetLogicDbId(v int64) *GetDBTopologyRequest {
	s.LogicDbId = &v
	return s
}

func (s *GetDBTopologyRequest) SetTid(v int64) *GetDBTopologyRequest {
	s.Tid = &v
	return s
}

type GetDBTopologyResponseBody struct {
	// Id of the request
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode    *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	DBTopology   *GetDBTopologyResponseBodyDBTopology `json:"DBTopology,omitempty" xml:"DBTopology,omitempty" type:"Struct"`
}

func (s GetDBTopologyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDBTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *GetDBTopologyResponseBody) SetRequestId(v string) *GetDBTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDBTopologyResponseBody) SetSuccess(v bool) *GetDBTopologyResponseBody {
	s.Success = &v
	return s
}

func (s *GetDBTopologyResponseBody) SetErrorMessage(v string) *GetDBTopologyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDBTopologyResponseBody) SetErrorCode(v string) *GetDBTopologyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDBTopologyResponseBody) SetDBTopology(v *GetDBTopologyResponseBodyDBTopology) *GetDBTopologyResponseBody {
	s.DBTopology = v
	return s
}

type GetDBTopologyResponseBodyDBTopology struct {
	LogicDbId          *int64                                                   `json:"LogicDbId,omitempty" xml:"LogicDbId,omitempty"`
	LogicDbName        *string                                                  `json:"LogicDbName,omitempty" xml:"LogicDbName,omitempty"`
	SearchName         *string                                                  `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	Alias              *string                                                  `json:"Alias,omitempty" xml:"Alias,omitempty"`
	DbType             *string                                                  `json:"DbType,omitempty" xml:"DbType,omitempty"`
	EnvType            *string                                                  `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	DBTopologyInfoList []*GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList `json:"DBTopologyInfoList,omitempty" xml:"DBTopologyInfoList,omitempty" type:"Repeated"`
}

func (s GetDBTopologyResponseBodyDBTopology) String() string {
	return tea.Prettify(s)
}

func (s GetDBTopologyResponseBodyDBTopology) GoString() string {
	return s.String()
}

func (s *GetDBTopologyResponseBodyDBTopology) SetLogicDbId(v int64) *GetDBTopologyResponseBodyDBTopology {
	s.LogicDbId = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopology) SetLogicDbName(v string) *GetDBTopologyResponseBodyDBTopology {
	s.LogicDbName = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopology) SetSearchName(v string) *GetDBTopologyResponseBodyDBTopology {
	s.SearchName = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopology) SetAlias(v string) *GetDBTopologyResponseBodyDBTopology {
	s.Alias = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopology) SetDbType(v string) *GetDBTopologyResponseBodyDBTopology {
	s.DbType = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopology) SetEnvType(v string) *GetDBTopologyResponseBodyDBTopology {
	s.EnvType = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopology) SetDBTopologyInfoList(v []*GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) *GetDBTopologyResponseBodyDBTopology {
	s.DBTopologyInfoList = v
	return s
}

type GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList struct {
	DbId               *int64  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	SchemaName         *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	CatalogName        *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	SearchName         *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	DbType             *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	EnvType            *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	InstanceId         *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceResourceId *string `json:"InstanceResourceId,omitempty" xml:"InstanceResourceId,omitempty"`
	InstanceSource     *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
}

func (s GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) GoString() string {
	return s.String()
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetDbId(v int64) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.DbId = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetSchemaName(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.SchemaName = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetCatalogName(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.CatalogName = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetSearchName(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.SearchName = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetDbType(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.DbType = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetEnvType(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.EnvType = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetInstanceId(v int64) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.InstanceId = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetRegionId(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.RegionId = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetInstanceResourceId(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.InstanceResourceId = &v
	return s
}

func (s *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList) SetInstanceSource(v string) *GetDBTopologyResponseBodyDBTopologyDBTopologyInfoList {
	s.InstanceSource = &v
	return s
}

type GetDBTopologyResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDBTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDBTopologyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDBTopologyResponse) GoString() string {
	return s.String()
}

func (s *GetDBTopologyResponse) SetHeaders(v map[string]*string) *GetDBTopologyResponse {
	s.Headers = v
	return s
}

func (s *GetDBTopologyResponse) SetBody(v *GetDBTopologyResponseBody) *GetDBTopologyResponse {
	s.Body = v
	return s
}

type GetSQLReviewCheckResultStatusRequest struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Tid     *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetSQLReviewCheckResultStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSQLReviewCheckResultStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSQLReviewCheckResultStatusRequest) SetOrderId(v int64) *GetSQLReviewCheckResultStatusRequest {
	s.OrderId = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusRequest) SetTid(v int64) *GetSQLReviewCheckResultStatusRequest {
	s.Tid = &v
	return s
}

type GetSQLReviewCheckResultStatusResponseBody struct {
	RequestId         *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode         *string                                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage      *string                                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success           *bool                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	CheckResultStatus *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus `json:"CheckResultStatus,omitempty" xml:"CheckResultStatus,omitempty" type:"Struct"`
}

func (s GetSQLReviewCheckResultStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSQLReviewCheckResultStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSQLReviewCheckResultStatusResponseBody) SetRequestId(v string) *GetSQLReviewCheckResultStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBody) SetErrorCode(v string) *GetSQLReviewCheckResultStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBody) SetErrorMessage(v string) *GetSQLReviewCheckResultStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBody) SetSuccess(v bool) *GetSQLReviewCheckResultStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBody) SetCheckResultStatus(v *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) *GetSQLReviewCheckResultStatusResponseBody {
	s.CheckResultStatus = v
	return s
}

type GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus struct {
	TotalSQLCount     *int64                                                                       `json:"TotalSQLCount,omitempty" xml:"TotalSQLCount,omitempty"`
	CheckedCount      *int64                                                                       `json:"CheckedCount,omitempty" xml:"CheckedCount,omitempty"`
	CheckStatusResult *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult `json:"CheckStatusResult,omitempty" xml:"CheckStatusResult,omitempty" type:"Struct"`
	SQLReviewResult   *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult   `json:"SQLReviewResult,omitempty" xml:"SQLReviewResult,omitempty" type:"Struct"`
}

func (s GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) String() string {
	return tea.Prettify(s)
}

func (s GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) GoString() string {
	return s.String()
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) SetTotalSQLCount(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus {
	s.TotalSQLCount = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) SetCheckedCount(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus {
	s.CheckedCount = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) SetCheckStatusResult(v *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus {
	s.CheckStatusResult = v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) SetSQLReviewResult(v *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus {
	s.SQLReviewResult = v
	return s
}

type GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult struct {
	New          *int64 `json:"New,omitempty" xml:"New,omitempty"`
	Unknown      *int64 `json:"Unknown,omitempty" xml:"Unknown,omitempty"`
	CheckNotPass *int64 `json:"CheckNotPass,omitempty" xml:"CheckNotPass,omitempty"`
	CheckPass    *int64 `json:"CheckPass,omitempty" xml:"CheckPass,omitempty"`
	ForcePass    *int64 `json:"ForcePass,omitempty" xml:"ForcePass,omitempty"`
	ForceNotPass *int64 `json:"ForceNotPass,omitempty" xml:"ForceNotPass,omitempty"`
}

func (s GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) String() string {
	return tea.Prettify(s)
}

func (s GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) GoString() string {
	return s.String()
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) SetNew(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult {
	s.New = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) SetUnknown(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult {
	s.Unknown = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) SetCheckNotPass(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult {
	s.CheckNotPass = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) SetCheckPass(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult {
	s.CheckPass = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) SetForcePass(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult {
	s.ForcePass = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) SetForceNotPass(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult {
	s.ForceNotPass = &v
	return s
}

type GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult struct {
	MustImprove       *int64 `json:"MustImprove,omitempty" xml:"MustImprove,omitempty"`
	PotentialIssue    *int64 `json:"PotentialIssue,omitempty" xml:"PotentialIssue,omitempty"`
	SuggestImprove    *int64 `json:"SuggestImprove,omitempty" xml:"SuggestImprove,omitempty"`
	UseDmsToolkit     *int64 `json:"UseDmsToolkit,omitempty" xml:"UseDmsToolkit,omitempty"`
	UseDmsDmlUnlock   *int64 `json:"UseDmsDmlUnlock,omitempty" xml:"UseDmsDmlUnlock,omitempty"`
	TableIndexSuggest *int64 `json:"TableIndexSuggest,omitempty" xml:"TableIndexSuggest,omitempty"`
}

func (s GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) String() string {
	return tea.Prettify(s)
}

func (s GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) GoString() string {
	return s.String()
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) SetMustImprove(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult {
	s.MustImprove = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) SetPotentialIssue(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult {
	s.PotentialIssue = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) SetSuggestImprove(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult {
	s.SuggestImprove = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) SetUseDmsToolkit(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult {
	s.UseDmsToolkit = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) SetUseDmsDmlUnlock(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult {
	s.UseDmsDmlUnlock = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) SetTableIndexSuggest(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult {
	s.TableIndexSuggest = &v
	return s
}

type GetSQLReviewCheckResultStatusResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSQLReviewCheckResultStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSQLReviewCheckResultStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSQLReviewCheckResultStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSQLReviewCheckResultStatusResponse) SetHeaders(v map[string]*string) *GetSQLReviewCheckResultStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponse) SetBody(v *GetSQLReviewCheckResultStatusResponseBody) *GetSQLReviewCheckResultStatusResponse {
	s.Body = v
	return s
}

type SyncDatabaseMetaRequest struct {
	Tid   *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	DbId  *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Logic *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s SyncDatabaseMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncDatabaseMetaRequest) GoString() string {
	return s.String()
}

func (s *SyncDatabaseMetaRequest) SetTid(v int64) *SyncDatabaseMetaRequest {
	s.Tid = &v
	return s
}

func (s *SyncDatabaseMetaRequest) SetDbId(v string) *SyncDatabaseMetaRequest {
	s.DbId = &v
	return s
}

func (s *SyncDatabaseMetaRequest) SetLogic(v bool) *SyncDatabaseMetaRequest {
	s.Logic = &v
	return s
}

type SyncDatabaseMetaResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncDatabaseMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncDatabaseMetaResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDatabaseMetaResponseBody) SetRequestId(v string) *SyncDatabaseMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncDatabaseMetaResponseBody) SetErrorCode(v string) *SyncDatabaseMetaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SyncDatabaseMetaResponseBody) SetErrorMessage(v string) *SyncDatabaseMetaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SyncDatabaseMetaResponseBody) SetSuccess(v bool) *SyncDatabaseMetaResponseBody {
	s.Success = &v
	return s
}

type SyncDatabaseMetaResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncDatabaseMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncDatabaseMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncDatabaseMetaResponse) GoString() string {
	return s.String()
}

func (s *SyncDatabaseMetaResponse) SetHeaders(v map[string]*string) *SyncDatabaseMetaResponse {
	s.Headers = v
	return s
}

func (s *SyncDatabaseMetaResponse) SetBody(v *SyncDatabaseMetaResponseBody) *SyncDatabaseMetaResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	Tid    *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	Uid    *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetTid(v int64) *GetUserRequest {
	s.Tid = &v
	return s
}

func (s *GetUserRequest) SetUid(v string) *GetUserRequest {
	s.Uid = &v
	return s
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

type GetUserResponseBody struct {
	User         *GetUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
	RequestId    *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetUser(v *GetUserResponseBodyUser) *GetUserResponseBody {
	s.User = v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetErrorCode(v string) *GetUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserResponseBody) SetErrorMessage(v string) *GetUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserResponseBody) SetSuccess(v bool) *GetUserResponseBody {
	s.Success = &v
	return s
}

type GetUserResponseBodyUser struct {
	State           *string                              `json:"State,omitempty" xml:"State,omitempty"`
	CurResultCount  *int64                               `json:"CurResultCount,omitempty" xml:"CurResultCount,omitempty"`
	UserId          *string                              `json:"UserId,omitempty" xml:"UserId,omitempty"`
	LastLoginTime   *string                              `json:"LastLoginTime,omitempty" xml:"LastLoginTime,omitempty"`
	MaxResultCount  *int64                               `json:"MaxResultCount,omitempty" xml:"MaxResultCount,omitempty"`
	ParentUid       *int64                               `json:"ParentUid,omitempty" xml:"ParentUid,omitempty"`
	RoleIdList      *GetUserResponseBodyUserRoleIdList   `json:"RoleIdList,omitempty" xml:"RoleIdList,omitempty" type:"Struct"`
	RoleNameList    *GetUserResponseBodyUserRoleNameList `json:"RoleNameList,omitempty" xml:"RoleNameList,omitempty" type:"Struct"`
	NickName        *string                              `json:"NickName,omitempty" xml:"NickName,omitempty"`
	MaxExecuteCount *int64                               `json:"MaxExecuteCount,omitempty" xml:"MaxExecuteCount,omitempty"`
	CurExecuteCount *int64                               `json:"CurExecuteCount,omitempty" xml:"CurExecuteCount,omitempty"`
	Mobile          *string                              `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Uid             *string                              `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUser) SetState(v string) *GetUserResponseBodyUser {
	s.State = &v
	return s
}

func (s *GetUserResponseBodyUser) SetCurResultCount(v int64) *GetUserResponseBodyUser {
	s.CurResultCount = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserId(v string) *GetUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetLastLoginTime(v string) *GetUserResponseBodyUser {
	s.LastLoginTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetMaxResultCount(v int64) *GetUserResponseBodyUser {
	s.MaxResultCount = &v
	return s
}

func (s *GetUserResponseBodyUser) SetParentUid(v int64) *GetUserResponseBodyUser {
	s.ParentUid = &v
	return s
}

func (s *GetUserResponseBodyUser) SetRoleIdList(v *GetUserResponseBodyUserRoleIdList) *GetUserResponseBodyUser {
	s.RoleIdList = v
	return s
}

func (s *GetUserResponseBodyUser) SetRoleNameList(v *GetUserResponseBodyUserRoleNameList) *GetUserResponseBodyUser {
	s.RoleNameList = v
	return s
}

func (s *GetUserResponseBodyUser) SetNickName(v string) *GetUserResponseBodyUser {
	s.NickName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetMaxExecuteCount(v int64) *GetUserResponseBodyUser {
	s.MaxExecuteCount = &v
	return s
}

func (s *GetUserResponseBodyUser) SetCurExecuteCount(v int64) *GetUserResponseBodyUser {
	s.CurExecuteCount = &v
	return s
}

func (s *GetUserResponseBodyUser) SetMobile(v string) *GetUserResponseBodyUser {
	s.Mobile = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUid(v string) *GetUserResponseBodyUser {
	s.Uid = &v
	return s
}

type GetUserResponseBodyUserRoleIdList struct {
	RoleIds []*int32 `json:"RoleIds,omitempty" xml:"RoleIds,omitempty" type:"Repeated"`
}

func (s GetUserResponseBodyUserRoleIdList) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUserRoleIdList) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserRoleIdList) SetRoleIds(v []*int32) *GetUserResponseBodyUserRoleIdList {
	s.RoleIds = v
	return s
}

type GetUserResponseBodyUserRoleNameList struct {
	RoleNames []*string `json:"RoleNames,omitempty" xml:"RoleNames,omitempty" type:"Repeated"`
}

func (s GetUserResponseBodyUserRoleNameList) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUserRoleNameList) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserRoleNameList) SetRoleNames(v []*string) *GetUserResponseBodyUserRoleNameList {
	s.RoleNames = v
	return s
}

type GetUserResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type ExecuteStructSyncRequest struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Tid     *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ExecuteStructSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStructSyncRequest) GoString() string {
	return s.String()
}

func (s *ExecuteStructSyncRequest) SetOrderId(v int64) *ExecuteStructSyncRequest {
	s.OrderId = &v
	return s
}

func (s *ExecuteStructSyncRequest) SetTid(v int64) *ExecuteStructSyncRequest {
	s.Tid = &v
	return s
}

type ExecuteStructSyncResponseBody struct {
	// Id of the request
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s ExecuteStructSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStructSyncResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteStructSyncResponseBody) SetRequestId(v string) *ExecuteStructSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteStructSyncResponseBody) SetSuccess(v bool) *ExecuteStructSyncResponseBody {
	s.Success = &v
	return s
}

func (s *ExecuteStructSyncResponseBody) SetErrorMessage(v string) *ExecuteStructSyncResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ExecuteStructSyncResponseBody) SetErrorCode(v string) *ExecuteStructSyncResponseBody {
	s.ErrorCode = &v
	return s
}

type ExecuteStructSyncResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteStructSyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteStructSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteStructSyncResponse) GoString() string {
	return s.String()
}

func (s *ExecuteStructSyncResponse) SetHeaders(v map[string]*string) *ExecuteStructSyncResponse {
	s.Headers = v
	return s
}

func (s *ExecuteStructSyncResponse) SetBody(v *ExecuteStructSyncResponseBody) *ExecuteStructSyncResponse {
	s.Body = v
	return s
}

type GetDataCorrectOrderDetailRequest struct {
	Tid     *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s GetDataCorrectOrderDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailRequest) SetTid(v int64) *GetDataCorrectOrderDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetDataCorrectOrderDetailRequest) SetOrderId(v int64) *GetDataCorrectOrderDetailRequest {
	s.OrderId = &v
	return s
}

type GetDataCorrectOrderDetailResponseBody struct {
	RequestId              *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DataCorrectOrderDetail *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail `json:"DataCorrectOrderDetail,omitempty" xml:"DataCorrectOrderDetail,omitempty" type:"Struct"`
	ErrorCode              *string                                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage           *string                                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success                *bool                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataCorrectOrderDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBody) SetRequestId(v string) *GetDataCorrectOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBody) SetDataCorrectOrderDetail(v *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) *GetDataCorrectOrderDetailResponseBody {
	s.DataCorrectOrderDetail = v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBody) SetErrorCode(v string) *GetDataCorrectOrderDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBody) SetErrorMessage(v string) *GetDataCorrectOrderDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBody) SetSuccess(v bool) *GetDataCorrectOrderDetailResponseBody {
	s.Success = &v
	return s
}

type GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail struct {
	PreCheckDetail *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail `json:"PreCheckDetail,omitempty" xml:"PreCheckDetail,omitempty" type:"Struct"`
	OrderDetail    *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail    `json:"OrderDetail,omitempty" xml:"OrderDetail,omitempty" type:"Struct"`
	DatabaseList   *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList   `json:"DatabaseList,omitempty" xml:"DatabaseList,omitempty" type:"Struct"`
	Status         *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) SetPreCheckDetail(v *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail {
	s.PreCheckDetail = v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) SetOrderDetail(v *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail {
	s.OrderDetail = v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) SetDatabaseList(v *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail {
	s.DatabaseList = v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) SetStatus(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail {
	s.Status = &v
	return s
}

type GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail struct {
	TaskCheckDO []*GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO `json:"TaskCheckDO,omitempty" xml:"TaskCheckDO,omitempty" type:"Repeated"`
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail) SetTaskCheckDO(v []*GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail {
	s.TaskCheckDO = v
	return s
}

type GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO struct {
	UserTip     *string `json:"UserTip,omitempty" xml:"UserTip,omitempty"`
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckStep   *string `json:"CheckStep,omitempty" xml:"CheckStep,omitempty"`
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) SetUserTip(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO {
	s.UserTip = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) SetCheckStatus(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO {
	s.CheckStatus = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) SetCheckStep(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO {
	s.CheckStep = &v
	return s
}

type GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail struct {
	RbSQL                  *string `json:"RbSQL,omitempty" xml:"RbSQL,omitempty"`
	RbAttachmentName       *string `json:"RbAttachmentName,omitempty" xml:"RbAttachmentName,omitempty"`
	Classify               *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	ExeSQL                 *string `json:"ExeSQL,omitempty" xml:"ExeSQL,omitempty"`
	EstimateAffectRows     *int64  `json:"EstimateAffectRows,omitempty" xml:"EstimateAffectRows,omitempty"`
	RbSQLType              *string `json:"RbSQLType,omitempty" xml:"RbSQLType,omitempty"`
	ActualAffectRows       *int64  `json:"ActualAffectRows,omitempty" xml:"ActualAffectRows,omitempty"`
	IgnoreAffectRows       *bool   `json:"IgnoreAffectRows,omitempty" xml:"IgnoreAffectRows,omitempty"`
	AttachmentName         *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	SqlType                *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	IgnoreAffectRowsReason *string `json:"IgnoreAffectRowsReason,omitempty" xml:"IgnoreAffectRowsReason,omitempty"`
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetRbSQL(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.RbSQL = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetRbAttachmentName(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.RbAttachmentName = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetClassify(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.Classify = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetExeSQL(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.ExeSQL = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetEstimateAffectRows(v int64) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.EstimateAffectRows = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetRbSQLType(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.RbSQLType = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetActualAffectRows(v int64) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.ActualAffectRows = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetIgnoreAffectRows(v bool) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.IgnoreAffectRows = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetAttachmentName(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.AttachmentName = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetSqlType(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.SqlType = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetIgnoreAffectRowsReason(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.IgnoreAffectRowsReason = &v
	return s
}

type GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList struct {
	Database []*GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Repeated"`
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList) SetDatabase(v []*GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList {
	s.Database = v
	return s
}

type GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase struct {
	DbId       *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	DbType     *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	Logic      *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	EnvType    *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) SetDbId(v int32) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase {
	s.DbId = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) SetDbType(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase {
	s.DbType = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) SetLogic(v bool) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase {
	s.Logic = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) SetSearchName(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase {
	s.SearchName = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) SetEnvType(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase {
	s.EnvType = &v
	return s
}

type GetDataCorrectOrderDetailResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDataCorrectOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDataCorrectOrderDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponse) SetHeaders(v map[string]*string) *GetDataCorrectOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDataCorrectOrderDetailResponse) SetBody(v *GetDataCorrectOrderDetailResponseBody) *GetDataCorrectOrderDetailResponse {
	s.Body = v
	return s
}

type ListColumnsRequest struct {
	Tid     *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	Logic   *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s ListColumnsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListColumnsRequest) GoString() string {
	return s.String()
}

func (s *ListColumnsRequest) SetTid(v int64) *ListColumnsRequest {
	s.Tid = &v
	return s
}

func (s *ListColumnsRequest) SetTableId(v string) *ListColumnsRequest {
	s.TableId = &v
	return s
}

func (s *ListColumnsRequest) SetLogic(v bool) *ListColumnsRequest {
	s.Logic = &v
	return s
}

type ListColumnsResponseBody struct {
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ColumnList   *ListColumnsResponseBodyColumnList `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Struct"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListColumnsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *ListColumnsResponseBody) SetRequestId(v string) *ListColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListColumnsResponseBody) SetErrorCode(v string) *ListColumnsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListColumnsResponseBody) SetColumnList(v *ListColumnsResponseBodyColumnList) *ListColumnsResponseBody {
	s.ColumnList = v
	return s
}

func (s *ListColumnsResponseBody) SetErrorMessage(v string) *ListColumnsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListColumnsResponseBody) SetSuccess(v bool) *ListColumnsResponseBody {
	s.Success = &v
	return s
}

type ListColumnsResponseBodyColumnList struct {
	Column []*ListColumnsResponseBodyColumnListColumn `json:"Column,omitempty" xml:"Column,omitempty" type:"Repeated"`
}

func (s ListColumnsResponseBodyColumnList) String() string {
	return tea.Prettify(s)
}

func (s ListColumnsResponseBodyColumnList) GoString() string {
	return s.String()
}

func (s *ListColumnsResponseBodyColumnList) SetColumn(v []*ListColumnsResponseBodyColumnListColumn) *ListColumnsResponseBodyColumnList {
	s.Column = v
	return s
}

type ListColumnsResponseBodyColumnListColumn struct {
	ColumnType    *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	AutoIncrement *bool   `json:"AutoIncrement,omitempty" xml:"AutoIncrement,omitempty"`
	ColumnId      *string `json:"ColumnId,omitempty" xml:"ColumnId,omitempty"`
	DefaultValue  *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Sensitive     *bool   `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	ColumnName    *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DataPrecision *int32  `json:"DataPrecision,omitempty" xml:"DataPrecision,omitempty"`
	DataScale     *int32  `json:"DataScale,omitempty" xml:"DataScale,omitempty"`
	FunctionType  *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	Nullable      *bool   `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	DataLength    *int64  `json:"DataLength,omitempty" xml:"DataLength,omitempty"`
}

func (s ListColumnsResponseBodyColumnListColumn) String() string {
	return tea.Prettify(s)
}

func (s ListColumnsResponseBodyColumnListColumn) GoString() string {
	return s.String()
}

func (s *ListColumnsResponseBodyColumnListColumn) SetColumnType(v string) *ListColumnsResponseBodyColumnListColumn {
	s.ColumnType = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetAutoIncrement(v bool) *ListColumnsResponseBodyColumnListColumn {
	s.AutoIncrement = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetColumnId(v string) *ListColumnsResponseBodyColumnListColumn {
	s.ColumnId = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetDefaultValue(v string) *ListColumnsResponseBodyColumnListColumn {
	s.DefaultValue = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetSensitive(v bool) *ListColumnsResponseBodyColumnListColumn {
	s.Sensitive = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetColumnName(v string) *ListColumnsResponseBodyColumnListColumn {
	s.ColumnName = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetSecurityLevel(v string) *ListColumnsResponseBodyColumnListColumn {
	s.SecurityLevel = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetDescription(v string) *ListColumnsResponseBodyColumnListColumn {
	s.Description = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetDataPrecision(v int32) *ListColumnsResponseBodyColumnListColumn {
	s.DataPrecision = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetDataScale(v int32) *ListColumnsResponseBodyColumnListColumn {
	s.DataScale = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetFunctionType(v string) *ListColumnsResponseBodyColumnListColumn {
	s.FunctionType = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetNullable(v bool) *ListColumnsResponseBodyColumnListColumn {
	s.Nullable = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetDataLength(v int64) *ListColumnsResponseBodyColumnListColumn {
	s.DataLength = &v
	return s
}

type ListColumnsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListColumnsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListColumnsResponse) GoString() string {
	return s.String()
}

func (s *ListColumnsResponse) SetHeaders(v map[string]*string) *ListColumnsResponse {
	s.Headers = v
	return s
}

func (s *ListColumnsResponse) SetBody(v *ListColumnsResponseBody) *ListColumnsResponse {
	s.Body = v
	return s
}

type RevokeUserPermissionRequest struct {
	Tid          *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DsType       *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
	DbId         *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Logic        *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	TableId      *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	TableName    *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	PermTypes    *string `json:"PermTypes,omitempty" xml:"PermTypes,omitempty"`
	UserAccessId *string `json:"UserAccessId,omitempty" xml:"UserAccessId,omitempty"`
}

func (s RevokeUserPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeUserPermissionRequest) GoString() string {
	return s.String()
}

func (s *RevokeUserPermissionRequest) SetTid(v int64) *RevokeUserPermissionRequest {
	s.Tid = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetUserId(v string) *RevokeUserPermissionRequest {
	s.UserId = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetDsType(v string) *RevokeUserPermissionRequest {
	s.DsType = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetDbId(v string) *RevokeUserPermissionRequest {
	s.DbId = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetLogic(v bool) *RevokeUserPermissionRequest {
	s.Logic = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetTableId(v string) *RevokeUserPermissionRequest {
	s.TableId = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetTableName(v string) *RevokeUserPermissionRequest {
	s.TableName = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetPermTypes(v string) *RevokeUserPermissionRequest {
	s.PermTypes = &v
	return s
}

func (s *RevokeUserPermissionRequest) SetUserAccessId(v string) *RevokeUserPermissionRequest {
	s.UserAccessId = &v
	return s
}

type RevokeUserPermissionResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokeUserPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeUserPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeUserPermissionResponseBody) SetRequestId(v string) *RevokeUserPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeUserPermissionResponseBody) SetErrorCode(v string) *RevokeUserPermissionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RevokeUserPermissionResponseBody) SetErrorMessage(v string) *RevokeUserPermissionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RevokeUserPermissionResponseBody) SetSuccess(v bool) *RevokeUserPermissionResponseBody {
	s.Success = &v
	return s
}

type RevokeUserPermissionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RevokeUserPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokeUserPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeUserPermissionResponse) GoString() string {
	return s.String()
}

func (s *RevokeUserPermissionResponse) SetHeaders(v map[string]*string) *RevokeUserPermissionResponse {
	s.Headers = v
	return s
}

func (s *RevokeUserPermissionResponse) SetBody(v *RevokeUserPermissionResponseBody) *RevokeUserPermissionResponse {
	s.Body = v
	return s
}

type GetMetaTableColumnRequest struct {
	Tid       *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s GetMetaTableColumnRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMetaTableColumnRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTableColumnRequest) SetTid(v int64) *GetMetaTableColumnRequest {
	s.Tid = &v
	return s
}

func (s *GetMetaTableColumnRequest) SetTableGuid(v string) *GetMetaTableColumnRequest {
	s.TableGuid = &v
	return s
}

type GetMetaTableColumnResponseBody struct {
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ColumnList   []*GetMetaTableColumnResponseBodyColumnList `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Repeated"`
	ErrorMessage *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetaTableColumnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMetaTableColumnResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaTableColumnResponseBody) SetRequestId(v string) *GetMetaTableColumnResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaTableColumnResponseBody) SetErrorCode(v string) *GetMetaTableColumnResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaTableColumnResponseBody) SetColumnList(v []*GetMetaTableColumnResponseBodyColumnList) *GetMetaTableColumnResponseBody {
	s.ColumnList = v
	return s
}

func (s *GetMetaTableColumnResponseBody) SetErrorMessage(v string) *GetMetaTableColumnResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaTableColumnResponseBody) SetSuccess(v bool) *GetMetaTableColumnResponseBody {
	s.Success = &v
	return s
}

type GetMetaTableColumnResponseBodyColumnList struct {
	ColumnType    *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	AutoIncrement *bool   `json:"AutoIncrement,omitempty" xml:"AutoIncrement,omitempty"`
	ColumnId      *string `json:"ColumnId,omitempty" xml:"ColumnId,omitempty"`
	ColumnName    *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	PrimaryKey    *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DataPrecision *int32  `json:"DataPrecision,omitempty" xml:"DataPrecision,omitempty"`
	DataScale     *int32  `json:"DataScale,omitempty" xml:"DataScale,omitempty"`
	Position      *int32  `json:"Position,omitempty" xml:"Position,omitempty"`
	Nullable      *bool   `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	DataLength    *int64  `json:"DataLength,omitempty" xml:"DataLength,omitempty"`
}

func (s GetMetaTableColumnResponseBodyColumnList) String() string {
	return tea.Prettify(s)
}

func (s GetMetaTableColumnResponseBodyColumnList) GoString() string {
	return s.String()
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetColumnType(v string) *GetMetaTableColumnResponseBodyColumnList {
	s.ColumnType = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetAutoIncrement(v bool) *GetMetaTableColumnResponseBodyColumnList {
	s.AutoIncrement = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetColumnId(v string) *GetMetaTableColumnResponseBodyColumnList {
	s.ColumnId = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetColumnName(v string) *GetMetaTableColumnResponseBodyColumnList {
	s.ColumnName = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetSecurityLevel(v string) *GetMetaTableColumnResponseBodyColumnList {
	s.SecurityLevel = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetPrimaryKey(v string) *GetMetaTableColumnResponseBodyColumnList {
	s.PrimaryKey = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetDescription(v string) *GetMetaTableColumnResponseBodyColumnList {
	s.Description = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetDataPrecision(v int32) *GetMetaTableColumnResponseBodyColumnList {
	s.DataPrecision = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetDataScale(v int32) *GetMetaTableColumnResponseBodyColumnList {
	s.DataScale = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetPosition(v int32) *GetMetaTableColumnResponseBodyColumnList {
	s.Position = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetNullable(v bool) *GetMetaTableColumnResponseBodyColumnList {
	s.Nullable = &v
	return s
}

func (s *GetMetaTableColumnResponseBodyColumnList) SetDataLength(v int64) *GetMetaTableColumnResponseBodyColumnList {
	s.DataLength = &v
	return s
}

type GetMetaTableColumnResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMetaTableColumnResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMetaTableColumnResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMetaTableColumnResponse) GoString() string {
	return s.String()
}

func (s *GetMetaTableColumnResponse) SetHeaders(v map[string]*string) *GetMetaTableColumnResponse {
	s.Headers = v
	return s
}

func (s *GetMetaTableColumnResponse) SetBody(v *GetMetaTableColumnResponseBody) *GetMetaTableColumnResponse {
	s.Body = v
	return s
}

type EnableUserRequest struct {
	Tid *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s EnableUserRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableUserRequest) GoString() string {
	return s.String()
}

func (s *EnableUserRequest) SetTid(v int64) *EnableUserRequest {
	s.Tid = &v
	return s
}

func (s *EnableUserRequest) SetUid(v string) *EnableUserRequest {
	s.Uid = &v
	return s
}

type EnableUserResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableUserResponseBody) GoString() string {
	return s.String()
}

func (s *EnableUserResponseBody) SetRequestId(v string) *EnableUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableUserResponseBody) SetErrorCode(v string) *EnableUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *EnableUserResponseBody) SetErrorMessage(v string) *EnableUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *EnableUserResponseBody) SetSuccess(v bool) *EnableUserResponseBody {
	s.Success = &v
	return s
}

type EnableUserResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableUserResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableUserResponse) GoString() string {
	return s.String()
}

func (s *EnableUserResponse) SetHeaders(v map[string]*string) *EnableUserResponse {
	s.Headers = v
	return s
}

func (s *EnableUserResponse) SetBody(v *EnableUserResponseBody) *EnableUserResponse {
	s.Body = v
	return s
}

type UpdateInstanceRequest struct {
	Tid              *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceSource   *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	EnvType          *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	EcsInstanceId    *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	VpcId            *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	EcsRegion        *string `json:"EcsRegion,omitempty" xml:"EcsRegion,omitempty"`
	Host             *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Sid              *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	DatabaseUser     *string `json:"DatabaseUser,omitempty" xml:"DatabaseUser,omitempty"`
	DatabasePassword *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	InstanceAlias    *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	DbaId            *string `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	SafeRuleId       *string `json:"SafeRuleId,omitempty" xml:"SafeRuleId,omitempty"`
	QueryTimeout     *int32  `json:"QueryTimeout,omitempty" xml:"QueryTimeout,omitempty"`
	ExportTimeout    *int32  `json:"ExportTimeout,omitempty" xml:"ExportTimeout,omitempty"`
	DataLinkName     *string `json:"DataLinkName,omitempty" xml:"DataLinkName,omitempty"`
	DdlOnline        *int32  `json:"DdlOnline,omitempty" xml:"DdlOnline,omitempty"`
	UseDsql          *int32  `json:"UseDsql,omitempty" xml:"UseDsql,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkipTest         *bool   `json:"SkipTest,omitempty" xml:"SkipTest,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetTid(v int64) *UpdateInstanceRequest {
	s.Tid = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceType(v string) *UpdateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceSource(v string) *UpdateInstanceRequest {
	s.InstanceSource = &v
	return s
}

func (s *UpdateInstanceRequest) SetEnvType(v string) *UpdateInstanceRequest {
	s.EnvType = &v
	return s
}

func (s *UpdateInstanceRequest) SetEcsInstanceId(v string) *UpdateInstanceRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetVpcId(v string) *UpdateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *UpdateInstanceRequest) SetEcsRegion(v string) *UpdateInstanceRequest {
	s.EcsRegion = &v
	return s
}

func (s *UpdateInstanceRequest) SetHost(v string) *UpdateInstanceRequest {
	s.Host = &v
	return s
}

func (s *UpdateInstanceRequest) SetPort(v int32) *UpdateInstanceRequest {
	s.Port = &v
	return s
}

func (s *UpdateInstanceRequest) SetSid(v string) *UpdateInstanceRequest {
	s.Sid = &v
	return s
}

func (s *UpdateInstanceRequest) SetDatabaseUser(v string) *UpdateInstanceRequest {
	s.DatabaseUser = &v
	return s
}

func (s *UpdateInstanceRequest) SetDatabasePassword(v string) *UpdateInstanceRequest {
	s.DatabasePassword = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceAlias(v string) *UpdateInstanceRequest {
	s.InstanceAlias = &v
	return s
}

func (s *UpdateInstanceRequest) SetDbaId(v string) *UpdateInstanceRequest {
	s.DbaId = &v
	return s
}

func (s *UpdateInstanceRequest) SetSafeRuleId(v string) *UpdateInstanceRequest {
	s.SafeRuleId = &v
	return s
}

func (s *UpdateInstanceRequest) SetQueryTimeout(v int32) *UpdateInstanceRequest {
	s.QueryTimeout = &v
	return s
}

func (s *UpdateInstanceRequest) SetExportTimeout(v int32) *UpdateInstanceRequest {
	s.ExportTimeout = &v
	return s
}

func (s *UpdateInstanceRequest) SetDataLinkName(v string) *UpdateInstanceRequest {
	s.DataLinkName = &v
	return s
}

func (s *UpdateInstanceRequest) SetDdlOnline(v int32) *UpdateInstanceRequest {
	s.DdlOnline = &v
	return s
}

func (s *UpdateInstanceRequest) SetUseDsql(v int32) *UpdateInstanceRequest {
	s.UseDsql = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceId(v string) *UpdateInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetSkipTest(v bool) *UpdateInstanceRequest {
	s.SkipTest = &v
	return s
}

type UpdateInstanceResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetErrorCode(v string) *UpdateInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetErrorMessage(v string) *UpdateInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetSuccess(v bool) *UpdateInstanceResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceResponse) SetBody(v *UpdateInstanceResponseBody) *UpdateInstanceResponse {
	s.Body = v
	return s
}

type ExecuteScriptRequest struct {
	DbId   *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	Logic  *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	Tid    *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ExecuteScriptRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteScriptRequest) GoString() string {
	return s.String()
}

func (s *ExecuteScriptRequest) SetDbId(v int32) *ExecuteScriptRequest {
	s.DbId = &v
	return s
}

func (s *ExecuteScriptRequest) SetScript(v string) *ExecuteScriptRequest {
	s.Script = &v
	return s
}

func (s *ExecuteScriptRequest) SetLogic(v bool) *ExecuteScriptRequest {
	s.Logic = &v
	return s
}

func (s *ExecuteScriptRequest) SetTid(v int64) *ExecuteScriptRequest {
	s.Tid = &v
	return s
}

type ExecuteScriptResponseBody struct {
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results      []*ExecuteScriptResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	ErrorCode    *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteScriptResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteScriptResponseBody) SetRequestId(v string) *ExecuteScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteScriptResponseBody) SetResults(v []*ExecuteScriptResponseBodyResults) *ExecuteScriptResponseBody {
	s.Results = v
	return s
}

func (s *ExecuteScriptResponseBody) SetErrorCode(v string) *ExecuteScriptResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ExecuteScriptResponseBody) SetErrorMessage(v string) *ExecuteScriptResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ExecuteScriptResponseBody) SetSuccess(v bool) *ExecuteScriptResponseBody {
	s.Success = &v
	return s
}

type ExecuteScriptResponseBodyResults struct {
	ColumnNames []*string                `json:"ColumnNames,omitempty" xml:"ColumnNames,omitempty" type:"Repeated"`
	Rows        []map[string]interface{} `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	Success     *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Message     *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RowCount    *int64                   `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
}

func (s ExecuteScriptResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ExecuteScriptResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ExecuteScriptResponseBodyResults) SetColumnNames(v []*string) *ExecuteScriptResponseBodyResults {
	s.ColumnNames = v
	return s
}

func (s *ExecuteScriptResponseBodyResults) SetRows(v []map[string]interface{}) *ExecuteScriptResponseBodyResults {
	s.Rows = v
	return s
}

func (s *ExecuteScriptResponseBodyResults) SetSuccess(v bool) *ExecuteScriptResponseBodyResults {
	s.Success = &v
	return s
}

func (s *ExecuteScriptResponseBodyResults) SetMessage(v string) *ExecuteScriptResponseBodyResults {
	s.Message = &v
	return s
}

func (s *ExecuteScriptResponseBodyResults) SetRowCount(v int64) *ExecuteScriptResponseBodyResults {
	s.RowCount = &v
	return s
}

type ExecuteScriptResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteScriptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteScriptResponse) GoString() string {
	return s.String()
}

func (s *ExecuteScriptResponse) SetHeaders(v map[string]*string) *ExecuteScriptResponse {
	s.Headers = v
	return s
}

func (s *ExecuteScriptResponse) SetBody(v *ExecuteScriptResponseBody) *ExecuteScriptResponse {
	s.Body = v
	return s
}

type ListDBTaskSQLJobDetailRequest struct {
	JobId      *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tid        *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListDBTaskSQLJobDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDBTaskSQLJobDetailRequest) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobDetailRequest) SetJobId(v int64) *ListDBTaskSQLJobDetailRequest {
	s.JobId = &v
	return s
}

func (s *ListDBTaskSQLJobDetailRequest) SetPageNumber(v int64) *ListDBTaskSQLJobDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDBTaskSQLJobDetailRequest) SetPageSize(v int64) *ListDBTaskSQLJobDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListDBTaskSQLJobDetailRequest) SetTid(v int64) *ListDBTaskSQLJobDetailRequest {
	s.Tid = &v
	return s
}

type ListDBTaskSQLJobDetailResponseBody struct {
	// Id of the request
	RequestId              *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                *bool                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage           *string                                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode              *string                                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	DBTaskSQLJobDetailList []*ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList `json:"DBTaskSQLJobDetailList,omitempty" xml:"DBTaskSQLJobDetailList,omitempty" type:"Repeated"`
	TotalCount             *int64                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDBTaskSQLJobDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDBTaskSQLJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobDetailResponseBody) SetRequestId(v string) *ListDBTaskSQLJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBody) SetSuccess(v bool) *ListDBTaskSQLJobDetailResponseBody {
	s.Success = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBody) SetErrorMessage(v string) *ListDBTaskSQLJobDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBody) SetErrorCode(v string) *ListDBTaskSQLJobDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBody) SetDBTaskSQLJobDetailList(v []*ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) *ListDBTaskSQLJobDetailResponseBody {
	s.DBTaskSQLJobDetailList = v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBody) SetTotalCount(v int64) *ListDBTaskSQLJobDetailResponseBody {
	s.TotalCount = &v
	return s
}

type ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList struct {
	JobDetailId  *int64  `json:"JobDetailId,omitempty" xml:"JobDetailId,omitempty"`
	CurrentSql   *string `json:"CurrentSql,omitempty" xml:"CurrentSql,omitempty"`
	ExecuteCount *int64  `json:"ExecuteCount,omitempty" xml:"ExecuteCount,omitempty"`
	JobId        *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	DbId         *int64  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Logic        *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	Skip         *bool   `json:"Skip,omitempty" xml:"Skip,omitempty"`
	SqlType      *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) String() string {
	return tea.Prettify(s)
}

func (s ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetJobDetailId(v int64) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.JobDetailId = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetCurrentSql(v string) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.CurrentSql = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetExecuteCount(v int64) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.ExecuteCount = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetJobId(v int64) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.JobId = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetDbId(v int64) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.DbId = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetLogic(v bool) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.Logic = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetSkip(v bool) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.Skip = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetSqlType(v string) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.SqlType = &v
	return s
}

func (s *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList) SetStatus(v string) *ListDBTaskSQLJobDetailResponseBodyDBTaskSQLJobDetailList {
	s.Status = &v
	return s
}

type ListDBTaskSQLJobDetailResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDBTaskSQLJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDBTaskSQLJobDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDBTaskSQLJobDetailResponse) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobDetailResponse) SetHeaders(v map[string]*string) *ListDBTaskSQLJobDetailResponse {
	s.Headers = v
	return s
}

func (s *ListDBTaskSQLJobDetailResponse) SetBody(v *ListDBTaskSQLJobDetailResponseBody) *ListDBTaskSQLJobDetailResponse {
	s.Body = v
	return s
}

type DisableUserRequest struct {
	Tid *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DisableUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableUserRequest) GoString() string {
	return s.String()
}

func (s *DisableUserRequest) SetTid(v int64) *DisableUserRequest {
	s.Tid = &v
	return s
}

func (s *DisableUserRequest) SetUid(v string) *DisableUserRequest {
	s.Uid = &v
	return s
}

type DisableUserResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableUserResponseBody) GoString() string {
	return s.String()
}

func (s *DisableUserResponseBody) SetRequestId(v string) *DisableUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableUserResponseBody) SetErrorCode(v string) *DisableUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DisableUserResponseBody) SetErrorMessage(v string) *DisableUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DisableUserResponseBody) SetSuccess(v bool) *DisableUserResponseBody {
	s.Success = &v
	return s
}

type DisableUserResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableUserResponse) GoString() string {
	return s.String()
}

func (s *DisableUserResponse) SetHeaders(v map[string]*string) *DisableUserResponse {
	s.Headers = v
	return s
}

func (s *DisableUserResponse) SetBody(v *DisableUserResponseBody) *DisableUserResponse {
	s.Body = v
	return s
}

type GetApprovalDetailRequest struct {
	Tid                *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	WorkflowInstanceId *int64 `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
}

func (s GetApprovalDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApprovalDetailRequest) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailRequest) SetTid(v int64) *GetApprovalDetailRequest {
	s.Tid = &v
	return s
}

func (s *GetApprovalDetailRequest) SetWorkflowInstanceId(v int64) *GetApprovalDetailRequest {
	s.WorkflowInstanceId = &v
	return s
}

type GetApprovalDetailResponseBody struct {
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode      *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ApprovalDetail *GetApprovalDetailResponseBodyApprovalDetail `json:"ApprovalDetail,omitempty" xml:"ApprovalDetail,omitempty" type:"Struct"`
	Success        *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetApprovalDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApprovalDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBody) SetRequestId(v string) *GetApprovalDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApprovalDetailResponseBody) SetErrorCode(v string) *GetApprovalDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetApprovalDetailResponseBody) SetErrorMessage(v string) *GetApprovalDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetApprovalDetailResponseBody) SetApprovalDetail(v *GetApprovalDetailResponseBodyApprovalDetail) *GetApprovalDetailResponseBody {
	s.ApprovalDetail = v
	return s
}

func (s *GetApprovalDetailResponseBody) SetSuccess(v bool) *GetApprovalDetailResponseBody {
	s.Success = &v
	return s
}

type GetApprovalDetailResponseBodyApprovalDetail struct {
	WorkflowNodes   *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes   `json:"WorkflowNodes,omitempty" xml:"WorkflowNodes,omitempty" type:"Struct"`
	Description     *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	CurrentHandlers *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers `json:"CurrentHandlers,omitempty" xml:"CurrentHandlers,omitempty" type:"Struct"`
	OrderType       *string                                                     `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	Title           *string                                                     `json:"Title,omitempty" xml:"Title,omitempty"`
	AuditId         *int64                                                      `json:"AuditId,omitempty" xml:"AuditId,omitempty"`
	OrderId         *int64                                                      `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	WorkflowInsCode *string                                                     `json:"WorkflowInsCode,omitempty" xml:"WorkflowInsCode,omitempty"`
	ReasonList      *GetApprovalDetailResponseBodyApprovalDetailReasonList      `json:"ReasonList,omitempty" xml:"ReasonList,omitempty" type:"Struct"`
}

func (s GetApprovalDetailResponseBodyApprovalDetail) String() string {
	return tea.Prettify(s)
}

func (s GetApprovalDetailResponseBodyApprovalDetail) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetWorkflowNodes(v *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes) *GetApprovalDetailResponseBodyApprovalDetail {
	s.WorkflowNodes = v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetDescription(v string) *GetApprovalDetailResponseBodyApprovalDetail {
	s.Description = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetCurrentHandlers(v *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers) *GetApprovalDetailResponseBodyApprovalDetail {
	s.CurrentHandlers = v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetOrderType(v string) *GetApprovalDetailResponseBodyApprovalDetail {
	s.OrderType = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetTitle(v string) *GetApprovalDetailResponseBodyApprovalDetail {
	s.Title = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetAuditId(v int64) *GetApprovalDetailResponseBodyApprovalDetail {
	s.AuditId = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetOrderId(v int64) *GetApprovalDetailResponseBodyApprovalDetail {
	s.OrderId = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetWorkflowInsCode(v string) *GetApprovalDetailResponseBodyApprovalDetail {
	s.WorkflowInsCode = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetail) SetReasonList(v *GetApprovalDetailResponseBodyApprovalDetailReasonList) *GetApprovalDetailResponseBodyApprovalDetail {
	s.ReasonList = v
	return s
}

type GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes struct {
	WorkflowNode []*GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode `json:"WorkflowNode,omitempty" xml:"WorkflowNode,omitempty" type:"Repeated"`
}

func (s GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes) String() string {
	return tea.Prettify(s)
}

func (s GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes) SetWorkflowNode(v []*GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodes {
	s.WorkflowNode = v
	return s
}

type GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode struct {
	OperateTime     *string                                                                              `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	OperatorId      *int64                                                                               `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	NodeName        *string                                                                              `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	AuditUserIdList *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList `json:"AuditUserIdList,omitempty" xml:"AuditUserIdList,omitempty" type:"Struct"`
	OperateComment  *string                                                                              `json:"OperateComment,omitempty" xml:"OperateComment,omitempty"`
	WorkflowInsCode *string                                                                              `json:"WorkflowInsCode,omitempty" xml:"WorkflowInsCode,omitempty"`
}

func (s GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) String() string {
	return tea.Prettify(s)
}

func (s GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) SetOperateTime(v string) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode {
	s.OperateTime = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) SetOperatorId(v int64) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode {
	s.OperatorId = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) SetNodeName(v string) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode {
	s.NodeName = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) SetAuditUserIdList(v *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode {
	s.AuditUserIdList = v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) SetOperateComment(v string) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode {
	s.OperateComment = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode) SetWorkflowInsCode(v string) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNode {
	s.WorkflowInsCode = &v
	return s
}

type GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList struct {
	AuditUserIds []*string `json:"AuditUserIds,omitempty" xml:"AuditUserIds,omitempty" type:"Repeated"`
}

func (s GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList) String() string {
	return tea.Prettify(s)
}

func (s GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList) SetAuditUserIds(v []*string) *GetApprovalDetailResponseBodyApprovalDetailWorkflowNodesWorkflowNodeAuditUserIdList {
	s.AuditUserIds = v
	return s
}

type GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers struct {
	CurrentHandler []*GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler `json:"CurrentHandler,omitempty" xml:"CurrentHandler,omitempty" type:"Repeated"`
}

func (s GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers) String() string {
	return tea.Prettify(s)
}

func (s GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers) SetCurrentHandler(v []*GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler) *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlers {
	s.CurrentHandler = v
	return s
}

type GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler struct {
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler) String() string {
	return tea.Prettify(s)
}

func (s GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler) SetNickName(v string) *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler {
	s.NickName = &v
	return s
}

func (s *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler) SetId(v int64) *GetApprovalDetailResponseBodyApprovalDetailCurrentHandlersCurrentHandler {
	s.Id = &v
	return s
}

type GetApprovalDetailResponseBodyApprovalDetailReasonList struct {
	Reasons []*string `json:"Reasons,omitempty" xml:"Reasons,omitempty" type:"Repeated"`
}

func (s GetApprovalDetailResponseBodyApprovalDetailReasonList) String() string {
	return tea.Prettify(s)
}

func (s GetApprovalDetailResponseBodyApprovalDetailReasonList) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponseBodyApprovalDetailReasonList) SetReasons(v []*string) *GetApprovalDetailResponseBodyApprovalDetailReasonList {
	s.Reasons = v
	return s
}

type GetApprovalDetailResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetApprovalDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetApprovalDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApprovalDetailResponse) GoString() string {
	return s.String()
}

func (s *GetApprovalDetailResponse) SetHeaders(v map[string]*string) *GetApprovalDetailResponse {
	s.Headers = v
	return s
}

func (s *GetApprovalDetailResponse) SetBody(v *GetApprovalDetailResponseBody) *GetApprovalDetailResponse {
	s.Body = v
	return s
}

type ListSQLReviewOriginSQLRequest struct {
	OrderId           *int64                                          `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderActionDetail *ListSQLReviewOriginSQLRequestOrderActionDetail `json:"OrderActionDetail,omitempty" xml:"OrderActionDetail,omitempty" type:"Struct"`
	Tid               *int64                                          `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListSQLReviewOriginSQLRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSQLReviewOriginSQLRequest) GoString() string {
	return s.String()
}

func (s *ListSQLReviewOriginSQLRequest) SetOrderId(v int64) *ListSQLReviewOriginSQLRequest {
	s.OrderId = &v
	return s
}

func (s *ListSQLReviewOriginSQLRequest) SetOrderActionDetail(v *ListSQLReviewOriginSQLRequestOrderActionDetail) *ListSQLReviewOriginSQLRequest {
	s.OrderActionDetail = v
	return s
}

func (s *ListSQLReviewOriginSQLRequest) SetTid(v int64) *ListSQLReviewOriginSQLRequest {
	s.Tid = &v
	return s
}

type ListSQLReviewOriginSQLRequestOrderActionDetail struct {
	FileId            *int64                                              `json:"FileId,omitempty" xml:"FileId,omitempty"`
	SQLReviewResult   *string                                             `json:"SQLReviewResult,omitempty" xml:"SQLReviewResult,omitempty"`
	CheckStatusResult *string                                             `json:"CheckStatusResult,omitempty" xml:"CheckStatusResult,omitempty"`
	Page              *ListSQLReviewOriginSQLRequestOrderActionDetailPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
}

func (s ListSQLReviewOriginSQLRequestOrderActionDetail) String() string {
	return tea.Prettify(s)
}

func (s ListSQLReviewOriginSQLRequestOrderActionDetail) GoString() string {
	return s.String()
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetail) SetFileId(v int64) *ListSQLReviewOriginSQLRequestOrderActionDetail {
	s.FileId = &v
	return s
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetail) SetSQLReviewResult(v string) *ListSQLReviewOriginSQLRequestOrderActionDetail {
	s.SQLReviewResult = &v
	return s
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetail) SetCheckStatusResult(v string) *ListSQLReviewOriginSQLRequestOrderActionDetail {
	s.CheckStatusResult = &v
	return s
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetail) SetPage(v *ListSQLReviewOriginSQLRequestOrderActionDetailPage) *ListSQLReviewOriginSQLRequestOrderActionDetail {
	s.Page = v
	return s
}

type ListSQLReviewOriginSQLRequestOrderActionDetailPage struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSQLReviewOriginSQLRequestOrderActionDetailPage) String() string {
	return tea.Prettify(s)
}

func (s ListSQLReviewOriginSQLRequestOrderActionDetailPage) GoString() string {
	return s.String()
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetailPage) SetPageNumber(v int32) *ListSQLReviewOriginSQLRequestOrderActionDetailPage {
	s.PageNumber = &v
	return s
}

func (s *ListSQLReviewOriginSQLRequestOrderActionDetailPage) SetPageSize(v int32) *ListSQLReviewOriginSQLRequestOrderActionDetailPage {
	s.PageSize = &v
	return s
}

type ListSQLReviewOriginSQLShrinkRequest struct {
	OrderId                 *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderActionDetailShrink *string `json:"OrderActionDetail,omitempty" xml:"OrderActionDetail,omitempty"`
	Tid                     *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListSQLReviewOriginSQLShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSQLReviewOriginSQLShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSQLReviewOriginSQLShrinkRequest) SetOrderId(v int64) *ListSQLReviewOriginSQLShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *ListSQLReviewOriginSQLShrinkRequest) SetOrderActionDetailShrink(v string) *ListSQLReviewOriginSQLShrinkRequest {
	s.OrderActionDetailShrink = &v
	return s
}

func (s *ListSQLReviewOriginSQLShrinkRequest) SetTid(v int64) *ListSQLReviewOriginSQLShrinkRequest {
	s.Tid = &v
	return s
}

type ListSQLReviewOriginSQLResponseBody struct {
	RequestId     *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode     *string                                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage  *string                                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success       *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount    *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	OriginSQLList []*ListSQLReviewOriginSQLResponseBodyOriginSQLList `json:"OriginSQLList,omitempty" xml:"OriginSQLList,omitempty" type:"Repeated"`
}

func (s ListSQLReviewOriginSQLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSQLReviewOriginSQLResponseBody) GoString() string {
	return s.String()
}

func (s *ListSQLReviewOriginSQLResponseBody) SetRequestId(v string) *ListSQLReviewOriginSQLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBody) SetErrorCode(v string) *ListSQLReviewOriginSQLResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBody) SetErrorMessage(v string) *ListSQLReviewOriginSQLResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBody) SetSuccess(v bool) *ListSQLReviewOriginSQLResponseBody {
	s.Success = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBody) SetTotalCount(v int32) *ListSQLReviewOriginSQLResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBody) SetOriginSQLList(v []*ListSQLReviewOriginSQLResponseBodyOriginSQLList) *ListSQLReviewOriginSQLResponseBody {
	s.OriginSQLList = v
	return s
}

type ListSQLReviewOriginSQLResponseBodyOriginSQLList struct {
	SQLId             *int64  `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	FileId            *int64  `json:"FileId,omitempty" xml:"FileId,omitempty"`
	FileName          *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	SQLContent        *string `json:"SQLContent,omitempty" xml:"SQLContent,omitempty"`
	CheckStatus       *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	StatusDesc        *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	CheckedTime       *string `json:"CheckedTime,omitempty" xml:"CheckedTime,omitempty"`
	SqlHash           *string `json:"SqlHash,omitempty" xml:"SqlHash,omitempty"`
	ReviewSummary     *string `json:"ReviewSummary,omitempty" xml:"ReviewSummary,omitempty"`
	SQLReviewQueryKey *string `json:"SQLReviewQueryKey,omitempty" xml:"SQLReviewQueryKey,omitempty"`
}

func (s ListSQLReviewOriginSQLResponseBodyOriginSQLList) String() string {
	return tea.Prettify(s)
}

func (s ListSQLReviewOriginSQLResponseBodyOriginSQLList) GoString() string {
	return s.String()
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetSQLId(v int64) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.SQLId = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetFileId(v int64) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.FileId = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetFileName(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.FileName = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetSQLContent(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.SQLContent = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetCheckStatus(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.CheckStatus = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetStatusDesc(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.StatusDesc = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetCheckedTime(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.CheckedTime = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetSqlHash(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.SqlHash = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetReviewSummary(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.ReviewSummary = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetSQLReviewQueryKey(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.SQLReviewQueryKey = &v
	return s
}

type ListSQLReviewOriginSQLResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSQLReviewOriginSQLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSQLReviewOriginSQLResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSQLReviewOriginSQLResponse) GoString() string {
	return s.String()
}

func (s *ListSQLReviewOriginSQLResponse) SetHeaders(v map[string]*string) *ListSQLReviewOriginSQLResponse {
	s.Headers = v
	return s
}

func (s *ListSQLReviewOriginSQLResponse) SetBody(v *ListSQLReviewOriginSQLResponseBody) *ListSQLReviewOriginSQLResponse {
	s.Body = v
	return s
}

type GetUserActiveTenantRequest struct {
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetUserActiveTenantRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserActiveTenantRequest) GoString() string {
	return s.String()
}

func (s *GetUserActiveTenantRequest) SetTid(v int64) *GetUserActiveTenantRequest {
	s.Tid = &v
	return s
}

type GetUserActiveTenantResponseBody struct {
	Tenant       *GetUserActiveTenantResponseBodyTenant `json:"Tenant,omitempty" xml:"Tenant,omitempty" type:"Struct"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserActiveTenantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserActiveTenantResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserActiveTenantResponseBody) SetTenant(v *GetUserActiveTenantResponseBodyTenant) *GetUserActiveTenantResponseBody {
	s.Tenant = v
	return s
}

func (s *GetUserActiveTenantResponseBody) SetRequestId(v string) *GetUserActiveTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserActiveTenantResponseBody) SetErrorCode(v string) *GetUserActiveTenantResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserActiveTenantResponseBody) SetErrorMessage(v string) *GetUserActiveTenantResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserActiveTenantResponseBody) SetSuccess(v bool) *GetUserActiveTenantResponseBody {
	s.Success = &v
	return s
}

type GetUserActiveTenantResponseBodyTenant struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Tid        *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s GetUserActiveTenantResponseBodyTenant) String() string {
	return tea.Prettify(s)
}

func (s GetUserActiveTenantResponseBodyTenant) GoString() string {
	return s.String()
}

func (s *GetUserActiveTenantResponseBodyTenant) SetStatus(v string) *GetUserActiveTenantResponseBodyTenant {
	s.Status = &v
	return s
}

func (s *GetUserActiveTenantResponseBodyTenant) SetTid(v int64) *GetUserActiveTenantResponseBodyTenant {
	s.Tid = &v
	return s
}

func (s *GetUserActiveTenantResponseBodyTenant) SetTenantName(v string) *GetUserActiveTenantResponseBodyTenant {
	s.TenantName = &v
	return s
}

type GetUserActiveTenantResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserActiveTenantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserActiveTenantResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserActiveTenantResponse) GoString() string {
	return s.String()
}

func (s *GetUserActiveTenantResponse) SetHeaders(v map[string]*string) *GetUserActiveTenantResponse {
	s.Headers = v
	return s
}

func (s *GetUserActiveTenantResponse) SetBody(v *GetUserActiveTenantResponseBody) *GetUserActiveTenantResponse {
	s.Body = v
	return s
}

type RegisterUserRequest struct {
	Tid       *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	Uid       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserNick  *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
	RoleNames *string `json:"RoleNames,omitempty" xml:"RoleNames,omitempty"`
	Mobile    *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
}

func (s RegisterUserRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterUserRequest) GoString() string {
	return s.String()
}

func (s *RegisterUserRequest) SetTid(v int64) *RegisterUserRequest {
	s.Tid = &v
	return s
}

func (s *RegisterUserRequest) SetUid(v string) *RegisterUserRequest {
	s.Uid = &v
	return s
}

func (s *RegisterUserRequest) SetUserNick(v string) *RegisterUserRequest {
	s.UserNick = &v
	return s
}

func (s *RegisterUserRequest) SetRoleNames(v string) *RegisterUserRequest {
	s.RoleNames = &v
	return s
}

func (s *RegisterUserRequest) SetMobile(v string) *RegisterUserRequest {
	s.Mobile = &v
	return s
}

type RegisterUserResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RegisterUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterUserResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterUserResponseBody) SetRequestId(v string) *RegisterUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterUserResponseBody) SetErrorCode(v string) *RegisterUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RegisterUserResponseBody) SetErrorMessage(v string) *RegisterUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RegisterUserResponseBody) SetSuccess(v bool) *RegisterUserResponseBody {
	s.Success = &v
	return s
}

type RegisterUserResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterUserResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterUserResponse) GoString() string {
	return s.String()
}

func (s *RegisterUserResponse) SetHeaders(v map[string]*string) *RegisterUserResponse {
	s.Headers = v
	return s
}

func (s *RegisterUserResponse) SetBody(v *RegisterUserResponseBody) *RegisterUserResponse {
	s.Body = v
	return s
}

type GetInstanceRequest struct {
	Tid  *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Sid  *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
}

func (s GetInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) SetTid(v int64) *GetInstanceRequest {
	s.Tid = &v
	return s
}

func (s *GetInstanceRequest) SetHost(v string) *GetInstanceRequest {
	s.Host = &v
	return s
}

func (s *GetInstanceRequest) SetPort(v int32) *GetInstanceRequest {
	s.Port = &v
	return s
}

func (s *GetInstanceRequest) SetSid(v string) *GetInstanceRequest {
	s.Sid = &v
	return s
}

type GetInstanceResponseBody struct {
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Instance     *GetInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	ErrorCode    *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstance(v *GetInstanceResponseBodyInstance) *GetInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *GetInstanceResponseBody) SetErrorCode(v string) *GetInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetErrorMessage(v string) *GetInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetInstanceResponseBody) SetSuccess(v bool) *GetInstanceResponseBody {
	s.Success = &v
	return s
}

type GetInstanceResponseBodyInstance struct {
	VpcId            *string                                       `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	DatabaseUser     *string                                       `json:"DatabaseUser,omitempty" xml:"DatabaseUser,omitempty"`
	DbaId            *string                                       `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	UseDsql          *int32                                        `json:"UseDsql,omitempty" xml:"UseDsql,omitempty"`
	Port             *int32                                        `json:"Port,omitempty" xml:"Port,omitempty"`
	EcsInstanceId    *string                                       `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	EnvType          *string                                       `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Sid              *string                                       `json:"Sid,omitempty" xml:"Sid,omitempty"`
	OwnerIdList      *GetInstanceResponseBodyInstanceOwnerIdList   `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	SafeRuleId       *string                                       `json:"SafeRuleId,omitempty" xml:"SafeRuleId,omitempty"`
	OwnerNameList    *GetInstanceResponseBodyInstanceOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	DbaNickName      *string                                       `json:"DbaNickName,omitempty" xml:"DbaNickName,omitempty"`
	QueryTimeout     *int32                                        `json:"QueryTimeout,omitempty" xml:"QueryTimeout,omitempty"`
	InstanceSource   *string                                       `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	Host             *string                                       `json:"Host,omitempty" xml:"Host,omitempty"`
	State            *string                                       `json:"State,omitempty" xml:"State,omitempty"`
	DataLinkName     *string                                       `json:"DataLinkName,omitempty" xml:"DataLinkName,omitempty"`
	ExportTimeout    *int32                                        `json:"ExportTimeout,omitempty" xml:"ExportTimeout,omitempty"`
	InstanceId       *string                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType     *string                                       `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	DatabasePassword *string                                       `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	InstanceAlias    *string                                       `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	DdlOnline        *int32                                        `json:"DdlOnline,omitempty" xml:"DdlOnline,omitempty"`
	EcsRegion        *string                                       `json:"EcsRegion,omitempty" xml:"EcsRegion,omitempty"`
}

func (s GetInstanceResponseBodyInstance) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstance) SetVpcId(v string) *GetInstanceResponseBodyInstance {
	s.VpcId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDatabaseUser(v string) *GetInstanceResponseBodyInstance {
	s.DatabaseUser = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDbaId(v string) *GetInstanceResponseBodyInstance {
	s.DbaId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetUseDsql(v int32) *GetInstanceResponseBodyInstance {
	s.UseDsql = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetPort(v int32) *GetInstanceResponseBodyInstance {
	s.Port = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetEcsInstanceId(v string) *GetInstanceResponseBodyInstance {
	s.EcsInstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetEnvType(v string) *GetInstanceResponseBodyInstance {
	s.EnvType = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetSid(v string) *GetInstanceResponseBodyInstance {
	s.Sid = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetOwnerIdList(v *GetInstanceResponseBodyInstanceOwnerIdList) *GetInstanceResponseBodyInstance {
	s.OwnerIdList = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetSafeRuleId(v string) *GetInstanceResponseBodyInstance {
	s.SafeRuleId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetOwnerNameList(v *GetInstanceResponseBodyInstanceOwnerNameList) *GetInstanceResponseBodyInstance {
	s.OwnerNameList = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDbaNickName(v string) *GetInstanceResponseBodyInstance {
	s.DbaNickName = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetQueryTimeout(v int32) *GetInstanceResponseBodyInstance {
	s.QueryTimeout = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceSource(v string) *GetInstanceResponseBodyInstance {
	s.InstanceSource = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetHost(v string) *GetInstanceResponseBodyInstance {
	s.Host = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetState(v string) *GetInstanceResponseBodyInstance {
	s.State = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDataLinkName(v string) *GetInstanceResponseBodyInstance {
	s.DataLinkName = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetExportTimeout(v int32) *GetInstanceResponseBodyInstance {
	s.ExportTimeout = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceId(v string) *GetInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceType(v string) *GetInstanceResponseBodyInstance {
	s.InstanceType = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDatabasePassword(v string) *GetInstanceResponseBodyInstance {
	s.DatabasePassword = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceAlias(v string) *GetInstanceResponseBodyInstance {
	s.InstanceAlias = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDdlOnline(v int32) *GetInstanceResponseBodyInstance {
	s.DdlOnline = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetEcsRegion(v string) *GetInstanceResponseBodyInstance {
	s.EcsRegion = &v
	return s
}

type GetInstanceResponseBodyInstanceOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBodyInstanceOwnerIdList) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceOwnerIdList) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceOwnerIdList) SetOwnerIds(v []*string) *GetInstanceResponseBodyInstanceOwnerIdList {
	s.OwnerIds = v
	return s
}

type GetInstanceResponseBodyInstanceOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBodyInstanceOwnerNameList) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceOwnerNameList) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceOwnerNameList) SetOwnerNames(v []*string) *GetInstanceResponseBodyInstanceOwnerNameList {
	s.OwnerNames = v
	return s
}

type GetInstanceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type GetPermApplyOrderDetailRequest struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Tid     *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetPermApplyOrderDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPermApplyOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailRequest) SetOrderId(v int64) *GetPermApplyOrderDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetPermApplyOrderDetailRequest) SetTid(v int64) *GetPermApplyOrderDetailRequest {
	s.Tid = &v
	return s
}

type GetPermApplyOrderDetailResponseBody struct {
	// Id of the request
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success              *bool                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage         *string                                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode            *string                                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	PermApplyOrderDetail *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail `json:"PermApplyOrderDetail,omitempty" xml:"PermApplyOrderDetail,omitempty" type:"Struct"`
}

func (s GetPermApplyOrderDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBody) SetRequestId(v string) *GetPermApplyOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBody) SetSuccess(v bool) *GetPermApplyOrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBody) SetErrorMessage(v string) *GetPermApplyOrderDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBody) SetErrorCode(v string) *GetPermApplyOrderDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBody) SetPermApplyOrderDetail(v *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) *GetPermApplyOrderDetailResponseBody {
	s.PermApplyOrderDetail = v
	return s
}

type GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail struct {
	Resources []*GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	PermType  *int64                                                              `json:"PermType,omitempty" xml:"PermType,omitempty"`
	Seconds   *int64                                                              `json:"Seconds,omitempty" xml:"Seconds,omitempty"`
	ApplyType *string                                                             `json:"ApplyType,omitempty" xml:"ApplyType,omitempty"`
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) String() string {
	return tea.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) SetResources(v []*GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail {
	s.Resources = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) SetPermType(v int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail {
	s.PermType = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) SetSeconds(v int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail {
	s.Seconds = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail) SetApplyType(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetail {
	s.ApplyType = &v
	return s
}

type GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources struct {
	DatabaseInfo *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo `json:"DatabaseInfo,omitempty" xml:"DatabaseInfo,omitempty" type:"Struct"`
	TableInfo    *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo    `json:"TableInfo,omitempty" xml:"TableInfo,omitempty" type:"Struct"`
	ColumnInfo   *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo   `json:"ColumnInfo,omitempty" xml:"ColumnInfo,omitempty" type:"Struct"`
	InstanceInfo *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Struct"`
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) String() string {
	return tea.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) SetDatabaseInfo(v *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources {
	s.DatabaseInfo = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) SetTableInfo(v *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources {
	s.TableInfo = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) SetColumnInfo(v *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources {
	s.ColumnInfo = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources) SetInstanceInfo(v *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResources {
	s.InstanceInfo = v
	return s
}

type GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo struct {
	DbId           *int64    `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Logic          *bool     `json:"Logic,omitempty" xml:"Logic,omitempty"`
	DbType         *string   `json:"DbType,omitempty" xml:"DbType,omitempty"`
	SearchName     *string   `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	OwnerIds       []*int64  `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
	OwnerNickNames []*string `json:"OwnerNickNames,omitempty" xml:"OwnerNickNames,omitempty" type:"Repeated"`
	EnvType        *string   `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) SetDbId(v int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo {
	s.DbId = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) SetLogic(v bool) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo {
	s.Logic = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) SetDbType(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo {
	s.DbType = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) SetSearchName(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo {
	s.SearchName = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) SetOwnerIds(v []*int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo {
	s.OwnerIds = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) SetOwnerNickNames(v []*string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo {
	s.OwnerNickNames = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo) SetEnvType(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesDatabaseInfo {
	s.EnvType = &v
	return s
}

type GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo struct {
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo) SetTableName(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesTableInfo {
	s.TableName = &v
	return s
}

type GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo struct {
	TableName  *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo) SetTableName(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo {
	s.TableName = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo) SetColumnName(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesColumnInfo {
	s.ColumnName = &v
	return s
}

type GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo struct {
	InstanceId    *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SearchName    *string   `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	DbType        *string   `json:"DbType,omitempty" xml:"DbType,omitempty"`
	EnvType       *string   `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Host          *string   `json:"Host,omitempty" xml:"Host,omitempty"`
	Port          *int64    `json:"Port,omitempty" xml:"Port,omitempty"`
	OwnerIds      []*int64  `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
	OwnerNickName []*string `json:"OwnerNickName,omitempty" xml:"OwnerNickName,omitempty" type:"Repeated"`
	DbaId         *int64    `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	DbaNickName   *string   `json:"DbaNickName,omitempty" xml:"DbaNickName,omitempty"`
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetInstanceId(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetSearchName(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.SearchName = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetDbType(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.DbType = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetEnvType(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.EnvType = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetHost(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.Host = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetPort(v int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.Port = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetOwnerIds(v []*int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.OwnerIds = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetOwnerNickName(v []*string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.OwnerNickName = v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetDbaId(v int64) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.DbaId = &v
	return s
}

func (s *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo) SetDbaNickName(v string) *GetPermApplyOrderDetailResponseBodyPermApplyOrderDetailResourcesInstanceInfo {
	s.DbaNickName = &v
	return s
}

type GetPermApplyOrderDetailResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPermApplyOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPermApplyOrderDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPermApplyOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponse) SetHeaders(v map[string]*string) *GetPermApplyOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetPermApplyOrderDetailResponse) SetBody(v *GetPermApplyOrderDetailResponseBody) *GetPermApplyOrderDetailResponse {
	s.Body = v
	return s
}

type ListIndexesRequest struct {
	Tid     *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	Logic   *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s ListIndexesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesRequest) GoString() string {
	return s.String()
}

func (s *ListIndexesRequest) SetTid(v int64) *ListIndexesRequest {
	s.Tid = &v
	return s
}

func (s *ListIndexesRequest) SetTableId(v string) *ListIndexesRequest {
	s.TableId = &v
	return s
}

func (s *ListIndexesRequest) SetLogic(v bool) *ListIndexesRequest {
	s.Logic = &v
	return s
}

type ListIndexesResponseBody struct {
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IndexList    *ListIndexesResponseBodyIndexList `json:"IndexList,omitempty" xml:"IndexList,omitempty" type:"Struct"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIndexesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBody) SetRequestId(v string) *ListIndexesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIndexesResponseBody) SetErrorCode(v string) *ListIndexesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListIndexesResponseBody) SetErrorMessage(v string) *ListIndexesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListIndexesResponseBody) SetIndexList(v *ListIndexesResponseBodyIndexList) *ListIndexesResponseBody {
	s.IndexList = v
	return s
}

func (s *ListIndexesResponseBody) SetSuccess(v bool) *ListIndexesResponseBody {
	s.Success = &v
	return s
}

type ListIndexesResponseBodyIndexList struct {
	Index []*ListIndexesResponseBodyIndexListIndex `json:"Index,omitempty" xml:"Index,omitempty" type:"Repeated"`
}

func (s ListIndexesResponseBodyIndexList) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBodyIndexList) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyIndexList) SetIndex(v []*ListIndexesResponseBodyIndexListIndex) *ListIndexesResponseBodyIndexList {
	s.Index = v
	return s
}

type ListIndexesResponseBodyIndexListIndex struct {
	IndexName    *string `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	IndexType    *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	TableId      *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	IndexId      *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	IndexComment *string `json:"IndexComment,omitempty" xml:"IndexComment,omitempty"`
}

func (s ListIndexesResponseBodyIndexListIndex) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponseBodyIndexListIndex) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyIndexListIndex) SetIndexName(v string) *ListIndexesResponseBodyIndexListIndex {
	s.IndexName = &v
	return s
}

func (s *ListIndexesResponseBodyIndexListIndex) SetIndexType(v string) *ListIndexesResponseBodyIndexListIndex {
	s.IndexType = &v
	return s
}

func (s *ListIndexesResponseBodyIndexListIndex) SetTableId(v string) *ListIndexesResponseBodyIndexListIndex {
	s.TableId = &v
	return s
}

func (s *ListIndexesResponseBodyIndexListIndex) SetIndexId(v string) *ListIndexesResponseBodyIndexListIndex {
	s.IndexId = &v
	return s
}

func (s *ListIndexesResponseBodyIndexListIndex) SetIndexComment(v string) *ListIndexesResponseBodyIndexListIndex {
	s.IndexComment = &v
	return s
}

type ListIndexesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListIndexesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIndexesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIndexesResponse) GoString() string {
	return s.String()
}

func (s *ListIndexesResponse) SetHeaders(v map[string]*string) *ListIndexesResponse {
	s.Headers = v
	return s
}

func (s *ListIndexesResponse) SetBody(v *ListIndexesResponseBody) *ListIndexesResponse {
	s.Body = v
	return s
}

type ListLogicTablesRequest struct {
	Tid        *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	ReturnGuid *bool   `json:"ReturnGuid,omitempty" xml:"ReturnGuid,omitempty"`
}

func (s ListLogicTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLogicTablesRequest) GoString() string {
	return s.String()
}

func (s *ListLogicTablesRequest) SetTid(v int64) *ListLogicTablesRequest {
	s.Tid = &v
	return s
}

func (s *ListLogicTablesRequest) SetDatabaseId(v string) *ListLogicTablesRequest {
	s.DatabaseId = &v
	return s
}

func (s *ListLogicTablesRequest) SetPageNumber(v int32) *ListLogicTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLogicTablesRequest) SetPageSize(v int32) *ListLogicTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLogicTablesRequest) SetSearchName(v string) *ListLogicTablesRequest {
	s.SearchName = &v
	return s
}

func (s *ListLogicTablesRequest) SetReturnGuid(v bool) *ListLogicTablesRequest {
	s.ReturnGuid = &v
	return s
}

type ListLogicTablesResponseBody struct {
	TotalCount     *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LogicTableList *ListLogicTablesResponseBodyLogicTableList `json:"LogicTableList,omitempty" xml:"LogicTableList,omitempty" type:"Struct"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode      *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success        *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLogicTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLogicTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogicTablesResponseBody) SetTotalCount(v int64) *ListLogicTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLogicTablesResponseBody) SetLogicTableList(v *ListLogicTablesResponseBodyLogicTableList) *ListLogicTablesResponseBody {
	s.LogicTableList = v
	return s
}

func (s *ListLogicTablesResponseBody) SetRequestId(v string) *ListLogicTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogicTablesResponseBody) SetErrorCode(v string) *ListLogicTablesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListLogicTablesResponseBody) SetErrorMessage(v string) *ListLogicTablesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListLogicTablesResponseBody) SetSuccess(v bool) *ListLogicTablesResponseBody {
	s.Success = &v
	return s
}

type ListLogicTablesResponseBodyLogicTableList struct {
	LogicTable []*ListLogicTablesResponseBodyLogicTableListLogicTable `json:"LogicTable,omitempty" xml:"LogicTable,omitempty" type:"Repeated"`
}

func (s ListLogicTablesResponseBodyLogicTableList) String() string {
	return tea.Prettify(s)
}

func (s ListLogicTablesResponseBodyLogicTableList) GoString() string {
	return s.String()
}

func (s *ListLogicTablesResponseBodyLogicTableList) SetLogicTable(v []*ListLogicTablesResponseBodyLogicTableListLogicTable) *ListLogicTablesResponseBodyLogicTableList {
	s.LogicTable = v
	return s
}

type ListLogicTablesResponseBodyLogicTableListLogicTable struct {
	DatabaseId    *string                                                           `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	TableName     *string                                                           `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableCount    *string                                                           `json:"TableCount,omitempty" xml:"TableCount,omitempty"`
	OwnerIdList   *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList   `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	OwnerNameList *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	SchemaName    *string                                                           `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	Logic         *bool                                                             `json:"Logic,omitempty" xml:"Logic,omitempty"`
	TableExpr     *string                                                           `json:"TableExpr,omitempty" xml:"TableExpr,omitempty"`
	TableGuid     *string                                                           `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	TableId       *string                                                           `json:"TableId,omitempty" xml:"TableId,omitempty"`
}

func (s ListLogicTablesResponseBodyLogicTableListLogicTable) String() string {
	return tea.Prettify(s)
}

func (s ListLogicTablesResponseBodyLogicTableListLogicTable) GoString() string {
	return s.String()
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetDatabaseId(v string) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.DatabaseId = &v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetTableName(v string) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.TableName = &v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetTableCount(v string) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.TableCount = &v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetOwnerIdList(v *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.OwnerIdList = v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetOwnerNameList(v *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.OwnerNameList = v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetSchemaName(v string) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.SchemaName = &v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetLogic(v bool) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.Logic = &v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetTableExpr(v string) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.TableExpr = &v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetTableGuid(v string) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.TableGuid = &v
	return s
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTable) SetTableId(v string) *ListLogicTablesResponseBodyLogicTableListLogicTable {
	s.TableId = &v
	return s
}

type ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList) String() string {
	return tea.Prettify(s)
}

func (s ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList) GoString() string {
	return s.String()
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList) SetOwnerIds(v []*string) *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerIdList {
	s.OwnerIds = v
	return s
}

type ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList) String() string {
	return tea.Prettify(s)
}

func (s ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList) GoString() string {
	return s.String()
}

func (s *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList) SetOwnerNames(v []*string) *ListLogicTablesResponseBodyLogicTableListLogicTableOwnerNameList {
	s.OwnerNames = v
	return s
}

type ListLogicTablesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLogicTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLogicTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLogicTablesResponse) GoString() string {
	return s.String()
}

func (s *ListLogicTablesResponse) SetHeaders(v map[string]*string) *ListLogicTablesResponse {
	s.Headers = v
	return s
}

func (s *ListLogicTablesResponse) SetBody(v *ListLogicTablesResponseBody) *ListLogicTablesResponse {
	s.Body = v
	return s
}

type GetTableTopologyRequest struct {
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	Tid       *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetTableTopologyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableTopologyRequest) GoString() string {
	return s.String()
}

func (s *GetTableTopologyRequest) SetTableGuid(v string) *GetTableTopologyRequest {
	s.TableGuid = &v
	return s
}

func (s *GetTableTopologyRequest) SetTid(v int64) *GetTableTopologyRequest {
	s.Tid = &v
	return s
}

type GetTableTopologyResponseBody struct {
	// Id of the request
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage  *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode     *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	TableTopology *GetTableTopologyResponseBodyTableTopology `json:"TableTopology,omitempty" xml:"TableTopology,omitempty" type:"Struct"`
}

func (s GetTableTopologyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableTopologyResponseBody) SetRequestId(v string) *GetTableTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableTopologyResponseBody) SetSuccess(v bool) *GetTableTopologyResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableTopologyResponseBody) SetErrorMessage(v string) *GetTableTopologyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTableTopologyResponseBody) SetErrorCode(v string) *GetTableTopologyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTableTopologyResponseBody) SetTableTopology(v *GetTableTopologyResponseBodyTableTopology) *GetTableTopologyResponseBody {
	s.TableTopology = v
	return s
}

type GetTableTopologyResponseBodyTableTopology struct {
	TableGuid             *string                                                           `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	TableName             *string                                                           `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableTopologyInfoList []*GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList `json:"TableTopologyInfoList,omitempty" xml:"TableTopologyInfoList,omitempty" type:"Repeated"`
	Logic                 *bool                                                             `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s GetTableTopologyResponseBodyTableTopology) String() string {
	return tea.Prettify(s)
}

func (s GetTableTopologyResponseBodyTableTopology) GoString() string {
	return s.String()
}

func (s *GetTableTopologyResponseBodyTableTopology) SetTableGuid(v string) *GetTableTopologyResponseBodyTableTopology {
	s.TableGuid = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopology) SetTableName(v string) *GetTableTopologyResponseBodyTableTopology {
	s.TableName = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopology) SetTableTopologyInfoList(v []*GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) *GetTableTopologyResponseBodyTableTopology {
	s.TableTopologyInfoList = v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopology) SetLogic(v bool) *GetTableTopologyResponseBodyTableTopology {
	s.Logic = &v
	return s
}

type GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList struct {
	TableNameList      *string `json:"TableNameList,omitempty" xml:"TableNameList,omitempty"`
	TableNameExpr      *string `json:"TableNameExpr,omitempty" xml:"TableNameExpr,omitempty"`
	TableCount         *int64  `json:"TableCount,omitempty" xml:"TableCount,omitempty"`
	DbId               *int64  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	DbSearchName       *string `json:"DbSearchName,omitempty" xml:"DbSearchName,omitempty"`
	InstanceId         *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceResourceId *string `json:"InstanceResourceId,omitempty" xml:"InstanceResourceId,omitempty"`
	InstanceSource     *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	DbName             *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbType             *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
}

func (s GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) GoString() string {
	return s.String()
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetTableNameList(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.TableNameList = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetTableNameExpr(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.TableNameExpr = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetTableCount(v int64) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.TableCount = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetDbId(v int64) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.DbId = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetDbSearchName(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.DbSearchName = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetInstanceId(v int64) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.InstanceId = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetRegionId(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.RegionId = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetInstanceResourceId(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.InstanceResourceId = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetInstanceSource(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.InstanceSource = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetDbName(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.DbName = &v
	return s
}

func (s *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList) SetDbType(v string) *GetTableTopologyResponseBodyTableTopologyTableTopologyInfoList {
	s.DbType = &v
	return s
}

type GetTableTopologyResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTableTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTableTopologyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableTopologyResponse) GoString() string {
	return s.String()
}

func (s *GetTableTopologyResponse) SetHeaders(v map[string]*string) *GetTableTopologyResponse {
	s.Headers = v
	return s
}

func (s *GetTableTopologyResponse) SetBody(v *GetTableTopologyResponseBody) *GetTableTopologyResponse {
	s.Body = v
	return s
}

type GetDataExportDownloadURLRequest struct {
	Tid     *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s GetDataExportDownloadURLRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataExportDownloadURLRequest) GoString() string {
	return s.String()
}

func (s *GetDataExportDownloadURLRequest) SetTid(v int64) *GetDataExportDownloadURLRequest {
	s.Tid = &v
	return s
}

func (s *GetDataExportDownloadURLRequest) SetOrderId(v int64) *GetDataExportDownloadURLRequest {
	s.OrderId = &v
	return s
}

type GetDataExportDownloadURLResponseBody struct {
	DownloadURLResult *GetDataExportDownloadURLResponseBodyDownloadURLResult `json:"DownloadURLResult,omitempty" xml:"DownloadURLResult,omitempty" type:"Struct"`
	RequestId         *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode         *string                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage      *string                                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success           *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataExportDownloadURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataExportDownloadURLResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataExportDownloadURLResponseBody) SetDownloadURLResult(v *GetDataExportDownloadURLResponseBodyDownloadURLResult) *GetDataExportDownloadURLResponseBody {
	s.DownloadURLResult = v
	return s
}

func (s *GetDataExportDownloadURLResponseBody) SetRequestId(v string) *GetDataExportDownloadURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataExportDownloadURLResponseBody) SetErrorCode(v string) *GetDataExportDownloadURLResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataExportDownloadURLResponseBody) SetErrorMessage(v string) *GetDataExportDownloadURLResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataExportDownloadURLResponseBody) SetSuccess(v bool) *GetDataExportDownloadURLResponseBody {
	s.Success = &v
	return s
}

type GetDataExportDownloadURLResponseBodyDownloadURLResult struct {
	HasResult  *bool   `json:"HasResult,omitempty" xml:"HasResult,omitempty"`
	TipMessage *string `json:"TipMessage,omitempty" xml:"TipMessage,omitempty"`
	URL        *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s GetDataExportDownloadURLResponseBodyDownloadURLResult) String() string {
	return tea.Prettify(s)
}

func (s GetDataExportDownloadURLResponseBodyDownloadURLResult) GoString() string {
	return s.String()
}

func (s *GetDataExportDownloadURLResponseBodyDownloadURLResult) SetHasResult(v bool) *GetDataExportDownloadURLResponseBodyDownloadURLResult {
	s.HasResult = &v
	return s
}

func (s *GetDataExportDownloadURLResponseBodyDownloadURLResult) SetTipMessage(v string) *GetDataExportDownloadURLResponseBodyDownloadURLResult {
	s.TipMessage = &v
	return s
}

func (s *GetDataExportDownloadURLResponseBodyDownloadURLResult) SetURL(v string) *GetDataExportDownloadURLResponseBodyDownloadURLResult {
	s.URL = &v
	return s
}

type GetDataExportDownloadURLResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDataExportDownloadURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDataExportDownloadURLResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataExportDownloadURLResponse) GoString() string {
	return s.String()
}

func (s *GetDataExportDownloadURLResponse) SetHeaders(v map[string]*string) *GetDataExportDownloadURLResponse {
	s.Headers = v
	return s
}

func (s *GetDataExportDownloadURLResponse) SetBody(v *GetDataExportDownloadURLResponseBody) *GetDataExportDownloadURLResponse {
	s.Body = v
	return s
}

type CreateDataCronClearOrderRequest struct {
	Comment         *string                               `json:"Comment,omitempty" xml:"Comment,omitempty"`
	RelatedUserList []*int64                              `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	Param           *CreateDataCronClearOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	AttachmentKey   *string                               `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	Tid             *int64                                `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataCronClearOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCronClearOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateDataCronClearOrderRequest) SetComment(v string) *CreateDataCronClearOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataCronClearOrderRequest) SetRelatedUserList(v []*int64) *CreateDataCronClearOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateDataCronClearOrderRequest) SetParam(v *CreateDataCronClearOrderRequestParam) *CreateDataCronClearOrderRequest {
	s.Param = v
	return s
}

func (s *CreateDataCronClearOrderRequest) SetAttachmentKey(v string) *CreateDataCronClearOrderRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDataCronClearOrderRequest) SetTid(v int64) *CreateDataCronClearOrderRequest {
	s.Tid = &v
	return s
}

type CreateDataCronClearOrderRequestParam struct {
	Classify          *string                                                  `json:"Classify,omitempty" xml:"Classify,omitempty"`
	DbItemList        []*CreateDataCronClearOrderRequestParamDbItemList        `json:"DbItemList,omitempty" xml:"DbItemList,omitempty" type:"Repeated"`
	CronFormat        *string                                                  `json:"CronFormat,omitempty" xml:"CronFormat,omitempty"`
	CronClearItemList []*CreateDataCronClearOrderRequestParamCronClearItemList `json:"CronClearItemList,omitempty" xml:"CronClearItemList,omitempty" type:"Repeated"`
	SpecifyDuration   *bool                                                    `json:"specifyDuration,omitempty" xml:"specifyDuration,omitempty"`
	DurationHour      *int64                                                   `json:"DurationHour,omitempty" xml:"DurationHour,omitempty"`
}

func (s CreateDataCronClearOrderRequestParam) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCronClearOrderRequestParam) GoString() string {
	return s.String()
}

func (s *CreateDataCronClearOrderRequestParam) SetClassify(v string) *CreateDataCronClearOrderRequestParam {
	s.Classify = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParam) SetDbItemList(v []*CreateDataCronClearOrderRequestParamDbItemList) *CreateDataCronClearOrderRequestParam {
	s.DbItemList = v
	return s
}

func (s *CreateDataCronClearOrderRequestParam) SetCronFormat(v string) *CreateDataCronClearOrderRequestParam {
	s.CronFormat = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParam) SetCronClearItemList(v []*CreateDataCronClearOrderRequestParamCronClearItemList) *CreateDataCronClearOrderRequestParam {
	s.CronClearItemList = v
	return s
}

func (s *CreateDataCronClearOrderRequestParam) SetSpecifyDuration(v bool) *CreateDataCronClearOrderRequestParam {
	s.SpecifyDuration = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParam) SetDurationHour(v int64) *CreateDataCronClearOrderRequestParam {
	s.DurationHour = &v
	return s
}

type CreateDataCronClearOrderRequestParamDbItemList struct {
	DbId  *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Logic *bool  `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s CreateDataCronClearOrderRequestParamDbItemList) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCronClearOrderRequestParamDbItemList) GoString() string {
	return s.String()
}

func (s *CreateDataCronClearOrderRequestParamDbItemList) SetDbId(v int64) *CreateDataCronClearOrderRequestParamDbItemList {
	s.DbId = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParamDbItemList) SetLogic(v bool) *CreateDataCronClearOrderRequestParamDbItemList {
	s.Logic = &v
	return s
}

type CreateDataCronClearOrderRequestParamCronClearItemList struct {
	TableName  *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	RemainDays *int64  `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
	TimeUnit   *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
	FilterSQL  *string `json:"FilterSQL,omitempty" xml:"FilterSQL,omitempty"`
}

func (s CreateDataCronClearOrderRequestParamCronClearItemList) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCronClearOrderRequestParamCronClearItemList) GoString() string {
	return s.String()
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) SetTableName(v string) *CreateDataCronClearOrderRequestParamCronClearItemList {
	s.TableName = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) SetColumnName(v string) *CreateDataCronClearOrderRequestParamCronClearItemList {
	s.ColumnName = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) SetRemainDays(v int64) *CreateDataCronClearOrderRequestParamCronClearItemList {
	s.RemainDays = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) SetTimeUnit(v string) *CreateDataCronClearOrderRequestParamCronClearItemList {
	s.TimeUnit = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) SetFilterSQL(v string) *CreateDataCronClearOrderRequestParamCronClearItemList {
	s.FilterSQL = &v
	return s
}

type CreateDataCronClearOrderShrinkRequest struct {
	Comment               *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	RelatedUserListShrink *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	ParamShrink           *string `json:"Param,omitempty" xml:"Param,omitempty"`
	AttachmentKey         *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	Tid                   *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataCronClearOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCronClearOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataCronClearOrderShrinkRequest) SetComment(v string) *CreateDataCronClearOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataCronClearOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateDataCronClearOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateDataCronClearOrderShrinkRequest) SetParamShrink(v string) *CreateDataCronClearOrderShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *CreateDataCronClearOrderShrinkRequest) SetAttachmentKey(v string) *CreateDataCronClearOrderShrinkRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDataCronClearOrderShrinkRequest) SetTid(v int64) *CreateDataCronClearOrderShrinkRequest {
	s.Tid = &v
	return s
}

type CreateDataCronClearOrderResponseBody struct {
	// Id of the request
	RequestId         *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success           *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage      *string  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode         *string  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	CreateOrderResult []*int64 `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Repeated"`
}

func (s CreateDataCronClearOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCronClearOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataCronClearOrderResponseBody) SetRequestId(v string) *CreateDataCronClearOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataCronClearOrderResponseBody) SetSuccess(v bool) *CreateDataCronClearOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataCronClearOrderResponseBody) SetErrorMessage(v string) *CreateDataCronClearOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataCronClearOrderResponseBody) SetErrorCode(v string) *CreateDataCronClearOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataCronClearOrderResponseBody) SetCreateOrderResult(v []*int64) *CreateDataCronClearOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

type CreateDataCronClearOrderResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDataCronClearOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataCronClearOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCronClearOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateDataCronClearOrderResponse) SetHeaders(v map[string]*string) *CreateDataCronClearOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateDataCronClearOrderResponse) SetBody(v *CreateDataCronClearOrderResponseBody) *CreateDataCronClearOrderResponse {
	s.Body = v
	return s
}

type CreatePublishGroupTaskRequest struct {
	Tid             *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	OrderId         *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	DbId            *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Logic           *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	PublishStrategy *string `json:"PublishStrategy,omitempty" xml:"PublishStrategy,omitempty"`
	PlanTime        *string `json:"PlanTime,omitempty" xml:"PlanTime,omitempty"`
}

func (s CreatePublishGroupTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePublishGroupTaskRequest) GoString() string {
	return s.String()
}

func (s *CreatePublishGroupTaskRequest) SetTid(v int64) *CreatePublishGroupTaskRequest {
	s.Tid = &v
	return s
}

func (s *CreatePublishGroupTaskRequest) SetOrderId(v int64) *CreatePublishGroupTaskRequest {
	s.OrderId = &v
	return s
}

func (s *CreatePublishGroupTaskRequest) SetDbId(v int32) *CreatePublishGroupTaskRequest {
	s.DbId = &v
	return s
}

func (s *CreatePublishGroupTaskRequest) SetLogic(v bool) *CreatePublishGroupTaskRequest {
	s.Logic = &v
	return s
}

func (s *CreatePublishGroupTaskRequest) SetPublishStrategy(v string) *CreatePublishGroupTaskRequest {
	s.PublishStrategy = &v
	return s
}

func (s *CreatePublishGroupTaskRequest) SetPlanTime(v string) *CreatePublishGroupTaskRequest {
	s.PlanTime = &v
	return s
}

type CreatePublishGroupTaskResponseBody struct {
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePublishGroupTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePublishGroupTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePublishGroupTaskResponseBody) SetTaskId(v int64) *CreatePublishGroupTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreatePublishGroupTaskResponseBody) SetRequestId(v string) *CreatePublishGroupTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePublishGroupTaskResponseBody) SetErrorCode(v string) *CreatePublishGroupTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreatePublishGroupTaskResponseBody) SetErrorMessage(v string) *CreatePublishGroupTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreatePublishGroupTaskResponseBody) SetSuccess(v bool) *CreatePublishGroupTaskResponseBody {
	s.Success = &v
	return s
}

type CreatePublishGroupTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePublishGroupTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePublishGroupTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePublishGroupTaskResponse) GoString() string {
	return s.String()
}

func (s *CreatePublishGroupTaskResponse) SetHeaders(v map[string]*string) *CreatePublishGroupTaskResponse {
	s.Headers = v
	return s
}

func (s *CreatePublishGroupTaskResponse) SetBody(v *CreatePublishGroupTaskResponseBody) *CreatePublishGroupTaskResponse {
	s.Body = v
	return s
}

type GetDatabaseRequest struct {
	Tid        *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	Host       *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port       *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Sid        *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s GetDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDatabaseRequest) GoString() string {
	return s.String()
}

func (s *GetDatabaseRequest) SetTid(v int64) *GetDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *GetDatabaseRequest) SetHost(v string) *GetDatabaseRequest {
	s.Host = &v
	return s
}

func (s *GetDatabaseRequest) SetPort(v int32) *GetDatabaseRequest {
	s.Port = &v
	return s
}

func (s *GetDatabaseRequest) SetSid(v string) *GetDatabaseRequest {
	s.Sid = &v
	return s
}

func (s *GetDatabaseRequest) SetSchemaName(v string) *GetDatabaseRequest {
	s.SchemaName = &v
	return s
}

type GetDatabaseResponseBody struct {
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Database     *GetDatabaseResponseBodyDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Struct"`
	ErrorCode    *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatabaseResponseBody) SetRequestId(v string) *GetDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatabaseResponseBody) SetDatabase(v *GetDatabaseResponseBodyDatabase) *GetDatabaseResponseBody {
	s.Database = v
	return s
}

func (s *GetDatabaseResponseBody) SetErrorCode(v string) *GetDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDatabaseResponseBody) SetErrorMessage(v string) *GetDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDatabaseResponseBody) SetSuccess(v bool) *GetDatabaseResponseBody {
	s.Success = &v
	return s
}

type GetDatabaseResponseBodyDatabase struct {
	DatabaseId    *string                                       `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	Host          *string                                       `json:"Host,omitempty" xml:"Host,omitempty"`
	CatalogName   *string                                       `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	DbaName       *string                                       `json:"DbaName,omitempty" xml:"DbaName,omitempty"`
	State         *string                                       `json:"State,omitempty" xml:"State,omitempty"`
	DbaId         *string                                       `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	SchemaName    *string                                       `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	InstanceId    *string                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Port          *int32                                        `json:"Port,omitempty" xml:"Port,omitempty"`
	EnvType       *string                                       `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Sid           *string                                       `json:"Sid,omitempty" xml:"Sid,omitempty"`
	OwnerIdList   *GetDatabaseResponseBodyDatabaseOwnerIdList   `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	Encoding      *string                                       `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	DbType        *string                                       `json:"DbType,omitempty" xml:"DbType,omitempty"`
	OwnerNameList *GetDatabaseResponseBodyDatabaseOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	SearchName    *string                                       `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s GetDatabaseResponseBodyDatabase) String() string {
	return tea.Prettify(s)
}

func (s GetDatabaseResponseBodyDatabase) GoString() string {
	return s.String()
}

func (s *GetDatabaseResponseBodyDatabase) SetDatabaseId(v string) *GetDatabaseResponseBodyDatabase {
	s.DatabaseId = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetHost(v string) *GetDatabaseResponseBodyDatabase {
	s.Host = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetCatalogName(v string) *GetDatabaseResponseBodyDatabase {
	s.CatalogName = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetDbaName(v string) *GetDatabaseResponseBodyDatabase {
	s.DbaName = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetState(v string) *GetDatabaseResponseBodyDatabase {
	s.State = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetDbaId(v string) *GetDatabaseResponseBodyDatabase {
	s.DbaId = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetSchemaName(v string) *GetDatabaseResponseBodyDatabase {
	s.SchemaName = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetInstanceId(v string) *GetDatabaseResponseBodyDatabase {
	s.InstanceId = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetPort(v int32) *GetDatabaseResponseBodyDatabase {
	s.Port = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetEnvType(v string) *GetDatabaseResponseBodyDatabase {
	s.EnvType = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetSid(v string) *GetDatabaseResponseBodyDatabase {
	s.Sid = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetOwnerIdList(v *GetDatabaseResponseBodyDatabaseOwnerIdList) *GetDatabaseResponseBodyDatabase {
	s.OwnerIdList = v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetEncoding(v string) *GetDatabaseResponseBodyDatabase {
	s.Encoding = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetDbType(v string) *GetDatabaseResponseBodyDatabase {
	s.DbType = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetOwnerNameList(v *GetDatabaseResponseBodyDatabaseOwnerNameList) *GetDatabaseResponseBodyDatabase {
	s.OwnerNameList = v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetSearchName(v string) *GetDatabaseResponseBodyDatabase {
	s.SearchName = &v
	return s
}

type GetDatabaseResponseBodyDatabaseOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s GetDatabaseResponseBodyDatabaseOwnerIdList) String() string {
	return tea.Prettify(s)
}

func (s GetDatabaseResponseBodyDatabaseOwnerIdList) GoString() string {
	return s.String()
}

func (s *GetDatabaseResponseBodyDatabaseOwnerIdList) SetOwnerIds(v []*string) *GetDatabaseResponseBodyDatabaseOwnerIdList {
	s.OwnerIds = v
	return s
}

type GetDatabaseResponseBodyDatabaseOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s GetDatabaseResponseBodyDatabaseOwnerNameList) String() string {
	return tea.Prettify(s)
}

func (s GetDatabaseResponseBodyDatabaseOwnerNameList) GoString() string {
	return s.String()
}

func (s *GetDatabaseResponseBodyDatabaseOwnerNameList) SetOwnerNames(v []*string) *GetDatabaseResponseBodyDatabaseOwnerNameList {
	s.OwnerNames = v
	return s
}

type GetDatabaseResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDatabaseResponse) GoString() string {
	return s.String()
}

func (s *GetDatabaseResponse) SetHeaders(v map[string]*string) *GetDatabaseResponse {
	s.Headers = v
	return s
}

func (s *GetDatabaseResponse) SetBody(v *GetDatabaseResponseBody) *GetDatabaseResponse {
	s.Body = v
	return s
}

type GetOwnerApplyOrderDetailRequest struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Tid     *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetOwnerApplyOrderDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOwnerApplyOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetOwnerApplyOrderDetailRequest) SetOrderId(v int64) *GetOwnerApplyOrderDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetOwnerApplyOrderDetailRequest) SetTid(v int64) *GetOwnerApplyOrderDetailRequest {
	s.Tid = &v
	return s
}

type GetOwnerApplyOrderDetailResponseBody struct {
	// Id of the request
	RequestId             *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success               *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage          *string                                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode             *string                                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	OwnerApplyOrderDetail *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail `json:"OwnerApplyOrderDetail,omitempty" xml:"OwnerApplyOrderDetail,omitempty" type:"Struct"`
}

func (s GetOwnerApplyOrderDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOwnerApplyOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetOwnerApplyOrderDetailResponseBody) SetRequestId(v string) *GetOwnerApplyOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBody) SetSuccess(v bool) *GetOwnerApplyOrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBody) SetErrorMessage(v string) *GetOwnerApplyOrderDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBody) SetErrorCode(v string) *GetOwnerApplyOrderDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBody) SetOwnerApplyOrderDetail(v *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail) *GetOwnerApplyOrderDetailResponseBody {
	s.OwnerApplyOrderDetail = v
	return s
}

type GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail struct {
	ApplyType *string                                                               `json:"ApplyType,omitempty" xml:"ApplyType,omitempty"`
	Resources []*GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail) String() string {
	return tea.Prettify(s)
}

func (s GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail) GoString() string {
	return s.String()
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail) SetApplyType(v string) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail {
	s.ApplyType = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail) SetResources(v []*GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetail {
	s.Resources = v
	return s
}

type GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources struct {
	Logic          *bool                                                                             `json:"Logic,omitempty" xml:"Logic,omitempty"`
	TargetId       *string                                                                           `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	ResourceDetail *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail `json:"ResourceDetail,omitempty" xml:"ResourceDetail,omitempty" type:"Struct"`
}

func (s GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) String() string {
	return tea.Prettify(s)
}

func (s GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) GoString() string {
	return s.String()
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) SetLogic(v bool) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources {
	s.Logic = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) SetTargetId(v string) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources {
	s.TargetId = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources) SetResourceDetail(v *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResources {
	s.ResourceDetail = v
	return s
}

type GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail struct {
	SearchName     *string   `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	DbType         *string   `json:"DbType,omitempty" xml:"DbType,omitempty"`
	EnvType        *string   `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	TableName      *string   `json:"TableName,omitempty" xml:"TableName,omitempty"`
	OwnerIds       []*int64  `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
	OwnerNickNames []*string `json:"OwnerNickNames,omitempty" xml:"OwnerNickNames,omitempty" type:"Repeated"`
}

func (s GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) String() string {
	return tea.Prettify(s)
}

func (s GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) GoString() string {
	return s.String()
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) SetSearchName(v string) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail {
	s.SearchName = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) SetDbType(v string) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail {
	s.DbType = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) SetEnvType(v string) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail {
	s.EnvType = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) SetTableName(v string) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail {
	s.TableName = &v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) SetOwnerIds(v []*int64) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail {
	s.OwnerIds = v
	return s
}

func (s *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail) SetOwnerNickNames(v []*string) *GetOwnerApplyOrderDetailResponseBodyOwnerApplyOrderDetailResourcesResourceDetail {
	s.OwnerNickNames = v
	return s
}

type GetOwnerApplyOrderDetailResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOwnerApplyOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOwnerApplyOrderDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOwnerApplyOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetOwnerApplyOrderDetailResponse) SetHeaders(v map[string]*string) *GetOwnerApplyOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetOwnerApplyOrderDetailResponse) SetBody(v *GetOwnerApplyOrderDetailResponseBody) *GetOwnerApplyOrderDetailResponse {
	s.Body = v
	return s
}

type GetOpLogRequest struct {
	Tid        *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	Module     *string `json:"Module,omitempty" xml:"Module,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s GetOpLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOpLogRequest) GoString() string {
	return s.String()
}

func (s *GetOpLogRequest) SetTid(v int64) *GetOpLogRequest {
	s.Tid = &v
	return s
}

func (s *GetOpLogRequest) SetModule(v string) *GetOpLogRequest {
	s.Module = &v
	return s
}

func (s *GetOpLogRequest) SetStartTime(v string) *GetOpLogRequest {
	s.StartTime = &v
	return s
}

func (s *GetOpLogRequest) SetEndTime(v string) *GetOpLogRequest {
	s.EndTime = &v
	return s
}

func (s *GetOpLogRequest) SetPageSize(v int32) *GetOpLogRequest {
	s.PageSize = &v
	return s
}

func (s *GetOpLogRequest) SetPageNumber(v int32) *GetOpLogRequest {
	s.PageNumber = &v
	return s
}

type GetOpLogResponseBody struct {
	TotalCount   *int64                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	OpLogDetails *GetOpLogResponseBodyOpLogDetails `json:"OpLogDetails,omitempty" xml:"OpLogDetails,omitempty" type:"Struct"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOpLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOpLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpLogResponseBody) SetTotalCount(v int64) *GetOpLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetOpLogResponseBody) SetRequestId(v string) *GetOpLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpLogResponseBody) SetErrorCode(v string) *GetOpLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOpLogResponseBody) SetOpLogDetails(v *GetOpLogResponseBodyOpLogDetails) *GetOpLogResponseBody {
	s.OpLogDetails = v
	return s
}

func (s *GetOpLogResponseBody) SetErrorMessage(v string) *GetOpLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOpLogResponseBody) SetSuccess(v bool) *GetOpLogResponseBody {
	s.Success = &v
	return s
}

type GetOpLogResponseBodyOpLogDetails struct {
	OpLogDetail []*GetOpLogResponseBodyOpLogDetailsOpLogDetail `json:"OpLogDetail,omitempty" xml:"OpLogDetail,omitempty" type:"Repeated"`
}

func (s GetOpLogResponseBodyOpLogDetails) String() string {
	return tea.Prettify(s)
}

func (s GetOpLogResponseBodyOpLogDetails) GoString() string {
	return s.String()
}

func (s *GetOpLogResponseBodyOpLogDetails) SetOpLogDetail(v []*GetOpLogResponseBodyOpLogDetailsOpLogDetail) *GetOpLogResponseBodyOpLogDetails {
	s.OpLogDetail = v
	return s
}

type GetOpLogResponseBodyOpLogDetailsOpLogDetail struct {
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	Database  *string `json:"Database,omitempty" xml:"Database,omitempty"`
	UserId    *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	OpContent *string `json:"OpContent,omitempty" xml:"OpContent,omitempty"`
	UserNick  *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OpTime    *string `json:"OpTime,omitempty" xml:"OpTime,omitempty"`
}

func (s GetOpLogResponseBodyOpLogDetailsOpLogDetail) String() string {
	return tea.Prettify(s)
}

func (s GetOpLogResponseBodyOpLogDetailsOpLogDetail) GoString() string {
	return s.String()
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) SetModule(v string) *GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	s.Module = &v
	return s
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) SetDatabase(v string) *GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	s.Database = &v
	return s
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) SetUserId(v int64) *GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	s.UserId = &v
	return s
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) SetOpContent(v string) *GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	s.OpContent = &v
	return s
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) SetUserNick(v string) *GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	s.UserNick = &v
	return s
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) SetOrderId(v int64) *GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	s.OrderId = &v
	return s
}

func (s *GetOpLogResponseBodyOpLogDetailsOpLogDetail) SetOpTime(v string) *GetOpLogResponseBodyOpLogDetailsOpLogDetail {
	s.OpTime = &v
	return s
}

type GetOpLogResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOpLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOpLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpLogResponse) GoString() string {
	return s.String()
}

func (s *GetOpLogResponse) SetHeaders(v map[string]*string) *GetOpLogResponse {
	s.Headers = v
	return s
}

func (s *GetOpLogResponse) SetBody(v *GetOpLogResponseBody) *GetOpLogResponse {
	s.Body = v
	return s
}

type SearchTableRequest struct {
	Tid          *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	SearchKey    *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	EnvType      *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	SearchRange  *string `json:"SearchRange,omitempty" xml:"SearchRange,omitempty"`
	SearchTarget *string `json:"SearchTarget,omitempty" xml:"SearchTarget,omitempty"`
	DbType       *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	ReturnGuid   *bool   `json:"ReturnGuid,omitempty" xml:"ReturnGuid,omitempty"`
}

func (s SearchTableRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTableRequest) GoString() string {
	return s.String()
}

func (s *SearchTableRequest) SetTid(v int64) *SearchTableRequest {
	s.Tid = &v
	return s
}

func (s *SearchTableRequest) SetSearchKey(v string) *SearchTableRequest {
	s.SearchKey = &v
	return s
}

func (s *SearchTableRequest) SetPageNumber(v int32) *SearchTableRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTableRequest) SetPageSize(v int32) *SearchTableRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTableRequest) SetEnvType(v string) *SearchTableRequest {
	s.EnvType = &v
	return s
}

func (s *SearchTableRequest) SetSearchRange(v string) *SearchTableRequest {
	s.SearchRange = &v
	return s
}

func (s *SearchTableRequest) SetSearchTarget(v string) *SearchTableRequest {
	s.SearchTarget = &v
	return s
}

func (s *SearchTableRequest) SetDbType(v string) *SearchTableRequest {
	s.DbType = &v
	return s
}

func (s *SearchTableRequest) SetReturnGuid(v bool) *SearchTableRequest {
	s.ReturnGuid = &v
	return s
}

type SearchTableResponseBody struct {
	SearchTableList *SearchTableResponseBodySearchTableList `json:"SearchTableList,omitempty" xml:"SearchTableList,omitempty" type:"Struct"`
	TotalCount      *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId       *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode       *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage    *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success         *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTableResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTableResponseBody) SetSearchTableList(v *SearchTableResponseBodySearchTableList) *SearchTableResponseBody {
	s.SearchTableList = v
	return s
}

func (s *SearchTableResponseBody) SetTotalCount(v int64) *SearchTableResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchTableResponseBody) SetRequestId(v string) *SearchTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTableResponseBody) SetErrorCode(v string) *SearchTableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchTableResponseBody) SetErrorMessage(v string) *SearchTableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SearchTableResponseBody) SetSuccess(v bool) *SearchTableResponseBody {
	s.Success = &v
	return s
}

type SearchTableResponseBodySearchTableList struct {
	SearchTable []*SearchTableResponseBodySearchTableListSearchTable `json:"SearchTable,omitempty" xml:"SearchTable,omitempty" type:"Repeated"`
}

func (s SearchTableResponseBodySearchTableList) String() string {
	return tea.Prettify(s)
}

func (s SearchTableResponseBodySearchTableList) GoString() string {
	return s.String()
}

func (s *SearchTableResponseBodySearchTableList) SetSearchTable(v []*SearchTableResponseBodySearchTableListSearchTable) *SearchTableResponseBodySearchTableList {
	s.SearchTable = v
	return s
}

type SearchTableResponseBodySearchTableListSearchTable struct {
	DatabaseId      *string                                                         `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	TableName       *string                                                         `json:"TableName,omitempty" xml:"TableName,omitempty"`
	DBSearchName    *string                                                         `json:"DBSearchName,omitempty" xml:"DBSearchName,omitempty"`
	Logic           *bool                                                           `json:"Logic,omitempty" xml:"Logic,omitempty"`
	EnvType         *string                                                         `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	DbName          *string                                                         `json:"DbName,omitempty" xml:"DbName,omitempty"`
	OwnerIdList     *SearchTableResponseBodySearchTableListSearchTableOwnerIdList   `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	Description     *string                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	Encoding        *string                                                         `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	DbType          *string                                                         `json:"DbType,omitempty" xml:"DbType,omitempty"`
	OwnerNameList   *SearchTableResponseBodySearchTableListSearchTableOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	TableSchemaName *string                                                         `json:"TableSchemaName,omitempty" xml:"TableSchemaName,omitempty"`
	TableGuid       *string                                                         `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	Engine          *string                                                         `json:"Engine,omitempty" xml:"Engine,omitempty"`
	TableId         *string                                                         `json:"TableId,omitempty" xml:"TableId,omitempty"`
}

func (s SearchTableResponseBodySearchTableListSearchTable) String() string {
	return tea.Prettify(s)
}

func (s SearchTableResponseBodySearchTableListSearchTable) GoString() string {
	return s.String()
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetDatabaseId(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.DatabaseId = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetTableName(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.TableName = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetDBSearchName(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.DBSearchName = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetLogic(v bool) *SearchTableResponseBodySearchTableListSearchTable {
	s.Logic = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetEnvType(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.EnvType = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetDbName(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.DbName = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetOwnerIdList(v *SearchTableResponseBodySearchTableListSearchTableOwnerIdList) *SearchTableResponseBodySearchTableListSearchTable {
	s.OwnerIdList = v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetDescription(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.Description = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetEncoding(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.Encoding = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetDbType(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.DbType = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetOwnerNameList(v *SearchTableResponseBodySearchTableListSearchTableOwnerNameList) *SearchTableResponseBodySearchTableListSearchTable {
	s.OwnerNameList = v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetTableSchemaName(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.TableSchemaName = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetTableGuid(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.TableGuid = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetEngine(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.Engine = &v
	return s
}

func (s *SearchTableResponseBodySearchTableListSearchTable) SetTableId(v string) *SearchTableResponseBodySearchTableListSearchTable {
	s.TableId = &v
	return s
}

type SearchTableResponseBodySearchTableListSearchTableOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s SearchTableResponseBodySearchTableListSearchTableOwnerIdList) String() string {
	return tea.Prettify(s)
}

func (s SearchTableResponseBodySearchTableListSearchTableOwnerIdList) GoString() string {
	return s.String()
}

func (s *SearchTableResponseBodySearchTableListSearchTableOwnerIdList) SetOwnerIds(v []*string) *SearchTableResponseBodySearchTableListSearchTableOwnerIdList {
	s.OwnerIds = v
	return s
}

type SearchTableResponseBodySearchTableListSearchTableOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s SearchTableResponseBodySearchTableListSearchTableOwnerNameList) String() string {
	return tea.Prettify(s)
}

func (s SearchTableResponseBodySearchTableListSearchTableOwnerNameList) GoString() string {
	return s.String()
}

func (s *SearchTableResponseBodySearchTableListSearchTableOwnerNameList) SetOwnerNames(v []*string) *SearchTableResponseBodySearchTableListSearchTableOwnerNameList {
	s.OwnerNames = v
	return s
}

type SearchTableResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTableResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTableResponse) GoString() string {
	return s.String()
}

func (s *SearchTableResponse) SetHeaders(v map[string]*string) *SearchTableResponse {
	s.Headers = v
	return s
}

func (s *SearchTableResponse) SetBody(v *SearchTableResponseBody) *SearchTableResponse {
	s.Body = v
	return s
}

type ListDBTaskSQLJobRequest struct {
	DBTaskGroupId *int64 `json:"DBTaskGroupId,omitempty" xml:"DBTaskGroupId,omitempty"`
	PageNumber    *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tid           *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListDBTaskSQLJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDBTaskSQLJobRequest) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobRequest) SetDBTaskGroupId(v int64) *ListDBTaskSQLJobRequest {
	s.DBTaskGroupId = &v
	return s
}

func (s *ListDBTaskSQLJobRequest) SetPageNumber(v int64) *ListDBTaskSQLJobRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDBTaskSQLJobRequest) SetPageSize(v int64) *ListDBTaskSQLJobRequest {
	s.PageSize = &v
	return s
}

func (s *ListDBTaskSQLJobRequest) SetTid(v int64) *ListDBTaskSQLJobRequest {
	s.Tid = &v
	return s
}

type ListDBTaskSQLJobResponseBody struct {
	// Id of the request
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage     *string                                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode        *string                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	DBTaskSQLJobList []*ListDBTaskSQLJobResponseBodyDBTaskSQLJobList `json:"DBTaskSQLJobList,omitempty" xml:"DBTaskSQLJobList,omitempty" type:"Repeated"`
	TotalCount       *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDBTaskSQLJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDBTaskSQLJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobResponseBody) SetRequestId(v string) *ListDBTaskSQLJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBody) SetSuccess(v bool) *ListDBTaskSQLJobResponseBody {
	s.Success = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBody) SetErrorMessage(v string) *ListDBTaskSQLJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBody) SetErrorCode(v string) *ListDBTaskSQLJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBody) SetDBTaskSQLJobList(v []*ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) *ListDBTaskSQLJobResponseBody {
	s.DBTaskSQLJobList = v
	return s
}

func (s *ListDBTaskSQLJobResponseBody) SetTotalCount(v int64) *ListDBTaskSQLJobResponseBody {
	s.TotalCount = &v
	return s
}

type ListDBTaskSQLJobResponseBodyDBTaskSQLJobList struct {
	JobId         *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobType       *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	DbSearchName  *string `json:"DbSearchName,omitempty" xml:"DbSearchName,omitempty"`
	DbId          *int64  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Logic         *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LastExecTime  *string `json:"LastExecTime,omitempty" xml:"LastExecTime,omitempty"`
	DbTaskGroupId *int64  `json:"DbTaskGroupId,omitempty" xml:"DbTaskGroupId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Transactional *bool   `json:"Transactional,omitempty" xml:"Transactional,omitempty"`
}

func (s ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) String() string {
	return tea.Prettify(s)
}

func (s ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetJobId(v int64) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.JobId = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetJobType(v string) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.JobType = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetComment(v string) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.Comment = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetDbSearchName(v string) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.DbSearchName = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetDbId(v int64) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.DbId = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetLogic(v bool) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.Logic = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetCreateTime(v string) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.CreateTime = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetLastExecTime(v string) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.LastExecTime = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetDbTaskGroupId(v int64) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.DbTaskGroupId = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetStatus(v string) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.Status = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetTransactional(v bool) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.Transactional = &v
	return s
}

type ListDBTaskSQLJobResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDBTaskSQLJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDBTaskSQLJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDBTaskSQLJobResponse) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobResponse) SetHeaders(v map[string]*string) *ListDBTaskSQLJobResponse {
	s.Headers = v
	return s
}

func (s *ListDBTaskSQLJobResponse) SetBody(v *ListDBTaskSQLJobResponseBody) *ListDBTaskSQLJobResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	Tid *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetTid(v int64) *DeleteUserRequest {
	s.Tid = &v
	return s
}

func (s *DeleteUserRequest) SetUid(v string) *DeleteUserRequest {
	s.Uid = &v
	return s
}

type DeleteUserResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserResponseBody) SetErrorCode(v string) *DeleteUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteUserResponseBody) SetErrorMessage(v string) *DeleteUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteUserResponseBody) SetSuccess(v bool) *DeleteUserResponseBody {
	s.Success = &v
	return s
}

type DeleteUserResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type GetDataCronClearTaskDetailListRequest struct {
	OrderId    *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tid        *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataCronClearTaskDetailListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataCronClearTaskDetailListRequest) GoString() string {
	return s.String()
}

func (s *GetDataCronClearTaskDetailListRequest) SetOrderId(v int64) *GetDataCronClearTaskDetailListRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataCronClearTaskDetailListRequest) SetPageNumber(v int64) *GetDataCronClearTaskDetailListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetDataCronClearTaskDetailListRequest) SetPageSize(v int64) *GetDataCronClearTaskDetailListRequest {
	s.PageSize = &v
	return s
}

func (s *GetDataCronClearTaskDetailListRequest) SetTid(v int64) *GetDataCronClearTaskDetailListRequest {
	s.Tid = &v
	return s
}

type GetDataCronClearTaskDetailListResponseBody struct {
	// Id of the request
	RequestId                   *string                                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                     *bool                                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage                *string                                                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode                   *string                                                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	DataCronClearTaskDetailList []*GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList `json:"DataCronClearTaskDetailList,omitempty" xml:"DataCronClearTaskDetailList,omitempty" type:"Repeated"`
	TotalCount                  *int64                                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetDataCronClearTaskDetailListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataCronClearTaskDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCronClearTaskDetailListResponseBody) SetRequestId(v string) *GetDataCronClearTaskDetailListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBody) SetSuccess(v bool) *GetDataCronClearTaskDetailListResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBody) SetErrorMessage(v string) *GetDataCronClearTaskDetailListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBody) SetErrorCode(v string) *GetDataCronClearTaskDetailListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBody) SetDataCronClearTaskDetailList(v []*GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) *GetDataCronClearTaskDetailListResponseBody {
	s.DataCronClearTaskDetailList = v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBody) SetTotalCount(v int64) *GetDataCronClearTaskDetailListResponseBody {
	s.TotalCount = &v
	return s
}

type GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList struct {
	DBTaskGroupId    *int64  `json:"DBTaskGroupId,omitempty" xml:"DBTaskGroupId,omitempty"`
	JobStatus        *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	ActualAffectRows *int64  `json:"ActualAffectRows,omitempty" xml:"ActualAffectRows,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) String() string {
	return tea.Prettify(s)
}

func (s GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) GoString() string {
	return s.String()
}

func (s *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) SetDBTaskGroupId(v int64) *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList {
	s.DBTaskGroupId = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) SetJobStatus(v string) *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList {
	s.JobStatus = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) SetActualAffectRows(v int64) *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList {
	s.ActualAffectRows = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) SetCreateTime(v string) *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList {
	s.CreateTime = &v
	return s
}

type GetDataCronClearTaskDetailListResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDataCronClearTaskDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDataCronClearTaskDetailListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataCronClearTaskDetailListResponse) GoString() string {
	return s.String()
}

func (s *GetDataCronClearTaskDetailListResponse) SetHeaders(v map[string]*string) *GetDataCronClearTaskDetailListResponse {
	s.Headers = v
	return s
}

func (s *GetDataCronClearTaskDetailListResponse) SetBody(v *GetDataCronClearTaskDetailListResponseBody) *GetDataCronClearTaskDetailListResponse {
	s.Body = v
	return s
}

type GetStructSyncJobAnalyzeResultRequest struct {
	OrderId     *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	CompareType *string `json:"CompareType,omitempty" xml:"CompareType,omitempty"`
	PageNumber  *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tid         *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetStructSyncJobAnalyzeResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncJobAnalyzeResultRequest) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobAnalyzeResultRequest) SetOrderId(v int64) *GetStructSyncJobAnalyzeResultRequest {
	s.OrderId = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultRequest) SetCompareType(v string) *GetStructSyncJobAnalyzeResultRequest {
	s.CompareType = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultRequest) SetPageNumber(v int64) *GetStructSyncJobAnalyzeResultRequest {
	s.PageNumber = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultRequest) SetPageSize(v int64) *GetStructSyncJobAnalyzeResultRequest {
	s.PageSize = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultRequest) SetTid(v int64) *GetStructSyncJobAnalyzeResultRequest {
	s.Tid = &v
	return s
}

type GetStructSyncJobAnalyzeResultResponseBody struct {
	// Id of the request
	RequestId                  *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                    *bool                                                                `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage               *string                                                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode                  *string                                                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	StructSyncJobAnalyzeResult *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult `json:"StructSyncJobAnalyzeResult,omitempty" xml:"StructSyncJobAnalyzeResult,omitempty" type:"Struct"`
}

func (s GetStructSyncJobAnalyzeResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncJobAnalyzeResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) SetRequestId(v string) *GetStructSyncJobAnalyzeResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) SetSuccess(v bool) *GetStructSyncJobAnalyzeResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) SetErrorMessage(v string) *GetStructSyncJobAnalyzeResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) SetErrorCode(v string) *GetStructSyncJobAnalyzeResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBody) SetStructSyncJobAnalyzeResult(v *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult) *GetStructSyncJobAnalyzeResultResponseBody {
	s.StructSyncJobAnalyzeResult = v
	return s
}

type GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult struct {
	SummaryList []*GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList `json:"SummaryList,omitempty" xml:"SummaryList,omitempty" type:"Repeated"`
	ResultList  []*GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList  `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
}

func (s GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult) SetSummaryList(v []*GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList) *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult {
	s.SummaryList = v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult) SetResultList(v []*GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResult {
	s.ResultList = v
	return s
}

type GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList struct {
	CompareType *string `json:"CompareType,omitempty" xml:"CompareType,omitempty"`
	Count       *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList) SetCompareType(v string) *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList {
	s.CompareType = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList) SetCount(v int64) *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultSummaryList {
	s.Count = &v
	return s
}

type GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList struct {
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
	Script          *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) SetSourceTableName(v string) *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList {
	s.SourceTableName = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) SetTargetTableName(v string) *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList {
	s.TargetTableName = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList) SetScript(v string) *GetStructSyncJobAnalyzeResultResponseBodyStructSyncJobAnalyzeResultResultList {
	s.Script = &v
	return s
}

type GetStructSyncJobAnalyzeResultResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStructSyncJobAnalyzeResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStructSyncJobAnalyzeResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncJobAnalyzeResultResponse) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobAnalyzeResultResponse) SetHeaders(v map[string]*string) *GetStructSyncJobAnalyzeResultResponse {
	s.Headers = v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponse) SetBody(v *GetStructSyncJobAnalyzeResultResponseBody) *GetStructSyncJobAnalyzeResultResponse {
	s.Body = v
	return s
}

type ApproveOrderRequest struct {
	Tid                *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	WorkflowInstanceId *int64  `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
	ApprovalType       *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
}

func (s ApproveOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveOrderRequest) GoString() string {
	return s.String()
}

func (s *ApproveOrderRequest) SetTid(v int64) *ApproveOrderRequest {
	s.Tid = &v
	return s
}

func (s *ApproveOrderRequest) SetWorkflowInstanceId(v int64) *ApproveOrderRequest {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ApproveOrderRequest) SetApprovalType(v string) *ApproveOrderRequest {
	s.ApprovalType = &v
	return s
}

func (s *ApproveOrderRequest) SetComment(v string) *ApproveOrderRequest {
	s.Comment = &v
	return s
}

type ApproveOrderResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApproveOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveOrderResponseBody) SetRequestId(v string) *ApproveOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApproveOrderResponseBody) SetErrorCode(v string) *ApproveOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ApproveOrderResponseBody) SetErrorMessage(v string) *ApproveOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ApproveOrderResponseBody) SetSuccess(v bool) *ApproveOrderResponseBody {
	s.Success = &v
	return s
}

type ApproveOrderResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApproveOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApproveOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveOrderResponse) GoString() string {
	return s.String()
}

func (s *ApproveOrderResponse) SetHeaders(v map[string]*string) *ApproveOrderResponse {
	s.Headers = v
	return s
}

func (s *ApproveOrderResponse) SetBody(v *ApproveOrderResponseBody) *ApproveOrderResponse {
	s.Body = v
	return s
}

type GetDataCorrectTaskDetailRequest struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Tid     *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataCorrectTaskDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDataCorrectTaskDetailRequest) SetOrderId(v int64) *GetDataCorrectTaskDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataCorrectTaskDetailRequest) SetTid(v int64) *GetDataCorrectTaskDetailRequest {
	s.Tid = &v
	return s
}

type GetDataCorrectTaskDetailResponseBody struct {
	// Id of the request
	RequestId             *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success               *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage          *string                                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode             *string                                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	DataCorrectTaskDetail *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail `json:"DataCorrectTaskDetail,omitempty" xml:"DataCorrectTaskDetail,omitempty" type:"Struct"`
}

func (s GetDataCorrectTaskDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCorrectTaskDetailResponseBody) SetRequestId(v string) *GetDataCorrectTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBody) SetSuccess(v bool) *GetDataCorrectTaskDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBody) SetErrorMessage(v string) *GetDataCorrectTaskDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBody) SetErrorCode(v string) *GetDataCorrectTaskDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBody) SetDataCorrectTaskDetail(v *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) *GetDataCorrectTaskDetailResponseBody {
	s.DataCorrectTaskDetail = v
	return s
}

type GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail struct {
	DBTaskGroupId    *int64  `json:"DBTaskGroupId,omitempty" xml:"DBTaskGroupId,omitempty"`
	JobStatus        *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	ActualAffectRows *int64  `json:"ActualAffectRows,omitempty" xml:"ActualAffectRows,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) GoString() string {
	return s.String()
}

func (s *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) SetDBTaskGroupId(v int64) *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail {
	s.DBTaskGroupId = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) SetJobStatus(v string) *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail {
	s.JobStatus = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) SetActualAffectRows(v int64) *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail {
	s.ActualAffectRows = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) SetCreateTime(v string) *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail {
	s.CreateTime = &v
	return s
}

type GetDataCorrectTaskDetailResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDataCorrectTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDataCorrectTaskDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDataCorrectTaskDetailResponse) SetHeaders(v map[string]*string) *GetDataCorrectTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDataCorrectTaskDetailResponse) SetBody(v *GetDataCorrectTaskDetailResponseBody) *GetDataCorrectTaskDetailResponse {
	s.Body = v
	return s
}

type CreateUploadFileJobRequest struct {
	FileSource *string `json:"FileSource,omitempty" xml:"FileSource,omitempty"`
	FileName   *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	UploadURL  *string `json:"UploadURL,omitempty" xml:"UploadURL,omitempty"`
	Tid        *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateUploadFileJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadFileJobRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadFileJobRequest) SetFileSource(v string) *CreateUploadFileJobRequest {
	s.FileSource = &v
	return s
}

func (s *CreateUploadFileJobRequest) SetFileName(v string) *CreateUploadFileJobRequest {
	s.FileName = &v
	return s
}

func (s *CreateUploadFileJobRequest) SetUploadURL(v string) *CreateUploadFileJobRequest {
	s.UploadURL = &v
	return s
}

func (s *CreateUploadFileJobRequest) SetTid(v int64) *CreateUploadFileJobRequest {
	s.Tid = &v
	return s
}

type CreateUploadFileJobResponseBody struct {
	// Id of the request
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	JobKey       *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
}

func (s CreateUploadFileJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadFileJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadFileJobResponseBody) SetRequestId(v string) *CreateUploadFileJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadFileJobResponseBody) SetSuccess(v bool) *CreateUploadFileJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUploadFileJobResponseBody) SetErrorMessage(v string) *CreateUploadFileJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateUploadFileJobResponseBody) SetErrorCode(v string) *CreateUploadFileJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateUploadFileJobResponseBody) SetJobKey(v string) *CreateUploadFileJobResponseBody {
	s.JobKey = &v
	return s
}

type CreateUploadFileJobResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUploadFileJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUploadFileJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadFileJobResponse) GoString() string {
	return s.String()
}

func (s *CreateUploadFileJobResponse) SetHeaders(v map[string]*string) *CreateUploadFileJobResponse {
	s.Headers = v
	return s
}

func (s *CreateUploadFileJobResponse) SetBody(v *CreateUploadFileJobResponseBody) *CreateUploadFileJobResponse {
	s.Body = v
	return s
}

type ListLogicDatabasesRequest struct {
	Tid        *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListLogicDatabasesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLogicDatabasesRequest) GoString() string {
	return s.String()
}

func (s *ListLogicDatabasesRequest) SetTid(v int64) *ListLogicDatabasesRequest {
	s.Tid = &v
	return s
}

func (s *ListLogicDatabasesRequest) SetPageNumber(v int32) *ListLogicDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLogicDatabasesRequest) SetPageSize(v int32) *ListLogicDatabasesRequest {
	s.PageSize = &v
	return s
}

type ListLogicDatabasesResponseBody struct {
	TotalCount        *int64                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LogicDatabaseList *ListLogicDatabasesResponseBodyLogicDatabaseList `json:"LogicDatabaseList,omitempty" xml:"LogicDatabaseList,omitempty" type:"Struct"`
	RequestId         *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode         *string                                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage      *string                                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success           *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLogicDatabasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLogicDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogicDatabasesResponseBody) SetTotalCount(v int64) *ListLogicDatabasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLogicDatabasesResponseBody) SetLogicDatabaseList(v *ListLogicDatabasesResponseBodyLogicDatabaseList) *ListLogicDatabasesResponseBody {
	s.LogicDatabaseList = v
	return s
}

func (s *ListLogicDatabasesResponseBody) SetRequestId(v string) *ListLogicDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogicDatabasesResponseBody) SetErrorCode(v string) *ListLogicDatabasesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListLogicDatabasesResponseBody) SetErrorMessage(v string) *ListLogicDatabasesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListLogicDatabasesResponseBody) SetSuccess(v bool) *ListLogicDatabasesResponseBody {
	s.Success = &v
	return s
}

type ListLogicDatabasesResponseBodyLogicDatabaseList struct {
	LogicDatabase []*ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase `json:"LogicDatabase,omitempty" xml:"LogicDatabase,omitempty" type:"Repeated"`
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseList) String() string {
	return tea.Prettify(s)
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseList) GoString() string {
	return s.String()
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseList) SetLogicDatabase(v []*ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) *ListLogicDatabasesResponseBodyLogicDatabaseList {
	s.LogicDatabase = v
	return s
}

type ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase struct {
	DatabaseId    *string                                                                    `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	OwnerIdList   *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList   `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	DbType        *string                                                                    `json:"DbType,omitempty" xml:"DbType,omitempty"`
	OwnerNameList *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	Logic         *bool                                                                      `json:"Logic,omitempty" xml:"Logic,omitempty"`
	SchemaName    *string                                                                    `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	SearchName    *string                                                                    `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	EnvType       *string                                                                    `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) String() string {
	return tea.Prettify(s)
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) GoString() string {
	return s.String()
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetDatabaseId(v string) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.DatabaseId = &v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetOwnerIdList(v *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.OwnerIdList = v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetDbType(v string) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.DbType = &v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetOwnerNameList(v *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.OwnerNameList = v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetLogic(v bool) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.Logic = &v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetSchemaName(v string) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.SchemaName = &v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetSearchName(v string) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.SearchName = &v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetEnvType(v string) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.EnvType = &v
	return s
}

type ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList) String() string {
	return tea.Prettify(s)
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList) GoString() string {
	return s.String()
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList) SetOwnerIds(v []*string) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList {
	s.OwnerIds = v
	return s
}

type ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList) String() string {
	return tea.Prettify(s)
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList) GoString() string {
	return s.String()
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList) SetOwnerNames(v []*string) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList {
	s.OwnerNames = v
	return s
}

type ListLogicDatabasesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLogicDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLogicDatabasesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLogicDatabasesResponse) GoString() string {
	return s.String()
}

func (s *ListLogicDatabasesResponse) SetHeaders(v map[string]*string) *ListLogicDatabasesResponse {
	s.Headers = v
	return s
}

func (s *ListLogicDatabasesResponse) SetBody(v *ListLogicDatabasesResponseBody) *ListLogicDatabasesResponse {
	s.Body = v
	return s
}

type CreateDataImportOrderRequest struct {
	Comment         *string                            `json:"Comment,omitempty" xml:"Comment,omitempty"`
	RelatedUserList []*int64                           `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	Param           *CreateDataImportOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	AttachmentKey   *string                            `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	Tid             *int64                             `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataImportOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataImportOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateDataImportOrderRequest) SetComment(v string) *CreateDataImportOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataImportOrderRequest) SetRelatedUserList(v []*int64) *CreateDataImportOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateDataImportOrderRequest) SetParam(v *CreateDataImportOrderRequestParam) *CreateDataImportOrderRequest {
	s.Param = v
	return s
}

func (s *CreateDataImportOrderRequest) SetAttachmentKey(v string) *CreateDataImportOrderRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDataImportOrderRequest) SetTid(v int64) *CreateDataImportOrderRequest {
	s.Tid = &v
	return s
}

type CreateDataImportOrderRequestParam struct {
	Classify               *string                                        `json:"Classify,omitempty" xml:"Classify,omitempty"`
	AttachmentName         *string                                        `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	DbItemList             []*CreateDataImportOrderRequestParamDbItemList `json:"DbItemList,omitempty" xml:"DbItemList,omitempty" type:"Repeated"`
	FileType               *string                                        `json:"FileType,omitempty" xml:"FileType,omitempty"`
	FileEncoding           *string                                        `json:"FileEncoding,omitempty" xml:"FileEncoding,omitempty"`
	TableName              *string                                        `json:"TableName,omitempty" xml:"TableName,omitempty"`
	InsertType             *string                                        `json:"InsertType,omitempty" xml:"InsertType,omitempty"`
	CsvFirstRowIsColumnDef *bool                                          `json:"CsvFirstRowIsColumnDef,omitempty" xml:"CsvFirstRowIsColumnDef,omitempty"`
	IgnoreError            *bool                                          `json:"IgnoreError,omitempty" xml:"IgnoreError,omitempty"`
	ImportMode             *string                                        `json:"ImportMode,omitempty" xml:"ImportMode,omitempty"`
	RollbackSQL            *string                                        `json:"RollbackSQL,omitempty" xml:"RollbackSQL,omitempty"`
	RollbackAttachmentName *string                                        `json:"RollbackAttachmentName,omitempty" xml:"RollbackAttachmentName,omitempty"`
	RollbackSqlType        *string                                        `json:"RollbackSqlType,omitempty" xml:"RollbackSqlType,omitempty"`
}

func (s CreateDataImportOrderRequestParam) String() string {
	return tea.Prettify(s)
}

func (s CreateDataImportOrderRequestParam) GoString() string {
	return s.String()
}

func (s *CreateDataImportOrderRequestParam) SetClassify(v string) *CreateDataImportOrderRequestParam {
	s.Classify = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetAttachmentName(v string) *CreateDataImportOrderRequestParam {
	s.AttachmentName = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetDbItemList(v []*CreateDataImportOrderRequestParamDbItemList) *CreateDataImportOrderRequestParam {
	s.DbItemList = v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetFileType(v string) *CreateDataImportOrderRequestParam {
	s.FileType = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetFileEncoding(v string) *CreateDataImportOrderRequestParam {
	s.FileEncoding = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetTableName(v string) *CreateDataImportOrderRequestParam {
	s.TableName = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetInsertType(v string) *CreateDataImportOrderRequestParam {
	s.InsertType = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetCsvFirstRowIsColumnDef(v bool) *CreateDataImportOrderRequestParam {
	s.CsvFirstRowIsColumnDef = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetIgnoreError(v bool) *CreateDataImportOrderRequestParam {
	s.IgnoreError = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetImportMode(v string) *CreateDataImportOrderRequestParam {
	s.ImportMode = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetRollbackSQL(v string) *CreateDataImportOrderRequestParam {
	s.RollbackSQL = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetRollbackAttachmentName(v string) *CreateDataImportOrderRequestParam {
	s.RollbackAttachmentName = &v
	return s
}

func (s *CreateDataImportOrderRequestParam) SetRollbackSqlType(v string) *CreateDataImportOrderRequestParam {
	s.RollbackSqlType = &v
	return s
}

type CreateDataImportOrderRequestParamDbItemList struct {
	DbId  *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Logic *bool  `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s CreateDataImportOrderRequestParamDbItemList) String() string {
	return tea.Prettify(s)
}

func (s CreateDataImportOrderRequestParamDbItemList) GoString() string {
	return s.String()
}

func (s *CreateDataImportOrderRequestParamDbItemList) SetDbId(v int64) *CreateDataImportOrderRequestParamDbItemList {
	s.DbId = &v
	return s
}

func (s *CreateDataImportOrderRequestParamDbItemList) SetLogic(v bool) *CreateDataImportOrderRequestParamDbItemList {
	s.Logic = &v
	return s
}

type CreateDataImportOrderShrinkRequest struct {
	Comment               *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	RelatedUserListShrink *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	ParamShrink           *string `json:"Param,omitempty" xml:"Param,omitempty"`
	AttachmentKey         *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	Tid                   *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataImportOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataImportOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataImportOrderShrinkRequest) SetComment(v string) *CreateDataImportOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataImportOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateDataImportOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateDataImportOrderShrinkRequest) SetParamShrink(v string) *CreateDataImportOrderShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *CreateDataImportOrderShrinkRequest) SetAttachmentKey(v string) *CreateDataImportOrderShrinkRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDataImportOrderShrinkRequest) SetTid(v int64) *CreateDataImportOrderShrinkRequest {
	s.Tid = &v
	return s
}

type CreateDataImportOrderResponseBody struct {
	// Id of the request
	RequestId         *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success           *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage      *string  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode         *string  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	CreateOrderResult []*int64 `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Repeated"`
}

func (s CreateDataImportOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataImportOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataImportOrderResponseBody) SetRequestId(v string) *CreateDataImportOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataImportOrderResponseBody) SetSuccess(v bool) *CreateDataImportOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataImportOrderResponseBody) SetErrorMessage(v string) *CreateDataImportOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataImportOrderResponseBody) SetErrorCode(v string) *CreateDataImportOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataImportOrderResponseBody) SetCreateOrderResult(v []*int64) *CreateDataImportOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

type CreateDataImportOrderResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDataImportOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataImportOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataImportOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateDataImportOrderResponse) SetHeaders(v map[string]*string) *CreateDataImportOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateDataImportOrderResponse) SetBody(v *CreateDataImportOrderResponseBody) *CreateDataImportOrderResponse {
	s.Body = v
	return s
}

type CloseOrderRequest struct {
	Tid         *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	OrderId     *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	CloseReason *string `json:"CloseReason,omitempty" xml:"CloseReason,omitempty"`
}

func (s CloseOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseOrderRequest) GoString() string {
	return s.String()
}

func (s *CloseOrderRequest) SetTid(v int64) *CloseOrderRequest {
	s.Tid = &v
	return s
}

func (s *CloseOrderRequest) SetOrderId(v int64) *CloseOrderRequest {
	s.OrderId = &v
	return s
}

func (s *CloseOrderRequest) SetCloseReason(v string) *CloseOrderRequest {
	s.CloseReason = &v
	return s
}

type CloseOrderResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CloseOrderResponseBody) SetRequestId(v string) *CloseOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseOrderResponseBody) SetErrorCode(v string) *CloseOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CloseOrderResponseBody) SetErrorMessage(v string) *CloseOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CloseOrderResponseBody) SetSuccess(v bool) *CloseOrderResponseBody {
	s.Success = &v
	return s
}

type CloseOrderResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloseOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseOrderResponse) GoString() string {
	return s.String()
}

func (s *CloseOrderResponse) SetHeaders(v map[string]*string) *CloseOrderResponse {
	s.Headers = v
	return s
}

func (s *CloseOrderResponse) SetBody(v *CloseOrderResponseBody) *CloseOrderResponse {
	s.Body = v
	return s
}

type SyncInstanceMetaRequest struct {
	Tid         *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IgnoreTable *bool   `json:"IgnoreTable,omitempty" xml:"IgnoreTable,omitempty"`
}

func (s SyncInstanceMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncInstanceMetaRequest) GoString() string {
	return s.String()
}

func (s *SyncInstanceMetaRequest) SetTid(v int64) *SyncInstanceMetaRequest {
	s.Tid = &v
	return s
}

func (s *SyncInstanceMetaRequest) SetInstanceId(v string) *SyncInstanceMetaRequest {
	s.InstanceId = &v
	return s
}

func (s *SyncInstanceMetaRequest) SetIgnoreTable(v bool) *SyncInstanceMetaRequest {
	s.IgnoreTable = &v
	return s
}

type SyncInstanceMetaResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncInstanceMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncInstanceMetaResponseBody) GoString() string {
	return s.String()
}

func (s *SyncInstanceMetaResponseBody) SetRequestId(v string) *SyncInstanceMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncInstanceMetaResponseBody) SetErrorCode(v string) *SyncInstanceMetaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SyncInstanceMetaResponseBody) SetErrorMessage(v string) *SyncInstanceMetaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SyncInstanceMetaResponseBody) SetSuccess(v bool) *SyncInstanceMetaResponseBody {
	s.Success = &v
	return s
}

type SyncInstanceMetaResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncInstanceMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncInstanceMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncInstanceMetaResponse) GoString() string {
	return s.String()
}

func (s *SyncInstanceMetaResponse) SetHeaders(v map[string]*string) *SyncInstanceMetaResponse {
	s.Headers = v
	return s
}

func (s *SyncInstanceMetaResponse) SetBody(v *SyncInstanceMetaResponseBody) *SyncInstanceMetaResponse {
	s.Body = v
	return s
}

type ListOrdersRequest struct {
	Tid             *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	PluginType      *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	OrderResultType *string `json:"OrderResultType,omitempty" xml:"OrderResultType,omitempty"`
	SearchDateType  *string `json:"SearchDateType,omitempty" xml:"SearchDateType,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SearchContent   *string `json:"SearchContent,omitempty" xml:"SearchContent,omitempty"`
	OrderStatus     *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListOrdersRequest) SetTid(v int64) *ListOrdersRequest {
	s.Tid = &v
	return s
}

func (s *ListOrdersRequest) SetPluginType(v string) *ListOrdersRequest {
	s.PluginType = &v
	return s
}

func (s *ListOrdersRequest) SetOrderResultType(v string) *ListOrdersRequest {
	s.OrderResultType = &v
	return s
}

func (s *ListOrdersRequest) SetSearchDateType(v string) *ListOrdersRequest {
	s.SearchDateType = &v
	return s
}

func (s *ListOrdersRequest) SetStartTime(v string) *ListOrdersRequest {
	s.StartTime = &v
	return s
}

func (s *ListOrdersRequest) SetEndTime(v string) *ListOrdersRequest {
	s.EndTime = &v
	return s
}

func (s *ListOrdersRequest) SetSearchContent(v string) *ListOrdersRequest {
	s.SearchContent = &v
	return s
}

func (s *ListOrdersRequest) SetOrderStatus(v string) *ListOrdersRequest {
	s.OrderStatus = &v
	return s
}

func (s *ListOrdersRequest) SetPageSize(v int32) *ListOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *ListOrdersRequest) SetPageNumber(v int32) *ListOrdersRequest {
	s.PageNumber = &v
	return s
}

type ListOrdersResponseBody struct {
	Orders       *ListOrdersResponseBodyOrders `json:"Orders,omitempty" xml:"Orders,omitempty" type:"Struct"`
	TotalCount   *int64                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId    *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrdersResponseBody) SetOrders(v *ListOrdersResponseBodyOrders) *ListOrdersResponseBody {
	s.Orders = v
	return s
}

func (s *ListOrdersResponseBody) SetTotalCount(v int64) *ListOrdersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrdersResponseBody) SetRequestId(v string) *ListOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrdersResponseBody) SetErrorCode(v string) *ListOrdersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListOrdersResponseBody) SetErrorMessage(v string) *ListOrdersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListOrdersResponseBody) SetSuccess(v bool) *ListOrdersResponseBody {
	s.Success = &v
	return s
}

type ListOrdersResponseBodyOrders struct {
	Order []*ListOrdersResponseBodyOrdersOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Repeated"`
}

func (s ListOrdersResponseBodyOrders) String() string {
	return tea.Prettify(s)
}

func (s ListOrdersResponseBodyOrders) GoString() string {
	return s.String()
}

func (s *ListOrdersResponseBodyOrders) SetOrder(v []*ListOrdersResponseBodyOrdersOrder) *ListOrdersResponseBodyOrders {
	s.Order = v
	return s
}

type ListOrdersResponseBodyOrdersOrder struct {
	Comment        *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	LastModifyTime *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	StatusCode     *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Committer      *string `json:"Committer,omitempty" xml:"Committer,omitempty"`
	CommitterId    *int64  `json:"CommitterId,omitempty" xml:"CommitterId,omitempty"`
	StatusDesc     *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	PluginType     *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	OrderId        *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ListOrdersResponseBodyOrdersOrder) String() string {
	return tea.Prettify(s)
}

func (s ListOrdersResponseBodyOrdersOrder) GoString() string {
	return s.String()
}

func (s *ListOrdersResponseBodyOrdersOrder) SetComment(v string) *ListOrdersResponseBodyOrdersOrder {
	s.Comment = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetLastModifyTime(v string) *ListOrdersResponseBodyOrdersOrder {
	s.LastModifyTime = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetStatusCode(v string) *ListOrdersResponseBodyOrdersOrder {
	s.StatusCode = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetCreateTime(v string) *ListOrdersResponseBodyOrdersOrder {
	s.CreateTime = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetCommitter(v string) *ListOrdersResponseBodyOrdersOrder {
	s.Committer = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetCommitterId(v int64) *ListOrdersResponseBodyOrdersOrder {
	s.CommitterId = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetStatusDesc(v string) *ListOrdersResponseBodyOrdersOrder {
	s.StatusDesc = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetPluginType(v string) *ListOrdersResponseBodyOrdersOrder {
	s.PluginType = &v
	return s
}

func (s *ListOrdersResponseBodyOrdersOrder) SetOrderId(v int64) *ListOrdersResponseBodyOrdersOrder {
	s.OrderId = &v
	return s
}

type ListOrdersResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListOrdersResponse) SetHeaders(v map[string]*string) *ListOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListOrdersResponse) SetBody(v *ListOrdersResponseBody) *ListOrdersResponse {
	s.Body = v
	return s
}

type CreateSQLReviewOrderRequest struct {
	Comment         *string                           `json:"Comment,omitempty" xml:"Comment,omitempty"`
	RelatedUserList []*int64                          `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	Param           *CreateSQLReviewOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	Tid             *int64                            `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateSQLReviewOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSQLReviewOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateSQLReviewOrderRequest) SetComment(v string) *CreateSQLReviewOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateSQLReviewOrderRequest) SetRelatedUserList(v []*int64) *CreateSQLReviewOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateSQLReviewOrderRequest) SetParam(v *CreateSQLReviewOrderRequestParam) *CreateSQLReviewOrderRequest {
	s.Param = v
	return s
}

func (s *CreateSQLReviewOrderRequest) SetTid(v int64) *CreateSQLReviewOrderRequest {
	s.Tid = &v
	return s
}

type CreateSQLReviewOrderRequestParam struct {
	ProjectName       *string   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	DbId              *int64    `json:"DbId,omitempty" xml:"DbId,omitempty"`
	AttachmentKeyList []*string `json:"AttachmentKeyList,omitempty" xml:"AttachmentKeyList,omitempty" type:"Repeated"`
}

func (s CreateSQLReviewOrderRequestParam) String() string {
	return tea.Prettify(s)
}

func (s CreateSQLReviewOrderRequestParam) GoString() string {
	return s.String()
}

func (s *CreateSQLReviewOrderRequestParam) SetProjectName(v string) *CreateSQLReviewOrderRequestParam {
	s.ProjectName = &v
	return s
}

func (s *CreateSQLReviewOrderRequestParam) SetDbId(v int64) *CreateSQLReviewOrderRequestParam {
	s.DbId = &v
	return s
}

func (s *CreateSQLReviewOrderRequestParam) SetAttachmentKeyList(v []*string) *CreateSQLReviewOrderRequestParam {
	s.AttachmentKeyList = v
	return s
}

type CreateSQLReviewOrderShrinkRequest struct {
	Comment               *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	RelatedUserListShrink *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	ParamShrink           *string `json:"Param,omitempty" xml:"Param,omitempty"`
	Tid                   *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateSQLReviewOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSQLReviewOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSQLReviewOrderShrinkRequest) SetComment(v string) *CreateSQLReviewOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateSQLReviewOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateSQLReviewOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateSQLReviewOrderShrinkRequest) SetParamShrink(v string) *CreateSQLReviewOrderShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *CreateSQLReviewOrderShrinkRequest) SetTid(v int64) *CreateSQLReviewOrderShrinkRequest {
	s.Tid = &v
	return s
}

type CreateSQLReviewOrderResponseBody struct {
	// Id of the request
	RequestId         *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success           *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage      *string  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode         *string  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	CreateOrderResult []*int64 `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Repeated"`
}

func (s CreateSQLReviewOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSQLReviewOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSQLReviewOrderResponseBody) SetRequestId(v string) *CreateSQLReviewOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSQLReviewOrderResponseBody) SetSuccess(v bool) *CreateSQLReviewOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSQLReviewOrderResponseBody) SetErrorMessage(v string) *CreateSQLReviewOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSQLReviewOrderResponseBody) SetErrorCode(v string) *CreateSQLReviewOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSQLReviewOrderResponseBody) SetCreateOrderResult(v []*int64) *CreateSQLReviewOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

type CreateSQLReviewOrderResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSQLReviewOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSQLReviewOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSQLReviewOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateSQLReviewOrderResponse) SetHeaders(v map[string]*string) *CreateSQLReviewOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateSQLReviewOrderResponse) SetBody(v *CreateSQLReviewOrderResponseBody) *CreateSQLReviewOrderResponse {
	s.Body = v
	return s
}

type GetOrderBaseInfoRequest struct {
	Tid     *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s GetOrderBaseInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrderBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *GetOrderBaseInfoRequest) SetTid(v int64) *GetOrderBaseInfoRequest {
	s.Tid = &v
	return s
}

func (s *GetOrderBaseInfoRequest) SetOrderId(v int64) *GetOrderBaseInfoRequest {
	s.OrderId = &v
	return s
}

type GetOrderBaseInfoResponseBody struct {
	OrderBaseInfo *GetOrderBaseInfoResponseBodyOrderBaseInfo `json:"OrderBaseInfo,omitempty" xml:"OrderBaseInfo,omitempty" type:"Struct"`
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode     *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage  *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success       *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOrderBaseInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrderBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrderBaseInfoResponseBody) SetOrderBaseInfo(v *GetOrderBaseInfoResponseBodyOrderBaseInfo) *GetOrderBaseInfoResponseBody {
	s.OrderBaseInfo = v
	return s
}

func (s *GetOrderBaseInfoResponseBody) SetRequestId(v string) *GetOrderBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrderBaseInfoResponseBody) SetErrorCode(v string) *GetOrderBaseInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOrderBaseInfoResponseBody) SetErrorMessage(v string) *GetOrderBaseInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOrderBaseInfoResponseBody) SetSuccess(v bool) *GetOrderBaseInfoResponseBody {
	s.Success = &v
	return s
}

type GetOrderBaseInfoResponseBodyOrderBaseInfo struct {
	RelatedUserNickList *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList `json:"RelatedUserNickList,omitempty" xml:"RelatedUserNickList,omitempty" type:"Struct"`
	Comment             *string                                                       `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreateTime          *string                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Committer           *string                                                       `json:"Committer,omitempty" xml:"Committer,omitempty"`
	WorkflowInstanceId  *int64                                                        `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
	CommitterId         *int64                                                        `json:"CommitterId,omitempty" xml:"CommitterId,omitempty"`
	LastModifyTime      *string                                                       `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	StatusCode          *string                                                       `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	RelatedUserList     *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList     `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Struct"`
	WorkflowStatusDesc  *string                                                       `json:"WorkflowStatusDesc,omitempty" xml:"WorkflowStatusDesc,omitempty"`
	StatusDesc          *string                                                       `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	PluginType          *string                                                       `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	OrderId             *int64                                                        `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s GetOrderBaseInfoResponseBodyOrderBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s GetOrderBaseInfoResponseBodyOrderBaseInfo) GoString() string {
	return s.String()
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetRelatedUserNickList(v *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.RelatedUserNickList = v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetComment(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.Comment = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetCreateTime(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetCommitter(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.Committer = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetWorkflowInstanceId(v int64) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.WorkflowInstanceId = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetCommitterId(v int64) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.CommitterId = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetLastModifyTime(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.LastModifyTime = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetStatusCode(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.StatusCode = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetRelatedUserList(v *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.RelatedUserList = v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetWorkflowStatusDesc(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.WorkflowStatusDesc = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetStatusDesc(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.StatusDesc = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetPluginType(v string) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.PluginType = &v
	return s
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfo) SetOrderId(v int64) *GetOrderBaseInfoResponseBodyOrderBaseInfo {
	s.OrderId = &v
	return s
}

type GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList struct {
	UserNicks []*string `json:"UserNicks,omitempty" xml:"UserNicks,omitempty" type:"Repeated"`
}

func (s GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList) String() string {
	return tea.Prettify(s)
}

func (s GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList) GoString() string {
	return s.String()
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList) SetUserNicks(v []*string) *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserNickList {
	s.UserNicks = v
	return s
}

type GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList struct {
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList) String() string {
	return tea.Prettify(s)
}

func (s GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList) GoString() string {
	return s.String()
}

func (s *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList) SetUserIds(v []*string) *GetOrderBaseInfoResponseBodyOrderBaseInfoRelatedUserList {
	s.UserIds = v
	return s
}

type GetOrderBaseInfoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOrderBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrderBaseInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrderBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *GetOrderBaseInfoResponse) SetHeaders(v map[string]*string) *GetOrderBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *GetOrderBaseInfoResponse) SetBody(v *GetOrderBaseInfoResponseBody) *GetOrderBaseInfoResponse {
	s.Body = v
	return s
}

type GetSQLReviewOptimizeDetailRequest struct {
	SQLReviewQueryKey *string `json:"SQLReviewQueryKey,omitempty" xml:"SQLReviewQueryKey,omitempty"`
	Tid               *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetSQLReviewOptimizeDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSQLReviewOptimizeDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSQLReviewOptimizeDetailRequest) SetSQLReviewQueryKey(v string) *GetSQLReviewOptimizeDetailRequest {
	s.SQLReviewQueryKey = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailRequest) SetTid(v int64) *GetSQLReviewOptimizeDetailRequest {
	s.Tid = &v
	return s
}

type GetSQLReviewOptimizeDetailResponseBody struct {
	RequestId      *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode      *string                                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success        *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	OptimizeDetail *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail `json:"OptimizeDetail,omitempty" xml:"OptimizeDetail,omitempty" type:"Struct"`
}

func (s GetSQLReviewOptimizeDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSQLReviewOptimizeDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSQLReviewOptimizeDetailResponseBody) SetRequestId(v string) *GetSQLReviewOptimizeDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBody) SetErrorCode(v string) *GetSQLReviewOptimizeDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBody) SetErrorMessage(v string) *GetSQLReviewOptimizeDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBody) SetSuccess(v bool) *GetSQLReviewOptimizeDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBody) SetOptimizeDetail(v *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) *GetSQLReviewOptimizeDetailResponseBody {
	s.OptimizeDetail = v
	return s
}

type GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail struct {
	QueryKey      *string                                                            `json:"QueryKey,omitempty" xml:"QueryKey,omitempty"`
	InstanceId    *int32                                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbId          *int32                                                             `json:"DbId,omitempty" xml:"DbId,omitempty"`
	SqlType       *string                                                            `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	QualityResult *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult `json:"QualityResult,omitempty" xml:"QualityResult,omitempty" type:"Struct"`
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) String() string {
	return tea.Prettify(s)
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) GoString() string {
	return s.String()
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) SetQueryKey(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail {
	s.QueryKey = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) SetInstanceId(v int32) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail {
	s.InstanceId = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) SetDbId(v int32) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail {
	s.DbId = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) SetSqlType(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail {
	s.SqlType = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) SetQualityResult(v *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail {
	s.QualityResult = v
	return s
}

type GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult struct {
	ErrorMessage *string                                                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Results      []*GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	OccurError   *bool                                                                       `json:"OccurError,omitempty" xml:"OccurError,omitempty"`
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) String() string {
	return tea.Prettify(s)
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) GoString() string {
	return s.String()
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) SetErrorMessage(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult {
	s.ErrorMessage = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) SetResults(v []*GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult {
	s.Results = v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) SetOccurError(v bool) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult {
	s.OccurError = &v
	return s
}

type GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults struct {
	RuleName *string                                                                            `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Feedback *string                                                                            `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	Comments *string                                                                            `json:"Comments,omitempty" xml:"Comments,omitempty"`
	Messages []*string                                                                          `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	Scripts  []*GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts `json:"Scripts,omitempty" xml:"Scripts,omitempty" type:"Repeated"`
	RuleType *string                                                                            `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) String() string {
	return tea.Prettify(s)
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) GoString() string {
	return s.String()
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) SetRuleName(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults {
	s.RuleName = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) SetFeedback(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults {
	s.Feedback = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) SetComments(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults {
	s.Comments = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) SetMessages(v []*string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults {
	s.Messages = v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) SetScripts(v []*GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults {
	s.Scripts = v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) SetRuleType(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults {
	s.RuleType = &v
	return s
}

type GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts struct {
	OpType    *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) String() string {
	return tea.Prettify(s)
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) GoString() string {
	return s.String()
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) SetOpType(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts {
	s.OpType = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) SetContent(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts {
	s.Content = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) SetTableName(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts {
	s.TableName = &v
	return s
}

type GetSQLReviewOptimizeDetailResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSQLReviewOptimizeDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSQLReviewOptimizeDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSQLReviewOptimizeDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSQLReviewOptimizeDetailResponse) SetHeaders(v map[string]*string) *GetSQLReviewOptimizeDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponse) SetBody(v *GetSQLReviewOptimizeDetailResponseBody) *GetSQLReviewOptimizeDetailResponse {
	s.Body = v
	return s
}

type ListUserTenantsRequest struct {
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListUserTenantsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserTenantsRequest) GoString() string {
	return s.String()
}

func (s *ListUserTenantsRequest) SetTid(v int64) *ListUserTenantsRequest {
	s.Tid = &v
	return s
}

type ListUserTenantsResponseBody struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantList   []*ListUserTenantsResponseBodyTenantList `json:"TenantList,omitempty" xml:"TenantList,omitempty" type:"Repeated"`
	ErrorCode    *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUserTenantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserTenantsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserTenantsResponseBody) SetRequestId(v string) *ListUserTenantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserTenantsResponseBody) SetTenantList(v []*ListUserTenantsResponseBodyTenantList) *ListUserTenantsResponseBody {
	s.TenantList = v
	return s
}

func (s *ListUserTenantsResponseBody) SetErrorCode(v string) *ListUserTenantsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUserTenantsResponseBody) SetErrorMessage(v string) *ListUserTenantsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListUserTenantsResponseBody) SetSuccess(v bool) *ListUserTenantsResponseBody {
	s.Success = &v
	return s
}

type ListUserTenantsResponseBodyTenantList struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Tid        *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s ListUserTenantsResponseBodyTenantList) String() string {
	return tea.Prettify(s)
}

func (s ListUserTenantsResponseBodyTenantList) GoString() string {
	return s.String()
}

func (s *ListUserTenantsResponseBodyTenantList) SetStatus(v string) *ListUserTenantsResponseBodyTenantList {
	s.Status = &v
	return s
}

func (s *ListUserTenantsResponseBodyTenantList) SetTid(v int64) *ListUserTenantsResponseBodyTenantList {
	s.Tid = &v
	return s
}

func (s *ListUserTenantsResponseBodyTenantList) SetTenantName(v string) *ListUserTenantsResponseBodyTenantList {
	s.TenantName = &v
	return s
}

type ListUserTenantsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserTenantsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserTenantsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserTenantsResponse) GoString() string {
	return s.String()
}

func (s *ListUserTenantsResponse) SetHeaders(v map[string]*string) *ListUserTenantsResponse {
	s.Headers = v
	return s
}

func (s *ListUserTenantsResponse) SetBody(v *ListUserTenantsResponseBody) *ListUserTenantsResponse {
	s.Body = v
	return s
}

type SetOwnersRequest struct {
	Tid        *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	OwnerType  *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	OwnerIds   *string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty"`
}

func (s SetOwnersRequest) String() string {
	return tea.Prettify(s)
}

func (s SetOwnersRequest) GoString() string {
	return s.String()
}

func (s *SetOwnersRequest) SetTid(v int64) *SetOwnersRequest {
	s.Tid = &v
	return s
}

func (s *SetOwnersRequest) SetResourceId(v string) *SetOwnersRequest {
	s.ResourceId = &v
	return s
}

func (s *SetOwnersRequest) SetOwnerType(v string) *SetOwnersRequest {
	s.OwnerType = &v
	return s
}

func (s *SetOwnersRequest) SetOwnerIds(v string) *SetOwnersRequest {
	s.OwnerIds = &v
	return s
}

type SetOwnersResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetOwnersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetOwnersResponseBody) GoString() string {
	return s.String()
}

func (s *SetOwnersResponseBody) SetRequestId(v string) *SetOwnersResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetOwnersResponseBody) SetErrorCode(v string) *SetOwnersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetOwnersResponseBody) SetErrorMessage(v string) *SetOwnersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetOwnersResponseBody) SetSuccess(v bool) *SetOwnersResponseBody {
	s.Success = &v
	return s
}

type SetOwnersResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetOwnersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetOwnersResponse) String() string {
	return tea.Prettify(s)
}

func (s SetOwnersResponse) GoString() string {
	return s.String()
}

func (s *SetOwnersResponse) SetHeaders(v map[string]*string) *SetOwnersResponse {
	s.Headers = v
	return s
}

func (s *SetOwnersResponse) SetBody(v *SetOwnersResponseBody) *SetOwnersResponse {
	s.Body = v
	return s
}

type CreateDataCorrectOrderRequest struct {
	Comment         *string                             `json:"Comment,omitempty" xml:"Comment,omitempty"`
	RelatedUserList []*int64                            `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	Param           *CreateDataCorrectOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	AttachmentKey   *string                             `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	Tid             *int64                              `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataCorrectOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCorrectOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateDataCorrectOrderRequest) SetComment(v string) *CreateDataCorrectOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataCorrectOrderRequest) SetRelatedUserList(v []*int64) *CreateDataCorrectOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateDataCorrectOrderRequest) SetParam(v *CreateDataCorrectOrderRequestParam) *CreateDataCorrectOrderRequest {
	s.Param = v
	return s
}

func (s *CreateDataCorrectOrderRequest) SetAttachmentKey(v string) *CreateDataCorrectOrderRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDataCorrectOrderRequest) SetTid(v int64) *CreateDataCorrectOrderRequest {
	s.Tid = &v
	return s
}

type CreateDataCorrectOrderRequestParam struct {
	Classify               *string                                         `json:"Classify,omitempty" xml:"Classify,omitempty"`
	EstimateAffectRows     *int64                                          `json:"EstimateAffectRows,omitempty" xml:"EstimateAffectRows,omitempty"`
	ExecSQL                *string                                         `json:"ExecSQL,omitempty" xml:"ExecSQL,omitempty"`
	SqlType                *string                                         `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	AttachmentName         *string                                         `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	RollbackSQL            *string                                         `json:"RollbackSQL,omitempty" xml:"RollbackSQL,omitempty"`
	RollbackAttachmentName *string                                         `json:"RollbackAttachmentName,omitempty" xml:"RollbackAttachmentName,omitempty"`
	RollbackSqlType        *string                                         `json:"RollbackSqlType,omitempty" xml:"RollbackSqlType,omitempty"`
	DbItemList             []*CreateDataCorrectOrderRequestParamDbItemList `json:"DbItemList,omitempty" xml:"DbItemList,omitempty" type:"Repeated"`
}

func (s CreateDataCorrectOrderRequestParam) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCorrectOrderRequestParam) GoString() string {
	return s.String()
}

func (s *CreateDataCorrectOrderRequestParam) SetClassify(v string) *CreateDataCorrectOrderRequestParam {
	s.Classify = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetEstimateAffectRows(v int64) *CreateDataCorrectOrderRequestParam {
	s.EstimateAffectRows = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetExecSQL(v string) *CreateDataCorrectOrderRequestParam {
	s.ExecSQL = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetSqlType(v string) *CreateDataCorrectOrderRequestParam {
	s.SqlType = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetAttachmentName(v string) *CreateDataCorrectOrderRequestParam {
	s.AttachmentName = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetRollbackSQL(v string) *CreateDataCorrectOrderRequestParam {
	s.RollbackSQL = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetRollbackAttachmentName(v string) *CreateDataCorrectOrderRequestParam {
	s.RollbackAttachmentName = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetRollbackSqlType(v string) *CreateDataCorrectOrderRequestParam {
	s.RollbackSqlType = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParam) SetDbItemList(v []*CreateDataCorrectOrderRequestParamDbItemList) *CreateDataCorrectOrderRequestParam {
	s.DbItemList = v
	return s
}

type CreateDataCorrectOrderRequestParamDbItemList struct {
	DbId  *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Logic *bool  `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s CreateDataCorrectOrderRequestParamDbItemList) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCorrectOrderRequestParamDbItemList) GoString() string {
	return s.String()
}

func (s *CreateDataCorrectOrderRequestParamDbItemList) SetDbId(v int64) *CreateDataCorrectOrderRequestParamDbItemList {
	s.DbId = &v
	return s
}

func (s *CreateDataCorrectOrderRequestParamDbItemList) SetLogic(v bool) *CreateDataCorrectOrderRequestParamDbItemList {
	s.Logic = &v
	return s
}

type CreateDataCorrectOrderShrinkRequest struct {
	Comment               *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	RelatedUserListShrink *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	ParamShrink           *string `json:"Param,omitempty" xml:"Param,omitempty"`
	AttachmentKey         *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	Tid                   *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataCorrectOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCorrectOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataCorrectOrderShrinkRequest) SetComment(v string) *CreateDataCorrectOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataCorrectOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateDataCorrectOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateDataCorrectOrderShrinkRequest) SetParamShrink(v string) *CreateDataCorrectOrderShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *CreateDataCorrectOrderShrinkRequest) SetAttachmentKey(v string) *CreateDataCorrectOrderShrinkRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDataCorrectOrderShrinkRequest) SetTid(v int64) *CreateDataCorrectOrderShrinkRequest {
	s.Tid = &v
	return s
}

type CreateDataCorrectOrderResponseBody struct {
	// Id of the request
	RequestId         *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success           *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage      *string  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode         *string  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	CreateOrderResult []*int64 `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Repeated"`
}

func (s CreateDataCorrectOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCorrectOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataCorrectOrderResponseBody) SetRequestId(v string) *CreateDataCorrectOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataCorrectOrderResponseBody) SetSuccess(v bool) *CreateDataCorrectOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataCorrectOrderResponseBody) SetErrorMessage(v string) *CreateDataCorrectOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataCorrectOrderResponseBody) SetErrorCode(v string) *CreateDataCorrectOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataCorrectOrderResponseBody) SetCreateOrderResult(v []*int64) *CreateDataCorrectOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

type CreateDataCorrectOrderResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDataCorrectOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataCorrectOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCorrectOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateDataCorrectOrderResponse) SetHeaders(v map[string]*string) *CreateDataCorrectOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateDataCorrectOrderResponse) SetBody(v *CreateDataCorrectOrderResponseBody) *CreateDataCorrectOrderResponse {
	s.Body = v
	return s
}

type GetLogicDatabaseRequest struct {
	Tid  *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
}

func (s GetLogicDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogicDatabaseRequest) GoString() string {
	return s.String()
}

func (s *GetLogicDatabaseRequest) SetTid(v int64) *GetLogicDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *GetLogicDatabaseRequest) SetDbId(v string) *GetLogicDatabaseRequest {
	s.DbId = &v
	return s
}

type GetLogicDatabaseResponseBody struct {
	LogicDatabase *GetLogicDatabaseResponseBodyLogicDatabase `json:"LogicDatabase,omitempty" xml:"LogicDatabase,omitempty" type:"Struct"`
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode     *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage  *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success       *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLogicDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLogicDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogicDatabaseResponseBody) SetLogicDatabase(v *GetLogicDatabaseResponseBodyLogicDatabase) *GetLogicDatabaseResponseBody {
	s.LogicDatabase = v
	return s
}

func (s *GetLogicDatabaseResponseBody) SetRequestId(v string) *GetLogicDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogicDatabaseResponseBody) SetErrorCode(v string) *GetLogicDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetLogicDatabaseResponseBody) SetErrorMessage(v string) *GetLogicDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetLogicDatabaseResponseBody) SetSuccess(v bool) *GetLogicDatabaseResponseBody {
	s.Success = &v
	return s
}

type GetLogicDatabaseResponseBodyLogicDatabase struct {
	DatabaseId    *string                                                 `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	OwnerIdList   *GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList   `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	DbType        *string                                                 `json:"DbType,omitempty" xml:"DbType,omitempty"`
	OwnerNameList *GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	Logic         *bool                                                   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	SchemaName    *string                                                 `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	SearchName    *string                                                 `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	EnvType       *string                                                 `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
}

func (s GetLogicDatabaseResponseBodyLogicDatabase) String() string {
	return tea.Prettify(s)
}

func (s GetLogicDatabaseResponseBodyLogicDatabase) GoString() string {
	return s.String()
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetDatabaseId(v string) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.DatabaseId = &v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetOwnerIdList(v *GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.OwnerIdList = v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetDbType(v string) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.DbType = &v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetOwnerNameList(v *GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.OwnerNameList = v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetLogic(v bool) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.Logic = &v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetSchemaName(v string) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.SchemaName = &v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetSearchName(v string) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.SearchName = &v
	return s
}

func (s *GetLogicDatabaseResponseBodyLogicDatabase) SetEnvType(v string) *GetLogicDatabaseResponseBodyLogicDatabase {
	s.EnvType = &v
	return s
}

type GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList) String() string {
	return tea.Prettify(s)
}

func (s GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList) GoString() string {
	return s.String()
}

func (s *GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList) SetOwnerIds(v []*string) *GetLogicDatabaseResponseBodyLogicDatabaseOwnerIdList {
	s.OwnerIds = v
	return s
}

type GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList) String() string {
	return tea.Prettify(s)
}

func (s GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList) GoString() string {
	return s.String()
}

func (s *GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList) SetOwnerNames(v []*string) *GetLogicDatabaseResponseBodyLogicDatabaseOwnerNameList {
	s.OwnerNames = v
	return s
}

type GetLogicDatabaseResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLogicDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLogicDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogicDatabaseResponse) GoString() string {
	return s.String()
}

func (s *GetLogicDatabaseResponse) SetHeaders(v map[string]*string) *GetLogicDatabaseResponse {
	s.Headers = v
	return s
}

func (s *GetLogicDatabaseResponse) SetBody(v *GetLogicDatabaseResponseBody) *GetLogicDatabaseResponse {
	s.Body = v
	return s
}

type GetDataCorrectBackupFilesRequest struct {
	Tid          *int64                 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	OrderId      *int64                 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ActionDetail map[string]interface{} `json:"ActionDetail,omitempty" xml:"ActionDetail,omitempty"`
}

func (s GetDataCorrectBackupFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectBackupFilesRequest) GoString() string {
	return s.String()
}

func (s *GetDataCorrectBackupFilesRequest) SetTid(v int64) *GetDataCorrectBackupFilesRequest {
	s.Tid = &v
	return s
}

func (s *GetDataCorrectBackupFilesRequest) SetOrderId(v int64) *GetDataCorrectBackupFilesRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataCorrectBackupFilesRequest) SetActionDetail(v map[string]interface{}) *GetDataCorrectBackupFilesRequest {
	s.ActionDetail = v
	return s
}

type GetDataCorrectBackupFilesShrinkRequest struct {
	Tid                *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	OrderId            *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ActionDetailShrink *string `json:"ActionDetail,omitempty" xml:"ActionDetail,omitempty"`
}

func (s GetDataCorrectBackupFilesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectBackupFilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDataCorrectBackupFilesShrinkRequest) SetTid(v int64) *GetDataCorrectBackupFilesShrinkRequest {
	s.Tid = &v
	return s
}

func (s *GetDataCorrectBackupFilesShrinkRequest) SetOrderId(v int64) *GetDataCorrectBackupFilesShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataCorrectBackupFilesShrinkRequest) SetActionDetailShrink(v string) *GetDataCorrectBackupFilesShrinkRequest {
	s.ActionDetailShrink = &v
	return s
}

type GetDataCorrectBackupFilesResponseBody struct {
	RequestId              *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode              *string                                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage           *string                                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success                *bool                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	DataCorrectBackupFiles *GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles `json:"DataCorrectBackupFiles,omitempty" xml:"DataCorrectBackupFiles,omitempty" type:"Struct"`
}

func (s GetDataCorrectBackupFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectBackupFilesResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCorrectBackupFilesResponseBody) SetRequestId(v string) *GetDataCorrectBackupFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCorrectBackupFilesResponseBody) SetErrorCode(v string) *GetDataCorrectBackupFilesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataCorrectBackupFilesResponseBody) SetErrorMessage(v string) *GetDataCorrectBackupFilesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataCorrectBackupFilesResponseBody) SetSuccess(v bool) *GetDataCorrectBackupFilesResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCorrectBackupFilesResponseBody) SetDataCorrectBackupFiles(v *GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles) *GetDataCorrectBackupFilesResponseBody {
	s.DataCorrectBackupFiles = v
	return s
}

type GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles struct {
	FileUrl []*string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty" type:"Repeated"`
}

func (s GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles) GoString() string {
	return s.String()
}

func (s *GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles) SetFileUrl(v []*string) *GetDataCorrectBackupFilesResponseBodyDataCorrectBackupFiles {
	s.FileUrl = v
	return s
}

type GetDataCorrectBackupFilesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDataCorrectBackupFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDataCorrectBackupFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataCorrectBackupFilesResponse) GoString() string {
	return s.String()
}

func (s *GetDataCorrectBackupFilesResponse) SetHeaders(v map[string]*string) *GetDataCorrectBackupFilesResponse {
	s.Headers = v
	return s
}

func (s *GetDataCorrectBackupFilesResponse) SetBody(v *GetDataCorrectBackupFilesResponseBody) *GetDataCorrectBackupFilesResponse {
	s.Body = v
	return s
}

type ListDataCorrectPreCheckSQLRequest struct {
	OrderId    *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	DbId       *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Tid        *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListDataCorrectPreCheckSQLRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataCorrectPreCheckSQLRequest) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckSQLRequest) SetOrderId(v int64) *ListDataCorrectPreCheckSQLRequest {
	s.OrderId = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLRequest) SetPageNumber(v int64) *ListDataCorrectPreCheckSQLRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLRequest) SetPageSize(v int64) *ListDataCorrectPreCheckSQLRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLRequest) SetDbId(v int64) *ListDataCorrectPreCheckSQLRequest {
	s.DbId = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLRequest) SetTid(v int64) *ListDataCorrectPreCheckSQLRequest {
	s.Tid = &v
	return s
}

type ListDataCorrectPreCheckSQLResponseBody struct {
	// Id of the request
	RequestId       *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage    *string                                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode       *string                                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	PreCheckSQLList []*ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList `json:"PreCheckSQLList,omitempty" xml:"PreCheckSQLList,omitempty" type:"Repeated"`
}

func (s ListDataCorrectPreCheckSQLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataCorrectPreCheckSQLResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckSQLResponseBody) SetRequestId(v string) *ListDataCorrectPreCheckSQLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBody) SetSuccess(v bool) *ListDataCorrectPreCheckSQLResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBody) SetErrorMessage(v string) *ListDataCorrectPreCheckSQLResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBody) SetErrorCode(v string) *ListDataCorrectPreCheckSQLResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBody) SetPreCheckSQLList(v []*ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) *ListDataCorrectPreCheckSQLResponseBody {
	s.PreCheckSQLList = v
	return s
}

type ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList struct {
	CheckSQL          *string `json:"CheckSQL,omitempty" xml:"CheckSQL,omitempty"`
	AffectRows        *int64  `json:"AffectRows,omitempty" xml:"AffectRows,omitempty"`
	DbId              *int64  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	SQLReviewQueryKey *string `json:"SQLReviewQueryKey,omitempty" xml:"SQLReviewQueryKey,omitempty"`
	SqlReviewStatus   *string `json:"SqlReviewStatus,omitempty" xml:"SqlReviewStatus,omitempty"`
	SqlType           *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
}

func (s ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) String() string {
	return tea.Prettify(s)
}

func (s ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) SetCheckSQL(v string) *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList {
	s.CheckSQL = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) SetAffectRows(v int64) *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList {
	s.AffectRows = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) SetDbId(v int64) *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList {
	s.DbId = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) SetSQLReviewQueryKey(v string) *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList {
	s.SQLReviewQueryKey = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) SetSqlReviewStatus(v string) *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList {
	s.SqlReviewStatus = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) SetSqlType(v string) *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList {
	s.SqlType = &v
	return s
}

type ListDataCorrectPreCheckSQLResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDataCorrectPreCheckSQLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataCorrectPreCheckSQLResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataCorrectPreCheckSQLResponse) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckSQLResponse) SetHeaders(v map[string]*string) *ListDataCorrectPreCheckSQLResponse {
	s.Headers = v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponse) SetBody(v *ListDataCorrectPreCheckSQLResponseBody) *ListDataCorrectPreCheckSQLResponse {
	s.Body = v
	return s
}

type RegisterInstanceRequest struct {
	Tid              *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceSource   *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	NetworkType      *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	EnvType          *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	EcsInstanceId    *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	VpcId            *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	EcsRegion        *string `json:"EcsRegion,omitempty" xml:"EcsRegion,omitempty"`
	Host             *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Sid              *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	DatabaseUser     *string `json:"DatabaseUser,omitempty" xml:"DatabaseUser,omitempty"`
	DatabasePassword *string `json:"DatabasePassword,omitempty" xml:"DatabasePassword,omitempty"`
	InstanceAlias    *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	DbaUid           *int64  `json:"DbaUid,omitempty" xml:"DbaUid,omitempty"`
	SafeRule         *string `json:"SafeRule,omitempty" xml:"SafeRule,omitempty"`
	QueryTimeout     *int32  `json:"QueryTimeout,omitempty" xml:"QueryTimeout,omitempty"`
	ExportTimeout    *int32  `json:"ExportTimeout,omitempty" xml:"ExportTimeout,omitempty"`
	DataLinkName     *string `json:"DataLinkName,omitempty" xml:"DataLinkName,omitempty"`
	DdlOnline        *int32  `json:"DdlOnline,omitempty" xml:"DdlOnline,omitempty"`
	UseDsql          *int32  `json:"UseDsql,omitempty" xml:"UseDsql,omitempty"`
	SkipTest         *bool   `json:"SkipTest,omitempty" xml:"SkipTest,omitempty"`
}

func (s RegisterInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterInstanceRequest) GoString() string {
	return s.String()
}

func (s *RegisterInstanceRequest) SetTid(v int64) *RegisterInstanceRequest {
	s.Tid = &v
	return s
}

func (s *RegisterInstanceRequest) SetInstanceType(v string) *RegisterInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *RegisterInstanceRequest) SetInstanceSource(v string) *RegisterInstanceRequest {
	s.InstanceSource = &v
	return s
}

func (s *RegisterInstanceRequest) SetNetworkType(v string) *RegisterInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *RegisterInstanceRequest) SetEnvType(v string) *RegisterInstanceRequest {
	s.EnvType = &v
	return s
}

func (s *RegisterInstanceRequest) SetEcsInstanceId(v string) *RegisterInstanceRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *RegisterInstanceRequest) SetVpcId(v string) *RegisterInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *RegisterInstanceRequest) SetEcsRegion(v string) *RegisterInstanceRequest {
	s.EcsRegion = &v
	return s
}

func (s *RegisterInstanceRequest) SetHost(v string) *RegisterInstanceRequest {
	s.Host = &v
	return s
}

func (s *RegisterInstanceRequest) SetPort(v int32) *RegisterInstanceRequest {
	s.Port = &v
	return s
}

func (s *RegisterInstanceRequest) SetSid(v string) *RegisterInstanceRequest {
	s.Sid = &v
	return s
}

func (s *RegisterInstanceRequest) SetDatabaseUser(v string) *RegisterInstanceRequest {
	s.DatabaseUser = &v
	return s
}

func (s *RegisterInstanceRequest) SetDatabasePassword(v string) *RegisterInstanceRequest {
	s.DatabasePassword = &v
	return s
}

func (s *RegisterInstanceRequest) SetInstanceAlias(v string) *RegisterInstanceRequest {
	s.InstanceAlias = &v
	return s
}

func (s *RegisterInstanceRequest) SetDbaUid(v int64) *RegisterInstanceRequest {
	s.DbaUid = &v
	return s
}

func (s *RegisterInstanceRequest) SetSafeRule(v string) *RegisterInstanceRequest {
	s.SafeRule = &v
	return s
}

func (s *RegisterInstanceRequest) SetQueryTimeout(v int32) *RegisterInstanceRequest {
	s.QueryTimeout = &v
	return s
}

func (s *RegisterInstanceRequest) SetExportTimeout(v int32) *RegisterInstanceRequest {
	s.ExportTimeout = &v
	return s
}

func (s *RegisterInstanceRequest) SetDataLinkName(v string) *RegisterInstanceRequest {
	s.DataLinkName = &v
	return s
}

func (s *RegisterInstanceRequest) SetDdlOnline(v int32) *RegisterInstanceRequest {
	s.DdlOnline = &v
	return s
}

func (s *RegisterInstanceRequest) SetUseDsql(v int32) *RegisterInstanceRequest {
	s.UseDsql = &v
	return s
}

func (s *RegisterInstanceRequest) SetSkipTest(v bool) *RegisterInstanceRequest {
	s.SkipTest = &v
	return s
}

type RegisterInstanceResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RegisterInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterInstanceResponseBody) SetRequestId(v string) *RegisterInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterInstanceResponseBody) SetErrorCode(v string) *RegisterInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RegisterInstanceResponseBody) SetErrorMessage(v string) *RegisterInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RegisterInstanceResponseBody) SetSuccess(v bool) *RegisterInstanceResponseBody {
	s.Success = &v
	return s
}

type RegisterInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterInstanceResponse) GoString() string {
	return s.String()
}

func (s *RegisterInstanceResponse) SetHeaders(v map[string]*string) *RegisterInstanceResponse {
	s.Headers = v
	return s
}

func (s *RegisterInstanceResponse) SetBody(v *RegisterInstanceResponseBody) *RegisterInstanceResponse {
	s.Body = v
	return s
}

type CreateStructSyncOrderRequest struct {
	Comment         *string                            `json:"Comment,omitempty" xml:"Comment,omitempty"`
	RelatedUserList []*int64                           `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	Param           *CreateStructSyncOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	AttachmentKey   *string                            `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	Tid             *int64                             `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateStructSyncOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStructSyncOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderRequest) SetComment(v string) *CreateStructSyncOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateStructSyncOrderRequest) SetRelatedUserList(v []*int64) *CreateStructSyncOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateStructSyncOrderRequest) SetParam(v *CreateStructSyncOrderRequestParam) *CreateStructSyncOrderRequest {
	s.Param = v
	return s
}

func (s *CreateStructSyncOrderRequest) SetAttachmentKey(v string) *CreateStructSyncOrderRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateStructSyncOrderRequest) SetTid(v int64) *CreateStructSyncOrderRequest {
	s.Tid = &v
	return s
}

type CreateStructSyncOrderRequestParam struct {
	Source        *CreateStructSyncOrderRequestParamSource          `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Target        *CreateStructSyncOrderRequestParamTarget          `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	IgnoreError   *bool                                             `json:"IgnoreError,omitempty" xml:"IgnoreError,omitempty"`
	TableInfoList []*CreateStructSyncOrderRequestParamTableInfoList `json:"TableInfoList,omitempty" xml:"TableInfoList,omitempty" type:"Repeated"`
}

func (s CreateStructSyncOrderRequestParam) String() string {
	return tea.Prettify(s)
}

func (s CreateStructSyncOrderRequestParam) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderRequestParam) SetSource(v *CreateStructSyncOrderRequestParamSource) *CreateStructSyncOrderRequestParam {
	s.Source = v
	return s
}

func (s *CreateStructSyncOrderRequestParam) SetTarget(v *CreateStructSyncOrderRequestParamTarget) *CreateStructSyncOrderRequestParam {
	s.Target = v
	return s
}

func (s *CreateStructSyncOrderRequestParam) SetIgnoreError(v bool) *CreateStructSyncOrderRequestParam {
	s.IgnoreError = &v
	return s
}

func (s *CreateStructSyncOrderRequestParam) SetTableInfoList(v []*CreateStructSyncOrderRequestParamTableInfoList) *CreateStructSyncOrderRequestParam {
	s.TableInfoList = v
	return s
}

type CreateStructSyncOrderRequestParamSource struct {
	DbId         *int64  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	DbSearchName *string `json:"DbSearchName,omitempty" xml:"DbSearchName,omitempty"`
	Logic        *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateStructSyncOrderRequestParamSource) String() string {
	return tea.Prettify(s)
}

func (s CreateStructSyncOrderRequestParamSource) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderRequestParamSource) SetDbId(v int64) *CreateStructSyncOrderRequestParamSource {
	s.DbId = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamSource) SetDbSearchName(v string) *CreateStructSyncOrderRequestParamSource {
	s.DbSearchName = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamSource) SetLogic(v bool) *CreateStructSyncOrderRequestParamSource {
	s.Logic = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamSource) SetVersionId(v string) *CreateStructSyncOrderRequestParamSource {
	s.VersionId = &v
	return s
}

type CreateStructSyncOrderRequestParamTarget struct {
	DbId         *int64  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	DbSearchName *string `json:"DbSearchName,omitempty" xml:"DbSearchName,omitempty"`
	Logic        *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateStructSyncOrderRequestParamTarget) String() string {
	return tea.Prettify(s)
}

func (s CreateStructSyncOrderRequestParamTarget) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderRequestParamTarget) SetDbId(v int64) *CreateStructSyncOrderRequestParamTarget {
	s.DbId = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamTarget) SetDbSearchName(v string) *CreateStructSyncOrderRequestParamTarget {
	s.DbSearchName = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamTarget) SetLogic(v bool) *CreateStructSyncOrderRequestParamTarget {
	s.Logic = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamTarget) SetVersionId(v string) *CreateStructSyncOrderRequestParamTarget {
	s.VersionId = &v
	return s
}

type CreateStructSyncOrderRequestParamTableInfoList struct {
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
}

func (s CreateStructSyncOrderRequestParamTableInfoList) String() string {
	return tea.Prettify(s)
}

func (s CreateStructSyncOrderRequestParamTableInfoList) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderRequestParamTableInfoList) SetSourceTableName(v string) *CreateStructSyncOrderRequestParamTableInfoList {
	s.SourceTableName = &v
	return s
}

func (s *CreateStructSyncOrderRequestParamTableInfoList) SetTargetTableName(v string) *CreateStructSyncOrderRequestParamTableInfoList {
	s.TargetTableName = &v
	return s
}

type CreateStructSyncOrderShrinkRequest struct {
	Comment               *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	RelatedUserListShrink *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	ParamShrink           *string `json:"Param,omitempty" xml:"Param,omitempty"`
	AttachmentKey         *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	Tid                   *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateStructSyncOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStructSyncOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderShrinkRequest) SetComment(v string) *CreateStructSyncOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateStructSyncOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateStructSyncOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateStructSyncOrderShrinkRequest) SetParamShrink(v string) *CreateStructSyncOrderShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *CreateStructSyncOrderShrinkRequest) SetAttachmentKey(v string) *CreateStructSyncOrderShrinkRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateStructSyncOrderShrinkRequest) SetTid(v int64) *CreateStructSyncOrderShrinkRequest {
	s.Tid = &v
	return s
}

type CreateStructSyncOrderResponseBody struct {
	// Id of the request
	RequestId         *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success           *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage      *string  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode         *string  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	CreateOrderResult []*int64 `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Repeated"`
}

func (s CreateStructSyncOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStructSyncOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderResponseBody) SetRequestId(v string) *CreateStructSyncOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStructSyncOrderResponseBody) SetSuccess(v bool) *CreateStructSyncOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStructSyncOrderResponseBody) SetErrorMessage(v string) *CreateStructSyncOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateStructSyncOrderResponseBody) SetErrorCode(v string) *CreateStructSyncOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateStructSyncOrderResponseBody) SetCreateOrderResult(v []*int64) *CreateStructSyncOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

type CreateStructSyncOrderResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateStructSyncOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStructSyncOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStructSyncOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderResponse) SetHeaders(v map[string]*string) *CreateStructSyncOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateStructSyncOrderResponse) SetBody(v *CreateStructSyncOrderResponseBody) *CreateStructSyncOrderResponse {
	s.Body = v
	return s
}

type ExecuteDataExportRequest struct {
	Tid          *int64                 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	OrderId      *int64                 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ActionDetail map[string]interface{} `json:"ActionDetail,omitempty" xml:"ActionDetail,omitempty"`
}

func (s ExecuteDataExportRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteDataExportRequest) GoString() string {
	return s.String()
}

func (s *ExecuteDataExportRequest) SetTid(v int64) *ExecuteDataExportRequest {
	s.Tid = &v
	return s
}

func (s *ExecuteDataExportRequest) SetOrderId(v int64) *ExecuteDataExportRequest {
	s.OrderId = &v
	return s
}

func (s *ExecuteDataExportRequest) SetActionDetail(v map[string]interface{}) *ExecuteDataExportRequest {
	s.ActionDetail = v
	return s
}

type ExecuteDataExportShrinkRequest struct {
	Tid                *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	OrderId            *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ActionDetailShrink *string `json:"ActionDetail,omitempty" xml:"ActionDetail,omitempty"`
}

func (s ExecuteDataExportShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteDataExportShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExecuteDataExportShrinkRequest) SetTid(v int64) *ExecuteDataExportShrinkRequest {
	s.Tid = &v
	return s
}

func (s *ExecuteDataExportShrinkRequest) SetOrderId(v int64) *ExecuteDataExportShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *ExecuteDataExportShrinkRequest) SetActionDetailShrink(v string) *ExecuteDataExportShrinkRequest {
	s.ActionDetailShrink = &v
	return s
}

type ExecuteDataExportResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteDataExportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteDataExportResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteDataExportResponseBody) SetRequestId(v string) *ExecuteDataExportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteDataExportResponseBody) SetErrorCode(v string) *ExecuteDataExportResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ExecuteDataExportResponseBody) SetErrorMessage(v string) *ExecuteDataExportResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ExecuteDataExportResponseBody) SetSuccess(v bool) *ExecuteDataExportResponseBody {
	s.Success = &v
	return s
}

type ExecuteDataExportResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteDataExportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteDataExportResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteDataExportResponse) GoString() string {
	return s.String()
}

func (s *ExecuteDataExportResponse) SetHeaders(v map[string]*string) *ExecuteDataExportResponse {
	s.Headers = v
	return s
}

func (s *ExecuteDataExportResponse) SetBody(v *ExecuteDataExportResponseBody) *ExecuteDataExportResponse {
	s.Body = v
	return s
}

type ExecuteDataCorrectRequest struct {
	Tid          *string                `json:"Tid,omitempty" xml:"Tid,omitempty"`
	OrderId      *int64                 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ActionDetail map[string]interface{} `json:"ActionDetail,omitempty" xml:"ActionDetail,omitempty"`
}

func (s ExecuteDataCorrectRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteDataCorrectRequest) GoString() string {
	return s.String()
}

func (s *ExecuteDataCorrectRequest) SetTid(v string) *ExecuteDataCorrectRequest {
	s.Tid = &v
	return s
}

func (s *ExecuteDataCorrectRequest) SetOrderId(v int64) *ExecuteDataCorrectRequest {
	s.OrderId = &v
	return s
}

func (s *ExecuteDataCorrectRequest) SetActionDetail(v map[string]interface{}) *ExecuteDataCorrectRequest {
	s.ActionDetail = v
	return s
}

type ExecuteDataCorrectShrinkRequest struct {
	Tid                *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
	OrderId            *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ActionDetailShrink *string `json:"ActionDetail,omitempty" xml:"ActionDetail,omitempty"`
}

func (s ExecuteDataCorrectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteDataCorrectShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExecuteDataCorrectShrinkRequest) SetTid(v string) *ExecuteDataCorrectShrinkRequest {
	s.Tid = &v
	return s
}

func (s *ExecuteDataCorrectShrinkRequest) SetOrderId(v int64) *ExecuteDataCorrectShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *ExecuteDataCorrectShrinkRequest) SetActionDetailShrink(v string) *ExecuteDataCorrectShrinkRequest {
	s.ActionDetailShrink = &v
	return s
}

type ExecuteDataCorrectResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteDataCorrectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteDataCorrectResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteDataCorrectResponseBody) SetRequestId(v string) *ExecuteDataCorrectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteDataCorrectResponseBody) SetErrorCode(v string) *ExecuteDataCorrectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ExecuteDataCorrectResponseBody) SetErrorMessage(v string) *ExecuteDataCorrectResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ExecuteDataCorrectResponseBody) SetSuccess(v bool) *ExecuteDataCorrectResponseBody {
	s.Success = &v
	return s
}

type ExecuteDataCorrectResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteDataCorrectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteDataCorrectResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteDataCorrectResponse) GoString() string {
	return s.String()
}

func (s *ExecuteDataCorrectResponse) SetHeaders(v map[string]*string) *ExecuteDataCorrectResponse {
	s.Headers = v
	return s
}

func (s *ExecuteDataCorrectResponse) SetBody(v *ExecuteDataCorrectResponseBody) *ExecuteDataCorrectResponse {
	s.Body = v
	return s
}

type ListDataCorrectPreCheckDBRequest struct {
	OrderId    *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tid        *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListDataCorrectPreCheckDBRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataCorrectPreCheckDBRequest) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckDBRequest) SetOrderId(v int64) *ListDataCorrectPreCheckDBRequest {
	s.OrderId = &v
	return s
}

func (s *ListDataCorrectPreCheckDBRequest) SetPageNumber(v int64) *ListDataCorrectPreCheckDBRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataCorrectPreCheckDBRequest) SetPageSize(v int64) *ListDataCorrectPreCheckDBRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataCorrectPreCheckDBRequest) SetTid(v int64) *ListDataCorrectPreCheckDBRequest {
	s.Tid = &v
	return s
}

type ListDataCorrectPreCheckDBResponseBody struct {
	// Id of the request
	RequestId      *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage   *string                                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode      *string                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	PreCheckDBList []*ListDataCorrectPreCheckDBResponseBodyPreCheckDBList `json:"PreCheckDBList,omitempty" xml:"PreCheckDBList,omitempty" type:"Repeated"`
}

func (s ListDataCorrectPreCheckDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataCorrectPreCheckDBResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckDBResponseBody) SetRequestId(v string) *ListDataCorrectPreCheckDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCorrectPreCheckDBResponseBody) SetSuccess(v bool) *ListDataCorrectPreCheckDBResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataCorrectPreCheckDBResponseBody) SetErrorMessage(v string) *ListDataCorrectPreCheckDBResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataCorrectPreCheckDBResponseBody) SetErrorCode(v string) *ListDataCorrectPreCheckDBResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataCorrectPreCheckDBResponseBody) SetPreCheckDBList(v []*ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) *ListDataCorrectPreCheckDBResponseBody {
	s.PreCheckDBList = v
	return s
}

type ListDataCorrectPreCheckDBResponseBodyPreCheckDBList struct {
	DbId       *int64  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	SqlNum     *int64  `json:"SqlNum,omitempty" xml:"SqlNum,omitempty"`
}

func (s ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) String() string {
	return tea.Prettify(s)
}

func (s ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) SetDbId(v int64) *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList {
	s.DbId = &v
	return s
}

func (s *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) SetSearchName(v string) *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList {
	s.SearchName = &v
	return s
}

func (s *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) SetSqlNum(v int64) *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList {
	s.SqlNum = &v
	return s
}

type ListDataCorrectPreCheckDBResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDataCorrectPreCheckDBResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataCorrectPreCheckDBResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataCorrectPreCheckDBResponse) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckDBResponse) SetHeaders(v map[string]*string) *ListDataCorrectPreCheckDBResponse {
	s.Headers = v
	return s
}

func (s *ListDataCorrectPreCheckDBResponse) SetBody(v *ListDataCorrectPreCheckDBResponseBody) *ListDataCorrectPreCheckDBResponse {
	s.Body = v
	return s
}

type ListTablesRequest struct {
	Tid        *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	ReturnGuid *bool   `json:"ReturnGuid,omitempty" xml:"ReturnGuid,omitempty"`
}

func (s ListTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTablesRequest) SetTid(v int64) *ListTablesRequest {
	s.Tid = &v
	return s
}

func (s *ListTablesRequest) SetDatabaseId(v string) *ListTablesRequest {
	s.DatabaseId = &v
	return s
}

func (s *ListTablesRequest) SetPageNumber(v int32) *ListTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTablesRequest) SetPageSize(v int32) *ListTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTablesRequest) SetSearchName(v string) *ListTablesRequest {
	s.SearchName = &v
	return s
}

func (s *ListTablesRequest) SetReturnGuid(v bool) *ListTablesRequest {
	s.ReturnGuid = &v
	return s
}

type ListTablesResponseBody struct {
	TotalCount   *int64                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	TableList    *ListTablesResponseBodyTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Struct"`
	ErrorMessage *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBody) SetTotalCount(v int64) *ListTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTablesResponseBody) SetRequestId(v string) *ListTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTablesResponseBody) SetErrorCode(v string) *ListTablesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTablesResponseBody) SetTableList(v *ListTablesResponseBodyTableList) *ListTablesResponseBody {
	s.TableList = v
	return s
}

func (s *ListTablesResponseBody) SetErrorMessage(v string) *ListTablesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTablesResponseBody) SetSuccess(v bool) *ListTablesResponseBody {
	s.Success = &v
	return s
}

type ListTablesResponseBodyTableList struct {
	Table []*ListTablesResponseBodyTableListTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s ListTablesResponseBodyTableList) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBodyTableList) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyTableList) SetTable(v []*ListTablesResponseBodyTableListTable) *ListTablesResponseBodyTableList {
	s.Table = v
	return s
}

type ListTablesResponseBodyTableListTable struct {
	DatabaseId      *string                                            `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	TableName       *string                                            `json:"TableName,omitempty" xml:"TableName,omitempty"`
	StoreCapacity   *int64                                             `json:"StoreCapacity,omitempty" xml:"StoreCapacity,omitempty"`
	OwnerIdList     *ListTablesResponseBodyTableListTableOwnerIdList   `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	Description     *string                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Encoding        *string                                            `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	OwnerNameList   *ListTablesResponseBodyTableListTableOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	TableSchemaName *string                                            `json:"TableSchemaName,omitempty" xml:"TableSchemaName,omitempty"`
	TableType       *string                                            `json:"TableType,omitempty" xml:"TableType,omitempty"`
	TableGuid       *string                                            `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	Engine          *string                                            `json:"Engine,omitempty" xml:"Engine,omitempty"`
	NumRows         *int64                                             `json:"NumRows,omitempty" xml:"NumRows,omitempty"`
	TableId         *string                                            `json:"TableId,omitempty" xml:"TableId,omitempty"`
}

func (s ListTablesResponseBodyTableListTable) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBodyTableListTable) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyTableListTable) SetDatabaseId(v string) *ListTablesResponseBodyTableListTable {
	s.DatabaseId = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetTableName(v string) *ListTablesResponseBodyTableListTable {
	s.TableName = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetStoreCapacity(v int64) *ListTablesResponseBodyTableListTable {
	s.StoreCapacity = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetOwnerIdList(v *ListTablesResponseBodyTableListTableOwnerIdList) *ListTablesResponseBodyTableListTable {
	s.OwnerIdList = v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetDescription(v string) *ListTablesResponseBodyTableListTable {
	s.Description = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetEncoding(v string) *ListTablesResponseBodyTableListTable {
	s.Encoding = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetOwnerNameList(v *ListTablesResponseBodyTableListTableOwnerNameList) *ListTablesResponseBodyTableListTable {
	s.OwnerNameList = v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetTableSchemaName(v string) *ListTablesResponseBodyTableListTable {
	s.TableSchemaName = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetTableType(v string) *ListTablesResponseBodyTableListTable {
	s.TableType = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetTableGuid(v string) *ListTablesResponseBodyTableListTable {
	s.TableGuid = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetEngine(v string) *ListTablesResponseBodyTableListTable {
	s.Engine = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetNumRows(v int64) *ListTablesResponseBodyTableListTable {
	s.NumRows = &v
	return s
}

func (s *ListTablesResponseBodyTableListTable) SetTableId(v string) *ListTablesResponseBodyTableListTable {
	s.TableId = &v
	return s
}

type ListTablesResponseBodyTableListTableOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s ListTablesResponseBodyTableListTableOwnerIdList) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBodyTableListTableOwnerIdList) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyTableListTableOwnerIdList) SetOwnerIds(v []*string) *ListTablesResponseBodyTableListTableOwnerIdList {
	s.OwnerIds = v
	return s
}

type ListTablesResponseBodyTableListTableOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s ListTablesResponseBodyTableListTableOwnerNameList) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBodyTableListTableOwnerNameList) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyTableListTableOwnerNameList) SetOwnerNames(v []*string) *ListTablesResponseBodyTableListTableOwnerNameList {
	s.OwnerNames = v
	return s
}

type ListTablesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponse) GoString() string {
	return s.String()
}

func (s *ListTablesResponse) SetHeaders(v map[string]*string) *ListTablesResponse {
	s.Headers = v
	return s
}

func (s *ListTablesResponse) SetBody(v *ListTablesResponseBody) *ListTablesResponse {
	s.Body = v
	return s
}

type ListWorkFlowNodesRequest struct {
	Tid        *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s ListWorkFlowNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkFlowNodesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkFlowNodesRequest) SetTid(v int64) *ListWorkFlowNodesRequest {
	s.Tid = &v
	return s
}

func (s *ListWorkFlowNodesRequest) SetSearchName(v string) *ListWorkFlowNodesRequest {
	s.SearchName = &v
	return s
}

type ListWorkFlowNodesResponseBody struct {
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode     *string                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage  *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	WorkflowNodes *ListWorkFlowNodesResponseBodyWorkflowNodes `json:"WorkflowNodes,omitempty" xml:"WorkflowNodes,omitempty" type:"Struct"`
	Success       *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWorkFlowNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkFlowNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkFlowNodesResponseBody) SetRequestId(v string) *ListWorkFlowNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkFlowNodesResponseBody) SetErrorCode(v string) *ListWorkFlowNodesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkFlowNodesResponseBody) SetErrorMessage(v string) *ListWorkFlowNodesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListWorkFlowNodesResponseBody) SetWorkflowNodes(v *ListWorkFlowNodesResponseBodyWorkflowNodes) *ListWorkFlowNodesResponseBody {
	s.WorkflowNodes = v
	return s
}

func (s *ListWorkFlowNodesResponseBody) SetSuccess(v bool) *ListWorkFlowNodesResponseBody {
	s.Success = &v
	return s
}

type ListWorkFlowNodesResponseBodyWorkflowNodes struct {
	WorkflowNode []*ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode `json:"WorkflowNode,omitempty" xml:"WorkflowNode,omitempty" type:"Repeated"`
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodes) String() string {
	return tea.Prettify(s)
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodes) GoString() string {
	return s.String()
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodes) SetWorkflowNode(v []*ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) *ListWorkFlowNodesResponseBodyWorkflowNodes {
	s.WorkflowNode = v
	return s
}

type ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode struct {
	Comment            *string                                                           `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreateUserNickName *string                                                           `json:"CreateUserNickName,omitempty" xml:"CreateUserNickName,omitempty"`
	NodeType           *string                                                           `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	NodeName           *string                                                           `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	AuditUsers         *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers `json:"AuditUsers,omitempty" xml:"AuditUsers,omitempty" type:"Struct"`
	CreateUserId       *int64                                                            `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	NodeId             *int64                                                            `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) String() string {
	return tea.Prettify(s)
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) GoString() string {
	return s.String()
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) SetComment(v string) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode {
	s.Comment = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) SetCreateUserNickName(v string) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode {
	s.CreateUserNickName = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) SetNodeType(v string) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode {
	s.NodeType = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) SetNodeName(v string) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode {
	s.NodeName = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) SetAuditUsers(v *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode {
	s.AuditUsers = v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) SetCreateUserId(v int64) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode {
	s.CreateUserId = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode) SetNodeId(v int64) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNode {
	s.NodeId = &v
	return s
}

type ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers struct {
	AuditUser []*ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser `json:"AuditUser,omitempty" xml:"AuditUser,omitempty" type:"Repeated"`
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers) String() string {
	return tea.Prettify(s)
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers) GoString() string {
	return s.String()
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers) SetAuditUser(v []*ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsers {
	s.AuditUser = v
	return s
}

type ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser struct {
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	UserId   *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) String() string {
	return tea.Prettify(s)
}

func (s ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) GoString() string {
	return s.String()
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) SetRealName(v string) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser {
	s.RealName = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) SetUserId(v int64) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser {
	s.UserId = &v
	return s
}

func (s *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser) SetNickName(v string) *ListWorkFlowNodesResponseBodyWorkflowNodesWorkflowNodeAuditUsersAuditUser {
	s.NickName = &v
	return s
}

type ListWorkFlowNodesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListWorkFlowNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWorkFlowNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkFlowNodesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkFlowNodesResponse) SetHeaders(v map[string]*string) *ListWorkFlowNodesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkFlowNodesResponse) SetBody(v *ListWorkFlowNodesResponseBody) *ListWorkFlowNodesResponse {
	s.Body = v
	return s
}

type GetStructSyncOrderDetailRequest struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Tid     *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetStructSyncOrderDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailRequest) SetOrderId(v int64) *GetStructSyncOrderDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetStructSyncOrderDetailRequest) SetTid(v int64) *GetStructSyncOrderDetailRequest {
	s.Tid = &v
	return s
}

type GetStructSyncOrderDetailResponseBody struct {
	// Id of the request
	RequestId             *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success               *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage          *string                                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode             *string                                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	StructSyncOrderDetail *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail `json:"StructSyncOrderDetail,omitempty" xml:"StructSyncOrderDetail,omitempty" type:"Struct"`
}

func (s GetStructSyncOrderDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponseBody) SetRequestId(v string) *GetStructSyncOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBody) SetSuccess(v bool) *GetStructSyncOrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBody) SetErrorMessage(v string) *GetStructSyncOrderDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBody) SetErrorCode(v string) *GetStructSyncOrderDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBody) SetStructSyncOrderDetail(v *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) *GetStructSyncOrderDetailResponseBody {
	s.StructSyncOrderDetail = v
	return s
}

type GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail struct {
	SourceDatabaseInfo *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo `json:"SourceDatabaseInfo,omitempty" xml:"SourceDatabaseInfo,omitempty" type:"Struct"`
	TargetDatabaseInfo *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo `json:"TargetDatabaseInfo,omitempty" xml:"TargetDatabaseInfo,omitempty" type:"Struct"`
	SourceType         *string                                                                      `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceVersionInfo  *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo  `json:"SourceVersionInfo,omitempty" xml:"SourceVersionInfo,omitempty" type:"Struct"`
	TargetType         *string                                                                      `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TargetVersionInfo  *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo  `json:"TargetVersionInfo,omitempty" xml:"TargetVersionInfo,omitempty" type:"Struct"`
	TableInfoList      []*GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList    `json:"TableInfoList,omitempty" xml:"TableInfoList,omitempty" type:"Repeated"`
	IgnoreError        *bool                                                                        `json:"IgnoreError,omitempty" xml:"IgnoreError,omitempty"`
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetSourceDatabaseInfo(v *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.SourceDatabaseInfo = v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetTargetDatabaseInfo(v *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.TargetDatabaseInfo = v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetSourceType(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.SourceType = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetSourceVersionInfo(v *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.SourceVersionInfo = v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetTargetType(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.TargetType = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetTargetVersionInfo(v *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.TargetVersionInfo = v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetTableInfoList(v []*GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.TableInfoList = v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail) SetIgnoreError(v bool) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetail {
	s.IgnoreError = &v
	return s
}

type GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo struct {
	DbId       *int64  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	DbType     *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	EnvType    *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Logic      *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) SetDbId(v int64) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo {
	s.DbId = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) SetSearchName(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo {
	s.SearchName = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) SetDbType(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo {
	s.DbType = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) SetEnvType(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo {
	s.EnvType = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo) SetLogic(v bool) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceDatabaseInfo {
	s.Logic = &v
	return s
}

type GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo struct {
	DbId       *int64  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	DbType     *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	EnvType    *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Logic      *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) SetDbId(v int64) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo {
	s.DbId = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) SetSearchName(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo {
	s.SearchName = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) SetDbType(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo {
	s.DbType = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) SetEnvType(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo {
	s.EnvType = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo) SetLogic(v bool) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetDatabaseInfo {
	s.Logic = &v
	return s
}

type GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo struct {
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo) SetVersionId(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailSourceVersionInfo {
	s.VersionId = &v
	return s
}

type GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo struct {
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo) SetVersionId(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTargetVersionInfo {
	s.VersionId = &v
	return s
}

type GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList struct {
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList) SetSourceTableName(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList {
	s.SourceTableName = &v
	return s
}

func (s *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList) SetTargetTableName(v string) *GetStructSyncOrderDetailResponseBodyStructSyncOrderDetailTableInfoList {
	s.TargetTableName = &v
	return s
}

type GetStructSyncOrderDetailResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStructSyncOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStructSyncOrderDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponse) SetHeaders(v map[string]*string) *GetStructSyncOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetStructSyncOrderDetailResponse) SetBody(v *GetStructSyncOrderDetailResponseBody) *GetStructSyncOrderDetailResponse {
	s.Body = v
	return s
}

type ListSensitiveColumnsDetailRequest struct {
	Tid        *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName  *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
}

func (s ListSensitiveColumnsDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveColumnsDetailRequest) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsDetailRequest) SetTid(v int64) *ListSensitiveColumnsDetailRequest {
	s.Tid = &v
	return s
}

func (s *ListSensitiveColumnsDetailRequest) SetSchemaName(v string) *ListSensitiveColumnsDetailRequest {
	s.SchemaName = &v
	return s
}

func (s *ListSensitiveColumnsDetailRequest) SetTableName(v string) *ListSensitiveColumnsDetailRequest {
	s.TableName = &v
	return s
}

func (s *ListSensitiveColumnsDetailRequest) SetColumnName(v string) *ListSensitiveColumnsDetailRequest {
	s.ColumnName = &v
	return s
}

type ListSensitiveColumnsDetailResponseBody struct {
	RequestId                  *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SensitiveColumnsDetailList *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList `json:"SensitiveColumnsDetailList,omitempty" xml:"SensitiveColumnsDetailList,omitempty" type:"Struct"`
	ErrorCode                  *string                                                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage               *string                                                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success                    *bool                                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSensitiveColumnsDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveColumnsDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsDetailResponseBody) SetRequestId(v string) *ListSensitiveColumnsDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBody) SetSensitiveColumnsDetailList(v *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList) *ListSensitiveColumnsDetailResponseBody {
	s.SensitiveColumnsDetailList = v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBody) SetErrorCode(v string) *ListSensitiveColumnsDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBody) SetErrorMessage(v string) *ListSensitiveColumnsDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBody) SetSuccess(v bool) *ListSensitiveColumnsDetailResponseBody {
	s.Success = &v
	return s
}

type ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList struct {
	SensitiveColumnsDetail []*ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail `json:"SensitiveColumnsDetail,omitempty" xml:"SensitiveColumnsDetail,omitempty" type:"Repeated"`
}

func (s ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList) SetSensitiveColumnsDetail(v []*ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList {
	s.SensitiveColumnsDetail = v
	return s
}

type ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail struct {
	DbId              *int64  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	ColumnName        *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	ColumnDescription *string `json:"ColumnDescription,omitempty" xml:"ColumnDescription,omitempty"`
	TableName         *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	DbType            *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	ColumnType        *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	Logic             *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	SchemaName        *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	SearchName        *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	EnvType           *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
}

func (s ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetDbId(v int64) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.DbId = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetColumnName(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.ColumnName = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetColumnDescription(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.ColumnDescription = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetTableName(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.TableName = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetDbType(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.DbType = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetColumnType(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.ColumnType = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetLogic(v bool) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.Logic = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetSchemaName(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.SchemaName = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetSearchName(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.SearchName = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetEnvType(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.EnvType = &v
	return s
}

type ListSensitiveColumnsDetailResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSensitiveColumnsDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSensitiveColumnsDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveColumnsDetailResponse) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsDetailResponse) SetHeaders(v map[string]*string) *ListSensitiveColumnsDetailResponse {
	s.Headers = v
	return s
}

func (s *ListSensitiveColumnsDetailResponse) SetBody(v *ListSensitiveColumnsDetailResponseBody) *ListSensitiveColumnsDetailResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	Tid             *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	Uid             *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserNick        *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
	RoleNames       *string `json:"RoleNames,omitempty" xml:"RoleNames,omitempty"`
	Mobile          *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	MaxExecuteCount *int64  `json:"MaxExecuteCount,omitempty" xml:"MaxExecuteCount,omitempty"`
	MaxResultCount  *int64  `json:"MaxResultCount,omitempty" xml:"MaxResultCount,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetTid(v int64) *UpdateUserRequest {
	s.Tid = &v
	return s
}

func (s *UpdateUserRequest) SetUid(v int64) *UpdateUserRequest {
	s.Uid = &v
	return s
}

func (s *UpdateUserRequest) SetUserNick(v string) *UpdateUserRequest {
	s.UserNick = &v
	return s
}

func (s *UpdateUserRequest) SetRoleNames(v string) *UpdateUserRequest {
	s.RoleNames = &v
	return s
}

func (s *UpdateUserRequest) SetMobile(v string) *UpdateUserRequest {
	s.Mobile = &v
	return s
}

func (s *UpdateUserRequest) SetMaxExecuteCount(v int64) *UpdateUserRequest {
	s.MaxExecuteCount = &v
	return s
}

func (s *UpdateUserRequest) SetMaxResultCount(v int64) *UpdateUserRequest {
	s.MaxResultCount = &v
	return s
}

type UpdateUserResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetErrorCode(v string) *UpdateUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateUserResponseBody) SetErrorMessage(v string) *UpdateUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateUserResponseBody) SetSuccess(v bool) *UpdateUserResponseBody {
	s.Success = &v
	return s
}

type UpdateUserResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
	s.Body = v
	return s
}

type GetPhysicalDatabaseRequest struct {
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	Tid  *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetPhysicalDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalDatabaseRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalDatabaseRequest) SetDbId(v int64) *GetPhysicalDatabaseRequest {
	s.DbId = &v
	return s
}

func (s *GetPhysicalDatabaseRequest) SetTid(v int64) *GetPhysicalDatabaseRequest {
	s.Tid = &v
	return s
}

type GetPhysicalDatabaseResponseBody struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Database     *GetPhysicalDatabaseResponseBodyDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Struct"`
	ErrorCode    *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPhysicalDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalDatabaseResponseBody) SetRequestId(v string) *GetPhysicalDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBody) SetDatabase(v *GetPhysicalDatabaseResponseBodyDatabase) *GetPhysicalDatabaseResponseBody {
	s.Database = v
	return s
}

func (s *GetPhysicalDatabaseResponseBody) SetErrorCode(v string) *GetPhysicalDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBody) SetErrorMessage(v string) *GetPhysicalDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBody) SetSuccess(v bool) *GetPhysicalDatabaseResponseBody {
	s.Success = &v
	return s
}

type GetPhysicalDatabaseResponseBodyDatabase struct {
	DatabaseId    *string                                               `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	Host          *string                                               `json:"Host,omitempty" xml:"Host,omitempty"`
	CatalogName   *string                                               `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	DbaName       *string                                               `json:"DbaName,omitempty" xml:"DbaName,omitempty"`
	State         *string                                               `json:"State,omitempty" xml:"State,omitempty"`
	DbaId         *string                                               `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	SchemaName    *string                                               `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	InstanceId    *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Port          *int32                                                `json:"Port,omitempty" xml:"Port,omitempty"`
	EnvType       *string                                               `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Sid           *string                                               `json:"Sid,omitempty" xml:"Sid,omitempty"`
	OwnerIdList   *GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList   `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	Encoding      *string                                               `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	DbType        *string                                               `json:"DbType,omitempty" xml:"DbType,omitempty"`
	OwnerNameList *GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	SearchName    *string                                               `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s GetPhysicalDatabaseResponseBodyDatabase) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalDatabaseResponseBodyDatabase) GoString() string {
	return s.String()
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetDatabaseId(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.DatabaseId = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetHost(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.Host = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetCatalogName(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.CatalogName = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetDbaName(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.DbaName = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetState(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.State = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetDbaId(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.DbaId = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetSchemaName(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.SchemaName = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetInstanceId(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.InstanceId = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetPort(v int32) *GetPhysicalDatabaseResponseBodyDatabase {
	s.Port = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetEnvType(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.EnvType = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetSid(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.Sid = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetOwnerIdList(v *GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList) *GetPhysicalDatabaseResponseBodyDatabase {
	s.OwnerIdList = v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetEncoding(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.Encoding = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetDbType(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.DbType = &v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetOwnerNameList(v *GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList) *GetPhysicalDatabaseResponseBodyDatabase {
	s.OwnerNameList = v
	return s
}

func (s *GetPhysicalDatabaseResponseBodyDatabase) SetSearchName(v string) *GetPhysicalDatabaseResponseBodyDatabase {
	s.SearchName = &v
	return s
}

type GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList) GoString() string {
	return s.String()
}

func (s *GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList) SetOwnerIds(v []*string) *GetPhysicalDatabaseResponseBodyDatabaseOwnerIdList {
	s.OwnerIds = v
	return s
}

type GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList) GoString() string {
	return s.String()
}

func (s *GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList) SetOwnerNames(v []*string) *GetPhysicalDatabaseResponseBodyDatabaseOwnerNameList {
	s.OwnerNames = v
	return s
}

type GetPhysicalDatabaseResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPhysicalDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPhysicalDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalDatabaseResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalDatabaseResponse) SetHeaders(v map[string]*string) *GetPhysicalDatabaseResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalDatabaseResponse) SetBody(v *GetPhysicalDatabaseResponseBody) *GetPhysicalDatabaseResponse {
	s.Body = v
	return s
}

type GetStructSyncExecSqlDetailRequest struct {
	OrderId    *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tid        *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetStructSyncExecSqlDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncExecSqlDetailRequest) GoString() string {
	return s.String()
}

func (s *GetStructSyncExecSqlDetailRequest) SetOrderId(v int64) *GetStructSyncExecSqlDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetStructSyncExecSqlDetailRequest) SetPageNumber(v int64) *GetStructSyncExecSqlDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *GetStructSyncExecSqlDetailRequest) SetPageSize(v int64) *GetStructSyncExecSqlDetailRequest {
	s.PageSize = &v
	return s
}

func (s *GetStructSyncExecSqlDetailRequest) SetTid(v int64) *GetStructSyncExecSqlDetailRequest {
	s.Tid = &v
	return s
}

type GetStructSyncExecSqlDetailResponseBody struct {
	// Id of the request
	RequestId               *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                 *bool                                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage            *string                                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode               *string                                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	StructSyncExecSqlDetail *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail `json:"StructSyncExecSqlDetail,omitempty" xml:"StructSyncExecSqlDetail,omitempty" type:"Struct"`
}

func (s GetStructSyncExecSqlDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncExecSqlDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetStructSyncExecSqlDetailResponseBody) SetRequestId(v string) *GetStructSyncExecSqlDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStructSyncExecSqlDetailResponseBody) SetSuccess(v bool) *GetStructSyncExecSqlDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetStructSyncExecSqlDetailResponseBody) SetErrorMessage(v string) *GetStructSyncExecSqlDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetStructSyncExecSqlDetailResponseBody) SetErrorCode(v string) *GetStructSyncExecSqlDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStructSyncExecSqlDetailResponseBody) SetStructSyncExecSqlDetail(v *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail) *GetStructSyncExecSqlDetailResponseBody {
	s.StructSyncExecSqlDetail = v
	return s
}

type GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail struct {
	TotalSqlCount *int64  `json:"TotalSqlCount,omitempty" xml:"TotalSqlCount,omitempty"`
	ExecSql       *string `json:"ExecSql,omitempty" xml:"ExecSql,omitempty"`
}

func (s GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail) GoString() string {
	return s.String()
}

func (s *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail) SetTotalSqlCount(v int64) *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail {
	s.TotalSqlCount = &v
	return s
}

func (s *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail) SetExecSql(v string) *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail {
	s.ExecSql = &v
	return s
}

type GetStructSyncExecSqlDetailResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStructSyncExecSqlDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStructSyncExecSqlDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStructSyncExecSqlDetailResponse) GoString() string {
	return s.String()
}

func (s *GetStructSyncExecSqlDetailResponse) SetHeaders(v map[string]*string) *GetStructSyncExecSqlDetailResponse {
	s.Headers = v
	return s
}

func (s *GetStructSyncExecSqlDetailResponse) SetBody(v *GetStructSyncExecSqlDetailResponseBody) *GetStructSyncExecSqlDetailResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	Tid  *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Sid  *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetTid(v int64) *DeleteInstanceRequest {
	s.Tid = &v
	return s
}

func (s *DeleteInstanceRequest) SetHost(v string) *DeleteInstanceRequest {
	s.Host = &v
	return s
}

func (s *DeleteInstanceRequest) SetPort(v int32) *DeleteInstanceRequest {
	s.Port = &v
	return s
}

func (s *DeleteInstanceRequest) SetSid(v string) *DeleteInstanceRequest {
	s.Sid = &v
	return s
}

type DeleteInstanceResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetErrorCode(v string) *DeleteInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetErrorMessage(v string) *DeleteInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetSuccess(v bool) *DeleteInstanceResponseBody {
	s.Success = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type GetTableDBTopologyRequest struct {
	Tid       *int64  `json:"Tid,omitempty" xml:"Tid,omitempty"`
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s GetTableDBTopologyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableDBTopologyRequest) GoString() string {
	return s.String()
}

func (s *GetTableDBTopologyRequest) SetTid(v int64) *GetTableDBTopologyRequest {
	s.Tid = &v
	return s
}

func (s *GetTableDBTopologyRequest) SetTableGuid(v string) *GetTableDBTopologyRequest {
	s.TableGuid = &v
	return s
}

type GetTableDBTopologyResponseBody struct {
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBTopology   *GetTableDBTopologyResponseBodyDBTopology `json:"DBTopology,omitempty" xml:"DBTopology,omitempty" type:"Struct"`
	ErrorCode    *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTableDBTopologyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableDBTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableDBTopologyResponseBody) SetRequestId(v string) *GetTableDBTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableDBTopologyResponseBody) SetDBTopology(v *GetTableDBTopologyResponseBodyDBTopology) *GetTableDBTopologyResponseBody {
	s.DBTopology = v
	return s
}

func (s *GetTableDBTopologyResponseBody) SetErrorCode(v string) *GetTableDBTopologyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTableDBTopologyResponseBody) SetErrorMessage(v string) *GetTableDBTopologyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTableDBTopologyResponseBody) SetSuccess(v bool) *GetTableDBTopologyResponseBody {
	s.Success = &v
	return s
}

type GetTableDBTopologyResponseBodyDBTopology struct {
	TableName      *string                                                   `json:"TableName,omitempty" xml:"TableName,omitempty"`
	DataSourceList []*GetTableDBTopologyResponseBodyDBTopologyDataSourceList `json:"DataSourceList,omitempty" xml:"DataSourceList,omitempty" type:"Repeated"`
	TableGuid      *string                                                   `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s GetTableDBTopologyResponseBodyDBTopology) String() string {
	return tea.Prettify(s)
}

func (s GetTableDBTopologyResponseBodyDBTopology) GoString() string {
	return s.String()
}

func (s *GetTableDBTopologyResponseBodyDBTopology) SetTableName(v string) *GetTableDBTopologyResponseBodyDBTopology {
	s.TableName = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopology) SetDataSourceList(v []*GetTableDBTopologyResponseBodyDBTopologyDataSourceList) *GetTableDBTopologyResponseBodyDBTopology {
	s.DataSourceList = v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopology) SetTableGuid(v string) *GetTableDBTopologyResponseBodyDBTopology {
	s.TableGuid = &v
	return s
}

type GetTableDBTopologyResponseBodyDBTopologyDataSourceList struct {
	Sid          *string                                                               `json:"Sid,omitempty" xml:"Sid,omitempty"`
	Host         *string                                                               `json:"Host,omitempty" xml:"Host,omitempty"`
	DbType       *string                                                               `json:"DbType,omitempty" xml:"DbType,omitempty"`
	DatabaseList []*GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList `json:"DatabaseList,omitempty" xml:"DatabaseList,omitempty" type:"Repeated"`
	Port         *int32                                                                `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s GetTableDBTopologyResponseBodyDBTopologyDataSourceList) String() string {
	return tea.Prettify(s)
}

func (s GetTableDBTopologyResponseBodyDBTopologyDataSourceList) GoString() string {
	return s.String()
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) SetSid(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceList {
	s.Sid = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) SetHost(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceList {
	s.Host = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) SetDbType(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceList {
	s.DbType = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) SetDatabaseList(v []*GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) *GetTableDBTopologyResponseBodyDBTopologyDataSourceList {
	s.DatabaseList = v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) SetPort(v int32) *GetTableDBTopologyResponseBodyDBTopologyDataSourceList {
	s.Port = &v
	return s
}

type GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList struct {
	DbId      *string                                                                        `json:"DbId,omitempty" xml:"DbId,omitempty"`
	DbName    *string                                                                        `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbType    *string                                                                        `json:"DbType,omitempty" xml:"DbType,omitempty"`
	TableList []*GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Repeated"`
	EnvType   *string                                                                        `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
}

func (s GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) String() string {
	return tea.Prettify(s)
}

func (s GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) GoString() string {
	return s.String()
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) SetDbId(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList {
	s.DbId = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) SetDbName(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList {
	s.DbName = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) SetDbType(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList {
	s.DbType = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) SetTableList(v []*GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList {
	s.TableList = v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) SetEnvType(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList {
	s.EnvType = &v
	return s
}

type GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList struct {
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
	TableId   *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
}

func (s GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) String() string {
	return tea.Prettify(s)
}

func (s GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) GoString() string {
	return s.String()
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) SetTableName(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList {
	s.TableName = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) SetTableType(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList {
	s.TableType = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) SetTableId(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList {
	s.TableId = &v
	return s
}

type GetTableDBTopologyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTableDBTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTableDBTopologyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableDBTopologyResponse) GoString() string {
	return s.String()
}

func (s *GetTableDBTopologyResponse) SetHeaders(v map[string]*string) *GetTableDBTopologyResponse {
	s.Headers = v
	return s
}

func (s *GetTableDBTopologyResponse) SetBody(v *GetTableDBTopologyResponseBody) *GetTableDBTopologyResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dms-enterprise"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) SubmitStructSyncOrderApprovalWithOptions(request *SubmitStructSyncOrderApprovalRequest, runtime *util.RuntimeOptions) (_result *SubmitStructSyncOrderApprovalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitStructSyncOrderApprovalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitStructSyncOrderApproval"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitStructSyncOrderApproval(request *SubmitStructSyncOrderApprovalRequest) (_result *SubmitStructSyncOrderApprovalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitStructSyncOrderApprovalResponse{}
	_body, _err := client.SubmitStructSyncOrderApprovalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDatabaseUserPermssionsWithOptions(request *ListDatabaseUserPermssionsRequest, runtime *util.RuntimeOptions) (_result *ListDatabaseUserPermssionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDatabaseUserPermssionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDatabaseUserPermssions"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDatabaseUserPermssions(request *ListDatabaseUserPermssionsRequest) (_result *ListDatabaseUserPermssionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDatabaseUserPermssionsResponse{}
	_body, _err := client.ListDatabaseUserPermssionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSensitiveColumnsWithOptions(request *ListSensitiveColumnsRequest, runtime *util.RuntimeOptions) (_result *ListSensitiveColumnsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSensitiveColumnsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSensitiveColumns"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSensitiveColumns(request *ListSensitiveColumnsRequest) (_result *ListSensitiveColumnsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSensitiveColumnsResponse{}
	_body, _err := client.ListSensitiveColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUsers"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitOrderApprovalWithOptions(request *SubmitOrderApprovalRequest, runtime *util.RuntimeOptions) (_result *SubmitOrderApprovalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitOrderApprovalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitOrderApproval"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitOrderApproval(request *SubmitOrderApprovalRequest) (_result *SubmitOrderApprovalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitOrderApprovalResponse{}
	_body, _err := client.SubmitOrderApprovalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantUserPermissionWithOptions(request *GrantUserPermissionRequest, runtime *util.RuntimeOptions) (_result *GrantUserPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GrantUserPermissionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GrantUserPermission"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantUserPermission(request *GrantUserPermissionRequest) (_result *GrantUserPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantUserPermissionResponse{}
	_body, _err := client.GrantUserPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMetaTableDetailInfoWithOptions(request *GetMetaTableDetailInfoRequest, runtime *util.RuntimeOptions) (_result *GetMetaTableDetailInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMetaTableDetailInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMetaTableDetailInfo"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMetaTableDetailInfo(request *GetMetaTableDetailInfoRequest) (_result *GetMetaTableDetailInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMetaTableDetailInfoResponse{}
	_body, _err := client.GetMetaTableDetailInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataCorrectSQLFileWithOptions(request *GetDataCorrectSQLFileRequest, runtime *util.RuntimeOptions) (_result *GetDataCorrectSQLFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDataCorrectSQLFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDataCorrectSQLFile"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataCorrectSQLFile(request *GetDataCorrectSQLFileRequest) (_result *GetDataCorrectSQLFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDataCorrectSQLFileResponse{}
	_body, _err := client.GetDataCorrectSQLFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFreeLockCorrectOrderWithOptions(tmpReq *CreateFreeLockCorrectOrderRequest, runtime *util.RuntimeOptions) (_result *CreateFreeLockCorrectOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateFreeLockCorrectOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedUserList)) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, tea.String("RelatedUserList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Param))) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Param), tea.String("Param"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFreeLockCorrectOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFreeLockCorrectOrder"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFreeLockCorrectOrder(request *CreateFreeLockCorrectOrderRequest) (_result *CreateFreeLockCorrectOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFreeLockCorrectOrderResponse{}
	_body, _err := client.CreateFreeLockCorrectOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrderWithOptions(tmpReq *CreateOrderRequest, runtime *util.RuntimeOptions) (_result *CreateOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PluginParam)) {
		request.PluginParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PluginParam, tea.String("PluginParam"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateOrder"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrder(request *CreateOrderRequest) (_result *CreateOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrderResponse{}
	_body, _err := client.CreateOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDatabasesWithOptions(request *ListDatabasesRequest, runtime *util.RuntimeOptions) (_result *ListDatabasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDatabasesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDatabases"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDatabases(request *ListDatabasesRequest) (_result *ListDatabasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDatabasesResponse{}
	_body, _err := client.ListDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserPermissionsWithOptions(request *ListUserPermissionsRequest, runtime *util.RuntimeOptions) (_result *ListUserPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUserPermissionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserPermissions"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserPermissions(request *ListUserPermissionsRequest) (_result *ListUserPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserPermissionsResponse{}
	_body, _err := client.ListUserPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkFlowTemplatesWithOptions(request *ListWorkFlowTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListWorkFlowTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListWorkFlowTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListWorkFlowTemplates"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkFlowTemplates(request *ListWorkFlowTemplatesRequest) (_result *ListWorkFlowTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWorkFlowTemplatesResponse{}
	_body, _err := client.ListWorkFlowTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataExportOrderDetailWithOptions(request *GetDataExportOrderDetailRequest, runtime *util.RuntimeOptions) (_result *GetDataExportOrderDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDataExportOrderDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDataExportOrderDetail"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataExportOrderDetail(request *GetDataExportOrderDetailRequest) (_result *GetDataExportOrderDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDataExportOrderDetailResponse{}
	_body, _err := client.GetDataExportOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstances"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserUploadFileJobWithOptions(request *GetUserUploadFileJobRequest, runtime *util.RuntimeOptions) (_result *GetUserUploadFileJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUserUploadFileJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUserUploadFileJob"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserUploadFileJob(request *GetUserUploadFileJobRequest) (_result *GetUserUploadFileJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserUploadFileJobResponse{}
	_body, _err := client.GetUserUploadFileJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDDLPublishRecordsWithOptions(request *ListDDLPublishRecordsRequest, runtime *util.RuntimeOptions) (_result *ListDDLPublishRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDDLPublishRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDDLPublishRecords"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDDLPublishRecords(request *ListDDLPublishRecordsRequest) (_result *ListDDLPublishRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDDLPublishRecordsResponse{}
	_body, _err := client.ListDDLPublishRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStructSyncJobDetailWithOptions(request *GetStructSyncJobDetailRequest, runtime *util.RuntimeOptions) (_result *GetStructSyncJobDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetStructSyncJobDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetStructSyncJobDetail"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStructSyncJobDetail(request *GetStructSyncJobDetailRequest) (_result *GetStructSyncJobDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStructSyncJobDetailResponse{}
	_body, _err := client.GetStructSyncJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUploadOSSFileJobWithOptions(tmpReq *CreateUploadOSSFileJobRequest, runtime *util.RuntimeOptions) (_result *CreateUploadOSSFileJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateUploadOSSFileJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UploadTarget))) {
		request.UploadTargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UploadTarget), tea.String("UploadTarget"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUploadOSSFileJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUploadOSSFileJob"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUploadOSSFileJob(request *CreateUploadOSSFileJobRequest) (_result *CreateUploadOSSFileJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUploadOSSFileJobResponse{}
	_body, _err := client.CreateUploadOSSFileJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchDatabaseWithOptions(request *SearchDatabaseRequest, runtime *util.RuntimeOptions) (_result *SearchDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchDatabaseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchDatabase"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchDatabase(request *SearchDatabaseRequest) (_result *SearchDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchDatabaseResponse{}
	_body, _err := client.SearchDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDBTopologyWithOptions(request *GetDBTopologyRequest, runtime *util.RuntimeOptions) (_result *GetDBTopologyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDBTopologyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDBTopology"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDBTopology(request *GetDBTopologyRequest) (_result *GetDBTopologyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDBTopologyResponse{}
	_body, _err := client.GetDBTopologyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSQLReviewCheckResultStatusWithOptions(request *GetSQLReviewCheckResultStatusRequest, runtime *util.RuntimeOptions) (_result *GetSQLReviewCheckResultStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSQLReviewCheckResultStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSQLReviewCheckResultStatus"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSQLReviewCheckResultStatus(request *GetSQLReviewCheckResultStatusRequest) (_result *GetSQLReviewCheckResultStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSQLReviewCheckResultStatusResponse{}
	_body, _err := client.GetSQLReviewCheckResultStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncDatabaseMetaWithOptions(request *SyncDatabaseMetaRequest, runtime *util.RuntimeOptions) (_result *SyncDatabaseMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SyncDatabaseMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SyncDatabaseMeta"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncDatabaseMeta(request *SyncDatabaseMetaRequest) (_result *SyncDatabaseMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncDatabaseMetaResponse{}
	_body, _err := client.SyncDatabaseMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUser"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteStructSyncWithOptions(request *ExecuteStructSyncRequest, runtime *util.RuntimeOptions) (_result *ExecuteStructSyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteStructSyncResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecuteStructSync"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteStructSync(request *ExecuteStructSyncRequest) (_result *ExecuteStructSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteStructSyncResponse{}
	_body, _err := client.ExecuteStructSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataCorrectOrderDetailWithOptions(request *GetDataCorrectOrderDetailRequest, runtime *util.RuntimeOptions) (_result *GetDataCorrectOrderDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDataCorrectOrderDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDataCorrectOrderDetail"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataCorrectOrderDetail(request *GetDataCorrectOrderDetailRequest) (_result *GetDataCorrectOrderDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDataCorrectOrderDetailResponse{}
	_body, _err := client.GetDataCorrectOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListColumnsWithOptions(request *ListColumnsRequest, runtime *util.RuntimeOptions) (_result *ListColumnsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListColumnsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListColumns"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListColumns(request *ListColumnsRequest) (_result *ListColumnsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListColumnsResponse{}
	_body, _err := client.ListColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeUserPermissionWithOptions(request *RevokeUserPermissionRequest, runtime *util.RuntimeOptions) (_result *RevokeUserPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RevokeUserPermissionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RevokeUserPermission"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeUserPermission(request *RevokeUserPermissionRequest) (_result *RevokeUserPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeUserPermissionResponse{}
	_body, _err := client.RevokeUserPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMetaTableColumnWithOptions(request *GetMetaTableColumnRequest, runtime *util.RuntimeOptions) (_result *GetMetaTableColumnResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMetaTableColumnResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMetaTableColumn"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMetaTableColumn(request *GetMetaTableColumnRequest) (_result *GetMetaTableColumnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMetaTableColumnResponse{}
	_body, _err := client.GetMetaTableColumnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableUserWithOptions(request *EnableUserRequest, runtime *util.RuntimeOptions) (_result *EnableUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableUser"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableUser(request *EnableUserRequest) (_result *EnableUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableUserResponse{}
	_body, _err := client.EnableUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceWithOptions(request *UpdateInstanceRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateInstance"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstance(request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteScriptWithOptions(request *ExecuteScriptRequest, runtime *util.RuntimeOptions) (_result *ExecuteScriptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteScriptResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecuteScript"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteScript(request *ExecuteScriptRequest) (_result *ExecuteScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteScriptResponse{}
	_body, _err := client.ExecuteScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDBTaskSQLJobDetailWithOptions(request *ListDBTaskSQLJobDetailRequest, runtime *util.RuntimeOptions) (_result *ListDBTaskSQLJobDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDBTaskSQLJobDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDBTaskSQLJobDetail"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDBTaskSQLJobDetail(request *ListDBTaskSQLJobDetailRequest) (_result *ListDBTaskSQLJobDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDBTaskSQLJobDetailResponse{}
	_body, _err := client.ListDBTaskSQLJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableUserWithOptions(request *DisableUserRequest, runtime *util.RuntimeOptions) (_result *DisableUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableUser"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableUser(request *DisableUserRequest) (_result *DisableUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableUserResponse{}
	_body, _err := client.DisableUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetApprovalDetailWithOptions(request *GetApprovalDetailRequest, runtime *util.RuntimeOptions) (_result *GetApprovalDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetApprovalDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetApprovalDetail"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApprovalDetail(request *GetApprovalDetailRequest) (_result *GetApprovalDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApprovalDetailResponse{}
	_body, _err := client.GetApprovalDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSQLReviewOriginSQLWithOptions(tmpReq *ListSQLReviewOriginSQLRequest, runtime *util.RuntimeOptions) (_result *ListSQLReviewOriginSQLResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListSQLReviewOriginSQLShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.OrderActionDetail))) {
		request.OrderActionDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.OrderActionDetail), tea.String("OrderActionDetail"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSQLReviewOriginSQLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSQLReviewOriginSQL"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSQLReviewOriginSQL(request *ListSQLReviewOriginSQLRequest) (_result *ListSQLReviewOriginSQLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSQLReviewOriginSQLResponse{}
	_body, _err := client.ListSQLReviewOriginSQLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserActiveTenantWithOptions(request *GetUserActiveTenantRequest, runtime *util.RuntimeOptions) (_result *GetUserActiveTenantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUserActiveTenantResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUserActiveTenant"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserActiveTenant(request *GetUserActiveTenantRequest) (_result *GetUserActiveTenantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserActiveTenantResponse{}
	_body, _err := client.GetUserActiveTenantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterUserWithOptions(request *RegisterUserRequest, runtime *util.RuntimeOptions) (_result *RegisterUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterUser"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterUser(request *RegisterUserRequest) (_result *RegisterUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterUserResponse{}
	_body, _err := client.RegisterUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetInstance"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstance(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPermApplyOrderDetailWithOptions(request *GetPermApplyOrderDetailRequest, runtime *util.RuntimeOptions) (_result *GetPermApplyOrderDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPermApplyOrderDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPermApplyOrderDetail"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPermApplyOrderDetail(request *GetPermApplyOrderDetailRequest) (_result *GetPermApplyOrderDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPermApplyOrderDetailResponse{}
	_body, _err := client.GetPermApplyOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIndexesWithOptions(request *ListIndexesRequest, runtime *util.RuntimeOptions) (_result *ListIndexesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListIndexesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListIndexes"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIndexes(request *ListIndexesRequest) (_result *ListIndexesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIndexesResponse{}
	_body, _err := client.ListIndexesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLogicTablesWithOptions(request *ListLogicTablesRequest, runtime *util.RuntimeOptions) (_result *ListLogicTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListLogicTablesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListLogicTables"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLogicTables(request *ListLogicTablesRequest) (_result *ListLogicTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLogicTablesResponse{}
	_body, _err := client.ListLogicTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTableTopologyWithOptions(request *GetTableTopologyRequest, runtime *util.RuntimeOptions) (_result *GetTableTopologyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTableTopologyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTableTopology"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTableTopology(request *GetTableTopologyRequest) (_result *GetTableTopologyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTableTopologyResponse{}
	_body, _err := client.GetTableTopologyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataExportDownloadURLWithOptions(request *GetDataExportDownloadURLRequest, runtime *util.RuntimeOptions) (_result *GetDataExportDownloadURLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDataExportDownloadURLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDataExportDownloadURL"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataExportDownloadURL(request *GetDataExportDownloadURLRequest) (_result *GetDataExportDownloadURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDataExportDownloadURLResponse{}
	_body, _err := client.GetDataExportDownloadURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDataCronClearOrderWithOptions(tmpReq *CreateDataCronClearOrderRequest, runtime *util.RuntimeOptions) (_result *CreateDataCronClearOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDataCronClearOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedUserList)) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, tea.String("RelatedUserList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Param))) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Param), tea.String("Param"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDataCronClearOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDataCronClearOrder"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataCronClearOrder(request *CreateDataCronClearOrderRequest) (_result *CreateDataCronClearOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataCronClearOrderResponse{}
	_body, _err := client.CreateDataCronClearOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePublishGroupTaskWithOptions(request *CreatePublishGroupTaskRequest, runtime *util.RuntimeOptions) (_result *CreatePublishGroupTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreatePublishGroupTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreatePublishGroupTask"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePublishGroupTask(request *CreatePublishGroupTaskRequest) (_result *CreatePublishGroupTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePublishGroupTaskResponse{}
	_body, _err := client.CreatePublishGroupTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDatabaseWithOptions(request *GetDatabaseRequest, runtime *util.RuntimeOptions) (_result *GetDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDatabaseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDatabase"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDatabase(request *GetDatabaseRequest) (_result *GetDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDatabaseResponse{}
	_body, _err := client.GetDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOwnerApplyOrderDetailWithOptions(request *GetOwnerApplyOrderDetailRequest, runtime *util.RuntimeOptions) (_result *GetOwnerApplyOrderDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOwnerApplyOrderDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOwnerApplyOrderDetail"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOwnerApplyOrderDetail(request *GetOwnerApplyOrderDetailRequest) (_result *GetOwnerApplyOrderDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOwnerApplyOrderDetailResponse{}
	_body, _err := client.GetOwnerApplyOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOpLogWithOptions(request *GetOpLogRequest, runtime *util.RuntimeOptions) (_result *GetOpLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOpLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOpLog"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOpLog(request *GetOpLogRequest) (_result *GetOpLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOpLogResponse{}
	_body, _err := client.GetOpLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTableWithOptions(request *SearchTableRequest, runtime *util.RuntimeOptions) (_result *SearchTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchTableResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchTable"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTable(request *SearchTableRequest) (_result *SearchTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTableResponse{}
	_body, _err := client.SearchTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDBTaskSQLJobWithOptions(request *ListDBTaskSQLJobRequest, runtime *util.RuntimeOptions) (_result *ListDBTaskSQLJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDBTaskSQLJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDBTaskSQLJob"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDBTaskSQLJob(request *ListDBTaskSQLJobRequest) (_result *ListDBTaskSQLJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDBTaskSQLJobResponse{}
	_body, _err := client.ListDBTaskSQLJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUser"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataCronClearTaskDetailListWithOptions(request *GetDataCronClearTaskDetailListRequest, runtime *util.RuntimeOptions) (_result *GetDataCronClearTaskDetailListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDataCronClearTaskDetailListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDataCronClearTaskDetailList"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataCronClearTaskDetailList(request *GetDataCronClearTaskDetailListRequest) (_result *GetDataCronClearTaskDetailListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDataCronClearTaskDetailListResponse{}
	_body, _err := client.GetDataCronClearTaskDetailListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStructSyncJobAnalyzeResultWithOptions(request *GetStructSyncJobAnalyzeResultRequest, runtime *util.RuntimeOptions) (_result *GetStructSyncJobAnalyzeResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetStructSyncJobAnalyzeResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetStructSyncJobAnalyzeResult"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStructSyncJobAnalyzeResult(request *GetStructSyncJobAnalyzeResultRequest) (_result *GetStructSyncJobAnalyzeResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStructSyncJobAnalyzeResultResponse{}
	_body, _err := client.GetStructSyncJobAnalyzeResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApproveOrderWithOptions(request *ApproveOrderRequest, runtime *util.RuntimeOptions) (_result *ApproveOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApproveOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApproveOrder"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApproveOrder(request *ApproveOrderRequest) (_result *ApproveOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveOrderResponse{}
	_body, _err := client.ApproveOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataCorrectTaskDetailWithOptions(request *GetDataCorrectTaskDetailRequest, runtime *util.RuntimeOptions) (_result *GetDataCorrectTaskDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDataCorrectTaskDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDataCorrectTaskDetail"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataCorrectTaskDetail(request *GetDataCorrectTaskDetailRequest) (_result *GetDataCorrectTaskDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDataCorrectTaskDetailResponse{}
	_body, _err := client.GetDataCorrectTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUploadFileJobWithOptions(request *CreateUploadFileJobRequest, runtime *util.RuntimeOptions) (_result *CreateUploadFileJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUploadFileJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUploadFileJob"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUploadFileJob(request *CreateUploadFileJobRequest) (_result *CreateUploadFileJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUploadFileJobResponse{}
	_body, _err := client.CreateUploadFileJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLogicDatabasesWithOptions(request *ListLogicDatabasesRequest, runtime *util.RuntimeOptions) (_result *ListLogicDatabasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListLogicDatabasesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListLogicDatabases"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLogicDatabases(request *ListLogicDatabasesRequest) (_result *ListLogicDatabasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLogicDatabasesResponse{}
	_body, _err := client.ListLogicDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDataImportOrderWithOptions(tmpReq *CreateDataImportOrderRequest, runtime *util.RuntimeOptions) (_result *CreateDataImportOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDataImportOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedUserList)) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, tea.String("RelatedUserList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Param))) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Param), tea.String("Param"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDataImportOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDataImportOrder"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataImportOrder(request *CreateDataImportOrderRequest) (_result *CreateDataImportOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataImportOrderResponse{}
	_body, _err := client.CreateDataImportOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseOrderWithOptions(request *CloseOrderRequest, runtime *util.RuntimeOptions) (_result *CloseOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CloseOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CloseOrder"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseOrder(request *CloseOrderRequest) (_result *CloseOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseOrderResponse{}
	_body, _err := client.CloseOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncInstanceMetaWithOptions(request *SyncInstanceMetaRequest, runtime *util.RuntimeOptions) (_result *SyncInstanceMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SyncInstanceMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SyncInstanceMeta"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncInstanceMeta(request *SyncInstanceMetaRequest) (_result *SyncInstanceMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncInstanceMetaResponse{}
	_body, _err := client.SyncInstanceMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrdersWithOptions(request *ListOrdersRequest, runtime *util.RuntimeOptions) (_result *ListOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOrdersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOrders"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrders(request *ListOrdersRequest) (_result *ListOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrdersResponse{}
	_body, _err := client.ListOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSQLReviewOrderWithOptions(tmpReq *CreateSQLReviewOrderRequest, runtime *util.RuntimeOptions) (_result *CreateSQLReviewOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSQLReviewOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedUserList)) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, tea.String("RelatedUserList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Param))) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Param), tea.String("Param"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSQLReviewOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSQLReviewOrder"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSQLReviewOrder(request *CreateSQLReviewOrderRequest) (_result *CreateSQLReviewOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSQLReviewOrderResponse{}
	_body, _err := client.CreateSQLReviewOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrderBaseInfoWithOptions(request *GetOrderBaseInfoRequest, runtime *util.RuntimeOptions) (_result *GetOrderBaseInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOrderBaseInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOrderBaseInfo"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrderBaseInfo(request *GetOrderBaseInfoRequest) (_result *GetOrderBaseInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOrderBaseInfoResponse{}
	_body, _err := client.GetOrderBaseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSQLReviewOptimizeDetailWithOptions(request *GetSQLReviewOptimizeDetailRequest, runtime *util.RuntimeOptions) (_result *GetSQLReviewOptimizeDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSQLReviewOptimizeDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSQLReviewOptimizeDetail"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSQLReviewOptimizeDetail(request *GetSQLReviewOptimizeDetailRequest) (_result *GetSQLReviewOptimizeDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSQLReviewOptimizeDetailResponse{}
	_body, _err := client.GetSQLReviewOptimizeDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserTenantsWithOptions(request *ListUserTenantsRequest, runtime *util.RuntimeOptions) (_result *ListUserTenantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUserTenantsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserTenants"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserTenants(request *ListUserTenantsRequest) (_result *ListUserTenantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserTenantsResponse{}
	_body, _err := client.ListUserTenantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetOwnersWithOptions(request *SetOwnersRequest, runtime *util.RuntimeOptions) (_result *SetOwnersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetOwnersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetOwners"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetOwners(request *SetOwnersRequest) (_result *SetOwnersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetOwnersResponse{}
	_body, _err := client.SetOwnersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDataCorrectOrderWithOptions(tmpReq *CreateDataCorrectOrderRequest, runtime *util.RuntimeOptions) (_result *CreateDataCorrectOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDataCorrectOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedUserList)) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, tea.String("RelatedUserList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Param))) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Param), tea.String("Param"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDataCorrectOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDataCorrectOrder"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataCorrectOrder(request *CreateDataCorrectOrderRequest) (_result *CreateDataCorrectOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataCorrectOrderResponse{}
	_body, _err := client.CreateDataCorrectOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLogicDatabaseWithOptions(request *GetLogicDatabaseRequest, runtime *util.RuntimeOptions) (_result *GetLogicDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLogicDatabaseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLogicDatabase"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogicDatabase(request *GetLogicDatabaseRequest) (_result *GetLogicDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLogicDatabaseResponse{}
	_body, _err := client.GetLogicDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataCorrectBackupFilesWithOptions(tmpReq *GetDataCorrectBackupFilesRequest, runtime *util.RuntimeOptions) (_result *GetDataCorrectBackupFilesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetDataCorrectBackupFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ActionDetail)) {
		request.ActionDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ActionDetail, tea.String("ActionDetail"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDataCorrectBackupFilesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDataCorrectBackupFiles"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataCorrectBackupFiles(request *GetDataCorrectBackupFilesRequest) (_result *GetDataCorrectBackupFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDataCorrectBackupFilesResponse{}
	_body, _err := client.GetDataCorrectBackupFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataCorrectPreCheckSQLWithOptions(request *ListDataCorrectPreCheckSQLRequest, runtime *util.RuntimeOptions) (_result *ListDataCorrectPreCheckSQLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDataCorrectPreCheckSQLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDataCorrectPreCheckSQL"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataCorrectPreCheckSQL(request *ListDataCorrectPreCheckSQLRequest) (_result *ListDataCorrectPreCheckSQLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataCorrectPreCheckSQLResponse{}
	_body, _err := client.ListDataCorrectPreCheckSQLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterInstanceWithOptions(request *RegisterInstanceRequest, runtime *util.RuntimeOptions) (_result *RegisterInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterInstance"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterInstance(request *RegisterInstanceRequest) (_result *RegisterInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterInstanceResponse{}
	_body, _err := client.RegisterInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStructSyncOrderWithOptions(tmpReq *CreateStructSyncOrderRequest, runtime *util.RuntimeOptions) (_result *CreateStructSyncOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateStructSyncOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RelatedUserList)) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, tea.String("RelatedUserList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Param))) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Param), tea.String("Param"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateStructSyncOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateStructSyncOrder"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStructSyncOrder(request *CreateStructSyncOrderRequest) (_result *CreateStructSyncOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStructSyncOrderResponse{}
	_body, _err := client.CreateStructSyncOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteDataExportWithOptions(tmpReq *ExecuteDataExportRequest, runtime *util.RuntimeOptions) (_result *ExecuteDataExportResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExecuteDataExportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ActionDetail)) {
		request.ActionDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ActionDetail, tea.String("ActionDetail"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteDataExportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecuteDataExport"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteDataExport(request *ExecuteDataExportRequest) (_result *ExecuteDataExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteDataExportResponse{}
	_body, _err := client.ExecuteDataExportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteDataCorrectWithOptions(tmpReq *ExecuteDataCorrectRequest, runtime *util.RuntimeOptions) (_result *ExecuteDataCorrectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExecuteDataCorrectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ActionDetail)) {
		request.ActionDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ActionDetail, tea.String("ActionDetail"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteDataCorrectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecuteDataCorrect"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteDataCorrect(request *ExecuteDataCorrectRequest) (_result *ExecuteDataCorrectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteDataCorrectResponse{}
	_body, _err := client.ExecuteDataCorrectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataCorrectPreCheckDBWithOptions(request *ListDataCorrectPreCheckDBRequest, runtime *util.RuntimeOptions) (_result *ListDataCorrectPreCheckDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDataCorrectPreCheckDBResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDataCorrectPreCheckDB"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataCorrectPreCheckDB(request *ListDataCorrectPreCheckDBRequest) (_result *ListDataCorrectPreCheckDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataCorrectPreCheckDBResponse{}
	_body, _err := client.ListDataCorrectPreCheckDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTablesWithOptions(request *ListTablesRequest, runtime *util.RuntimeOptions) (_result *ListTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTablesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTables"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTables(request *ListTablesRequest) (_result *ListTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTablesResponse{}
	_body, _err := client.ListTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkFlowNodesWithOptions(request *ListWorkFlowNodesRequest, runtime *util.RuntimeOptions) (_result *ListWorkFlowNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListWorkFlowNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListWorkFlowNodes"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkFlowNodes(request *ListWorkFlowNodesRequest) (_result *ListWorkFlowNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWorkFlowNodesResponse{}
	_body, _err := client.ListWorkFlowNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStructSyncOrderDetailWithOptions(request *GetStructSyncOrderDetailRequest, runtime *util.RuntimeOptions) (_result *GetStructSyncOrderDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetStructSyncOrderDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetStructSyncOrderDetail"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStructSyncOrderDetail(request *GetStructSyncOrderDetailRequest) (_result *GetStructSyncOrderDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStructSyncOrderDetailResponse{}
	_body, _err := client.GetStructSyncOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSensitiveColumnsDetailWithOptions(request *ListSensitiveColumnsDetailRequest, runtime *util.RuntimeOptions) (_result *ListSensitiveColumnsDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSensitiveColumnsDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSensitiveColumnsDetail"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSensitiveColumnsDetail(request *ListSensitiveColumnsDetailRequest) (_result *ListSensitiveColumnsDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSensitiveColumnsDetailResponse{}
	_body, _err := client.ListSensitiveColumnsDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateUser"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPhysicalDatabaseWithOptions(request *GetPhysicalDatabaseRequest, runtime *util.RuntimeOptions) (_result *GetPhysicalDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPhysicalDatabaseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPhysicalDatabase"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPhysicalDatabase(request *GetPhysicalDatabaseRequest) (_result *GetPhysicalDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhysicalDatabaseResponse{}
	_body, _err := client.GetPhysicalDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStructSyncExecSqlDetailWithOptions(request *GetStructSyncExecSqlDetailRequest, runtime *util.RuntimeOptions) (_result *GetStructSyncExecSqlDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetStructSyncExecSqlDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetStructSyncExecSqlDetail"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStructSyncExecSqlDetail(request *GetStructSyncExecSqlDetailRequest) (_result *GetStructSyncExecSqlDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStructSyncExecSqlDetailResponse{}
	_body, _err := client.GetStructSyncExecSqlDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteInstance"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTableDBTopologyWithOptions(request *GetTableDBTopologyRequest, runtime *util.RuntimeOptions) (_result *GetTableDBTopologyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTableDBTopologyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTableDBTopology"), tea.String("2018-11-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTableDBTopology(request *GetTableDBTopologyRequest) (_result *GetTableDBTopologyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTableDBTopologyResponse{}
	_body, _err := client.GetTableDBTopologyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

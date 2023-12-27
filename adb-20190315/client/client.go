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

type AllocateClusterPublicConnectionRequest struct {
	// The prefix of the public endpoint.
	//
	// *   The prefix must contain lowercase letters, digits, and hyphens (-). It must start with a lowercase letter.
	// *   The prefix can be up to 30 characters in length.
	// *   By default, the cluster name is used as the prefix of the public endpoint.
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// You can call the [DescribeDBClusters](~~129857~~) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocateClusterPublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateClusterPublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateClusterPublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateClusterPublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetDBClusterId(v string) *AllocateClusterPublicConnectionRequest {
	s.DBClusterId = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetOwnerAccount(v string) *AllocateClusterPublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetOwnerId(v int64) *AllocateClusterPublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetResourceOwnerAccount(v string) *AllocateClusterPublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetResourceOwnerId(v int64) *AllocateClusterPublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

type AllocateClusterPublicConnectionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateClusterPublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocateClusterPublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateClusterPublicConnectionResponseBody) SetRequestId(v string) *AllocateClusterPublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type AllocateClusterPublicConnectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AllocateClusterPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AllocateClusterPublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocateClusterPublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocateClusterPublicConnectionResponse) SetHeaders(v map[string]*string) *AllocateClusterPublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocateClusterPublicConnectionResponse) SetStatusCode(v int32) *AllocateClusterPublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateClusterPublicConnectionResponse) SetBody(v *AllocateClusterPublicConnectionResponseBody) *AllocateClusterPublicConnectionResponse {
	s.Body = v
	return s
}

type ApplyAdviceByIdRequest struct {
	// The date when the suggestion is generated. Specify the date in the yyyyMMdd format. The date must be in UTC.
	AdviceDate *int64 `json:"AdviceDate,omitempty" xml:"AdviceDate,omitempty"`
	// The suggestion ID.
	AdviceId *string `json:"AdviceId,omitempty" xml:"AdviceId,omitempty"`
	// The ID of the cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of Data Warehouse Edition (V3.0) clusters.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ApplyAdviceByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyAdviceByIdRequest) GoString() string {
	return s.String()
}

func (s *ApplyAdviceByIdRequest) SetAdviceDate(v int64) *ApplyAdviceByIdRequest {
	s.AdviceDate = &v
	return s
}

func (s *ApplyAdviceByIdRequest) SetAdviceId(v string) *ApplyAdviceByIdRequest {
	s.AdviceId = &v
	return s
}

func (s *ApplyAdviceByIdRequest) SetDBClusterId(v string) *ApplyAdviceByIdRequest {
	s.DBClusterId = &v
	return s
}

func (s *ApplyAdviceByIdRequest) SetRegionId(v string) *ApplyAdviceByIdRequest {
	s.RegionId = &v
	return s
}

type ApplyAdviceByIdResponseBody struct {
	// The message returned for the operation. Valid values:
	//
	// *   **SUCCESS** is returned if the operation is successful.
	// *   An error message is returned if the operation fails.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyAdviceByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyAdviceByIdResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyAdviceByIdResponseBody) SetMessage(v string) *ApplyAdviceByIdResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyAdviceByIdResponseBody) SetRequestId(v string) *ApplyAdviceByIdResponseBody {
	s.RequestId = &v
	return s
}

type ApplyAdviceByIdResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyAdviceByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyAdviceByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyAdviceByIdResponse) GoString() string {
	return s.String()
}

func (s *ApplyAdviceByIdResponse) SetHeaders(v map[string]*string) *ApplyAdviceByIdResponse {
	s.Headers = v
	return s
}

func (s *ApplyAdviceByIdResponse) SetStatusCode(v int32) *ApplyAdviceByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyAdviceByIdResponse) SetBody(v *ApplyAdviceByIdResponseBody) *ApplyAdviceByIdResponse {
	s.Body = v
	return s
}

type AttachUserENIRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query cluster IDs.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AttachUserENIRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachUserENIRequest) GoString() string {
	return s.String()
}

func (s *AttachUserENIRequest) SetDBClusterId(v string) *AttachUserENIRequest {
	s.DBClusterId = &v
	return s
}

func (s *AttachUserENIRequest) SetOwnerAccount(v string) *AttachUserENIRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AttachUserENIRequest) SetOwnerId(v int64) *AttachUserENIRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachUserENIRequest) SetResourceOwnerAccount(v string) *AttachUserENIRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachUserENIRequest) SetResourceOwnerId(v int64) *AttachUserENIRequest {
	s.ResourceOwnerId = &v
	return s
}

type AttachUserENIResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachUserENIResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachUserENIResponseBody) GoString() string {
	return s.String()
}

func (s *AttachUserENIResponseBody) SetRequestId(v string) *AttachUserENIResponseBody {
	s.RequestId = &v
	return s
}

type AttachUserENIResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachUserENIResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachUserENIResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachUserENIResponse) GoString() string {
	return s.String()
}

func (s *AttachUserENIResponse) SetHeaders(v map[string]*string) *AttachUserENIResponse {
	s.Headers = v
	return s
}

func (s *AttachUserENIResponse) SetStatusCode(v int32) *AttachUserENIResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachUserENIResponse) SetBody(v *AttachUserENIResponseBody) *AttachUserENIResponse {
	s.Body = v
	return s
}

type BatchApplyAdviceByIdListRequest struct {
	AdviceDate   *int64  `json:"AdviceDate,omitempty" xml:"AdviceDate,omitempty"`
	AdviceIdList *string `json:"AdviceIdList,omitempty" xml:"AdviceIdList,omitempty"`
	// The message returned for the operation. Valid values:
	//
	// *   **SUCCESS** is returned if the operation is successful.
	// *   An error message is returned if the operation fails.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the request.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BatchApplyAdviceByIdListRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchApplyAdviceByIdListRequest) GoString() string {
	return s.String()
}

func (s *BatchApplyAdviceByIdListRequest) SetAdviceDate(v int64) *BatchApplyAdviceByIdListRequest {
	s.AdviceDate = &v
	return s
}

func (s *BatchApplyAdviceByIdListRequest) SetAdviceIdList(v string) *BatchApplyAdviceByIdListRequest {
	s.AdviceIdList = &v
	return s
}

func (s *BatchApplyAdviceByIdListRequest) SetDBClusterId(v string) *BatchApplyAdviceByIdListRequest {
	s.DBClusterId = &v
	return s
}

func (s *BatchApplyAdviceByIdListRequest) SetRegionId(v string) *BatchApplyAdviceByIdListRequest {
	s.RegionId = &v
	return s
}

type BatchApplyAdviceByIdListResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchApplyAdviceByIdListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchApplyAdviceByIdListResponseBody) GoString() string {
	return s.String()
}

func (s *BatchApplyAdviceByIdListResponseBody) SetMessage(v string) *BatchApplyAdviceByIdListResponseBody {
	s.Message = &v
	return s
}

func (s *BatchApplyAdviceByIdListResponseBody) SetRequestId(v string) *BatchApplyAdviceByIdListResponseBody {
	s.RequestId = &v
	return s
}

type BatchApplyAdviceByIdListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchApplyAdviceByIdListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchApplyAdviceByIdListResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchApplyAdviceByIdListResponse) GoString() string {
	return s.String()
}

func (s *BatchApplyAdviceByIdListResponse) SetHeaders(v map[string]*string) *BatchApplyAdviceByIdListResponse {
	s.Headers = v
	return s
}

func (s *BatchApplyAdviceByIdListResponse) SetStatusCode(v int32) *BatchApplyAdviceByIdListResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchApplyAdviceByIdListResponse) SetBody(v *BatchApplyAdviceByIdListResponseBody) *BatchApplyAdviceByIdListResponse {
	s.Body = v
	return s
}

type BindDBResourceGroupWithUserRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The database account with which to associate the resource group. It can be a standard account or a privileged account.
	GroupUser            *string `json:"GroupUser,omitempty" xml:"GroupUser,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BindDBResourceGroupWithUserRequest) String() string {
	return tea.Prettify(s)
}

func (s BindDBResourceGroupWithUserRequest) GoString() string {
	return s.String()
}

func (s *BindDBResourceGroupWithUserRequest) SetDBClusterId(v string) *BindDBResourceGroupWithUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetGroupName(v string) *BindDBResourceGroupWithUserRequest {
	s.GroupName = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetGroupUser(v string) *BindDBResourceGroupWithUserRequest {
	s.GroupUser = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetOwnerAccount(v string) *BindDBResourceGroupWithUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetOwnerId(v int64) *BindDBResourceGroupWithUserRequest {
	s.OwnerId = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetResourceOwnerAccount(v string) *BindDBResourceGroupWithUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetResourceOwnerId(v int64) *BindDBResourceGroupWithUserRequest {
	s.ResourceOwnerId = &v
	return s
}

type BindDBResourceGroupWithUserResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindDBResourceGroupWithUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindDBResourceGroupWithUserResponseBody) GoString() string {
	return s.String()
}

func (s *BindDBResourceGroupWithUserResponseBody) SetRequestId(v string) *BindDBResourceGroupWithUserResponseBody {
	s.RequestId = &v
	return s
}

type BindDBResourceGroupWithUserResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindDBResourceGroupWithUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindDBResourceGroupWithUserResponse) String() string {
	return tea.Prettify(s)
}

func (s BindDBResourceGroupWithUserResponse) GoString() string {
	return s.String()
}

func (s *BindDBResourceGroupWithUserResponse) SetHeaders(v map[string]*string) *BindDBResourceGroupWithUserResponse {
	s.Headers = v
	return s
}

func (s *BindDBResourceGroupWithUserResponse) SetStatusCode(v int32) *BindDBResourceGroupWithUserResponse {
	s.StatusCode = &v
	return s
}

func (s *BindDBResourceGroupWithUserResponse) SetBody(v *BindDBResourceGroupWithUserResponseBody) *BindDBResourceGroupWithUserResponse {
	s.Body = v
	return s
}

type BindDBResourcePoolWithUserRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the resource group.
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The database account with which to associate the resource group. It can be a standard account or a privileged account.
	PoolUser             *string `json:"PoolUser,omitempty" xml:"PoolUser,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BindDBResourcePoolWithUserRequest) String() string {
	return tea.Prettify(s)
}

func (s BindDBResourcePoolWithUserRequest) GoString() string {
	return s.String()
}

func (s *BindDBResourcePoolWithUserRequest) SetDBClusterId(v string) *BindDBResourcePoolWithUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetOwnerAccount(v string) *BindDBResourcePoolWithUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetOwnerId(v int64) *BindDBResourcePoolWithUserRequest {
	s.OwnerId = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetPoolName(v string) *BindDBResourcePoolWithUserRequest {
	s.PoolName = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetPoolUser(v string) *BindDBResourcePoolWithUserRequest {
	s.PoolUser = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetResourceOwnerAccount(v string) *BindDBResourcePoolWithUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetResourceOwnerId(v int64) *BindDBResourcePoolWithUserRequest {
	s.ResourceOwnerId = &v
	return s
}

type BindDBResourcePoolWithUserResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindDBResourcePoolWithUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindDBResourcePoolWithUserResponseBody) GoString() string {
	return s.String()
}

func (s *BindDBResourcePoolWithUserResponseBody) SetRequestId(v string) *BindDBResourcePoolWithUserResponseBody {
	s.RequestId = &v
	return s
}

type BindDBResourcePoolWithUserResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindDBResourcePoolWithUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindDBResourcePoolWithUserResponse) String() string {
	return tea.Prettify(s)
}

func (s BindDBResourcePoolWithUserResponse) GoString() string {
	return s.String()
}

func (s *BindDBResourcePoolWithUserResponse) SetHeaders(v map[string]*string) *BindDBResourcePoolWithUserResponse {
	s.Headers = v
	return s
}

func (s *BindDBResourcePoolWithUserResponse) SetStatusCode(v int32) *BindDBResourcePoolWithUserResponse {
	s.StatusCode = &v
	return s
}

func (s *BindDBResourcePoolWithUserResponse) SetBody(v *BindDBResourcePoolWithUserResponseBody) *BindDBResourcePoolWithUserResponse {
	s.Body = v
	return s
}

type CreateAccountRequest struct {
	// The description of the database account.
	//
	// *   The description cannot start with `http://` or `https://`.
	// *   The description can be up to 256 characters in length.
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	//
	// *   The name must start with a lowercase letter and end with a lowercase letter or a digit.
	// *   The name can contain lowercase letters, digits, and underscores (\_).
	// *   The name must be 2 to 16 characters in length.
	// *   Reserved account names such as root, admin, and opsadmin cannot be used.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the database account.
	//
	// *   The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	// *   Special characters include `! @ # $ % ^ & * ( ) _ + - =`
	// *   The password must be 8 to 32 characters in length.
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The type of the database account. Valid values:
	//
	// *   **Normal**: standard account. Up to 256 standard accounts can be created for a cluster.
	// *   **Super** (default): privileged account. Only a single privileged account can be created for a cluster.
	//
	// >  If a cluster does not have accounts, you can specify this parameter to create a privileged account or standard account. If a cluster has a privileged account, you must set this parameter to Normal to create a standard account. Otherwise, the operation fails. After an account is created, the privileged account has permissions on all databases of the cluster. The standard account does not have permissions and must be granted permissions on specific databases by the privileged account. For more information, see GRANT.
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to view cluster IDs.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetAccountType(v string) *CreateAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountRequest) SetDBClusterId(v string) *CreateAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerAccount(v string) *CreateAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerId(v int64) *CreateAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerAccount(v string) *CreateAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerId(v int64) *CreateAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateAccountResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) SetDBClusterId(v string) *CreateAccountResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountResponseBody) SetTaskId(v int32) *CreateAccountResponseBody {
	s.TaskId = &v
	return s
}

type CreateAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountResponse) SetHeaders(v map[string]*string) *CreateAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateAccountResponse) SetStatusCode(v int32) *CreateAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccountResponse) SetBody(v *CreateAccountResponseBody) *CreateAccountResponse {
	s.Body = v
	return s
}

type CreateDBClusterRequest struct {
	// A reserved parameter.
	BackupSetID *string `json:"BackupSetID,omitempty" xml:"BackupSetID,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The value is case-sensitive and can contain a maximum of 64 ASCII characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The computing resources of the cluster. This parameter is required if the Mode parameter is set to **Flexible**.
	//
	// >  You can call the [DescribeAvailableResource](~~190632~~) operation to query the computing resources that are available within a specific region.
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// *   **Cluster**: reserved mode for Cluster Edition
	//
	// <!---->
	//
	// *   **MixedStorage**: elastic mode for Cluster Edition
	//
	// >  If the DBClusterCategory parameter is set to Cluster, you must set the Mode parameter to Reserver. If the DBClusterCategory parameter is set to MixedStorage, you must set the Mode parameter to Flexible. Otherwise, the cluster fails to be created.
	DBClusterCategory *string `json:"DBClusterCategory,omitempty" xml:"DBClusterCategory,omitempty"`
	// The specification of the cluster. Valid values:
	//
	// *   **C8**
	// *   **C32**
	//
	// >  This parameter is required if the Mode parameter is set to Reserver.
	DBClusterClass *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	// The description of the cluster.
	//
	// *   The description cannot start with `http://` or `https`.
	// *   The description must be 2 to 256 characters in length.
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The network type of the cluster. Set the value to **VPC**.
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The version of the cluster. Set the value to **3.0**.
	DBClusterVersion *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	// The number of node groups. Valid values: 1 to 200 (integer).
	//
	// >  This parameter is required if the Mode parameter is set to Reserver.
	DBNodeGroupCount *string `json:"DBNodeGroupCount,omitempty" xml:"DBNodeGroupCount,omitempty"`
	// The storage capacity of the cluster. Unit: GB.
	//
	// *   Valid values when DBClusterClass is set to C8: 100 to 1000
	// *   Valid values when DBClusterClass is set to C32: 100 to 8000
	//
	// > * This parameter is required if the Mode parameter is set to Reserver.
	// > * 1000 The storage capacity less than 1,000 GB increases in 100 GB increments. The storage capacity greater than 1,000 GB increases in 1,000 GB increments.
	DBNodeStorage *string `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	// Specifies whether to enable disk encryption.
	//
	// Valid values:
	//
	// *   true
	// *   false
	DiskEncryption *string `json:"DiskEncryption,omitempty" xml:"DiskEncryption,omitempty"`
	// The number of elastic I/O units (EIUs). For more information, see [Use EIUs to scale up storage resources](~~189505~~).
	ElasticIOResource *string `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	// A reserved parameter.
	ExecutorCount *string `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	// The Key Management Service (KMS) ID that is used for disk encryption. This parameter is valid only when DiskEncryption is set to true.
	KmsId *string `json:"KmsId,omitempty" xml:"KmsId,omitempty"`
	// The mode of the cluster. Valid values:
	//
	// *   **Reserver**: the reserved mode.
	// *   **Flexible**: the elastic mode.
	Mode         *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// *   **Postpaid**: pay-as-you-go
	// *   **Prepaid**: subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription type of the subscription cluster. Valid values:
	//
	// *   **Year**: subscription on a yearly basis
	// *   **Month**: subscription on a monthly basis
	//
	// >  This parameter is required if the PayType parameter is set to Prepaid.
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the cluster belongs.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// A reserved parameter.
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// A reserved parameter.
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// A reserved parameter.
	SourceDBInstanceName *string `json:"SourceDBInstanceName,omitempty" xml:"SourceDBInstanceName,omitempty"`
	// A reserved parameter.
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	// A reserved parameter.
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The tags to add to the cluster.
	Tag []*CreateDBClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The subscription period of the subscription cluster.
	//
	// *   Valid values when Period is set to Year: 1, 2, 3, and 5 (integer)
	// *   Valid values when Period is set to Month: 1 to 11 (integer)
	//
	// > * This parameter is required if the PayType parameter is set to Prepaid.
	// > * Longer subscription periods offer more savings. Purchasing a cluster for one year is more cost-effective than purchasing the cluster for 10 or 11 months.
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The virtual private cloud (VPC) ID of the cluster.
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the cluster.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent zone list.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterRequest) SetBackupSetID(v string) *CreateDBClusterRequest {
	s.BackupSetID = &v
	return s
}

func (s *CreateDBClusterRequest) SetClientToken(v string) *CreateDBClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBClusterRequest) SetComputeResource(v string) *CreateDBClusterRequest {
	s.ComputeResource = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterCategory(v string) *CreateDBClusterRequest {
	s.DBClusterCategory = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterClass(v string) *CreateDBClusterRequest {
	s.DBClusterClass = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterDescription(v string) *CreateDBClusterRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterNetworkType(v string) *CreateDBClusterRequest {
	s.DBClusterNetworkType = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterVersion(v string) *CreateDBClusterRequest {
	s.DBClusterVersion = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBNodeGroupCount(v string) *CreateDBClusterRequest {
	s.DBNodeGroupCount = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBNodeStorage(v string) *CreateDBClusterRequest {
	s.DBNodeStorage = &v
	return s
}

func (s *CreateDBClusterRequest) SetDiskEncryption(v string) *CreateDBClusterRequest {
	s.DiskEncryption = &v
	return s
}

func (s *CreateDBClusterRequest) SetElasticIOResource(v string) *CreateDBClusterRequest {
	s.ElasticIOResource = &v
	return s
}

func (s *CreateDBClusterRequest) SetExecutorCount(v string) *CreateDBClusterRequest {
	s.ExecutorCount = &v
	return s
}

func (s *CreateDBClusterRequest) SetKmsId(v string) *CreateDBClusterRequest {
	s.KmsId = &v
	return s
}

func (s *CreateDBClusterRequest) SetMode(v string) *CreateDBClusterRequest {
	s.Mode = &v
	return s
}

func (s *CreateDBClusterRequest) SetOwnerAccount(v string) *CreateDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBClusterRequest) SetOwnerId(v int64) *CreateDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBClusterRequest) SetPayType(v string) *CreateDBClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBClusterRequest) SetPeriod(v string) *CreateDBClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateDBClusterRequest) SetRegionId(v string) *CreateDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceGroupId(v string) *CreateDBClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceOwnerAccount(v string) *CreateDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceOwnerId(v int64) *CreateDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBClusterRequest) SetRestoreTime(v string) *CreateDBClusterRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateDBClusterRequest) SetRestoreType(v string) *CreateDBClusterRequest {
	s.RestoreType = &v
	return s
}

func (s *CreateDBClusterRequest) SetSourceDBInstanceName(v string) *CreateDBClusterRequest {
	s.SourceDBInstanceName = &v
	return s
}

func (s *CreateDBClusterRequest) SetStorageResource(v string) *CreateDBClusterRequest {
	s.StorageResource = &v
	return s
}

func (s *CreateDBClusterRequest) SetStorageType(v string) *CreateDBClusterRequest {
	s.StorageType = &v
	return s
}

func (s *CreateDBClusterRequest) SetTag(v []*CreateDBClusterRequestTag) *CreateDBClusterRequest {
	s.Tag = v
	return s
}

func (s *CreateDBClusterRequest) SetUsedTime(v string) *CreateDBClusterRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBClusterRequest) SetVPCId(v string) *CreateDBClusterRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBClusterRequest) SetVSwitchId(v string) *CreateDBClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBClusterRequest) SetZoneId(v string) *CreateDBClusterRequest {
	s.ZoneId = &v
	return s
}

type CreateDBClusterRequestTag struct {
	// The key of tag N to add to the cluster. You can use tags to filter clusters. Valid values of N: 1 to 20. The values that you specify for N must be unique and consecutive integers that start from 1. Each value of `Tag.N.Key` is paired with a value of `Tag.N.Value`.
	//
	// >  The tag key can be up to 64 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the cluster. You can use tags to filter clusters. Valid values of N: 1 to 20. The values that you specify for N must be unique and consecutive integers that start from 1. Each value of `Tag.N.Key` is paired with a value of `Tag.N.Value`.
	//
	// >  The tag value can be up to 64 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDBClusterRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDBClusterRequestTag) SetKey(v string) *CreateDBClusterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDBClusterRequestTag) SetValue(v string) *CreateDBClusterRequestTag {
	s.Value = &v
	return s
}

type CreateDBClusterResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the order.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the cluster belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponseBody) SetDBClusterId(v string) *CreateDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetOrderId(v string) *CreateDBClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetRequestId(v string) *CreateDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetResourceGroupId(v string) *CreateDBClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

type CreateDBClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponse) SetHeaders(v map[string]*string) *CreateDBClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateDBClusterResponse) SetStatusCode(v int32) *CreateDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBClusterResponse) SetBody(v *CreateDBClusterResponseBody) *CreateDBClusterResponse {
	s.Body = v
	return s
}

type CreateDBResourceGroupRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// *   The name can be up to 255 characters in length.
	// *   The name must start with a letter or a digit.
	// *   The name can contain letters, digits, hyphens (\_), and underscores (\_).
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The query execution mode. Default value: batch. Valid values:
	//
	// *   **interactive**
	// *   **batch**
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The number of nodes. Default value: 0.
	//
	// *   Each node is configured with the resources of 16 cores and 64 GB memory.
	// *   Make sure that the amount of resources of the nodes (Number of nodes × 16 cores and 64 GB memory) is less than or equal to the amount of unused resources of the cluster.
	NodeNum              *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDBResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupRequest) SetDBClusterId(v string) *CreateDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetGroupName(v string) *CreateDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetGroupType(v string) *CreateDBResourceGroupRequest {
	s.GroupType = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetNodeNum(v int32) *CreateDBResourceGroupRequest {
	s.NodeNum = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetOwnerAccount(v string) *CreateDBResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetOwnerId(v int64) *CreateDBResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetResourceOwnerAccount(v string) *CreateDBResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetResourceOwnerId(v int64) *CreateDBResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateDBResourceGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupResponseBody) SetRequestId(v string) *CreateDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateDBResourceGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDBResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupResponse) SetHeaders(v map[string]*string) *CreateDBResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDBResourceGroupResponse) SetStatusCode(v int32) *CreateDBResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBResourceGroupResponse) SetBody(v *CreateDBResourceGroupResponseBody) *CreateDBResourceGroupResponse {
	s.Body = v
	return s
}

type CreateDBResourcePoolRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The number of nodes. Default value: 0.
	//
	// *   Each node provides 16 cores and 64 GB memory.
	// *   The total amount of resources provided by the nodes (number of nodes × 16 cores, number of nodes × 64 GB memory) cannot exceed the total amount of resources in the cluster. Set this parameter to a proper value.
	NodeNum      *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the resource group.
	//
	// *   The name can be up to 255 characters in length.
	// *   The name must start with a letter or a digit.
	// *   The name can contain letters, digits, hyphens (\_), and underscores (\_).
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The mode in which to execute SQL statements.
	//
	// *   **batch**
	//
	// *   **interactive**
	//
	// > For more information, see [Query execution modes](~~189502~~).
	QueryType            *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDBResourcePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *CreateDBResourcePoolRequest) SetDBClusterId(v string) *CreateDBResourcePoolRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetNodeNum(v int32) *CreateDBResourcePoolRequest {
	s.NodeNum = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetOwnerAccount(v string) *CreateDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetOwnerId(v int64) *CreateDBResourcePoolRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetPoolName(v string) *CreateDBResourcePoolRequest {
	s.PoolName = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetQueryType(v string) *CreateDBResourcePoolRequest {
	s.QueryType = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetResourceOwnerAccount(v string) *CreateDBResourcePoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBResourcePoolRequest) SetResourceOwnerId(v int64) *CreateDBResourcePoolRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateDBResourcePoolResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBResourcePoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResourcePoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBResourcePoolResponseBody) SetRequestId(v string) *CreateDBResourcePoolResponseBody {
	s.RequestId = &v
	return s
}

type CreateDBResourcePoolResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDBResourcePoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBResourcePoolResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBResourcePoolResponse) GoString() string {
	return s.String()
}

func (s *CreateDBResourcePoolResponse) SetHeaders(v map[string]*string) *CreateDBResourcePoolResponse {
	s.Headers = v
	return s
}

func (s *CreateDBResourcePoolResponse) SetStatusCode(v int32) *CreateDBResourcePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBResourcePoolResponse) SetBody(v *CreateDBResourcePoolResponseBody) *CreateDBResourcePoolResponse {
	s.Body = v
	return s
}

type CreateElasticPlanRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether the scaling plan takes effect. Valid values:
	//
	// *   **true** (default)
	// *   **false**
	ElasticPlanEnable *bool `json:"ElasticPlanEnable,omitempty" xml:"ElasticPlanEnable,omitempty"`
	// The end date of the scaling plan. Specify the date in the yyyy-MM-dd format.
	ElasticPlanEndDay        *string `json:"ElasticPlanEndDay,omitempty" xml:"ElasticPlanEndDay,omitempty"`
	ElasticPlanMonthlyRepeat *string `json:"ElasticPlanMonthlyRepeat,omitempty" xml:"ElasticPlanMonthlyRepeat,omitempty"`
	// The name of the scaling plan.
	//
	// *   The name must be 2 to 30 characters in length.
	// *   The name can contain letters, digits, and underscores (\_).
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// The number of nodes that are involved in the scaling plan.
	//
	// *   If ElasticPlanType is set to **worker**, you can set this parameter to 0 or leave this parameter empty.
	// *   If ElasticPlanType is set to **executorcombineworker** or **executor**, you must set this parameter to a value that is greater than 0.
	ElasticPlanNodeNum *int32 `json:"ElasticPlanNodeNum,omitempty" xml:"ElasticPlanNodeNum,omitempty"`
	// The start date of the scaling plan. Specify the date in the yyyy-MM-dd format.
	ElasticPlanStartDay *string `json:"ElasticPlanStartDay,omitempty" xml:"ElasticPlanStartDay,omitempty"`
	// The restoration time of the scaling plan. Specify the time on the hour in the HH:mm:ss format. The interval between the scale-up time and the restoration time cannot be more than 24 hours.
	ElasticPlanTimeEnd *string `json:"ElasticPlanTimeEnd,omitempty" xml:"ElasticPlanTimeEnd,omitempty"`
	// The scale-up time of the scaling plan. Specify the time on the hour in the HH:mm:ss format.
	ElasticPlanTimeStart *string `json:"ElasticPlanTimeStart,omitempty" xml:"ElasticPlanTimeStart,omitempty"`
	// The type of the scaling plan. Valid values:
	//
	// *   **worker**: scales only elastic I/O resources.
	// *   **executor**: scales only computing resources.
	// *   **executorcombineworker** (default): scales both elastic I/O resources and computing resources by proportion.
	// > - If you want to set this parameter to **executorcombineworker**, make sure that the cluster runs a minor version of 3.1.3.2 or later.
	// > - If you want to set this parameter to **worker** or **executor**, make sure that the cluster runs a minor version of 3.1.6.1 or later and a ticket is submitted. After your request is approved, you can set this parameter to **worker** or **executor**.
	ElasticPlanType *string `json:"ElasticPlanType,omitempty" xml:"ElasticPlanType,omitempty"`
	// The days of the week when you want to execute the scaling plan. Valid values: 0 to 6, which indicates Sunday to Saturday. Separate multiple values with commas (,).
	ElasticPlanWeeklyRepeat *string `json:"ElasticPlanWeeklyRepeat,omitempty" xml:"ElasticPlanWeeklyRepeat,omitempty"`
	// The resource specifications that can be scaled up by the scaling plan. Valid values:
	//
	// *   8 Core 64 GB (default)
	// *   16 Core 64 GB
	// *   32 Core 64 GB
	// *   64 Core 128 GB
	// *   12 Core 96 GB
	// *   24 Core 96 GB
	// *   52 Core 86 GB
	ElasticPlanWorkerSpec *string `json:"ElasticPlanWorkerSpec,omitempty" xml:"ElasticPlanWorkerSpec,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the resource group.
	//
	// > You can call the [DescribeDBResourceGroup](~~466685~~) operation to query the resource group name.
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
}

func (s CreateElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanRequest) SetDBClusterId(v string) *CreateElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanEnable(v bool) *CreateElasticPlanRequest {
	s.ElasticPlanEnable = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanEndDay(v string) *CreateElasticPlanRequest {
	s.ElasticPlanEndDay = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanMonthlyRepeat(v string) *CreateElasticPlanRequest {
	s.ElasticPlanMonthlyRepeat = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanName(v string) *CreateElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanNodeNum(v int32) *CreateElasticPlanRequest {
	s.ElasticPlanNodeNum = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanStartDay(v string) *CreateElasticPlanRequest {
	s.ElasticPlanStartDay = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanTimeEnd(v string) *CreateElasticPlanRequest {
	s.ElasticPlanTimeEnd = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanTimeStart(v string) *CreateElasticPlanRequest {
	s.ElasticPlanTimeStart = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanType(v string) *CreateElasticPlanRequest {
	s.ElasticPlanType = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanWeeklyRepeat(v string) *CreateElasticPlanRequest {
	s.ElasticPlanWeeklyRepeat = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanWorkerSpec(v string) *CreateElasticPlanRequest {
	s.ElasticPlanWorkerSpec = &v
	return s
}

func (s *CreateElasticPlanRequest) SetOwnerAccount(v string) *CreateElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateElasticPlanRequest) SetOwnerId(v int64) *CreateElasticPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateElasticPlanRequest) SetResourceOwnerAccount(v string) *CreateElasticPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateElasticPlanRequest) SetResourceOwnerId(v int64) *CreateElasticPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateElasticPlanRequest) SetResourcePoolName(v string) *CreateElasticPlanRequest {
	s.ResourcePoolName = &v
	return s
}

type CreateElasticPlanResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateElasticPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanResponseBody) SetRequestId(v string) *CreateElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

type CreateElasticPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateElasticPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanResponse) SetHeaders(v map[string]*string) *CreateElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateElasticPlanResponse) SetStatusCode(v int32) *CreateElasticPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateElasticPlanResponse) SetBody(v *CreateElasticPlanResponseBody) *CreateElasticPlanResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	// The account of the database.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// *   Normal: standard account
	// *   Super: privileged account
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The ID of the cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAccountRequest) SetAccountType(v string) *DeleteAccountRequest {
	s.AccountType = &v
	return s
}

func (s *DeleteAccountRequest) SetDBClusterId(v string) *DeleteAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAccountRequest) SetOwnerAccount(v string) *DeleteAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAccountRequest) SetOwnerId(v int64) *DeleteAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAccountRequest) SetResourceOwnerAccount(v string) *DeleteAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAccountRequest) SetResourceOwnerId(v int64) *DeleteAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteAccountResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBody) SetRequestId(v string) *DeleteAccountResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponse) SetHeaders(v map[string]*string) *DeleteAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccountResponse) SetStatusCode(v int32) *DeleteAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccountResponse) SetBody(v *DeleteAccountResponseBody) *DeleteAccountResponse {
	s.Body = v
	return s
}

type DeleteDBClusterRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterRequest) SetDBClusterId(v string) *DeleteDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetOwnerAccount(v string) *DeleteDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBClusterRequest) SetOwnerId(v int64) *DeleteDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetResourceOwnerAccount(v string) *DeleteDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBClusterRequest) SetResourceOwnerId(v int64) *DeleteDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteDBClusterResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponseBody) SetDBClusterId(v string) *DeleteDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterResponseBody) SetRequestId(v string) *DeleteDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBClusterResponseBody) SetTaskId(v int32) *DeleteDBClusterResponseBody {
	s.TaskId = &v
	return s
}

type DeleteDBClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponse) SetHeaders(v map[string]*string) *DeleteDBClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBClusterResponse) SetStatusCode(v int32) *DeleteDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBClusterResponse) SetBody(v *DeleteDBClusterResponseBody) *DeleteDBClusterResponse {
	s.Body = v
	return s
}

type DeleteDBResourceGroupRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBResourceGroupRequest) SetDBClusterId(v string) *DeleteDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) SetGroupName(v string) *DeleteDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) SetOwnerAccount(v string) *DeleteDBResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) SetOwnerId(v int64) *DeleteDBResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) SetResourceOwnerAccount(v string) *DeleteDBResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) SetResourceOwnerId(v int64) *DeleteDBResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteDBResourceGroupResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBResourceGroupResponseBody) SetRequestId(v string) *DeleteDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBResourceGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDBResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBResourceGroupResponse) SetHeaders(v map[string]*string) *DeleteDBResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBResourceGroupResponse) SetStatusCode(v int32) *DeleteDBResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBResourceGroupResponse) SetBody(v *DeleteDBResourceGroupResponseBody) *DeleteDBResourceGroupResponse {
	s.Body = v
	return s
}

type DeleteDBResourcePoolRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the resource group.
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBResourcePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBResourcePoolRequest) SetDBClusterId(v string) *DeleteDBResourcePoolRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetOwnerAccount(v string) *DeleteDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetOwnerId(v int64) *DeleteDBResourcePoolRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetPoolName(v string) *DeleteDBResourcePoolRequest {
	s.PoolName = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetResourceOwnerAccount(v string) *DeleteDBResourcePoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetResourceOwnerId(v int64) *DeleteDBResourcePoolRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteDBResourcePoolResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBResourcePoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResourcePoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBResourcePoolResponseBody) SetRequestId(v string) *DeleteDBResourcePoolResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBResourcePoolResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDBResourcePoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBResourcePoolResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBResourcePoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBResourcePoolResponse) SetHeaders(v map[string]*string) *DeleteDBResourcePoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBResourcePoolResponse) SetStatusCode(v int32) *DeleteDBResourcePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBResourcePoolResponse) SetBody(v *DeleteDBResourcePoolResponseBody) *DeleteDBResourcePoolResponse {
	s.Body = v
	return s
}

type DeleteElasticPlanRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~612241~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the scaling plan.
	//
	// > You can call the [DescribeElasticPlans](~~601334~~) operation to query the names of scaling plans.
	ElasticPlanName      *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteElasticPlanRequest) SetDBClusterId(v string) *DeleteElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetElasticPlanName(v string) *DeleteElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetOwnerAccount(v string) *DeleteElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetOwnerId(v int64) *DeleteElasticPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetResourceOwnerAccount(v string) *DeleteElasticPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetResourceOwnerId(v int64) *DeleteElasticPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteElasticPlanResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteElasticPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteElasticPlanResponseBody) SetRequestId(v string) *DeleteElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

type DeleteElasticPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteElasticPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteElasticPlanResponse) SetHeaders(v map[string]*string) *DeleteElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteElasticPlanResponse) SetStatusCode(v int32) *DeleteElasticPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteElasticPlanResponse) SetBody(v *DeleteElasticPlanResponseBody) *DeleteElasticPlanResponse {
	s.Body = v
	return s
}

type DescribeAccountsRequest struct {
	// The name of the database account.
	//
	// >  If you do not specify this parameter, the information about all database accounts is returned.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The type of the database account. If you do not specify this parameter, the information about all account types is returned. Valid values:
	//
	// *   **Normal**: standard account.
	// *   **Super**: privileged account.
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsRequest) SetAccountType(v string) *DescribeAccountsRequest {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsRequest) SetDBClusterId(v string) *DescribeAccountsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountsRequest) SetOwnerAccount(v string) *DescribeAccountsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAccountsRequest) SetOwnerId(v int64) *DescribeAccountsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccountsRequest) SetResourceOwnerAccount(v string) *DescribeAccountsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccountsRequest) SetResourceOwnerId(v int64) *DescribeAccountsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeAccountsResponseBody struct {
	// The queried database accounts.
	AccountList *DescribeAccountsResponseBodyAccountList `json:"AccountList,omitempty" xml:"AccountList,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) SetAccountList(v *DescribeAccountsResponseBodyAccountList) *DescribeAccountsResponseBody {
	s.AccountList = v
	return s
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAccountsResponseBodyAccountList struct {
	DBAccount []*DescribeAccountsResponseBodyAccountListDBAccount `json:"DBAccount,omitempty" xml:"DBAccount,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccountList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountList) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountList) SetDBAccount(v []*DescribeAccountsResponseBodyAccountListDBAccount) *DescribeAccountsResponseBodyAccountList {
	s.DBAccount = v
	return s
}

type DescribeAccountsResponseBodyAccountListDBAccount struct {
	// The description of the database account.
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The state of the database account. Valid values:
	//
	// *   **Creating**
	// *   **Available**
	// *   **Deleting**
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The type of the database account. Valid values:
	//
	// *   **Normal**: standard account.
	// *   **Super**: privileged account.
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountListDBAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountListDBAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountDescription(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountName(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountListDBAccount) SetAccountType(v string) *DescribeAccountsResponseBodyAccountListDBAccount {
	s.AccountType = &v
	return s
}

type DescribeAccountsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponse) SetHeaders(v map[string]*string) *DescribeAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountsResponse) SetStatusCode(v int32) *DescribeAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountsResponse) SetBody(v *DescribeAccountsResponseBody) *DescribeAccountsResponse {
	s.Body = v
	return s
}

type DescribeAdviceServiceEnabledRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of Data Warehouse Edition (V3.0) clusters.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAdviceServiceEnabledRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdviceServiceEnabledRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdviceServiceEnabledRequest) SetDBClusterId(v string) *DescribeAdviceServiceEnabledRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAdviceServiceEnabledRequest) SetRegionId(v string) *DescribeAdviceServiceEnabledRequest {
	s.RegionId = &v
	return s
}

type DescribeAdviceServiceEnabledResponseBody struct {
	// The returned message.
	//
	// *   If the request was successful, **Success** is returned.
	// *   If the request failed, an error message is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the suggestion feature is enabled. Valid values:
	//
	// *   **True**
	// *   **False**
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeAdviceServiceEnabledResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdviceServiceEnabledResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdviceServiceEnabledResponseBody) SetMessage(v string) *DescribeAdviceServiceEnabledResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAdviceServiceEnabledResponseBody) SetRequestId(v string) *DescribeAdviceServiceEnabledResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdviceServiceEnabledResponseBody) SetResult(v bool) *DescribeAdviceServiceEnabledResponseBody {
	s.Result = &v
	return s
}

type DescribeAdviceServiceEnabledResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAdviceServiceEnabledResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAdviceServiceEnabledResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdviceServiceEnabledResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdviceServiceEnabledResponse) SetHeaders(v map[string]*string) *DescribeAdviceServiceEnabledResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdviceServiceEnabledResponse) SetStatusCode(v int32) *DescribeAdviceServiceEnabledResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdviceServiceEnabledResponse) SetBody(v *DescribeAdviceServiceEnabledResponseBody) *DescribeAdviceServiceEnabledResponse {
	s.Body = v
	return s
}

type DescribeAllAccountsRequest struct {
	// The ID of the cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAllAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllAccountsRequest) SetDBClusterId(v string) *DescribeAllAccountsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllAccountsRequest) SetOwnerAccount(v string) *DescribeAllAccountsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAllAccountsRequest) SetOwnerId(v int64) *DescribeAllAccountsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAllAccountsRequest) SetResourceOwnerAccount(v string) *DescribeAllAccountsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAllAccountsRequest) SetResourceOwnerId(v int64) *DescribeAllAccountsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeAllAccountsResponseBody struct {
	// The list of accounts.
	AccountList []*DescribeAllAccountsResponseBodyAccountList `json:"AccountList,omitempty" xml:"AccountList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAllAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllAccountsResponseBody) SetAccountList(v []*DescribeAllAccountsResponseBodyAccountList) *DescribeAllAccountsResponseBody {
	s.AccountList = v
	return s
}

func (s *DescribeAllAccountsResponseBody) SetRequestId(v string) *DescribeAllAccountsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAllAccountsResponseBodyAccountList struct {
	// The name of the account.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeAllAccountsResponseBodyAccountList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllAccountsResponseBodyAccountList) GoString() string {
	return s.String()
}

func (s *DescribeAllAccountsResponseBodyAccountList) SetUser(v string) *DescribeAllAccountsResponseBodyAccountList {
	s.User = &v
	return s
}

type DescribeAllAccountsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAllAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAllAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllAccountsResponse) SetHeaders(v map[string]*string) *DescribeAllAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllAccountsResponse) SetStatusCode(v int32) *DescribeAllAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllAccountsResponse) SetBody(v *DescribeAllAccountsResponseBody) *DescribeAllAccountsResponse {
	s.Body = v
	return s
}

type DescribeAllDataSourceRequest struct {
	// The ID of the cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAllDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceRequest) SetDBClusterId(v string) *DescribeAllDataSourceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetOwnerAccount(v string) *DescribeAllDataSourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetOwnerId(v int64) *DescribeAllDataSourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetResourceOwnerAccount(v string) *DescribeAllDataSourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetResourceOwnerId(v int64) *DescribeAllDataSourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetSchemaName(v string) *DescribeAllDataSourceRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetTableName(v string) *DescribeAllDataSourceRequest {
	s.TableName = &v
	return s
}

type DescribeAllDataSourceResponseBody struct {
	// The queried columns.
	Columns *DescribeAllDataSourceResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried databases.
	Schemas *DescribeAllDataSourceResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Struct"`
	// The queried tables.
	Tables *DescribeAllDataSourceResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
}

func (s DescribeAllDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBody) SetColumns(v *DescribeAllDataSourceResponseBodyColumns) *DescribeAllDataSourceResponseBody {
	s.Columns = v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetRequestId(v string) *DescribeAllDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetSchemas(v *DescribeAllDataSourceResponseBodySchemas) *DescribeAllDataSourceResponseBody {
	s.Schemas = v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetTables(v *DescribeAllDataSourceResponseBodyTables) *DescribeAllDataSourceResponseBody {
	s.Tables = v
	return s
}

type DescribeAllDataSourceResponseBodyColumns struct {
	Column []*DescribeAllDataSourceResponseBodyColumnsColumn `json:"Column,omitempty" xml:"Column,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourceResponseBodyColumns) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyColumns) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyColumns) SetColumn(v []*DescribeAllDataSourceResponseBodyColumnsColumn) *DescribeAllDataSourceResponseBodyColumns {
	s.Column = v
	return s
}

type DescribeAllDataSourceResponseBodyColumnsColumn struct {
	// Indicates whether the column is an auto-increment column. Valid values:
	//
	// *   **true**
	// *   **false**
	AutoIncrementColumn *bool `json:"AutoIncrementColumn,omitempty" xml:"AutoIncrementColumn,omitempty"`
	// The name of the column.
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Indicates whether the column is the primary key of the table. Valid values:
	//
	// *   **true**
	// *   **false**
	PrimaryKey *bool `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The data type of the column.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAllDataSourceResponseBodyColumnsColumn) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyColumnsColumn) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetAutoIncrementColumn(v bool) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.AutoIncrementColumn = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetColumnName(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.ColumnName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetDBClusterId(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetPrimaryKey(v bool) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetSchemaName(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetTableName(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.TableName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetType(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.Type = &v
	return s
}

type DescribeAllDataSourceResponseBodySchemas struct {
	Schema []*DescribeAllDataSourceResponseBodySchemasSchema `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourceResponseBodySchemas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodySchemas) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodySchemas) SetSchema(v []*DescribeAllDataSourceResponseBodySchemasSchema) *DescribeAllDataSourceResponseBodySchemas {
	s.Schema = v
	return s
}

type DescribeAllDataSourceResponseBodySchemasSchema struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeAllDataSourceResponseBodySchemasSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodySchemasSchema) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodySchemasSchema) SetDBClusterId(v string) *DescribeAllDataSourceResponseBodySchemasSchema {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodySchemasSchema) SetSchemaName(v string) *DescribeAllDataSourceResponseBodySchemasSchema {
	s.SchemaName = &v
	return s
}

type DescribeAllDataSourceResponseBodyTables struct {
	Table []*DescribeAllDataSourceResponseBodyTablesTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourceResponseBodyTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyTables) SetTable(v []*DescribeAllDataSourceResponseBodyTablesTable) *DescribeAllDataSourceResponseBodyTables {
	s.Table = v
	return s
}

type DescribeAllDataSourceResponseBodyTablesTable struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAllDataSourceResponseBodyTablesTable) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyTablesTable) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) SetDBClusterId(v string) *DescribeAllDataSourceResponseBodyTablesTable {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) SetSchemaName(v string) *DescribeAllDataSourceResponseBodyTablesTable {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) SetTableName(v string) *DescribeAllDataSourceResponseBodyTablesTable {
	s.TableName = &v
	return s
}

type DescribeAllDataSourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAllDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAllDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponse) SetHeaders(v map[string]*string) *DescribeAllDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllDataSourceResponse) SetStatusCode(v int32) *DescribeAllDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllDataSourceResponse) SetBody(v *DescribeAllDataSourceResponseBody) *DescribeAllDataSourceResponse {
	s.Body = v
	return s
}

type DescribeAppliedAdvicesRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyyMMdd format. The time must be in UTC.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The display language of the suggestion. Valid values:
	//
	// *   **zh** (default): simplified Chinese.
	// *   **en**: English.
	// *   **ja**: Japanese.
	// *   **zh-tw**: traditional Chinese.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyyMMdd format. The time must be in UTC.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAppliedAdvicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppliedAdvicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppliedAdvicesRequest) SetDBClusterId(v string) *DescribeAppliedAdvicesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetEndTime(v int64) *DescribeAppliedAdvicesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetLang(v string) *DescribeAppliedAdvicesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetPageNumber(v int64) *DescribeAppliedAdvicesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetPageSize(v int64) *DescribeAppliedAdvicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetRegionId(v string) *DescribeAppliedAdvicesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAppliedAdvicesRequest) SetStartTime(v int64) *DescribeAppliedAdvicesRequest {
	s.StartTime = &v
	return s
}

type DescribeAppliedAdvicesResponseBody struct {
	// The queried suggestions.
	Items []*DescribeAppliedAdvicesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number. Pages start from page 1. Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned. The value is an integer that is greater than or equal to 0. Default value: 0.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppliedAdvicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppliedAdvicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppliedAdvicesResponseBody) SetItems(v []*DescribeAppliedAdvicesResponseBodyItems) *DescribeAppliedAdvicesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAppliedAdvicesResponseBody) SetPageNumber(v int64) *DescribeAppliedAdvicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBody) SetPageSize(v int64) *DescribeAppliedAdvicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBody) SetRequestId(v string) *DescribeAppliedAdvicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBody) SetTotalCount(v int64) *DescribeAppliedAdvicesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAppliedAdvicesResponseBodyItems struct {
	// The suggestion ID.
	AdviceId *string `json:"AdviceId,omitempty" xml:"AdviceId,omitempty"`
	// The benefit of the suggestion.
	Benefit *string `json:"Benefit,omitempty" xml:"Benefit,omitempty"`
	// The SQL statement that is used to execute the BUILD job.
	BuildSQL *string `json:"BuildSQL,omitempty" xml:"BuildSQL,omitempty"`
	// The state of the suggestion execution job. Valid values:
	//
	// *   **SUCCEED**
	// *   **FAILED**
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The SQL statement that is used to roll back the suggestion.
	RollbackSQL *string `json:"RollbackSQL,omitempty" xml:"RollbackSQL,omitempty"`
	// The SQL statement that is used to apply the suggestion.
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	// The submission state of the suggestion. Valid values:
	//
	// *   **SUCCEED**
	// *   **FAILED**
	SubmitStatus *string `json:"SubmitStatus,omitempty" xml:"SubmitStatus,omitempty"`
	// The time when the suggestion was submitted. The time follows the ISO 8601 standard in the yyMMddHHmm format. The time is displayed in UTC.
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// The total number of entries returned. Minimum value: 0. Default value: 0.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppliedAdvicesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppliedAdvicesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetAdviceId(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.AdviceId = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetBenefit(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.Benefit = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetBuildSQL(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.BuildSQL = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetJobStatus(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.JobStatus = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetPageNumber(v int64) *DescribeAppliedAdvicesResponseBodyItems {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetPageSize(v int64) *DescribeAppliedAdvicesResponseBodyItems {
	s.PageSize = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetRollbackSQL(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.RollbackSQL = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetSQL(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.SQL = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetSubmitStatus(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.SubmitStatus = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetSubmitTime(v string) *DescribeAppliedAdvicesResponseBodyItems {
	s.SubmitTime = &v
	return s
}

func (s *DescribeAppliedAdvicesResponseBodyItems) SetTotalCount(v int64) *DescribeAppliedAdvicesResponseBodyItems {
	s.TotalCount = &v
	return s
}

type DescribeAppliedAdvicesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppliedAdvicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppliedAdvicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppliedAdvicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppliedAdvicesResponse) SetHeaders(v map[string]*string) *DescribeAppliedAdvicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppliedAdvicesResponse) SetStatusCode(v int32) *DescribeAppliedAdvicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppliedAdvicesResponse) SetBody(v *DescribeAppliedAdvicesResponseBody) *DescribeAppliedAdvicesResponse {
	s.Body = v
	return s
}

type DescribeAuditLogConfigRequest struct {
	// The ID of the cluster.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAuditLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogConfigRequest) SetDBClusterId(v string) *DescribeAuditLogConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetOwnerAccount(v string) *DescribeAuditLogConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetOwnerId(v int64) *DescribeAuditLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetRegionId(v string) *DescribeAuditLogConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetResourceOwnerAccount(v string) *DescribeAuditLogConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetResourceOwnerId(v int64) *DescribeAuditLogConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeAuditLogConfigResponseBody struct {
	// The status of SQL audit. Valid values:
	//
	// *   **on**: SQL audit is enabled.
	// *   **off**: SQL audit is disabled.
	AuditLogStatus *string `json:"AuditLogStatus,omitempty" xml:"AuditLogStatus,omitempty"`
	// The ID of the cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAuditLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogConfigResponseBody) SetAuditLogStatus(v string) *DescribeAuditLogConfigResponseBody {
	s.AuditLogStatus = &v
	return s
}

func (s *DescribeAuditLogConfigResponseBody) SetDBClusterId(v string) *DescribeAuditLogConfigResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogConfigResponseBody) SetRequestId(v string) *DescribeAuditLogConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAuditLogConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAuditLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAuditLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogConfigResponse) SetHeaders(v map[string]*string) *DescribeAuditLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditLogConfigResponse) SetStatusCode(v int32) *DescribeAuditLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditLogConfigResponse) SetBody(v *DescribeAuditLogConfigResponseBody) *DescribeAuditLogConfigResponse {
	s.Body = v
	return s
}

type DescribeAuditLogRecordsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database on which you want to execute the SQL statement.
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// > - The end time must be later than the start time.
	// > - The maximum time range that can be specified is 24 hours.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address and port number of the client that is used to execute the SQL statement.
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The order in which specified fields are sorted. Specify this parameter as an ordered JSON array that consists of the Field and Type fields.
	//
	// *   Field specifies the field that is used to sort the retrieved entries. Valid values:
	//
	//     *   HostAddress: the IP address of the client that is used to connect to the database.
	//     *   Succeed: specifies whether the SQL statement is successfully executed.
	//     *   TotalTime: the total amount of time that is consumed to execute the SQL statement.
	//     *   DBName: the name of the database on which the SQL statement is executed.
	//     *   SQLType: the type of the SQL statement.
	//     *   User: the username that is used to execute the SQL statement.
	//     *   ExecuteTime: the time to start executing the SQL statement.
	//
	// *   Type specifies the sorting order. Valid values:
	//
	//     *   Desc: descending order.
	//     *   Asc: ascending order.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The sorting order of the retrieved entries. Valid values:
	//
	// *   **asc**: sorts the retrieved entries by time in ascending order.
	// *   **desc**: sorts the retrieved entries by time in descending order.
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value is an integer that is greater than 0. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   **10**
	// *   **30**
	// *   **50**
	// *   **100**
	//
	// > If you do not specify this parameter, the value 10 is used.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keywords that are included in the SQL statement to query.
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the SQL statement. Valid values:
	//
	// *   **DELETE**
	// *   **SELECT**
	// *   **UPDATE**
	// *   **INSERT_INTO_SELECT**
	// *   **ALTER**
	// *   **DROP**
	// *   **CREATE**
	//
	// > You can query only a single type of SQL statements at a time. If you leave this parameter empty, the **SELECT** statements are queried.
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// > SQL audit logs can be queried only when SQL audit is enabled. Only SQL audit logs within the last 30 days can be queried. If SQL audit was disabled and re-enabled, only SQL audit logs from the time when SQL audit was re-enabled can be queried.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Specifies whether the execution of the SQL statement succeeds. Valid values:
	//
	// *   **true**
	// *   **false**
	Succeed *string `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// The name of the user who executed the SQL statement.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeAuditLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsRequest) SetDBClusterId(v string) *DescribeAuditLogRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetDBName(v string) *DescribeAuditLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetEndTime(v string) *DescribeAuditLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetHostAddress(v string) *DescribeAuditLogRecordsRequest {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOrder(v string) *DescribeAuditLogRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOrderType(v string) *DescribeAuditLogRecordsRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOwnerAccount(v string) *DescribeAuditLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOwnerId(v int64) *DescribeAuditLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetPageNumber(v int32) *DescribeAuditLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetPageSize(v int32) *DescribeAuditLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetQueryKeyword(v string) *DescribeAuditLogRecordsRequest {
	s.QueryKeyword = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetRegionId(v string) *DescribeAuditLogRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeAuditLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeAuditLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetSqlType(v string) *DescribeAuditLogRecordsRequest {
	s.SqlType = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetStartTime(v string) *DescribeAuditLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetSucceed(v string) *DescribeAuditLogRecordsRequest {
	s.Succeed = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetUser(v string) *DescribeAuditLogRecordsRequest {
	s.User = &v
	return s
}

type DescribeAuditLogRecordsResponseBody struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried SQL audit logs.
	Items []*DescribeAuditLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAuditLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsResponseBody) SetDBClusterId(v string) *DescribeAuditLogRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetItems(v []*DescribeAuditLogRecordsResponseBodyItems) *DescribeAuditLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetPageNumber(v string) *DescribeAuditLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetPageSize(v string) *DescribeAuditLogRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetRequestId(v string) *DescribeAuditLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBody) SetTotalCount(v string) *DescribeAuditLogRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAuditLogRecordsResponseBodyItems struct {
	// This parameter is unavailable.
	ConnId *string `json:"ConnId,omitempty" xml:"ConnId,omitempty"`
	// The name of the database on which the SQL statement was executed.
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The start time of the execution of the SQL statement. The time is displayed in the ISO 8601 standard in the yyyy-MM-dd HH:mm:ss format. The time must be in UTC.
	ExecuteTime *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The IP address and port number of the client that is used to execute the SQL statement.
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The task ID.
	ProcessID *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
	// Details of the SQL statement.
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// The type of the SQL statement.
	SQLType *string `json:"SQLType,omitempty" xml:"SQLType,omitempty"`
	// Indicates whether the SQL statement was successfully executed. Valid values:
	//
	// *   **true**
	// *   **false**
	Succeed *string `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// The amount of time that is consumed to execute the SQL statement. Unit: milliseconds.
	TotalTime *string `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	// The name of the user who executed the SQL statement.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeAuditLogRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetConnId(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.ConnId = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetDBName(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetExecuteTime(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetHostAddress(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetProcessID(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.ProcessID = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetSQLText(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.SQLText = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetSQLType(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.SQLType = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetSucceed(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.Succeed = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetTotalTime(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.TotalTime = &v
	return s
}

func (s *DescribeAuditLogRecordsResponseBodyItems) SetUser(v string) *DescribeAuditLogRecordsResponseBodyItems {
	s.User = &v
	return s
}

type DescribeAuditLogRecordsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAuditLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAuditLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeAuditLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditLogRecordsResponse) SetStatusCode(v int32) *DescribeAuditLogRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditLogRecordsResponse) SetBody(v *DescribeAuditLogRecordsResponseBody) *DescribeAuditLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeAutoRenewAttributeRequest struct {
	// The cluster ID. Separate multiple clusters with commas (,).
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   30
	// *   50
	// *   100
	//
	// Default value: 30.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeRequest) SetDBClusterIds(v string) *DescribeAutoRenewAttributeRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetOwnerAccount(v string) *DescribeAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetOwnerId(v int64) *DescribeAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetPageNumber(v int32) *DescribeAutoRenewAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetPageSize(v int32) *DescribeAutoRenewAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetRegionId(v string) *DescribeAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetResourceGroupId(v string) *DescribeAutoRenewAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *DescribeAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *DescribeAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeAutoRenewAttributeResponseBody struct {
	// The renewal information of the cluster.
	Items *DescribeAutoRenewAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeAutoRenewAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseBody) SetItems(v *DescribeAutoRenewAttributeResponseBodyItems) *DescribeAutoRenewAttributeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetPageNumber(v int32) *DescribeAutoRenewAttributeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetPageRecordCount(v int32) *DescribeAutoRenewAttributeResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetRequestId(v string) *DescribeAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetTotalRecordCount(v int32) *DescribeAutoRenewAttributeResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeAutoRenewAttributeResponseBodyItems struct {
	AutoRenewAttribute []*DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute `json:"AutoRenewAttribute,omitempty" xml:"AutoRenewAttribute,omitempty" type:"Repeated"`
}

func (s DescribeAutoRenewAttributeResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseBodyItems) SetAutoRenewAttribute(v []*DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) *DescribeAutoRenewAttributeResponseBodyItems {
	s.AutoRenewAttribute = v
	return s
}

type DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute struct {
	// Indicates whether auto-renewal is enabled for the cluster. Valid values:
	//
	// *   **true**
	// *   **false**
	AutoRenewEnabled *bool `json:"AutoRenewEnabled,omitempty" xml:"AutoRenewEnabled,omitempty"`
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The renewal duration.
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The unit of the renewal duration. Valid values:
	//
	// *   **Year**
	// *   **Month**
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The renewal status of the cluster. Valid values:
	//
	// *   **AutoRenewal**: The cluster is automatically renewed.
	// *   **Normal**: The cluster is manually renewed. Before the cluster expires, the system sends you a reminder by SMS message.
	// *   **NotRenewal**: The cluster is not renewed. Three days before the cluster expires, the system sends you a reminder by SMS message to remind you that the cluster is not renewed. However, the system does not send you a reminder when the cluster expires.
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
}

func (s DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetAutoRenewEnabled(v bool) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.AutoRenewEnabled = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetDBClusterId(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetDuration(v int32) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.Duration = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetPeriodUnit(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetRegionId(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetRenewalStatus(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.RenewalStatus = &v
	return s
}

type DescribeAutoRenewAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutoRenewAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *DescribeAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoRenewAttributeResponse) SetStatusCode(v int32) *DescribeAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponse) SetBody(v *DescribeAutoRenewAttributeResponseBody) *DescribeAutoRenewAttributeResponse {
	s.Body = v
	return s
}

type DescribeAvailableAdvicesRequest struct {
	// The date when the suggestion is generated. Specify the date in the yyyyMMdd format. The date must be in UTC.
	AdviceDate *int64 `json:"AdviceDate,omitempty" xml:"AdviceDate,omitempty"`
	// The ID of the cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of Data Warehouse Edition (V3.0) clusters.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The display language of the suggestion. Default value: zh. Valid values:
	//
	// *   **zh**: simplified Chinese
	// *   **en**: English
	// *   **ja**: Japanese
	// *   **zh-tw**: traditional Chinese
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 30. Valid values:
	//
	// *   **30**
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAvailableAdvicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableAdvicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableAdvicesRequest) SetAdviceDate(v int64) *DescribeAvailableAdvicesRequest {
	s.AdviceDate = &v
	return s
}

func (s *DescribeAvailableAdvicesRequest) SetDBClusterId(v string) *DescribeAvailableAdvicesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAvailableAdvicesRequest) SetLang(v string) *DescribeAvailableAdvicesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAvailableAdvicesRequest) SetPageNumber(v int64) *DescribeAvailableAdvicesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAvailableAdvicesRequest) SetPageSize(v int64) *DescribeAvailableAdvicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAvailableAdvicesRequest) SetRegionId(v string) *DescribeAvailableAdvicesRequest {
	s.RegionId = &v
	return s
}

type DescribeAvailableAdvicesResponseBody struct {
	// The queried suggestions.
	Items []*DescribeAvailableAdvicesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page. The value must be an integer that is greater than 0. Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Default value: 30. Valid values:
	//
	// *   **30**
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned. The value must be an integer that is greater than or equal to 0. Default value: 0.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAvailableAdvicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableAdvicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableAdvicesResponseBody) SetItems(v []*DescribeAvailableAdvicesResponseBodyItems) *DescribeAvailableAdvicesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAvailableAdvicesResponseBody) SetPageNumber(v int64) *DescribeAvailableAdvicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBody) SetPageSize(v int64) *DescribeAvailableAdvicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBody) SetRequestId(v string) *DescribeAvailableAdvicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBody) SetTotalCount(v int64) *DescribeAvailableAdvicesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAvailableAdvicesResponseBodyItems struct {
	// The time when the suggestion was generated. The time follows the ISO 8601 standard in the yyyyMMdd format. The time is displayed in UTC.
	AdviceDate *string `json:"AdviceDate,omitempty" xml:"AdviceDate,omitempty"`
	// The suggestion ID.
	AdviceId *string `json:"AdviceId,omitempty" xml:"AdviceId,omitempty"`
	// The type of the suggestion. Valid values:
	//
	// *   **Index**: index optimization.
	// *   **Tiering**: hot and cold data optimization.
	AdviceType *string `json:"AdviceType,omitempty" xml:"AdviceType,omitempty"`
	// The benefit of the suggestion.
	Benefit *string `json:"Benefit,omitempty" xml:"Benefit,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The reason why the suggestion was generated.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The SQL statement that is used to apply the suggestion.
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	// The total number of entries returned. Minimum value: 0. Default value: 0.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAvailableAdvicesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableAdvicesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetAdviceDate(v string) *DescribeAvailableAdvicesResponseBodyItems {
	s.AdviceDate = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetAdviceId(v string) *DescribeAvailableAdvicesResponseBodyItems {
	s.AdviceId = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetAdviceType(v string) *DescribeAvailableAdvicesResponseBodyItems {
	s.AdviceType = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetBenefit(v string) *DescribeAvailableAdvicesResponseBodyItems {
	s.Benefit = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetPageNumber(v int64) *DescribeAvailableAdvicesResponseBodyItems {
	s.PageNumber = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetPageSize(v int64) *DescribeAvailableAdvicesResponseBodyItems {
	s.PageSize = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetReason(v string) *DescribeAvailableAdvicesResponseBodyItems {
	s.Reason = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetSQL(v string) *DescribeAvailableAdvicesResponseBodyItems {
	s.SQL = &v
	return s
}

func (s *DescribeAvailableAdvicesResponseBodyItems) SetTotalCount(v int64) *DescribeAvailableAdvicesResponseBodyItems {
	s.TotalCount = &v
	return s
}

type DescribeAvailableAdvicesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAvailableAdvicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableAdvicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableAdvicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableAdvicesResponse) SetHeaders(v map[string]*string) *DescribeAvailableAdvicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableAdvicesResponse) SetStatusCode(v int32) *DescribeAvailableAdvicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableAdvicesResponse) SetBody(v *DescribeAvailableAdvicesResponseBody) *DescribeAvailableAdvicesResponse {
	s.Body = v
	return s
}

type DescribeAvailableResourceRequest struct {
	// The language of query results. Valid values:
	//
	// *   **zh-CN** (default): Chinese.
	// *   **en-US**: English.
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The resources available in the supported modes.
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The version of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterVersion *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The zone ID.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent zone list.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) SetAcceptLanguage(v string) *DescribeAvailableResourceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetChargeType(v string) *DescribeAvailableResourceRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetDBClusterVersion(v string) *DescribeAvailableResourceRequest {
	s.DBClusterVersion = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetRegionId(v string) *DescribeAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetZoneId(v string) *DescribeAvailableResourceRequest {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourceResponseBody struct {
	// The supported zones.
	AvailableZoneList []*DescribeAvailableResourceResponseBodyAvailableZoneList `json:"AvailableZoneList,omitempty" xml:"AvailableZoneList,omitempty" type:"Repeated"`
	// The resources available in the supported editions.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The supported edition. Valid values:
	//
	// *   **basic**: Basic Edition
	// *   **cluster**: Cluster Edition
	// *   **mixed_storage**: elastic mode for Cluster Edition
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) SetAvailableZoneList(v []*DescribeAvailableResourceResponseBodyAvailableZoneList) *DescribeAvailableResourceResponseBody {
	s.AvailableZoneList = v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRegionId(v string) *DescribeAvailableResourceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneList struct {
	// A reserved parameter.
	SupportedComputeResource []*string `json:"SupportedComputeResource,omitempty" xml:"SupportedComputeResource,omitempty" type:"Repeated"`
	// The supported modes.
	SupportedMode []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode `json:"SupportedMode,omitempty" xml:"SupportedMode,omitempty" type:"Repeated"`
	// A reserved parameter.
	SupportedStorageResource []*string `json:"SupportedStorageResource,omitempty" xml:"SupportedStorageResource,omitempty" type:"Repeated"`
	// The zone ID.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) SetSupportedComputeResource(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneList {
	s.SupportedComputeResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) SetSupportedMode(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) *DescribeAvailableResourceResponseBodyAvailableZoneList {
	s.SupportedMode = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) SetSupportedStorageResource(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneList {
	s.SupportedStorageResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) SetZoneId(v string) *DescribeAvailableResourceResponseBodyAvailableZoneList {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode struct {
	// The supported mode. Valid values:
	//
	// *   **flexible**: elastic mode.
	// *   **reserver**: reserved mode.
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The supported editions.
	SupportedSerialList []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList `json:"SupportedSerialList,omitempty" xml:"SupportedSerialList,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) SetMode(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode {
	s.Mode = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) SetSupportedSerialList(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode {
	s.SupportedSerialList = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList struct {
	// The supported edition. Valid values:
	//
	// *   **basic**: Basic Edition.
	// *   **cluster**: Cluster Edition.
	// *   **mixed_storage**: elastic mode for Cluster Edition.
	Serial *string `json:"Serial,omitempty" xml:"Serial,omitempty"`
	// The supported resources in elastic mode.
	SupportedFlexibleResource []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource `json:"SupportedFlexibleResource,omitempty" xml:"SupportedFlexibleResource,omitempty" type:"Repeated"`
	// The supported resources in reserved mode.
	SupportedInstanceClassList []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList `json:"SupportedInstanceClassList,omitempty" xml:"SupportedInstanceClassList,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) SetSerial(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList {
	s.Serial = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) SetSupportedFlexibleResource(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList {
	s.SupportedFlexibleResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) SetSupportedInstanceClassList(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList {
	s.SupportedInstanceClassList = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource struct {
	// The disk storage type. Valid values:
	//
	// *   **hdd**
	// *   **ssd**
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The supported computing resources.
	SupportedComputeResource []*string `json:"SupportedComputeResource,omitempty" xml:"SupportedComputeResource,omitempty" type:"Repeated"`
	// The supported elastic I/O resources.
	SupportedElasticIOResource *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource `json:"SupportedElasticIOResource,omitempty" xml:"SupportedElasticIOResource,omitempty" type:"Struct"`
	// The supported storage resources.
	SupportedStorageResource []*string `json:"SupportedStorageResource,omitempty" xml:"SupportedStorageResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetStorageType(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.StorageType = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetSupportedComputeResource(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.SupportedComputeResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetSupportedElasticIOResource(v *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.SupportedElasticIOResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetSupportedStorageResource(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.SupportedStorageResource = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource struct {
	// The maximum amount of elastic I/O resources.
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The minimum amount of elastic I/O resources.
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// The step size.
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) SetMaxCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) SetMinCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) SetStep(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource {
	s.Step = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList struct {
	// The supported instance type.
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// A reserved parameter.
	SupportedExecutorList []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList `json:"SupportedExecutorList,omitempty" xml:"SupportedExecutorList,omitempty" type:"Repeated"`
	// The supported compute nodes.
	SupportedNodeCountList []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList `json:"SupportedNodeCountList,omitempty" xml:"SupportedNodeCountList,omitempty" type:"Repeated"`
	// The description of the instance type.
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) SetInstanceClass(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList {
	s.InstanceClass = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) SetSupportedExecutorList(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList {
	s.SupportedExecutorList = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) SetSupportedNodeCountList(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList {
	s.SupportedNodeCountList = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) SetTips(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList {
	s.Tips = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList struct {
	// The information about the supported compute nodes.
	NodeCount *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount `json:"NodeCount,omitempty" xml:"NodeCount,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList) SetNodeCount(v *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList {
	s.NodeCount = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount struct {
	// A reserved parameter.
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// A reserved parameter.
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// A reserved parameter.
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) SetMaxCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) SetMinCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) SetStep(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount {
	s.Step = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList struct {
	// The number of the supported compute nodes.
	NodeCount *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount `json:"NodeCount,omitempty" xml:"NodeCount,omitempty" type:"Struct"`
	// The support storage capacity. Unit: GB.
	StorageSize []*string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) SetNodeCount(v *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList {
	s.NodeCount = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) SetStorageSize(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList {
	s.StorageSize = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount struct {
	// The maximum number of compute nodes.
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The minimum number of compute nodes.
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// The step size.
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) SetMaxCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) SetMinCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) SetStep(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount {
	s.Step = &v
	return s
}

type DescribeAvailableResourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponse) SetHeaders(v map[string]*string) *DescribeAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableResourceResponse) SetStatusCode(v int32) *DescribeAvailableResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableResourceResponse) SetBody(v *DescribeAvailableResourceResponseBody) *DescribeAvailableResourceResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	// The ID of the cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) SetDBClusterId(v string) *DescribeBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeBackupPolicyResponseBody struct {
	// The number of days for which data backup files are retained.
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// Specifies whether to enable the origin protocol policy.
	//
	// *   true: enabled
	// *   false: disabled
	EnableBackupLog *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// The number of days for which log backup files are retained.
	LogBackupRetentionPeriod *int32 `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	// The cycle based on which backups are performed. If more than one day of the week is specified, the days of the week are separated by commas (,). Valid values:
	//
	// *   Monday
	// *   Tuesday
	// *   Wednesday
	// *   Thursday
	// *   Friday
	// *   Saturday
	// *   Sunday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The backup time. Specify the time in the HH:mmZ-HH:mmZ format.
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetEnableBackupLog(v string) *DescribeBackupPolicyResponseBody {
	s.EnableBackupLog = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetLogBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBody {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBackupPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPolicyResponse) SetStatusCode(v int32) *DescribeBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse {
	s.Body = v
	return s
}

type DescribeBackupsRequest struct {
	// The ID of the backup set.
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC. The end time must be later than the start time.
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupsRequest) SetBackupId(v string) *DescribeBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsRequest) SetDBClusterId(v string) *DescribeBackupsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTime(v string) *DescribeBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetOwnerAccount(v string) *DescribeBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupsRequest) SetOwnerId(v int64) *DescribeBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageNumber(v int32) *DescribeBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageSize(v int32) *DescribeBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsRequest) SetResourceOwnerAccount(v string) *DescribeBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupsRequest) SetResourceOwnerId(v int64) *DescribeBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTime(v string) *DescribeBackupsRequest {
	s.StartTime = &v
	return s
}

type DescribeBackupsResponseBody struct {
	// The queried backup sets.
	Items *DescribeBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) SetItems(v *DescribeBackupsResponseBodyItems) *DescribeBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageNumber(v string) *DescribeBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageSize(v string) *DescribeBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetRequestId(v string) *DescribeBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetTotalCount(v string) *DescribeBackupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeBackupsResponseBodyItems struct {
	Backup []*DescribeBackupsResponseBodyItemsBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Repeated"`
}

func (s DescribeBackupsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItems) SetBackup(v []*DescribeBackupsResponseBodyItemsBackup) *DescribeBackupsResponseBodyItems {
	s.Backup = v
	return s
}

type DescribeBackupsResponseBodyItemsBackup struct {
	// The end time of the backup.
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The backup set ID.
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup method. Only Snapshot is returned.
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The size of the backup set. Unit: bytes.
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The start time of the backup.
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The backup type. Valid values:
	//
	// *   **FullBackup**
	// *   **IncrementalBackup**
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeBackupsResponseBodyItemsBackup) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyItemsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupMethod(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupSize(v int64) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupStartTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupType(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetDBClusterId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.DBClusterId = &v
	return s
}

type DescribeBackupsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponse) SetHeaders(v map[string]*string) *DescribeBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupsResponse) SetStatusCode(v int32) *DescribeBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupsResponse) SetBody(v *DescribeBackupsResponseBody) *DescribeBackupsResponse {
	s.Body = v
	return s
}

type DescribeColumnsRequest struct {
	// The ID of the cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeColumnsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsRequest) GoString() string {
	return s.String()
}

func (s *DescribeColumnsRequest) SetDBClusterId(v string) *DescribeColumnsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeColumnsRequest) SetOwnerAccount(v string) *DescribeColumnsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeColumnsRequest) SetOwnerId(v int64) *DescribeColumnsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeColumnsRequest) SetResourceOwnerAccount(v string) *DescribeColumnsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeColumnsRequest) SetResourceOwnerId(v int64) *DescribeColumnsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeColumnsRequest) SetSchemaName(v string) *DescribeColumnsRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeColumnsRequest) SetTableName(v string) *DescribeColumnsRequest {
	s.TableName = &v
	return s
}

type DescribeColumnsResponseBody struct {
	// The list of columns.
	Items *DescribeColumnsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeColumnsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBody) SetItems(v *DescribeColumnsResponseBodyItems) *DescribeColumnsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeColumnsResponseBody) SetRequestId(v string) *DescribeColumnsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeColumnsResponseBodyItems struct {
	Column []*DescribeColumnsResponseBodyItemsColumn `json:"Column,omitempty" xml:"Column,omitempty" type:"Repeated"`
}

func (s DescribeColumnsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBodyItems) SetColumn(v []*DescribeColumnsResponseBodyItemsColumn) *DescribeColumnsResponseBodyItems {
	s.Column = v
	return s
}

type DescribeColumnsResponseBodyItemsColumn struct {
	// Indicates whether the columns are auto-incremented.
	AutoIncrementColumn *bool `json:"AutoIncrementColumn,omitempty" xml:"AutoIncrementColumn,omitempty"`
	// The name of the column.
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The ID of the cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Indicates whether the column is a primary key.
	PrimaryKey *bool `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The data type of the column.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeColumnsResponseBodyItemsColumn) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponseBodyItemsColumn) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetAutoIncrementColumn(v bool) *DescribeColumnsResponseBodyItemsColumn {
	s.AutoIncrementColumn = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetColumnName(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.ColumnName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetDBClusterId(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.DBClusterId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetPrimaryKey(v bool) *DescribeColumnsResponseBodyItemsColumn {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetSchemaName(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.SchemaName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetTableName(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.TableName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetType(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.Type = &v
	return s
}

type DescribeColumnsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeColumnsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponse) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponse) SetHeaders(v map[string]*string) *DescribeColumnsResponse {
	s.Headers = v
	return s
}

func (s *DescribeColumnsResponse) SetStatusCode(v int32) *DescribeColumnsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeColumnsResponse) SetBody(v *DescribeColumnsResponseBody) *DescribeColumnsResponse {
	s.Body = v
	return s
}

type DescribeComputeResourceRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The version of the AnalyticDB for MySQL Data Warehouse Edition cluster. Set the value to **3**.
	DBClusterVersion *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	// The available computing resources for migrating AnalyticDB MySQL Data Warehouse Edition to AnalyticDB MySQL Lakehouse Edition. Possible values are:
	// - **true**
	// - **false**(default value)
	Migrate      *bool   `json:"Migrate,omitempty" xml:"Migrate,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The zone ID of the cluster.
	//
	// > You can call the [DescribeRegions](~~129857~~) operation to query the most recent zone list.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeComputeResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComputeResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceRequest) SetDBClusterId(v string) *DescribeComputeResourceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetDBClusterVersion(v string) *DescribeComputeResourceRequest {
	s.DBClusterVersion = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetMigrate(v bool) *DescribeComputeResourceRequest {
	s.Migrate = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetOwnerAccount(v string) *DescribeComputeResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetOwnerId(v int64) *DescribeComputeResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetRegionId(v string) *DescribeComputeResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetResourceOwnerAccount(v string) *DescribeComputeResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetResourceOwnerId(v int64) *DescribeComputeResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeComputeResourceRequest) SetZoneId(v string) *DescribeComputeResourceRequest {
	s.ZoneId = &v
	return s
}

type DescribeComputeResourceResponseBody struct {
	// The queried specifications of computing resources.
	ComputeResource []*DescribeComputeResourceResponseBodyComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComputeResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComputeResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceResponseBody) SetComputeResource(v []*DescribeComputeResourceResponseBodyComputeResource) *DescribeComputeResourceResponseBody {
	s.ComputeResource = v
	return s
}

func (s *DescribeComputeResourceResponseBody) SetRequestId(v string) *DescribeComputeResourceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeComputeResourceResponseBodyComputeResource struct {
	// The specifications of computing resources displayed in the console.
	DisplayValue *string `json:"DisplayValue,omitempty" xml:"DisplayValue,omitempty"`
	// The actual specifications of computing resources.
	RealValue *string `json:"RealValue,omitempty" xml:"RealValue,omitempty"`
}

func (s DescribeComputeResourceResponseBodyComputeResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeComputeResourceResponseBodyComputeResource) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceResponseBodyComputeResource) SetDisplayValue(v string) *DescribeComputeResourceResponseBodyComputeResource {
	s.DisplayValue = &v
	return s
}

func (s *DescribeComputeResourceResponseBodyComputeResource) SetRealValue(v string) *DescribeComputeResourceResponseBodyComputeResource {
	s.RealValue = &v
	return s
}

type DescribeComputeResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeComputeResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeComputeResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComputeResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceResponse) SetHeaders(v map[string]*string) *DescribeComputeResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeComputeResourceResponse) SetStatusCode(v int32) *DescribeComputeResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComputeResourceResponse) SetBody(v *DescribeComputeResourceResponseBody) *DescribeComputeResourceResponse {
	s.Body = v
	return s
}

type DescribeConnectionCountRecordsRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeConnectionCountRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionCountRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsRequest) SetDBClusterId(v string) *DescribeConnectionCountRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeConnectionCountRecordsRequest) SetOwnerAccount(v string) *DescribeConnectionCountRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeConnectionCountRecordsRequest) SetOwnerId(v int64) *DescribeConnectionCountRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeConnectionCountRecordsRequest) SetResourceOwnerAccount(v string) *DescribeConnectionCountRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeConnectionCountRecordsRequest) SetResourceOwnerId(v int64) *DescribeConnectionCountRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeConnectionCountRecordsResponseBody struct {
	// The queried client IP addresses.
	AccessIpRecords []*DescribeConnectionCountRecordsResponseBodyAccessIpRecords `json:"AccessIpRecords,omitempty" xml:"AccessIpRecords,omitempty" type:"Repeated"`
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried database accounts.
	UserRecords []*DescribeConnectionCountRecordsResponseBodyUserRecords `json:"UserRecords,omitempty" xml:"UserRecords,omitempty" type:"Repeated"`
}

func (s DescribeConnectionCountRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionCountRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsResponseBody) SetAccessIpRecords(v []*DescribeConnectionCountRecordsResponseBodyAccessIpRecords) *DescribeConnectionCountRecordsResponseBody {
	s.AccessIpRecords = v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBody) SetDBClusterId(v string) *DescribeConnectionCountRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBody) SetRequestId(v string) *DescribeConnectionCountRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBody) SetUserRecords(v []*DescribeConnectionCountRecordsResponseBodyUserRecords) *DescribeConnectionCountRecordsResponseBody {
	s.UserRecords = v
	return s
}

type DescribeConnectionCountRecordsResponseBodyAccessIpRecords struct {
	// The IP address of the client.
	AccessIp *string `json:"AccessIp,omitempty" xml:"AccessIp,omitempty"`
	// The number of connections.
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeConnectionCountRecordsResponseBodyAccessIpRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionCountRecordsResponseBodyAccessIpRecords) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsResponseBodyAccessIpRecords) SetAccessIp(v string) *DescribeConnectionCountRecordsResponseBodyAccessIpRecords {
	s.AccessIp = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBodyAccessIpRecords) SetCount(v int64) *DescribeConnectionCountRecordsResponseBodyAccessIpRecords {
	s.Count = &v
	return s
}

type DescribeConnectionCountRecordsResponseBodyUserRecords struct {
	// The number of connections.
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The username of the database account.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeConnectionCountRecordsResponseBodyUserRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionCountRecordsResponseBodyUserRecords) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsResponseBodyUserRecords) SetCount(v int64) *DescribeConnectionCountRecordsResponseBodyUserRecords {
	s.Count = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponseBodyUserRecords) SetUser(v string) *DescribeConnectionCountRecordsResponseBodyUserRecords {
	s.User = &v
	return s
}

type DescribeConnectionCountRecordsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeConnectionCountRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConnectionCountRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionCountRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeConnectionCountRecordsResponse) SetHeaders(v map[string]*string) *DescribeConnectionCountRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeConnectionCountRecordsResponse) SetStatusCode(v int32) *DescribeConnectionCountRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConnectionCountRecordsResponse) SetBody(v *DescribeConnectionCountRecordsResponseBody) *DescribeConnectionCountRecordsResponse {
	s.Body = v
	return s
}

type DescribeDBClusterAccessWhiteListRequest struct {
	// The cluster ID.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterAccessWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetDBClusterId(v string) *DescribeDBClusterAccessWhiteListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetOwnerAccount(v string) *DescribeDBClusterAccessWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetOwnerId(v int64) *DescribeDBClusterAccessWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAccessWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAccessWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBClusterAccessWhiteListResponseBody struct {
	// The queried IP address whitelists.
	Items *DescribeDBClusterAccessWhiteListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterAccessWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) SetItems(v *DescribeDBClusterAccessWhiteListResponseBodyItems) *DescribeDBClusterAccessWhiteListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) SetRequestId(v string) *DescribeDBClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterAccessWhiteListResponseBodyItems struct {
	IPArray []*DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray `json:"IPArray,omitempty" xml:"IPArray,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAccessWhiteListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItems) SetIPArray(v []*DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) *DescribeDBClusterAccessWhiteListResponseBodyItems {
	s.IPArray = v
	return s
}

type DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray struct {
	// The attribute of the IP address whitelist. By default, this parameter is empty.
	//
	// >  The IP address whitelists that have the **hidden** attribute are not displayed in the console. These IP address whitelists are used to access services such as Data Transmission Service (DTS) and PolarDB-X.
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	// The name of the IP address whitelist.
	//
	// *   The name of an IP address whitelist must be 2 to 32 characters in length. The name can contain lowercase letters, digits, and underscores (\_). The name must start with a lowercase letter and end with a lowercase letter or digit.
	// *   Each cluster supports up to 50 IP address whitelists.
	DBClusterIPArrayName *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	// The IP addresses in the IP address whitelist. Up to 1,000 IP addresses can be returned. Multiple IP addresses are separated by commas (,).
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) SetDBClusterIPArrayAttribute(v string) *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) SetDBClusterIPArrayName(v string) *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) SetSecurityIPList(v string) *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray {
	s.SecurityIPList = &v
	return s
}

type DescribeDBClusterAccessWhiteListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBClusterAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterAccessWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponse) SetHeaders(v map[string]*string) *DescribeDBClusterAccessWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponse) SetStatusCode(v int32) *DescribeDBClusterAccessWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponse) SetBody(v *DescribeDBClusterAccessWhiteListResponseBody) *DescribeDBClusterAccessWhiteListResponse {
	s.Body = v
	return s
}

type DescribeDBClusterAttributeRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeRequest) SetDBClusterId(v string) *DescribeDBClusterAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetOwnerAccount(v string) *DescribeDBClusterAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetOwnerId(v int64) *DescribeDBClusterAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBClusterAttributeResponseBody struct {
	// The queried cluster information.
	Items *DescribeDBClusterAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBody) SetItems(v *DescribeDBClusterAttributeResponseBodyItems) *DescribeDBClusterAttributeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetRequestId(v string) *DescribeDBClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterAttributeResponseBodyItems struct {
	DBCluster []*DescribeDBClusterAttributeResponseBodyItemsDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAttributeResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItems) SetDBCluster(v []*DescribeDBClusterAttributeResponseBodyItemsDBCluster) *DescribeDBClusterAttributeResponseBodyItems {
	s.DBCluster = v
	return s
}

type DescribeDBClusterAttributeResponseBodyItemsDBCluster struct {
	// The edition of the cluster. Valid values:
	//
	// *   **BASIC**: reserved mode for Basic Edition.
	// *   **CLUSTER**: reserved mode for Cluster Edition.
	// *   **MIXED_STORAGE**: elastic mode for Cluster Edition.
	//
	// >  For more information about cluster editions, see [Editions](~~205001~~).
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// *   **ads**: pay-as-you-go.
	// *   **ads_pre**: subscription.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The specifications of computing resources that are used in the cluster in elastic mode. Computing resources are used to compute data. The increase in the computing resources can accelerate queries. You can scale computing resources based on your business requirements.
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The Virtual Private Cloud (VPC) endpoint that is used to connect to the cluster.
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The time when the cluster was created. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the cluster.
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The network type of the cluster. **VPC** is returned.
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The state of the cluster. For more information, see [Cluster states](~~143075~~).
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The cluster type. Valid values:
	//
	// *   **Common**: common cluster.
	// *   **RDS_ANALYSIS**: MySQL analytic instance.
	DBClusterType *string `json:"DBClusterType,omitempty" xml:"DBClusterType,omitempty"`
	// The instance type of the cluster.
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of node groups.
	DBNodeCount *int64 `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	// The storage capacity of the cluster. Unit: GB.
	DBNodeStorage *int64 `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	// The engine version of the cluster. **3.0** is returned.
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// Indicates whether disk encryption is enabled. Valid values:
	//
	// *   true
	// *   false
	DiskEncryption *string `json:"DiskEncryption,omitempty" xml:"DiskEncryption,omitempty"`
	// The ESSD performance level.
	DiskPerformanceLevel *string `json:"DiskPerformanceLevel,omitempty" xml:"DiskPerformanceLevel,omitempty"`
	// The disk type of the cluster. Valid values:
	//
	// *   **local_ssd**: local disk.
	// *   **cloud**: basic disk.
	// *   **cloud_ssd**: standard SSD.
	// *   **cloud_efficiency**: ultra disk.
	// *   **cloud_essd0**: PL0 enhanced SSD (ESSD).
	// *   **cloud_essd**: PL1 ESSD.
	// *   **cloud_essd2**: PL2 ESSD.
	// *   **cloud_essd3**: PL3 ESSD.
	//
	// >  For more information about ESSDs, see [ESSDs](~~122389~~).
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the Data Transmission Service (DTS) synchronization job. This parameter is returned only for MySQL analytic instances.
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The number of elastic I/O units (EIUs).
	ElasticIOResource *int32 `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	// The single-node specifications of an EIU. Valid values:
	//
	// *   8Core64GB: If this value is returned, the specifications of an EIU are 24 cores and 192 GB memory.
	// *   12Core96GB: If this value is returned, the specifications of an EIU are 36 cores and 288 GB memory.
	ElasticIOResourceSize *string `json:"ElasticIOResourceSize,omitempty" xml:"ElasticIOResourceSize,omitempty"`
	// Indicates whether an Airflow cluster was created. Valid values:
	//
	// *   **true**
	// *   **false**
	EnableAirflow *bool `json:"EnableAirflow,omitempty" xml:"EnableAirflow,omitempty"`
	// Indicates whether a Spark cluster was created. Valid values:
	//
	// *   **true**
	// *   **false**
	EnableSpark *bool `json:"EnableSpark,omitempty" xml:"EnableSpark,omitempty"`
	// The engine of the cluster. **AnalyticDB** is returned.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The minor version of the cluster.
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The number of compute nodes that are used by the cluster in elastic mode.
	ExecutorCount *string `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	// The expiration time of the cluster. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC. Example: *2999-09-08T16:00:00Z*.
	//
	// >
	//
	// *   If the billing method of the cluster is subscription, the actual expiration time is returned.
	//
	// *   If the billing method of the cluster is pay-as-you-go, **2999-09-08T16:00:00Z** is returned.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the cluster has expired. Valid values:
	//
	// *   **true**
	// *   **false**
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The public IP address of the cluster.
	InnerIp *string `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	// The public port number.
	InnerPort *string `json:"InnerPort,omitempty" xml:"InnerPort,omitempty"`
	// The ID of the key that is used to encrypt disk data.
	//
	// >  This parameter is returned only when disk encryption is enabled.
	KmsId *string `json:"KmsId,omitempty" xml:"KmsId,omitempty"`
	// The lock mode of the cluster. Valid values:
	//
	// *   **Unlock**: The cluster is not locked.
	// *   **ManualLock**: The cluster is manually locked.
	// *   **LockByExpiration**: The cluster is automatically locked due to cluster expiration.
	// *   **LockByRestoration**: The cluster is automatically locked due to cluster restoration.
	// *   **LockByDiskQuota**: The cluster is automatically locked when 90% of the cluster storage is used.
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the cluster is locked.
	//
	// >  This parameter is returned only when the cluster was locked. **instance_expire** is returned.
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The maintenance window of the cluster. The time is displayed in the *HH:mmZ-HH:mmZ* format in UTC. An example is *04:00Z-05:00Z*, which indicates that routine maintenance is performed from 04:00 to 05:00.
	//
	// >  For more information about maintenance windows, see [Configure a maintenance window](~~122569~~).
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	// The mode of the cluster. Valid values:
	//
	// *   **flexible**: elastic mode.
	// *   **reserver**: reserved mode.
	//
	// >  For more information about cluster modes, see [Editions](~~205001~~).
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// *   **Postpaid**: pay-as-you-go.
	// *   **Prepaid**: subscription.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port number that is used to connect to the cluster.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the ApsaraDB RDS instance from which data is synchronized to the cluster. This parameter is returned only for MySQL analytic instances.
	RdsInstanceId *string `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The specifications of storage resources that are used in the cluster in elastic mode. Storage resources are used to read and write data. The increase in the storage resources can improve the read and write performance of the cluster.
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	// The tags that are added to the cluster.
	Tags *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// Indicates whether Elastic Network Interface (ENI) is enabled. Valid values:
	//
	// *   **true**
	// *   **false**
	UserENIStatus *bool `json:"UserENIStatus,omitempty" xml:"UserENIStatus,omitempty"`
	// The ID of the cluster that resides in the VPC.
	VPCCloudInstanceId *string `json:"VPCCloudInstanceId,omitempty" xml:"VPCCloudInstanceId,omitempty"`
	// The VPC ID of the cluster.
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the cluster.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the cluster.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetCategory(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Category = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetCommodityCode(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetComputeResource(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ComputeResource = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetConnectionString(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetCreationTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBClusterType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBClusterType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBNodeClass(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBNodeCount(v int64) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBNodeStorage(v int64) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBNodeStorage = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDiskEncryption(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DiskEncryption = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDiskPerformanceLevel(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DiskPerformanceLevel = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDiskType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DiskType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetDtsJobId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetElasticIOResource(v int32) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ElasticIOResource = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetElasticIOResourceSize(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ElasticIOResourceSize = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEnableAirflow(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.EnableAirflow = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEnableSpark(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.EnableSpark = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEngine(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetEngineVersion(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExecutorCount(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ExecutorCount = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetExpired(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetInnerIp(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.InnerIp = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetInnerPort(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.InnerPort = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetKmsId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.KmsId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetLockReason(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.LockReason = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetMaintainTime(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.MaintainTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetMode(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Mode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetPayType(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetPort(v int32) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetRdsInstanceId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.RdsInstanceId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetResourceGroupId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetStorageResource(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.StorageResource = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetTags(v *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetUserENIStatus(v bool) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.UserENIStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetVPCCloudInstanceId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.VPCCloudInstanceId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetVPCId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetVSwitchId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeDBClusterAttributeResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

type DescribeDBClusterAttributeResponseBodyItemsDBClusterTags struct {
	Tag []*DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags) SetTag(v []*DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTags {
	s.Tag = v
	return s
}

type DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag struct {
	// The tag key.
	//
	// >  You can call the [TagResources](~~179253~~) operation to add a tag to the cluster.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) SetKey(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag) SetValue(v string) *DescribeDBClusterAttributeResponseBodyItemsDBClusterTagsTag {
	s.Value = &v
	return s
}

type DescribeDBClusterAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponse) SetHeaders(v map[string]*string) *DescribeDBClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetStatusCode(v int32) *DescribeDBClusterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetBody(v *DescribeDBClusterAttributeResponseBody) *DescribeDBClusterAttributeResponse {
	s.Body = v
	return s
}

type DescribeDBClusterHealthStatusRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the region.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBClusterHealthStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusRequest) SetDBClusterId(v string) *DescribeDBClusterHealthStatusRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterHealthStatusRequest) SetRegionId(v string) *DescribeDBClusterHealthStatusRequest {
	s.RegionId = &v
	return s
}

type DescribeDBClusterHealthStatusResponseBody struct {
	// Health state details of access nodes.
	CS *DescribeDBClusterHealthStatusResponseBodyCS `json:"CS,omitempty" xml:"CS,omitempty" type:"Struct"`
	// Health state details of compute node groups.
	Executor *DescribeDBClusterHealthStatusResponseBodyExecutor `json:"Executor,omitempty" xml:"Executor,omitempty" type:"Struct"`
	// The health state of the cluster. Valid values:
	//
	// *   **RISK**: risky
	//
	// *   **NORMAL**: healthy
	//
	// *   **UNAVAILABLE**: unavailable
	//
	// > If the health states of access nodes, compute node groups, and storage node groups are all **healthy** and the cluster is detected to be alive, the health state of the cluster is **healthy**. If the preceding three health states include **risky**, the health state of the cluster is **risky**. If the preceding three health states include **unavailable**, the health state of the cluster is **unavailable**.
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Health state details of storage node groups.
	Worker *DescribeDBClusterHealthStatusResponseBodyWorker `json:"Worker,omitempty" xml:"Worker,omitempty" type:"Struct"`
}

func (s DescribeDBClusterHealthStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetCS(v *DescribeDBClusterHealthStatusResponseBodyCS) *DescribeDBClusterHealthStatusResponseBody {
	s.CS = v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetExecutor(v *DescribeDBClusterHealthStatusResponseBodyExecutor) *DescribeDBClusterHealthStatusResponseBody {
	s.Executor = v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetInstanceStatus(v string) *DescribeDBClusterHealthStatusResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetRequestId(v string) *DescribeDBClusterHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetWorker(v *DescribeDBClusterHealthStatusResponseBodyWorker) *DescribeDBClusterHealthStatusResponseBody {
	s.Worker = v
	return s
}

type DescribeDBClusterHealthStatusResponseBodyCS struct {
	// The number of healthy access nodes.
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The total number of access nodes.
	ExpectedCount *int64 `json:"ExpectedCount,omitempty" xml:"ExpectedCount,omitempty"`
	// The number of risky access nodes.
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The health state of access nodes. Valid values:
	//
	// *   **RISK**: risky
	// *   **NORMAL**: healthy
	// *   **UNAVAILABLE**: unavailable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of unavailable access nodes.
	UnavailableCount *int64 `json:"UnavailableCount,omitempty" xml:"UnavailableCount,omitempty"`
}

func (s DescribeDBClusterHealthStatusResponseBodyCS) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthStatusResponseBodyCS) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) SetActiveCount(v int64) *DescribeDBClusterHealthStatusResponseBodyCS {
	s.ActiveCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) SetExpectedCount(v int64) *DescribeDBClusterHealthStatusResponseBodyCS {
	s.ExpectedCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) SetRiskCount(v int64) *DescribeDBClusterHealthStatusResponseBodyCS {
	s.RiskCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) SetStatus(v string) *DescribeDBClusterHealthStatusResponseBodyCS {
	s.Status = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) SetUnavailableCount(v int64) *DescribeDBClusterHealthStatusResponseBodyCS {
	s.UnavailableCount = &v
	return s
}

type DescribeDBClusterHealthStatusResponseBodyExecutor struct {
	// The number of healthy compute node groups.
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The total number of compute node groups.
	ExpectedCount *int64 `json:"ExpectedCount,omitempty" xml:"ExpectedCount,omitempty"`
	// The number of risky compute node groups.
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The health state of compute node groups. Valid values:
	//
	// *   **RISK**: risky
	// *   **NORMAL**: healthy
	// *   **UNAVAILABLE**: unavailable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of unavailable compute node groups.
	UnavailableCount *int64 `json:"UnavailableCount,omitempty" xml:"UnavailableCount,omitempty"`
}

func (s DescribeDBClusterHealthStatusResponseBodyExecutor) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthStatusResponseBodyExecutor) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) SetActiveCount(v int64) *DescribeDBClusterHealthStatusResponseBodyExecutor {
	s.ActiveCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) SetExpectedCount(v int64) *DescribeDBClusterHealthStatusResponseBodyExecutor {
	s.ExpectedCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) SetRiskCount(v int64) *DescribeDBClusterHealthStatusResponseBodyExecutor {
	s.RiskCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) SetStatus(v string) *DescribeDBClusterHealthStatusResponseBodyExecutor {
	s.Status = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) SetUnavailableCount(v int64) *DescribeDBClusterHealthStatusResponseBodyExecutor {
	s.UnavailableCount = &v
	return s
}

type DescribeDBClusterHealthStatusResponseBodyWorker struct {
	// The number of healthy storage node groups.
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The total number of storage node groups.
	ExpectedCount *int64 `json:"ExpectedCount,omitempty" xml:"ExpectedCount,omitempty"`
	// The number of risky storage node groups.
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The health state of storage node groups. Valid values:
	//
	// *   **RISK**: risky
	// *   **NORMAL**: healthy
	// *   **UNAVAILABLE**: unavailable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of unavailable storage node groups.
	UnavailableCount *int64 `json:"UnavailableCount,omitempty" xml:"UnavailableCount,omitempty"`
}

func (s DescribeDBClusterHealthStatusResponseBodyWorker) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthStatusResponseBodyWorker) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) SetActiveCount(v int64) *DescribeDBClusterHealthStatusResponseBodyWorker {
	s.ActiveCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) SetExpectedCount(v int64) *DescribeDBClusterHealthStatusResponseBodyWorker {
	s.ExpectedCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) SetRiskCount(v int64) *DescribeDBClusterHealthStatusResponseBodyWorker {
	s.RiskCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) SetStatus(v string) *DescribeDBClusterHealthStatusResponseBodyWorker {
	s.Status = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) SetUnavailableCount(v int64) *DescribeDBClusterHealthStatusResponseBodyWorker {
	s.UnavailableCount = &v
	return s
}

type DescribeDBClusterHealthStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBClusterHealthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterHealthStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterHealthStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusResponse) SetHeaders(v map[string]*string) *DescribeDBClusterHealthStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterHealthStatusResponse) SetStatusCode(v int32) *DescribeDBClusterHealthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponse) SetBody(v *DescribeDBClusterHealthStatusResponseBody) *DescribeDBClusterHealthStatusResponse {
	s.Body = v
	return s
}

type DescribeDBClusterNetInfoRequest struct {
	// The ID of the cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterNetInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoRequest) SetDBClusterId(v string) *DescribeDBClusterNetInfoRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) SetOwnerAccount(v string) *DescribeDBClusterNetInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) SetOwnerId(v int64) *DescribeDBClusterNetInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterNetInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterNetInfoRequest) SetResourceOwnerId(v int64) *DescribeDBClusterNetInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBClusterNetInfoResponseBody struct {
	// The network type of the cluster.
	ClusterNetworkType *string `json:"ClusterNetworkType,omitempty" xml:"ClusterNetworkType,omitempty"`
	// The network information of the cluster.
	Items *DescribeDBClusterNetInfoResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterNetInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBody) SetClusterNetworkType(v string) *DescribeDBClusterNetInfoResponseBody {
	s.ClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBody) SetItems(v *DescribeDBClusterNetInfoResponseBodyItems) *DescribeDBClusterNetInfoResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBody) SetRequestId(v string) *DescribeDBClusterNetInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterNetInfoResponseBodyItems struct {
	Address []*DescribeDBClusterNetInfoResponseBodyItemsAddress `json:"Address,omitempty" xml:"Address,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterNetInfoResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBodyItems) SetAddress(v []*DescribeDBClusterNetInfoResponseBodyItemsAddress) *DescribeDBClusterNetInfoResponseBodyItems {
	s.Address = v
	return s
}

type DescribeDBClusterNetInfoResponseBodyItemsAddress struct {
	// The endpoint of the cluster.
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The prefix of the cluster endpoint.
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The IP address.
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The network type of the endpoint. Valid values:
	//
	// *   **Public**: public endpoint
	// *   **VPC**: Virtual Private Cloud (VPC) endpoint
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The port number that is used to connect to the cluster.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the VPC.
	//
	// >  This parameter is empty when Public is returned for NetType.
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch.
	//
	// >  This parameter is empty when Public is returned for NetType.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeDBClusterNetInfoResponseBodyItemsAddress) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBodyItemsAddress) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetConnectionString(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetConnectionStringPrefix(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetIPAddress(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetNetType(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetPort(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetVPCId(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetVSwitchId(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.VSwitchId = &v
	return s
}

type DescribeDBClusterNetInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBClusterNetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterNetInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponse) SetHeaders(v map[string]*string) *DescribeDBClusterNetInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterNetInfoResponse) SetStatusCode(v int32) *DescribeDBClusterNetInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponse) SetBody(v *DescribeDBClusterNetInfoResponseBody) *DescribeDBClusterNetInfoResponse {
	s.Body = v
	return s
}

type DescribeDBClusterPerformanceRequest struct {
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ* format. The time must be in UTC.
	//
	// > The end time must be later than the start time. The maximum time range that can be specified is two days.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance metrics to be queried. Separate multiple values with commas (,). Valid values:
	//
	// *   CPU
	//
	//     *   **AnalyticDB_CPU**: the average CPU utilization.
	//
	// *   Connections
	//
	//     *   **AnalyticDB_Connections**: the number of database connections.
	//
	// *   Writes
	//
	//     *   **AnalyticDB_TPS**: the write transactions per second (TPS).
	//     *   **AnalyticDB_InsertRT**: the write response time.
	//     *   **AnalyticDB_InsertBytes**: the write throughput.
	//
	// *   Updates
	//
	//     *   **AnalyticDB_UpdateRT**: the update response time.
	//
	// *   Deletion
	//
	//     *   **AnalyticDB_DeleteRT**: the delete response time.
	//
	// *   Queries
	//
	//     *   **AnalyticDB_QPS**: the queries per second (QPS).
	//     *   **AnalyticDB_QueryRT**: the query response time.
	//     *   **AnalyticDB_QueryWaitTime**: the query wait time.
	//
	// *   Disks
	//
	//     *   **AnalyticDB_IO**: the disk I/O throughput.
	//     *   **AnalyticDB_IO_UTIL**: the I/O utilization.
	//     *   **AnalyticDB_IO_WAIT**: the I/O wait time.
	//     *   **AnalyticDB_IOPS**: the disk input/output operations per second (IOPS).
	//     *   **AnalyticDB_DiskUsage**: the disk space that is used.
	//     *   **AnalyticDB_HotDataDiskUsage**: the disk space that is used by hot data.
	//     *   **AnalyticDB_ColdDataDiskUsage**: the disk space that is used by cold data.
	//
	// >  If you leave this parameter empty, the values of all the preceding performance metrics are returned.
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the resource group.
	ResourcePools *string `json:"ResourcePools,omitempty" xml:"ResourcePools,omitempty"`
	// The start time of the query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ* format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceRequest) SetDBClusterId(v string) *DescribeDBClusterPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetEndTime(v string) *DescribeDBClusterPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetKey(v string) *DescribeDBClusterPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetOwnerAccount(v string) *DescribeDBClusterPerformanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetOwnerId(v int64) *DescribeDBClusterPerformanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetRegionId(v string) *DescribeDBClusterPerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterPerformanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetResourceOwnerId(v int64) *DescribeDBClusterPerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetResourcePools(v string) *DescribeDBClusterPerformanceRequest {
	s.ResourcePools = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetStartTime(v string) *DescribeDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeDBClusterPerformanceResponseBody struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The queried performance metrics.
	Performances []*DescribeDBClusterPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBody) SetDBClusterId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetEndTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetPerformances(v []*DescribeDBClusterPerformanceResponseBodyPerformances) *DescribeDBClusterPerformanceResponseBody {
	s.Performances = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetRequestId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetStartTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformances struct {
	// The name of the performance metric.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The queried performance metric data.
	Series []*DescribeDBClusterPerformanceResponseBodyPerformancesSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	// The unit of the performance metrics.
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetKey(v string) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetSeries(v []*DescribeDBClusterPerformanceResponseBodyPerformancesSeries) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Series = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetUnit(v string) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Unit = &v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformancesSeries struct {
	// The name of the performance metric value.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tags that are added to the cluster.
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The values of the queried performance metrics.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformancesSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformancesSeries) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetName(v string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetTags(v string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Tags = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetValues(v []*string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Values = v
	return s
}

type DescribeDBClusterPerformanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBClusterPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBClusterPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetStatusCode(v int32) *DescribeDBClusterPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetBody(v *DescribeDBClusterPerformanceResponseBody) *DescribeDBClusterPerformanceResponse {
	s.Body = v
	return s
}

type DescribeDBClusterResourcePoolPerformanceRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to monitor the resource group. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ* format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The metrics of the resource group. You can enter multiple metrics at the same time to query the monitoring information. Separate multiple metrics with commas (,). Valid values:
	//
	// *   **AnalyticDB_RP_CPU**: the average CPU utilization. Unit: %.
	// *   **AnalyticDB_RP_RT**: the query response time (RT). Unit: milliseconds.
	// *   **AnalyticDB_RP_QPS**: the queries per second (QPS). The value of this parameter must be a numeric value.
	// *   **AnalyticDB_RP_WaitTime**: the query waiting time. Unit: milliseconds.
	// *   **AnalyticDB_RP_OriginalNode**: the number of basic nodes in the resource group.
	// *   **AnalyticDB_RP_ActualNode**: the number of scheduled nodes that are scaled out in the resource group.
	// *   **AnalyticDB_RP_PlanNode**: the number of scheduled nodes to be scaled out in the resource group.
	// *   **AnalyticDB_RP_TotalNode**: the total number of nodes in the resource group. Total number of nodes = Number of basic nodes + Number of scheduled nodes that are scaled out.
	//
	// >
	//
	// *   If you leave this parameter empty, the monitoring information about all metrics is returned.
	//
	// *   For more information about scaling plans, see [Create a resource scaling plan](~~189507~~).
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The names of the resource groups that you want to query. You can enter multiple names of resource groups. Separate multiple names with commas (,).
	//
	// >
	//
	// *   The value of this parameter is case-insensitive. For example, `USER_DEFAULT` and `user_default` specify the same resource group.
	//
	// *   If you leave this parameter empty, the monitoring information about the `USER_DEFAULT` resource group is returned.
	ResourcePools *string `json:"ResourcePools,omitempty" xml:"ResourcePools,omitempty"`
	// The beginning of the time range to monitor the resource group. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ* format. The time must be in UTC.
	//
	// > You can view only the monitoring information about the resource groups within the last two days.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterResourcePoolPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetDBClusterId(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetEndTime(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetKey(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetOwnerAccount(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetOwnerId(v int64) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetResourceOwnerId(v int64) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetResourcePools(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.ResourcePools = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetStartTime(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeDBClusterResourcePoolPerformanceResponseBody struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range for monitoring the resource group. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The queried monitoring information about the metrics.
	Performances []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range for monitoring the resource group. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetDBClusterId(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetEndTime(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetPerformances(v []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.Performances = v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetRequestId(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBody) SetStartTime(v string) *DescribeDBClusterResourcePoolPerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances struct {
	// The metric of the resource group.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The queried monitoring information about the resource groups.
	ResourcePoolPerformances []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances `json:"ResourcePoolPerformances,omitempty" xml:"ResourcePoolPerformances,omitempty" type:"Repeated"`
	// The unit of the metric value.
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) SetKey(v string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) SetResourcePoolPerformances(v []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances {
	s.ResourcePoolPerformances = v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances) SetUnit(v string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformances {
	s.Unit = &v
	return s
}

type DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances struct {
	// The name of the resource group.
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
	// The sequential monitoring information about the resource groups.
	ResourcePoolSeries []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries `json:"ResourcePoolSeries,omitempty" xml:"ResourcePoolSeries,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) SetResourcePoolName(v string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances {
	s.ResourcePoolName = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances) SetResourcePoolSeries(v []*DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformances {
	s.ResourcePoolSeries = v
	return s
}

type DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries struct {
	// The name of the metric.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the metric.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) SetName(v string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries) SetValues(v []*string) *DescribeDBClusterResourcePoolPerformanceResponseBodyPerformancesResourcePoolPerformancesResourcePoolSeries {
	s.Values = v
	return s
}

type DescribeDBClusterResourcePoolPerformanceResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBClusterResourcePoolPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterResourcePoolPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBClusterResourcePoolPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponse) SetStatusCode(v int32) *DescribeDBClusterResourcePoolPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceResponse) SetBody(v *DescribeDBClusterResourcePoolPerformanceResponseBody) *DescribeDBClusterResourcePoolPerformanceResponse {
	s.Body = v
	return s
}

type DescribeDBClusterStatusRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// >  You can call [DescribeRegions](~~143074~~) to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStatusRequest) SetOwnerAccount(v string) *DescribeDBClusterStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterStatusRequest) SetOwnerId(v int64) *DescribeDBClusterStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterStatusRequest) SetRegionId(v string) *DescribeDBClusterStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterStatusRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterStatusRequest) SetResourceOwnerId(v int64) *DescribeDBClusterStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBClusterStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The states of clusters. Valid values:
	//
	// *   **Preparing**: The cluster is being prepared.
	// *   **Creating**: The cluster is being created.
	// *   **Restoring**: The cluster is being restored from a backup.
	// *   **Running**: The cluster is running.
	// *   **Deleting**: The cluster is being deleted.
	// *   **ClassChanging**: The cluster configurations are being changed.
	// *   **NetAddressCreating**: A network connection is being created.
	// *   **NetAddressDeleting**: A network connection is being released.
	// *   **NetAddressModifying**: A network connection is being modified.
	// *   **EngineVersionUpgrading**: The engine version is being updated.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStatusResponseBody) SetRequestId(v string) *DescribeDBClusterStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterStatusResponseBody) SetStatus(v []*string) *DescribeDBClusterStatusResponseBody {
	s.Status = v
	return s
}

type DescribeDBClusterStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBClusterStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStatusResponse) SetHeaders(v map[string]*string) *DescribeDBClusterStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterStatusResponse) SetStatusCode(v int32) *DescribeDBClusterStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterStatusResponse) SetBody(v *DescribeDBClusterStatusResponseBody) *DescribeDBClusterStatusResponse {
	s.Body = v
	return s
}

type DescribeDBClustersRequest struct {
	// The description of the cluster.
	//
	// *   The description cannot start with `http://` or `https://`.
	// *   The description must be 2 to 256 characters in length
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The cluster IDs.
	//
	// > You can specify the ID of one cluster or IDs of more clusters within the preceding region.
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	// The state of the cluster. Valid values:
	//
	// *   **Preparing**: The cluster is being prepared.
	// *   **Creating**: The cluster is being created.
	// *   **Restoring**: The cluster is being restored from a backup.
	// *   **Running**: The cluster is running.
	// *   **Deleting**: The cluster is being deleted.
	// *   **ClassChanging**: The cluster specifications are being changed.
	// *   **NetAddressCreating**: A network connection is being created.
	// *   **NetAddressDeleting**: A network connection is being deleted.
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The version of the cluster. Set the value to **3.0**.
	DBVersion    *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the clusters.
	//
	// > You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags that are added to the cluster.
	Tag []*DescribeDBClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequest) SetDBClusterDescription(v string) *DescribeDBClustersRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterIds(v string) *DescribeDBClustersRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterStatus(v string) *DescribeDBClustersRequest {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBVersion(v string) *DescribeDBClustersRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersRequest) SetOwnerAccount(v string) *DescribeDBClustersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClustersRequest) SetOwnerId(v int64) *DescribeDBClustersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageNumber(v int32) *DescribeDBClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageSize(v int32) *DescribeDBClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersRequest) SetRegionId(v string) *DescribeDBClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceGroupId(v string) *DescribeDBClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceOwnerAccount(v string) *DescribeDBClustersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceOwnerId(v int64) *DescribeDBClustersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetTag(v []*DescribeDBClustersRequestTag) *DescribeDBClustersRequest {
	s.Tag = v
	return s
}

type DescribeDBClustersRequestTag struct {
	// The key of tag N that is added to the cluster. You can use tags to filter clusters. A tag is a key-value pair. You can specify up to 20 tags in one request. The letter N specifies the sequence number of each key-value pair and must be unique. The values of N must be consecutive integers that start from 1. Each value of `Tag.N.Key` is paired with a value of `Tag.N.Value`.
	//
	// > The tag key can be up to 64 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is added to the cluster. You can use tags to filter clusters. A tag is a key-value pair. You can specify up to 20 tags in one request. The letter N specifies the sequence number of each key-value pair and must be unique. The values of N must be consecutive integers that start from 1. Each value of `Tag.N.Key` is paired with a value of `Tag.N.Value`.
	//
	// > The tag key can be up to 64 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequestTag) SetKey(v string) *DescribeDBClustersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersRequestTag) SetValue(v string) *DescribeDBClustersRequestTag {
	s.Value = &v
	return s
}

type DescribeDBClustersResponseBody struct {
	// The queried clusters.
	Items *DescribeDBClustersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBody) SetItems(v *DescribeDBClustersResponseBodyItems) *DescribeDBClustersResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageNumber(v int32) *DescribeDBClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageSize(v int32) *DescribeDBClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetRequestId(v string) *DescribeDBClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetTotalCount(v int32) *DescribeDBClustersResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDBClustersResponseBodyItems struct {
	DBCluster []*DescribeDBClustersResponseBodyItemsDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItems) SetDBCluster(v []*DescribeDBClustersResponseBodyItemsDBCluster) *DescribeDBClustersResponseBodyItems {
	s.DBCluster = v
	return s
}

type DescribeDBClustersResponseBodyItemsDBCluster struct {
	// The edition of the cluster. Valid values:
	//
	// *   **BASIC**: reserved mode for Basic Edition.
	// *   **CLUSTER**: reserved mode for Cluster Edition.
	// *   **MIXED_STORAGE**: elastic mode for Cluster Edition.
	//
	// >  For more information about cluster editions, see [Editions](~~205001~~).
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The commodity code. **ads** is returned.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The specifications of computing resources that are used in the cluster in elastic mode. The increase of computing resources can speed up queries. You can adjust the value of this parameter to scale the cluster.
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The public endpoint that is used to connect to the cluster.
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The time when the cluster was created. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC. Example: *2021-04-01T09:50:18Z*.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the cluster.
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The network type of the cluster. **VPC** is returned.
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The state of the cluster. For more information, see [Cluster states](~~143075~~).
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The type of the cluster. Valid values:
	//
	// *   **Common**: common cluster.
	// *   **RDS_ANALYSIS**: MySQL analytic instance.
	DBClusterType *string `json:"DBClusterType,omitempty" xml:"DBClusterType,omitempty"`
	// The instance type of the cluster.
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of node groups.
	DBNodeCount *int64 `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	// The storage capacity of the cluster. Unit: GB.
	DBNodeStorage *int64 `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	// The version of the database engine. **3.0** is returned.
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The disk type of the cluster. Valid values:
	//
	// *   **local_ssd**: local disk.
	// *   **cloud**: basic disk.
	// *   **cloud_ssd**: standard SSD.
	// *   **cloud_efficiency**: ultra disk.
	// *   **cloud_essd**: PL0 enhanced SSD (ESSD).
	// *   **cloud_essd**: PL1 ESSD.
	// *   **cloud_essd2**: PL2 ESSD.
	// *   **cloud_essd3**: PL3 ESSD.
	//
	// >  For more information, see [ESSDs](~~122389~~).
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the Data Transmission Service (DTS) synchronization task. This parameter is returned only for MySQL analytic instances.
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The number of elastic I/O units (EIUs). For more information, see [Use EIUs to scale up storage resources](~~189505~~).
	//
	// >  This parameter is returned only for clusters in elastic mode.
	ElasticIOResource *int32 `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	// The engine of the cluster. **AnalyticDB** is returned.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The number of compute nodes that are used by the cluster in elastic mode.
	ExecutorCount *string `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	// The time when the cluster expires. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC. Example: *2999-09-08T16:00:00Z*.
	//
	// >
	//
	// *   If the billing method of the cluster is subscription, the actual expiration time is returned.
	//
	// *   If the billing method of the cluster is pay-as-you-go, **2999-09-08T16:00:00Z** is returned.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the cluster has expired. Valid values:
	//
	// *   **true**
	// *   **false**
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The public IP address of the cluster.
	InnerIp *string `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	// The port number that is used to connect to the cluster.
	InnerPort *string `json:"InnerPort,omitempty" xml:"InnerPort,omitempty"`
	// The lock mode of the cluster. Valid values:
	//
	// *   **Unlock**: The cluster is not locked.
	// *   **ManualLock**: The cluster is manually locked.
	// *   **LockByExpiration**: The cluster is automatically locked due to cluster expiration.
	// *   **LockByRestoration**: The cluster is automatically locked due to cluster restoration.
	// *   **LockByDiskQuota**: The cluster is automatically locked when it has used 90% of its storage.
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the cluster is locked.
	//
	// >  This parameter is returned only when the cluster was locked. **instance_expire** is returned.
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The mode of the cluster. Valid values:
	//
	// *   **flexible**: elastic mode.
	// *   **reserver**: reserved mode.
	//
	// >
	//
	// *   For more information about cluster modes, see [Editions](~~205001~~).
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// *   **Postpaid**: pay-as-you-go.
	// *   **Prepaid**: subscription.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port number that is used to connect to the cluster. Default value: 3306.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the ApsaraDB RDS instance from which data is synchronized to the cluster. This parameter is returned only for MySQL analytic instances.
	RdsInstanceId *string `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	// The region ID of the cluster.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The specifications of storage resources that are used in the cluster in elastic mode. These resources are used to read and write data. You can increase the value of this parameter to improve the read and write performance of the cluster.
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	// The tags that are added to the cluster.
	Tags *DescribeDBClustersResponseBodyItemsDBClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the cluster that is deployed in the VPC.
	VPCCloudInstanceId *string `json:"VPCCloudInstanceId,omitempty" xml:"VPCCloudInstanceId,omitempty"`
	// The virtual private cloud (VPC) ID of the cluster.
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the cluster.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the cluster.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClustersResponseBodyItemsDBCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCategory(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Category = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCommodityCode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetComputeResource(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ComputeResource = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetConnectionString(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCreateTime(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeClass(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeCount(v int64) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeStorage(v int64) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeStorage = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDiskType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DiskType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDtsJobId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetElasticIOResource(v int32) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ElasticIOResource = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetEngine(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Engine = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExecutorCount(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ExecutorCount = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExpired(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetInnerIp(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.InnerIp = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetInnerPort(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.InnerPort = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetLockReason(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.LockReason = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetMode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Mode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetPayType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetPort(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Port = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetRdsInstanceId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.RdsInstanceId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetResourceGroupId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetStorageResource(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.StorageResource = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetTags(v *DescribeDBClustersResponseBodyItemsDBClusterTags) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVPCCloudInstanceId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VPCCloudInstanceId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVPCId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVSwitchId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

type DescribeDBClustersResponseBodyItemsDBClusterTags struct {
	Tag []*DescribeDBClustersResponseBodyItemsDBClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTags) SetTag(v []*DescribeDBClustersResponseBodyItemsDBClusterTagsTag) *DescribeDBClustersResponseBodyItemsDBClusterTags {
	s.Tag = v
	return s
}

type DescribeDBClustersResponseBodyItemsDBClusterTagsTag struct {
	// The tag key.
	//
	// >  You can call the [TagResources](~~179253~~) operation to add tags to a cluster.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTagsTag) SetKey(v string) *DescribeDBClustersResponseBodyItemsDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTagsTag) SetValue(v string) *DescribeDBClustersResponseBodyItemsDBClusterTagsTag {
	s.Value = &v
	return s
}

type DescribeDBClustersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponse) SetHeaders(v map[string]*string) *DescribeDBClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClustersResponse) SetStatusCode(v int32) *DescribeDBClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClustersResponse) SetBody(v *DescribeDBClustersResponseBody) *DescribeDBClustersResponse {
	s.Body = v
	return s
}

type DescribeDBResourceGroupRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupRequest) SetDBClusterId(v string) *DescribeDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetGroupName(v string) *DescribeDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetOwnerAccount(v string) *DescribeDBResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetOwnerId(v int64) *DescribeDBResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetResourceOwnerAccount(v string) *DescribeDBResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetResourceOwnerId(v int64) *DescribeDBResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBResourceGroupResponseBody struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried resource group.
	GroupsInfo []*DescribeDBResourceGroupResponseBodyGroupsInfo `json:"GroupsInfo,omitempty" xml:"GroupsInfo,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBody) SetDBClusterId(v string) *DescribeDBResourceGroupResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBody) SetGroupsInfo(v []*DescribeDBResourceGroupResponseBodyGroupsInfo) *DescribeDBResourceGroupResponseBody {
	s.GroupsInfo = v
	return s
}

func (s *DescribeDBResourceGroupResponseBody) SetRequestId(v string) *DescribeDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBResourceGroupResponseBodyGroupsInfo struct {
	// The time when the resource group was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the resource group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The query execution mode. Valid values:
	//
	// *   **interactive**
	// *   **batch** (default)
	//
	// > For more information, see [Query execution modes](~~189502~~).
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The database accounts that are associated with the resource group.
	GroupUsers *string `json:"GroupUsers,omitempty" xml:"GroupUsers,omitempty"`
	// The number of nodes. Each node provides 16 cores and 64 GB memory.
	NodeNum *int32 `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// The time when the resource group was updated.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetCreateTime(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetGroupName(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetGroupType(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.GroupType = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetGroupUsers(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.GroupUsers = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetNodeNum(v int32) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.NodeNum = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetUpdateTime(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.UpdateTime = &v
	return s
}

type DescribeDBResourceGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponse) SetHeaders(v map[string]*string) *DescribeDBResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBResourceGroupResponse) SetStatusCode(v int32) *DescribeDBResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBResourceGroupResponse) SetBody(v *DescribeDBResourceGroupResponseBody) *DescribeDBResourceGroupResponse {
	s.Body = v
	return s
}

type DescribeDBResourcePoolRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the resource group.
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBResourcePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBResourcePoolRequest) SetDBClusterId(v string) *DescribeDBResourcePoolRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetOwnerAccount(v string) *DescribeDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetOwnerId(v int64) *DescribeDBResourcePoolRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetPoolName(v string) *DescribeDBResourcePoolRequest {
	s.PoolName = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetResourceOwnerAccount(v string) *DescribeDBResourcePoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetResourceOwnerId(v int64) *DescribeDBResourcePoolRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBResourcePoolResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Details of the resource group.
	PoolsInfo []*DescribeDBResourcePoolResponseBodyPoolsInfo `json:"PoolsInfo,omitempty" xml:"PoolsInfo,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBResourcePoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourcePoolResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBResourcePoolResponseBody) SetDBClusterId(v string) *DescribeDBResourcePoolResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBody) SetPoolsInfo(v []*DescribeDBResourcePoolResponseBodyPoolsInfo) *DescribeDBResourcePoolResponseBody {
	s.PoolsInfo = v
	return s
}

func (s *DescribeDBResourcePoolResponseBody) SetRequestId(v string) *DescribeDBResourcePoolResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBResourcePoolResponseBodyPoolsInfo struct {
	// The time when the resource group was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of nodes.
	//
	// >  Each node consumes 16 cores and 64 GB memory.
	NodeNum *int32 `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// The name of the resource group.
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The database accounts that are associated with the resource group.
	PoolUsers *string `json:"PoolUsers,omitempty" xml:"PoolUsers,omitempty"`
	// The mode in which SQL statements are executed.
	//
	// *   **batch**
	// *   **interactive**
	//
	// >  For more information, see [Query execution modes](~~189502~~).
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The time when the resource group was updated.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDBResourcePoolResponseBodyPoolsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourcePoolResponseBodyPoolsInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetCreateTime(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetNodeNum(v int32) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.NodeNum = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetPoolName(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.PoolName = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetPoolUsers(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.PoolUsers = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetQueryType(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.QueryType = &v
	return s
}

func (s *DescribeDBResourcePoolResponseBodyPoolsInfo) SetUpdateTime(v string) *DescribeDBResourcePoolResponseBodyPoolsInfo {
	s.UpdateTime = &v
	return s
}

type DescribeDBResourcePoolResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBResourcePoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBResourcePoolResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBResourcePoolResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBResourcePoolResponse) SetHeaders(v map[string]*string) *DescribeDBResourcePoolResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBResourcePoolResponse) SetStatusCode(v int32) *DescribeDBResourcePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBResourcePoolResponse) SetBody(v *DescribeDBResourcePoolResponseBody) *DescribeDBResourcePoolResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisDimensionsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >
	//
	// *   The end time must be later than the start time.
	//
	// *   The maximum time range that can be specified is 24 hours.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of file titles and error messages. Valid values:
	//
	// *   **zh** (default): simplified Chinese.
	// *   **en**: English.
	// *   **ja**: Japanese.
	// *   **zh-tw**: traditional Chinese.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The query condition for SQL statements, which can contain the `Type`, `Value`, and `Min` or `Max` fields. Specify the condition in the JSON format. `Type` specifies the query dimension. Valid values for Type: `maxCost`, `status`, and `cost`. `Value`, `Min`, or `Max` specifies the query range for the dimension. Valid values:
	//
	// *   `{"Type":"maxCost","Value":"100"}`: queries the top 100 most time-consuming SQL statements. Set `Value` to 100.
	// *   `{"Type":"status","Value":"finished"}`: queries executed SQL statements. You can set `Value` to `running` to query SQL statements that are being executed. You can also set Value to `failed` to query SQL statements that failed to be executed.
	// *   `{"Type":"cost","Min":"10","Max":"200"}`: queries SQL statements whose execution durations are in the range of 10 to 200 milliseconds. You can also customize the maximum and minimum execution durations.
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// > Only data within the last 14 days can be queried. If you call this operation to query data that is earlier than 14 days, an empty string is returned.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDiagnosisDimensionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisDimensionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsRequest) SetDBClusterId(v string) *DescribeDiagnosisDimensionsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) SetEndTime(v string) *DescribeDiagnosisDimensionsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) SetLang(v string) *DescribeDiagnosisDimensionsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) SetQueryCondition(v string) *DescribeDiagnosisDimensionsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) SetRegionId(v string) *DescribeDiagnosisDimensionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) SetStartTime(v string) *DescribeDiagnosisDimensionsRequest {
	s.StartTime = &v
	return s
}

type DescribeDiagnosisDimensionsResponseBody struct {
	// The source IP addresses.
	ClientIps []*string `json:"ClientIps,omitempty" xml:"ClientIps,omitempty" type:"Repeated"`
	// The databases.
	Databases []*string `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource groups.
	ResourceGroups []*string `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
	// The usernames.
	UserNames []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
}

func (s DescribeDiagnosisDimensionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisDimensionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetClientIps(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.ClientIps = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetDatabases(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetRequestId(v string) *DescribeDiagnosisDimensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetResourceGroups(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.ResourceGroups = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetUserNames(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.UserNames = v
	return s
}

type DescribeDiagnosisDimensionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiagnosisDimensionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisDimensionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisDimensionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisDimensionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponse) SetStatusCode(v int32) *DescribeDiagnosisDimensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisDimensionsResponse) SetBody(v *DescribeDiagnosisDimensionsResponseBody) *DescribeDiagnosisDimensionsResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisMonitorPerformanceRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the UNIX timestamp format. Unit: milliseconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of file titles and error messages. Default value: zh. Valid values:
	//
	// *   **zh**: simplified Chinese
	// *   **en**: English
	// *   **ja**: Japanese
	// *   **zh-tw**: traditional Chinese
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The query conditions for SQL statements, which can be a combination of the `Type` and `Value` fields or a combination of the Type, `Min`, and `Max` fields. Specify the conditions in the JSON format. `Type` specifies the query dimension. Valid values for Type: `maxCost`, `status`, and `cost`. `Value`, `Min`, or `Max` specifies the query range for the dimension. Valid values:
	//
	// *   `{"Type":"maxCost","Value":"100"}`: queries the top 100 most time-consuming SQL statements. Set `Value` to 100.
	// *   `{"Type":"status","Value":"finished"}`: queries executed SQL statements. You can set `Value` to `running` to query SQL statements that are being executed. You can also set Value to `failed` to query SQL statements that failed to be executed.
	// *   `{"Type":"cost","Min":"10","Max":"200"}`: queries SQL statements whose execution durations are in the range of 10 to 200 milliseconds. You can also customize the maximum and minimum execution durations.
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the UNIX timestamp format. Unit: milliseconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDiagnosisMonitorPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetDBClusterId(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetEndTime(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetLang(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetQueryCondition(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetRegionId(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceRequest) SetStartTime(v string) *DescribeDiagnosisMonitorPerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeDiagnosisMonitorPerformanceResponseBody struct {
	// The monitoring information of queries displayed in Gantt charts.
	Performances []*DescribeDiagnosisMonitorPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The threshold for the number of queries displayed in a Gantt chart. The default value is 10000.
	//
	// >  A maximum of 10,000 queries can be displayed in a Gantt chart even if more queries exist.
	PerformancesThreshold *int32 `json:"PerformancesThreshold,omitempty" xml:"PerformancesThreshold,omitempty"`
	// Indicates whether all queries are returned. Valid values:
	//
	// *   true: All queries are returned.
	// *   false: Only a specified number of queries are returned.
	PerformancesTruncated *bool `json:"PerformancesTruncated,omitempty" xml:"PerformancesTruncated,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiagnosisMonitorPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetPerformances(v []*DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.Performances = v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetPerformancesThreshold(v int32) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.PerformancesThreshold = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetPerformancesTruncated(v bool) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.PerformancesTruncated = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBody) SetRequestId(v string) *DescribeDiagnosisMonitorPerformanceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDiagnosisMonitorPerformanceResponseBodyPerformances struct {
	// The total amount of time consumed by the query. Unit: milliseconds.
	//
	// >  This parameter indicates the sum of `QueuedTime`, `TotalPlanningTime`, and `ExecutionTime`.
	Cost *int64 `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// The peak memory of the query. Unit: bytes.
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The ID of the query.
	//
	// >  You can call the [DescribeProcessList](~~143382~~) operation to query the IDs of queries that are being executed.
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The IP address of the AnalyticDB for MySQL frontend node on which the SQL statement is executed.
	RcHost *string `json:"RcHost,omitempty" xml:"RcHost,omitempty"`
	// The number of entries scanned.
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// The amount of scanned data. Unit: bytes.
	ScanSize *int64 `json:"ScanSize,omitempty" xml:"ScanSize,omitempty"`
	// The execution start time of the SQL statement. The time is in the UNIX timestamp format. Unit: milliseconds.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the SQL statement. Valid values:
	//
	// *   **running**
	// *   **finished**
	// *   **failed**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The database account that is used to submit the query.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetCost(v int64) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.Cost = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetPeakMemory(v int64) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.PeakMemory = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetProcessId(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.ProcessId = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetRcHost(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.RcHost = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetScanRows(v int64) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.ScanRows = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetScanSize(v int64) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.ScanSize = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetStartTime(v int64) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetStatus(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances) SetUserName(v string) *DescribeDiagnosisMonitorPerformanceResponseBodyPerformances {
	s.UserName = &v
	return s
}

type DescribeDiagnosisMonitorPerformanceResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiagnosisMonitorPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisMonitorPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisMonitorPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) SetStatusCode(v int32) *DescribeDiagnosisMonitorPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) SetBody(v *DescribeDiagnosisMonitorPerformanceResponseBody) *DescribeDiagnosisMonitorPerformanceResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisRecordsRequest struct {
	// The source IP address.
	//
	// > You can call the [DescribeDiagnosisDimensions](~~308210~~) operation to query the resource group, database name, username, and source IP address of the SQL statements to be queried.
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database on which the SQL statements are executed.
	//
	// > You can call the [DescribeDiagnosisDimensions](~~308210~~) operation to query the resource group, database name, username, and source IP address of the SQL statements to be queried.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >
	//
	// *   The end time must be later than the start time.
	//
	// *   The maximum time range that can be specified is 24 hours.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword for the query.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of file titles and error messages. Valid values:
	//
	// *   **zh** (default): simplified Chinese.
	// *   **en**: English.
	// *   **ja**: Japanese.
	// *   **zh-tw**: traditional Chinese.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum peak memory of the SQL statements. Unit: bytes.
	MaxPeakMemory *int64 `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	// The maximum scan size of the SQL statements. Unit: bytes.
	MaxScanSize *int64 `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	// The minimum peak memory of the SQL statements. Unit: bytes.
	MinPeakMemory *int64 `json:"MinPeakMemory,omitempty" xml:"MinPeakMemory,omitempty"`
	// The minimum scan size of the SQL statements. Unit: bytes.
	MinScanSize *int64 `json:"MinScanSize,omitempty" xml:"MinScanSize,omitempty"`
	// The order in which to sort the retrieved SQL statements by field. Specify this value in the JSON format. The value is an ordered array that uses the order of the input array and contains the `Field` and `Type` fields. Example: `[{"Field":"StartTime", "Type": "desc" }]`. Fields:
	//
	// *   `Field` specifies the field that is used to sort the retrieved SQL statements. Valid values:
	//
	//     *   `StartTime`: the start time of the execution.
	//     *   `Status`: the execution state.
	//     *   `UserName`: the username.
	//     *   `Cost`: the execution duration.
	//     *   `PeakMemory`: the peak memory.
	//     *   `ScanSize`: the amount of data to be scanned.
	//     *   `Database`: the name of the database.
	//     *   `ClientIp`: the source IP address.
	//     *   `ResourceGroup`: the name of the resource group.
	//     *   `QueueTime`: the amount of time that is consumed for queuing.
	//     *   `OutputRows`: the number of output rows.
	//     *   `OutputDataSize`: the amount of output data.
	//     *   `ResourceCostRank`: the execution duration rank of operators that are used in the SQL statements. This field takes effect only when `QueryCondition` is set to `{"Type":"status","Value":"running"}`.
	//
	// *   `Type` specifies the sorting order. Valid values (case-insensitive):
	//
	//     *   `Desc`: descending order.
	//     *   `Asc`: ascending order.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30**, **50**, and **100**. Default value: 30.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the SQL pattern.[](~~321868~~)
	PatternId *string `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	// The query condition for SQL statements, which can contain the `Type`, `Value`, and `Min` or `Max` fields. Specify the condition in the JSON format. `Type` specifies the query dimension. Valid values for Type: `maxCost`, `status`, and `cost`. `Value`, `Min`, or `Max` specifies the query range for the dimension. Valid values:
	//
	// *   `{"Type":"maxCost","Value":"100"}`: queries the top 100 most time-consuming SQL statements. Set `Value` to 100.
	// *   `{"Type":"status","Value":"finished"}`: queries executed SQL statements. You can set `Value` to `running` to query SQL statements that are being executed. You can also set Value to `failed` to query SQL statements that failed to be executed.
	// *   `{"Type":"cost","Min":"10","Max":"200"}`: queries SQL statements whose execution durations are in the range of 10 to 200 milliseconds. You can also customize the maximum and minimum execution durations.
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group to which the SQL statements belong.
	//
	// > You can call the [DescribeDiagnosisDimensions](~~308210~~) operation to query the resource group, database name, username, and source IP address of the SQL statements to be queried.
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The beginning of the time range to query. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// > Only data within the last 14 days can be queried.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The username that is used to execute the SQL statements.
	//
	// > You can call the [DescribeDiagnosisDimensions](~~308210~~) operation to query the resource group, database name, username, and source IP address of the SQL statements to be queried.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDiagnosisRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsRequest) SetClientIp(v string) *DescribeDiagnosisRecordsRequest {
	s.ClientIp = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetDBClusterId(v string) *DescribeDiagnosisRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetDatabase(v string) *DescribeDiagnosisRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetEndTime(v string) *DescribeDiagnosisRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetKeyword(v string) *DescribeDiagnosisRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetLang(v string) *DescribeDiagnosisRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMaxPeakMemory(v int64) *DescribeDiagnosisRecordsRequest {
	s.MaxPeakMemory = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMaxScanSize(v int64) *DescribeDiagnosisRecordsRequest {
	s.MaxScanSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMinPeakMemory(v int64) *DescribeDiagnosisRecordsRequest {
	s.MinPeakMemory = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetMinScanSize(v int64) *DescribeDiagnosisRecordsRequest {
	s.MinScanSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetOrder(v string) *DescribeDiagnosisRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetPageNumber(v int32) *DescribeDiagnosisRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetPageSize(v int32) *DescribeDiagnosisRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetPatternId(v string) *DescribeDiagnosisRecordsRequest {
	s.PatternId = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetQueryCondition(v string) *DescribeDiagnosisRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetRegionId(v string) *DescribeDiagnosisRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetResourceGroup(v string) *DescribeDiagnosisRecordsRequest {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetStartTime(v string) *DescribeDiagnosisRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsRequest) SetUserName(v string) *DescribeDiagnosisRecordsRequest {
	s.UserName = &v
	return s
}

type DescribeDiagnosisRecordsResponseBody struct {
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The queried SQL statements.
	Querys []*DescribeDiagnosisRecordsResponseBodyQuerys `json:"Querys,omitempty" xml:"Querys,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiagnosisRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponseBody) SetPageNumber(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetPageSize(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetQuerys(v []*DescribeDiagnosisRecordsResponseBodyQuerys) *DescribeDiagnosisRecordsResponseBody {
	s.Querys = v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetRequestId(v string) *DescribeDiagnosisRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetTotalCount(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDiagnosisRecordsResponseBodyQuerys struct {
	// The source IP address.
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The total execution duration. Unit: milliseconds.
	//
	// >  This value is the cumulative value of the `QueuedTime`, `TotalPlanningTime`, and `ExecutionTime` parameters.
	Cost *int64 `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// The name of the database on which the SQL statement is executed.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The number of rows written to the table by an extract, transform, and load (ETL) task.
	EtlWriteRows *int64 `json:"EtlWriteRows,omitempty" xml:"EtlWriteRows,omitempty"`
	// The execution duration. Unit: milliseconds.
	ExecutionTime *int64 `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// The amount of returned data. Unit: bytes.
	OutputDataSize *int64 `json:"OutputDataSize,omitempty" xml:"OutputDataSize,omitempty"`
	// The number of rows returned.
	OutputRows *int64 `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	// The peak memory. Unit: bytes.
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The query ID.
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The amount of time that is consumed for queuing. Unit: milliseconds.
	QueueTime *int64 `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	// The IP address and port number of the AnalyticDB for MySQL frontend node on which the SQL statement is executed.
	RcHost *string `json:"RcHost,omitempty" xml:"RcHost,omitempty"`
	// The execution duration rank of operators that are used in the SQL statement.
	//
	// > This field is returned only for SQL statements that have the `Status` parameter set to `running`.
	ResourceCostRank *int32 `json:"ResourceCostRank,omitempty" xml:"ResourceCostRank,omitempty"`
	// The resource group to which the SQL statement belongs.
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The SQL statement.
	//
	// > For performance considerations, an SQL statement cannot exceed 5,120 characters in length. Otherwise, the SQL statement is truncated. You can call the [DownloadDiagnosisRecords](~~308212~~) operation to download the diagnostic information about SQL statements that meet a condition in an AnalyticDB for MySQL cluster, including the complete SQL statements.
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	// Indicates whether the SQL statement is truncated. Valid values:
	//
	// *   **true**
	// *   **false**
	SQLTruncated *bool `json:"SQLTruncated,omitempty" xml:"SQLTruncated,omitempty"`
	// The maximum length of the SQL statement. 5120 is returned. Unit: character. SQL statements that exceed this limit are truncated.
	SQLTruncatedThreshold *int64 `json:"SQLTruncatedThreshold,omitempty" xml:"SQLTruncatedThreshold,omitempty"`
	// The number of entries scanned.
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// The amount of scanned data. Unit: bytes.
	ScanSize *int64 `json:"ScanSize,omitempty" xml:"ScanSize,omitempty"`
	// The beginning of the time range in which the SQL statement is executed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the SQL statement. Valid values:
	//
	// *   **running**
	// *   **finished**
	// *   **failed**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The amount of time that is consumed to generate an execution plan. Unit: milliseconds.
	TotalPlanningTime *int64 `json:"TotalPlanningTime,omitempty" xml:"TotalPlanningTime,omitempty"`
	// The total number of stages generated.
	TotalStages *int32 `json:"TotalStages,omitempty" xml:"TotalStages,omitempty"`
	// The username that is used to execute the SQL statement.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDiagnosisRecordsResponseBodyQuerys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponseBodyQuerys) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetClientIp(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ClientIp = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetCost(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.Cost = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetDatabase(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetEtlWriteRows(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.EtlWriteRows = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetExecutionTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ExecutionTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetOutputDataSize(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.OutputDataSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetOutputRows(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.OutputRows = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetPeakMemory(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.PeakMemory = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetProcessId(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ProcessId = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetQueueTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.QueueTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetRcHost(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.RcHost = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetResourceCostRank(v int32) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ResourceCostRank = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetResourceGroup(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetSQL(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.SQL = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetSQLTruncated(v bool) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.SQLTruncated = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetSQLTruncatedThreshold(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.SQLTruncatedThreshold = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetScanRows(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ScanRows = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetScanSize(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.ScanSize = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetStartTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetStatus(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetTotalPlanningTime(v int64) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.TotalPlanningTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetTotalStages(v int32) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.TotalStages = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyQuerys) SetUserName(v string) *DescribeDiagnosisRecordsResponseBodyQuerys {
	s.UserName = &v
	return s
}

type DescribeDiagnosisRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiagnosisRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisRecordsResponse) SetStatusCode(v int32) *DescribeDiagnosisRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponse) SetBody(v *DescribeDiagnosisRecordsResponseBody) *DescribeDiagnosisRecordsResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisSQLInfoRequest struct {
	// The ID of the cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a specific region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The language of file titles and error messages. Valid values:
	//
	// *   **zh**: simplified Chinese
	// *   **en**: English
	// *   **ja**: Japanese
	// *   **zh-tw**: traditional Chinese
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the query.
	//
	// >  You can call the [DescribeDiagnosisRecords](~~308207~~) operation to query the SQL summary information of a specified AnalyticDB for MySQL cluster, including the query ID.
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The IP address and port number of the AnalyticDB for MySQL frontend node on which the SQL statement is executed.
	//
	// >  You can call the [DescribeDiagnosisRecords](~~308207~~) operation to query the SQL summary information of a specified AnalyticDB for MySQL cluster, including the IP address and port number of the frontend node.
	ProcessRcHost *string `json:"ProcessRcHost,omitempty" xml:"ProcessRcHost,omitempty"`
	// The execution start time of the SQL statement. Specify the time in the UNIX timestamp format. Unit: milliseconds.
	//
	// >  You can call the [DescribeDiagnosisRecords](~~308207~~) operation to query the SQL summary information of a specified AnalyticDB for MySQL cluster, including the execution start time of the SQL statement.
	ProcessStartTime *int64 `json:"ProcessStartTime,omitempty" xml:"ProcessStartTime,omitempty"`
	// The state of the SQL statement. Valid values:
	//
	// *   **running**
	//
	// *   **finished**
	//
	// *   **failed**
	//
	// > You can call the [DescribeDiagnosisRecords](~~308207~~) operation to query the SQL summary information of a specified AnalyticDB for MySQL cluster, including the state of the SQL statement.
	ProcessState *string `json:"ProcessState,omitempty" xml:"ProcessState,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDiagnosisSQLInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoRequest) SetDBClusterId(v string) *DescribeDiagnosisSQLInfoRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetLang(v string) *DescribeDiagnosisSQLInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetProcessId(v string) *DescribeDiagnosisSQLInfoRequest {
	s.ProcessId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetProcessRcHost(v string) *DescribeDiagnosisSQLInfoRequest {
	s.ProcessRcHost = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetProcessStartTime(v int64) *DescribeDiagnosisSQLInfoRequest {
	s.ProcessStartTime = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetProcessState(v string) *DescribeDiagnosisSQLInfoRequest {
	s.ProcessState = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetRegionId(v string) *DescribeDiagnosisSQLInfoRequest {
	s.RegionId = &v
	return s
}

type DescribeDiagnosisSQLInfoResponseBody struct {
	// Execution details of the SQL statement, including the SQL statement text, statistics, execution plan, and operator information.
	DiagnosisSQLInfo *string `json:"DiagnosisSQLInfo,omitempty" xml:"DiagnosisSQLInfo,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Execution details of the query by stage.
	StageInfos []*DescribeDiagnosisSQLInfoResponseBodyStageInfos `json:"StageInfos,omitempty" xml:"StageInfos,omitempty" type:"Repeated"`
}

func (s DescribeDiagnosisSQLInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetDiagnosisSQLInfo(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.DiagnosisSQLInfo = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetRequestId(v string) *DescribeDiagnosisSQLInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBody) SetStageInfos(v []*DescribeDiagnosisSQLInfoResponseBodyStageInfos) *DescribeDiagnosisSQLInfoResponseBody {
	s.StageInfos = v
	return s
}

type DescribeDiagnosisSQLInfoResponseBodyStageInfos struct {
	// The total amount of input data in the stage. Unit: bytes.
	InputDataSize *int64 `json:"InputDataSize,omitempty" xml:"InputDataSize,omitempty"`
	// The total number of input rows in the stage.
	InputRows *int64 `json:"InputRows,omitempty" xml:"InputRows,omitempty"`
	// The total amount of time consumed by all operators in the stage. Unit: milliseconds.
	OperatorCost *int64 `json:"OperatorCost,omitempty" xml:"OperatorCost,omitempty"`
	// The total amount of output data in the stage. Unit: bytes.
	OutputDataSize *int64 `json:"OutputDataSize,omitempty" xml:"OutputDataSize,omitempty"`
	// The total number of output rows in the stage.
	OutputRows *int64 `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	// The total peak memory of the stage. Unit: bytes.
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The execution progress of the stage.
	Progress *float64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the stage.
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The state of the stage.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeDiagnosisSQLInfoResponseBodyStageInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoResponseBodyStageInfos) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetInputDataSize(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.InputDataSize = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetInputRows(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.InputRows = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetOperatorCost(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.OperatorCost = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetOutputDataSize(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.OutputDataSize = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetOutputRows(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.OutputRows = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetPeakMemory(v int64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.PeakMemory = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetProgress(v float64) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.Progress = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetStageId(v string) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.StageId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponseBodyStageInfos) SetState(v string) *DescribeDiagnosisSQLInfoResponseBodyStageInfos {
	s.State = &v
	return s
}

type DescribeDiagnosisSQLInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiagnosisSQLInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisSQLInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisSQLInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponse) SetStatusCode(v int32) *DescribeDiagnosisSQLInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoResponse) SetBody(v *DescribeDiagnosisSQLInfoResponseBody) *DescribeDiagnosisSQLInfoResponse {
	s.Body = v
	return s
}

type DescribeDiagnosisTasksRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The IP address from which the query was initiated.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The order in which to sort the tasks by field. Specify the value in the JSON format. Example: `[{"Field":"CreateTime", "Type":"desc"}]`.
	//
	// >
	//
	// *   `Field` indicates the field that is used to sort the tasks. Valid values of Field: `State`, `CreateTime`, `DBName`, `ProcessID`, `UpdateTime`, `JobName`, and `ProcessRows`.
	//
	// *   `Type` indicates the sort type. Valid values of Type: `Desc` and `Asc`. The values are case-insensitive.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   30 (default)
	// *   50
	// *   100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query ID.
	//
	// > You can call the [DescribeProcessList](~~190092~~) operation to query the IDs of queries that are being executed.
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of a stage in the query that is specified by the `ProcessId` parameter.
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The state of the asynchronous import or export task to be queried. Valid values:
	//
	// *   **RUNNING**
	// *   **FINISHED**
	// *   **FAILED**
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeDiagnosisTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisTasksRequest) SetDBClusterId(v string) *DescribeDiagnosisTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetHost(v string) *DescribeDiagnosisTasksRequest {
	s.Host = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetOrder(v string) *DescribeDiagnosisTasksRequest {
	s.Order = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetPageNumber(v int32) *DescribeDiagnosisTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetPageSize(v int32) *DescribeDiagnosisTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetProcessId(v string) *DescribeDiagnosisTasksRequest {
	s.ProcessId = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetRegionId(v string) *DescribeDiagnosisTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetStageId(v string) *DescribeDiagnosisTasksRequest {
	s.StageId = &v
	return s
}

func (s *DescribeDiagnosisTasksRequest) SetState(v string) *DescribeDiagnosisTasksRequest {
	s.State = &v
	return s
}

type DescribeDiagnosisTasksResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried tasks.
	TaskList []*DescribeDiagnosisTasksResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
	// The total number of tasks in the stage.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiagnosisTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisTasksResponseBody) SetRequestId(v string) *DescribeDiagnosisTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBody) SetTaskList(v []*DescribeDiagnosisTasksResponseBodyTaskList) *DescribeDiagnosisTasksResponseBody {
	s.TaskList = v
	return s
}

func (s *DescribeDiagnosisTasksResponseBody) SetTotalCount(v int32) *DescribeDiagnosisTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDiagnosisTasksResponseBodyTaskList struct {
	// The compute time ratio, which can be used to determine whether the task is really time-consuming. This parameter can be calculated by using the following formula: OperatorCost/Drivers/ElapsedTime. A greater value indicates that the task was executed for computing for most of the task time. A less value indicates that the task was waiting for scheduling or blocked due to other reasons for most of the task time.
	ComputeTimeRatio *string `json:"ComputeTimeRatio,omitempty" xml:"ComputeTimeRatio,omitempty"`
	// The number of tasks that can be executed concurrently.
	Drivers *string `json:"Drivers,omitempty" xml:"Drivers,omitempty"`
	// The amount of time that elapsed from when the task was created to when the task was completed. Unit: milliseconds.
	ElapsedTime *int64 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// The amount of input data in the task. Unit: bytes.
	InputDataSize *int64 `json:"InputDataSize,omitempty" xml:"InputDataSize,omitempty"`
	// The number of input rows in the task.
	InputRows *int64 `json:"InputRows,omitempty" xml:"InputRows,omitempty"`
	// The total amount of time that is consumed by all operators in the task on a node. This parameter can be used to determine whether long tails occur in computing. Unit: milliseconds.
	OperatorCost *int64 `json:"OperatorCost,omitempty" xml:"OperatorCost,omitempty"`
	// The amount of output data in the task. Unit: bytes.
	OutputDataSize *int64 `json:"OutputDataSize,omitempty" xml:"OutputDataSize,omitempty"`
	// The number of output rows in the task.
	OutputRows *int64 `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	// The peak memory of the task. Unit: bytes.
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The queuing duration of the task. Unit: milliseconds.
	QueuedTime *string `json:"QueuedTime,omitempty" xml:"QueuedTime,omitempty"`
	// The amount of time that is consumed to scan data from a data source in the task. Unit: milliseconds.
	ScanCost *int64 `json:"ScanCost,omitempty" xml:"ScanCost,omitempty"`
	// The amount of scanned data in the task. Unit: bytes.
	ScanDataSize *int64 `json:"ScanDataSize,omitempty" xml:"ScanDataSize,omitempty"`
	// The number of rows that are scanned from a data source in the task.
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// The final execution state of the task. Valid values:
	//
	// *   FINISHED
	// *   CANCELED
	// *   ABORTED
	// *   FAILED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The timestamp when the task was created.
	TaskCreateTime *int64 `json:"TaskCreateTime,omitempty" xml:"TaskCreateTime,omitempty"`
	// The timestamp when the task ends.
	TaskEndTime *int64 `json:"TaskEndTime,omitempty" xml:"TaskEndTime,omitempty"`
	// The IP address of the host where the task was executed.
	TaskHost *string `json:"TaskHost,omitempty" xml:"TaskHost,omitempty"`
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDiagnosisTasksResponseBodyTaskList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisTasksResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetComputeTimeRatio(v string) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.ComputeTimeRatio = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetDrivers(v string) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.Drivers = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetElapsedTime(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetInputDataSize(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.InputDataSize = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetInputRows(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.InputRows = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetOperatorCost(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.OperatorCost = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetOutputDataSize(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.OutputDataSize = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetOutputRows(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.OutputRows = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetPeakMemory(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.PeakMemory = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetQueuedTime(v string) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.QueuedTime = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetScanCost(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.ScanCost = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetScanDataSize(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.ScanDataSize = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetScanRows(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.ScanRows = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetState(v string) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.State = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetTaskCreateTime(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.TaskCreateTime = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetTaskEndTime(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.TaskEndTime = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetTaskHost(v string) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.TaskHost = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetTaskId(v string) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.TaskId = &v
	return s
}

type DescribeDiagnosisTasksResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiagnosisTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiagnosisTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosisTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisTasksResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisTasksResponse) SetStatusCode(v int32) *DescribeDiagnosisTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisTasksResponse) SetBody(v *DescribeDiagnosisTasksResponseBody) *DescribeDiagnosisTasksResponse {
	s.Body = v
	return s
}

type DescribeDownloadRecordsRequest struct {
	// The ID of the cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the detailed information of all AnalyticDB for MySQL clusters within a specific region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The languages available for file titles and some error messages. Default value: zh. Valid values:
	//
	// *   **zh**: Simplified Chinese
	// *   **en**: English
	// *   **ja**: Japanese
	// *   **zh-tw**: Traditional Chinese
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the regions and zones supported by AnalyticDB for MySQL, including region IDs.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDownloadRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsRequest) SetDBClusterId(v string) *DescribeDownloadRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDownloadRecordsRequest) SetLang(v string) *DescribeDownloadRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDownloadRecordsRequest) SetRegionId(v string) *DescribeDownloadRecordsRequest {
	s.RegionId = &v
	return s
}

type DescribeDownloadRecordsResponseBody struct {
	// Details about the download tasks.
	Records []*DescribeDownloadRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDownloadRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponseBody) SetRecords(v []*DescribeDownloadRecordsResponseBodyRecords) *DescribeDownloadRecordsResponseBody {
	s.Records = v
	return s
}

func (s *DescribeDownloadRecordsResponseBody) SetRequestId(v string) *DescribeDownloadRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDownloadRecordsResponseBodyRecords struct {
	// The ID of the download task.
	DownloadId *int64 `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	// The error message returned when the download task has failed.
	ExceptionMsg *string `json:"ExceptionMsg,omitempty" xml:"ExceptionMsg,omitempty"`
	// The name of the downloaded file.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The status of the download task.
	//
	// *   **running**: The download task is currently in progress.
	// *   **finished**: The download task is complete.
	// *   **failed**: The download task has failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The download URL of the file.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeDownloadRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetDownloadId(v int64) *DescribeDownloadRecordsResponseBodyRecords {
	s.DownloadId = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetExceptionMsg(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.ExceptionMsg = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetFileName(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.FileName = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetStatus(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.Status = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetUrl(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.Url = &v
	return s
}

type DescribeDownloadRecordsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDownloadRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDownloadRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponse) SetHeaders(v map[string]*string) *DescribeDownloadRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadRecordsResponse) SetStatusCode(v int32) *DescribeDownloadRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadRecordsResponse) SetBody(v *DescribeDownloadRecordsResponseBody) *DescribeDownloadRecordsResponse {
	s.Body = v
	return s
}

type DescribeEIURangeRequest struct {
	// The specifications of computing resources.
	//
	// >  You can call the [DescribeComputeResource](~~469002~~) operation to query the specifications of computing resources.
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// *   This parameter can be left empty when **Operation** is set to **Buy**.
	// *   This parameter must be specified when **Operation** is set to **Upgrade** or **Downgrade**.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The version of the AnalyticDB for MySQL Data Warehouse Edition cluster. Set the value to **3.0**.
	DBClusterVersion *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	// The type of the operation. Valid values:
	//
	// *   **Buy**: purchases a cluster.
	// *   **Upgrade**: upgrades a cluster.
	// *   **Downgrade**: downgrades a cluster.
	Operation    *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The zone ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~612293~~) operation to query the most recent zone list.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeEIURangeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEIURangeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEIURangeRequest) SetComputeResource(v string) *DescribeEIURangeRequest {
	s.ComputeResource = &v
	return s
}

func (s *DescribeEIURangeRequest) SetDBClusterId(v string) *DescribeEIURangeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeEIURangeRequest) SetDBClusterVersion(v string) *DescribeEIURangeRequest {
	s.DBClusterVersion = &v
	return s
}

func (s *DescribeEIURangeRequest) SetOperation(v string) *DescribeEIURangeRequest {
	s.Operation = &v
	return s
}

func (s *DescribeEIURangeRequest) SetOwnerAccount(v string) *DescribeEIURangeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeEIURangeRequest) SetOwnerId(v int64) *DescribeEIURangeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEIURangeRequest) SetRegionId(v string) *DescribeEIURangeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEIURangeRequest) SetResourceGroupId(v string) *DescribeEIURangeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEIURangeRequest) SetResourceOwnerAccount(v string) *DescribeEIURangeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEIURangeRequest) SetResourceOwnerId(v int64) *DescribeEIURangeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEIURangeRequest) SetZoneId(v string) *DescribeEIURangeRequest {
	s.ZoneId = &v
	return s
}

type DescribeEIURangeResponseBody struct {
	// The queried information about the number of EIUs.
	EIUInfo *DescribeEIURangeResponseBodyEIUInfo `json:"EIUInfo,omitempty" xml:"EIUInfo,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEIURangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEIURangeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEIURangeResponseBody) SetEIUInfo(v *DescribeEIURangeResponseBodyEIUInfo) *DescribeEIURangeResponseBody {
	s.EIUInfo = v
	return s
}

func (s *DescribeEIURangeResponseBody) SetRequestId(v string) *DescribeEIURangeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEIURangeResponseBodyEIUInfo struct {
	// The suggested value for the number of EIUs.
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The queried range for the number of EIUs.
	EIURange []*int64 `json:"EIURange,omitempty" xml:"EIURange,omitempty" type:"Repeated"`
	// A reserved parameter.
	StorageResourceRange []*string `json:"StorageResourceRange,omitempty" xml:"StorageResourceRange,omitempty" type:"Repeated"`
}

func (s DescribeEIURangeResponseBodyEIUInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeEIURangeResponseBodyEIUInfo) GoString() string {
	return s.String()
}

func (s *DescribeEIURangeResponseBodyEIUInfo) SetDefaultValue(v string) *DescribeEIURangeResponseBodyEIUInfo {
	s.DefaultValue = &v
	return s
}

func (s *DescribeEIURangeResponseBodyEIUInfo) SetEIURange(v []*int64) *DescribeEIURangeResponseBodyEIUInfo {
	s.EIURange = v
	return s
}

func (s *DescribeEIURangeResponseBodyEIUInfo) SetStorageResourceRange(v []*string) *DescribeEIURangeResponseBodyEIUInfo {
	s.StorageResourceRange = v
	return s
}

type DescribeEIURangeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEIURangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEIURangeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEIURangeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEIURangeResponse) SetHeaders(v map[string]*string) *DescribeEIURangeResponse {
	s.Headers = v
	return s
}

func (s *DescribeEIURangeResponse) SetStatusCode(v int32) *DescribeEIURangeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEIURangeResponse) SetBody(v *DescribeEIURangeResponseBody) *DescribeEIURangeResponse {
	s.Body = v
	return s
}

type DescribeElasticDailyPlanRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The start date of the current-day scaling plan. Specify the date in the yyyy-MM-dd format.
	ElasticDailyPlanDay *string `json:"ElasticDailyPlanDay,omitempty" xml:"ElasticDailyPlanDay,omitempty"`
	// The execution state of the current-day scaling plan. Separate multiple values with commas (,). Valid values:
	//
	// *   **1**: The scaling plan is not executed.
	// *   **2**: The scaling plan is being executed.
	// *   **3**: The scaling plan is executed.
	// *   **4**: The scaling plan fails to be executed.
	ElasticDailyPlanStatusList *string `json:"ElasticDailyPlanStatusList,omitempty" xml:"ElasticDailyPlanStatusList,omitempty"`
	// The name of the scaling plan. Valid values:
	//
	// *   The name must be 2 to 30 characters in length.
	// *   The name can contain letters, digits, and underscores (\_).
	ElasticPlanName      *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the resource group.
	//
	// >  You can call the [DescribeDBResourceGroup](~~466685~~) operation to query the resource group name.
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
}

func (s DescribeElasticDailyPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticDailyPlanRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticDailyPlanRequest) SetDBClusterId(v string) *DescribeElasticDailyPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetElasticDailyPlanDay(v string) *DescribeElasticDailyPlanRequest {
	s.ElasticDailyPlanDay = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetElasticDailyPlanStatusList(v string) *DescribeElasticDailyPlanRequest {
	s.ElasticDailyPlanStatusList = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetElasticPlanName(v string) *DescribeElasticDailyPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetOwnerAccount(v string) *DescribeElasticDailyPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetOwnerId(v int64) *DescribeElasticDailyPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetResourceOwnerAccount(v string) *DescribeElasticDailyPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetResourceOwnerId(v int64) *DescribeElasticDailyPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetResourcePoolName(v string) *DescribeElasticDailyPlanRequest {
	s.ResourcePoolName = &v
	return s
}

type DescribeElasticDailyPlanResponseBody struct {
	// Details of the current-day scaling plans.
	ElasticDailyPlanList []*DescribeElasticDailyPlanResponseBodyElasticDailyPlanList `json:"ElasticDailyPlanList,omitempty" xml:"ElasticDailyPlanList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeElasticDailyPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticDailyPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticDailyPlanResponseBody) SetElasticDailyPlanList(v []*DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) *DescribeElasticDailyPlanResponseBody {
	s.ElasticDailyPlanList = v
	return s
}

func (s *DescribeElasticDailyPlanResponseBody) SetRequestId(v string) *DescribeElasticDailyPlanResponseBody {
	s.RequestId = &v
	return s
}

type DescribeElasticDailyPlanResponseBodyElasticDailyPlanList struct {
	// The start date of the current-day scaling plan. The date is in the yyyy-MM-dd format.
	Day *string `json:"Day,omitempty" xml:"Day,omitempty"`
	// The number of nodes involved in the scaling plan.
	//
	// *   If ElasticPlanType is set to **worker**, a value of 0 or null is returned.
	// *   If ElasticPlanType is set to **executorcombineworker** or **executor**, a value greater than 0 is returned.
	ElasticNodeNum *int32 `json:"ElasticNodeNum,omitempty" xml:"ElasticNodeNum,omitempty"`
	// The type of the scaling plan. Default value: executorcombineworker. Valid values:
	//
	// *   **worker**: scales only elastic I/O resources.
	// *   **executor**: scales only computing resources.
	// *   **executorcombineworker**: scales both elastic I/O resources and computing resources by proportion.
	ElasticPlanType *string `json:"ElasticPlanType,omitempty" xml:"ElasticPlanType,omitempty"`
	// The resource specifications that can be scaled up by the scaling plan. Default value: 8 Core 64 GB. Valid values:
	//
	// *   8 Core 64 GB
	// *   16 Core 64 GB
	// *   32 Core 64 GB
	// *   64 Core 128 GB
	// *   12 Core 96 GB
	// *   24 Core 96 GB
	// *   52 Core 86 GB
	ElasticPlanWorkerSpec *string `json:"ElasticPlanWorkerSpec,omitempty" xml:"ElasticPlanWorkerSpec,omitempty"`
	// The actual restoration time. The time is in the yyyy-MM-dd hh:mm:ss format. The time is displayed in UTC.
	EndTs *string `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// The scheduled restoration time. The time is in the yyyy-MM-dd hh:mm:ss format. The time is displayed in UTC.
	PlanEndTs *string `json:"PlanEndTs,omitempty" xml:"PlanEndTs,omitempty"`
	// The name of the scaling plan.
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The scheduled scale-up time. The time is in the yyyy-MM-dd hh:mm:ss format. The time is displayed in UTC.
	PlanStartTs *string `json:"PlanStartTs,omitempty" xml:"PlanStartTs,omitempty"`
	// The name of the resource group.
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
	// The actual scale-up time. The time is in the yyyy-MM-dd hh:mm:ss format. The time is displayed in UTC.
	StartTs *string `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// The execution state of the current-day scaling plan. Multiple values are separated by commas (,). Valid values:
	//
	// *   **1**: The scaling plan is not executed.
	// *   **2**: The scaling plan is being executed.
	// *   **3**: The scaling plan is executed.
	// *   **4**: The scaling plan fails to be executed.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) GoString() string {
	return s.String()
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetDay(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.Day = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetElasticNodeNum(v int32) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.ElasticNodeNum = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetElasticPlanType(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.ElasticPlanType = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetElasticPlanWorkerSpec(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.ElasticPlanWorkerSpec = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetEndTs(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.EndTs = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetPlanEndTs(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.PlanEndTs = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetPlanName(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.PlanName = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetPlanStartTs(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.PlanStartTs = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetResourcePoolName(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.ResourcePoolName = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetStartTs(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.StartTs = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetStatus(v int32) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.Status = &v
	return s
}

type DescribeElasticDailyPlanResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeElasticDailyPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeElasticDailyPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticDailyPlanResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticDailyPlanResponse) SetHeaders(v map[string]*string) *DescribeElasticDailyPlanResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticDailyPlanResponse) SetStatusCode(v int32) *DescribeElasticDailyPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticDailyPlanResponse) SetBody(v *DescribeElasticDailyPlanResponseBody) *DescribeElasticDailyPlanResponse {
	s.Body = v
	return s
}

type DescribeElasticPlanRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether the scaling plan takes effect. Valid values:
	//
	// *   **true** (default)
	// *   **false**
	ElasticPlanEnable *bool `json:"ElasticPlanEnable,omitempty" xml:"ElasticPlanEnable,omitempty"`
	// The name of the scaling plan.
	//
	// *   The name must be 2 to 30 characters in length.
	// *   The name can contain letters, digits, and underscores (\_).
	//
	// > If you do not specify this parameter, the information about all scaling plans for the specified cluster is returned.
	ElasticPlanName      *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the resource group.
	//
	// > You can call the [DescribeDBResourceGroup](~~466685~~) operation to query the resource group name.
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
}

func (s DescribeElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanRequest) SetDBClusterId(v string) *DescribeElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetElasticPlanEnable(v bool) *DescribeElasticPlanRequest {
	s.ElasticPlanEnable = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetElasticPlanName(v string) *DescribeElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetOwnerAccount(v string) *DescribeElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetOwnerId(v int64) *DescribeElasticPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetResourceOwnerAccount(v string) *DescribeElasticPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetResourceOwnerId(v int64) *DescribeElasticPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetResourcePoolName(v string) *DescribeElasticPlanRequest {
	s.ResourcePoolName = &v
	return s
}

type DescribeElasticPlanResponseBody struct {
	// The queried scaling plans.
	ElasticPlanList []*DescribeElasticPlanResponseBodyElasticPlanList `json:"ElasticPlanList,omitempty" xml:"ElasticPlanList,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeElasticPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanResponseBody) SetElasticPlanList(v []*DescribeElasticPlanResponseBodyElasticPlanList) *DescribeElasticPlanResponseBody {
	s.ElasticPlanList = v
	return s
}

func (s *DescribeElasticPlanResponseBody) SetRequestId(v string) *DescribeElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

type DescribeElasticPlanResponseBodyElasticPlanList struct {
	// The number of nodes that are involved in the scaling plan.
	//
	// *   If ElasticPlanType is set to **worker**, a value of 0 or null is returned.
	// *   If ElasticPlanType is set to **executorcombineworker** or **executor**, a value greater than 0 is returned.
	ElasticNodeNum *int32 `json:"ElasticNodeNum,omitempty" xml:"ElasticNodeNum,omitempty"`
	// The type of the scaling plan. Valid values:
	//
	// *   **worker**: scales only elastic I/O resources.
	// *   **executor**: scales only computing resources.
	// *   **executorcombineworker** (default): scales both elastic I/O resources and computing resources by proportion.
	ElasticPlanType *string `json:"ElasticPlanType,omitempty" xml:"ElasticPlanType,omitempty"`
	// The resource specifications that can be scaled up by the scaling plan. Valid values:
	//
	// *   8 Core 64 GB (default)
	// *   16 Core 64 GB
	// *   32 Core 64 GB
	// *   64 Core 128 GB
	// *   12 Core 96 GB
	// *   24 Core 96 GB
	// *   52 Core 86 GB
	ElasticPlanWorkerSpec *string `json:"ElasticPlanWorkerSpec,omitempty" xml:"ElasticPlanWorkerSpec,omitempty"`
	// Indicates whether the scaling plan takes effect. Default value: true. Valid values:
	//
	// *   **true** (default)
	// *   **false**
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The end date of the scaling plan. This parameter is returned only if the end date of the scaling plan is set. The date is in the yyyy-MM-dd format.
	EndDay *string `json:"EndDay,omitempty" xml:"EndDay,omitempty"`
	// The restoration time of the scaling plan. The interval between the scale-up time and the restoration time cannot be more than 24 hours. The time is in the HH:mm:ss format.
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MonthlyRepeat *string `json:"MonthlyRepeat,omitempty" xml:"MonthlyRepeat,omitempty"`
	// The name of the scaling plan.
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The name of the resource group.
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
	// The start date of the scaling plan. This parameter is returned only if the start date of the scaling plan is set. The date is in the yyyy-MM-dd format.
	StartDay *string `json:"StartDay,omitempty" xml:"StartDay,omitempty"`
	// The scale-up time of the scaling plan. The time is in the HH:mm:ss format.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The days of the week when the scaling plan was executed. Valid values: 0 to 6, which indicate Sunday to Saturday. Multiple values are separated by commas (,).
	WeeklyRepeat *string `json:"WeeklyRepeat,omitempty" xml:"WeeklyRepeat,omitempty"`
}

func (s DescribeElasticPlanResponseBodyElasticPlanList) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanResponseBodyElasticPlanList) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetElasticNodeNum(v int32) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.ElasticNodeNum = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetElasticPlanType(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.ElasticPlanType = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetElasticPlanWorkerSpec(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.ElasticPlanWorkerSpec = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetEnable(v bool) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.Enable = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetEndDay(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.EndDay = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetEndTime(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.EndTime = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetMonthlyRepeat(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.MonthlyRepeat = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetPlanName(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.PlanName = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetResourcePoolName(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.ResourcePoolName = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetStartDay(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.StartDay = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetStartTime(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.StartTime = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetWeeklyRepeat(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.WeeklyRepeat = &v
	return s
}

type DescribeElasticPlanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeElasticPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanResponse) SetHeaders(v map[string]*string) *DescribeElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticPlanResponse) SetStatusCode(v int32) *DescribeElasticPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticPlanResponse) SetBody(v *DescribeElasticPlanResponseBody) *DescribeElasticPlanResponse {
	s.Body = v
	return s
}

type DescribeInclinedTablesRequest struct {
	// The ID of the cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order in which queries are sorted in the JSON format based on the specified fields. Specify the fields used to sort the queries and the order type.
	//
	// Example:
	//
	// ```
	//
	// [
	//
	//     {
	//
	//         "Field":"Name",
	//
	//         "Type":"Asc"
	//
	//     }
	//
	// ]
	// ```
	//
	// In the preceding code, Field indicates the field used to sort queries. Set the value of Field to Name.
	//
	// Type indicates the order type. Valid values of Type: Desc and Asc. A value of Desc indicates a descending order. A value of Asc indicates an ascending order.
	//
	// Both fields are not case-sensitive.
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   30
	// *   50
	// *   100
	//
	// Default value: 30.
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the table. Valid values:
	//
	// *   FactTable
	// *   DimensionTable
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s DescribeInclinedTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInclinedTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesRequest) SetDBClusterId(v string) *DescribeInclinedTablesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetOrder(v string) *DescribeInclinedTablesRequest {
	s.Order = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetOwnerAccount(v string) *DescribeInclinedTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetOwnerId(v int64) *DescribeInclinedTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetPageNumber(v int32) *DescribeInclinedTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetPageSize(v int32) *DescribeInclinedTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetResourceOwnerAccount(v string) *DescribeInclinedTablesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetResourceOwnerId(v int64) *DescribeInclinedTablesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetTableType(v string) *DescribeInclinedTablesRequest {
	s.TableType = &v
	return s
}

type DescribeInclinedTablesResponseBody struct {
	// The monitoring information about tables.
	Items *DescribeInclinedTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInclinedTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInclinedTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBody) SetItems(v *DescribeInclinedTablesResponseBodyItems) *DescribeInclinedTablesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetPageNumber(v string) *DescribeInclinedTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetPageSize(v string) *DescribeInclinedTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetRequestId(v string) *DescribeInclinedTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetTotalCount(v string) *DescribeInclinedTablesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeInclinedTablesResponseBodyItems struct {
	Table []*DescribeInclinedTablesResponseBodyItemsTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeInclinedTablesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeInclinedTablesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBodyItems) SetTable(v []*DescribeInclinedTablesResponseBodyItemsTable) *DescribeInclinedTablesResponseBodyItems {
	s.Table = v
	return s
}

type DescribeInclinedTablesResponseBodyItemsTable struct {
	// Indicates whether data is skewed in partitions of the table. Valid values:
	//
	// *   **true**
	// *   **false**
	IsIncline *bool `json:"IsIncline,omitempty" xml:"IsIncline,omitempty"`
	// The name of the table.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the database.
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The number of rows in the table.
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The type of the table. Valid values:
	//
	// *   **FactTable**
	// *   **DimensionTable**
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeInclinedTablesResponseBodyItemsTable) String() string {
	return tea.Prettify(s)
}

func (s DescribeInclinedTablesResponseBodyItemsTable) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetIsIncline(v bool) *DescribeInclinedTablesResponseBodyItemsTable {
	s.IsIncline = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetName(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Name = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetSchema(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Schema = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetSize(v int64) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Size = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetType(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Type = &v
	return s
}

type DescribeInclinedTablesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInclinedTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInclinedTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInclinedTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponse) SetHeaders(v map[string]*string) *DescribeInclinedTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInclinedTablesResponse) SetStatusCode(v int32) *DescribeInclinedTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInclinedTablesResponse) SetBody(v *DescribeInclinedTablesResponseBody) *DescribeInclinedTablesResponse {
	s.Body = v
	return s
}

type DescribeLoadTasksRecordsRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL clusters in a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database that is involved in the import or export task.
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The order in which to sort the tasks by field. Specify the field and the sort order in the JSON format. Example: `[{"Field":"CreateTime", "Type":"desc"}]`.
	//
	// >
	//
	// *   `Field` specifies the field that is used to sort the tasks. Valid values of Field: `State`, `CreateTime`, `DBName`, `ProcessID`, `UpdateTime`, `JobName`, and `ProcessRows`.
	//
	// *   `Type` specifies the sort order. Valid values of Type: `Desc` and `Asc`. The values are case-insensitive.
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30** (default)
	// *   **50**
	// *   **100**
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time must be in UTC.
	//
	// > We recommend that you set the query start time to any point in time within 30 days.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the asynchronous import or export task to be queried. Valid values:
	//
	// *   **INIT**: The task is being initialized.
	// *   **RUNNING**: The task is running.
	// *   **FINISH**: The task is successful.
	// *   **FAILED**: The task fails.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeLoadTasksRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadTasksRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadTasksRecordsRequest) SetDBClusterId(v string) *DescribeLoadTasksRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetDBName(v string) *DescribeLoadTasksRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetEndTime(v string) *DescribeLoadTasksRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetOrder(v string) *DescribeLoadTasksRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetOwnerAccount(v string) *DescribeLoadTasksRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetOwnerId(v int64) *DescribeLoadTasksRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetPageNumber(v int32) *DescribeLoadTasksRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetPageSize(v int32) *DescribeLoadTasksRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetRegionId(v string) *DescribeLoadTasksRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetResourceOwnerAccount(v string) *DescribeLoadTasksRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetResourceOwnerId(v int64) *DescribeLoadTasksRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetStartTime(v string) *DescribeLoadTasksRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLoadTasksRecordsRequest) SetState(v string) *DescribeLoadTasksRecordsRequest {
	s.State = &v
	return s
}

type DescribeLoadTasksRecordsResponseBody struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried asynchronous import and export tasks.
	LoadTasksRecords []*DescribeLoadTasksRecordsResponseBodyLoadTasksRecords `json:"LoadTasksRecords,omitempty" xml:"LoadTasksRecords,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLoadTasksRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadTasksRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadTasksRecordsResponseBody) SetDBClusterId(v string) *DescribeLoadTasksRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetLoadTasksRecords(v []*DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) *DescribeLoadTasksRecordsResponseBody {
	s.LoadTasksRecords = v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetPageNumber(v string) *DescribeLoadTasksRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetPageSize(v string) *DescribeLoadTasksRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetRequestId(v string) *DescribeLoadTasksRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBody) SetTotalCount(v string) *DescribeLoadTasksRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeLoadTasksRecordsResponseBodyLoadTasksRecords struct {
	// The start time of the task. The time is accurate to milliseconds. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ss.SSSZ* format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the database that is involved in the import or export task.
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The task ID.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The process ID.
	ProcessID *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
	// The number of rows that are processed in the asynchronous import or export task.
	ProcessRows *int64 `json:"ProcessRows,omitempty" xml:"ProcessRows,omitempty"`
	// The SQL statement that is used in the asynchronous import or export task.
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The state of the task.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the task state was updated. The time is accurate to milliseconds. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ss.SSSZ* format. The time is displayed in UTC.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) GoString() string {
	return s.String()
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetCreateTime(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.CreateTime = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetDBName(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.DBName = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetJobName(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.JobName = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetProcessID(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.ProcessID = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetProcessRows(v int64) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.ProcessRows = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetSql(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.Sql = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetState(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.State = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords) SetUpdateTime(v string) *DescribeLoadTasksRecordsResponseBodyLoadTasksRecords {
	s.UpdateTime = &v
	return s
}

type DescribeLoadTasksRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLoadTasksRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLoadTasksRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoadTasksRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadTasksRecordsResponse) SetHeaders(v map[string]*string) *DescribeLoadTasksRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadTasksRecordsResponse) SetStatusCode(v int32) *DescribeLoadTasksRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponse) SetBody(v *DescribeLoadTasksRecordsResponseBody) *DescribeLoadTasksRecordsResponse {
	s.Body = v
	return s
}

type DescribeMaintenanceActionRequest struct {
	// Specifies whether to return the information about pending or historical O\&M events. Valid values:
	//
	// *   **0**: returns the information about pending O\&M event.
	// *   **1**: returns the information about historical O\&M event.
	//
	// If you do not specify this parameter, the information about pending O\&M event are returned.
	IsHistory    *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30**, **50**, and **100**. Default value: 30.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. Valid values:
	//
	// *   The ID of the region where the O\&M event occurs. Example: `cn-hangzhou`. You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	// *   You can also set Region to `all` to query the O\&M events in all regions. If you set `Region` to `all`, you must set `TaskType` to `all`.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the region where the O\&M event occurs.
	//
	// > You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the O\&M event. Valid values:
	//
	// *   **rds_apsaradb_upgrade**: database software upgrades.
	// *   **all**: all the O\&M events in all regions within the current account. If you set `Region` to `all`, you must set `TaskType` to `all`.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeMaintenanceActionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMaintenanceActionRequest) GoString() string {
	return s.String()
}

func (s *DescribeMaintenanceActionRequest) SetIsHistory(v int32) *DescribeMaintenanceActionRequest {
	s.IsHistory = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetOwnerAccount(v string) *DescribeMaintenanceActionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetOwnerId(v int64) *DescribeMaintenanceActionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetPageNumber(v int32) *DescribeMaintenanceActionRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetPageSize(v int32) *DescribeMaintenanceActionRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetRegion(v string) *DescribeMaintenanceActionRequest {
	s.Region = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetRegionId(v string) *DescribeMaintenanceActionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetResourceOwnerAccount(v string) *DescribeMaintenanceActionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetResourceOwnerId(v int64) *DescribeMaintenanceActionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetTaskType(v string) *DescribeMaintenanceActionRequest {
	s.TaskType = &v
	return s
}

type DescribeMaintenanceActionResponseBody struct {
	// The queried O\&M events.
	Items []*DescribeMaintenanceActionResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeMaintenanceActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMaintenanceActionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMaintenanceActionResponseBody) SetItems(v []*DescribeMaintenanceActionResponseBodyItems) *DescribeMaintenanceActionResponseBody {
	s.Items = v
	return s
}

func (s *DescribeMaintenanceActionResponseBody) SetPageNumber(v int32) *DescribeMaintenanceActionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBody) SetPageSize(v int32) *DescribeMaintenanceActionResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBody) SetRequestId(v string) *DescribeMaintenanceActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBody) SetTotalRecordCount(v int32) *DescribeMaintenanceActionResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeMaintenanceActionResponseBodyItems struct {
	// The time when the O\&M event was created. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the cluster that is involved in the O\&M event.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine.
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine.
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The deadline before which the event can be executed. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	Deadline *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// The ID of the O\&M event.
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The point in time at which the switchover time of the O\&M event was modified. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The preparation time that is required before the pending O\&M event can be switched. The time is in the `HH:mm:ss` format.
	PrepareInterval *string `json:"PrepareInterval,omitempty" xml:"PrepareInterval,omitempty"`
	// The ID of the region where the O\&M event occurs.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The execution result of the O\&M event.
	//
	// > This parameter is returned only when the value of `Status` is **FAILED** or **CANCEL**.
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// The time when the task was executed in the backend. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the event.
	//
	// *   If you set `IsHistory` to **0**, the state of the pending O\&M event is returned. Valid values:
	//
	//     *   **WAITING_MODIFY**: The start time of the O\&M event is waiting to be set.
	//     *   **WAITING**: The O\&M event is waiting to be processed.
	//     *   **PROCESSING**: The O\&M event is being processed. The switchover time of an event in this state cannot be changed.
	//
	// *   If you set `IsHistory` to **1**, the state of the historical O\&M event is returned. Valid values:
	//
	//     *   **SUCCESS**: The event ended and the execution succeeded.
	//     *   **FAILED**: The event ended but the execution failed.
	//     *   **CANCEL**: The event was canceled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the pending O\&M event is switched. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The type of the O\&M event.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeMaintenanceActionResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeMaintenanceActionResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetCreatedTime(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.CreatedTime = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetDBClusterId(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetDBType(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.DBType = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetDBVersion(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.DBVersion = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetDeadline(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.Deadline = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetId(v int32) *DescribeMaintenanceActionResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetModifiedTime(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetPrepareInterval(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.PrepareInterval = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetRegion(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.Region = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetResultInfo(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.ResultInfo = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetStartTime(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetStatus(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetSwitchTime(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.SwitchTime = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetTaskType(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.TaskType = &v
	return s
}

type DescribeMaintenanceActionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMaintenanceActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMaintenanceActionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMaintenanceActionResponse) GoString() string {
	return s.String()
}

func (s *DescribeMaintenanceActionResponse) SetHeaders(v map[string]*string) *DescribeMaintenanceActionResponse {
	s.Headers = v
	return s
}

func (s *DescribeMaintenanceActionResponse) SetStatusCode(v int32) *DescribeMaintenanceActionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMaintenanceActionResponse) SetBody(v *DescribeMaintenanceActionResponseBody) *DescribeMaintenanceActionResponse {
	s.Body = v
	return s
}

type DescribeOperatorPermissionRequest struct {
	// The ID of the cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeOperatorPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOperatorPermissionRequest) GoString() string {
	return s.String()
}

func (s *DescribeOperatorPermissionRequest) SetDBClusterId(v string) *DescribeOperatorPermissionRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetOwnerAccount(v string) *DescribeOperatorPermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetOwnerId(v int64) *DescribeOperatorPermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetResourceOwnerAccount(v string) *DescribeOperatorPermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetResourceOwnerId(v int64) *DescribeOperatorPermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeOperatorPermissionResponseBody struct {
	// The time when the authorization takes effect.
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The time when the authorization expires.
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The type of authorization. Valid values: Control | Data.
	Privileges *string `json:"Privileges,omitempty" xml:"Privileges,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOperatorPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOperatorPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOperatorPermissionResponseBody) SetCreatedTime(v string) *DescribeOperatorPermissionResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeOperatorPermissionResponseBody) SetDBClusterId(v string) *DescribeOperatorPermissionResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeOperatorPermissionResponseBody) SetExpiredTime(v string) *DescribeOperatorPermissionResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeOperatorPermissionResponseBody) SetPrivileges(v string) *DescribeOperatorPermissionResponseBody {
	s.Privileges = &v
	return s
}

func (s *DescribeOperatorPermissionResponseBody) SetRequestId(v string) *DescribeOperatorPermissionResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOperatorPermissionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOperatorPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOperatorPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOperatorPermissionResponse) GoString() string {
	return s.String()
}

func (s *DescribeOperatorPermissionResponse) SetHeaders(v map[string]*string) *DescribeOperatorPermissionResponse {
	s.Headers = v
	return s
}

func (s *DescribeOperatorPermissionResponse) SetStatusCode(v int32) *DescribeOperatorPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOperatorPermissionResponse) SetBody(v *DescribeOperatorPermissionResponseBody) *DescribeOperatorPermissionResponse {
	s.Body = v
	return s
}

type DescribePatternPerformanceRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The SQL pattern ID.
	//
	// > You can call the [DescribeSQLPatterns](~~321868~~) operation to query the information about all SQL patterns in an AnalyticDB for MySQL cluster within a period of time, including SQL pattern IDs.
	PatternId *string `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time must be in UTC.
	//
	// >
	//
	// *   Only data within the last 14 days can be queried. For example, if the current date is November 22 (UTC+8), you can query data on a day as early as November 9 by setting StartTime to 2021-11-08T16:00:00Z. If you set StartTime to a value earlier than 2021-11-08T16:00:00Z, the Performances parameter is empty.
	//
	// *   The maximum time range that can be specified is 24 hours.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePatternPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePatternPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceRequest) SetDBClusterId(v string) *DescribePatternPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetEndTime(v string) *DescribePatternPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetPatternId(v string) *DescribePatternPerformanceRequest {
	s.PatternId = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetRegionId(v string) *DescribePatternPerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetStartTime(v string) *DescribePatternPerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribePatternPerformanceResponseBody struct {
	// The end time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The queried performance metrics.
	Performances []*DescribePatternPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePatternPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePatternPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponseBody) SetEndTime(v string) *DescribePatternPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetPerformances(v []*DescribePatternPerformanceResponseBodyPerformances) *DescribePatternPerformanceResponseBody {
	s.Performances = v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetRequestId(v string) *DescribePatternPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePatternPerformanceResponseBody) SetStartTime(v string) *DescribePatternPerformanceResponseBody {
	s.StartTime = &v
	return s
}

type DescribePatternPerformanceResponseBodyPerformances struct {
	// The performance metric that was queried. Valid values:
	//
	// *   **AnalyticDB_PatternQueryCount**: the total number of queries executed in association with the SQL pattern.
	// *   **AnalyticDB_PatternQueryTime**: the total amount of time consumed by the queries executed in association with the SQL pattern.
	// *   **AnalyticDB_PatternExecutionTime**: the total execution duration of the queries executed in association with the SQL pattern.
	// *   **AnalyticDB_PatternPeakMemory**: the peak memory usage of the queries executed in association with the SQL pattern.
	// *   **AnalyticDB_PatternScanSize**: the amount of data scanned in the queries executed in association with the SQL pattern.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The queried performance metrics.
	Series []*DescribePatternPerformanceResponseBodyPerformancesSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	// The unit of the performance metric. Valid values:
	//
	// *   When the performance metric is related to the query duration (the `Key` value is `AnalyticDB_PatternQueryTime` or `AnalyticDB_PatternExecutionTime`), **ms** is returned.
	// *   When the performance metric is related to the memory usage (the `Key` value is `AnalyticDB_PatternPeakMemory`), **MB** is returned.
	// *   When the performance metric is related to the amount of data scanned (the `Key` value is `AnalyticDB_PatternScanSize`), **MB** is returned.
	// *   When the performance metric is related to the number of queries (the `Key` value is `AnalyticDB_PatternQueryCount`), this parameter is empty.
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribePatternPerformanceResponseBodyPerformances) String() string {
	return tea.Prettify(s)
}

func (s DescribePatternPerformanceResponseBodyPerformances) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponseBodyPerformances) SetKey(v string) *DescribePatternPerformanceResponseBodyPerformances {
	s.Key = &v
	return s
}

func (s *DescribePatternPerformanceResponseBodyPerformances) SetSeries(v []*DescribePatternPerformanceResponseBodyPerformancesSeries) *DescribePatternPerformanceResponseBodyPerformances {
	s.Series = v
	return s
}

func (s *DescribePatternPerformanceResponseBodyPerformances) SetUnit(v string) *DescribePatternPerformanceResponseBodyPerformances {
	s.Unit = &v
	return s
}

type DescribePatternPerformanceResponseBodyPerformancesSeries struct {
	// The name of the performance metric value. Valid values:
	//
	// *   When the `Key` parameter is set to `AnalyticDB_PatternQueryCount`, `pattern_query_count` is returned, which indicates the number of executions of the SQL statements in association with the SQL pattern.
	//
	// *   When the `Key` parameter is set to `AnalyticDB_PatternQueryTime`, the following values are returned:
	//
	//     *   `average_query_time`, which indicates the average total amount of time consumed by the SQL statements in association with the SQL pattern.
	//     *   `max_query_time`, which indicates the maximum total amount of time consumed by the SQL statements in association with the SQL pattern.
	//
	// *   When the `Key` parameter is set to `AnalyticDB_PatternExecutionTime`, the following values are returned:
	//
	//     *   `average_execution_time`, which indicates the average execution duration of the SQL statements in association with the SQL pattern.
	//     *   `max_execution_time`, which indicates the maximum execution duration of the SQL statements in association with the SQL pattern.
	//
	// *   When the `Key` parameter is set to `AnalyticDB_PatternPeakMemory`, the following values are returned:
	//
	//     *   `average_peak_memory`, which indicates the average peak memory usage of the SQL statements in association with the SQL pattern.
	//     *   `max_peak_memory`, which indicates the maximum peak memory usage of the SQL statements in association with the SQL pattern.
	//
	// *   When the `Key` parameter is set `AnalyticDB_PatternScanSize`, the following values are returned:
	//
	//     *   `average_scan_size`, which indicates the average amount of data scanned by the SQL statements in association with the SQL pattern.
	//     *   `max_scan_size`, which indicates the maximum amount of data scanned by the SQL statements in association with the SQL pattern.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The queried performance metrics.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribePatternPerformanceResponseBodyPerformancesSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribePatternPerformanceResponseBodyPerformancesSeries) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponseBodyPerformancesSeries) SetName(v string) *DescribePatternPerformanceResponseBodyPerformancesSeries {
	s.Name = &v
	return s
}

func (s *DescribePatternPerformanceResponseBodyPerformancesSeries) SetValues(v []*string) *DescribePatternPerformanceResponseBodyPerformancesSeries {
	s.Values = v
	return s
}

type DescribePatternPerformanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePatternPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePatternPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePatternPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceResponse) SetHeaders(v map[string]*string) *DescribePatternPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribePatternPerformanceResponse) SetStatusCode(v int32) *DescribePatternPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePatternPerformanceResponse) SetBody(v *DescribePatternPerformanceResponseBody) *DescribePatternPerformanceResponse {
	s.Body = v
	return s
}

type DescribeProcessListRequest struct {
	// The ID of the cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The keyword in an SQL statement, which is used to filter queries. Set the value to **SELECT**.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The order in which queries are sorted based on the specified fields. Specify this parameter as an ordered JSON array in the `[{"Field":"Time","Type":"Desc" },{ "Field":"User", "Type":"Asc" }]` format.
	//
	// *   **Field** specifies the field used to sort queries. Valid values: Time, User, Host, and DB.
	// *   **Type** specifies the sorting sequence. Valid values: **Desc** and **Asc**.
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 30. Valid values:
	//
	// *   **30**
	// *   **50**
	// *   **100**
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The execution duration used to filter queries. Queries that take a longer time than the specified execution duration are displayed. Unit: seconds.
	RunningTime *int32 `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	// Specifies whether to show a complete SQL statement. Valid values:
	//
	// *   **True**: shows a complete SQL statement.
	// *   **False**: shows only the first 100 characters of an SQL statement.
	//
	// >  The default value is False.
	ShowFull *bool `json:"ShowFull,omitempty" xml:"ShowFull,omitempty"`
	// The name of the user used to filter queries.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeProcessListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListRequest) GoString() string {
	return s.String()
}

func (s *DescribeProcessListRequest) SetDBClusterId(v string) *DescribeProcessListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeProcessListRequest) SetKeyword(v string) *DescribeProcessListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeProcessListRequest) SetOrder(v string) *DescribeProcessListRequest {
	s.Order = &v
	return s
}

func (s *DescribeProcessListRequest) SetOwnerAccount(v string) *DescribeProcessListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeProcessListRequest) SetOwnerId(v int64) *DescribeProcessListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeProcessListRequest) SetPageNumber(v int32) *DescribeProcessListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessListRequest) SetPageSize(v int32) *DescribeProcessListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessListRequest) SetResourceOwnerAccount(v string) *DescribeProcessListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeProcessListRequest) SetResourceOwnerId(v int64) *DescribeProcessListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeProcessListRequest) SetRunningTime(v int32) *DescribeProcessListRequest {
	s.RunningTime = &v
	return s
}

func (s *DescribeProcessListRequest) SetShowFull(v bool) *DescribeProcessListRequest {
	s.ShowFull = &v
	return s
}

func (s *DescribeProcessListRequest) SetUser(v string) *DescribeProcessListRequest {
	s.User = &v
	return s
}

type DescribeProcessListResponseBody struct {
	// Details of the queries.
	Items *DescribeProcessListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The total number of pages returned.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProcessListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBody) SetItems(v *DescribeProcessListResponseBodyItems) *DescribeProcessListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeProcessListResponseBody) SetPageNumber(v string) *DescribeProcessListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessListResponseBody) SetPageSize(v string) *DescribeProcessListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessListResponseBody) SetRequestId(v string) *DescribeProcessListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProcessListResponseBody) SetTotalCount(v string) *DescribeProcessListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeProcessListResponseBodyItems struct {
	Process []*DescribeProcessListResponseBodyItemsProcess `json:"Process,omitempty" xml:"Process,omitempty" type:"Repeated"`
}

func (s DescribeProcessListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyItems) SetProcess(v []*DescribeProcessListResponseBodyItemsProcess) *DescribeProcessListResponseBodyItems {
	s.Process = v
	return s
}

type DescribeProcessListResponseBodyItemsProcess struct {
	// The type of the statement. Only SELECT can be returned.
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The name of the database.
	DB *string `json:"DB,omitempty" xml:"DB,omitempty"`
	// The IP address from which the query was initiated.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ID of the worker thread.
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The SQL statement that is being executed. By default, the first 100 characters of the SQL statement are returned. If the ShowFull parameter is set to True, the complete SQL statement is returned.
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	// The unique ID of the query. You must specify this parameter when you use the KILL PROCESS statement.
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The amount of time that has elapsed from the start time of the query. Unit: seconds.
	Time *int32 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The username.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeProcessListResponseBodyItemsProcess) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBodyItemsProcess) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetCommand(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.Command = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetDB(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.DB = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetHost(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.Host = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetId(v int32) *DescribeProcessListResponseBodyItemsProcess {
	s.Id = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetInfo(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.Info = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetProcessId(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.ProcessId = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetStartTime(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.StartTime = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetTime(v int32) *DescribeProcessListResponseBodyItemsProcess {
	s.Time = &v
	return s
}

func (s *DescribeProcessListResponseBodyItemsProcess) SetUser(v string) *DescribeProcessListResponseBodyItemsProcess {
	s.User = &v
	return s
}

type DescribeProcessListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeProcessListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProcessListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponse) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponse) SetHeaders(v map[string]*string) *DescribeProcessListResponse {
	s.Headers = v
	return s
}

func (s *DescribeProcessListResponse) SetStatusCode(v int32) *DescribeProcessListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProcessListResponse) SetBody(v *DescribeProcessListResponseBody) *DescribeProcessListResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// Th language of the `LocalName` response parameter. Valid values:
	//
	// *   **zh-CN**: Chinese.
	// *   **en-US**: English.
	// *   **ja**: Japanese.
	//
	// > If you do not specify this parameter, the Chinese language is used.
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// The queried regions.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	// The region name.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the region.
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The queried zones.
	Zones *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetZones(v *DescribeRegionsResponseBodyRegionsRegionZones) *DescribeRegionsResponseBodyRegionsRegion {
	s.Zones = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegionZones struct {
	Zone []*DescribeRegionsResponseBodyRegionsRegionZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZones) SetZone(v []*DescribeRegionsResponseBodyRegionsRegionZonesZone) *DescribeRegionsResponseBodyRegionsRegionZones {
	s.Zone = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegionZonesZone struct {
	// The zone name.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// Indicates whether Virtual Private Cloud (VPC) is supported in the zone. Valid values:
	//
	// *   **true**
	// *   **false**
	VpcEnabled *bool `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
	// The zone ID.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetVpcEnabled(v bool) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.VpcEnabled = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.ZoneId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeResubmitConfigRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// >  You can call the [DescribeDBResourceGroup](~~459446~~) operation to query the resource group name of a cluster.
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeResubmitConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResubmitConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeResubmitConfigRequest) SetDBClusterId(v string) *DescribeResubmitConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeResubmitConfigRequest) SetGroupName(v string) *DescribeResubmitConfigRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeResubmitConfigRequest) SetOwnerAccount(v string) *DescribeResubmitConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeResubmitConfigRequest) SetOwnerId(v int64) *DescribeResubmitConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeResubmitConfigRequest) SetResourceGroupId(v string) *DescribeResubmitConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeResubmitConfigRequest) SetResourceOwnerAccount(v string) *DescribeResubmitConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeResubmitConfigRequest) SetResourceOwnerId(v int64) *DescribeResubmitConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeResubmitConfigResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job resubmission rules.
	Rules []*DescribeResubmitConfigResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeResubmitConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResubmitConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResubmitConfigResponseBody) SetDBClusterId(v string) *DescribeResubmitConfigResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeResubmitConfigResponseBody) SetRequestId(v string) *DescribeResubmitConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResubmitConfigResponseBody) SetRules(v []*DescribeResubmitConfigResponseBodyRules) *DescribeResubmitConfigResponseBody {
	s.Rules = v
	return s
}

type DescribeResubmitConfigResponseBodyRules struct {
	// Indicates whether out-of-memory (OOM) check is configured.
	ExceedMemoryException *bool `json:"ExceedMemoryException,omitempty" xml:"ExceedMemoryException,omitempty"`
	// The name of the source resource group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The peak memory usage.
	PeakMemory *string `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The duration of the SQL statement. Unit: milliseconds.
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// The name of the destination resource group.
	TargetGroupName *string `json:"TargetGroupName,omitempty" xml:"TargetGroupName,omitempty"`
}

func (s DescribeResubmitConfigResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeResubmitConfigResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeResubmitConfigResponseBodyRules) SetExceedMemoryException(v bool) *DescribeResubmitConfigResponseBodyRules {
	s.ExceedMemoryException = &v
	return s
}

func (s *DescribeResubmitConfigResponseBodyRules) SetGroupName(v string) *DescribeResubmitConfigResponseBodyRules {
	s.GroupName = &v
	return s
}

func (s *DescribeResubmitConfigResponseBodyRules) SetPeakMemory(v string) *DescribeResubmitConfigResponseBodyRules {
	s.PeakMemory = &v
	return s
}

func (s *DescribeResubmitConfigResponseBodyRules) SetQueryTime(v string) *DescribeResubmitConfigResponseBodyRules {
	s.QueryTime = &v
	return s
}

func (s *DescribeResubmitConfigResponseBodyRules) SetTargetGroupName(v string) *DescribeResubmitConfigResponseBodyRules {
	s.TargetGroupName = &v
	return s
}

type DescribeResubmitConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeResubmitConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResubmitConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResubmitConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeResubmitConfigResponse) SetHeaders(v map[string]*string) *DescribeResubmitConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeResubmitConfigResponse) SetStatusCode(v int32) *DescribeResubmitConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResubmitConfigResponse) SetBody(v *DescribeResubmitConfigResponseBody) *DescribeResubmitConfigResponse {
	s.Body = v
	return s
}

type DescribeSQAConfigRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// >  You can call the [DescribeDBResourceGroup](~~612410~~) operation to query the resource group name of a cluster.
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSQAConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQAConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQAConfigRequest) SetDBClusterId(v string) *DescribeSQAConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQAConfigRequest) SetGroupName(v string) *DescribeSQAConfigRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeSQAConfigRequest) SetOwnerAccount(v string) *DescribeSQAConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQAConfigRequest) SetOwnerId(v int64) *DescribeSQAConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQAConfigRequest) SetRegionId(v string) *DescribeSQAConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSQAConfigRequest) SetResourceGroupId(v string) *DescribeSQAConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSQAConfigRequest) SetResourceOwnerAccount(v string) *DescribeSQAConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQAConfigRequest) SetResourceOwnerId(v int64) *DescribeSQAConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeSQAConfigResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether short query acceleration (SQA) is enabled.
	SQAStatus *string `json:"SQAStatus,omitempty" xml:"SQAStatus,omitempty"`
}

func (s DescribeSQAConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQAConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQAConfigResponseBody) SetDBClusterId(v string) *DescribeSQAConfigResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQAConfigResponseBody) SetGroupName(v string) *DescribeSQAConfigResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeSQAConfigResponseBody) SetRequestId(v string) *DescribeSQAConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQAConfigResponseBody) SetSQAStatus(v string) *DescribeSQAConfigResponseBody {
	s.SQAStatus = &v
	return s
}

type DescribeSQAConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQAConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQAConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQAConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQAConfigResponse) SetHeaders(v map[string]*string) *DescribeSQAConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQAConfigResponse) SetStatusCode(v int32) *DescribeSQAConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQAConfigResponse) SetBody(v *DescribeSQAConfigResponseBody) *DescribeSQAConfigResponse {
	s.Body = v
	return s
}

type DescribeSQLPatternsRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL clusters in a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword that is used for the query.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of file titles and error messages. Valid values:
	//
	// *   **zh** (default): simplified Chinese.
	// *   **en**: English.
	// *   **ja**: Japanese.
	// *   **zh-tw**: traditional Chinese.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The order by which to sort query results. Specify the parameter value in the JSON format. Example: `[{"Field":"AverageQueryTime","Type":"Asc"}]`.
	//
	// *   `Field` specifies the field by which to sort the query results. Valid values:
	//
	//     *   `PatternCreationTime`: the earliest commit time of the SQL pattern within the time range to query.
	//     *   `AverageQueryTime`: the average total amount of time consumed by the SQL pattern within the time range to query.
	//     *   `MaxQueryTime`: the maximum total amount of time consumed by the SQL pattern within the time range to query.
	//     *   `AverageExecutionTime`: the average execution duration of the SQL pattern within the time range to query.
	//     *   `MaxExecutionTime`: the maximum execution duration of the SQL pattern within the time range to query.
	//     *   `AveragePeakMemory`: the average peak memory usage of the SQL pattern within the time range to query.
	//     *   `MaxPeakMemory`: the maximum peak memory usage of the SQL pattern within the time range to query.
	//     *   `AverageScanSize`: the average amount of data scanned based on the SQL pattern within the time range to query.
	//     *   `MaxScanSize`: the maximum amount of data scanned based on the SQL pattern within the time range to query.
	//     *   `QueryCount`: the number of queries performed in association with the SQL pattern within the time range to query.
	//     *   `FailedCount`: the number of failed queries performed in association with the SQL pattern within the time range to query.
	//
	// *   `Type` specifies the sorting order. Valid values (case-insensitive):
	//
	//     *   `Asc`: ascending order.
	//     *   `Desc`: descending order.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1.
	//
	// > If you do not specify this parameter, the value **1** is used.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// *   **30**
	// *   **50**
	// *   **100**
	//
	// > If you do not specify this parameter, the value **30** is used.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time must be in UTC.
	//
	// >
	//
	// *   Only data within the last 14 days can be queried. For example, if the current time is 2021-11-22T12:00:00Z, you can query SQL patterns at a point in time as early as 2021-11-09T12:00:00Z.
	//
	// *   The maximum time range that can be specified is 24 hours.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSQLPatternsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsRequest) SetDBClusterId(v string) *DescribeSQLPatternsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetEndTime(v string) *DescribeSQLPatternsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetKeyword(v string) *DescribeSQLPatternsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetLang(v string) *DescribeSQLPatternsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetOrder(v string) *DescribeSQLPatternsRequest {
	s.Order = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetPageNumber(v int32) *DescribeSQLPatternsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetPageSize(v int32) *DescribeSQLPatternsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetRegionId(v string) *DescribeSQLPatternsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSQLPatternsRequest) SetStartTime(v string) *DescribeSQLPatternsRequest {
	s.StartTime = &v
	return s
}

type DescribeSQLPatternsResponseBody struct {
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The queried SQL patterns.
	PatternDetails []*DescribeSQLPatternsResponseBodyPatternDetails `json:"PatternDetails,omitempty" xml:"PatternDetails,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSQLPatternsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsResponseBody) SetPageNumber(v int32) *DescribeSQLPatternsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetPageSize(v int32) *DescribeSQLPatternsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetPatternDetails(v []*DescribeSQLPatternsResponseBodyPatternDetails) *DescribeSQLPatternsResponseBody {
	s.PatternDetails = v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetRequestId(v string) *DescribeSQLPatternsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLPatternsResponseBody) SetTotalCount(v int32) *DescribeSQLPatternsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSQLPatternsResponseBodyPatternDetails struct {
	// The IP address of the SQL client that commits the SQL pattern.
	AccessIp *string `json:"AccessIp,omitempty" xml:"AccessIp,omitempty"`
	// The average execution duration of the SQL pattern within the query time range. Unit: milliseconds.
	AverageExecutionTime *float64 `json:"AverageExecutionTime,omitempty" xml:"AverageExecutionTime,omitempty"`
	// The average peak memory usage of the SQL pattern within the query time range. Unit: bytes.
	AveragePeakMemory *float64 `json:"AveragePeakMemory,omitempty" xml:"AveragePeakMemory,omitempty"`
	// The average total amount of time consumed by the SQL pattern within the query time range. Unit: milliseconds.
	AverageQueryTime *float64 `json:"AverageQueryTime,omitempty" xml:"AverageQueryTime,omitempty"`
	// The average amount of data scanned based on the SQL pattern within the query time range. Unit: bytes.
	AverageScanSize *float64 `json:"AverageScanSize,omitempty" xml:"AverageScanSize,omitempty"`
	// Indicates whether the execution of the SQL pattern can be blocked. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// > Only SELECT and INSERT statements can be blocked.
	Blockable *bool `json:"Blockable,omitempty" xml:"Blockable,omitempty"`
	// The number of failed queries executed in association with the SQL pattern within the query time range.
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The maximum execution duration of the SQL pattern within the query time range. Unit: milliseconds.
	MaxExecutionTime *int64 `json:"MaxExecutionTime,omitempty" xml:"MaxExecutionTime,omitempty"`
	// The maximum peak memory usage of the SQL pattern within the query time range. Unit: bytes.
	MaxPeakMemory *int64 `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	// The maximum total amount of time consumed by the SQL pattern within the query time range. Unit: milliseconds.
	MaxQueryTime *int64 `json:"MaxQueryTime,omitempty" xml:"MaxQueryTime,omitempty"`
	// The maximum amount of data scanned based on the SQL pattern within the query time range. Unit: bytes.
	MaxScanSize *int64 `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	// The earliest commit time of the SQL pattern within the query time range. Unit: milliseconds.
	PatternCreationTime *string `json:"PatternCreationTime,omitempty" xml:"PatternCreationTime,omitempty"`
	// The ID of the SQL pattern.
	PatternId *string `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	// The number of queries executed in association with the SQL pattern within the query time range.
	QueryCount *int64 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The statement of the SQL pattern.
	SQLPattern *string `json:"SQLPattern,omitempty" xml:"SQLPattern,omitempty"`
	// The tables scanned based on the SQL pattern.
	Tables *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
	// The database username that is used to commit the SQL pattern.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLPatternsResponseBodyPatternDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternsResponseBodyPatternDetails) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAccessIp(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AccessIp = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAverageExecutionTime(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AverageExecutionTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAveragePeakMemory(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AveragePeakMemory = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAverageQueryTime(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AverageQueryTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetAverageScanSize(v float64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.AverageScanSize = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetBlockable(v bool) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.Blockable = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetFailedCount(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.FailedCount = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxExecutionTime(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxExecutionTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxPeakMemory(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxPeakMemory = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxQueryTime(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxQueryTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetMaxScanSize(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.MaxScanSize = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetPatternCreationTime(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.PatternCreationTime = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetPatternId(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.PatternId = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetQueryCount(v int64) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.QueryCount = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetSQLPattern(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.SQLPattern = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetTables(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.Tables = &v
	return s
}

func (s *DescribeSQLPatternsResponseBodyPatternDetails) SetUser(v string) *DescribeSQLPatternsResponseBodyPatternDetails {
	s.User = &v
	return s
}

type DescribeSQLPatternsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLPatternsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLPatternsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPatternsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLPatternsResponse) SetHeaders(v map[string]*string) *DescribeSQLPatternsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLPatternsResponse) SetStatusCode(v int32) *DescribeSQLPatternsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLPatternsResponse) SetBody(v *DescribeSQLPatternsResponseBody) *DescribeSQLPatternsResponse {
	s.Body = v
	return s
}

type DescribeSQLPlanRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The query ID.
	//
	// > You can call the [DescribeProcessList](~~143382~~) operation to query the IDs of queries that are being executed.
	ProcessId            *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSQLPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanRequest) SetDBClusterId(v string) *DescribeSQLPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetOwnerAccount(v string) *DescribeSQLPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetOwnerId(v int64) *DescribeSQLPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetProcessId(v string) *DescribeSQLPlanRequest {
	s.ProcessId = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetResourceOwnerAccount(v string) *DescribeSQLPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetResourceOwnerId(v int64) *DescribeSQLPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeSQLPlanResponseBody struct {
	// The execution information about the SQL statement.
	Detail *DescribeSQLPlanResponseBodyDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	// The original information about the SQL statement.
	OriginInfo *string `json:"OriginInfo,omitempty" xml:"OriginInfo,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried plan in different stages.
	StageList []*DescribeSQLPlanResponseBodyStageList `json:"StageList,omitempty" xml:"StageList,omitempty" type:"Repeated"`
}

func (s DescribeSQLPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanResponseBody) SetDetail(v *DescribeSQLPlanResponseBodyDetail) *DescribeSQLPlanResponseBody {
	s.Detail = v
	return s
}

func (s *DescribeSQLPlanResponseBody) SetOriginInfo(v string) *DescribeSQLPlanResponseBody {
	s.OriginInfo = &v
	return s
}

func (s *DescribeSQLPlanResponseBody) SetRequestId(v string) *DescribeSQLPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLPlanResponseBody) SetStageList(v []*DescribeSQLPlanResponseBodyStageList) *DescribeSQLPlanResponseBody {
	s.StageList = v
	return s
}

type DescribeSQLPlanResponseBodyDetail struct {
	// The total CPU time consumed by all operators on multithreaded servers when the SQL statement is executed. Unit: milliseconds.
	CPUTime *int64 `json:"CPUTime,omitempty" xml:"CPUTime,omitempty"`
	// The IP address of the client that is used to execute the SQL statement.
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	// The name of the database on which the SQL statement was executed.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The total number of rows generated by the SQL statement.
	OutputRows *int64 `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	// The total amount of data generated by the SQL statement. Unit: bytes.
	OutputSize *int64 `json:"OutputSize,omitempty" xml:"OutputSize,omitempty"`
	// The maximum memory usage when the SQL statement is executed. Unit: bytes.
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The amount of time consumed to generate the execution plan of the SQL statement. Unit: milliseconds.
	PlanningTime *int64 `json:"PlanningTime,omitempty" xml:"PlanningTime,omitempty"`
	// The amount of time consumed to queue the SQL statement. Unit: milliseconds.
	QueuedTime *int64 `json:"QueuedTime,omitempty" xml:"QueuedTime,omitempty"`
	// The SQL statement.
	SQL *string `json:"SQL,omitempty" xml:"SQL,omitempty"`
	// The execution start time of the SQL statement. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The final execution state of the SQL statement. Valid values:
	//
	// *   FINISHED
	// *   FAILED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The total number of stages in the SQL statement.
	TotalStage *int64 `json:"TotalStage,omitempty" xml:"TotalStage,omitempty"`
	// The total number of tasks in the SQL statement.
	TotalTask *int64 `json:"TotalTask,omitempty" xml:"TotalTask,omitempty"`
	// The total amount of time consumed to execute the SQL statement. Unit: milliseconds.
	TotalTime *int64 `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	// The name of the user who submitted the SQL statement.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLPlanResponseBodyDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanResponseBodyDetail) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanResponseBodyDetail) SetCPUTime(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.CPUTime = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetClientIP(v string) *DescribeSQLPlanResponseBodyDetail {
	s.ClientIP = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetDatabase(v string) *DescribeSQLPlanResponseBodyDetail {
	s.Database = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetOutputRows(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.OutputRows = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetOutputSize(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.OutputSize = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetPeakMemory(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.PeakMemory = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetPlanningTime(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.PlanningTime = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetQueuedTime(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.QueuedTime = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetSQL(v string) *DescribeSQLPlanResponseBodyDetail {
	s.SQL = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetStartTime(v string) *DescribeSQLPlanResponseBodyDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetState(v string) *DescribeSQLPlanResponseBodyDetail {
	s.State = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetTotalStage(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.TotalStage = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetTotalTask(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.TotalTask = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetTotalTime(v int64) *DescribeSQLPlanResponseBodyDetail {
	s.TotalTime = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyDetail) SetUser(v string) *DescribeSQLPlanResponseBodyDetail {
	s.User = &v
	return s
}

type DescribeSQLPlanResponseBodyStageList struct {
	// The average `CPU Time` value on each compute node in the stage. Unit: milliseconds.
	CPUTimeAvg *int64 `json:"CPUTimeAvg,omitempty" xml:"CPUTimeAvg,omitempty"`
	// The maximum `CPU Time` value on each compute node in the stage. Unit: milliseconds.
	CPUTimeMax *int64 `json:"CPUTimeMax,omitempty" xml:"CPUTimeMax,omitempty"`
	// The minimum `CPU Time` value on each compute node in the stage. Unit: milliseconds.
	CPUTimeMin *int64 `json:"CPUTimeMin,omitempty" xml:"CPUTimeMin,omitempty"`
	// The average amount of input data on each compute node in the stage. Unit: bytes.
	InputSizeAvg *int64 `json:"InputSizeAvg,omitempty" xml:"InputSizeAvg,omitempty"`
	// The maximum amount of input data on each compute node in the stage. Unit: byte.
	InputSizeMax *int64 `json:"InputSizeMax,omitempty" xml:"InputSizeMax,omitempty"`
	// The minimum amount of input data on each compute node in the stage. Unit: bytes.
	InputSizeMin *int64 `json:"InputSizeMin,omitempty" xml:"InputSizeMin,omitempty"`
	// The total CPU time consumed by all operators in the stage, which is equivalent to the total CPU time of the stage. You can use this parameter to determine which parts of the stage consume a large amount of computing resources. Unit: milliseconds.
	OperatorCost *int64 `json:"OperatorCost,omitempty" xml:"OperatorCost,omitempty"`
	// The maximum memory usage when the SQL statement is executed. Unit: bytes.
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The average amount of data scanned by a scan operator on each storage node in the stage. Unit: bytes.
	ScanSizeAvg *int64 `json:"ScanSizeAvg,omitempty" xml:"ScanSizeAvg,omitempty"`
	// The maximum amount of data scanned by a scan operator on each storage node in the stage. Unit: bytes.
	ScanSizeMax *int64 `json:"ScanSizeMax,omitempty" xml:"ScanSizeMax,omitempty"`
	// The minimum amount of data scanned by a scan operator on each storage node in the stage. Unit: bytes.
	ScanSizeMin *int64 `json:"ScanSizeMin,omitempty" xml:"ScanSizeMin,omitempty"`
	// The average amount of time consumed by a scan operator to read data on each storage node in the stage. Unit: milliseconds.
	ScanTimeAvg *int64 `json:"ScanTimeAvg,omitempty" xml:"ScanTimeAvg,omitempty"`
	// The maximum amount of time consumed by a scan operator to read data on each storage node in the stage. Unit: milliseconds.
	ScanTimeMax *int64 `json:"ScanTimeMax,omitempty" xml:"ScanTimeMax,omitempty"`
	// The minimum amount of time consumed by a scan operator to read data on each storage node in the stage. Unit: milliseconds.
	ScanTimeMin *int64 `json:"ScanTimeMin,omitempty" xml:"ScanTimeMin,omitempty"`
	// The stage ID.
	StageId *int32 `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The final execution state of the stage. Valid values:
	//
	// *   FINISHED
	// *   CANCELED
	// *   ABORTED
	// *   FAILED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeSQLPlanResponseBodyStageList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanResponseBodyStageList) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanResponseBodyStageList) SetCPUTimeAvg(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.CPUTimeAvg = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetCPUTimeMax(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.CPUTimeMax = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetCPUTimeMin(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.CPUTimeMin = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetInputSizeAvg(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.InputSizeAvg = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetInputSizeMax(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.InputSizeMax = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetInputSizeMin(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.InputSizeMin = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetOperatorCost(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.OperatorCost = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetPeakMemory(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.PeakMemory = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanSizeAvg(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanSizeAvg = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanSizeMax(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanSizeMax = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanSizeMin(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanSizeMin = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanTimeAvg(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanTimeAvg = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanTimeMax(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanTimeMax = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetScanTimeMin(v int64) *DescribeSQLPlanResponseBodyStageList {
	s.ScanTimeMin = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetStageId(v int32) *DescribeSQLPlanResponseBodyStageList {
	s.StageId = &v
	return s
}

func (s *DescribeSQLPlanResponseBodyStageList) SetState(v string) *DescribeSQLPlanResponseBodyStageList {
	s.State = &v
	return s
}

type DescribeSQLPlanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanResponse) SetHeaders(v map[string]*string) *DescribeSQLPlanResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLPlanResponse) SetStatusCode(v int32) *DescribeSQLPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLPlanResponse) SetBody(v *DescribeSQLPlanResponseBody) *DescribeSQLPlanResponse {
	s.Body = v
	return s
}

type DescribeSQLPlanTaskRequest struct {
	// The ID of the cluster.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the task.
	ProcessId            *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The stage of the task.
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
}

func (s DescribeSQLPlanTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanTaskRequest) SetDBClusterId(v string) *DescribeSQLPlanTaskRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetOwnerAccount(v string) *DescribeSQLPlanTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetOwnerId(v int64) *DescribeSQLPlanTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetProcessId(v string) *DescribeSQLPlanTaskRequest {
	s.ProcessId = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetResourceOwnerAccount(v string) *DescribeSQLPlanTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetResourceOwnerId(v int64) *DescribeSQLPlanTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetStageId(v string) *DescribeSQLPlanTaskRequest {
	s.StageId = &v
	return s
}

type DescribeSQLPlanTaskResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of tasks.
	TaskList []*DescribeSQLPlanTaskResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s DescribeSQLPlanTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanTaskResponseBody) SetRequestId(v string) *DescribeSQLPlanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBody) SetTaskList(v []*DescribeSQLPlanTaskResponseBodyTaskList) *DescribeSQLPlanTaskResponseBody {
	s.TaskList = v
	return s
}

type DescribeSQLPlanTaskResponseBodyTaskList struct {
	// The time elapsed from when the task was created to when the task was complete. Unit: milliseconds.
	ElapsedTime *int64 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// The number of input rows in the task.
	InputRows *int64 `json:"InputRows,omitempty" xml:"InputRows,omitempty"`
	// The amount of input data in the task. Unit: bytes.
	InputSize *int64 `json:"InputSize,omitempty" xml:"InputSize,omitempty"`
	// The total amount of time consumed by operators in the task on a specific node. This parameter can be used to determine whether long tails occur in computing. Unit: milliseconds.
	OperatorCost *int64 `json:"OperatorCost,omitempty" xml:"OperatorCost,omitempty"`
	// The number of output rows in the task.
	OutputRows *int64 `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	// The amount of output data in the task. Unit: bytes.
	OutputSize *int64 `json:"OutputSize,omitempty" xml:"OutputSize,omitempty"`
	// The peak memory of the task on a specific node. Unit: bytes.
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The time consumed to scan data from a data source in the task. Unit: milliseconds.
	ScanCost *int64 `json:"ScanCost,omitempty" xml:"ScanCost,omitempty"`
	// The number of rows scanned from a data source in the task.
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// The amount of data scanned from a data source in the task. Unit: bytes.
	ScanSize *int64 `json:"ScanSize,omitempty" xml:"ScanSize,omitempty"`
	// The final execution status of the task. Valid values:
	//
	// *   FINISHED
	// *   CANCELED
	// *   ABORTED
	// *   FAILED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The ID of the task.
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeSQLPlanTaskResponseBodyTaskList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanTaskResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetElapsedTime(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetInputRows(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.InputRows = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetInputSize(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.InputSize = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetOperatorCost(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.OperatorCost = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetOutputRows(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.OutputRows = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetOutputSize(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.OutputSize = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetPeakMemory(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.PeakMemory = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetScanCost(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ScanCost = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetScanRows(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ScanRows = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetScanSize(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ScanSize = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetState(v string) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.State = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetTaskId(v int32) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.TaskId = &v
	return s
}

type DescribeSQLPlanTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLPlanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLPlanTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlanTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanTaskResponse) SetHeaders(v map[string]*string) *DescribeSQLPlanTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLPlanTaskResponse) SetStatusCode(v int32) *DescribeSQLPlanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLPlanTaskResponse) SetBody(v *DescribeSQLPlanTaskResponseBody) *DescribeSQLPlanTaskResponse {
	s.Body = v
	return s
}

type DescribeSchemasRequest struct {
	// The ID of the cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSchemasRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasRequest) GoString() string {
	return s.String()
}

func (s *DescribeSchemasRequest) SetDBClusterId(v string) *DescribeSchemasRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSchemasRequest) SetOwnerAccount(v string) *DescribeSchemasRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSchemasRequest) SetOwnerId(v int64) *DescribeSchemasRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSchemasRequest) SetResourceOwnerAccount(v string) *DescribeSchemasRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSchemasRequest) SetResourceOwnerId(v int64) *DescribeSchemasRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeSchemasResponseBody struct {
	// The databases.
	Items *DescribeSchemasResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSchemasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponseBody) SetItems(v *DescribeSchemasResponseBodyItems) *DescribeSchemasResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSchemasResponseBody) SetRequestId(v string) *DescribeSchemasResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSchemasResponseBodyItems struct {
	Schema []*DescribeSchemasResponseBodyItemsSchema `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Repeated"`
}

func (s DescribeSchemasResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponseBodyItems) SetSchema(v []*DescribeSchemasResponseBodyItemsSchema) *DescribeSchemasResponseBodyItems {
	s.Schema = v
	return s
}

type DescribeSchemasResponseBodyItemsSchema struct {
	// The ID of the cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeSchemasResponseBodyItemsSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasResponseBodyItemsSchema) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponseBodyItemsSchema) SetDBClusterId(v string) *DescribeSchemasResponseBodyItemsSchema {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSchemasResponseBodyItemsSchema) SetSchemaName(v string) *DescribeSchemasResponseBodyItemsSchema {
	s.SchemaName = &v
	return s
}

type DescribeSchemasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSchemasResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchemasResponse) GoString() string {
	return s.String()
}

func (s *DescribeSchemasResponse) SetHeaders(v map[string]*string) *DescribeSchemasResponse {
	s.Headers = v
	return s
}

func (s *DescribeSchemasResponse) SetStatusCode(v int32) *DescribeSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSchemasResponse) SetBody(v *DescribeSchemasResponseBody) *DescribeSchemasResponse {
	s.Body = v
	return s
}

type DescribeSlowLogRecordsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time must be in UTC.
	//
	// >  The end time must be later than the start time. The specified time range must be less than seven days.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The order in which to sort the retrieved entries by field. Specify this parameter in the JSON format. The value is an ordered array that uses the order of the input array and contains `Field` and `Type`. Example: `[{"Field":"ExecutionStartTime","Type":"Desc"},{"Field":"ScanRows","Type":"Asc"}]`.
	//
	// *   `Field`: the field that is used to sort the retrieved entries. Valid values:
	//
	//     *   **HostAddress**: the IP address of the client that is used to connect to the database.
	//     *   **UserName**: the username.
	//     *   **ExecutionStartTime**: the start time of the query execution.
	//     *   **QueryTime**: the amount of time consumed to execute the SQL statement.
	//     *   **PeakMemoryUsage**: the maximum memory usage when the SQL statement is executed.
	//     *   **ScanRows**: the number of rows to be scanned from a data source in the task.
	//     *   **ScanSize**: the amount of data to be scanned.
	//     *   **ScanTime**: the total amount of time consumed to scan data.
	//     *   **PlanningTime**: the amount of time consumed to generate execution plans.
	//     *   **WallTime**: the accumulated CPU Time values of all operators in the query on each node.
	//     *   **ProcessID**: the ID of the process.
	//
	// *   `Type`: the sorting type of the retrieved entries. Valid values:
	//
	//     *   **Desc**: descending order
	//     *   **Asc**: ascending order
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **30**, **50**, and **100**. Default value: 30.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the process.
	ProcessID *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
	// The range conditions used to filter specified fields, including `Max` and `Min`. Specify this parameter in the JSON format. Example: `[{"Field":"ScanSize","Min":"1000000","Max":"10000000"},{"Field":"QueryTime","Min":"1000","Max":"10000"}]`.
	//
	// `Field`: the field to be filtered. Valid values:
	//
	// *   **ScanSize**: the amount of data to be scanned. Unit: KB.
	// *   **QueryTime**: the amount of time consumed to execute the statement. Unit: milliseconds.
	// *   **PeakMemoryUsage**: the maximum memory usage when the SQL statement is executed. Unit: KB.
	//
	// >  `Min` indicates the minimum value of the query range (left operand). `Max` indicates the maximum value of the query range (right operand). Max and Min are both of the String type.
	Range                *string `json:"Range,omitempty" xml:"Range,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the query. Valid values:
	//
	// *   **Successed**: successful
	// *   **Failed**: failed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsRequest) SetDBClusterId(v string) *DescribeSlowLogRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBName(v string) *DescribeSlowLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetEndTime(v string) *DescribeSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOrder(v string) *DescribeSlowLogRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOwnerAccount(v string) *DescribeSlowLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOwnerId(v int64) *DescribeSlowLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageNumber(v int32) *DescribeSlowLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageSize(v int32) *DescribeSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetProcessID(v string) *DescribeSlowLogRecordsRequest {
	s.ProcessID = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetRange(v string) *DescribeSlowLogRecordsRequest {
	s.Range = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeSlowLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeSlowLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetStartTime(v string) *DescribeSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetState(v string) *DescribeSlowLogRecordsRequest {
	s.State = &v
	return s
}

type DescribeSlowLogRecordsResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Details of the slow query logs.
	Items *DescribeSlowLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) SetDBClusterId(v string) *DescribeSlowLogRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetItems(v *DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageNumber(v string) *DescribeSlowLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageSize(v string) *DescribeSlowLogRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetTotalCount(v string) *DescribeSlowLogRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSlowLogRecordsResponseBodyItems struct {
	SlowLogRecord []*DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord `json:"SlowLogRecord,omitempty" xml:"SlowLogRecord,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetSlowLogRecord(v []*DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) *DescribeSlowLogRecordsResponseBodyItems {
	s.SlowLogRecord = v
	return s
}

type DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord struct {
	// The name of the database.
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The time when the execution started. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ* format. The time is displayed in UTC.
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	// The IP address of the client that is used to connect to the database.
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The amount of output data in the task. Unit: bytes.
	OutputSize *string `json:"OutputSize,omitempty" xml:"OutputSize,omitempty"`
	// The number of rows parsed by the SQL statement.
	ParseRowCounts *int64 `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty"`
	// The maximum memory usage when the SQL statement is executed. Unit: KB.
	PeakMemoryUsage *string `json:"PeakMemoryUsage,omitempty" xml:"PeakMemoryUsage,omitempty"`
	// The amount of time consumed to generate execution plans. Unit: milliseconds.
	PlanningTime *int64 `json:"PlanningTime,omitempty" xml:"PlanningTime,omitempty"`
	// The ID of the process.
	ProcessID *string `json:"ProcessID,omitempty" xml:"ProcessID,omitempty"`
	// The time consumed to execute the SQL statement. Unit: milliseconds.
	QueryTime *int64 `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// The queuing duration before the query is executed. Unit: milliseconds.
	QueueTime *int64 `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	// The number of rows returned by the SQL statement.
	ReturnRowCounts *int64 `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	// Details of the SQL statement.
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	// The number of rows scanned from a data source in the task.
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// The amount of scanned data. Unit: KB.
	ScanSize *string `json:"ScanSize,omitempty" xml:"ScanSize,omitempty"`
	// The total amount of time consumed to scan data. It is an accumulated value collected from multiple TableScanNode nodes. Unit: milliseconds.
	ScanTime *int64 `json:"ScanTime,omitempty" xml:"ScanTime,omitempty"`
	// The execution state of the SQL statement.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The username.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The accumulated CPU Time value of all operators collected from all nodes. Unit: milliseconds.
	WallTime *int64 `json:"WallTime,omitempty" xml:"WallTime,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetExecutionStartTime(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ExecutionStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetHostAddress(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetOutputSize(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.OutputSize = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetParseRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ParseRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetPeakMemoryUsage(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.PeakMemoryUsage = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetPlanningTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.PlanningTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetProcessID(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ProcessID = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetQueryTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.QueryTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetQueueTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.QueueTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetReturnRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetSQLText(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetScanRows(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ScanRows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetScanSize(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ScanSize = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetScanTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.ScanTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetState(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.State = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetUserName(v string) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.UserName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord) SetWallTime(v int64) *DescribeSlowLogRecordsResponseBodyItemsSlowLogRecord {
	s.WallTime = &v
	return s
}

type DescribeSlowLogRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSlowLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlowLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeSlowLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetStatusCode(v int32) *DescribeSlowLogRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetBody(v *DescribeSlowLogRecordsResponseBody) *DescribeSlowLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeSlowLogTrendRequest struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. The maximum time range that can be specified is seven days. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendRequest) SetDBClusterId(v string) *DescribeSlowLogTrendRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetDBName(v string) *DescribeSlowLogTrendRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetEndTime(v string) *DescribeSlowLogTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetOwnerAccount(v string) *DescribeSlowLogTrendRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetOwnerId(v int64) *DescribeSlowLogTrendRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetResourceOwnerAccount(v string) *DescribeSlowLogTrendRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetResourceOwnerId(v int64) *DescribeSlowLogTrendRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetStartTime(v string) *DescribeSlowLogTrendRequest {
	s.StartTime = &v
	return s
}

type DescribeSlowLogTrendResponseBody struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The end time must be later than the start time. The maximum time range that can be specified is seven days. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The information about the trend of slow query logs.
	Items *DescribeSlowLogTrendResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBody) SetDBClusterId(v string) *DescribeSlowLogTrendResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBody) SetEndTime(v string) *DescribeSlowLogTrendResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBody) SetItems(v *DescribeSlowLogTrendResponseBodyItems) *DescribeSlowLogTrendResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSlowLogTrendResponseBody) SetRequestId(v string) *DescribeSlowLogTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBody) SetStartTime(v string) *DescribeSlowLogTrendResponseBody {
	s.StartTime = &v
	return s
}

type DescribeSlowLogTrendResponseBodyItems struct {
	SlowLogTrendItem []*DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem `json:"SlowLogTrendItem,omitempty" xml:"SlowLogTrendItem,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogTrendResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodyItems) SetSlowLogTrendItem(v []*DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem) *DescribeSlowLogTrendResponseBodyItems {
	s.SlowLogTrendItem = v
	return s
}

type DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem struct {
	// The trend of slow query logs. The value is AnalyticDB_SlowLogTrend.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The performance metrics.
	Series *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Struct"`
	// The unit of performance metrics.
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem) SetKey(v string) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem {
	s.Key = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem) SetSeries(v *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem {
	s.Series = v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem) SetUnit(v string) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItem {
	s.Unit = &v
	return s
}

type DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries struct {
	SeriesItem []*DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem `json:"SeriesItem,omitempty" xml:"SeriesItem,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries) SetSeriesItem(v []*DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeries {
	s.SeriesItem = v
	return s
}

type DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem struct {
	// The name of the performance metric.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The values of the performance metric.
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem) SetName(v string) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem {
	s.Name = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem) SetValues(v string) *DescribeSlowLogTrendResponseBodyItemsSlowLogTrendItemSeriesSeriesItem {
	s.Values = &v
	return s
}

type DescribeSlowLogTrendResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSlowLogTrendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlowLogTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponse) SetHeaders(v map[string]*string) *DescribeSlowLogTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowLogTrendResponse) SetStatusCode(v int32) *DescribeSlowLogTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowLogTrendResponse) SetBody(v *DescribeSlowLogTrendResponseBody) *DescribeSlowLogTrendResponse {
	s.Body = v
	return s
}

type DescribeSqlPatternRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order by which to sort query results. Specify the parameter value in the JSON string format. Example: `[{"Field":"Pattern","Type":"Asc"}]`. Parameters:
	//
	// *   `Field` specifies the field by which to sort the query results. Valid values:
	//
	//     *   `Pattern`: the SQL pattern.
	//     *   `AccessIP`: the IP address of the client.
	//     *   `User`: the username.
	//     *   `QueryCount`: the number of queries performed in association with the SQL pattern within the time range to query.
	//     *   `AvgPeakMemory`: the average peak memory usage of the SQL pattern within the time range to query. Unit: KB.
	//     *   `MaxPeakMemory`: the maximum peak memory usage of the SQL pattern within the time range to query. Unit: KB.
	//     *   `AvgCpuTime`: the average execution duration of the SQL pattern within the time range to query. Unit: milliseconds.
	//     *   `MaxCpuTime`: the maximum execution duration of the SQL pattern within the time range to query. Unit: milliseconds.
	//     *   `AvgStageCount`: the average number of stages.
	//     *   `MaxStageCount`: the maximum number of stages.
	//     *   `AvgTaskCount`: the average number of tasks.
	//     *   `MaxTaskCount`: the maximum number of tasks.
	//     *   `AvgScanSize`: the average amount of data scanned based on the SQL pattern within the time range to query. Unit: KB.
	//     *   `MaxScanSize`: the maximum amount of data scanned based on the SQL pattern within the time range to query. Unit: KB.
	//
	// *   `Type` specifies the sorting order. Valid values:
	//
	//     *   `Asc`: ascending order.
	//     *   `Desc`: descending order.
	//
	// >
	//
	// *   If you do not specify this parameter, query results are sorted in ascending order of `Pattern`.
	//
	// *   If you want to sort query results by `AccessIP`, you must set the `Type` parameter to `accessip`. If you want to sort query results by `User`, you must leave the `Type` parameter empty or set it to `user`.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The value must be a positive integer. Default value: **30**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The keyword that is used for the query.
	//
	// > If you do not specify this parameter, all SQL patterns of the AnalyticDB for MySQL cluster within the time period specified by `StartTime` are returned.
	SqlPattern *string `json:"SqlPattern,omitempty" xml:"SqlPattern,omitempty"`
	// The start date to query. Specify the time in the *yyyy-MM-dd* format. The time must be in UTC.
	//
	// > Only data within the last 30 days can be queried.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The dimension by which to aggregate the SQL patterns. Valid values:
	//
	// *   `user`: aggregates the SQL patterns by user.
	// *   `accessip`: aggregates the SQL patterns by client IP address.
	//
	// > If you do not specify this parameter, the SQL patterns are aggregated by `user`.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSqlPatternRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlPatternRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternRequest) SetDBClusterId(v string) *DescribeSqlPatternRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetOrder(v string) *DescribeSqlPatternRequest {
	s.Order = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetPageNumber(v int32) *DescribeSqlPatternRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetPageSize(v int32) *DescribeSqlPatternRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetRegionId(v string) *DescribeSqlPatternRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetSqlPattern(v string) *DescribeSqlPatternRequest {
	s.SqlPattern = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetStartTime(v string) *DescribeSqlPatternRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSqlPatternRequest) SetType(v string) *DescribeSqlPatternRequest {
	s.Type = &v
	return s
}

type DescribeSqlPatternResponseBody struct {
	// The queried SQL pattern.
	Items []*DescribeSqlPatternResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSqlPatternResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlPatternResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternResponseBody) SetItems(v []*DescribeSqlPatternResponseBodyItems) *DescribeSqlPatternResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetPageNumber(v int32) *DescribeSqlPatternResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetPageSize(v int32) *DescribeSqlPatternResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetRequestId(v string) *DescribeSqlPatternResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlPatternResponseBody) SetTotalCount(v int32) *DescribeSqlPatternResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSqlPatternResponseBodyItems struct {
	// The IP address of the client.
	//
	// > This parameter is returned only when `Type` is set to `accessip`.
	AccessIP *string `json:"AccessIP,omitempty" xml:"AccessIP,omitempty"`
	// The average execution duration of the SQL pattern within the time range to query. Unit: milliseconds.
	AvgCpuTime *string `json:"AvgCpuTime,omitempty" xml:"AvgCpuTime,omitempty"`
	// The average peak memory usage of the SQL pattern within the query time range. Unit: KB.
	AvgPeakMemory *string `json:"AvgPeakMemory,omitempty" xml:"AvgPeakMemory,omitempty"`
	// The average amount of data scanned based on the SQL pattern within the query time range. Unit: KB.
	AvgScanSize *string `json:"AvgScanSize,omitempty" xml:"AvgScanSize,omitempty"`
	// The average number of stages.
	AvgStageCount *string `json:"AvgStageCount,omitempty" xml:"AvgStageCount,omitempty"`
	// The average number of tasks.
	AvgTaskCount *string `json:"AvgTaskCount,omitempty" xml:"AvgTaskCount,omitempty"`
	// The cluster ID.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The maximum execution duration of the SQL pattern within the time range to query. Unit: milliseconds.
	MaxCpuTime *string `json:"MaxCpuTime,omitempty" xml:"MaxCpuTime,omitempty"`
	// The maximum peak memory usage of the SQL pattern within the query time range. Unit: KB.
	MaxPeakMemory *string `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	// The maximum amount of data scanned based on the SQL pattern within the query time range. Unit: KB.
	MaxScanSize *string `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	// The maximum number of stages.
	MaxStageCount *string `json:"MaxStageCount,omitempty" xml:"MaxStageCount,omitempty"`
	// The maximum number of tasks.
	MaxTaskCount *string `json:"MaxTaskCount,omitempty" xml:"MaxTaskCount,omitempty"`
	// The SQL pattern.
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The number of queries performed in association with the SQL pattern within the query time range.
	QueryCount *string `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// The start date of the query.
	ReportDate *string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
	// The username.
	//
	// > This parameter is returned only when `Type` is left empty or set to `user`.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSqlPatternResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlPatternResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternResponseBodyItems) SetAccessIP(v string) *DescribeSqlPatternResponseBodyItems {
	s.AccessIP = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgCpuTime(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgCpuTime = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgPeakMemory(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgPeakMemory = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgScanSize(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgScanSize = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgStageCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgStageCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetAvgTaskCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.AvgTaskCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetInstanceName(v string) *DescribeSqlPatternResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxCpuTime(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxCpuTime = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxPeakMemory(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxPeakMemory = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxScanSize(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxScanSize = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxStageCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxStageCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetMaxTaskCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.MaxTaskCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetPattern(v string) *DescribeSqlPatternResponseBodyItems {
	s.Pattern = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetQueryCount(v string) *DescribeSqlPatternResponseBodyItems {
	s.QueryCount = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetReportDate(v string) *DescribeSqlPatternResponseBodyItems {
	s.ReportDate = &v
	return s
}

func (s *DescribeSqlPatternResponseBodyItems) SetUser(v string) *DescribeSqlPatternResponseBodyItems {
	s.User = &v
	return s
}

type DescribeSqlPatternResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSqlPatternResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSqlPatternResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSqlPatternResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternResponse) SetHeaders(v map[string]*string) *DescribeSqlPatternResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlPatternResponse) SetStatusCode(v int32) *DescribeSqlPatternResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlPatternResponse) SetBody(v *DescribeSqlPatternResponseBody) *DescribeSqlPatternResponse {
	s.Body = v
	return s
}

type DescribeTableAccessCountRequest struct {
	// The ID of the cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the details of all AnalyticDB for MySQL clusters within a specified region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order by which to sort query results. Specify the parameter value in the JSON string format. Example: `[{"Field":"TableSchema","Type":"Asc"}]`.
	//
	// *   `Field` indicates the field that is used to sort the retrieved entries. Valid values:
	//
	//     *   `TableSchema`: the name of the database to which the table belongs.
	//     *   `TableName`: the name of the table.
	//     *   `AccessCount`: the number of accesses to the table.
	//
	// *   `Type` indicates the sorting method. Valid values:
	//
	//     *   `Asc`: ascending order.
	//     *   `Desc`: descending order.
	//
	// >  If this parameter is not specified, query results are sorted in ascending order of the database to which a specific table belongs.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The value must be a positive integer. Default value: **30**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the regions and zones supported by AnalyticDB for MySQL, including region IDs.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The date to query. Specify the time in the *yyyy-MM-dd* format. The time must be in UTC.
	//
	// >  Only data for the last 30 days can be queried.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the specific table.
	//
	// >  If this parameter is not specified, the number of accesses to all tables within the specified cluster for a specified date is returned.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableAccessCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableAccessCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountRequest) SetDBClusterId(v string) *DescribeTableAccessCountRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetOrder(v string) *DescribeTableAccessCountRequest {
	s.Order = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetPageNumber(v int32) *DescribeTableAccessCountRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetPageSize(v int32) *DescribeTableAccessCountRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetRegionId(v string) *DescribeTableAccessCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetStartTime(v string) *DescribeTableAccessCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTableAccessCountRequest) SetTableName(v string) *DescribeTableAccessCountRequest {
	s.TableName = &v
	return s
}

type DescribeTableAccessCountResponseBody struct {
	// Details about the table usage.
	Items []*DescribeTableAccessCountResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTableAccessCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableAccessCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountResponseBody) SetItems(v []*DescribeTableAccessCountResponseBodyItems) *DescribeTableAccessCountResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetPageNumber(v int32) *DescribeTableAccessCountResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetPageSize(v int32) *DescribeTableAccessCountResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetRequestId(v string) *DescribeTableAccessCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableAccessCountResponseBody) SetTotalCount(v int32) *DescribeTableAccessCountResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTableAccessCountResponseBodyItems struct {
	// The number of accesses to the table.
	AccessCount *string `json:"AccessCount,omitempty" xml:"AccessCount,omitempty"`
	// The ID of the cluster to which the table belongs.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The date when the table was used.
	ReportDate *string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The database to which the table belongs.
	TableSchema *string `json:"TableSchema,omitempty" xml:"TableSchema,omitempty"`
}

func (s DescribeTableAccessCountResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableAccessCountResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountResponseBodyItems) SetAccessCount(v string) *DescribeTableAccessCountResponseBodyItems {
	s.AccessCount = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetInstanceName(v string) *DescribeTableAccessCountResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetReportDate(v string) *DescribeTableAccessCountResponseBodyItems {
	s.ReportDate = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetTableName(v string) *DescribeTableAccessCountResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeTableAccessCountResponseBodyItems) SetTableSchema(v string) *DescribeTableAccessCountResponseBodyItems {
	s.TableSchema = &v
	return s
}

type DescribeTableAccessCountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTableAccessCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTableAccessCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableAccessCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountResponse) SetHeaders(v map[string]*string) *DescribeTableAccessCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableAccessCountResponse) SetStatusCode(v int32) *DescribeTableAccessCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTableAccessCountResponse) SetBody(v *DescribeTableAccessCountResponseBody) *DescribeTableAccessCountResponse {
	s.Body = v
	return s
}

type DescribeTableDetailRequest struct {
	// The ID of the cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailRequest) SetDBClusterId(v string) *DescribeTableDetailRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableDetailRequest) SetOwnerAccount(v string) *DescribeTableDetailRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTableDetailRequest) SetOwnerId(v int64) *DescribeTableDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTableDetailRequest) SetResourceOwnerAccount(v string) *DescribeTableDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTableDetailRequest) SetResourceOwnerId(v int64) *DescribeTableDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTableDetailRequest) SetSchemaName(v string) *DescribeTableDetailRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeTableDetailRequest) SetTableName(v string) *DescribeTableDetailRequest {
	s.TableName = &v
	return s
}

type DescribeTableDetailResponseBody struct {
	// The average number of rows in partitions.
	AvgSize *int64 `json:"AvgSize,omitempty" xml:"AvgSize,omitempty"`
	// The list of partitions.
	Items *DescribeTableDetailResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTableDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBody) SetAvgSize(v int64) *DescribeTableDetailResponseBody {
	s.AvgSize = &v
	return s
}

func (s *DescribeTableDetailResponseBody) SetItems(v *DescribeTableDetailResponseBodyItems) *DescribeTableDetailResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTableDetailResponseBody) SetRequestId(v string) *DescribeTableDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTableDetailResponseBodyItems struct {
	Shard []*DescribeTableDetailResponseBodyItemsShard `json:"Shard,omitempty" xml:"Shard,omitempty" type:"Repeated"`
}

func (s DescribeTableDetailResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBodyItems) SetShard(v []*DescribeTableDetailResponseBodyItemsShard) *DescribeTableDetailResponseBodyItems {
	s.Shard = v
	return s
}

type DescribeTableDetailResponseBodyItemsShard struct {
	// The ID of the partition. Only the numeric part of the partition name is returned.
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of rows in the table.
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeTableDetailResponseBodyItemsShard) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailResponseBodyItemsShard) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBodyItemsShard) SetId(v int32) *DescribeTableDetailResponseBodyItemsShard {
	s.Id = &v
	return s
}

func (s *DescribeTableDetailResponseBodyItemsShard) SetSize(v int64) *DescribeTableDetailResponseBodyItemsShard {
	s.Size = &v
	return s
}

type DescribeTableDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTableDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTableDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponse) SetHeaders(v map[string]*string) *DescribeTableDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableDetailResponse) SetStatusCode(v int32) *DescribeTableDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTableDetailResponse) SetBody(v *DescribeTableDetailResponseBody) *DescribeTableDetailResponse {
	s.Body = v
	return s
}

type DescribeTablePartitionDiagnoseRequest struct {
	// The ID of the cluster.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 30. Valid values:
	//
	// *   30
	// *   50
	// *   100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeTablePartitionDiagnoseRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseRequest) SetDBClusterId(v string) *DescribeTablePartitionDiagnoseRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetOwnerAccount(v string) *DescribeTablePartitionDiagnoseRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetOwnerId(v int64) *DescribeTablePartitionDiagnoseRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetPageNumber(v int32) *DescribeTablePartitionDiagnoseRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetPageSize(v int32) *DescribeTablePartitionDiagnoseRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetRegionId(v string) *DescribeTablePartitionDiagnoseRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetResourceOwnerAccount(v string) *DescribeTablePartitionDiagnoseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetResourceOwnerId(v int64) *DescribeTablePartitionDiagnoseRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeTablePartitionDiagnoseResponseBody struct {
	// The ID of the cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The information of tables.
	Items []*DescribeTablePartitionDiagnoseResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The recommended maximum number of rows in each list partition.
	SuggestMaxRecordsPerPartition *int64 `json:"SuggestMaxRecordsPerPartition,omitempty" xml:"SuggestMaxRecordsPerPartition,omitempty"`
	// The recommended minimum number of rows in each list partition.
	SuggestMinRecordsPerPartition *int64 `json:"SuggestMinRecordsPerPartition,omitempty" xml:"SuggestMinRecordsPerPartition,omitempty"`
	// The total number of entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTablePartitionDiagnoseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetDBClusterId(v string) *DescribeTablePartitionDiagnoseResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetItems(v []*DescribeTablePartitionDiagnoseResponseBodyItems) *DescribeTablePartitionDiagnoseResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetPageNumber(v int32) *DescribeTablePartitionDiagnoseResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetPageSize(v int32) *DescribeTablePartitionDiagnoseResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetRequestId(v string) *DescribeTablePartitionDiagnoseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetSuggestMaxRecordsPerPartition(v int64) *DescribeTablePartitionDiagnoseResponseBody {
	s.SuggestMaxRecordsPerPartition = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetSuggestMinRecordsPerPartition(v int64) *DescribeTablePartitionDiagnoseResponseBody {
	s.SuggestMinRecordsPerPartition = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBody) SetTotalCount(v int32) *DescribeTablePartitionDiagnoseResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTablePartitionDiagnoseResponseBodyItems struct {
	// Details of the inappropriate partitions.
	PartitionDetail *string `json:"PartitionDetail,omitempty" xml:"PartitionDetail,omitempty"`
	// The number of partitions.
	PartitionNumber *int32 `json:"PartitionNumber,omitempty" xml:"PartitionNumber,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTablePartitionDiagnoseResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetPartitionDetail(v string) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.PartitionDetail = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetPartitionNumber(v int32) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.PartitionNumber = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetSchemaName(v string) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponseBodyItems) SetTableName(v string) *DescribeTablePartitionDiagnoseResponseBodyItems {
	s.TableName = &v
	return s
}

type DescribeTablePartitionDiagnoseResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTablePartitionDiagnoseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTablePartitionDiagnoseResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseResponse) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseResponse) SetHeaders(v map[string]*string) *DescribeTablePartitionDiagnoseResponse {
	s.Headers = v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponse) SetStatusCode(v int32) *DescribeTablePartitionDiagnoseResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponse) SetBody(v *DescribeTablePartitionDiagnoseResponseBody) *DescribeTablePartitionDiagnoseResponse {
	s.Body = v
	return s
}

type DescribeTableStatisticsRequest struct {
	// The ID of the cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query details about all AnalyticDB for MySQL clusters in a specific region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order in which to sort the retrieved records by field. Specify this value in the JSON format. The value is an ordered array that uses the order of the input array and contains `Field` and `Type`. Example: `[{ "Field":"TableName", "Type":"Asc" }]`.
	//
	// *   In the example, `Field` indicates the field that is used to sort the retrieved records. Set the value of Field to `TableName`.
	//
	// *   `Type` indicates the sort type. Valid values (case-insensitive):
	//
	//     *   **Desc**: The entries are sorted in descending order.
	//     *   **Asc**: The entries are sorted in ascending order.
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **30**, **50**, and **100**. Default value: 30.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeTableStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsRequest) SetDBClusterId(v string) *DescribeTableStatisticsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetOrder(v string) *DescribeTableStatisticsRequest {
	s.Order = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetOwnerAccount(v string) *DescribeTableStatisticsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetOwnerId(v int64) *DescribeTableStatisticsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetPageNumber(v int32) *DescribeTableStatisticsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetPageSize(v int32) *DescribeTableStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetRegionId(v string) *DescribeTableStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetResourceOwnerAccount(v string) *DescribeTableStatisticsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTableStatisticsRequest) SetResourceOwnerId(v int64) *DescribeTableStatisticsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeTableStatisticsResponseBody struct {
	// The ID of the cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Details about table statistics.
	Items *DescribeTableStatisticsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTableStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsResponseBody) SetDBClusterId(v string) *DescribeTableStatisticsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetItems(v *DescribeTableStatisticsResponseBodyItems) *DescribeTableStatisticsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetPageNumber(v string) *DescribeTableStatisticsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetPageSize(v string) *DescribeTableStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetRequestId(v string) *DescribeTableStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetTotalCount(v string) *DescribeTableStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTableStatisticsResponseBodyItems struct {
	TableStatisticRecords []*DescribeTableStatisticsResponseBodyItemsTableStatisticRecords `json:"TableStatisticRecords,omitempty" xml:"TableStatisticRecords,omitempty" type:"Repeated"`
}

func (s DescribeTableStatisticsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableStatisticsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsResponseBodyItems) SetTableStatisticRecords(v []*DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) *DescribeTableStatisticsResponseBodyItems {
	s.TableStatisticRecords = v
	return s
}

type DescribeTableStatisticsResponseBodyItemsTableStatisticRecords struct {
	// The total amount of cold data. Unit: byte.
	//
	// >  The parameter is returned only when the engine version of the cluster is 3.1.3.4 or later.
	ColdDataSize *int64 `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty"`
	// The amount of data in the table. Unit: byte.
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The amount of data in indexes. Unit: byte.
	IndexSize *int64 `json:"IndexSize,omitempty" xml:"IndexSize,omitempty"`
	// The number of partitions.
	PartitionCount *int64 `json:"PartitionCount,omitempty" xml:"PartitionCount,omitempty"`
	// The amount of data in primary key indexes. Unit: byte.
	PrimaryKeyIndexSize *int64 `json:"PrimaryKeyIndexSize,omitempty" xml:"PrimaryKeyIndexSize,omitempty"`
	// The number of rows in the table.
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetColdDataSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.ColdDataSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetDataSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.DataSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetIndexSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.IndexSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetPartitionCount(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.PartitionCount = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetPrimaryKeyIndexSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.PrimaryKeyIndexSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetRowCount(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.RowCount = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetSchemaName(v string) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.SchemaName = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetTableName(v string) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.TableName = &v
	return s
}

type DescribeTableStatisticsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTableStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTableStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTableStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsResponse) SetHeaders(v map[string]*string) *DescribeTableStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableStatisticsResponse) SetStatusCode(v int32) *DescribeTableStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTableStatisticsResponse) SetBody(v *DescribeTableStatisticsResponseBody) *DescribeTableStatisticsResponse {
	s.Body = v
	return s
}

type DescribeTablesRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablesRequest) SetDBClusterId(v string) *DescribeTablesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablesRequest) SetOwnerAccount(v string) *DescribeTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTablesRequest) SetOwnerId(v int64) *DescribeTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTablesRequest) SetRegionId(v string) *DescribeTablesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTablesRequest) SetResourceOwnerAccount(v string) *DescribeTablesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTablesRequest) SetResourceOwnerId(v int64) *DescribeTablesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTablesRequest) SetSchemaName(v string) *DescribeTablesRequest {
	s.SchemaName = &v
	return s
}

type DescribeTablesResponseBody struct {
	// The queried tables.
	Items *DescribeTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBody) SetItems(v *DescribeTablesResponseBodyItems) *DescribeTablesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTablesResponseBody) SetRequestId(v string) *DescribeTablesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTablesResponseBodyItems struct {
	Table []*DescribeTablesResponseBodyItemsTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeTablesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyItems) SetTable(v []*DescribeTablesResponseBodyItemsTable) *DescribeTablesResponseBodyItems {
	s.Table = v
	return s
}

type DescribeTablesResponseBodyItemsTable struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTablesResponseBodyItemsTable) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBodyItemsTable) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyItemsTable) SetDBClusterId(v string) *DescribeTablesResponseBodyItemsTable {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsTable) SetSchemaName(v string) *DescribeTablesResponseBodyItemsTable {
	s.SchemaName = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsTable) SetTableName(v string) *DescribeTablesResponseBodyItemsTable {
	s.TableName = &v
	return s
}

type DescribeTablesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponse) SetHeaders(v map[string]*string) *DescribeTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTablesResponse) SetStatusCode(v int32) *DescribeTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTablesResponse) SetBody(v *DescribeTablesResponseBody) *DescribeTablesResponse {
	s.Body = v
	return s
}

type DescribeTaskInfoRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task ID.
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoRequest) SetDBClusterId(v string) *DescribeTaskInfoRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetOwnerAccount(v string) *DescribeTaskInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetOwnerId(v int64) *DescribeTaskInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetRegionId(v string) *DescribeTaskInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetResourceOwnerAccount(v string) *DescribeTaskInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetResourceOwnerId(v int64) *DescribeTaskInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTaskInfoRequest) SetTaskId(v int32) *DescribeTaskInfoRequest {
	s.TaskId = &v
	return s
}

type DescribeTaskInfoResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried task.
	TaskInfo *DescribeTaskInfoResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s DescribeTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoResponseBody) SetRequestId(v string) *DescribeTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskInfoResponseBody) SetTaskInfo(v *DescribeTaskInfoResponseBodyTaskInfo) *DescribeTaskInfoResponseBody {
	s.TaskInfo = v
	return s
}

type DescribeTaskInfoResponseBodyTaskInfo struct {
	// The start time of the task. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end time of the task. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The progress of the task. Unit: %.
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status. Valid values:
	//
	// *   Waiting
	// *   Running
	// *   Finished
	// *   Failed
	// *   Closed
	// *   Cancel
	// *   Retry
	// *   Pause
	// *   Stop
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTaskInfoResponseBodyTaskInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetBeginTime(v string) *DescribeTaskInfoResponseBodyTaskInfo {
	s.BeginTime = &v
	return s
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetFinishTime(v string) *DescribeTaskInfoResponseBodyTaskInfo {
	s.FinishTime = &v
	return s
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetProgress(v string) *DescribeTaskInfoResponseBodyTaskInfo {
	s.Progress = &v
	return s
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetStatus(v string) *DescribeTaskInfoResponseBodyTaskInfo {
	s.Status = &v
	return s
}

func (s *DescribeTaskInfoResponseBodyTaskInfo) SetTaskId(v int32) *DescribeTaskInfoResponseBodyTaskInfo {
	s.TaskId = &v
	return s
}

type DescribeTaskInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoResponse) SetHeaders(v map[string]*string) *DescribeTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskInfoResponse) SetStatusCode(v int32) *DescribeTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskInfoResponse) SetBody(v *DescribeTaskInfoResponseBody) *DescribeTaskInfoResponse {
	s.Body = v
	return s
}

type DescribeVSwitchesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the VPC ID.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	VswId *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
	// The zone ID.
	//
	// > You can call the [DescribeRegions](~~129857~~) operation to query the most recent zone list.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVSwitchesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesRequest) SetOwnerAccount(v string) *DescribeVSwitchesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetOwnerId(v int64) *DescribeVSwitchesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetRegionId(v string) *DescribeVSwitchesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetResourceOwnerAccount(v string) *DescribeVSwitchesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetResourceOwnerId(v int64) *DescribeVSwitchesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetSecurityToken(v string) *DescribeVSwitchesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVpcId(v string) *DescribeVSwitchesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVswId(v string) *DescribeVSwitchesRequest {
	s.VswId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetZoneId(v string) *DescribeVSwitchesRequest {
	s.ZoneId = &v
	return s
}

type DescribeVSwitchesResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried vSwitches.
	VSwitches *DescribeVSwitchesResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Struct"`
}

func (s DescribeVSwitchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBody) SetRequestId(v string) *DescribeVSwitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetVSwitches(v *DescribeVSwitchesResponseBodyVSwitches) *DescribeVSwitchesResponseBody {
	s.VSwitches = v
	return s
}

type DescribeVSwitchesResponseBodyVSwitches struct {
	// The queried vSwitch.
	VSwitch []*DescribeVSwitchesResponseBodyVSwitchesVSwitch `json:"VSwitch,omitempty" xml:"VSwitch,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchesResponseBodyVSwitches) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetVSwitch(v []*DescribeVSwitchesResponseBodyVSwitchesVSwitch) *DescribeVSwitchesResponseBodyVSwitches {
	s.VSwitch = v
	return s
}

type DescribeVSwitchesResponseBodyVSwitchesVSwitch struct {
	// The ID of the Resource Access Management (RAM) user.
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The ID of the user channel.
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The IPv4 CIDR block of the vSwitch.
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The time when the vSwitch was created.
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the vSwitch was modified.
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the vSwitch is the default vSwitch. Valid values: **true**: The vSwitch is the default vSwitch. **false**: The vSwitch is not the default vSwitch.
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The zone ID of the vSwitch.
	IzNo *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	// The region ID of the vSwitch.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The state of the vSwitch. Valid values: **Pending**: the vSwitch is being configured. **Available**: the vSwitch is available.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitch) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetAliUid(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.AliUid = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetBid(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Bid = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetCidrBlock(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetGmtCreate(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetGmtModified(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.GmtModified = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetIsDefault(v bool) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.IsDefault = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetIzNo(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.IzNo = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetRegionNo(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.RegionNo = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetStatus(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetVSwitchId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetVSwitchName(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.VSwitchName = &v
	return s
}

type DescribeVSwitchesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVSwitchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVSwitchesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVSwitchesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponse) SetHeaders(v map[string]*string) *DescribeVSwitchesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVSwitchesResponse) SetStatusCode(v int32) *DescribeVSwitchesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVSwitchesResponse) SetBody(v *DescribeVSwitchesResponseBody) *DescribeVSwitchesResponse {
	s.Body = v
	return s
}

type DetachUserENIRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query cluster IDs.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DetachUserENIRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachUserENIRequest) GoString() string {
	return s.String()
}

func (s *DetachUserENIRequest) SetDBClusterId(v string) *DetachUserENIRequest {
	s.DBClusterId = &v
	return s
}

func (s *DetachUserENIRequest) SetOwnerAccount(v string) *DetachUserENIRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DetachUserENIRequest) SetOwnerId(v int64) *DetachUserENIRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachUserENIRequest) SetResourceOwnerAccount(v string) *DetachUserENIRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachUserENIRequest) SetResourceOwnerId(v int64) *DetachUserENIRequest {
	s.ResourceOwnerId = &v
	return s
}

type DetachUserENIResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachUserENIResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachUserENIResponseBody) GoString() string {
	return s.String()
}

func (s *DetachUserENIResponseBody) SetRequestId(v string) *DetachUserENIResponseBody {
	s.RequestId = &v
	return s
}

type DetachUserENIResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachUserENIResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachUserENIResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachUserENIResponse) GoString() string {
	return s.String()
}

func (s *DetachUserENIResponse) SetHeaders(v map[string]*string) *DetachUserENIResponse {
	s.Headers = v
	return s
}

func (s *DetachUserENIResponse) SetStatusCode(v int32) *DetachUserENIResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachUserENIResponse) SetBody(v *DetachUserENIResponseBody) *DetachUserENIResponse {
	s.Body = v
	return s
}

type DisableAdviceServiceRequest struct {
	// The ID of the cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of Data Warehouse Edition (V3.0) clusters.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableAdviceServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableAdviceServiceRequest) GoString() string {
	return s.String()
}

func (s *DisableAdviceServiceRequest) SetDBClusterId(v string) *DisableAdviceServiceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DisableAdviceServiceRequest) SetRegionId(v string) *DisableAdviceServiceRequest {
	s.RegionId = &v
	return s
}

type DisableAdviceServiceResponseBody struct {
	// The message returned for the operation. Valid values:
	//
	// *   **Success** is returned if the operation is successful.
	// *   An error message is returned if the operation fails.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAdviceServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableAdviceServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAdviceServiceResponseBody) SetMessage(v string) *DisableAdviceServiceResponseBody {
	s.Message = &v
	return s
}

func (s *DisableAdviceServiceResponseBody) SetRequestId(v string) *DisableAdviceServiceResponseBody {
	s.RequestId = &v
	return s
}

type DisableAdviceServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableAdviceServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableAdviceServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableAdviceServiceResponse) GoString() string {
	return s.String()
}

func (s *DisableAdviceServiceResponse) SetHeaders(v map[string]*string) *DisableAdviceServiceResponse {
	s.Headers = v
	return s
}

func (s *DisableAdviceServiceResponse) SetStatusCode(v int32) *DisableAdviceServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAdviceServiceResponse) SetBody(v *DisableAdviceServiceResponseBody) *DisableAdviceServiceResponse {
	s.Body = v
	return s
}

type DownloadDiagnosisRecordsRequest struct {
	// The source IP addresses.
	//
	// > You can call the [DescribeDiagnosisDimensions](~~308210~~) operation to query the resource group, database name, username, and source IP address of the SQL statements to be queried.
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database on which the SQL statements are executed.
	//
	// > You can call the [DescribeDiagnosisDimensions](~~308210~~) operation to query the resource group, database name, username, and source IP address of the SQL statements to be queried.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >
	//
	// *   The end time must be later than the start time.
	//
	// *   The maximum time range that can be specified is 24 hours.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword that is used for the query.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of file titles and error messages. Valid values:
	//
	// *   **zh** (default): simplified Chinese.
	// *   **en**: English.
	// *   **ja**: Japanese.
	// *   **zh-tw**: traditional Chinese.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum peak memory of the SQL statements. Unit: bytes.
	MaxPeakMemory *int64 `json:"MaxPeakMemory,omitempty" xml:"MaxPeakMemory,omitempty"`
	// The maximum scan size of the SQL statements. Unit: bytes.
	MaxScanSize *int64 `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	// The minimum peak memory of the SQL statements. Unit: bytes.
	MinPeakMemory *int64 `json:"MinPeakMemory,omitempty" xml:"MinPeakMemory,omitempty"`
	// The minimum scan size of the SQL statements. Unit: bytes.
	MinScanSize *int64 `json:"MinScanSize,omitempty" xml:"MinScanSize,omitempty"`
	// The SQL query condition, which can be a combination of the `Type` and `Value` fields or a combination of the Type, `Min`, and `Max` fields. Specify the condition in the JSON format. `Type` specifies the SQL query dimension. Valid values for Type: `maxCost`, `status`, and `cost`. `Value`, `Min`, or `Max` specifies the SQL query range for the dimension. Valid values:
	//
	// *   `{"Type":"maxCost","Value":"100"}`: queries the top 100 most time-consuming SQL statements. Set `Value` to 100.
	// *   `{"Type":"status","Value":"finished"}`: queries executed SQL statements. You can set `Value` to `running` to query SQL statements that are being executed. You can also set Value to `failed` to query SQL statements that failed to be executed.
	// *   `{"Type":"cost","Min":"10","Max":"200"}`: queries SQL statements whose execution durations are in the range of 10 to 200 milliseconds. You can also customize the maximum and minimum execution durations.
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group to which the SQL statements belong.
	//
	// > You can call the [DescribeDiagnosisDimensions](~~308210~~) operation to query the resource group, database name, username, and source IP address of the SQL statements to be queried.
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The beginning of the time range to query. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// > Only data within the last 14 days can be queried.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The username that is used to execute the SQL statements.
	//
	// > You can call the [DescribeDiagnosisDimensions](~~308210~~) operation to query the resource group, database name, username, and source IP address of the SQL statements to be queried.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DownloadDiagnosisRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadDiagnosisRecordsRequest) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsRequest) SetClientIp(v string) *DownloadDiagnosisRecordsRequest {
	s.ClientIp = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetDBClusterId(v string) *DownloadDiagnosisRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetDatabase(v string) *DownloadDiagnosisRecordsRequest {
	s.Database = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetEndTime(v string) *DownloadDiagnosisRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetKeyword(v string) *DownloadDiagnosisRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetLang(v string) *DownloadDiagnosisRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMaxPeakMemory(v int64) *DownloadDiagnosisRecordsRequest {
	s.MaxPeakMemory = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMaxScanSize(v int64) *DownloadDiagnosisRecordsRequest {
	s.MaxScanSize = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMinPeakMemory(v int64) *DownloadDiagnosisRecordsRequest {
	s.MinPeakMemory = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetMinScanSize(v int64) *DownloadDiagnosisRecordsRequest {
	s.MinScanSize = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetQueryCondition(v string) *DownloadDiagnosisRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetRegionId(v string) *DownloadDiagnosisRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetResourceGroup(v string) *DownloadDiagnosisRecordsRequest {
	s.ResourceGroup = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetStartTime(v string) *DownloadDiagnosisRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetUserName(v string) *DownloadDiagnosisRecordsRequest {
	s.UserName = &v
	return s
}

type DownloadDiagnosisRecordsResponseBody struct {
	// The ID of the download task.
	DownloadId *int32 `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadDiagnosisRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadDiagnosisRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsResponseBody) SetDownloadId(v int32) *DownloadDiagnosisRecordsResponseBody {
	s.DownloadId = &v
	return s
}

func (s *DownloadDiagnosisRecordsResponseBody) SetRequestId(v string) *DownloadDiagnosisRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DownloadDiagnosisRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DownloadDiagnosisRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadDiagnosisRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadDiagnosisRecordsResponse) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsResponse) SetHeaders(v map[string]*string) *DownloadDiagnosisRecordsResponse {
	s.Headers = v
	return s
}

func (s *DownloadDiagnosisRecordsResponse) SetStatusCode(v int32) *DownloadDiagnosisRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadDiagnosisRecordsResponse) SetBody(v *DownloadDiagnosisRecordsResponseBody) *DownloadDiagnosisRecordsResponse {
	s.Body = v
	return s
}

type EnableAdviceServiceRequest struct {
	// The ID of the cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of Data Warehouse Edition (V3.0) clusters.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableAdviceServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableAdviceServiceRequest) GoString() string {
	return s.String()
}

func (s *EnableAdviceServiceRequest) SetDBClusterId(v string) *EnableAdviceServiceRequest {
	s.DBClusterId = &v
	return s
}

func (s *EnableAdviceServiceRequest) SetRegionId(v string) *EnableAdviceServiceRequest {
	s.RegionId = &v
	return s
}

type EnableAdviceServiceResponseBody struct {
	// The message returned for the operation. Valid values:
	//
	// *   **Success** is returned if the operation is successful.
	// *   An error message is returned if the operation fails.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAdviceServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableAdviceServiceResponseBody) GoString() string {
	return s.String()
}

func (s *EnableAdviceServiceResponseBody) SetMessage(v string) *EnableAdviceServiceResponseBody {
	s.Message = &v
	return s
}

func (s *EnableAdviceServiceResponseBody) SetRequestId(v string) *EnableAdviceServiceResponseBody {
	s.RequestId = &v
	return s
}

type EnableAdviceServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableAdviceServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableAdviceServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableAdviceServiceResponse) GoString() string {
	return s.String()
}

func (s *EnableAdviceServiceResponse) SetHeaders(v map[string]*string) *EnableAdviceServiceResponse {
	s.Headers = v
	return s
}

func (s *EnableAdviceServiceResponse) SetStatusCode(v int32) *EnableAdviceServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableAdviceServiceResponse) SetBody(v *EnableAdviceServiceResponseBody) *EnableAdviceServiceResponse {
	s.Body = v
	return s
}

type GrantOperatorPermissionRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The expiration time of the service account permissions. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	ExpiredTime  *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the permissions. Valid values:
	//
	// *   **Control**: configuration permissions. The service account is granted permissions to query and modify cluster configurations.
	// *   **Data**: data permissions. The service account is granted permissions to query schemas, indexes, and SQL statements.
	Privileges           *string `json:"Privileges,omitempty" xml:"Privileges,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GrantOperatorPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantOperatorPermissionRequest) GoString() string {
	return s.String()
}

func (s *GrantOperatorPermissionRequest) SetDBClusterId(v string) *GrantOperatorPermissionRequest {
	s.DBClusterId = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetExpiredTime(v string) *GrantOperatorPermissionRequest {
	s.ExpiredTime = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetOwnerAccount(v string) *GrantOperatorPermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetOwnerId(v int64) *GrantOperatorPermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetPrivileges(v string) *GrantOperatorPermissionRequest {
	s.Privileges = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetResourceOwnerAccount(v string) *GrantOperatorPermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetResourceOwnerId(v int64) *GrantOperatorPermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

type GrantOperatorPermissionResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantOperatorPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantOperatorPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GrantOperatorPermissionResponseBody) SetRequestId(v string) *GrantOperatorPermissionResponseBody {
	s.RequestId = &v
	return s
}

type GrantOperatorPermissionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GrantOperatorPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantOperatorPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantOperatorPermissionResponse) GoString() string {
	return s.String()
}

func (s *GrantOperatorPermissionResponse) SetHeaders(v map[string]*string) *GrantOperatorPermissionResponse {
	s.Headers = v
	return s
}

func (s *GrantOperatorPermissionResponse) SetStatusCode(v int32) *GrantOperatorPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantOperatorPermissionResponse) SetBody(v *GrantOperatorPermissionResponseBody) *GrantOperatorPermissionResponse {
	s.Body = v
	return s
}

type KillProcessRequest struct {
	// The ID of the cluster.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The unique ID of the process. You can call the [DescribeProcessList](~~190092~~) operation to obtain the ID.
	ProcessId            *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s KillProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s KillProcessRequest) GoString() string {
	return s.String()
}

func (s *KillProcessRequest) SetDBClusterId(v string) *KillProcessRequest {
	s.DBClusterId = &v
	return s
}

func (s *KillProcessRequest) SetOwnerAccount(v string) *KillProcessRequest {
	s.OwnerAccount = &v
	return s
}

func (s *KillProcessRequest) SetOwnerId(v int64) *KillProcessRequest {
	s.OwnerId = &v
	return s
}

func (s *KillProcessRequest) SetProcessId(v string) *KillProcessRequest {
	s.ProcessId = &v
	return s
}

func (s *KillProcessRequest) SetResourceOwnerAccount(v string) *KillProcessRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *KillProcessRequest) SetResourceOwnerId(v int64) *KillProcessRequest {
	s.ResourceOwnerId = &v
	return s
}

type KillProcessResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KillProcessResponseBody) GoString() string {
	return s.String()
}

func (s *KillProcessResponseBody) SetRequestId(v string) *KillProcessResponseBody {
	s.RequestId = &v
	return s
}

type KillProcessResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *KillProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s KillProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s KillProcessResponse) GoString() string {
	return s.String()
}

func (s *KillProcessResponse) SetHeaders(v map[string]*string) *KillProcessResponse {
	s.Headers = v
	return s
}

func (s *KillProcessResponse) SetStatusCode(v int32) *KillProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *KillProcessResponse) SetBody(v *KillProcessResponseBody) *KillProcessResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster. You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The cluster ID. You can specify multiple cluster IDs. Valid values of N: 1 to 50.
	//
	// > You must specify at least one of the following parameters: ResourceId.N and Tag.N.Key.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Set the value to **cluster**.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags that are added to clusters.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerId(v int64) *ListTagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The key of the tag. You can specify multiple tag keys. The tag key cannot be an empty string. Valid values of N: 1 to 20.
	//
	// > You must specify at least one of the following parameters: ResourceId.N and Tag.N.Key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. You can specify multiple tag values. The tag value can be an empty string. Valid values of N: 1 to 20.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results. You must specify the token that is obtained from the previous query as the value of NextToken.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags that are added to clusters.
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	// The cluster ID.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. A value of cluster indicates an AnalyticDB for MySQL cluster.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type MigrateDBClusterRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s MigrateDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s MigrateDBClusterRequest) GoString() string {
	return s.String()
}

func (s *MigrateDBClusterRequest) SetDBClusterId(v string) *MigrateDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *MigrateDBClusterRequest) SetOwnerAccount(v string) *MigrateDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MigrateDBClusterRequest) SetOwnerId(v int64) *MigrateDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *MigrateDBClusterRequest) SetResourceOwnerAccount(v string) *MigrateDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MigrateDBClusterRequest) SetResourceOwnerId(v int64) *MigrateDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

type MigrateDBClusterResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MigrateDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateDBClusterResponseBody) SetRequestId(v string) *MigrateDBClusterResponseBody {
	s.RequestId = &v
	return s
}

type MigrateDBClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MigrateDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MigrateDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s MigrateDBClusterResponse) GoString() string {
	return s.String()
}

func (s *MigrateDBClusterResponse) SetHeaders(v map[string]*string) *MigrateDBClusterResponse {
	s.Headers = v
	return s
}

func (s *MigrateDBClusterResponse) SetStatusCode(v int32) *MigrateDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateDBClusterResponse) SetBody(v *MigrateDBClusterResponseBody) *MigrateDBClusterResponse {
	s.Body = v
	return s
}

type ModifyAccountDescriptionRequest struct {
	// The description of the account. The description must meet the following requirements:
	//
	// *   The description must start with a letter.
	// *   The description can contain letters, digits, underscores (\_), and hyphens (-).
	// *   The description cannot start with `http://` or `https://`.
	// *   The description must be 2 to 256 characters in length.
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) SetAccountDescription(v string) *ModifyAccountDescriptionRequest {
	s.AccountDescription = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDBClusterId(v string) *ModifyAccountDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetOwnerAccount(v string) *ModifyAccountDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetOwnerId(v int64) *ModifyAccountDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyAccountDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetResourceOwnerId(v int64) *ModifyAccountDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyAccountDescriptionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponseBody) SetRequestId(v string) *ModifyAccountDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccountDescriptionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAccountDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccountDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponse) SetHeaders(v map[string]*string) *ModifyAccountDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountDescriptionResponse) SetStatusCode(v int32) *ModifyAccountDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountDescriptionResponse) SetBody(v *ModifyAccountDescriptionResponseBody) *ModifyAccountDescriptionResponse {
	s.Body = v
	return s
}

type ModifyAuditLogConfigRequest struct {
	// The status of SQL audit. Valid values:
	//
	// *   **on**: SQL audit is enabled.
	// *   **off**: SQL audit is disabled.
	AuditLogStatus *string `json:"AuditLogStatus,omitempty" xml:"AuditLogStatus,omitempty"`
	// The ID of the cluster.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAuditLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditLogConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogConfigRequest) SetAuditLogStatus(v string) *ModifyAuditLogConfigRequest {
	s.AuditLogStatus = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetDBClusterId(v string) *ModifyAuditLogConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetOwnerAccount(v string) *ModifyAuditLogConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetOwnerId(v int64) *ModifyAuditLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetRegionId(v string) *ModifyAuditLogConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetResourceOwnerAccount(v string) *ModifyAuditLogConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetResourceOwnerId(v int64) *ModifyAuditLogConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyAuditLogConfigResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the status of SQL audit is updated. Valid values:
	//
	// *   **true**: The status of SQL audit is updated.
	// *   **false**: The status of SQL audit is not updated.
	UpdateSucceed *bool `json:"UpdateSucceed,omitempty" xml:"UpdateSucceed,omitempty"`
}

func (s ModifyAuditLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogConfigResponseBody) SetRequestId(v string) *ModifyAuditLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAuditLogConfigResponseBody) SetUpdateSucceed(v bool) *ModifyAuditLogConfigResponseBody {
	s.UpdateSucceed = &v
	return s
}

type ModifyAuditLogConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAuditLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAuditLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAuditLogConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogConfigResponse) SetHeaders(v map[string]*string) *ModifyAuditLogConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyAuditLogConfigResponse) SetStatusCode(v int32) *ModifyAuditLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAuditLogConfigResponse) SetBody(v *ModifyAuditLogConfigResponseBody) *ModifyAuditLogConfigResponse {
	s.Body = v
	return s
}

type ModifyAutoRenewAttributeRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The renewal duration. Default value: **1**.
	//
	// *   Valid values when PeriodUnit is set to **Month**: 1 to 11. Data type: INTEGER.
	// *   Valid values when PeriodUnit is set to **Year**: 1, 2, 3, and 5. Data type: INTEGER.
	//
	// > Longer subscription durations offer more savings. Purchasing a cluster for one year is more cost-effective than purchasing the cluster for 10 or 11 months.
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The unit of the renewal period. Default value: **Month**. Valid values:
	//
	// *   **Year**
	// *   **Month**
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The renewal status of the cluster. Valid values:
	//
	// *   **AutoRenewal**: The cluster is automatically renewed.
	// *   **Normal**: The cluster is manually renewed.
	// *   **NotRenewal**: The cluster is not renewed.
	RenewalStatus        *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewAttributeRequest) SetDBClusterId(v string) *ModifyAutoRenewAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetDuration(v string) *ModifyAutoRenewAttributeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetOwnerAccount(v string) *ModifyAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetOwnerId(v int64) *ModifyAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetPeriodUnit(v string) *ModifyAutoRenewAttributeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetRegionId(v string) *ModifyAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetRenewalStatus(v string) *ModifyAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *ModifyAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *ModifyAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyAutoRenewAttributeResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAutoRenewAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewAttributeResponseBody) SetRequestId(v string) *ModifyAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAutoRenewAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAutoRenewAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *ModifyAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoRenewAttributeResponse) SetStatusCode(v int32) *ModifyAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAutoRenewAttributeResponse) SetBody(v *ModifyAutoRenewAttributeResponseBody) *ModifyAutoRenewAttributeResponse {
	s.Body = v
	return s
}

type ModifyBackupPolicyRequest struct {
	// The number of days for which to retain full backup files. Valid values: 7 to 730.
	//
	// >  If you leave this parameter empty, the default value 7 is used.
	BackupRetentionPeriod *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The ID of the cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable real-time log backup. Valid values:
	//
	// *   **Enable**
	//
	// *   **Disable**
	//
	// > If you leave this parameter empty, the default value Enable is used.
	EnableBackupLog *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// The number of days for which to retain log backup files. Valid values: 7 to 730.
	//
	// >  If you leave this parameter empty, the default value 7 is used.
	LogBackupRetentionPeriod *int32  `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The days of the week on which to perform full backup. Separate multiple values with commas (,). Valid values:
	//
	// *   **Monday**
	// *   **Tuesday**
	// *   **Wednesday**
	// *   **Thursday**
	// *   **Friday**
	// *   **Saturday**
	// *   **Sunday**
	//
	// >  To ensure data security, we recommend that you specify at least two values.
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The start time of the full backup within a time range. Specify the time range in the HH:mmZ-HH:mmZ format. The time must be in UTC.
	//
	// >  The time range is 1 hour.
	PreferredBackupTime  *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDBClusterId(v string) *ModifyBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableBackupLog(v string) *ModifyBackupPolicyRequest {
	s.EnableBackupLog = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetLogBackupRetentionPeriod(v int32) *ModifyBackupPolicyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyBackupPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponseBody) SetRequestId(v string) *ModifyBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBackupPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupPolicyResponse) SetStatusCode(v int32) *ModifyBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupPolicyResponse) SetBody(v *ModifyBackupPolicyResponseBody) *ModifyBackupPolicyResponse {
	s.Body = v
	return s
}

type ModifyClusterConnectionStringRequest struct {
	// The prefix of public endpoints.
	//
	// *   The prefix can contain lowercase letters, digits, and hyphens (-). It must start with a lowercase letter.
	// *   The prefix can be up to 30 characters in length.
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The current public endpoint of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusterNetInfo](~~143384~~) operation to query the public endpoint of the cluster.
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port number. Set the value to **3306**.
	Port                 *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyClusterConnectionStringRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterConnectionStringRequest) SetConnectionStringPrefix(v string) *ModifyClusterConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyClusterConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetDBClusterId(v string) *ModifyClusterConnectionStringRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetOwnerAccount(v string) *ModifyClusterConnectionStringRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetOwnerId(v int64) *ModifyClusterConnectionStringRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetPort(v int32) *ModifyClusterConnectionStringRequest {
	s.Port = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetResourceOwnerAccount(v string) *ModifyClusterConnectionStringRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyClusterConnectionStringRequest) SetResourceOwnerId(v int64) *ModifyClusterConnectionStringRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyClusterConnectionStringResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterConnectionStringResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConnectionStringResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterConnectionStringResponseBody) SetRequestId(v string) *ModifyClusterConnectionStringResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterConnectionStringResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyClusterConnectionStringResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterConnectionStringResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterConnectionStringResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterConnectionStringResponse) SetHeaders(v map[string]*string) *ModifyClusterConnectionStringResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterConnectionStringResponse) SetStatusCode(v int32) *ModifyClusterConnectionStringResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterConnectionStringResponse) SetBody(v *ModifyClusterConnectionStringResponseBody) *ModifyClusterConnectionStringResponse {
	s.Body = v
	return s
}

type ModifyDBClusterRequest struct {
	// The computing resources of the cluster. You can call the [DescribeAvailableResource](~~190632~~) operation to query the computing resources that are available within a region.
	//
	// > This parameter must be specified when Mode is set to Flexible.
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// *   **Cluster**: reserved mode for Cluster Edition.
	// *   **MixedStorage**: elastic mode for Cluster Edition.
	//
	// > If you set DBClusterCategory to Cluster, you must set Mode to Reserver. If you set DBClusterCategory to MixedStorage, you must set Mode to Flexible. Otherwise, you fail to change the specifications of the cluster.
	DBClusterCategory *string `json:"DBClusterCategory,omitempty" xml:"DBClusterCategory,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The node specifications of the cluster. Valid values:
	//
	// *   **C8**
	// *   **C32**
	//
	// > This parameter must be specified when Mode is set to Reserver.
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of node groups. Valid values: 1 to 200.
	//
	// > This parameter must be specified when Mode is set to Reserver.
	DBNodeGroupCount *string `json:"DBNodeGroupCount,omitempty" xml:"DBNodeGroupCount,omitempty"`
	// The storage capacity per node. Unit: GB.
	//
	// *   Valid values when DBClusterClass is set to C8: 100 to 2000.
	// *   Valid values when DBClusterClass is set to C32: 100 to 8000.
	//
	// >
	//
	// *   This parameter must be specified when Mode is set to Reserver.
	//
	// *   The storage capacity less than 1,000 GB increases in 100 GB increments. The storage capacity greater than 1,000 GB increases in 1,000 GB increments.
	DBNodeStorage *string `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	// The enhanced SSD (ESSD) performance level of the cluster. Valid values:
	//
	// *   PL0
	// *   PL1
	// *   PL2
	// *   PL3
	DiskPerformanceLevel *string `json:"DiskPerformanceLevel,omitempty" xml:"DiskPerformanceLevel,omitempty"`
	// The number of EIUs. The number of EIUs that you can purchase varies based on the single-node EIU specifications.
	//
	// *   If the single-node EIU specifications are 8 cores and 64 GB, you can purchase up to 32 EIUs.
	// *   If the single-node EIU specifications are 12 cores and 96 GB, you can purchase up to 16 EIUs.
	ElasticIOResource *int32 `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	// The single-node specifications of an elastic I/O unit (EIU). Valid values:
	//
	// *   **8Core64GB**: If you set the parameter to **8Core64GB**, the specifications of an EIU are 24 cores and 192 GB memory.
	// *   **12Core96GB**: If you set the parameter to **12Core96GB**, the specifications of an EIU are 36 cores and 288 GB memory.
	//
	// >  This parameter takes effect only when your cluster meets the following requirements:
	//
	// *   The cluster is in elastic mode.
	//
	// *   If the cluster resides in the China (Guangzhou), China (Shenzhen), China (Hangzhou), China (Shanghai), China (Qingdao), China (Beijing), or China (Zhangjiakou) region, the cluster has 16 cores and 64 GB memory or higher specifications.
	//
	// *   If the cluster resides in another region, the cluster has 32 cores and 128 GB memory or higher specifications.
	ElasticIOResourceSize *string `json:"ElasticIOResourceSize,omitempty" xml:"ElasticIOResourceSize,omitempty"`
	// N/A
	ExecutorCount *string `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	// The mode of the cluster. Valid values:
	//
	// *   **Reserver**: the reserved mode.
	// *   **Flexible**: the elastic mode.
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The change type. Valid values:
	//
	// *   **Upgrade**
	// *   **Downgrade**
	ModifyType   *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster. You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// N/A
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
}

func (s ModifyDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterRequest) SetComputeResource(v string) *ModifyDBClusterRequest {
	s.ComputeResource = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterCategory(v string) *ModifyDBClusterRequest {
	s.DBClusterCategory = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterId(v string) *ModifyDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBNodeClass(v string) *ModifyDBClusterRequest {
	s.DBNodeClass = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBNodeGroupCount(v string) *ModifyDBClusterRequest {
	s.DBNodeGroupCount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBNodeStorage(v string) *ModifyDBClusterRequest {
	s.DBNodeStorage = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDiskPerformanceLevel(v string) *ModifyDBClusterRequest {
	s.DiskPerformanceLevel = &v
	return s
}

func (s *ModifyDBClusterRequest) SetElasticIOResource(v int32) *ModifyDBClusterRequest {
	s.ElasticIOResource = &v
	return s
}

func (s *ModifyDBClusterRequest) SetElasticIOResourceSize(v string) *ModifyDBClusterRequest {
	s.ElasticIOResourceSize = &v
	return s
}

func (s *ModifyDBClusterRequest) SetExecutorCount(v string) *ModifyDBClusterRequest {
	s.ExecutorCount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetMode(v string) *ModifyDBClusterRequest {
	s.Mode = &v
	return s
}

func (s *ModifyDBClusterRequest) SetModifyType(v string) *ModifyDBClusterRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyDBClusterRequest) SetOwnerAccount(v string) *ModifyDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetOwnerId(v int64) *ModifyDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetRegionId(v string) *ModifyDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetResourceOwnerId(v int64) *ModifyDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetStorageResource(v string) *ModifyDBClusterRequest {
	s.StorageResource = &v
	return s
}

type ModifyDBClusterResponseBody struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order ID.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponseBody) SetDBClusterId(v string) *ModifyDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterResponseBody) SetOrderId(v string) *ModifyDBClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBClusterResponseBody) SetRequestId(v string) *ModifyDBClusterResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponse) SetHeaders(v map[string]*string) *ModifyDBClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterResponse) SetStatusCode(v int32) *ModifyDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterResponse) SetBody(v *ModifyDBClusterResponseBody) *ModifyDBClusterResponse {
	s.Body = v
	return s
}

type ModifyDBClusterAccessWhiteListRequest struct {
	// The attribute of the IP address whitelist. By default, this parameter is empty. The IP address whitelists that have the **hidden** attribute are not displayed in the console. These IP address whitelists are used to access services such as Data Transmission Service (DTS) and PolarDB-X.
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	// The name of the IP address whitelist that you want to modify. Default value: **Default**. The name of an IP address whitelist must be 2 to 32 characters in length. The name can contain lowercase letters, digits, and underscores (\_). The name must start with a lowercase letter and end with a lowercase letter or digit.
	//
	// Each cluster supports up to 50 IP address whitelists.
	DBClusterIPArrayName *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The method that you want to use to modify the IP address whitelist. Valid values:
	//
	// *   Cover: overwrites the IP address whitelist.
	// *   Append: adds IP addresses to the IP address whitelist.
	// *   Delete: removes IP addresses from the IP address whitelist.
	//
	// Default value: Cover.
	ModifyMode           *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP addresses that you want to use to modify the IP address whitelist of the cluster. Separate multiple IP addresses with commas (,). You can specify up to 500 distinct IP addresses. The following formats are supported:
	//
	// *   IP address. Example: 10.23.12.24.
	// *   CIDR block. Example: 10.23.12.24/24. In this example, 24 indicates that the prefix of the CIDR block is 24 bits in length. You can replace 24 with a value that ranges from 1 to 32.
	//
	// >  This parameter must be specified unless ModifyMode is set to Delete.
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s ModifyDBClusterAccessWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAccessWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetDBClusterIPArrayAttribute(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetDBClusterIPArrayName(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetDBClusterId(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetModifyMode(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetOwnerAccount(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetOwnerId(v int64) *ModifyDBClusterAccessWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetResourceOwnerId(v int64) *ModifyDBClusterAccessWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListRequest) SetSecurityIps(v string) *ModifyDBClusterAccessWhiteListRequest {
	s.SecurityIps = &v
	return s
}

type ModifyDBClusterAccessWhiteListResponseBody struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDBClusterAccessWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) SetDBClusterId(v string) *ModifyDBClusterAccessWhiteListResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) SetRequestId(v string) *ModifyDBClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) SetTaskId(v int32) *ModifyDBClusterAccessWhiteListResponseBody {
	s.TaskId = &v
	return s
}

type ModifyDBClusterAccessWhiteListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBClusterAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterAccessWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAccessWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhiteListResponse) SetHeaders(v map[string]*string) *ModifyDBClusterAccessWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponse) SetStatusCode(v int32) *ModifyDBClusterAccessWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponse) SetBody(v *ModifyDBClusterAccessWhiteListResponseBody) *ModifyDBClusterAccessWhiteListResponse {
	s.Body = v
	return s
}

type ModifyDBClusterDescriptionRequest struct {
	// The description of the cluster.
	//
	// *   The description cannot start with `http://` or `https`.
	// *   The description must be 2 to 256 characters in length.
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterDescription(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterId(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetOwnerAccount(v string) *ModifyDBClusterDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetOwnerId(v int64) *ModifyDBClusterDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetResourceOwnerId(v int64) *ModifyDBClusterDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyDBClusterDescriptionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionResponseBody) SetRequestId(v string) *ModifyDBClusterDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterDescriptionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBClusterDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDBClusterDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterDescriptionResponse) SetStatusCode(v int32) *ModifyDBClusterDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterDescriptionResponse) SetBody(v *ModifyDBClusterDescriptionResponseBody) *ModifyDBClusterDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDBClusterMaintainTimeRequest struct {
	// The ID of cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the details of all AnalyticDB for MySQL clusters within a specific region, including cluster IDs.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The maintenance window of the cluster. It is in the hh:mmZ-hh:mmZ format.
	//
	// >  The maintenance window lasts only 1 hour. Specify the beginning and end of the time range on the hour.
	MaintainTime         *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterMaintainTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeRequest) SetDBClusterId(v string) *ModifyDBClusterMaintainTimeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetMaintainTime(v string) *ModifyDBClusterMaintainTimeRequest {
	s.MaintainTime = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetOwnerAccount(v string) *ModifyDBClusterMaintainTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetOwnerId(v int64) *ModifyDBClusterMaintainTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterMaintainTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetResourceOwnerId(v int64) *ModifyDBClusterMaintainTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyDBClusterMaintainTimeResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterMaintainTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeResponseBody) SetRequestId(v string) *ModifyDBClusterMaintainTimeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterMaintainTimeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBClusterMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterMaintainTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeResponse) SetHeaders(v map[string]*string) *ModifyDBClusterMaintainTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterMaintainTimeResponse) SetStatusCode(v int32) *ModifyDBClusterMaintainTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeResponse) SetBody(v *ModifyDBClusterMaintainTimeResponseBody) *ModifyDBClusterMaintainTimeResponse {
	s.Body = v
	return s
}

type ModifyDBClusterPayTypeRequest struct {
	// The cluster ID.
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// The billing method. Valid values:
	//
	// *   **Postpaid**: pay-as-you-go.
	// *   **Prepaid**: subscription.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription type of the subscription cluster. Valid values:
	//
	// *   **Year**: subscription on a yearly basis.
	// *   **Month**: subscription on a monthly basis.
	//
	// > This parameter must be specified when PayType is set to Prepaid.
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The subscription duration of the subscription cluster.
	//
	// *   Valid values when Period is set to Year: 1, 2, 3, and 5 (integer).
	// *   Valid values when Period is set to Month: 1 to 11 (integer).
	//
	// >
	//
	// *   This parameter must be specified when PayType is set to Prepaid.
	//
	// *   Longer subscription durations offer more savings. Purchasing a cluster for one year is more cost-effective than purchasing the cluster for 10 or 11 months.
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s ModifyDBClusterPayTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterPayTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterPayTypeRequest) SetDbClusterId(v string) *ModifyDBClusterPayTypeRequest {
	s.DbClusterId = &v
	return s
}

func (s *ModifyDBClusterPayTypeRequest) SetPayType(v string) *ModifyDBClusterPayTypeRequest {
	s.PayType = &v
	return s
}

func (s *ModifyDBClusterPayTypeRequest) SetPeriod(v string) *ModifyDBClusterPayTypeRequest {
	s.Period = &v
	return s
}

func (s *ModifyDBClusterPayTypeRequest) SetUsedTime(v string) *ModifyDBClusterPayTypeRequest {
	s.UsedTime = &v
	return s
}

type ModifyDBClusterPayTypeResponseBody struct {
	// The cluster ID.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order ID.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The billing method. Valid values:
	//
	// *   **Postpaid**: pay-as-you-go.
	// *   **Prepaid**: subscription.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterPayTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterPayTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterPayTypeResponseBody) SetDBClusterId(v string) *ModifyDBClusterPayTypeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterPayTypeResponseBody) SetOrderId(v string) *ModifyDBClusterPayTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDBClusterPayTypeResponseBody) SetPayType(v string) *ModifyDBClusterPayTypeResponseBody {
	s.PayType = &v
	return s
}

func (s *ModifyDBClusterPayTypeResponseBody) SetRequestId(v string) *ModifyDBClusterPayTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterPayTypeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBClusterPayTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterPayTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterPayTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterPayTypeResponse) SetHeaders(v map[string]*string) *ModifyDBClusterPayTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterPayTypeResponse) SetStatusCode(v int32) *ModifyDBClusterPayTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterPayTypeResponse) SetBody(v *ModifyDBClusterPayTypeResponseBody) *ModifyDBClusterPayTypeResponse {
	s.Body = v
	return s
}

type ModifyDBClusterResourceGroupRequest struct {
	// The ID of the cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the resource group. For more information, see [View basic information of a resource group](~~151181#task-2398293~~ "This topic describes how to view basic information of a resource group, including the resource group ID, resource group name, and resource group display name.").
	NewResourceGroupId   *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResourceGroupRequest) SetDBClusterId(v string) *ModifyDBClusterResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetNewResourceGroupId(v string) *ModifyDBClusterResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetOwnerAccount(v string) *ModifyDBClusterResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetOwnerId(v int64) *ModifyDBClusterResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetResourceOwnerId(v int64) *ModifyDBClusterResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyDBClusterResourceGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResourceGroupResponseBody) SetRequestId(v string) *ModifyDBClusterResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterResourceGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBClusterResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResourceGroupResponse) SetHeaders(v map[string]*string) *ModifyDBClusterResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterResourceGroupResponse) SetStatusCode(v int32) *ModifyDBClusterResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterResourceGroupResponse) SetBody(v *ModifyDBClusterResourceGroupResponseBody) *ModifyDBClusterResourceGroupResponse {
	s.Body = v
	return s
}

type ModifyDBResourceGroupRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The query execution mode. Valid values:
	//
	// *   **interactive**
	// *   **batch**
	//
	// >  For more information, see [Query execution modes](~~189502~~).
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The number of nodes. Default value: 0.
	//
	// *   Each node is configured with the resources of 16 cores and 64 GB memory.
	// *   Make sure that the amount of resources of the nodes (Number of nodes × 16 cores and 64 GB memory) is less than or equal to the amount of unused resources of the cluster.
	NodeNum              *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequest) SetDBClusterId(v string) *ModifyDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetGroupName(v string) *ModifyDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetGroupType(v string) *ModifyDBResourceGroupRequest {
	s.GroupType = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetNodeNum(v int32) *ModifyDBResourceGroupRequest {
	s.NodeNum = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetOwnerAccount(v string) *ModifyDBResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetOwnerId(v int64) *ModifyDBResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetResourceOwnerAccount(v string) *ModifyDBResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetResourceOwnerId(v int64) *ModifyDBResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyDBResourceGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupResponseBody) SetRequestId(v string) *ModifyDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBResourceGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupResponse) SetHeaders(v map[string]*string) *ModifyDBResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBResourceGroupResponse) SetStatusCode(v int32) *ModifyDBResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBResourceGroupResponse) SetBody(v *ModifyDBResourceGroupResponseBody) *ModifyDBResourceGroupResponse {
	s.Body = v
	return s
}

type ModifyDBResourcePoolRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The number of nodes.
	//
	// *   Each node provides 16 cores and 64 GB memory.
	// *   The amount of resources that you want to add to or remove from the cluster cannot exceed the total amount of resources in the cluster.
	//
	// > - If you do not specify this parameter, the original value is retained.
	// > - You must specify at least one of the QueryType and NodeNum parameters.
	NodeNum      *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the resource group.
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The mode in which SQL statements are executed. Valid values:
	//
	// *   **batch**
	// *   **interactive**
	//
	// > If you do not specify this parameter, the original value is retained.
	QueryType            *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBResourcePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBResourcePoolRequest) SetDBClusterId(v string) *ModifyDBResourcePoolRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetNodeNum(v int32) *ModifyDBResourcePoolRequest {
	s.NodeNum = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetOwnerAccount(v string) *ModifyDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetOwnerId(v int64) *ModifyDBResourcePoolRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetPoolName(v string) *ModifyDBResourcePoolRequest {
	s.PoolName = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetQueryType(v string) *ModifyDBResourcePoolRequest {
	s.QueryType = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetResourceOwnerAccount(v string) *ModifyDBResourcePoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBResourcePoolRequest) SetResourceOwnerId(v int64) *ModifyDBResourcePoolRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyDBResourcePoolResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBResourcePoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBResourcePoolResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBResourcePoolResponseBody) SetRequestId(v string) *ModifyDBResourcePoolResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBResourcePoolResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBResourcePoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBResourcePoolResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBResourcePoolResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBResourcePoolResponse) SetHeaders(v map[string]*string) *ModifyDBResourcePoolResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBResourcePoolResponse) SetStatusCode(v int32) *ModifyDBResourcePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBResourcePoolResponse) SetBody(v *ModifyDBResourcePoolResponseBody) *ModifyDBResourcePoolResponse {
	s.Body = v
	return s
}

type ModifyElasticPlanRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether the scaling plan takes effect. Valid values:
	//
	// *   **true** (default)
	// *   **false**
	ElasticPlanEnable *bool `json:"ElasticPlanEnable,omitempty" xml:"ElasticPlanEnable,omitempty"`
	// The end date of the scaling plan. Specify the date in the yyyy-MM-dd format.
	ElasticPlanEndDay        *string `json:"ElasticPlanEndDay,omitempty" xml:"ElasticPlanEndDay,omitempty"`
	ElasticPlanMonthlyRepeat *string `json:"ElasticPlanMonthlyRepeat,omitempty" xml:"ElasticPlanMonthlyRepeat,omitempty"`
	// The name of the scaling plan.
	//
	// *   The name must be 2 to 30 characters in length.
	// *   The name can contain letters, digits, and underscores (\_).
	//
	// > You can call the [DescribeElasticPlan](~~190596~~) operation to query the information about all scaling plans of a cluster, including the scaling plan names.
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// The number of nodes that are involved in the scaling plan.
	//
	// *   If ElasticPlanType is set to **worker**, you can set this parameter to 0 or leave this parameter empty.
	// *   If ElasticPlanType is set to **executorcombineworker** or **executor**, you must set this parameter to a value that is greater than 0.
	ElasticPlanNodeNum *int32 `json:"ElasticPlanNodeNum,omitempty" xml:"ElasticPlanNodeNum,omitempty"`
	// The start date of the scaling plan. Specify the date in the yyyy-MM-dd format.
	ElasticPlanStartDay *string `json:"ElasticPlanStartDay,omitempty" xml:"ElasticPlanStartDay,omitempty"`
	// The restoration time of the scaling plan. Specify the time on the hour in the HH:mm:ss format. The interval between the scale-up time and the restoration time cannot be more than 24 hours.
	ElasticPlanTimeEnd *string `json:"ElasticPlanTimeEnd,omitempty" xml:"ElasticPlanTimeEnd,omitempty"`
	// The scale-up time of the scaling plan. Specify the time on the hour in the HH:mm:ss format.
	ElasticPlanTimeStart *string `json:"ElasticPlanTimeStart,omitempty" xml:"ElasticPlanTimeStart,omitempty"`
	// The type of the scaling plan. Valid values:
	//
	// *   **worker**: scales only elastic I/O resources.
	// *   **executor**: scales only computing resources.
	// *   **executorcombineworker** (default): scales both elastic I/O resources and computing resources by proportion.
	//
	// >
	//
	// *   If you want to set this parameter to **executorcombineworker**, make sure that the cluster runs a minor version of 3.1.3.2 or later.
	//
	// *   If you want to set this parameter to **worker** or **executor**, make sure that the cluster runs a minor version of 3.1.6.1 or later and a ticket is submitted. After your request is approved, you can set this parameter to **worker** or **executor**.
	ElasticPlanType *string `json:"ElasticPlanType,omitempty" xml:"ElasticPlanType,omitempty"`
	// The days of the week when you want to execute the scaling plan. Valid values: 0 to 6, which indicate Sunday to Saturday. Separate multiple values with commas (,).
	ElasticPlanWeeklyRepeat *string `json:"ElasticPlanWeeklyRepeat,omitempty" xml:"ElasticPlanWeeklyRepeat,omitempty"`
	// The resource specifications that can be scaled up by the scaling plan. Valid values:
	//
	// *   8 Core 64 GB (default)
	// *   16 Core 64 GB
	// *   32 Core 64 GB
	// *   64 Core 128 GB
	// *   12 Core 96 GB
	// *   24 Core 96 GB
	// *   52 Core 86 GB
	ElasticPlanWorkerSpec *string `json:"ElasticPlanWorkerSpec,omitempty" xml:"ElasticPlanWorkerSpec,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the resource group.
	//
	// > You can call the [DescribeDBResourceGroup](~~466685~~) operation to query the resource group name.
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
}

func (s ModifyElasticPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanRequest) SetDBClusterId(v string) *ModifyElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanEnable(v bool) *ModifyElasticPlanRequest {
	s.ElasticPlanEnable = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanEndDay(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanEndDay = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanMonthlyRepeat(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanMonthlyRepeat = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanName(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanNodeNum(v int32) *ModifyElasticPlanRequest {
	s.ElasticPlanNodeNum = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanStartDay(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanStartDay = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanTimeEnd(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanTimeEnd = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanTimeStart(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanTimeStart = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanType(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanType = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanWeeklyRepeat(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanWeeklyRepeat = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanWorkerSpec(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanWorkerSpec = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetOwnerAccount(v string) *ModifyElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetOwnerId(v int64) *ModifyElasticPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetResourceOwnerAccount(v string) *ModifyElasticPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetResourceOwnerId(v int64) *ModifyElasticPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetResourcePoolName(v string) *ModifyElasticPlanRequest {
	s.ResourcePoolName = &v
	return s
}

type ModifyElasticPlanResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyElasticPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanResponseBody) SetRequestId(v string) *ModifyElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

type ModifyElasticPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyElasticPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanResponse) SetHeaders(v map[string]*string) *ModifyElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *ModifyElasticPlanResponse) SetStatusCode(v int32) *ModifyElasticPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyElasticPlanResponse) SetBody(v *ModifyElasticPlanResponseBody) *ModifyElasticPlanResponse {
	s.Body = v
	return s
}

type ModifyLogBackupPolicyRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable log backup. Valid values:
	//
	// *   **Enable**
	// *   **Disable**
	EnableBackupLog *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// The number of days for which to retain backup files. Valid values: 7 to 730.
	//
	// > The default value is 7.
	LogBackupRetentionPeriod *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyLogBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyRequest) SetDBClusterId(v string) *ModifyLogBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetEnableBackupLog(v string) *ModifyLogBackupPolicyRequest {
	s.EnableBackupLog = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetLogBackupRetentionPeriod(v string) *ModifyLogBackupPolicyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetOwnerAccount(v string) *ModifyLogBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetOwnerId(v int64) *ModifyLogBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetResourceGroupId(v string) *ModifyLogBackupPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyLogBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyLogBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyLogBackupPolicyResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLogBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyResponseBody) SetRequestId(v string) *ModifyLogBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLogBackupPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyLogBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLogBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyLogBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyLogBackupPolicyResponse) SetStatusCode(v int32) *ModifyLogBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLogBackupPolicyResponse) SetBody(v *ModifyLogBackupPolicyResponseBody) *ModifyLogBackupPolicyResponse {
	s.Body = v
	return s
}

type ModifyMaintenanceActionRequest struct {
	// The ID of the pending O\&M event. You can specify multiple IDs to batch change the switchover time. Separate multiple IDs with commas (,).
	// > - You can call the [DescribeMaintenanceAction](~~271738~~) operation to query the information about pending O\&M events, including the event ID.
	// > - You can change the switchover time only for pending O\&M events. The switchover time of historical O\&M events cannot be changed. For more information about the status of pending and historical O\&M events, see [DescribeMaintenanceAction](~~271738~~).
	Ids          *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the pending O\&M event occurs.
	//
	// > - You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time when you want the system to perform operations on the pending O\&M event. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s ModifyMaintenanceActionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMaintenanceActionRequest) GoString() string {
	return s.String()
}

func (s *ModifyMaintenanceActionRequest) SetIds(v string) *ModifyMaintenanceActionRequest {
	s.Ids = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetOwnerAccount(v string) *ModifyMaintenanceActionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetOwnerId(v int64) *ModifyMaintenanceActionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetRegionId(v string) *ModifyMaintenanceActionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetResourceGroupId(v string) *ModifyMaintenanceActionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetResourceOwnerAccount(v string) *ModifyMaintenanceActionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetResourceOwnerId(v int64) *ModifyMaintenanceActionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetSwitchTime(v string) *ModifyMaintenanceActionRequest {
	s.SwitchTime = &v
	return s
}

type ModifyMaintenanceActionResponseBody struct {
	// The O\&M event ID.
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMaintenanceActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMaintenanceActionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMaintenanceActionResponseBody) SetIds(v string) *ModifyMaintenanceActionResponseBody {
	s.Ids = &v
	return s
}

func (s *ModifyMaintenanceActionResponseBody) SetRequestId(v string) *ModifyMaintenanceActionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyMaintenanceActionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyMaintenanceActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMaintenanceActionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMaintenanceActionResponse) GoString() string {
	return s.String()
}

func (s *ModifyMaintenanceActionResponse) SetHeaders(v map[string]*string) *ModifyMaintenanceActionResponse {
	s.Headers = v
	return s
}

func (s *ModifyMaintenanceActionResponse) SetStatusCode(v int32) *ModifyMaintenanceActionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMaintenanceActionResponse) SetBody(v *ModifyMaintenanceActionResponseBody) *ModifyMaintenanceActionResponse {
	s.Body = v
	return s
}

type ModifyResubmitConfigRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The job resubmission rules.
	Rules []*ModifyResubmitConfigRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s ModifyResubmitConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyResubmitConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyResubmitConfigRequest) SetDBClusterId(v string) *ModifyResubmitConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyResubmitConfigRequest) SetOwnerAccount(v string) *ModifyResubmitConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyResubmitConfigRequest) SetOwnerId(v int64) *ModifyResubmitConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyResubmitConfigRequest) SetResourceGroupId(v string) *ModifyResubmitConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyResubmitConfigRequest) SetResourceOwnerAccount(v string) *ModifyResubmitConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyResubmitConfigRequest) SetResourceOwnerId(v int64) *ModifyResubmitConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyResubmitConfigRequest) SetRules(v []*ModifyResubmitConfigRequestRules) *ModifyResubmitConfigRequest {
	s.Rules = v
	return s
}

type ModifyResubmitConfigRequestRules struct {
	// Specifies whether to configure out-of-memory (OOM) check.
	ExceedMemoryException *bool `json:"ExceedMemoryException,omitempty" xml:"ExceedMemoryException,omitempty"`
	// The name of the source resource group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The peak memory usage.
	PeakMemory *string `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The duration of the SQL statement. Unit: milliseconds.
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// The name of the destination resource group.
	TargetGroupName *string `json:"TargetGroupName,omitempty" xml:"TargetGroupName,omitempty"`
}

func (s ModifyResubmitConfigRequestRules) String() string {
	return tea.Prettify(s)
}

func (s ModifyResubmitConfigRequestRules) GoString() string {
	return s.String()
}

func (s *ModifyResubmitConfigRequestRules) SetExceedMemoryException(v bool) *ModifyResubmitConfigRequestRules {
	s.ExceedMemoryException = &v
	return s
}

func (s *ModifyResubmitConfigRequestRules) SetGroupName(v string) *ModifyResubmitConfigRequestRules {
	s.GroupName = &v
	return s
}

func (s *ModifyResubmitConfigRequestRules) SetPeakMemory(v string) *ModifyResubmitConfigRequestRules {
	s.PeakMemory = &v
	return s
}

func (s *ModifyResubmitConfigRequestRules) SetQueryTime(v string) *ModifyResubmitConfigRequestRules {
	s.QueryTime = &v
	return s
}

func (s *ModifyResubmitConfigRequestRules) SetTargetGroupName(v string) *ModifyResubmitConfigRequestRules {
	s.TargetGroupName = &v
	return s
}

type ModifyResubmitConfigShrinkRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The job resubmission rules.
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s ModifyResubmitConfigShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyResubmitConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyResubmitConfigShrinkRequest) SetDBClusterId(v string) *ModifyResubmitConfigShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyResubmitConfigShrinkRequest) SetOwnerAccount(v string) *ModifyResubmitConfigShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyResubmitConfigShrinkRequest) SetOwnerId(v int64) *ModifyResubmitConfigShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyResubmitConfigShrinkRequest) SetResourceGroupId(v string) *ModifyResubmitConfigShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyResubmitConfigShrinkRequest) SetResourceOwnerAccount(v string) *ModifyResubmitConfigShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyResubmitConfigShrinkRequest) SetResourceOwnerId(v int64) *ModifyResubmitConfigShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyResubmitConfigShrinkRequest) SetRulesShrink(v string) *ModifyResubmitConfigShrinkRequest {
	s.RulesShrink = &v
	return s
}

type ModifyResubmitConfigResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyResubmitConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyResubmitConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyResubmitConfigResponseBody) SetRequestId(v string) *ModifyResubmitConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyResubmitConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyResubmitConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyResubmitConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyResubmitConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyResubmitConfigResponse) SetHeaders(v map[string]*string) *ModifyResubmitConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyResubmitConfigResponse) SetStatusCode(v int32) *ModifyResubmitConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyResubmitConfigResponse) SetBody(v *ModifyResubmitConfigResponseBody) *ModifyResubmitConfigResponse {
	s.Body = v
	return s
}

type ModifySQAConfigRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// >  You can call the [DescribeDBResourceGroup](~~459446~~) operation to query the resource group name of a cluster.
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable short query acceleration (SQA).
	//
	// Valid values:
	//
	// *   on
	// *   off
	SQAStatus *string `json:"SQAStatus,omitempty" xml:"SQAStatus,omitempty"`
}

func (s ModifySQAConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySQAConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifySQAConfigRequest) SetDBClusterId(v string) *ModifySQAConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifySQAConfigRequest) SetGroupName(v string) *ModifySQAConfigRequest {
	s.GroupName = &v
	return s
}

func (s *ModifySQAConfigRequest) SetOwnerAccount(v string) *ModifySQAConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySQAConfigRequest) SetOwnerId(v int64) *ModifySQAConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySQAConfigRequest) SetResourceGroupId(v string) *ModifySQAConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifySQAConfigRequest) SetResourceOwnerAccount(v string) *ModifySQAConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySQAConfigRequest) SetResourceOwnerId(v int64) *ModifySQAConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySQAConfigRequest) SetSQAStatus(v string) *ModifySQAConfigRequest {
	s.SQAStatus = &v
	return s
}

type ModifySQAConfigResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySQAConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySQAConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySQAConfigResponseBody) SetRequestId(v string) *ModifySQAConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifySQAConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySQAConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySQAConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySQAConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifySQAConfigResponse) SetHeaders(v map[string]*string) *ModifySQAConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifySQAConfigResponse) SetStatusCode(v int32) *ModifySQAConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySQAConfigResponse) SetBody(v *ModifySQAConfigResponseBody) *ModifySQAConfigResponse {
	s.Body = v
	return s
}

type ReleaseClusterPublicConnectionRequest struct {
	// The ID of the cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleaseClusterPublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterPublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseClusterPublicConnectionRequest) SetDBClusterId(v string) *ReleaseClusterPublicConnectionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) SetOwnerAccount(v string) *ReleaseClusterPublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) SetOwnerId(v int64) *ReleaseClusterPublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) SetResourceOwnerAccount(v string) *ReleaseClusterPublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseClusterPublicConnectionRequest) SetResourceOwnerId(v int64) *ReleaseClusterPublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

type ReleaseClusterPublicConnectionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseClusterPublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterPublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseClusterPublicConnectionResponseBody) SetRequestId(v string) *ReleaseClusterPublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseClusterPublicConnectionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseClusterPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseClusterPublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterPublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseClusterPublicConnectionResponse) SetHeaders(v map[string]*string) *ReleaseClusterPublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseClusterPublicConnectionResponse) SetStatusCode(v int32) *ReleaseClusterPublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseClusterPublicConnectionResponse) SetBody(v *ReleaseClusterPublicConnectionResponseBody) *ReleaseClusterPublicConnectionResponse {
	s.Body = v
	return s
}

type ResetAccountPasswordRequest struct {
	// The account of the database.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The account and password of the database.
	//
	// *   The password must contain uppercase letters, lowercase letters, digits, and special characters.
	// *   Special characters include ! @ # $ % ^ & \* () \_ + - and =
	// *   A password must be 8 to 32 characters in length.
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// Normal: standard account
	//
	// Super: privileged account
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The ID of the cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ResetAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRequest) SetAccountName(v string) *ResetAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountPassword(v string) *ResetAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountType(v string) *ResetAccountPasswordRequest {
	s.AccountType = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetDBClusterId(v string) *ResetAccountPasswordRequest {
	s.DBClusterId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetOwnerAccount(v string) *ResetAccountPasswordRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetOwnerId(v int64) *ResetAccountPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetResourceOwnerAccount(v string) *ResetAccountPasswordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetResourceOwnerId(v int64) *ResetAccountPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

type ResetAccountPasswordResponseBody struct {
	// The ID of the cluster.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ResetAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBody) SetDBClusterId(v string) *ResetAccountPasswordResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetRequestId(v string) *ResetAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAccountPasswordResponseBody) SetTaskId(v int32) *ResetAccountPasswordResponseBody {
	s.TaskId = &v
	return s
}

type ResetAccountPasswordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponse) SetHeaders(v map[string]*string) *ResetAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetAccountPasswordResponse) SetStatusCode(v int32) *ResetAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAccountPasswordResponse) SetBody(v *ResetAccountPasswordResponseBody) *ResetAccountPasswordResponse {
	s.Body = v
	return s
}

type RevokeOperatorPermissionRequest struct {
	// The ID of the cluster.
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RevokeOperatorPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeOperatorPermissionRequest) GoString() string {
	return s.String()
}

func (s *RevokeOperatorPermissionRequest) SetDBClusterId(v string) *RevokeOperatorPermissionRequest {
	s.DBClusterId = &v
	return s
}

func (s *RevokeOperatorPermissionRequest) SetOwnerAccount(v string) *RevokeOperatorPermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeOperatorPermissionRequest) SetOwnerId(v int64) *RevokeOperatorPermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *RevokeOperatorPermissionRequest) SetResourceOwnerAccount(v string) *RevokeOperatorPermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RevokeOperatorPermissionRequest) SetResourceOwnerId(v int64) *RevokeOperatorPermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

type RevokeOperatorPermissionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeOperatorPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeOperatorPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeOperatorPermissionResponseBody) SetRequestId(v string) *RevokeOperatorPermissionResponseBody {
	s.RequestId = &v
	return s
}

type RevokeOperatorPermissionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RevokeOperatorPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokeOperatorPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeOperatorPermissionResponse) GoString() string {
	return s.String()
}

func (s *RevokeOperatorPermissionResponse) SetHeaders(v map[string]*string) *RevokeOperatorPermissionResponse {
	s.Headers = v
	return s
}

func (s *RevokeOperatorPermissionResponse) SetStatusCode(v int32) *RevokeOperatorPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeOperatorPermissionResponse) SetBody(v *RevokeOperatorPermissionResponseBody) *RevokeOperatorPermissionResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster. You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the cluster to which to add a tag. If you want to add a tag to multiple clusters, click **Add** and enter the cluster IDs.
	//
	// >
	//
	// *   You can add tags to up to 50 clusters at a time.
	//
	// *   You can call the [DescribeDBClusters](~~129857~~) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the cluster. Set the value to **ALIYUN::ADB::CLUSTER**.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags to add to the cluster.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The key of the tag. If you want to add multiple tags to a single cluster at a time, click **Add** and enter tag keys and values.
	//
	// > You can add up to 20 tags at a time.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. If you want to add multiple tags to a single cluster at a time, click **Add** and enter tag keys and values.
	//
	// > You can add up to 20 tags at a time.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnbindDBResourceGroupWithUserRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](~~129857~~) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The database account with which the resource group is associated.
	GroupUser            *string `json:"GroupUser,omitempty" xml:"GroupUser,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnbindDBResourceGroupWithUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindDBResourceGroupWithUserRequest) GoString() string {
	return s.String()
}

func (s *UnbindDBResourceGroupWithUserRequest) SetDBClusterId(v string) *UnbindDBResourceGroupWithUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetGroupName(v string) *UnbindDBResourceGroupWithUserRequest {
	s.GroupName = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetGroupUser(v string) *UnbindDBResourceGroupWithUserRequest {
	s.GroupUser = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetOwnerAccount(v string) *UnbindDBResourceGroupWithUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetOwnerId(v int64) *UnbindDBResourceGroupWithUserRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetResourceOwnerAccount(v string) *UnbindDBResourceGroupWithUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetResourceOwnerId(v int64) *UnbindDBResourceGroupWithUserRequest {
	s.ResourceOwnerId = &v
	return s
}

type UnbindDBResourceGroupWithUserResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindDBResourceGroupWithUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindDBResourceGroupWithUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDBResourceGroupWithUserResponseBody) SetRequestId(v string) *UnbindDBResourceGroupWithUserResponseBody {
	s.RequestId = &v
	return s
}

type UnbindDBResourceGroupWithUserResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindDBResourceGroupWithUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindDBResourceGroupWithUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindDBResourceGroupWithUserResponse) GoString() string {
	return s.String()
}

func (s *UnbindDBResourceGroupWithUserResponse) SetHeaders(v map[string]*string) *UnbindDBResourceGroupWithUserResponse {
	s.Headers = v
	return s
}

func (s *UnbindDBResourceGroupWithUserResponse) SetStatusCode(v int32) *UnbindDBResourceGroupWithUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserResponse) SetBody(v *UnbindDBResourceGroupWithUserResponseBody) *UnbindDBResourceGroupWithUserResponse {
	s.Body = v
	return s
}

type UnbindDBResourcePoolWithUserRequest struct {
	// The ID of the cluster.
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the resource pool. You cannot unbind users from the default resource pool named USER_DEFAULT.
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The user bound to the resource pool.
	PoolUser             *string `json:"PoolUser,omitempty" xml:"PoolUser,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnbindDBResourcePoolWithUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindDBResourcePoolWithUserRequest) GoString() string {
	return s.String()
}

func (s *UnbindDBResourcePoolWithUserRequest) SetDBClusterId(v string) *UnbindDBResourcePoolWithUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetOwnerAccount(v string) *UnbindDBResourcePoolWithUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetOwnerId(v int64) *UnbindDBResourcePoolWithUserRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetPoolName(v string) *UnbindDBResourcePoolWithUserRequest {
	s.PoolName = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetPoolUser(v string) *UnbindDBResourcePoolWithUserRequest {
	s.PoolUser = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetResourceOwnerAccount(v string) *UnbindDBResourcePoolWithUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetResourceOwnerId(v int64) *UnbindDBResourcePoolWithUserRequest {
	s.ResourceOwnerId = &v
	return s
}

type UnbindDBResourcePoolWithUserResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindDBResourcePoolWithUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindDBResourcePoolWithUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDBResourcePoolWithUserResponseBody) SetRequestId(v string) *UnbindDBResourcePoolWithUserResponseBody {
	s.RequestId = &v
	return s
}

type UnbindDBResourcePoolWithUserResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindDBResourcePoolWithUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindDBResourcePoolWithUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindDBResourcePoolWithUserResponse) GoString() string {
	return s.String()
}

func (s *UnbindDBResourcePoolWithUserResponse) SetHeaders(v map[string]*string) *UnbindDBResourcePoolWithUserResponse {
	s.Headers = v
	return s
}

func (s *UnbindDBResourcePoolWithUserResponse) SetStatusCode(v int32) *UnbindDBResourcePoolWithUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserResponse) SetBody(v *UnbindDBResourcePoolWithUserResponseBody) *UnbindDBResourcePoolWithUserResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from clusters. Default value: false. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// >  If you specify TagKey and this parameter, this parameter does not take effect.
	All          *bool   `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](~~143074~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of cluster N. Valid values of N: 1 to 50.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Set the value to **ALIYUN::ADB::CLUSTER**.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of tag N. Valid values of N: 1 to 20.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerId(v int64) *UntagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":                  tea.String("adb.aliyuncs.com"),
		"cn-beijing":                  tea.String("adb.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("adb.aliyuncs.com"),
		"cn-shanghai":                 tea.String("adb.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("adb.aliyuncs.com"),
		"cn-hongkong":                 tea.String("adb.aliyuncs.com"),
		"ap-southeast-1":              tea.String("adb.aliyuncs.com"),
		"us-west-1":                   tea.String("adb.aliyuncs.com"),
		"us-east-1":                   tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("adb.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("adb.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("adb.ap-northeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("adb.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("adb.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("adb.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("adb.aliyuncs.com"),
		"cn-edge-1":                   tea.String("adb.aliyuncs.com"),
		"cn-fujian":                   tea.String("adb.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("adb.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("adb.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("adb.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("adb.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("adb.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("adb.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("adb.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("adb.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("adb.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("adb.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("adb.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("adb.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("adb.aliyuncs.com"),
		"cn-wuhan":                    tea.String("adb.aliyuncs.com"),
		"cn-yushanfang":               tea.String("adb.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("adb.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("adb.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("adb.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("adb.ap-northeast-1.aliyuncs.com"),
		"me-east-1":                   tea.String("adb.ap-northeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("adb.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("adb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AllocateClusterPublicConnectionWithOptions(request *AllocateClusterPublicConnectionRequest, runtime *util.RuntimeOptions) (_result *AllocateClusterPublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionStringPrefix)) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AllocateClusterPublicConnection"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AllocateClusterPublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocateClusterPublicConnection(request *AllocateClusterPublicConnectionRequest) (_result *AllocateClusterPublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocateClusterPublicConnectionResponse{}
	_body, _err := client.AllocateClusterPublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyAdviceByIdWithOptions(request *ApplyAdviceByIdRequest, runtime *util.RuntimeOptions) (_result *ApplyAdviceByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdviceDate)) {
		query["AdviceDate"] = request.AdviceDate
	}

	if !tea.BoolValue(util.IsUnset(request.AdviceId)) {
		query["AdviceId"] = request.AdviceId
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyAdviceById"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyAdviceByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyAdviceById(request *ApplyAdviceByIdRequest) (_result *ApplyAdviceByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyAdviceByIdResponse{}
	_body, _err := client.ApplyAdviceByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation only for AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters in elastic mode for Cluster Edition.
 *
 * @param request AttachUserENIRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AttachUserENIResponse
 */
func (client *Client) AttachUserENIWithOptions(request *AttachUserENIRequest, runtime *util.RuntimeOptions) (_result *AttachUserENIResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachUserENI"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachUserENIResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation only for AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters in elastic mode for Cluster Edition.
 *
 * @param request AttachUserENIRequest
 * @return AttachUserENIResponse
 */
func (client *Client) AttachUserENI(request *AttachUserENIRequest) (_result *AttachUserENIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachUserENIResponse{}
	_body, _err := client.AttachUserENIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchApplyAdviceByIdListWithOptions(request *BatchApplyAdviceByIdListRequest, runtime *util.RuntimeOptions) (_result *BatchApplyAdviceByIdListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdviceDate)) {
		query["AdviceDate"] = request.AdviceDate
	}

	if !tea.BoolValue(util.IsUnset(request.AdviceIdList)) {
		query["AdviceIdList"] = request.AdviceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchApplyAdviceByIdList"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchApplyAdviceByIdListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchApplyAdviceByIdList(request *BatchApplyAdviceByIdListRequest) (_result *BatchApplyAdviceByIdListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchApplyAdviceByIdListResponse{}
	_body, _err := client.BatchApplyAdviceByIdListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Precautions
 * *   This operation is applicable only for elastic clusters of 32 cores or more.
 * *   The default resource group USER_DEFAULT cannot be associated with a database account.
 *
 * @param request BindDBResourceGroupWithUserRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return BindDBResourceGroupWithUserResponse
 */
func (client *Client) BindDBResourceGroupWithUserWithOptions(request *BindDBResourceGroupWithUserRequest, runtime *util.RuntimeOptions) (_result *BindDBResourceGroupWithUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupUser)) {
		query["GroupUser"] = request.GroupUser
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindDBResourceGroupWithUser"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindDBResourceGroupWithUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Precautions
 * *   This operation is applicable only for elastic clusters of 32 cores or more.
 * *   The default resource group USER_DEFAULT cannot be associated with a database account.
 *
 * @param request BindDBResourceGroupWithUserRequest
 * @return BindDBResourceGroupWithUserResponse
 */
func (client *Client) BindDBResourceGroupWithUser(request *BindDBResourceGroupWithUserRequest) (_result *BindDBResourceGroupWithUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindDBResourceGroupWithUserResponse{}
	_body, _err := client.BindDBResourceGroupWithUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   This operation is available only for AnalyticDB for MySQL clusters in elastic mode for Cluster Edition that have 32 cores or more.
 * *   The default resource group USER_DEFAULT cannot be associated with a database account.
 *
 * @param request BindDBResourcePoolWithUserRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return BindDBResourcePoolWithUserResponse
 */
func (client *Client) BindDBResourcePoolWithUserWithOptions(request *BindDBResourcePoolWithUserRequest, runtime *util.RuntimeOptions) (_result *BindDBResourcePoolWithUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.PoolUser)) {
		query["PoolUser"] = request.PoolUser
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindDBResourcePoolWithUser"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindDBResourcePoolWithUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   This operation is available only for AnalyticDB for MySQL clusters in elastic mode for Cluster Edition that have 32 cores or more.
 * *   The default resource group USER_DEFAULT cannot be associated with a database account.
 *
 * @param request BindDBResourcePoolWithUserRequest
 * @return BindDBResourcePoolWithUserResponse
 */
func (client *Client) BindDBResourcePoolWithUser(request *BindDBResourcePoolWithUserRequest) (_result *BindDBResourcePoolWithUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindDBResourcePoolWithUserResponse{}
	_body, _err := client.BindDBResourcePoolWithUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountDescription)) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccount"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccount(request *CreateAccountRequest) (_result *CreateAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccountResponse{}
	_body, _err := client.CreateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After you create a cluster, you are billed for the cluster specifications that you select. For more information about the billable items and pricing for Data Warehouse Edition (V3.0) clusters, see [Billable items of Data Warehouse Edition (V3.0)](~~303131~~) and [Pricing for Data Warehouse Edition (V3.0)](~~135229~~).
 *
 * @param request CreateDBClusterRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateDBClusterResponse
 */
func (client *Client) CreateDBClusterWithOptions(request *CreateDBClusterRequest, runtime *util.RuntimeOptions) (_result *CreateDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupSetID)) {
		query["BackupSetID"] = request.BackupSetID
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ComputeResource)) {
		query["ComputeResource"] = request.ComputeResource
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterCategory)) {
		query["DBClusterCategory"] = request.DBClusterCategory
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterClass)) {
		query["DBClusterClass"] = request.DBClusterClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterDescription)) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterNetworkType)) {
		query["DBClusterNetworkType"] = request.DBClusterNetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterVersion)) {
		query["DBClusterVersion"] = request.DBClusterVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeGroupCount)) {
		query["DBNodeGroupCount"] = request.DBNodeGroupCount
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeStorage)) {
		query["DBNodeStorage"] = request.DBNodeStorage
	}

	if !tea.BoolValue(util.IsUnset(request.DiskEncryption)) {
		query["DiskEncryption"] = request.DiskEncryption
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticIOResource)) {
		query["ElasticIOResource"] = request.ElasticIOResource
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorCount)) {
		query["ExecutorCount"] = request.ExecutorCount
	}

	if !tea.BoolValue(util.IsUnset(request.KmsId)) {
		query["KmsId"] = request.KmsId
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreTime)) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreType)) {
		query["RestoreType"] = request.RestoreType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDBInstanceName)) {
		query["SourceDBInstanceName"] = request.SourceDBInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.StorageResource)) {
		query["StorageResource"] = request.StorageResource
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.UsedTime)) {
		query["UsedTime"] = request.UsedTime
	}

	if !tea.BoolValue(util.IsUnset(request.VPCId)) {
		query["VPCId"] = request.VPCId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBCluster"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After you create a cluster, you are billed for the cluster specifications that you select. For more information about the billable items and pricing for Data Warehouse Edition (V3.0) clusters, see [Billable items of Data Warehouse Edition (V3.0)](~~303131~~) and [Pricing for Data Warehouse Edition (V3.0)](~~135229~~).
 *
 * @param request CreateDBClusterRequest
 * @return CreateDBClusterResponse
 */
func (client *Client) CreateDBCluster(request *CreateDBClusterRequest) (_result *CreateDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CreateDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Precautions
 * This operation is applicable only for elastic clusters of 32 cores or more.
 *
 * @param request CreateDBResourceGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateDBResourceGroupResponse
 */
func (client *Client) CreateDBResourceGroupWithOptions(request *CreateDBResourceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateDBResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["GroupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeNum)) {
		query["NodeNum"] = request.NodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBResourceGroup"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Precautions
 * This operation is applicable only for elastic clusters of 32 cores or more.
 *
 * @param request CreateDBResourceGroupRequest
 * @return CreateDBResourceGroupResponse
 */
func (client *Client) CreateDBResourceGroup(request *CreateDBResourceGroupRequest) (_result *CreateDBResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBResourceGroupResponse{}
	_body, _err := client.CreateDBResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This operation is applicable only for elastic clusters of 32 cores or more.
 *
 * @param request CreateDBResourcePoolRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateDBResourcePoolResponse
 */
func (client *Client) CreateDBResourcePoolWithOptions(request *CreateDBResourcePoolRequest, runtime *util.RuntimeOptions) (_result *CreateDBResourcePoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeNum)) {
		query["NodeNum"] = request.NodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBResourcePool"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBResourcePoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation is applicable only for elastic clusters of 32 cores or more.
 *
 * @param request CreateDBResourcePoolRequest
 * @return CreateDBResourcePoolResponse
 */
func (client *Client) CreateDBResourcePool(request *CreateDBResourcePoolRequest) (_result *CreateDBResourcePoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBResourcePoolResponse{}
	_body, _err := client.CreateDBResourcePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * You can call this operation only for AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters in elastic mode for Cluster Edition that have 32 cores or more.
 *
 * @param request CreateElasticPlanRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateElasticPlanResponse
 */
func (client *Client) CreateElasticPlanWithOptions(request *CreateElasticPlanRequest, runtime *util.RuntimeOptions) (_result *CreateElasticPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanEnable)) {
		query["ElasticPlanEnable"] = request.ElasticPlanEnable
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanEndDay)) {
		query["ElasticPlanEndDay"] = request.ElasticPlanEndDay
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanMonthlyRepeat)) {
		query["ElasticPlanMonthlyRepeat"] = request.ElasticPlanMonthlyRepeat
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanNodeNum)) {
		query["ElasticPlanNodeNum"] = request.ElasticPlanNodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanStartDay)) {
		query["ElasticPlanStartDay"] = request.ElasticPlanStartDay
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanTimeEnd)) {
		query["ElasticPlanTimeEnd"] = request.ElasticPlanTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanTimeStart)) {
		query["ElasticPlanTimeStart"] = request.ElasticPlanTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanType)) {
		query["ElasticPlanType"] = request.ElasticPlanType
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanWeeklyRepeat)) {
		query["ElasticPlanWeeklyRepeat"] = request.ElasticPlanWeeklyRepeat
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanWorkerSpec)) {
		query["ElasticPlanWorkerSpec"] = request.ElasticPlanWorkerSpec
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePoolName)) {
		query["ResourcePoolName"] = request.ResourcePoolName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateElasticPlan"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * You can call this operation only for AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters in elastic mode for Cluster Edition that have 32 cores or more.
 *
 * @param request CreateElasticPlanRequest
 * @return CreateElasticPlanResponse
 */
func (client *Client) CreateElasticPlan(request *CreateElasticPlanRequest) (_result *CreateElasticPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateElasticPlanResponse{}
	_body, _err := client.CreateElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccount"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccount(request *DeleteAccountRequest) (_result *DeleteAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DeleteAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Subscription clusters cannot be deleted by using API operations. After expiration, subscription clusters are automatically released. If you no longer need a cluster, you can submit a request to unsubscribe from the cluster in the Billing Management console. For more information about cluster refunds, see [Refund policy](~~471477~~).
 * *   After you delete a cluster, resources of the cluster are immediately released, and data of the cluster is no longer retained and cannot be recovered. Proceed with caution.
 *
 * @param request DeleteDBClusterRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteDBClusterResponse
 */
func (client *Client) DeleteDBClusterWithOptions(request *DeleteDBClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBCluster"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Subscription clusters cannot be deleted by using API operations. After expiration, subscription clusters are automatically released. If you no longer need a cluster, you can submit a request to unsubscribe from the cluster in the Billing Management console. For more information about cluster refunds, see [Refund policy](~~471477~~).
 * *   After you delete a cluster, resources of the cluster are immediately released, and data of the cluster is no longer retained and cannot be recovered. Proceed with caution.
 *
 * @param request DeleteDBClusterRequest
 * @return DeleteDBClusterResponse
 */
func (client *Client) DeleteDBCluster(request *DeleteDBClusterRequest) (_result *DeleteDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.DeleteDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Precautions
 * *   You can call this operation only for AnalyticDB for MySQL clusters in elastic mode for Cluster Edition that have 32 cores or more.
 * *   The default resource group USER_DEFAULT cannot be deleted.
 *
 * @param request DeleteDBResourceGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteDBResourceGroupResponse
 */
func (client *Client) DeleteDBResourceGroupWithOptions(request *DeleteDBResourceGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteDBResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBResourceGroup"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDBResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Precautions
 * *   You can call this operation only for AnalyticDB for MySQL clusters in elastic mode for Cluster Edition that have 32 cores or more.
 * *   The default resource group USER_DEFAULT cannot be deleted.
 *
 * @param request DeleteDBResourceGroupRequest
 * @return DeleteDBResourceGroupResponse
 */
func (client *Client) DeleteDBResourceGroup(request *DeleteDBResourceGroupRequest) (_result *DeleteDBResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBResourceGroupResponse{}
	_body, _err := client.DeleteDBResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * **Precautions**
 * *   This operation is available only for AnalyticDB for MySQL clusters in elastic mode for Cluster Edition that have 32 cores or more.
 * *   The default resource group USER_DEFAULT cannot be deleted.
 *
 * @param request DeleteDBResourcePoolRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteDBResourcePoolResponse
 */
func (client *Client) DeleteDBResourcePoolWithOptions(request *DeleteDBResourcePoolRequest, runtime *util.RuntimeOptions) (_result *DeleteDBResourcePoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBResourcePool"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDBResourcePoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * **Precautions**
 * *   This operation is available only for AnalyticDB for MySQL clusters in elastic mode for Cluster Edition that have 32 cores or more.
 * *   The default resource group USER_DEFAULT cannot be deleted.
 *
 * @param request DeleteDBResourcePoolRequest
 * @return DeleteDBResourcePoolResponse
 */
func (client *Client) DeleteDBResourcePool(request *DeleteDBResourcePoolRequest) (_result *DeleteDBResourcePoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBResourcePoolResponse{}
	_body, _err := client.DeleteDBResourcePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteElasticPlanWithOptions(request *DeleteElasticPlanRequest, runtime *util.RuntimeOptions) (_result *DeleteElasticPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteElasticPlan"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteElasticPlan(request *DeleteElasticPlanRequest) (_result *DeleteElasticPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteElasticPlanResponse{}
	_body, _err := client.DeleteElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountsWithOptions(request *DescribeAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccounts"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccounts(request *DescribeAccountsRequest) (_result *DescribeAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DescribeAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAdviceServiceEnabledWithOptions(request *DescribeAdviceServiceEnabledRequest, runtime *util.RuntimeOptions) (_result *DescribeAdviceServiceEnabledResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAdviceServiceEnabled"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAdviceServiceEnabledResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAdviceServiceEnabled(request *DescribeAdviceServiceEnabledRequest) (_result *DescribeAdviceServiceEnabledResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAdviceServiceEnabledResponse{}
	_body, _err := client.DescribeAdviceServiceEnabledWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAllAccountsWithOptions(request *DescribeAllAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAllAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAllAccounts"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAllAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAllAccounts(request *DescribeAllAccountsRequest) (_result *DescribeAllAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAllAccountsResponse{}
	_body, _err := client.DescribeAllAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAllDataSourceWithOptions(request *DescribeAllDataSourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAllDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAllDataSource"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAllDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAllDataSource(request *DescribeAllDataSourceRequest) (_result *DescribeAllDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAllDataSourceResponse{}
	_body, _err := client.DescribeAllDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppliedAdvicesWithOptions(request *DescribeAppliedAdvicesRequest, runtime *util.RuntimeOptions) (_result *DescribeAppliedAdvicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppliedAdvices"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppliedAdvicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppliedAdvices(request *DescribeAppliedAdvicesRequest) (_result *DescribeAppliedAdvicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppliedAdvicesResponse{}
	_body, _err := client.DescribeAppliedAdvicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAuditLogConfigWithOptions(request *DescribeAuditLogConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeAuditLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAuditLogConfig"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAuditLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuditLogConfig(request *DescribeAuditLogConfigRequest) (_result *DescribeAuditLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditLogConfigResponse{}
	_body, _err := client.DescribeAuditLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you call the DescribeAuditLogRecords operation to query the SQL audit logs of an AnalyticDB for MySQL cluster, you must enable SQL audit for the cluster. You can call the [DescribeAuditLogConfig](~~190628~~) operation to query the status of SQL audit. If SQL audit is disabled, you can call the [ModifyAuditLogConfig](~~190629~~) operation to enable SQL audit.
 * SQL audit logs can be queried only when SQL audit is enabled. Only SQL audit logs within the last 30 days can be queried. If SQL audit was disabled and re-enabled, only SQL audit logs from the time when SQL audit was re-enabled can be queried. The following operations are not recorded in SQL audit logs: **INSERT INTO VALUES**, **REPLACE INTO VALUES**, and **UPSERT INTO VALUES**.
 *
 * @param request DescribeAuditLogRecordsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeAuditLogRecordsResponse
 */
func (client *Client) DescribeAuditLogRecordsWithOptions(request *DescribeAuditLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeAuditLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		query["DBName"] = request.DBName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.HostAddress)) {
		query["HostAddress"] = request.HostAddress
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeyword)) {
		query["QueryKeyword"] = request.QueryKeyword
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SqlType)) {
		query["SqlType"] = request.SqlType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Succeed)) {
		query["Succeed"] = request.Succeed
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAuditLogRecords"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAuditLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you call the DescribeAuditLogRecords operation to query the SQL audit logs of an AnalyticDB for MySQL cluster, you must enable SQL audit for the cluster. You can call the [DescribeAuditLogConfig](~~190628~~) operation to query the status of SQL audit. If SQL audit is disabled, you can call the [ModifyAuditLogConfig](~~190629~~) operation to enable SQL audit.
 * SQL audit logs can be queried only when SQL audit is enabled. Only SQL audit logs within the last 30 days can be queried. If SQL audit was disabled and re-enabled, only SQL audit logs from the time when SQL audit was re-enabled can be queried. The following operations are not recorded in SQL audit logs: **INSERT INTO VALUES**, **REPLACE INTO VALUES**, and **UPSERT INTO VALUES**.
 *
 * @param request DescribeAuditLogRecordsRequest
 * @return DescribeAuditLogRecordsResponse
 */
func (client *Client) DescribeAuditLogRecords(request *DescribeAuditLogRecordsRequest) (_result *DescribeAuditLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditLogRecordsResponse{}
	_body, _err := client.DescribeAuditLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutoRenewAttributeWithOptions(request *DescribeAutoRenewAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoRenewAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterIds)) {
		query["DBClusterIds"] = request.DBClusterIds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAutoRenewAttribute"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutoRenewAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutoRenewAttribute(request *DescribeAutoRenewAttributeRequest) (_result *DescribeAutoRenewAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoRenewAttributeResponse{}
	_body, _err := client.DescribeAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableAdvicesWithOptions(request *DescribeAvailableAdvicesRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableAdvicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdviceDate)) {
		query["AdviceDate"] = request.AdviceDate
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAvailableAdvices"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAvailableAdvicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableAdvices(request *DescribeAvailableAdvicesRequest) (_result *DescribeAvailableAdvicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableAdvicesResponse{}
	_body, _err := client.DescribeAvailableAdvicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableResourceWithOptions(request *DescribeAvailableResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterVersion)) {
		query["DBClusterVersion"] = request.DBClusterVersion
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAvailableResource"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableResource(request *DescribeAvailableResourceRequest) (_result *DescribeAvailableResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DescribeAvailableResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupPolicyWithOptions(request *DescribeBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupPolicy"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupPolicy(request *DescribeBackupPolicyRequest) (_result *DescribeBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DescribeBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupsWithOptions(request *DescribeBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackups"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackups(request *DescribeBackupsRequest) (_result *DescribeBackupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.DescribeBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeColumnsWithOptions(request *DescribeColumnsRequest, runtime *util.RuntimeOptions) (_result *DescribeColumnsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeColumns"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeColumns(request *DescribeColumnsRequest) (_result *DescribeColumnsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeColumnsResponse{}
	_body, _err := client.DescribeColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeComputeResourceWithOptions(request *DescribeComputeResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeComputeResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterVersion)) {
		query["DBClusterVersion"] = request.DBClusterVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Migrate)) {
		query["Migrate"] = request.Migrate
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComputeResource"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeComputeResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeComputeResource(request *DescribeComputeResourceRequest) (_result *DescribeComputeResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComputeResourceResponse{}
	_body, _err := client.DescribeComputeResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConnectionCountRecordsWithOptions(request *DescribeConnectionCountRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeConnectionCountRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeConnectionCountRecords"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeConnectionCountRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConnectionCountRecords(request *DescribeConnectionCountRecordsRequest) (_result *DescribeConnectionCountRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConnectionCountRecordsResponse{}
	_body, _err := client.DescribeConnectionCountRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterAccessWhiteListWithOptions(request *DescribeDBClusterAccessWhiteListRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterAccessWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterAccessWhiteList"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterAccessWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterAccessWhiteList(request *DescribeDBClusterAccessWhiteListRequest) (_result *DescribeDBClusterAccessWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterAccessWhiteListResponse{}
	_body, _err := client.DescribeDBClusterAccessWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterAttributeWithOptions(request *DescribeDBClusterAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterAttribute"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterAttribute(request *DescribeDBClusterAttributeRequest) (_result *DescribeDBClusterAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterAttributeResponse{}
	_body, _err := client.DescribeDBClusterAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterHealthStatusWithOptions(request *DescribeDBClusterHealthStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterHealthStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterHealthStatus"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterHealthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterHealthStatus(request *DescribeDBClusterHealthStatusRequest) (_result *DescribeDBClusterHealthStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterHealthStatusResponse{}
	_body, _err := client.DescribeDBClusterHealthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterNetInfoWithOptions(request *DescribeDBClusterNetInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterNetInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterNetInfo"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterNetInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterNetInfo(request *DescribeDBClusterNetInfoRequest) (_result *DescribeDBClusterNetInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterNetInfoResponse{}
	_body, _err := client.DescribeDBClusterNetInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the performance data of a cluster over a time range based on its performance metrics. The data is collected every 30 seconds. This operation allows you to query information about slow queries, such as the SQL query duration, number of scanned rows, and amount of scanned data.
 *
 * @param request DescribeDBClusterPerformanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDBClusterPerformanceResponse
 */
func (client *Client) DescribeDBClusterPerformanceWithOptions(request *DescribeDBClusterPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePools)) {
		query["ResourcePools"] = request.ResourcePools
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterPerformance"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the performance data of a cluster over a time range based on its performance metrics. The data is collected every 30 seconds. This operation allows you to query information about slow queries, such as the SQL query duration, number of scanned rows, and amount of scanned data.
 *
 * @param request DescribeDBClusterPerformanceRequest
 * @return DescribeDBClusterPerformanceResponse
 */
func (client *Client) DescribeDBClusterPerformance(request *DescribeDBClusterPerformanceRequest) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.DescribeDBClusterPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * > You can also view the monitoring information about resource groups within an AnalyticDB for MySQL cluster in elastic mode for Cluster Edition in the form of graphs in the console. For more information, see [View monitoring information](~~188721~~).
 *
 * @param request DescribeDBClusterResourcePoolPerformanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDBClusterResourcePoolPerformanceResponse
 */
func (client *Client) DescribeDBClusterResourcePoolPerformanceWithOptions(request *DescribeDBClusterResourcePoolPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterResourcePoolPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePools)) {
		query["ResourcePools"] = request.ResourcePools
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterResourcePoolPerformance"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterResourcePoolPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * > You can also view the monitoring information about resource groups within an AnalyticDB for MySQL cluster in elastic mode for Cluster Edition in the form of graphs in the console. For more information, see [View monitoring information](~~188721~~).
 *
 * @param request DescribeDBClusterResourcePoolPerformanceRequest
 * @return DescribeDBClusterResourcePoolPerformanceResponse
 */
func (client *Client) DescribeDBClusterResourcePoolPerformance(request *DescribeDBClusterResourcePoolPerformanceRequest) (_result *DescribeDBClusterResourcePoolPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterResourcePoolPerformanceResponse{}
	_body, _err := client.DescribeDBClusterResourcePoolPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterStatusWithOptions(request *DescribeDBClusterStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterStatus"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterStatus(request *DescribeDBClusterStatusRequest) (_result *DescribeDBClusterStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterStatusResponse{}
	_body, _err := client.DescribeDBClusterStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClustersWithOptions(request *DescribeDBClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterDescription)) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterIds)) {
		query["DBClusterIds"] = request.DBClusterIds
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterStatus)) {
		query["DBClusterStatus"] = request.DBClusterStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DBVersion)) {
		query["DBVersion"] = request.DBVersion
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusters"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusters(request *DescribeDBClustersRequest) (_result *DescribeDBClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClustersResponse{}
	_body, _err := client.DescribeDBClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * You can call this operation only for AnalyticDB for MySQL clusters in elastic mode for Cluster Edition that have 32 cores or more.
 *
 * @param request DescribeDBResourceGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDBResourceGroupResponse
 */
func (client *Client) DescribeDBResourceGroupWithOptions(request *DescribeDBResourceGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeDBResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBResourceGroup"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * You can call this operation only for AnalyticDB for MySQL clusters in elastic mode for Cluster Edition that have 32 cores or more.
 *
 * @param request DescribeDBResourceGroupRequest
 * @return DescribeDBResourceGroupResponse
 */
func (client *Client) DescribeDBResourceGroup(request *DescribeDBResourceGroupRequest) (_result *DescribeDBResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBResourceGroupResponse{}
	_body, _err := client.DescribeDBResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This operation is applicable only for elastic clusters of 32 cores or more.
 *
 * @param request DescribeDBResourcePoolRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDBResourcePoolResponse
 */
func (client *Client) DescribeDBResourcePoolWithOptions(request *DescribeDBResourcePoolRequest, runtime *util.RuntimeOptions) (_result *DescribeDBResourcePoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBResourcePool"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBResourcePoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation is applicable only for elastic clusters of 32 cores or more.
 *
 * @param request DescribeDBResourcePoolRequest
 * @return DescribeDBResourcePoolResponse
 */
func (client *Client) DescribeDBResourcePool(request *DescribeDBResourcePoolRequest) (_result *DescribeDBResourcePoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBResourcePoolResponse{}
	_body, _err := client.DescribeDBResourcePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosisDimensionsWithOptions(request *DescribeDiagnosisDimensionsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisDimensionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisDimensions"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisDimensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosisDimensions(request *DescribeDiagnosisDimensionsRequest) (_result *DescribeDiagnosisDimensionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisDimensionsResponse{}
	_body, _err := client.DescribeDiagnosisDimensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosisMonitorPerformanceWithOptions(request *DescribeDiagnosisMonitorPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisMonitorPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisMonitorPerformance"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisMonitorPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosisMonitorPerformance(request *DescribeDiagnosisMonitorPerformanceRequest) (_result *DescribeDiagnosisMonitorPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisMonitorPerformanceResponse{}
	_body, _err := client.DescribeDiagnosisMonitorPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosisRecordsWithOptions(request *DescribeDiagnosisRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MaxPeakMemory)) {
		query["MaxPeakMemory"] = request.MaxPeakMemory
	}

	if !tea.BoolValue(util.IsUnset(request.MaxScanSize)) {
		query["MaxScanSize"] = request.MaxScanSize
	}

	if !tea.BoolValue(util.IsUnset(request.MinPeakMemory)) {
		query["MinPeakMemory"] = request.MinPeakMemory
	}

	if !tea.BoolValue(util.IsUnset(request.MinScanSize)) {
		query["MinScanSize"] = request.MinScanSize
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PatternId)) {
		query["PatternId"] = request.PatternId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroup)) {
		query["ResourceGroup"] = request.ResourceGroup
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisRecords"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosisRecords(request *DescribeDiagnosisRecordsRequest) (_result *DescribeDiagnosisRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisRecordsResponse{}
	_body, _err := client.DescribeDiagnosisRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosisSQLInfoWithOptions(request *DescribeDiagnosisSQLInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisSQLInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisSQLInfo"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisSQLInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosisSQLInfo(request *DescribeDiagnosisSQLInfoRequest) (_result *DescribeDiagnosisSQLInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisSQLInfoResponse{}
	_body, _err := client.DescribeDiagnosisSQLInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiagnosisTasksWithOptions(request *DescribeDiagnosisTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosisTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosisTasks"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosisTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiagnosisTasks(request *DescribeDiagnosisTasksRequest) (_result *DescribeDiagnosisTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosisTasksResponse{}
	_body, _err := client.DescribeDiagnosisTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDownloadRecordsWithOptions(request *DescribeDownloadRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeDownloadRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDownloadRecords"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDownloadRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDownloadRecords(request *DescribeDownloadRecordsRequest) (_result *DescribeDownloadRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDownloadRecordsResponse{}
	_body, _err := client.DescribeDownloadRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEIURangeWithOptions(request *DescribeEIURangeRequest, runtime *util.RuntimeOptions) (_result *DescribeEIURangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComputeResource)) {
		query["ComputeResource"] = request.ComputeResource
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterVersion)) {
		query["DBClusterVersion"] = request.DBClusterVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		query["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEIURange"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEIURangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEIURange(request *DescribeEIURangeRequest) (_result *DescribeEIURangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEIURangeResponse{}
	_body, _err := client.DescribeEIURangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This operation is available only for AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters in elastic mode for Cluster Edition that have 32 cores or more.
 *
 * @param request DescribeElasticDailyPlanRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeElasticDailyPlanResponse
 */
func (client *Client) DescribeElasticDailyPlanWithOptions(request *DescribeElasticDailyPlanRequest, runtime *util.RuntimeOptions) (_result *DescribeElasticDailyPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticDailyPlanDay)) {
		query["ElasticDailyPlanDay"] = request.ElasticDailyPlanDay
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticDailyPlanStatusList)) {
		query["ElasticDailyPlanStatusList"] = request.ElasticDailyPlanStatusList
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePoolName)) {
		query["ResourcePoolName"] = request.ResourcePoolName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeElasticDailyPlan"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeElasticDailyPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation is available only for AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters in elastic mode for Cluster Edition that have 32 cores or more.
 *
 * @param request DescribeElasticDailyPlanRequest
 * @return DescribeElasticDailyPlanResponse
 */
func (client *Client) DescribeElasticDailyPlan(request *DescribeElasticDailyPlanRequest) (_result *DescribeElasticDailyPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeElasticDailyPlanResponse{}
	_body, _err := client.DescribeElasticDailyPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * You can call this operation only for AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters in elastic mode for Cluster Edition that have 32 cores or more.
 *
 * @param request DescribeElasticPlanRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeElasticPlanResponse
 */
func (client *Client) DescribeElasticPlanWithOptions(request *DescribeElasticPlanRequest, runtime *util.RuntimeOptions) (_result *DescribeElasticPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanEnable)) {
		query["ElasticPlanEnable"] = request.ElasticPlanEnable
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePoolName)) {
		query["ResourcePoolName"] = request.ResourcePoolName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeElasticPlan"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * You can call this operation only for AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters in elastic mode for Cluster Edition that have 32 cores or more.
 *
 * @param request DescribeElasticPlanRequest
 * @return DescribeElasticPlanResponse
 */
func (client *Client) DescribeElasticPlan(request *DescribeElasticPlanRequest) (_result *DescribeElasticPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeElasticPlanResponse{}
	_body, _err := client.DescribeElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInclinedTablesWithOptions(request *DescribeInclinedTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeInclinedTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TableType)) {
		query["TableType"] = request.TableType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInclinedTables"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInclinedTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInclinedTables(request *DescribeInclinedTablesRequest) (_result *DescribeInclinedTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInclinedTablesResponse{}
	_body, _err := client.DescribeInclinedTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * For information about how to asynchronously submit import and export tasks, see [Asynchronously submit an import or export task](~~160291~~).
 *
 * @param request DescribeLoadTasksRecordsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeLoadTasksRecordsResponse
 */
func (client *Client) DescribeLoadTasksRecordsWithOptions(request *DescribeLoadTasksRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeLoadTasksRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		query["DBName"] = request.DBName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLoadTasksRecords"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLoadTasksRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * For information about how to asynchronously submit import and export tasks, see [Asynchronously submit an import or export task](~~160291~~).
 *
 * @param request DescribeLoadTasksRecordsRequest
 * @return DescribeLoadTasksRecordsResponse
 */
func (client *Client) DescribeLoadTasksRecords(request *DescribeLoadTasksRecordsRequest) (_result *DescribeLoadTasksRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLoadTasksRecordsResponse{}
	_body, _err := client.DescribeLoadTasksRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMaintenanceActionWithOptions(request *DescribeMaintenanceActionRequest, runtime *util.RuntimeOptions) (_result *DescribeMaintenanceActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsHistory)) {
		query["IsHistory"] = request.IsHistory
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMaintenanceAction"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMaintenanceActionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMaintenanceAction(request *DescribeMaintenanceActionRequest) (_result *DescribeMaintenanceActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMaintenanceActionResponse{}
	_body, _err := client.DescribeMaintenanceActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOperatorPermissionWithOptions(request *DescribeOperatorPermissionRequest, runtime *util.RuntimeOptions) (_result *DescribeOperatorPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOperatorPermission"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOperatorPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOperatorPermission(request *DescribeOperatorPermissionRequest) (_result *DescribeOperatorPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOperatorPermissionResponse{}
	_body, _err := client.DescribeOperatorPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePatternPerformanceWithOptions(request *DescribePatternPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribePatternPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PatternId)) {
		query["PatternId"] = request.PatternId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePatternPerformance"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePatternPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePatternPerformance(request *DescribePatternPerformanceRequest) (_result *DescribePatternPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePatternPerformanceResponse{}
	_body, _err := client.DescribePatternPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProcessListWithOptions(request *DescribeProcessListRequest, runtime *util.RuntimeOptions) (_result *DescribeProcessListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RunningTime)) {
		query["RunningTime"] = request.RunningTime
	}

	if !tea.BoolValue(util.IsUnset(request.ShowFull)) {
		query["ShowFull"] = request.ShowFull
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProcessList"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProcessListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProcessList(request *DescribeProcessListRequest) (_result *DescribeProcessListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProcessListResponse{}
	_body, _err := client.DescribeProcessListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResubmitConfigWithOptions(request *DescribeResubmitConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeResubmitConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResubmitConfig"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResubmitConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResubmitConfig(request *DescribeResubmitConfigRequest) (_result *DescribeResubmitConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResubmitConfigResponse{}
	_body, _err := client.DescribeResubmitConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQAConfigWithOptions(request *DescribeSQAConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeSQAConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQAConfig"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQAConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQAConfig(request *DescribeSQAConfigRequest) (_result *DescribeSQAConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQAConfigResponse{}
	_body, _err := client.DescribeSQAConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLPatternsWithOptions(request *DescribeSQLPatternsRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLPatternsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLPatterns"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLPatternsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLPatterns(request *DescribeSQLPatternsRequest) (_result *DescribeSQLPatternsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLPatternsResponse{}
	_body, _err := client.DescribeSQLPatternsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLPlanWithOptions(request *DescribeSQLPlanRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessId)) {
		query["ProcessId"] = request.ProcessId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLPlan"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLPlan(request *DescribeSQLPlanRequest) (_result *DescribeSQLPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLPlanResponse{}
	_body, _err := client.DescribeSQLPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLPlanTaskWithOptions(request *DescribeSQLPlanTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLPlanTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessId)) {
		query["ProcessId"] = request.ProcessId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StageId)) {
		query["StageId"] = request.StageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLPlanTask"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLPlanTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLPlanTask(request *DescribeSQLPlanTaskRequest) (_result *DescribeSQLPlanTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLPlanTaskResponse{}
	_body, _err := client.DescribeSQLPlanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSchemasWithOptions(request *DescribeSchemasRequest, runtime *util.RuntimeOptions) (_result *DescribeSchemasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSchemas"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSchemas(request *DescribeSchemasRequest) (_result *DescribeSchemasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSchemasResponse{}
	_body, _err := client.DescribeSchemasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowLogRecordsWithOptions(request *DescribeSlowLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		query["DBName"] = request.DBName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessID)) {
		query["ProcessID"] = request.ProcessID
	}

	if !tea.BoolValue(util.IsUnset(request.Range)) {
		query["Range"] = request.Range
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlowLogRecords"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowLogRecords(request *DescribeSlowLogRecordsRequest) (_result *DescribeSlowLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DescribeSlowLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowLogTrendWithOptions(request *DescribeSlowLogTrendRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowLogTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBName)) {
		query["DBName"] = request.DBName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlowLogTrend"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlowLogTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowLogTrend(request *DescribeSlowLogTrendRequest) (_result *DescribeSlowLogTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowLogTrendResponse{}
	_body, _err := client.DescribeSlowLogTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSqlPatternWithOptions(request *DescribeSqlPatternRequest, runtime *util.RuntimeOptions) (_result *DescribeSqlPatternResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SqlPattern)) {
		query["SqlPattern"] = request.SqlPattern
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSqlPattern"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSqlPatternResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSqlPattern(request *DescribeSqlPatternRequest) (_result *DescribeSqlPatternResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSqlPatternResponse{}
	_body, _err := client.DescribeSqlPatternWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTableAccessCountWithOptions(request *DescribeTableAccessCountRequest, runtime *util.RuntimeOptions) (_result *DescribeTableAccessCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTableAccessCount"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTableAccessCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTableAccessCount(request *DescribeTableAccessCountRequest) (_result *DescribeTableAccessCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTableAccessCountResponse{}
	_body, _err := client.DescribeTableAccessCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTableDetailWithOptions(request *DescribeTableDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeTableDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTableDetail"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTableDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTableDetail(request *DescribeTableDetailRequest) (_result *DescribeTableDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTableDetailResponse{}
	_body, _err := client.DescribeTableDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTablePartitionDiagnoseWithOptions(request *DescribeTablePartitionDiagnoseRequest, runtime *util.RuntimeOptions) (_result *DescribeTablePartitionDiagnoseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTablePartitionDiagnose"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTablePartitionDiagnoseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTablePartitionDiagnose(request *DescribeTablePartitionDiagnoseRequest) (_result *DescribeTablePartitionDiagnoseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTablePartitionDiagnoseResponse{}
	_body, _err := client.DescribeTablePartitionDiagnoseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * >  For more information about table statistics, see [View monitoring information of resource pools](~~188721~~).
 *
 * @param request DescribeTableStatisticsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeTableStatisticsResponse
 */
func (client *Client) DescribeTableStatisticsWithOptions(request *DescribeTableStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeTableStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTableStatistics"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTableStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * >  For more information about table statistics, see [View monitoring information of resource pools](~~188721~~).
 *
 * @param request DescribeTableStatisticsRequest
 * @return DescribeTableStatisticsResponse
 */
func (client *Client) DescribeTableStatistics(request *DescribeTableStatisticsRequest) (_result *DescribeTableStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTableStatisticsResponse{}
	_body, _err := client.DescribeTableStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTablesWithOptions(request *DescribeTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["SchemaName"] = request.SchemaName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTables"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTables(request *DescribeTablesRequest) (_result *DescribeTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTablesResponse{}
	_body, _err := client.DescribeTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTaskInfoWithOptions(request *DescribeTaskInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeTaskInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTaskInfo"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTaskInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTaskInfo(request *DescribeTaskInfoRequest) (_result *DescribeTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTaskInfoResponse{}
	_body, _err := client.DescribeTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVSwitchesWithOptions(request *DescribeVSwitchesRequest, runtime *util.RuntimeOptions) (_result *DescribeVSwitchesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VswId)) {
		query["VswId"] = request.VswId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVSwitches"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVSwitches(request *DescribeVSwitchesRequest) (_result *DescribeVSwitchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.DescribeVSwitchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation only for AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters in elastic mode for Cluster Edition.
 *
 * @param request DetachUserENIRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DetachUserENIResponse
 */
func (client *Client) DetachUserENIWithOptions(request *DetachUserENIRequest, runtime *util.RuntimeOptions) (_result *DetachUserENIResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachUserENI"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachUserENIResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation only for AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters in elastic mode for Cluster Edition.
 *
 * @param request DetachUserENIRequest
 * @return DetachUserENIResponse
 */
func (client *Client) DetachUserENI(request *DetachUserENIRequest) (_result *DetachUserENIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachUserENIResponse{}
	_body, _err := client.DetachUserENIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableAdviceServiceWithOptions(request *DisableAdviceServiceRequest, runtime *util.RuntimeOptions) (_result *DisableAdviceServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableAdviceService"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableAdviceServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableAdviceService(request *DisableAdviceServiceRequest) (_result *DisableAdviceServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableAdviceServiceResponse{}
	_body, _err := client.DisableAdviceServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadDiagnosisRecordsWithOptions(request *DownloadDiagnosisRecordsRequest, runtime *util.RuntimeOptions) (_result *DownloadDiagnosisRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Database)) {
		query["Database"] = request.Database
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MaxPeakMemory)) {
		query["MaxPeakMemory"] = request.MaxPeakMemory
	}

	if !tea.BoolValue(util.IsUnset(request.MaxScanSize)) {
		query["MaxScanSize"] = request.MaxScanSize
	}

	if !tea.BoolValue(util.IsUnset(request.MinPeakMemory)) {
		query["MinPeakMemory"] = request.MinPeakMemory
	}

	if !tea.BoolValue(util.IsUnset(request.MinScanSize)) {
		query["MinScanSize"] = request.MinScanSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCondition)) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroup)) {
		query["ResourceGroup"] = request.ResourceGroup
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DownloadDiagnosisRecords"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DownloadDiagnosisRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadDiagnosisRecords(request *DownloadDiagnosisRecordsRequest) (_result *DownloadDiagnosisRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadDiagnosisRecordsResponse{}
	_body, _err := client.DownloadDiagnosisRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableAdviceServiceWithOptions(request *EnableAdviceServiceRequest, runtime *util.RuntimeOptions) (_result *EnableAdviceServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableAdviceService"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableAdviceServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableAdviceService(request *EnableAdviceServiceRequest) (_result *EnableAdviceServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableAdviceServiceResponse{}
	_body, _err := client.EnableAdviceServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * If you need Alibaba Cloud technical support to perform operations on your AnalyticDB for MySQL cluster, you must grant permissions to the service account of your cluster. When the validity period of the authorization ends, the granted permissions are automatically revoked.
 *
 * @param request GrantOperatorPermissionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GrantOperatorPermissionResponse
 */
func (client *Client) GrantOperatorPermissionWithOptions(request *GrantOperatorPermissionRequest, runtime *util.RuntimeOptions) (_result *GrantOperatorPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTime)) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Privileges)) {
		query["Privileges"] = request.Privileges
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantOperatorPermission"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantOperatorPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * If you need Alibaba Cloud technical support to perform operations on your AnalyticDB for MySQL cluster, you must grant permissions to the service account of your cluster. When the validity period of the authorization ends, the granted permissions are automatically revoked.
 *
 * @param request GrantOperatorPermissionRequest
 * @return GrantOperatorPermissionResponse
 */
func (client *Client) GrantOperatorPermission(request *GrantOperatorPermissionRequest) (_result *GrantOperatorPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantOperatorPermissionResponse{}
	_body, _err := client.GrantOperatorPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KillProcessWithOptions(request *KillProcessRequest, runtime *util.RuntimeOptions) (_result *KillProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessId)) {
		query["ProcessId"] = request.ProcessId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("KillProcess"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &KillProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KillProcess(request *KillProcessRequest) (_result *KillProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KillProcessResponse{}
	_body, _err := client.KillProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MigrateDBClusterWithOptions(request *MigrateDBClusterRequest, runtime *util.RuntimeOptions) (_result *MigrateDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MigrateDBCluster"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MigrateDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MigrateDBCluster(request *MigrateDBClusterRequest) (_result *MigrateDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MigrateDBClusterResponse{}
	_body, _err := client.MigrateDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccountDescriptionWithOptions(request *ModifyAccountDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountDescription)) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountDescription"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (_result *ModifyAccountDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.ModifyAccountDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAuditLogConfigWithOptions(request *ModifyAuditLogConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyAuditLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditLogStatus)) {
		query["AuditLogStatus"] = request.AuditLogStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAuditLogConfig"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAuditLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAuditLogConfig(request *ModifyAuditLogConfigRequest) (_result *ModifyAuditLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAuditLogConfigResponse{}
	_body, _err := client.ModifyAuditLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAutoRenewAttributeWithOptions(request *ModifyAutoRenewAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyAutoRenewAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RenewalStatus)) {
		query["RenewalStatus"] = request.RenewalStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAutoRenewAttribute"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAutoRenewAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAutoRenewAttribute(request *ModifyAutoRenewAttributeRequest) (_result *ModifyAutoRenewAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAutoRenewAttributeResponse{}
	_body, _err := client.ModifyAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBackupPolicyWithOptions(request *ModifyBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupRetentionPeriod)) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableBackupLog)) {
		query["EnableBackupLog"] = request.EnableBackupLog
	}

	if !tea.BoolValue(util.IsUnset(request.LogBackupRetentionPeriod)) {
		query["LogBackupRetentionPeriod"] = request.LogBackupRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupPeriod)) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupTime)) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyBackupPolicy"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBackupPolicy(request *ModifyBackupPolicyRequest) (_result *ModifyBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.ModifyBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterConnectionStringWithOptions(request *ModifyClusterConnectionStringRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterConnectionStringResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionStringPrefix)) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentConnectionString)) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterConnectionString"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterConnectionStringResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterConnectionString(request *ModifyClusterConnectionStringRequest) (_result *ModifyClusterConnectionStringResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterConnectionStringResponse{}
	_body, _err := client.ModifyClusterConnectionStringWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterWithOptions(request *ModifyDBClusterRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComputeResource)) {
		query["ComputeResource"] = request.ComputeResource
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterCategory)) {
		query["DBClusterCategory"] = request.DBClusterCategory
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeClass)) {
		query["DBNodeClass"] = request.DBNodeClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeGroupCount)) {
		query["DBNodeGroupCount"] = request.DBNodeGroupCount
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeStorage)) {
		query["DBNodeStorage"] = request.DBNodeStorage
	}

	if !tea.BoolValue(util.IsUnset(request.DiskPerformanceLevel)) {
		query["DiskPerformanceLevel"] = request.DiskPerformanceLevel
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticIOResource)) {
		query["ElasticIOResource"] = request.ElasticIOResource
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticIOResourceSize)) {
		query["ElasticIOResourceSize"] = request.ElasticIOResourceSize
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorCount)) {
		query["ExecutorCount"] = request.ExecutorCount
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyType)) {
		query["ModifyType"] = request.ModifyType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StorageResource)) {
		query["StorageResource"] = request.StorageResource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBCluster"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBCluster(request *ModifyDBClusterRequest) (_result *ModifyDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterResponse{}
	_body, _err := client.ModifyDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterAccessWhiteListWithOptions(request *ModifyDBClusterAccessWhiteListRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterAccessWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterIPArrayAttribute)) {
		query["DBClusterIPArrayAttribute"] = request.DBClusterIPArrayAttribute
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterIPArrayName)) {
		query["DBClusterIPArrayName"] = request.DBClusterIPArrayName
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyMode)) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIps)) {
		query["SecurityIps"] = request.SecurityIps
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterAccessWhiteList"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterAccessWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterAccessWhiteList(request *ModifyDBClusterAccessWhiteListRequest) (_result *ModifyDBClusterAccessWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterAccessWhiteListResponse{}
	_body, _err := client.ModifyDBClusterAccessWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterDescriptionWithOptions(request *ModifyDBClusterDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterDescription)) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterDescription"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterDescription(request *ModifyDBClusterDescriptionRequest) (_result *ModifyDBClusterDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterDescriptionResponse{}
	_body, _err := client.ModifyDBClusterDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterMaintainTimeWithOptions(request *ModifyDBClusterMaintainTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MaintainTime)) {
		query["MaintainTime"] = request.MaintainTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterMaintainTime"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterMaintainTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterMaintainTime(request *ModifyDBClusterMaintainTimeRequest) (_result *ModifyDBClusterMaintainTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterMaintainTimeResponse{}
	_body, _err := client.ModifyDBClusterMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterPayTypeWithOptions(request *ModifyDBClusterPayTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterPayTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbClusterId)) {
		query["DbClusterId"] = request.DbClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.UsedTime)) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterPayType"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterPayTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterPayType(request *ModifyDBClusterPayTypeRequest) (_result *ModifyDBClusterPayTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterPayTypeResponse{}
	_body, _err := client.ModifyDBClusterPayTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Resource Management enables you to build an organizational structure for resources based on your business needs. You can use a resource directory, folders, accounts, and resource groups to hierarchically organize and manage resources. For more information, see [What is Resource Management?](~~94475#concept-zyn-3p1-dhb~~ "Resource Management provides a collection of resource management services that support enterprise IT administration. The services include Resource Directory, Resource Group, and Tag. Resource Directory allows you to build an organizational structure for resources based on your business requirements. Resource Group and Tag allow you to hierarchically manage the resources. Resource Sharing allows you to share the resources among your accounts.")
 *
 * @param request ModifyDBClusterResourceGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyDBClusterResourceGroupResponse
 */
func (client *Client) ModifyDBClusterResourceGroupWithOptions(request *ModifyDBClusterResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterResourceGroup"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Resource Management enables you to build an organizational structure for resources based on your business needs. You can use a resource directory, folders, accounts, and resource groups to hierarchically organize and manage resources. For more information, see [What is Resource Management?](~~94475#concept-zyn-3p1-dhb~~ "Resource Management provides a collection of resource management services that support enterprise IT administration. The services include Resource Directory, Resource Group, and Tag. Resource Directory allows you to build an organizational structure for resources based on your business requirements. Resource Group and Tag allow you to hierarchically manage the resources. Resource Sharing allows you to share the resources among your accounts.")
 *
 * @param request ModifyDBClusterResourceGroupRequest
 * @return ModifyDBClusterResourceGroupResponse
 */
func (client *Client) ModifyDBClusterResourceGroup(request *ModifyDBClusterResourceGroupRequest) (_result *ModifyDBClusterResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterResourceGroupResponse{}
	_body, _err := client.ModifyDBClusterResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Precautions
 * *   This operation is applicable only for elastic clusters of 32 cores or more.
 * *   The number of nodes cannot be changed for the default resource group USER_DEFAULT.
 *
 * @param request ModifyDBResourceGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyDBResourceGroupResponse
 */
func (client *Client) ModifyDBResourceGroupWithOptions(request *ModifyDBResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDBResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["GroupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeNum)) {
		query["NodeNum"] = request.NodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBResourceGroup"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Precautions
 * *   This operation is applicable only for elastic clusters of 32 cores or more.
 * *   The number of nodes cannot be changed for the default resource group USER_DEFAULT.
 *
 * @param request ModifyDBResourceGroupRequest
 * @return ModifyDBResourceGroupResponse
 */
func (client *Client) ModifyDBResourceGroup(request *ModifyDBResourceGroupRequest) (_result *ModifyDBResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBResourceGroupResponse{}
	_body, _err := client.ModifyDBResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * *   You can call this operation only for elastic clusters of 32 cores or more.
 * *   You cannot change the number of nodes for the USER_DEFAULT resource group.
 *
 * @param request ModifyDBResourcePoolRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyDBResourcePoolResponse
 */
func (client *Client) ModifyDBResourcePoolWithOptions(request *ModifyDBResourcePoolRequest, runtime *util.RuntimeOptions) (_result *ModifyDBResourcePoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeNum)) {
		query["NodeNum"] = request.NodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBResourcePool"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBResourcePoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * *   You can call this operation only for elastic clusters of 32 cores or more.
 * *   You cannot change the number of nodes for the USER_DEFAULT resource group.
 *
 * @param request ModifyDBResourcePoolRequest
 * @return ModifyDBResourcePoolResponse
 */
func (client *Client) ModifyDBResourcePool(request *ModifyDBResourcePoolRequest) (_result *ModifyDBResourcePoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBResourcePoolResponse{}
	_body, _err := client.ModifyDBResourcePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation only for AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters in elastic mode for Cluster Edition that have 32 cores or more.
 *
 * @param request ModifyElasticPlanRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyElasticPlanResponse
 */
func (client *Client) ModifyElasticPlanWithOptions(request *ModifyElasticPlanRequest, runtime *util.RuntimeOptions) (_result *ModifyElasticPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanEnable)) {
		query["ElasticPlanEnable"] = request.ElasticPlanEnable
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanEndDay)) {
		query["ElasticPlanEndDay"] = request.ElasticPlanEndDay
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanMonthlyRepeat)) {
		query["ElasticPlanMonthlyRepeat"] = request.ElasticPlanMonthlyRepeat
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanName)) {
		query["ElasticPlanName"] = request.ElasticPlanName
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanNodeNum)) {
		query["ElasticPlanNodeNum"] = request.ElasticPlanNodeNum
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanStartDay)) {
		query["ElasticPlanStartDay"] = request.ElasticPlanStartDay
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanTimeEnd)) {
		query["ElasticPlanTimeEnd"] = request.ElasticPlanTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanTimeStart)) {
		query["ElasticPlanTimeStart"] = request.ElasticPlanTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanType)) {
		query["ElasticPlanType"] = request.ElasticPlanType
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanWeeklyRepeat)) {
		query["ElasticPlanWeeklyRepeat"] = request.ElasticPlanWeeklyRepeat
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticPlanWorkerSpec)) {
		query["ElasticPlanWorkerSpec"] = request.ElasticPlanWorkerSpec
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePoolName)) {
		query["ResourcePoolName"] = request.ResourcePoolName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyElasticPlan"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyElasticPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation only for AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters in elastic mode for Cluster Edition that have 32 cores or more.
 *
 * @param request ModifyElasticPlanRequest
 * @return ModifyElasticPlanResponse
 */
func (client *Client) ModifyElasticPlan(request *ModifyElasticPlanRequest) (_result *ModifyElasticPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyElasticPlanResponse{}
	_body, _err := client.ModifyElasticPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLogBackupPolicyWithOptions(request *ModifyLogBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyLogBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableBackupLog)) {
		query["EnableBackupLog"] = request.EnableBackupLog
	}

	if !tea.BoolValue(util.IsUnset(request.LogBackupRetentionPeriod)) {
		query["LogBackupRetentionPeriod"] = request.LogBackupRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLogBackupPolicy"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyLogBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLogBackupPolicy(request *ModifyLogBackupPolicyRequest) (_result *ModifyLogBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLogBackupPolicyResponse{}
	_body, _err := client.ModifyLogBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMaintenanceActionWithOptions(request *ModifyMaintenanceActionRequest, runtime *util.RuntimeOptions) (_result *ModifyMaintenanceActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTime)) {
		query["SwitchTime"] = request.SwitchTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMaintenanceAction"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMaintenanceActionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMaintenanceAction(request *ModifyMaintenanceActionRequest) (_result *ModifyMaintenanceActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMaintenanceActionResponse{}
	_body, _err := client.ModifyMaintenanceActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyResubmitConfigWithOptions(tmpReq *ModifyResubmitConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyResubmitConfigResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyResubmitConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Rules)) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, tea.String("Rules"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RulesShrink)) {
		query["Rules"] = request.RulesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyResubmitConfig"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyResubmitConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyResubmitConfig(request *ModifyResubmitConfigRequest) (_result *ModifyResubmitConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyResubmitConfigResponse{}
	_body, _err := client.ModifyResubmitConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySQAConfigWithOptions(request *ModifySQAConfigRequest, runtime *util.RuntimeOptions) (_result *ModifySQAConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SQAStatus)) {
		query["SQAStatus"] = request.SQAStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySQAConfig"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySQAConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySQAConfig(request *ModifySQAConfigRequest) (_result *ModifySQAConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySQAConfigResponse{}
	_body, _err := client.ModifySQAConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseClusterPublicConnectionWithOptions(request *ReleaseClusterPublicConnectionRequest, runtime *util.RuntimeOptions) (_result *ReleaseClusterPublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseClusterPublicConnection"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseClusterPublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseClusterPublicConnection(request *ReleaseClusterPublicConnectionRequest) (_result *ReleaseClusterPublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseClusterPublicConnectionResponse{}
	_body, _err := client.ReleaseClusterPublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetAccountPasswordWithOptions(request *ResetAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAccountPassword"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (_result *ResetAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.ResetAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeOperatorPermissionWithOptions(request *RevokeOperatorPermissionRequest, runtime *util.RuntimeOptions) (_result *RevokeOperatorPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeOperatorPermission"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeOperatorPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeOperatorPermission(request *RevokeOperatorPermissionRequest) (_result *RevokeOperatorPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeOperatorPermissionResponse{}
	_body, _err := client.RevokeOperatorPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindDBResourceGroupWithUserWithOptions(request *UnbindDBResourceGroupWithUserRequest, runtime *util.RuntimeOptions) (_result *UnbindDBResourceGroupWithUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupUser)) {
		query["GroupUser"] = request.GroupUser
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindDBResourceGroupWithUser"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindDBResourceGroupWithUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindDBResourceGroupWithUser(request *UnbindDBResourceGroupWithUserRequest) (_result *UnbindDBResourceGroupWithUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindDBResourceGroupWithUserResponse{}
	_body, _err := client.UnbindDBResourceGroupWithUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindDBResourcePoolWithUserWithOptions(request *UnbindDBResourcePoolWithUserRequest, runtime *util.RuntimeOptions) (_result *UnbindDBResourcePoolWithUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.PoolUser)) {
		query["PoolUser"] = request.PoolUser
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindDBResourcePoolWithUser"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindDBResourcePoolWithUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindDBResourcePoolWithUser(request *UnbindDBResourcePoolWithUserRequest) (_result *UnbindDBResourcePoolWithUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindDBResourcePoolWithUserResponse{}
	_body, _err := client.UnbindDBResourcePoolWithUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

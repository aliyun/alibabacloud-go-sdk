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

type AllocateClusterPublicConnectionRequest struct {
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	DBClusterId            *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AllocateClusterPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AllocateClusterPublicConnectionResponse) SetBody(v *AllocateClusterPublicConnectionResponseBody) *AllocateClusterPublicConnectionResponse {
	s.Body = v
	return s
}

type CheckClickhouseToRDSRequest struct {
	CkPassword           *string `json:"CkPassword,omitempty" xml:"CkPassword,omitempty"`
	CkUserName           *string `json:"CkUserName,omitempty" xml:"CkUserName,omitempty"`
	ClickhousePort       *int64  `json:"ClickhousePort,omitempty" xml:"ClickhousePort,omitempty"`
	DbClusterId          *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RdsId                *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	RdsPassword          *string `json:"RdsPassword,omitempty" xml:"RdsPassword,omitempty"`
	RdsPort              *int64  `json:"RdsPort,omitempty" xml:"RdsPort,omitempty"`
	RdsUserName          *string `json:"RdsUserName,omitempty" xml:"RdsUserName,omitempty"`
	RdsVpcId             *string `json:"RdsVpcId,omitempty" xml:"RdsVpcId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckClickhouseToRDSRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckClickhouseToRDSRequest) GoString() string {
	return s.String()
}

func (s *CheckClickhouseToRDSRequest) SetCkPassword(v string) *CheckClickhouseToRDSRequest {
	s.CkPassword = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetCkUserName(v string) *CheckClickhouseToRDSRequest {
	s.CkUserName = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetClickhousePort(v int64) *CheckClickhouseToRDSRequest {
	s.ClickhousePort = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetDbClusterId(v string) *CheckClickhouseToRDSRequest {
	s.DbClusterId = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetOwnerAccount(v string) *CheckClickhouseToRDSRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetOwnerId(v int64) *CheckClickhouseToRDSRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetRdsId(v string) *CheckClickhouseToRDSRequest {
	s.RdsId = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetRdsPassword(v string) *CheckClickhouseToRDSRequest {
	s.RdsPassword = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetRdsPort(v int64) *CheckClickhouseToRDSRequest {
	s.RdsPort = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetRdsUserName(v string) *CheckClickhouseToRDSRequest {
	s.RdsUserName = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetRdsVpcId(v string) *CheckClickhouseToRDSRequest {
	s.RdsVpcId = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetResourceOwnerAccount(v string) *CheckClickhouseToRDSRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckClickhouseToRDSRequest) SetResourceOwnerId(v int64) *CheckClickhouseToRDSRequest {
	s.ResourceOwnerId = &v
	return s
}

type CheckClickhouseToRDSResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckClickhouseToRDSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckClickhouseToRDSResponseBody) GoString() string {
	return s.String()
}

func (s *CheckClickhouseToRDSResponseBody) SetErrorCode(v string) *CheckClickhouseToRDSResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckClickhouseToRDSResponseBody) SetRequestId(v string) *CheckClickhouseToRDSResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckClickhouseToRDSResponseBody) SetStatus(v bool) *CheckClickhouseToRDSResponseBody {
	s.Status = &v
	return s
}

type CheckClickhouseToRDSResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckClickhouseToRDSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckClickhouseToRDSResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckClickhouseToRDSResponse) GoString() string {
	return s.String()
}

func (s *CheckClickhouseToRDSResponse) SetHeaders(v map[string]*string) *CheckClickhouseToRDSResponse {
	s.Headers = v
	return s
}

func (s *CheckClickhouseToRDSResponse) SetBody(v *CheckClickhouseToRDSResponseBody) *CheckClickhouseToRDSResponse {
	s.Body = v
	return s
}

type CheckScaleOutBalancedRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckScaleOutBalancedRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckScaleOutBalancedRequest) GoString() string {
	return s.String()
}

func (s *CheckScaleOutBalancedRequest) SetDBClusterId(v string) *CheckScaleOutBalancedRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckScaleOutBalancedRequest) SetOwnerAccount(v string) *CheckScaleOutBalancedRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckScaleOutBalancedRequest) SetOwnerId(v int64) *CheckScaleOutBalancedRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckScaleOutBalancedRequest) SetPageNumber(v int32) *CheckScaleOutBalancedRequest {
	s.PageNumber = &v
	return s
}

func (s *CheckScaleOutBalancedRequest) SetPageSize(v int32) *CheckScaleOutBalancedRequest {
	s.PageSize = &v
	return s
}

func (s *CheckScaleOutBalancedRequest) SetRegionId(v string) *CheckScaleOutBalancedRequest {
	s.RegionId = &v
	return s
}

func (s *CheckScaleOutBalancedRequest) SetResourceOwnerAccount(v string) *CheckScaleOutBalancedRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckScaleOutBalancedRequest) SetResourceOwnerId(v int64) *CheckScaleOutBalancedRequest {
	s.ResourceOwnerId = &v
	return s
}

type CheckScaleOutBalancedResponseBody struct {
	CheckCode    *string                                        `json:"CheckCode,omitempty" xml:"CheckCode,omitempty"`
	PageNumber   *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TableDetails *CheckScaleOutBalancedResponseBodyTableDetails `json:"TableDetails,omitempty" xml:"TableDetails,omitempty" type:"Struct"`
	TimeDuration *string                                        `json:"TimeDuration,omitempty" xml:"TimeDuration,omitempty"`
	TotalCount   *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CheckScaleOutBalancedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckScaleOutBalancedResponseBody) GoString() string {
	return s.String()
}

func (s *CheckScaleOutBalancedResponseBody) SetCheckCode(v string) *CheckScaleOutBalancedResponseBody {
	s.CheckCode = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBody) SetPageNumber(v int32) *CheckScaleOutBalancedResponseBody {
	s.PageNumber = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBody) SetPageSize(v int32) *CheckScaleOutBalancedResponseBody {
	s.PageSize = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBody) SetRequestId(v string) *CheckScaleOutBalancedResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBody) SetTableDetails(v *CheckScaleOutBalancedResponseBodyTableDetails) *CheckScaleOutBalancedResponseBody {
	s.TableDetails = v
	return s
}

func (s *CheckScaleOutBalancedResponseBody) SetTimeDuration(v string) *CheckScaleOutBalancedResponseBody {
	s.TimeDuration = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBody) SetTotalCount(v int32) *CheckScaleOutBalancedResponseBody {
	s.TotalCount = &v
	return s
}

type CheckScaleOutBalancedResponseBodyTableDetails struct {
	TableDetail []*CheckScaleOutBalancedResponseBodyTableDetailsTableDetail `json:"TableDetail,omitempty" xml:"TableDetail,omitempty" type:"Repeated"`
}

func (s CheckScaleOutBalancedResponseBodyTableDetails) String() string {
	return tea.Prettify(s)
}

func (s CheckScaleOutBalancedResponseBodyTableDetails) GoString() string {
	return s.String()
}

func (s *CheckScaleOutBalancedResponseBodyTableDetails) SetTableDetail(v []*CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) *CheckScaleOutBalancedResponseBodyTableDetails {
	s.TableDetail = v
	return s
}

type CheckScaleOutBalancedResponseBodyTableDetailsTableDetail struct {
	Cluster   *string `json:"Cluster,omitempty" xml:"Cluster,omitempty"`
	Database  *string `json:"Database,omitempty" xml:"Database,omitempty"`
	Detail    *int32  `json:"Detail,omitempty" xml:"Detail,omitempty"`
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) String() string {
	return tea.Prettify(s)
}

func (s CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) GoString() string {
	return s.String()
}

func (s *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) SetCluster(v string) *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail {
	s.Cluster = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) SetDatabase(v string) *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail {
	s.Database = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) SetDetail(v int32) *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail {
	s.Detail = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) SetTableName(v string) *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail {
	s.TableName = &v
	return s
}

type CheckScaleOutBalancedResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckScaleOutBalancedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckScaleOutBalancedResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckScaleOutBalancedResponse) GoString() string {
	return s.String()
}

func (s *CheckScaleOutBalancedResponse) SetHeaders(v map[string]*string) *CheckScaleOutBalancedResponse {
	s.Headers = v
	return s
}

func (s *CheckScaleOutBalancedResponse) SetBody(v *CheckScaleOutBalancedResponseBody) *CheckScaleOutBalancedResponse {
	s.Body = v
	return s
}

type CheckServiceLinkedRoleRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleRequest) SetOwnerAccount(v string) *CheckServiceLinkedRoleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) SetOwnerId(v int64) *CheckServiceLinkedRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) SetResourceOwnerAccount(v string) *CheckServiceLinkedRoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) SetResourceOwnerId(v int64) *CheckServiceLinkedRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

type CheckServiceLinkedRoleResponseBody struct {
	HasServiceLinkedRole *bool   `json:"HasServiceLinkedRole,omitempty" xml:"HasServiceLinkedRole,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponseBody) SetHasServiceLinkedRole(v bool) *CheckServiceLinkedRoleResponseBody {
	s.HasServiceLinkedRole = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

type CheckServiceLinkedRoleResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *CheckServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceLinkedRoleResponse) SetBody(v *CheckServiceLinkedRoleResponseBody) *CheckServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type CheckVersionTransferRequest struct {
	CheckAccount         *bool   `json:"CheckAccount,omitempty" xml:"CheckAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SourceAccount        *string `json:"SourceAccount,omitempty" xml:"SourceAccount,omitempty"`
	SourcePassword       *string `json:"SourcePassword,omitempty" xml:"SourcePassword,omitempty"`
	TargetAccount        *string `json:"TargetAccount,omitempty" xml:"TargetAccount,omitempty"`
	TargetDbClusterId    *string `json:"TargetDbClusterId,omitempty" xml:"TargetDbClusterId,omitempty"`
	TargetPassword       *string `json:"TargetPassword,omitempty" xml:"TargetPassword,omitempty"`
}

func (s CheckVersionTransferRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckVersionTransferRequest) GoString() string {
	return s.String()
}

func (s *CheckVersionTransferRequest) SetCheckAccount(v bool) *CheckVersionTransferRequest {
	s.CheckAccount = &v
	return s
}

func (s *CheckVersionTransferRequest) SetDBClusterId(v string) *CheckVersionTransferRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckVersionTransferRequest) SetOwnerAccount(v string) *CheckVersionTransferRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckVersionTransferRequest) SetOwnerId(v int64) *CheckVersionTransferRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckVersionTransferRequest) SetPageNumber(v int32) *CheckVersionTransferRequest {
	s.PageNumber = &v
	return s
}

func (s *CheckVersionTransferRequest) SetPageSize(v int32) *CheckVersionTransferRequest {
	s.PageSize = &v
	return s
}

func (s *CheckVersionTransferRequest) SetRegionId(v string) *CheckVersionTransferRequest {
	s.RegionId = &v
	return s
}

func (s *CheckVersionTransferRequest) SetResourceOwnerAccount(v string) *CheckVersionTransferRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckVersionTransferRequest) SetResourceOwnerId(v int64) *CheckVersionTransferRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckVersionTransferRequest) SetSourceAccount(v string) *CheckVersionTransferRequest {
	s.SourceAccount = &v
	return s
}

func (s *CheckVersionTransferRequest) SetSourcePassword(v string) *CheckVersionTransferRequest {
	s.SourcePassword = &v
	return s
}

func (s *CheckVersionTransferRequest) SetTargetAccount(v string) *CheckVersionTransferRequest {
	s.TargetAccount = &v
	return s
}

func (s *CheckVersionTransferRequest) SetTargetDbClusterId(v string) *CheckVersionTransferRequest {
	s.TargetDbClusterId = &v
	return s
}

func (s *CheckVersionTransferRequest) SetTargetPassword(v string) *CheckVersionTransferRequest {
	s.TargetPassword = &v
	return s
}

type CheckVersionTransferResponseBody struct {
	CheckAccess     *CheckVersionTransferResponseBodyCheckAccess `json:"CheckAccess,omitempty" xml:"CheckAccess,omitempty" type:"Struct"`
	CheckCategory   *bool                                        `json:"CheckCategory,omitempty" xml:"CheckCategory,omitempty"`
	CheckCode       *CheckVersionTransferResponseBodyCheckCode   `json:"CheckCode,omitempty" xml:"CheckCode,omitempty" type:"Struct"`
	CheckConnection *bool                                        `json:"CheckConnection,omitempty" xml:"CheckConnection,omitempty"`
	CheckStatus     *CheckVersionTransferResponseBodyCheckStatus `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty" type:"Struct"`
	CheckStorage    *bool                                        `json:"CheckStorage,omitempty" xml:"CheckStorage,omitempty"`
	RequestId       *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeDuration    *string                                      `json:"TimeDuration,omitempty" xml:"TimeDuration,omitempty"`
	TotalCount      *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CheckVersionTransferResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckVersionTransferResponseBody) GoString() string {
	return s.String()
}

func (s *CheckVersionTransferResponseBody) SetCheckAccess(v *CheckVersionTransferResponseBodyCheckAccess) *CheckVersionTransferResponseBody {
	s.CheckAccess = v
	return s
}

func (s *CheckVersionTransferResponseBody) SetCheckCategory(v bool) *CheckVersionTransferResponseBody {
	s.CheckCategory = &v
	return s
}

func (s *CheckVersionTransferResponseBody) SetCheckCode(v *CheckVersionTransferResponseBodyCheckCode) *CheckVersionTransferResponseBody {
	s.CheckCode = v
	return s
}

func (s *CheckVersionTransferResponseBody) SetCheckConnection(v bool) *CheckVersionTransferResponseBody {
	s.CheckConnection = &v
	return s
}

func (s *CheckVersionTransferResponseBody) SetCheckStatus(v *CheckVersionTransferResponseBodyCheckStatus) *CheckVersionTransferResponseBody {
	s.CheckStatus = v
	return s
}

func (s *CheckVersionTransferResponseBody) SetCheckStorage(v bool) *CheckVersionTransferResponseBody {
	s.CheckStorage = &v
	return s
}

func (s *CheckVersionTransferResponseBody) SetRequestId(v string) *CheckVersionTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckVersionTransferResponseBody) SetTimeDuration(v string) *CheckVersionTransferResponseBody {
	s.TimeDuration = &v
	return s
}

func (s *CheckVersionTransferResponseBody) SetTotalCount(v int32) *CheckVersionTransferResponseBody {
	s.TotalCount = &v
	return s
}

type CheckVersionTransferResponseBodyCheckAccess struct {
	ErrCode    *int64  `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s CheckVersionTransferResponseBodyCheckAccess) String() string {
	return tea.Prettify(s)
}

func (s CheckVersionTransferResponseBodyCheckAccess) GoString() string {
	return s.String()
}

func (s *CheckVersionTransferResponseBodyCheckAccess) SetErrCode(v int64) *CheckVersionTransferResponseBodyCheckAccess {
	s.ErrCode = &v
	return s
}

func (s *CheckVersionTransferResponseBodyCheckAccess) SetErrMessage(v string) *CheckVersionTransferResponseBodyCheckAccess {
	s.ErrMessage = &v
	return s
}

type CheckVersionTransferResponseBodyCheckCode struct {
	ErrCode    *int64  `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s CheckVersionTransferResponseBodyCheckCode) String() string {
	return tea.Prettify(s)
}

func (s CheckVersionTransferResponseBodyCheckCode) GoString() string {
	return s.String()
}

func (s *CheckVersionTransferResponseBodyCheckCode) SetErrCode(v int64) *CheckVersionTransferResponseBodyCheckCode {
	s.ErrCode = &v
	return s
}

func (s *CheckVersionTransferResponseBodyCheckCode) SetErrMessage(v string) *CheckVersionTransferResponseBodyCheckCode {
	s.ErrMessage = &v
	return s
}

type CheckVersionTransferResponseBodyCheckStatus struct {
	ErrCode    *int64  `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s CheckVersionTransferResponseBodyCheckStatus) String() string {
	return tea.Prettify(s)
}

func (s CheckVersionTransferResponseBodyCheckStatus) GoString() string {
	return s.String()
}

func (s *CheckVersionTransferResponseBodyCheckStatus) SetErrCode(v int64) *CheckVersionTransferResponseBodyCheckStatus {
	s.ErrCode = &v
	return s
}

func (s *CheckVersionTransferResponseBodyCheckStatus) SetErrMessage(v string) *CheckVersionTransferResponseBodyCheckStatus {
	s.ErrMessage = &v
	return s
}

type CheckVersionTransferResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckVersionTransferResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckVersionTransferResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckVersionTransferResponse) GoString() string {
	return s.String()
}

func (s *CheckVersionTransferResponse) SetHeaders(v map[string]*string) *CheckVersionTransferResponse {
	s.Headers = v
	return s
}

func (s *CheckVersionTransferResponse) SetBody(v *CheckVersionTransferResponseBody) *CheckVersionTransferResponse {
	s.Body = v
	return s
}

type CreateAccountRequest struct {
	AccountDescription   *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccountResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAccountResponse) SetBody(v *CreateAccountResponseBody) *CreateAccountResponse {
	s.Body = v
	return s
}

type CreateAccountAndAuthorityRequest struct {
	AccountDescription   *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	AllowDatabases       *string `json:"AllowDatabases,omitempty" xml:"AllowDatabases,omitempty"`
	AllowDictionaries    *string `json:"AllowDictionaries,omitempty" xml:"AllowDictionaries,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DdlAuthority         *bool   `json:"DdlAuthority,omitempty" xml:"DdlAuthority,omitempty"`
	DmlAuthority         *string `json:"DmlAuthority,omitempty" xml:"DmlAuthority,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TotalDatabases       *string `json:"TotalDatabases,omitempty" xml:"TotalDatabases,omitempty"`
	TotalDictionaries    *string `json:"TotalDictionaries,omitempty" xml:"TotalDictionaries,omitempty"`
}

func (s CreateAccountAndAuthorityRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountAndAuthorityRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountAndAuthorityRequest) SetAccountDescription(v string) *CreateAccountAndAuthorityRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetAccountName(v string) *CreateAccountAndAuthorityRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetAccountPassword(v string) *CreateAccountAndAuthorityRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetAllowDatabases(v string) *CreateAccountAndAuthorityRequest {
	s.AllowDatabases = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetAllowDictionaries(v string) *CreateAccountAndAuthorityRequest {
	s.AllowDictionaries = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetDBClusterId(v string) *CreateAccountAndAuthorityRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetDdlAuthority(v bool) *CreateAccountAndAuthorityRequest {
	s.DdlAuthority = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetDmlAuthority(v string) *CreateAccountAndAuthorityRequest {
	s.DmlAuthority = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetOwnerAccount(v string) *CreateAccountAndAuthorityRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetOwnerId(v int64) *CreateAccountAndAuthorityRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetRegionId(v string) *CreateAccountAndAuthorityRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetResourceOwnerAccount(v string) *CreateAccountAndAuthorityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetResourceOwnerId(v int64) *CreateAccountAndAuthorityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetTotalDatabases(v string) *CreateAccountAndAuthorityRequest {
	s.TotalDatabases = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetTotalDictionaries(v string) *CreateAccountAndAuthorityRequest {
	s.TotalDictionaries = &v
	return s
}

type CreateAccountAndAuthorityResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccountAndAuthorityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountAndAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountAndAuthorityResponseBody) SetRequestId(v string) *CreateAccountAndAuthorityResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccountAndAuthorityResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAccountAndAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccountAndAuthorityResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountAndAuthorityResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountAndAuthorityResponse) SetHeaders(v map[string]*string) *CreateAccountAndAuthorityResponse {
	s.Headers = v
	return s
}

func (s *CreateAccountAndAuthorityResponse) SetBody(v *CreateAccountAndAuthorityResponseBody) *CreateAccountAndAuthorityResponse {
	s.Body = v
	return s
}

type CreateBackupPolicyRequest struct {
	BackupRetentionPeriod *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	DBClusterId           *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupTime   *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupPolicyRequest) SetBackupRetentionPeriod(v string) *CreateBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetDBClusterId(v string) *CreateBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetOwnerAccount(v string) *CreateBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetOwnerId(v int64) *CreateBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetPreferredBackupPeriod(v string) *CreateBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetPreferredBackupTime(v string) *CreateBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetResourceOwnerAccount(v string) *CreateBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetResourceOwnerId(v int64) *CreateBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateBackupPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupPolicyResponseBody) SetRequestId(v string) *CreateBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreateBackupPolicyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateBackupPolicyResponse) SetHeaders(v map[string]*string) *CreateBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateBackupPolicyResponse) SetBody(v *CreateBackupPolicyResponseBody) *CreateBackupPolicyResponse {
	s.Body = v
	return s
}

type CreateDBInstanceRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBClusterCategory    *string `json:"DBClusterCategory,omitempty" xml:"DBClusterCategory,omitempty"`
	DBClusterClass       *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterVersion     *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	DBNodeGroupCount     *string `json:"DBNodeGroupCount,omitempty" xml:"DBNodeGroupCount,omitempty"`
	DBNodeStorage        *string `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	DbNodeStorageType    *string `json:"DbNodeStorageType,omitempty" xml:"DbNodeStorageType,omitempty"`
	EncryptionKey        *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	EncryptionType       *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UsedTime             *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VPCId                *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBClusterCategory(v string) *CreateDBInstanceRequest {
	s.DBClusterCategory = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBClusterClass(v string) *CreateDBInstanceRequest {
	s.DBClusterClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBClusterDescription(v string) *CreateDBInstanceRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBClusterNetworkType(v string) *CreateDBInstanceRequest {
	s.DBClusterNetworkType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBClusterVersion(v string) *CreateDBInstanceRequest {
	s.DBClusterVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBNodeGroupCount(v string) *CreateDBInstanceRequest {
	s.DBNodeGroupCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBNodeStorage(v string) *CreateDBInstanceRequest {
	s.DBNodeStorage = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDbNodeStorageType(v string) *CreateDBInstanceRequest {
	s.DbNodeStorageType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEncryptionKey(v string) *CreateDBInstanceRequest {
	s.EncryptionKey = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEncryptionType(v string) *CreateDBInstanceRequest {
	s.EncryptionType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetOwnerAccount(v string) *CreateDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetOwnerId(v int64) *CreateDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPayType(v string) *CreateDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPeriod(v string) *CreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRegionId(v string) *CreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceOwnerAccount(v string) *CreateDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceOwnerId(v int64) *CreateDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetUsedTime(v string) *CreateDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVPCId(v string) *CreateDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchId(v string) *CreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

type CreateDBInstanceResponseBody struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OrderId     *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBody) SetDBClusterId(v string) *CreateDBInstanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetOrderId(v string) *CreateDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetRequestId(v string) *CreateDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateDBInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponse) SetHeaders(v map[string]*string) *CreateDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateDBInstanceResponse) SetBody(v *CreateDBInstanceResponseBody) *CreateDBInstanceResponse {
	s.Body = v
	return s
}

type CreateOSSStorageRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateOSSStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOSSStorageRequest) GoString() string {
	return s.String()
}

func (s *CreateOSSStorageRequest) SetDBClusterId(v string) *CreateOSSStorageRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateOSSStorageRequest) SetOwnerAccount(v string) *CreateOSSStorageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateOSSStorageRequest) SetOwnerId(v int64) *CreateOSSStorageRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateOSSStorageRequest) SetRegionId(v string) *CreateOSSStorageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOSSStorageRequest) SetResourceOwnerAccount(v string) *CreateOSSStorageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateOSSStorageRequest) SetResourceOwnerId(v int64) *CreateOSSStorageRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateOSSStorageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOSSStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOSSStorageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOSSStorageResponseBody) SetRequestId(v string) *CreateOSSStorageResponseBody {
	s.RequestId = &v
	return s
}

type CreateOSSStorageResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOSSStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOSSStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOSSStorageResponse) GoString() string {
	return s.String()
}

func (s *CreateOSSStorageResponse) SetHeaders(v map[string]*string) *CreateOSSStorageResponse {
	s.Headers = v
	return s
}

func (s *CreateOSSStorageResponse) SetBody(v *CreateOSSStorageResponseBody) *CreateOSSStorageResponse {
	s.Body = v
	return s
}

type CreatePortsForClickHouseRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PortType             *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreatePortsForClickHouseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePortsForClickHouseRequest) GoString() string {
	return s.String()
}

func (s *CreatePortsForClickHouseRequest) SetDBClusterId(v string) *CreatePortsForClickHouseRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreatePortsForClickHouseRequest) SetOwnerAccount(v string) *CreatePortsForClickHouseRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreatePortsForClickHouseRequest) SetOwnerId(v int64) *CreatePortsForClickHouseRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePortsForClickHouseRequest) SetPortType(v string) *CreatePortsForClickHouseRequest {
	s.PortType = &v
	return s
}

func (s *CreatePortsForClickHouseRequest) SetRegionId(v string) *CreatePortsForClickHouseRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePortsForClickHouseRequest) SetResourceOwnerAccount(v string) *CreatePortsForClickHouseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePortsForClickHouseRequest) SetResourceOwnerId(v int64) *CreatePortsForClickHouseRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreatePortsForClickHouseResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePortsForClickHouseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePortsForClickHouseResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePortsForClickHouseResponseBody) SetRequestId(v string) *CreatePortsForClickHouseResponseBody {
	s.RequestId = &v
	return s
}

type CreatePortsForClickHouseResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePortsForClickHouseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePortsForClickHouseResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePortsForClickHouseResponse) GoString() string {
	return s.String()
}

func (s *CreatePortsForClickHouseResponse) SetHeaders(v map[string]*string) *CreatePortsForClickHouseResponse {
	s.Headers = v
	return s
}

func (s *CreatePortsForClickHouseResponse) SetBody(v *CreatePortsForClickHouseResponseBody) *CreatePortsForClickHouseResponse {
	s.Body = v
	return s
}

type CreateRDSToClickhouseDbRequest struct {
	CkPassword           *string `json:"CkPassword,omitempty" xml:"CkPassword,omitempty"`
	CkUserName           *string `json:"CkUserName,omitempty" xml:"CkUserName,omitempty"`
	ClickhousePort       *int64  `json:"ClickhousePort,omitempty" xml:"ClickhousePort,omitempty"`
	DbClusterId          *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	LimitUpper           *int64  `json:"LimitUpper,omitempty" xml:"LimitUpper,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RdsId                *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	RdsPassword          *string `json:"RdsPassword,omitempty" xml:"RdsPassword,omitempty"`
	RdsPort              *int64  `json:"RdsPort,omitempty" xml:"RdsPort,omitempty"`
	RdsUserName          *string `json:"RdsUserName,omitempty" xml:"RdsUserName,omitempty"`
	RdsVpcId             *string `json:"RdsVpcId,omitempty" xml:"RdsVpcId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SkipUnsupported      *bool   `json:"SkipUnsupported,omitempty" xml:"SkipUnsupported,omitempty"`
	SynDbTables          *string `json:"SynDbTables,omitempty" xml:"SynDbTables,omitempty"`
}

func (s CreateRDSToClickhouseDbRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRDSToClickhouseDbRequest) GoString() string {
	return s.String()
}

func (s *CreateRDSToClickhouseDbRequest) SetCkPassword(v string) *CreateRDSToClickhouseDbRequest {
	s.CkPassword = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetCkUserName(v string) *CreateRDSToClickhouseDbRequest {
	s.CkUserName = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetClickhousePort(v int64) *CreateRDSToClickhouseDbRequest {
	s.ClickhousePort = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetDbClusterId(v string) *CreateRDSToClickhouseDbRequest {
	s.DbClusterId = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetLimitUpper(v int64) *CreateRDSToClickhouseDbRequest {
	s.LimitUpper = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetOwnerAccount(v string) *CreateRDSToClickhouseDbRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetOwnerId(v int64) *CreateRDSToClickhouseDbRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetRdsId(v string) *CreateRDSToClickhouseDbRequest {
	s.RdsId = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetRdsPassword(v string) *CreateRDSToClickhouseDbRequest {
	s.RdsPassword = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetRdsPort(v int64) *CreateRDSToClickhouseDbRequest {
	s.RdsPort = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetRdsUserName(v string) *CreateRDSToClickhouseDbRequest {
	s.RdsUserName = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetRdsVpcId(v string) *CreateRDSToClickhouseDbRequest {
	s.RdsVpcId = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetResourceOwnerAccount(v string) *CreateRDSToClickhouseDbRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetResourceOwnerId(v int64) *CreateRDSToClickhouseDbRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetSkipUnsupported(v bool) *CreateRDSToClickhouseDbRequest {
	s.SkipUnsupported = &v
	return s
}

func (s *CreateRDSToClickhouseDbRequest) SetSynDbTables(v string) *CreateRDSToClickhouseDbRequest {
	s.SynDbTables = &v
	return s
}

type CreateRDSToClickhouseDbResponseBody struct {
	ErrorMsg    *string   `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RepeatedDbs []*string `json:"RepeatedDbs,omitempty" xml:"RepeatedDbs,omitempty" type:"Repeated"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status      *int64    `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateRDSToClickhouseDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRDSToClickhouseDbResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRDSToClickhouseDbResponseBody) SetErrorMsg(v string) *CreateRDSToClickhouseDbResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateRDSToClickhouseDbResponseBody) SetRepeatedDbs(v []*string) *CreateRDSToClickhouseDbResponseBody {
	s.RepeatedDbs = v
	return s
}

func (s *CreateRDSToClickhouseDbResponseBody) SetRequestId(v string) *CreateRDSToClickhouseDbResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRDSToClickhouseDbResponseBody) SetStatus(v int64) *CreateRDSToClickhouseDbResponseBody {
	s.Status = &v
	return s
}

type CreateRDSToClickhouseDbResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRDSToClickhouseDbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRDSToClickhouseDbResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRDSToClickhouseDbResponse) GoString() string {
	return s.String()
}

func (s *CreateRDSToClickhouseDbResponse) SetHeaders(v map[string]*string) *CreateRDSToClickhouseDbResponse {
	s.Headers = v
	return s
}

func (s *CreateRDSToClickhouseDbResponse) SetBody(v *CreateRDSToClickhouseDbResponseBody) *CreateRDSToClickhouseDbResponse {
	s.Body = v
	return s
}

type CreateServiceLinkedRoleRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) SetOwnerAccount(v string) *CreateServiceLinkedRoleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetOwnerId(v int64) *CreateServiceLinkedRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetResourceOwnerAccount(v string) *CreateServiceLinkedRoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetResourceOwnerId(v int64) *CreateServiceLinkedRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateServiceLinkedRoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceLinkedRoleResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *CreateServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetBody(v *CreateServiceLinkedRoleResponseBody) *CreateServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAccountResponse) SetBody(v *DeleteAccountResponseBody) *DeleteAccountResponse {
	s.Body = v
	return s
}

type DeleteDBClusterRequest struct {
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponseBody) SetRequestId(v string) *DeleteDBClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteDBClusterResponse) SetBody(v *DeleteDBClusterResponseBody) *DeleteDBClusterResponse {
	s.Body = v
	return s
}

type DeleteLorneTaskRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteLorneTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLorneTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteLorneTaskRequest) SetDBClusterId(v string) *DeleteLorneTaskRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteLorneTaskRequest) SetOwnerAccount(v string) *DeleteLorneTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteLorneTaskRequest) SetOwnerId(v int64) *DeleteLorneTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLorneTaskRequest) SetRegionId(v string) *DeleteLorneTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLorneTaskRequest) SetResourceOwnerAccount(v string) *DeleteLorneTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteLorneTaskRequest) SetResourceOwnerId(v int64) *DeleteLorneTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteLorneTaskRequest) SetTaskId(v string) *DeleteLorneTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteLorneTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLorneTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLorneTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLorneTaskResponseBody) SetRequestId(v string) *DeleteLorneTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLorneTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLorneTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLorneTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLorneTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteLorneTaskResponse) SetHeaders(v map[string]*string) *DeleteLorneTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteLorneTaskResponse) SetBody(v *DeleteLorneTaskResponseBody) *DeleteLorneTaskResponse {
	s.Body = v
	return s
}

type DeleteSyndbRequest struct {
	DbClusterId          *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SynDb                *string `json:"SynDb,omitempty" xml:"SynDb,omitempty"`
}

func (s DeleteSyndbRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSyndbRequest) GoString() string {
	return s.String()
}

func (s *DeleteSyndbRequest) SetDbClusterId(v string) *DeleteSyndbRequest {
	s.DbClusterId = &v
	return s
}

func (s *DeleteSyndbRequest) SetOwnerAccount(v string) *DeleteSyndbRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSyndbRequest) SetOwnerId(v int64) *DeleteSyndbRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSyndbRequest) SetResourceOwnerAccount(v string) *DeleteSyndbRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSyndbRequest) SetResourceOwnerId(v int64) *DeleteSyndbRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSyndbRequest) SetSynDb(v string) *DeleteSyndbRequest {
	s.SynDb = &v
	return s
}

type DeleteSyndbResponseBody struct {
	ErrorCode *int64  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteSyndbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSyndbResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSyndbResponseBody) SetErrorCode(v int64) *DeleteSyndbResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSyndbResponseBody) SetErrorMsg(v string) *DeleteSyndbResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteSyndbResponseBody) SetRequestId(v string) *DeleteSyndbResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSyndbResponseBody) SetStatus(v bool) *DeleteSyndbResponseBody {
	s.Status = &v
	return s
}

type DeleteSyndbResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSyndbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSyndbResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSyndbResponse) GoString() string {
	return s.String()
}

func (s *DeleteSyndbResponse) SetHeaders(v map[string]*string) *DeleteSyndbResponse {
	s.Headers = v
	return s
}

func (s *DeleteSyndbResponse) SetBody(v *DeleteSyndbResponseBody) *DeleteSyndbResponse {
	s.Body = v
	return s
}

type DescribeAccountAuthorityRequest struct {
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAccountAuthorityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAuthorityRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountAuthorityRequest) SetAccountName(v string) *DescribeAccountAuthorityRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetDBClusterId(v string) *DescribeAccountAuthorityRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetOwnerAccount(v string) *DescribeAccountAuthorityRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetOwnerId(v int64) *DescribeAccountAuthorityRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetRegionId(v string) *DescribeAccountAuthorityRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetResourceOwnerAccount(v string) *DescribeAccountAuthorityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetResourceOwnerId(v int64) *DescribeAccountAuthorityRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeAccountAuthorityResponseBody struct {
	AccountName       *string   `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AllowDatabases    []*string `json:"AllowDatabases,omitempty" xml:"AllowDatabases,omitempty" type:"Repeated"`
	AllowDictionaries []*string `json:"AllowDictionaries,omitempty" xml:"AllowDictionaries,omitempty" type:"Repeated"`
	DdlAuthority      *bool     `json:"DdlAuthority,omitempty" xml:"DdlAuthority,omitempty"`
	DmlAuthority      *string   `json:"DmlAuthority,omitempty" xml:"DmlAuthority,omitempty"`
	RequestId         *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalDatabases    []*string `json:"TotalDatabases,omitempty" xml:"TotalDatabases,omitempty" type:"Repeated"`
	TotalDictionaries []*string `json:"TotalDictionaries,omitempty" xml:"TotalDictionaries,omitempty" type:"Repeated"`
}

func (s DescribeAccountAuthorityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountAuthorityResponseBody) SetAccountName(v string) *DescribeAccountAuthorityResponseBody {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetAllowDatabases(v []*string) *DescribeAccountAuthorityResponseBody {
	s.AllowDatabases = v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetAllowDictionaries(v []*string) *DescribeAccountAuthorityResponseBody {
	s.AllowDictionaries = v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetDdlAuthority(v bool) *DescribeAccountAuthorityResponseBody {
	s.DdlAuthority = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetDmlAuthority(v string) *DescribeAccountAuthorityResponseBody {
	s.DmlAuthority = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetRequestId(v string) *DescribeAccountAuthorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetTotalDatabases(v []*string) *DescribeAccountAuthorityResponseBody {
	s.TotalDatabases = v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetTotalDictionaries(v []*string) *DescribeAccountAuthorityResponseBody {
	s.TotalDictionaries = v
	return s
}

type DescribeAccountAuthorityResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccountAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccountAuthorityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAuthorityResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountAuthorityResponse) SetHeaders(v map[string]*string) *DescribeAccountAuthorityResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountAuthorityResponse) SetBody(v *DescribeAccountAuthorityResponseBody) *DescribeAccountAuthorityResponse {
	s.Body = v
	return s
}

type DescribeAccountsRequest struct {
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *DescribeAccountsRequest) SetPageNumber(v int32) *DescribeAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountsRequest) SetPageSize(v int32) *DescribeAccountsRequest {
	s.PageSize = &v
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
	Accounts   *DescribeAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) SetAccounts(v *DescribeAccountsResponseBodyAccounts) *DescribeAccountsResponseBody {
	s.Accounts = v
	return s
}

func (s *DescribeAccountsResponseBody) SetPageNumber(v int32) *DescribeAccountsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetPageSize(v int32) *DescribeAccountsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetTotalCount(v int32) *DescribeAccountsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAccountsResponseBodyAccounts struct {
	Account []*DescribeAccountsResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccount(v []*DescribeAccountsResponseBodyAccountsAccount) *DescribeAccountsResponseBodyAccounts {
	s.Account = v
	return s
}

type DescribeAccountsResponseBodyAccountsAccount struct {
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountStatus      *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	AccountType        *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetAccountDescription(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetAccountName(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetAccountType(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.AccountType = &v
	return s
}

type DescribeAccountsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAccountsResponse) SetBody(v *DescribeAccountsResponseBody) *DescribeAccountsResponse {
	s.Body = v
	return s
}

type DescribeAllDataSourceRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemaName           *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName            *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
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
	Columns   *DescribeAllDataSourceResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Schemas   *DescribeAllDataSourceResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Struct"`
	Tables    *DescribeAllDataSourceResponseBodyTables  `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
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
	AutoIncrementColumn *bool   `json:"AutoIncrementColumn,omitempty" xml:"AutoIncrementColumn,omitempty"`
	ColumnName          *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	DBClusterId         *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PrimaryKey          *bool   `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	SchemaName          *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName           *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName  *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
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
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName  *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName   *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAllDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAllDataSourceResponse) SetBody(v *DescribeAllDataSourceResponseBody) *DescribeAllDataSourceResponse {
	s.Body = v
	return s
}

type DescribeAllDataSourcesRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemaName           *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName            *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAllDataSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesRequest) SetDBClusterId(v string) *DescribeAllDataSourcesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourcesRequest) SetOwnerAccount(v string) *DescribeAllDataSourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAllDataSourcesRequest) SetOwnerId(v int64) *DescribeAllDataSourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAllDataSourcesRequest) SetResourceOwnerAccount(v string) *DescribeAllDataSourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAllDataSourcesRequest) SetResourceOwnerId(v int64) *DescribeAllDataSourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAllDataSourcesRequest) SetSchemaName(v string) *DescribeAllDataSourcesRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourcesRequest) SetTableName(v string) *DescribeAllDataSourcesRequest {
	s.TableName = &v
	return s
}

type DescribeAllDataSourcesResponseBody struct {
	Columns   *DescribeAllDataSourcesResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Schemas   *DescribeAllDataSourcesResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Struct"`
	Tables    *DescribeAllDataSourcesResponseBodyTables  `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
}

func (s DescribeAllDataSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponseBody) SetColumns(v *DescribeAllDataSourcesResponseBodyColumns) *DescribeAllDataSourcesResponseBody {
	s.Columns = v
	return s
}

func (s *DescribeAllDataSourcesResponseBody) SetRequestId(v string) *DescribeAllDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBody) SetSchemas(v *DescribeAllDataSourcesResponseBodySchemas) *DescribeAllDataSourcesResponseBody {
	s.Schemas = v
	return s
}

func (s *DescribeAllDataSourcesResponseBody) SetTables(v *DescribeAllDataSourcesResponseBodyTables) *DescribeAllDataSourcesResponseBody {
	s.Tables = v
	return s
}

type DescribeAllDataSourcesResponseBodyColumns struct {
	Column []*DescribeAllDataSourcesResponseBodyColumnsColumn `json:"Column,omitempty" xml:"Column,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourcesResponseBodyColumns) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourcesResponseBodyColumns) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponseBodyColumns) SetColumn(v []*DescribeAllDataSourcesResponseBodyColumnsColumn) *DescribeAllDataSourcesResponseBodyColumns {
	s.Column = v
	return s
}

type DescribeAllDataSourcesResponseBodyColumnsColumn struct {
	AutoIncrementColumn *bool   `json:"AutoIncrementColumn,omitempty" xml:"AutoIncrementColumn,omitempty"`
	ColumnName          *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	DBClusterId         *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PrimaryKey          *bool   `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	SchemaName          *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName           *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAllDataSourcesResponseBodyColumnsColumn) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourcesResponseBodyColumnsColumn) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) SetAutoIncrementColumn(v bool) *DescribeAllDataSourcesResponseBodyColumnsColumn {
	s.AutoIncrementColumn = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) SetColumnName(v string) *DescribeAllDataSourcesResponseBodyColumnsColumn {
	s.ColumnName = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) SetDBClusterId(v string) *DescribeAllDataSourcesResponseBodyColumnsColumn {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) SetPrimaryKey(v bool) *DescribeAllDataSourcesResponseBodyColumnsColumn {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) SetSchemaName(v string) *DescribeAllDataSourcesResponseBodyColumnsColumn {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) SetTableName(v string) *DescribeAllDataSourcesResponseBodyColumnsColumn {
	s.TableName = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) SetType(v string) *DescribeAllDataSourcesResponseBodyColumnsColumn {
	s.Type = &v
	return s
}

type DescribeAllDataSourcesResponseBodySchemas struct {
	Schema []*DescribeAllDataSourcesResponseBodySchemasSchema `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourcesResponseBodySchemas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourcesResponseBodySchemas) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponseBodySchemas) SetSchema(v []*DescribeAllDataSourcesResponseBodySchemasSchema) *DescribeAllDataSourcesResponseBodySchemas {
	s.Schema = v
	return s
}

type DescribeAllDataSourcesResponseBodySchemasSchema struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName  *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeAllDataSourcesResponseBodySchemasSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourcesResponseBodySchemasSchema) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponseBodySchemasSchema) SetDBClusterId(v string) *DescribeAllDataSourcesResponseBodySchemasSchema {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodySchemasSchema) SetSchemaName(v string) *DescribeAllDataSourcesResponseBodySchemasSchema {
	s.SchemaName = &v
	return s
}

type DescribeAllDataSourcesResponseBodyTables struct {
	Table []*DescribeAllDataSourcesResponseBodyTablesTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourcesResponseBodyTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourcesResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponseBodyTables) SetTable(v []*DescribeAllDataSourcesResponseBodyTablesTable) *DescribeAllDataSourcesResponseBodyTables {
	s.Table = v
	return s
}

type DescribeAllDataSourcesResponseBodyTablesTable struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName  *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName   *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAllDataSourcesResponseBodyTablesTable) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourcesResponseBodyTablesTable) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponseBodyTablesTable) SetDBClusterId(v string) *DescribeAllDataSourcesResponseBodyTablesTable {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyTablesTable) SetSchemaName(v string) *DescribeAllDataSourcesResponseBodyTablesTable {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyTablesTable) SetTableName(v string) *DescribeAllDataSourcesResponseBodyTablesTable {
	s.TableName = &v
	return s
}

type DescribeAllDataSourcesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAllDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAllDataSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponse) SetHeaders(v map[string]*string) *DescribeAllDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllDataSourcesResponse) SetBody(v *DescribeAllDataSourcesResponseBody) *DescribeAllDataSourcesResponse {
	s.Body = v
	return s
}

type DescribeAvailableResourceRequest struct {
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	ChargeType           *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	AvailableZoneList *DescribeAvailableResourceResponseBodyAvailableZoneList `json:"AvailableZoneList,omitempty" xml:"AvailableZoneList,omitempty" type:"Struct"`
	RegionId          *string                                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) SetAvailableZoneList(v *DescribeAvailableResourceResponseBodyAvailableZoneList) *DescribeAvailableResourceResponseBody {
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
	AvailableZone []*DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZone `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) SetAvailableZone(v []*DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZone) *DescribeAvailableResourceResponseBodyAvailableZoneList {
	s.AvailableZone = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZone struct {
	SupportedSerialList *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialList `json:"SupportedSerialList,omitempty" xml:"SupportedSerialList,omitempty" type:"Struct"`
	ZoneId              *string                                                                                 `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZone) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZone) SetSupportedSerialList(v *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialList) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZone {
	s.SupportedSerialList = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZone) SetZoneId(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZone {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialList struct {
	SupportedSerial []*DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerial `json:"SupportedSerial,omitempty" xml:"SupportedSerial,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialList) SetSupportedSerial(v []*DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerial) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialList {
	s.SupportedSerial = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerial struct {
	Serial                     *string                                                                                                                          `json:"Serial,omitempty" xml:"Serial,omitempty"`
	SupportedInstanceClassList *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassList `json:"SupportedInstanceClassList,omitempty" xml:"SupportedInstanceClassList,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerial) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerial) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerial) SetSerial(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerial {
	s.Serial = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerial) SetSupportedInstanceClassList(v *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassList) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerial {
	s.SupportedInstanceClassList = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassList struct {
	SupportedInstanceClass []*DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClass `json:"SupportedInstanceClass,omitempty" xml:"SupportedInstanceClass,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassList) SetSupportedInstanceClass(v []*DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClass) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassList {
	s.SupportedInstanceClass = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClass struct {
	InstanceClass          *string                                                                                                                                                                      `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	SupportedExecutorList  *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorList  `json:"SupportedExecutorList,omitempty" xml:"SupportedExecutorList,omitempty" type:"Struct"`
	SupportedNodeCountList *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountList `json:"SupportedNodeCountList,omitempty" xml:"SupportedNodeCountList,omitempty" type:"Struct"`
	Tips                   *string                                                                                                                                                                      `json:"Tips,omitempty" xml:"Tips,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClass) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClass) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClass) SetInstanceClass(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClass {
	s.InstanceClass = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClass) SetSupportedExecutorList(v *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorList) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClass {
	s.SupportedExecutorList = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClass) SetSupportedNodeCountList(v *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountList) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClass {
	s.SupportedNodeCountList = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClass) SetTips(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClass {
	s.Tips = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorList struct {
	SupportedExecutor []*DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutor `json:"SupportedExecutor,omitempty" xml:"SupportedExecutor,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorList) SetSupportedExecutor(v []*DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutor) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorList {
	s.SupportedExecutor = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutor struct {
	NodeCount *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutorNodeCount `json:"NodeCount,omitempty" xml:"NodeCount,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutor) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutor) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutor) SetNodeCount(v *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutorNodeCount) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutor {
	s.NodeCount = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutorNodeCount struct {
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	Step     *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutorNodeCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutorNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutorNodeCount) SetMaxCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutorNodeCount {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutorNodeCount) SetMinCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutorNodeCount {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutorNodeCount) SetStep(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedExecutorListSupportedExecutorNodeCount {
	s.Step = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountList struct {
	SupportedNodeCount []*DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCount `json:"SupportedNodeCount,omitempty" xml:"SupportedNodeCount,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountList) SetSupportedNodeCount(v []*DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCount) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountList {
	s.SupportedNodeCount = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCount struct {
	NodeCount   *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountNodeCount   `json:"NodeCount,omitempty" xml:"NodeCount,omitempty" type:"Struct"`
	StorageSize *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountStorageSize `json:"StorageSize,omitempty" xml:"StorageSize,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCount) SetNodeCount(v *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountNodeCount) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCount {
	s.NodeCount = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCount) SetStorageSize(v *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountStorageSize) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCount {
	s.StorageSize = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountNodeCount struct {
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	Step     *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountNodeCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountNodeCount) SetMaxCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountNodeCount {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountNodeCount) SetMinCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountNodeCount {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountNodeCount) SetStep(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountNodeCount {
	s.Step = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountStorageSize struct {
	StorageSize []*string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountStorageSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountStorageSize) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountStorageSize) SetStorageSize(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneListAvailableZoneSupportedSerialListSupportedSerialSupportedInstanceClassListSupportedInstanceClassSupportedNodeCountListSupportedNodeCountStorageSize {
	s.StorageSize = v
	return s
}

type DescribeAvailableResourceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAvailableResourceResponse) SetBody(v *DescribeAvailableResourceResponseBody) *DescribeAvailableResourceResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
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
	BackupRetentionPeriod *int32  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupTime   *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Switch                *string `json:"Switch,omitempty" xml:"Switch,omitempty"`
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

func (s *DescribeBackupPolicyResponseBody) SetSwitch(v string) *DescribeBackupPolicyResponseBody {
	s.Switch = &v
	return s
}

type DescribeBackupPolicyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeBackupPolicyResponse) SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse {
	s.Body = v
	return s
}

type DescribeBackupsRequest struct {
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	Items      []*DescribeBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *string                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) SetItems(v []*DescribeBackupsResponseBodyItems) *DescribeBackupsResponseBody {
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
	BackupEndTime   *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupId        *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupMethod    *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	BackupSetInfo   *string `json:"BackupSetInfo,omitempty" xml:"BackupSetInfo,omitempty"`
	BackupSize      *int32  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupStatus    *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	BackupType      *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	DBClusterId     *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeBackupsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItems) SetBackupEndTime(v string) *DescribeBackupsResponseBodyItems {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetBackupId(v string) *DescribeBackupsResponseBodyItems {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetBackupMethod(v string) *DescribeBackupsResponseBodyItems {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetBackupSetInfo(v string) *DescribeBackupsResponseBodyItems {
	s.BackupSetInfo = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetBackupSize(v int32) *DescribeBackupsResponseBodyItems {
	s.BackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetBackupStartTime(v string) *DescribeBackupsResponseBodyItems {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetBackupStatus(v string) *DescribeBackupsResponseBodyItems {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetBackupType(v string) *DescribeBackupsResponseBodyItems {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetDBClusterId(v string) *DescribeBackupsResponseBodyItems {
	s.DBClusterId = &v
	return s
}

type DescribeBackupsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeBackupsResponse) SetBody(v *DescribeBackupsResponseBody) *DescribeBackupsResponse {
	s.Body = v
	return s
}

type DescribeColumnsRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemaName           *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName            *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
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
	Items     *DescribeColumnsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AutoIncrementColumn *bool   `json:"AutoIncrementColumn,omitempty" xml:"AutoIncrementColumn,omitempty"`
	ColumnName          *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	DBClusterId         *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PrimaryKey          *bool   `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	SchemaName          *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName           *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeColumnsResponse) SetBody(v *DescribeColumnsResponseBody) *DescribeColumnsResponse {
	s.Body = v
	return s
}

type DescribeDBClusterAccessWhiteListRequest struct {
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
	DBClusterAccessWhiteList *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList `json:"DBClusterAccessWhiteList,omitempty" xml:"DBClusterAccessWhiteList,omitempty" type:"Struct"`
	RequestId                *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterAccessWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) SetDBClusterAccessWhiteList(v *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList) *DescribeDBClusterAccessWhiteListResponseBody {
	s.DBClusterAccessWhiteList = v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) SetRequestId(v string) *DescribeDBClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList struct {
	IPArray []*DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray `json:"IPArray,omitempty" xml:"IPArray,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList) SetIPArray(v []*DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList {
	s.IPArray = v
	return s
}

type DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray struct {
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	DBClusterIPArrayName      *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	SecurityIPList            *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) SetDBClusterIPArrayAttribute(v string) *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) SetDBClusterIPArrayName(v string) *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) SetSecurityIPList(v string) *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray {
	s.SecurityIPList = &v
	return s
}

type DescribeDBClusterAccessWhiteListResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDBClusterAccessWhiteListResponse) SetBody(v *DescribeDBClusterAccessWhiteListResponseBody) *DescribeDBClusterAccessWhiteListResponse {
	s.Body = v
	return s
}

type DescribeDBClusterAttributeRequest struct {
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
	DBCluster *DescribeDBClusterAttributeResponseBodyDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Struct"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBCluster(v *DescribeDBClusterAttributeResponseBodyDBCluster) *DescribeDBClusterAttributeResponseBody {
	s.DBCluster = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetRequestId(v string) *DescribeDBClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterAttributeResponseBodyDBCluster struct {
	AliUid                 *string                                                        `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Bid                    *string                                                        `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Category               *string                                                        `json:"Category,omitempty" xml:"Category,omitempty"`
	CommodityCode          *string                                                        `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ConnectionString       *string                                                        `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	ControlVersion         *string                                                        `json:"ControlVersion,omitempty" xml:"ControlVersion,omitempty"`
	CreateTime             *string                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBClusterDescription   *string                                                        `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterId            *string                                                        `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBClusterNetworkType   *string                                                        `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterStatus        *string                                                        `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	DBClusterType          *string                                                        `json:"DBClusterType,omitempty" xml:"DBClusterType,omitempty"`
	DBNodeClass            *string                                                        `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBNodeCount            *int64                                                         `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DBNodeStorage          *int64                                                         `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	EncryptionKey          *string                                                        `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	EncryptionType         *string                                                        `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	Engine                 *string                                                        `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion          *string                                                        `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpireTime             *string                                                        `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	IsExpired              *string                                                        `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	LockMode               *string                                                        `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason             *string                                                        `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MaintainTime           *string                                                        `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	PayType                *string                                                        `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port                   *int32                                                         `json:"Port,omitempty" xml:"Port,omitempty"`
	PublicConnectionString *string                                                        `json:"PublicConnectionString,omitempty" xml:"PublicConnectionString,omitempty"`
	PublicPort             *string                                                        `json:"PublicPort,omitempty" xml:"PublicPort,omitempty"`
	RegionId               *string                                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScaleOutStatus         *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus `json:"ScaleOutStatus,omitempty" xml:"ScaleOutStatus,omitempty" type:"Struct"`
	StorageType            *string                                                        `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	SupportBackup          *int32                                                         `json:"SupportBackup,omitempty" xml:"SupportBackup,omitempty"`
	SupportHttpsPort       *bool                                                          `json:"SupportHttpsPort,omitempty" xml:"SupportHttpsPort,omitempty"`
	SupportMysqlPort       *bool                                                          `json:"SupportMysqlPort,omitempty" xml:"SupportMysqlPort,omitempty"`
	Tags                   *DescribeDBClusterAttributeResponseBodyDBClusterTags           `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VSwitchId              *string                                                        `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcCloudInstanceId     *string                                                        `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	VpcId                  *string                                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                 *string                                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyDBCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetAliUid(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.AliUid = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetBid(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.Bid = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetCategory(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.Category = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetCommodityCode(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetConnectionString(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetControlVersion(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ControlVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetCreateTime(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBClusterDescription(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBClusterId(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBClusterStatus(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBClusterType(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBClusterType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBNodeClass(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBNodeCount(v int64) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBNodeStorage(v int64) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBNodeStorage = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetEncryptionKey(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetEncryptionType(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.EncryptionType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetEngine(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetEngineVersion(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetExpireTime(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetIsExpired(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.IsExpired = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetLockMode(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetLockReason(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.LockReason = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetMaintainTime(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.MaintainTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetPayType(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetPort(v int32) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetPublicConnectionString(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.PublicConnectionString = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetPublicPort(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.PublicPort = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetRegionId(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetScaleOutStatus(v *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ScaleOutStatus = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetStorageType(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.StorageType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetSupportBackup(v int32) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.SupportBackup = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetSupportHttpsPort(v bool) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.SupportHttpsPort = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetSupportMysqlPort(v bool) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.SupportMysqlPort = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetTags(v *DescribeDBClusterAttributeResponseBodyDBClusterTags) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetVSwitchId(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetVpcCloudInstanceId(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetVpcId(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.VpcId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetZoneId(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ZoneId = &v
	return s
}

type DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus struct {
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Ratio    *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus) SetProgress(v string) *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus) SetRatio(v string) *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus {
	s.Ratio = &v
	return s
}

type DescribeDBClusterAttributeResponseBodyDBClusterTags struct {
	Tag []*DescribeDBClusterAttributeResponseBodyDBClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAttributeResponseBodyDBClusterTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterTags) SetTag(v []*DescribeDBClusterAttributeResponseBodyDBClusterTagsTag) *DescribeDBClusterAttributeResponseBodyDBClusterTags {
	s.Tag = v
	return s
}

type DescribeDBClusterAttributeResponseBodyDBClusterTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyDBClusterTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterTagsTag) SetKey(v string) *DescribeDBClusterAttributeResponseBodyDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterTagsTag) SetValue(v string) *DescribeDBClusterAttributeResponseBodyDBClusterTagsTag {
	s.Value = &v
	return s
}

type DescribeDBClusterAttributeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDBClusterAttributeResponse) SetBody(v *DescribeDBClusterAttributeResponseBody) *DescribeDBClusterAttributeResponse {
	s.Body = v
	return s
}

type DescribeDBClusterConfigRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigRequest) SetDBClusterId(v string) *DescribeDBClusterConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetOwnerAccount(v string) *DescribeDBClusterConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetOwnerId(v int64) *DescribeDBClusterConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetRegionId(v string) *DescribeDBClusterConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterConfigRequest) SetResourceOwnerId(v int64) *DescribeDBClusterConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBClusterConfigResponseBody struct {
	Config    *string `json:"Config,omitempty" xml:"Config,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigResponseBody) SetConfig(v string) *DescribeDBClusterConfigResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeDBClusterConfigResponseBody) SetRequestId(v string) *DescribeDBClusterConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigResponse) SetHeaders(v map[string]*string) *DescribeDBClusterConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterConfigResponse) SetBody(v *DescribeDBClusterConfigResponseBody) *DescribeDBClusterConfigResponse {
	s.Body = v
	return s
}

type DescribeDBClusterNetInfoItemsRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterNetInfoItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoItemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoItemsRequest) SetDBClusterId(v string) *DescribeDBClusterNetInfoItemsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsRequest) SetOwnerAccount(v string) *DescribeDBClusterNetInfoItemsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsRequest) SetOwnerId(v int64) *DescribeDBClusterNetInfoItemsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterNetInfoItemsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsRequest) SetResourceOwnerId(v int64) *DescribeDBClusterNetInfoItemsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBClusterNetInfoItemsResponseBody struct {
	ClusterNetworkType *string                                                `json:"ClusterNetworkType,omitempty" xml:"ClusterNetworkType,omitempty"`
	NetInfoItems       *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems `json:"NetInfoItems,omitempty" xml:"NetInfoItems,omitempty" type:"Struct"`
	RequestId          *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterNetInfoItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoItemsResponseBody) SetClusterNetworkType(v string) *DescribeDBClusterNetInfoItemsResponseBody {
	s.ClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBody) SetNetInfoItems(v *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems) *DescribeDBClusterNetInfoItemsResponseBody {
	s.NetInfoItems = v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBody) SetRequestId(v string) *DescribeDBClusterNetInfoItemsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems struct {
	NetInfoItem []*DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem `json:"NetInfoItem,omitempty" xml:"NetInfoItem,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems) SetNetInfoItem(v []*DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems {
	s.NetInfoItem = v
	return s
}

type DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem struct {
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	HttpPort         *string `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	IPAddress        *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	JdbcPort         *string `json:"JdbcPort,omitempty" xml:"JdbcPort,omitempty"`
	NetType          *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId            *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetConnectionString(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetHttpPort(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.HttpPort = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetIPAddress(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetJdbcPort(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.JdbcPort = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetNetType(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetVSwitchId(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetVpcId(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.VpcId = &v
	return s
}

type DescribeDBClusterNetInfoItemsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterNetInfoItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterNetInfoItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterNetInfoItemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoItemsResponse) SetHeaders(v map[string]*string) *DescribeDBClusterNetInfoItemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponse) SetBody(v *DescribeDBClusterNetInfoItemsResponseBody) *DescribeDBClusterNetInfoItemsResponse {
	s.Body = v
	return s
}

type DescribeDBClusterPerformanceRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s *DescribeDBClusterPerformanceRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterPerformanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetResourceOwnerId(v int64) *DescribeDBClusterPerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetStartTime(v string) *DescribeDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

type DescribeDBClusterPerformanceResponseBody struct {
	DBClusterId  *string                                                 `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime      *string                                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Performances []*DescribeDBClusterPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	Key    *string                                                       `json:"Key,omitempty" xml:"Key,omitempty"`
	Name   *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	Series []*DescribeDBClusterPerformanceResponseBodyPerformancesSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	Unit   *string                                                       `json:"Unit,omitempty" xml:"Unit,omitempty"`
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

func (s *DescribeDBClusterPerformanceResponseBodyPerformances) SetName(v string) *DescribeDBClusterPerformanceResponseBodyPerformances {
	s.Name = &v
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
	Name   *string                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Values []*DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
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

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeries) SetValues(v []*DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues) *DescribeDBClusterPerformanceResponseBodyPerformancesSeries {
	s.Values = v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues struct {
	Point []*string `json:"Point,omitempty" xml:"Point,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues) SetPoint(v []*string) *DescribeDBClusterPerformanceResponseBodyPerformancesSeriesValues {
	s.Point = v
	return s
}

type DescribeDBClusterPerformanceResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDBClusterPerformanceResponse) SetBody(v *DescribeDBClusterPerformanceResponseBody) *DescribeDBClusterPerformanceResponse {
	s.Body = v
	return s
}

type DescribeDBClustersRequest struct {
	DBClusterDescription *string                         `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterIds         *string                         `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	DBClusterStatus      *string                         `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	OwnerAccount         *string                         `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                         `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                          `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tag                  []*DescribeDBClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	DBClusters *DescribeDBClustersResponseBodyDBClusters `json:"DBClusters,omitempty" xml:"DBClusters,omitempty" type:"Struct"`
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBody) SetDBClusters(v *DescribeDBClustersResponseBodyDBClusters) *DescribeDBClustersResponseBody {
	s.DBClusters = v
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

type DescribeDBClustersResponseBodyDBClusters struct {
	DBCluster []*DescribeDBClustersResponseBodyDBClustersDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyDBClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyDBClusters) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyDBClusters) SetDBCluster(v []*DescribeDBClustersResponseBodyDBClustersDBCluster) *DescribeDBClustersResponseBodyDBClusters {
	s.DBCluster = v
	return s
}

type DescribeDBClustersResponseBodyDBClustersDBCluster struct {
	AliUid               *string                                                          `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Bid                  *string                                                          `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Category             *string                                                          `json:"Category,omitempty" xml:"Category,omitempty"`
	CommodityCode        *string                                                          `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ConnectionString     *string                                                          `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	ControlVersion       *string                                                          `json:"ControlVersion,omitempty" xml:"ControlVersion,omitempty"`
	CreateTime           *string                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBClusterDescription *string                                                          `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterId          *string                                                          `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBClusterNetworkType *string                                                          `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterStatus      *string                                                          `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	DBNodeClass          *string                                                          `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBNodeCount          *int64                                                           `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DBNodeStorage        *int64                                                           `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	ExpireTime           *string                                                          `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	IsExpired            *string                                                          `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	LockMode             *string                                                          `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason           *string                                                          `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	PayType              *string                                                          `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port                 *int32                                                           `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId             *string                                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScaleOutStatus       *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus `json:"ScaleOutStatus,omitempty" xml:"ScaleOutStatus,omitempty" type:"Struct"`
	StorageType          *string                                                          `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Tags                 *DescribeDBClustersResponseBodyDBClustersDBClusterTags           `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VSwitchId            *string                                                          `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcCloudInstanceId   *string                                                          `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	VpcId                *string                                                          `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string                                                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClustersResponseBodyDBClustersDBCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyDBClustersDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetAliUid(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.AliUid = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetBid(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.Bid = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetCategory(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.Category = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetCommodityCode(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetConnectionString(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetControlVersion(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ControlVersion = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetCreateTime(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBClusterDescription(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBClusterId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBClusterStatus(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBNodeClass(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBNodeCount(v int64) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBNodeStorage(v int64) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBNodeStorage = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetExpireTime(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetIsExpired(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.IsExpired = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetLockMode(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetLockReason(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.LockReason = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetPayType(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetPort(v int32) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.Port = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetRegionId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetScaleOutStatus(v *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ScaleOutStatus = v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetStorageType(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.StorageType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetTags(v *DescribeDBClustersResponseBodyDBClustersDBClusterTags) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetVSwitchId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetVpcCloudInstanceId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetVpcId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.VpcId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetZoneId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ZoneId = &v
	return s
}

type DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus struct {
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Ratio    *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) SetProgress(v string) *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) SetRatio(v string) *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus {
	s.Ratio = &v
	return s
}

type DescribeDBClustersResponseBodyDBClustersDBClusterTags struct {
	Tag []*DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTags) SetTag(v []*DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) *DescribeDBClustersResponseBodyDBClustersDBClusterTags {
	s.Tag = v
	return s
}

type DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) SetKey(v string) *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) SetValue(v string) *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag {
	s.Value = &v
	return s
}

type DescribeDBClustersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDBClustersResponse) SetBody(v *DescribeDBClustersResponseBody) *DescribeDBClustersResponse {
	s.Body = v
	return s
}

type DescribeDBConfigRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBConfigRequest) SetDBClusterId(v string) *DescribeDBConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBConfigRequest) SetOwnerAccount(v string) *DescribeDBConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBConfigRequest) SetOwnerId(v int64) *DescribeDBConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBConfigRequest) SetRegionId(v string) *DescribeDBConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBConfigRequest) SetResourceOwnerAccount(v string) *DescribeDBConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBConfigRequest) SetResourceOwnerId(v int64) *DescribeDBConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDBConfigResponseBody struct {
	Config    *string `json:"Config,omitempty" xml:"Config,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBConfigResponseBody) SetConfig(v string) *DescribeDBConfigResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeDBConfigResponseBody) SetRequestId(v string) *DescribeDBConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBConfigResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBConfigResponse) SetHeaders(v map[string]*string) *DescribeDBConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBConfigResponse) SetBody(v *DescribeDBConfigResponseBody) *DescribeDBConfigResponse {
	s.Body = v
	return s
}

type DescribeLogHubAttributeRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DeliverName          *string `json:"DeliverName,omitempty" xml:"DeliverName,omitempty"`
	LogStoreName         *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectName          *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeLogHubAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogHubAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogHubAttributeRequest) SetDBClusterId(v string) *DescribeLogHubAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetDeliverName(v string) *DescribeLogHubAttributeRequest {
	s.DeliverName = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetLogStoreName(v string) *DescribeLogHubAttributeRequest {
	s.LogStoreName = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetOwnerAccount(v string) *DescribeLogHubAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetOwnerId(v int64) *DescribeLogHubAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetProjectName(v string) *DescribeLogHubAttributeRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetRegionId(v string) *DescribeLogHubAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetResourceOwnerAccount(v string) *DescribeLogHubAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetResourceOwnerId(v int64) *DescribeLogHubAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLogHubAttributeRequest) SetTaskId(v string) *DescribeLogHubAttributeRequest {
	s.TaskId = &v
	return s
}

type DescribeLogHubAttributeResponseBody struct {
	LoghubInfo *DescribeLogHubAttributeResponseBodyLoghubInfo `json:"LoghubInfo,omitempty" xml:"LoghubInfo,omitempty" type:"Struct"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLogHubAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogHubAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogHubAttributeResponseBody) SetLoghubInfo(v *DescribeLogHubAttributeResponseBodyLoghubInfo) *DescribeLogHubAttributeResponseBody {
	s.LoghubInfo = v
	return s
}

func (s *DescribeLogHubAttributeResponseBody) SetRequestId(v string) *DescribeLogHubAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeLogHubAttributeResponseBodyLoghubInfo struct {
	DBClusterId     *string                                                    `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBType          *string                                                    `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DeliverName     *string                                                    `json:"DeliverName,omitempty" xml:"DeliverName,omitempty"`
	DeliverTime     *string                                                    `json:"DeliverTime,omitempty" xml:"DeliverTime,omitempty"`
	Description     *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainUrl       *string                                                    `json:"DomainUrl,omitempty" xml:"DomainUrl,omitempty"`
	FilterDirtyData *string                                                    `json:"FilterDirtyData,omitempty" xml:"FilterDirtyData,omitempty"`
	Id              *string                                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	LogHubStores    *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores `json:"LogHubStores,omitempty" xml:"LogHubStores,omitempty" type:"Struct"`
	LogStoreName    *string                                                    `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	Password        *string                                                    `json:"Password,omitempty" xml:"Password,omitempty"`
	ProjectName     *string                                                    `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RegionId        *string                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SchemaName      *string                                                    `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName       *string                                                    `json:"TableName,omitempty" xml:"TableName,omitempty"`
	UserName        *string                                                    `json:"UserName,omitempty" xml:"UserName,omitempty"`
	ZoneId          *string                                                    `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeLogHubAttributeResponseBodyLoghubInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogHubAttributeResponseBodyLoghubInfo) GoString() string {
	return s.String()
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetDBClusterId(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetDBType(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.DBType = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetDeliverName(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.DeliverName = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetDeliverTime(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.DeliverTime = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetDescription(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.Description = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetDomainUrl(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.DomainUrl = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetFilterDirtyData(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.FilterDirtyData = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetId(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.Id = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetLogHubStores(v *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.LogHubStores = v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetLogStoreName(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.LogStoreName = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetPassword(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.Password = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetProjectName(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.ProjectName = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetRegionId(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetSchemaName(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.SchemaName = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetTableName(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.TableName = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetUserName(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.UserName = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfo) SetZoneId(v string) *DescribeLogHubAttributeResponseBodyLoghubInfo {
	s.ZoneId = &v
	return s
}

type DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores struct {
	LogHubStore []*DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore `json:"LogHubStore,omitempty" xml:"LogHubStore,omitempty" type:"Repeated"`
}

func (s DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores) GoString() string {
	return s.String()
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores) SetLogHubStore(v []*DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore) *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStores {
	s.LogHubStore = v
	return s
}

type DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore struct {
	FieldKey *string `json:"FieldKey,omitempty" xml:"FieldKey,omitempty"`
	LogKey   *string `json:"LogKey,omitempty" xml:"LogKey,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore) GoString() string {
	return s.String()
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore) SetFieldKey(v string) *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore {
	s.FieldKey = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore) SetLogKey(v string) *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore {
	s.LogKey = &v
	return s
}

func (s *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore) SetType(v string) *DescribeLogHubAttributeResponseBodyLoghubInfoLogHubStoresLogHubStore {
	s.Type = &v
	return s
}

type DescribeLogHubAttributeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLogHubAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogHubAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogHubAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogHubAttributeResponse) SetHeaders(v map[string]*string) *DescribeLogHubAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogHubAttributeResponse) SetBody(v *DescribeLogHubAttributeResponseBody) *DescribeLogHubAttributeResponse {
	s.Body = v
	return s
}

type DescribeLogStoreKeysRequest struct {
	LogStoreName         *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectName          *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLogStoreKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreKeysRequest) SetLogStoreName(v string) *DescribeLogStoreKeysRequest {
	s.LogStoreName = &v
	return s
}

func (s *DescribeLogStoreKeysRequest) SetOwnerAccount(v string) *DescribeLogStoreKeysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLogStoreKeysRequest) SetOwnerId(v int64) *DescribeLogStoreKeysRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLogStoreKeysRequest) SetProjectName(v string) *DescribeLogStoreKeysRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeLogStoreKeysRequest) SetRegionId(v string) *DescribeLogStoreKeysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogStoreKeysRequest) SetResourceOwnerAccount(v string) *DescribeLogStoreKeysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLogStoreKeysRequest) SetResourceOwnerId(v int64) *DescribeLogStoreKeysRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeLogStoreKeysResponseBody struct {
	LogStoreKeys *DescribeLogStoreKeysResponseBodyLogStoreKeys `json:"LogStoreKeys,omitempty" xml:"LogStoreKeys,omitempty" type:"Struct"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLogStoreKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreKeysResponseBody) SetLogStoreKeys(v *DescribeLogStoreKeysResponseBodyLogStoreKeys) *DescribeLogStoreKeysResponseBody {
	s.LogStoreKeys = v
	return s
}

func (s *DescribeLogStoreKeysResponseBody) SetRequestId(v string) *DescribeLogStoreKeysResponseBody {
	s.RequestId = &v
	return s
}

type DescribeLogStoreKeysResponseBodyLogStoreKeys struct {
	LogStoreKey []*string `json:"LogStoreKey,omitempty" xml:"LogStoreKey,omitempty" type:"Repeated"`
}

func (s DescribeLogStoreKeysResponseBodyLogStoreKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreKeysResponseBodyLogStoreKeys) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreKeysResponseBodyLogStoreKeys) SetLogStoreKey(v []*string) *DescribeLogStoreKeysResponseBodyLogStoreKeys {
	s.LogStoreKey = v
	return s
}

type DescribeLogStoreKeysResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLogStoreKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogStoreKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreKeysResponse) SetHeaders(v map[string]*string) *DescribeLogStoreKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogStoreKeysResponse) SetBody(v *DescribeLogStoreKeysResponseBody) *DescribeLogStoreKeysResponse {
	s.Body = v
	return s
}

type DescribeLoghubDetailRequest struct {
	ExportName           *string `json:"ExportName,omitempty" xml:"ExportName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectName          *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLoghubDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoghubDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoghubDetailRequest) SetExportName(v string) *DescribeLoghubDetailRequest {
	s.ExportName = &v
	return s
}

func (s *DescribeLoghubDetailRequest) SetOwnerAccount(v string) *DescribeLoghubDetailRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoghubDetailRequest) SetOwnerId(v int64) *DescribeLoghubDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoghubDetailRequest) SetProjectName(v string) *DescribeLoghubDetailRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeLoghubDetailRequest) SetRegionId(v string) *DescribeLoghubDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoghubDetailRequest) SetResourceOwnerAccount(v string) *DescribeLoghubDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoghubDetailRequest) SetResourceOwnerId(v int64) *DescribeLoghubDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeLoghubDetailResponseBody struct {
	LoghubInfo *DescribeLoghubDetailResponseBodyLoghubInfo `json:"LoghubInfo,omitempty" xml:"LoghubInfo,omitempty" type:"Struct"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLoghubDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoghubDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoghubDetailResponseBody) SetLoghubInfo(v *DescribeLoghubDetailResponseBodyLoghubInfo) *DescribeLoghubDetailResponseBody {
	s.LoghubInfo = v
	return s
}

func (s *DescribeLoghubDetailResponseBody) SetRequestId(v string) *DescribeLoghubDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeLoghubDetailResponseBodyLoghubInfo struct {
	AccessKey       *string                                                 `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	AccessSecret    *string                                                 `json:"AccessSecret,omitempty" xml:"AccessSecret,omitempty"`
	DBClusterId     *string                                                 `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBType          *string                                                 `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DeliverName     *string                                                 `json:"DeliverName,omitempty" xml:"DeliverName,omitempty"`
	DeliverTime     *string                                                 `json:"DeliverTime,omitempty" xml:"DeliverTime,omitempty"`
	Description     *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainUrl       *string                                                 `json:"DomainUrl,omitempty" xml:"DomainUrl,omitempty"`
	FilterDirtyData *bool                                                   `json:"FilterDirtyData,omitempty" xml:"FilterDirtyData,omitempty"`
	LogHubStores    *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores `json:"LogHubStores,omitempty" xml:"LogHubStores,omitempty" type:"Struct"`
	LogStoreName    *string                                                 `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	Password        *string                                                 `json:"Password,omitempty" xml:"Password,omitempty"`
	ProjectName     *string                                                 `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RegionId        *string                                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SchemaName      *string                                                 `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName       *string                                                 `json:"TableName,omitempty" xml:"TableName,omitempty"`
	UserName        *string                                                 `json:"UserName,omitempty" xml:"UserName,omitempty"`
	ZoneId          *string                                                 `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeLoghubDetailResponseBodyLoghubInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoghubDetailResponseBodyLoghubInfo) GoString() string {
	return s.String()
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetAccessKey(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.AccessKey = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetAccessSecret(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.AccessSecret = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetDBClusterId(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetDBType(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.DBType = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetDeliverName(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.DeliverName = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetDeliverTime(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.DeliverTime = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetDescription(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.Description = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetDomainUrl(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.DomainUrl = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetFilterDirtyData(v bool) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.FilterDirtyData = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetLogHubStores(v *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.LogHubStores = v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetLogStoreName(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.LogStoreName = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetPassword(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.Password = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetProjectName(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.ProjectName = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetRegionId(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetSchemaName(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.SchemaName = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetTableName(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.TableName = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetUserName(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.UserName = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfo) SetZoneId(v string) *DescribeLoghubDetailResponseBodyLoghubInfo {
	s.ZoneId = &v
	return s
}

type DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores struct {
	LogHubStore []*DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore `json:"LogHubStore,omitempty" xml:"LogHubStore,omitempty" type:"Repeated"`
}

func (s DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores) GoString() string {
	return s.String()
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores) SetLogHubStore(v []*DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore) *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStores {
	s.LogHubStore = v
	return s
}

type DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore struct {
	FieldKey *string `json:"FieldKey,omitempty" xml:"FieldKey,omitempty"`
	LogKey   *string `json:"LogKey,omitempty" xml:"LogKey,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore) GoString() string {
	return s.String()
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore) SetFieldKey(v string) *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore {
	s.FieldKey = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore) SetLogKey(v string) *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore {
	s.LogKey = &v
	return s
}

func (s *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore) SetType(v string) *DescribeLoghubDetailResponseBodyLoghubInfoLogHubStoresLogHubStore {
	s.Type = &v
	return s
}

type DescribeLoghubDetailResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLoghubDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLoghubDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLoghubDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoghubDetailResponse) SetHeaders(v map[string]*string) *DescribeLoghubDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoghubDetailResponse) SetBody(v *DescribeLoghubDetailResponseBody) *DescribeLoghubDetailResponse {
	s.Body = v
	return s
}

type DescribeLorneLogRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeLorneLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLorneLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeLorneLogRequest) SetDBClusterId(v string) *DescribeLorneLogRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLorneLogRequest) SetEndTime(v string) *DescribeLorneLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLorneLogRequest) SetOwnerAccount(v string) *DescribeLorneLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLorneLogRequest) SetOwnerId(v int64) *DescribeLorneLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLorneLogRequest) SetPageNumber(v int32) *DescribeLorneLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLorneLogRequest) SetPageSize(v int32) *DescribeLorneLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLorneLogRequest) SetRegionId(v string) *DescribeLorneLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLorneLogRequest) SetResourceOwnerAccount(v string) *DescribeLorneLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLorneLogRequest) SetResourceOwnerId(v int64) *DescribeLorneLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLorneLogRequest) SetStartTime(v string) *DescribeLorneLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLorneLogRequest) SetTaskId(v string) *DescribeLorneLogRequest {
	s.TaskId = &v
	return s
}

type DescribeLorneLogResponseBody struct {
	Data       []*DescribeLorneLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLorneLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLorneLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLorneLogResponseBody) SetData(v []*DescribeLorneLogResponseBodyData) *DescribeLorneLogResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLorneLogResponseBody) SetPageNumber(v int32) *DescribeLorneLogResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLorneLogResponseBody) SetPageSize(v int32) *DescribeLorneLogResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLorneLogResponseBody) SetRequestId(v string) *DescribeLorneLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLorneLogResponseBody) SetTotalCount(v int32) *DescribeLorneLogResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeLorneLogResponseBodyData struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Count     *string `json:"Count,omitempty" xml:"Count,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLorneLogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLorneLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLorneLogResponseBodyData) SetCode(v string) *DescribeLorneLogResponseBodyData {
	s.Code = &v
	return s
}

func (s *DescribeLorneLogResponseBodyData) SetCount(v string) *DescribeLorneLogResponseBodyData {
	s.Count = &v
	return s
}

func (s *DescribeLorneLogResponseBodyData) SetEndTime(v string) *DescribeLorneLogResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeLorneLogResponseBodyData) SetMessage(v string) *DescribeLorneLogResponseBodyData {
	s.Message = &v
	return s
}

func (s *DescribeLorneLogResponseBodyData) SetStartTime(v string) *DescribeLorneLogResponseBodyData {
	s.StartTime = &v
	return s
}

type DescribeLorneLogResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLorneLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLorneLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLorneLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeLorneLogResponse) SetHeaders(v map[string]*string) *DescribeLorneLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeLorneLogResponse) SetBody(v *DescribeLorneLogResponseBody) *DescribeLorneLogResponse {
	s.Body = v
	return s
}

type DescribeLorneTasksRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLorneTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLorneTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeLorneTasksRequest) SetDBClusterId(v string) *DescribeLorneTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLorneTasksRequest) SetOwnerAccount(v string) *DescribeLorneTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLorneTasksRequest) SetOwnerId(v int64) *DescribeLorneTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLorneTasksRequest) SetPageNumber(v int32) *DescribeLorneTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLorneTasksRequest) SetPageSize(v int32) *DescribeLorneTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLorneTasksRequest) SetRegionId(v string) *DescribeLorneTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLorneTasksRequest) SetResourceOwnerAccount(v string) *DescribeLorneTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLorneTasksRequest) SetResourceOwnerId(v int64) *DescribeLorneTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeLorneTasksResponseBody struct {
	PageNumber  *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskDetails []*DescribeLorneTasksResponseBodyTaskDetails `json:"TaskDetails,omitempty" xml:"TaskDetails,omitempty" type:"Repeated"`
	TotalCount  *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLorneTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLorneTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLorneTasksResponseBody) SetPageNumber(v int32) *DescribeLorneTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLorneTasksResponseBody) SetPageSize(v int32) *DescribeLorneTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLorneTasksResponseBody) SetRequestId(v string) *DescribeLorneTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLorneTasksResponseBody) SetTaskDetails(v []*DescribeLorneTasksResponseBodyTaskDetails) *DescribeLorneTasksResponseBody {
	s.TaskDetails = v
	return s
}

func (s *DescribeLorneTasksResponseBody) SetTotalCount(v int32) *DescribeLorneTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeLorneTasksResponseBodyTaskDetails struct {
	Checkpoint    *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	ColumnMapper  *string `json:"ColumnMapper,omitempty" xml:"ColumnMapper,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SinkInstance  *string `json:"SinkInstance,omitempty" xml:"SinkInstance,omitempty"`
	SinkRegion    *string `json:"SinkRegion,omitempty" xml:"SinkRegion,omitempty"`
	SinkSchema    *string `json:"SinkSchema,omitempty" xml:"SinkSchema,omitempty"`
	SinkTable     *string `json:"SinkTable,omitempty" xml:"SinkTable,omitempty"`
	SinkType      *string `json:"SinkType,omitempty" xml:"SinkType,omitempty"`
	SinkUser      *string `json:"SinkUser,omitempty" xml:"SinkUser,omitempty"`
	SinkVpcId     *string `json:"SinkVpcId,omitempty" xml:"SinkVpcId,omitempty"`
	SourceProject *string `json:"SourceProject,omitempty" xml:"SourceProject,omitempty"`
	SourceRegion  *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	SourceTopic   *string `json:"SourceTopic,omitempty" xml:"SourceTopic,omitempty"`
	SourceType    *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
	Strict        *string `json:"Strict,omitempty" xml:"Strict,omitempty"`
}

func (s DescribeLorneTasksResponseBodyTaskDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeLorneTasksResponseBodyTaskDetails) GoString() string {
	return s.String()
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetCheckpoint(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.Checkpoint = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetColumnMapper(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.ColumnMapper = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetCreateTime(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.CreateTime = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetDescription(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.Description = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetId(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.Id = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetMessage(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.Message = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetName(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.Name = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetSinkInstance(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.SinkInstance = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetSinkRegion(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.SinkRegion = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetSinkSchema(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.SinkSchema = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetSinkTable(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.SinkTable = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetSinkType(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.SinkType = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetSinkUser(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.SinkUser = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetSinkVpcId(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.SinkVpcId = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetSourceProject(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.SourceProject = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetSourceRegion(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.SourceRegion = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetSourceTopic(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.SourceTopic = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetSourceType(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.SourceType = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetState(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.State = &v
	return s
}

func (s *DescribeLorneTasksResponseBodyTaskDetails) SetStrict(v string) *DescribeLorneTasksResponseBodyTaskDetails {
	s.Strict = &v
	return s
}

type DescribeLorneTasksResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLorneTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLorneTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLorneTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeLorneTasksResponse) SetHeaders(v map[string]*string) *DescribeLorneTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeLorneTasksResponse) SetBody(v *DescribeLorneTasksResponseBody) *DescribeLorneTasksResponse {
	s.Body = v
	return s
}

type DescribeLorneTasksMCountRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MetricName           *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeLorneTasksMCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLorneTasksMCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeLorneTasksMCountRequest) SetDBClusterId(v string) *DescribeLorneTasksMCountRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLorneTasksMCountRequest) SetEndTime(v string) *DescribeLorneTasksMCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLorneTasksMCountRequest) SetMetricName(v string) *DescribeLorneTasksMCountRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeLorneTasksMCountRequest) SetOwnerAccount(v string) *DescribeLorneTasksMCountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLorneTasksMCountRequest) SetOwnerId(v int64) *DescribeLorneTasksMCountRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLorneTasksMCountRequest) SetRegionId(v string) *DescribeLorneTasksMCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLorneTasksMCountRequest) SetResourceOwnerAccount(v string) *DescribeLorneTasksMCountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLorneTasksMCountRequest) SetResourceOwnerId(v int64) *DescribeLorneTasksMCountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLorneTasksMCountRequest) SetStartTime(v string) *DescribeLorneTasksMCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLorneTasksMCountRequest) SetTaskId(v string) *DescribeLorneTasksMCountRequest {
	s.TaskId = &v
	return s
}

type DescribeLorneTasksMCountResponseBody struct {
	Data      *float32 `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLorneTasksMCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLorneTasksMCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLorneTasksMCountResponseBody) SetData(v float32) *DescribeLorneTasksMCountResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeLorneTasksMCountResponseBody) SetRequestId(v string) *DescribeLorneTasksMCountResponseBody {
	s.RequestId = &v
	return s
}

type DescribeLorneTasksMCountResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLorneTasksMCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLorneTasksMCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLorneTasksMCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeLorneTasksMCountResponse) SetHeaders(v map[string]*string) *DescribeLorneTasksMCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeLorneTasksMCountResponse) SetBody(v *DescribeLorneTasksMCountResponseBody) *DescribeLorneTasksMCountResponse {
	s.Body = v
	return s
}

type DescribeLorneTasksMetricsRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MetricName           *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeLorneTasksMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLorneTasksMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLorneTasksMetricsRequest) SetDBClusterId(v string) *DescribeLorneTasksMetricsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLorneTasksMetricsRequest) SetEndTime(v string) *DescribeLorneTasksMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLorneTasksMetricsRequest) SetMetricName(v string) *DescribeLorneTasksMetricsRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeLorneTasksMetricsRequest) SetOwnerAccount(v string) *DescribeLorneTasksMetricsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLorneTasksMetricsRequest) SetOwnerId(v int64) *DescribeLorneTasksMetricsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLorneTasksMetricsRequest) SetRegionId(v string) *DescribeLorneTasksMetricsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLorneTasksMetricsRequest) SetResourceOwnerAccount(v string) *DescribeLorneTasksMetricsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLorneTasksMetricsRequest) SetResourceOwnerId(v int64) *DescribeLorneTasksMetricsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLorneTasksMetricsRequest) SetStartTime(v string) *DescribeLorneTasksMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLorneTasksMetricsRequest) SetTaskId(v string) *DescribeLorneTasksMetricsRequest {
	s.TaskId = &v
	return s
}

type DescribeLorneTasksMetricsResponseBody struct {
	Data      *DescribeLorneTasksMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLorneTasksMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLorneTasksMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLorneTasksMetricsResponseBody) SetData(v *DescribeLorneTasksMetricsResponseBodyData) *DescribeLorneTasksMetricsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLorneTasksMetricsResponseBody) SetRequestId(v string) *DescribeLorneTasksMetricsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeLorneTasksMetricsResponseBodyData struct {
	Columns []*string `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	TaskId  *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Values  []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeLorneTasksMetricsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLorneTasksMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLorneTasksMetricsResponseBodyData) SetColumns(v []*string) *DescribeLorneTasksMetricsResponseBodyData {
	s.Columns = v
	return s
}

func (s *DescribeLorneTasksMetricsResponseBodyData) SetName(v string) *DescribeLorneTasksMetricsResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeLorneTasksMetricsResponseBodyData) SetTaskId(v string) *DescribeLorneTasksMetricsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeLorneTasksMetricsResponseBodyData) SetValues(v []*string) *DescribeLorneTasksMetricsResponseBodyData {
	s.Values = v
	return s
}

type DescribeLorneTasksMetricsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLorneTasksMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLorneTasksMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLorneTasksMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLorneTasksMetricsResponse) SetHeaders(v map[string]*string) *DescribeLorneTasksMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLorneTasksMetricsResponse) SetBody(v *DescribeLorneTasksMetricsResponseBody) *DescribeLorneTasksMetricsResponse {
	s.Body = v
	return s
}

type DescribeOSSStorageRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeOSSStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOSSStorageRequest) GoString() string {
	return s.String()
}

func (s *DescribeOSSStorageRequest) SetDBClusterId(v string) *DescribeOSSStorageRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeOSSStorageRequest) SetOwnerAccount(v string) *DescribeOSSStorageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeOSSStorageRequest) SetOwnerId(v int64) *DescribeOSSStorageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeOSSStorageRequest) SetRegionId(v string) *DescribeOSSStorageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOSSStorageRequest) SetResourceOwnerAccount(v string) *DescribeOSSStorageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeOSSStorageRequest) SetResourceOwnerId(v int64) *DescribeOSSStorageRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeOSSStorageResponseBody struct {
	ColdStorage  *bool   `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	Policy       *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	StorageUsage *string `json:"StorageUsage,omitempty" xml:"StorageUsage,omitempty"`
}

func (s DescribeOSSStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOSSStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOSSStorageResponseBody) SetColdStorage(v bool) *DescribeOSSStorageResponseBody {
	s.ColdStorage = &v
	return s
}

func (s *DescribeOSSStorageResponseBody) SetPolicy(v string) *DescribeOSSStorageResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeOSSStorageResponseBody) SetRequestId(v string) *DescribeOSSStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOSSStorageResponseBody) SetState(v string) *DescribeOSSStorageResponseBody {
	s.State = &v
	return s
}

func (s *DescribeOSSStorageResponseBody) SetStorageUsage(v string) *DescribeOSSStorageResponseBody {
	s.StorageUsage = &v
	return s
}

type DescribeOSSStorageResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOSSStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOSSStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOSSStorageResponse) GoString() string {
	return s.String()
}

func (s *DescribeOSSStorageResponse) SetHeaders(v map[string]*string) *DescribeOSSStorageResponse {
	s.Headers = v
	return s
}

func (s *DescribeOSSStorageResponse) SetBody(v *DescribeOSSStorageResponseBody) *DescribeOSSStorageResponse {
	s.Body = v
	return s
}

type DescribeProcessListRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	InitialQueryId       *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	InitialUser          *string `json:"InitialUser,omitempty" xml:"InitialUser,omitempty"`
	Keyword              *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Order                *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryDurationMs      *int32  `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *DescribeProcessListRequest) SetInitialQueryId(v string) *DescribeProcessListRequest {
	s.InitialQueryId = &v
	return s
}

func (s *DescribeProcessListRequest) SetInitialUser(v string) *DescribeProcessListRequest {
	s.InitialUser = &v
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

func (s *DescribeProcessListRequest) SetQueryDurationMs(v int32) *DescribeProcessListRequest {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeProcessListRequest) SetRegionId(v string) *DescribeProcessListRequest {
	s.RegionId = &v
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

type DescribeProcessListResponseBody struct {
	ProcessList *DescribeProcessListResponseBodyProcessList `json:"ProcessList,omitempty" xml:"ProcessList,omitempty" type:"Struct"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProcessListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBody) SetProcessList(v *DescribeProcessListResponseBodyProcessList) *DescribeProcessListResponseBody {
	s.ProcessList = v
	return s
}

func (s *DescribeProcessListResponseBody) SetRequestId(v string) *DescribeProcessListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProcessListResponseBodyProcessList struct {
	Data                   *DescribeProcessListResponseBodyProcessListData        `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Rows                   *string                                                `json:"Rows,omitempty" xml:"Rows,omitempty"`
	RowsBeforeLimitAtLeast *string                                                `json:"RowsBeforeLimitAtLeast,omitempty" xml:"RowsBeforeLimitAtLeast,omitempty"`
	Statistics             *DescribeProcessListResponseBodyProcessListStatistics  `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
	TableSchema            *DescribeProcessListResponseBodyProcessListTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s DescribeProcessListResponseBodyProcessList) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBodyProcessList) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyProcessList) SetData(v *DescribeProcessListResponseBodyProcessListData) *DescribeProcessListResponseBodyProcessList {
	s.Data = v
	return s
}

func (s *DescribeProcessListResponseBodyProcessList) SetRows(v string) *DescribeProcessListResponseBodyProcessList {
	s.Rows = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessList) SetRowsBeforeLimitAtLeast(v string) *DescribeProcessListResponseBodyProcessList {
	s.RowsBeforeLimitAtLeast = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessList) SetStatistics(v *DescribeProcessListResponseBodyProcessListStatistics) *DescribeProcessListResponseBodyProcessList {
	s.Statistics = v
	return s
}

func (s *DescribeProcessListResponseBodyProcessList) SetTableSchema(v *DescribeProcessListResponseBodyProcessListTableSchema) *DescribeProcessListResponseBodyProcessList {
	s.TableSchema = v
	return s
}

type DescribeProcessListResponseBodyProcessListData struct {
	ResultSet []*DescribeProcessListResponseBodyProcessListDataResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
}

func (s DescribeProcessListResponseBodyProcessListData) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBodyProcessListData) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyProcessListData) SetResultSet(v []*DescribeProcessListResponseBodyProcessListDataResultSet) *DescribeProcessListResponseBodyProcessListData {
	s.ResultSet = v
	return s
}

type DescribeProcessListResponseBodyProcessListDataResultSet struct {
	InitialAddress  *string `json:"InitialAddress,omitempty" xml:"InitialAddress,omitempty"`
	InitialQueryId  *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	InitialUser     *string `json:"InitialUser,omitempty" xml:"InitialUser,omitempty"`
	Query           *string `json:"Query,omitempty" xml:"Query,omitempty"`
	QueryDurationMs *string `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	QueryStartTime  *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
}

func (s DescribeProcessListResponseBodyProcessListDataResultSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBodyProcessListDataResultSet) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) SetInitialAddress(v string) *DescribeProcessListResponseBodyProcessListDataResultSet {
	s.InitialAddress = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) SetInitialQueryId(v string) *DescribeProcessListResponseBodyProcessListDataResultSet {
	s.InitialQueryId = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) SetInitialUser(v string) *DescribeProcessListResponseBodyProcessListDataResultSet {
	s.InitialUser = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) SetQuery(v string) *DescribeProcessListResponseBodyProcessListDataResultSet {
	s.Query = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) SetQueryDurationMs(v string) *DescribeProcessListResponseBodyProcessListDataResultSet {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListDataResultSet) SetQueryStartTime(v string) *DescribeProcessListResponseBodyProcessListDataResultSet {
	s.QueryStartTime = &v
	return s
}

type DescribeProcessListResponseBodyProcessListStatistics struct {
	BytesRead   *int32   `json:"BytesRead,omitempty" xml:"BytesRead,omitempty"`
	ElapsedTime *float32 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	RowsRead    *int32   `json:"RowsRead,omitempty" xml:"RowsRead,omitempty"`
}

func (s DescribeProcessListResponseBodyProcessListStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBodyProcessListStatistics) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyProcessListStatistics) SetBytesRead(v int32) *DescribeProcessListResponseBodyProcessListStatistics {
	s.BytesRead = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListStatistics) SetElapsedTime(v float32) *DescribeProcessListResponseBodyProcessListStatistics {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListStatistics) SetRowsRead(v int32) *DescribeProcessListResponseBodyProcessListStatistics {
	s.RowsRead = &v
	return s
}

type DescribeProcessListResponseBodyProcessListTableSchema struct {
	ResultSet []*DescribeProcessListResponseBodyProcessListTableSchemaResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
}

func (s DescribeProcessListResponseBodyProcessListTableSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBodyProcessListTableSchema) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyProcessListTableSchema) SetResultSet(v []*DescribeProcessListResponseBodyProcessListTableSchemaResultSet) *DescribeProcessListResponseBodyProcessListTableSchema {
	s.ResultSet = v
	return s
}

type DescribeProcessListResponseBodyProcessListTableSchemaResultSet struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeProcessListResponseBodyProcessListTableSchemaResultSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessListResponseBodyProcessListTableSchemaResultSet) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyProcessListTableSchemaResultSet) SetName(v string) *DescribeProcessListResponseBodyProcessListTableSchemaResultSet {
	s.Name = &v
	return s
}

func (s *DescribeProcessListResponseBodyProcessListTableSchemaResultSet) SetType(v string) *DescribeProcessListResponseBodyProcessListTableSchemaResultSet {
	s.Type = &v
	return s
}

type DescribeProcessListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeProcessListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeProcessListResponse) SetBody(v *DescribeProcessListResponseBody) *DescribeProcessListResponse {
	s.Body = v
	return s
}

type DescribeRDSTablesRequest struct {
	DbClusterId          *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RdsId                *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	RdsPassword          *string `json:"RdsPassword,omitempty" xml:"RdsPassword,omitempty"`
	RdsPort              *int64  `json:"RdsPort,omitempty" xml:"RdsPort,omitempty"`
	RdsUserName          *string `json:"RdsUserName,omitempty" xml:"RdsUserName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Schema               *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s DescribeRDSTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRDSTablesRequest) SetDbClusterId(v string) *DescribeRDSTablesRequest {
	s.DbClusterId = &v
	return s
}

func (s *DescribeRDSTablesRequest) SetOwnerAccount(v string) *DescribeRDSTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRDSTablesRequest) SetOwnerId(v int64) *DescribeRDSTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRDSTablesRequest) SetRdsId(v string) *DescribeRDSTablesRequest {
	s.RdsId = &v
	return s
}

func (s *DescribeRDSTablesRequest) SetRdsPassword(v string) *DescribeRDSTablesRequest {
	s.RdsPassword = &v
	return s
}

func (s *DescribeRDSTablesRequest) SetRdsPort(v int64) *DescribeRDSTablesRequest {
	s.RdsPort = &v
	return s
}

func (s *DescribeRDSTablesRequest) SetRdsUserName(v string) *DescribeRDSTablesRequest {
	s.RdsUserName = &v
	return s
}

func (s *DescribeRDSTablesRequest) SetResourceOwnerAccount(v string) *DescribeRDSTablesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRDSTablesRequest) SetResourceOwnerId(v int64) *DescribeRDSTablesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRDSTablesRequest) SetSchema(v string) *DescribeRDSTablesRequest {
	s.Schema = &v
	return s
}

type DescribeRDSTablesResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tables    []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s DescribeRDSTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRDSTablesResponseBody) SetRequestId(v string) *DescribeRDSTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRDSTablesResponseBody) SetTables(v []*string) *DescribeRDSTablesResponseBody {
	s.Tables = v
	return s
}

type DescribeRDSTablesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRDSTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRDSTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRDSTablesResponse) SetHeaders(v map[string]*string) *DescribeRDSTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRDSTablesResponse) SetBody(v *DescribeRDSTablesResponseBody) *DescribeRDSTablesResponse {
	s.Body = v
	return s
}

type DescribeRDSVpcRequest struct {
	DbClusterId          *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RdsId                *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRDSVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSVpcRequest) GoString() string {
	return s.String()
}

func (s *DescribeRDSVpcRequest) SetDbClusterId(v string) *DescribeRDSVpcRequest {
	s.DbClusterId = &v
	return s
}

func (s *DescribeRDSVpcRequest) SetOwnerAccount(v string) *DescribeRDSVpcRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRDSVpcRequest) SetOwnerId(v int64) *DescribeRDSVpcRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRDSVpcRequest) SetRdsId(v string) *DescribeRDSVpcRequest {
	s.RdsId = &v
	return s
}

func (s *DescribeRDSVpcRequest) SetRegionId(v string) *DescribeRDSVpcRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRDSVpcRequest) SetResourceOwnerAccount(v string) *DescribeRDSVpcRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRDSVpcRequest) SetResourceOwnerId(v int64) *DescribeRDSVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeRDSVpcResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRDSVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSVpcResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRDSVpcResponseBody) SetRequestId(v string) *DescribeRDSVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRDSVpcResponseBody) SetVpcId(v string) *DescribeRDSVpcResponseBody {
	s.VpcId = &v
	return s
}

type DescribeRDSVpcResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRDSVpcResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRDSVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSVpcResponse) GoString() string {
	return s.String()
}

func (s *DescribeRDSVpcResponse) SetHeaders(v map[string]*string) *DescribeRDSVpcResponse {
	s.Headers = v
	return s
}

func (s *DescribeRDSVpcResponse) SetBody(v *DescribeRDSVpcResponseBody) *DescribeRDSVpcResponse {
	s.Body = v
	return s
}

type DescribeRDSschemasRequest struct {
	DbClusterId          *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RdsId                *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	RdsPassword          *string `json:"RdsPassword,omitempty" xml:"RdsPassword,omitempty"`
	RdsPort              *int64  `json:"RdsPort,omitempty" xml:"RdsPort,omitempty"`
	RdsUserName          *string `json:"RdsUserName,omitempty" xml:"RdsUserName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRDSschemasRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSschemasRequest) GoString() string {
	return s.String()
}

func (s *DescribeRDSschemasRequest) SetDbClusterId(v string) *DescribeRDSschemasRequest {
	s.DbClusterId = &v
	return s
}

func (s *DescribeRDSschemasRequest) SetOwnerAccount(v string) *DescribeRDSschemasRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRDSschemasRequest) SetOwnerId(v int64) *DescribeRDSschemasRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRDSschemasRequest) SetRdsId(v string) *DescribeRDSschemasRequest {
	s.RdsId = &v
	return s
}

func (s *DescribeRDSschemasRequest) SetRdsPassword(v string) *DescribeRDSschemasRequest {
	s.RdsPassword = &v
	return s
}

func (s *DescribeRDSschemasRequest) SetRdsPort(v int64) *DescribeRDSschemasRequest {
	s.RdsPort = &v
	return s
}

func (s *DescribeRDSschemasRequest) SetRdsUserName(v string) *DescribeRDSschemasRequest {
	s.RdsUserName = &v
	return s
}

func (s *DescribeRDSschemasRequest) SetResourceOwnerAccount(v string) *DescribeRDSschemasRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRDSschemasRequest) SetResourceOwnerId(v int64) *DescribeRDSschemasRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeRDSschemasResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Schemas   []*string `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Repeated"`
}

func (s DescribeRDSschemasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSschemasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRDSschemasResponseBody) SetRequestId(v string) *DescribeRDSschemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRDSschemasResponseBody) SetSchemas(v []*string) *DescribeRDSschemasResponseBody {
	s.Schemas = v
	return s
}

type DescribeRDSschemasResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRDSschemasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRDSschemasResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRDSschemasResponse) GoString() string {
	return s.String()
}

func (s *DescribeRDSschemasResponse) SetHeaders(v map[string]*string) *DescribeRDSschemasResponse {
	s.Headers = v
	return s
}

func (s *DescribeRDSschemasResponse) SetBody(v *DescribeRDSschemasResponseBody) *DescribeRDSschemasResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
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
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	RegionId *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Zones    *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
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
	VpcEnabled *bool   `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) GoString() string {
	return s.String()
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeSchemasRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemaName           *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName            *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
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

func (s *DescribeSchemasRequest) SetSchemaName(v string) *DescribeSchemasRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeSchemasRequest) SetTableName(v string) *DescribeSchemasRequest {
	s.TableName = &v
	return s
}

type DescribeSchemasResponseBody struct {
	Items     *DescribeSchemasResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName  *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSchemasResponse) SetBody(v *DescribeSchemasResponseBody) *DescribeSchemasResponse {
	s.Body = v
	return s
}

type DescribeSlowLogRecordsRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryDurationMs      *int32  `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s *DescribeSlowLogRecordsRequest) SetEndTime(v string) *DescribeSlowLogRecordsRequest {
	s.EndTime = &v
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

func (s *DescribeSlowLogRecordsRequest) SetQueryDurationMs(v int32) *DescribeSlowLogRecordsRequest {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetRegionId(v string) *DescribeSlowLogRecordsRequest {
	s.RegionId = &v
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

type DescribeSlowLogRecordsResponseBody struct {
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlowLogRecords *DescribeSlowLogRecordsResponseBodySlowLogRecords `json:"SlowLogRecords,omitempty" xml:"SlowLogRecords,omitempty" type:"Struct"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetSlowLogRecords(v *DescribeSlowLogRecordsResponseBodySlowLogRecords) *DescribeSlowLogRecordsResponseBody {
	s.SlowLogRecords = v
	return s
}

type DescribeSlowLogRecordsResponseBodySlowLogRecords struct {
	Data                   *DescribeSlowLogRecordsResponseBodySlowLogRecordsData        `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Rows                   *string                                                      `json:"Rows,omitempty" xml:"Rows,omitempty"`
	RowsBeforeLimitAtLeast *string                                                      `json:"RowsBeforeLimitAtLeast,omitempty" xml:"RowsBeforeLimitAtLeast,omitempty"`
	Statistics             *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics  `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
	TableSchema            *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecords) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) SetData(v *DescribeSlowLogRecordsResponseBodySlowLogRecordsData) *DescribeSlowLogRecordsResponseBodySlowLogRecords {
	s.Data = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) SetRows(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecords {
	s.Rows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) SetRowsBeforeLimitAtLeast(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecords {
	s.RowsBeforeLimitAtLeast = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) SetStatistics(v *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) *DescribeSlowLogRecordsResponseBodySlowLogRecords {
	s.Statistics = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecords) SetTableSchema(v *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema) *DescribeSlowLogRecordsResponseBodySlowLogRecords {
	s.TableSchema = v
	return s
}

type DescribeSlowLogRecordsResponseBodySlowLogRecordsData struct {
	ResultSet []*DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsData) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsData) SetResultSet(v []*DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) *DescribeSlowLogRecordsResponseBodySlowLogRecordsData {
	s.ResultSet = v
	return s
}

type DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet struct {
	InitialAddress  *string `json:"InitialAddress,omitempty" xml:"InitialAddress,omitempty"`
	InitialQueryId  *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	InitialUser     *string `json:"InitialUser,omitempty" xml:"InitialUser,omitempty"`
	MemoryUsage     *string `json:"MemoryUsage,omitempty" xml:"MemoryUsage,omitempty"`
	Query           *string `json:"Query,omitempty" xml:"Query,omitempty"`
	QueryDurationMs *string `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	QueryStartTime  *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	ReadBytes       *string `json:"ReadBytes,omitempty" xml:"ReadBytes,omitempty"`
	ReadRows        *string `json:"ReadRows,omitempty" xml:"ReadRows,omitempty"`
	ResultBytes     *string `json:"ResultBytes,omitempty" xml:"ResultBytes,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetInitialAddress(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.InitialAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetInitialQueryId(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.InitialQueryId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetInitialUser(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.InitialUser = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetMemoryUsage(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.MemoryUsage = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetQuery(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.Query = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetQueryDurationMs(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetQueryStartTime(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.QueryStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetReadBytes(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.ReadBytes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetReadRows(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.ReadRows = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetResultBytes(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.ResultBytes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet) SetType(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsDataResultSet {
	s.Type = &v
	return s
}

type DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics struct {
	BytesRead   *int32   `json:"BytesRead,omitempty" xml:"BytesRead,omitempty"`
	ElapsedTime *float32 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	RowsRead    *int32   `json:"RowsRead,omitempty" xml:"RowsRead,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) SetBytesRead(v int32) *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics {
	s.BytesRead = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) SetElapsedTime(v float32) *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics) SetRowsRead(v int32) *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics {
	s.RowsRead = &v
	return s
}

type DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema struct {
	ResultSet []*DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema) SetResultSet(v []*DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet) *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema {
	s.ResultSet = v
	return s
}

type DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet) SetName(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet {
	s.Name = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet) SetType(v string) *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchemaResultSet {
	s.Type = &v
	return s
}

type DescribeSlowLogRecordsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlowLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSlowLogRecordsResponse) SetBody(v *DescribeSlowLogRecordsResponseBody) *DescribeSlowLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeSlowLogTrendRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	QueryDurationMs      *int32  `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s *DescribeSlowLogTrendRequest) SetQueryDurationMs(v int32) *DescribeSlowLogTrendRequest {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetRegionId(v string) *DescribeSlowLogTrendRequest {
	s.RegionId = &v
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
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlowLogTrend *DescribeSlowLogTrendResponseBodySlowLogTrend `json:"SlowLogTrend,omitempty" xml:"SlowLogTrend,omitempty" type:"Struct"`
}

func (s DescribeSlowLogTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBody) SetRequestId(v string) *DescribeSlowLogTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBody) SetSlowLogTrend(v *DescribeSlowLogTrendResponseBodySlowLogTrend) *DescribeSlowLogTrendResponseBody {
	s.SlowLogTrend = v
	return s
}

type DescribeSlowLogTrendResponseBodySlowLogTrend struct {
	Data                   *DescribeSlowLogTrendResponseBodySlowLogTrendData        `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Rows                   *string                                                  `json:"Rows,omitempty" xml:"Rows,omitempty"`
	RowsBeforeLimitAtLeast *string                                                  `json:"RowsBeforeLimitAtLeast,omitempty" xml:"RowsBeforeLimitAtLeast,omitempty"`
	Statistics             *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics  `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
	TableSchema            *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrend) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrend) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) SetData(v *DescribeSlowLogTrendResponseBodySlowLogTrendData) *DescribeSlowLogTrendResponseBodySlowLogTrend {
	s.Data = v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) SetRows(v string) *DescribeSlowLogTrendResponseBodySlowLogTrend {
	s.Rows = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) SetRowsBeforeLimitAtLeast(v string) *DescribeSlowLogTrendResponseBodySlowLogTrend {
	s.RowsBeforeLimitAtLeast = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) SetStatistics(v *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) *DescribeSlowLogTrendResponseBodySlowLogTrend {
	s.Statistics = v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrend) SetTableSchema(v *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema) *DescribeSlowLogTrendResponseBodySlowLogTrend {
	s.TableSchema = v
	return s
}

type DescribeSlowLogTrendResponseBodySlowLogTrendData struct {
	ResultSet []*DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendData) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendData) SetResultSet(v []*DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) *DescribeSlowLogTrendResponseBodySlowLogTrendData {
	s.ResultSet = v
	return s
}

type DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet struct {
	AvgQueryDurationMs *string `json:"AvgQueryDurationMs,omitempty" xml:"AvgQueryDurationMs,omitempty"`
	Count              *string `json:"Count,omitempty" xml:"Count,omitempty"`
	MaxQueryDurationMs *string `json:"MaxQueryDurationMs,omitempty" xml:"MaxQueryDurationMs,omitempty"`
	MinQueryDurationMs *string `json:"MinQueryDurationMs,omitempty" xml:"MinQueryDurationMs,omitempty"`
	QueryStartTime     *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) SetAvgQueryDurationMs(v string) *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet {
	s.AvgQueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) SetCount(v string) *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet {
	s.Count = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) SetMaxQueryDurationMs(v string) *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet {
	s.MaxQueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) SetMinQueryDurationMs(v string) *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet {
	s.MinQueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet) SetQueryStartTime(v string) *DescribeSlowLogTrendResponseBodySlowLogTrendDataResultSet {
	s.QueryStartTime = &v
	return s
}

type DescribeSlowLogTrendResponseBodySlowLogTrendStatistics struct {
	BytesRead   *int32   `json:"BytesRead,omitempty" xml:"BytesRead,omitempty"`
	ElapsedTime *float32 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	RowsRead    *int32   `json:"RowsRead,omitempty" xml:"RowsRead,omitempty"`
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) SetBytesRead(v int32) *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics {
	s.BytesRead = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) SetElapsedTime(v float32) *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics) SetRowsRead(v int32) *DescribeSlowLogTrendResponseBodySlowLogTrendStatistics {
	s.RowsRead = &v
	return s
}

type DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema struct {
	ResultSet []*DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema) SetResultSet(v []*DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet) *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchema {
	s.ResultSet = v
	return s
}

type DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet) SetName(v string) *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet {
	s.Name = &v
	return s
}

func (s *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet) SetType(v string) *DescribeSlowLogTrendResponseBodySlowLogTrendTableSchemaResultSet {
	s.Type = &v
	return s
}

type DescribeSlowLogTrendResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlowLogTrendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSlowLogTrendResponse) SetBody(v *DescribeSlowLogTrendResponseBody) *DescribeSlowLogTrendResponse {
	s.Body = v
	return s
}

type DescribeSynDbTablesRequest struct {
	DbClusterId          *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SynDb                *string `json:"SynDb,omitempty" xml:"SynDb,omitempty"`
}

func (s DescribeSynDbTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynDbTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynDbTablesRequest) SetDbClusterId(v string) *DescribeSynDbTablesRequest {
	s.DbClusterId = &v
	return s
}

func (s *DescribeSynDbTablesRequest) SetOwnerAccount(v string) *DescribeSynDbTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSynDbTablesRequest) SetOwnerId(v int64) *DescribeSynDbTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynDbTablesRequest) SetResourceOwnerAccount(v string) *DescribeSynDbTablesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSynDbTablesRequest) SetResourceOwnerId(v int64) *DescribeSynDbTablesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSynDbTablesRequest) SetSynDb(v string) *DescribeSynDbTablesRequest {
	s.SynDb = &v
	return s
}

type DescribeSynDbTablesResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tables    []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s DescribeSynDbTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynDbTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynDbTablesResponseBody) SetRequestId(v string) *DescribeSynDbTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynDbTablesResponseBody) SetTables(v []*string) *DescribeSynDbTablesResponseBody {
	s.Tables = v
	return s
}

type DescribeSynDbTablesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSynDbTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSynDbTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynDbTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynDbTablesResponse) SetHeaders(v map[string]*string) *DescribeSynDbTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynDbTablesResponse) SetBody(v *DescribeSynDbTablesResponseBody) *DescribeSynDbTablesResponse {
	s.Body = v
	return s
}

type DescribeSynDbsRequest struct {
	DbClusterId          *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSynDbsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynDbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynDbsRequest) SetDbClusterId(v string) *DescribeSynDbsRequest {
	s.DbClusterId = &v
	return s
}

func (s *DescribeSynDbsRequest) SetOwnerAccount(v string) *DescribeSynDbsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSynDbsRequest) SetOwnerId(v int64) *DescribeSynDbsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynDbsRequest) SetResourceOwnerAccount(v string) *DescribeSynDbsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSynDbsRequest) SetResourceOwnerId(v int64) *DescribeSynDbsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeSynDbsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SynDbs    []*DescribeSynDbsResponseBodySynDbs `json:"SynDbs,omitempty" xml:"SynDbs,omitempty" type:"Repeated"`
}

func (s DescribeSynDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynDbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynDbsResponseBody) SetRequestId(v string) *DescribeSynDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynDbsResponseBody) SetSynDbs(v []*DescribeSynDbsResponseBodySynDbs) *DescribeSynDbsResponseBody {
	s.SynDbs = v
	return s
}

type DescribeSynDbsResponseBodySynDbs struct {
	ErrorMsg    *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RdsId       *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	RdsPassword *string `json:"RdsPassword,omitempty" xml:"RdsPassword,omitempty"`
	RdsUserName *string `json:"RdsUserName,omitempty" xml:"RdsUserName,omitempty"`
	SynDb       *string `json:"SynDb,omitempty" xml:"SynDb,omitempty"`
	SynStatus   *bool   `json:"SynStatus,omitempty" xml:"SynStatus,omitempty"`
}

func (s DescribeSynDbsResponseBodySynDbs) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynDbsResponseBodySynDbs) GoString() string {
	return s.String()
}

func (s *DescribeSynDbsResponseBodySynDbs) SetErrorMsg(v string) *DescribeSynDbsResponseBodySynDbs {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeSynDbsResponseBodySynDbs) SetRdsId(v string) *DescribeSynDbsResponseBodySynDbs {
	s.RdsId = &v
	return s
}

func (s *DescribeSynDbsResponseBodySynDbs) SetRdsPassword(v string) *DescribeSynDbsResponseBodySynDbs {
	s.RdsPassword = &v
	return s
}

func (s *DescribeSynDbsResponseBodySynDbs) SetRdsUserName(v string) *DescribeSynDbsResponseBodySynDbs {
	s.RdsUserName = &v
	return s
}

func (s *DescribeSynDbsResponseBodySynDbs) SetSynDb(v string) *DescribeSynDbsResponseBodySynDbs {
	s.SynDb = &v
	return s
}

func (s *DescribeSynDbsResponseBodySynDbs) SetSynStatus(v bool) *DescribeSynDbsResponseBodySynDbs {
	s.SynStatus = &v
	return s
}

type DescribeSynDbsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSynDbsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSynDbsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynDbsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynDbsResponse) SetHeaders(v map[string]*string) *DescribeSynDbsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynDbsResponse) SetBody(v *DescribeSynDbsResponseBody) *DescribeSynDbsResponse {
	s.Body = v
	return s
}

type DescribeTablesRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemaName           *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName            *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
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

func (s *DescribeTablesRequest) SetTableName(v string) *DescribeTablesRequest {
	s.TableName = &v
	return s
}

type DescribeTablesResponseBody struct {
	Items     *DescribeTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SchemaName  *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName   *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeTablesResponse) SetBody(v *DescribeTablesResponseBody) *DescribeTablesResponse {
	s.Body = v
	return s
}

type DescribeTransferHistoryRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeTransferHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransferHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeTransferHistoryRequest) SetDBClusterId(v string) *DescribeTransferHistoryRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTransferHistoryRequest) SetOwnerAccount(v string) *DescribeTransferHistoryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTransferHistoryRequest) SetOwnerId(v int64) *DescribeTransferHistoryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTransferHistoryRequest) SetResourceOwnerAccount(v string) *DescribeTransferHistoryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTransferHistoryRequest) SetResourceOwnerId(v int64) *DescribeTransferHistoryRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeTransferHistoryResponseBody struct {
	HistoryDetails *DescribeTransferHistoryResponseBodyHistoryDetails `json:"HistoryDetails,omitempty" xml:"HistoryDetails,omitempty" type:"Struct"`
	RequestId      *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTransferHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransferHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTransferHistoryResponseBody) SetHistoryDetails(v *DescribeTransferHistoryResponseBodyHistoryDetails) *DescribeTransferHistoryResponseBody {
	s.HistoryDetails = v
	return s
}

func (s *DescribeTransferHistoryResponseBody) SetRequestId(v string) *DescribeTransferHistoryResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTransferHistoryResponseBodyHistoryDetails struct {
	HistoryDetail []*DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail `json:"HistoryDetail,omitempty" xml:"HistoryDetail,omitempty" type:"Repeated"`
}

func (s DescribeTransferHistoryResponseBodyHistoryDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransferHistoryResponseBodyHistoryDetails) GoString() string {
	return s.String()
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetails) SetHistoryDetail(v []*DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) *DescribeTransferHistoryResponseBodyHistoryDetails {
	s.HistoryDetail = v
	return s
}

type DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail struct {
	Progress        *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	SourceDBCluster *string `json:"SourceDBCluster,omitempty" xml:"SourceDBCluster,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetDBCluster *string `json:"TargetDBCluster,omitempty" xml:"TargetDBCluster,omitempty"`
}

func (s DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) GoString() string {
	return s.String()
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetProgress(v string) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.Progress = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetSourceDBCluster(v string) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.SourceDBCluster = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetStatus(v string) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.Status = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetTargetDBCluster(v string) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.TargetDBCluster = &v
	return s
}

type DescribeTransferHistoryResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTransferHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTransferHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransferHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeTransferHistoryResponse) SetHeaders(v map[string]*string) *DescribeTransferHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeTransferHistoryResponse) SetBody(v *DescribeTransferHistoryResponseBody) *DescribeTransferHistoryResponse {
	s.Body = v
	return s
}

type KillProcessRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	InitialQueryId       *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *KillProcessRequest) SetInitialQueryId(v string) *KillProcessRequest {
	s.InitialQueryId = &v
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

func (s *KillProcessRequest) SetRegionId(v string) *KillProcessRequest {
	s.RegionId = &v
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *KillProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *KillProcessResponse) SetBody(v *KillProcessResponseBody) *KillProcessResponse {
	s.Body = v
	return s
}

type ModifyAccountAuthorityRequest struct {
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AllowDatabases       *string `json:"AllowDatabases,omitempty" xml:"AllowDatabases,omitempty"`
	AllowDictionaries    *string `json:"AllowDictionaries,omitempty" xml:"AllowDictionaries,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DdlAuthority         *bool   `json:"DdlAuthority,omitempty" xml:"DdlAuthority,omitempty"`
	DmlAuthority         *string `json:"DmlAuthority,omitempty" xml:"DmlAuthority,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TotalDatabases       *string `json:"TotalDatabases,omitempty" xml:"TotalDatabases,omitempty"`
	TotalDictionaries    *string `json:"TotalDictionaries,omitempty" xml:"TotalDictionaries,omitempty"`
}

func (s ModifyAccountAuthorityRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountAuthorityRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityRequest) SetAccountName(v string) *ModifyAccountAuthorityRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetAllowDatabases(v string) *ModifyAccountAuthorityRequest {
	s.AllowDatabases = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetAllowDictionaries(v string) *ModifyAccountAuthorityRequest {
	s.AllowDictionaries = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetDBClusterId(v string) *ModifyAccountAuthorityRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetDdlAuthority(v bool) *ModifyAccountAuthorityRequest {
	s.DdlAuthority = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetDmlAuthority(v string) *ModifyAccountAuthorityRequest {
	s.DmlAuthority = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetOwnerAccount(v string) *ModifyAccountAuthorityRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetOwnerId(v int64) *ModifyAccountAuthorityRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetRegionId(v string) *ModifyAccountAuthorityRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetResourceOwnerAccount(v string) *ModifyAccountAuthorityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetResourceOwnerId(v int64) *ModifyAccountAuthorityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetTotalDatabases(v string) *ModifyAccountAuthorityRequest {
	s.TotalDatabases = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetTotalDictionaries(v string) *ModifyAccountAuthorityRequest {
	s.TotalDictionaries = &v
	return s
}

type ModifyAccountAuthorityResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountAuthorityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityResponseBody) SetRequestId(v string) *ModifyAccountAuthorityResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccountAuthorityResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAccountAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccountAuthorityResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountAuthorityResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityResponse) SetHeaders(v map[string]*string) *ModifyAccountAuthorityResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountAuthorityResponse) SetBody(v *ModifyAccountAuthorityResponseBody) *ModifyAccountAuthorityResponse {
	s.Body = v
	return s
}

type ModifyAccountDescriptionRequest struct {
	AccountDescription   *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAccountDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyAccountDescriptionResponse) SetBody(v *ModifyAccountDescriptionResponseBody) *ModifyAccountDescriptionResponse {
	s.Body = v
	return s
}

type ModifyBackupPolicyRequest struct {
	BackupRetentionPeriod *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	DBClusterId           *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupTime   *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyBackupPolicyResponse) SetBody(v *ModifyBackupPolicyResponseBody) *ModifyBackupPolicyResponse {
	s.Body = v
	return s
}

type ModifyDBClusterRequest struct {
	DBClusterClass       *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBNodeGroupCount     *string `json:"DBNodeGroupCount,omitempty" xml:"DBNodeGroupCount,omitempty"`
	DBNodeStorage        *string `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterRequest) SetDBClusterClass(v string) *ModifyDBClusterRequest {
	s.DBClusterClass = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterId(v string) *ModifyDBClusterRequest {
	s.DBClusterId = &v
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

type ModifyDBClusterResponseBody struct {
	DBCluster *ModifyDBClusterResponseBodyDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponseBody) SetDBCluster(v *ModifyDBClusterResponseBodyDBCluster) *ModifyDBClusterResponseBody {
	s.DBCluster = v
	return s
}

func (s *ModifyDBClusterResponseBody) SetRequestId(v string) *ModifyDBClusterResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterResponseBodyDBCluster struct {
	DbClusterId *string `json:"dbClusterId,omitempty" xml:"dbClusterId,omitempty"`
	OrderId     *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ModifyDBClusterResponseBodyDBCluster) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterResponseBodyDBCluster) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResponseBodyDBCluster) SetDbClusterId(v string) *ModifyDBClusterResponseBodyDBCluster {
	s.DbClusterId = &v
	return s
}

func (s *ModifyDBClusterResponseBodyDBCluster) SetOrderId(v string) *ModifyDBClusterResponseBodyDBCluster {
	s.OrderId = &v
	return s
}

type ModifyDBClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDBClusterResponse) SetBody(v *ModifyDBClusterResponseBody) *ModifyDBClusterResponse {
	s.Body = v
	return s
}

type ModifyDBClusterAccessWhiteListRequest struct {
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	DBClusterIPArrayName      *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	DBClusterId               *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ModifyMode                *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityIps               *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterAccessWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) SetRequestId(v string) *ModifyDBClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterAccessWhiteListResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDBClusterAccessWhiteListResponse) SetBody(v *ModifyDBClusterAccessWhiteListResponseBody) *ModifyDBClusterAccessWhiteListResponse {
	s.Body = v
	return s
}

type ModifyDBClusterConfigRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserConfig           *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
}

func (s ModifyDBClusterConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigRequest) SetDBClusterId(v string) *ModifyDBClusterConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetOwnerAccount(v string) *ModifyDBClusterConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetOwnerId(v int64) *ModifyDBClusterConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetRegionId(v string) *ModifyDBClusterConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetResourceOwnerId(v int64) *ModifyDBClusterConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetUserConfig(v string) *ModifyDBClusterConfigRequest {
	s.UserConfig = &v
	return s
}

type ModifyDBClusterConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigResponseBody) SetRequestId(v string) *ModifyDBClusterConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterConfigResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigResponse) SetHeaders(v map[string]*string) *ModifyDBClusterConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterConfigResponse) SetBody(v *ModifyDBClusterConfigResponseBody) *ModifyDBClusterConfigResponse {
	s.Body = v
	return s
}

type ModifyDBClusterDescriptionRequest struct {
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDBClusterDescriptionResponse) SetBody(v *ModifyDBClusterDescriptionResponseBody) *ModifyDBClusterDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDBClusterMaintainTimeRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDBClusterMaintainTimeResponse) SetBody(v *ModifyDBClusterMaintainTimeResponseBody) *ModifyDBClusterMaintainTimeResponse {
	s.Body = v
	return s
}

type ModifyDBConfigRequest struct {
	Config               *string `json:"Config,omitempty" xml:"Config,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBConfigRequest) SetConfig(v string) *ModifyDBConfigRequest {
	s.Config = &v
	return s
}

func (s *ModifyDBConfigRequest) SetDBClusterId(v string) *ModifyDBConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBConfigRequest) SetOwnerAccount(v string) *ModifyDBConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBConfigRequest) SetOwnerId(v int64) *ModifyDBConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBConfigRequest) SetRegionId(v string) *ModifyDBConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBConfigRequest) SetResourceOwnerAccount(v string) *ModifyDBConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBConfigRequest) SetResourceOwnerId(v int64) *ModifyDBConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyDBConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBConfigResponseBody) SetRequestId(v string) *ModifyDBConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBConfigResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBConfigResponse) SetHeaders(v map[string]*string) *ModifyDBConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBConfigResponse) SetBody(v *ModifyDBConfigResponseBody) *ModifyDBConfigResponse {
	s.Body = v
	return s
}

type ModifyRDSToClickhouseDbRequest struct {
	CkPassword           *string `json:"CkPassword,omitempty" xml:"CkPassword,omitempty"`
	CkUserName           *string `json:"CkUserName,omitempty" xml:"CkUserName,omitempty"`
	ClickhousePort       *int64  `json:"ClickhousePort,omitempty" xml:"ClickhousePort,omitempty"`
	DbClusterId          *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	LimitUpper           *int64  `json:"LimitUpper,omitempty" xml:"LimitUpper,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RdsId                *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	RdsPassword          *string `json:"RdsPassword,omitempty" xml:"RdsPassword,omitempty"`
	RdsPort              *int64  `json:"RdsPort,omitempty" xml:"RdsPort,omitempty"`
	RdsSynDb             *string `json:"RdsSynDb,omitempty" xml:"RdsSynDb,omitempty"`
	RdsSynTables         *string `json:"RdsSynTables,omitempty" xml:"RdsSynTables,omitempty"`
	RdsUserName          *string `json:"RdsUserName,omitempty" xml:"RdsUserName,omitempty"`
	RdsVpcId             *string `json:"RdsVpcId,omitempty" xml:"RdsVpcId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SkipUnsupported      *bool   `json:"SkipUnsupported,omitempty" xml:"SkipUnsupported,omitempty"`
}

func (s ModifyRDSToClickhouseDbRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRDSToClickhouseDbRequest) GoString() string {
	return s.String()
}

func (s *ModifyRDSToClickhouseDbRequest) SetCkPassword(v string) *ModifyRDSToClickhouseDbRequest {
	s.CkPassword = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetCkUserName(v string) *ModifyRDSToClickhouseDbRequest {
	s.CkUserName = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetClickhousePort(v int64) *ModifyRDSToClickhouseDbRequest {
	s.ClickhousePort = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetDbClusterId(v string) *ModifyRDSToClickhouseDbRequest {
	s.DbClusterId = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetLimitUpper(v int64) *ModifyRDSToClickhouseDbRequest {
	s.LimitUpper = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetOwnerAccount(v string) *ModifyRDSToClickhouseDbRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetOwnerId(v int64) *ModifyRDSToClickhouseDbRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetRdsId(v string) *ModifyRDSToClickhouseDbRequest {
	s.RdsId = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetRdsPassword(v string) *ModifyRDSToClickhouseDbRequest {
	s.RdsPassword = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetRdsPort(v int64) *ModifyRDSToClickhouseDbRequest {
	s.RdsPort = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetRdsSynDb(v string) *ModifyRDSToClickhouseDbRequest {
	s.RdsSynDb = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetRdsSynTables(v string) *ModifyRDSToClickhouseDbRequest {
	s.RdsSynTables = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetRdsUserName(v string) *ModifyRDSToClickhouseDbRequest {
	s.RdsUserName = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetRdsVpcId(v string) *ModifyRDSToClickhouseDbRequest {
	s.RdsVpcId = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetResourceOwnerAccount(v string) *ModifyRDSToClickhouseDbRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetResourceOwnerId(v int64) *ModifyRDSToClickhouseDbRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyRDSToClickhouseDbRequest) SetSkipUnsupported(v bool) *ModifyRDSToClickhouseDbRequest {
	s.SkipUnsupported = &v
	return s
}

type ModifyRDSToClickhouseDbResponseBody struct {
	ErrorCode *int64  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyRDSToClickhouseDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRDSToClickhouseDbResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRDSToClickhouseDbResponseBody) SetErrorCode(v int64) *ModifyRDSToClickhouseDbResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyRDSToClickhouseDbResponseBody) SetErrorMsg(v string) *ModifyRDSToClickhouseDbResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ModifyRDSToClickhouseDbResponseBody) SetRequestId(v string) *ModifyRDSToClickhouseDbResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRDSToClickhouseDbResponseBody) SetStatus(v int64) *ModifyRDSToClickhouseDbResponseBody {
	s.Status = &v
	return s
}

type ModifyRDSToClickhouseDbResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyRDSToClickhouseDbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyRDSToClickhouseDbResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRDSToClickhouseDbResponse) GoString() string {
	return s.String()
}

func (s *ModifyRDSToClickhouseDbResponse) SetHeaders(v map[string]*string) *ModifyRDSToClickhouseDbResponse {
	s.Headers = v
	return s
}

func (s *ModifyRDSToClickhouseDbResponse) SetBody(v *ModifyRDSToClickhouseDbResponseBody) *ModifyRDSToClickhouseDbResponse {
	s.Body = v
	return s
}

type OperateLogHubRequest struct {
	AccessKey            *string                             `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	AccessSecret         *string                             `json:"AccessSecret,omitempty" xml:"AccessSecret,omitempty"`
	Create               *bool                               `json:"Create,omitempty" xml:"Create,omitempty"`
	DBClusterId          *string                             `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DeliverName          *string                             `json:"DeliverName,omitempty" xml:"DeliverName,omitempty"`
	DeliverTime          *string                             `json:"DeliverTime,omitempty" xml:"DeliverTime,omitempty"`
	Description          *string                             `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainUrl            *string                             `json:"DomainUrl,omitempty" xml:"DomainUrl,omitempty"`
	FilterDirtyData      *bool                               `json:"FilterDirtyData,omitempty" xml:"FilterDirtyData,omitempty"`
	LogHubStores         []*OperateLogHubRequestLogHubStores `json:"LogHubStores,omitempty" xml:"LogHubStores,omitempty" type:"Repeated"`
	LogStoreName         *string                             `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	OwnerAccount         *string                             `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Password             *string                             `json:"Password,omitempty" xml:"Password,omitempty"`
	ProjectName          *string                             `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ResourceOwnerAccount *string                             `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                              `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SchemaName           *string                             `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName            *string                             `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TaskId               *string                             `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UseLorne             *bool                               `json:"UseLorne,omitempty" xml:"UseLorne,omitempty"`
	UserName             *string                             `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s OperateLogHubRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateLogHubRequest) GoString() string {
	return s.String()
}

func (s *OperateLogHubRequest) SetAccessKey(v string) *OperateLogHubRequest {
	s.AccessKey = &v
	return s
}

func (s *OperateLogHubRequest) SetAccessSecret(v string) *OperateLogHubRequest {
	s.AccessSecret = &v
	return s
}

func (s *OperateLogHubRequest) SetCreate(v bool) *OperateLogHubRequest {
	s.Create = &v
	return s
}

func (s *OperateLogHubRequest) SetDBClusterId(v string) *OperateLogHubRequest {
	s.DBClusterId = &v
	return s
}

func (s *OperateLogHubRequest) SetDeliverName(v string) *OperateLogHubRequest {
	s.DeliverName = &v
	return s
}

func (s *OperateLogHubRequest) SetDeliverTime(v string) *OperateLogHubRequest {
	s.DeliverTime = &v
	return s
}

func (s *OperateLogHubRequest) SetDescription(v string) *OperateLogHubRequest {
	s.Description = &v
	return s
}

func (s *OperateLogHubRequest) SetDomainUrl(v string) *OperateLogHubRequest {
	s.DomainUrl = &v
	return s
}

func (s *OperateLogHubRequest) SetFilterDirtyData(v bool) *OperateLogHubRequest {
	s.FilterDirtyData = &v
	return s
}

func (s *OperateLogHubRequest) SetLogHubStores(v []*OperateLogHubRequestLogHubStores) *OperateLogHubRequest {
	s.LogHubStores = v
	return s
}

func (s *OperateLogHubRequest) SetLogStoreName(v string) *OperateLogHubRequest {
	s.LogStoreName = &v
	return s
}

func (s *OperateLogHubRequest) SetOwnerAccount(v string) *OperateLogHubRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OperateLogHubRequest) SetOwnerId(v int64) *OperateLogHubRequest {
	s.OwnerId = &v
	return s
}

func (s *OperateLogHubRequest) SetPassword(v string) *OperateLogHubRequest {
	s.Password = &v
	return s
}

func (s *OperateLogHubRequest) SetProjectName(v string) *OperateLogHubRequest {
	s.ProjectName = &v
	return s
}

func (s *OperateLogHubRequest) SetResourceOwnerAccount(v string) *OperateLogHubRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OperateLogHubRequest) SetResourceOwnerId(v int64) *OperateLogHubRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OperateLogHubRequest) SetSchemaName(v string) *OperateLogHubRequest {
	s.SchemaName = &v
	return s
}

func (s *OperateLogHubRequest) SetTableName(v string) *OperateLogHubRequest {
	s.TableName = &v
	return s
}

func (s *OperateLogHubRequest) SetTaskId(v string) *OperateLogHubRequest {
	s.TaskId = &v
	return s
}

func (s *OperateLogHubRequest) SetUseLorne(v bool) *OperateLogHubRequest {
	s.UseLorne = &v
	return s
}

func (s *OperateLogHubRequest) SetUserName(v string) *OperateLogHubRequest {
	s.UserName = &v
	return s
}

type OperateLogHubRequestLogHubStores struct {
	FieldKey *string `json:"FieldKey,omitempty" xml:"FieldKey,omitempty"`
	LogKey   *string `json:"LogKey,omitempty" xml:"LogKey,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OperateLogHubRequestLogHubStores) String() string {
	return tea.Prettify(s)
}

func (s OperateLogHubRequestLogHubStores) GoString() string {
	return s.String()
}

func (s *OperateLogHubRequestLogHubStores) SetFieldKey(v string) *OperateLogHubRequestLogHubStores {
	s.FieldKey = &v
	return s
}

func (s *OperateLogHubRequestLogHubStores) SetLogKey(v string) *OperateLogHubRequestLogHubStores {
	s.LogKey = &v
	return s
}

func (s *OperateLogHubRequestLogHubStores) SetType(v string) *OperateLogHubRequestLogHubStores {
	s.Type = &v
	return s
}

type OperateLogHubResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateLogHubResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateLogHubResponseBody) GoString() string {
	return s.String()
}

func (s *OperateLogHubResponseBody) SetRequestId(v string) *OperateLogHubResponseBody {
	s.RequestId = &v
	return s
}

type OperateLogHubResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OperateLogHubResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OperateLogHubResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateLogHubResponse) GoString() string {
	return s.String()
}

func (s *OperateLogHubResponse) SetHeaders(v map[string]*string) *OperateLogHubResponse {
	s.Headers = v
	return s
}

func (s *OperateLogHubResponse) SetBody(v *OperateLogHubResponseBody) *OperateLogHubResponse {
	s.Body = v
	return s
}

type OperateLorneTaskStatusRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	LorneStatus          *string `json:"LorneStatus,omitempty" xml:"LorneStatus,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OperateLorneTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateLorneTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *OperateLorneTaskStatusRequest) SetDBClusterId(v string) *OperateLorneTaskStatusRequest {
	s.DBClusterId = &v
	return s
}

func (s *OperateLorneTaskStatusRequest) SetLorneStatus(v string) *OperateLorneTaskStatusRequest {
	s.LorneStatus = &v
	return s
}

func (s *OperateLorneTaskStatusRequest) SetOwnerAccount(v string) *OperateLorneTaskStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OperateLorneTaskStatusRequest) SetOwnerId(v int64) *OperateLorneTaskStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *OperateLorneTaskStatusRequest) SetResourceOwnerAccount(v string) *OperateLorneTaskStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OperateLorneTaskStatusRequest) SetResourceOwnerId(v int64) *OperateLorneTaskStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OperateLorneTaskStatusRequest) SetTaskId(v string) *OperateLorneTaskStatusRequest {
	s.TaskId = &v
	return s
}

type OperateLorneTaskStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateLorneTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateLorneTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *OperateLorneTaskStatusResponseBody) SetRequestId(v string) *OperateLorneTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

type OperateLorneTaskStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OperateLorneTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OperateLorneTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateLorneTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *OperateLorneTaskStatusResponse) SetHeaders(v map[string]*string) *OperateLorneTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *OperateLorneTaskStatusResponse) SetBody(v *OperateLorneTaskStatusResponseBody) *OperateLorneTaskStatusResponse {
	s.Body = v
	return s
}

type ReleaseClusterPublicConnectionRequest struct {
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseClusterPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ReleaseClusterPublicConnectionResponse) SetBody(v *ReleaseClusterPublicConnectionResponseBody) *ReleaseClusterPublicConnectionResponse {
	s.Body = v
	return s
}

type ResetAccountPasswordRequest struct {
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponseBody) SetRequestId(v string) *ResetAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ResetAccountPasswordResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ResetAccountPasswordResponse) SetBody(v *ResetAccountPasswordResponseBody) *ResetAccountPasswordResponse {
	s.Body = v
	return s
}

type RestartInstanceRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RestartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceRequest) SetDBClusterId(v string) *RestartInstanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *RestartInstanceRequest) SetOwnerAccount(v string) *RestartInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartInstanceRequest) SetOwnerId(v int64) *RestartInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartInstanceRequest) SetPageNumber(v int32) *RestartInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *RestartInstanceRequest) SetPageSize(v int32) *RestartInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *RestartInstanceRequest) SetRegionId(v string) *RestartInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RestartInstanceRequest) SetResourceOwnerAccount(v string) *RestartInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartInstanceRequest) SetResourceOwnerId(v int64) *RestartInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

type RestartInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBody) SetRequestId(v string) *RestartInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestartInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponse) SetHeaders(v map[string]*string) *RestartInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartInstanceResponse) SetBody(v *RestartInstanceResponseBody) *RestartInstanceResponse {
	s.Body = v
	return s
}

type SearchRDSTablesRequest struct {
	DbClusterId          *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RdsId                *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	RdsPassword          *string `json:"RdsPassword,omitempty" xml:"RdsPassword,omitempty"`
	RdsPort              *int64  `json:"RdsPort,omitempty" xml:"RdsPort,omitempty"`
	RdsUserName          *string `json:"RdsUserName,omitempty" xml:"RdsUserName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TableName            *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s SearchRDSTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchRDSTablesRequest) GoString() string {
	return s.String()
}

func (s *SearchRDSTablesRequest) SetDbClusterId(v string) *SearchRDSTablesRequest {
	s.DbClusterId = &v
	return s
}

func (s *SearchRDSTablesRequest) SetOwnerAccount(v string) *SearchRDSTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SearchRDSTablesRequest) SetOwnerId(v int64) *SearchRDSTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *SearchRDSTablesRequest) SetRdsId(v string) *SearchRDSTablesRequest {
	s.RdsId = &v
	return s
}

func (s *SearchRDSTablesRequest) SetRdsPassword(v string) *SearchRDSTablesRequest {
	s.RdsPassword = &v
	return s
}

func (s *SearchRDSTablesRequest) SetRdsPort(v int64) *SearchRDSTablesRequest {
	s.RdsPort = &v
	return s
}

func (s *SearchRDSTablesRequest) SetRdsUserName(v string) *SearchRDSTablesRequest {
	s.RdsUserName = &v
	return s
}

func (s *SearchRDSTablesRequest) SetResourceOwnerAccount(v string) *SearchRDSTablesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SearchRDSTablesRequest) SetResourceOwnerId(v int64) *SearchRDSTablesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SearchRDSTablesRequest) SetTableName(v string) *SearchRDSTablesRequest {
	s.TableName = &v
	return s
}

type SearchRDSTablesResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Schemas   []*SearchRDSTablesResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Repeated"`
}

func (s SearchRDSTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchRDSTablesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchRDSTablesResponseBody) SetRequestId(v string) *SearchRDSTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchRDSTablesResponseBody) SetSchemas(v []*SearchRDSTablesResponseBodySchemas) *SearchRDSTablesResponseBody {
	s.Schemas = v
	return s
}

type SearchRDSTablesResponseBodySchemas struct {
	DbName *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Tables []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s SearchRDSTablesResponseBodySchemas) String() string {
	return tea.Prettify(s)
}

func (s SearchRDSTablesResponseBodySchemas) GoString() string {
	return s.String()
}

func (s *SearchRDSTablesResponseBodySchemas) SetDbName(v string) *SearchRDSTablesResponseBodySchemas {
	s.DbName = &v
	return s
}

func (s *SearchRDSTablesResponseBodySchemas) SetTables(v []*string) *SearchRDSTablesResponseBodySchemas {
	s.Tables = v
	return s
}

type SearchRDSTablesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchRDSTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchRDSTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchRDSTablesResponse) GoString() string {
	return s.String()
}

func (s *SearchRDSTablesResponse) SetHeaders(v map[string]*string) *SearchRDSTablesResponse {
	s.Headers = v
	return s
}

func (s *SearchRDSTablesResponse) SetBody(v *SearchRDSTablesResponseBody) *SearchRDSTablesResponse {
	s.Body = v
	return s
}

type TransferVersionRequest struct {
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SourceAccount        *string `json:"SourceAccount,omitempty" xml:"SourceAccount,omitempty"`
	SourcePassword       *string `json:"SourcePassword,omitempty" xml:"SourcePassword,omitempty"`
	TargetAccount        *string `json:"TargetAccount,omitempty" xml:"TargetAccount,omitempty"`
	TargetDbClusterId    *string `json:"TargetDbClusterId,omitempty" xml:"TargetDbClusterId,omitempty"`
	TargetPassword       *string `json:"TargetPassword,omitempty" xml:"TargetPassword,omitempty"`
}

func (s TransferVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferVersionRequest) GoString() string {
	return s.String()
}

func (s *TransferVersionRequest) SetDBClusterId(v string) *TransferVersionRequest {
	s.DBClusterId = &v
	return s
}

func (s *TransferVersionRequest) SetOwnerAccount(v string) *TransferVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TransferVersionRequest) SetOwnerId(v int64) *TransferVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *TransferVersionRequest) SetPageNumber(v int32) *TransferVersionRequest {
	s.PageNumber = &v
	return s
}

func (s *TransferVersionRequest) SetPageSize(v int32) *TransferVersionRequest {
	s.PageSize = &v
	return s
}

func (s *TransferVersionRequest) SetRegionId(v string) *TransferVersionRequest {
	s.RegionId = &v
	return s
}

func (s *TransferVersionRequest) SetResourceOwnerAccount(v string) *TransferVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TransferVersionRequest) SetResourceOwnerId(v int64) *TransferVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TransferVersionRequest) SetSourceAccount(v string) *TransferVersionRequest {
	s.SourceAccount = &v
	return s
}

func (s *TransferVersionRequest) SetSourcePassword(v string) *TransferVersionRequest {
	s.SourcePassword = &v
	return s
}

func (s *TransferVersionRequest) SetTargetAccount(v string) *TransferVersionRequest {
	s.TargetAccount = &v
	return s
}

func (s *TransferVersionRequest) SetTargetDbClusterId(v string) *TransferVersionRequest {
	s.TargetDbClusterId = &v
	return s
}

func (s *TransferVersionRequest) SetTargetPassword(v string) *TransferVersionRequest {
	s.TargetPassword = &v
	return s
}

type TransferVersionResponseBody struct {
	DBInstanceID   *int32  `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	DBInstanceName *int64  `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId         *bool   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TransferVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferVersionResponseBody) GoString() string {
	return s.String()
}

func (s *TransferVersionResponseBody) SetDBInstanceID(v int32) *TransferVersionResponseBody {
	s.DBInstanceID = &v
	return s
}

func (s *TransferVersionResponseBody) SetDBInstanceName(v int64) *TransferVersionResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *TransferVersionResponseBody) SetRequestId(v string) *TransferVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferVersionResponseBody) SetTaskId(v bool) *TransferVersionResponseBody {
	s.TaskId = &v
	return s
}

type TransferVersionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransferVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferVersionResponse) GoString() string {
	return s.String()
}

func (s *TransferVersionResponse) SetHeaders(v map[string]*string) *TransferVersionResponse {
	s.Headers = v
	return s
}

func (s *TransferVersionResponse) SetBody(v *TransferVersionResponseBody) *TransferVersionResponse {
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
		"ap-northeast-2-pop":          tea.String("clickhouse.aliyuncs.com"),
		"ap-southeast-1":              tea.String("clickhouse.aliyuncs.com"),
		"cn-beijing":                  tea.String("clickhouse.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("clickhouse.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("clickhouse.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("clickhouse.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("clickhouse.aliyuncs.com"),
		"cn-edge-1":                   tea.String("clickhouse.aliyuncs.com"),
		"cn-fujian":                   tea.String("clickhouse.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("clickhouse.aliyuncs.com"),
		"cn-hongkong":                 tea.String("clickhouse.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("clickhouse.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("clickhouse.aliyuncs.com"),
		"cn-qingdao":                  tea.String("clickhouse.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("clickhouse.aliyuncs.com"),
		"cn-shanghai":                 tea.String("clickhouse.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("clickhouse.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("clickhouse.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("clickhouse.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("clickhouse.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("clickhouse.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("clickhouse.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("clickhouse.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("clickhouse.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("clickhouse.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("clickhouse.aliyuncs.com"),
		"cn-wuhan":                    tea.String("clickhouse.aliyuncs.com"),
		"cn-yushanfang":               tea.String("clickhouse.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("clickhouse.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("clickhouse.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("clickhouse.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("clickhouse.aliyuncs.com"),
		"me-east-1":                   tea.String("clickhouse.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("clickhouse.aliyuncs.com"),
		"us-east-1":                   tea.String("clickhouse.aliyuncs.com"),
		"us-west-1":                   tea.String("clickhouse.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("clickhouse"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
	query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("AllocateClusterPublicConnection"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) CheckClickhouseToRDSWithOptions(request *CheckClickhouseToRDSRequest, runtime *util.RuntimeOptions) (_result *CheckClickhouseToRDSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CkPassword"] = request.CkPassword
	query["CkUserName"] = request.CkUserName
	query["ClickhousePort"] = request.ClickhousePort
	query["DbClusterId"] = request.DbClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RdsId"] = request.RdsId
	query["RdsPassword"] = request.RdsPassword
	query["RdsPort"] = request.RdsPort
	query["RdsUserName"] = request.RdsUserName
	query["RdsVpcId"] = request.RdsVpcId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckClickhouseToRDS"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckClickhouseToRDSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckClickhouseToRDS(request *CheckClickhouseToRDSRequest) (_result *CheckClickhouseToRDSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckClickhouseToRDSResponse{}
	_body, _err := client.CheckClickhouseToRDSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckScaleOutBalancedWithOptions(request *CheckScaleOutBalancedRequest, runtime *util.RuntimeOptions) (_result *CheckScaleOutBalancedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckScaleOutBalanced"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckScaleOutBalancedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckScaleOutBalanced(request *CheckScaleOutBalancedRequest) (_result *CheckScaleOutBalancedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckScaleOutBalancedResponse{}
	_body, _err := client.CheckScaleOutBalancedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckServiceLinkedRoleWithOptions(request *CheckServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CheckServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckServiceLinkedRole"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckServiceLinkedRole(request *CheckServiceLinkedRoleRequest) (_result *CheckServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.CheckServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckVersionTransferWithOptions(request *CheckVersionTransferRequest, runtime *util.RuntimeOptions) (_result *CheckVersionTransferResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CheckAccount"] = request.CheckAccount
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SourceAccount"] = request.SourceAccount
	query["SourcePassword"] = request.SourcePassword
	query["TargetAccount"] = request.TargetAccount
	query["TargetDbClusterId"] = request.TargetDbClusterId
	query["TargetPassword"] = request.TargetPassword
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckVersionTransfer"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckVersionTransferResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckVersionTransfer(request *CheckVersionTransferRequest) (_result *CheckVersionTransferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckVersionTransferResponse{}
	_body, _err := client.CheckVersionTransferWithOptions(request, runtime)
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
	query["AccountDescription"] = request.AccountDescription
	query["AccountName"] = request.AccountName
	query["AccountPassword"] = request.AccountPassword
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccount"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) CreateAccountAndAuthorityWithOptions(request *CreateAccountAndAuthorityRequest, runtime *util.RuntimeOptions) (_result *CreateAccountAndAuthorityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccountDescription"] = request.AccountDescription
	query["AccountName"] = request.AccountName
	query["AccountPassword"] = request.AccountPassword
	query["AllowDatabases"] = request.AllowDatabases
	query["AllowDictionaries"] = request.AllowDictionaries
	query["DBClusterId"] = request.DBClusterId
	query["DdlAuthority"] = request.DdlAuthority
	query["DmlAuthority"] = request.DmlAuthority
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TotalDatabases"] = request.TotalDatabases
	query["TotalDictionaries"] = request.TotalDictionaries
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccountAndAuthority"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccountAndAuthorityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccountAndAuthority(request *CreateAccountAndAuthorityRequest) (_result *CreateAccountAndAuthorityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccountAndAuthorityResponse{}
	_body, _err := client.CreateAccountAndAuthorityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBackupPolicyWithOptions(request *CreateBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	query["PreferredBackupTime"] = request.PreferredBackupTime
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBackupPolicy"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBackupPolicy(request *CreateBackupPolicyRequest) (_result *CreateBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBackupPolicyResponse{}
	_body, _err := client.CreateBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBInstanceWithOptions(request *CreateDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["DBClusterCategory"] = request.DBClusterCategory
	query["DBClusterClass"] = request.DBClusterClass
	query["DBClusterDescription"] = request.DBClusterDescription
	query["DBClusterNetworkType"] = request.DBClusterNetworkType
	query["DBClusterVersion"] = request.DBClusterVersion
	query["DBNodeGroupCount"] = request.DBNodeGroupCount
	query["DBNodeStorage"] = request.DBNodeStorage
	query["DbNodeStorageType"] = request.DbNodeStorageType
	query["EncryptionKey"] = request.EncryptionKey
	query["EncryptionType"] = request.EncryptionType
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PayType"] = request.PayType
	query["Period"] = request.Period
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["UsedTime"] = request.UsedTime
	query["VPCId"] = request.VPCId
	query["VSwitchId"] = request.VSwitchId
	query["ZoneId"] = request.ZoneId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBInstance"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBInstance(request *CreateDBInstanceRequest) (_result *CreateDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.CreateDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOSSStorageWithOptions(request *CreateOSSStorageRequest, runtime *util.RuntimeOptions) (_result *CreateOSSStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOSSStorage"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOSSStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOSSStorage(request *CreateOSSStorageRequest) (_result *CreateOSSStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOSSStorageResponse{}
	_body, _err := client.CreateOSSStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePortsForClickHouseWithOptions(request *CreatePortsForClickHouseRequest, runtime *util.RuntimeOptions) (_result *CreatePortsForClickHouseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PortType"] = request.PortType
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePortsForClickHouse"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePortsForClickHouseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePortsForClickHouse(request *CreatePortsForClickHouseRequest) (_result *CreatePortsForClickHouseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePortsForClickHouseResponse{}
	_body, _err := client.CreatePortsForClickHouseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRDSToClickhouseDbWithOptions(request *CreateRDSToClickhouseDbRequest, runtime *util.RuntimeOptions) (_result *CreateRDSToClickhouseDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CkPassword"] = request.CkPassword
	query["CkUserName"] = request.CkUserName
	query["ClickhousePort"] = request.ClickhousePort
	query["DbClusterId"] = request.DbClusterId
	query["LimitUpper"] = request.LimitUpper
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RdsId"] = request.RdsId
	query["RdsPassword"] = request.RdsPassword
	query["RdsPort"] = request.RdsPort
	query["RdsUserName"] = request.RdsUserName
	query["RdsVpcId"] = request.RdsVpcId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SkipUnsupported"] = request.SkipUnsupported
	query["SynDbTables"] = request.SynDbTables
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRDSToClickhouseDb"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRDSToClickhouseDbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRDSToClickhouseDb(request *CreateRDSToClickhouseDbRequest) (_result *CreateRDSToClickhouseDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRDSToClickhouseDbResponse{}
	_body, _err := client.CreateRDSToClickhouseDbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceLinkedRole"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(request, runtime)
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
	query["AccountName"] = request.AccountName
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccount"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DeleteDBClusterWithOptions(request *DeleteDBClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDBCluster"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DeleteLorneTaskWithOptions(request *DeleteLorneTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteLorneTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLorneTask"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLorneTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLorneTask(request *DeleteLorneTaskRequest) (_result *DeleteLorneTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLorneTaskResponse{}
	_body, _err := client.DeleteLorneTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSyndbWithOptions(request *DeleteSyndbRequest, runtime *util.RuntimeOptions) (_result *DeleteSyndbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbClusterId"] = request.DbClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SynDb"] = request.SynDb
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSyndb"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSyndbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSyndb(request *DeleteSyndbRequest) (_result *DeleteSyndbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSyndbResponse{}
	_body, _err := client.DeleteSyndbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountAuthorityWithOptions(request *DescribeAccountAuthorityRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountAuthorityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccountName"] = request.AccountName
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccountAuthority"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccountAuthorityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccountAuthority(request *DescribeAccountAuthorityRequest) (_result *DescribeAccountAuthorityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountAuthorityResponse{}
	_body, _err := client.DescribeAccountAuthorityWithOptions(request, runtime)
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
	query["AccountName"] = request.AccountName
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccounts"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeAllDataSourceWithOptions(request *DescribeAllDataSourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAllDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SchemaName"] = request.SchemaName
	query["TableName"] = request.TableName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAllDataSource"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeAllDataSourcesWithOptions(request *DescribeAllDataSourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeAllDataSourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SchemaName"] = request.SchemaName
	query["TableName"] = request.TableName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAllDataSources"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAllDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAllDataSources(request *DescribeAllDataSourcesRequest) (_result *DescribeAllDataSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAllDataSourcesResponse{}
	_body, _err := client.DescribeAllDataSourcesWithOptions(request, runtime)
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
	query["AcceptLanguage"] = request.AcceptLanguage
	query["ChargeType"] = request.ChargeType
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["ZoneId"] = request.ZoneId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAvailableResource"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupPolicy"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["BackupId"] = request.BackupId
	query["DBClusterId"] = request.DBClusterId
	query["EndTime"] = request.EndTime
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["StartTime"] = request.StartTime
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackups"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SchemaName"] = request.SchemaName
	query["TableName"] = request.TableName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeColumns"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeDBClusterAccessWhiteListWithOptions(request *DescribeDBClusterAccessWhiteListRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterAccessWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterAccessWhiteList"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterAttribute"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeDBClusterConfigWithOptions(request *DescribeDBClusterConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterConfig"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterConfig(request *DescribeDBClusterConfigRequest) (_result *DescribeDBClusterConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterConfigResponse{}
	_body, _err := client.DescribeDBClusterConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterNetInfoItemsWithOptions(request *DescribeDBClusterNetInfoItemsRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterNetInfoItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterNetInfoItems"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterNetInfoItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterNetInfoItems(request *DescribeDBClusterNetInfoItemsRequest) (_result *DescribeDBClusterNetInfoItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterNetInfoItemsResponse{}
	_body, _err := client.DescribeDBClusterNetInfoItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterPerformanceWithOptions(request *DescribeDBClusterPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["EndTime"] = request.EndTime
	query["Key"] = request.Key
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["StartTime"] = request.StartTime
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusterPerformance"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeDBClustersWithOptions(request *DescribeDBClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterDescription"] = request.DBClusterDescription
	query["DBClusterIds"] = request.DBClusterIds
	query["DBClusterStatus"] = request.DBClusterStatus
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBClusters"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeDBConfigWithOptions(request *DescribeDBConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDBConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBConfig"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBConfig(request *DescribeDBConfigRequest) (_result *DescribeDBConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBConfigResponse{}
	_body, _err := client.DescribeDBConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogHubAttributeWithOptions(request *DescribeLogHubAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeLogHubAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["DeliverName"] = request.DeliverName
	query["LogStoreName"] = request.LogStoreName
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ProjectName"] = request.ProjectName
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogHubAttribute"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogHubAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogHubAttribute(request *DescribeLogHubAttributeRequest) (_result *DescribeLogHubAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogHubAttributeResponse{}
	_body, _err := client.DescribeLogHubAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogStoreKeysWithOptions(request *DescribeLogStoreKeysRequest, runtime *util.RuntimeOptions) (_result *DescribeLogStoreKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["LogStoreName"] = request.LogStoreName
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ProjectName"] = request.ProjectName
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogStoreKeys"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogStoreKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogStoreKeys(request *DescribeLogStoreKeysRequest) (_result *DescribeLogStoreKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogStoreKeysResponse{}
	_body, _err := client.DescribeLogStoreKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLoghubDetailWithOptions(request *DescribeLoghubDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeLoghubDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ExportName"] = request.ExportName
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ProjectName"] = request.ProjectName
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLoghubDetail"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLoghubDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLoghubDetail(request *DescribeLoghubDetailRequest) (_result *DescribeLoghubDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLoghubDetailResponse{}
	_body, _err := client.DescribeLoghubDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLorneLogWithOptions(request *DescribeLorneLogRequest, runtime *util.RuntimeOptions) (_result *DescribeLorneLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["EndTime"] = request.EndTime
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["StartTime"] = request.StartTime
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLorneLog"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLorneLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLorneLog(request *DescribeLorneLogRequest) (_result *DescribeLorneLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLorneLogResponse{}
	_body, _err := client.DescribeLorneLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLorneTasksWithOptions(request *DescribeLorneTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeLorneTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLorneTasks"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLorneTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLorneTasks(request *DescribeLorneTasksRequest) (_result *DescribeLorneTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLorneTasksResponse{}
	_body, _err := client.DescribeLorneTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLorneTasksMCountWithOptions(request *DescribeLorneTasksMCountRequest, runtime *util.RuntimeOptions) (_result *DescribeLorneTasksMCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["EndTime"] = request.EndTime
	query["MetricName"] = request.MetricName
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["StartTime"] = request.StartTime
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLorneTasksMCount"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLorneTasksMCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLorneTasksMCount(request *DescribeLorneTasksMCountRequest) (_result *DescribeLorneTasksMCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLorneTasksMCountResponse{}
	_body, _err := client.DescribeLorneTasksMCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLorneTasksMetricsWithOptions(request *DescribeLorneTasksMetricsRequest, runtime *util.RuntimeOptions) (_result *DescribeLorneTasksMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["EndTime"] = request.EndTime
	query["MetricName"] = request.MetricName
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["StartTime"] = request.StartTime
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLorneTasksMetrics"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLorneTasksMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLorneTasksMetrics(request *DescribeLorneTasksMetricsRequest) (_result *DescribeLorneTasksMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLorneTasksMetricsResponse{}
	_body, _err := client.DescribeLorneTasksMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOSSStorageWithOptions(request *DescribeOSSStorageRequest, runtime *util.RuntimeOptions) (_result *DescribeOSSStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOSSStorage"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOSSStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOSSStorage(request *DescribeOSSStorageRequest) (_result *DescribeOSSStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOSSStorageResponse{}
	_body, _err := client.DescribeOSSStorageWithOptions(request, runtime)
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
	query["DBClusterId"] = request.DBClusterId
	query["InitialQueryId"] = request.InitialQueryId
	query["InitialUser"] = request.InitialUser
	query["Keyword"] = request.Keyword
	query["Order"] = request.Order
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["QueryDurationMs"] = request.QueryDurationMs
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProcessList"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeRDSTablesWithOptions(request *DescribeRDSTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeRDSTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbClusterId"] = request.DbClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RdsId"] = request.RdsId
	query["RdsPassword"] = request.RdsPassword
	query["RdsPort"] = request.RdsPort
	query["RdsUserName"] = request.RdsUserName
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Schema"] = request.Schema
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRDSTables"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRDSTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRDSTables(request *DescribeRDSTablesRequest) (_result *DescribeRDSTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRDSTablesResponse{}
	_body, _err := client.DescribeRDSTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRDSVpcWithOptions(request *DescribeRDSVpcRequest, runtime *util.RuntimeOptions) (_result *DescribeRDSVpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbClusterId"] = request.DbClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RdsId"] = request.RdsId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRDSVpc"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRDSVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRDSVpc(request *DescribeRDSVpcRequest) (_result *DescribeRDSVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRDSVpcResponse{}
	_body, _err := client.DescribeRDSVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRDSschemasWithOptions(request *DescribeRDSschemasRequest, runtime *util.RuntimeOptions) (_result *DescribeRDSschemasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbClusterId"] = request.DbClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RdsId"] = request.RdsId
	query["RdsPassword"] = request.RdsPassword
	query["RdsPort"] = request.RdsPort
	query["RdsUserName"] = request.RdsUserName
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRDSschemas"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRDSschemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRDSschemas(request *DescribeRDSschemasRequest) (_result *DescribeRDSschemasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRDSschemasResponse{}
	_body, _err := client.DescribeRDSschemasWithOptions(request, runtime)
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
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeSchemasWithOptions(request *DescribeSchemasRequest, runtime *util.RuntimeOptions) (_result *DescribeSchemasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SchemaName"] = request.SchemaName
	query["TableName"] = request.TableName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSchemas"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["DBClusterId"] = request.DBClusterId
	query["EndTime"] = request.EndTime
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["QueryDurationMs"] = request.QueryDurationMs
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["StartTime"] = request.StartTime
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlowLogRecords"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["DBClusterId"] = request.DBClusterId
	query["EndTime"] = request.EndTime
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["QueryDurationMs"] = request.QueryDurationMs
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["StartTime"] = request.StartTime
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlowLogTrend"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeSynDbTablesWithOptions(request *DescribeSynDbTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeSynDbTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbClusterId"] = request.DbClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SynDb"] = request.SynDb
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSynDbTables"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSynDbTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSynDbTables(request *DescribeSynDbTablesRequest) (_result *DescribeSynDbTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSynDbTablesResponse{}
	_body, _err := client.DescribeSynDbTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSynDbsWithOptions(request *DescribeSynDbsRequest, runtime *util.RuntimeOptions) (_result *DescribeSynDbsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbClusterId"] = request.DbClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSynDbs"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSynDbsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSynDbs(request *DescribeSynDbsRequest) (_result *DescribeSynDbsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSynDbsResponse{}
	_body, _err := client.DescribeSynDbsWithOptions(request, runtime)
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
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SchemaName"] = request.SchemaName
	query["TableName"] = request.TableName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTables"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeTransferHistoryWithOptions(request *DescribeTransferHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeTransferHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTransferHistory"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTransferHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTransferHistory(request *DescribeTransferHistoryRequest) (_result *DescribeTransferHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTransferHistoryResponse{}
	_body, _err := client.DescribeTransferHistoryWithOptions(request, runtime)
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
	query["DBClusterId"] = request.DBClusterId
	query["InitialQueryId"] = request.InitialQueryId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("KillProcess"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ModifyAccountAuthorityWithOptions(request *ModifyAccountAuthorityRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountAuthorityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccountName"] = request.AccountName
	query["AllowDatabases"] = request.AllowDatabases
	query["AllowDictionaries"] = request.AllowDictionaries
	query["DBClusterId"] = request.DBClusterId
	query["DdlAuthority"] = request.DdlAuthority
	query["DmlAuthority"] = request.DmlAuthority
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TotalDatabases"] = request.TotalDatabases
	query["TotalDictionaries"] = request.TotalDictionaries
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountAuthority"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccountAuthorityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountAuthority(request *ModifyAccountAuthorityRequest) (_result *ModifyAccountAuthorityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountAuthorityResponse{}
	_body, _err := client.ModifyAccountAuthorityWithOptions(request, runtime)
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
	query["AccountDescription"] = request.AccountDescription
	query["AccountName"] = request.AccountName
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountDescription"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ModifyBackupPolicyWithOptions(request *ModifyBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	query["PreferredBackupTime"] = request.PreferredBackupTime
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyBackupPolicy"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ModifyDBClusterWithOptions(request *ModifyDBClusterRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterClass"] = request.DBClusterClass
	query["DBClusterId"] = request.DBClusterId
	query["DBNodeGroupCount"] = request.DBNodeGroupCount
	query["DBNodeStorage"] = request.DBNodeStorage
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBCluster"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["DBClusterIPArrayAttribute"] = request.DBClusterIPArrayAttribute
	query["DBClusterIPArrayName"] = request.DBClusterIPArrayName
	query["DBClusterId"] = request.DBClusterId
	query["ModifyMode"] = request.ModifyMode
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SecurityIps"] = request.SecurityIps
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterAccessWhiteList"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ModifyDBClusterConfigWithOptions(request *ModifyDBClusterConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["UserConfig"] = request.UserConfig
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterConfig"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterConfig(request *ModifyDBClusterConfigRequest) (_result *ModifyDBClusterConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterConfigResponse{}
	_body, _err := client.ModifyDBClusterConfigWithOptions(request, runtime)
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
	query["DBClusterDescription"] = request.DBClusterDescription
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterDescription"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["DBClusterId"] = request.DBClusterId
	query["MaintainTime"] = request.MaintainTime
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterMaintainTime"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ModifyDBConfigWithOptions(request *ModifyDBConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyDBConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Config"] = request.Config
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBConfig"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBConfig(request *ModifyDBConfigRequest) (_result *ModifyDBConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBConfigResponse{}
	_body, _err := client.ModifyDBConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRDSToClickhouseDbWithOptions(request *ModifyRDSToClickhouseDbRequest, runtime *util.RuntimeOptions) (_result *ModifyRDSToClickhouseDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CkPassword"] = request.CkPassword
	query["CkUserName"] = request.CkUserName
	query["ClickhousePort"] = request.ClickhousePort
	query["DbClusterId"] = request.DbClusterId
	query["LimitUpper"] = request.LimitUpper
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RdsId"] = request.RdsId
	query["RdsPassword"] = request.RdsPassword
	query["RdsPort"] = request.RdsPort
	query["RdsSynDb"] = request.RdsSynDb
	query["RdsSynTables"] = request.RdsSynTables
	query["RdsUserName"] = request.RdsUserName
	query["RdsVpcId"] = request.RdsVpcId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SkipUnsupported"] = request.SkipUnsupported
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyRDSToClickhouseDb"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyRDSToClickhouseDbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRDSToClickhouseDb(request *ModifyRDSToClickhouseDbRequest) (_result *ModifyRDSToClickhouseDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRDSToClickhouseDbResponse{}
	_body, _err := client.ModifyRDSToClickhouseDbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OperateLogHubWithOptions(request *OperateLogHubRequest, runtime *util.RuntimeOptions) (_result *OperateLogHubResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessKey"] = request.AccessKey
	query["AccessSecret"] = request.AccessSecret
	query["Create"] = request.Create
	query["DBClusterId"] = request.DBClusterId
	query["DeliverName"] = request.DeliverName
	query["DeliverTime"] = request.DeliverTime
	query["Description"] = request.Description
	query["DomainUrl"] = request.DomainUrl
	query["FilterDirtyData"] = request.FilterDirtyData
	query["LogHubStores"] = request.LogHubStores
	query["LogStoreName"] = request.LogStoreName
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["Password"] = request.Password
	query["ProjectName"] = request.ProjectName
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SchemaName"] = request.SchemaName
	query["TableName"] = request.TableName
	query["TaskId"] = request.TaskId
	query["UseLorne"] = request.UseLorne
	query["UserName"] = request.UserName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateLogHub"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &OperateLogHubResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OperateLogHub(request *OperateLogHubRequest) (_result *OperateLogHubResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateLogHubResponse{}
	_body, _err := client.OperateLogHubWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OperateLorneTaskStatusWithOptions(request *OperateLorneTaskStatusRequest, runtime *util.RuntimeOptions) (_result *OperateLorneTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["LorneStatus"] = request.LorneStatus
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateLorneTaskStatus"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &OperateLorneTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OperateLorneTaskStatus(request *OperateLorneTaskStatusRequest) (_result *OperateLorneTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateLorneTaskStatusResponse{}
	_body, _err := client.OperateLorneTaskStatusWithOptions(request, runtime)
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
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseClusterPublicConnection"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["AccountName"] = request.AccountName
	query["AccountPassword"] = request.AccountPassword
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAccountPassword"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) RestartInstanceWithOptions(request *RestartInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartInstance"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartInstance(request *RestartInstanceRequest) (_result *RestartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartInstanceResponse{}
	_body, _err := client.RestartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchRDSTablesWithOptions(request *SearchRDSTablesRequest, runtime *util.RuntimeOptions) (_result *SearchRDSTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DbClusterId"] = request.DbClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["RdsId"] = request.RdsId
	query["RdsPassword"] = request.RdsPassword
	query["RdsPort"] = request.RdsPort
	query["RdsUserName"] = request.RdsUserName
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TableName"] = request.TableName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchRDSTables"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchRDSTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchRDSTables(request *SearchRDSTablesRequest) (_result *SearchRDSTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchRDSTablesResponse{}
	_body, _err := client.SearchRDSTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferVersionWithOptions(request *TransferVersionRequest, runtime *util.RuntimeOptions) (_result *TransferVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DBClusterId"] = request.DBClusterId
	query["OwnerAccount"] = request.OwnerAccount
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["RegionId"] = request.RegionId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SourceAccount"] = request.SourceAccount
	query["SourcePassword"] = request.SourcePassword
	query["TargetAccount"] = request.TargetAccount
	query["TargetDbClusterId"] = request.TargetDbClusterId
	query["TargetPassword"] = request.TargetPassword
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferVersion"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferVersion(request *TransferVersionRequest) (_result *TransferVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferVersionResponse{}
	_body, _err := client.TransferVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

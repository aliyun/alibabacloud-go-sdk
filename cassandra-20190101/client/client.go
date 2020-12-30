// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AllocatePublicContactPointsRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s AllocatePublicContactPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicContactPointsRequest) GoString() string {
	return s.String()
}

func (s *AllocatePublicContactPointsRequest) SetClusterId(v string) *AllocatePublicContactPointsRequest {
	s.ClusterId = &v
	return s
}

func (s *AllocatePublicContactPointsRequest) SetDataCenterId(v string) *AllocatePublicContactPointsRequest {
	s.DataCenterId = &v
	return s
}

func (s *AllocatePublicContactPointsRequest) SetClientToken(v string) *AllocatePublicContactPointsRequest {
	s.ClientToken = &v
	return s
}

type AllocatePublicContactPointsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocatePublicContactPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicContactPointsResponseBody) GoString() string {
	return s.String()
}

func (s *AllocatePublicContactPointsResponseBody) SetRequestId(v string) *AllocatePublicContactPointsResponseBody {
	s.RequestId = &v
	return s
}

type AllocatePublicContactPointsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AllocatePublicContactPointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AllocatePublicContactPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicContactPointsResponse) GoString() string {
	return s.String()
}

func (s *AllocatePublicContactPointsResponse) SetHeaders(v map[string]*string) *AllocatePublicContactPointsResponse {
	s.Headers = v
	return s
}

func (s *AllocatePublicContactPointsResponse) SetBody(v *AllocatePublicContactPointsResponseBody) *AllocatePublicContactPointsResponse {
	s.Body = v
	return s
}

type CreateBackupPlanRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId    *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
	BackupTime      *string `json:"BackupTime,omitempty" xml:"BackupTime,omitempty"`
	BackupPeriod    *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	RetentionPeriod *int32  `json:"RetentionPeriod,omitempty" xml:"RetentionPeriod,omitempty"`
	Active          *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanRequest) SetClusterId(v string) *CreateBackupPlanRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetDataCenterId(v string) *CreateBackupPlanRequest {
	s.DataCenterId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetBackupTime(v string) *CreateBackupPlanRequest {
	s.BackupTime = &v
	return s
}

func (s *CreateBackupPlanRequest) SetBackupPeriod(v string) *CreateBackupPlanRequest {
	s.BackupPeriod = &v
	return s
}

func (s *CreateBackupPlanRequest) SetRetentionPeriod(v int32) *CreateBackupPlanRequest {
	s.RetentionPeriod = &v
	return s
}

func (s *CreateBackupPlanRequest) SetActive(v bool) *CreateBackupPlanRequest {
	s.Active = &v
	return s
}

func (s *CreateBackupPlanRequest) SetClientToken(v string) *CreateBackupPlanRequest {
	s.ClientToken = &v
	return s
}

type CreateBackupPlanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanResponseBody) SetRequestId(v string) *CreateBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

type CreateBackupPlanResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanResponse) SetHeaders(v map[string]*string) *CreateBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateBackupPlanResponse) SetBody(v *CreateBackupPlanResponseBody) *CreateBackupPlanResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	PayType         *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PeriodUnit      *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Period          *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	AutoRenew       *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId          *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ClusterName     *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	DataCenterName  *string `json:"DataCenterName,omitempty" xml:"DataCenterName,omitempty"`
	MajorVersion    *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	NodeCount       *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	DiskType        *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	DiskSize        *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId       *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetPayType(v string) *CreateClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateClusterRequest) SetPeriodUnit(v string) *CreateClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateClusterRequest) SetPeriod(v int32) *CreateClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateClusterRequest) SetAutoRenew(v bool) *CreateClusterRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateClusterRequest) SetAutoRenewPeriod(v int32) *CreateClusterRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateClusterRequest) SetClientToken(v string) *CreateClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterRequest) SetRegionId(v string) *CreateClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateClusterRequest) SetZoneId(v string) *CreateClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetDataCenterName(v string) *CreateClusterRequest {
	s.DataCenterName = &v
	return s
}

func (s *CreateClusterRequest) SetMajorVersion(v string) *CreateClusterRequest {
	s.MajorVersion = &v
	return s
}

func (s *CreateClusterRequest) SetInstanceType(v string) *CreateClusterRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateClusterRequest) SetNodeCount(v int32) *CreateClusterRequest {
	s.NodeCount = &v
	return s
}

func (s *CreateClusterRequest) SetDiskType(v string) *CreateClusterRequest {
	s.DiskType = &v
	return s
}

func (s *CreateClusterRequest) SetDiskSize(v int32) *CreateClusterRequest {
	s.DiskSize = &v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) SetVswitchId(v string) *CreateClusterRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateClusterRequest) SetPassword(v string) *CreateClusterRequest {
	s.Password = &v
	return s
}

func (s *CreateClusterRequest) SetResourceGroupId(v string) *CreateClusterRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

type CreateClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateDataCenterRequest struct {
	PayType         *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PeriodUnit      *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Period          *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	AutoRenew       *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId          *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterName  *string `json:"DataCenterName,omitempty" xml:"DataCenterName,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	NodeCount       *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	DiskType        *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	DiskSize        *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId       *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s CreateDataCenterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCenterRequest) GoString() string {
	return s.String()
}

func (s *CreateDataCenterRequest) SetPayType(v string) *CreateDataCenterRequest {
	s.PayType = &v
	return s
}

func (s *CreateDataCenterRequest) SetPeriodUnit(v string) *CreateDataCenterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDataCenterRequest) SetPeriod(v int32) *CreateDataCenterRequest {
	s.Period = &v
	return s
}

func (s *CreateDataCenterRequest) SetAutoRenew(v bool) *CreateDataCenterRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDataCenterRequest) SetAutoRenewPeriod(v int32) *CreateDataCenterRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateDataCenterRequest) SetClientToken(v string) *CreateDataCenterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDataCenterRequest) SetRegionId(v string) *CreateDataCenterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDataCenterRequest) SetZoneId(v string) *CreateDataCenterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDataCenterRequest) SetClusterId(v string) *CreateDataCenterRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateDataCenterRequest) SetDataCenterName(v string) *CreateDataCenterRequest {
	s.DataCenterName = &v
	return s
}

func (s *CreateDataCenterRequest) SetInstanceType(v string) *CreateDataCenterRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateDataCenterRequest) SetNodeCount(v int32) *CreateDataCenterRequest {
	s.NodeCount = &v
	return s
}

func (s *CreateDataCenterRequest) SetDiskType(v string) *CreateDataCenterRequest {
	s.DiskType = &v
	return s
}

func (s *CreateDataCenterRequest) SetDiskSize(v int32) *CreateDataCenterRequest {
	s.DiskSize = &v
	return s
}

func (s *CreateDataCenterRequest) SetVpcId(v string) *CreateDataCenterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDataCenterRequest) SetVswitchId(v string) *CreateDataCenterRequest {
	s.VswitchId = &v
	return s
}

type CreateDataCenterResponseBody struct {
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCenterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataCenterResponseBody) SetDataCenterId(v string) *CreateDataCenterResponseBody {
	s.DataCenterId = &v
	return s
}

func (s *CreateDataCenterResponseBody) SetRequestId(v string) *CreateDataCenterResponseBody {
	s.RequestId = &v
	return s
}

type CreateDataCenterResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDataCenterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataCenterResponse) GoString() string {
	return s.String()
}

func (s *CreateDataCenterResponse) SetHeaders(v map[string]*string) *CreateDataCenterResponse {
	s.Headers = v
	return s
}

func (s *CreateDataCenterResponse) SetBody(v *CreateDataCenterResponseBody) *CreateDataCenterResponse {
	s.Body = v
	return s
}

type DeleteBackupPlanRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
}

func (s DeleteBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupPlanRequest) SetClusterId(v string) *DeleteBackupPlanRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteBackupPlanRequest) SetDataCenterId(v string) *DeleteBackupPlanRequest {
	s.DataCenterId = &v
	return s
}

type DeleteBackupPlanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupPlanResponseBody) SetRequestId(v string) *DeleteBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBackupPlanResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupPlanResponse) SetHeaders(v map[string]*string) *DeleteBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupPlanResponse) SetBody(v *DeleteBackupPlanResponseBody) *DeleteBackupPlanResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

type DeleteClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DeleteDataCenterRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
}

func (s DeleteDataCenterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataCenterRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataCenterRequest) SetClusterId(v string) *DeleteDataCenterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteDataCenterRequest) SetDataCenterId(v string) *DeleteDataCenterRequest {
	s.DataCenterId = &v
	return s
}

type DeleteDataCenterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataCenterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataCenterResponseBody) SetRequestId(v string) *DeleteDataCenterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDataCenterResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDataCenterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDataCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataCenterResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataCenterResponse) SetHeaders(v map[string]*string) *DeleteDataCenterResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataCenterResponse) SetBody(v *DeleteDataCenterResponseBody) *DeleteDataCenterResponse {
	s.Body = v
	return s
}

type DeleteNodeToolExecutionHistoryRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
}

func (s DeleteNodeToolExecutionHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeToolExecutionHistoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodeToolExecutionHistoryRequest) SetClusterId(v string) *DeleteNodeToolExecutionHistoryRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteNodeToolExecutionHistoryRequest) SetJobId(v string) *DeleteNodeToolExecutionHistoryRequest {
	s.JobId = &v
	return s
}

func (s *DeleteNodeToolExecutionHistoryRequest) SetDataCenterId(v string) *DeleteNodeToolExecutionHistoryRequest {
	s.DataCenterId = &v
	return s
}

type DeleteNodeToolExecutionHistoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNodeToolExecutionHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeToolExecutionHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodeToolExecutionHistoryResponseBody) SetRequestId(v string) *DeleteNodeToolExecutionHistoryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNodeToolExecutionHistoryResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNodeToolExecutionHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNodeToolExecutionHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeToolExecutionHistoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodeToolExecutionHistoryResponse) SetHeaders(v map[string]*string) *DeleteNodeToolExecutionHistoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodeToolExecutionHistoryResponse) SetBody(v *DeleteNodeToolExecutionHistoryResponseBody) *DeleteNodeToolExecutionHistoryResponse {
	s.Body = v
	return s
}

type DescribeAccountsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) SetClusterId(v string) *DescribeAccountsRequest {
	s.ClusterId = &v
	return s
}

type DescribeAccountsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Accounts  *DescribeAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
}

func (s DescribeAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetAccounts(v *DescribeAccountsResponseBodyAccounts) *DescribeAccountsResponseBody {
	s.Accounts = v
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
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsAccount) SetName(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.Name = &v
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

type DescribeBackupRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
	BackupId     *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupType   *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
}

func (s DescribeBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupRequest) SetClusterId(v string) *DescribeBackupRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupRequest) SetDataCenterId(v string) *DescribeBackupRequest {
	s.DataCenterId = &v
	return s
}

func (s *DescribeBackupRequest) SetBackupId(v string) *DescribeBackupRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupRequest) SetBackupType(v string) *DescribeBackupRequest {
	s.BackupType = &v
	return s
}

type DescribeBackupResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Backup    *DescribeBackupResponseBodyBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Struct"`
}

func (s DescribeBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupResponseBody) SetRequestId(v string) *DescribeBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupResponseBody) SetBackup(v *DescribeBackupResponseBodyBackup) *DescribeBackupResponseBody {
	s.Backup = v
	return s
}

type DescribeBackupResponseBodyBackup struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	BackupType   *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupId     *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	Details      *string `json:"Details,omitempty" xml:"Details,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeBackupResponseBodyBackup) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupResponseBodyBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupResponseBodyBackup) SetEndTime(v string) *DescribeBackupResponseBodyBackup {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupResponseBodyBackup) SetStatus(v string) *DescribeBackupResponseBodyBackup {
	s.Status = &v
	return s
}

func (s *DescribeBackupResponseBodyBackup) SetStartTime(v string) *DescribeBackupResponseBodyBackup {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupResponseBodyBackup) SetSize(v int64) *DescribeBackupResponseBodyBackup {
	s.Size = &v
	return s
}

func (s *DescribeBackupResponseBodyBackup) SetBackupType(v string) *DescribeBackupResponseBodyBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupResponseBodyBackup) SetBackupId(v string) *DescribeBackupResponseBodyBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupResponseBodyBackup) SetDetails(v string) *DescribeBackupResponseBodyBackup {
	s.Details = &v
	return s
}

func (s *DescribeBackupResponseBodyBackup) SetDataCenterId(v string) *DescribeBackupResponseBodyBackup {
	s.DataCenterId = &v
	return s
}

func (s *DescribeBackupResponseBodyBackup) SetClusterId(v string) *DescribeBackupResponseBodyBackup {
	s.ClusterId = &v
	return s
}

type DescribeBackupResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupResponse) SetHeaders(v map[string]*string) *DescribeBackupResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupResponse) SetBody(v *DescribeBackupResponseBody) *DescribeBackupResponse {
	s.Body = v
	return s
}

type DescribeBackupPlanRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
}

func (s DescribeBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanRequest) SetClusterId(v string) *DescribeBackupPlanRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupPlanRequest) SetDataCenterId(v string) *DescribeBackupPlanRequest {
	s.DataCenterId = &v
	return s
}

type DescribeBackupPlanResponseBody struct {
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BackupPlan *DescribeBackupPlanResponseBodyBackupPlan `json:"BackupPlan,omitempty" xml:"BackupPlan,omitempty" type:"Struct"`
}

func (s DescribeBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanResponseBody) SetRequestId(v string) *DescribeBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPlanResponseBody) SetBackupPlan(v *DescribeBackupPlanResponseBodyBackupPlan) *DescribeBackupPlanResponseBody {
	s.BackupPlan = v
	return s
}

type DescribeBackupPlanResponseBodyBackupPlan struct {
	Active          *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	BackupPeriod    *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	RetentionPeriod *int32  `json:"RetentionPeriod,omitempty" xml:"RetentionPeriod,omitempty"`
	CreatedTime     *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	BackupTime      *string `json:"BackupTime,omitempty" xml:"BackupTime,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId    *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
}

func (s DescribeBackupPlanResponseBodyBackupPlan) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanResponseBodyBackupPlan) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanResponseBodyBackupPlan) SetActive(v bool) *DescribeBackupPlanResponseBodyBackupPlan {
	s.Active = &v
	return s
}

func (s *DescribeBackupPlanResponseBodyBackupPlan) SetBackupPeriod(v string) *DescribeBackupPlanResponseBodyBackupPlan {
	s.BackupPeriod = &v
	return s
}

func (s *DescribeBackupPlanResponseBodyBackupPlan) SetRetentionPeriod(v int32) *DescribeBackupPlanResponseBodyBackupPlan {
	s.RetentionPeriod = &v
	return s
}

func (s *DescribeBackupPlanResponseBodyBackupPlan) SetCreatedTime(v string) *DescribeBackupPlanResponseBodyBackupPlan {
	s.CreatedTime = &v
	return s
}

func (s *DescribeBackupPlanResponseBodyBackupPlan) SetBackupTime(v string) *DescribeBackupPlanResponseBodyBackupPlan {
	s.BackupTime = &v
	return s
}

func (s *DescribeBackupPlanResponseBodyBackupPlan) SetClusterId(v string) *DescribeBackupPlanResponseBodyBackupPlan {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupPlanResponseBodyBackupPlan) SetDataCenterId(v string) *DescribeBackupPlanResponseBodyBackupPlan {
	s.DataCenterId = &v
	return s
}

type DescribeBackupPlanResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanResponse) SetHeaders(v map[string]*string) *DescribeBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPlanResponse) SetBody(v *DescribeBackupPlanResponseBody) *DescribeBackupPlanResponse {
	s.Body = v
	return s
}

type DescribeBackupPlansRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeBackupPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansRequest) SetClusterId(v string) *DescribeBackupPlansRequest {
	s.ClusterId = &v
	return s
}

type DescribeBackupPlansResponseBody struct {
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BackupPlans *DescribeBackupPlansResponseBodyBackupPlans `json:"BackupPlans,omitempty" xml:"BackupPlans,omitempty" type:"Struct"`
}

func (s DescribeBackupPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBody) SetRequestId(v string) *DescribeBackupPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPlansResponseBody) SetBackupPlans(v *DescribeBackupPlansResponseBodyBackupPlans) *DescribeBackupPlansResponseBody {
	s.BackupPlans = v
	return s
}

type DescribeBackupPlansResponseBodyBackupPlans struct {
	BackupPlan []*DescribeBackupPlansResponseBodyBackupPlansBackupPlan `json:"BackupPlan,omitempty" xml:"BackupPlan,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlansResponseBodyBackupPlans) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlans) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlans) SetBackupPlan(v []*DescribeBackupPlansResponseBodyBackupPlansBackupPlan) *DescribeBackupPlansResponseBodyBackupPlans {
	s.BackupPlan = v
	return s
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlan struct {
	Active          *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	BackupPeriod    *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	RetentionPeriod *int32  `json:"RetentionPeriod,omitempty" xml:"RetentionPeriod,omitempty"`
	CreatedTime     *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	BackupTime      *string `json:"BackupTime,omitempty" xml:"BackupTime,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId    *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlan) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetActive(v bool) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Active = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetBackupPeriod(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.BackupPeriod = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetRetentionPeriod(v int32) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.RetentionPeriod = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetCreatedTime(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.CreatedTime = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetBackupTime(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.BackupTime = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetClusterId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetDataCenterId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.DataCenterId = &v
	return s
}

type DescribeBackupPlansResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponse) SetHeaders(v map[string]*string) *DescribeBackupPlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPlansResponse) SetBody(v *DescribeBackupPlansResponseBody) *DescribeBackupPlansResponse {
	s.Body = v
	return s
}

type DescribeBackupsRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
	BackupType   *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
}

func (s DescribeBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupsRequest) SetClusterId(v string) *DescribeBackupsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupsRequest) SetDataCenterId(v string) *DescribeBackupsRequest {
	s.DataCenterId = &v
	return s
}

func (s *DescribeBackupsRequest) SetBackupType(v string) *DescribeBackupsRequest {
	s.BackupType = &v
	return s
}

type DescribeBackupsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Backups   *DescribeBackupsResponseBodyBackups `json:"Backups,omitempty" xml:"Backups,omitempty" type:"Struct"`
}

func (s DescribeBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) SetRequestId(v string) *DescribeBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetBackups(v *DescribeBackupsResponseBodyBackups) *DescribeBackupsResponseBody {
	s.Backups = v
	return s
}

type DescribeBackupsResponseBodyBackups struct {
	Backup []*DescribeBackupsResponseBodyBackupsBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Repeated"`
}

func (s DescribeBackupsResponseBodyBackups) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyBackups) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyBackups) SetBackup(v []*DescribeBackupsResponseBodyBackupsBackup) *DescribeBackupsResponseBodyBackups {
	s.Backup = v
	return s
}

type DescribeBackupsResponseBodyBackupsBackup struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	BackupType   *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupId     *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
}

func (s DescribeBackupsResponseBodyBackupsBackup) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyBackupsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetEndTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetStatus(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.Status = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetStartTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetSize(v int64) *DescribeBackupsResponseBodyBackupsBackup {
	s.Size = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupType(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupId(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetClusterId(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetDataCenterId(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.DataCenterId = &v
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

type DescribeClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterRequest) SetClusterId(v string) *DescribeClusterRequest {
	s.ClusterId = &v
	return s
}

type DescribeClusterResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Cluster   *DescribeClusterResponseBodyCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Struct"`
}

func (s DescribeClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBody) SetRequestId(v string) *DescribeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterResponseBody) SetCluster(v *DescribeClusterResponseBodyCluster) *DescribeClusterResponseBody {
	s.Cluster = v
	return s
}

type DescribeClusterResponseBodyCluster struct {
	Status            *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	ExpireTime        *string                                 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	MaintainStartTime *string                                 `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	PayType           *string                                 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	MaintainEndTime   *string                                 `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	Tags              *DescribeClusterResponseBodyClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	LockMode          *string                                 `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MinorVersion      *string                                 `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	AutoRenewPeriod   *int32                                  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	IsLatestVersion   *bool                                   `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	DataCenterCount   *int32                                  `json:"DataCenterCount,omitempty" xml:"DataCenterCount,omitempty"`
	AutoRenewal       *bool                                   `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	ResourceGroupId   *string                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ClusterName       *string                                 `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	MajorVersion      *string                                 `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	CreatedTime       *string                                 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	ClusterId         *string                                 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterResponseBodyCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyCluster) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyCluster) SetStatus(v string) *DescribeClusterResponseBodyCluster {
	s.Status = &v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetExpireTime(v string) *DescribeClusterResponseBodyCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetMaintainStartTime(v string) *DescribeClusterResponseBodyCluster {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetPayType(v string) *DescribeClusterResponseBodyCluster {
	s.PayType = &v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetMaintainEndTime(v string) *DescribeClusterResponseBodyCluster {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetTags(v *DescribeClusterResponseBodyClusterTags) *DescribeClusterResponseBodyCluster {
	s.Tags = v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetLockMode(v string) *DescribeClusterResponseBodyCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetMinorVersion(v string) *DescribeClusterResponseBodyCluster {
	s.MinorVersion = &v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetAutoRenewPeriod(v int32) *DescribeClusterResponseBodyCluster {
	s.AutoRenewPeriod = &v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetIsLatestVersion(v bool) *DescribeClusterResponseBodyCluster {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetDataCenterCount(v int32) *DescribeClusterResponseBodyCluster {
	s.DataCenterCount = &v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetAutoRenewal(v bool) *DescribeClusterResponseBodyCluster {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetResourceGroupId(v string) *DescribeClusterResponseBodyCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetClusterName(v string) *DescribeClusterResponseBodyCluster {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetMajorVersion(v string) *DescribeClusterResponseBodyCluster {
	s.MajorVersion = &v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetCreatedTime(v string) *DescribeClusterResponseBodyCluster {
	s.CreatedTime = &v
	return s
}

func (s *DescribeClusterResponseBodyCluster) SetClusterId(v string) *DescribeClusterResponseBodyCluster {
	s.ClusterId = &v
	return s
}

type DescribeClusterResponseBodyClusterTags struct {
	Tag []*DescribeClusterResponseBodyClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeClusterResponseBodyClusterTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterTags) SetTag(v []*DescribeClusterResponseBodyClusterTagsTag) *DescribeClusterResponseBodyClusterTags {
	s.Tag = v
	return s
}

type DescribeClusterResponseBodyClusterTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeClusterResponseBodyClusterTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterTagsTag) SetKey(v string) *DescribeClusterResponseBodyClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterTagsTag) SetValue(v string) *DescribeClusterResponseBodyClusterTagsTag {
	s.Value = &v
	return s
}

type DescribeClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponse) SetHeaders(v map[string]*string) *DescribeClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterResponse) SetBody(v *DescribeClusterResponseBody) *DescribeClusterResponse {
	s.Body = v
	return s
}

type DescribeClusterDashboardRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterDashboardRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDashboardRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterDashboardRequest) SetClusterId(v string) *DescribeClusterDashboardRequest {
	s.ClusterId = &v
	return s
}

type DescribeClusterDashboardResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Dashboard *DescribeClusterDashboardResponseBodyDashboard `json:"Dashboard,omitempty" xml:"Dashboard,omitempty" type:"Struct"`
}

func (s DescribeClusterDashboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterDashboardResponseBody) SetRequestId(v string) *DescribeClusterDashboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterDashboardResponseBody) SetDashboard(v *DescribeClusterDashboardResponseBodyDashboard) *DescribeClusterDashboardResponseBody {
	s.Dashboard = v
	return s
}

type DescribeClusterDashboardResponseBodyDashboard struct {
	DataCenters *DescribeClusterDashboardResponseBodyDashboardDataCenters `json:"DataCenters,omitempty" xml:"DataCenters,omitempty" type:"Struct"`
	ClusterId   *string                                                   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterDashboardResponseBodyDashboard) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDashboardResponseBodyDashboard) GoString() string {
	return s.String()
}

func (s *DescribeClusterDashboardResponseBodyDashboard) SetDataCenters(v *DescribeClusterDashboardResponseBodyDashboardDataCenters) *DescribeClusterDashboardResponseBodyDashboard {
	s.DataCenters = v
	return s
}

func (s *DescribeClusterDashboardResponseBodyDashboard) SetClusterId(v string) *DescribeClusterDashboardResponseBodyDashboard {
	s.ClusterId = &v
	return s
}

type DescribeClusterDashboardResponseBodyDashboardDataCenters struct {
	DataCenter []*DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenter `json:"DataCenter,omitempty" xml:"DataCenter,omitempty" type:"Repeated"`
}

func (s DescribeClusterDashboardResponseBodyDashboardDataCenters) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDashboardResponseBodyDashboardDataCenters) GoString() string {
	return s.String()
}

func (s *DescribeClusterDashboardResponseBodyDashboardDataCenters) SetDataCenter(v []*DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenter) *DescribeClusterDashboardResponseBodyDashboardDataCenters {
	s.DataCenter = v
	return s
}

type DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenter struct {
	Nodes        *DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	DataCenterId *string                                                                  `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
}

func (s DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenter) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenter) GoString() string {
	return s.String()
}

func (s *DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenter) SetNodes(v *DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodes) *DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenter {
	s.Nodes = v
	return s
}

func (s *DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenter) SetDataCenterId(v string) *DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenter {
	s.DataCenterId = &v
	return s
}

type DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodes struct {
	Node []*DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Repeated"`
}

func (s DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodes) GoString() string {
	return s.String()
}

func (s *DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodes) SetNode(v []*DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodesNode) *DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodes {
	s.Node = v
	return s
}

type DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodesNode struct {
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Load    *string `json:"Load,omitempty" xml:"Load,omitempty"`
}

func (s DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodesNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodesNode) GoString() string {
	return s.String()
}

func (s *DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodesNode) SetStatus(v string) *DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodesNode {
	s.Status = &v
	return s
}

func (s *DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodesNode) SetAddress(v string) *DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodesNode {
	s.Address = &v
	return s
}

func (s *DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodesNode) SetLoad(v string) *DescribeClusterDashboardResponseBodyDashboardDataCentersDataCenterNodesNode {
	s.Load = &v
	return s
}

type DescribeClusterDashboardResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterDashboardResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDashboardResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterDashboardResponse) SetHeaders(v map[string]*string) *DescribeClusterDashboardResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterDashboardResponse) SetBody(v *DescribeClusterDashboardResponseBody) *DescribeClusterDashboardResponse {
	s.Body = v
	return s
}

type DescribeClustersRequest struct {
	RegionId        *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber      *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ClusterName     *string                       `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ResourceGroupId *string                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*DescribeClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeClustersRequest) SetRegionId(v string) *DescribeClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClustersRequest) SetPageNumber(v int32) *DescribeClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClustersRequest) SetPageSize(v int32) *DescribeClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClustersRequest) SetClusterName(v string) *DescribeClustersRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeClustersRequest) SetResourceGroupId(v string) *DescribeClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClustersRequest) SetTag(v []*DescribeClustersRequestTag) *DescribeClustersRequest {
	s.Tag = v
	return s
}

type DescribeClustersRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeClustersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeClustersRequestTag) SetKey(v string) *DescribeClustersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeClustersRequestTag) SetValue(v string) *DescribeClustersRequestTag {
	s.Value = &v
	return s
}

type DescribeClustersResponseBody struct {
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Clusters   *DescribeClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
}

func (s DescribeClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBody) SetTotalCount(v int64) *DescribeClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeClustersResponseBody) SetPageSize(v int32) *DescribeClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeClustersResponseBody) SetRequestId(v string) *DescribeClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetPageNumber(v int32) *DescribeClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeClustersResponseBody) SetClusters(v *DescribeClustersResponseBodyClusters) *DescribeClustersResponseBody {
	s.Clusters = v
	return s
}

type DescribeClustersResponseBodyClusters struct {
	Cluster []*DescribeClustersResponseBodyClustersCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s DescribeClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyClusters) SetCluster(v []*DescribeClustersResponseBodyClustersCluster) *DescribeClustersResponseBodyClusters {
	s.Cluster = v
	return s
}

type DescribeClustersResponseBodyClustersCluster struct {
	Status          *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	ExpireTime      *string                                          `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	PayType         *string                                          `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Tags            *DescribeClustersResponseBodyClustersClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	LockMode        *string                                          `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	AutoRenewPeriod *int32                                           `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	MinorVersion    *string                                          `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	DataCenterCount *int32                                           `json:"DataCenterCount,omitempty" xml:"DataCenterCount,omitempty"`
	AutoRenewal     *bool                                            `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	ResourceGroupId *string                                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ClusterName     *string                                          `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	MajorVersion    *string                                          `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	CreatedTime     *string                                          `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	ClusterId       *string                                          `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClustersResponseBodyClustersCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBodyClustersCluster) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyClustersCluster) SetStatus(v string) *DescribeClustersResponseBodyClustersCluster {
	s.Status = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersCluster) SetExpireTime(v string) *DescribeClustersResponseBodyClustersCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersCluster) SetPayType(v string) *DescribeClustersResponseBodyClustersCluster {
	s.PayType = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersCluster) SetTags(v *DescribeClustersResponseBodyClustersClusterTags) *DescribeClustersResponseBodyClustersCluster {
	s.Tags = v
	return s
}

func (s *DescribeClustersResponseBodyClustersCluster) SetLockMode(v string) *DescribeClustersResponseBodyClustersCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersCluster) SetAutoRenewPeriod(v int32) *DescribeClustersResponseBodyClustersCluster {
	s.AutoRenewPeriod = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersCluster) SetMinorVersion(v string) *DescribeClustersResponseBodyClustersCluster {
	s.MinorVersion = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersCluster) SetDataCenterCount(v int32) *DescribeClustersResponseBodyClustersCluster {
	s.DataCenterCount = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersCluster) SetAutoRenewal(v bool) *DescribeClustersResponseBodyClustersCluster {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersCluster) SetResourceGroupId(v string) *DescribeClustersResponseBodyClustersCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersCluster) SetClusterName(v string) *DescribeClustersResponseBodyClustersCluster {
	s.ClusterName = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersCluster) SetMajorVersion(v string) *DescribeClustersResponseBodyClustersCluster {
	s.MajorVersion = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersCluster) SetCreatedTime(v string) *DescribeClustersResponseBodyClustersCluster {
	s.CreatedTime = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersCluster) SetClusterId(v string) *DescribeClustersResponseBodyClustersCluster {
	s.ClusterId = &v
	return s
}

type DescribeClustersResponseBodyClustersClusterTags struct {
	Tag []*DescribeClustersResponseBodyClustersClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeClustersResponseBodyClustersClusterTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBodyClustersClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyClustersClusterTags) SetTag(v []*DescribeClustersResponseBodyClustersClusterTagsTag) *DescribeClustersResponseBodyClustersClusterTags {
	s.Tag = v
	return s
}

type DescribeClustersResponseBodyClustersClusterTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeClustersResponseBodyClustersClusterTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBodyClustersClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyClustersClusterTagsTag) SetKey(v string) *DescribeClustersResponseBodyClustersClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersClusterTagsTag) SetValue(v string) *DescribeClustersResponseBodyClustersClusterTagsTag {
	s.Value = &v
	return s
}

type DescribeClustersResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponse) SetHeaders(v map[string]*string) *DescribeClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeClustersResponse) SetBody(v *DescribeClustersResponseBody) *DescribeClustersResponse {
	s.Body = v
	return s
}

type DescribeClusterStatusRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterStatusRequest) SetClusterId(v string) *DescribeClusterStatusRequest {
	s.ClusterId = &v
	return s
}

type DescribeClusterStatusResponseBody struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterStatusResponseBody) SetStatus(v string) *DescribeClusterStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeClusterStatusResponseBody) SetCreatedTime(v string) *DescribeClusterStatusResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeClusterStatusResponseBody) SetRequestId(v string) *DescribeClusterStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClusterStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterStatusResponse) SetHeaders(v map[string]*string) *DescribeClusterStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterStatusResponse) SetBody(v *DescribeClusterStatusResponseBody) *DescribeClusterStatusResponse {
	s.Body = v
	return s
}

type DescribeContactPointsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeContactPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactPointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContactPointsRequest) SetClusterId(v string) *DescribeContactPointsRequest {
	s.ClusterId = &v
	return s
}

type DescribeContactPointsResponseBody struct {
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ContactPoints *DescribeContactPointsResponseBodyContactPoints `json:"ContactPoints,omitempty" xml:"ContactPoints,omitempty" type:"Struct"`
}

func (s DescribeContactPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactPointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContactPointsResponseBody) SetRequestId(v string) *DescribeContactPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContactPointsResponseBody) SetContactPoints(v *DescribeContactPointsResponseBodyContactPoints) *DescribeContactPointsResponseBody {
	s.ContactPoints = v
	return s
}

type DescribeContactPointsResponseBodyContactPoints struct {
	ContactPoint []*DescribeContactPointsResponseBodyContactPointsContactPoint `json:"ContactPoint,omitempty" xml:"ContactPoint,omitempty" type:"Repeated"`
}

func (s DescribeContactPointsResponseBodyContactPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactPointsResponseBodyContactPoints) GoString() string {
	return s.String()
}

func (s *DescribeContactPointsResponseBodyContactPoints) SetContactPoint(v []*DescribeContactPointsResponseBodyContactPointsContactPoint) *DescribeContactPointsResponseBodyContactPoints {
	s.ContactPoint = v
	return s
}

type DescribeContactPointsResponseBodyContactPointsContactPoint struct {
	PublicAddresses  *DescribeContactPointsResponseBodyContactPointsContactPointPublicAddresses  `json:"PublicAddresses,omitempty" xml:"PublicAddresses,omitempty" type:"Struct"`
	Port             *int32                                                                      `json:"Port,omitempty" xml:"Port,omitempty"`
	PrivateAddresses *DescribeContactPointsResponseBodyContactPointsContactPointPrivateAddresses `json:"PrivateAddresses,omitempty" xml:"PrivateAddresses,omitempty" type:"Struct"`
	DataCenterId     *string                                                                     `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
}

func (s DescribeContactPointsResponseBodyContactPointsContactPoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactPointsResponseBodyContactPointsContactPoint) GoString() string {
	return s.String()
}

func (s *DescribeContactPointsResponseBodyContactPointsContactPoint) SetPublicAddresses(v *DescribeContactPointsResponseBodyContactPointsContactPointPublicAddresses) *DescribeContactPointsResponseBodyContactPointsContactPoint {
	s.PublicAddresses = v
	return s
}

func (s *DescribeContactPointsResponseBodyContactPointsContactPoint) SetPort(v int32) *DescribeContactPointsResponseBodyContactPointsContactPoint {
	s.Port = &v
	return s
}

func (s *DescribeContactPointsResponseBodyContactPointsContactPoint) SetPrivateAddresses(v *DescribeContactPointsResponseBodyContactPointsContactPointPrivateAddresses) *DescribeContactPointsResponseBodyContactPointsContactPoint {
	s.PrivateAddresses = v
	return s
}

func (s *DescribeContactPointsResponseBodyContactPointsContactPoint) SetDataCenterId(v string) *DescribeContactPointsResponseBodyContactPointsContactPoint {
	s.DataCenterId = &v
	return s
}

type DescribeContactPointsResponseBodyContactPointsContactPointPublicAddresses struct {
	PublicAddress []*string `json:"PublicAddress,omitempty" xml:"PublicAddress,omitempty" type:"Repeated"`
}

func (s DescribeContactPointsResponseBodyContactPointsContactPointPublicAddresses) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactPointsResponseBodyContactPointsContactPointPublicAddresses) GoString() string {
	return s.String()
}

func (s *DescribeContactPointsResponseBodyContactPointsContactPointPublicAddresses) SetPublicAddress(v []*string) *DescribeContactPointsResponseBodyContactPointsContactPointPublicAddresses {
	s.PublicAddress = v
	return s
}

type DescribeContactPointsResponseBodyContactPointsContactPointPrivateAddresses struct {
	PrivateAddress []*string `json:"PrivateAddress,omitempty" xml:"PrivateAddress,omitempty" type:"Repeated"`
}

func (s DescribeContactPointsResponseBodyContactPointsContactPointPrivateAddresses) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactPointsResponseBodyContactPointsContactPointPrivateAddresses) GoString() string {
	return s.String()
}

func (s *DescribeContactPointsResponseBodyContactPointsContactPointPrivateAddresses) SetPrivateAddress(v []*string) *DescribeContactPointsResponseBodyContactPointsContactPointPrivateAddresses {
	s.PrivateAddress = v
	return s
}

type DescribeContactPointsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeContactPointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContactPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactPointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContactPointsResponse) SetHeaders(v map[string]*string) *DescribeContactPointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContactPointsResponse) SetBody(v *DescribeContactPointsResponseBody) *DescribeContactPointsResponse {
	s.Body = v
	return s
}

type DescribeDataCenterRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
}

func (s DescribeDataCenterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCenterRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataCenterRequest) SetClusterId(v string) *DescribeDataCenterRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeDataCenterRequest) SetDataCenterId(v string) *DescribeDataCenterRequest {
	s.DataCenterId = &v
	return s
}

type DescribeDataCenterResponseBody struct {
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	AutoRenewPeriod   *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	DataCenterId      *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
	CommodityInstance *string `json:"CommodityInstance,omitempty" xml:"CommodityInstance,omitempty"`
	CreatedTime       *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NodeCount         *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	ZoneId            *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ClusterId         *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PayType           *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	LockMode          *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	VswitchId         *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	DataCenterName    *string `json:"DataCenterName,omitempty" xml:"DataCenterName,omitempty"`
	DiskType          *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	AutoRenewal       *bool   `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	DiskSize          *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ExpireTime        *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceType      *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeDataCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCenterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataCenterResponseBody) SetStatus(v string) *DescribeDataCenterResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetAutoRenewPeriod(v int32) *DescribeDataCenterResponseBody {
	s.AutoRenewPeriod = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetDataCenterId(v string) *DescribeDataCenterResponseBody {
	s.DataCenterId = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetCommodityInstance(v string) *DescribeDataCenterResponseBody {
	s.CommodityInstance = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetCreatedTime(v string) *DescribeDataCenterResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetRequestId(v string) *DescribeDataCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetNodeCount(v int32) *DescribeDataCenterResponseBody {
	s.NodeCount = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetZoneId(v string) *DescribeDataCenterResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetClusterId(v string) *DescribeDataCenterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetPayType(v string) *DescribeDataCenterResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetLockMode(v string) *DescribeDataCenterResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetVswitchId(v string) *DescribeDataCenterResponseBody {
	s.VswitchId = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetDataCenterName(v string) *DescribeDataCenterResponseBody {
	s.DataCenterName = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetDiskType(v string) *DescribeDataCenterResponseBody {
	s.DiskType = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetVpcId(v string) *DescribeDataCenterResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetAutoRenewal(v bool) *DescribeDataCenterResponseBody {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetDiskSize(v int32) *DescribeDataCenterResponseBody {
	s.DiskSize = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetRegionId(v string) *DescribeDataCenterResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetExpireTime(v string) *DescribeDataCenterResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDataCenterResponseBody) SetInstanceType(v string) *DescribeDataCenterResponseBody {
	s.InstanceType = &v
	return s
}

type DescribeDataCenterResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataCenterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCenterResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataCenterResponse) SetHeaders(v map[string]*string) *DescribeDataCenterResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataCenterResponse) SetBody(v *DescribeDataCenterResponseBody) *DescribeDataCenterResponse {
	s.Body = v
	return s
}

type DescribeDataCentersRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeDataCentersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCentersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataCentersRequest) SetClusterId(v string) *DescribeDataCentersRequest {
	s.ClusterId = &v
	return s
}

type DescribeDataCentersResponseBody struct {
	DataCenters *DescribeDataCentersResponseBodyDataCenters `json:"DataCenters,omitempty" xml:"DataCenters,omitempty" type:"Struct"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataCentersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCentersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataCentersResponseBody) SetDataCenters(v *DescribeDataCentersResponseBodyDataCenters) *DescribeDataCentersResponseBody {
	s.DataCenters = v
	return s
}

func (s *DescribeDataCentersResponseBody) SetRequestId(v string) *DescribeDataCentersResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDataCentersResponseBodyDataCenters struct {
	DataCenter []*DescribeDataCentersResponseBodyDataCentersDataCenter `json:"DataCenter,omitempty" xml:"DataCenter,omitempty" type:"Repeated"`
}

func (s DescribeDataCentersResponseBodyDataCenters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCentersResponseBodyDataCenters) GoString() string {
	return s.String()
}

func (s *DescribeDataCentersResponseBodyDataCenters) SetDataCenter(v []*DescribeDataCentersResponseBodyDataCentersDataCenter) *DescribeDataCentersResponseBodyDataCenters {
	s.DataCenter = v
	return s
}

type DescribeDataCentersResponseBodyDataCentersDataCenter struct {
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId         *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ExpireTime        *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	DiskSize          *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	PayType           *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DiskType          *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	InstanceType      *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	LockMode          *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	AutoRenewPeriod   *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AutoRenewal       *bool   `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	CommodityInstance *string `json:"CommodityInstance,omitempty" xml:"CommodityInstance,omitempty"`
	NodeCount         *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	DataCenterName    *string `json:"DataCenterName,omitempty" xml:"DataCenterName,omitempty"`
	ZoneId            *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	CreatedTime       *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	ClusterId         *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId      *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
}

func (s DescribeDataCentersResponseBodyDataCentersDataCenter) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCentersResponseBodyDataCentersDataCenter) GoString() string {
	return s.String()
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetStatus(v string) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.Status = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetVpcId(v string) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.VpcId = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetVswitchId(v string) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.VswitchId = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetExpireTime(v string) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetDiskSize(v int32) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.DiskSize = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetPayType(v string) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.PayType = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetDiskType(v string) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.DiskType = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetInstanceType(v string) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.InstanceType = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetLockMode(v string) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.LockMode = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetAutoRenewPeriod(v int32) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.AutoRenewPeriod = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetRegionId(v string) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.RegionId = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetAutoRenewal(v bool) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetCommodityInstance(v string) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.CommodityInstance = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetNodeCount(v int32) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.NodeCount = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetDataCenterName(v string) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.DataCenterName = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetZoneId(v string) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.ZoneId = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetCreatedTime(v string) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetClusterId(v string) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.ClusterId = &v
	return s
}

func (s *DescribeDataCentersResponseBodyDataCentersDataCenter) SetDataCenterId(v string) *DescribeDataCentersResponseBodyDataCentersDataCenter {
	s.DataCenterId = &v
	return s
}

type DescribeDataCentersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataCentersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataCentersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCentersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataCentersResponse) SetHeaders(v map[string]*string) *DescribeDataCentersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataCentersResponse) SetBody(v *DescribeDataCentersResponseBody) *DescribeDataCentersResponse {
	s.Body = v
	return s
}

type DescribeDeletedClustersRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDeletedClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeletedClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeletedClustersRequest) SetRegionId(v string) *DescribeDeletedClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDeletedClustersRequest) SetPageNumber(v int32) *DescribeDeletedClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeletedClustersRequest) SetPageSize(v int32) *DescribeDeletedClustersRequest {
	s.PageSize = &v
	return s
}

type DescribeDeletedClustersResponseBody struct {
	TotalCount *int64                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Clusters   *DescribeDeletedClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
}

func (s DescribeDeletedClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeletedClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeletedClustersResponseBody) SetTotalCount(v int64) *DescribeDeletedClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDeletedClustersResponseBody) SetPageSize(v int32) *DescribeDeletedClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDeletedClustersResponseBody) SetRequestId(v string) *DescribeDeletedClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeletedClustersResponseBody) SetPageNumber(v int32) *DescribeDeletedClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeletedClustersResponseBody) SetClusters(v *DescribeDeletedClustersResponseBodyClusters) *DescribeDeletedClustersResponseBody {
	s.Clusters = v
	return s
}

type DescribeDeletedClustersResponseBodyClusters struct {
	Cluster []*DescribeDeletedClustersResponseBodyClustersCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s DescribeDeletedClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeletedClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeDeletedClustersResponseBodyClusters) SetCluster(v []*DescribeDeletedClustersResponseBodyClustersCluster) *DescribeDeletedClustersResponseBodyClusters {
	s.Cluster = v
	return s
}

type DescribeDeletedClustersResponseBodyClustersCluster struct {
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DataCenterCount *int32  `json:"DataCenterCount,omitempty" xml:"DataCenterCount,omitempty"`
	ExpireTime      *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	PayType         *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	ClusterName     *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	MajorVersion    *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	CreatedTime     *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	MinorVersion    *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeDeletedClustersResponseBodyClustersCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeletedClustersResponseBodyClustersCluster) GoString() string {
	return s.String()
}

func (s *DescribeDeletedClustersResponseBodyClustersCluster) SetStatus(v string) *DescribeDeletedClustersResponseBodyClustersCluster {
	s.Status = &v
	return s
}

func (s *DescribeDeletedClustersResponseBodyClustersCluster) SetDataCenterCount(v int32) *DescribeDeletedClustersResponseBodyClustersCluster {
	s.DataCenterCount = &v
	return s
}

func (s *DescribeDeletedClustersResponseBodyClustersCluster) SetExpireTime(v string) *DescribeDeletedClustersResponseBodyClustersCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDeletedClustersResponseBodyClustersCluster) SetPayType(v string) *DescribeDeletedClustersResponseBodyClustersCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDeletedClustersResponseBodyClustersCluster) SetClusterName(v string) *DescribeDeletedClustersResponseBodyClustersCluster {
	s.ClusterName = &v
	return s
}

func (s *DescribeDeletedClustersResponseBodyClustersCluster) SetMajorVersion(v string) *DescribeDeletedClustersResponseBodyClustersCluster {
	s.MajorVersion = &v
	return s
}

func (s *DescribeDeletedClustersResponseBodyClustersCluster) SetCreatedTime(v string) *DescribeDeletedClustersResponseBodyClustersCluster {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDeletedClustersResponseBodyClustersCluster) SetMinorVersion(v string) *DescribeDeletedClustersResponseBodyClustersCluster {
	s.MinorVersion = &v
	return s
}

func (s *DescribeDeletedClustersResponseBodyClustersCluster) SetClusterId(v string) *DescribeDeletedClustersResponseBodyClustersCluster {
	s.ClusterId = &v
	return s
}

type DescribeDeletedClustersResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeletedClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeletedClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeletedClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeletedClustersResponse) SetHeaders(v map[string]*string) *DescribeDeletedClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeletedClustersResponse) SetBody(v *DescribeDeletedClustersResponseBody) *DescribeDeletedClustersResponse {
	s.Body = v
	return s
}

type DescribeInstanceTypeRequest struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeInstanceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeRequest) SetInstanceType(v string) *DescribeInstanceTypeRequest {
	s.InstanceType = &v
	return s
}

type DescribeInstanceTypeResponseBody struct {
	RequestId            *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceTypeSpecList *DescribeInstanceTypeResponseBodyInstanceTypeSpecList `json:"InstanceTypeSpecList,omitempty" xml:"InstanceTypeSpecList,omitempty" type:"Struct"`
}

func (s DescribeInstanceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeResponseBody) SetRequestId(v string) *DescribeInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceTypeResponseBody) SetInstanceTypeSpecList(v *DescribeInstanceTypeResponseBodyInstanceTypeSpecList) *DescribeInstanceTypeResponseBody {
	s.InstanceTypeSpecList = v
	return s
}

type DescribeInstanceTypeResponseBodyInstanceTypeSpecList struct {
	InstanceTypeSpec []*DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec `json:"InstanceTypeSpec,omitempty" xml:"InstanceTypeSpec,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTypeResponseBodyInstanceTypeSpecList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypeResponseBodyInstanceTypeSpecList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecList) SetInstanceTypeSpec(v []*DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) *DescribeInstanceTypeResponseBodyInstanceTypeSpecList {
	s.InstanceTypeSpec = v
	return s
}

type DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec struct {
	CpuSize      *int64  `json:"CpuSize,omitempty" xml:"CpuSize,omitempty"`
	MemSize      *int64  `json:"MemSize,omitempty" xml:"MemSize,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) SetCpuSize(v int64) *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec {
	s.CpuSize = &v
	return s
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) SetMemSize(v int64) *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec {
	s.MemSize = &v
	return s
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) SetInstanceType(v string) *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec {
	s.InstanceType = &v
	return s
}

type DescribeInstanceTypeResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeResponse) SetHeaders(v map[string]*string) *DescribeInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceTypeResponse) SetBody(v *DescribeInstanceTypeResponseBody) *DescribeInstanceTypeResponse {
	s.Body = v
	return s
}

type DescribeIpWhitelistRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeIpWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistRequest) SetClusterId(v string) *DescribeIpWhitelistRequest {
	s.ClusterId = &v
	return s
}

type DescribeIpWhitelistResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IpList    *DescribeIpWhitelistResponseBodyIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Struct"`
}

func (s DescribeIpWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponseBody) SetRequestId(v string) *DescribeIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpWhitelistResponseBody) SetIpList(v *DescribeIpWhitelistResponseBodyIpList) *DescribeIpWhitelistResponseBody {
	s.IpList = v
	return s
}

type DescribeIpWhitelistResponseBodyIpList struct {
	IP []*string `json:"IP,omitempty" xml:"IP,omitempty" type:"Repeated"`
}

func (s DescribeIpWhitelistResponseBodyIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistResponseBodyIpList) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponseBodyIpList) SetIP(v []*string) *DescribeIpWhitelistResponseBodyIpList {
	s.IP = v
	return s
}

type DescribeIpWhitelistResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIpWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponse) SetHeaders(v map[string]*string) *DescribeIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpWhitelistResponse) SetBody(v *DescribeIpWhitelistResponseBody) *DescribeIpWhitelistResponse {
	s.Body = v
	return s
}

type DescribeIpWhitelistGroupsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeIpWhitelistGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistGroupsRequest) SetClusterId(v string) *DescribeIpWhitelistGroupsRequest {
	s.ClusterId = &v
	return s
}

type DescribeIpWhitelistGroupsResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Groups    *DescribeIpWhitelistGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
}

func (s DescribeIpWhitelistGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistGroupsResponseBody) SetRequestId(v string) *DescribeIpWhitelistGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpWhitelistGroupsResponseBody) SetGroups(v *DescribeIpWhitelistGroupsResponseBodyGroups) *DescribeIpWhitelistGroupsResponseBody {
	s.Groups = v
	return s
}

type DescribeIpWhitelistGroupsResponseBodyGroups struct {
	Group []*DescribeIpWhitelistGroupsResponseBodyGroupsGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
}

func (s DescribeIpWhitelistGroupsResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistGroupsResponseBodyGroups) SetGroup(v []*DescribeIpWhitelistGroupsResponseBodyGroupsGroup) *DescribeIpWhitelistGroupsResponseBodyGroups {
	s.Group = v
	return s
}

type DescribeIpWhitelistGroupsResponseBodyGroupsGroup struct {
	IpVersion *int32                                                  `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	GroupName *string                                                 `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IpList    *DescribeIpWhitelistGroupsResponseBodyGroupsGroupIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Struct"`
}

func (s DescribeIpWhitelistGroupsResponseBodyGroupsGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistGroupsResponseBodyGroupsGroup) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistGroupsResponseBodyGroupsGroup) SetIpVersion(v int32) *DescribeIpWhitelistGroupsResponseBodyGroupsGroup {
	s.IpVersion = &v
	return s
}

func (s *DescribeIpWhitelistGroupsResponseBodyGroupsGroup) SetGroupName(v string) *DescribeIpWhitelistGroupsResponseBodyGroupsGroup {
	s.GroupName = &v
	return s
}

func (s *DescribeIpWhitelistGroupsResponseBodyGroupsGroup) SetIpList(v *DescribeIpWhitelistGroupsResponseBodyGroupsGroupIpList) *DescribeIpWhitelistGroupsResponseBodyGroupsGroup {
	s.IpList = v
	return s
}

type DescribeIpWhitelistGroupsResponseBodyGroupsGroupIpList struct {
	IP []*string `json:"IP,omitempty" xml:"IP,omitempty" type:"Repeated"`
}

func (s DescribeIpWhitelistGroupsResponseBodyGroupsGroupIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistGroupsResponseBodyGroupsGroupIpList) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistGroupsResponseBodyGroupsGroupIpList) SetIP(v []*string) *DescribeIpWhitelistGroupsResponseBodyGroupsGroupIpList {
	s.IP = v
	return s
}

type DescribeIpWhitelistGroupsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIpWhitelistGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIpWhitelistGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistGroupsResponse) SetHeaders(v map[string]*string) *DescribeIpWhitelistGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpWhitelistGroupsResponse) SetBody(v *DescribeIpWhitelistGroupsResponseBody) *DescribeIpWhitelistGroupsResponse {
	s.Body = v
	return s
}

type DescribeNodeToolExecutionHistoriesRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeNodeToolExecutionHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeToolExecutionHistoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeToolExecutionHistoriesRequest) SetClusterId(v string) *DescribeNodeToolExecutionHistoriesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesRequest) SetPageNumber(v int32) *DescribeNodeToolExecutionHistoriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesRequest) SetPageSize(v int32) *DescribeNodeToolExecutionHistoriesRequest {
	s.PageSize = &v
	return s
}

type DescribeNodeToolExecutionHistoriesResponseBody struct {
	TotalCount *int64                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Histories  *DescribeNodeToolExecutionHistoriesResponseBodyHistories `json:"Histories,omitempty" xml:"Histories,omitempty" type:"Struct"`
}

func (s DescribeNodeToolExecutionHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeToolExecutionHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeToolExecutionHistoriesResponseBody) SetTotalCount(v int64) *DescribeNodeToolExecutionHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesResponseBody) SetRequestId(v string) *DescribeNodeToolExecutionHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesResponseBody) SetPageSize(v int32) *DescribeNodeToolExecutionHistoriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesResponseBody) SetPageNumber(v int32) *DescribeNodeToolExecutionHistoriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesResponseBody) SetHistories(v *DescribeNodeToolExecutionHistoriesResponseBodyHistories) *DescribeNodeToolExecutionHistoriesResponseBody {
	s.Histories = v
	return s
}

type DescribeNodeToolExecutionHistoriesResponseBodyHistories struct {
	History []*DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory `json:"History,omitempty" xml:"History,omitempty" type:"Repeated"`
}

func (s DescribeNodeToolExecutionHistoriesResponseBodyHistories) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeToolExecutionHistoriesResponseBodyHistories) GoString() string {
	return s.String()
}

func (s *DescribeNodeToolExecutionHistoriesResponseBodyHistories) SetHistory(v []*DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory) *DescribeNodeToolExecutionHistoriesResponseBodyHistories {
	s.History = v
	return s
}

type DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory struct {
	Nodes        *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsEnded      *bool   `json:"IsEnded,omitempty" xml:"IsEnded,omitempty"`
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Command      *string `json:"Command,omitempty" xml:"Command,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
	Arguments    *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ModifyTime   *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory) GoString() string {
	return s.String()
}

func (s *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory) SetNodes(v string) *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory {
	s.Nodes = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory) SetErrorMessage(v string) *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory) SetIsEnded(v bool) *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory {
	s.IsEnded = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory) SetCreateTime(v int64) *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory {
	s.CreateTime = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory) SetJobId(v string) *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory {
	s.JobId = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory) SetCommand(v string) *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory {
	s.Command = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory) SetDataCenterId(v string) *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory {
	s.DataCenterId = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory) SetArguments(v string) *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory {
	s.Arguments = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory) SetRegionId(v string) *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory {
	s.RegionId = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory) SetModifyTime(v int64) *DescribeNodeToolExecutionHistoriesResponseBodyHistoriesHistory {
	s.ModifyTime = &v
	return s
}

type DescribeNodeToolExecutionHistoriesResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNodeToolExecutionHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNodeToolExecutionHistoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeToolExecutionHistoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeToolExecutionHistoriesResponse) SetHeaders(v map[string]*string) *DescribeNodeToolExecutionHistoriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeToolExecutionHistoriesResponse) SetBody(v *DescribeNodeToolExecutionHistoriesResponseBody) *DescribeNodeToolExecutionHistoriesResponse {
	s.Body = v
	return s
}

type DescribeNodeToolExecutionHistoryRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DcId      *string `json:"DcId,omitempty" xml:"DcId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DescribeNodeToolExecutionHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeToolExecutionHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeToolExecutionHistoryRequest) SetClusterId(v string) *DescribeNodeToolExecutionHistoryRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoryRequest) SetDcId(v string) *DescribeNodeToolExecutionHistoryRequest {
	s.DcId = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoryRequest) SetJobId(v string) *DescribeNodeToolExecutionHistoryRequest {
	s.JobId = &v
	return s
}

type DescribeNodeToolExecutionHistoryResponseBody struct {
	ModifyTime   *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IsEnded      *bool   `json:"IsEnded,omitempty" xml:"IsEnded,omitempty"`
	Command      *string `json:"Command,omitempty" xml:"Command,omitempty"`
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Arguments    *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Nodes        *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeNodeToolExecutionHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeToolExecutionHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeToolExecutionHistoryResponseBody) SetModifyTime(v int64) *DescribeNodeToolExecutionHistoryResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoryResponseBody) SetDataCenterId(v string) *DescribeNodeToolExecutionHistoryResponseBody {
	s.DataCenterId = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoryResponseBody) SetRequestId(v string) *DescribeNodeToolExecutionHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoryResponseBody) SetIsEnded(v bool) *DescribeNodeToolExecutionHistoryResponseBody {
	s.IsEnded = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoryResponseBody) SetCommand(v string) *DescribeNodeToolExecutionHistoryResponseBody {
	s.Command = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoryResponseBody) SetCreateTime(v int64) *DescribeNodeToolExecutionHistoryResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoryResponseBody) SetArguments(v string) *DescribeNodeToolExecutionHistoryResponseBody {
	s.Arguments = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoryResponseBody) SetRegionId(v string) *DescribeNodeToolExecutionHistoryResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoryResponseBody) SetErrorMessage(v string) *DescribeNodeToolExecutionHistoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoryResponseBody) SetNodes(v string) *DescribeNodeToolExecutionHistoryResponseBody {
	s.Nodes = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoryResponseBody) SetJobId(v string) *DescribeNodeToolExecutionHistoryResponseBody {
	s.JobId = &v
	return s
}

func (s *DescribeNodeToolExecutionHistoryResponseBody) SetResult(v string) *DescribeNodeToolExecutionHistoryResponseBody {
	s.Result = &v
	return s
}

type DescribeNodeToolExecutionHistoryResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNodeToolExecutionHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNodeToolExecutionHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeToolExecutionHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeToolExecutionHistoryResponse) SetHeaders(v map[string]*string) *DescribeNodeToolExecutionHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeToolExecutionHistoryResponse) SetBody(v *DescribeNodeToolExecutionHistoryResponseBody) *DescribeNodeToolExecutionHistoryResponse {
	s.Body = v
	return s
}

type DescribeParameterModificationHistoriesRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeParameterModificationHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterModificationHistoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoriesRequest) SetClusterId(v string) *DescribeParameterModificationHistoriesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeParameterModificationHistoriesRequest) SetPageNumber(v int32) *DescribeParameterModificationHistoriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeParameterModificationHistoriesRequest) SetPageSize(v int32) *DescribeParameterModificationHistoriesRequest {
	s.PageSize = &v
	return s
}

type DescribeParameterModificationHistoriesResponseBody struct {
	TotalCount *int64                                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Histories  *DescribeParameterModificationHistoriesResponseBodyHistories `json:"Histories,omitempty" xml:"Histories,omitempty" type:"Struct"`
}

func (s DescribeParameterModificationHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterModificationHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoriesResponseBody) SetTotalCount(v int64) *DescribeParameterModificationHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeParameterModificationHistoriesResponseBody) SetRequestId(v string) *DescribeParameterModificationHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterModificationHistoriesResponseBody) SetPageSize(v int32) *DescribeParameterModificationHistoriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeParameterModificationHistoriesResponseBody) SetPageNumber(v int32) *DescribeParameterModificationHistoriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeParameterModificationHistoriesResponseBody) SetHistories(v *DescribeParameterModificationHistoriesResponseBodyHistories) *DescribeParameterModificationHistoriesResponseBody {
	s.Histories = v
	return s
}

type DescribeParameterModificationHistoriesResponseBodyHistories struct {
	History []*DescribeParameterModificationHistoriesResponseBodyHistoriesHistory `json:"History,omitempty" xml:"History,omitempty" type:"Repeated"`
}

func (s DescribeParameterModificationHistoriesResponseBodyHistories) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterModificationHistoriesResponseBodyHistories) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoriesResponseBodyHistories) SetHistory(v []*DescribeParameterModificationHistoriesResponseBodyHistoriesHistory) *DescribeParameterModificationHistoriesResponseBodyHistories {
	s.History = v
	return s
}

type DescribeParameterModificationHistoriesResponseBodyHistoriesHistory struct {
	Time     *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	OldValue *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NewValue *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
}

func (s DescribeParameterModificationHistoriesResponseBodyHistoriesHistory) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterModificationHistoriesResponseBodyHistoriesHistory) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoriesResponseBodyHistoriesHistory) SetTime(v int64) *DescribeParameterModificationHistoriesResponseBodyHistoriesHistory {
	s.Time = &v
	return s
}

func (s *DescribeParameterModificationHistoriesResponseBodyHistoriesHistory) SetOldValue(v string) *DescribeParameterModificationHistoriesResponseBodyHistoriesHistory {
	s.OldValue = &v
	return s
}

func (s *DescribeParameterModificationHistoriesResponseBodyHistoriesHistory) SetName(v string) *DescribeParameterModificationHistoriesResponseBodyHistoriesHistory {
	s.Name = &v
	return s
}

func (s *DescribeParameterModificationHistoriesResponseBodyHistoriesHistory) SetNewValue(v string) *DescribeParameterModificationHistoriesResponseBodyHistoriesHistory {
	s.NewValue = &v
	return s
}

type DescribeParameterModificationHistoriesResponse struct {
	Headers map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeParameterModificationHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParameterModificationHistoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterModificationHistoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoriesResponse) SetHeaders(v map[string]*string) *DescribeParameterModificationHistoriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeParameterModificationHistoriesResponse) SetBody(v *DescribeParameterModificationHistoriesResponseBody) *DescribeParameterModificationHistoriesResponse {
	s.Body = v
	return s
}

type DescribeParametersRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeParametersRequest) SetClusterId(v string) *DescribeParametersRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeParametersRequest) SetPageNumber(v int32) *DescribeParametersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeParametersRequest) SetPageSize(v int32) *DescribeParametersRequest {
	s.PageSize = &v
	return s
}

type DescribeParametersResponseBody struct {
	TotalCount *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Parameters *DescribeParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBody) SetTotalCount(v int64) *DescribeParametersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeParametersResponseBody) SetParameters(v *DescribeParametersResponseBodyParameters) *DescribeParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeParametersResponseBody) SetRequestId(v string) *DescribeParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParametersResponseBody) SetPageSize(v int32) *DescribeParametersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeParametersResponseBody) SetPageNumber(v int32) *DescribeParametersResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeParametersResponseBodyParameters struct {
	Parameter []*DescribeParametersResponseBodyParametersParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Repeated"`
}

func (s DescribeParametersResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyParameters) SetParameter(v []*DescribeParametersResponseBodyParametersParameter) *DescribeParametersResponseBodyParameters {
	s.Parameter = v
	return s
}

type DescribeParametersResponseBodyParametersParameter struct {
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
	DataType      *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	DefaultValue  *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	AllowedValues *string `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty"`
}

func (s DescribeParametersResponseBodyParametersParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBodyParametersParameter) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyParametersParameter) SetValue(v string) *DescribeParametersResponseBodyParametersParameter {
	s.Value = &v
	return s
}

func (s *DescribeParametersResponseBodyParametersParameter) SetDataType(v string) *DescribeParametersResponseBodyParametersParameter {
	s.DataType = &v
	return s
}

func (s *DescribeParametersResponseBodyParametersParameter) SetDescription(v string) *DescribeParametersResponseBodyParametersParameter {
	s.Description = &v
	return s
}

func (s *DescribeParametersResponseBodyParametersParameter) SetName(v string) *DescribeParametersResponseBodyParametersParameter {
	s.Name = &v
	return s
}

func (s *DescribeParametersResponseBodyParametersParameter) SetDefaultValue(v string) *DescribeParametersResponseBodyParametersParameter {
	s.DefaultValue = &v
	return s
}

func (s *DescribeParametersResponseBodyParametersParameter) SetAllowedValues(v string) *DescribeParametersResponseBodyParametersParameter {
	s.AllowedValues = &v
	return s
}

type DescribeParametersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponse) SetHeaders(v map[string]*string) *DescribeParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeParametersResponse) SetBody(v *DescribeParametersResponseBody) *DescribeParametersResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
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

type DescribeRegionsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
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
	Zones          *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
	LocalName      *string                                        `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string                                        `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetZones(v *DescribeRegionsResponseBodyRegionsRegionZones) *DescribeRegionsResponseBodyRegionsRegion {
	s.Zones = v
	return s
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
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetId(v string) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.Id = &v
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

type DescribeSecurityGroupsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeSecurityGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsRequest) SetClusterId(v string) *DescribeSecurityGroupsRequest {
	s.ClusterId = &v
	return s
}

type DescribeSecurityGroupsResponseBody struct {
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityGroupIds *DescribeSecurityGroupsResponseBodySecurityGroupIds `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
}

func (s DescribeSecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBody) SetRequestId(v string) *DescribeSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetSecurityGroupIds(v *DescribeSecurityGroupsResponseBodySecurityGroupIds) *DescribeSecurityGroupsResponseBody {
	s.SecurityGroupIds = v
	return s
}

type DescribeSecurityGroupsResponseBodySecurityGroupIds struct {
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupIds) SetSecurityGroupId(v []*string) *DescribeSecurityGroupsResponseBodySecurityGroupIds {
	s.SecurityGroupId = v
	return s
}

type DescribeSecurityGroupsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponse) SetHeaders(v map[string]*string) *DescribeSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityGroupsResponse) SetBody(v *DescribeSecurityGroupsResponseBody) *DescribeSecurityGroupsResponse {
	s.Body = v
	return s
}

type ExecuteNodeToolRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Command      *string `json:"Command,omitempty" xml:"Command,omitempty"`
	Arguments    *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
	ExecuteNodes *string `json:"ExecuteNodes,omitempty" xml:"ExecuteNodes,omitempty"`
}

func (s ExecuteNodeToolRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteNodeToolRequest) GoString() string {
	return s.String()
}

func (s *ExecuteNodeToolRequest) SetClusterId(v string) *ExecuteNodeToolRequest {
	s.ClusterId = &v
	return s
}

func (s *ExecuteNodeToolRequest) SetCommand(v string) *ExecuteNodeToolRequest {
	s.Command = &v
	return s
}

func (s *ExecuteNodeToolRequest) SetArguments(v string) *ExecuteNodeToolRequest {
	s.Arguments = &v
	return s
}

func (s *ExecuteNodeToolRequest) SetDataCenterId(v string) *ExecuteNodeToolRequest {
	s.DataCenterId = &v
	return s
}

func (s *ExecuteNodeToolRequest) SetExecuteNodes(v string) *ExecuteNodeToolRequest {
	s.ExecuteNodes = &v
	return s
}

type ExecuteNodeToolResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteNodeToolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteNodeToolResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteNodeToolResponseBody) SetRequestId(v string) *ExecuteNodeToolResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteNodeToolResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteNodeToolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteNodeToolResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteNodeToolResponse) GoString() string {
	return s.String()
}

func (s *ExecuteNodeToolResponse) SetHeaders(v map[string]*string) *ExecuteNodeToolResponse {
	s.Headers = v
	return s
}

func (s *ExecuteNodeToolResponse) SetBody(v *ExecuteNodeToolResponseBody) *ExecuteNodeToolResponse {
	s.Body = v
	return s
}

type GetCmsUrlRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetCmsUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCmsUrlRequest) GoString() string {
	return s.String()
}

func (s *GetCmsUrlRequest) SetClusterId(v string) *GetCmsUrlRequest {
	s.ClusterId = &v
	return s
}

type GetCmsUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetCmsUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCmsUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetCmsUrlResponseBody) SetRequestId(v string) *GetCmsUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCmsUrlResponseBody) SetUrl(v string) *GetCmsUrlResponseBody {
	s.Url = &v
	return s
}

type GetCmsUrlResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCmsUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCmsUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCmsUrlResponse) GoString() string {
	return s.String()
}

func (s *GetCmsUrlResponse) SetHeaders(v map[string]*string) *GetCmsUrlResponse {
	s.Headers = v
	return s
}

func (s *GetCmsUrlResponse) SetBody(v *GetCmsUrlResponseBody) *GetCmsUrlResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	RegionId   *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	NextToken  *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag        []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagsRequest) GoString() string {
	return s.String()
}

func (s *ListTagsRequest) SetRegionId(v string) *ListTagsRequest {
	s.RegionId = &v
	return s
}

type ListTagsResponseBody struct {
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      *ListTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s ListTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBody) SetRequestId(v string) *ListTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagsResponseBody) SetTags(v *ListTagsResponseBodyTags) *ListTagsResponseBody {
	s.Tags = v
	return s
}

type ListTagsResponseBodyTags struct {
	Tag []*ListTagsResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTags) SetTag(v []*ListTagsResponseBodyTagsTag) *ListTagsResponseBodyTags {
	s.Tag = v
	return s
}

type ListTagsResponseBodyTagsTag struct {
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagsResponseBodyTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTagsTag) SetTagValue(v string) *ListTagsResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *ListTagsResponseBodyTagsTag) SetTagKey(v string) *ListTagsResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

type ListTagsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponse) GoString() string {
	return s.String()
}

func (s *ListTagsResponse) SetHeaders(v map[string]*string) *ListTagsResponse {
	s.Headers = v
	return s
}

func (s *ListTagsResponse) SetBody(v *ListTagsResponseBody) *ListTagsResponse {
	s.Body = v
	return s
}

type ModifyAccountPasswordRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Account   *string `json:"Account,omitempty" xml:"Account,omitempty"`
	Password  *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ModifyAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordRequest) SetClusterId(v string) *ModifyAccountPasswordRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetAccount(v string) *ModifyAccountPasswordRequest {
	s.Account = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetPassword(v string) *ModifyAccountPasswordRequest {
	s.Password = &v
	return s
}

type ModifyAccountPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordResponseBody) SetRequestId(v string) *ModifyAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccountPasswordResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordResponse) SetHeaders(v map[string]*string) *ModifyAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountPasswordResponse) SetBody(v *ModifyAccountPasswordResponseBody) *ModifyAccountPasswordResponse {
	s.Body = v
	return s
}

type ModifyBackupPlanRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId    *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
	BackupTime      *string `json:"BackupTime,omitempty" xml:"BackupTime,omitempty"`
	BackupPeriod    *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	RetentionPeriod *int32  `json:"RetentionPeriod,omitempty" xml:"RetentionPeriod,omitempty"`
	Active          *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
}

func (s ModifyBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanRequest) SetClusterId(v string) *ModifyBackupPlanRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyBackupPlanRequest) SetDataCenterId(v string) *ModifyBackupPlanRequest {
	s.DataCenterId = &v
	return s
}

func (s *ModifyBackupPlanRequest) SetBackupTime(v string) *ModifyBackupPlanRequest {
	s.BackupTime = &v
	return s
}

func (s *ModifyBackupPlanRequest) SetBackupPeriod(v string) *ModifyBackupPlanRequest {
	s.BackupPeriod = &v
	return s
}

func (s *ModifyBackupPlanRequest) SetRetentionPeriod(v int32) *ModifyBackupPlanRequest {
	s.RetentionPeriod = &v
	return s
}

func (s *ModifyBackupPlanRequest) SetActive(v bool) *ModifyBackupPlanRequest {
	s.Active = &v
	return s
}

type ModifyBackupPlanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanResponseBody) SetRequestId(v string) *ModifyBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBackupPlanResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanResponse) SetHeaders(v map[string]*string) *ModifyBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupPlanResponse) SetBody(v *ModifyBackupPlanResponseBody) *ModifyBackupPlanResponse {
	s.Body = v
	return s
}

type ModifyClusterRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
}

func (s ModifyClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterRequest) SetClusterId(v string) *ModifyClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterRequest) SetClusterName(v string) *ModifyClusterRequest {
	s.ClusterName = &v
	return s
}

type ModifyClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterResponseBody) SetRequestId(v string) *ModifyClusterResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterResponse) SetHeaders(v map[string]*string) *ModifyClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterResponse) SetBody(v *ModifyClusterResponseBody) *ModifyClusterResponse {
	s.Body = v
	return s
}

type ModifyDataCenterRequest struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId   *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
	DataCenterName *string `json:"DataCenterName,omitempty" xml:"DataCenterName,omitempty"`
}

func (s ModifyDataCenterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataCenterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataCenterRequest) SetClusterId(v string) *ModifyDataCenterRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyDataCenterRequest) SetDataCenterId(v string) *ModifyDataCenterRequest {
	s.DataCenterId = &v
	return s
}

func (s *ModifyDataCenterRequest) SetDataCenterName(v string) *ModifyDataCenterRequest {
	s.DataCenterName = &v
	return s
}

type ModifyDataCenterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDataCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataCenterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataCenterResponseBody) SetRequestId(v string) *ModifyDataCenterResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDataCenterResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDataCenterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDataCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataCenterResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataCenterResponse) SetHeaders(v map[string]*string) *ModifyDataCenterResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataCenterResponse) SetBody(v *ModifyDataCenterResponseBody) *ModifyDataCenterResponse {
	s.Body = v
	return s
}

type ModifyInstanceMaintainTimeRequest struct {
	ClusterId         *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	MaintainEndTime   *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
}

func (s ModifyInstanceMaintainTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeRequest) SetClusterId(v string) *ModifyInstanceMaintainTimeRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetMaintainStartTime(v string) *ModifyInstanceMaintainTimeRequest {
	s.MaintainStartTime = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetMaintainEndTime(v string) *ModifyInstanceMaintainTimeRequest {
	s.MaintainEndTime = &v
	return s
}

type ModifyInstanceMaintainTimeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceMaintainTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMaintainTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeResponseBody) SetRequestId(v string) *ModifyInstanceMaintainTimeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceMaintainTimeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceMaintainTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeResponse) SetHeaders(v map[string]*string) *ModifyInstanceMaintainTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceMaintainTimeResponse) SetBody(v *ModifyInstanceMaintainTimeResponseBody) *ModifyInstanceMaintainTimeResponse {
	s.Body = v
	return s
}

type ModifyInstanceTypeRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ModifyInstanceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTypeRequest) SetClusterId(v string) *ModifyInstanceTypeRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyInstanceTypeRequest) SetDataCenterId(v string) *ModifyInstanceTypeRequest {
	s.DataCenterId = &v
	return s
}

func (s *ModifyInstanceTypeRequest) SetInstanceType(v string) *ModifyInstanceTypeRequest {
	s.InstanceType = &v
	return s
}

type ModifyInstanceTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTypeResponseBody) SetRequestId(v string) *ModifyInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceTypeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTypeResponse) SetHeaders(v map[string]*string) *ModifyInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceTypeResponse) SetBody(v *ModifyInstanceTypeResponseBody) *ModifyInstanceTypeResponse {
	s.Body = v
	return s
}

type ModifyIpWhitelistRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	IpList    *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
}

func (s ModifyIpWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpWhitelistRequest) SetClusterId(v string) *ModifyIpWhitelistRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyIpWhitelistRequest) SetIpList(v string) *ModifyIpWhitelistRequest {
	s.IpList = &v
	return s
}

type ModifyIpWhitelistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpWhitelistResponseBody) SetRequestId(v string) *ModifyIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type ModifyIpWhitelistResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyIpWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpWhitelistResponse) SetHeaders(v map[string]*string) *ModifyIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpWhitelistResponse) SetBody(v *ModifyIpWhitelistResponseBody) *ModifyIpWhitelistResponse {
	s.Body = v
	return s
}

type ModifyIpWhitelistGroupRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	IpList    *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IpVersion *int32  `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
}

func (s ModifyIpWhitelistGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpWhitelistGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpWhitelistGroupRequest) SetClusterId(v string) *ModifyIpWhitelistGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyIpWhitelistGroupRequest) SetIpList(v string) *ModifyIpWhitelistGroupRequest {
	s.IpList = &v
	return s
}

func (s *ModifyIpWhitelistGroupRequest) SetGroupName(v string) *ModifyIpWhitelistGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyIpWhitelistGroupRequest) SetIpVersion(v int32) *ModifyIpWhitelistGroupRequest {
	s.IpVersion = &v
	return s
}

type ModifyIpWhitelistGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpWhitelistGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpWhitelistGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpWhitelistGroupResponseBody) SetRequestId(v string) *ModifyIpWhitelistGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyIpWhitelistGroupResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyIpWhitelistGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyIpWhitelistGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpWhitelistGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpWhitelistGroupResponse) SetHeaders(v map[string]*string) *ModifyIpWhitelistGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpWhitelistGroupResponse) SetBody(v *ModifyIpWhitelistGroupResponseBody) *ModifyIpWhitelistGroupResponse {
	s.Body = v
	return s
}

type ModifyParameterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyParameterRequest) GoString() string {
	return s.String()
}

func (s *ModifyParameterRequest) SetClusterId(v string) *ModifyParameterRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyParameterRequest) SetName(v string) *ModifyParameterRequest {
	s.Name = &v
	return s
}

func (s *ModifyParameterRequest) SetValue(v string) *ModifyParameterRequest {
	s.Value = &v
	return s
}

type ModifyParameterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyParameterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyParameterResponseBody) SetRequestId(v string) *ModifyParameterResponseBody {
	s.RequestId = &v
	return s
}

type ModifyParameterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyParameterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyParameterResponse) GoString() string {
	return s.String()
}

func (s *ModifyParameterResponse) SetHeaders(v map[string]*string) *ModifyParameterResponse {
	s.Headers = v
	return s
}

func (s *ModifyParameterResponse) SetBody(v *ModifyParameterResponseBody) *ModifyParameterResponse {
	s.Body = v
	return s
}

type ModifySecurityGroupsRequest struct {
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
}

func (s ModifySecurityGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupsRequest) SetClusterId(v string) *ModifySecurityGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifySecurityGroupsRequest) SetSecurityGroupIds(v string) *ModifySecurityGroupsRequest {
	s.SecurityGroupIds = &v
	return s
}

type ModifySecurityGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupsResponseBody) SetRequestId(v string) *ModifySecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ModifySecurityGroupsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySecurityGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupsResponse) SetHeaders(v map[string]*string) *ModifySecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityGroupsResponse) SetBody(v *ModifySecurityGroupsResponseBody) *ModifySecurityGroupsResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	ClusterId          *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetClusterId(v string) *MoveResourceGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetNewResourceGroupId(v string) *MoveResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type PurgeClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s PurgeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s PurgeClusterRequest) GoString() string {
	return s.String()
}

func (s *PurgeClusterRequest) SetClusterId(v string) *PurgeClusterRequest {
	s.ClusterId = &v
	return s
}

type PurgeClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PurgeClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PurgeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *PurgeClusterResponseBody) SetRequestId(v string) *PurgeClusterResponseBody {
	s.RequestId = &v
	return s
}

type PurgeClusterResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PurgeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PurgeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s PurgeClusterResponse) GoString() string {
	return s.String()
}

func (s *PurgeClusterResponse) SetHeaders(v map[string]*string) *PurgeClusterResponse {
	s.Headers = v
	return s
}

func (s *PurgeClusterResponse) SetBody(v *PurgeClusterResponseBody) *PurgeClusterResponse {
	s.Body = v
	return s
}

type RebootClusterRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
}

func (s RebootClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootClusterRequest) GoString() string {
	return s.String()
}

func (s *RebootClusterRequest) SetClusterId(v string) *RebootClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *RebootClusterRequest) SetDataCenterId(v string) *RebootClusterRequest {
	s.DataCenterId = &v
	return s
}

type RebootClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebootClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RebootClusterResponseBody) SetRequestId(v string) *RebootClusterResponseBody {
	s.RequestId = &v
	return s
}

type RebootClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RebootClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebootClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootClusterResponse) GoString() string {
	return s.String()
}

func (s *RebootClusterResponse) SetHeaders(v map[string]*string) *RebootClusterResponse {
	s.Headers = v
	return s
}

func (s *RebootClusterResponse) SetBody(v *RebootClusterResponseBody) *RebootClusterResponse {
	s.Body = v
	return s
}

type ReleasePublicContactPointsRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
}

func (s ReleasePublicContactPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicContactPointsRequest) GoString() string {
	return s.String()
}

func (s *ReleasePublicContactPointsRequest) SetClusterId(v string) *ReleasePublicContactPointsRequest {
	s.ClusterId = &v
	return s
}

func (s *ReleasePublicContactPointsRequest) SetDataCenterId(v string) *ReleasePublicContactPointsRequest {
	s.DataCenterId = &v
	return s
}

type ReleasePublicContactPointsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleasePublicContactPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicContactPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePublicContactPointsResponseBody) SetRequestId(v string) *ReleasePublicContactPointsResponseBody {
	s.RequestId = &v
	return s
}

type ReleasePublicContactPointsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleasePublicContactPointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleasePublicContactPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicContactPointsResponse) GoString() string {
	return s.String()
}

func (s *ReleasePublicContactPointsResponse) SetHeaders(v map[string]*string) *ReleasePublicContactPointsResponse {
	s.Headers = v
	return s
}

func (s *ReleasePublicContactPointsResponse) SetBody(v *ReleasePublicContactPointsResponseBody) *ReleasePublicContactPointsResponse {
	s.Body = v
	return s
}

type ResizeDiskSizeRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
	DiskSize     *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
}

func (s ResizeDiskSizeRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeDiskSizeRequest) GoString() string {
	return s.String()
}

func (s *ResizeDiskSizeRequest) SetClusterId(v string) *ResizeDiskSizeRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeDiskSizeRequest) SetDataCenterId(v string) *ResizeDiskSizeRequest {
	s.DataCenterId = &v
	return s
}

func (s *ResizeDiskSizeRequest) SetDiskSize(v int32) *ResizeDiskSizeRequest {
	s.DiskSize = &v
	return s
}

type ResizeDiskSizeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeDiskSizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeDiskSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeDiskSizeResponseBody) SetRequestId(v string) *ResizeDiskSizeResponseBody {
	s.RequestId = &v
	return s
}

type ResizeDiskSizeResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResizeDiskSizeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResizeDiskSizeResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeDiskSizeResponse) GoString() string {
	return s.String()
}

func (s *ResizeDiskSizeResponse) SetHeaders(v map[string]*string) *ResizeDiskSizeResponse {
	s.Headers = v
	return s
}

func (s *ResizeDiskSizeResponse) SetBody(v *ResizeDiskSizeResponseBody) *ResizeDiskSizeResponse {
	s.Body = v
	return s
}

type ResizeNodeCountRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DataCenterId *string `json:"DataCenterId,omitempty" xml:"DataCenterId,omitempty"`
	NodeCount    *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
}

func (s ResizeNodeCountRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeNodeCountRequest) GoString() string {
	return s.String()
}

func (s *ResizeNodeCountRequest) SetClusterId(v string) *ResizeNodeCountRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeNodeCountRequest) SetDataCenterId(v string) *ResizeNodeCountRequest {
	s.DataCenterId = &v
	return s
}

func (s *ResizeNodeCountRequest) SetNodeCount(v int32) *ResizeNodeCountRequest {
	s.NodeCount = &v
	return s
}

type ResizeNodeCountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeNodeCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeNodeCountResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeNodeCountResponseBody) SetRequestId(v string) *ResizeNodeCountResponseBody {
	s.RequestId = &v
	return s
}

type ResizeNodeCountResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResizeNodeCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResizeNodeCountResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeNodeCountResponse) GoString() string {
	return s.String()
}

func (s *ResizeNodeCountResponse) SetHeaders(v map[string]*string) *ResizeNodeCountResponse {
	s.Headers = v
	return s
}

func (s *ResizeNodeCountResponse) SetBody(v *ResizeNodeCountResponseBody) *ResizeNodeCountResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId   *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag        []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnTagResourcesRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	All        *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey     []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) SetRegionId(v string) *UnTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetTagKey(v []*string) *UnTagResourcesRequest {
	s.TagKey = v
	return s
}

type UnTagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponseBody) SetRequestId(v string) *UnTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UnTagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponse) SetHeaders(v map[string]*string) *UnTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UnTagResourcesResponse) SetBody(v *UnTagResourcesResponseBody) *UnTagResourcesResponse {
	s.Body = v
	return s
}

type UpgradeClusterVersionRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UpgradeClusterVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClusterVersionRequest) SetClusterId(v string) *UpgradeClusterVersionRequest {
	s.ClusterId = &v
	return s
}

type UpgradeClusterVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeClusterVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeClusterVersionResponseBody) SetRequestId(v string) *UpgradeClusterVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeClusterVersionResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeClusterVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeClusterVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeClusterVersionResponse) SetHeaders(v map[string]*string) *UpgradeClusterVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeClusterVersionResponse) SetBody(v *UpgradeClusterVersionResponseBody) *UpgradeClusterVersionResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cassandra"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AllocatePublicContactPointsWithOptions(request *AllocatePublicContactPointsRequest, runtime *util.RuntimeOptions) (_result *AllocatePublicContactPointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AllocatePublicContactPointsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AllocatePublicContactPoints"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocatePublicContactPoints(request *AllocatePublicContactPointsRequest) (_result *AllocatePublicContactPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocatePublicContactPointsResponse{}
	_body, _err := client.AllocatePublicContactPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBackupPlanWithOptions(request *CreateBackupPlanRequest, runtime *util.RuntimeOptions) (_result *CreateBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateBackupPlan"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBackupPlan(request *CreateBackupPlanRequest) (_result *CreateBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.CreateBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCluster"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDataCenterWithOptions(request *CreateDataCenterRequest, runtime *util.RuntimeOptions) (_result *CreateDataCenterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDataCenterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDataCenter"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataCenter(request *CreateDataCenterRequest) (_result *CreateDataCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataCenterResponse{}
	_body, _err := client.CreateDataCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBackupPlanWithOptions(request *DeleteBackupPlanRequest, runtime *util.RuntimeOptions) (_result *DeleteBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteBackupPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteBackupPlan"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBackupPlan(request *DeleteBackupPlanRequest) (_result *DeleteBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBackupPlanResponse{}
	_body, _err := client.DeleteBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCluster"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDataCenterWithOptions(request *DeleteDataCenterRequest, runtime *util.RuntimeOptions) (_result *DeleteDataCenterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDataCenterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDataCenter"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataCenter(request *DeleteDataCenterRequest) (_result *DeleteDataCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataCenterResponse{}
	_body, _err := client.DeleteDataCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNodeToolExecutionHistoryWithOptions(request *DeleteNodeToolExecutionHistoryRequest, runtime *util.RuntimeOptions) (_result *DeleteNodeToolExecutionHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteNodeToolExecutionHistoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteNodeToolExecutionHistory"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNodeToolExecutionHistory(request *DeleteNodeToolExecutionHistoryRequest) (_result *DeleteNodeToolExecutionHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNodeToolExecutionHistoryResponse{}
	_body, _err := client.DeleteNodeToolExecutionHistoryWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccounts"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeBackupWithOptions(request *DescribeBackupRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackup"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackup(request *DescribeBackupRequest) (_result *DescribeBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupResponse{}
	_body, _err := client.DescribeBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupPlanWithOptions(request *DescribeBackupPlanRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupPlan"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupPlan(request *DescribeBackupPlanRequest) (_result *DescribeBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPlanResponse{}
	_body, _err := client.DescribeBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupPlansWithOptions(request *DescribeBackupPlansRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupPlansResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupPlans"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupPlans(request *DescribeBackupPlansRequest) (_result *DescribeBackupPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPlansResponse{}
	_body, _err := client.DescribeBackupPlansWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackups"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeClusterWithOptions(request *DescribeClusterRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCluster"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCluster(request *DescribeClusterRequest) (_result *DescribeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterResponse{}
	_body, _err := client.DescribeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterDashboardWithOptions(request *DescribeClusterDashboardRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterDashboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClusterDashboardResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeClusterDashboard"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterDashboard(request *DescribeClusterDashboardRequest) (_result *DescribeClusterDashboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterDashboardResponse{}
	_body, _err := client.DescribeClusterDashboardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClustersWithOptions(request *DescribeClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClustersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeClusters"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusters(request *DescribeClustersRequest) (_result *DescribeClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClustersResponse{}
	_body, _err := client.DescribeClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterStatusWithOptions(request *DescribeClusterStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClusterStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeClusterStatus"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterStatus(request *DescribeClusterStatusRequest) (_result *DescribeClusterStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterStatusResponse{}
	_body, _err := client.DescribeClusterStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContactPointsWithOptions(request *DescribeContactPointsRequest, runtime *util.RuntimeOptions) (_result *DescribeContactPointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeContactPointsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeContactPoints"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContactPoints(request *DescribeContactPointsRequest) (_result *DescribeContactPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContactPointsResponse{}
	_body, _err := client.DescribeContactPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataCenterWithOptions(request *DescribeDataCenterRequest, runtime *util.RuntimeOptions) (_result *DescribeDataCenterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDataCenterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDataCenter"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataCenter(request *DescribeDataCenterRequest) (_result *DescribeDataCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataCenterResponse{}
	_body, _err := client.DescribeDataCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataCentersWithOptions(request *DescribeDataCentersRequest, runtime *util.RuntimeOptions) (_result *DescribeDataCentersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDataCentersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDataCenters"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataCenters(request *DescribeDataCentersRequest) (_result *DescribeDataCentersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataCentersResponse{}
	_body, _err := client.DescribeDataCentersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeletedClustersWithOptions(request *DescribeDeletedClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeDeletedClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDeletedClustersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDeletedClusters"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeletedClusters(request *DescribeDeletedClustersRequest) (_result *DescribeDeletedClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeletedClustersResponse{}
	_body, _err := client.DescribeDeletedClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceTypeWithOptions(request *DescribeInstanceTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceType"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceType(request *DescribeInstanceTypeRequest) (_result *DescribeInstanceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceTypeResponse{}
	_body, _err := client.DescribeInstanceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIpWhitelistWithOptions(request *DescribeIpWhitelistRequest, runtime *util.RuntimeOptions) (_result *DescribeIpWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeIpWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIpWhitelist"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpWhitelist(request *DescribeIpWhitelistRequest) (_result *DescribeIpWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpWhitelistResponse{}
	_body, _err := client.DescribeIpWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIpWhitelistGroupsWithOptions(request *DescribeIpWhitelistGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeIpWhitelistGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeIpWhitelistGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIpWhitelistGroups"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpWhitelistGroups(request *DescribeIpWhitelistGroupsRequest) (_result *DescribeIpWhitelistGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpWhitelistGroupsResponse{}
	_body, _err := client.DescribeIpWhitelistGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNodeToolExecutionHistoriesWithOptions(request *DescribeNodeToolExecutionHistoriesRequest, runtime *util.RuntimeOptions) (_result *DescribeNodeToolExecutionHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeNodeToolExecutionHistoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNodeToolExecutionHistories"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNodeToolExecutionHistories(request *DescribeNodeToolExecutionHistoriesRequest) (_result *DescribeNodeToolExecutionHistoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNodeToolExecutionHistoriesResponse{}
	_body, _err := client.DescribeNodeToolExecutionHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNodeToolExecutionHistoryWithOptions(request *DescribeNodeToolExecutionHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeNodeToolExecutionHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeNodeToolExecutionHistoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNodeToolExecutionHistory"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNodeToolExecutionHistory(request *DescribeNodeToolExecutionHistoryRequest) (_result *DescribeNodeToolExecutionHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNodeToolExecutionHistoryResponse{}
	_body, _err := client.DescribeNodeToolExecutionHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeParameterModificationHistoriesWithOptions(request *DescribeParameterModificationHistoriesRequest, runtime *util.RuntimeOptions) (_result *DescribeParameterModificationHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeParameterModificationHistoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeParameterModificationHistories"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeParameterModificationHistories(request *DescribeParameterModificationHistoriesRequest) (_result *DescribeParameterModificationHistoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParameterModificationHistoriesResponse{}
	_body, _err := client.DescribeParameterModificationHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeParametersWithOptions(request *DescribeParametersRequest, runtime *util.RuntimeOptions) (_result *DescribeParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeParametersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeParameters"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeParameters(request *DescribeParametersRequest) (_result *DescribeParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParametersResponse{}
	_body, _err := client.DescribeParametersWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeSecurityGroupsWithOptions(request *DescribeSecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSecurityGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityGroups"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (_result *DescribeSecurityGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityGroupsResponse{}
	_body, _err := client.DescribeSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteNodeToolWithOptions(request *ExecuteNodeToolRequest, runtime *util.RuntimeOptions) (_result *ExecuteNodeToolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteNodeToolResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecuteNodeTool"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteNodeTool(request *ExecuteNodeToolRequest) (_result *ExecuteNodeToolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteNodeToolResponse{}
	_body, _err := client.ExecuteNodeToolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCmsUrlWithOptions(request *GetCmsUrlRequest, runtime *util.RuntimeOptions) (_result *GetCmsUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCmsUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCmsUrl"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCmsUrl(request *GetCmsUrlRequest) (_result *GetCmsUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCmsUrlResponse{}
	_body, _err := client.GetCmsUrlWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListTagsWithOptions(request *ListTagsRequest, runtime *util.RuntimeOptions) (_result *ListTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTags"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTags(request *ListTagsRequest) (_result *ListTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagsResponse{}
	_body, _err := client.ListTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccountPasswordWithOptions(request *ModifyAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAccountPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAccountPassword"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountPassword(request *ModifyAccountPasswordRequest) (_result *ModifyAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountPasswordResponse{}
	_body, _err := client.ModifyAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBackupPlanWithOptions(request *ModifyBackupPlanRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBackupPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBackupPlan"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBackupPlan(request *ModifyBackupPlanRequest) (_result *ModifyBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupPlanResponse{}
	_body, _err := client.ModifyBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterWithOptions(request *ModifyClusterRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCluster"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCluster(request *ModifyClusterRequest) (_result *ModifyClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterResponse{}
	_body, _err := client.ModifyClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDataCenterWithOptions(request *ModifyDataCenterRequest, runtime *util.RuntimeOptions) (_result *ModifyDataCenterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDataCenterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDataCenter"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDataCenter(request *ModifyDataCenterRequest) (_result *ModifyDataCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDataCenterResponse{}
	_body, _err := client.ModifyDataCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceMaintainTimeWithOptions(request *ModifyInstanceMaintainTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceMaintainTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceMaintainTime"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceMaintainTime(request *ModifyInstanceMaintainTimeRequest) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceMaintainTimeResponse{}
	_body, _err := client.ModifyInstanceMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceTypeWithOptions(request *ModifyInstanceTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceType"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceType(request *ModifyInstanceTypeRequest) (_result *ModifyInstanceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceTypeResponse{}
	_body, _err := client.ModifyInstanceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyIpWhitelistWithOptions(request *ModifyIpWhitelistRequest, runtime *util.RuntimeOptions) (_result *ModifyIpWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyIpWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyIpWhitelist"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyIpWhitelist(request *ModifyIpWhitelistRequest) (_result *ModifyIpWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyIpWhitelistResponse{}
	_body, _err := client.ModifyIpWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyIpWhitelistGroupWithOptions(request *ModifyIpWhitelistGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyIpWhitelistGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyIpWhitelistGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyIpWhitelistGroup"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyIpWhitelistGroup(request *ModifyIpWhitelistGroupRequest) (_result *ModifyIpWhitelistGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyIpWhitelistGroupResponse{}
	_body, _err := client.ModifyIpWhitelistGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyParameterWithOptions(request *ModifyParameterRequest, runtime *util.RuntimeOptions) (_result *ModifyParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyParameterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyParameter"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyParameter(request *ModifyParameterRequest) (_result *ModifyParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyParameterResponse{}
	_body, _err := client.ModifyParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySecurityGroupsWithOptions(request *ModifySecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *ModifySecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySecurityGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySecurityGroups"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySecurityGroups(request *ModifySecurityGroupsRequest) (_result *ModifySecurityGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySecurityGroupsResponse{}
	_body, _err := client.ModifySecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MoveResourceGroup"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PurgeClusterWithOptions(request *PurgeClusterRequest, runtime *util.RuntimeOptions) (_result *PurgeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PurgeClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PurgeCluster"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PurgeCluster(request *PurgeClusterRequest) (_result *PurgeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PurgeClusterResponse{}
	_body, _err := client.PurgeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebootClusterWithOptions(request *RebootClusterRequest, runtime *util.RuntimeOptions) (_result *RebootClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RebootClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RebootCluster"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebootCluster(request *RebootClusterRequest) (_result *RebootClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootClusterResponse{}
	_body, _err := client.RebootClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleasePublicContactPointsWithOptions(request *ReleasePublicContactPointsRequest, runtime *util.RuntimeOptions) (_result *ReleasePublicContactPointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleasePublicContactPointsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleasePublicContactPoints"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleasePublicContactPoints(request *ReleasePublicContactPointsRequest) (_result *ReleasePublicContactPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleasePublicContactPointsResponse{}
	_body, _err := client.ReleasePublicContactPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResizeDiskSizeWithOptions(request *ResizeDiskSizeRequest, runtime *util.RuntimeOptions) (_result *ResizeDiskSizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResizeDiskSizeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResizeDiskSize"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResizeDiskSize(request *ResizeDiskSizeRequest) (_result *ResizeDiskSizeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeDiskSizeResponse{}
	_body, _err := client.ResizeDiskSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResizeNodeCountWithOptions(request *ResizeNodeCountRequest, runtime *util.RuntimeOptions) (_result *ResizeNodeCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResizeNodeCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResizeNodeCount"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResizeNodeCount(request *ResizeNodeCountRequest) (_result *ResizeNodeCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeNodeCountResponse{}
	_body, _err := client.ResizeNodeCountWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, runtime *util.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnTagResources"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.UnTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeClusterVersionWithOptions(request *UpgradeClusterVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeClusterVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeClusterVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeClusterVersion"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeClusterVersion(request *UpgradeClusterVersionRequest) (_result *UpgradeClusterVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeClusterVersionResponse{}
	_body, _err := client.UpgradeClusterVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

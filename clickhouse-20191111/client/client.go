// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AllocateClusterPublicConnectionRequest struct {
	// The prefix of the endpoint that is used to connect to the database. Set the value to the cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
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
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateClusterPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CheckClickhouseToRDSRequest struct {
	// The password of the account that is used to log on to the database in the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Aa
	CkPassword *string `json:"CkPassword,omitempty" xml:"CkPassword,omitempty"`
	// The account that is used to log on to the database in the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// user1
	CkUserName *string `json:"CkUserName,omitempty" xml:"CkUserName,omitempty"`
	// The port number of the ApsaraDB for ClickHouse cluster.
	//
	// example:
	//
	// 8123
	ClickhousePort *int64 `json:"ClickhousePort,omitempty" xml:"ClickhousePort,omitempty"`
	// The ID of the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-2zeyy362b5sbm****
	DbClusterId  *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp13v4bnwlu8j****
	RdsId *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	// The password of the account that is used to log on to the database in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Rr
	RdsPassword *string `json:"RdsPassword,omitempty" xml:"RdsPassword,omitempty"`
	// The port number of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// 3306
	RdsPort *int64 `json:"RdsPort,omitempty" xml:"RdsPort,omitempty"`
	// The account that is used to log on to the database in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// user2
	RdsUserName *string `json:"RdsUserName,omitempty" xml:"RdsUserName,omitempty"`
	// The ID of the VPC in which the ApsaraDB RDS for MySQL instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-wz9mm0qka0winfl47****
	RdsVpcId *string `json:"RdsVpcId,omitempty" xml:"RdsVpcId,omitempty"`
	// The internal endpoint of the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp16t9h3999xb0a7****.mysql.rds.aliyuncs.com
	RdsVpcUrl            *string `json:"RdsVpcUrl,omitempty" xml:"RdsVpcUrl,omitempty"`
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

func (s *CheckClickhouseToRDSRequest) SetRdsVpcUrl(v string) *CheckClickhouseToRDSRequest {
	s.RdsVpcUrl = &v
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
	// 	- When the value **true*	- is returned for the **Status*	- parameter, the system does not return the ErrorCode parameter.
	//
	// 	- When the value **false*	- is returned for the **Status*	- parameter, the system returns for the ErrorCode parameter the reason why the ApsaraDB for ClickHouse cluster cannot be connected to the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// NotSameVpc
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A82758F8-E793-5610-BE11-0E46664305C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the ApsaraDB for ClickHouse cluster can be connected to the ApsaraDB RDS for MySQL instance.
	//
	// 	- **true**: The ApsaraDB for ClickHouse cluster can be connected to the ApsaraDB RDS for MySQL instance.
	//
	// 	- **false**: The ApsaraDB for ClickHouse cluster cannot be connected to the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// false
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckClickhouseToRDSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CheckClickhouseToRDSResponse) SetStatusCode(v int32) *CheckClickhouseToRDSResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckClickhouseToRDSResponse) SetBody(v *CheckClickhouseToRDSResponseBody) *CheckClickhouseToRDSResponse {
	s.Body = v
	return s
}

type CheckModifyConfigNeedRestartRequest struct {
	// The configuration parameters whose settings are modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// <?xml version="1.0"?>
	//
	// <yandex>
	//
	//     <keep_alive_timeout>400</keep_alive_timeout>
	//
	//     <listen_backlog>4096</listen_backlog>
	//
	//     <logger>
	//
	//         <level>debug</level>
	//
	//         <size>1000M</size>
	//
	//     </logger>
	//
	//     <mark_cache_size>5368709120</mark_cache_size>
	//
	//     <max_concurrent_queries>201</max_concurrent_queries>
	//
	//     <max_connections>4096</max_connections>
	//
	//     <max_partition_size_to_drop>0</max_partition_size_to_drop>
	//
	//     <max_table_size_to_drop>0</max_table_size_to_drop>
	//
	//     <merge_tree>
	//
	//         <max_delay_to_insert>256</max_delay_to_insert>
	//
	//         <max_part_loading_threads>auto</max_part_loading_threads>
	//
	//         <max_suspicious_broken_parts>100</max_suspicious_broken_parts>
	//
	//         <zookeeper_session_expiration_check_period>1</zookeeper_session_expiration_check_period>
	//
	//     </merge_tree>
	//
	//     <uncompressed_cache_size>1717986918</uncompressed_cache_size>
	//
	// </yandex>
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all the clusters that are deployed in a specific region. The information includes the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1tm8zf130ew****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s CheckModifyConfigNeedRestartRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckModifyConfigNeedRestartRequest) GoString() string {
	return s.String()
}

func (s *CheckModifyConfigNeedRestartRequest) SetConfig(v string) *CheckModifyConfigNeedRestartRequest {
	s.Config = &v
	return s
}

func (s *CheckModifyConfigNeedRestartRequest) SetDBClusterId(v string) *CheckModifyConfigNeedRestartRequest {
	s.DBClusterId = &v
	return s
}

type CheckModifyConfigNeedRestartResponseBody struct {
	// Indicates whether the cluster was restarted after you modified the configuration parameters. Valid values:
	//
	// 	- **true**: The cluster was restarted.
	//
	// 	- **false**: The cluster was not restarted.
	//
	// example:
	//
	// true
	NeedRestart *bool `json:"NeedRestart,omitempty" xml:"NeedRestart,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06798FEE-BEF2-5FAF-A30D-728973BBE97C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckModifyConfigNeedRestartResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckModifyConfigNeedRestartResponseBody) GoString() string {
	return s.String()
}

func (s *CheckModifyConfigNeedRestartResponseBody) SetNeedRestart(v bool) *CheckModifyConfigNeedRestartResponseBody {
	s.NeedRestart = &v
	return s
}

func (s *CheckModifyConfigNeedRestartResponseBody) SetRequestId(v string) *CheckModifyConfigNeedRestartResponseBody {
	s.RequestId = &v
	return s
}

type CheckModifyConfigNeedRestartResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckModifyConfigNeedRestartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckModifyConfigNeedRestartResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckModifyConfigNeedRestartResponse) GoString() string {
	return s.String()
}

func (s *CheckModifyConfigNeedRestartResponse) SetHeaders(v map[string]*string) *CheckModifyConfigNeedRestartResponse {
	s.Headers = v
	return s
}

func (s *CheckModifyConfigNeedRestartResponse) SetStatusCode(v int32) *CheckModifyConfigNeedRestartResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckModifyConfigNeedRestartResponse) SetBody(v *CheckModifyConfigNeedRestartResponseBody) *CheckModifyConfigNeedRestartResponse {
	s.Body = v
	return s
}

type CheckMonitorAlertRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp13s14l8498l****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckMonitorAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckMonitorAlertRequest) GoString() string {
	return s.String()
}

func (s *CheckMonitorAlertRequest) SetDBClusterId(v string) *CheckMonitorAlertRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckMonitorAlertRequest) SetOwnerAccount(v string) *CheckMonitorAlertRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckMonitorAlertRequest) SetOwnerId(v int64) *CheckMonitorAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckMonitorAlertRequest) SetRegionId(v string) *CheckMonitorAlertRequest {
	s.RegionId = &v
	return s
}

func (s *CheckMonitorAlertRequest) SetResourceOwnerAccount(v string) *CheckMonitorAlertRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckMonitorAlertRequest) SetResourceOwnerId(v int64) *CheckMonitorAlertRequest {
	s.ResourceOwnerId = &v
	return s
}

type CheckMonitorAlertResponseBody struct {
	// The parameters that are used to configure the monitoring and alerting feature.
	//
	// example:
	//
	// {   "monitor":{     "key1":"value1",     "key2":"value2"   },   "alert":{     "key1":"value1",     "key2":"value2"   } }
	Parameter *string `json:"Parameter,omitempty" xml:"Parameter,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 94F92113-FF63-5E57-8401-6FE123AD11DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the monitoring and alerting feature is enabled. Valid values:
	//
	// 	- **enable**: The monitoring and alerting feature is enabled.
	//
	// 	- **disable**: The monitoring and alerting feature is disabled.
	//
	// example:
	//
	// enable
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CheckMonitorAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckMonitorAlertResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMonitorAlertResponseBody) SetParameter(v string) *CheckMonitorAlertResponseBody {
	s.Parameter = &v
	return s
}

func (s *CheckMonitorAlertResponseBody) SetRequestId(v string) *CheckMonitorAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckMonitorAlertResponseBody) SetState(v string) *CheckMonitorAlertResponseBody {
	s.State = &v
	return s
}

type CheckMonitorAlertResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckMonitorAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckMonitorAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckMonitorAlertResponse) GoString() string {
	return s.String()
}

func (s *CheckMonitorAlertResponse) SetHeaders(v map[string]*string) *CheckMonitorAlertResponse {
	s.Headers = v
	return s
}

func (s *CheckMonitorAlertResponse) SetStatusCode(v int32) *CheckMonitorAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckMonitorAlertResponse) SetBody(v *CheckMonitorAlertResponseBody) *CheckMonitorAlertResponse {
	s.Body = v
	return s
}

type CheckScaleOutBalancedRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// The check result. Valid values:
	//
	// 	- **400**: The cluster failed the check.
	//
	// 	- **200**: The cluster passed the check.
	//
	// example:
	//
	// 400
	CheckCode *string `json:"CheckCode,omitempty" xml:"CheckCode,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error information returned for a check failure.
	TableDetails *CheckScaleOutBalancedResponseBodyTableDetails `json:"TableDetails,omitempty" xml:"TableDetails,omitempty" type:"Struct"`
	// The amount of time that is required for the migration and scale-out. Unit: minutes.
	//
	// example:
	//
	// 21
	TimeDuration *string `json:"TimeDuration,omitempty" xml:"TimeDuration,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The cluster. The value is fixed as **default**.
	//
	// example:
	//
	// default
	Cluster *string `json:"Cluster,omitempty" xml:"Cluster,omitempty"`
	// The database name.
	//
	// example:
	//
	// db_name
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The error details. Valid values:
	//
	// 	- **1**: The unique distributed table is missing.
	//
	// 	- **2**: More than one distributed table exists for the local table.
	//
	// example:
	//
	// 1
	Detail *int32 `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The name of the local table.
	//
	// example:
	//
	// test
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckScaleOutBalancedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CheckScaleOutBalancedResponse) SetStatusCode(v int32) *CheckScaleOutBalancedResponse {
	s.StatusCode = &v
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
	// The role.
	//
	// example:
	//
	// xxxx
	HasServiceLinkedRole *bool `json:"HasServiceLinkedRole,omitempty" xml:"HasServiceLinkedRole,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CheckServiceLinkedRoleResponse) SetStatusCode(v int32) *CheckServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckServiceLinkedRoleResponse) SetBody(v *CheckServiceLinkedRoleResponseBody) *CheckServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type CreateAccountRequest struct {
	// The description of the database account.
	//
	// >
	//
	// 	- The description cannot start with http:// or https://.
	//
	// 	- The description must be 0 to 256 characters in length.
	//
	// example:
	//
	// ceshi
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	//
	// >
	//
	// 	- The name must be unique in the cluster.
	//
	// 	- The name can contain lowercase letters, digits, and underscores (_).
	//
	// 	- The name must start with a lowercase letter and end with a lowercase letter or a digit.
	//
	// 	- The name must be 2 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the database account.
	//
	// >
	//
	// 	- The password must contain at least three types of the following characters: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- The password can contain the following special characters: ! @ # $ % ^ & \\	- ( ) _ + - =
	//
	// 	- The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123789Ff!
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
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
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateAccountAndAuthorityRequest struct {
	// example:
	//
	// ceshi
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456Aa
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// db1
	AllowDatabases *string `json:"AllowDatabases,omitempty" xml:"AllowDatabases,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dt1
	AllowDictionaries *string `json:"AllowDictionaries,omitempty" xml:"AllowDictionaries,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1p816075e21****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	DdlAuthority *bool `json:"DdlAuthority,omitempty" xml:"DdlAuthority,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// all
	DmlAuthority *string `json:"DmlAuthority,omitempty" xml:"DmlAuthority,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// db1,db2
	TotalDatabases *string `json:"TotalDatabases,omitempty" xml:"TotalDatabases,omitempty"`
	// example:
	//
	// dt1,dt2
	TotalDictionaries *string `json:"TotalDictionaries,omitempty" xml:"TotalDictionaries,omitempty"`
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
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccountAndAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateAccountAndAuthorityResponse) SetStatusCode(v int32) *CreateAccountAndAuthorityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccountAndAuthorityResponse) SetBody(v *CreateAccountAndAuthorityResponseBody) *CreateAccountAndAuthorityResponse {
	s.Body = v
	return s
}

type CreateBackupPolicyRequest struct {
	// The retention period for the backup data. By default, the backup data is retained for seven days. Valid values: 7 to 730. Unit: day.
	//
	// example:
	//
	// 8
	BackupRetentionPeriod *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The day of a week when the system regularly backs up data. If you specify multiple days of a week, separate them with commas (,). Valid values:
	//
	// 	- **Monday**
	//
	// 	- **Tuesday**
	//
	// 	- **Wednesday**
	//
	// 	- **Thursday**
	//
	// 	- **Friday**
	//
	// 	- **Saturday**
	//
	// 	- **Sunday**
	//
	// This parameter is required.
	//
	// example:
	//
	// Monday,Friday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The backup window. Specify the time in the ISO 8601 standard in the HH:mmZ-HH:mmZ format. The time must be in Coordinated Universal Time (UTC).
	//
	// For example, if you set the backup window to 00:00Z-01:00Z, the data of the cluster can be backed up from 08:00 (UTC+8) to 09:00 (UTC+8).
	//
	// This parameter is required.
	//
	// example:
	//
	// 10:00Z-11:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *CreateBackupPolicyRequest) SetRegionId(v string) *CreateBackupPolicyRequest {
	s.RegionId = &v
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
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateBackupPolicyResponse) SetStatusCode(v int32) *CreateBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBackupPolicyResponse) SetBody(v *CreateBackupPolicyResponseBody) *CreateBackupPolicyResponse {
	s.Body = v
	return s
}

type CreateDBInstanceRequest struct {
	// Specifies whether to enable auto-renewal.
	//
	// >  This parameter is valid only if the value of PayType is set to Prepaid.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the backup set. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/360339.html) operation to query the backup sets.
	//
	// >  If you want to restore the data of an ApsaraDB for ClickHouse cluster, this parameter is required.
	//
	// example:
	//
	// b-12af23adsf
	BackupSetID *string `json:"BackupSetID,omitempty" xml:"BackupSetID,omitempty"`
	// The client token that is used to ensure the idempotence of the request. The value is a string and can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// AB
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// 	- **Basic**: Single-replica Edition
	//
	// 	- **HighAvailability**: Double-replica Edition
	//
	// This parameter is required.
	//
	// example:
	//
	// Basic
	DBClusterCategory *string `json:"DBClusterCategory,omitempty" xml:"DBClusterCategory,omitempty"`
	// The specifications of the cluster.
	//
	// 	- Valid values when the cluster is of Single-replica Edition:
	//
	//     - **S4**: 4 CPU cores and 16 GB of memory
	//
	//     - **S8**: 8 CPU cores and 32 GB of memory
	//
	//     -  **S16**: 16 CPU cores and 64 GB of memory
	//
	//     	- **S32**: 32 CPU cores and 128 GB of memory
	//
	//     	- **S64**: 64 CPU cores and 256 GB of memory
	//
	//     	- **S104**: 104 CPU cores and 384 GB of memory
	//
	// 	- Valid values when the cluster is of Double-replica Edition:
	//
	//     - **C4**: 4 CPU cores and 16 GB of memory
	//
	//     - **C8**: 8 CPU cores and 32 GB of memory
	//
	//     - **C16**: 16 CPU cores and 64 GB of memory
	//
	//     - **C32**: 32 CPU cores and 128 GB of memory
	//
	//     - **C64**: 64 CPU cores and 256 GB of memory
	//
	//     - **C104**: 104 CPU cores and 384 GB of memory
	//
	// This parameter is required.
	//
	// example:
	//
	// S8
	DBClusterClass *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The network type of the cluster. Only Virtual Private Cloud (VPC) is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The kernel version. Valid values:
	//
	// 	- **21.8.10.19**
	//
	// 	- **22.8.5.29**
	//
	// This parameter is required.
	//
	// example:
	//
	// 21.8.10.19
	DBClusterVersion *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	// The number of nodes.
	//
	// 	- Valid values when the cluster is of Single-replica Edition: 1 to 48.
	//
	// 	- Valid values when the cluster is of Double-replica Edition: 1 to 24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DBNodeGroupCount *string `json:"DBNodeGroupCount,omitempty" xml:"DBNodeGroupCount,omitempty"`
	// The storage capacity of a single node. Valid values: 100 to 32000. Unit: GB.
	//
	// >  This value is a multiple of 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	DBNodeStorage *string `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	// The storage type of the cluster. Valid values:
	//
	// 	- **CloudESSD**: The cluster uses an enhanced SSD (ESSD) of performance level 1 (PL1).
	//
	// 	- **CloudESSD_PL2**: The cluster uses an ESSD of PL2.
	//
	// 	- **CloudESSD_PL3**: The cluster uses an ESSD of PL3.
	//
	// 	- **CloudEfficiency**: The cluster uses an ultra disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudESSD_PL2
	DbNodeStorageType *string `json:"DbNodeStorageType,omitempty" xml:"DbNodeStorageType,omitempty"`
	// You must specify this parameter when EncryptionType is set to CloudDisk.
	//
	// The ID of the key that is used to encrypt data on disks. You can obtain the ID of the key from the Key Management Service (KMS) console. You can also create a key.
	//
	// >  If EncryptionType is empty, you do not need to specify this parameter.
	//
	// example:
	//
	// 0d2470df-da7b-4786-b981-9a164dae****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The encryption type. Set the value to **CloudDisk**, which indicates that only disk encryption is supported.
	//
	// >  If this parameter is not specified, data is not encrypted.
	//
	// example:
	//
	// CloudDisk
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// 	- **Postpaid**: The cluster uses the pay-as-you-go billing method.
	//
	// 	- **Prepaid**: The cluster uses the subscription billing method.
	//
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration. This parameter is required when PayType is set to Prepaid.
	//
	// Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the cluster belongs.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the source cluster. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query backup set IDs.
	//
	// >  If you want to restore the data of an ApsaraDB for ClickHouse cluster, this parameter is required.
	//
	// example:
	//
	// cc-bp1lxbo89u950****
	SourceDBClusterId *string `json:"SourceDBClusterId,omitempty" xml:"SourceDBClusterId,omitempty"`
	// The subscription duration of the subscription cluster. This parameter is required when PayType is set to Prepaid.
	//
	// Valid values:
	//
	// 	- If Period is set to Year, the value of UsedTime must be an integer that ranges from 1 to 3.
	//
	// 	- If Period is set to Month, the value of UsedTime must be an integer that ranges from 1 to 9.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The VPC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch in the secondary zone for the VPC.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchBak *string `json:"VSwitchBak,omitempty" xml:"VSwitchBak,omitempty"`
	// The vSwitch in secondary zone 2 for the VPC.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchBak2 *string `json:"VSwitchBak2,omitempty" xml:"VSwitchBak2,omitempty"`
	// The vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Secondary zone 2.
	//
	// example:
	//
	// cn-hangzhou-j
	ZondIdBak2 *string `json:"ZondIdBak2,omitempty" xml:"ZondIdBak2,omitempty"`
	// The zone ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The secondary zone.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneIdBak *string `json:"ZoneIdBak,omitempty" xml:"ZoneIdBak,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) SetAutoRenew(v bool) *CreateDBInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBInstanceRequest) SetBackupSetID(v string) *CreateDBInstanceRequest {
	s.BackupSetID = &v
	return s
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

func (s *CreateDBInstanceRequest) SetResourceGroupId(v string) *CreateDBInstanceRequest {
	s.ResourceGroupId = &v
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

func (s *CreateDBInstanceRequest) SetSourceDBClusterId(v string) *CreateDBInstanceRequest {
	s.SourceDBClusterId = &v
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

func (s *CreateDBInstanceRequest) SetVSwitchBak(v string) *CreateDBInstanceRequest {
	s.VSwitchBak = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchBak2(v string) *CreateDBInstanceRequest {
	s.VSwitchBak2 = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchId(v string) *CreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZondIdBak2(v string) *CreateDBInstanceRequest {
	s.ZondIdBak2 = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneIdBak(v string) *CreateDBInstanceRequest {
	s.ZoneIdBak = &v
	return s
}

type CreateDBInstanceResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp1qx68m06981****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 21137950671****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateDBInstanceResponse) SetStatusCode(v int32) *CreateDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBInstanceResponse) SetBody(v *CreateDBInstanceResponseBody) *CreateDBInstanceResponse {
	s.Body = v
	return s
}

type CreateMonitorDataReportRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp13s14l8498l****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateMonitorDataReportRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorDataReportRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorDataReportRequest) SetDBClusterId(v string) *CreateMonitorDataReportRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateMonitorDataReportRequest) SetOwnerAccount(v string) *CreateMonitorDataReportRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateMonitorDataReportRequest) SetOwnerId(v int64) *CreateMonitorDataReportRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMonitorDataReportRequest) SetRegionId(v string) *CreateMonitorDataReportRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMonitorDataReportRequest) SetResourceOwnerAccount(v string) *CreateMonitorDataReportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMonitorDataReportRequest) SetResourceOwnerId(v int64) *CreateMonitorDataReportRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateMonitorDataReportResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 94F92113-FF63-5E57-8401-6FE123AD11DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMonitorDataReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorDataReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitorDataReportResponseBody) SetRequestId(v string) *CreateMonitorDataReportResponseBody {
	s.RequestId = &v
	return s
}

type CreateMonitorDataReportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMonitorDataReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorDataReportResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorDataReportResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorDataReportResponse) SetHeaders(v map[string]*string) *CreateMonitorDataReportResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorDataReportResponse) SetStatusCode(v int32) *CreateMonitorDataReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorDataReportResponse) SetBody(v *CreateMonitorDataReportResponseBody) *CreateMonitorDataReportResponse {
	s.Body = v
	return s
}

type CreateOSSStorageRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1z3a2hc8dmt****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
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
	// The request ID.
	//
	// example:
	//
	// 1F488A93-83FD-540F-9B67-0333AF64E6A0
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOSSStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateOSSStorageResponse) SetStatusCode(v int32) *CreateOSSStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOSSStorageResponse) SetBody(v *CreateOSSStorageResponseBody) *CreateOSSStorageResponse {
	s.Body = v
	return s
}

type CreatePortsForClickHouseRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port type. Set the value to mysql_port.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql_port
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
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
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePortsForClickHouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreatePortsForClickHouseResponse) SetStatusCode(v int32) *CreatePortsForClickHouseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePortsForClickHouseResponse) SetBody(v *CreatePortsForClickHouseResponseBody) *CreatePortsForClickHouseResponse {
	s.Body = v
	return s
}

type CreateRDSToClickhouseDbRequest struct {
	// The password of the account that is used to log on to the database in the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Aa
	CkPassword *string `json:"CkPassword,omitempty" xml:"CkPassword,omitempty"`
	// The account that is used to log on to the database in the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// user1
	CkUserName *string `json:"CkUserName,omitempty" xml:"CkUserName,omitempty"`
	// The port number of the ApsaraDB for ClickHouse cluster.
	//
	// example:
	//
	// 8123
	ClickhousePort *int64 `json:"ClickhousePort,omitempty" xml:"ClickhousePort,omitempty"`
	// The ID of the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-2ze5zeyl72188****
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// The maximum number of rows that can be synchronized per second.
	//
	// example:
	//
	// 50000
	LimitUpper   *int64  `json:"LimitUpper,omitempty" xml:"LimitUpper,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-8vb989qj9roh0****
	RdsId *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	// The password of the account that is used to log on to the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Rr
	RdsPassword *string `json:"RdsPassword,omitempty" xml:"RdsPassword,omitempty"`
	// The port number of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// 3306
	RdsPort *int64 `json:"RdsPort,omitempty" xml:"RdsPort,omitempty"`
	// The account that is used to log on to the database in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// user2
	RdsUserName *string `json:"RdsUserName,omitempty" xml:"RdsUserName,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the ApsaraDB RDS for MySQL instance belongs.
	//
	// example:
	//
	// vpc-2zen93xryil99jsfy****
	RdsVpcId *string `json:"RdsVpcId,omitempty" xml:"RdsVpcId,omitempty"`
	// The private endpoint of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// rm-bp16t9h3999xb0a7****.mysql.rds.aliyuncs.com
	RdsVpcUrl            *string `json:"RdsVpcUrl,omitempty" xml:"RdsVpcUrl,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to ignore the table schemas that do not support synchronization. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	SkipUnsupported *bool `json:"SkipUnsupported,omitempty" xml:"SkipUnsupported,omitempty"`
	// The tables whose data you want to synchronize.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"Schema":"recommend","Tables":["mr_platform_cpm","mr_platform_ecpm","p_monitor_record"]}]
	SynDbTables *string `json:"SynDbTables,omitempty" xml:"SynDbTables,omitempty"`
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

func (s *CreateRDSToClickhouseDbRequest) SetRdsVpcUrl(v string) *CreateRDSToClickhouseDbRequest {
	s.RdsVpcUrl = &v
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
	// If -1 is returned for the **Status*	- parameter, the cause of the creation failure is returned.
	//
	// example:
	//
	// ClickHouse exception, code: 49, host: 100.100.79.102, port: 14540; Code: 49, e.displayText() = DB::Exception: Logical error: there is no global context (version 20.8.17.25)n
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Duplicate tables in the synchronization task.
	RepeatedDbs []*string `json:"RepeatedDbs,omitempty" xml:"RepeatedDbs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 66676F54-1994-5DCF-993F-74536649628A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the synchronization task was created. Valid values:
	//
	// 	- **1**: Created.
	//
	// 	- **0**: Creation failed. The tables in the synchronization task are duplicate. The duplicate tables are returned for the **RepeatedDbs*	- parameter.
	//
	// 	- **-1**: Creation failed. The cause why the creation failed is returned for the **ErrorMsg*	- parameter.
	//
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRDSToClickhouseDbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateRDSToClickhouseDbResponse) SetStatusCode(v int32) *CreateRDSToClickhouseDbResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRDSToClickhouseDbResponse) SetBody(v *CreateRDSToClickhouseDbResponseBody) *CreateRDSToClickhouseDbResponse {
	s.Body = v
	return s
}

type CreateSQLAccountRequest struct {
	// The description of the database account.
	//
	// 	- The description cannot start with http:// or https://.
	//
	// 	- The description can be up to 256 characters in length or be an empty string.
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the database account.
	//
	// 	- The name must be unique in the cluster.
	//
	// 	- The name can contain lowercase letters, digits, or underscores (_).
	//
	// 	- The name must start with a lowercase letter and end with a lowercase letter or a digit.
	//
	// 	- The name must be 2 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the database account.
	//
	// 	- The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- Special characters include ! @ # $ % ^ & \\	- ( ) _ + - =
	//
	// 	- The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test1234
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The type of the database account. Valid values:
	//
	// 	- **Super**: privileged account.
	//
	// 	- **Normal**: standard account.
	//
	// This parameter is required.
	//
	// example:
	//
	// Super
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1p816075e21****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateSQLAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSQLAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateSQLAccountRequest) SetAccountDescription(v string) *CreateSQLAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateSQLAccountRequest) SetAccountName(v string) *CreateSQLAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateSQLAccountRequest) SetAccountPassword(v string) *CreateSQLAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateSQLAccountRequest) SetAccountType(v string) *CreateSQLAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateSQLAccountRequest) SetDBClusterId(v string) *CreateSQLAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateSQLAccountRequest) SetOwnerAccount(v string) *CreateSQLAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSQLAccountRequest) SetOwnerId(v int64) *CreateSQLAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSQLAccountRequest) SetResourceOwnerAccount(v string) *CreateSQLAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSQLAccountRequest) SetResourceOwnerId(v int64) *CreateSQLAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateSQLAccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSQLAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSQLAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSQLAccountResponseBody) SetRequestId(v string) *CreateSQLAccountResponseBody {
	s.RequestId = &v
	return s
}

type CreateSQLAccountResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSQLAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSQLAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSQLAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateSQLAccountResponse) SetHeaders(v map[string]*string) *CreateSQLAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateSQLAccountResponse) SetStatusCode(v int32) *CreateSQLAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSQLAccountResponse) SetBody(v *CreateSQLAccountResponseBody) *CreateSQLAccountResponse {
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
	// The request ID.
	//
	// example:
	//
	// 66428721-FFEC-5023-B4E5-3BD1B67837E0
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateServiceLinkedRoleResponse) SetStatusCode(v int32) *CreateServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetBody(v *CreateServiceLinkedRoleResponseBody) *CreateServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
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
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the pay-as-you-go ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
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
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DeleteSyndbRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp158i5wvj436****
	DbClusterId          *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the database in the ApsaraDB RDS for MySQL instance. The database is used for data synchronization.
	//
	// This parameter is required.
	//
	// example:
	//
	// database
	SynDb *string `json:"SynDb,omitempty" xml:"SynDb,omitempty"`
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
	// The error code.
	//
	// example:
	//
	// 0
	ErrorCode *int64 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 	- If the value **true*	- is returned for the **Status*	- parameter, the system does not return the ErrorMsg parameter.
	//
	// 	- If the value **false*	- is returned for the **Status*	- parameter, the system returns the deletion failure cause for the ErrorMsg parameter.
	//
	// example:
	//
	// ClickHouse exception, code: 49, host: 100.100.xx.xx, port: 49670; Code: 49, e.displayText() = DB::Exception: Logical error: there is no global context (version 20.8.17.25)n
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2C7393F1-5FD1-5CEE-A2EA-270A2CF99693
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the database used for data synchronization was deleted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSyndbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteSyndbResponse) SetStatusCode(v int32) *DeleteSyndbResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSyndbResponse) SetBody(v *DeleteSyndbResponseBody) *DeleteSyndbResponse {
	s.Body = v
	return s
}

type DescribeAccountAuthorityRequest struct {
	// The name of the database account.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1p816075e21****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
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
	// The name of the database account.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Databases to which permissions have been granted.
	AllowDatabases []*string `json:"AllowDatabases,omitempty" xml:"AllowDatabases,omitempty" type:"Repeated"`
	// Dictionaries to which permissions have been granted.
	AllowDictionaries []*string `json:"AllowDictionaries,omitempty" xml:"AllowDictionaries,omitempty" type:"Repeated"`
	// Indicates whether the database account has DDL permissions. Valid values:
	//
	// 	- **true**: has DDL permissions.
	//
	// 	- **false**: does not have DDL permissions.
	//
	// example:
	//
	// true
	DdlAuthority *bool `json:"DdlAuthority,omitempty" xml:"DdlAuthority,omitempty"`
	// Indicates whether the database account has DML permissions. Valid values:
	//
	// 	- **all**
	//
	// 	- **readOnly,modify**
	//
	// example:
	//
	// all
	DmlAuthority *string `json:"DmlAuthority,omitempty" xml:"DmlAuthority,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// All databases.
	TotalDatabases []*string `json:"TotalDatabases,omitempty" xml:"TotalDatabases,omitempty" type:"Repeated"`
	// All dictionaries.
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAccountAuthorityResponse) SetStatusCode(v int32) *DescribeAccountAuthorityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountAuthorityResponse) SetBody(v *DescribeAccountAuthorityResponseBody) *DescribeAccountAuthorityResponse {
	s.Body = v
	return s
}

type DescribeAccountsRequest struct {
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
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
	Accounts *DescribeAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// test
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// Creating
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// example:
	//
	// Super
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	ConfigType  *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
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

func (s *DescribeAccountsResponseBodyAccountsAccount) SetConfigType(v string) *DescribeAccountsResponseBodyAccountsAccount {
	s.ConfigType = &v
	return s
}

type DescribeAccountsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeAllDataSourceRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name.
	//
	// example:
	//
	// table
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
	// The information about the columns.
	Columns *DescribeAllDataSourceResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the databases.
	Schemas *DescribeAllDataSourceResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Struct"`
	// The information about the tables.
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
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoIncrementColumn *bool `json:"AutoIncrementColumn,omitempty" xml:"AutoIncrementColumn,omitempty"`
	// The column name.
	//
	// example:
	//
	// name
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Indicates whether the column is the primary key of the table. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	PrimaryKey *bool `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name.
	//
	// example:
	//
	// table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The type of the column.
	//
	// example:
	//
	// String
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
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
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
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name.
	//
	// example:
	//
	// table
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAllDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeAllDataSourcesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name.
	//
	// example:
	//
	// table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
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
	// Details of the columns.
	Columns *DescribeAllDataSourcesResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 75EA41D7-025A-50A6-9287-359A91399F1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the databases.
	Schemas *DescribeAllDataSourcesResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Struct"`
	// The information about the tables.
	Tables *DescribeAllDataSourcesResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
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
	// Indicates whether the column is an auto-increment column. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoIncrementColumn *bool `json:"AutoIncrementColumn,omitempty" xml:"AutoIncrementColumn,omitempty"`
	// The column name.
	//
	// example:
	//
	// name
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Indicates whether the column is the primary key of the table. Valid values:
	//
	// 	- **true**: The column is the primary key of the table.
	//
	// 	- **false**: The column is not the primary key of the table.
	//
	// example:
	//
	// true
	PrimaryKey *bool `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name.
	//
	// example:
	//
	// table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The column type.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
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
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name.
	//
	// example:
	//
	// table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAllDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAllDataSourcesResponse) SetStatusCode(v int32) *DescribeAllDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllDataSourcesResponse) SetBody(v *DescribeAllDataSourcesResponseBody) *DescribeAllDataSourcesResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1qx68m06981****
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
	// The retention period for the backup data. By default, the backup data is retained for seven days. Valid values: 7 to 730. Unit: day.
	//
	// example:
	//
	// 7
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The size of the backup data. Unit: MB.
	//
	// example:
	//
	// 123124
	BackupSize *string `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The day of a week when the system regularly backs up data. Valid values:
	//
	// 	- **Monday**
	//
	// 	- **Tuesday**
	//
	// 	- **Wednesday**
	//
	// 	- **Thursday**
	//
	// 	- **Friday**
	//
	// 	- **Saturday**
	//
	// 	- **Sunday**
	//
	// example:
	//
	// Monday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The backup window. The time is displayed in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// 15:00Z-16:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the backup feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Switch *string `json:"Switch,omitempty" xml:"Switch,omitempty"`
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

func (s *DescribeBackupPolicyResponseBody) SetBackupSize(v string) *DescribeBackupPolicyResponseBody {
	s.BackupSize = &v
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 117403****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1qx68m06981****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-25T16:00Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-21T16:00Z
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
	Items []*DescribeBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// 2021-11-22T18:28:41Z
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// example:
	//
	// 117403****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// example:
	//
	// Physical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// example:
	//
	// {"shard_count"：4}
	BackupSetInfo *string `json:"BackupSetInfo,omitempty" xml:"BackupSetInfo,omitempty"`
	// example:
	//
	// 131072
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// example:
	//
	// 2021-11-22T18:28:22Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// example:
	//
	// Success
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// example:
	//
	// IncrementalBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// example:
	//
	// cc-bp1qx68m06981****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 2022-07-22T18:28:41Z
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
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

func (s *DescribeBackupsResponseBodyItems) SetBackupSize(v int64) *DescribeBackupsResponseBodyItems {
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

func (s *DescribeBackupsResponseBodyItems) SetExpireDate(v string) *DescribeBackupsResponseBodyItems {
	s.ExpireDate = &v
	return s
}

type DescribeBackupsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-2zeux3ua25242****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The database name. You can call the [DescribeSchemas](https://help.aliyun.com/document_detail/350931.html) operation to query database names.
	//
	// This parameter is required.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name. You can call the [DescribeTables](https://help.aliyun.com/document_detail/350932.html) operation to query table names.
	//
	// This parameter is required.
	//
	// example:
	//
	// table
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
	// Details of the columns.
	Items *DescribeColumnsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 66428721-FFEC-5023-B4E5-3BD1B67837E0
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
	// Indicates whether the column is an auto-increment column. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoIncrementColumn *bool `json:"AutoIncrementColumn,omitempty" xml:"AutoIncrementColumn,omitempty"`
	// The column name.
	//
	// example:
	//
	// name
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-2zeux3ua25242****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Indicates whether the column is the primary key of the table. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	PrimaryKey *bool `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name.
	//
	// example:
	//
	// table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The column type.
	//
	// example:
	//
	// String
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeConfigHistoryRequest struct {
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all the clusters that are deployed in a specific region. The information includes the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1p816075e21****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-08-22T10:06:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in Coordinated Universal Time (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-08-11T06:27:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeConfigHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigHistoryRequest) SetDBClusterId(v string) *DescribeConfigHistoryRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeConfigHistoryRequest) SetEndTime(v string) *DescribeConfigHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeConfigHistoryRequest) SetStartTime(v string) *DescribeConfigHistoryRequest {
	s.StartTime = &v
	return s
}

type DescribeConfigHistoryResponseBody struct {
	// The change records of the configuration parameters.
	ConfigHistoryItems []*DescribeConfigHistoryResponseBodyConfigHistoryItems `json:"ConfigHistoryItems,omitempty" xml:"ConfigHistoryItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeConfigHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigHistoryResponseBody) SetConfigHistoryItems(v []*DescribeConfigHistoryResponseBodyConfigHistoryItems) *DescribeConfigHistoryResponseBody {
	s.ConfigHistoryItems = v
	return s
}

func (s *DescribeConfigHistoryResponseBody) SetRequestId(v string) *DescribeConfigHistoryResponseBody {
	s.RequestId = &v
	return s
}

type DescribeConfigHistoryResponseBodyConfigHistoryItems struct {
	// The ID of the change record.
	//
	// example:
	//
	// 1
	ChangeId *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	// The user ID (UID) of the Alibaba Cloud account.
	//
	// example:
	//
	// 253460731706911258
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The reason for the setting modification of the configuration parameters.
	//
	// example:
	//
	// test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Indicates whether the setting modification of the configuration parameters took effect. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The time when the values of the configuration parameters were changed. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-08-22T10:00:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeConfigHistoryResponseBodyConfigHistoryItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigHistoryResponseBodyConfigHistoryItems) GoString() string {
	return s.String()
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) SetChangeId(v string) *DescribeConfigHistoryResponseBodyConfigHistoryItems {
	s.ChangeId = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) SetOwnerId(v string) *DescribeConfigHistoryResponseBodyConfigHistoryItems {
	s.OwnerId = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) SetReason(v string) *DescribeConfigHistoryResponseBodyConfigHistoryItems {
	s.Reason = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) SetSuccess(v bool) *DescribeConfigHistoryResponseBodyConfigHistoryItems {
	s.Success = &v
	return s
}

func (s *DescribeConfigHistoryResponseBodyConfigHistoryItems) SetTime(v string) *DescribeConfigHistoryResponseBodyConfigHistoryItems {
	s.Time = &v
	return s
}

type DescribeConfigHistoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConfigHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConfigHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigHistoryResponse) SetHeaders(v map[string]*string) *DescribeConfigHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigHistoryResponse) SetStatusCode(v int32) *DescribeConfigHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfigHistoryResponse) SetBody(v *DescribeConfigHistoryResponseBody) *DescribeConfigHistoryResponse {
	s.Body = v
	return s
}

type DescribeConfigVersionDifferenceRequest struct {
	// The ID of the change record. You can call the [DescribeConfigHistory](https://help.aliyun.com/document_detail/452209.html) operation to query the ID of the change record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ChangeId *string `json:"ChangeId,omitempty" xml:"ChangeId,omitempty"`
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all the clusters that are deployed in a specific region. The information includes the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1tm8zf130ew****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeConfigVersionDifferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigVersionDifferenceRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigVersionDifferenceRequest) SetChangeId(v string) *DescribeConfigVersionDifferenceRequest {
	s.ChangeId = &v
	return s
}

func (s *DescribeConfigVersionDifferenceRequest) SetDBClusterId(v string) *DescribeConfigVersionDifferenceRequest {
	s.DBClusterId = &v
	return s
}

type DescribeConfigVersionDifferenceResponseBody struct {
	// The values of the configuration parameters after the values of the configuration parameters are changed.
	//
	// example:
	//
	// "<?xml version="1.0"?>
	//
	// <yandex>
	//
	//     <keep_alive_timeout>400</keep_alive_timeout>
	//
	//     <listen_backlog>4096</listen_backlog>
	//
	//     <logger>
	//
	//         <level>debug</level>
	//
	//         <size>1000M</size>
	//
	//     </logger>
	//
	//     <mark_cache_size>5368709120</mark_cache_size>
	//
	//     <max_concurrent_queries>201</max_concurrent_queries>
	//
	//     <max_connections>4096</max_connections>
	//
	//     <max_partition_size_to_drop>0</max_partition_size_to_drop>
	//
	//     <max_table_size_to_drop>0</max_table_size_to_drop>
	//
	//     <merge_tree>
	//
	//         <max_delay_to_insert>256</max_delay_to_insert>
	//
	//         <max_part_loading_threads>auto</max_part_loading_threads>
	//
	//         <max_suspicious_broken_parts>100</max_suspicious_broken_parts>
	//
	//         <zookeeper_session_expiration_check_period>1</zookeeper_session_expiration_check_period>
	//
	//     </merge_tree>
	//
	//     <uncompressed_cache_size>1717986918</uncompressed_cache_size>
	//
	// </yandex>"
	NewConfigXML *string `json:"NewConfigXML,omitempty" xml:"NewConfigXML,omitempty"`
	// The values of the configuration parameters before the values of the configuration parameters are changed.
	//
	// example:
	//
	// "<?xml version="1.0"?>
	//
	// <yandex>
	//
	//     <keep_alive_timeout>300</keep_alive_timeout>
	//
	//     <listen_backlog>4096</listen_backlog>
	//
	//     <logger>
	//
	//         <level>debug</level>
	//
	//         <size>1000M</size>
	//
	//     </logger>
	//
	//     <mark_cache_size>5368709120</mark_cache_size>
	//
	//     <max_concurrent_queries>150</max_concurrent_queries>
	//
	//     <max_connections>4096</max_connections>
	//
	//     <max_partition_size_to_drop>0</max_partition_size_to_drop>
	//
	//     <max_table_size_to_drop>0</max_table_size_to_drop>
	//
	//     <merge_tree>
	//
	//         <max_delay_to_insert>256</max_delay_to_insert>
	//
	//         <max_part_loading_threads>auto</max_part_loading_threads>
	//
	//         <max_suspicious_broken_parts>100</max_suspicious_broken_parts>
	//
	//         <zookeeper_session_expiration_check_period>1</zookeeper_session_expiration_check_period>
	//
	//     </merge_tree>
	//
	//     <uncompressed_cache_size>1717986918</uncompressed_cache_size>
	//
	// </yandex>"
	OldConfigXML *string `json:"OldConfigXML,omitempty" xml:"OldConfigXML,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeConfigVersionDifferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigVersionDifferenceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigVersionDifferenceResponseBody) SetNewConfigXML(v string) *DescribeConfigVersionDifferenceResponseBody {
	s.NewConfigXML = &v
	return s
}

func (s *DescribeConfigVersionDifferenceResponseBody) SetOldConfigXML(v string) *DescribeConfigVersionDifferenceResponseBody {
	s.OldConfigXML = &v
	return s
}

func (s *DescribeConfigVersionDifferenceResponseBody) SetRequestId(v string) *DescribeConfigVersionDifferenceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeConfigVersionDifferenceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConfigVersionDifferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConfigVersionDifferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigVersionDifferenceResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigVersionDifferenceResponse) SetHeaders(v map[string]*string) *DescribeConfigVersionDifferenceResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigVersionDifferenceResponse) SetStatusCode(v int32) *DescribeConfigVersionDifferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfigVersionDifferenceResponse) SetBody(v *DescribeConfigVersionDifferenceResponseBody) *DescribeConfigVersionDifferenceResponse {
	s.Body = v
	return s
}

type DescribeDBClusterAccessWhiteListRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
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
	// The details about the IP address whitelist.
	DBClusterAccessWhiteList *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList `json:"DBClusterAccessWhiteList,omitempty" xml:"DBClusterAccessWhiteList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 905F13A4-5097-4897-A84D-527EC75FFF4F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The attribute of the IP address whitelist.
	//
	// example:
	//
	// default
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	// The name of the IP address whitelist.
	//
	// example:
	//
	// default
	DBClusterIPArrayName *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	// The IP addresses in the IP address whitelist.
	//
	// example:
	//
	// 192.168.xx.xx,192.168.xx.xx
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all the clusters that are deployed in a specific region. The information includes the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
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
	// The information about the cluster.
	DBCluster *DescribeDBClusterAttributeResponseBodyDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 140692647406****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The scheduled restart time. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// 2023-11-06T12:00:00Z
	AppointmentRestartTime       *string                `json:"AppointmentRestartTime,omitempty" xml:"AppointmentRestartTime,omitempty"`
	AvailableUpgradeMajorVersion map[string]interface{} `json:"AvailableUpgradeMajorVersion,omitempty" xml:"AvailableUpgradeMajorVersion,omitempty"`
	// The site ID. Valid values:
	//
	// 	- **26842**: the China site (aliyun.com)
	//
	// 	- **26888**: the international site (alibabacloud.com)
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// 	- **Basic**: Single-replica Edition
	//
	// 	- **HighAvailability**: Double-replica Edition
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The commodity code of the cluster.
	//
	// example:
	//
	// clickhouse_go_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The VPC endpoint of the cluster.
	//
	// example:
	//
	// cc-bp1qx68m06981****.ads.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The version of the ApsaraDB for ClickHouse console that is used to manage the cluster. Valid values:
	//
	// 	- **v1**
	//
	// 	- **v2**
	//
	// example:
	//
	// v1
	ControlVersion *string `json:"ControlVersion,omitempty" xml:"ControlVersion,omitempty"`
	// The time when the cluster was created. The value is in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-13T11:33:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The network type of the cluster. Only VPC is supported.
	//
	// example:
	//
	// vpc
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The cluster state. Valid values:
	//
	// 	- **Preparing**: The cluster is being prepared.
	//
	// 	- **Creating**: The cluster is being created.
	//
	// 	- **Running**: The cluster is running.
	//
	// 	- **Deleting**: The cluster is being deleted.
	//
	// 	- **SCALING_OUT**: The storage capacity of the cluster is being expanded.
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The type of the cluster. Valid values:
	//
	// 	- **Common**: a common cluster
	//
	// 	- **Readonly**: a read-only cluster
	//
	// 	- **Guard**: a disaster recovery cluster
	//
	// example:
	//
	// Common
	DBClusterType *string `json:"DBClusterType,omitempty" xml:"DBClusterType,omitempty"`
	// The specifications of the cluster.
	//
	// 	- Valid values when the cluster is of Single-replica Edition:
	//
	//     	- **S4-NEW**
	//
	//     	- **S8**
	//
	//     	- **S16**
	//
	//     	- **S32**
	//
	//     	- **S64**
	//
	//     	- **S104**
	//
	// 	- Valid values when the cluster is of Double-replica Edition:
	//
	//     	- **C4-NEW**
	//
	//     	- **C8**
	//
	//     	- **C16**
	//
	//     	- **C32**
	//
	//     	- **C64**
	//
	//     	- **C104**
	//
	// example:
	//
	// C8
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of nodes.
	//
	// 	- Valid values when the cluster is of Single-replica Edition: 1 to 48.
	//
	// 	- Valid values when the cluster is of Double-replica Edition: 1 to 24.
	//
	// example:
	//
	// 1
	DBNodeCount *int64 `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	// The storage capacity of a single node of the cluster. Unit: GB.
	//
	// Valid values: 100 to 32000.
	//
	// >  This value is a multiple of 100.
	//
	// example:
	//
	// 100
	DBNodeStorage *int64 `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	// The Key Management Service (KMS) key that is used to encrypt data.
	//
	// >  If the value of the EncryptionType parameter is off, an empty string is returned for this parameter.
	//
	// example:
	//
	// 685f416f-87c9-4554-8d3a-75b6ce25****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The encryption type. Valid values:
	//
	// 	- **CloudDisk**: Disk encryption is enabled.
	//
	// 	- **off**: Data is not encrypted.
	//
	// example:
	//
	// CloudDisk
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// ClickHouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The latest minor version to which the cluster can be updated.
	//
	// example:
	//
	// 1.34.0
	EngineLatestMinorVersion *string `json:"EngineLatestMinorVersion,omitempty" xml:"EngineLatestMinorVersion,omitempty"`
	// The current minor version.
	//
	// example:
	//
	// 1.6.0
	EngineMinorVersion *string `json:"EngineMinorVersion,omitempty" xml:"EngineMinorVersion,omitempty"`
	// The engine version.
	//
	// example:
	//
	// 21.8.10.19
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the cluster expired. The time is in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// >  Pay-as-you-go clusters never expire. If the cluster is a pay-as-you-go cluster, an empty string is returned for this parameter.
	//
	// example:
	//
	// 2022-11-11T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The extended storage space. Unit: GB.
	//
	// example:
	//
	// 500
	ExtStorageSize *int32 `json:"ExtStorageSize,omitempty" xml:"ExtStorageSize,omitempty"`
	// The extended storage type. Valid values:
	//
	// 	- **CloudESSD**: The cluster uses an enhanced SSD (ESSD) of performance level (PL) 1.
	//
	// 	- **CloudESSD_PL2**: The cluster uses an ESSD of PL 2.
	//
	// 	- **CloudESSD_PL3**: The cluster uses an ESSD of PL 3.
	//
	// 	- **CloudEfficiency**: The cluster uses an ultra disk.
	//
	// example:
	//
	// CloudESSD
	ExtStorageType *string `json:"ExtStorageType,omitempty" xml:"ExtStorageType,omitempty"`
	// Indicates whether the cluster has expired. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsExpired *string `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	// The lock mode of the cluster. Valid values:
	//
	// 	- **Unlock**: The cluster is not locked.
	//
	// 	- **ManualLock**: The cluster is manually locked.
	//
	// 	- **LockByExpiration**: The cluster is automatically locked due to cluster expiration.
	//
	// 	- **LockByRestoration**: The cluster is automatically locked because the cluster is about to be rolled back.
	//
	// 	- **LockByDiskQuota**: The cluster is automatically locked because the disk space is exhausted.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The cause why the cluster was locked.
	//
	// >  If the value of the LockMode parameter is Unlock, an empty string is returned for this parameter.
	//
	// example:
	//
	// DISK_FULL
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The update type. If the value of the parameter is **false**, it indicates a manual update.
	//
	// example:
	//
	// false
	MaintainAutoType *bool `json:"MaintainAutoType,omitempty" xml:"MaintainAutoType,omitempty"`
	// The maintenance window of the cluster. The value is in the HH:mmZ-HH:mmZ format. The time is displayed in UTC.
	//
	// For example, if you set the maintenance window to 00:00Z-01:00Z, the cluster can be maintained from 08:00 (UTC+8) to 09:00 (UTC+8).
	//
	// example:
	//
	// 00:00Z-01:00Z
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// 	- **Postpaid**: indicates the pay-as-you-go billing method.
	//
	// 	- **Prepaid**: indicates the subscription billing method.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The HTTP port number.
	//
	// example:
	//
	// 8123
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The public endpoint.
	//
	// example:
	//
	// cc-bp1199ya710s7****.public.clickhouse.ads.aliyuncs.com
	PublicConnectionString *string `json:"PublicConnectionString,omitempty" xml:"PublicConnectionString,omitempty"`
	// The IP address that is used to connect to the cluster over the Internet.
	//
	// example:
	//
	// 121.40.xx.xx
	PublicIpAddr *string `json:"PublicIpAddr,omitempty" xml:"PublicIpAddr,omitempty"`
	// The TCP port number in the public endpoint.
	//
	// example:
	//
	// 3306
	PublicPort *string `json:"PublicPort,omitempty" xml:"PublicPort,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmyf65je6****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the data migration task.
	ScaleOutStatus *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus `json:"ScaleOutStatus,omitempty" xml:"ScaleOutStatus,omitempty" type:"Struct"`
	// The storage type of the cluster. Valid values:
	//
	// 	- **CloudESSD**: The cluster uses an enhanced SSD (ESSD) of performance level (PL) 1.
	//
	// 	- **CloudESSD_PL2**: The cluster uses an ESSD of PL 2.
	//
	// 	- **CloudESSD_PL3**: The cluster uses an ESSD of PL 3.
	//
	// 	- **CloudEfficiency**: The cluster uses an ultra disk.
	//
	// example:
	//
	// CloudESSD
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// Indicates whether data backup is supported. Valid values:
	//
	// 	- **1**: Data backup is supported.
	//
	// 	- **2**: Data backup is not supported.
	//
	// example:
	//
	// 1
	SupportBackup *int32 `json:"SupportBackup,omitempty" xml:"SupportBackup,omitempty"`
	// Indicates whether HTTPS ports are supported. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	SupportHttpsPort *bool `json:"SupportHttpsPort,omitempty" xml:"SupportHttpsPort,omitempty"`
	// Indicates whether the cluster supports a MySQL port. Valid values:
	//
	// 	- **true**: A MySQL port is supported.
	//
	// 	- **false**: A MySQL port is not supported.
	//
	// example:
	//
	// false
	SupportMysqlPort *bool `json:"SupportMysqlPort,omitempty" xml:"SupportMysqlPort,omitempty"`
	// Indicates whether tiered storage of hot data and cold data is supported. Valid values:
	//
	// 	- **1**: Tiered storage of hot data and cold data is supported.
	//
	// 	- **2**: Tiered storage of hot data and cold data is not supported.
	//
	// example:
	//
	// 1
	SupportOss *int32 `json:"SupportOss,omitempty" xml:"SupportOss,omitempty"`
	// The tags.
	Tags *DescribeDBClusterAttributeResponseBodyDBClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1n874li1t5y57wi****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC in which the cluster is deployed.
	//
	// example:
	//
	// vpc-bp10tr8k9qasioaty****
	VpcCloudInstanceId *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-bp10tr8k9qasioaty****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The IP address that is used to connect to the cluster over the VPC.
	//
	// example:
	//
	// 192.168.xx.xx
	VpcIpAddr *string `json:"VpcIpAddr,omitempty" xml:"VpcIpAddr,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The list of vSwitch IDs in multi-zone clusters.
	//
	// example:
	//
	// cn-shanghai-f: vsw-zm0n42d5vvuo****
	ZoneIdVswitchMap map[string]interface{} `json:"ZoneIdVswitchMap,omitempty" xml:"ZoneIdVswitchMap,omitempty"`
	// The ZooKeeper specifications.
	//
	// example:
	//
	// 4 Core 8 GB
	ZookeeperClass *string `json:"ZookeeperClass,omitempty" xml:"ZookeeperClass,omitempty"`
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

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetAppointmentRestartTime(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.AppointmentRestartTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetAvailableUpgradeMajorVersion(v map[string]interface{}) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.AvailableUpgradeMajorVersion = v
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

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetEngineLatestMinorVersion(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.EngineLatestMinorVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetEngineMinorVersion(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.EngineMinorVersion = &v
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

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetExtStorageSize(v int32) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ExtStorageSize = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetExtStorageType(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ExtStorageType = &v
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

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetMaintainAutoType(v bool) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.MaintainAutoType = &v
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

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetPublicIpAddr(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.PublicIpAddr = &v
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

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetResourceGroupId(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ResourceGroupId = &v
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

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetSupportOss(v int32) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.SupportOss = &v
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

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetVpcIpAddr(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.VpcIpAddr = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetZoneId(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetZoneIdVswitchMap(v map[string]interface{}) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ZoneIdVswitchMap = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetZookeeperClass(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ZookeeperClass = &v
	return s
}

type DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus struct {
	// The progress of the data migration task in percentage.
	//
	// >  This parameter is returned only when the cluster is in the SCALING_OUT state.
	//
	// example:
	//
	// 0
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The progress of the data migration task. This value is displayed in the following format: Data volume that has been migrated/Total data volume.
	//
	// >  This parameter is returned only when the cluster is in the SCALING_OUT state.
	//
	// example:
	//
	// 0MB/60469MB
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
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
	// The tag name.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// it
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeDBClusterConfigRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-wz988vja2mor4****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
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
	// The information about the parameter settings of the cluster.
	//
	// example:
	//
	// [ { "name": "keep_alive_timeout", "defaultValue": 300, "currentValue": 300, "restart": true, "valueRange": ">0", "desc": "The number of seconds that ClickHouse waits for incoming requests before closing the connection." }, ... ,{ "name": "max_partition_size_to_drop", "defaultValue": 0, "currentValue": 0, "restart": true, "valueRange": ">=0", "desc": "If the size of a MergeTree partition exceeds max_partition_size_to_drop (in bytes), you can’t delete it using a DROP query." } ]
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A23C87D-87DF-4DA0-A50E-CB13F4F7923D
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDBClusterConfigResponse) SetStatusCode(v int32) *DescribeDBClusterConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterConfigResponse) SetBody(v *DescribeDBClusterConfigResponseBody) *DescribeDBClusterConfigResponse {
	s.Body = v
	return s
}

type DescribeDBClusterConfigInXMLRequest struct {
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all the clusters that are deployed in a specific region. The information includes the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBClusterConfigInXMLRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigInXMLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigInXMLRequest) SetDBClusterId(v string) *DescribeDBClusterConfigInXMLRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterConfigInXMLRequest) SetRegionId(v string) *DescribeDBClusterConfigInXMLRequest {
	s.RegionId = &v
	return s
}

type DescribeDBClusterConfigInXMLResponseBody struct {
	// The values of the configuration parameters.
	//
	// example:
	//
	// <?xml version="1.0"?>
	//
	// <yandex>
	//
	//     <keep_alive_timeout>300</keep_alive_timeout>
	//
	//     <listen_backlog>64</listen_backlog>
	//
	//     <logger>
	//
	//         <level>debug</level>
	//
	//         <size>1000M</size>
	//
	//     </logger>
	//
	//     <mark_cache_size>6871947673</mark_cache_size>
	//
	//     <max_concurrent_queries>100</max_concurrent_queries>
	//
	//     <max_connections>4096</max_connections>
	//
	//     <max_partition_size_to_drop>0</max_partition_size_to_drop>
	//
	//     <max_server_memory_usage>0</max_server_memory_usage>
	//
	// <max_server_memory_usage_to_ram_ratio>0.9</max_server_memory_usage_to_ram_ratio>
	//
	//     <max_table_size_to_drop>0</max_table_size_to_drop>
	//
	//     <max_thread_pool_size>10000</max_thread_pool_size>
	//
	//     <merge_tree>
	//
	//         <max_delay_to_insert>256</max_delay_to_insert>
	//
	//         <max_part_loading_threads>16</max_part_loading_threads>      <zookeeper_session_expiration_check_period>1</zookeeper_session_expiration_check_period>
	//
	//     </merge_tree>
	//
	//     <total_memory_profiler_step>4194304</total_memory_profiler_step>
	//
	// <total_memory_tracker_sample_probability>0</total_memory_tracker_sample_probability>
	//
	//     <uncompressed_cache_size>3435973836</uncompressed_cache_size>
	//
	// </yandex>
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE042911-2B00-134C-9F42-816871BBAFB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterConfigInXMLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigInXMLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigInXMLResponseBody) SetConfig(v string) *DescribeDBClusterConfigInXMLResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeDBClusterConfigInXMLResponseBody) SetRequestId(v string) *DescribeDBClusterConfigInXMLResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterConfigInXMLResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterConfigInXMLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterConfigInXMLResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterConfigInXMLResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigInXMLResponse) SetHeaders(v map[string]*string) *DescribeDBClusterConfigInXMLResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterConfigInXMLResponse) SetStatusCode(v int32) *DescribeDBClusterConfigInXMLResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterConfigInXMLResponse) SetBody(v *DescribeDBClusterConfigInXMLResponseBody) *DescribeDBClusterConfigInXMLResponse {
	s.Body = v
	return s
}

type DescribeDBClusterNetInfoItemsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1qx68m06981****
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
	// The network type of the cluster. Only VPC is supported.
	//
	// example:
	//
	// VPC
	ClusterNetworkType *string `json:"ClusterNetworkType,omitempty" xml:"ClusterNetworkType,omitempty"`
	// Indicates whether Server Load Balancer (SLB) is activated in the VPC. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableSLB *bool `json:"EnableSLB,omitempty" xml:"EnableSLB,omitempty"`
	// The network information about the cluster.
	NetInfoItems *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems `json:"NetInfoItems,omitempty" xml:"NetInfoItems,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9A23C87D-87DF-4DA0-A50E-CB13F4F7923D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DescribeDBClusterNetInfoItemsResponseBody) SetEnableSLB(v bool) *DescribeDBClusterNetInfoItemsResponseBody {
	s.EnableSLB = &v
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
	// The endpoint that is used to connect to the database.
	//
	// example:
	//
	// cc-bp1554t789i8e****.clickhouse.ads.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The HTTP port number.
	//
	// example:
	//
	// 8123
	HttpPort *string `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	// The HTTPS port number.
	//
	// example:
	//
	// 8443
	HttpsPort *string `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 10.255.234.251
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The port number that is used in Java Database Connectivity (JDBC).
	//
	// example:
	//
	// 3306
	JdbcPort *string `json:"JdbcPort,omitempty" xml:"JdbcPort,omitempty"`
	// The port of the MySQL instance.
	//
	// example:
	//
	// 9004
	MySQLPort *string `json:"MySQLPort,omitempty" xml:"MySQLPort,omitempty"`
	// The network type of the endpoint. Valid values:
	//
	// 	- Public: public endpoint
	//
	// 	- VPC: VPC
	//
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The vSwitch ID.
	//
	// >  If the value of the NetType parameter is set to Public, an empty string is returned.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// >  If the value of the NetType parameter is set to Public, an empty string is returned.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetHttpsPort(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.HttpsPort = &v
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

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetMySQLPort(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.MySQLPort = &v
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterNetInfoItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDBClusterNetInfoItemsResponse) SetStatusCode(v int32) *DescribeDBClusterNetInfoItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponse) SetBody(v *DescribeDBClusterNetInfoItemsResponseBody) *DescribeDBClusterNetInfoItemsResponse {
	s.Body = v
	return s
}

type DescribeDBClusterPerformanceRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp125e3uu94wo****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// >  The end time must be later than the start time. The interval cannot be more than 32 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-27T16:38Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance metrics that you want to query. Separate multiple performance metrics with commas (,). You can query up to five performance metrics at a time. You can query the following performance metrics:
	//
	// >  The **Key*	- parameter is required.
	//
	// 	- **CPU**:
	//
	//     	- **CPU_USAGE**: the CPU utilization
	//
	// 	- **Memory**:
	//
	//     	- **MEM_USAGE**: the memory usage
	//
	//     	- **MEM_USAGE_SIZE**: the used memory. Unit: MB
	//
	// 	- **Disk**:
	//
	//     	- **DISK_USAGE**: the disk usage
	//
	//     	- **DISK_USAGE_SIZE**: the size of the used disks. Unit: MB
	//
	//     	- **IOPS**: the disk Input/Output Operations per Second (IOPS)
	//
	// 	- **Connection**:
	//
	//     	- **CONN_USAGE**: the database connection usage
	//
	//     	- **CONN_USAGE_COUNT**: the number of database connections used
	//
	// 	- **Write**:
	//
	//     	- **TPS**: the number of rows written per second
	//
	//     	- **INSERT_SIZE**: the amount of data written per second. Unit: MB
	//
	// 	- **Query**:
	//
	//     	- **QPS**: the queries per second
	//
	//     	- **AVG_SEEK**: the average number of random seek calls
	//
	// 	- **WAIT**:
	//
	//     	- **ZK_WAIT**: the average ZooKeeper wait time. Unit: ms
	//
	//     	- **IO_WAIT**: the average I/O wait time. Unit: ms
	//
	//     	- **CPU_WAIT**: the average CPU wait time. Unit: ms
	//
	// example:
	//
	// MEM_USAGE
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in Coordinated Universal Time (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-27T16:37Z
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
	// The cluster ID.
	//
	// example:
	//
	// cc-bp125e3uu94wo****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// 2021-11-27T16:38Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The values of the queried performance metrics of the cluster.
	Performances []*DescribeDBClusterPerformanceResponseBodyPerformances `json:"Performances,omitempty" xml:"Performances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// FE242962-6DA3-5FC8-9691-37B62A3210F7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range to query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-27T16:37Z
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
	//
	// example:
	//
	// MEM_USAGE
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the performance metric value.
	//
	// example:
	//
	// mem_usage
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The queried performance pamaters.
	Series []*DescribeDBClusterPerformanceResponseBodyPerformancesSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	// The unit of the performance metric.
	//
	// example:
	//
	// %
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
	// The name of the list of performance metric values.
	//
	// example:
	//
	// cc-bp125e3uu94wo1s0k16****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The values of the performance parameter. Each value of the performance parameter is collected at a point in time.
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
	// The values of a metric.
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeDBClustersRequest struct {
	// The description of the cluster.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The cluster ID.
	//
	// >  If you do not specify this parameter, the information about all clusters is queried.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	// The state of the cluster. Valid values:
	//
	// 	- **Preparing**: The cluster is being prepared.
	//
	// 	- **Creating**: The cluster is being created.
	//
	// 	- **Running**: The cluster is running.
	//
	// 	- **Deleting**: The cluster is being deleted.
	//
	// 	- **SCALING_OUT**: The storage capacity of the cluster is being expanded.
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the cluster belongs.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
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
	// The tag name.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// it
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
	// The details of the clusters.
	DBClusters *DescribeDBClustersResponseBodyDBClusters `json:"DBClusters,omitempty" xml:"DBClusters,omitempty" type:"Struct"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 140692647406****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The site ID. Valid values:
	//
	// 	- **26842**: the China site (aliyun.com)
	//
	// 	- **26888**: the international site (alibabacloud.com)
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// 	- **Basic**: Single-replica Edition
	//
	// 	- **HighAvailability**: Double-replica Edition
	//
	// example:
	//
	// Basic
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The commodity code of the cluster.
	//
	// example:
	//
	// clickhouse_go_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The VPC endpoint of the cluster.
	//
	// example:
	//
	// cc-bp1fs5o051c61****.clickhouse.ads.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The version number of the backend management system of ApsaraDB for ClickHouse. Valid values:
	//
	// 	- **v1**
	//
	// 	- **v2**
	//
	// example:
	//
	// v1
	ControlVersion *string `json:"ControlVersion,omitempty" xml:"ControlVersion,omitempty"`
	// The time when the cluster was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2021-10-28T07:24:45Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The network type of the cluster. Only VPC is supported.
	//
	// example:
	//
	// VPC
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The state of the cluster. Valid values:
	//
	// 	- **Preparing**: The cluster is being prepared.
	//
	// 	- **Creating**: The cluster is being created.
	//
	// 	- **Running**: The cluster is running.
	//
	// 	- **Deleting**: The cluster is being deleted.
	//
	// 	- **SCALING_OUT**: The storage capacity of the cluster is being expanded.
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The specifications of the cluster.
	//
	// 	- Valid values when the cluster is of Single-replica Edition: -**S4**: 4 CPU cores and 16 GB of memory -**S8**: 8 CPU cores and 32 GB of memory
	//
	//     	- **S16**: 16 CPU cores and 64 GB of memory
	//
	//     	- **S32**: 32 CPU cores and 128 GB of memory
	//
	//     	- **S64**: 64 CPU cores and 256 GB of memory
	//
	//     	- **S104**: 104 CPU cores and 384 GB of memory
	//
	// 	- Valid values when the cluster is of Double-replica Edition: -**C4**: 4 CPU cores and 16 GB of memory -**C8**: 8 CPU cores and 32 GB of memory -**C16**: 16 CPU cores and 64 GB of memory -**C32**: 32 CPU cores and 128 GB of memory -**C64**: 64 CPU cores and 256 GB of memory -**C104**: 104 CPU cores and 384 GB of memory
	//
	// example:
	//
	// C8
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of nodes.
	//
	// 	- Valid values when the cluster is of Single-replica Edition: 1 to 48.
	//
	// 	- Valid values when the cluster is of Double-replica Edition: 1 to 24.
	//
	// example:
	//
	// 2
	DBNodeCount *int64 `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	// The storage capacity of each node. Valid values: 100 to 32000. Unit: GB.
	//
	// >  This value is a multiple of 100.
	//
	// example:
	//
	// 100
	DBNodeStorage *int64 `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	// The time when the cluster expired. The time is in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// >  Pay-as-you-go clusters never expire. If the cluster is a pay-as-you-go cluster, an empty string is returned for this parameter.
	//
	// example:
	//
	// 2011-05-30T12:11:4Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The extended storage space.
	//
	// example:
	//
	// 100GB
	ExtStorageSize *int32 `json:"ExtStorageSize,omitempty" xml:"ExtStorageSize,omitempty"`
	// The extended storage type. Valid values:
	//
	// 	- **CloudSSD**: standard SSD.
	//
	// 	- **CloudESSD**: The cluster uses an enhanced SSD (ESSD) of performance level (PL) 1.
	//
	// 	- **CloudESSD_PL2**: The cluster uses an ESSD of PL 2.
	//
	// 	- **CloudESSD_PL3**: The cluster uses an ESSD of PL 3.
	//
	// 	- **CloudEfficiency**: The cluster uses an ultra disk.
	//
	// example:
	//
	// CloudESSD
	ExtStorageType *string `json:"ExtStorageType,omitempty" xml:"ExtStorageType,omitempty"`
	// Indicates whether the cluster has expired. Valid values:
	//
	// 	- **true**: The cluster has expired.
	//
	// 	- **false**: The cluster has not expired.
	//
	// example:
	//
	// false
	IsExpired *string `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	// The lock mode of the cluster. Valid values:
	//
	// 	- **Unlock**: The cluster is not locked.
	//
	// 	- **ManualLock**: The cluster is manually locked.
	//
	// 	- **LockByExpiration**: The cluster is automatically locked due to cluster expiration.
	//
	// 	- **LockByRestoration**: The cluster is automatically locked because the cluster is about to be rolled back.
	//
	// 	- **LockByDiskQuota**: The cluster is automatically locked because the disk space is exhausted.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The cause why the cluster was locked.
	//
	// >  If the value of the LockMode parameter is Unlock, an empty string is returned for this parameter.
	//
	// example:
	//
	// DISK_FULL
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// 	- **Postpaid**: The cluster uses the pay-as-you-go billing method.
	//
	// 	- **Prepaid**: The cluster uses the subscription billing method.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The HTTP port number.
	//
	// example:
	//
	// 8123
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the cluster belongs.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of a data migration task.
	ScaleOutStatus *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus `json:"ScaleOutStatus,omitempty" xml:"ScaleOutStatus,omitempty" type:"Struct"`
	// The storage type of the cluster. Valid values:
	//
	// 	- **CloudESSD**: The cluster uses an enhanced SSD (ESSD) of performance level (PL) 1.
	//
	// 	- **CloudESSD_PL2**: The cluster uses an ESSD of PL 2.
	//
	// 	- **CloudESSD_PL3**: The cluster uses an ESSD of PL 3.
	//
	// 	- **CloudEfficiency**: The cluster uses an ultra disk.
	//
	// example:
	//
	// CloudESSD
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The tags.
	Tags *DescribeDBClustersResponseBodyDBClustersDBClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC in which the cluster is deployed.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcCloudInstanceId *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the cluster is deployed.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetExtStorageSize(v int32) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ExtStorageSize = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetExtStorageType(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ExtStorageType = &v
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

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetResourceGroupId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ResourceGroupId = &v
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
	// The progress of the data migration task in percentage.
	//
	// >  This parameter is returned only when the cluster is in the SCALING_OUT state.
	//
	// example:
	//
	// 0
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The progress of the data migration task. This value is displayed in the following format: Data volume that has been migrated/Total data volume.
	//
	// >  This parameter is returned only when the cluster is in the SCALING_OUT state.
	//
	// example:
	//
	// 0MB/60469MB
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
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
	// The tag name.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// it
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeDBConfigRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-t4nw17nh2e4t2****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
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
	// The configuration information about the cluster.
	//
	// example:
	//
	// <dictionaries><dictionary><name>test</name><source><clickhouse><host>10.37.XX.XX</host><port>9000</port><user>default</user><password></password><db>default</db><table>t_organization</table><where>id=1</where><invalidate_query>SQL_QUERY</invalidate_query></clickhouse></source><lifetime><min>31</min><max>33</max></lifetime><layout><flat></flat></layout><structure><key><attribute><name>field1</name><type>String</type></attribute></key></structure></dictionary></dictionaries>
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16060117-92E1-5F3B-BF42-28B172D0F869
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDBConfigResponse) SetStatusCode(v int32) *DescribeDBConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBConfigResponse) SetBody(v *DescribeDBConfigResponseBody) *DescribeDBConfigResponse {
	s.Body = v
	return s
}

type DescribeOSSStorageRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
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
	// Indicates whether tiered storage of hot data and cold data is supported. Valid values:
	//
	// 	- **true**: Tiered storage of hot data and cold data is supported.
	//
	// 	- **false**: Tiered storage of hot data and cold data is not supported.
	//
	// example:
	//
	// true
	ColdStorage *bool `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The parameters for tiered storage of hot data and cold data.
	//
	// example:
	//
	// [{ "currentValue":"0.1", "defaultValue":"0.1", "desc":"Ratio of free disk space. When the ratio exceeds the value of configuration parameter, ClickHouse start to move data to the cold storage", "name":"move_factor", "restart":true, "valueRange":"(0, 1]" },{ "currentValue":"true", "defaultValue":"true", "desc":"Disables merging of data parts on cold storage", "name":"prefer_not_to_merge", "restart":true, "valueRange":"true|false" }]
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The request ID.
	//
	// example:
	//
	// aadbb456-ebf7-4ed8-9671-fad9f3846ca4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of tiered storage of hot data and cold data. Valid values:
	//
	// 	- **CREATING**: Tiered storage of hot data and cold data is being enabled.
	//
	// 	- **DISABLE**: Tiered storage of hot data and cold data is not enabled.
	//
	// 	- **ENABLE**: Tiered storage of hot data and cold data is enabled.
	//
	// example:
	//
	// ENABLE
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The space used for tiered storage of hot data and cold data. Unit: GB.
	//
	// example:
	//
	// 0.00
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOSSStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeOSSStorageResponse) SetStatusCode(v int32) *DescribeOSSStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOSSStorageResponse) SetBody(v *DescribeOSSStorageResponseBody) *DescribeOSSStorageResponse {
	s.Body = v
	return s
}

type DescribeProcessListRequest struct {
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all the clusters that are deployed in a specific region. The information includes the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1190tj036am****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the query statement.
	//
	// example:
	//
	// 6c69d508-f05f-4c74-857c-d982b7e7e79f
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	// The account that is used to log on to the database.
	//
	// example:
	//
	// user
	InitialUser *string `json:"InitialUser,omitempty" xml:"InitialUser,omitempty"`
	// The keyword that is used to query.
	//
	// example:
	//
	// join
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Sorting by the specified column name. Valid values:
	//
	// 	- elapsed: the cumulative execution time
	//
	// 	- written_rows: the number of written rows
	//
	// 	- read_rows: the number of read rows
	//
	// 	- memory_usage: the memory usage
	//
	// example:
	//
	// elapsed
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The minimum query duration. The minimum value is **1000**, and the default value is **1000**. Unit: milliseconds. Queries that last longer than this duration are returned in response parameters.
	//
	// example:
	//
	// 500
	QueryDurationMs *int32 `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
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
	// The queries.
	ProcessList *DescribeProcessListResponseBodyProcessList `json:"ProcessList,omitempty" xml:"ProcessList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FD61BB0D-788A-5185-A8E3-1B90BA8F6F04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The details of the query.
	Data *DescribeProcessListResponseBodyProcessListData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The number of rows returned for the query.
	//
	// example:
	//
	// 1145700
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
	RowsBeforeLimitAtLeast *string `json:"RowsBeforeLimitAtLeast,omitempty" xml:"RowsBeforeLimitAtLeast,omitempty"`
	// The statistics of the results.
	Statistics *DescribeProcessListResponseBodyProcessListStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
	// Details of the columns.
	TableSchema *DescribeProcessListResponseBodyProcessListTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
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
	// The IP address of the client that initiates the query.
	//
	// example:
	//
	// ::ffff:10.1.XX.XX
	InitialAddress *string `json:"InitialAddress,omitempty" xml:"InitialAddress,omitempty"`
	// The query ID.
	//
	// example:
	//
	// 2dd144fd-4230-4249-b15c-e63f964fbb5a
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	// The database account.
	//
	// example:
	//
	// test
	InitialUser *string `json:"InitialUser,omitempty" xml:"InitialUser,omitempty"`
	// The SQL statement that is executed in the query.
	//
	// example:
	//
	// select 	- from test order by score limit 1;
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The execution duration of the query. Unit: milliseconds.
	//
	// example:
	//
	// 2000
	QueryDurationMs *string `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// The beginning of the time range to query. The value is in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// 2021-02-02T09:14:48Z
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
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
	// The size of the data that was scanned. Unit: bytes.
	//
	// example:
	//
	// 9141300000
	BytesRead *int32 `json:"BytesRead,omitempty" xml:"BytesRead,omitempty"`
	// The average response time.
	//
	// example:
	//
	// 4156
	ElapsedTime *float32 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// The number of scanned rows.
	//
	// example:
	//
	// 1000000
	RowsRead *int32 `json:"RowsRead,omitempty" xml:"RowsRead,omitempty"`
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
	// The column name.
	//
	// example:
	//
	// InitialUser
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The column type.
	//
	// example:
	//
	// String
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProcessListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The queried regions.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
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
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zones.
	Zones *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
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
	// Indicates whether Virtual Private Cloud (VPC) is supported in the zone. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	VpcEnabled *bool `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeSchemasRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
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
	// The information about the databases of the cluster.
	Items *DescribeSchemasResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
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
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all the clusters that are deployed in a specific region. The information includes the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1z58t881wcx****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-dd hh:mm:ss format. The time must be in UTC.
	//
	// >  The end time must be later than the start time. The specified time range that can be specified must be less than seven days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-05-27 16:00:00
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The minimum query duration. The minimum value is **1000**, and the default value is **1000**. Unit: milliseconds. Queries that last longer than this duration are returned in response parameters.
	//
	// example:
	//
	// 1000
	QueryDurationMs *int32 `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-dd hh:mm:ss format. The time must be in Coordinated Universal Time (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-05-20 16:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// DF203CC8-5F68-5E3F-8050-3C77DD65731A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the slow query logs.
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
	// Details about the slow query logs.
	Data *DescribeSlowLogRecordsResponseBodySlowLogRecordsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The number of rows in the result set.
	//
	// example:
	//
	// 1
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	RowsBeforeLimitAtLeast *string `json:"RowsBeforeLimitAtLeast,omitempty" xml:"RowsBeforeLimitAtLeast,omitempty"`
	// The statistics of the results.
	Statistics *DescribeSlowLogRecordsResponseBodySlowLogRecordsStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
	// The schema of the table in the database.
	TableSchema *DescribeSlowLogRecordsResponseBodySlowLogRecordsTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
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
	// The IP address of the client that initiated the query.
	//
	// example:
	//
	// ::ffff:100.104.XX.XX
	InitialAddress *string `json:"InitialAddress,omitempty" xml:"InitialAddress,omitempty"`
	// The query ID.
	//
	// example:
	//
	// \\"b51496f2-6b0b-4546-aff9-e17951cb9410\\"
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	// The username that is used to initiate the query.
	//
	// example:
	//
	// test_users
	InitialUser *string `json:"InitialUser,omitempty" xml:"InitialUser,omitempty"`
	// The peak memory usage for the query. Unit: bytes.
	//
	// example:
	//
	// 1048576
	MemoryUsage *string `json:"MemoryUsage,omitempty" xml:"MemoryUsage,omitempty"`
	// The statement that was executed in the query.
	//
	// example:
	//
	// Select 	- from table
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The duration of the query. Unit: milliseconds.
	//
	// example:
	//
	// 2000
	QueryDurationMs *string `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// The beginning of the time range to query. The time is in the yyyy-MM-dd hh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-05-22 20:00:01
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	// The size of the data read by executing the statement. Unit: bytes.
	//
	// example:
	//
	// 1048576
	ReadBytes *string `json:"ReadBytes,omitempty" xml:"ReadBytes,omitempty"`
	// The number of rows read by executing the statement.
	//
	// example:
	//
	// 10027008
	ReadRows *string `json:"ReadRows,omitempty" xml:"ReadRows,omitempty"`
	// The size of the result data. Unit: bytes.
	//
	// example:
	//
	// 1024
	ResultBytes *string `json:"ResultBytes,omitempty" xml:"ResultBytes,omitempty"`
	// The query status. Valid values:
	//
	// 	- **QueryFinish**: The query is complete.
	//
	// 	- **Processing**: The query is running.
	//
	// example:
	//
	// QueryFinish
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// The total size of data that were read. Unit: bytes.
	//
	// example:
	//
	// 123456
	BytesRead *int32 `json:"BytesRead,omitempty" xml:"BytesRead,omitempty"`
	// The time consumed by the slow query. Unit: milliseconds.
	//
	// example:
	//
	// 21.35
	ElapsedTime *float32 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// The total number of rows that were read.
	//
	// example:
	//
	// 2016722
	RowsRead *int32 `json:"RowsRead,omitempty" xml:"RowsRead,omitempty"`
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
	// The name of the column.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the column.
	//
	// example:
	//
	// String
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlowLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeSynDbTablesRequest struct {
	// The ID of the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp158i5wvj436****
	DbClusterId          *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// database
	SynDb *string `json:"SynDb,omitempty" xml:"SynDb,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 851D11EA-681C-5B38-A065-C3F90BBD49DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried tables.
	Tables []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSynDbTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeSynDbTablesResponse) SetStatusCode(v int32) *DescribeSynDbTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynDbTablesResponse) SetBody(v *DescribeSynDbTablesResponseBody) *DescribeSynDbTablesResponse {
	s.Body = v
	return s
}

type DescribeSynDbsRequest struct {
	// The ID of the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1ab22b80814****
	DbClusterId  *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// 20
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *DescribeSynDbsRequest) SetPageNumber(v int32) *DescribeSynDbsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSynDbsRequest) SetPageSize(v int32) *DescribeSynDbsRequest {
	s.PageSize = &v
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7655F5F9-1313-5ABA-8516-F6EB79605A5F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about data synchronization between the ApsaraDB for ClickHouse cluster and an ApsaraDB RDS for MySQL instance.
	SynDbs []*DescribeSynDbsResponseBodySynDbs `json:"SynDbs,omitempty" xml:"SynDbs,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSynDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynDbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynDbsResponseBody) SetPageNumber(v int32) *DescribeSynDbsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSynDbsResponseBody) SetPageSize(v int32) *DescribeSynDbsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSynDbsResponseBody) SetRequestId(v string) *DescribeSynDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynDbsResponseBody) SetSynDbs(v []*DescribeSynDbsResponseBodySynDbs) *DescribeSynDbsResponseBody {
	s.SynDbs = v
	return s
}

func (s *DescribeSynDbsResponseBody) SetTotalCount(v int32) *DescribeSynDbsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSynDbsResponseBodySynDbs struct {
	// 	- When the value **true*	- is returned for the **SynStatus*	- parameter, the system does not return the ErrorMsg parameter.
	//
	// 	- When the value **false*	- is returned for the **SynStatus*	- parameter, the system returns for the ErrorMsg parameter the cause why the data synchronization failed.
	//
	// example:
	//
	// ClickHouse exception, code: 49, host: 100.100.118.132, port: 49670; Code: 49, e.displayText() = DB::Exception: Logical error: there is no global context (version 20.8.17.25)n
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The ID of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// rm-wz9d11qg1j0h4****
	RdsId *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	// The database account that is used to log on to the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// test
	RdsUserName *string `json:"RdsUserName,omitempty" xml:"RdsUserName,omitempty"`
	// The internal endpoint of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// rm-bp16t9h3999xb0a711****.mysql.rds.aliyuncs.com:3306
	RdsVpcUrl *string `json:"RdsVpcUrl,omitempty" xml:"RdsVpcUrl,omitempty"`
	// The name of the database in the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// database
	SynDb *string `json:"SynDb,omitempty" xml:"SynDb,omitempty"`
	// Indicates whether the data synchronization succeeded. Valid values:
	//
	// 	- **true**: The data synchronization succeeded.
	//
	// 	- **false**: The data synchronization failed.
	//
	// example:
	//
	// true
	SynStatus *bool `json:"SynStatus,omitempty" xml:"SynStatus,omitempty"`
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

func (s *DescribeSynDbsResponseBodySynDbs) SetRdsUserName(v string) *DescribeSynDbsResponseBodySynDbs {
	s.RdsUserName = &v
	return s
}

func (s *DescribeSynDbsResponseBodySynDbs) SetRdsVpcUrl(v string) *DescribeSynDbsResponseBodySynDbs {
	s.RdsVpcUrl = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSynDbsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeSynDbsResponse) SetStatusCode(v int32) *DescribeSynDbsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynDbsResponse) SetBody(v *DescribeSynDbsResponseBody) *DescribeSynDbsResponse {
	s.Body = v
	return s
}

type DescribeTablesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The database name.
	//
	// This parameter is required.
	//
	// example:
	//
	// database
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
	// The information about the tables.
	Items *DescribeTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
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
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name.
	//
	// example:
	//
	// test
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeTransferHistoryRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
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
	// The migration information.
	HistoryDetails *DescribeTransferHistoryResponseBodyHistoryDetails `json:"HistoryDetails,omitempty" xml:"HistoryDetails,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The progress of the data migration.
	//
	// example:
	//
	// 100%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the source cluster.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	SourceDBCluster *string `json:"SourceDBCluster,omitempty" xml:"SourceDBCluster,omitempty"`
	// The status of the data migration task. Valid values:
	//
	// 	- **Finished**: The data migration task is complete.
	//
	// 	- **Processing**: The data migration task is in progress.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the destination cluster.
	//
	// example:
	//
	// cc-bp13zkh9uw523****
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTransferHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeTransferHistoryResponse) SetStatusCode(v int32) *DescribeTransferHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTransferHistoryResponse) SetBody(v *DescribeTransferHistoryResponseBody) *DescribeTransferHistoryResponse {
	s.Body = v
	return s
}

type KillProcessRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The query statement or query statements that you want to stop executing. If you want to stop executing multiple query statements, separate the statements with commas (,).
	//
	// >  If you do not set this parameter, all query statements are stopped by default.
	//
	// example:
	//
	// SELECT 	- FROM `test01` ;
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
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
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KillProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ModifyAccountAuthorityRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The databases to which you want to grant permissions. Separate databases with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// db1
	AllowDatabases *string `json:"AllowDatabases,omitempty" xml:"AllowDatabases,omitempty"`
	// The dictionaries to which you want to grant permissions. Separate dictionaries with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// dt1
	AllowDictionaries *string `json:"AllowDictionaries,omitempty" xml:"AllowDictionaries,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to grant DDL permissions to the database account. Valid values:
	//
	// 	- **true**: grants DDL permissions to the database account.
	//
	// 	- **false**: does not grant DDL permissions to the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DdlAuthority *bool `json:"DdlAuthority,omitempty" xml:"DdlAuthority,omitempty"`
	// Specifies whether to grant DML permissions to the database account. Valid values:
	//
	// 	- **all**
	//
	// 	- **readonly,modify**
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	DmlAuthority *string `json:"DmlAuthority,omitempty" xml:"DmlAuthority,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// All databases. Separate databases with commas (,).
	//
	// example:
	//
	// db1,db2
	TotalDatabases *string `json:"TotalDatabases,omitempty" xml:"TotalDatabases,omitempty"`
	// All dictionaries. Separate dictionaries with commas (,).
	//
	// example:
	//
	// dt1,dt2
	TotalDictionaries *string `json:"TotalDictionaries,omitempty" xml:"TotalDictionaries,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyAccountAuthorityResponse) SetStatusCode(v int32) *ModifyAccountAuthorityResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountAuthorityResponse) SetBody(v *ModifyAccountAuthorityResponseBody) *ModifyAccountAuthorityResponse {
	s.Body = v
	return s
}

type ModifyAccountDescriptionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ceshi
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
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
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ModifyBackupPolicyRequest struct {
	// The retention period for the backup data. Valid values: 7 to 730. Unit: day.
	//
	// example:
	//
	// 7
	BackupRetentionPeriod *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1qx68m06981****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The day of a week when the system regularly backs up data. If you specify multiple days of a week, separate them with commas (,). Valid values:
	//
	// 	- **Monday**
	//
	// 	- **Tuesday**
	//
	// 	- **Wednesday**
	//
	// 	- **Thursday**
	//
	// 	- **Friday**
	//
	// 	- **Saturday**
	//
	// 	- **Sunday**
	//
	// This parameter is required.
	//
	// example:
	//
	// Monday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The backup window. Specify the time in the ISO 8601 standard in the HH:mmZ-HH:mmZ format. The time must be in Coordinated Universal Time (UTC).
	//
	// For example, if you set the backup window to 00:00Z-01:00Z, the data of the cluster can be backed up from 08:00 (UTC+8) to 09:00 (UTC+8).
	//
	// This parameter is required.
	//
	// example:
	//
	// 15:00Z-16:00Z
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
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ModifyDBClusterRequest struct {
	// The specifications of the cluster.
	//
	// 	- Valid values when the cluster is of Single-replica Edition:
	//
	//     	- **S4-NEW**
	//
	//     	- **S8**
	//
	//     	- **S16**
	//
	//     	- **S32**
	//
	//     	- **S64**
	//
	//     	- **S104**
	//
	// 	- Valid values when the cluster is of Double-replica Edition:
	//
	//     	- **C4-NEW**
	//
	//     	- **C8**
	//
	//     	- **C16**
	//
	//     	- **C32**
	//
	//     	- **C64**
	//
	//     	- **C104**
	//
	// This parameter is required.
	//
	// example:
	//
	// S4-NEW
	DBClusterClass *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp19lo45sy98x****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The number of nodes in the cluster.
	//
	// 	- If the cluster is of Single-replica Edition, the value must be an integer that ranges from 1 to 48.
	//
	// 	- If the cluster is of Double-replica Edition, the value must be an integer that ranges from 1 to 24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DBNodeGroupCount *string `json:"DBNodeGroupCount,omitempty" xml:"DBNodeGroupCount,omitempty"`
	// The storage capacity of a single node of the cluster. Unit: GB.
	//
	// Valid values: 100 to 32000.
	//
	// >  This value is a multiple of 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	DBNodeStorage     *string `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	DbNodeStorageType *string `json:"DbNodeStorageType,omitempty" xml:"DbNodeStorageType,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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

func (s *ModifyDBClusterRequest) SetDbNodeStorageType(v string) *ModifyDBClusterRequest {
	s.DbNodeStorageType = &v
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
	// The information about the cluster.
	DBCluster *ModifyDBClusterResponseBodyDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BA30A000-3A4D-5B97-9420-E5D0D49F7016
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The cluster ID.
	//
	// example:
	//
	// cc-bp19lo45sy98x****
	DbClusterId *string `json:"dbClusterId,omitempty" xml:"dbClusterId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 21417210003****
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// NULL
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	// example:
	//
	// default
	DBClusterIPArrayName *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1qx68m06981****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// Cover
	ModifyMode           *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 192.168.xx.xx
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
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterAccessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ModifyDBClusterConfigRequest struct {
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all the clusters that are deployed in a specific region. The information includes the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1t9lbb7a4z7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The reason for the change.
	//
	// example:
	//
	// test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The names of the parameters and the new values that you want to specify for the parameters.
	//
	// >  You can change the value of a single parameter. The values of parameters that are not specified will not be changed.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"keep_alive_timeout":"301"}
	UserConfig *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
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

func (s *ModifyDBClusterConfigRequest) SetReason(v string) *ModifyDBClusterConfigRequest {
	s.Reason = &v
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
	// The request ID.
	//
	// example:
	//
	// A9AA1E0A-2AEE-5847-87CF-E4FDC0E66052
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyDBClusterConfigResponse) SetStatusCode(v int32) *ModifyDBClusterConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterConfigResponse) SetBody(v *ModifyDBClusterConfigResponseBody) *ModifyDBClusterConfigResponse {
	s.Body = v
	return s
}

type ModifyDBClusterConfigInXMLRequest struct {
	// The configuration parameters whose settings you want to modify. You can call the [DescribeDBClusterConfigInXML](https://help.aliyun.com/document_detail/452210.html) operation to query configuration parameters, and modify the settings of the returned configuration parameters.
	//
	// > You must specify all configuration parameters even when you want to modify the setting of a single parameter. If a configuration parameter is not specified, the original value of this parameter is retained or the modification fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// <?xml version="1.0"?>
	//
	// <yandex>
	//
	//     <keep_alive_timeout>400</keep_alive_timeout>
	//
	//     <listen_backlog>4096</listen_backlog>
	//
	//     <logger>
	//
	//         <level>debug</level>
	//
	//         <size>1000M</size>
	//
	//     </logger>
	//
	//     <mark_cache_size>5368709120</mark_cache_size>
	//
	//     <max_concurrent_queries>201</max_concurrent_queries>
	//
	//     <max_connections>4096</max_connections>
	//
	//     <max_partition_size_to_drop>0</max_partition_size_to_drop>
	//
	//     <max_table_size_to_drop>0</max_table_size_to_drop>
	//
	//     <merge_tree>
	//
	//         <max_delay_to_insert>256</max_delay_to_insert>
	//
	//         <max_part_loading_threads>auto</max_part_loading_threads>
	//
	//         <max_suspicious_broken_parts>100</max_suspicious_broken_parts>
	//
	//         <zookeeper_session_expiration_check_period>1</zookeeper_session_expiration_check_period>
	//
	//     </merge_tree>
	//
	//     <uncompressed_cache_size>1717986918</uncompressed_cache_size>
	//
	// </yandex>
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all the clusters that are deployed in a specific region. The information includes the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The reason for the modification.
	//
	// example:
	//
	// test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The region ID of the cluster. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDBClusterConfigInXMLRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterConfigInXMLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigInXMLRequest) SetConfig(v string) *ModifyDBClusterConfigInXMLRequest {
	s.Config = &v
	return s
}

func (s *ModifyDBClusterConfigInXMLRequest) SetDBClusterId(v string) *ModifyDBClusterConfigInXMLRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterConfigInXMLRequest) SetReason(v string) *ModifyDBClusterConfigInXMLRequest {
	s.Reason = &v
	return s
}

func (s *ModifyDBClusterConfigInXMLRequest) SetRegionId(v string) *ModifyDBClusterConfigInXMLRequest {
	s.RegionId = &v
	return s
}

type ModifyDBClusterConfigInXMLResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BDD29EB1-BE76-5CFA-9068-D34B696310BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterConfigInXMLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterConfigInXMLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigInXMLResponseBody) SetRequestId(v string) *ModifyDBClusterConfigInXMLResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterConfigInXMLResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterConfigInXMLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterConfigInXMLResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterConfigInXMLResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigInXMLResponse) SetHeaders(v map[string]*string) *ModifyDBClusterConfigInXMLResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterConfigInXMLResponse) SetStatusCode(v int32) *ModifyDBClusterConfigInXMLResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterConfigInXMLResponse) SetBody(v *ModifyDBClusterConfigInXMLResponseBody) *ModifyDBClusterConfigInXMLResponse {
	s.Body = v
	return s
}

type ModifyDBClusterDescriptionRequest struct {
	// The cluster name. When you set the cluster name, take note of the following rules:
	//
	// 	- The cluster name cannot start with http:// or https://.
	//
	// 	- The cluster name must be 2 to 256 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// ClusterDescriptionTest
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
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
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The maintenance window of the cluster. Specify the time in the HH:mmZ-HH:mmZ format. The time must be in Coordinated Universal Time (UTC).
	//
	// For example, a value of 00:00Z-01:00Z indicates that routine maintenance can be performed on the cluster from 08:00 (UTC+8) to 09:00 (UTC+8).
	//
	// >  You can set the start time and end time of the maintenance window to the time on the hour, and the maintenance window is 1 hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00:00Z-01:00Z
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
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ModifyDBConfigRequest struct {
	// The dictionary configuration.
	//
	// example:
	//
	// {"name":"test"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1r59y779o04****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
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
	// The request ID.
	//
	// example:
	//
	// BF3844B6-1B12-57A0-A259-476D2079EE83
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyDBConfigResponse) SetStatusCode(v int32) *ModifyDBConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBConfigResponse) SetBody(v *ModifyDBConfigResponseBody) *ModifyDBConfigResponse {
	s.Body = v
	return s
}

type ModifyMinorVersionGreadeTypeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1qx68m06981****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	MaintainAutoType     *bool   `json:"MaintainAutoType,omitempty" xml:"MaintainAutoType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyMinorVersionGreadeTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMinorVersionGreadeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyMinorVersionGreadeTypeRequest) SetDBClusterId(v string) *ModifyMinorVersionGreadeTypeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyMinorVersionGreadeTypeRequest) SetMaintainAutoType(v bool) *ModifyMinorVersionGreadeTypeRequest {
	s.MaintainAutoType = &v
	return s
}

func (s *ModifyMinorVersionGreadeTypeRequest) SetOwnerAccount(v string) *ModifyMinorVersionGreadeTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyMinorVersionGreadeTypeRequest) SetOwnerId(v int64) *ModifyMinorVersionGreadeTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyMinorVersionGreadeTypeRequest) SetResourceOwnerAccount(v string) *ModifyMinorVersionGreadeTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyMinorVersionGreadeTypeRequest) SetResourceOwnerId(v int64) *ModifyMinorVersionGreadeTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyMinorVersionGreadeTypeResponseBody struct {
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMinorVersionGreadeTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMinorVersionGreadeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMinorVersionGreadeTypeResponseBody) SetRequestId(v string) *ModifyMinorVersionGreadeTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyMinorVersionGreadeTypeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMinorVersionGreadeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMinorVersionGreadeTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMinorVersionGreadeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyMinorVersionGreadeTypeResponse) SetHeaders(v map[string]*string) *ModifyMinorVersionGreadeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyMinorVersionGreadeTypeResponse) SetStatusCode(v int32) *ModifyMinorVersionGreadeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMinorVersionGreadeTypeResponse) SetBody(v *ModifyMinorVersionGreadeTypeResponseBody) *ModifyMinorVersionGreadeTypeResponse {
	s.Body = v
	return s
}

type ModifyRDSToClickhouseDbRequest struct {
	// The password of the account that is used to log on to the database in the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Aa
	CkPassword *string `json:"CkPassword,omitempty" xml:"CkPassword,omitempty"`
	// The account that is used to log on to the database in the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// user1
	CkUserName *string `json:"CkUserName,omitempty" xml:"CkUserName,omitempty"`
	// The port number of the ApsaraDB for ClickHouse cluster.
	//
	// example:
	//
	// 8123
	ClickhousePort *int64 `json:"ClickhousePort,omitempty" xml:"ClickhousePort,omitempty"`
	// The ID of the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp158i5wvj436****
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// The maximum number of rows that can be synchronized per second.
	//
	// example:
	//
	// 50000
	LimitUpper   *int64  `json:"LimitUpper,omitempty" xml:"LimitUpper,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6x3qq4t90ok****
	RdsId *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	// The password of the account that is used to log on to the database in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Rr
	RdsPassword *string `json:"RdsPassword,omitempty" xml:"RdsPassword,omitempty"`
	// The port number of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// 3306
	RdsPort *int64 `json:"RdsPort,omitempty" xml:"RdsPort,omitempty"`
	// The database in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// database
	RdsSynDb *string `json:"RdsSynDb,omitempty" xml:"RdsSynDb,omitempty"`
	// The table in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// table
	RdsSynTables *string `json:"RdsSynTables,omitempty" xml:"RdsSynTables,omitempty"`
	// The account that is used to log on to the database in the ApsaraDB RDS for MySQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// user2
	RdsUserName *string `json:"RdsUserName,omitempty" xml:"RdsUserName,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the ApsaraDB RDS for MySQL instance belongs.
	//
	// example:
	//
	// vpc-bp1v9dtwmqqjhwwg****
	RdsVpcId             *string `json:"RdsVpcId,omitempty" xml:"RdsVpcId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to ignore databases that do not support synchronization. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	SkipUnsupported *bool `json:"SkipUnsupported,omitempty" xml:"SkipUnsupported,omitempty"`
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
	// The error code.
	//
	// example:
	//
	// 0
	ErrorCode *int64 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 	- If the value **1*	- is returned for the **Status*	- parameter, the system does not return the ErrorMsg parameter.
	//
	// 	- If the value **0*	- is returned for the **Status*	- parameter, the ErrorMsg parameter returns the cause for the modification failure.
	//
	// example:
	//
	// ClickHouse exception, code: 49, host: 100.100.118.132, port: 49670; Code: 49, e.displayText() = DB::Exception: Logical error: there is no global context (version 20.8.17.25)n
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 746CD303-0B82-5E8D-886D-93A9FAF3A876
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the modification was successful. Valid values:
	//
	// 	- **1**: The modification was successful.
	//
	// 	- **0**: The modification failed.
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRDSToClickhouseDbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyRDSToClickhouseDbResponse) SetStatusCode(v int32) *ModifyRDSToClickhouseDbResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRDSToClickhouseDbResponse) SetBody(v *ModifyRDSToClickhouseDbResponseBody) *ModifyRDSToClickhouseDbResponse {
	s.Body = v
	return s
}

type ReleaseClusterPublicConnectionRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
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
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseClusterPublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The new password for the database account.
	//
	// >
	//
	// 	- The password must contain at least three types of the following characters: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- The password can contain the following special characters: ! @ # $ % ^ & \\	- ( ) _ + - =
	//
	// 	- The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Ff
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
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
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type RestartInstanceRequest struct {
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all the clusters that are deployed in a specific region. The information includes the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- 30 (default)
	//
	// 	- 50
	//
	// 	- 100
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The scheduled restart time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in Coordinated Universal Time (UTC).
	//
	// >  If this parameter is left empty or the time specified by this parameter is earlier than the current time, the cluster is immediately restarted.
	//
	// example:
	//
	// 2023-03-22T00:00:50Z
	RestartTime *string `json:"RestartTime,omitempty" xml:"RestartTime,omitempty"`
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

func (s *RestartInstanceRequest) SetRestartTime(v string) *RestartInstanceRequest {
	s.RestartTime = &v
	return s
}

type RestartInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *RestartInstanceResponse) SetStatusCode(v int32) *RestartInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartInstanceResponse) SetBody(v *RestartInstanceResponseBody) *RestartInstanceResponse {
	s.Body = v
	return s
}

type TransferVersionRequest struct {
	// The ID of the source ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1tm8zf130ew****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The database account that is used to log on to the database in the source ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	SourceAccount *string `json:"SourceAccount,omitempty" xml:"SourceAccount,omitempty"`
	// The password that corresponds to the database account for logging on to the database in the source ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Aa
	SourcePassword *string `json:"SourcePassword,omitempty" xml:"SourcePassword,omitempty"`
	// The database account that is used to log on to the database in the destination ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// test01
	TargetAccount *string `json:"TargetAccount,omitempty" xml:"TargetAccount,omitempty"`
	// The ID of the destination ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp14b39djx7zg****
	TargetDbClusterId *string `json:"TargetDbClusterId,omitempty" xml:"TargetDbClusterId,omitempty"`
	// The password that corresponds to the database account for logging on to the database in the destination ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Ff
	TargetPassword *string `json:"TargetPassword,omitempty" xml:"TargetPassword,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 7084CDB5-308F-5D0B-8F9B-5F7D83E09738
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransferVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferVersionResponseBody) GoString() string {
	return s.String()
}

func (s *TransferVersionResponseBody) SetRequestId(v string) *TransferVersionResponseBody {
	s.RequestId = &v
	return s
}

type TransferVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *TransferVersionResponse) SetStatusCode(v int32) *TransferVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferVersionResponse) SetBody(v *TransferVersionResponseBody) *TransferVersionResponse {
	s.Body = v
	return s
}

type UpgradeMinorVersionRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to update the minor engine version of the ApsaraDB for ClickHouse cluster immediately. Valid values:
	//
	// 	- **true**: updates the minor engine version of the ApsaraDB for ClickHouse cluster immediately.
	//
	// 	- **false**: updates the minor engine version of the ApsaraDB for ClickHouse cluster at the specified time or within the specified maintenance window.
	//
	// >  If you want to update the minor engine version of the ApsaraDB for ClickHouse cluster at the specified time, **UpgradeTime*	- is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	UpgradeImmediately *bool `json:"UpgradeImmediately,omitempty" xml:"UpgradeImmediately,omitempty"`
	// The update time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in Coordinated Universal Time (UTC).
	//
	// >  If you do not set this parameter, the minor engine version of an ApsaraDB for ClickHouse cluster is updated within the specified maintenance window.
	//
	// example:
	//
	// 2022-08-07T16:38Z
	UpgradeTime *string `json:"UpgradeTime,omitempty" xml:"UpgradeTime,omitempty"`
	// The minor engine version to which you want to update.
	//
	// >  By default, UpgradeVersion is not set and the minor engine version of the ApsaraDB for ClickHouse cluster is updated to the latest version.
	//
	// example:
	//
	// 1.37.0
	UpgradeVersion *string `json:"UpgradeVersion,omitempty" xml:"UpgradeVersion,omitempty"`
}

func (s UpgradeMinorVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMinorVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionRequest) SetDBClusterId(v string) *UpgradeMinorVersionRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetOwnerAccount(v string) *UpgradeMinorVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetOwnerId(v int64) *UpgradeMinorVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetResourceOwnerAccount(v string) *UpgradeMinorVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetResourceOwnerId(v int64) *UpgradeMinorVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetUpgradeImmediately(v bool) *UpgradeMinorVersionRequest {
	s.UpgradeImmediately = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetUpgradeTime(v string) *UpgradeMinorVersionRequest {
	s.UpgradeTime = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetUpgradeVersion(v string) *UpgradeMinorVersionRequest {
	s.UpgradeVersion = &v
	return s
}

type UpgradeMinorVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FE242962-6DA3-5FC8-9691-37B62A3210F7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeMinorVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMinorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionResponseBody) SetRequestId(v string) *UpgradeMinorVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeMinorVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeMinorVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeMinorVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMinorVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionResponse) SetHeaders(v map[string]*string) *UpgradeMinorVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeMinorVersionResponse) SetStatusCode(v int32) *UpgradeMinorVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeMinorVersionResponse) SetBody(v *UpgradeMinorVersionResponseBody) *UpgradeMinorVersionResponse {
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

// Summary:
//
// Creates a public endpoint for an ApsaraDB for ClickHouse cluster.
//
// @param request - AllocateClusterPublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateClusterPublicConnectionResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Creates a public endpoint for an ApsaraDB for ClickHouse cluster.
//
// @param request - AllocateClusterPublicConnectionRequest
//
// @return AllocateClusterPublicConnectionResponse
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

// Summary:
//
// Checks the connectivity between an ApsaraDB for ClickHouse cluster and an ApsaraDB RDS for MySQL instance.
//
// @param request - CheckClickhouseToRDSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckClickhouseToRDSResponse
func (client *Client) CheckClickhouseToRDSWithOptions(request *CheckClickhouseToRDSRequest, runtime *util.RuntimeOptions) (_result *CheckClickhouseToRDSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CkPassword)) {
		query["CkPassword"] = request.CkPassword
	}

	if !tea.BoolValue(util.IsUnset(request.CkUserName)) {
		query["CkUserName"] = request.CkUserName
	}

	if !tea.BoolValue(util.IsUnset(request.ClickhousePort)) {
		query["ClickhousePort"] = request.ClickhousePort
	}

	if !tea.BoolValue(util.IsUnset(request.DbClusterId)) {
		query["DbClusterId"] = request.DbClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RdsId)) {
		query["RdsId"] = request.RdsId
	}

	if !tea.BoolValue(util.IsUnset(request.RdsPassword)) {
		query["RdsPassword"] = request.RdsPassword
	}

	if !tea.BoolValue(util.IsUnset(request.RdsPort)) {
		query["RdsPort"] = request.RdsPort
	}

	if !tea.BoolValue(util.IsUnset(request.RdsUserName)) {
		query["RdsUserName"] = request.RdsUserName
	}

	if !tea.BoolValue(util.IsUnset(request.RdsVpcId)) {
		query["RdsVpcId"] = request.RdsVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.RdsVpcUrl)) {
		query["RdsVpcUrl"] = request.RdsVpcUrl
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
		Action:      tea.String("CheckClickhouseToRDS"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Checks the connectivity between an ApsaraDB for ClickHouse cluster and an ApsaraDB RDS for MySQL instance.
//
// @param request - CheckClickhouseToRDSRequest
//
// @return CheckClickhouseToRDSResponse
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

// Summary:
//
// Queries whether an ApsaraDB for ClickHouse cluster needs to be restarted after you change the values of the configuration parameters in XML mode.
//
// Description:
//
// >  You can call this operation only for ApsaraDB for ClickHouse clusters that were created after December 1, 2021.
//
// @param request - CheckModifyConfigNeedRestartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckModifyConfigNeedRestartResponse
func (client *Client) CheckModifyConfigNeedRestartWithOptions(request *CheckModifyConfigNeedRestartRequest, runtime *util.RuntimeOptions) (_result *CheckModifyConfigNeedRestartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		query["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckModifyConfigNeedRestart"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckModifyConfigNeedRestartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether an ApsaraDB for ClickHouse cluster needs to be restarted after you change the values of the configuration parameters in XML mode.
//
// Description:
//
// >  You can call this operation only for ApsaraDB for ClickHouse clusters that were created after December 1, 2021.
//
// @param request - CheckModifyConfigNeedRestartRequest
//
// @return CheckModifyConfigNeedRestartResponse
func (client *Client) CheckModifyConfigNeedRestart(request *CheckModifyConfigNeedRestartRequest) (_result *CheckModifyConfigNeedRestartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckModifyConfigNeedRestartResponse{}
	_body, _err := client.CheckModifyConfigNeedRestartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether the monitoring and alerting feature that is provided by Application Real-Time Monitoring Service (ARMS) is enabled for an ApsaraDB for ClickHouse cluster.
//
// @param request - CheckMonitorAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckMonitorAlertResponse
func (client *Client) CheckMonitorAlertWithOptions(request *CheckMonitorAlertRequest, runtime *util.RuntimeOptions) (_result *CheckMonitorAlertResponse, _err error) {
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
		Action:      tea.String("CheckMonitorAlert"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckMonitorAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the monitoring and alerting feature that is provided by Application Real-Time Monitoring Service (ARMS) is enabled for an ApsaraDB for ClickHouse cluster.
//
// @param request - CheckMonitorAlertRequest
//
// @return CheckMonitorAlertResponse
func (client *Client) CheckMonitorAlert(request *CheckMonitorAlertRequest) (_result *CheckMonitorAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckMonitorAlertResponse{}
	_body, _err := client.CheckMonitorAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs migration and scale-out detection on an ApsaraDB for ClickHouse cluster.
//
// @param request - CheckScaleOutBalancedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckScaleOutBalancedResponse
func (client *Client) CheckScaleOutBalancedWithOptions(request *CheckScaleOutBalancedRequest, runtime *util.RuntimeOptions) (_result *CheckScaleOutBalancedResponse, _err error) {
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
		Action:      tea.String("CheckScaleOutBalanced"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Performs migration and scale-out detection on an ApsaraDB for ClickHouse cluster.
//
// @param request - CheckScaleOutBalancedRequest
//
// @return CheckScaleOutBalancedResponse
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

// Summary:
//
// Queries the service-linked role of ApsaraDB for ClickHouse.
//
// @param request - CheckServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckServiceLinkedRoleResponse
func (client *Client) CheckServiceLinkedRoleWithOptions(request *CheckServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CheckServiceLinkedRoleResponse, _err error) {
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
		Action:      tea.String("CheckServiceLinkedRole"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Queries the service-linked role of ApsaraDB for ClickHouse.
//
// @param request - CheckServiceLinkedRoleRequest
//
// @return CheckServiceLinkedRoleResponse
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

// Summary:
//
// Creates a database account for an ApsaraDB for ClickHouse cluster.
//
// @param request - CreateAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Creates a database account for an ApsaraDB for ClickHouse cluster.
//
// @param request - CreateAccountRequest
//
// @return CreateAccountResponse
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

// @param request - CreateAccountAndAuthorityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountAndAuthorityResponse
func (client *Client) CreateAccountAndAuthorityWithOptions(request *CreateAccountAndAuthorityRequest, runtime *util.RuntimeOptions) (_result *CreateAccountAndAuthorityResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.AllowDatabases)) {
		query["AllowDatabases"] = request.AllowDatabases
	}

	if !tea.BoolValue(util.IsUnset(request.AllowDictionaries)) {
		query["AllowDictionaries"] = request.AllowDictionaries
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DdlAuthority)) {
		query["DdlAuthority"] = request.DdlAuthority
	}

	if !tea.BoolValue(util.IsUnset(request.DmlAuthority)) {
		query["DmlAuthority"] = request.DmlAuthority
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

	if !tea.BoolValue(util.IsUnset(request.TotalDatabases)) {
		query["TotalDatabases"] = request.TotalDatabases
	}

	if !tea.BoolValue(util.IsUnset(request.TotalDictionaries)) {
		query["TotalDictionaries"] = request.TotalDictionaries
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccountAndAuthority"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// @param request - CreateAccountAndAuthorityRequest
//
// @return CreateAccountAndAuthorityResponse
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

// Summary:
//
// Creates a backup policy.
//
// Description:
//
// >  This operation is available only for the ApsaraDB for ClickHouse clusters of versions 20.3, 20.8, and 21.8.
//
// @param request - CreateBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackupPolicyResponse
func (client *Client) CreateBackupPolicyWithOptions(request *CreateBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateBackupPolicyResponse, _err error) {
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
		Action:      tea.String("CreateBackupPolicy"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Creates a backup policy.
//
// Description:
//
// >  This operation is available only for the ApsaraDB for ClickHouse clusters of versions 20.3, 20.8, and 21.8.
//
// @param request - CreateBackupPolicyRequest
//
// @return CreateBackupPolicyResponse
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

// Summary:
//
// Creates an ApsaraDB for ClickHouse cluster.
//
// @param request - CreateDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBInstanceResponse
func (client *Client) CreateDBInstanceWithOptions(request *CreateDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.BackupSetID)) {
		query["BackupSetID"] = request.BackupSetID
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
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

	if !tea.BoolValue(util.IsUnset(request.DbNodeStorageType)) {
		query["DbNodeStorageType"] = request.DbNodeStorageType
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptionKey)) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptionType)) {
		query["EncryptionType"] = request.EncryptionType
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

	if !tea.BoolValue(util.IsUnset(request.SourceDBClusterId)) {
		query["SourceDBClusterId"] = request.SourceDBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.UsedTime)) {
		query["UsedTime"] = request.UsedTime
	}

	if !tea.BoolValue(util.IsUnset(request.VPCId)) {
		query["VPCId"] = request.VPCId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchBak)) {
		query["VSwitchBak"] = request.VSwitchBak
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchBak2)) {
		query["VSwitchBak2"] = request.VSwitchBak2
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZondIdBak2)) {
		query["ZondIdBak2"] = request.ZondIdBak2
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneIdBak)) {
		query["ZoneIdBak"] = request.ZoneIdBak
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDBInstance"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Creates an ApsaraDB for ClickHouse cluster.
//
// @param request - CreateDBInstanceRequest
//
// @return CreateDBInstanceResponse
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

// Summary:
//
// Creates a monitoring data report for an ApsaraDB for ClickHouse cluster.
//
// @param request - CreateMonitorDataReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorDataReportResponse
func (client *Client) CreateMonitorDataReportWithOptions(request *CreateMonitorDataReportRequest, runtime *util.RuntimeOptions) (_result *CreateMonitorDataReportResponse, _err error) {
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
		Action:      tea.String("CreateMonitorDataReport"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMonitorDataReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a monitoring data report for an ApsaraDB for ClickHouse cluster.
//
// @param request - CreateMonitorDataReportRequest
//
// @return CreateMonitorDataReportResponse
func (client *Client) CreateMonitorDataReport(request *CreateMonitorDataReportRequest) (_result *CreateMonitorDataReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMonitorDataReportResponse{}
	_body, _err := client.CreateMonitorDataReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a storage task for cold data.
//
// Description:
//
// Only an ApsaraDB for ClickHouse cluster of V20.8 or later supports tiered storage of hot data and cold data. If your data is in an ApsaraDB for ClickHouse cluster of a version earlier than V20.8 and you want to use tiered storage of hot data and cold data to store the data, you can migrate the data to an ApsaraDB for ClickHouse cluster of V20.8 or later and use tiered storage of hot data and cold data. For more information about how to migrate data between ApsaraDB for ClickHouse clusters, see [Migrate data between ApsaraDB for ClickHouse clusters](https://help.aliyun.com/document_detail/276926.html).
//
// @param request - CreateOSSStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOSSStorageResponse
func (client *Client) CreateOSSStorageWithOptions(request *CreateOSSStorageRequest, runtime *util.RuntimeOptions) (_result *CreateOSSStorageResponse, _err error) {
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
		Action:      tea.String("CreateOSSStorage"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Creates a storage task for cold data.
//
// Description:
//
// Only an ApsaraDB for ClickHouse cluster of V20.8 or later supports tiered storage of hot data and cold data. If your data is in an ApsaraDB for ClickHouse cluster of a version earlier than V20.8 and you want to use tiered storage of hot data and cold data to store the data, you can migrate the data to an ApsaraDB for ClickHouse cluster of V20.8 or later and use tiered storage of hot data and cold data. For more information about how to migrate data between ApsaraDB for ClickHouse clusters, see [Migrate data between ApsaraDB for ClickHouse clusters](https://help.aliyun.com/document_detail/276926.html).
//
// @param request - CreateOSSStorageRequest
//
// @return CreateOSSStorageResponse
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

// Summary:
//
// Enables the MySQL port for an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  For an ApsaraDB for ClickHouse cluster of V20.8 or later that was created before December 1, 2021, you must manually enable the MySQL port. For an ApsaraDB for ClickHouse cluster of V20.8 or later that was created after December 1, 2021, the MySQL port is automatically enabled.
//
// @param request - CreatePortsForClickHouseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePortsForClickHouseResponse
func (client *Client) CreatePortsForClickHouseWithOptions(request *CreatePortsForClickHouseRequest, runtime *util.RuntimeOptions) (_result *CreatePortsForClickHouseResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.PortType)) {
		query["PortType"] = request.PortType
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
		Action:      tea.String("CreatePortsForClickHouse"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Enables the MySQL port for an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  For an ApsaraDB for ClickHouse cluster of V20.8 or later that was created before December 1, 2021, you must manually enable the MySQL port. For an ApsaraDB for ClickHouse cluster of V20.8 or later that was created after December 1, 2021, the MySQL port is automatically enabled.
//
// @param request - CreatePortsForClickHouseRequest
//
// @return CreatePortsForClickHouseResponse
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

// Summary:
//
// Creates a task to synchronize data from an ApsaraDB RDS for MySQL instance to an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  This operation is only applicable to ApsaraDB for ClickHouse clusters.
//
// @param request - CreateRDSToClickhouseDbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRDSToClickhouseDbResponse
func (client *Client) CreateRDSToClickhouseDbWithOptions(request *CreateRDSToClickhouseDbRequest, runtime *util.RuntimeOptions) (_result *CreateRDSToClickhouseDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CkPassword)) {
		query["CkPassword"] = request.CkPassword
	}

	if !tea.BoolValue(util.IsUnset(request.CkUserName)) {
		query["CkUserName"] = request.CkUserName
	}

	if !tea.BoolValue(util.IsUnset(request.ClickhousePort)) {
		query["ClickhousePort"] = request.ClickhousePort
	}

	if !tea.BoolValue(util.IsUnset(request.DbClusterId)) {
		query["DbClusterId"] = request.DbClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.LimitUpper)) {
		query["LimitUpper"] = request.LimitUpper
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RdsId)) {
		query["RdsId"] = request.RdsId
	}

	if !tea.BoolValue(util.IsUnset(request.RdsPassword)) {
		query["RdsPassword"] = request.RdsPassword
	}

	if !tea.BoolValue(util.IsUnset(request.RdsPort)) {
		query["RdsPort"] = request.RdsPort
	}

	if !tea.BoolValue(util.IsUnset(request.RdsUserName)) {
		query["RdsUserName"] = request.RdsUserName
	}

	if !tea.BoolValue(util.IsUnset(request.RdsVpcId)) {
		query["RdsVpcId"] = request.RdsVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.RdsVpcUrl)) {
		query["RdsVpcUrl"] = request.RdsVpcUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SkipUnsupported)) {
		query["SkipUnsupported"] = request.SkipUnsupported
	}

	if !tea.BoolValue(util.IsUnset(request.SynDbTables)) {
		query["SynDbTables"] = request.SynDbTables
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRDSToClickhouseDb"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Creates a task to synchronize data from an ApsaraDB RDS for MySQL instance to an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  This operation is only applicable to ApsaraDB for ClickHouse clusters.
//
// @param request - CreateRDSToClickhouseDbRequest
//
// @return CreateRDSToClickhouseDbResponse
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

// Summary:
//
// Creates a privileged account or a standard account for an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  This operation is applicable only to ApsaraDB for ClickHouse clusters of V20.8 or later that were created after December 1, 2021,
//
// @param request - CreateSQLAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSQLAccountResponse
func (client *Client) CreateSQLAccountWithOptions(request *CreateSQLAccountRequest, runtime *util.RuntimeOptions) (_result *CreateSQLAccountResponse, _err error) {
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
		Action:      tea.String("CreateSQLAccount"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSQLAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a privileged account or a standard account for an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  This operation is applicable only to ApsaraDB for ClickHouse clusters of V20.8 or later that were created after December 1, 2021,
//
// @param request - CreateSQLAccountRequest
//
// @return CreateSQLAccountResponse
func (client *Client) CreateSQLAccount(request *CreateSQLAccountRequest) (_result *CreateSQLAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSQLAccountResponse{}
	_body, _err := client.CreateSQLAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a service-linked role.
//
// @param request - CreateServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceLinkedRoleResponse
func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
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
		Action:      tea.String("CreateServiceLinkedRole"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Creates a service-linked role.
//
// @param request - CreateServiceLinkedRoleRequest
//
// @return CreateServiceLinkedRoleResponse
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

// Summary:
//
// Deletes a database account of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  After you delete a database account, you cannot use the account to log on to the ApsaraDB for ClickHouse cluster. Exercise caution when performing this operation.
//
// @param request - DeleteAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("DeleteAccount"),
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Deletes a database account of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  After you delete a database account, you cannot use the account to log on to the ApsaraDB for ClickHouse cluster. Exercise caution when performing this operation.
//
// @param request - DeleteAccountRequest
//
// @return DeleteAccountResponse
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

// Summary:
//
// Releases a pay-as-you-go ApsaraDB for ClickHouse cluster.
//
// Description:
//
// *Warning*	- After an ApsaraDB for ClickHouse cluster is deleted, all data in the cluster is deleted and cannot be recovered. Exercise caution when performing this operation.
//
// @param request - DeleteDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBClusterResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Releases a pay-as-you-go ApsaraDB for ClickHouse cluster.
//
// Description:
//
// *Warning*	- After an ApsaraDB for ClickHouse cluster is deleted, all data in the cluster is deleted and cannot be recovered. Exercise caution when performing this operation.
//
// @param request - DeleteDBClusterRequest
//
// @return DeleteDBClusterResponse
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

// Summary:
//
// Deletes a database used for data synchronization.
//
// @param request - DeleteSyndbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSyndbResponse
func (client *Client) DeleteSyndbWithOptions(request *DeleteSyndbRequest, runtime *util.RuntimeOptions) (_result *DeleteSyndbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbClusterId)) {
		query["DbClusterId"] = request.DbClusterId
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

	if !tea.BoolValue(util.IsUnset(request.SynDb)) {
		query["SynDb"] = request.SynDb
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSyndb"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Deletes a database used for data synchronization.
//
// @param request - DeleteSyndbRequest
//
// @return DeleteSyndbResponse
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

// Summary:
//
// Queries the permissions of an account.
//
// @param request - DescribeAccountAuthorityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountAuthorityResponse
func (client *Client) DescribeAccountAuthorityWithOptions(request *DescribeAccountAuthorityRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountAuthorityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("DescribeAccountAuthority"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Queries the permissions of an account.
//
// @param request - DescribeAccountAuthorityRequest
//
// @return DescribeAccountAuthorityResponse
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

// Summary:
//
// DescribeAccounts
//
// @param request - DescribeAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountsResponse
func (client *Client) DescribeAccountsWithOptions(request *DescribeAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccounts"),
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// DescribeAccounts
//
// @param request - DescribeAccountsRequest
//
// @return DescribeAccountsResponse
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

// Summary:
//
// Queries a list of databases, tables, and columns in an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeAllDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAllDataSourceResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Queries a list of databases, tables, and columns in an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeAllDataSourceRequest
//
// @return DescribeAllDataSourceResponse
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

// Summary:
//
// Queries the data sources of an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeAllDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAllDataSourcesResponse
func (client *Client) DescribeAllDataSourcesWithOptions(request *DescribeAllDataSourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeAllDataSourcesResponse, _err error) {
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
		Action:      tea.String("DescribeAllDataSources"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Queries the data sources of an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeAllDataSourcesRequest
//
// @return DescribeAllDataSourcesResponse
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

// Summary:
//
// Queries the backup settings of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  This operation is available only for the ApsaraDB for ClickHouse clusters of versions 20.3, 20.8, and 21.8.
//
// @param request - DescribeBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPolicyResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Queries the backup settings of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  This operation is available only for the ApsaraDB for ClickHouse clusters of versions 20.3, 20.8, and 21.8.
//
// @param request - DescribeBackupPolicyRequest
//
// @return DescribeBackupPolicyResponse
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

// @param request - DescribeBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupsResponse
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
		Version:     tea.String("2019-11-11"),
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

// @param request - DescribeBackupsRequest
//
// @return DescribeBackupsResponse
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

// Summary:
//
// Queries information about columns.
//
// @param request - DescribeColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColumnsResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Queries information about columns.
//
// @param request - DescribeColumnsRequest
//
// @return DescribeColumnsResponse
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

// Summary:
//
// Queries the change records of the configuration parameters of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  You can call this operation only for ApsaraDB for ClickHouse clusters that were created after December 1, 2021.
//
// @param request - DescribeConfigHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConfigHistoryResponse
func (client *Client) DescribeConfigHistoryWithOptions(request *DescribeConfigHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeConfigHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeConfigHistory"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeConfigHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the change records of the configuration parameters of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  You can call this operation only for ApsaraDB for ClickHouse clusters that were created after December 1, 2021.
//
// @param request - DescribeConfigHistoryRequest
//
// @return DescribeConfigHistoryResponse
func (client *Client) DescribeConfigHistory(request *DescribeConfigHistoryRequest) (_result *DescribeConfigHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConfigHistoryResponse{}
	_body, _err := client.DescribeConfigHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the values of the configuration parameters of an ApsaraDB for ClickHouse cluster before and after the values of the configuration parameters are changed.
//
// Description:
//
// >  You can call this operation only for ApsaraDB for ClickHouse clusters that were created after December 1, 2021.
//
// @param request - DescribeConfigVersionDifferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConfigVersionDifferenceResponse
func (client *Client) DescribeConfigVersionDifferenceWithOptions(request *DescribeConfigVersionDifferenceRequest, runtime *util.RuntimeOptions) (_result *DescribeConfigVersionDifferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeConfigVersionDifference"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeConfigVersionDifferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the values of the configuration parameters of an ApsaraDB for ClickHouse cluster before and after the values of the configuration parameters are changed.
//
// Description:
//
// >  You can call this operation only for ApsaraDB for ClickHouse clusters that were created after December 1, 2021.
//
// @param request - DescribeConfigVersionDifferenceRequest
//
// @return DescribeConfigVersionDifferenceResponse
func (client *Client) DescribeConfigVersionDifference(request *DescribeConfigVersionDifferenceRequest) (_result *DescribeConfigVersionDifferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConfigVersionDifferenceResponse{}
	_body, _err := client.DescribeConfigVersionDifferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IP address whitelist of an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeDBClusterAccessWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterAccessWhiteListResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Queries the IP address whitelist of an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeDBClusterAccessWhiteListRequest
//
// @return DescribeDBClusterAccessWhiteListResponse
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

// Summary:
//
// Queries the information about an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeDBClusterAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterAttributeResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Queries the information about an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeDBClusterAttributeRequest
//
// @return DescribeDBClusterAttributeResponse
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

// Summary:
//
// Queries information about the parameter settings of an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeDBClusterConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterConfigResponse
func (client *Client) DescribeDBClusterConfigWithOptions(request *DescribeDBClusterConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterConfigResponse, _err error) {
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
		Action:      tea.String("DescribeDBClusterConfig"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Queries information about the parameter settings of an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeDBClusterConfigRequest
//
// @return DescribeDBClusterConfigResponse
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

// Summary:
//
// Queries the values of the configuration parameters in the config.xml file of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  You can call this operation only for ApsaraDB for ClickHouse clusters that were created after December 1, 2021.
//
// @param request - DescribeDBClusterConfigInXMLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterConfigInXMLResponse
func (client *Client) DescribeDBClusterConfigInXMLWithOptions(request *DescribeDBClusterConfigInXMLRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterConfigInXMLResponse, _err error) {
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
		Action:      tea.String("DescribeDBClusterConfigInXML"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBClusterConfigInXMLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the values of the configuration parameters in the config.xml file of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  You can call this operation only for ApsaraDB for ClickHouse clusters that were created after December 1, 2021.
//
// @param request - DescribeDBClusterConfigInXMLRequest
//
// @return DescribeDBClusterConfigInXMLResponse
func (client *Client) DescribeDBClusterConfigInXML(request *DescribeDBClusterConfigInXMLRequest) (_result *DescribeDBClusterConfigInXMLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterConfigInXMLResponse{}
	_body, _err := client.DescribeDBClusterConfigInXMLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the network information about an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeDBClusterNetInfoItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterNetInfoItemsResponse
func (client *Client) DescribeDBClusterNetInfoItemsWithOptions(request *DescribeDBClusterNetInfoItemsRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterNetInfoItemsResponse, _err error) {
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
		Action:      tea.String("DescribeDBClusterNetInfoItems"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Queries the network information about an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeDBClusterNetInfoItemsRequest
//
// @return DescribeDBClusterNetInfoItemsResponse
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

// Summary:
//
// Queries performance data about an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// You can query the performance data of a specified cluster over a specific time range based on the performance metrics. The data is collected every 30 seconds.
//
// >  You can call this operation only for ApsaraDB for ClickHouse clusters that were created before December 1, 2021.
//
// @param request - DescribeDBClusterPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterPerformanceResponse
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
		Action:      tea.String("DescribeDBClusterPerformance"),
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Queries performance data about an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// You can query the performance data of a specified cluster over a specific time range based on the performance metrics. The data is collected every 30 seconds.
//
// >  You can call this operation only for ApsaraDB for ClickHouse clusters that were created before December 1, 2021.
//
// @param request - DescribeDBClusterPerformanceRequest
//
// @return DescribeDBClusterPerformanceResponse
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

// Summary:
//
// Queries the information about ApsaraDB for ClickHouse clusters in a region.
//
// @param request - DescribeDBClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClustersResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Queries the information about ApsaraDB for ClickHouse clusters in a region.
//
// @param request - DescribeDBClustersRequest
//
// @return DescribeDBClustersResponse
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

// Summary:
//
// Queries configuration information about an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeDBConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBConfigResponse
func (client *Client) DescribeDBConfigWithOptions(request *DescribeDBConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDBConfigResponse, _err error) {
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
		Action:      tea.String("DescribeDBConfig"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Queries configuration information about an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeDBConfigRequest
//
// @return DescribeDBConfigResponse
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

// Summary:
//
// Queries the storage of cold data.
//
// @param request - DescribeOSSStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOSSStorageResponse
func (client *Client) DescribeOSSStorageWithOptions(request *DescribeOSSStorageRequest, runtime *util.RuntimeOptions) (_result *DescribeOSSStorageResponse, _err error) {
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
		Action:      tea.String("DescribeOSSStorage"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Queries the storage of cold data.
//
// @param request - DescribeOSSStorageRequest
//
// @return DescribeOSSStorageResponse
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

// Summary:
//
// Queries the details of queries that are being executed in an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeProcessListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProcessListResponse
func (client *Client) DescribeProcessListWithOptions(request *DescribeProcessListRequest, runtime *util.RuntimeOptions) (_result *DescribeProcessListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.InitialQueryId)) {
		query["InitialQueryId"] = request.InitialQueryId
	}

	if !tea.BoolValue(util.IsUnset(request.InitialUser)) {
		query["InitialUser"] = request.InitialUser
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

	if !tea.BoolValue(util.IsUnset(request.QueryDurationMs)) {
		query["QueryDurationMs"] = request.QueryDurationMs
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
		Action:      tea.String("DescribeProcessList"),
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Queries the details of queries that are being executed in an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeProcessListRequest
//
// @return DescribeProcessListResponse
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

// Summary:
//
// Queries the information about all regions and zones of ApsaraDB for ClickHouse clusters.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Queries the information about all regions and zones of ApsaraDB for ClickHouse clusters.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
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

// Summary:
//
// Queries a list of all databases in an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeSchemasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSchemasResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Queries a list of all databases in an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeSchemasRequest
//
// @return DescribeSchemasResponse
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

// Summary:
//
// Queries the details about slow query logs.
//
// @param request - DescribeSlowLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowLogRecordsResponse
func (client *Client) DescribeSlowLogRecordsWithOptions(request *DescribeSlowLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowLogRecordsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.QueryDurationMs)) {
		query["QueryDurationMs"] = request.QueryDurationMs
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlowLogRecords"),
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Queries the details about slow query logs.
//
// @param request - DescribeSlowLogRecordsRequest
//
// @return DescribeSlowLogRecordsResponse
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

// Summary:
//
// Queries information about tables that are synchronized from an ApsaraDB RDS for MySQL instance to an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeSynDbTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSynDbTablesResponse
func (client *Client) DescribeSynDbTablesWithOptions(request *DescribeSynDbTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeSynDbTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbClusterId)) {
		query["DbClusterId"] = request.DbClusterId
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

	if !tea.BoolValue(util.IsUnset(request.SynDb)) {
		query["SynDb"] = request.SynDb
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSynDbTables"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Queries information about tables that are synchronized from an ApsaraDB RDS for MySQL instance to an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeSynDbTablesRequest
//
// @return DescribeSynDbTablesResponse
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

// Summary:
//
// Queries the information about data synchronization between an ApsaraDB for ClickHouse cluster and an ApsaraDB RDS for MySQL instance.
//
// @param request - DescribeSynDbsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSynDbsResponse
func (client *Client) DescribeSynDbsWithOptions(request *DescribeSynDbsRequest, runtime *util.RuntimeOptions) (_result *DescribeSynDbsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbClusterId)) {
		query["DbClusterId"] = request.DbClusterId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSynDbs"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Queries the information about data synchronization between an ApsaraDB for ClickHouse cluster and an ApsaraDB RDS for MySQL instance.
//
// @param request - DescribeSynDbsRequest
//
// @return DescribeSynDbsResponse
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

// Summary:
//
// Queries the information about tables in a database of an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTablesResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Queries the information about tables in a database of an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeTablesRequest
//
// @return DescribeTablesResponse
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

// Summary:
//
// Queries information about data migration from an ApsaraDB for ClickHouse cluster of an earlier version to an ApsaraDB for ClickHouse cluster of a later version
//
// Description:
//
// >  You can call this operation to query information about only data migration from an ApsaraDB for ClickHouse cluster of an earlier version to an ApsaraDB for ClickHouse cluster of a later version.
//
// @param request - DescribeTransferHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTransferHistoryResponse
func (client *Client) DescribeTransferHistoryWithOptions(request *DescribeTransferHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeTransferHistoryResponse, _err error) {
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
		Action:      tea.String("DescribeTransferHistory"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Queries information about data migration from an ApsaraDB for ClickHouse cluster of an earlier version to an ApsaraDB for ClickHouse cluster of a later version
//
// Description:
//
// >  You can call this operation to query information about only data migration from an ApsaraDB for ClickHouse cluster of an earlier version to an ApsaraDB for ClickHouse cluster of a later version.
//
// @param request - DescribeTransferHistoryRequest
//
// @return DescribeTransferHistoryResponse
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

// Summary:
//
// Terminates an ongoing task.
//
// @param request - KillProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KillProcessResponse
func (client *Client) KillProcessWithOptions(request *KillProcessRequest, runtime *util.RuntimeOptions) (_result *KillProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.InitialQueryId)) {
		query["InitialQueryId"] = request.InitialQueryId
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
		Action:      tea.String("KillProcess"),
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Terminates an ongoing task.
//
// @param request - KillProcessRequest
//
// @return KillProcessResponse
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

// Summary:
//
// Modifies the permissions of an account.
//
// @param request - ModifyAccountAuthorityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountAuthorityResponse
func (client *Client) ModifyAccountAuthorityWithOptions(request *ModifyAccountAuthorityRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountAuthorityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AllowDatabases)) {
		query["AllowDatabases"] = request.AllowDatabases
	}

	if !tea.BoolValue(util.IsUnset(request.AllowDictionaries)) {
		query["AllowDictionaries"] = request.AllowDictionaries
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DdlAuthority)) {
		query["DdlAuthority"] = request.DdlAuthority
	}

	if !tea.BoolValue(util.IsUnset(request.DmlAuthority)) {
		query["DmlAuthority"] = request.DmlAuthority
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

	if !tea.BoolValue(util.IsUnset(request.TotalDatabases)) {
		query["TotalDatabases"] = request.TotalDatabases
	}

	if !tea.BoolValue(util.IsUnset(request.TotalDictionaries)) {
		query["TotalDictionaries"] = request.TotalDictionaries
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountAuthority"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Modifies the permissions of an account.
//
// @param request - ModifyAccountAuthorityRequest
//
// @return ModifyAccountAuthorityResponse
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

// @param request - ModifyAccountDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountDescriptionResponse
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
		Version:     tea.String("2019-11-11"),
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

// @param request - ModifyAccountDescriptionRequest
//
// @return ModifyAccountDescriptionResponse
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

// Summary:
//
// Modifies the backup settings of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  This operation is available only for the ApsaraDB for ClickHouse clusters of versions 20.3, 20.8, and 21.8.
//
// @param request - ModifyBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupPolicyResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Modifies the backup settings of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  This operation is available only for the ApsaraDB for ClickHouse clusters of versions 20.3, 20.8, and 21.8.
//
// @param request - ModifyBackupPolicyRequest
//
// @return ModifyBackupPolicyResponse
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

// Summary:
//
// Upgrades or downgrades an ApsaraDB for ClickHouse cluster.
//
// @param request - ModifyDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterResponse
func (client *Client) ModifyDBClusterWithOptions(request *ModifyDBClusterRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterClass)) {
		query["DBClusterClass"] = request.DBClusterClass
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeGroupCount)) {
		query["DBNodeGroupCount"] = request.DBNodeGroupCount
	}

	if !tea.BoolValue(util.IsUnset(request.DBNodeStorage)) {
		query["DBNodeStorage"] = request.DBNodeStorage
	}

	if !tea.BoolValue(util.IsUnset(request.DbNodeStorageType)) {
		query["DbNodeStorageType"] = request.DbNodeStorageType
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
		Action:      tea.String("ModifyDBCluster"),
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Upgrades or downgrades an ApsaraDB for ClickHouse cluster.
//
// @param request - ModifyDBClusterRequest
//
// @return ModifyDBClusterResponse
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

// @param request - ModifyDBClusterAccessWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterAccessWhiteListResponse
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
		Version:     tea.String("2019-11-11"),
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

// @param request - ModifyDBClusterAccessWhiteListRequest
//
// @return ModifyDBClusterAccessWhiteListResponse
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

// Summary:
//
// Modifies the configurations of an ApsaraDB for ClickHouse cluster.
//
// @param request - ModifyDBClusterConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterConfigResponse
func (client *Client) ModifyDBClusterConfigWithOptions(request *ModifyDBClusterConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterConfigResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		query["Reason"] = request.Reason
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

	if !tea.BoolValue(util.IsUnset(request.UserConfig)) {
		query["UserConfig"] = request.UserConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterConfig"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Modifies the configurations of an ApsaraDB for ClickHouse cluster.
//
// @param request - ModifyDBClusterConfigRequest
//
// @return ModifyDBClusterConfigResponse
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

// Summary:
//
// Modifies the settings of the configuration parameters for an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  You can call this operation only for ApsaraDB for ClickHouse clusters that were created after December 1, 2021.
//
// @param request - ModifyDBClusterConfigInXMLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterConfigInXMLResponse
func (client *Client) ModifyDBClusterConfigInXMLWithOptions(request *ModifyDBClusterConfigInXMLRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterConfigInXMLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		query["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		query["Reason"] = request.Reason
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBClusterConfigInXML"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBClusterConfigInXMLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the settings of the configuration parameters for an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  You can call this operation only for ApsaraDB for ClickHouse clusters that were created after December 1, 2021.
//
// @param request - ModifyDBClusterConfigInXMLRequest
//
// @return ModifyDBClusterConfigInXMLResponse
func (client *Client) ModifyDBClusterConfigInXML(request *ModifyDBClusterConfigInXMLRequest) (_result *ModifyDBClusterConfigInXMLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterConfigInXMLResponse{}
	_body, _err := client.ModifyDBClusterConfigInXMLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the name of an ApsaraDB for ClickHouse cluster.
//
// @param request - ModifyDBClusterDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterDescriptionResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Changes the name of an ApsaraDB for ClickHouse cluster.
//
// @param request - ModifyDBClusterDescriptionRequest
//
// @return ModifyDBClusterDescriptionResponse
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

// Summary:
//
// Modifies the maintenance window of an ApsaraDB for ClickHouse cluster.
//
// @param request - ModifyDBClusterMaintainTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterMaintainTimeResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Modifies the maintenance window of an ApsaraDB for ClickHouse cluster.
//
// @param request - ModifyDBClusterMaintainTimeRequest
//
// @return ModifyDBClusterMaintainTimeResponse
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

// Summary:
//
// Modifies the dictionary configuration of an ApsaraDB for ClickHouse cluster.
//
// @param request - ModifyDBConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBConfigResponse
func (client *Client) ModifyDBConfigWithOptions(request *ModifyDBConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyDBConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		query["Config"] = request.Config
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
		Action:      tea.String("ModifyDBConfig"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Modifies the dictionary configuration of an ApsaraDB for ClickHouse cluster.
//
// @param request - ModifyDBConfigRequest
//
// @return ModifyDBConfigResponse
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

// Summary:
//
// Modifies the type of a minor version update in ApsaraDB for ClickHouse.
//
// @param request - ModifyMinorVersionGreadeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMinorVersionGreadeTypeResponse
func (client *Client) ModifyMinorVersionGreadeTypeWithOptions(request *ModifyMinorVersionGreadeTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyMinorVersionGreadeTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBClusterId)) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MaintainAutoType)) {
		query["MaintainAutoType"] = request.MaintainAutoType
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
		Action:      tea.String("ModifyMinorVersionGreadeType"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMinorVersionGreadeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the type of a minor version update in ApsaraDB for ClickHouse.
//
// @param request - ModifyMinorVersionGreadeTypeRequest
//
// @return ModifyMinorVersionGreadeTypeResponse
func (client *Client) ModifyMinorVersionGreadeType(request *ModifyMinorVersionGreadeTypeRequest) (_result *ModifyMinorVersionGreadeTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMinorVersionGreadeTypeResponse{}
	_body, _err := client.ModifyMinorVersionGreadeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the synchronization task of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  This operation is applicable only to ApsaraDB for ClickHouse clusters.
//
// @param request - ModifyRDSToClickhouseDbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRDSToClickhouseDbResponse
func (client *Client) ModifyRDSToClickhouseDbWithOptions(request *ModifyRDSToClickhouseDbRequest, runtime *util.RuntimeOptions) (_result *ModifyRDSToClickhouseDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CkPassword)) {
		query["CkPassword"] = request.CkPassword
	}

	if !tea.BoolValue(util.IsUnset(request.CkUserName)) {
		query["CkUserName"] = request.CkUserName
	}

	if !tea.BoolValue(util.IsUnset(request.ClickhousePort)) {
		query["ClickhousePort"] = request.ClickhousePort
	}

	if !tea.BoolValue(util.IsUnset(request.DbClusterId)) {
		query["DbClusterId"] = request.DbClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.LimitUpper)) {
		query["LimitUpper"] = request.LimitUpper
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RdsId)) {
		query["RdsId"] = request.RdsId
	}

	if !tea.BoolValue(util.IsUnset(request.RdsPassword)) {
		query["RdsPassword"] = request.RdsPassword
	}

	if !tea.BoolValue(util.IsUnset(request.RdsPort)) {
		query["RdsPort"] = request.RdsPort
	}

	if !tea.BoolValue(util.IsUnset(request.RdsSynDb)) {
		query["RdsSynDb"] = request.RdsSynDb
	}

	if !tea.BoolValue(util.IsUnset(request.RdsSynTables)) {
		query["RdsSynTables"] = request.RdsSynTables
	}

	if !tea.BoolValue(util.IsUnset(request.RdsUserName)) {
		query["RdsUserName"] = request.RdsUserName
	}

	if !tea.BoolValue(util.IsUnset(request.RdsVpcId)) {
		query["RdsVpcId"] = request.RdsVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SkipUnsupported)) {
		query["SkipUnsupported"] = request.SkipUnsupported
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyRDSToClickhouseDb"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Modifies the synchronization task of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  This operation is applicable only to ApsaraDB for ClickHouse clusters.
//
// @param request - ModifyRDSToClickhouseDbRequest
//
// @return ModifyRDSToClickhouseDbResponse
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

// Summary:
//
// Releases the public endpoint of an ApsaraDB for ClickHouse cluster.
//
// @param request - ReleaseClusterPublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseClusterPublicConnectionResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Releases the public endpoint of an ApsaraDB for ClickHouse cluster.
//
// @param request - ReleaseClusterPublicConnectionRequest
//
// @return ReleaseClusterPublicConnectionResponse
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

// Summary:
//
// Resets the password of a database account for an ApsaraDB for ClickHouse cluster.
//
// @param request - ResetAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAccountPasswordResponse
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
		Version:     tea.String("2019-11-11"),
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

// Summary:
//
// Resets the password of a database account for an ApsaraDB for ClickHouse cluster.
//
// @param request - ResetAccountPasswordRequest
//
// @return ResetAccountPasswordResponse
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

// Summary:
//
// Restarts an ApsaraDB for ClickHouse cluster.
//
// @param request - RestartInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartInstanceResponse
func (client *Client) RestartInstanceWithOptions(request *RestartInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RestartTime)) {
		query["RestartTime"] = request.RestartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartInstance"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Restarts an ApsaraDB for ClickHouse cluster.
//
// @param request - RestartInstanceRequest
//
// @return RestartInstanceResponse
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

// Summary:
//
// Migrates the data of a source ApsaraDB for ClickHouse cluster to a destination ApsaraDB for ClickHouse cluster.
//
// Description:
//
// ## [](#)Prerequisites
//
// 	- The IP address of the source ApsaraDB for ClickHouse cluster is added to the IP address whitelist of the destination ApsaraDB for ClickHouse cluster.
//
// 	- The IP address of the destination ApsaraDB for ClickHouse cluster is added to the IP address whitelist of the source ApsaraDB for ClickHouse cluster.
//
// >  You can execute the `select 	- from system.clusters;` statement to query the IP address of an ApsaraDB for ClickHouse cluster.
//
// @param request - TransferVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferVersionResponse
func (client *Client) TransferVersionWithOptions(request *TransferVersionRequest, runtime *util.RuntimeOptions) (_result *TransferVersionResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SourceAccount)) {
		query["SourceAccount"] = request.SourceAccount
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePassword)) {
		query["SourcePassword"] = request.SourcePassword
	}

	if !tea.BoolValue(util.IsUnset(request.TargetAccount)) {
		query["TargetAccount"] = request.TargetAccount
	}

	if !tea.BoolValue(util.IsUnset(request.TargetDbClusterId)) {
		query["TargetDbClusterId"] = request.TargetDbClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetPassword)) {
		query["TargetPassword"] = request.TargetPassword
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferVersion"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

// Summary:
//
// Migrates the data of a source ApsaraDB for ClickHouse cluster to a destination ApsaraDB for ClickHouse cluster.
//
// Description:
//
// ## [](#)Prerequisites
//
// 	- The IP address of the source ApsaraDB for ClickHouse cluster is added to the IP address whitelist of the destination ApsaraDB for ClickHouse cluster.
//
// 	- The IP address of the destination ApsaraDB for ClickHouse cluster is added to the IP address whitelist of the source ApsaraDB for ClickHouse cluster.
//
// >  You can execute the `select 	- from system.clusters;` statement to query the IP address of an ApsaraDB for ClickHouse cluster.
//
// @param request - TransferVersionRequest
//
// @return TransferVersionResponse
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

// Summary:
//
// Updates the minor engine version of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  You can call this operation only for ApsaraDB for ClickHouse clusters that were purchased after December 1, 2021.
//
// @param request - UpgradeMinorVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeMinorVersionResponse
func (client *Client) UpgradeMinorVersionWithOptions(request *UpgradeMinorVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeMinorVersionResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.UpgradeImmediately)) {
		query["UpgradeImmediately"] = request.UpgradeImmediately
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeTime)) {
		query["UpgradeTime"] = request.UpgradeTime
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeVersion)) {
		query["UpgradeVersion"] = request.UpgradeVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeMinorVersion"),
		Version:     tea.String("2019-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeMinorVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the minor engine version of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >  You can call this operation only for ApsaraDB for ClickHouse clusters that were purchased after December 1, 2021.
//
// @param request - UpgradeMinorVersionRequest
//
// @return UpgradeMinorVersionResponse
func (client *Client) UpgradeMinorVersion(request *UpgradeMinorVersionRequest) (_result *UpgradeMinorVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeMinorVersionResponse{}
	_body, _err := client.UpgradeMinorVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

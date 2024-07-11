// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddUserHdfsInfoRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16o0pd5****582s
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"dfs.nameservices":"hdfs_test","dfs.ha.namenodes":"nn1,nn2","dfs.namenode.http-address.hdfs_test.nn1":"TEST-xxx1.com:50070","dfs.namenode.http-address.hdfs_test.nn2":"TEST-xxx2.com:50070","dfs.namenode.rpc-address.hdfs_test.nn1":"TEST-xxx1.com:8020","dfs.namenode.rpc-address.hdfs_test.nn2":"TEST-xxx2.com:8020"}
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s AddUserHdfsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserHdfsInfoRequest) GoString() string {
	return s.String()
}

func (s *AddUserHdfsInfoRequest) SetClientToken(v string) *AddUserHdfsInfoRequest {
	s.ClientToken = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetClusterId(v string) *AddUserHdfsInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *AddUserHdfsInfoRequest) SetExtInfo(v string) *AddUserHdfsInfoRequest {
	s.ExtInfo = &v
	return s
}

type AddUserHdfsInfoResponseBody struct {
	// example:
	//
	// FB0B7918-198C-46A8-AB9B-FE15403B1F0A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserHdfsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserHdfsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserHdfsInfoResponseBody) SetRequestId(v string) *AddUserHdfsInfoResponseBody {
	s.RequestId = &v
	return s
}

type AddUserHdfsInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserHdfsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserHdfsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserHdfsInfoResponse) GoString() string {
	return s.String()
}

func (s *AddUserHdfsInfoResponse) SetHeaders(v map[string]*string) *AddUserHdfsInfoResponse {
	s.Headers = v
	return s
}

func (s *AddUserHdfsInfoResponse) SetStatusCode(v int32) *AddUserHdfsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserHdfsInfoResponse) SetBody(v *AddUserHdfsInfoResponseBody) *AddUserHdfsInfoResponse {
	s.Body = v
	return s
}

type AllocatePublicNetworkAddressRequest struct {
	// example:
	//
	// 83b2b5e117a5b8bce0fae88d90576a84_6452320_82718582
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s AllocatePublicNetworkAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocatePublicNetworkAddressRequest) SetClientToken(v string) *AllocatePublicNetworkAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetClusterId(v string) *AllocatePublicNetworkAddressRequest {
	s.ClusterId = &v
	return s
}

type AllocatePublicNetworkAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocatePublicNetworkAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocatePublicNetworkAddressResponseBody) SetRequestId(v string) *AllocatePublicNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

type AllocatePublicNetworkAddressResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocatePublicNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocatePublicNetworkAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *AllocatePublicNetworkAddressResponse) SetHeaders(v map[string]*string) *AllocatePublicNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *AllocatePublicNetworkAddressResponse) SetStatusCode(v int32) *AllocatePublicNetworkAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocatePublicNetworkAddressResponse) SetBody(v *AllocatePublicNetworkAddressResponseBody) *AllocatePublicNetworkAddressResponse {
	s.Body = v
	return s
}

type CancelActiveOperationTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11111,22222
	Ids                  *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CancelActiveOperationTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelActiveOperationTasksRequest) GoString() string {
	return s.String()
}

func (s *CancelActiveOperationTasksRequest) SetIds(v string) *CancelActiveOperationTasksRequest {
	s.Ids = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetOwnerAccount(v string) *CancelActiveOperationTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetOwnerId(v int64) *CancelActiveOperationTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetResourceOwnerAccount(v string) *CancelActiveOperationTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetResourceOwnerId(v int64) *CancelActiveOperationTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelActiveOperationTasksRequest) SetSecurityToken(v string) *CancelActiveOperationTasksRequest {
	s.SecurityToken = &v
	return s
}

type CancelActiveOperationTasksResponseBody struct {
	// example:
	//
	// 11111,22222
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// example:
	//
	// AE4F6C34-065F-45AA-F5BN-4B8D816F6305
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelActiveOperationTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelActiveOperationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *CancelActiveOperationTasksResponseBody) SetIds(v string) *CancelActiveOperationTasksResponseBody {
	s.Ids = &v
	return s
}

func (s *CancelActiveOperationTasksResponseBody) SetRequestId(v string) *CancelActiveOperationTasksResponseBody {
	s.RequestId = &v
	return s
}

type CancelActiveOperationTasksResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelActiveOperationTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelActiveOperationTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelActiveOperationTasksResponse) GoString() string {
	return s.String()
}

func (s *CancelActiveOperationTasksResponse) SetHeaders(v map[string]*string) *CancelActiveOperationTasksResponse {
	s.Headers = v
	return s
}

func (s *CancelActiveOperationTasksResponse) SetStatusCode(v int32) *CancelActiveOperationTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelActiveOperationTasksResponse) SetBody(v *CancelActiveOperationTasksResponseBody) *CancelActiveOperationTasksResponse {
	s.Body = v
	return s
}

type CheckComponentsVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HBASE,HADOOP
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
}

func (s CheckComponentsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckComponentsVersionRequest) GoString() string {
	return s.String()
}

func (s *CheckComponentsVersionRequest) SetClusterId(v string) *CheckComponentsVersionRequest {
	s.ClusterId = &v
	return s
}

func (s *CheckComponentsVersionRequest) SetComponents(v string) *CheckComponentsVersionRequest {
	s.Components = &v
	return s
}

type CheckComponentsVersionResponseBody struct {
	Components *CheckComponentsVersionResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Struct"`
	// example:
	//
	// E3537EB4-1100-41CA-A147-C74CCC8BB12C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckComponentsVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckComponentsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckComponentsVersionResponseBody) SetComponents(v *CheckComponentsVersionResponseBodyComponents) *CheckComponentsVersionResponseBody {
	s.Components = v
	return s
}

func (s *CheckComponentsVersionResponseBody) SetRequestId(v string) *CheckComponentsVersionResponseBody {
	s.RequestId = &v
	return s
}

type CheckComponentsVersionResponseBodyComponents struct {
	Component []*CheckComponentsVersionResponseBodyComponentsComponent `json:"Component,omitempty" xml:"Component,omitempty" type:"Repeated"`
}

func (s CheckComponentsVersionResponseBodyComponents) String() string {
	return tea.Prettify(s)
}

func (s CheckComponentsVersionResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *CheckComponentsVersionResponseBodyComponents) SetComponent(v []*CheckComponentsVersionResponseBodyComponentsComponent) *CheckComponentsVersionResponseBodyComponents {
	s.Component = v
	return s
}

type CheckComponentsVersionResponseBodyComponentsComponent struct {
	// example:
	//
	// HBASE
	Component *string `json:"Component,omitempty" xml:"Component,omitempty"`
	// example:
	//
	// true
	IsLatestVersion *string `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
}

func (s CheckComponentsVersionResponseBodyComponentsComponent) String() string {
	return tea.Prettify(s)
}

func (s CheckComponentsVersionResponseBodyComponentsComponent) GoString() string {
	return s.String()
}

func (s *CheckComponentsVersionResponseBodyComponentsComponent) SetComponent(v string) *CheckComponentsVersionResponseBodyComponentsComponent {
	s.Component = &v
	return s
}

func (s *CheckComponentsVersionResponseBodyComponentsComponent) SetIsLatestVersion(v string) *CheckComponentsVersionResponseBodyComponentsComponent {
	s.IsLatestVersion = &v
	return s
}

type CheckComponentsVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckComponentsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckComponentsVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckComponentsVersionResponse) GoString() string {
	return s.String()
}

func (s *CheckComponentsVersionResponse) SetHeaders(v map[string]*string) *CheckComponentsVersionResponse {
	s.Headers = v
	return s
}

func (s *CheckComponentsVersionResponse) SetStatusCode(v int32) *CheckComponentsVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckComponentsVersionResponse) SetBody(v *CheckComponentsVersionResponseBody) *CheckComponentsVersionResponse {
	s.Body = v
	return s
}

type CloseBackupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CloseBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseBackupRequest) GoString() string {
	return s.String()
}

func (s *CloseBackupRequest) SetClusterId(v string) *CloseBackupRequest {
	s.ClusterId = &v
	return s
}

type CloseBackupResponseBody struct {
	// example:
	//
	// F1A11940-0C34-4385-864F-A01E29B55F6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CloseBackupResponseBody) SetRequestId(v string) *CloseBackupResponseBody {
	s.RequestId = &v
	return s
}

type CloseBackupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseBackupResponse) GoString() string {
	return s.String()
}

func (s *CloseBackupResponse) SetHeaders(v map[string]*string) *CloseBackupResponse {
	s.Headers = v
	return s
}

func (s *CloseBackupResponse) SetStatusCode(v int32) *CloseBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseBackupResponse) SetBody(v *CloseBackupResponseBody) *CloseBackupResponse {
	s.Body = v
	return s
}

type ConvertInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16o0pd52e3y****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 7
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
}

func (s ConvertInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequest) SetClusterId(v string) *ConvertInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *ConvertInstanceRequest) SetDuration(v int32) *ConvertInstanceRequest {
	s.Duration = &v
	return s
}

func (s *ConvertInstanceRequest) SetPayType(v string) *ConvertInstanceRequest {
	s.PayType = &v
	return s
}

func (s *ConvertInstanceRequest) SetPricingCycle(v string) *ConvertInstanceRequest {
	s.PricingCycle = &v
	return s
}

type ConvertInstanceResponseBody struct {
	// example:
	//
	// 54124548879
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 50373857-C47B-4B64-9332-D0B5280B59EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConvertInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertInstanceResponseBody) SetOrderId(v int64) *ConvertInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *ConvertInstanceResponseBody) SetRequestId(v string) *ConvertInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ConvertInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceResponse) GoString() string {
	return s.String()
}

func (s *ConvertInstanceResponse) SetHeaders(v map[string]*string) *ConvertInstanceResponse {
	s.Headers = v
	return s
}

func (s *ConvertInstanceResponse) SetStatusCode(v int32) *ConvertInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertInstanceResponse) SetBody(v *ConvertInstanceResponseBody) *ConvertInstanceResponse {
	s.Body = v
	return s
}

type CreateAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test01
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// **********
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetClusterId(v string) *CreateAccountRequest {
	s.ClusterId = &v
	return s
}

type CreateAccountResponseBody struct {
	// example:
	//
	// 50373857-C47B-4B64-9332-D0B5280B59EA
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

type CreateBackupPlanRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-wz94lbcqc****4x93
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

type CreateBackupPlanResponseBody struct {
	// example:
	//
	// 50373857-C47B-4B64-9332-D0B5280B59EA
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateBackupPlanResponse) SetStatusCode(v int32) *CreateBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBackupPlanResponse) SetBody(v *CreateBackupPlanResponseBody) *CreateBackupPlanResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	// example:
	//
	// 2
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// hbase_test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// 1024
	ColdStorageSize *int32 `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase.sn1.medium
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// example:
	//
	// 400
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// example:
	//
	// 0d2470df-da7b-4786-b981-9a164dae****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// hbase.sn1.medium
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 6
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-j4d53glb3****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 116.62.XX.XX/24
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// example:
	//
	// vsw-bp191otqj1ssyl****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp120k6ixs4eog****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetAutoRenewPeriod(v int32) *CreateClusterRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateClusterRequest) SetClientToken(v string) *CreateClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetColdStorageSize(v int32) *CreateClusterRequest {
	s.ColdStorageSize = &v
	return s
}

func (s *CreateClusterRequest) SetCoreInstanceType(v string) *CreateClusterRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *CreateClusterRequest) SetDiskSize(v int32) *CreateClusterRequest {
	s.DiskSize = &v
	return s
}

func (s *CreateClusterRequest) SetDiskType(v string) *CreateClusterRequest {
	s.DiskType = &v
	return s
}

func (s *CreateClusterRequest) SetEncryptionKey(v string) *CreateClusterRequest {
	s.EncryptionKey = &v
	return s
}

func (s *CreateClusterRequest) SetEngine(v string) *CreateClusterRequest {
	s.Engine = &v
	return s
}

func (s *CreateClusterRequest) SetEngineVersion(v string) *CreateClusterRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateClusterRequest) SetMasterInstanceType(v string) *CreateClusterRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *CreateClusterRequest) SetNodeCount(v int32) *CreateClusterRequest {
	s.NodeCount = &v
	return s
}

func (s *CreateClusterRequest) SetPayType(v string) *CreateClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateClusterRequest) SetPeriod(v int32) *CreateClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateClusterRequest) SetPeriodUnit(v string) *CreateClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateClusterRequest) SetRegionId(v string) *CreateClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateClusterRequest) SetResourceGroupId(v string) *CreateClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityIPList(v string) *CreateClusterRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateClusterRequest) SetVSwitchId(v string) *CreateClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) SetZoneId(v string) *CreateClusterRequest {
	s.ZoneId = &v
	return s
}

type CreateClusterResponseBody struct {
	// example:
	//
	// hb-bp1hy2sjf8gd****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 23232069786****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 3E19E345-101D-4014-946C-A205
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetOrderId(v string) *CreateClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

type CreateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateClusterResponse) SetStatusCode(v int32) *CreateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateGlobalResourceRequest struct {
	// example:
	//
	// xxxxx-xxxxx-xxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PubPhoenixSLBQueryServerVip
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GLOBAL_VIP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateGlobalResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGlobalResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateGlobalResourceRequest) SetClientToken(v string) *CreateGlobalResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetClusterId(v string) *CreateGlobalResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetRegionId(v string) *CreateGlobalResourceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetResourceName(v string) *CreateGlobalResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetResourceType(v string) *CreateGlobalResourceRequest {
	s.ResourceType = &v
	return s
}

type CreateGlobalResourceResponseBody struct {
	// example:
	//
	// 1AB9ABDF-7E1E-44AD-8610-70A005115DD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGlobalResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGlobalResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGlobalResourceResponseBody) SetRequestId(v string) *CreateGlobalResourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateGlobalResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGlobalResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGlobalResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGlobalResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateGlobalResourceResponse) SetHeaders(v map[string]*string) *CreateGlobalResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateGlobalResourceResponse) SetStatusCode(v int32) *CreateGlobalResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGlobalResourceResponse) SetBody(v *CreateGlobalResourceResponseBody) *CreateGlobalResourceResponse {
	s.Body = v
	return s
}

type CreateHBaseSlbServerRequest struct {
	// example:
	//
	// xxxxx-xxxxx-xxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// thrift
	SlbServer *string `json:"SlbServer,omitempty" xml:"SlbServer,omitempty"`
}

func (s CreateHBaseSlbServerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHBaseSlbServerRequest) GoString() string {
	return s.String()
}

func (s *CreateHBaseSlbServerRequest) SetClientToken(v string) *CreateHBaseSlbServerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateHBaseSlbServerRequest) SetClusterId(v string) *CreateHBaseSlbServerRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateHBaseSlbServerRequest) SetSlbServer(v string) *CreateHBaseSlbServerRequest {
	s.SlbServer = &v
	return s
}

type CreateHBaseSlbServerResponseBody struct {
	// example:
	//
	// 61FC5B21-87B0-41BC-9686-9DA395EB40B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHBaseSlbServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHBaseSlbServerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHBaseSlbServerResponseBody) SetRequestId(v string) *CreateHBaseSlbServerResponseBody {
	s.RequestId = &v
	return s
}

type CreateHBaseSlbServerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHBaseSlbServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHBaseSlbServerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHBaseSlbServerResponse) GoString() string {
	return s.String()
}

func (s *CreateHBaseSlbServerResponse) SetHeaders(v map[string]*string) *CreateHBaseSlbServerResponse {
	s.Headers = v
	return s
}

func (s *CreateHBaseSlbServerResponse) SetStatusCode(v int32) *CreateHBaseSlbServerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHBaseSlbServerResponse) SetBody(v *CreateHBaseSlbServerResponseBody) *CreateHBaseSlbServerResponse {
	s.Body = v
	return s
}

type CreateHbaseHaSlbRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bds-t4n3496whj23****
	BdsId *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ha-v21tmnxjwh2yu****
	HaId *string `json:"HaId,omitempty" xml:"HaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// thrift
	HaTypes *string `json:"HaTypes,omitempty" xml:"HaTypes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Active
	HbaseType *string `json:"HbaseType,omitempty" xml:"HbaseType,omitempty"`
}

func (s CreateHbaseHaSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHbaseHaSlbRequest) GoString() string {
	return s.String()
}

func (s *CreateHbaseHaSlbRequest) SetBdsId(v string) *CreateHbaseHaSlbRequest {
	s.BdsId = &v
	return s
}

func (s *CreateHbaseHaSlbRequest) SetClientToken(v string) *CreateHbaseHaSlbRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateHbaseHaSlbRequest) SetHaId(v string) *CreateHbaseHaSlbRequest {
	s.HaId = &v
	return s
}

func (s *CreateHbaseHaSlbRequest) SetHaTypes(v string) *CreateHbaseHaSlbRequest {
	s.HaTypes = &v
	return s
}

func (s *CreateHbaseHaSlbRequest) SetHbaseType(v string) *CreateHbaseHaSlbRequest {
	s.HbaseType = &v
	return s
}

type CreateHbaseHaSlbResponseBody struct {
	// example:
	//
	// C9D568D9-A59C-4AF2-8FBB-F086A841D58E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHbaseHaSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHbaseHaSlbResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHbaseHaSlbResponseBody) SetRequestId(v string) *CreateHbaseHaSlbResponseBody {
	s.RequestId = &v
	return s
}

type CreateHbaseHaSlbResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHbaseHaSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHbaseHaSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHbaseHaSlbResponse) GoString() string {
	return s.String()
}

func (s *CreateHbaseHaSlbResponse) SetHeaders(v map[string]*string) *CreateHbaseHaSlbResponse {
	s.Headers = v
	return s
}

func (s *CreateHbaseHaSlbResponse) SetStatusCode(v int32) *CreateHbaseHaSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHbaseHaSlbResponse) SetBody(v *CreateHbaseHaSlbResponseBody) *CreateHbaseHaSlbResponse {
	s.Body = v
	return s
}

type CreateMultiZoneClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// vsw-hangxzhouxb*****
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.0
	ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	// example:
	//
	// 0
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// example:
	//
	// dfh3sf5gslfksfk****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// hbaseue_test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 400
	CoreDiskSize *int32 `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cloud_ssd
	CoreDiskType *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase.sn1.medium
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	CoreNodeCount *int32 `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbaseue
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 400
	LogDiskSize *int32 `json:"LogDiskSize,omitempty" xml:"LogDiskSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cloud_ssd
	LogDiskType *string `json:"LogDiskType,omitempty" xml:"LogDiskType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase.sn1.medium
	LogInstanceType *string `json:"LogInstanceType,omitempty" xml:"LogInstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	LogNodeCount *int32 `json:"LogNodeCount,omitempty" xml:"LogNodeCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase.sn1.medium
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-bef-aliyun-com
	MultiZoneCombination *string `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-hangxzhouxe****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-e
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-gg3f4f5d5g5w****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-hangxzhouxf****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-f
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp120k6ixs4eog****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateMultiZoneClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiZoneClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateMultiZoneClusterRequest) SetArbiterVSwitchId(v string) *CreateMultiZoneClusterRequest {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetArbiterZoneId(v string) *CreateMultiZoneClusterRequest {
	s.ArbiterZoneId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetArchVersion(v string) *CreateMultiZoneClusterRequest {
	s.ArchVersion = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetAutoRenewPeriod(v int32) *CreateMultiZoneClusterRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetClientToken(v string) *CreateMultiZoneClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetClusterName(v string) *CreateMultiZoneClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetCoreDiskSize(v int32) *CreateMultiZoneClusterRequest {
	s.CoreDiskSize = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetCoreDiskType(v string) *CreateMultiZoneClusterRequest {
	s.CoreDiskType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetCoreInstanceType(v string) *CreateMultiZoneClusterRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetCoreNodeCount(v int32) *CreateMultiZoneClusterRequest {
	s.CoreNodeCount = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetEngine(v string) *CreateMultiZoneClusterRequest {
	s.Engine = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetEngineVersion(v string) *CreateMultiZoneClusterRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetLogDiskSize(v int32) *CreateMultiZoneClusterRequest {
	s.LogDiskSize = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetLogDiskType(v string) *CreateMultiZoneClusterRequest {
	s.LogDiskType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetLogInstanceType(v string) *CreateMultiZoneClusterRequest {
	s.LogInstanceType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetLogNodeCount(v int32) *CreateMultiZoneClusterRequest {
	s.LogNodeCount = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetMasterInstanceType(v string) *CreateMultiZoneClusterRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetMultiZoneCombination(v string) *CreateMultiZoneClusterRequest {
	s.MultiZoneCombination = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetPayType(v string) *CreateMultiZoneClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetPeriod(v int32) *CreateMultiZoneClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetPeriodUnit(v string) *CreateMultiZoneClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetPrimaryVSwitchId(v string) *CreateMultiZoneClusterRequest {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetPrimaryZoneId(v string) *CreateMultiZoneClusterRequest {
	s.PrimaryZoneId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetRegionId(v string) *CreateMultiZoneClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetResourceGroupId(v string) *CreateMultiZoneClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetSecurityIPList(v string) *CreateMultiZoneClusterRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetStandbyVSwitchId(v string) *CreateMultiZoneClusterRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetStandbyZoneId(v string) *CreateMultiZoneClusterRequest {
	s.StandbyZoneId = &v
	return s
}

func (s *CreateMultiZoneClusterRequest) SetVpcId(v string) *CreateMultiZoneClusterRequest {
	s.VpcId = &v
	return s
}

type CreateMultiZoneClusterResponseBody struct {
	// example:
	//
	// ld-t4nn71xa0yn56****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 23232453****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 7F68E8F5-0377-4CF8-8B1D-FFFD6F5804D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMultiZoneClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiZoneClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultiZoneClusterResponseBody) SetClusterId(v string) *CreateMultiZoneClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateMultiZoneClusterResponseBody) SetOrderId(v string) *CreateMultiZoneClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateMultiZoneClusterResponseBody) SetRequestId(v string) *CreateMultiZoneClusterResponseBody {
	s.RequestId = &v
	return s
}

type CreateMultiZoneClusterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMultiZoneClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMultiZoneClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMultiZoneClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateMultiZoneClusterResponse) SetHeaders(v map[string]*string) *CreateMultiZoneClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateMultiZoneClusterResponse) SetStatusCode(v int32) *CreateMultiZoneClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMultiZoneClusterResponse) SetBody(v *CreateMultiZoneClusterResponseBody) *CreateMultiZoneClusterResponse {
	s.Body = v
	return s
}

type CreateRestorePlanRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	RestoreAllTable *bool `json:"RestoreAllTable,omitempty" xml:"RestoreAllTable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	RestoreByCopy *bool `json:"RestoreByCopy,omitempty" xml:"RestoreByCopy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020-11-05T05:49:42Z
	RestoreToDate *string `json:"RestoreToDate,omitempty" xml:"RestoreToDate,omitempty"`
	// example:
	//
	// test_ns:test_table/test_ns:test_table2
	Tables *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-bp169l540vc6c****
	TargetClusterId *string `json:"TargetClusterId,omitempty" xml:"TargetClusterId,omitempty"`
}

func (s CreateRestorePlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRestorePlanRequest) GoString() string {
	return s.String()
}

func (s *CreateRestorePlanRequest) SetClusterId(v string) *CreateRestorePlanRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateRestorePlanRequest) SetRestoreAllTable(v bool) *CreateRestorePlanRequest {
	s.RestoreAllTable = &v
	return s
}

func (s *CreateRestorePlanRequest) SetRestoreByCopy(v bool) *CreateRestorePlanRequest {
	s.RestoreByCopy = &v
	return s
}

func (s *CreateRestorePlanRequest) SetRestoreToDate(v string) *CreateRestorePlanRequest {
	s.RestoreToDate = &v
	return s
}

func (s *CreateRestorePlanRequest) SetTables(v string) *CreateRestorePlanRequest {
	s.Tables = &v
	return s
}

func (s *CreateRestorePlanRequest) SetTargetClusterId(v string) *CreateRestorePlanRequest {
	s.TargetClusterId = &v
	return s
}

type CreateRestorePlanResponseBody struct {
	// example:
	//
	// A0598673-EB6E-4F6D-9961-E0F2012090C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRestorePlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRestorePlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRestorePlanResponseBody) SetRequestId(v string) *CreateRestorePlanResponseBody {
	s.RequestId = &v
	return s
}

type CreateRestorePlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRestorePlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRestorePlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRestorePlanResponse) GoString() string {
	return s.String()
}

func (s *CreateRestorePlanResponse) SetHeaders(v map[string]*string) *CreateRestorePlanResponse {
	s.Headers = v
	return s
}

func (s *CreateRestorePlanResponse) SetStatusCode(v int32) *CreateRestorePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRestorePlanResponse) SetBody(v *CreateRestorePlanResponseBody) *CreateRestorePlanResponse {
	s.Body = v
	return s
}

type CreateServerlessClusterRequest struct {
	// example:
	//
	// 2
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// xx
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// serverless-name
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// example:
	//
	// serverlesshbase
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 6
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-j4d53glb3****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1000
	ServerlessCapability *int32 `json:"ServerlessCapability,omitempty" xml:"ServerlessCapability,omitempty"`
	// example:
	//
	// serverless.small
	ServerlessSpec *string `json:"ServerlessSpec,omitempty" xml:"ServerlessSpec,omitempty"`
	// example:
	//
	// 100
	ServerlessStorage *int32 `json:"ServerlessStorage,omitempty" xml:"ServerlessStorage,omitempty"`
	// example:
	//
	// vsw-bp191ipotqj1ssyl*****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp120k6ixs4eog****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateServerlessClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServerlessClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateServerlessClusterRequest) SetAutoRenewPeriod(v int32) *CreateServerlessClusterRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetClientToken(v string) *CreateServerlessClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetClientType(v string) *CreateServerlessClusterRequest {
	s.ClientType = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetClusterName(v string) *CreateServerlessClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetDiskType(v string) *CreateServerlessClusterRequest {
	s.DiskType = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetEngine(v string) *CreateServerlessClusterRequest {
	s.Engine = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetEngineVersion(v string) *CreateServerlessClusterRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetPayType(v string) *CreateServerlessClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetPeriod(v int32) *CreateServerlessClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetPeriodUnit(v string) *CreateServerlessClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetRegionId(v string) *CreateServerlessClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetResourceGroupId(v string) *CreateServerlessClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetServerlessCapability(v int32) *CreateServerlessClusterRequest {
	s.ServerlessCapability = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetServerlessSpec(v string) *CreateServerlessClusterRequest {
	s.ServerlessSpec = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetServerlessStorage(v int32) *CreateServerlessClusterRequest {
	s.ServerlessStorage = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetVSwitchId(v string) *CreateServerlessClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetVpcId(v string) *CreateServerlessClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateServerlessClusterRequest) SetZoneId(v string) *CreateServerlessClusterRequest {
	s.ZoneId = &v
	return s
}

type CreateServerlessClusterResponseBody struct {
	// example:
	//
	// sh-bp1a969y7681****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 23232453233*****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// *********
	PassWord *string `json:"PassWord,omitempty" xml:"PassWord,omitempty"`
	// example:
	//
	// 3E19E345-101D-4014-946C-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServerlessClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServerlessClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServerlessClusterResponseBody) SetClusterId(v string) *CreateServerlessClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateServerlessClusterResponseBody) SetOrderId(v string) *CreateServerlessClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateServerlessClusterResponseBody) SetPassWord(v string) *CreateServerlessClusterResponseBody {
	s.PassWord = &v
	return s
}

func (s *CreateServerlessClusterResponseBody) SetRequestId(v string) *CreateServerlessClusterResponseBody {
	s.RequestId = &v
	return s
}

type CreateServerlessClusterResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServerlessClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServerlessClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServerlessClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateServerlessClusterResponse) SetHeaders(v map[string]*string) *CreateServerlessClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateServerlessClusterResponse) SetStatusCode(v int32) *CreateServerlessClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServerlessClusterResponse) SetBody(v *CreateServerlessClusterResponseBody) *CreateServerlessClusterResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test01
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

func (s *DeleteAccountRequest) SetClusterId(v string) *DeleteAccountRequest {
	s.ClusterId = &v
	return s
}

type DeleteAccountResponseBody struct {
	// example:
	//
	// 729CB2A7-3065-53A9-B27C-7033CA4881D9
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

type DeleteGlobalResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PubPhoenixSLBQueryServerVip
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GLOBAL_VIP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DeleteGlobalResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGlobalResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGlobalResourceRequest) SetClusterId(v string) *DeleteGlobalResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetRegionId(v string) *DeleteGlobalResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetResourceName(v string) *DeleteGlobalResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetResourceType(v string) *DeleteGlobalResourceRequest {
	s.ResourceType = &v
	return s
}

type DeleteGlobalResourceResponseBody struct {
	// example:
	//
	// BD0B0B9A-79E8-4FDD-9C51-93443490B784
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGlobalResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGlobalResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGlobalResourceResponseBody) SetRequestId(v string) *DeleteGlobalResourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGlobalResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGlobalResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGlobalResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGlobalResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGlobalResourceResponse) SetHeaders(v map[string]*string) *DeleteGlobalResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGlobalResourceResponse) SetStatusCode(v int32) *DeleteGlobalResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGlobalResourceResponse) SetBody(v *DeleteGlobalResourceResponseBody) *DeleteGlobalResourceResponse {
	s.Body = v
	return s
}

type DeleteHBaseHaDBRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bds-bp14112fd7g52s1****
	BdsId *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ha-sw2o0l01s4r76****
	HaId *string `json:"HaId,omitempty" xml:"HaId,omitempty"`
}

func (s DeleteHBaseHaDBRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHBaseHaDBRequest) GoString() string {
	return s.String()
}

func (s *DeleteHBaseHaDBRequest) SetBdsId(v string) *DeleteHBaseHaDBRequest {
	s.BdsId = &v
	return s
}

func (s *DeleteHBaseHaDBRequest) SetHaId(v string) *DeleteHBaseHaDBRequest {
	s.HaId = &v
	return s
}

type DeleteHBaseHaDBResponseBody struct {
	// example:
	//
	// B409CF51-E01F-4551-BE40-123678FA9026
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHBaseHaDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHBaseHaDBResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHBaseHaDBResponseBody) SetRequestId(v string) *DeleteHBaseHaDBResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHBaseHaDBResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHBaseHaDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHBaseHaDBResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHBaseHaDBResponse) GoString() string {
	return s.String()
}

func (s *DeleteHBaseHaDBResponse) SetHeaders(v map[string]*string) *DeleteHBaseHaDBResponse {
	s.Headers = v
	return s
}

func (s *DeleteHBaseHaDBResponse) SetStatusCode(v int32) *DeleteHBaseHaDBResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHBaseHaDBResponse) SetBody(v *DeleteHBaseHaDBResponseBody) *DeleteHBaseHaDBResponse {
	s.Body = v
	return s
}

type DeleteHBaseSlbServerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// thrift
	SlbServer *string `json:"SlbServer,omitempty" xml:"SlbServer,omitempty"`
}

func (s DeleteHBaseSlbServerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHBaseSlbServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteHBaseSlbServerRequest) SetClusterId(v string) *DeleteHBaseSlbServerRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteHBaseSlbServerRequest) SetSlbServer(v string) *DeleteHBaseSlbServerRequest {
	s.SlbServer = &v
	return s
}

type DeleteHBaseSlbServerResponseBody struct {
	// example:
	//
	// 7242130A-82CF-49BF-AB32-30DCB819EBA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHBaseSlbServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHBaseSlbServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHBaseSlbServerResponseBody) SetRequestId(v string) *DeleteHBaseSlbServerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHBaseSlbServerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHBaseSlbServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHBaseSlbServerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHBaseSlbServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteHBaseSlbServerResponse) SetHeaders(v map[string]*string) *DeleteHBaseSlbServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteHBaseSlbServerResponse) SetStatusCode(v int32) *DeleteHBaseSlbServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHBaseSlbServerResponse) SetBody(v *DeleteHBaseSlbServerResponseBody) *DeleteHBaseSlbServerResponse {
	s.Body = v
	return s
}

type DeleteHbaseHaSlbRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bds-t4n3496whj23ia4k
	BdsId *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ha-v21tmnxjwh2yuy1il
	HaId *string `json:"HaId,omitempty" xml:"HaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// thrift
	HaTypes *string `json:"HaTypes,omitempty" xml:"HaTypes,omitempty"`
}

func (s DeleteHbaseHaSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHbaseHaSlbRequest) GoString() string {
	return s.String()
}

func (s *DeleteHbaseHaSlbRequest) SetBdsId(v string) *DeleteHbaseHaSlbRequest {
	s.BdsId = &v
	return s
}

func (s *DeleteHbaseHaSlbRequest) SetHaId(v string) *DeleteHbaseHaSlbRequest {
	s.HaId = &v
	return s
}

func (s *DeleteHbaseHaSlbRequest) SetHaTypes(v string) *DeleteHbaseHaSlbRequest {
	s.HaTypes = &v
	return s
}

type DeleteHbaseHaSlbResponseBody struct {
	// example:
	//
	// C9D568D9-A59C-4AF2-8FBB-F086A841D58E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHbaseHaSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHbaseHaSlbResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHbaseHaSlbResponseBody) SetRequestId(v string) *DeleteHbaseHaSlbResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHbaseHaSlbResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHbaseHaSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHbaseHaSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHbaseHaSlbResponse) GoString() string {
	return s.String()
}

func (s *DeleteHbaseHaSlbResponse) SetHeaders(v map[string]*string) *DeleteHbaseHaSlbResponse {
	s.Headers = v
	return s
}

func (s *DeleteHbaseHaSlbResponse) SetStatusCode(v int32) *DeleteHbaseHaSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHbaseHaSlbResponse) SetBody(v *DeleteHbaseHaSlbResponseBody) *DeleteHbaseHaSlbResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16o0pd52e3y****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// false
	ImmediateDeleteFlag *bool `json:"ImmediateDeleteFlag,omitempty" xml:"ImmediateDeleteFlag,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetClusterId(v string) *DeleteInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteInstanceRequest) SetImmediateDeleteFlag(v bool) *DeleteInstanceRequest {
	s.ImmediateDeleteFlag = &v
	return s
}

type DeleteInstanceResponseBody struct {
	// example:
	//
	// 50373857-C47B-4B64-9332-D0B5280B59EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type DeleteInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteInstanceResponse) SetStatusCode(v int32) *DeleteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DeleteMultiZoneClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// d-t4nn71xa0yn56****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// false
	ImmediateDeleteFlag *bool `json:"ImmediateDeleteFlag,omitempty" xml:"ImmediateDeleteFlag,omitempty"`
}

func (s DeleteMultiZoneClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMultiZoneClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultiZoneClusterRequest) SetClusterId(v string) *DeleteMultiZoneClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteMultiZoneClusterRequest) SetImmediateDeleteFlag(v bool) *DeleteMultiZoneClusterRequest {
	s.ImmediateDeleteFlag = &v
	return s
}

type DeleteMultiZoneClusterResponseBody struct {
	// example:
	//
	// 169A3910-A39E-4BC2-AA9F-E7AD8D473527
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMultiZoneClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMultiZoneClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMultiZoneClusterResponseBody) SetRequestId(v string) *DeleteMultiZoneClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMultiZoneClusterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMultiZoneClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMultiZoneClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMultiZoneClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteMultiZoneClusterResponse) SetHeaders(v map[string]*string) *DeleteMultiZoneClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteMultiZoneClusterResponse) SetStatusCode(v int32) *DeleteMultiZoneClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMultiZoneClusterResponse) SetBody(v *DeleteMultiZoneClusterResponseBody) *DeleteMultiZoneClusterResponse {
	s.Body = v
	return s
}

type DeleteServerlessClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sh-bp1pj13wh9****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteServerlessClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerlessClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteServerlessClusterRequest) SetClusterId(v string) *DeleteServerlessClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteServerlessClusterRequest) SetRegionId(v string) *DeleteServerlessClusterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServerlessClusterRequest) SetZoneId(v string) *DeleteServerlessClusterRequest {
	s.ZoneId = &v
	return s
}

type DeleteServerlessClusterResponseBody struct {
	// example:
	//
	// 46950E74-59C4-4E3E-9B38-A33B*********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServerlessClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerlessClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServerlessClusterResponseBody) SetRequestId(v string) *DeleteServerlessClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServerlessClusterResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServerlessClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServerlessClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerlessClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteServerlessClusterResponse) SetHeaders(v map[string]*string) *DeleteServerlessClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteServerlessClusterResponse) SetStatusCode(v int32) *DeleteServerlessClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServerlessClusterResponse) SetBody(v *DeleteServerlessClusterResponseBody) *DeleteServerlessClusterResponse {
	s.Body = v
	return s
}

type DeleteUserHdfsInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16o0pd52e3y****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hdfs_test
	NameService *string `json:"NameService,omitempty" xml:"NameService,omitempty"`
}

func (s DeleteUserHdfsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserHdfsInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserHdfsInfoRequest) SetClusterId(v string) *DeleteUserHdfsInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteUserHdfsInfoRequest) SetNameService(v string) *DeleteUserHdfsInfoRequest {
	s.NameService = &v
	return s
}

type DeleteUserHdfsInfoResponseBody struct {
	// example:
	//
	// 50373857-C47B-4B64-9332-D0B5280B59EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserHdfsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserHdfsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserHdfsInfoResponseBody) SetRequestId(v string) *DeleteUserHdfsInfoResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserHdfsInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserHdfsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserHdfsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserHdfsInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserHdfsInfoResponse) SetHeaders(v map[string]*string) *DeleteUserHdfsInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserHdfsInfoResponse) SetStatusCode(v int32) *DeleteUserHdfsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserHdfsInfoResponse) SetBody(v *DeleteUserHdfsInfoResponseBody) *DeleteUserHdfsInfoResponse {
	s.Body = v
	return s
}

type DescribeAccountsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1uoihlf82e8****
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
	Accounts *DescribeAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// example:
	//
	// F744E939-D08D-5623-82C8-9D1F9F7685D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAccountsResponseBodyAccounts struct {
	Account []*string `json:"account,omitempty" xml:"account,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBodyAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccount(v []*string) *DescribeAccountsResponseBodyAccounts {
	s.Account = v
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

type DescribeActiveOperationTaskTypeRequest struct {
	// example:
	//
	// 0
	IsHistory            *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeActiveOperationTaskTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskTypeRequest) SetIsHistory(v int32) *DescribeActiveOperationTaskTypeRequest {
	s.IsHistory = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetOwnerAccount(v string) *DescribeActiveOperationTaskTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetOwnerId(v int64) *DescribeActiveOperationTaskTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetResourceOwnerAccount(v string) *DescribeActiveOperationTaskTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetResourceOwnerId(v int64) *DescribeActiveOperationTaskTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeRequest) SetSecurityToken(v string) *DescribeActiveOperationTaskTypeRequest {
	s.SecurityToken = &v
	return s
}

type DescribeActiveOperationTaskTypeResponseBody struct {
	// example:
	//
	// EC7E27FC-58F8-4722-89CF-D1B6B0971956
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TypeList  []*DescribeActiveOperationTaskTypeResponseBodyTypeList `json:"TypeList,omitempty" xml:"TypeList,omitempty" type:"Repeated"`
}

func (s DescribeActiveOperationTaskTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskTypeResponseBody) SetRequestId(v string) *DescribeActiveOperationTaskTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponseBody) SetTypeList(v []*DescribeActiveOperationTaskTypeResponseBodyTypeList) *DescribeActiveOperationTaskTypeResponseBody {
	s.TypeList = v
	return s
}

type DescribeActiveOperationTaskTypeResponseBodyTypeList struct {
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// rds_apsaradb_upgrade
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// rds_apsaradb_upgrade
	TaskTypeInfoEn *string `json:"TaskTypeInfoEn,omitempty" xml:"TaskTypeInfoEn,omitempty"`
	TaskTypeInfoZh *string `json:"TaskTypeInfoZh,omitempty" xml:"TaskTypeInfoZh,omitempty"`
}

func (s DescribeActiveOperationTaskTypeResponseBodyTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskTypeResponseBodyTypeList) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskTypeResponseBodyTypeList) SetCount(v int32) *DescribeActiveOperationTaskTypeResponseBodyTypeList {
	s.Count = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponseBodyTypeList) SetTaskType(v string) *DescribeActiveOperationTaskTypeResponseBodyTypeList {
	s.TaskType = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponseBodyTypeList) SetTaskTypeInfoEn(v string) *DescribeActiveOperationTaskTypeResponseBodyTypeList {
	s.TaskTypeInfoEn = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponseBodyTypeList) SetTaskTypeInfoZh(v string) *DescribeActiveOperationTaskTypeResponseBodyTypeList {
	s.TaskTypeInfoZh = &v
	return s
}

type DescribeActiveOperationTaskTypeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActiveOperationTaskTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActiveOperationTaskTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTaskTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskTypeResponse) SetHeaders(v map[string]*string) *DescribeActiveOperationTaskTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponse) SetStatusCode(v int32) *DescribeActiveOperationTaskTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActiveOperationTaskTypeResponse) SetBody(v *DescribeActiveOperationTaskTypeResponseBody) *DescribeActiveOperationTaskTypeResponse {
	s.Body = v
	return s
}

type DescribeActiveOperationTasksRequest struct {
	// example:
	//
	// 1
	AllowCancel *int32 `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// example:
	//
	// 1
	AllowChange *int32 `json:"AllowChange,omitempty" xml:"AllowChange,omitempty"`
	// example:
	//
	// S1
	ChangeLevel *string `json:"ChangeLevel,omitempty" xml:"ChangeLevel,omitempty"`
	// example:
	//
	// hbaseue
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// ld-bp150tns0sjxs****
	InsName      *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// hbase
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// 5
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// rds_apsaradb_upgrade
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeActiveOperationTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTasksRequest) SetAllowCancel(v int32) *DescribeActiveOperationTasksRequest {
	s.AllowCancel = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetAllowChange(v int32) *DescribeActiveOperationTasksRequest {
	s.AllowChange = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetChangeLevel(v string) *DescribeActiveOperationTasksRequest {
	s.ChangeLevel = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetDbType(v string) *DescribeActiveOperationTasksRequest {
	s.DbType = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetInsName(v string) *DescribeActiveOperationTasksRequest {
	s.InsName = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetOwnerAccount(v string) *DescribeActiveOperationTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetOwnerId(v int64) *DescribeActiveOperationTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetPageNumber(v int32) *DescribeActiveOperationTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetPageSize(v int32) *DescribeActiveOperationTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetProductId(v string) *DescribeActiveOperationTasksRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetRegion(v string) *DescribeActiveOperationTasksRequest {
	s.Region = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetResourceOwnerAccount(v string) *DescribeActiveOperationTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetResourceOwnerId(v int64) *DescribeActiveOperationTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetSecurityToken(v string) *DescribeActiveOperationTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetStatus(v int32) *DescribeActiveOperationTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetTaskType(v string) *DescribeActiveOperationTasksRequest {
	s.TaskType = &v
	return s
}

type DescribeActiveOperationTasksResponseBody struct {
	Items []*DescribeActiveOperationTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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
	// C634E813-42FA-53D2-A7EB-B881C4B264CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeActiveOperationTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTasksResponseBody) SetItems(v []*DescribeActiveOperationTasksResponseBodyItems) *DescribeActiveOperationTasksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetPageNumber(v int32) *DescribeActiveOperationTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetPageSize(v int32) *DescribeActiveOperationTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetRequestId(v string) *DescribeActiveOperationTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetTotalRecordCount(v int32) *DescribeActiveOperationTasksResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeActiveOperationTasksResponseBodyItems struct {
	// example:
	//
	// 1
	AllowCancel *string `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// example:
	//
	// 1
	AllowChange *string `json:"AllowChange,omitempty" xml:"AllowChange,omitempty"`
	// example:
	//
	// all
	ChangeLevel *string `json:"ChangeLevel,omitempty" xml:"ChangeLevel,omitempty"`
	// example:
	//
	// Risk repairment
	ChangeLevelEn *string `json:"ChangeLevelEn,omitempty" xml:"ChangeLevelEn,omitempty"`
	ChangeLevelZh *string `json:"ChangeLevelZh,omitempty" xml:"ChangeLevelZh,omitempty"`
	// example:
	//
	// 2022-02-15 23:59:59
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// cn-shanghai-et-a
	CurrentAVZ *string `json:"CurrentAVZ,omitempty" xml:"CurrentAVZ,omitempty"`
	// example:
	//
	// hbaseue
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// 2.0
	DbVersion *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// example:
	//
	// 2022-02-19 23:59:59
	Deadline *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// example:
	//
	// 111111
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Transient instance disconnection
	ImpactEn *string `json:"ImpactEn,omitempty" xml:"ImpactEn,omitempty"`
	ImpactZh *string `json:"ImpactZh,omitempty" xml:"ImpactZh,omitempty"`
	// example:
	//
	// --
	InsComment *string `json:"InsComment,omitempty" xml:"InsComment,omitempty"`
	// example:
	//
	// ld-bp150tns0sjxs****
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// example:
	//
	// 2022-02-19 14:00:00
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 14:00:00
	PrepareInterval *string `json:"PrepareInterval,omitempty" xml:"PrepareInterval,omitempty"`
	// example:
	//
	// cn-hanghzou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// --
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// example:
	//
	// 2022-02-19 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 5
	Status      *int32    `json:"Status,omitempty" xml:"Status,omitempty"`
	SubInsNames []*string `json:"SubInsNames,omitempty" xml:"SubInsNames,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-02-19 14:00:00
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// example:
	//
	// rds_apsaradb_upgrade
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// rds_apsaradb_upgrade
	TaskTypeEn *string `json:"TaskTypeEn,omitempty" xml:"TaskTypeEn,omitempty"`
	TaskTypeZh *string `json:"TaskTypeZh,omitempty" xml:"TaskTypeZh,omitempty"`
}

func (s DescribeActiveOperationTasksResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetAllowCancel(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.AllowCancel = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetAllowChange(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.AllowChange = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetChangeLevel(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ChangeLevel = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetChangeLevelEn(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ChangeLevelEn = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetChangeLevelZh(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ChangeLevelZh = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetCreatedTime(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.CreatedTime = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetCurrentAVZ(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.CurrentAVZ = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetDbType(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.DbType = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetDbVersion(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.DbVersion = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetDeadline(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.Deadline = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetId(v int32) *DescribeActiveOperationTasksResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetImpactEn(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ImpactEn = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetImpactZh(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ImpactZh = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetInsComment(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.InsComment = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetInsName(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.InsName = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetModifiedTime(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetPrepareInterval(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.PrepareInterval = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetRegion(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.Region = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetResultInfo(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ResultInfo = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetStartTime(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetStatus(v int32) *DescribeActiveOperationTasksResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetSubInsNames(v []*string) *DescribeActiveOperationTasksResponseBodyItems {
	s.SubInsNames = v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetSwitchTime(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.SwitchTime = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetTaskType(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.TaskType = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetTaskTypeEn(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.TaskTypeEn = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetTaskTypeZh(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.TaskTypeZh = &v
	return s
}

type DescribeActiveOperationTasksResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActiveOperationTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActiveOperationTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveOperationTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTasksResponse) SetHeaders(v map[string]*string) *DescribeActiveOperationTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveOperationTasksResponse) SetStatusCode(v int32) *DescribeActiveOperationTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActiveOperationTasksResponse) SetBody(v *DescribeActiveOperationTasksResponseBody) *DescribeActiveOperationTasksResponse {
	s.Body = v
	return s
}

type DescribeAvailableResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// hbase.sn1.large
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// example:
	//
	// hbaseue
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) SetChargeType(v string) *DescribeAvailableResourceRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetCoreInstanceType(v string) *DescribeAvailableResourceRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetDiskType(v string) *DescribeAvailableResourceRequest {
	s.DiskType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetEngine(v string) *DescribeAvailableResourceRequest {
	s.Engine = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetEngineVersion(v string) *DescribeAvailableResourceRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetRegionId(v string) *DescribeAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetZoneId(v string) *DescribeAvailableResourceRequest {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourceResponseBody struct {
	AvailableZones *DescribeAvailableResourceResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Struct"`
	// example:
	//
	// EA76F208-E334-592A-A0C6-41E15EC87ED0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) SetAvailableZones(v *DescribeAvailableResourceResponseBodyAvailableZones) *DescribeAvailableResourceResponseBody {
	s.AvailableZones = v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZones struct {
	AvailableZone []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZones) SetAvailableZone(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) *DescribeAvailableResourceResponseBodyAvailableZones {
	s.AvailableZone = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone struct {
	MasterResources *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources `json:"MasterResources,omitempty" xml:"MasterResources,omitempty" type:"Struct"`
	// example:
	//
	// cn-shenzhen
	RegionId         *string                                                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SupportedEngines *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines `json:"SupportedEngines,omitempty" xml:"SupportedEngines,omitempty" type:"Struct"`
	// example:
	//
	// cn-shenzhen-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetMasterResources(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.MasterResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetRegionId(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetSupportedEngines(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.SupportedEngines = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetZoneId(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources struct {
	MasterResource []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource `json:"MasterResource,omitempty" xml:"MasterResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) SetMasterResource(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources {
	s.MasterResource = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource struct {
	// example:
	//
	// hbase.sn1.medium
	InstanceType       *string                                                                                                          `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypeDetail *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail `json:"InstanceTypeDetail,omitempty" xml:"InstanceTypeDetail,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) SetInstanceType(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource {
	s.InstanceType = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) SetInstanceTypeDetail(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource {
	s.InstanceTypeDetail = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail struct {
	// example:
	//
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 8
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) SetCpu(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail {
	s.Cpu = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) SetMem(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail {
	s.Mem = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines struct {
	SupportedEngine []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine `json:"SupportedEngine,omitempty" xml:"SupportedEngine,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) SetSupportedEngine(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines {
	s.SupportedEngine = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine struct {
	// example:
	//
	// hbase
	Engine                  *string                                                                                                                 `json:"Engine,omitempty" xml:"Engine,omitempty"`
	SupportedEngineVersions *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions `json:"SupportedEngineVersions,omitempty" xml:"SupportedEngineVersions,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) SetEngine(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine {
	s.Engine = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) SetSupportedEngineVersions(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine {
	s.SupportedEngineVersions = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions struct {
	SupportedEngineVersion []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion `json:"SupportedEngineVersion,omitempty" xml:"SupportedEngineVersion,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) SetSupportedEngineVersion(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions {
	s.SupportedEngineVersion = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion struct {
	SupportedCategories *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories `json:"SupportedCategories,omitempty" xml:"SupportedCategories,omitempty" type:"Struct"`
	// example:
	//
	// 2.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) SetSupportedCategories(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion {
	s.SupportedCategories = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) SetVersion(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion {
	s.Version = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories struct {
	SupportedCategories []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories `json:"SupportedCategories,omitempty" xml:"SupportedCategories,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) SetSupportedCategories(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories {
	s.SupportedCategories = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories struct {
	// example:
	//
	// cluster
	Category              *string                                                                                                                                                                                                  `json:"Category,omitempty" xml:"Category,omitempty"`
	SupportedStorageTypes *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes `json:"SupportedStorageTypes,omitempty" xml:"SupportedStorageTypes,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) SetCategory(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories {
	s.Category = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) SetSupportedStorageTypes(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories {
	s.SupportedStorageTypes = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes struct {
	SupportedStorageType []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType `json:"SupportedStorageType,omitempty" xml:"SupportedStorageType,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) SetSupportedStorageType(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes {
	s.SupportedStorageType = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType struct {
	CoreResources *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources `json:"CoreResources,omitempty" xml:"CoreResources,omitempty" type:"Struct"`
	// example:
	//
	// cloud_ssd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) SetCoreResources(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType {
	s.CoreResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) SetStorageType(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType {
	s.StorageType = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources struct {
	CoreResource []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource `json:"CoreResource,omitempty" xml:"CoreResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) SetCoreResource(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources {
	s.CoreResource = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource struct {
	DBInstanceStorageRange *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange `json:"DBInstanceStorageRange,omitempty" xml:"DBInstanceStorageRange,omitempty" type:"Struct"`
	// example:
	//
	// hbase.sn1.large
	InstanceType       *string                                                                                                                                                                                                                                                                 `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypeDetail *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail `json:"InstanceTypeDetail,omitempty" xml:"InstanceTypeDetail,omitempty" type:"Struct"`
	// example:
	//
	// 16
	MaxCoreCount *int32 `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetDBInstanceStorageRange(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.DBInstanceStorageRange = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetInstanceType(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.InstanceType = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetInstanceTypeDetail(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.InstanceTypeDetail = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetMaxCoreCount(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.MaxCoreCount = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange struct {
	// example:
	//
	// 8000
	MaxSize *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// example:
	//
	// 400
	MinSize *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	// example:
	//
	// 40
	StepSize *int32 `json:"StepSize,omitempty" xml:"StepSize,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) SetMaxSize(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	s.MaxSize = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) SetMinSize(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	s.MinSize = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) SetStepSize(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	s.StepSize = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail struct {
	// example:
	//
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 8
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) SetCpu(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail {
	s.Cpu = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) SetMem(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail {
	s.Mem = &v
	return s
}

type DescribeAvailableResourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeBackupPlanConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeBackupPlanConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanConfigRequest) SetClusterId(v string) *DescribeBackupPlanConfigRequest {
	s.ClusterId = &v
	return s
}

type DescribeBackupPlanConfigResponseBody struct {
	// example:
	//
	// 7
	FullBackupCycle *int32 `json:"FullBackupCycle,omitempty" xml:"FullBackupCycle,omitempty"`
	// example:
	//
	// 3
	MinHFileBackupCount *int32 `json:"MinHFileBackupCount,omitempty" xml:"MinHFileBackupCount,omitempty"`
	// example:
	//
	// 2020-11-09T18:00:00Z
	NextFullBackupDate *string `json:"NextFullBackupDate,omitempty" xml:"NextFullBackupDate,omitempty"`
	// example:
	//
	// 33A23201-6038-4A6A-B76A-61047EA04E6A
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tables    *DescribeBackupPlanConfigResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
}

func (s DescribeBackupPlanConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanConfigResponseBody) SetFullBackupCycle(v int32) *DescribeBackupPlanConfigResponseBody {
	s.FullBackupCycle = &v
	return s
}

func (s *DescribeBackupPlanConfigResponseBody) SetMinHFileBackupCount(v int32) *DescribeBackupPlanConfigResponseBody {
	s.MinHFileBackupCount = &v
	return s
}

func (s *DescribeBackupPlanConfigResponseBody) SetNextFullBackupDate(v string) *DescribeBackupPlanConfigResponseBody {
	s.NextFullBackupDate = &v
	return s
}

func (s *DescribeBackupPlanConfigResponseBody) SetRequestId(v string) *DescribeBackupPlanConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPlanConfigResponseBody) SetTables(v *DescribeBackupPlanConfigResponseBodyTables) *DescribeBackupPlanConfigResponseBody {
	s.Tables = v
	return s
}

type DescribeBackupPlanConfigResponseBodyTables struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlanConfigResponseBodyTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanConfigResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanConfigResponseBodyTables) SetTable(v []*string) *DescribeBackupPlanConfigResponseBodyTables {
	s.Table = v
	return s
}

type DescribeBackupPlanConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupPlanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupPlanConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanConfigResponse) SetHeaders(v map[string]*string) *DescribeBackupPlanConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPlanConfigResponse) SetStatusCode(v int32) *DescribeBackupPlanConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPlanConfigResponse) SetBody(v *DescribeBackupPlanConfigResponseBody) *DescribeBackupPlanConfigResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) SetClusterId(v string) *DescribeBackupPolicyRequest {
	s.ClusterId = &v
	return s
}

type DescribeBackupPolicyResponseBody struct {
	// example:
	//
	// 10
	BackupRetentionPeriod *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// example:
	//
	// 18:00Z
	PreferredBackupEndTimeUTC *string `json:"PreferredBackupEndTimeUTC,omitempty" xml:"PreferredBackupEndTimeUTC,omitempty"`
	// example:
	//
	// Friday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// example:
	//
	// 17:00Z
	PreferredBackupStartTimeUTC *string `json:"PreferredBackupStartTimeUTC,omitempty" xml:"PreferredBackupStartTimeUTC,omitempty"`
	// example:
	//
	// 01:00-02:00
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// example:
	//
	// 94AC38B6-7C6D-45B2-BC03-B8750071A482
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupEndTimeUTC(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupEndTimeUTC = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupStartTimeUTC(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupStartTimeUTC = &v
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

type DescribeBackupStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeBackupStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupStatusRequest) SetClusterId(v string) *DescribeBackupStatusRequest {
	s.ClusterId = &v
	return s
}

type DescribeBackupStatusResponseBody struct {
	// example:
	//
	// opened
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// example:
	//
	// bds-m5e54q06ceyhxxxx
	BdsClusterId *string `json:"BdsClusterId,omitempty" xml:"BdsClusterId,omitempty"`
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// F7E71430-A825-470A-B40B-DF3F3AAC9BEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupStatusResponseBody) SetBackupStatus(v string) *DescribeBackupStatusResponseBody {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupStatusResponseBody) SetBdsClusterId(v string) *DescribeBackupStatusResponseBody {
	s.BdsClusterId = &v
	return s
}

func (s *DescribeBackupStatusResponseBody) SetClusterId(v string) *DescribeBackupStatusResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupStatusResponseBody) SetRequestId(v string) *DescribeBackupStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBackupStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupStatusResponse) SetHeaders(v map[string]*string) *DescribeBackupStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupStatusResponse) SetStatusCode(v int32) *DescribeBackupStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupStatusResponse) SetBody(v *DescribeBackupStatusResponseBody) *DescribeBackupStatusResponse {
	s.Body = v
	return s
}

type DescribeBackupSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp169l540vc6c****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeBackupSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryRequest) SetClusterId(v string) *DescribeBackupSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupSummaryRequest) SetPageNumber(v int32) *DescribeBackupSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupSummaryRequest) SetPageSize(v int32) *DescribeBackupSummaryRequest {
	s.PageSize = &v
	return s
}

type DescribeBackupSummaryResponseBody struct {
	Full *DescribeBackupSummaryResponseBodyFull `json:"Full,omitempty" xml:"Full,omitempty" type:"Struct"`
	Incr *DescribeBackupSummaryResponseBodyIncr `json:"Incr,omitempty" xml:"Incr,omitempty" type:"Struct"`
	// example:
	//
	// 168793CB-7B31-43E7-ADAB-FE3E8D584D6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponseBody) SetFull(v *DescribeBackupSummaryResponseBodyFull) *DescribeBackupSummaryResponseBody {
	s.Full = v
	return s
}

func (s *DescribeBackupSummaryResponseBody) SetIncr(v *DescribeBackupSummaryResponseBodyIncr) *DescribeBackupSummaryResponseBody {
	s.Incr = v
	return s
}

func (s *DescribeBackupSummaryResponseBody) SetRequestId(v string) *DescribeBackupSummaryResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBackupSummaryResponseBodyFull struct {
	// example:
	//
	// false
	HasMore *string `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// example:
	//
	// 2020-11-09T18:00:00Z
	NextFullBackupDate *string `json:"NextFullBackupDate,omitempty" xml:"NextFullBackupDate,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  *DescribeBackupSummaryResponseBodyFullRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeBackupSummaryResponseBodyFull) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSummaryResponseBodyFull) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponseBodyFull) SetHasMore(v string) *DescribeBackupSummaryResponseBodyFull {
	s.HasMore = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) SetNextFullBackupDate(v string) *DescribeBackupSummaryResponseBodyFull {
	s.NextFullBackupDate = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) SetPageNumber(v int32) *DescribeBackupSummaryResponseBodyFull {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) SetPageSize(v int32) *DescribeBackupSummaryResponseBodyFull {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) SetRecords(v *DescribeBackupSummaryResponseBodyFullRecords) *DescribeBackupSummaryResponseBodyFull {
	s.Records = v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFull) SetTotal(v int32) *DescribeBackupSummaryResponseBodyFull {
	s.Total = &v
	return s
}

type DescribeBackupSummaryResponseBodyFullRecords struct {
	Record []*DescribeBackupSummaryResponseBodyFullRecordsRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s DescribeBackupSummaryResponseBodyFullRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSummaryResponseBodyFullRecords) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponseBodyFullRecords) SetRecord(v []*DescribeBackupSummaryResponseBodyFullRecordsRecord) *DescribeBackupSummaryResponseBodyFullRecords {
	s.Record = v
	return s
}

type DescribeBackupSummaryResponseBodyFullRecordsRecord struct {
	// example:
	//
	// 2020-11-02T18:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1.2 kB
	DataSize *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// example:
	//
	// 2020-11-02T18:02:04Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// 1/1
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// 20201103020000
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// 0.00 MB/s
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBackupSummaryResponseBodyFullRecordsRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSummaryResponseBodyFullRecordsRecord) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetCreateTime(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.CreateTime = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetDataSize(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.DataSize = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetFinishTime(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.FinishTime = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetProcess(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.Process = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetRecordId(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.RecordId = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetSpeed(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.Speed = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyFullRecordsRecord) SetStatus(v string) *DescribeBackupSummaryResponseBodyFullRecordsRecord {
	s.Status = &v
	return s
}

type DescribeBackupSummaryResponseBodyIncr struct {
	// example:
	//
	// 266 B
	BackupLogSize *string `json:"BackupLogSize,omitempty" xml:"BackupLogSize,omitempty"`
	// example:
	//
	// 2020-11-05T01:20:31Z
	Pos *string `json:"Pos,omitempty" xml:"Pos,omitempty"`
	// example:
	//
	// 0
	QueueLogNum *string `json:"QueueLogNum,omitempty" xml:"QueueLogNum,omitempty"`
	// example:
	//
	// 2
	RunningLogNum *string `json:"RunningLogNum,omitempty" xml:"RunningLogNum,omitempty"`
	// example:
	//
	// 0.00 MB/s
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBackupSummaryResponseBodyIncr) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSummaryResponseBodyIncr) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetBackupLogSize(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.BackupLogSize = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetPos(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.Pos = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetQueueLogNum(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.QueueLogNum = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetRunningLogNum(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.RunningLogNum = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetSpeed(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.Speed = &v
	return s
}

func (s *DescribeBackupSummaryResponseBodyIncr) SetStatus(v string) *DescribeBackupSummaryResponseBodyIncr {
	s.Status = &v
	return s
}

type DescribeBackupSummaryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponse) SetHeaders(v map[string]*string) *DescribeBackupSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupSummaryResponse) SetStatusCode(v int32) *DescribeBackupSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupSummaryResponse) SetBody(v *DescribeBackupSummaryResponseBody) *DescribeBackupSummaryResponse {
	s.Body = v
	return s
}

type DescribeBackupTablesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2020110302xxxx
	BackupRecordId *string `json:"BackupRecordId,omitempty" xml:"BackupRecordId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeBackupTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesRequest) SetBackupRecordId(v string) *DescribeBackupTablesRequest {
	s.BackupRecordId = &v
	return s
}

func (s *DescribeBackupTablesRequest) SetClusterId(v string) *DescribeBackupTablesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupTablesRequest) SetPageNumber(v int32) *DescribeBackupTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupTablesRequest) SetPageSize(v int32) *DescribeBackupTablesRequest {
	s.PageSize = &v
	return s
}

type DescribeBackupTablesResponseBody struct {
	BackupRecords *DescribeBackupTablesResponseBodyBackupRecords `json:"BackupRecords,omitempty" xml:"BackupRecords,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 01262E9C-B0CC-4663-82FA-D50173649F92
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tables    *DescribeBackupTablesResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeBackupTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesResponseBody) SetBackupRecords(v *DescribeBackupTablesResponseBodyBackupRecords) *DescribeBackupTablesResponseBody {
	s.BackupRecords = v
	return s
}

func (s *DescribeBackupTablesResponseBody) SetPageNumber(v int32) *DescribeBackupTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupTablesResponseBody) SetPageSize(v int32) *DescribeBackupTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupTablesResponseBody) SetRequestId(v string) *DescribeBackupTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupTablesResponseBody) SetTables(v *DescribeBackupTablesResponseBodyTables) *DescribeBackupTablesResponseBody {
	s.Tables = v
	return s
}

func (s *DescribeBackupTablesResponseBody) SetTotal(v int64) *DescribeBackupTablesResponseBody {
	s.Total = &v
	return s
}

type DescribeBackupTablesResponseBodyBackupRecords struct {
	BackupRecord []*DescribeBackupTablesResponseBodyBackupRecordsBackupRecord `json:"BackupRecord,omitempty" xml:"BackupRecord,omitempty" type:"Repeated"`
}

func (s DescribeBackupTablesResponseBodyBackupRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTablesResponseBodyBackupRecords) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesResponseBodyBackupRecords) SetBackupRecord(v []*DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) *DescribeBackupTablesResponseBodyBackupRecords {
	s.BackupRecord = v
	return s
}

type DescribeBackupTablesResponseBodyBackupRecordsBackupRecord struct {
	// example:
	//
	// 1.2 kB
	DataSize *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// example:
	//
	// 2020-11-02T18:00:05Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 14/14
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// 0.00 MB/s
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// example:
	//
	// 2020-11-02T18:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCEEDED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// default:test1
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetDataSize(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.DataSize = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetEndTime(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetMessage(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.Message = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetProcess(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.Process = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetSpeed(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.Speed = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetStartTime(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetState(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.State = &v
	return s
}

func (s *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord) SetTable(v string) *DescribeBackupTablesResponseBodyBackupRecordsBackupRecord {
	s.Table = &v
	return s
}

type DescribeBackupTablesResponseBodyTables struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeBackupTablesResponseBodyTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTablesResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesResponseBodyTables) SetTable(v []*string) *DescribeBackupTablesResponseBodyTables {
	s.Table = v
	return s
}

type DescribeBackupTablesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupTablesResponse) SetHeaders(v map[string]*string) *DescribeBackupTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupTablesResponse) SetStatusCode(v int32) *DescribeBackupTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupTablesResponse) SetBody(v *DescribeBackupTablesResponseBody) *DescribeBackupTablesResponse {
	s.Body = v
	return s
}

type DescribeBackupsRequest struct {
	// example:
	//
	// job-xxxx
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 2020-12-23 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2020-12-23T15:59:59Z
	EndTimeUTC *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2020-12-13 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 2020-12-12T16:00:00Z
	StartTimeUTC *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
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

func (s *DescribeBackupsRequest) SetClusterId(v string) *DescribeBackupsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTime(v string) *DescribeBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTimeUTC(v string) *DescribeBackupsRequest {
	s.EndTimeUTC = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageNumber(v string) *DescribeBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageSize(v string) *DescribeBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTime(v string) *DescribeBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTimeUTC(v string) *DescribeBackupsRequest {
	s.StartTimeUTC = &v
	return s
}

type DescribeBackupsResponseBody struct {
	Backups *DescribeBackupsResponseBodyBackups `json:"Backups,omitempty" xml:"Backups,omitempty" type:"Struct"`
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// A93CE373-0FDE-4CCB-9DBA-6700906825ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) SetBackups(v *DescribeBackupsResponseBodyBackups) *DescribeBackupsResponseBody {
	s.Backups = v
	return s
}

func (s *DescribeBackupsResponseBody) SetEnableStatus(v string) *DescribeBackupsResponseBody {
	s.EnableStatus = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageNumber(v int32) *DescribeBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageSize(v int32) *DescribeBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetRequestId(v string) *DescribeBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetTotalCount(v int32) *DescribeBackupsResponseBody {
	s.TotalCount = &v
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
	// example:
	//
	// 1
	BackupDBNames *string `json:"BackupDBNames,omitempty" xml:"BackupDBNames,omitempty"`
	// example:
	//
	// 1
	BackupDownloadURL *string `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	// example:
	//
	// 2020-12-23 17:25:24
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// example:
	//
	// 2020-12-23T09:25:24Z
	BackupEndTimeUTC *string `json:"BackupEndTimeUTC,omitempty" xml:"BackupEndTimeUTC,omitempty"`
	// example:
	//
	// 511876087
	BackupId *int32 `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// example:
	//
	// P
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// example:
	//
	// 0
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// example:
	//
	// 0.00
	BackupSize *string `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// example:
	//
	// 2020-12-23 17:25:08
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// example:
	//
	// 2020-12-23T09:25:08Z
	BackupStartTimeUTC *string `json:"BackupStartTimeUTC,omitempty" xml:"BackupStartTimeUTC,omitempty"`
	// example:
	//
	// 0
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// example:
	//
	// F
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
}

func (s DescribeBackupsResponseBodyBackupsBackup) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyBackupsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupDBNames(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupDBNames = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupDownloadURL(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupDownloadURL = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupEndTimeUTC(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupEndTimeUTC = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupId(v int32) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupMethod(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupMode(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupSize(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStartTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStartTimeUTC(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStartTimeUTC = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStatus(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupType(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupType = &v
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

type DescribeClusterConnectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeClusterConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionRequest) SetClusterId(v string) *DescribeClusterConnectionRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterConnectionRequest) SetRegionId(v string) *DescribeClusterConnectionRequest {
	s.RegionId = &v
	return s
}

type DescribeClusterConnectionResponseBody struct {
	// example:
	//
	// hbaseue
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// true
	IsMultimod *string `json:"IsMultimod,omitempty" xml:"IsMultimod,omitempty"`
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// example:
	//
	// 70220050-A465-5DCC-8C0C-C38C6E3DB24D
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceConnAddrs    *DescribeClusterConnectionResponseBodyServiceConnAddrs    `json:"ServiceConnAddrs,omitempty" xml:"ServiceConnAddrs,omitempty" type:"Struct"`
	SlbConnAddrs        *DescribeClusterConnectionResponseBodySlbConnAddrs        `json:"SlbConnAddrs,omitempty" xml:"SlbConnAddrs,omitempty" type:"Struct"`
	ThriftConn          *DescribeClusterConnectionResponseBodyThriftConn          `json:"ThriftConn,omitempty" xml:"ThriftConn,omitempty" type:"Struct"`
	UiProxyConnAddrInfo *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo `json:"UiProxyConnAddrInfo,omitempty" xml:"UiProxyConnAddrInfo,omitempty" type:"Struct"`
	// example:
	//
	// vsw-bp1foll427ze3d4ps****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp15s22y1a7sff5gj****
	VpcId       *string                                           `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZkConnAddrs *DescribeClusterConnectionResponseBodyZkConnAddrs `json:"ZkConnAddrs,omitempty" xml:"ZkConnAddrs,omitempty" type:"Struct"`
}

func (s DescribeClusterConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBody) SetDbType(v string) *DescribeClusterConnectionResponseBody {
	s.DbType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetIsMultimod(v string) *DescribeClusterConnectionResponseBody {
	s.IsMultimod = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetNetType(v string) *DescribeClusterConnectionResponseBody {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetRequestId(v string) *DescribeClusterConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetServiceConnAddrs(v *DescribeClusterConnectionResponseBodyServiceConnAddrs) *DescribeClusterConnectionResponseBody {
	s.ServiceConnAddrs = v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetSlbConnAddrs(v *DescribeClusterConnectionResponseBodySlbConnAddrs) *DescribeClusterConnectionResponseBody {
	s.SlbConnAddrs = v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetThriftConn(v *DescribeClusterConnectionResponseBodyThriftConn) *DescribeClusterConnectionResponseBody {
	s.ThriftConn = v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetUiProxyConnAddrInfo(v *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) *DescribeClusterConnectionResponseBody {
	s.UiProxyConnAddrInfo = v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetVSwitchId(v string) *DescribeClusterConnectionResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetVpcId(v string) *DescribeClusterConnectionResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetZkConnAddrs(v *DescribeClusterConnectionResponseBodyZkConnAddrs) *DescribeClusterConnectionResponseBody {
	s.ZkConnAddrs = v
	return s
}

type DescribeClusterConnectionResponseBodyServiceConnAddrs struct {
	ServiceConnAddr []*DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr `json:"ServiceConnAddr,omitempty" xml:"ServiceConnAddr,omitempty" type:"Repeated"`
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrs) SetServiceConnAddr(v []*DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) *DescribeClusterConnectionResponseBodyServiceConnAddrs {
	s.ServiceConnAddr = v
	return s
}

type DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr struct {
	ConnAddrInfo *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo `json:"ConnAddrInfo,omitempty" xml:"ConnAddrInfo,omitempty" type:"Struct"`
	// example:
	//
	// PhoenixConnAddr
	ConnType *string `json:"ConnType,omitempty" xml:"ConnType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) SetConnAddrInfo(v *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr {
	s.ConnAddrInfo = v
	return s
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) SetConnType(v string) *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr {
	s.ConnType = &v
	return s
}

type DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo struct {
	// example:
	//
	// hb-****-proxy-phoenix.hbase.rds.aliyuncs.com
	ConnAddr *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	// example:
	//
	// 8765
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	// example:
	//
	// 2
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) SetConnAddr(v string) *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) SetConnAddrPort(v string) *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) SetNetType(v string) *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	s.NetType = &v
	return s
}

type DescribeClusterConnectionResponseBodySlbConnAddrs struct {
	SlbConnAddr []*DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr `json:"SlbConnAddr,omitempty" xml:"SlbConnAddr,omitempty" type:"Repeated"`
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrs) SetSlbConnAddr(v []*DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) *DescribeClusterConnectionResponseBodySlbConnAddrs {
	s.SlbConnAddr = v
	return s
}

type DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr struct {
	ConnAddrInfo *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo `json:"ConnAddrInfo,omitempty" xml:"ConnAddrInfo,omitempty" type:"Struct"`
	// example:
	//
	// hbaseue
	SlbType *string `json:"SlbType,omitempty" xml:"SlbType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) SetConnAddrInfo(v *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr {
	s.ConnAddrInfo = v
	return s
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) SetSlbType(v string) *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr {
	s.SlbType = &v
	return s
}

type DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo struct {
	// example:
	//
	// ld-bp150tns0sjxs****-proxy-hbaseue-pub.hbaseue.rds.aliyuncs.com
	ConnAddr *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	// example:
	//
	// 9190
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	// example:
	//
	// 0
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) SetConnAddr(v string) *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) SetConnAddrPort(v string) *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) SetNetType(v string) *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	s.NetType = &v
	return s
}

type DescribeClusterConnectionResponseBodyThriftConn struct {
	// example:
	//
	// hb-bp1u0639js2h7****-proxy-thrift.hbase.rds.aliyuncs.com
	ConnAddr *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	// example:
	//
	// 9099
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	// example:
	//
	// 2
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodyThriftConn) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyThriftConn) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyThriftConn) SetConnAddr(v string) *DescribeClusterConnectionResponseBodyThriftConn {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyThriftConn) SetConnAddrPort(v string) *DescribeClusterConnectionResponseBodyThriftConn {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyThriftConn) SetNetType(v string) *DescribeClusterConnectionResponseBodyThriftConn {
	s.NetType = &v
	return s
}

type DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo struct {
	// example:
	//
	// ld-bp150tns0sjxs****-master1-001.hbaseue.rds.aliyuncs.com
	ConnAddr *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	// example:
	//
	// 443
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	// example:
	//
	// PUBLIC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) SetConnAddr(v string) *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) SetConnAddrPort(v string) *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) SetNetType(v string) *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo {
	s.NetType = &v
	return s
}

type DescribeClusterConnectionResponseBodyZkConnAddrs struct {
	ZkConnAddr []*DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr `json:"ZkConnAddr,omitempty" xml:"ZkConnAddr,omitempty" type:"Repeated"`
}

func (s DescribeClusterConnectionResponseBodyZkConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyZkConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrs) SetZkConnAddr(v []*DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) *DescribeClusterConnectionResponseBodyZkConnAddrs {
	s.ZkConnAddr = v
	return s
}

type DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr struct {
	// example:
	//
	// ld-bp150tns0sjxs****-master1-001.hbaseue.rds.aliyuncs.com
	ConnAddr *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	// example:
	//
	// 2181
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	// example:
	//
	// 2
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) SetConnAddr(v string) *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) SetConnAddrPort(v string) *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) SetNetType(v string) *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr {
	s.NetType = &v
	return s
}

type DescribeClusterConnectionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterConnectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponse) SetHeaders(v map[string]*string) *DescribeClusterConnectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterConnectionResponse) SetStatusCode(v int32) *DescribeClusterConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterConnectionResponse) SetBody(v *DescribeClusterConnectionResponseBody) *DescribeClusterConnectionResponse {
	s.Body = v
	return s
}

type DescribeColdStorageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1uoihlf82e8****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeColdStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeColdStorageRequest) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageRequest) SetClusterId(v string) *DescribeColdStorageRequest {
	s.ClusterId = &v
	return s
}

type DescribeColdStorageResponseBody struct {
	// example:
	//
	// ld-bp1uoihlf82e8****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 800
	ColdStorageSize *string `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	// example:
	//
	// BdsColdStorage
	ColdStorageType *string `json:"ColdStorageType,omitempty" xml:"ColdStorageType,omitempty"`
	// example:
	//
	// 20.00
	ColdStorageUseAmount *string `json:"ColdStorageUseAmount,omitempty" xml:"ColdStorageUseAmount,omitempty"`
	// example:
	//
	// 20.00
	ColdStorageUsePercent *string `json:"ColdStorageUsePercent,omitempty" xml:"ColdStorageUsePercent,omitempty"`
	// example:
	//
	// open
	OpenStatus *string `json:"OpenStatus,omitempty" xml:"OpenStatus,omitempty"`
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// DCB9479E-F05F-4D1C-AFB7-C639B87764B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeColdStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeColdStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageResponseBody) SetClusterId(v string) *DescribeColdStorageResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetColdStorageSize(v string) *DescribeColdStorageResponseBody {
	s.ColdStorageSize = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetColdStorageType(v string) *DescribeColdStorageResponseBody {
	s.ColdStorageType = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetColdStorageUseAmount(v string) *DescribeColdStorageResponseBody {
	s.ColdStorageUseAmount = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetColdStorageUsePercent(v string) *DescribeColdStorageResponseBody {
	s.ColdStorageUsePercent = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetOpenStatus(v string) *DescribeColdStorageResponseBody {
	s.OpenStatus = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetPayType(v string) *DescribeColdStorageResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeColdStorageResponseBody) SetRequestId(v string) *DescribeColdStorageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeColdStorageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeColdStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeColdStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeColdStorageResponse) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageResponse) SetHeaders(v map[string]*string) *DescribeColdStorageResponse {
	s.Headers = v
	return s
}

func (s *DescribeColdStorageResponse) SetStatusCode(v int32) *DescribeColdStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeColdStorageResponse) SetBody(v *DescribeColdStorageResponseBody) *DescribeColdStorageResponse {
	s.Body = v
	return s
}

type DescribeDBInstanceUsageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp1u0639js2h7****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeDBInstanceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceUsageRequest) SetClusterId(v string) *DescribeDBInstanceUsageRequest {
	s.ClusterId = &v
	return s
}

type DescribeDBInstanceUsageResponseBody struct {
	// example:
	//
	// A2D841CE-D066-53E8-B9AC-3731DCC85397
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {\\"disk_usage_record\\":{\\"disk_used\\":\\"0.9GB\\",\\"disk_total\\":\\"1156.1GB\\",\\"usage_rate\\":\\"1%\\"}}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeDBInstanceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceUsageResponseBody) SetRequestId(v string) *DescribeDBInstanceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceUsageResponseBody) SetResult(v string) *DescribeDBInstanceUsageResponseBody {
	s.Result = &v
	return s
}

type DescribeDBInstanceUsageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceUsageResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceUsageResponse) SetStatusCode(v int32) *DescribeDBInstanceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceUsageResponse) SetBody(v *DescribeDBInstanceUsageResponseBody) *DescribeDBInstanceUsageResponse {
	s.Body = v
	return s
}

type DescribeDeletedInstancesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDeletedInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeletedInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeletedInstancesRequest) SetPageNumber(v int32) *DescribeDeletedInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeletedInstancesRequest) SetPageSize(v int32) *DescribeDeletedInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeletedInstancesRequest) SetRegionId(v string) *DescribeDeletedInstancesRequest {
	s.RegionId = &v
	return s
}

type DescribeDeletedInstancesResponseBody struct {
	Instances *DescribeDeletedInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0CAC5702-C862-44C0-AD54-C9CE70F4B246
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeletedInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeletedInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeletedInstancesResponseBody) SetInstances(v *DescribeDeletedInstancesResponseBodyInstances) *DescribeDeletedInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeDeletedInstancesResponseBody) SetPageNumber(v int32) *DescribeDeletedInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBody) SetPageSize(v int32) *DescribeDeletedInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBody) SetRequestId(v string) *DescribeDeletedInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBody) SetTotalCount(v int64) *DescribeDeletedInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDeletedInstancesResponseBodyInstances struct {
	Instance []*DescribeDeletedInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeDeletedInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeletedInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDeletedInstancesResponseBodyInstances) SetInstance(v []*DescribeDeletedInstancesResponseBodyInstancesInstance) *DescribeDeletedInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

type DescribeDeletedInstancesResponseBodyInstancesInstance struct {
	// example:
	//
	// cluster
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// example:
	//
	// 2020-11-02T07:16:07Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 2020-11-02T07:27:24Z
	DeleteTime *string `json:"DeleteTime,omitempty" xml:"DeleteTime,omitempty"`
	// example:
	//
	// hbase
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// hb-bp10q7n2zdw12xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// e2e-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 2.0
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// example:
	//
	// null
	ModuleStackVersion *string `json:"ModuleStackVersion,omitempty" xml:"ModuleStackVersion,omitempty"`
	// example:
	//
	// null
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// DELETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDeletedInstancesResponseBodyInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeletedInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetClusterType(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.ClusterType = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetCreatedTime(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetDeleteTime(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.DeleteTime = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetEngine(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetMajorVersion(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.MajorVersion = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetModuleStackVersion(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.ModuleStackVersion = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetParentId(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.ParentId = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetRegionId(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetStatus(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeDeletedInstancesResponseBodyInstancesInstance) SetZoneId(v string) *DescribeDeletedInstancesResponseBodyInstancesInstance {
	s.ZoneId = &v
	return s
}

type DescribeDeletedInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeletedInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeletedInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeletedInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeletedInstancesResponse) SetHeaders(v map[string]*string) *DescribeDeletedInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeletedInstancesResponse) SetStatusCode(v int32) *DescribeDeletedInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeletedInstancesResponse) SetBody(v *DescribeDeletedInstancesResponseBody) *DescribeDeletedInstancesResponse {
	s.Body = v
	return s
}

type DescribeDiskWarningLineRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp1bl7iqzkahmyxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeDiskWarningLineRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskWarningLineRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskWarningLineRequest) SetClusterId(v string) *DescribeDiskWarningLineRequest {
	s.ClusterId = &v
	return s
}

type DescribeDiskWarningLineResponseBody struct {
	// example:
	//
	// 08DF8283-D290-4107-931E-7913D6D3480D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 80
	WarningLine *string `json:"WarningLine,omitempty" xml:"WarningLine,omitempty"`
}

func (s DescribeDiskWarningLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskWarningLineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskWarningLineResponseBody) SetRequestId(v string) *DescribeDiskWarningLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskWarningLineResponseBody) SetWarningLine(v string) *DescribeDiskWarningLineResponseBody {
	s.WarningLine = &v
	return s
}

type DescribeDiskWarningLineResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskWarningLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskWarningLineResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskWarningLineResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskWarningLineResponse) SetHeaders(v map[string]*string) *DescribeDiskWarningLineResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskWarningLineResponse) SetStatusCode(v int32) *DescribeDiskWarningLineResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskWarningLineResponse) SetBody(v *DescribeDiskWarningLineResponseBody) *DescribeDiskWarningLineResponse {
	s.Body = v
	return s
}

type DescribeEndpointsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsRequest) SetClusterId(v string) *DescribeEndpointsRequest {
	s.ClusterId = &v
	return s
}

type DescribeEndpointsResponseBody struct {
	ConnAddrs *DescribeEndpointsResponseBodyConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Struct"`
	// example:
	//
	// hbaseue
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// example:
	//
	// F072593C-5234-5B56-9F63-3C7A3AD85D66
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// vsw-bp1foll427ze3d4ps****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp15s22y1a7sff5gj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBody) SetConnAddrs(v *DescribeEndpointsResponseBodyConnAddrs) *DescribeEndpointsResponseBody {
	s.ConnAddrs = v
	return s
}

func (s *DescribeEndpointsResponseBody) SetEngine(v string) *DescribeEndpointsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeEndpointsResponseBody) SetNetType(v string) *DescribeEndpointsResponseBody {
	s.NetType = &v
	return s
}

func (s *DescribeEndpointsResponseBody) SetRequestId(v string) *DescribeEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndpointsResponseBody) SetVSwitchId(v string) *DescribeEndpointsResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeEndpointsResponseBody) SetVpcId(v string) *DescribeEndpointsResponseBody {
	s.VpcId = &v
	return s
}

type DescribeEndpointsResponseBodyConnAddrs struct {
	ConnAddrInfo []*DescribeEndpointsResponseBodyConnAddrsConnAddrInfo `json:"ConnAddrInfo,omitempty" xml:"ConnAddrInfo,omitempty" type:"Repeated"`
}

func (s DescribeEndpointsResponseBodyConnAddrs) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponseBodyConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBodyConnAddrs) SetConnAddrInfo(v []*DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) *DescribeEndpointsResponseBodyConnAddrs {
	s.ConnAddrInfo = v
	return s
}

type DescribeEndpointsResponseBodyConnAddrsConnAddrInfo struct {
	// example:
	//
	// ****
	ConnAddr *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	// example:
	//
	// ****
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	// example:
	//
	// zkConn
	ConnType *string `json:"ConnType,omitempty" xml:"ConnType,omitempty"`
	// example:
	//
	// 2
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) SetConnAddr(v string) *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo {
	s.ConnAddr = &v
	return s
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) SetConnAddrPort(v string) *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) SetConnType(v string) *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo {
	s.ConnType = &v
	return s
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) SetNetType(v string) *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo {
	s.NetType = &v
	return s
}

type DescribeEndpointsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponse) SetHeaders(v map[string]*string) *DescribeEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndpointsResponse) SetStatusCode(v int32) *DescribeEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEndpointsResponse) SetBody(v *DescribeEndpointsResponseBody) *DescribeEndpointsResponse {
	s.Body = v
	return s
}

type DescribeInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) SetClusterId(v string) *DescribeInstanceRequest {
	s.ClusterId = &v
	return s
}

type DescribeInstanceResponseBody struct {
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// example:
	//
	// open
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// testhbase
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// cluster
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// example:
	//
	// 800
	ColdStorageSize *int32 `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	// example:
	//
	// open
	ColdStorageStatus *string `json:"ColdStorageStatus,omitempty" xml:"ColdStorageStatus,omitempty"`
	// example:
	//
	// true
	ConfirmMaintainTime *string `json:"ConfirmMaintainTime,omitempty" xml:"ConfirmMaintainTime,omitempty"`
	// example:
	//
	// 4
	CoreDiskCount *string `json:"CoreDiskCount,omitempty" xml:"CoreDiskCount,omitempty"`
	// example:
	//
	// 100
	CoreDiskSize *int32 `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	// example:
	//
	// cloud_ssd
	CoreDiskType *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	// example:
	//
	// hbase.sn2.2xlarge
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// example:
	//
	// 2
	CoreNodeCount *int32 `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	// example:
	//
	// 2021-07-19T11:23:22
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 2021-07-19T03:23:22Z
	CreatedTimeUTC *string `json:"CreatedTimeUTC,omitempty" xml:"CreatedTimeUTC,omitempty"`
	// example:
	//
	// 12
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	EnableHbaseProxy *bool `json:"EnableHbaseProxy,omitempty" xml:"EnableHbaseProxy,omitempty"`
	// example:
	//
	// 0d2470df-da7b-4786-b981-9a164dae****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// example:
	//
	// NoEncryption
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// example:
	//
	// hbaseue
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2022-02-24T00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2022-02-23T16:00:00Z
	ExpireTimeUTC *string `json:"ExpireTimeUTC,omitempty" xml:"ExpireTimeUTC,omitempty"`
	// example:
	//
	// ld-bp150tns0sjxs****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// testhbase
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// false
	IsDeletionProtection *bool `json:"IsDeletionProtection,omitempty" xml:"IsDeletionProtection,omitempty"`
	// example:
	//
	// true
	IsHa *bool `json:"IsHa,omitempty" xml:"IsHa,omitempty"`
	// example:
	//
	// true
	IsLatestVersion *bool `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// example:
	//
	// true
	IsMultiModel *bool `json:"IsMultiModel,omitempty" xml:"IsMultiModel,omitempty"`
	// example:
	//
	// 2.3.2
	LproxyMinorVersion *string `json:"LproxyMinorVersion,omitempty" xml:"LproxyMinorVersion,omitempty"`
	// example:
	//
	// 22:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// example:
	//
	// 18:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// example:
	//
	// 2.0
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// example:
	//
	// 0
	MasterDiskSize *int32 `json:"MasterDiskSize,omitempty" xml:"MasterDiskSize,omitempty"`
	// example:
	//
	// cloud_efficiency
	MasterDiskType *string `json:"MasterDiskType,omitempty" xml:"MasterDiskType,omitempty"`
	// example:
	//
	// hbase.sn2.large
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	// example:
	//
	// 2
	MasterNodeCount *int32 `json:"MasterNodeCount,omitempty" xml:"MasterNodeCount,omitempty"`
	// example:
	//
	// 2.2.9.1
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// example:
	//
	// 0
	ModuleId *int32 `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// example:
	//
	// phoenxi:4.0
	ModuleStackVersion *string `json:"ModuleStackVersion,omitempty" xml:"ModuleStackVersion,omitempty"`
	// example:
	//
	// false
	NeedUpgrade      *bool                                         `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	NeedUpgradeComps *DescribeInstanceResponseBodyNeedUpgradeComps `json:"NeedUpgradeComps,omitempty" xml:"NeedUpgradeComps,omitempty" type:"Struct"`
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// ld-uf699153o1m2l****
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 3F429923-B6F6-52C5-9C2A-5B8A8C6BBA66
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// ACTIVATION
	Status *string                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// example:
	//
	// 25.00
	TaskProgress *string `json:"TaskProgress,omitempty" xml:"TaskProgress,omitempty"`
	TaskStatus   *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// vpc-bp15s22y1a7sff5gj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-bp1foll427ze3d4ps****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetAutoRenewal(v bool) *DescribeInstanceResponseBody {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetBackupStatus(v string) *DescribeInstanceResponseBody {
	s.BackupStatus = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetClusterId(v string) *DescribeInstanceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetClusterName(v string) *DescribeInstanceResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetClusterType(v string) *DescribeInstanceResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetColdStorageSize(v int32) *DescribeInstanceResponseBody {
	s.ColdStorageSize = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetColdStorageStatus(v string) *DescribeInstanceResponseBody {
	s.ColdStorageStatus = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetConfirmMaintainTime(v string) *DescribeInstanceResponseBody {
	s.ConfirmMaintainTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCoreDiskCount(v string) *DescribeInstanceResponseBody {
	s.CoreDiskCount = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCoreDiskSize(v int32) *DescribeInstanceResponseBody {
	s.CoreDiskSize = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCoreDiskType(v string) *DescribeInstanceResponseBody {
	s.CoreDiskType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCoreInstanceType(v string) *DescribeInstanceResponseBody {
	s.CoreInstanceType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCoreNodeCount(v int32) *DescribeInstanceResponseBody {
	s.CoreNodeCount = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCreatedTime(v string) *DescribeInstanceResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCreatedTimeUTC(v string) *DescribeInstanceResponseBody {
	s.CreatedTimeUTC = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetDuration(v int32) *DescribeInstanceResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEnableHbaseProxy(v bool) *DescribeInstanceResponseBody {
	s.EnableHbaseProxy = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEncryptionKey(v string) *DescribeInstanceResponseBody {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEncryptionType(v string) *DescribeInstanceResponseBody {
	s.EncryptionType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEngine(v string) *DescribeInstanceResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetExpireTime(v string) *DescribeInstanceResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetExpireTimeUTC(v string) *DescribeInstanceResponseBody {
	s.ExpireTimeUTC = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceId(v string) *DescribeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceName(v string) *DescribeInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIsDeletionProtection(v bool) *DescribeInstanceResponseBody {
	s.IsDeletionProtection = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIsHa(v bool) *DescribeInstanceResponseBody {
	s.IsHa = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIsLatestVersion(v bool) *DescribeInstanceResponseBody {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIsMultiModel(v bool) *DescribeInstanceResponseBody {
	s.IsMultiModel = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetLproxyMinorVersion(v string) *DescribeInstanceResponseBody {
	s.LproxyMinorVersion = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMaintainEndTime(v string) *DescribeInstanceResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMaintainStartTime(v string) *DescribeInstanceResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMajorVersion(v string) *DescribeInstanceResponseBody {
	s.MajorVersion = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMasterDiskSize(v int32) *DescribeInstanceResponseBody {
	s.MasterDiskSize = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMasterDiskType(v string) *DescribeInstanceResponseBody {
	s.MasterDiskType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMasterInstanceType(v string) *DescribeInstanceResponseBody {
	s.MasterInstanceType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMasterNodeCount(v int32) *DescribeInstanceResponseBody {
	s.MasterNodeCount = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMinorVersion(v string) *DescribeInstanceResponseBody {
	s.MinorVersion = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetModuleId(v int32) *DescribeInstanceResponseBody {
	s.ModuleId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetModuleStackVersion(v string) *DescribeInstanceResponseBody {
	s.ModuleStackVersion = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetNeedUpgrade(v bool) *DescribeInstanceResponseBody {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetNeedUpgradeComps(v *DescribeInstanceResponseBodyNeedUpgradeComps) *DescribeInstanceResponseBody {
	s.NeedUpgradeComps = v
	return s
}

func (s *DescribeInstanceResponseBody) SetNetworkType(v string) *DescribeInstanceResponseBody {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetParentId(v string) *DescribeInstanceResponseBody {
	s.ParentId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetPayType(v string) *DescribeInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRegionId(v string) *DescribeInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetResourceGroupId(v string) *DescribeInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetStatus(v string) *DescribeInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetTags(v *DescribeInstanceResponseBodyTags) *DescribeInstanceResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeInstanceResponseBody) SetTaskProgress(v string) *DescribeInstanceResponseBody {
	s.TaskProgress = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetTaskStatus(v string) *DescribeInstanceResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetVpcId(v string) *DescribeInstanceResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetVswitchId(v string) *DescribeInstanceResponseBody {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetZoneId(v string) *DescribeInstanceResponseBody {
	s.ZoneId = &v
	return s
}

type DescribeInstanceResponseBodyNeedUpgradeComps struct {
	Comps []*string `json:"Comps,omitempty" xml:"Comps,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyNeedUpgradeComps) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyNeedUpgradeComps) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyNeedUpgradeComps) SetComps(v []*string) *DescribeInstanceResponseBodyNeedUpgradeComps {
	s.Comps = v
	return s
}

type DescribeInstanceResponseBodyTags struct {
	Tag []*DescribeInstanceResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyTags) SetTag(v []*DescribeInstanceResponseBodyTagsTag) *DescribeInstanceResponseBodyTags {
	s.Tag = v
	return s
}

type DescribeInstanceResponseBodyTagsTag struct {
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceResponseBodyTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyTagsTag) SetKey(v string) *DescribeInstanceResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeInstanceResponseBodyTagsTag) SetValue(v string) *DescribeInstanceResponseBodyTagsTag {
	s.Value = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetStatusCode(v int32) *DescribeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeInstanceTypeRequest struct {
	// example:
	//
	// hbase.n2.4xlarge
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
	InstanceTypeSpecList *DescribeInstanceTypeResponseBodyInstanceTypeSpecList `json:"InstanceTypeSpecList,omitempty" xml:"InstanceTypeSpecList,omitempty" type:"Struct"`
	// example:
	//
	// DD23BBB4-64C2-42A4-B2E2-7E56C7AA815A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeResponseBody) SetInstanceTypeSpecList(v *DescribeInstanceTypeResponseBodyInstanceTypeSpecList) *DescribeInstanceTypeResponseBody {
	s.InstanceTypeSpecList = v
	return s
}

func (s *DescribeInstanceTypeResponseBody) SetRequestId(v string) *DescribeInstanceTypeResponseBody {
	s.RequestId = &v
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
	// example:
	//
	// 8
	CpuSize *int64 `json:"CpuSize,omitempty" xml:"CpuSize,omitempty"`
	// example:
	//
	// hbase.n2.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 16
	MemSize *int64 `json:"MemSize,omitempty" xml:"MemSize,omitempty"`
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

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) SetInstanceType(v string) *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) SetMemSize(v int64) *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec {
	s.MemSize = &v
	return s
}

type DescribeInstanceTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeInstanceTypeResponse) SetStatusCode(v int32) *DescribeInstanceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceTypeResponse) SetBody(v *DescribeInstanceTypeResponseBody) *DescribeInstanceTypeResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	// example:
	//
	// hb-bp1u0639js2h7****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// hbase
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-4f51d54g5****
	ResourceGroupId *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetClusterId(v string) *DescribeInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstancesRequest) SetClusterName(v string) *DescribeInstancesRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeInstancesRequest) SetDbType(v string) *DescribeInstancesRequest {
	s.DbType = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v int32) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetRegionId(v string) *DescribeInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

type DescribeInstancesRequestTag struct {
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTag) SetKey(v string) *DescribeInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTag) SetValue(v string) *DescribeInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeInstancesResponseBody struct {
	Instances *DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// EBECBF12-2E34-41BE-8DE9-FC3700D4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 18
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetInstances(v *DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageNumber(v int32) *DescribeInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	Instance []*DescribeInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetInstance(v []*DescribeInstancesResponseBodyInstancesInstance) *DescribeInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

type DescribeInstancesResponseBodyInstancesInstance struct {
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// example:
	//
	// open
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// example:
	//
	// hb-bp1u0639js2h7****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// cluster
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// example:
	//
	// open
	ColdStorageStatus *string `json:"ColdStorageStatus,omitempty" xml:"ColdStorageStatus,omitempty"`
	// example:
	//
	// 2
	CoreDiskCount *string `json:"CoreDiskCount,omitempty" xml:"CoreDiskCount,omitempty"`
	// example:
	//
	// 100
	CoreDiskSize *int32 `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	// example:
	//
	// cloud_efficiency
	CoreDiskType *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	// example:
	//
	// hbase.sn1.large
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// example:
	//
	// 2
	CoreNodeCount *int32 `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	// example:
	//
	// 2019-09-12T14:40:46
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 2019-09-12T14:40:46Z
	CreatedTimeUTC *string `json:"CreatedTimeUTC,omitempty" xml:"CreatedTimeUTC,omitempty"`
	// example:
	//
	// 12
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// hbase
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2019-10-12T14:40:46
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2019-10-12T14:40:46Z
	ExpireTimeUTC *string `json:"ExpireTimeUTC,omitempty" xml:"ExpireTimeUTC,omitempty"`
	// example:
	//
	// hb-bp1u0639js2h7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// true
	IsDeletionProtection *bool `json:"IsDeletionProtection,omitempty" xml:"IsDeletionProtection,omitempty"`
	// example:
	//
	// true
	IsHa *bool `json:"IsHa,omitempty" xml:"IsHa,omitempty"`
	// example:
	//
	// 2.0
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// example:
	//
	// 100
	MasterDiskSize *int32 `json:"MasterDiskSize,omitempty" xml:"MasterDiskSize,omitempty"`
	// example:
	//
	// cloud_efficiency
	MasterDiskType *string `json:"MasterDiskType,omitempty" xml:"MasterDiskType,omitempty"`
	// example:
	//
	// hbase.sn1.large
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	// example:
	//
	// 2
	MasterNodeCount *int32 `json:"MasterNodeCount,omitempty" xml:"MasterNodeCount,omitempty"`
	// example:
	//
	// 0
	ModuleId *int32 `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// example:
	//
	// 1.0
	ModuleStackVersion *string `json:"ModuleStackVersion,omitempty" xml:"ModuleStackVersion,omitempty"`
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// 2980****2123
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-4f51d54g5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// ACTIVATION
	Status *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeInstancesResponseBodyInstancesInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// example:
	//
	// vpc-bp120k6ixs4eoghz*****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-bp191ipotq****dbqf
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetAutoRenewal(v bool) *DescribeInstancesResponseBodyInstancesInstance {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetBackupStatus(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.BackupStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetClusterId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetClusterName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ClusterName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetClusterType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ClusterType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetColdStorageStatus(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ColdStorageStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCoreDiskCount(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CoreDiskCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCoreDiskSize(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.CoreDiskSize = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCoreDiskType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CoreDiskType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCoreInstanceType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CoreInstanceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCoreNodeCount(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.CoreNodeCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCreatedTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CreatedTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCreatedTimeUTC(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CreatedTimeUTC = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDuration(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.Duration = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetEngine(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.Engine = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetExpireTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetExpireTimeUTC(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ExpireTimeUTC = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetIsDeletionProtection(v bool) *DescribeInstancesResponseBodyInstancesInstance {
	s.IsDeletionProtection = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetIsHa(v bool) *DescribeInstancesResponseBodyInstancesInstance {
	s.IsHa = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMajorVersion(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.MajorVersion = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMasterDiskSize(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.MasterDiskSize = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMasterDiskType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.MasterDiskType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMasterInstanceType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.MasterInstanceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMasterNodeCount(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.MasterNodeCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetModuleId(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.ModuleId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetModuleStackVersion(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ModuleStackVersion = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetNetworkType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetParentId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ParentId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetPayType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.PayType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetRegionId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetResourceGroupId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetStatus(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetTags(v *DescribeInstancesResponseBodyInstancesInstanceTags) *DescribeInstancesResponseBodyInstancesInstance {
	s.Tags = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetVpcId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetVswitchId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetZoneId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ZoneId = &v
	return s
}

type DescribeInstancesResponseBodyInstancesInstanceTags struct {
	Tag []*DescribeInstancesResponseBodyInstancesInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTags) SetTag(v []*DescribeInstancesResponseBodyInstancesInstanceTagsTag) *DescribeInstancesResponseBodyInstancesInstanceTags {
	s.Tag = v
	return s
}

type DescribeInstancesResponseBodyInstancesInstanceTagsTag struct {
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTag) SetKey(v string) *DescribeInstancesResponseBodyInstancesInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTag) SetValue(v string) *DescribeInstancesResponseBodyInstancesInstanceTagsTag {
	s.Value = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetHeaders(v map[string]*string) *DescribeInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesResponse) SetStatusCode(v int32) *DescribeInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeIpWhitelistRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
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
	Groups *DescribeIpWhitelistResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	// example:
	//
	// AFAA617B-3268-5883-982B-DB8EC8CC1F1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIpWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponseBody) SetGroups(v *DescribeIpWhitelistResponseBodyGroups) *DescribeIpWhitelistResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeIpWhitelistResponseBody) SetRequestId(v string) *DescribeIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type DescribeIpWhitelistResponseBodyGroups struct {
	Group []*DescribeIpWhitelistResponseBodyGroupsGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
}

func (s DescribeIpWhitelistResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponseBodyGroups) SetGroup(v []*DescribeIpWhitelistResponseBodyGroupsGroup) *DescribeIpWhitelistResponseBodyGroups {
	s.Group = v
	return s
}

type DescribeIpWhitelistResponseBodyGroupsGroup struct {
	// example:
	//
	// default
	GroupName *string                                           `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IpList    *DescribeIpWhitelistResponseBodyGroupsGroupIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Struct"`
	// example:
	//
	// 4
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
}

func (s DescribeIpWhitelistResponseBodyGroupsGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistResponseBodyGroupsGroup) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroup) SetGroupName(v string) *DescribeIpWhitelistResponseBodyGroupsGroup {
	s.GroupName = &v
	return s
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroup) SetIpList(v *DescribeIpWhitelistResponseBodyGroupsGroupIpList) *DescribeIpWhitelistResponseBodyGroupsGroup {
	s.IpList = v
	return s
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroup) SetIpVersion(v int32) *DescribeIpWhitelistResponseBodyGroupsGroup {
	s.IpVersion = &v
	return s
}

type DescribeIpWhitelistResponseBodyGroupsGroupIpList struct {
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
}

func (s DescribeIpWhitelistResponseBodyGroupsGroupIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpWhitelistResponseBodyGroupsGroupIpList) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistResponseBodyGroupsGroupIpList) SetIp(v []*string) *DescribeIpWhitelistResponseBodyGroupsGroupIpList {
	s.Ip = v
	return s
}

type DescribeIpWhitelistResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeIpWhitelistResponse) SetStatusCode(v int32) *DescribeIpWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpWhitelistResponse) SetBody(v *DescribeIpWhitelistResponseBody) *DescribeIpWhitelistResponse {
	s.Body = v
	return s
}

type DescribeMultiZoneAvailableRegionsRequest struct {
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeMultiZoneAvailableRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsRequest) SetAcceptLanguage(v string) *DescribeMultiZoneAvailableRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeMultiZoneAvailableRegionsResponseBody struct {
	Regions *DescribeMultiZoneAvailableRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// example:
	//
	// F03BB273-45EE-4B6C-A329-A6E6A8D15856
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBody) SetRegions(v *DescribeMultiZoneAvailableRegionsResponseBodyRegions) *DescribeMultiZoneAvailableRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBody) SetRequestId(v string) *DescribeMultiZoneAvailableRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMultiZoneAvailableRegionsResponseBodyRegions struct {
	Region []*DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegions) SetRegion(v []*DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) *DescribeMultiZoneAvailableRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion struct {
	AvailableCombines *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines `json:"AvailableCombines,omitempty" xml:"AvailableCombines,omitempty" type:"Struct"`
	LocalName         *string                                                                      `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// example:
	//
	// hbase.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) SetAvailableCombines(v *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion {
	s.AvailableCombines = v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines struct {
	AvailableCombine []*DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine `json:"AvailableCombine,omitempty" xml:"AvailableCombine,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines) SetAvailableCombine(v []*DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines {
	s.AvailableCombine = v
	return s
}

type DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine struct {
	// example:
	//
	// cn-shenzhen-****-aliyun
	Id    *string                                                                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	Zones *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) SetId(v string) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine {
	s.Id = &v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) SetZones(v *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine {
	s.Zones = v
	return s
}

type DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones struct {
	Zone []*string `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones) SetZone(v []*string) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones {
	s.Zone = v
	return s
}

type DescribeMultiZoneAvailableRegionsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMultiZoneAvailableRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMultiZoneAvailableRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponse) SetHeaders(v map[string]*string) *DescribeMultiZoneAvailableRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponse) SetStatusCode(v int32) *DescribeMultiZoneAvailableRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponse) SetBody(v *DescribeMultiZoneAvailableRegionsResponseBody) *DescribeMultiZoneAvailableRegionsResponse {
	s.Body = v
	return s
}

type DescribeMultiZoneAvailableResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-hangzhou-bef-aliyun
	ZoneCombination *string `json:"ZoneCombination,omitempty" xml:"ZoneCombination,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceRequest) SetChargeType(v string) *DescribeMultiZoneAvailableResourceRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceRequest) SetRegionId(v string) *DescribeMultiZoneAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceRequest) SetZoneCombination(v string) *DescribeMultiZoneAvailableResourceRequest {
	s.ZoneCombination = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBody struct {
	AvailableZones *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Struct"`
	// example:
	//
	// B2EEBBA9-C627-4415-81A0-B77BC54F1D52
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBody) SetAvailableZones(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones) *DescribeMultiZoneAvailableResourceResponseBody {
	s.AvailableZones = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBody) SetRequestId(v string) *DescribeMultiZoneAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZones struct {
	AvailableZone []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones) SetAvailableZone(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones {
	s.AvailableZone = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone struct {
	MasterResources *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources `json:"MasterResources,omitempty" xml:"MasterResources,omitempty" type:"Struct"`
	// example:
	//
	// cn-hangzhou
	RegionId         *string                                                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SupportedEngines *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines `json:"SupportedEngines,omitempty" xml:"SupportedEngines,omitempty" type:"Struct"`
	// example:
	//
	// cn-hangzhou-bef-aliyun
	ZoneCombination *string `json:"ZoneCombination,omitempty" xml:"ZoneCombination,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) SetMasterResources(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.MasterResources = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) SetRegionId(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.RegionId = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) SetSupportedEngines(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.SupportedEngines = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) SetZoneCombination(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.ZoneCombination = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources struct {
	MasterResource []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource `json:"MasterResource,omitempty" xml:"MasterResource,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) SetMasterResource(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources {
	s.MasterResource = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource struct {
	// example:
	//
	// hbase.sn2.large
	InstanceType       *string                                                                                                                   `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypeDetail *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail `json:"InstanceTypeDetail,omitempty" xml:"InstanceTypeDetail,omitempty" type:"Struct"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) SetInstanceType(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource {
	s.InstanceType = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) SetInstanceTypeDetail(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource {
	s.InstanceTypeDetail = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail struct {
	// example:
	//
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 16
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) SetCpu(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail {
	s.Cpu = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) SetMem(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail {
	s.Mem = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines struct {
	SupportedEngine []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine `json:"SupportedEngine,omitempty" xml:"SupportedEngine,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) SetSupportedEngine(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines {
	s.SupportedEngine = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine struct {
	// example:
	//
	// hbaseue
	Engine                  *string                                                                                                                          `json:"Engine,omitempty" xml:"Engine,omitempty"`
	SupportedEngineVersions *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions `json:"SupportedEngineVersions,omitempty" xml:"SupportedEngineVersions,omitempty" type:"Struct"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) SetEngine(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine {
	s.Engine = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) SetSupportedEngineVersions(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine {
	s.SupportedEngineVersions = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions struct {
	SupportedEngineVersion []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion `json:"SupportedEngineVersion,omitempty" xml:"SupportedEngineVersion,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) SetSupportedEngineVersion(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions {
	s.SupportedEngineVersion = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion struct {
	SupportedCategories *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories `json:"SupportedCategories,omitempty" xml:"SupportedCategories,omitempty" type:"Struct"`
	// example:
	//
	// 2.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) SetSupportedCategories(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion {
	s.SupportedCategories = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) SetVersion(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion {
	s.Version = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories struct {
	SupportedCategories []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories `json:"SupportedCategories,omitempty" xml:"SupportedCategories,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) SetSupportedCategories(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories {
	s.SupportedCategories = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories struct {
	// example:
	//
	// cluster
	Category              *string                                                                                                                                                                                                           `json:"Category,omitempty" xml:"Category,omitempty"`
	SupportedStorageTypes *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes `json:"SupportedStorageTypes,omitempty" xml:"SupportedStorageTypes,omitempty" type:"Struct"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) SetCategory(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories {
	s.Category = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) SetSupportedStorageTypes(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories {
	s.SupportedStorageTypes = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes struct {
	SupportedStorageType []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType `json:"SupportedStorageType,omitempty" xml:"SupportedStorageType,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) SetSupportedStorageType(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes {
	s.SupportedStorageType = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType struct {
	CoreResources *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources `json:"CoreResources,omitempty" xml:"CoreResources,omitempty" type:"Struct"`
	// example:
	//
	// cloud_efficiency
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) SetCoreResources(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType {
	s.CoreResources = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) SetStorageType(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType {
	s.StorageType = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources struct {
	CoreResource []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource `json:"CoreResource,omitempty" xml:"CoreResource,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) SetCoreResource(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources {
	s.CoreResource = v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource struct {
	DBInstanceStorageRange *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange `json:"DBInstanceStorageRange,omitempty" xml:"DBInstanceStorageRange,omitempty" type:"Struct"`
	// example:
	//
	// hbase.sn2.2xlarge
	InstanceType       *string                                                                                                                                                                                                                                                                          `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypeDetail *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail `json:"InstanceTypeDetail,omitempty" xml:"InstanceTypeDetail,omitempty" type:"Struct"`
	// example:
	//
	// 30
	MaxCoreCount *int32 `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetDBInstanceStorageRange(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.DBInstanceStorageRange = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetInstanceType(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.InstanceType = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetInstanceTypeDetail(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.InstanceTypeDetail = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetMaxCoreCount(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.MaxCoreCount = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange struct {
	// example:
	//
	// 64000
	MaxSize *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// example:
	//
	// 400
	MinSize *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	// example:
	//
	// 40
	StepSize *int32 `json:"StepSize,omitempty" xml:"StepSize,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) SetMaxSize(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	s.MaxSize = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) SetMinSize(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	s.MinSize = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) SetStepSize(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	s.StepSize = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail struct {
	// example:
	//
	// 32
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 8
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) SetCpu(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail {
	s.Cpu = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) SetMem(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail {
	s.Mem = &v
	return s
}

type DescribeMultiZoneAvailableResourceResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMultiZoneAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponse) SetHeaders(v map[string]*string) *DescribeMultiZoneAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponse) SetStatusCode(v int32) *DescribeMultiZoneAvailableResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponse) SetBody(v *DescribeMultiZoneAvailableResourceResponseBody) *DescribeMultiZoneAvailableResourceResponse {
	s.Body = v
	return s
}

type DescribeMultiZoneClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-t4nn71xa0yn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeMultiZoneClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterRequest) SetClusterId(v string) *DescribeMultiZoneClusterRequest {
	s.ClusterId = &v
	return s
}

type DescribeMultiZoneClusterResponseBody struct {
	// example:
	//
	// vsw-t4nax9mp3wk0czn****
	ArbiterVSwitchIds *string `json:"ArbiterVSwitchIds,omitempty" xml:"ArbiterVSwitchIds,omitempty"`
	// example:
	//
	// ap-southeast-1c
	ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// example:
	//
	// ld-t4nn71xa0yn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// mz_test
	ClusterName     *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ColdStorageSize *int32  `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	// example:
	//
	// 4
	CoreDiskCount *string `json:"CoreDiskCount,omitempty" xml:"CoreDiskCount,omitempty"`
	// example:
	//
	// 100
	CoreDiskSize *int32 `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	// example:
	//
	// cloud_efficiency
	CoreDiskType *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	// example:
	//
	// hbase.sn1.large
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// example:
	//
	// 6
	CoreNodeCount *int32 `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	// example:
	//
	// 2020-10-15T18:04:52
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 2020-10-15T10:04:52Z
	CreatedTimeUTC *string `json:"CreatedTimeUTC,omitempty" xml:"CreatedTimeUTC,omitempty"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2a****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// example:
	//
	// CloudDisk
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// example:
	//
	// hbaseue
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2020-11-16T08:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2020-11-16T00:00:00Z
	ExpireTimeUTC *string `json:"ExpireTimeUTC,omitempty" xml:"ExpireTimeUTC,omitempty"`
	// example:
	//
	// ld-t4nn71xa0yn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// mz_test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// false
	IsDeletionProtection *bool `json:"IsDeletionProtection,omitempty" xml:"IsDeletionProtection,omitempty"`
	// example:
	//
	// 4
	LogDiskCount *string `json:"LogDiskCount,omitempty" xml:"LogDiskCount,omitempty"`
	// example:
	//
	// 100
	LogDiskSize *int32 `json:"LogDiskSize,omitempty" xml:"LogDiskSize,omitempty"`
	// example:
	//
	// cloud_efficiency
	LogDiskType *string `json:"LogDiskType,omitempty" xml:"LogDiskType,omitempty"`
	// example:
	//
	// hbase.sn1.large
	LogInstanceType *string `json:"LogInstanceType,omitempty" xml:"LogInstanceType,omitempty"`
	// example:
	//
	// 4
	LogNodeCount *int32 `json:"LogNodeCount,omitempty" xml:"LogNodeCount,omitempty"`
	// example:
	//
	// 06:00:00
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// example:
	//
	// 02:00:00
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// example:
	//
	// 2.0
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// example:
	//
	// 50
	MasterDiskSize *int32 `json:"MasterDiskSize,omitempty" xml:"MasterDiskSize,omitempty"`
	// example:
	//
	// cloud_efficiency
	MasterDiskType *string `json:"MasterDiskType,omitempty" xml:"MasterDiskType,omitempty"`
	// example:
	//
	// hbase.sn1.large
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	// example:
	//
	// 2
	MasterNodeCount *int32 `json:"MasterNodeCount,omitempty" xml:"MasterNodeCount,omitempty"`
	// example:
	//
	// 0
	ModuleId *int32 `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// example:
	//
	// 2.0
	ModuleStackVersion *string `json:"ModuleStackVersion,omitempty" xml:"ModuleStackVersion,omitempty"`
	// example:
	//
	// ap-southeast-1-abc-aliyun
	MultiZoneCombination    *string                                                      `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
	MultiZoneInstanceModels *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels `json:"MultiZoneInstanceModels,omitempty" xml:"MultiZoneInstanceModels,omitempty" type:"Struct"`
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// ld-fls1gf31y5s35****
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// vsw-t4n3s1zd2gtidg****
	PrimaryVSwitchIds *string `json:"PrimaryVSwitchIds,omitempty" xml:"PrimaryVSwitchIds,omitempty"`
	// example:
	//
	// ap-southeast-1a
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// example:
	//
	// ap-southeast-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// A02C0E6D-3A47-4FA0-BA7E-60793CE256DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-lk51f5fer315e****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// vsw-t4nvvk7xur3rdi****
	StandbyVSwitchIds *string `json:"StandbyVSwitchIds,omitempty" xml:"StandbyVSwitchIds,omitempty"`
	// example:
	//
	// ap-southeast-1b
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// example:
	//
	// ACTIVATION
	Status       *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags         *DescribeMultiZoneClusterResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TaskProgress *string                                   `json:"TaskProgress,omitempty" xml:"TaskProgress,omitempty"`
	TaskStatus   *string                                   `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// vpc-t4nx81tmlixcq5****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeMultiZoneClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponseBody) SetArbiterVSwitchIds(v string) *DescribeMultiZoneClusterResponseBody {
	s.ArbiterVSwitchIds = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetArbiterZoneId(v string) *DescribeMultiZoneClusterResponseBody {
	s.ArbiterZoneId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetAutoRenewal(v bool) *DescribeMultiZoneClusterResponseBody {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetClusterId(v string) *DescribeMultiZoneClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetClusterName(v string) *DescribeMultiZoneClusterResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetColdStorageSize(v int32) *DescribeMultiZoneClusterResponseBody {
	s.ColdStorageSize = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCoreDiskCount(v string) *DescribeMultiZoneClusterResponseBody {
	s.CoreDiskCount = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCoreDiskSize(v int32) *DescribeMultiZoneClusterResponseBody {
	s.CoreDiskSize = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCoreDiskType(v string) *DescribeMultiZoneClusterResponseBody {
	s.CoreDiskType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCoreInstanceType(v string) *DescribeMultiZoneClusterResponseBody {
	s.CoreInstanceType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCoreNodeCount(v int32) *DescribeMultiZoneClusterResponseBody {
	s.CoreNodeCount = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCreatedTime(v string) *DescribeMultiZoneClusterResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCreatedTimeUTC(v string) *DescribeMultiZoneClusterResponseBody {
	s.CreatedTimeUTC = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetDuration(v int32) *DescribeMultiZoneClusterResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetEncryptionKey(v string) *DescribeMultiZoneClusterResponseBody {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetEncryptionType(v string) *DescribeMultiZoneClusterResponseBody {
	s.EncryptionType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetEngine(v string) *DescribeMultiZoneClusterResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetExpireTime(v string) *DescribeMultiZoneClusterResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetExpireTimeUTC(v string) *DescribeMultiZoneClusterResponseBody {
	s.ExpireTimeUTC = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetInstanceId(v string) *DescribeMultiZoneClusterResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetInstanceName(v string) *DescribeMultiZoneClusterResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetIsDeletionProtection(v bool) *DescribeMultiZoneClusterResponseBody {
	s.IsDeletionProtection = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetLogDiskCount(v string) *DescribeMultiZoneClusterResponseBody {
	s.LogDiskCount = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetLogDiskSize(v int32) *DescribeMultiZoneClusterResponseBody {
	s.LogDiskSize = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetLogDiskType(v string) *DescribeMultiZoneClusterResponseBody {
	s.LogDiskType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetLogInstanceType(v string) *DescribeMultiZoneClusterResponseBody {
	s.LogInstanceType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetLogNodeCount(v int32) *DescribeMultiZoneClusterResponseBody {
	s.LogNodeCount = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMaintainEndTime(v string) *DescribeMultiZoneClusterResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMaintainStartTime(v string) *DescribeMultiZoneClusterResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMajorVersion(v string) *DescribeMultiZoneClusterResponseBody {
	s.MajorVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMasterDiskSize(v int32) *DescribeMultiZoneClusterResponseBody {
	s.MasterDiskSize = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMasterDiskType(v string) *DescribeMultiZoneClusterResponseBody {
	s.MasterDiskType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMasterInstanceType(v string) *DescribeMultiZoneClusterResponseBody {
	s.MasterInstanceType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMasterNodeCount(v int32) *DescribeMultiZoneClusterResponseBody {
	s.MasterNodeCount = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetModuleId(v int32) *DescribeMultiZoneClusterResponseBody {
	s.ModuleId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetModuleStackVersion(v string) *DescribeMultiZoneClusterResponseBody {
	s.ModuleStackVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMultiZoneCombination(v string) *DescribeMultiZoneClusterResponseBody {
	s.MultiZoneCombination = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMultiZoneInstanceModels(v *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels) *DescribeMultiZoneClusterResponseBody {
	s.MultiZoneInstanceModels = v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetNetworkType(v string) *DescribeMultiZoneClusterResponseBody {
	s.NetworkType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetParentId(v string) *DescribeMultiZoneClusterResponseBody {
	s.ParentId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetPayType(v string) *DescribeMultiZoneClusterResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetPrimaryVSwitchIds(v string) *DescribeMultiZoneClusterResponseBody {
	s.PrimaryVSwitchIds = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetPrimaryZoneId(v string) *DescribeMultiZoneClusterResponseBody {
	s.PrimaryZoneId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetRegionId(v string) *DescribeMultiZoneClusterResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetRequestId(v string) *DescribeMultiZoneClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetResourceGroupId(v string) *DescribeMultiZoneClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetStandbyVSwitchIds(v string) *DescribeMultiZoneClusterResponseBody {
	s.StandbyVSwitchIds = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetStandbyZoneId(v string) *DescribeMultiZoneClusterResponseBody {
	s.StandbyZoneId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetStatus(v string) *DescribeMultiZoneClusterResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetTags(v *DescribeMultiZoneClusterResponseBodyTags) *DescribeMultiZoneClusterResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetTaskProgress(v string) *DescribeMultiZoneClusterResponseBody {
	s.TaskProgress = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetTaskStatus(v string) *DescribeMultiZoneClusterResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetVpcId(v string) *DescribeMultiZoneClusterResponseBody {
	s.VpcId = &v
	return s
}

type DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels struct {
	MultiZoneInstanceModel []*DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel `json:"MultiZoneInstanceModel,omitempty" xml:"MultiZoneInstanceModel,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels) SetMultiZoneInstanceModel(v []*DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels {
	s.MultiZoneInstanceModel = v
	return s
}

type DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel struct {
	HdfsMinorVersion *string `json:"HdfsMinorVersion,omitempty" xml:"HdfsMinorVersion,omitempty"`
	// example:
	//
	// ld-t4nn71xa0yn****-az-a
	InsName             *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	IsHdfsLatestVersion *string `json:"IsHdfsLatestVersion,omitempty" xml:"IsHdfsLatestVersion,omitempty"`
	// example:
	//
	// true
	IsLatestVersion        *bool   `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	LatestHdfsMinorVersion *string `json:"LatestHdfsMinorVersion,omitempty" xml:"LatestHdfsMinorVersion,omitempty"`
	LatestMinorVersion     *string `json:"LatestMinorVersion,omitempty" xml:"LatestMinorVersion,omitempty"`
	// example:
	//
	// 2.1.24
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// example:
	//
	// primary
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetHdfsMinorVersion(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.HdfsMinorVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetInsName(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.InsName = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetIsHdfsLatestVersion(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.IsHdfsLatestVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetIsLatestVersion(v bool) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetLatestHdfsMinorVersion(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.LatestHdfsMinorVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetLatestMinorVersion(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.LatestMinorVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetMinorVersion(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.MinorVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetRole(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.Role = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetStatus(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.Status = &v
	return s
}

type DescribeMultiZoneClusterResponseBodyTags struct {
	Tag []*DescribeMultiZoneClusterResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneClusterResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneClusterResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponseBodyTags) SetTag(v []*DescribeMultiZoneClusterResponseBodyTagsTag) *DescribeMultiZoneClusterResponseBodyTags {
	s.Tag = v
	return s
}

type DescribeMultiZoneClusterResponseBodyTagsTag struct {
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMultiZoneClusterResponseBodyTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneClusterResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponseBodyTagsTag) SetKey(v string) *DescribeMultiZoneClusterResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyTagsTag) SetValue(v string) *DescribeMultiZoneClusterResponseBodyTagsTag {
	s.Value = &v
	return s
}

type DescribeMultiZoneClusterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMultiZoneClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMultiZoneClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiZoneClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponse) SetHeaders(v map[string]*string) *DescribeMultiZoneClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiZoneClusterResponse) SetStatusCode(v int32) *DescribeMultiZoneClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMultiZoneClusterResponse) SetBody(v *DescribeMultiZoneClusterResponseBody) *DescribeMultiZoneClusterResponse {
	s.Body = v
	return s
}

type DescribeRecoverableTimeRangeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeRecoverableTimeRangeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecoverableTimeRangeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableTimeRangeRequest) SetClusterId(v string) *DescribeRecoverableTimeRangeRequest {
	s.ClusterId = &v
	return s
}

type DescribeRecoverableTimeRangeResponseBody struct {
	// example:
	//
	// A1A51D18-96DC-465C-9F1B-47180CA22524
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2020-10-26T18:02:03Z
	TimeBegin *string `json:"TimeBegin,omitempty" xml:"TimeBegin,omitempty"`
	// example:
	//
	// 2020-11-05T01:20:31Z
	TimeEnd *string `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty"`
}

func (s DescribeRecoverableTimeRangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecoverableTimeRangeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableTimeRangeResponseBody) SetRequestId(v string) *DescribeRecoverableTimeRangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecoverableTimeRangeResponseBody) SetTimeBegin(v string) *DescribeRecoverableTimeRangeResponseBody {
	s.TimeBegin = &v
	return s
}

func (s *DescribeRecoverableTimeRangeResponseBody) SetTimeEnd(v string) *DescribeRecoverableTimeRangeResponseBody {
	s.TimeEnd = &v
	return s
}

type DescribeRecoverableTimeRangeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecoverableTimeRangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecoverableTimeRangeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecoverableTimeRangeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableTimeRangeResponse) SetHeaders(v map[string]*string) *DescribeRecoverableTimeRangeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecoverableTimeRangeResponse) SetStatusCode(v int32) *DescribeRecoverableTimeRangeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecoverableTimeRangeResponse) SetBody(v *DescribeRecoverableTimeRangeResponseBody) *DescribeRecoverableTimeRangeResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// hbase
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
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

func (s *DescribeRegionsRequest) SetEngine(v string) *DescribeRegionsRequest {
	s.Engine = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// example:
	//
	// 14D3924C-4FD8-4EE9-9B34-DA949D104F42
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
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// example:
	//
	// hbase.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Zones    *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
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
	// example:
	//
	// cn-hangzhou-b
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

type DescribeRestoreFullDetailsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020110514xxxx
	RestoreRecordId *string `json:"RestoreRecordId,omitempty" xml:"RestoreRecordId,omitempty"`
}

func (s DescribeRestoreFullDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreFullDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsRequest) SetClusterId(v string) *DescribeRestoreFullDetailsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreFullDetailsRequest) SetPageNumber(v int32) *DescribeRestoreFullDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreFullDetailsRequest) SetPageSize(v int32) *DescribeRestoreFullDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreFullDetailsRequest) SetRestoreRecordId(v string) *DescribeRestoreFullDetailsRequest {
	s.RestoreRecordId = &v
	return s
}

type DescribeRestoreFullDetailsResponseBody struct {
	// example:
	//
	// CFE525CF-C691-4140-A981-D004DAA7A840
	RequestId   *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreFull *DescribeRestoreFullDetailsResponseBodyRestoreFull `json:"RestoreFull,omitempty" xml:"RestoreFull,omitempty" type:"Struct"`
}

func (s DescribeRestoreFullDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreFullDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsResponseBody) SetRequestId(v string) *DescribeRestoreFullDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBody) SetRestoreFull(v *DescribeRestoreFullDetailsResponseBodyRestoreFull) *DescribeRestoreFullDetailsResponseBody {
	s.RestoreFull = v
	return s
}

type DescribeRestoreFullDetailsResponseBodyRestoreFull struct {
	// example:
	//
	// 1.2 kB
	DataSize *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// example:
	//
	// 0
	Fail *int32 `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize           *int32                                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RestoreFullDetails *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails `json:"RestoreFullDetails,omitempty" xml:"RestoreFullDetails,omitempty" type:"Struct"`
	// example:
	//
	// 0.00 MB/s
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// example:
	//
	// 1
	Succeed *int32 `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFull) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFull) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetDataSize(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.DataSize = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetFail(v int32) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.Fail = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetPageNumber(v int32) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetPageSize(v int32) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetRestoreFullDetails(v *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.RestoreFullDetails = v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetSpeed(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.Speed = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetSucceed(v int32) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.Succeed = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFull) SetTotal(v int64) *DescribeRestoreFullDetailsResponseBodyRestoreFull {
	s.Total = &v
	return s
}

type DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails struct {
	RestoreFullDetail []*DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail `json:"RestoreFullDetail,omitempty" xml:"RestoreFullDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails) SetRestoreFullDetail(v []*DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetails {
	s.RestoreFullDetail = v
	return s
}

type DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail struct {
	// example:
	//
	// 1.2 kB
	DataSize *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:51Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 14/14
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// 0.00 MB/s
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:45Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCEEDED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// default:test1
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetDataSize(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.DataSize = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetEndTime(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetMessage(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Message = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetProcess(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Process = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetSpeed(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Speed = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetStartTime(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetState(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetTable(v string) *DescribeRestoreFullDetailsResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Table = &v
	return s
}

type DescribeRestoreFullDetailsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreFullDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreFullDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreFullDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsResponse) SetHeaders(v map[string]*string) *DescribeRestoreFullDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreFullDetailsResponse) SetStatusCode(v int32) *DescribeRestoreFullDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreFullDetailsResponse) SetBody(v *DescribeRestoreFullDetailsResponseBody) *DescribeRestoreFullDetailsResponse {
	s.Body = v
	return s
}

type DescribeRestoreIncrDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eyf188hw481xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020110514xxxx
	RestoreRecordId *string `json:"RestoreRecordId,omitempty" xml:"RestoreRecordId,omitempty"`
}

func (s DescribeRestoreIncrDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreIncrDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreIncrDetailRequest) SetClusterId(v string) *DescribeRestoreIncrDetailRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreIncrDetailRequest) SetRestoreRecordId(v string) *DescribeRestoreIncrDetailRequest {
	s.RestoreRecordId = &v
	return s
}

type DescribeRestoreIncrDetailResponseBody struct {
	// example:
	//
	// D0FE2717-E194-465A-B27B-7373F96E580B
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreIncrDetail *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail `json:"RestoreIncrDetail,omitempty" xml:"RestoreIncrDetail,omitempty" type:"Struct"`
}

func (s DescribeRestoreIncrDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreIncrDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreIncrDetailResponseBody) SetRequestId(v string) *DescribeRestoreIncrDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBody) SetRestoreIncrDetail(v *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) *DescribeRestoreIncrDetailResponseBody {
	s.RestoreIncrDetail = v
	return s
}

type DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail struct {
	// example:
	//
	// 2020-11-05T06:45:44Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 0/0
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// 0 ms
	RestoreDelay *string `json:"RestoreDelay,omitempty" xml:"RestoreDelay,omitempty"`
	// example:
	//
	// 2020-11-02T18:00:00Z
	RestoreStartTs *string `json:"RestoreStartTs,omitempty" xml:"RestoreStartTs,omitempty"`
	// example:
	//
	// \\"\\"
	RestoredTs *string `json:"RestoredTs,omitempty" xml:"RestoredTs,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:44Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCEEDED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetEndTime(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetProcess(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.Process = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetRestoreDelay(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.RestoreDelay = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetRestoreStartTs(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.RestoreStartTs = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetRestoredTs(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.RestoredTs = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetStartTime(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetState(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.State = &v
	return s
}

type DescribeRestoreIncrDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreIncrDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreIncrDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreIncrDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreIncrDetailResponse) SetHeaders(v map[string]*string) *DescribeRestoreIncrDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreIncrDetailResponse) SetStatusCode(v int32) *DescribeRestoreIncrDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponse) SetBody(v *DescribeRestoreIncrDetailResponseBody) *DescribeRestoreIncrDetailResponse {
	s.Body = v
	return s
}

type DescribeRestoreSchemaDetailsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020110514xxxx
	RestoreRecordId *string `json:"RestoreRecordId,omitempty" xml:"RestoreRecordId,omitempty"`
}

func (s DescribeRestoreSchemaDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsRequest) SetClusterId(v string) *DescribeRestoreSchemaDetailsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsRequest) SetPageNumber(v int32) *DescribeRestoreSchemaDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsRequest) SetPageSize(v int32) *DescribeRestoreSchemaDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsRequest) SetRestoreRecordId(v string) *DescribeRestoreSchemaDetailsRequest {
	s.RestoreRecordId = &v
	return s
}

type DescribeRestoreSchemaDetailsResponseBody struct {
	// example:
	//
	// BC682A80-7677-4294-975C-CFEA425381DE
	RequestId     *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreSchema *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema `json:"RestoreSchema,omitempty" xml:"RestoreSchema,omitempty" type:"Struct"`
}

func (s DescribeRestoreSchemaDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsResponseBody) SetRequestId(v string) *DescribeRestoreSchemaDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBody) SetRestoreSchema(v *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) *DescribeRestoreSchemaDetailsResponseBody {
	s.RestoreSchema = v
	return s
}

type DescribeRestoreSchemaDetailsResponseBodyRestoreSchema struct {
	// example:
	//
	// 0
	Fail *int32 `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize             *int32                                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RestoreSchemaDetails *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails `json:"RestoreSchemaDetails,omitempty" xml:"RestoreSchemaDetails,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Succeed *int32 `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetFail(v int32) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.Fail = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetPageNumber(v int32) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetPageSize(v int32) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetRestoreSchemaDetails(v *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.RestoreSchemaDetails = v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetSucceed(v int32) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.Succeed = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema) SetTotal(v int64) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchema {
	s.Total = &v
	return s
}

type DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails struct {
	RestoreSchemaDetail []*DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail `json:"RestoreSchemaDetail,omitempty" xml:"RestoreSchemaDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails) SetRestoreSchemaDetail(v []*DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetails {
	s.RestoreSchemaDetail = v
	return s
}

type DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail struct {
	// example:
	//
	// 2020-11-05T06:45:18Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:14Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCEEDED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// default:test1
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetEndTime(v string) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetMessage(v string) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.Message = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetStartTime(v string) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetState(v string) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetTable(v string) *DescribeRestoreSchemaDetailsResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.Table = &v
	return s
}

type DescribeRestoreSchemaDetailsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreSchemaDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreSchemaDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsResponse) SetHeaders(v map[string]*string) *DescribeRestoreSchemaDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponse) SetStatusCode(v int32) *DescribeRestoreSchemaDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsResponse) SetBody(v *DescribeRestoreSchemaDetailsResponseBody) *DescribeRestoreSchemaDetailsResponse {
	s.Body = v
	return s
}

type DescribeRestoreSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRestoreSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSummaryRequest) SetClusterId(v string) *DescribeRestoreSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreSummaryRequest) SetPageNumber(v int32) *DescribeRestoreSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreSummaryRequest) SetPageSize(v int32) *DescribeRestoreSummaryRequest {
	s.PageSize = &v
	return s
}

type DescribeRestoreSummaryResponseBody struct {
	// example:
	//
	// 0
	HasMoreRestoreRecord *int32 `json:"HasMoreRestoreRecord,omitempty" xml:"HasMoreRestoreRecord,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AE639ED7-F0F3-4A71-911E-CF8EC088816E
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rescords  *DescribeRestoreSummaryResponseBodyRescords `json:"Rescords,omitempty" xml:"Rescords,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRestoreSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSummaryResponseBody) SetHasMoreRestoreRecord(v int32) *DescribeRestoreSummaryResponseBody {
	s.HasMoreRestoreRecord = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) SetPageNumber(v int32) *DescribeRestoreSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) SetPageSize(v int32) *DescribeRestoreSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) SetRequestId(v string) *DescribeRestoreSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) SetRescords(v *DescribeRestoreSummaryResponseBodyRescords) *DescribeRestoreSummaryResponseBody {
	s.Rescords = v
	return s
}

func (s *DescribeRestoreSummaryResponseBody) SetTotal(v int32) *DescribeRestoreSummaryResponseBody {
	s.Total = &v
	return s
}

type DescribeRestoreSummaryResponseBodyRescords struct {
	Rescord []*DescribeRestoreSummaryResponseBodyRescordsRescord `json:"Rescord,omitempty" xml:"Rescord,omitempty" type:"Repeated"`
}

func (s DescribeRestoreSummaryResponseBodyRescords) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSummaryResponseBodyRescords) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSummaryResponseBodyRescords) SetRescord(v []*DescribeRestoreSummaryResponseBodyRescordsRescord) *DescribeRestoreSummaryResponseBodyRescords {
	s.Rescord = v
	return s
}

type DescribeRestoreSummaryResponseBodyRescordsRescord struct {
	// example:
	//
	// 1/1
	BulkLoadProcess *string `json:"BulkLoadProcess,omitempty" xml:"BulkLoadProcess,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:14Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:51Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// 1/1
	HfileRestoreProcess *string `json:"HfileRestoreProcess,omitempty" xml:"HfileRestoreProcess,omitempty"`
	// example:
	//
	// 0/0
	LogProcess *string `json:"LogProcess,omitempty" xml:"LogProcess,omitempty"`
	// example:
	//
	// 20201105144514
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// 1/1
	SchemaProcess *string `json:"SchemaProcess,omitempty" xml:"SchemaProcess,omitempty"`
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRestoreSummaryResponseBodyRescordsRescord) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSummaryResponseBodyRescordsRescord) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetBulkLoadProcess(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.BulkLoadProcess = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetCreateTime(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.CreateTime = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetFinishTime(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.FinishTime = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetHfileRestoreProcess(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.HfileRestoreProcess = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetLogProcess(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.LogProcess = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetRecordId(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.RecordId = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetSchemaProcess(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.SchemaProcess = &v
	return s
}

func (s *DescribeRestoreSummaryResponseBodyRescordsRescord) SetStatus(v string) *DescribeRestoreSummaryResponseBodyRescordsRescord {
	s.Status = &v
	return s
}

type DescribeRestoreSummaryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSummaryResponse) SetHeaders(v map[string]*string) *DescribeRestoreSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreSummaryResponse) SetStatusCode(v int32) *DescribeRestoreSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreSummaryResponse) SetBody(v *DescribeRestoreSummaryResponseBody) *DescribeRestoreSummaryResponse {
	s.Body = v
	return s
}

type DescribeRestoreTablesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020110514xxxx
	RestoreRecordId *string `json:"RestoreRecordId,omitempty" xml:"RestoreRecordId,omitempty"`
}

func (s DescribeRestoreTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesRequest) SetClusterId(v string) *DescribeRestoreTablesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreTablesRequest) SetRestoreRecordId(v string) *DescribeRestoreTablesRequest {
	s.RestoreRecordId = &v
	return s
}

type DescribeRestoreTablesResponseBody struct {
	// example:
	//
	// 18D9CC47-D913-48BF-AB6B-4FA9B28FBDB1
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreFull       *DescribeRestoreTablesResponseBodyRestoreFull       `json:"RestoreFull,omitempty" xml:"RestoreFull,omitempty" type:"Struct"`
	RestoreIncrDetail *DescribeRestoreTablesResponseBodyRestoreIncrDetail `json:"RestoreIncrDetail,omitempty" xml:"RestoreIncrDetail,omitempty" type:"Struct"`
	RestoreSchema     *DescribeRestoreTablesResponseBodyRestoreSchema     `json:"RestoreSchema,omitempty" xml:"RestoreSchema,omitempty" type:"Struct"`
	RestoreSummary    *DescribeRestoreTablesResponseBodyRestoreSummary    `json:"RestoreSummary,omitempty" xml:"RestoreSummary,omitempty" type:"Struct"`
	Tables            *DescribeRestoreTablesResponseBodyTables            `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
}

func (s DescribeRestoreTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBody) SetRequestId(v string) *DescribeRestoreTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreTablesResponseBody) SetRestoreFull(v *DescribeRestoreTablesResponseBodyRestoreFull) *DescribeRestoreTablesResponseBody {
	s.RestoreFull = v
	return s
}

func (s *DescribeRestoreTablesResponseBody) SetRestoreIncrDetail(v *DescribeRestoreTablesResponseBodyRestoreIncrDetail) *DescribeRestoreTablesResponseBody {
	s.RestoreIncrDetail = v
	return s
}

func (s *DescribeRestoreTablesResponseBody) SetRestoreSchema(v *DescribeRestoreTablesResponseBodyRestoreSchema) *DescribeRestoreTablesResponseBody {
	s.RestoreSchema = v
	return s
}

func (s *DescribeRestoreTablesResponseBody) SetRestoreSummary(v *DescribeRestoreTablesResponseBodyRestoreSummary) *DescribeRestoreTablesResponseBody {
	s.RestoreSummary = v
	return s
}

func (s *DescribeRestoreTablesResponseBody) SetTables(v *DescribeRestoreTablesResponseBodyTables) *DescribeRestoreTablesResponseBody {
	s.Tables = v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreFull struct {
	// example:
	//
	// 1.2 kB
	DataSize *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// example:
	//
	// 0
	Fail *int32 `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize           *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RestoreFullDetails *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails `json:"RestoreFullDetails,omitempty" xml:"RestoreFullDetails,omitempty" type:"Struct"`
	// example:
	//
	// 0.00 MB/s
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// example:
	//
	// 1
	Succeed *int32 `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRestoreTablesResponseBodyRestoreFull) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreFull) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetDataSize(v string) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.DataSize = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetFail(v int32) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.Fail = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetPageNumber(v int32) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetPageSize(v int32) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetRestoreFullDetails(v *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.RestoreFullDetails = v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetSpeed(v string) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.Speed = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetSucceed(v int32) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.Succeed = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFull) SetTotal(v int64) *DescribeRestoreTablesResponseBodyRestoreFull {
	s.Total = &v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails struct {
	RestoreFullDetail []*DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail `json:"RestoreFullDetail,omitempty" xml:"RestoreFullDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails) SetRestoreFullDetail(v []*DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetails {
	s.RestoreFullDetail = v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail struct {
	// example:
	//
	// 1.2 kB
	DataSize *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:51Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// “”
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 14/14
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// 0.00 MB/s
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:45Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCEEDED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// default:test1
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetDataSize(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.DataSize = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetEndTime(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetMessage(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Message = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetProcess(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Process = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetSpeed(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Speed = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetStartTime(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetState(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail) SetTable(v string) *DescribeRestoreTablesResponseBodyRestoreFullRestoreFullDetailsRestoreFullDetail {
	s.Table = &v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreIncrDetail struct {
	// example:
	//
	// 2020-11-05T06:45:44Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 0/0
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// 0 ms
	RestoreDelay *string `json:"RestoreDelay,omitempty" xml:"RestoreDelay,omitempty"`
	// example:
	//
	// 2020-11-02T18:00:00Z
	RestoreStartTs *string `json:"RestoreStartTs,omitempty" xml:"RestoreStartTs,omitempty"`
	// example:
	//
	// “”
	RestoredTs *string `json:"RestoredTs,omitempty" xml:"RestoredTs,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:44Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCEEDED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeRestoreTablesResponseBodyRestoreIncrDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreIncrDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetEndTime(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetProcess(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.Process = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetRestoreDelay(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.RestoreDelay = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetRestoreStartTs(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.RestoreStartTs = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetRestoredTs(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.RestoredTs = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetStartTime(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreIncrDetail) SetState(v string) *DescribeRestoreTablesResponseBodyRestoreIncrDetail {
	s.State = &v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreSchema struct {
	// example:
	//
	// 0
	Fail *int32 `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize             *int32                                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RestoreSchemaDetails *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails `json:"RestoreSchemaDetails,omitempty" xml:"RestoreSchemaDetails,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Succeed *int32 `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRestoreTablesResponseBodyRestoreSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreSchema) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetFail(v int32) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.Fail = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetPageNumber(v int32) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetPageSize(v int32) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetRestoreSchemaDetails(v *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.RestoreSchemaDetails = v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetSucceed(v int32) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.Succeed = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchema) SetTotal(v int64) *DescribeRestoreTablesResponseBodyRestoreSchema {
	s.Total = &v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails struct {
	RestoreSchemaDetail []*DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail `json:"RestoreSchemaDetail,omitempty" xml:"RestoreSchemaDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails) SetRestoreSchemaDetail(v []*DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetails {
	s.RestoreSchemaDetail = v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail struct {
	// example:
	//
	// 2020-11-05T06:45:18Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:14Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCEEDED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// default:test1
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetEndTime(v string) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetMessage(v string) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.Message = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetStartTime(v string) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetState(v string) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail) SetTable(v string) *DescribeRestoreTablesResponseBodyRestoreSchemaRestoreSchemaDetailsRestoreSchemaDetail {
	s.Table = &v
	return s
}

type DescribeRestoreTablesResponseBodyRestoreSummary struct {
	// example:
	//
	// 2020-11-05T06:45:51Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 20201105144514
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// 2020-11-04T05:15:00Z
	RestoreToDate *string `json:"RestoreToDate,omitempty" xml:"RestoreToDate,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:14Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCEEDED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// ld-m5e2t34kr54wgxxxx
	TargetCluster *string `json:"TargetCluster,omitempty" xml:"TargetCluster,omitempty"`
}

func (s DescribeRestoreTablesResponseBodyRestoreSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyRestoreSummary) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetEndTime(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetRecordId(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.RecordId = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetRestoreToDate(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.RestoreToDate = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetStartTime(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetState(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.State = &v
	return s
}

func (s *DescribeRestoreTablesResponseBodyRestoreSummary) SetTargetCluster(v string) *DescribeRestoreTablesResponseBodyRestoreSummary {
	s.TargetCluster = &v
	return s
}

type DescribeRestoreTablesResponseBodyTables struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeRestoreTablesResponseBodyTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponseBodyTables) SetTable(v []*string) *DescribeRestoreTablesResponseBodyTables {
	s.Table = v
	return s
}

type DescribeRestoreTablesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesResponse) SetHeaders(v map[string]*string) *DescribeRestoreTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreTablesResponse) SetStatusCode(v int32) *DescribeRestoreTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreTablesResponse) SetBody(v *DescribeRestoreTablesResponseBody) *DescribeRestoreTablesResponse {
	s.Body = v
	return s
}

type DescribeSecurityGroupsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp161ax8i03c4uq**
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
	// example:
	//
	// 50373857-C47B-4B64-9332-D0B5280B59EA
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeSecurityGroupsResponse) SetStatusCode(v int32) *DescribeSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityGroupsResponse) SetBody(v *DescribeSecurityGroupsResponseBody) *DescribeSecurityGroupsResponse {
	s.Body = v
	return s
}

type DescribeServerlessClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16f1441y6p2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sdh0b7f4k5f****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeServerlessClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServerlessClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeServerlessClusterRequest) SetClusterId(v string) *DescribeServerlessClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeServerlessClusterRequest) SetZoneId(v string) *DescribeServerlessClusterRequest {
	s.ZoneId = &v
	return s
}

type DescribeServerlessClusterResponseBody struct {
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// single
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// example:
	//
	// 2019-10-12T14:40:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 150
	CuSize *string `json:"CuSize,omitempty" xml:"CuSize,omitempty"`
	// example:
	//
	// 200
	DiskSize *string `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// 2019-10-12T14:40:46
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// false
	HaType *string `json:"HaType,omitempty" xml:"HaType,omitempty"`
	// example:
	//
	// false
	HasUser *string `json:"HasUser,omitempty" xml:"HasUser,omitempty"`
	// example:
	//
	// https://sh-wz91452kg946i****-lindorm-serverless-in.lindorm.rds.aliyuncs.com:443
	InnerEndpoint *string `json:"InnerEndpoint,omitempty" xml:"InnerEndpoint,omitempty"`
	// example:
	//
	// hb-bp16f1441y6p2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// true
	IsDeletionProtection *string `json:"IsDeletionProtection,omitempty" xml:"IsDeletionProtection,omitempty"`
	LockMode             *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// 2.0.8
	MainVersion *string `json:"MainVersion,omitempty" xml:"MainVersion,omitempty"`
	// example:
	//
	// https://sh-wz91452kg946i****-lindorm-serverless.lindorm.rds.aliyuncs.com:443
	OuterEndpoint *string `json:"OuterEndpoint,omitempty" xml:"OuterEndpoint,omitempty"`
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 89F81C30-320B-4550-91DB-C37C81D2358F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	ReserverMaxQpsNum *string `json:"ReserverMaxQpsNum,omitempty" xml:"ReserverMaxQpsNum,omitempty"`
	// example:
	//
	// 50
	ReserverMinQpsNum *string `json:"ReserverMinQpsNum,omitempty" xml:"ReserverMinQpsNum,omitempty"`
	// example:
	//
	// rg-fjm2d4v7sf****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// NO
	UpdateStatus *string `json:"UpdateStatus,omitempty" xml:"UpdateStatus,omitempty"`
	// example:
	//
	// vsw-bp191ipotqf****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp120k6ixs4eoghz****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeServerlessClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServerlessClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServerlessClusterResponseBody) SetAutoRenew(v string) *DescribeServerlessClusterResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetClusterType(v string) *DescribeServerlessClusterResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetCreateTime(v string) *DescribeServerlessClusterResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetCuSize(v string) *DescribeServerlessClusterResponseBody {
	s.CuSize = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetDiskSize(v string) *DescribeServerlessClusterResponseBody {
	s.DiskSize = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetExpireTime(v string) *DescribeServerlessClusterResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetHaType(v string) *DescribeServerlessClusterResponseBody {
	s.HaType = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetHasUser(v string) *DescribeServerlessClusterResponseBody {
	s.HasUser = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetInnerEndpoint(v string) *DescribeServerlessClusterResponseBody {
	s.InnerEndpoint = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetInstanceId(v string) *DescribeServerlessClusterResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetInstanceName(v string) *DescribeServerlessClusterResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetIsDeletionProtection(v string) *DescribeServerlessClusterResponseBody {
	s.IsDeletionProtection = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetLockMode(v string) *DescribeServerlessClusterResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetMainVersion(v string) *DescribeServerlessClusterResponseBody {
	s.MainVersion = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetOuterEndpoint(v string) *DescribeServerlessClusterResponseBody {
	s.OuterEndpoint = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetPayType(v string) *DescribeServerlessClusterResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetRegionId(v string) *DescribeServerlessClusterResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetRequestId(v string) *DescribeServerlessClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetReserverMaxQpsNum(v string) *DescribeServerlessClusterResponseBody {
	s.ReserverMaxQpsNum = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetReserverMinQpsNum(v string) *DescribeServerlessClusterResponseBody {
	s.ReserverMinQpsNum = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetResourceGroupId(v string) *DescribeServerlessClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetStatus(v string) *DescribeServerlessClusterResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetUpdateStatus(v string) *DescribeServerlessClusterResponseBody {
	s.UpdateStatus = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetVSwitchId(v string) *DescribeServerlessClusterResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetVpcId(v string) *DescribeServerlessClusterResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetZoneId(v string) *DescribeServerlessClusterResponseBody {
	s.ZoneId = &v
	return s
}

type DescribeServerlessClusterResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServerlessClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServerlessClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServerlessClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeServerlessClusterResponse) SetHeaders(v map[string]*string) *DescribeServerlessClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeServerlessClusterResponse) SetStatusCode(v int32) *DescribeServerlessClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServerlessClusterResponse) SetBody(v *DescribeServerlessClusterResponseBody) *DescribeServerlessClusterResponse {
	s.Body = v
	return s
}

type DescribeSubDomainRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSubDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainRequest) SetRegionId(v string) *DescribeSubDomainRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubDomainRequest) SetZoneId(v string) *DescribeSubDomainRequest {
	s.ZoneId = &v
	return s
}

type DescribeSubDomainResponseBody struct {
	// example:
	//
	// F4208C83-B9BC-4A64-A739-8F88E98DA469
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// cn-hangzhou-h-aliyun
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s DescribeSubDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainResponseBody) SetRequestId(v string) *DescribeSubDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubDomainResponseBody) SetSubDomain(v string) *DescribeSubDomainResponseBody {
	s.SubDomain = &v
	return s
}

type DescribeSubDomainResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainResponse) SetHeaders(v map[string]*string) *DescribeSubDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubDomainResponse) SetStatusCode(v int32) *DescribeSubDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubDomainResponse) SetBody(v *DescribeSubDomainResponseBody) *DescribeSubDomainResponse {
	s.Body = v
	return s
}

type EnableHBaseueBackupRequest struct {
	// example:
	//
	// xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 800
	ColdStorageSize *int32 `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	HbaseueClusterId *string `json:"HbaseueClusterId,omitempty" xml:"HbaseueClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
}

func (s EnableHBaseueBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableHBaseueBackupRequest) GoString() string {
	return s.String()
}

func (s *EnableHBaseueBackupRequest) SetClientToken(v string) *EnableHBaseueBackupRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableHBaseueBackupRequest) SetColdStorageSize(v int32) *EnableHBaseueBackupRequest {
	s.ColdStorageSize = &v
	return s
}

func (s *EnableHBaseueBackupRequest) SetHbaseueClusterId(v string) *EnableHBaseueBackupRequest {
	s.HbaseueClusterId = &v
	return s
}

func (s *EnableHBaseueBackupRequest) SetNodeCount(v int32) *EnableHBaseueBackupRequest {
	s.NodeCount = &v
	return s
}

type EnableHBaseueBackupResponseBody struct {
	// example:
	//
	// bds-m5e54q06ceyhxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1449xxx
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 15272D5D-46E8-4400-9CC8-A7E7B589F575
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableHBaseueBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableHBaseueBackupResponseBody) GoString() string {
	return s.String()
}

func (s *EnableHBaseueBackupResponseBody) SetClusterId(v string) *EnableHBaseueBackupResponseBody {
	s.ClusterId = &v
	return s
}

func (s *EnableHBaseueBackupResponseBody) SetOrderId(v string) *EnableHBaseueBackupResponseBody {
	s.OrderId = &v
	return s
}

func (s *EnableHBaseueBackupResponseBody) SetRequestId(v string) *EnableHBaseueBackupResponseBody {
	s.RequestId = &v
	return s
}

type EnableHBaseueBackupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableHBaseueBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableHBaseueBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableHBaseueBackupResponse) GoString() string {
	return s.String()
}

func (s *EnableHBaseueBackupResponse) SetHeaders(v map[string]*string) *EnableHBaseueBackupResponse {
	s.Headers = v
	return s
}

func (s *EnableHBaseueBackupResponse) SetStatusCode(v int32) *EnableHBaseueBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableHBaseueBackupResponse) SetBody(v *EnableHBaseueBackupResponseBody) *EnableHBaseueBackupResponse {
	s.Body = v
	return s
}

type EnableHBaseueModuleRequest struct {
	// example:
	//
	// 2
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// example:
	//
	// bds-bp174pm3tsk3****
	BdsId *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase.sn1.large
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// example:
	//
	// 400
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	HbaseueClusterId *string `json:"HbaseueClusterId,omitempty" xml:"HbaseueClusterId,omitempty"`
	// example:
	//
	// hbase.sn1.large
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	// example:
	//
	// cluster-name
	ModuleClusterName *string `json:"ModuleClusterName,omitempty" xml:"ModuleClusterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// solr
	ModuleTypeName *string `json:"ModuleTypeName,omitempty" xml:"ModuleTypeName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 6
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp120k6ixs4eog*****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp191ipotqj1ssyl*****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s EnableHBaseueModuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableHBaseueModuleRequest) GoString() string {
	return s.String()
}

func (s *EnableHBaseueModuleRequest) SetAutoRenewPeriod(v int32) *EnableHBaseueModuleRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetBdsId(v string) *EnableHBaseueModuleRequest {
	s.BdsId = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetClientToken(v string) *EnableHBaseueModuleRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetCoreInstanceType(v string) *EnableHBaseueModuleRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetDiskSize(v int32) *EnableHBaseueModuleRequest {
	s.DiskSize = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetDiskType(v string) *EnableHBaseueModuleRequest {
	s.DiskType = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetHbaseueClusterId(v string) *EnableHBaseueModuleRequest {
	s.HbaseueClusterId = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetMasterInstanceType(v string) *EnableHBaseueModuleRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetModuleClusterName(v string) *EnableHBaseueModuleRequest {
	s.ModuleClusterName = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetModuleTypeName(v string) *EnableHBaseueModuleRequest {
	s.ModuleTypeName = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetNodeCount(v int32) *EnableHBaseueModuleRequest {
	s.NodeCount = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetPayType(v string) *EnableHBaseueModuleRequest {
	s.PayType = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetPeriod(v int32) *EnableHBaseueModuleRequest {
	s.Period = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetPeriodUnit(v string) *EnableHBaseueModuleRequest {
	s.PeriodUnit = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetRegionId(v string) *EnableHBaseueModuleRequest {
	s.RegionId = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetVpcId(v string) *EnableHBaseueModuleRequest {
	s.VpcId = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetVswitchId(v string) *EnableHBaseueModuleRequest {
	s.VswitchId = &v
	return s
}

func (s *EnableHBaseueModuleRequest) SetZoneId(v string) *EnableHBaseueModuleRequest {
	s.ZoneId = &v
	return s
}

type EnableHBaseueModuleResponseBody struct {
	// example:
	//
	// ld-bp150tns0sjxs****-m1-ps
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 21474915573****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 407075EA-47F5-5A2D-888F-C1F90B8F3FCA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableHBaseueModuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableHBaseueModuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableHBaseueModuleResponseBody) SetClusterId(v string) *EnableHBaseueModuleResponseBody {
	s.ClusterId = &v
	return s
}

func (s *EnableHBaseueModuleResponseBody) SetOrderId(v string) *EnableHBaseueModuleResponseBody {
	s.OrderId = &v
	return s
}

func (s *EnableHBaseueModuleResponseBody) SetRequestId(v string) *EnableHBaseueModuleResponseBody {
	s.RequestId = &v
	return s
}

type EnableHBaseueModuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableHBaseueModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableHBaseueModuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableHBaseueModuleResponse) GoString() string {
	return s.String()
}

func (s *EnableHBaseueModuleResponse) SetHeaders(v map[string]*string) *EnableHBaseueModuleResponse {
	s.Headers = v
	return s
}

func (s *EnableHBaseueModuleResponse) SetStatusCode(v int32) *EnableHBaseueModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableHBaseueModuleResponse) SetBody(v *EnableHBaseueModuleResponseBody) *EnableHBaseueModuleResponse {
	s.Body = v
	return s
}

type EvaluateMultiZoneResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// vsw-hangxzhouxb****
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.0
	ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	// example:
	//
	// 0
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// example:
	//
	// f4g8t5rd2gr94****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// hbaseue_test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 400
	CoreDiskSize *int32 `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cloud_ssd
	CoreDiskType *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase.sn1.medium
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	CoreNodeCount *int32 `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbaseue
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// 400
	LogDiskSize *int32 `json:"LogDiskSize,omitempty" xml:"LogDiskSize,omitempty"`
	// example:
	//
	// cloud_ssd
	LogDiskType *string `json:"LogDiskType,omitempty" xml:"LogDiskType,omitempty"`
	// example:
	//
	// hbase.sn1.medium
	LogInstanceType *string `json:"LogInstanceType,omitempty" xml:"LogInstanceType,omitempty"`
	// example:
	//
	// 4
	LogNodeCount *int32 `json:"LogNodeCount,omitempty" xml:"LogNodeCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase.sn1.medium
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-bef-aliyun-com
	MultiZoneCombination *string `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-hangxzhouxe*****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-e
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-hangxzhouxf****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-f
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp120k6ixs4eog*****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s EvaluateMultiZoneResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s EvaluateMultiZoneResourceRequest) GoString() string {
	return s.String()
}

func (s *EvaluateMultiZoneResourceRequest) SetArbiterVSwitchId(v string) *EvaluateMultiZoneResourceRequest {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetArbiterZoneId(v string) *EvaluateMultiZoneResourceRequest {
	s.ArbiterZoneId = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetArchVersion(v string) *EvaluateMultiZoneResourceRequest {
	s.ArchVersion = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetAutoRenewPeriod(v int32) *EvaluateMultiZoneResourceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetClientToken(v string) *EvaluateMultiZoneResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetClusterName(v string) *EvaluateMultiZoneResourceRequest {
	s.ClusterName = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetCoreDiskSize(v int32) *EvaluateMultiZoneResourceRequest {
	s.CoreDiskSize = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetCoreDiskType(v string) *EvaluateMultiZoneResourceRequest {
	s.CoreDiskType = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetCoreInstanceType(v string) *EvaluateMultiZoneResourceRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetCoreNodeCount(v int32) *EvaluateMultiZoneResourceRequest {
	s.CoreNodeCount = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetEngine(v string) *EvaluateMultiZoneResourceRequest {
	s.Engine = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetEngineVersion(v string) *EvaluateMultiZoneResourceRequest {
	s.EngineVersion = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetLogDiskSize(v int32) *EvaluateMultiZoneResourceRequest {
	s.LogDiskSize = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetLogDiskType(v string) *EvaluateMultiZoneResourceRequest {
	s.LogDiskType = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetLogInstanceType(v string) *EvaluateMultiZoneResourceRequest {
	s.LogInstanceType = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetLogNodeCount(v int32) *EvaluateMultiZoneResourceRequest {
	s.LogNodeCount = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetMasterInstanceType(v string) *EvaluateMultiZoneResourceRequest {
	s.MasterInstanceType = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetMultiZoneCombination(v string) *EvaluateMultiZoneResourceRequest {
	s.MultiZoneCombination = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetPayType(v string) *EvaluateMultiZoneResourceRequest {
	s.PayType = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetPeriod(v int32) *EvaluateMultiZoneResourceRequest {
	s.Period = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetPeriodUnit(v string) *EvaluateMultiZoneResourceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetPrimaryVSwitchId(v string) *EvaluateMultiZoneResourceRequest {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetPrimaryZoneId(v string) *EvaluateMultiZoneResourceRequest {
	s.PrimaryZoneId = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetRegionId(v string) *EvaluateMultiZoneResourceRequest {
	s.RegionId = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetSecurityIPList(v string) *EvaluateMultiZoneResourceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetStandbyVSwitchId(v string) *EvaluateMultiZoneResourceRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetStandbyZoneId(v string) *EvaluateMultiZoneResourceRequest {
	s.StandbyZoneId = &v
	return s
}

func (s *EvaluateMultiZoneResourceRequest) SetVpcId(v string) *EvaluateMultiZoneResourceRequest {
	s.VpcId = &v
	return s
}

type EvaluateMultiZoneResourceResponseBody struct {
	// example:
	//
	// FB703B69-D4D4-4879-B9FE-6A37F67C46FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EvaluateMultiZoneResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EvaluateMultiZoneResourceResponseBody) GoString() string {
	return s.String()
}

func (s *EvaluateMultiZoneResourceResponseBody) SetRequestId(v string) *EvaluateMultiZoneResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *EvaluateMultiZoneResourceResponseBody) SetSuccess(v bool) *EvaluateMultiZoneResourceResponseBody {
	s.Success = &v
	return s
}

type EvaluateMultiZoneResourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EvaluateMultiZoneResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EvaluateMultiZoneResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s EvaluateMultiZoneResourceResponse) GoString() string {
	return s.String()
}

func (s *EvaluateMultiZoneResourceResponse) SetHeaders(v map[string]*string) *EvaluateMultiZoneResourceResponse {
	s.Headers = v
	return s
}

func (s *EvaluateMultiZoneResourceResponse) SetStatusCode(v int32) *EvaluateMultiZoneResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *EvaluateMultiZoneResourceResponse) SetBody(v *EvaluateMultiZoneResourceResponseBody) *EvaluateMultiZoneResourceResponse {
	s.Body = v
	return s
}

type GetMultimodeCmsUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetMultimodeCmsUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMultimodeCmsUrlRequest) GoString() string {
	return s.String()
}

func (s *GetMultimodeCmsUrlRequest) SetClusterId(v string) *GetMultimodeCmsUrlRequest {
	s.ClusterId = &v
	return s
}

func (s *GetMultimodeCmsUrlRequest) SetRegionId(v string) *GetMultimodeCmsUrlRequest {
	s.RegionId = &v
	return s
}

type GetMultimodeCmsUrlResponseBody struct {
	// example:
	//
	// hb-t4naqsay5gn******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// htpp://***********
	MultimodCmsUrl *string `json:"MultimodCmsUrl,omitempty" xml:"MultimodCmsUrl,omitempty"`
	// example:
	//
	// 44183B05-852E-4716-B902-52977140190F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMultimodeCmsUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultimodeCmsUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultimodeCmsUrlResponseBody) SetClusterId(v string) *GetMultimodeCmsUrlResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetMultimodeCmsUrlResponseBody) SetMultimodCmsUrl(v string) *GetMultimodeCmsUrlResponseBody {
	s.MultimodCmsUrl = &v
	return s
}

func (s *GetMultimodeCmsUrlResponseBody) SetRequestId(v string) *GetMultimodeCmsUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetMultimodeCmsUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultimodeCmsUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultimodeCmsUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMultimodeCmsUrlResponse) GoString() string {
	return s.String()
}

func (s *GetMultimodeCmsUrlResponse) SetHeaders(v map[string]*string) *GetMultimodeCmsUrlResponse {
	s.Headers = v
	return s
}

func (s *GetMultimodeCmsUrlResponse) SetStatusCode(v int32) *GetMultimodeCmsUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultimodeCmsUrlResponse) SetBody(v *GetMultimodeCmsUrlResponseBody) *GetMultimodeCmsUrlResponse {
	s.Body = v
	return s
}

type GrantRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test01
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// READ,WRITE
	AclActions *string `json:"AclActions,omitempty" xml:"AclActions,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GrantRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantRequest) GoString() string {
	return s.String()
}

func (s *GrantRequest) SetAccountName(v string) *GrantRequest {
	s.AccountName = &v
	return s
}

func (s *GrantRequest) SetAclActions(v string) *GrantRequest {
	s.AclActions = &v
	return s
}

func (s *GrantRequest) SetClusterId(v string) *GrantRequest {
	s.ClusterId = &v
	return s
}

func (s *GrantRequest) SetNamespace(v string) *GrantRequest {
	s.Namespace = &v
	return s
}

func (s *GrantRequest) SetTableName(v string) *GrantRequest {
	s.TableName = &v
	return s
}

type GrantResponseBody struct {
	// example:
	//
	// 9CBF8DF0-4931-4A54-9B60-4C6E1AB5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantResponseBody) GoString() string {
	return s.String()
}

func (s *GrantResponseBody) SetRequestId(v string) *GrantResponseBody {
	s.RequestId = &v
	return s
}

type GrantResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantResponse) GoString() string {
	return s.String()
}

func (s *GrantResponse) SetHeaders(v map[string]*string) *GrantResponse {
	s.Headers = v
	return s
}

func (s *GrantResponse) SetStatusCode(v int32) *GrantResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantResponse) SetBody(v *GrantResponseBody) *GrantResponse {
	s.Body = v
	return s
}

type ListHBaseInstancesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// vpc-t4nx81tmlixcq5i****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListHBaseInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHBaseInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListHBaseInstancesRequest) SetVpcId(v string) *ListHBaseInstancesRequest {
	s.VpcId = &v
	return s
}

type ListHBaseInstancesResponseBody struct {
	Instances *ListHBaseInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// example:
	//
	// 89F81C30-320B-4550-91DB-C37C81D2358F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHBaseInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHBaseInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHBaseInstancesResponseBody) SetInstances(v *ListHBaseInstancesResponseBodyInstances) *ListHBaseInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListHBaseInstancesResponseBody) SetRequestId(v string) *ListHBaseInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ListHBaseInstancesResponseBodyInstances struct {
	Instance []*ListHBaseInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ListHBaseInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListHBaseInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListHBaseInstancesResponseBodyInstances) SetInstance(v []*ListHBaseInstancesResponseBodyInstancesInstance) *ListHBaseInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

type ListHBaseInstancesResponseBodyInstancesInstance struct {
	// example:
	//
	// hb-t4naqsay5gn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// name_test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
}

func (s ListHBaseInstancesResponseBodyInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s ListHBaseInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *ListHBaseInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *ListHBaseInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *ListHBaseInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *ListHBaseInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *ListHBaseInstancesResponseBodyInstancesInstance) SetIsDefault(v bool) *ListHBaseInstancesResponseBodyInstancesInstance {
	s.IsDefault = &v
	return s
}

type ListHBaseInstancesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHBaseInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHBaseInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHBaseInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListHBaseInstancesResponse) SetHeaders(v map[string]*string) *ListHBaseInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListHBaseInstancesResponse) SetStatusCode(v int32) *ListHBaseInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHBaseInstancesResponse) SetBody(v *ListHBaseInstancesResponseBody) *ListHBaseInstancesResponse {
	s.Body = v
	return s
}

type ListInstanceServiceConfigHistoriesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInstanceServiceConfigHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigHistoriesRequest) SetClusterId(v string) *ListInstanceServiceConfigHistoriesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesRequest) SetPageNumber(v int32) *ListInstanceServiceConfigHistoriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesRequest) SetPageSize(v int32) *ListInstanceServiceConfigHistoriesRequest {
	s.PageSize = &v
	return s
}

type ListInstanceServiceConfigHistoriesResponseBody struct {
	ConfigureHistoryList *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList `json:"ConfigureHistoryList,omitempty" xml:"ConfigureHistoryList,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 658C1549-2C02-4FD9-9490-EB3B285F9DCA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListInstanceServiceConfigHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) SetConfigureHistoryList(v *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList) *ListInstanceServiceConfigHistoriesResponseBody {
	s.ConfigureHistoryList = v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) SetPageNumber(v int32) *ListInstanceServiceConfigHistoriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) SetPageRecordCount(v int32) *ListInstanceServiceConfigHistoriesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) SetRequestId(v string) *ListInstanceServiceConfigHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBody) SetTotalRecordCount(v int64) *ListInstanceServiceConfigHistoriesResponseBody {
	s.TotalRecordCount = &v
	return s
}

type ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList struct {
	Config []*ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
}

func (s ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList) SetConfig(v []*ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryList {
	s.Config = v
	return s
}

type ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig struct {
	// example:
	//
	// hbase#hbase-site.xml#hbase.client.keyvalue.maxsize
	ConfigureName *string `json:"ConfigureName,omitempty" xml:"ConfigureName,omitempty"`
	// example:
	//
	// 1608708923000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// false
	Effective *string `json:"Effective,omitempty" xml:"Effective,omitempty"`
	// example:
	//
	// 10485770
	NewValue *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	// example:
	//
	// 10485760
	OldValue *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
}

func (s ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) SetConfigureName(v string) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	s.ConfigureName = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) SetCreateTime(v string) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	s.CreateTime = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) SetEffective(v string) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	s.Effective = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) SetNewValue(v string) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	s.NewValue = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig) SetOldValue(v string) *ListInstanceServiceConfigHistoriesResponseBodyConfigureHistoryListConfig {
	s.OldValue = &v
	return s
}

type ListInstanceServiceConfigHistoriesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceServiceConfigHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceServiceConfigHistoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigHistoriesResponse) SetHeaders(v map[string]*string) *ListInstanceServiceConfigHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponse) SetStatusCode(v int32) *ListInstanceServiceConfigHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesResponse) SetBody(v *ListInstanceServiceConfigHistoriesResponseBody) *ListInstanceServiceConfigHistoriesResponse {
	s.Body = v
	return s
}

type ListInstanceServiceConfigurationsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInstanceServiceConfigurationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigurationsRequest) SetClusterId(v string) *ListInstanceServiceConfigurationsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInstanceServiceConfigurationsRequest) SetPageNumber(v int32) *ListInstanceServiceConfigurationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceServiceConfigurationsRequest) SetPageSize(v int32) *ListInstanceServiceConfigurationsRequest {
	s.PageSize = &v
	return s
}

type ListInstanceServiceConfigurationsResponseBody struct {
	ConfigureList *ListInstanceServiceConfigurationsResponseBodyConfigureList `json:"ConfigureList,omitempty" xml:"ConfigureList,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 5B381E36-BCA3-4377-8638-B65C236617D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 42
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListInstanceServiceConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigurationsResponseBody) SetConfigureList(v *ListInstanceServiceConfigurationsResponseBodyConfigureList) *ListInstanceServiceConfigurationsResponseBody {
	s.ConfigureList = v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBody) SetPageNumber(v int32) *ListInstanceServiceConfigurationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBody) SetPageRecordCount(v int32) *ListInstanceServiceConfigurationsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBody) SetRequestId(v string) *ListInstanceServiceConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBody) SetTotalRecordCount(v int64) *ListInstanceServiceConfigurationsResponseBody {
	s.TotalRecordCount = &v
	return s
}

type ListInstanceServiceConfigurationsResponseBodyConfigureList struct {
	Config []*ListInstanceServiceConfigurationsResponseBodyConfigureListConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
}

func (s ListInstanceServiceConfigurationsResponseBodyConfigureList) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigurationsResponseBodyConfigureList) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureList) SetConfig(v []*ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) *ListInstanceServiceConfigurationsResponseBodyConfigureList {
	s.Config = v
	return s
}

type ListInstanceServiceConfigurationsResponseBodyConfigureListConfig struct {
	// example:
	//
	// hbase#hbase-site.xml#hbase.client.keyvalue.maxsize
	ConfigureName *string `json:"ConfigureName,omitempty" xml:"ConfigureName,omitempty"`
	// example:
	//
	// INT
	ConfigureUnit *string `json:"ConfigureUnit,omitempty" xml:"ConfigureUnit,omitempty"`
	// example:
	//
	// 10485760
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// hbase client keyvalue maxsize
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	NeedRestart *string `json:"NeedRestart,omitempty" xml:"NeedRestart,omitempty"`
	// example:
	//
	// 10485760
	RunningValue *string `json:"RunningValue,omitempty" xml:"RunningValue,omitempty"`
	// example:
	//
	// R[10485760,52428800]
	ValueRange *string `json:"ValueRange,omitempty" xml:"ValueRange,omitempty"`
}

func (s ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetConfigureName(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.ConfigureName = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetConfigureUnit(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.ConfigureUnit = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetDefaultValue(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.DefaultValue = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetDescription(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.Description = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetNeedRestart(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.NeedRestart = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetRunningValue(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.RunningValue = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig) SetValueRange(v string) *ListInstanceServiceConfigurationsResponseBodyConfigureListConfig {
	s.ValueRange = &v
	return s
}

type ListInstanceServiceConfigurationsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceServiceConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceServiceConfigurationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceServiceConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigurationsResponse) SetHeaders(v map[string]*string) *ListInstanceServiceConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceServiceConfigurationsResponse) SetStatusCode(v int32) *ListInstanceServiceConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceServiceConfigurationsResponse) SetBody(v *ListInstanceServiceConfigurationsResponseBody) *ListInstanceServiceConfigurationsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// example:
	//
	// NextToken
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// bds-bp15e022622f****
	ResourceId []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag        []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
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
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
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
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 74F02441-9A8D-48F6-933F-E317AEB28DBF
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
	// example:
	//
	// bds-bp15e022622f****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// ALIYUN::MULTIMOD::CLUSTER
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// k1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// v2
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListTagsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
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
	// example:
	//
	// 36D1BE9B-3C4A-425B-947A-69E3D77999C4
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
	// example:
	//
	// k1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// v2
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagsResponseBodyTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTagsTag) SetTagKey(v string) *ListTagsResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListTagsResponseBodyTagsTag) SetTagValue(v string) *ListTagsResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

type ListTagsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListTagsResponse) SetStatusCode(v int32) *ListTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagsResponse) SetBody(v *ListTagsResponseBody) *ListTagsResponse {
	s.Body = v
	return s
}

type ModifyAccountPasswordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test01
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// **********
	NewAccountPassword *string `json:"NewAccountPassword,omitempty" xml:"NewAccountPassword,omitempty"`
}

func (s ModifyAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordRequest) SetAccountName(v string) *ModifyAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetClusterId(v string) *ModifyAccountPasswordRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetNewAccountPassword(v string) *ModifyAccountPasswordRequest {
	s.NewAccountPassword = &v
	return s
}

type ModifyAccountPasswordResponseBody struct {
	// example:
	//
	// AFAA617B-3268-5883-982B-DB8EC8CC1F1B
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyAccountPasswordResponse) SetStatusCode(v int32) *ModifyAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountPasswordResponse) SetBody(v *ModifyAccountPasswordResponseBody) *ModifyAccountPasswordResponse {
	s.Body = v
	return s
}

type ModifyActiveOperationTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1111,2222
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// example:
	//
	// 1
	ImmediateStart       *int32  `json:"ImmediateStart,omitempty" xml:"ImmediateStart,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-02-17T18:50:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s ModifyActiveOperationTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyActiveOperationTasksRequest) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTasksRequest) SetIds(v string) *ModifyActiveOperationTasksRequest {
	s.Ids = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetImmediateStart(v int32) *ModifyActiveOperationTasksRequest {
	s.ImmediateStart = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetOwnerAccount(v string) *ModifyActiveOperationTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetOwnerId(v int64) *ModifyActiveOperationTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetResourceOwnerAccount(v string) *ModifyActiveOperationTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetResourceOwnerId(v int64) *ModifyActiveOperationTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetSecurityToken(v string) *ModifyActiveOperationTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetSwitchTime(v string) *ModifyActiveOperationTasksRequest {
	s.SwitchTime = &v
	return s
}

type ModifyActiveOperationTasksResponseBody struct {
	// example:
	//
	// 1111,2222
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// example:
	//
	// 8C9CC46A-9532-4752-B59F-580112C5A45B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyActiveOperationTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyActiveOperationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTasksResponseBody) SetIds(v string) *ModifyActiveOperationTasksResponseBody {
	s.Ids = &v
	return s
}

func (s *ModifyActiveOperationTasksResponseBody) SetRequestId(v string) *ModifyActiveOperationTasksResponseBody {
	s.RequestId = &v
	return s
}

type ModifyActiveOperationTasksResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyActiveOperationTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyActiveOperationTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyActiveOperationTasksResponse) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTasksResponse) SetHeaders(v map[string]*string) *ModifyActiveOperationTasksResponse {
	s.Headers = v
	return s
}

func (s *ModifyActiveOperationTasksResponse) SetStatusCode(v int32) *ModifyActiveOperationTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyActiveOperationTasksResponse) SetBody(v *ModifyActiveOperationTasksResponseBody) *ModifyActiveOperationTasksResponse {
	s.Body = v
	return s
}

type ModifyBackupPlanConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7
	FullBackupCycle *string `json:"FullBackupCycle,omitempty" xml:"FullBackupCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	MinHFileBackupCount *string `json:"MinHFileBackupCount,omitempty" xml:"MinHFileBackupCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020-11-09T18:00:00Z
	NextFullBackupDate *string `json:"NextFullBackupDate,omitempty" xml:"NextFullBackupDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// *
	Tables *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
}

func (s ModifyBackupPlanConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPlanConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanConfigRequest) SetClusterId(v string) *ModifyBackupPlanConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyBackupPlanConfigRequest) SetFullBackupCycle(v string) *ModifyBackupPlanConfigRequest {
	s.FullBackupCycle = &v
	return s
}

func (s *ModifyBackupPlanConfigRequest) SetMinHFileBackupCount(v string) *ModifyBackupPlanConfigRequest {
	s.MinHFileBackupCount = &v
	return s
}

func (s *ModifyBackupPlanConfigRequest) SetNextFullBackupDate(v string) *ModifyBackupPlanConfigRequest {
	s.NextFullBackupDate = &v
	return s
}

func (s *ModifyBackupPlanConfigRequest) SetTables(v string) *ModifyBackupPlanConfigRequest {
	s.Tables = &v
	return s
}

type ModifyBackupPlanConfigResponseBody struct {
	// example:
	//
	// 50F4A8C2-076F-4703-9813-2FCD7FBB91C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupPlanConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPlanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanConfigResponseBody) SetRequestId(v string) *ModifyBackupPlanConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBackupPlanConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackupPlanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBackupPlanConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPlanConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanConfigResponse) SetHeaders(v map[string]*string) *ModifyBackupPlanConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupPlanConfigResponse) SetStatusCode(v int32) *ModifyBackupPlanConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupPlanConfigResponse) SetBody(v *ModifyBackupPlanConfigResponseBody) *ModifyBackupPlanConfigResponse {
	s.Body = v
	return s
}

type ModifyBackupPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 18:00Z
	PreferredBackupEndTimeUTC *string `json:"PreferredBackupEndTimeUTC,omitempty" xml:"PreferredBackupEndTimeUTC,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Thursday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// example:
	//
	// 17:00Z
	PreferredBackupStartTimeUTC *string `json:"PreferredBackupStartTimeUTC,omitempty" xml:"PreferredBackupStartTimeUTC,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 01:00-02:00
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) SetClusterId(v string) *ModifyBackupPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupEndTimeUTC(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupEndTimeUTC = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupStartTimeUTC(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupStartTimeUTC = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

type ModifyBackupPolicyResponseBody struct {
	// example:
	//
	// 17E3AC63-300D-4B69-9108-45EC20E50E85
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

type ModifyClusterDeletionProtectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-****************
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Protection *bool `json:"Protection,omitempty" xml:"Protection,omitempty"`
}

func (s ModifyClusterDeletionProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterDeletionProtectionRequest) SetClusterId(v string) *ModifyClusterDeletionProtectionRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterDeletionProtectionRequest) SetProtection(v bool) *ModifyClusterDeletionProtectionRequest {
	s.Protection = &v
	return s
}

type ModifyClusterDeletionProtectionResponseBody struct {
	// example:
	//
	// 24C80BD8-C710-4138-893A-D2AFED4FC13D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterDeletionProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterDeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterDeletionProtectionResponseBody) SetRequestId(v string) *ModifyClusterDeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterDeletionProtectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterDeletionProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterDeletionProtectionResponse) SetHeaders(v map[string]*string) *ModifyClusterDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterDeletionProtectionResponse) SetStatusCode(v int32) *ModifyClusterDeletionProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterDeletionProtectionResponse) SetBody(v *ModifyClusterDeletionProtectionResponseBody) *ModifyClusterDeletionProtectionResponse {
	s.Body = v
	return s
}

type ModifyDiskWarningLineRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 80
	WarningLine *int32 `json:"WarningLine,omitempty" xml:"WarningLine,omitempty"`
}

func (s ModifyDiskWarningLineRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskWarningLineRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskWarningLineRequest) SetClusterId(v string) *ModifyDiskWarningLineRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyDiskWarningLineRequest) SetWarningLine(v int32) *ModifyDiskWarningLineRequest {
	s.WarningLine = &v
	return s
}

type ModifyDiskWarningLineResponseBody struct {
	// example:
	//
	// FC4A930D-3AEE-4C9D-BC70-C0F2EEEAA174
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDiskWarningLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskWarningLineResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskWarningLineResponseBody) SetRequestId(v string) *ModifyDiskWarningLineResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDiskWarningLineResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskWarningLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskWarningLineResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskWarningLineResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskWarningLineResponse) SetHeaders(v map[string]*string) *ModifyDiskWarningLineResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskWarningLineResponse) SetStatusCode(v int32) *ModifyDiskWarningLineResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskWarningLineResponse) SetBody(v *ModifyDiskWarningLineResponseBody) *ModifyDiskWarningLineResponse {
	s.Body = v
	return s
}

type ModifyInstanceMaintainTimeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1b**6jco89****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 06:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 02:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
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

func (s *ModifyInstanceMaintainTimeRequest) SetMaintainEndTime(v string) *ModifyInstanceMaintainTimeRequest {
	s.MaintainEndTime = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetMaintainStartTime(v string) *ModifyInstanceMaintainTimeRequest {
	s.MaintainStartTime = &v
	return s
}

type ModifyInstanceMaintainTimeResponseBody struct {
	// example:
	//
	// C9085433-A56A-4089-B49A-DF5A4E2B7B06
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyInstanceMaintainTimeResponse) SetStatusCode(v int32) *ModifyInstanceMaintainTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceMaintainTimeResponse) SetBody(v *ModifyInstanceMaintainTimeResponseBody) *ModifyInstanceMaintainTimeResponse {
	s.Body = v
	return s
}

type ModifyInstanceNameRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testhbaseone
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyInstanceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameRequest) SetClientToken(v string) *ModifyInstanceNameRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetClusterId(v string) *ModifyInstanceNameRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetClusterName(v string) *ModifyInstanceNameRequest {
	s.ClusterName = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetRegionId(v string) *ModifyInstanceNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetZoneId(v string) *ModifyInstanceNameRequest {
	s.ZoneId = &v
	return s
}

type ModifyInstanceNameResponseBody struct {
	// example:
	//
	// 959DA199-54E5-569D-AD46-92BED8515E62
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameResponseBody) SetRequestId(v string) *ModifyInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameResponse) SetHeaders(v map[string]*string) *ModifyInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceNameResponse) SetStatusCode(v int32) *ModifyInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceNameResponse) SetBody(v *ModifyInstanceNameResponseBody) *ModifyInstanceNameResponse {
	s.Body = v
	return s
}

type ModifyInstanceServiceConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase#hbase-site.xml#hbase.client.keyvalue.maxsize
	ConfigureName *string `json:"ConfigureName,omitempty" xml:"ConfigureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10485770
	ConfigureValue *string `json:"ConfigureValue,omitempty" xml:"ConfigureValue,omitempty"`
	// example:
	//
	// {"key1=value1", "key2=value2"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// false
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s ModifyInstanceServiceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceServiceConfigRequest) SetClusterId(v string) *ModifyInstanceServiceConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyInstanceServiceConfigRequest) SetConfigureName(v string) *ModifyInstanceServiceConfigRequest {
	s.ConfigureName = &v
	return s
}

func (s *ModifyInstanceServiceConfigRequest) SetConfigureValue(v string) *ModifyInstanceServiceConfigRequest {
	s.ConfigureValue = &v
	return s
}

func (s *ModifyInstanceServiceConfigRequest) SetParameters(v string) *ModifyInstanceServiceConfigRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyInstanceServiceConfigRequest) SetRestart(v bool) *ModifyInstanceServiceConfigRequest {
	s.Restart = &v
	return s
}

type ModifyInstanceServiceConfigResponseBody struct {
	// example:
	//
	// F008B7AB-025D-4C20-AE12-047C8F8C3D97
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceServiceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceServiceConfigResponseBody) SetRequestId(v string) *ModifyInstanceServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceServiceConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceServiceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceServiceConfigResponse) SetHeaders(v map[string]*string) *ModifyInstanceServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceServiceConfigResponse) SetStatusCode(v int32) *ModifyInstanceServiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceServiceConfigResponse) SetBody(v *ModifyInstanceServiceConfigResponseBody) *ModifyInstanceServiceConfigResponse {
	s.Body = v
	return s
}

type ModifyInstanceTypeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp1x940uh********
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// hbase.sn1.8xlarge
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// example:
	//
	// hbase.sn1.large
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
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

func (s *ModifyInstanceTypeRequest) SetCoreInstanceType(v string) *ModifyInstanceTypeRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *ModifyInstanceTypeRequest) SetMasterInstanceType(v string) *ModifyInstanceTypeRequest {
	s.MasterInstanceType = &v
	return s
}

type ModifyInstanceTypeResponseBody struct {
	// example:
	//
	// 123412341234123
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 3E19E345-101D-4014-946C-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTypeResponseBody) SetOrderId(v string) *ModifyInstanceTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyInstanceTypeResponseBody) SetRequestId(v string) *ModifyInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceTypeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyInstanceTypeResponse) SetStatusCode(v int32) *ModifyInstanceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceTypeResponse) SetBody(v *ModifyInstanceTypeResponseBody) *ModifyInstanceTypeResponse {
	s.Body = v
	return s
}

type ModifyIpWhitelistRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1uoihlf82e8****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// group_01
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 42.120.XX.XX
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
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

func (s *ModifyIpWhitelistRequest) SetGroupName(v string) *ModifyIpWhitelistRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyIpWhitelistRequest) SetIpList(v string) *ModifyIpWhitelistRequest {
	s.IpList = &v
	return s
}

func (s *ModifyIpWhitelistRequest) SetIpVersion(v string) *ModifyIpWhitelistRequest {
	s.IpVersion = &v
	return s
}

type ModifyIpWhitelistResponseBody struct {
	// example:
	//
	// 101CFA8A-FB88-5014-A10C-3A0DA9AD8B0B
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyIpWhitelistResponse) SetStatusCode(v int32) *ModifyIpWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIpWhitelistResponse) SetBody(v *ModifyIpWhitelistResponseBody) *ModifyIpWhitelistResponse {
	s.Body = v
	return s
}

type ModifyMultiZoneClusterNodeTypeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-dj45g7d6rbrd****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// hbase.sn1.large
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// example:
	//
	// hbase.sn1.2xlarge
	LogInstanceType *string `json:"LogInstanceType,omitempty" xml:"LogInstanceType,omitempty"`
	// example:
	//
	// hbase.sn1.8xlarge
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
}

func (s ModifyMultiZoneClusterNodeTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMultiZoneClusterNodeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) SetClusterId(v string) *ModifyMultiZoneClusterNodeTypeRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) SetCoreInstanceType(v string) *ModifyMultiZoneClusterNodeTypeRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) SetLogInstanceType(v string) *ModifyMultiZoneClusterNodeTypeRequest {
	s.LogInstanceType = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeRequest) SetMasterInstanceType(v string) *ModifyMultiZoneClusterNodeTypeRequest {
	s.MasterInstanceType = &v
	return s
}

type ModifyMultiZoneClusterNodeTypeResponseBody struct {
	// example:
	//
	// 12341234123****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 06CF7A6F-A81C-431D-BACD-793F24A67C54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMultiZoneClusterNodeTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMultiZoneClusterNodeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMultiZoneClusterNodeTypeResponseBody) SetOrderId(v string) *ModifyMultiZoneClusterNodeTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeResponseBody) SetRequestId(v string) *ModifyMultiZoneClusterNodeTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyMultiZoneClusterNodeTypeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMultiZoneClusterNodeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMultiZoneClusterNodeTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMultiZoneClusterNodeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyMultiZoneClusterNodeTypeResponse) SetHeaders(v map[string]*string) *ModifyMultiZoneClusterNodeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeResponse) SetStatusCode(v int32) *ModifyMultiZoneClusterNodeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMultiZoneClusterNodeTypeResponse) SetBody(v *ModifyMultiZoneClusterNodeTypeResponseBody) *ModifyMultiZoneClusterNodeTypeResponse {
	s.Body = v
	return s
}

type ModifySecurityGroupsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16f1441y6p2kv**
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sg-t4ng4yyc916o81nu****,sg-x4gg4dyc9d6w********
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
	// example:
	//
	// F4AD2E65-482B-46B6-942E-765989B1C8A3
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifySecurityGroupsResponse) SetStatusCode(v int32) *ModifySecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityGroupsResponse) SetBody(v *ModifySecurityGroupsResponseBody) *ModifySecurityGroupsResponse {
	s.Body = v
	return s
}

type ModifyUIAccountPasswordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test01
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// **********
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ModifyUIAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUIAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyUIAccountPasswordRequest) SetAccountName(v string) *ModifyUIAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyUIAccountPasswordRequest) SetAccountPassword(v string) *ModifyUIAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ModifyUIAccountPasswordRequest) SetClusterId(v string) *ModifyUIAccountPasswordRequest {
	s.ClusterId = &v
	return s
}

type ModifyUIAccountPasswordResponseBody struct {
	// example:
	//
	// BED4ADEB-4EA9-507E-892C-84112D6AC7C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUIAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUIAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUIAccountPasswordResponseBody) SetRequestId(v string) *ModifyUIAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUIAccountPasswordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUIAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUIAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUIAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyUIAccountPasswordResponse) SetHeaders(v map[string]*string) *ModifyUIAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyUIAccountPasswordResponse) SetStatusCode(v int32) *ModifyUIAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUIAccountPasswordResponse) SetBody(v *ModifyUIAccountPasswordResponseBody) *ModifyUIAccountPasswordResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp169l540vc6c****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rg-aekzrk6zzsy****
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
	// example:
	//
	// 8CD9BFBC-D575-5FCC-BA7E-956BF0D0****
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *MoveResourceGroupResponse) SetStatusCode(v int32) *MoveResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type OpenBackupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s OpenBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenBackupRequest) GoString() string {
	return s.String()
}

func (s *OpenBackupRequest) SetClusterId(v string) *OpenBackupRequest {
	s.ClusterId = &v
	return s
}

type OpenBackupResponseBody struct {
	// example:
	//
	// C977DF60-7D06-4E34-A27D-8BC696C5112A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenBackupResponseBody) GoString() string {
	return s.String()
}

func (s *OpenBackupResponseBody) SetRequestId(v string) *OpenBackupResponseBody {
	s.RequestId = &v
	return s
}

type OpenBackupResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenBackupResponse) GoString() string {
	return s.String()
}

func (s *OpenBackupResponse) SetHeaders(v map[string]*string) *OpenBackupResponse {
	s.Headers = v
	return s
}

func (s *OpenBackupResponse) SetStatusCode(v int32) *OpenBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenBackupResponse) SetBody(v *OpenBackupResponseBody) *OpenBackupResponse {
	s.Body = v
	return s
}

type PurgeInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-m5ek15uzs7613xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s PurgeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s PurgeInstanceRequest) GoString() string {
	return s.String()
}

func (s *PurgeInstanceRequest) SetClusterId(v string) *PurgeInstanceRequest {
	s.ClusterId = &v
	return s
}

type PurgeInstanceResponseBody struct {
	// example:
	//
	// 276F899F-E952-496F-81B8-BD46D86854E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PurgeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PurgeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PurgeInstanceResponseBody) SetRequestId(v string) *PurgeInstanceResponseBody {
	s.RequestId = &v
	return s
}

type PurgeInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurgeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PurgeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s PurgeInstanceResponse) GoString() string {
	return s.String()
}

func (s *PurgeInstanceResponse) SetHeaders(v map[string]*string) *PurgeInstanceResponse {
	s.Headers = v
	return s
}

func (s *PurgeInstanceResponse) SetStatusCode(v int32) *PurgeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *PurgeInstanceResponse) SetBody(v *PurgeInstanceResponseBody) *PurgeInstanceResponse {
	s.Body = v
	return s
}

type QueryHBaseHaDBRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bds-t4n3496whj23ia4k
	BdsId *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
}

func (s QueryHBaseHaDBRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBRequest) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBRequest) SetBdsId(v string) *QueryHBaseHaDBRequest {
	s.BdsId = &v
	return s
}

type QueryHBaseHaDBResponseBody struct {
	ClusterList *QueryHBaseHaDBResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 963355AD-A3B1-4654-AFFC-B5186EB8F889
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryHBaseHaDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBody) SetClusterList(v *QueryHBaseHaDBResponseBodyClusterList) *QueryHBaseHaDBResponseBody {
	s.ClusterList = v
	return s
}

func (s *QueryHBaseHaDBResponseBody) SetPageNumber(v int32) *QueryHBaseHaDBResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryHBaseHaDBResponseBody) SetPageSize(v int32) *QueryHBaseHaDBResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryHBaseHaDBResponseBody) SetRequestId(v string) *QueryHBaseHaDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHBaseHaDBResponseBody) SetTotalCount(v int64) *QueryHBaseHaDBResponseBody {
	s.TotalCount = &v
	return s
}

type QueryHBaseHaDBResponseBodyClusterList struct {
	Cluster []*QueryHBaseHaDBResponseBodyClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s QueryHBaseHaDBResponseBodyClusterList) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBodyClusterList) SetCluster(v []*QueryHBaseHaDBResponseBodyClusterListCluster) *QueryHBaseHaDBResponseBodyClusterList {
	s.Cluster = v
	return s
}

type QueryHBaseHaDBResponseBodyClusterListCluster struct {
	// example:
	//
	// hb-t4nn7dy1u1etbzmzm
	ActiveName *string `json:"ActiveName,omitempty" xml:"ActiveName,omitempty"`
	// bdsId
	//
	// example:
	//
	// bds-t4n3496whj23ia4k
	BdsName *string `json:"BdsName,omitempty" xml:"BdsName,omitempty"`
	// example:
	//
	// ha-v21tmnxjwh2yuy1il
	HaName        *string                                                    `json:"HaName,omitempty" xml:"HaName,omitempty"`
	HaSlbConnList *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList `json:"HaSlbConnList,omitempty" xml:"HaSlbConnList,omitempty" type:"Struct"`
	// example:
	//
	// hb-t4n0ye37832tx22vz
	StandbyName *string `json:"StandbyName,omitempty" xml:"StandbyName,omitempty"`
}

func (s QueryHBaseHaDBResponseBodyClusterListCluster) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBResponseBodyClusterListCluster) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetActiveName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.ActiveName = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetBdsName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.BdsName = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetHaName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.HaName = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetHaSlbConnList(v *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.HaSlbConnList = v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetStandbyName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.StandbyName = &v
	return s
}

type QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList struct {
	HaSlbConn []*QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn `json:"HaSlbConn,omitempty" xml:"HaSlbConn,omitempty" type:"Repeated"`
}

func (s QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList) SetHaSlbConn(v []*QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList {
	s.HaSlbConn = v
	return s
}

type QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn struct {
	// example:
	//
	// Standby
	HbaseType *string `json:"HbaseType,omitempty" xml:"HbaseType,omitempty"`
	// example:
	//
	// ha-v21tmnxjwh2yuy1il-phoenix.bds.9b78df04-b.rds.aliyuncs.com:8765
	SlbConnAddr *string `json:"SlbConnAddr,omitempty" xml:"SlbConnAddr,omitempty"`
	// example:
	//
	// phoenix
	SlbType *string `json:"SlbType,omitempty" xml:"SlbType,omitempty"`
}

func (s QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) SetHbaseType(v string) *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn {
	s.HbaseType = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) SetSlbConnAddr(v string) *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn {
	s.SlbConnAddr = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) SetSlbType(v string) *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn {
	s.SlbType = &v
	return s
}

type QueryHBaseHaDBResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHBaseHaDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHBaseHaDBResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHBaseHaDBResponse) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponse) SetHeaders(v map[string]*string) *QueryHBaseHaDBResponse {
	s.Headers = v
	return s
}

func (s *QueryHBaseHaDBResponse) SetStatusCode(v int32) *QueryHBaseHaDBResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHBaseHaDBResponse) SetBody(v *QueryHBaseHaDBResponseBody) *QueryHBaseHaDBResponse {
	s.Body = v
	return s
}

type QueryXpackRelateDBRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ap-bp1qtz9rcbbt3p6ng
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// false
	HasSingleNode *bool `json:"HasSingleNode,omitempty" xml:"HasSingleNode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase
	RelateDbType *string `json:"RelateDbType,omitempty" xml:"RelateDbType,omitempty"`
}

func (s QueryXpackRelateDBRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryXpackRelateDBRequest) GoString() string {
	return s.String()
}

func (s *QueryXpackRelateDBRequest) SetClusterId(v string) *QueryXpackRelateDBRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryXpackRelateDBRequest) SetHasSingleNode(v bool) *QueryXpackRelateDBRequest {
	s.HasSingleNode = &v
	return s
}

func (s *QueryXpackRelateDBRequest) SetRelateDbType(v string) *QueryXpackRelateDBRequest {
	s.RelateDbType = &v
	return s
}

type QueryXpackRelateDBResponseBody struct {
	ClusterList *QueryXpackRelateDBResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	// example:
	//
	// 288E9010-36DD-499C-B4DA-61E4362DA4CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryXpackRelateDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryXpackRelateDBResponseBody) GoString() string {
	return s.String()
}

func (s *QueryXpackRelateDBResponseBody) SetClusterList(v *QueryXpackRelateDBResponseBodyClusterList) *QueryXpackRelateDBResponseBody {
	s.ClusterList = v
	return s
}

func (s *QueryXpackRelateDBResponseBody) SetRequestId(v string) *QueryXpackRelateDBResponseBody {
	s.RequestId = &v
	return s
}

type QueryXpackRelateDBResponseBodyClusterList struct {
	Cluster []*QueryXpackRelateDBResponseBodyClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s QueryXpackRelateDBResponseBodyClusterList) String() string {
	return tea.Prettify(s)
}

func (s QueryXpackRelateDBResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *QueryXpackRelateDBResponseBodyClusterList) SetCluster(v []*QueryXpackRelateDBResponseBodyClusterListCluster) *QueryXpackRelateDBResponseBodyClusterList {
	s.Cluster = v
	return s
}

type QueryXpackRelateDBResponseBodyClusterListCluster struct {
	// example:
	//
	// hb-bp16o0pd52e3y582s
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// hbase_test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// hbase
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 2.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// example:
	//
	// false
	IsRelated *bool `json:"IsRelated,omitempty" xml:"IsRelated,omitempty"`
	// example:
	//
	// ..
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryXpackRelateDBResponseBodyClusterListCluster) String() string {
	return tea.Prettify(s)
}

func (s QueryXpackRelateDBResponseBodyClusterListCluster) GoString() string {
	return s.String()
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetClusterId(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.ClusterId = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetClusterName(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.ClusterName = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetDBType(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.DBType = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetDBVersion(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.DBVersion = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetIsRelated(v bool) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.IsRelated = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetLockMode(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.LockMode = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetStatus(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.Status = &v
	return s
}

type QueryXpackRelateDBResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryXpackRelateDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryXpackRelateDBResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryXpackRelateDBResponse) GoString() string {
	return s.String()
}

func (s *QueryXpackRelateDBResponse) SetHeaders(v map[string]*string) *QueryXpackRelateDBResponse {
	s.Headers = v
	return s
}

func (s *QueryXpackRelateDBResponse) SetStatusCode(v int32) *QueryXpackRelateDBResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryXpackRelateDBResponse) SetBody(v *QueryXpackRelateDBResponseBody) *QueryXpackRelateDBResponse {
	s.Body = v
	return s
}

type RelateDbForHBaseHaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bds-t4nj9v2x85******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hb-bp1w6krp539******
	HaActive *string `json:"HaActive,omitempty" xml:"HaActive,omitempty"`
	// example:
	//
	// hb-t4naqsay5gn******-master1-001.hbase.singapore.rds.aliyuncs.com,hb-t4naqsay5gn******-master3-001.hbase.singapore.rds.aliyuncs.com,hb-t4naqsay5gn******-master2-001.hbase.singapore.rds.aliyuncs.com:2181:/hbase
	HaActiveClusterKey *string `json:"HaActiveClusterKey,omitempty" xml:"HaActiveClusterKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase
	HaActiveDBType *string `json:"HaActiveDBType,omitempty" xml:"HaActiveDBType,omitempty"`
	// example:
	//
	// /hbase
	HaActiveHbaseFsDir *string `json:"HaActiveHbaseFsDir,omitempty" xml:"HaActiveHbaseFsDir,omitempty"`
	// example:
	//
	// hdfs://hb-t4naqsay5gn******-master1-001.hbase.rds.aliyuncs.com:8020,hb-t4naqsay5gn******-master2-001.hbase.rds.aliyuncs.com:8020
	HaActiveHdfsUri *string `json:"HaActiveHdfsUri,omitempty" xml:"HaActiveHdfsUri,omitempty"`
	// example:
	//
	// root
	HaActivePassword *string `json:"HaActivePassword,omitempty" xml:"HaActivePassword,omitempty"`
	// example:
	//
	// root
	HaActiveUser *string `json:"HaActiveUser,omitempty" xml:"HaActiveUser,omitempty"`
	// example:
	//
	// HBase2x
	HaActiveVersion *string `json:"HaActiveVersion,omitempty" xml:"HaActiveVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CLUSTER
	HaMigrateType *string `json:"HaMigrateType,omitempty" xml:"HaMigrateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hb-bp1bl7iqzka******
	HaStandby *string `json:"HaStandby,omitempty" xml:"HaStandby,omitempty"`
	// example:
	//
	// hb-bp1w6krp539******-master1-001.hbase.singapore.rds.aliyuncs.com,hb-bp1w6krp539******-master3-001.hbase.singapore.rds.aliyuncs.com,hb-t4naqsay5gn******-master2-001.hbase.singapore.rds.aliyuncs.com:2181:/hbase
	HaStandbyClusterKey *string `json:"HaStandbyClusterKey,omitempty" xml:"HaStandbyClusterKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase
	HaStandbyDBType *string `json:"HaStandbyDBType,omitempty" xml:"HaStandbyDBType,omitempty"`
	// example:
	//
	// /hbase
	HaStandbyHbaseFsDir *string `json:"HaStandbyHbaseFsDir,omitempty" xml:"HaStandbyHbaseFsDir,omitempty"`
	// example:
	//
	// hdfs://hb-bp1w6krp539******-master1-001.hbase.rds.aliyuncs.com:8020,hb-bp1w6krp539******-master2-001.hbase.rds.aliyuncs.com:8020
	HaStandbyHdfsUri *string `json:"HaStandbyHdfsUri,omitempty" xml:"HaStandbyHdfsUri,omitempty"`
	// example:
	//
	// root
	HaStandbyPassword *string `json:"HaStandbyPassword,omitempty" xml:"HaStandbyPassword,omitempty"`
	// example:
	//
	// root
	HaStandbyUser *string `json:"HaStandbyUser,omitempty" xml:"HaStandbyUser,omitempty"`
	// example:
	//
	// HBase2x
	HaStandbyVersion *string `json:"HaStandbyVersion,omitempty" xml:"HaStandbyVersion,omitempty"`
	// example:
	//
	// test,test1
	HaTables *string `json:"HaTables,omitempty" xml:"HaTables,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	IsActiveStandard *bool `json:"IsActiveStandard,omitempty" xml:"IsActiveStandard,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	IsStandbyStandard *bool `json:"IsStandbyStandard,omitempty" xml:"IsStandbyStandard,omitempty"`
}

func (s RelateDbForHBaseHaRequest) String() string {
	return tea.Prettify(s)
}

func (s RelateDbForHBaseHaRequest) GoString() string {
	return s.String()
}

func (s *RelateDbForHBaseHaRequest) SetClusterId(v string) *RelateDbForHBaseHaRequest {
	s.ClusterId = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActive(v string) *RelateDbForHBaseHaRequest {
	s.HaActive = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveClusterKey(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveClusterKey = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveDBType(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveDBType = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveHbaseFsDir(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveHbaseFsDir = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveHdfsUri(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveHdfsUri = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActivePassword(v string) *RelateDbForHBaseHaRequest {
	s.HaActivePassword = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveUser(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveUser = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveVersion(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveVersion = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaMigrateType(v string) *RelateDbForHBaseHaRequest {
	s.HaMigrateType = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandby(v string) *RelateDbForHBaseHaRequest {
	s.HaStandby = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyClusterKey(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyClusterKey = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyDBType(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyDBType = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyHbaseFsDir(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyHbaseFsDir = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyHdfsUri(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyHdfsUri = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyPassword(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyPassword = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyUser(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyUser = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyVersion(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyVersion = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaTables(v string) *RelateDbForHBaseHaRequest {
	s.HaTables = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetIsActiveStandard(v bool) *RelateDbForHBaseHaRequest {
	s.IsActiveStandard = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetIsStandbyStandard(v bool) *RelateDbForHBaseHaRequest {
	s.IsStandbyStandard = &v
	return s
}

type RelateDbForHBaseHaResponseBody struct {
	// example:
	//
	// DC654531-0799-4502-AFA5-80EE1C16829A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RelateDbForHBaseHaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RelateDbForHBaseHaResponseBody) GoString() string {
	return s.String()
}

func (s *RelateDbForHBaseHaResponseBody) SetRequestId(v string) *RelateDbForHBaseHaResponseBody {
	s.RequestId = &v
	return s
}

type RelateDbForHBaseHaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RelateDbForHBaseHaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RelateDbForHBaseHaResponse) String() string {
	return tea.Prettify(s)
}

func (s RelateDbForHBaseHaResponse) GoString() string {
	return s.String()
}

func (s *RelateDbForHBaseHaResponse) SetHeaders(v map[string]*string) *RelateDbForHBaseHaResponse {
	s.Headers = v
	return s
}

func (s *RelateDbForHBaseHaResponse) SetStatusCode(v int32) *RelateDbForHBaseHaResponse {
	s.StatusCode = &v
	return s
}

func (s *RelateDbForHBaseHaResponse) SetBody(v *RelateDbForHBaseHaResponseBody) *RelateDbForHBaseHaResponse {
	s.Body = v
	return s
}

type ReleasePublicNetworkAddressRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ReleasePublicNetworkAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleasePublicNetworkAddressRequest) SetClusterId(v string) *ReleasePublicNetworkAddressRequest {
	s.ClusterId = &v
	return s
}

type ReleasePublicNetworkAddressResponseBody struct {
	// example:
	//
	// B18D4390-A968-4444-B323-4360B8E5DA3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleasePublicNetworkAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePublicNetworkAddressResponseBody) SetRequestId(v string) *ReleasePublicNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

type ReleasePublicNetworkAddressResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleasePublicNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleasePublicNetworkAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleasePublicNetworkAddressResponse) SetHeaders(v map[string]*string) *ReleasePublicNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleasePublicNetworkAddressResponse) SetStatusCode(v int32) *ReleasePublicNetworkAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleasePublicNetworkAddressResponse) SetBody(v *ReleasePublicNetworkAddressResponseBody) *ReleasePublicNetworkAddressResponse {
	s.Body = v
	return s
}

type RenewInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp1u0639js2h7****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) SetClusterId(v string) *RenewInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *RenewInstanceRequest) SetDuration(v int32) *RenewInstanceRequest {
	s.Duration = &v
	return s
}

func (s *RenewInstanceRequest) SetPricingCycle(v string) *RenewInstanceRequest {
	s.PricingCycle = &v
	return s
}

type RenewInstanceResponseBody struct {
	// example:
	//
	// 211235614240728
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 729CB2A7-3065-53A9-B27C-7033CA4881D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) SetOrderId(v int64) *RenewInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RenewInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponse) SetHeaders(v map[string]*string) *RenewInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewInstanceResponse) SetStatusCode(v int32) *RenewInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
	s.Body = v
	return s
}

type ResizeColdStorageSizeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp169l540vc6c****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 900
	ColdStorageSize *int32 `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
}

func (s ResizeColdStorageSizeRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeColdStorageSizeRequest) GoString() string {
	return s.String()
}

func (s *ResizeColdStorageSizeRequest) SetClusterId(v string) *ResizeColdStorageSizeRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeColdStorageSizeRequest) SetColdStorageSize(v int32) *ResizeColdStorageSizeRequest {
	s.ColdStorageSize = &v
	return s
}

type ResizeColdStorageSizeResponseBody struct {
	// example:
	//
	// 21711518427****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 5AA6F80E-535C-5611-BD13-3832D96A4D0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeColdStorageSizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeColdStorageSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeColdStorageSizeResponseBody) SetOrderId(v string) *ResizeColdStorageSizeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ResizeColdStorageSizeResponseBody) SetRequestId(v string) *ResizeColdStorageSizeResponseBody {
	s.RequestId = &v
	return s
}

type ResizeColdStorageSizeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeColdStorageSizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResizeColdStorageSizeResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeColdStorageSizeResponse) GoString() string {
	return s.String()
}

func (s *ResizeColdStorageSizeResponse) SetHeaders(v map[string]*string) *ResizeColdStorageSizeResponse {
	s.Headers = v
	return s
}

func (s *ResizeColdStorageSizeResponse) SetStatusCode(v int32) *ResizeColdStorageSizeResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeColdStorageSizeResponse) SetBody(v *ResizeColdStorageSizeResponseBody) *ResizeColdStorageSizeResponse {
	s.Body = v
	return s
}

type ResizeDiskSizeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16o0pd52e3y****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 120
	NodeDiskSize *int32 `json:"NodeDiskSize,omitempty" xml:"NodeDiskSize,omitempty"`
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

func (s *ResizeDiskSizeRequest) SetNodeDiskSize(v int32) *ResizeDiskSizeRequest {
	s.NodeDiskSize = &v
	return s
}

type ResizeDiskSizeResponseBody struct {
	// example:
	//
	// 3C22622B-8555-42BF-AD8A-1B960743****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 493A762B-E4A6-44E9-B877-CA6D0CAF8B29
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeDiskSizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeDiskSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeDiskSizeResponseBody) SetOrderId(v string) *ResizeDiskSizeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ResizeDiskSizeResponseBody) SetRequestId(v string) *ResizeDiskSizeResponseBody {
	s.RequestId = &v
	return s
}

type ResizeDiskSizeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeDiskSizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ResizeDiskSizeResponse) SetStatusCode(v int32) *ResizeDiskSizeResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeDiskSizeResponse) SetBody(v *ResizeDiskSizeResponseBody) *ResizeDiskSizeResponse {
	s.Body = v
	return s
}

type ResizeMultiZoneClusterDiskSizeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-f5d6vc2r8d6****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 480
	CoreDiskSize *int32 `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	// example:
	//
	// 440
	LogDiskSize *int32 `json:"LogDiskSize,omitempty" xml:"LogDiskSize,omitempty"`
}

func (s ResizeMultiZoneClusterDiskSizeRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeMultiZoneClusterDiskSizeRequest) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterDiskSizeRequest) SetClusterId(v string) *ResizeMultiZoneClusterDiskSizeRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeRequest) SetCoreDiskSize(v int32) *ResizeMultiZoneClusterDiskSizeRequest {
	s.CoreDiskSize = &v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeRequest) SetLogDiskSize(v int32) *ResizeMultiZoneClusterDiskSizeRequest {
	s.LogDiskSize = &v
	return s
}

type ResizeMultiZoneClusterDiskSizeResponseBody struct {
	// example:
	//
	// 123412341****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 568339C4-9F71-43D0-994E-E039CD826E56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeMultiZoneClusterDiskSizeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeMultiZoneClusterDiskSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterDiskSizeResponseBody) SetOrderId(v string) *ResizeMultiZoneClusterDiskSizeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeResponseBody) SetRequestId(v string) *ResizeMultiZoneClusterDiskSizeResponseBody {
	s.RequestId = &v
	return s
}

type ResizeMultiZoneClusterDiskSizeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeMultiZoneClusterDiskSizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResizeMultiZoneClusterDiskSizeResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeMultiZoneClusterDiskSizeResponse) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterDiskSizeResponse) SetHeaders(v map[string]*string) *ResizeMultiZoneClusterDiskSizeResponse {
	s.Headers = v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeResponse) SetStatusCode(v int32) *ResizeMultiZoneClusterDiskSizeResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeResponse) SetBody(v *ResizeMultiZoneClusterDiskSizeResponseBody) *ResizeMultiZoneClusterDiskSizeResponse {
	s.Body = v
	return s
}

type ResizeMultiZoneClusterNodeCountRequest struct {
	// example:
	//
	// vsw-hangxzhouxb*****
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-f5d8d6s4s2a1****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 6
	CoreNodeCount *int32 `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	// example:
	//
	// 8
	LogNodeCount *int32 `json:"LogNodeCount,omitempty" xml:"LogNodeCount,omitempty"`
	// example:
	//
	// 6
	PrimaryCoreNodeCount *int32 `json:"PrimaryCoreNodeCount,omitempty" xml:"PrimaryCoreNodeCount,omitempty"`
	// example:
	//
	// vsw-hangxzhouxe*****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// example:
	//
	// 6
	StandbyCoreNodeCount *int32 `json:"StandbyCoreNodeCount,omitempty" xml:"StandbyCoreNodeCount,omitempty"`
	// example:
	//
	// vsw-hangxzhouxf****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
}

func (s ResizeMultiZoneClusterNodeCountRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeMultiZoneClusterNodeCountRequest) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetArbiterVSwitchId(v string) *ResizeMultiZoneClusterNodeCountRequest {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetClusterId(v string) *ResizeMultiZoneClusterNodeCountRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetCoreNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest {
	s.CoreNodeCount = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetLogNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest {
	s.LogNodeCount = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetPrimaryCoreNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest {
	s.PrimaryCoreNodeCount = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetPrimaryVSwitchId(v string) *ResizeMultiZoneClusterNodeCountRequest {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetStandbyCoreNodeCount(v int32) *ResizeMultiZoneClusterNodeCountRequest {
	s.StandbyCoreNodeCount = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountRequest) SetStandbyVSwitchId(v string) *ResizeMultiZoneClusterNodeCountRequest {
	s.StandbyVSwitchId = &v
	return s
}

type ResizeMultiZoneClusterNodeCountResponseBody struct {
	// example:
	//
	// 1234123412****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// E2B7E9DA-1575-4B9D-A0E4-9468BAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeMultiZoneClusterNodeCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeMultiZoneClusterNodeCountResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterNodeCountResponseBody) SetOrderId(v string) *ResizeMultiZoneClusterNodeCountResponseBody {
	s.OrderId = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountResponseBody) SetRequestId(v string) *ResizeMultiZoneClusterNodeCountResponseBody {
	s.RequestId = &v
	return s
}

type ResizeMultiZoneClusterNodeCountResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeMultiZoneClusterNodeCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResizeMultiZoneClusterNodeCountResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeMultiZoneClusterNodeCountResponse) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterNodeCountResponse) SetHeaders(v map[string]*string) *ResizeMultiZoneClusterNodeCountResponse {
	s.Headers = v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountResponse) SetStatusCode(v int32) *ResizeMultiZoneClusterNodeCountResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountResponse) SetBody(v *ResizeMultiZoneClusterNodeCountResponseBody) *ResizeMultiZoneClusterNodeCountResponse {
	s.Body = v
	return s
}

type ResizeNodeCountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16o0pd52e3y****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// example:
	//
	// vsw-bp191otqj1ssyl****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *ResizeNodeCountRequest) SetNodeCount(v int32) *ResizeNodeCountRequest {
	s.NodeCount = &v
	return s
}

func (s *ResizeNodeCountRequest) SetVSwitchId(v string) *ResizeNodeCountRequest {
	s.VSwitchId = &v
	return s
}

func (s *ResizeNodeCountRequest) SetZoneId(v string) *ResizeNodeCountRequest {
	s.ZoneId = &v
	return s
}

type ResizeNodeCountResponseBody struct {
	// example:
	//
	// 20470860005****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// B288B41F-6681-42A6-8905-47C3C42B19B0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeNodeCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeNodeCountResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeNodeCountResponseBody) SetOrderId(v string) *ResizeNodeCountResponseBody {
	s.OrderId = &v
	return s
}

func (s *ResizeNodeCountResponseBody) SetRequestId(v string) *ResizeNodeCountResponseBody {
	s.RequestId = &v
	return s
}

type ResizeNodeCountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeNodeCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ResizeNodeCountResponse) SetStatusCode(v int32) *ResizeNodeCountResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeNodeCountResponse) SetBody(v *ResizeNodeCountResponseBody) *ResizeNodeCountResponse {
	s.Body = v
	return s
}

type RestartInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// THRIFT
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
}

func (s RestartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceRequest) SetClusterId(v string) *RestartInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *RestartInstanceRequest) SetComponents(v string) *RestartInstanceRequest {
	s.Components = &v
	return s
}

type RestartInstanceResponseBody struct {
	// example:
	//
	// F744E939-D08D-5623-82C8-9D1F9F7685D1
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

type RevokeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test01
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// READ,WRITE
	AclActions *string `json:"AclActions,omitempty" xml:"AclActions,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s RevokeRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeRequest) GoString() string {
	return s.String()
}

func (s *RevokeRequest) SetAccountName(v string) *RevokeRequest {
	s.AccountName = &v
	return s
}

func (s *RevokeRequest) SetAclActions(v string) *RevokeRequest {
	s.AclActions = &v
	return s
}

func (s *RevokeRequest) SetClusterId(v string) *RevokeRequest {
	s.ClusterId = &v
	return s
}

func (s *RevokeRequest) SetNamespace(v string) *RevokeRequest {
	s.Namespace = &v
	return s
}

func (s *RevokeRequest) SetTableName(v string) *RevokeRequest {
	s.TableName = &v
	return s
}

type RevokeResponseBody struct {
	// example:
	//
	// C9085433-A56A-4089-B49A-DF5A4E2B7B06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeResponseBody) SetRequestId(v string) *RevokeResponseBody {
	s.RequestId = &v
	return s
}

type RevokeResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeResponse) GoString() string {
	return s.String()
}

func (s *RevokeResponse) SetHeaders(v map[string]*string) *RevokeResponse {
	s.Headers = v
	return s
}

func (s *RevokeResponse) SetStatusCode(v int32) *RevokeResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeResponse) SetBody(v *RevokeResponseBody) *RevokeResponse {
	s.Body = v
	return s
}

type SwitchHbaseHaSlbRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bds-t4n3496whj23ia4k
	BdsId *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ha-v21tmnxjwh2yuy1il
	HaId *string `json:"HaId,omitempty" xml:"HaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// thrift
	HaTypes *string `json:"HaTypes,omitempty" xml:"HaTypes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Active
	HbaseType *string `json:"HbaseType,omitempty" xml:"HbaseType,omitempty"`
}

func (s SwitchHbaseHaSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchHbaseHaSlbRequest) GoString() string {
	return s.String()
}

func (s *SwitchHbaseHaSlbRequest) SetBdsId(v string) *SwitchHbaseHaSlbRequest {
	s.BdsId = &v
	return s
}

func (s *SwitchHbaseHaSlbRequest) SetHaId(v string) *SwitchHbaseHaSlbRequest {
	s.HaId = &v
	return s
}

func (s *SwitchHbaseHaSlbRequest) SetHaTypes(v string) *SwitchHbaseHaSlbRequest {
	s.HaTypes = &v
	return s
}

func (s *SwitchHbaseHaSlbRequest) SetHbaseType(v string) *SwitchHbaseHaSlbRequest {
	s.HbaseType = &v
	return s
}

type SwitchHbaseHaSlbResponseBody struct {
	// example:
	//
	// C9D568D9-A59C-4AF2-8FBB-F086A841D58E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchHbaseHaSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchHbaseHaSlbResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchHbaseHaSlbResponseBody) SetRequestId(v string) *SwitchHbaseHaSlbResponseBody {
	s.RequestId = &v
	return s
}

type SwitchHbaseHaSlbResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchHbaseHaSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchHbaseHaSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchHbaseHaSlbResponse) GoString() string {
	return s.String()
}

func (s *SwitchHbaseHaSlbResponse) SetHeaders(v map[string]*string) *SwitchHbaseHaSlbResponse {
	s.Headers = v
	return s
}

func (s *SwitchHbaseHaSlbResponse) SetStatusCode(v int32) *SwitchHbaseHaSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchHbaseHaSlbResponse) SetBody(v *SwitchHbaseHaSlbResponseBody) *SwitchHbaseHaSlbResponse {
	s.Body = v
	return s
}

type SwitchServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-uf6r2hn2zrxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// open
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HBaseProxy
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s SwitchServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchServiceRequest) GoString() string {
	return s.String()
}

func (s *SwitchServiceRequest) SetClusterId(v string) *SwitchServiceRequest {
	s.ClusterId = &v
	return s
}

func (s *SwitchServiceRequest) SetOperate(v string) *SwitchServiceRequest {
	s.Operate = &v
	return s
}

func (s *SwitchServiceRequest) SetServiceName(v string) *SwitchServiceRequest {
	s.ServiceName = &v
	return s
}

type SwitchServiceResponseBody struct {
	// example:
	//
	// F1005DE4-D981-559F-9E37-5172DXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchServiceResponseBody) SetRequestId(v string) *SwitchServiceResponseBody {
	s.RequestId = &v
	return s
}

type SwitchServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchServiceResponse) GoString() string {
	return s.String()
}

func (s *SwitchServiceResponse) SetHeaders(v map[string]*string) *SwitchServiceResponse {
	s.Headers = v
	return s
}

func (s *SwitchServiceResponse) SetStatusCode(v int32) *SwitchServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchServiceResponse) SetBody(v *SwitchServiceResponseBody) *SwitchServiceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bds-bp15e022622fk0w1
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
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
	// example:
	//
	// 2656FA19-6059-40C8-A157-3FFBEAEC2369
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UnTagResourcesRequest struct {
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bds-bp15e022622fk0w1
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// example:
	//
	// key1
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetRegionId(v string) *UnTagResourcesRequest {
	s.RegionId = &v
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
	// example:
	//
	// 9CBF8DF0-4931-4A54-9B60-4C6E1AB59286
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UnTagResourcesResponse) SetStatusCode(v int32) *UnTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UnTagResourcesResponse) SetBody(v *UnTagResourcesResponseBody) *UnTagResourcesResponse {
	s.Body = v
	return s
}

type UpgradeMinorVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// HADOOP
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
}

func (s UpgradeMinorVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMinorVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionRequest) SetClusterId(v string) *UpgradeMinorVersionRequest {
	s.ClusterId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetComponents(v string) *UpgradeMinorVersionRequest {
	s.Components = &v
	return s
}

type UpgradeMinorVersionResponseBody struct {
	// example:
	//
	// 7B8EC240-BB13-4DBC-B955-F90170E82609
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// HADOOP
	UpgradingComponents *string `json:"UpgradingComponents,omitempty" xml:"UpgradingComponents,omitempty"`
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

func (s *UpgradeMinorVersionResponseBody) SetUpgradingComponents(v string) *UpgradeMinorVersionResponseBody {
	s.UpgradingComponents = &v
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

type UpgradeMultiZoneClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-***************
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// LINDORM
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// example:
	//
	// LPROXY
	RestartComponents *string `json:"RestartComponents,omitempty" xml:"RestartComponents,omitempty"`
	// example:
	//
	// serial
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// example:
	//
	// ld-t4n40m3171t4******-az-b
	UpgradeInsName *string `json:"UpgradeInsName,omitempty" xml:"UpgradeInsName,omitempty"`
	// example:
	//
	// t-apsara-lindorm-2.1.20-20200518175539.alios7.x86_64
	Versions *string `json:"Versions,omitempty" xml:"Versions,omitempty"`
}

func (s UpgradeMultiZoneClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMultiZoneClusterRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMultiZoneClusterRequest) SetClusterId(v string) *UpgradeMultiZoneClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) SetComponents(v string) *UpgradeMultiZoneClusterRequest {
	s.Components = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) SetRestartComponents(v string) *UpgradeMultiZoneClusterRequest {
	s.RestartComponents = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) SetRunMode(v string) *UpgradeMultiZoneClusterRequest {
	s.RunMode = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) SetUpgradeInsName(v string) *UpgradeMultiZoneClusterRequest {
	s.UpgradeInsName = &v
	return s
}

func (s *UpgradeMultiZoneClusterRequest) SetVersions(v string) *UpgradeMultiZoneClusterRequest {
	s.Versions = &v
	return s
}

type UpgradeMultiZoneClusterResponseBody struct {
	// example:
	//
	// C532A4D4-9451-4460-BB3E-300FEC852D3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// LINDORM
	UpgradingComponents *string `json:"UpgradingComponents,omitempty" xml:"UpgradingComponents,omitempty"`
}

func (s UpgradeMultiZoneClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMultiZoneClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeMultiZoneClusterResponseBody) SetRequestId(v string) *UpgradeMultiZoneClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeMultiZoneClusterResponseBody) SetUpgradingComponents(v string) *UpgradeMultiZoneClusterResponseBody {
	s.UpgradingComponents = &v
	return s
}

type UpgradeMultiZoneClusterResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeMultiZoneClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeMultiZoneClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeMultiZoneClusterResponse) GoString() string {
	return s.String()
}

func (s *UpgradeMultiZoneClusterResponse) SetHeaders(v map[string]*string) *UpgradeMultiZoneClusterResponse {
	s.Headers = v
	return s
}

func (s *UpgradeMultiZoneClusterResponse) SetStatusCode(v int32) *UpgradeMultiZoneClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeMultiZoneClusterResponse) SetBody(v *UpgradeMultiZoneClusterResponseBody) *UpgradeMultiZoneClusterResponse {
	s.Body = v
	return s
}

type XpackRelateDBRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ap-bp1qtz9rcbbt3****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16o0pd52e3****
	DbClusterIds *string `json:"DbClusterIds,omitempty" xml:"DbClusterIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hbase
	RelateDbType *string `json:"RelateDbType,omitempty" xml:"RelateDbType,omitempty"`
}

func (s XpackRelateDBRequest) String() string {
	return tea.Prettify(s)
}

func (s XpackRelateDBRequest) GoString() string {
	return s.String()
}

func (s *XpackRelateDBRequest) SetClusterId(v string) *XpackRelateDBRequest {
	s.ClusterId = &v
	return s
}

func (s *XpackRelateDBRequest) SetDbClusterIds(v string) *XpackRelateDBRequest {
	s.DbClusterIds = &v
	return s
}

func (s *XpackRelateDBRequest) SetRelateDbType(v string) *XpackRelateDBRequest {
	s.RelateDbType = &v
	return s
}

type XpackRelateDBResponseBody struct {
	// example:
	//
	// 50373857-C47B-4B64-9332-D0B5280B59EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s XpackRelateDBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s XpackRelateDBResponseBody) GoString() string {
	return s.String()
}

func (s *XpackRelateDBResponseBody) SetRequestId(v string) *XpackRelateDBResponseBody {
	s.RequestId = &v
	return s
}

type XpackRelateDBResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *XpackRelateDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s XpackRelateDBResponse) String() string {
	return tea.Prettify(s)
}

func (s XpackRelateDBResponse) GoString() string {
	return s.String()
}

func (s *XpackRelateDBResponse) SetHeaders(v map[string]*string) *XpackRelateDBResponse {
	s.Headers = v
	return s
}

func (s *XpackRelateDBResponse) SetStatusCode(v int32) *XpackRelateDBResponse {
	s.StatusCode = &v
	return s
}

func (s *XpackRelateDBResponse) SetBody(v *XpackRelateDBResponseBody) *XpackRelateDBResponse {
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
		"ap-southeast-1":        tea.String("hbase.aliyuncs.com"),
		"cn-beijing":            tea.String("hbase.aliyuncs.com"),
		"cn-hangzhou":           tea.String("hbase.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("hbase.aliyuncs.com"),
		"cn-hongkong":           tea.String("hbase.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("hbase.aliyuncs.com"),
		"cn-qingdao":            tea.String("hbase.aliyuncs.com"),
		"cn-shanghai":           tea.String("hbase.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("hbase.aliyuncs.com"),
		"cn-shenzhen":           tea.String("hbase.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("hbase.aliyuncs.com"),
		"cn-guangzhou":          tea.String("hbase.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("hbase"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AddUserHdfsInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserHdfsInfoResponse
func (client *Client) AddUserHdfsInfoWithOptions(request *AddUserHdfsInfoRequest, runtime *util.RuntimeOptions) (_result *AddUserHdfsInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		query["ExtInfo"] = request.ExtInfo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserHdfsInfo"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserHdfsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddUserHdfsInfoRequest
//
// @return AddUserHdfsInfoResponse
func (client *Client) AddUserHdfsInfo(request *AddUserHdfsInfoRequest) (_result *AddUserHdfsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserHdfsInfoResponse{}
	_body, _err := client.AddUserHdfsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AllocatePublicNetworkAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocatePublicNetworkAddressResponse
func (client *Client) AllocatePublicNetworkAddressWithOptions(request *AllocatePublicNetworkAddressRequest, runtime *util.RuntimeOptions) (_result *AllocatePublicNetworkAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AllocatePublicNetworkAddress"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AllocatePublicNetworkAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AllocatePublicNetworkAddressRequest
//
// @return AllocatePublicNetworkAddressResponse
func (client *Client) AllocatePublicNetworkAddress(request *AllocatePublicNetworkAddressRequest) (_result *AllocatePublicNetworkAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocatePublicNetworkAddressResponse{}
	_body, _err := client.AllocatePublicNetworkAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CancelActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelActiveOperationTasksResponse
func (client *Client) CancelActiveOperationTasksWithOptions(request *CancelActiveOperationTasksRequest, runtime *util.RuntimeOptions) (_result *CancelActiveOperationTasksResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelActiveOperationTasks"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelActiveOperationTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CancelActiveOperationTasksRequest
//
// @return CancelActiveOperationTasksResponse
func (client *Client) CancelActiveOperationTasks(request *CancelActiveOperationTasksRequest) (_result *CancelActiveOperationTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelActiveOperationTasksResponse{}
	_body, _err := client.CancelActiveOperationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CheckComponentsVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckComponentsVersionResponse
func (client *Client) CheckComponentsVersionWithOptions(request *CheckComponentsVersionRequest, runtime *util.RuntimeOptions) (_result *CheckComponentsVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Components)) {
		query["Components"] = request.Components
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckComponentsVersion"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckComponentsVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckComponentsVersionRequest
//
// @return CheckComponentsVersionResponse
func (client *Client) CheckComponentsVersion(request *CheckComponentsVersionRequest) (_result *CheckComponentsVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckComponentsVersionResponse{}
	_body, _err := client.CheckComponentsVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CloseBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseBackupResponse
func (client *Client) CloseBackupWithOptions(request *CloseBackupRequest, runtime *util.RuntimeOptions) (_result *CloseBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseBackup"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CloseBackupRequest
//
// @return CloseBackupResponse
func (client *Client) CloseBackup(request *CloseBackupRequest) (_result *CloseBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseBackupResponse{}
	_body, _err := client.CloseBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ConvertInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertInstanceResponse
func (client *Client) ConvertInstanceWithOptions(request *ConvertInstanceRequest, runtime *util.RuntimeOptions) (_result *ConvertInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConvertInstance"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConvertInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ConvertInstanceRequest
//
// @return ConvertInstanceResponse
func (client *Client) ConvertInstance(request *ConvertInstanceRequest) (_result *ConvertInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConvertInstanceResponse{}
	_body, _err := client.ConvertInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建账户
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
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccount"),
		Version:     tea.String("2019-01-01"),
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
// 新建账户
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

// @param request - CreateBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackupPlanResponse
func (client *Client) CreateBackupPlanWithOptions(request *CreateBackupPlanRequest, runtime *util.RuntimeOptions) (_result *CreateBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBackupPlan"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateBackupPlanRequest
//
// @return CreateBackupPlanResponse
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

// @param request - CreateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.ColdStorageSize)) {
		query["ColdStorageSize"] = request.ColdStorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.CoreInstanceType)) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.DiskSize)) {
		query["DiskSize"] = request.DiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.DiskType)) {
		query["DiskType"] = request.DiskType
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptionKey)) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.EngineVersion)) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.MasterInstanceType)) {
		query["MasterInstanceType"] = request.MasterInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeCount)) {
		query["NodeCount"] = request.NodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIPList)) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateClusterRequest
//
// @return CreateClusterResponse
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

// @param request - CreateGlobalResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGlobalResourceResponse
func (client *Client) CreateGlobalResourceWithOptions(request *CreateGlobalResourceRequest, runtime *util.RuntimeOptions) (_result *CreateGlobalResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGlobalResource"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGlobalResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateGlobalResourceRequest
//
// @return CreateGlobalResourceResponse
func (client *Client) CreateGlobalResource(request *CreateGlobalResourceRequest) (_result *CreateGlobalResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGlobalResourceResponse{}
	_body, _err := client.CreateGlobalResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateHBaseSlbServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHBaseSlbServerResponse
func (client *Client) CreateHBaseSlbServerWithOptions(request *CreateHBaseSlbServerRequest, runtime *util.RuntimeOptions) (_result *CreateHBaseSlbServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.SlbServer)) {
		query["SlbServer"] = request.SlbServer
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHBaseSlbServer"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHBaseSlbServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateHBaseSlbServerRequest
//
// @return CreateHBaseSlbServerResponse
func (client *Client) CreateHBaseSlbServer(request *CreateHBaseSlbServerRequest) (_result *CreateHBaseSlbServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHBaseSlbServerResponse{}
	_body, _err := client.CreateHBaseSlbServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateHbaseHaSlbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHbaseHaSlbResponse
func (client *Client) CreateHbaseHaSlbWithOptions(request *CreateHbaseHaSlbRequest, runtime *util.RuntimeOptions) (_result *CreateHbaseHaSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BdsId)) {
		query["BdsId"] = request.BdsId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.HaId)) {
		query["HaId"] = request.HaId
	}

	if !tea.BoolValue(util.IsUnset(request.HaTypes)) {
		query["HaTypes"] = request.HaTypes
	}

	if !tea.BoolValue(util.IsUnset(request.HbaseType)) {
		query["HbaseType"] = request.HbaseType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHbaseHaSlb"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHbaseHaSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateHbaseHaSlbRequest
//
// @return CreateHbaseHaSlbResponse
func (client *Client) CreateHbaseHaSlb(request *CreateHbaseHaSlbRequest) (_result *CreateHbaseHaSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHbaseHaSlbResponse{}
	_body, _err := client.CreateHbaseHaSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateMultiZoneClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultiZoneClusterResponse
func (client *Client) CreateMultiZoneClusterWithOptions(request *CreateMultiZoneClusterRequest, runtime *util.RuntimeOptions) (_result *CreateMultiZoneClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArbiterVSwitchId)) {
		query["ArbiterVSwitchId"] = request.ArbiterVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ArbiterZoneId)) {
		query["ArbiterZoneId"] = request.ArbiterZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.ArchVersion)) {
		query["ArchVersion"] = request.ArchVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.CoreDiskSize)) {
		query["CoreDiskSize"] = request.CoreDiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.CoreDiskType)) {
		query["CoreDiskType"] = request.CoreDiskType
	}

	if !tea.BoolValue(util.IsUnset(request.CoreInstanceType)) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.CoreNodeCount)) {
		query["CoreNodeCount"] = request.CoreNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.EngineVersion)) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.LogDiskSize)) {
		query["LogDiskSize"] = request.LogDiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.LogDiskType)) {
		query["LogDiskType"] = request.LogDiskType
	}

	if !tea.BoolValue(util.IsUnset(request.LogInstanceType)) {
		query["LogInstanceType"] = request.LogInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.LogNodeCount)) {
		query["LogNodeCount"] = request.LogNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.MasterInstanceType)) {
		query["MasterInstanceType"] = request.MasterInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.MultiZoneCombination)) {
		query["MultiZoneCombination"] = request.MultiZoneCombination
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryVSwitchId)) {
		query["PrimaryVSwitchId"] = request.PrimaryVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryZoneId)) {
		query["PrimaryZoneId"] = request.PrimaryZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIPList)) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !tea.BoolValue(util.IsUnset(request.StandbyVSwitchId)) {
		query["StandbyVSwitchId"] = request.StandbyVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.StandbyZoneId)) {
		query["StandbyZoneId"] = request.StandbyZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMultiZoneCluster"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMultiZoneClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMultiZoneClusterRequest
//
// @return CreateMultiZoneClusterResponse
func (client *Client) CreateMultiZoneCluster(request *CreateMultiZoneClusterRequest) (_result *CreateMultiZoneClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMultiZoneClusterResponse{}
	_body, _err := client.CreateMultiZoneClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateRestorePlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRestorePlanResponse
func (client *Client) CreateRestorePlanWithOptions(request *CreateRestorePlanRequest, runtime *util.RuntimeOptions) (_result *CreateRestorePlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreAllTable)) {
		query["RestoreAllTable"] = request.RestoreAllTable
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreByCopy)) {
		query["RestoreByCopy"] = request.RestoreByCopy
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreToDate)) {
		query["RestoreToDate"] = request.RestoreToDate
	}

	if !tea.BoolValue(util.IsUnset(request.Tables)) {
		query["Tables"] = request.Tables
	}

	if !tea.BoolValue(util.IsUnset(request.TargetClusterId)) {
		query["TargetClusterId"] = request.TargetClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRestorePlan"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRestorePlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateRestorePlanRequest
//
// @return CreateRestorePlanResponse
func (client *Client) CreateRestorePlan(request *CreateRestorePlanRequest) (_result *CreateRestorePlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRestorePlanResponse{}
	_body, _err := client.CreateRestorePlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateServerlessClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServerlessClusterResponse
func (client *Client) CreateServerlessClusterWithOptions(request *CreateServerlessClusterRequest, runtime *util.RuntimeOptions) (_result *CreateServerlessClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.DiskType)) {
		query["DiskType"] = request.DiskType
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.EngineVersion)) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerlessCapability)) {
		query["ServerlessCapability"] = request.ServerlessCapability
	}

	if !tea.BoolValue(util.IsUnset(request.ServerlessSpec)) {
		query["ServerlessSpec"] = request.ServerlessSpec
	}

	if !tea.BoolValue(util.IsUnset(request.ServerlessStorage)) {
		query["ServerlessStorage"] = request.ServerlessStorage
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServerlessCluster"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServerlessClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateServerlessClusterRequest
//
// @return CreateServerlessClusterResponse
func (client *Client) CreateServerlessCluster(request *CreateServerlessClusterRequest) (_result *CreateServerlessClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServerlessClusterResponse{}
	_body, _err := client.CreateServerlessClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除账户
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

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccount"),
		Version:     tea.String("2019-01-01"),
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
// 删除账户
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

// @param request - DeleteGlobalResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGlobalResourceResponse
func (client *Client) DeleteGlobalResourceWithOptions(request *DeleteGlobalResourceRequest, runtime *util.RuntimeOptions) (_result *DeleteGlobalResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGlobalResource"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGlobalResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteGlobalResourceRequest
//
// @return DeleteGlobalResourceResponse
func (client *Client) DeleteGlobalResource(request *DeleteGlobalResourceRequest) (_result *DeleteGlobalResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGlobalResourceResponse{}
	_body, _err := client.DeleteGlobalResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteHBaseHaDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHBaseHaDBResponse
func (client *Client) DeleteHBaseHaDBWithOptions(request *DeleteHBaseHaDBRequest, runtime *util.RuntimeOptions) (_result *DeleteHBaseHaDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BdsId)) {
		query["BdsId"] = request.BdsId
	}

	if !tea.BoolValue(util.IsUnset(request.HaId)) {
		query["HaId"] = request.HaId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHBaseHaDB"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHBaseHaDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteHBaseHaDBRequest
//
// @return DeleteHBaseHaDBResponse
func (client *Client) DeleteHBaseHaDB(request *DeleteHBaseHaDBRequest) (_result *DeleteHBaseHaDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHBaseHaDBResponse{}
	_body, _err := client.DeleteHBaseHaDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteHBaseSlbServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHBaseSlbServerResponse
func (client *Client) DeleteHBaseSlbServerWithOptions(request *DeleteHBaseSlbServerRequest, runtime *util.RuntimeOptions) (_result *DeleteHBaseSlbServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.SlbServer)) {
		query["SlbServer"] = request.SlbServer
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHBaseSlbServer"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHBaseSlbServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteHBaseSlbServerRequest
//
// @return DeleteHBaseSlbServerResponse
func (client *Client) DeleteHBaseSlbServer(request *DeleteHBaseSlbServerRequest) (_result *DeleteHBaseSlbServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHBaseSlbServerResponse{}
	_body, _err := client.DeleteHBaseSlbServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteHbaseHaSlbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHbaseHaSlbResponse
func (client *Client) DeleteHbaseHaSlbWithOptions(request *DeleteHbaseHaSlbRequest, runtime *util.RuntimeOptions) (_result *DeleteHbaseHaSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BdsId)) {
		query["BdsId"] = request.BdsId
	}

	if !tea.BoolValue(util.IsUnset(request.HaId)) {
		query["HaId"] = request.HaId
	}

	if !tea.BoolValue(util.IsUnset(request.HaTypes)) {
		query["HaTypes"] = request.HaTypes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHbaseHaSlb"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHbaseHaSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteHbaseHaSlbRequest
//
// @return DeleteHbaseHaSlbResponse
func (client *Client) DeleteHbaseHaSlb(request *DeleteHbaseHaSlbRequest) (_result *DeleteHbaseHaSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHbaseHaSlbResponse{}
	_body, _err := client.DeleteHbaseHaSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ImmediateDeleteFlag)) {
		query["ImmediateDeleteFlag"] = request.ImmediateDeleteFlag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
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

// @param request - DeleteMultiZoneClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMultiZoneClusterResponse
func (client *Client) DeleteMultiZoneClusterWithOptions(request *DeleteMultiZoneClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteMultiZoneClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ImmediateDeleteFlag)) {
		query["ImmediateDeleteFlag"] = request.ImmediateDeleteFlag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMultiZoneCluster"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMultiZoneClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMultiZoneClusterRequest
//
// @return DeleteMultiZoneClusterResponse
func (client *Client) DeleteMultiZoneCluster(request *DeleteMultiZoneClusterRequest) (_result *DeleteMultiZoneClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMultiZoneClusterResponse{}
	_body, _err := client.DeleteMultiZoneClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteServerlessClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServerlessClusterResponse
func (client *Client) DeleteServerlessClusterWithOptions(request *DeleteServerlessClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteServerlessClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServerlessCluster"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServerlessClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteServerlessClusterRequest
//
// @return DeleteServerlessClusterResponse
func (client *Client) DeleteServerlessCluster(request *DeleteServerlessClusterRequest) (_result *DeleteServerlessClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServerlessClusterResponse{}
	_body, _err := client.DeleteServerlessClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteUserHdfsInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserHdfsInfoResponse
func (client *Client) DeleteUserHdfsInfoWithOptions(request *DeleteUserHdfsInfoRequest, runtime *util.RuntimeOptions) (_result *DeleteUserHdfsInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NameService)) {
		query["NameService"] = request.NameService
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserHdfsInfo"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserHdfsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteUserHdfsInfoRequest
//
// @return DeleteUserHdfsInfoResponse
func (client *Client) DeleteUserHdfsInfo(request *DeleteUserHdfsInfoRequest) (_result *DeleteUserHdfsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserHdfsInfoResponse{}
	_body, _err := client.DeleteUserHdfsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询账户列表
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
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccounts"),
		Version:     tea.String("2019-01-01"),
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
// 查询账户列表
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

// @param request - DescribeActiveOperationTaskTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationTaskTypeResponse
func (client *Client) DescribeActiveOperationTaskTypeWithOptions(request *DescribeActiveOperationTaskTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeActiveOperationTaskTypeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeActiveOperationTaskType"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeActiveOperationTaskTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeActiveOperationTaskTypeRequest
//
// @return DescribeActiveOperationTaskTypeResponse
func (client *Client) DescribeActiveOperationTaskType(request *DescribeActiveOperationTaskTypeRequest) (_result *DescribeActiveOperationTaskTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeActiveOperationTaskTypeResponse{}
	_body, _err := client.DescribeActiveOperationTaskTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationTasksResponse
func (client *Client) DescribeActiveOperationTasksWithOptions(request *DescribeActiveOperationTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeActiveOperationTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowCancel)) {
		query["AllowCancel"] = request.AllowCancel
	}

	if !tea.BoolValue(util.IsUnset(request.AllowChange)) {
		query["AllowChange"] = request.AllowChange
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeLevel)) {
		query["ChangeLevel"] = request.ChangeLevel
	}

	if !tea.BoolValue(util.IsUnset(request.DbType)) {
		query["DbType"] = request.DbType
	}

	if !tea.BoolValue(util.IsUnset(request.InsName)) {
		query["InsName"] = request.InsName
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

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
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

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeActiveOperationTasks"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeActiveOperationTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeActiveOperationTasksRequest
//
// @return DescribeActiveOperationTasksResponse
func (client *Client) DescribeActiveOperationTasks(request *DescribeActiveOperationTasksRequest) (_result *DescribeActiveOperationTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeActiveOperationTasksResponse{}
	_body, _err := client.DescribeActiveOperationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAvailableResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableResourceResponse
func (client *Client) DescribeAvailableResourceWithOptions(request *DescribeAvailableResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.CoreInstanceType)) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.DiskType)) {
		query["DiskType"] = request.DiskType
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.EngineVersion)) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAvailableResource"),
		Version:     tea.String("2019-01-01"),
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

// @param request - DescribeAvailableResourceRequest
//
// @return DescribeAvailableResourceResponse
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

// @param request - DescribeBackupPlanConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPlanConfigResponse
func (client *Client) DescribeBackupPlanConfigWithOptions(request *DescribeBackupPlanConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPlanConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupPlanConfig"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupPlanConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeBackupPlanConfigRequest
//
// @return DescribeBackupPlanConfigResponse
func (client *Client) DescribeBackupPlanConfig(request *DescribeBackupPlanConfigRequest) (_result *DescribeBackupPlanConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPlanConfigResponse{}
	_body, _err := client.DescribeBackupPlanConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupPolicy"),
		Version:     tea.String("2019-01-01"),
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

// @param request - DescribeBackupStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupStatusResponse
func (client *Client) DescribeBackupStatusWithOptions(request *DescribeBackupStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupStatus"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeBackupStatusRequest
//
// @return DescribeBackupStatusResponse
func (client *Client) DescribeBackupStatus(request *DescribeBackupStatusRequest) (_result *DescribeBackupStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupStatusResponse{}
	_body, _err := client.DescribeBackupStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeBackupSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupSummaryResponse
func (client *Client) DescribeBackupSummaryWithOptions(request *DescribeBackupSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupSummary"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeBackupSummaryRequest
//
// @return DescribeBackupSummaryResponse
func (client *Client) DescribeBackupSummary(request *DescribeBackupSummaryRequest) (_result *DescribeBackupSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupSummaryResponse{}
	_body, _err := client.DescribeBackupSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeBackupTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupTablesResponse
func (client *Client) DescribeBackupTablesWithOptions(request *DescribeBackupTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupRecordId)) {
		query["BackupRecordId"] = request.BackupRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupTables"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeBackupTablesRequest
//
// @return DescribeBackupTablesResponse
func (client *Client) DescribeBackupTables(request *DescribeBackupTablesRequest) (_result *DescribeBackupTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupTablesResponse{}
	_body, _err := client.DescribeBackupTablesWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeUTC)) {
		query["EndTimeUTC"] = request.EndTimeUTC
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeUTC)) {
		query["StartTimeUTC"] = request.StartTimeUTC
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackups"),
		Version:     tea.String("2019-01-01"),
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

// @param request - DescribeClusterConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterConnectionResponse
func (client *Client) DescribeClusterConnectionWithOptions(request *DescribeClusterConnectionRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeClusterConnection"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeClusterConnectionRequest
//
// @return DescribeClusterConnectionResponse
func (client *Client) DescribeClusterConnection(request *DescribeClusterConnectionRequest) (_result *DescribeClusterConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterConnectionResponse{}
	_body, _err := client.DescribeClusterConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeColdStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColdStorageResponse
func (client *Client) DescribeColdStorageWithOptions(request *DescribeColdStorageRequest, runtime *util.RuntimeOptions) (_result *DescribeColdStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeColdStorage"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeColdStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeColdStorageRequest
//
// @return DescribeColdStorageResponse
func (client *Client) DescribeColdStorage(request *DescribeColdStorageRequest) (_result *DescribeColdStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeColdStorageResponse{}
	_body, _err := client.DescribeColdStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDBInstanceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceUsageResponse
func (client *Client) DescribeDBInstanceUsageWithOptions(request *DescribeDBInstanceUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBInstanceUsage"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBInstanceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDBInstanceUsageRequest
//
// @return DescribeDBInstanceUsageResponse
func (client *Client) DescribeDBInstanceUsage(request *DescribeDBInstanceUsageRequest) (_result *DescribeDBInstanceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceUsageResponse{}
	_body, _err := client.DescribeDBInstanceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDeletedInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeletedInstancesResponse
func (client *Client) DescribeDeletedInstancesWithOptions(request *DescribeDeletedInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDeletedInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("DescribeDeletedInstances"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeletedInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDeletedInstancesRequest
//
// @return DescribeDeletedInstancesResponse
func (client *Client) DescribeDeletedInstances(request *DescribeDeletedInstancesRequest) (_result *DescribeDeletedInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeletedInstancesResponse{}
	_body, _err := client.DescribeDeletedInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDiskWarningLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiskWarningLineResponse
func (client *Client) DescribeDiskWarningLineWithOptions(request *DescribeDiskWarningLineRequest, runtime *util.RuntimeOptions) (_result *DescribeDiskWarningLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiskWarningLine"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiskWarningLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDiskWarningLineRequest
//
// @return DescribeDiskWarningLineResponse
func (client *Client) DescribeDiskWarningLine(request *DescribeDiskWarningLineRequest) (_result *DescribeDiskWarningLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiskWarningLineResponse{}
	_body, _err := client.DescribeDiskWarningLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEndpointsResponse
func (client *Client) DescribeEndpointsWithOptions(request *DescribeEndpointsRequest, runtime *util.RuntimeOptions) (_result *DescribeEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEndpoints"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeEndpointsRequest
//
// @return DescribeEndpointsResponse
func (client *Client) DescribeEndpoints(request *DescribeEndpointsRequest) (_result *DescribeEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEndpointsResponse{}
	_body, _err := client.DescribeEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstance"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeInstanceRequest
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeInstanceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceTypeResponse
func (client *Client) DescribeInstanceTypeWithOptions(request *DescribeInstanceTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceType"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeInstanceTypeRequest
//
// @return DescribeInstanceTypeResponse
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

// @param request - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.DbType)) {
		query["DbType"] = request.DbType
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

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstances"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeInstancesRequest
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeIpWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpWhitelistResponse
func (client *Client) DescribeIpWhitelistWithOptions(request *DescribeIpWhitelistRequest, runtime *util.RuntimeOptions) (_result *DescribeIpWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIpWhitelist"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIpWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeIpWhitelistRequest
//
// @return DescribeIpWhitelistResponse
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

// @param request - DescribeMultiZoneAvailableRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMultiZoneAvailableRegionsResponse
func (client *Client) DescribeMultiZoneAvailableRegionsWithOptions(request *DescribeMultiZoneAvailableRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeMultiZoneAvailableRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMultiZoneAvailableRegions"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMultiZoneAvailableRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeMultiZoneAvailableRegionsRequest
//
// @return DescribeMultiZoneAvailableRegionsResponse
func (client *Client) DescribeMultiZoneAvailableRegions(request *DescribeMultiZoneAvailableRegionsRequest) (_result *DescribeMultiZoneAvailableRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMultiZoneAvailableRegionsResponse{}
	_body, _err := client.DescribeMultiZoneAvailableRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeMultiZoneAvailableResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMultiZoneAvailableResourceResponse
func (client *Client) DescribeMultiZoneAvailableResourceWithOptions(request *DescribeMultiZoneAvailableResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeMultiZoneAvailableResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneCombination)) {
		query["ZoneCombination"] = request.ZoneCombination
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMultiZoneAvailableResource"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMultiZoneAvailableResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeMultiZoneAvailableResourceRequest
//
// @return DescribeMultiZoneAvailableResourceResponse
func (client *Client) DescribeMultiZoneAvailableResource(request *DescribeMultiZoneAvailableResourceRequest) (_result *DescribeMultiZoneAvailableResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMultiZoneAvailableResourceResponse{}
	_body, _err := client.DescribeMultiZoneAvailableResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeMultiZoneClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMultiZoneClusterResponse
func (client *Client) DescribeMultiZoneClusterWithOptions(request *DescribeMultiZoneClusterRequest, runtime *util.RuntimeOptions) (_result *DescribeMultiZoneClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMultiZoneCluster"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMultiZoneClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeMultiZoneClusterRequest
//
// @return DescribeMultiZoneClusterResponse
func (client *Client) DescribeMultiZoneCluster(request *DescribeMultiZoneClusterRequest) (_result *DescribeMultiZoneClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMultiZoneClusterResponse{}
	_body, _err := client.DescribeMultiZoneClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRecoverableTimeRangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecoverableTimeRangeResponse
func (client *Client) DescribeRecoverableTimeRangeWithOptions(request *DescribeRecoverableTimeRangeRequest, runtime *util.RuntimeOptions) (_result *DescribeRecoverableTimeRangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRecoverableTimeRange"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRecoverableTimeRangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRecoverableTimeRangeRequest
//
// @return DescribeRecoverableTimeRangeResponse
func (client *Client) DescribeRecoverableTimeRange(request *DescribeRecoverableTimeRangeRequest) (_result *DescribeRecoverableTimeRangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecoverableTimeRangeResponse{}
	_body, _err := client.DescribeRecoverableTimeRangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2019-01-01"),
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

// @param request - DescribeRestoreFullDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreFullDetailsResponse
func (client *Client) DescribeRestoreFullDetailsWithOptions(request *DescribeRestoreFullDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeRestoreFullDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreRecordId)) {
		query["RestoreRecordId"] = request.RestoreRecordId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRestoreFullDetails"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRestoreFullDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRestoreFullDetailsRequest
//
// @return DescribeRestoreFullDetailsResponse
func (client *Client) DescribeRestoreFullDetails(request *DescribeRestoreFullDetailsRequest) (_result *DescribeRestoreFullDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRestoreFullDetailsResponse{}
	_body, _err := client.DescribeRestoreFullDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRestoreIncrDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreIncrDetailResponse
func (client *Client) DescribeRestoreIncrDetailWithOptions(request *DescribeRestoreIncrDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeRestoreIncrDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreRecordId)) {
		query["RestoreRecordId"] = request.RestoreRecordId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRestoreIncrDetail"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRestoreIncrDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRestoreIncrDetailRequest
//
// @return DescribeRestoreIncrDetailResponse
func (client *Client) DescribeRestoreIncrDetail(request *DescribeRestoreIncrDetailRequest) (_result *DescribeRestoreIncrDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRestoreIncrDetailResponse{}
	_body, _err := client.DescribeRestoreIncrDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRestoreSchemaDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreSchemaDetailsResponse
func (client *Client) DescribeRestoreSchemaDetailsWithOptions(request *DescribeRestoreSchemaDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeRestoreSchemaDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreRecordId)) {
		query["RestoreRecordId"] = request.RestoreRecordId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRestoreSchemaDetails"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRestoreSchemaDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRestoreSchemaDetailsRequest
//
// @return DescribeRestoreSchemaDetailsResponse
func (client *Client) DescribeRestoreSchemaDetails(request *DescribeRestoreSchemaDetailsRequest) (_result *DescribeRestoreSchemaDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRestoreSchemaDetailsResponse{}
	_body, _err := client.DescribeRestoreSchemaDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRestoreSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreSummaryResponse
func (client *Client) DescribeRestoreSummaryWithOptions(request *DescribeRestoreSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeRestoreSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRestoreSummary"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRestoreSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRestoreSummaryRequest
//
// @return DescribeRestoreSummaryResponse
func (client *Client) DescribeRestoreSummary(request *DescribeRestoreSummaryRequest) (_result *DescribeRestoreSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRestoreSummaryResponse{}
	_body, _err := client.DescribeRestoreSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRestoreTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreTablesResponse
func (client *Client) DescribeRestoreTablesWithOptions(request *DescribeRestoreTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeRestoreTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RestoreRecordId)) {
		query["RestoreRecordId"] = request.RestoreRecordId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRestoreTables"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRestoreTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRestoreTablesRequest
//
// @return DescribeRestoreTablesResponse
func (client *Client) DescribeRestoreTables(request *DescribeRestoreTablesRequest) (_result *DescribeRestoreTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRestoreTablesResponse{}
	_body, _err := client.DescribeRestoreTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeSecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityGroupsResponse
func (client *Client) DescribeSecurityGroupsWithOptions(request *DescribeSecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSecurityGroups"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSecurityGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeSecurityGroupsRequest
//
// @return DescribeSecurityGroupsResponse
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

// @param request - DescribeServerlessClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServerlessClusterResponse
func (client *Client) DescribeServerlessClusterWithOptions(request *DescribeServerlessClusterRequest, runtime *util.RuntimeOptions) (_result *DescribeServerlessClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServerlessCluster"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServerlessClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeServerlessClusterRequest
//
// @return DescribeServerlessClusterResponse
func (client *Client) DescribeServerlessCluster(request *DescribeServerlessClusterRequest) (_result *DescribeServerlessClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServerlessClusterResponse{}
	_body, _err := client.DescribeServerlessClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeSubDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSubDomainResponse
func (client *Client) DescribeSubDomainWithOptions(request *DescribeSubDomainRequest, runtime *util.RuntimeOptions) (_result *DescribeSubDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSubDomain"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSubDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeSubDomainRequest
//
// @return DescribeSubDomainResponse
func (client *Client) DescribeSubDomain(request *DescribeSubDomainRequest) (_result *DescribeSubDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubDomainResponse{}
	_body, _err := client.DescribeSubDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EnableHBaseueBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableHBaseueBackupResponse
func (client *Client) EnableHBaseueBackupWithOptions(request *EnableHBaseueBackupRequest, runtime *util.RuntimeOptions) (_result *EnableHBaseueBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ColdStorageSize)) {
		query["ColdStorageSize"] = request.ColdStorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.HbaseueClusterId)) {
		query["HbaseueClusterId"] = request.HbaseueClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeCount)) {
		query["NodeCount"] = request.NodeCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableHBaseueBackup"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableHBaseueBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableHBaseueBackupRequest
//
// @return EnableHBaseueBackupResponse
func (client *Client) EnableHBaseueBackup(request *EnableHBaseueBackupRequest) (_result *EnableHBaseueBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableHBaseueBackupResponse{}
	_body, _err := client.EnableHBaseueBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EnableHBaseueModuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableHBaseueModuleResponse
func (client *Client) EnableHBaseueModuleWithOptions(request *EnableHBaseueModuleRequest, runtime *util.RuntimeOptions) (_result *EnableHBaseueModuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.BdsId)) {
		query["BdsId"] = request.BdsId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CoreInstanceType)) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.DiskSize)) {
		query["DiskSize"] = request.DiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.DiskType)) {
		query["DiskType"] = request.DiskType
	}

	if !tea.BoolValue(util.IsUnset(request.HbaseueClusterId)) {
		query["HbaseueClusterId"] = request.HbaseueClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MasterInstanceType)) {
		query["MasterInstanceType"] = request.MasterInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleClusterName)) {
		query["ModuleClusterName"] = request.ModuleClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleTypeName)) {
		query["ModuleTypeName"] = request.ModuleTypeName
	}

	if !tea.BoolValue(util.IsUnset(request.NodeCount)) {
		query["NodeCount"] = request.NodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VswitchId)) {
		query["VswitchId"] = request.VswitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableHBaseueModule"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableHBaseueModuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableHBaseueModuleRequest
//
// @return EnableHBaseueModuleResponse
func (client *Client) EnableHBaseueModule(request *EnableHBaseueModuleRequest) (_result *EnableHBaseueModuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableHBaseueModuleResponse{}
	_body, _err := client.EnableHBaseueModuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EvaluateMultiZoneResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvaluateMultiZoneResourceResponse
func (client *Client) EvaluateMultiZoneResourceWithOptions(request *EvaluateMultiZoneResourceRequest, runtime *util.RuntimeOptions) (_result *EvaluateMultiZoneResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArbiterVSwitchId)) {
		query["ArbiterVSwitchId"] = request.ArbiterVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ArbiterZoneId)) {
		query["ArbiterZoneId"] = request.ArbiterZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.ArchVersion)) {
		query["ArchVersion"] = request.ArchVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.CoreDiskSize)) {
		query["CoreDiskSize"] = request.CoreDiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.CoreDiskType)) {
		query["CoreDiskType"] = request.CoreDiskType
	}

	if !tea.BoolValue(util.IsUnset(request.CoreInstanceType)) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.CoreNodeCount)) {
		query["CoreNodeCount"] = request.CoreNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.EngineVersion)) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !tea.BoolValue(util.IsUnset(request.LogDiskSize)) {
		query["LogDiskSize"] = request.LogDiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.LogDiskType)) {
		query["LogDiskType"] = request.LogDiskType
	}

	if !tea.BoolValue(util.IsUnset(request.LogInstanceType)) {
		query["LogInstanceType"] = request.LogInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.LogNodeCount)) {
		query["LogNodeCount"] = request.LogNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.MasterInstanceType)) {
		query["MasterInstanceType"] = request.MasterInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.MultiZoneCombination)) {
		query["MultiZoneCombination"] = request.MultiZoneCombination
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryVSwitchId)) {
		query["PrimaryVSwitchId"] = request.PrimaryVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryZoneId)) {
		query["PrimaryZoneId"] = request.PrimaryZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIPList)) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !tea.BoolValue(util.IsUnset(request.StandbyVSwitchId)) {
		query["StandbyVSwitchId"] = request.StandbyVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.StandbyZoneId)) {
		query["StandbyZoneId"] = request.StandbyZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EvaluateMultiZoneResource"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EvaluateMultiZoneResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - EvaluateMultiZoneResourceRequest
//
// @return EvaluateMultiZoneResourceResponse
func (client *Client) EvaluateMultiZoneResource(request *EvaluateMultiZoneResourceRequest) (_result *EvaluateMultiZoneResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EvaluateMultiZoneResourceResponse{}
	_body, _err := client.EvaluateMultiZoneResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetMultimodeCmsUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultimodeCmsUrlResponse
func (client *Client) GetMultimodeCmsUrlWithOptions(request *GetMultimodeCmsUrlRequest, runtime *util.RuntimeOptions) (_result *GetMultimodeCmsUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMultimodeCmsUrl"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMultimodeCmsUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMultimodeCmsUrlRequest
//
// @return GetMultimodeCmsUrlResponse
func (client *Client) GetMultimodeCmsUrl(request *GetMultimodeCmsUrlRequest) (_result *GetMultimodeCmsUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMultimodeCmsUrlResponse{}
	_body, _err := client.GetMultimodeCmsUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 授权账户权限
//
// @param request - GrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantResponse
func (client *Client) GrantWithOptions(request *GrantRequest, runtime *util.RuntimeOptions) (_result *GrantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AclActions)) {
		query["AclActions"] = request.AclActions
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Grant"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 授权账户权限
//
// @param request - GrantRequest
//
// @return GrantResponse
func (client *Client) Grant(request *GrantRequest) (_result *GrantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantResponse{}
	_body, _err := client.GrantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListHBaseInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHBaseInstancesResponse
func (client *Client) ListHBaseInstancesWithOptions(request *ListHBaseInstancesRequest, runtime *util.RuntimeOptions) (_result *ListHBaseInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHBaseInstances"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHBaseInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListHBaseInstancesRequest
//
// @return ListHBaseInstancesResponse
func (client *Client) ListHBaseInstances(request *ListHBaseInstancesRequest) (_result *ListHBaseInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHBaseInstancesResponse{}
	_body, _err := client.ListHBaseInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListInstanceServiceConfigHistoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceServiceConfigHistoriesResponse
func (client *Client) ListInstanceServiceConfigHistoriesWithOptions(request *ListInstanceServiceConfigHistoriesRequest, runtime *util.RuntimeOptions) (_result *ListInstanceServiceConfigHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceServiceConfigHistories"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceServiceConfigHistoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListInstanceServiceConfigHistoriesRequest
//
// @return ListInstanceServiceConfigHistoriesResponse
func (client *Client) ListInstanceServiceConfigHistories(request *ListInstanceServiceConfigHistoriesRequest) (_result *ListInstanceServiceConfigHistoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstanceServiceConfigHistoriesResponse{}
	_body, _err := client.ListInstanceServiceConfigHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListInstanceServiceConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceServiceConfigurationsResponse
func (client *Client) ListInstanceServiceConfigurationsWithOptions(request *ListInstanceServiceConfigurationsRequest, runtime *util.RuntimeOptions) (_result *ListInstanceServiceConfigurationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceServiceConfigurations"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceServiceConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListInstanceServiceConfigurationsRequest
//
// @return ListInstanceServiceConfigurationsResponse
func (client *Client) ListInstanceServiceConfigurations(request *ListInstanceServiceConfigurationsRequest) (_result *ListInstanceServiceConfigurationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstanceServiceConfigurationsResponse{}
	_body, _err := client.ListInstanceServiceConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2019-01-01"),
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

// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
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

// @param request - ListTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsResponse
func (client *Client) ListTagsWithOptions(request *ListTagsRequest, runtime *util.RuntimeOptions) (_result *ListTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTags"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListTagsRequest
//
// @return ListTagsResponse
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

// Summary:
//
// 更改账户密码
//
// @param request - ModifyAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountPasswordResponse
func (client *Client) ModifyAccountPasswordWithOptions(request *ModifyAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NewAccountPassword)) {
		query["NewAccountPassword"] = request.NewAccountPassword
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccountPassword"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更改账户密码
//
// @param request - ModifyAccountPasswordRequest
//
// @return ModifyAccountPasswordResponse
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

// @param request - ModifyActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyActiveOperationTasksResponse
func (client *Client) ModifyActiveOperationTasksWithOptions(request *ModifyActiveOperationTasksRequest, runtime *util.RuntimeOptions) (_result *ModifyActiveOperationTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.ImmediateStart)) {
		query["ImmediateStart"] = request.ImmediateStart
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

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SwitchTime)) {
		query["SwitchTime"] = request.SwitchTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyActiveOperationTasks"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyActiveOperationTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyActiveOperationTasksRequest
//
// @return ModifyActiveOperationTasksResponse
func (client *Client) ModifyActiveOperationTasks(request *ModifyActiveOperationTasksRequest) (_result *ModifyActiveOperationTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyActiveOperationTasksResponse{}
	_body, _err := client.ModifyActiveOperationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyBackupPlanConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupPlanConfigResponse
func (client *Client) ModifyBackupPlanConfigWithOptions(request *ModifyBackupPlanConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupPlanConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.FullBackupCycle)) {
		query["FullBackupCycle"] = request.FullBackupCycle
	}

	if !tea.BoolValue(util.IsUnset(request.MinHFileBackupCount)) {
		query["MinHFileBackupCount"] = request.MinHFileBackupCount
	}

	if !tea.BoolValue(util.IsUnset(request.NextFullBackupDate)) {
		query["NextFullBackupDate"] = request.NextFullBackupDate
	}

	if !tea.BoolValue(util.IsUnset(request.Tables)) {
		query["Tables"] = request.Tables
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyBackupPlanConfig"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyBackupPlanConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyBackupPlanConfigRequest
//
// @return ModifyBackupPlanConfigResponse
func (client *Client) ModifyBackupPlanConfig(request *ModifyBackupPlanConfigRequest) (_result *ModifyBackupPlanConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupPlanConfigResponse{}
	_body, _err := client.ModifyBackupPlanConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupEndTimeUTC)) {
		query["PreferredBackupEndTimeUTC"] = request.PreferredBackupEndTimeUTC
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupPeriod)) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupStartTimeUTC)) {
		query["PreferredBackupStartTimeUTC"] = request.PreferredBackupStartTimeUTC
	}

	if !tea.BoolValue(util.IsUnset(request.PreferredBackupTime)) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyBackupPolicy"),
		Version:     tea.String("2019-01-01"),
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

// @param request - ModifyClusterDeletionProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterDeletionProtectionResponse
func (client *Client) ModifyClusterDeletionProtectionWithOptions(request *ModifyClusterDeletionProtectionRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterDeletionProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Protection)) {
		query["Protection"] = request.Protection
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterDeletionProtection"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterDeletionProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyClusterDeletionProtectionRequest
//
// @return ModifyClusterDeletionProtectionResponse
func (client *Client) ModifyClusterDeletionProtection(request *ModifyClusterDeletionProtectionRequest) (_result *ModifyClusterDeletionProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterDeletionProtectionResponse{}
	_body, _err := client.ModifyClusterDeletionProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyDiskWarningLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDiskWarningLineResponse
func (client *Client) ModifyDiskWarningLineWithOptions(request *ModifyDiskWarningLineRequest, runtime *util.RuntimeOptions) (_result *ModifyDiskWarningLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.WarningLine)) {
		query["WarningLine"] = request.WarningLine
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDiskWarningLine"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDiskWarningLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyDiskWarningLineRequest
//
// @return ModifyDiskWarningLineResponse
func (client *Client) ModifyDiskWarningLine(request *ModifyDiskWarningLineRequest) (_result *ModifyDiskWarningLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDiskWarningLineResponse{}
	_body, _err := client.ModifyDiskWarningLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyInstanceMaintainTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceMaintainTimeResponse
func (client *Client) ModifyInstanceMaintainTimeWithOptions(request *ModifyInstanceMaintainTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MaintainEndTime)) {
		query["MaintainEndTime"] = request.MaintainEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaintainStartTime)) {
		query["MaintainStartTime"] = request.MaintainStartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceMaintainTime"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceMaintainTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyInstanceMaintainTimeRequest
//
// @return ModifyInstanceMaintainTimeResponse
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

// @param request - ModifyInstanceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceNameResponse
func (client *Client) ModifyInstanceNameWithOptions(request *ModifyInstanceNameRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceName"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyInstanceNameRequest
//
// @return ModifyInstanceNameResponse
func (client *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (_result *ModifyInstanceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceNameResponse{}
	_body, _err := client.ModifyInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyInstanceServiceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceServiceConfigResponse
func (client *Client) ModifyInstanceServiceConfigWithOptions(request *ModifyInstanceServiceConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceServiceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigureName)) {
		query["ConfigureName"] = request.ConfigureName
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigureValue)) {
		query["ConfigureValue"] = request.ConfigureValue
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.Restart)) {
		query["Restart"] = request.Restart
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceServiceConfig"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceServiceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyInstanceServiceConfigRequest
//
// @return ModifyInstanceServiceConfigResponse
func (client *Client) ModifyInstanceServiceConfig(request *ModifyInstanceServiceConfigRequest) (_result *ModifyInstanceServiceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceServiceConfigResponse{}
	_body, _err := client.ModifyInstanceServiceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyInstanceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceTypeResponse
func (client *Client) ModifyInstanceTypeWithOptions(request *ModifyInstanceTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.CoreInstanceType)) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.MasterInstanceType)) {
		query["MasterInstanceType"] = request.MasterInstanceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceType"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyInstanceTypeRequest
//
// @return ModifyInstanceTypeResponse
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

// @param request - ModifyIpWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIpWhitelistResponse
func (client *Client) ModifyIpWhitelistWithOptions(request *ModifyIpWhitelistRequest, runtime *util.RuntimeOptions) (_result *ModifyIpWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyIpWhitelist"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyIpWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyIpWhitelistRequest
//
// @return ModifyIpWhitelistResponse
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

// @param request - ModifyMultiZoneClusterNodeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMultiZoneClusterNodeTypeResponse
func (client *Client) ModifyMultiZoneClusterNodeTypeWithOptions(request *ModifyMultiZoneClusterNodeTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyMultiZoneClusterNodeTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.CoreInstanceType)) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.LogInstanceType)) {
		query["LogInstanceType"] = request.LogInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.MasterInstanceType)) {
		query["MasterInstanceType"] = request.MasterInstanceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMultiZoneClusterNodeType"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMultiZoneClusterNodeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyMultiZoneClusterNodeTypeRequest
//
// @return ModifyMultiZoneClusterNodeTypeResponse
func (client *Client) ModifyMultiZoneClusterNodeType(request *ModifyMultiZoneClusterNodeTypeRequest) (_result *ModifyMultiZoneClusterNodeTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMultiZoneClusterNodeTypeResponse{}
	_body, _err := client.ModifyMultiZoneClusterNodeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifySecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecurityGroupsResponse
func (client *Client) ModifySecurityGroupsWithOptions(request *ModifySecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *ModifySecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupIds)) {
		query["SecurityGroupIds"] = request.SecurityGroupIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySecurityGroups"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySecurityGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifySecurityGroupsRequest
//
// @return ModifySecurityGroupsResponse
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

// @param request - ModifyUIAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUIAccountPasswordResponse
func (client *Client) ModifyUIAccountPasswordWithOptions(request *ModifyUIAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyUIAccountPasswordResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUIAccountPassword"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUIAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyUIAccountPasswordRequest
//
// @return ModifyUIAccountPasswordResponse
func (client *Client) ModifyUIAccountPassword(request *ModifyUIAccountPasswordRequest) (_result *ModifyUIAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUIAccountPasswordResponse{}
	_body, _err := client.ModifyUIAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResourceGroup"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
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

// @param request - OpenBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenBackupResponse
func (client *Client) OpenBackupWithOptions(request *OpenBackupRequest, runtime *util.RuntimeOptions) (_result *OpenBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenBackup"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - OpenBackupRequest
//
// @return OpenBackupResponse
func (client *Client) OpenBackup(request *OpenBackupRequest) (_result *OpenBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenBackupResponse{}
	_body, _err := client.OpenBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PurgeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PurgeInstanceResponse
func (client *Client) PurgeInstanceWithOptions(request *PurgeInstanceRequest, runtime *util.RuntimeOptions) (_result *PurgeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PurgeInstance"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PurgeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - PurgeInstanceRequest
//
// @return PurgeInstanceResponse
func (client *Client) PurgeInstance(request *PurgeInstanceRequest) (_result *PurgeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PurgeInstanceResponse{}
	_body, _err := client.PurgeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryHBaseHaDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryHBaseHaDBResponse
func (client *Client) QueryHBaseHaDBWithOptions(request *QueryHBaseHaDBRequest, runtime *util.RuntimeOptions) (_result *QueryHBaseHaDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BdsId)) {
		query["BdsId"] = request.BdsId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryHBaseHaDB"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHBaseHaDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryHBaseHaDBRequest
//
// @return QueryHBaseHaDBResponse
func (client *Client) QueryHBaseHaDB(request *QueryHBaseHaDBRequest) (_result *QueryHBaseHaDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryHBaseHaDBResponse{}
	_body, _err := client.QueryHBaseHaDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryXpackRelateDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryXpackRelateDBResponse
func (client *Client) QueryXpackRelateDBWithOptions(request *QueryXpackRelateDBRequest, runtime *util.RuntimeOptions) (_result *QueryXpackRelateDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.HasSingleNode)) {
		query["HasSingleNode"] = request.HasSingleNode
	}

	if !tea.BoolValue(util.IsUnset(request.RelateDbType)) {
		query["RelateDbType"] = request.RelateDbType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryXpackRelateDB"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryXpackRelateDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryXpackRelateDBRequest
//
// @return QueryXpackRelateDBResponse
func (client *Client) QueryXpackRelateDB(request *QueryXpackRelateDBRequest) (_result *QueryXpackRelateDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryXpackRelateDBResponse{}
	_body, _err := client.QueryXpackRelateDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RelateDbForHBaseHaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RelateDbForHBaseHaResponse
func (client *Client) RelateDbForHBaseHaWithOptions(request *RelateDbForHBaseHaRequest, runtime *util.RuntimeOptions) (_result *RelateDbForHBaseHaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.HaActive)) {
		query["HaActive"] = request.HaActive
	}

	if !tea.BoolValue(util.IsUnset(request.HaActiveClusterKey)) {
		query["HaActiveClusterKey"] = request.HaActiveClusterKey
	}

	if !tea.BoolValue(util.IsUnset(request.HaActiveDBType)) {
		query["HaActiveDBType"] = request.HaActiveDBType
	}

	if !tea.BoolValue(util.IsUnset(request.HaActiveHbaseFsDir)) {
		query["HaActiveHbaseFsDir"] = request.HaActiveHbaseFsDir
	}

	if !tea.BoolValue(util.IsUnset(request.HaActiveHdfsUri)) {
		query["HaActiveHdfsUri"] = request.HaActiveHdfsUri
	}

	if !tea.BoolValue(util.IsUnset(request.HaActivePassword)) {
		query["HaActivePassword"] = request.HaActivePassword
	}

	if !tea.BoolValue(util.IsUnset(request.HaActiveUser)) {
		query["HaActiveUser"] = request.HaActiveUser
	}

	if !tea.BoolValue(util.IsUnset(request.HaActiveVersion)) {
		query["HaActiveVersion"] = request.HaActiveVersion
	}

	if !tea.BoolValue(util.IsUnset(request.HaMigrateType)) {
		query["HaMigrateType"] = request.HaMigrateType
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandby)) {
		query["HaStandby"] = request.HaStandby
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandbyClusterKey)) {
		query["HaStandbyClusterKey"] = request.HaStandbyClusterKey
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandbyDBType)) {
		query["HaStandbyDBType"] = request.HaStandbyDBType
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandbyHbaseFsDir)) {
		query["HaStandbyHbaseFsDir"] = request.HaStandbyHbaseFsDir
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandbyHdfsUri)) {
		query["HaStandbyHdfsUri"] = request.HaStandbyHdfsUri
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandbyPassword)) {
		query["HaStandbyPassword"] = request.HaStandbyPassword
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandbyUser)) {
		query["HaStandbyUser"] = request.HaStandbyUser
	}

	if !tea.BoolValue(util.IsUnset(request.HaStandbyVersion)) {
		query["HaStandbyVersion"] = request.HaStandbyVersion
	}

	if !tea.BoolValue(util.IsUnset(request.HaTables)) {
		query["HaTables"] = request.HaTables
	}

	if !tea.BoolValue(util.IsUnset(request.IsActiveStandard)) {
		query["IsActiveStandard"] = request.IsActiveStandard
	}

	if !tea.BoolValue(util.IsUnset(request.IsStandbyStandard)) {
		query["IsStandbyStandard"] = request.IsStandbyStandard
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RelateDbForHBaseHa"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RelateDbForHBaseHaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - RelateDbForHBaseHaRequest
//
// @return RelateDbForHBaseHaResponse
func (client *Client) RelateDbForHBaseHa(request *RelateDbForHBaseHaRequest) (_result *RelateDbForHBaseHaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RelateDbForHBaseHaResponse{}
	_body, _err := client.RelateDbForHBaseHaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ReleasePublicNetworkAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleasePublicNetworkAddressResponse
func (client *Client) ReleasePublicNetworkAddressWithOptions(request *ReleasePublicNetworkAddressRequest, runtime *util.RuntimeOptions) (_result *ReleasePublicNetworkAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleasePublicNetworkAddress"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleasePublicNetworkAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ReleasePublicNetworkAddressRequest
//
// @return ReleasePublicNetworkAddressResponse
func (client *Client) ReleasePublicNetworkAddress(request *ReleasePublicNetworkAddressRequest) (_result *ReleasePublicNetworkAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleasePublicNetworkAddressResponse{}
	_body, _err := client.ReleasePublicNetworkAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RenewInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewInstance"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - RenewInstanceRequest
//
// @return RenewInstanceResponse
func (client *Client) RenewInstance(request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResizeColdStorageSizeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResizeColdStorageSizeResponse
func (client *Client) ResizeColdStorageSizeWithOptions(request *ResizeColdStorageSizeRequest, runtime *util.RuntimeOptions) (_result *ResizeColdStorageSizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ColdStorageSize)) {
		query["ColdStorageSize"] = request.ColdStorageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResizeColdStorageSize"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResizeColdStorageSizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResizeColdStorageSizeRequest
//
// @return ResizeColdStorageSizeResponse
func (client *Client) ResizeColdStorageSize(request *ResizeColdStorageSizeRequest) (_result *ResizeColdStorageSizeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeColdStorageSizeResponse{}
	_body, _err := client.ResizeColdStorageSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResizeDiskSizeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResizeDiskSizeResponse
func (client *Client) ResizeDiskSizeWithOptions(request *ResizeDiskSizeRequest, runtime *util.RuntimeOptions) (_result *ResizeDiskSizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeDiskSize)) {
		query["NodeDiskSize"] = request.NodeDiskSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResizeDiskSize"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResizeDiskSizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResizeDiskSizeRequest
//
// @return ResizeDiskSizeResponse
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

// @param request - ResizeMultiZoneClusterDiskSizeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResizeMultiZoneClusterDiskSizeResponse
func (client *Client) ResizeMultiZoneClusterDiskSizeWithOptions(request *ResizeMultiZoneClusterDiskSizeRequest, runtime *util.RuntimeOptions) (_result *ResizeMultiZoneClusterDiskSizeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.CoreDiskSize)) {
		query["CoreDiskSize"] = request.CoreDiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.LogDiskSize)) {
		query["LogDiskSize"] = request.LogDiskSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResizeMultiZoneClusterDiskSize"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResizeMultiZoneClusterDiskSizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResizeMultiZoneClusterDiskSizeRequest
//
// @return ResizeMultiZoneClusterDiskSizeResponse
func (client *Client) ResizeMultiZoneClusterDiskSize(request *ResizeMultiZoneClusterDiskSizeRequest) (_result *ResizeMultiZoneClusterDiskSizeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeMultiZoneClusterDiskSizeResponse{}
	_body, _err := client.ResizeMultiZoneClusterDiskSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResizeMultiZoneClusterNodeCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResizeMultiZoneClusterNodeCountResponse
func (client *Client) ResizeMultiZoneClusterNodeCountWithOptions(request *ResizeMultiZoneClusterNodeCountRequest, runtime *util.RuntimeOptions) (_result *ResizeMultiZoneClusterNodeCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArbiterVSwitchId)) {
		query["ArbiterVSwitchId"] = request.ArbiterVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.CoreNodeCount)) {
		query["CoreNodeCount"] = request.CoreNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.LogNodeCount)) {
		query["LogNodeCount"] = request.LogNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryCoreNodeCount)) {
		query["PrimaryCoreNodeCount"] = request.PrimaryCoreNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryVSwitchId)) {
		query["PrimaryVSwitchId"] = request.PrimaryVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.StandbyCoreNodeCount)) {
		query["StandbyCoreNodeCount"] = request.StandbyCoreNodeCount
	}

	if !tea.BoolValue(util.IsUnset(request.StandbyVSwitchId)) {
		query["StandbyVSwitchId"] = request.StandbyVSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResizeMultiZoneClusterNodeCount"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResizeMultiZoneClusterNodeCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResizeMultiZoneClusterNodeCountRequest
//
// @return ResizeMultiZoneClusterNodeCountResponse
func (client *Client) ResizeMultiZoneClusterNodeCount(request *ResizeMultiZoneClusterNodeCountRequest) (_result *ResizeMultiZoneClusterNodeCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeMultiZoneClusterNodeCountResponse{}
	_body, _err := client.ResizeMultiZoneClusterNodeCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResizeNodeCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResizeNodeCountResponse
func (client *Client) ResizeNodeCountWithOptions(request *ResizeNodeCountRequest, runtime *util.RuntimeOptions) (_result *ResizeNodeCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeCount)) {
		query["NodeCount"] = request.NodeCount
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
		Action:      tea.String("ResizeNodeCount"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResizeNodeCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResizeNodeCountRequest
//
// @return ResizeNodeCountResponse
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
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Components)) {
		query["Components"] = request.Components
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartInstance"),
		Version:     tea.String("2019-01-01"),
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
// 回收账户权限
//
// @param request - RevokeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeResponse
func (client *Client) RevokeWithOptions(request *RevokeRequest, runtime *util.RuntimeOptions) (_result *RevokeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AclActions)) {
		query["AclActions"] = request.AclActions
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Revoke"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 回收账户权限
//
// @param request - RevokeRequest
//
// @return RevokeResponse
func (client *Client) Revoke(request *RevokeRequest) (_result *RevokeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeResponse{}
	_body, _err := client.RevokeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SwitchHbaseHaSlbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchHbaseHaSlbResponse
func (client *Client) SwitchHbaseHaSlbWithOptions(request *SwitchHbaseHaSlbRequest, runtime *util.RuntimeOptions) (_result *SwitchHbaseHaSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BdsId)) {
		query["BdsId"] = request.BdsId
	}

	if !tea.BoolValue(util.IsUnset(request.HaId)) {
		query["HaId"] = request.HaId
	}

	if !tea.BoolValue(util.IsUnset(request.HaTypes)) {
		query["HaTypes"] = request.HaTypes
	}

	if !tea.BoolValue(util.IsUnset(request.HbaseType)) {
		query["HbaseType"] = request.HbaseType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchHbaseHaSlb"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SwitchHbaseHaSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SwitchHbaseHaSlbRequest
//
// @return SwitchHbaseHaSlbResponse
func (client *Client) SwitchHbaseHaSlb(request *SwitchHbaseHaSlbRequest) (_result *SwitchHbaseHaSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchHbaseHaSlbResponse{}
	_body, _err := client.SwitchHbaseHaSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开通/关闭 扩展服务
//
// @param request - SwitchServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchServiceResponse
func (client *Client) SwitchServiceWithOptions(request *SwitchServiceRequest, runtime *util.RuntimeOptions) (_result *SwitchServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Operate)) {
		query["Operate"] = request.Operate
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchService"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SwitchServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开通/关闭 扩展服务
//
// @param request - SwitchServiceRequest
//
// @return SwitchServiceResponse
func (client *Client) SwitchService(request *SwitchServiceRequest) (_result *SwitchServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchServiceResponse{}
	_body, _err := client.SwitchServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2019-01-01"),
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

// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
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

// @param request - UnTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, runtime *util.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnTagResources"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UnTagResourcesRequest
//
// @return UnTagResourcesResponse
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
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Components)) {
		query["Components"] = request.Components
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeMinorVersion"),
		Version:     tea.String("2019-01-01"),
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

// @param request - UpgradeMultiZoneClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeMultiZoneClusterResponse
func (client *Client) UpgradeMultiZoneClusterWithOptions(request *UpgradeMultiZoneClusterRequest, runtime *util.RuntimeOptions) (_result *UpgradeMultiZoneClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Components)) {
		query["Components"] = request.Components
	}

	if !tea.BoolValue(util.IsUnset(request.RestartComponents)) {
		query["RestartComponents"] = request.RestartComponents
	}

	if !tea.BoolValue(util.IsUnset(request.RunMode)) {
		query["RunMode"] = request.RunMode
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeInsName)) {
		query["UpgradeInsName"] = request.UpgradeInsName
	}

	if !tea.BoolValue(util.IsUnset(request.Versions)) {
		query["Versions"] = request.Versions
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeMultiZoneCluster"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeMultiZoneClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpgradeMultiZoneClusterRequest
//
// @return UpgradeMultiZoneClusterResponse
func (client *Client) UpgradeMultiZoneCluster(request *UpgradeMultiZoneClusterRequest) (_result *UpgradeMultiZoneClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeMultiZoneClusterResponse{}
	_body, _err := client.UpgradeMultiZoneClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - XpackRelateDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return XpackRelateDBResponse
func (client *Client) XpackRelateDBWithOptions(request *XpackRelateDBRequest, runtime *util.RuntimeOptions) (_result *XpackRelateDBResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DbClusterIds)) {
		query["DbClusterIds"] = request.DbClusterIds
	}

	if !tea.BoolValue(util.IsUnset(request.RelateDbType)) {
		query["RelateDbType"] = request.RelateDbType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("XpackRelateDB"),
		Version:     tea.String("2019-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &XpackRelateDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - XpackRelateDBRequest
//
// @return XpackRelateDBResponse
func (client *Client) XpackRelateDB(request *XpackRelateDBRequest) (_result *XpackRelateDBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &XpackRelateDBResponse{}
	_body, _err := client.XpackRelateDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
